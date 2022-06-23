package mch

import (
	"context"
	"crypto/md5"
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/chenghonour/gochat/urls"
	"github.com/chenghonour/gochat/wx"
)

// Mch 微信支付
type Mch struct {
	mchid  string
	apikey string
	nonce  func() string
	client wx.Client
	tlscli wx.Client
}

// New returns new wechat pay
// [证书参考](https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=4_3)
func New(mchid, apikey string, certs ...tls.Certificate) *Mch {
	return &Mch{
		mchid:  mchid,
		apikey: apikey,
		nonce: func() string {
			return wx.Nonce(16)
		},
		client: wx.DefaultClient(),
		tlscli: wx.DefaultClient(certs...),
	}
}

// SetClient sets options for wechat client
func (mch *Mch) SetClient(options ...wx.ClientOption) {
	mch.client.Set(options...)
}

// SetTLSClient sets options for wechat tls client
func (mch *Mch) SetTLSClient(options ...wx.ClientOption) {
	mch.tlscli.Set(options...)
}

// MchID returns mchid
func (mch *Mch) MchID() string {
	return mch.mchid
}

// ApiKey returns apikey
func (mch *Mch) ApiKey() string {
	return mch.apikey
}

// Do exec action
func (mch *Mch) Do(ctx context.Context, action wx.Action, options ...wx.HTTPOption) (wx.WXML, error) {
	m, err := action.WXML(mch.mchid, mch.nonce())

	if err != nil {
		return nil, err
	}

	// 签名
	m["sign"] = mch.Sign(SignType(m["sign_type"]), m, true)

	if len(action.Method()) == 0 {
		if len(action.URL()) == 0 {
			return m, nil
		}

		query := url.Values{}

		for k, v := range m {
			query.Add(k, v)
		}

		return wx.WXML{"entrust_url": fmt.Sprintf("%s?%s", action.URL(), query.Encode())}, nil
	}

	body, err := wx.FormatMap2XML(m)

	if err != nil {
		return nil, err
	}

	var resp []byte

	if action.IsTLS() {
		resp, err = mch.tlscli.Do(ctx, action.Method(), action.URL(), body, options...)
	} else {
		resp, err = mch.client.Do(ctx, action.Method(), action.URL(), body, options...)
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
func (mch *Mch) APPAPI(appid, prepayID string) wx.WXML {
	m := wx.WXML{
		"appid":     appid,
		"partnerid": mch.mchid,
		"prepayid":  prepayID,
		"package":   "Sign=WXPay",
		"noncestr":  mch.nonce(),
		"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
	}

	m["sign"] = mch.Sign(SignMD5, m, true)

	return m
}

// JSAPI 用于JS拉起支付
func (mch *Mch) JSAPI(appid, prepayID string) wx.WXML {
	m := wx.WXML{
		"appId":     appid,
		"nonceStr":  mch.nonce(),
		"package":   fmt.Sprintf("prepay_id=%s", prepayID),
		"signType":  string(SignMD5),
		"timeStamp": strconv.FormatInt(time.Now().Unix(), 10),
	}

	m["paySign"] = mch.Sign(SignMD5, m, true)

	return m
}

// MinipRedpackJSAPI 小程序领取红包
func (mch *Mch) MinipRedpackJSAPI(appid, pkg string) wx.WXML {
	m := wx.WXML{
		"nonceStr":  mch.nonce(),
		"package":   url.QueryEscape(pkg),
		"timeStamp": strconv.FormatInt(time.Now().Unix(), 10),
		"signType":  string(SignMD5),
	}

	signStr := fmt.Sprintf("appId=%s&nonceStr=%s&package=%s&timeStamp=%s&key=%s", appid, m["nonceStr"], m["package"], m["timeStamp"], mch.apikey)

	m["paySign"] = wx.MD5(signStr)

	return m
}

// DownloadBill 下载交易账单
// 账单日期格式：20140603
func (mch *Mch) DownloadBill(ctx context.Context, appid, billDate, billType string) ([]byte, error) {
	m := wx.WXML{
		"appid":     appid,
		"mch_id":    mch.mchid,
		"bill_date": billDate,
		"bill_type": billType,
		"nonce_str": mch.nonce(),
	}

	m["sign"] = mch.Sign(SignMD5, m, true)

	body, err := wx.FormatMap2XML(m)

	if err != nil {
		return nil, err
	}

	resp, err := mch.client.Do(ctx, http.MethodPost, urls.MchDownloadBill, body, wx.WithHTTPClose())

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
func (mch *Mch) DownloadFundFlow(ctx context.Context, appid string, billDate, accountType string) ([]byte, error) {
	m := wx.WXML{
		"appid":        appid,
		"mch_id":       mch.mchid,
		"bill_date":    billDate,
		"account_type": accountType,
		"nonce_str":    mch.nonce(),
	}

	m["sign"] = mch.Sign(SignHMacSHA256, m, true)

	body, err := wx.FormatMap2XML(m)

	if err != nil {
		return nil, err
	}

	resp, err := mch.tlscli.Do(ctx, http.MethodPost, urls.MchDownloadFundFlow, body, wx.WithHTTPClose())

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
func (mch *Mch) BatchQueryComment(ctx context.Context, appid, beginTime, endTime string, offset int, limit ...int) ([]byte, error) {
	m := wx.WXML{
		"appid":      appid,
		"mch_id":     mch.mchid,
		"begin_time": beginTime,
		"end_time":   endTime,
		"offset":     strconv.Itoa(offset),
		"nonce_str":  mch.nonce(),
	}

	if len(limit) != 0 {
		m["limit"] = strconv.Itoa(limit[0])
	}

	m["sign"] = mch.Sign(SignHMacSHA256, m, true)

	body, err := wx.FormatMap2XML(m)

	if err != nil {
		return nil, err
	}

	resp, err := mch.tlscli.Do(ctx, http.MethodPost, urls.MchBatchQueryComment, body, wx.WithHTTPClose())

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

// Sign 生成签名
func (mch *Mch) Sign(t SignType, m wx.WXML, toUpper bool) string {
	str := mch.buildSignStr(m)
	sign := ""

	if t == SignHMacSHA256 {
		sign = wx.HMacSHA256(str, mch.apikey)
	} else {
		sign = wx.MD5(str)
	}

	if toUpper {
		sign = strings.ToUpper(sign)
	}

	return sign
}

// VerifyWXMLResult 微信请求/回调通知签名验证
func (mch *Mch) VerifyWXMLResult(m wx.WXML) error {
	if wxsign, ok := m["sign"]; ok {
		t := SignMD5

		if v, ok := m["sign_type"]; ok {
			t = SignType(v)
		}

		if signature := mch.Sign(t, m, true); wxsign != signature {
			return fmt.Errorf("signature verified failed, want: %s, got: %s", signature, wxsign)
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
