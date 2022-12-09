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
	"strconv"
	"strings"
	"time"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

// Mch 微信支付
type Mch struct {
	mchid  string
	apikey string
	nonce  func() string
	client wx.HTTPClient
	tlscli wx.HTTPClient
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
	m, err := action.WXML(mch.mchid, mch.apikey, mch.nonce())

	if err != nil {
		return nil, err
	}

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
	// body, err := wx.FormatMap2XMLForTest(m) // 运行单元测试时使用

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

	m["sign"] = wx.SignMD5.Do(mch.apikey, m, true)

	return m
}

// JSAPI 用于JS拉起支付
func (mch *Mch) JSAPI(appid, prepayID string) wx.WXML {
	m := wx.WXML{
		"appId":     appid,
		"nonceStr":  mch.nonce(),
		"package":   fmt.Sprintf("prepay_id=%s", prepayID),
		"signType":  "MD5",
		"timeStamp": strconv.FormatInt(time.Now().Unix(), 10),
	}

	m["paySign"] = wx.SignMD5.Do(mch.apikey, m, true)

	return m
}

// MinipRedpackJSAPI 小程序领取红包
func (mch *Mch) MinipRedpackJSAPI(appid, pkg string) wx.WXML {
	m := wx.WXML{
		"appId":     appid,
		"nonceStr":  mch.nonce(),
		"package":   url.QueryEscape(pkg),
		"timeStamp": strconv.FormatInt(time.Now().Unix(), 10),
		"signType":  "MD5",
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

	m["sign"] = wx.SignMD5.Do(mch.apikey, m, true)

	body, err := wx.FormatMap2XML(m)
	// body, err := wx.FormatMap2XMLForTest(m) // 运行单元测试时使用

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

	m["sign"] = wx.SignHMacSHA256.Do(mch.apikey, m, true)

	body, err := wx.FormatMap2XML(m)
	// body, err := wx.FormatMap2XMLForTest(m) // 运行单元测试时使用

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

	m["sign"] = wx.SignHMacSHA256.Do(mch.apikey, m, true)

	body, err := wx.FormatMap2XML(m)
	// body, err := wx.FormatMap2XMLForTest(m) // 运行单元测试时使用

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

// VerifyWXMLResult 微信请求/回调通知签名验证
func (mch *Mch) VerifyWXMLResult(m wx.WXML) error {
	if wxsign, ok := m["sign"]; ok {
		st := wx.SignMD5

		if v, ok := m["sign_type"]; ok {
			st = wx.SignType(strings.ToUpper(v))
		}

		if signature := st.Do(mch.apikey, m, true); wxsign != signature {
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

type MchOption func(mch *Mch)

// WithTLSCert 设置TLS证书
func WithTLSCert(cert tls.Certificate) MchOption {
	return func(mch *Mch) {
		mch.tlscli = wx.NewDefaultClient(cert)
	}
}

// WithNonce 设置 Nonce（加密随机串）
func WithNonce(f func() string) MchOption {
	return func(mch *Mch) {
		mch.nonce = f
	}
}

// WithClient 设置 HTTP Client
func WithClient(c *http.Client) MchOption {
	return func(mch *Mch) {
		mch.client = wx.NewHTTPClient(c)
	}
}

// WithTLSClient 设置 TLS HTTP Client（带证书）
func WithTLSClient(c *http.Client) MchOption {
	return func(mch *Mch) {
		mch.tlscli = wx.NewHTTPClient(c)
	}
}

// WithMockClient 设置 Mock Client
func WithMockClient(c wx.HTTPClient) MchOption {
	return func(mch *Mch) {
		mch.client = c
		mch.tlscli = c
	}
}

// New returns new wechat pay
// [证书参考](https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=4_3)
func New(mchid, apikey string, options ...MchOption) *Mch {
	mch := &Mch{
		mchid:  mchid,
		apikey: apikey,
		nonce: func() string {
			return wx.Nonce(16)
		},
		client: wx.NewDefaultClient(),
		tlscli: wx.NewDefaultClient(),
	}

	for _, f := range options {
		f(mch)
	}

	return mch
}
