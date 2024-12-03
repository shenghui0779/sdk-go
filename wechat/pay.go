package wechat

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/shenghui0779/sdk-go/lib"
	"github.com/shenghui0779/sdk-go/lib/value"
	"github.com/shenghui0779/sdk-go/lib/xcrypto"
	"github.com/shenghui0779/sdk-go/lib/xhash"
)

// Pay 微信支付
type Pay struct {
	host      string
	mchid     string
	apikey    string
	client    *resty.Client
	clientTls *resty.Client
	logger    func(ctx context.Context, err error, data map[string]string)
}

// MchID 返回mchid
func (p *Pay) MchID() string {
	return p.mchid
}

// ApiKey 返回apikey
func (p *Pay) ApiKey() string {
	return p.apikey
}

func (p *Pay) url(path string, query url.Values) string {
	var builder strings.Builder

	builder.WriteString(p.host)
	if len(path) != 0 && path[0] != '/' {
		builder.WriteString("/")
	}
	builder.WriteString(path)
	if len(query) != 0 {
		builder.WriteString("?")
		builder.WriteString(query.Encode())
	}

	return builder.String()
}

func (p *Pay) do(ctx context.Context, path string, params value.V) ([]byte, error) {
	reqURL := p.url(path, nil)

	log := lib.NewReqLog(http.MethodPost, reqURL)
	defer log.Do(ctx, p.logger)

	params.Set("sign", p.Sign(params))

	body, err := ValueToXML(params)
	if err != nil {
		log.SetError(err)
		return nil, err
	}
	log.SetReqBody(body)

	resp, err := p.client.R().
		SetContext(ctx).
		SetBody(body).
		Post(reqURL)
	if err != nil {
		log.SetError(err)
		return nil, err
	}
	log.SetRespHeader(resp.Header())
	log.SetStatusCode(resp.StatusCode())
	log.SetRespBody(string(resp.Body()))
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", resp.StatusCode())
	}
	return resp.Body(), nil
}

func (p *Pay) doTls(ctx context.Context, path string, params value.V) ([]byte, error) {
	reqURL := p.url(path, nil)

	log := lib.NewReqLog(http.MethodPost, reqURL)
	defer log.Do(ctx, p.logger)

	params.Set("sign", p.Sign(params))

	body, err := ValueToXML(params)
	if err != nil {
		log.SetError(err)
		return nil, err
	}
	log.SetReqBody(body)

	resp, err := p.clientTls.R().
		SetContext(ctx).
		SetBody(body).
		Post(reqURL)
	if err != nil {
		log.SetError(err)
		return nil, err
	}
	log.SetRespHeader(resp.Header())
	log.SetStatusCode(resp.StatusCode())
	log.SetRespBody(string(resp.Body()))
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", resp.StatusCode())
	}
	return resp.Body(), nil
}

// PostXML POST请求XML数据 (无证书请求)
func (p *Pay) PostXML(ctx context.Context, path string, params value.V) (value.V, error) {
	b, err := p.do(ctx, path, params)
	if err != nil {
		return nil, err
	}

	ret, err := XMLToValue(b)
	if err != nil {
		return nil, err
	}
	if code := ret.Get("return_code"); code != ResultSuccess {
		return nil, fmt.Errorf("%s | %s", code, ret.Get("return_msg"))
	}
	if err = p.Verify(ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// PostTLSXML POST请求XML数据 (带证书请求)
func (p *Pay) PostTLSXML(ctx context.Context, path string, params value.V) (value.V, error) {
	b, err := p.doTls(ctx, path, params)
	if err != nil {
		return nil, err
	}

	ret, err := XMLToValue(b)
	if err != nil {
		return nil, err
	}
	if code := ret.Get("return_code"); code != ResultSuccess {
		return nil, fmt.Errorf("%s | %s", code, ret.Get("return_msg"))
	}
	if err = p.Verify(ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// PostBuffer POST请求获取buffer (无证书请求，如：下载交易订单)
func (p *Pay) PostBuffer(ctx context.Context, path string, params value.V) ([]byte, error) {
	b, err := p.do(ctx, path, params)
	if err != nil {
		return nil, err
	}

	ret, err := XMLToValue(b)
	if err != nil {
		return nil, err
	}
	// 能解析出XML，说明发生错误
	if len(ret) != 0 {
		return nil, fmt.Errorf("%s | %s (error_code = %s, err_code_des = %s)", ret.Get("return_code"), ret.Get("return_msg"), ret.Get("error_code"), ret.Get("err_code_des"))
	}
	return b, nil
}

// PostBuffer POST请求获取buffer (带证书请求，如：下载资金账单)
func (p *Pay) PostTlsBuffer(ctx context.Context, path string, params value.V) ([]byte, error) {
	b, err := p.doTls(ctx, path, params)
	if err != nil {
		return nil, err
	}

	ret, err := XMLToValue(b)
	if err != nil {
		return nil, err
	}
	// 能解析出XML，说明发生错误
	if len(ret) != 0 {
		return nil, fmt.Errorf("%s | %s | %s", ret.Get("return_code"), ret.Get("return_msg"), ret.Get("error_code"))
	}
	return b, nil
}

func (p *Pay) Sign(v value.V) string {
	signStr := v.Encode("=", "&", value.WithIgnoreKeys("sign"), value.WithEmptyMode(value.EmptyIgnore)) + "&key=" + p.apikey
	signType := v.Get("sign_type")
	if len(signType) == 0 {
		signType = v.Get("signType")
	}
	if len(signType) != 0 && SignAlgo(strings.ToUpper(signType)) == SignHMacSHA256 {
		return strings.ToUpper(xhash.HMacSHA256(p.apikey, signStr))
	}
	return strings.ToUpper(xhash.MD5(signStr))
}

func (p *Pay) Verify(v value.V) error {
	wxsign := v.Get("sign")
	signType := v.Get("sign_type")
	if len(signType) == 0 {
		signType = v.Get("signType")
	}
	signStr := v.Encode("=", "&", value.WithIgnoreKeys("sign"), value.WithEmptyMode(value.EmptyIgnore)) + "&key=" + p.apikey
	// hmac-sha256
	if len(signType) != 0 && SignAlgo(strings.ToUpper(signType)) == SignHMacSHA256 {
		if sign := strings.ToUpper(xhash.HMacSHA256(p.apikey, signStr)); sign != wxsign {
			return fmt.Errorf("sign verify failed, expect = %s, actual = %s", sign, wxsign)
		}
		return nil
	}
	// md5
	if sign := strings.ToUpper(xhash.MD5(signStr)); sign != wxsign {
		return fmt.Errorf("sign verify failed, expect = %s, actual = %s", sign, wxsign)
	}
	return nil
}

// DecryptRefund 退款结果通知解密
func (p *Pay) DecryptRefund(encrypt string) (value.V, error) {
	cipherText, err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		return nil, err
	}
	plainText, err := xcrypto.AESDecryptECB([]byte(xhash.MD5(p.apikey)), cipherText)
	if err != nil {
		return nil, err
	}
	return XMLToValue(plainText)
}

// APPAPI 用于APP拉起支付
func (p *Pay) APPAPI(appid, prepayID string) value.V {
	v := value.V{}

	v.Set("appid", appid)
	v.Set("partnerid", p.mchid)
	v.Set("prepayid", prepayID)
	v.Set("package", "Sign=WXPay")
	v.Set("noncestr", lib.Nonce(16))
	v.Set("timestamp", strconv.FormatInt(time.Now().Unix(), 10))

	v.Set("sign", p.Sign(v))

	return v
}

// JSAPI 用于JS拉起支付
func (p *Pay) JSAPI(appid, prepayID string) value.V {
	v := value.V{}

	v.Set("appId", appid)
	v.Set("nonceStr", lib.Nonce(16))
	v.Set("package", "prepay_id="+prepayID)
	v.Set("signType", "MD5")
	v.Set("timeStamp", strconv.FormatInt(time.Now().Unix(), 10))

	v.Set("paySign", p.Sign(v))

	return v
}

// MinipRedpackJSAPI 小程序领取红包
func (p *Pay) MinipRedpackJSAPI(appid, pkg string) value.V {
	v := value.V{}

	v.Set("appId", appid)
	v.Set("nonceStr", lib.Nonce(16))
	v.Set("package", url.QueryEscape(pkg))
	v.Set("timeStamp", strconv.FormatInt(time.Now().Unix(), 10))
	v.Set("signType", "MD5")

	signStr := fmt.Sprintf("appId=%s&nonceStr=%s&package=%s&timeStamp=%s&key=%s", appid, v.Get("nonceStr"), v.Get("package"), v.Get("timeStamp"), p.apikey)

	v.Set("paySign", xhash.MD5(signStr))

	return v
}

// PayOption 微信支付设置项
type PayOption func(p *Pay)

// WithPayTLSCert 设置支付TLS证书
func WithPayCert(cert tls.Certificate) PayOption {
	return func(p *Pay) {
		p.clientTls.SetCertificates(cert)
	}
}

// WithPayClient 设置支付无证书 HTTP Client
func WithPayClient(cli *http.Client) PayOption {
	return func(p *Pay) {
		p.client = resty.NewWithClient(cli)
	}
}

// WithPayTlsClient 设置支付带证书 HTTP Client
func WithPayTlsClient(cli *http.Client) PayOption {
	return func(p *Pay) {
		p.clientTls = resty.NewWithClient(cli)
	}
}

// WithPayLogger 设置支付日志记录
func WithPayLogger(fn func(ctx context.Context, err error, data map[string]string)) PayOption {
	return func(p *Pay) {
		p.logger = fn
	}
}

// NewPay 生成一个微信支付实例
func NewPay(mchid, apikey string, options ...PayOption) *Pay {
	pay := &Pay{
		host:      "https://api.mch.weixin.qq.com",
		mchid:     mchid,
		apikey:    apikey,
		client:    lib.NewClient(),
		clientTls: lib.NewClient(),
	}
	for _, f := range options {
		f(pay)
	}
	return pay
}
