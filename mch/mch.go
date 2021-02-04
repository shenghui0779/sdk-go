package mch

import (
	"context"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/shenghui0779/gochat/wx"
	"golang.org/x/crypto/pkcs12"
)

// Mch 微信支付
type Mch struct {
	appid     string
	mchid     string
	apikey    string
	nonce     func(size int) string
	client    wx.HTTPClient
	tlsClient wx.HTTPClient
}

// New returns new wechat pay
func New(appid, mchid, apikey string) *Mch {
	c := wx.NewHTTPClient(&tls.Config{InsecureSkipVerify: true})

	return &Mch{
		appid:  appid,
		mchid:  mchid,
		apikey: apikey,
		nonce: func(size int) string {
			nonce := make([]byte, size/2)
			io.ReadFull(rand.Reader, nonce)

			return hex.EncodeToString(nonce)
		},
		client:    c,
		tlsClient: c,
	}
}

// LoadCertFromP12File load cert from p12(pfx) file
func (mch *Mch) LoadCertFromP12File(path string) error {
	p12, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	cert, err := mch.pkcs12ToPem(p12)

	if err != nil {
		return err
	}

	mch.tlsClient = wx.NewHTTPClient(&tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	})

	return nil
}

// LoadCertFromPemFile load cert from PEM file
func (mch *Mch) LoadCertFromPemFile(certFile, keyFile string) error {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)

	if err != nil {
		return err
	}

	mch.tlsClient = wx.NewHTTPClient(&tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	})

	return nil
}

// LoadCertFromPemBlock load cert from a pair of PEM encoded data
func (mch *Mch) LoadCertFromPemBlock(certPEMBlock, keyPEMBlock []byte) error {
	cert, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)

	if err != nil {
		return err
	}

	mch.tlsClient = wx.NewHTTPClient(&tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	})

	return nil
}

// Do exec action
func (mch *Mch) Do(ctx context.Context, action wx.Action, options ...wx.HTTPOption) (wx.WXML, error) {
	m, err := action.WXML(mch.appid, mch.mchid, mch.nonce(16))

	if err != nil {
		return nil, err
	}

	// 签名
	if v, ok := m["sign_type"]; ok && v == SignHMacSHA256 {
		m["sign"] = mch.SignWithHMacSHA256(m, true)
	} else {
		m["sign"] = mch.SignWithMD5(m, true)
	}

	reqURL := action.URL()

	switch reqURL {
	case ContractOAEntrust: // 公众号签约
		query := url.Values{}

		for k, v := range m {
			query.Add(k, v)
		}

		return wx.WXML{"entrust_url": fmt.Sprintf("%s?%s", PappayOAEntrustURL, query.Encode())}, nil
	case ContractMPEntrust: // 小程序签约
		return m, nil
	case ContractH5Entrust: // H5签约
		query := url.Values{}

		for k, v := range m {
			query.Add(k, v)
		}

		return wx.WXML{"entrust_url": fmt.Sprintf("%s?%s", PappayH5EntrustURL, query.Encode())}, nil
	}

	var resp []byte

	if action.TLS() {
		resp, err = mch.tlsClient.PostXML(ctx, reqURL, m, options...)
	} else {
		resp, err = mch.client.PostXML(ctx, reqURL, m, options...)
	}

	if err != nil {
		return nil, err
	}

	// XML解析
	result, err := wx.ParseXML2Map(resp)

	if err != nil {
		return nil, err
	}

	if result["return_code"] != ResultSuccess {
		return nil, errors.New(result["return_msg"])
	}

	// 签名验证
	if err := mch.VerifyWXMLResult(result); err != nil {
		return nil, err
	}

	return result, nil
}

// APPAPI 用于APP拉起支付
func (mch *Mch) APPAPI(prepayID string) wx.WXML {
	m := wx.WXML{
		"appid":     mch.appid,
		"partnerid": mch.mchid,
		"prepayid":  prepayID,
		"package":   "Sign=WXPay",
		"noncestr":  mch.nonce(16),
		"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
	}

	m["sign"] = mch.SignWithMD5(m, true)

	return m
}

// JSAPI 用于JS拉起支付
func (mch *Mch) JSAPI(prepayID string) wx.WXML {
	m := wx.WXML{
		"appId":     mch.appid,
		"nonceStr":  mch.nonce(16),
		"package":   fmt.Sprintf("prepay_id=%s", prepayID),
		"signType":  SignMD5,
		"timeStamp": strconv.FormatInt(time.Now().Unix(), 10),
	}

	m["paySign"] = mch.SignWithMD5(m, true)

	return m
}

// MinipRedpackJSAPI 小程序领取红包
func (mch *Mch) MinipRedpackJSAPI(pkg string) wx.WXML {
	m := wx.WXML{
		"appId":     mch.appid,
		"nonceStr":  mch.nonce(16),
		"package":   url.QueryEscape(pkg),
		"timeStamp": strconv.FormatInt(time.Now().Unix(), 10),
	}

	m["paySign"] = mch.SignWithMD5(m, false)

	delete(m, "appId")
	m["signType"] = SignMD5

	return m
}

// DownloadBill 下载交易账单
// 账单日期格式：20140603
func (mch *Mch) DownloadBill(ctx context.Context, billDate, billType string) ([]byte, error) {
	m := wx.WXML{
		"appid":     mch.appid,
		"mch_id":    mch.mchid,
		"bill_date": billDate,
		"bill_type": billType,
		"nonce_str": mch.nonce(16),
	}

	m["sign"] = mch.SignWithMD5(m, true)

	resp, err := mch.client.PostXML(ctx, DownloadBillURL, m, wx.WithHTTPClose())

	if err != nil {
		return nil, err
	}

	// XML解析
	result, err := wx.ParseXML2Map(resp)

	if err != nil {
		return nil, err
	}

	if len(result) != 0 && result["return_code"] != ResultSuccess {
		return nil, errors.New(result["return_msg"])
	}

	return resp, nil
}

// DownloadFundFlow 下载资金账单
// 账单日期格式：20140603
func (mch *Mch) DownloadFundFlow(ctx context.Context, billDate, accountType string) ([]byte, error) {
	m := wx.WXML{
		"appid":        mch.appid,
		"mch_id":       mch.mchid,
		"bill_date":    billDate,
		"account_type": accountType,
		"nonce_str":    mch.nonce(16),
	}

	m["sign"] = mch.SignWithHMacSHA256(m, true)

	resp, err := mch.tlsClient.PostXML(ctx, DownloadFundFlowURL, m, wx.WithHTTPClose())

	if err != nil {
		return nil, err
	}

	// XML解析
	result, err := wx.ParseXML2Map(resp)

	if err != nil {
		return nil, err
	}

	if len(result) != 0 && result["return_code"] != ResultSuccess {
		return nil, errors.New(result["return_msg"])
	}

	return resp, nil
}

// BatchQueryComment 拉取订单评价数据
// 时间格式：yyyyMMddHHmmss
// 默认一次且最多拉取200条
func (mch *Mch) BatchQueryComment(ctx context.Context, beginTime, endTime string, offset int, limit ...int) ([]byte, error) {
	m := wx.WXML{
		"appid":      mch.appid,
		"mch_id":     mch.mchid,
		"begin_time": beginTime,
		"end_time":   endTime,
		"offset":     strconv.Itoa(offset),
		"nonce_str":  mch.nonce(16),
	}

	if len(limit) != 0 {
		m["limit"] = strconv.Itoa(limit[0])
	}

	m["sign"] = mch.SignWithHMacSHA256(m, true)

	resp, err := mch.tlsClient.PostXML(ctx, BatchQueryCommentURL, m, wx.WithHTTPClose())

	if err != nil {
		return nil, err
	}

	// XML解析
	result, err := wx.ParseXML2Map(resp)

	if err != nil {
		return nil, err
	}

	if len(result) != 0 && result["return_code"] != ResultSuccess {
		return nil, errors.New(result["return_msg"])
	}

	return resp, nil
}

// SignWithMD5 生成MD5签名
func (mch *Mch) SignWithMD5(m wx.WXML, toUpper bool) string {
	h := md5.New()
	h.Write([]byte(mch.buildSignStr(m)))

	sign := hex.EncodeToString(h.Sum(nil))

	if toUpper {
		sign = strings.ToUpper(sign)
	}

	return sign
}

// SignWithHMacSHA256 生成HMAC-SHA256签名
func (mch *Mch) SignWithHMacSHA256(m wx.WXML, toUpper bool) string {
	h := hmac.New(sha256.New, []byte(mch.apikey))
	h.Write([]byte(mch.buildSignStr(m)))

	sign := hex.EncodeToString(h.Sum(nil))

	if toUpper {
		sign = strings.ToUpper(sign)
	}

	return sign
}

// VerifyWXMLResult 微信请求/回调通知签名验证
func (mch *Mch) VerifyWXMLResult(m wx.WXML) error {
	if wxsign, ok := m["sign"]; ok {
		signature := ""

		if v, ok := m["sign_type"]; ok && v == SignHMacSHA256 {
			signature = mch.SignWithHMacSHA256(m, true)
		} else {
			signature = mch.SignWithMD5(m, true)
		}

		if wxsign != signature {
			return fmt.Errorf("signature verified failed, want: %s, got: %s", signature, wxsign)
		}
	}

	if appid, ok := m["appid"]; ok {
		if appid != mch.appid {
			return fmt.Errorf("appid mismatch, want: %s, got: %s", mch.appid, m["appid"])
		}
	}

	if mchid, ok := m["mch_id"]; ok {
		if mchid != mch.mchid {
			return fmt.Errorf("mchid mismatch, want: %s, got: %s", mch.mchid, m["mch_id"])
		}
	}

	return nil
}

// DecryptWithAES256ECB AES-256-ECB解密（主要用于退款结果通知）
func (mch *Mch) DecryptWithAES256ECB(encrypt string) (wx.WXML, error) {
	cipherText, err := base64.StdEncoding.DecodeString(encrypt)

	if err != nil {
		return nil, err
	}

	h := md5.New()
	h.Write([]byte(mch.apikey))

	ecb := wx.NewECBCrypto([]byte(hex.EncodeToString(h.Sum(nil))), wx.PKCS7)

	plainText, err := ecb.Decrypt(cipherText)

	if err != nil {
		return nil, err
	}

	return wx.ParseXML2Map(plainText)
}

func (mch *Mch) pkcs12ToPem(p12 []byte) (tls.Certificate, error) {
	blocks, err := pkcs12.ToPEM(p12, mch.mchid)

	if err != nil {
		return tls.Certificate{}, err
	}

	pemData := make([]byte, 0)

	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}

	// then use PEM data for tls to construct tls certificate:
	return tls.X509KeyPair(pemData, pemData)
}

// Sign 生成签名
func (mch *Mch) buildSignStr(m wx.WXML) string {
	l := len(m)

	ks := make([]string, 0, l)
	kvs := make([]string, 0, l)

	for k := range m {
		if k == "sign" {
			continue
		}

		ks = append(ks, k)
	}

	sort.Strings(ks)

	for _, k := range ks {
		if v, ok := m[k]; ok && v != "" {
			kvs = append(kvs, fmt.Sprintf("%s=%s", k, v))
		}
	}

	kvs = append(kvs, fmt.Sprintf("key=%s", mch.apikey))

	return strings.Join(kvs, "&")
}
