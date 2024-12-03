package alipay

import (
	"context"
	"crypto"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"

	"github.com/shenghui0779/sdk-go/lib"
	"github.com/shenghui0779/sdk-go/lib/value"
	"github.com/shenghui0779/sdk-go/lib/xcrypto"
)

// Client 支付宝客户端
type Client struct {
	gateway string
	appid   string
	aesKey  string
	prvKey  *xcrypto.PrivateKey
	pubKey  *xcrypto.PublicKey
	client  *resty.Client
	logger  func(ctx context.Context, err error, data map[string]string)
}

// AppID 返回appid
func (c *Client) AppID() string {
	return c.appid
}

// Do 向支付宝网关发送请求
func (c *Client) Do(ctx context.Context, method string, options ...ActionOption) (gjson.Result, error) {
	reqURL := c.gateway + "?charset=utf-8"

	log := lib.NewReqLog(http.MethodPost, reqURL)
	defer log.Do(ctx, c.logger)

	action := NewAction(method, options...)

	body, err := action.Encode(c)
	if err != nil {
		log.SetError(err)
		return lib.Fail(err)
	}
	log.SetReqBody(body)

	resp, err := c.client.R().
		SetContext(ctx).
		SetHeader(lib.HeaderAccept, lib.ContentJSON).
		SetHeader(lib.HeaderContentType, lib.ContentForm).
		SetBody(body).
		Post(reqURL)
	if err != nil {
		log.SetError(err)
		return lib.Fail(err)
	}
	log.SetRespHeader(resp.Header())
	log.SetStatusCode(resp.StatusCode())
	log.SetRespBody(string(resp.Body()))
	if !resp.IsSuccess() {
		return lib.Fail(fmt.Errorf("HTTP Request Error, StatusCode = %d", resp.StatusCode()))
	}

	// 签名校验
	ret, err := c.verifyResp(action.RespKey(), resp.Body())
	if err != nil {
		log.SetError(err)
		return lib.Fail(err)
	}

	// JSON串，无需解密
	if strings.HasPrefix(ret.String(), "{") {
		if code := ret.Get("code").String(); code != CodeOK {
			return lib.Fail(fmt.Errorf("%s | %s (sub_code = %s, sub_msg = %s)", code, ret.Get("msg").String(), ret.Get("sub_code").String(), ret.Get("sub_msg").String()))
		}
		return ret, nil
	}

	// 非JSON串，需解密
	data, err := c.Decrypt(ret.String())
	if err != nil {
		log.SetError(err)
		return lib.Fail(err)
	}
	log.Set("decrypt", string(data))

	return gjson.ParseBytes(data), nil
}

// Upload 文件上传，参考：https://opendocs.alipay.com/apis/api_4/alipay.merchant.item.file.upload
func (c *Client) Upload(ctx context.Context, method string, fieldName, filePath string, formData map[string]string, options ...ActionOption) (gjson.Result, error) {
	log := lib.NewReqLog(http.MethodPost, c.gateway)
	defer log.Do(ctx, c.logger)

	action := NewAction(method, options...)

	query, err := action.Encode(c)
	if err != nil {
		log.SetError(err)
		return lib.Fail(err)
	}
	log.Set("query", query)

	resp, err := c.client.R().
		SetContext(ctx).
		SetHeader(lib.HeaderAccept, lib.ContentJSON).
		SetFile(fieldName, filePath).
		SetMultipartFormData(formData).
		Post(c.gateway + "?" + query)
	if err != nil {
		log.SetError(err)
		return lib.Fail(err)
	}
	log.SetRespHeader(resp.Header())
	log.SetStatusCode(resp.StatusCode())
	log.SetRespBody(string(resp.Body()))
	if !resp.IsSuccess() {
		return lib.Fail(fmt.Errorf("HTTP Request Error, StatusCode = %d", resp.StatusCode()))
	}

	// 签名校验
	ret, err := c.verifyResp(action.RespKey(), resp.Body())
	if err != nil {
		log.SetError(err)
		return lib.Fail(err)
	}

	// JSON串，无需解密
	if strings.HasPrefix(ret.String(), "{") {
		if code := ret.Get("code").String(); code != CodeOK {
			return lib.Fail(fmt.Errorf("%s | %s (sub_code = %s, sub_msg = %s)", code, ret.Get("msg").String(), ret.Get("sub_code").String(), ret.Get("sub_msg").String()))
		}
		return ret, nil
	}

	// 非JSON串，需解密
	data, err := c.Decrypt(ret.String())
	if err != nil {
		log.SetError(err)
		return lib.Fail(err)
	}
	log.Set("decrypt", string(data))

	return gjson.ParseBytes(data), nil
}

// UploadWithReader 文件上传，参考：https://opendocs.alipay.com/apis/api_4/alipay.merchant.item.file.upload
func (c *Client) UploadWithReader(ctx context.Context, method string, fieldName, fileName string, reader io.Reader, formData map[string]string, options ...ActionOption) (gjson.Result, error) {
	log := lib.NewReqLog(http.MethodPost, c.gateway)
	defer log.Do(ctx, c.logger)

	action := NewAction(method, options...)

	query, err := action.Encode(c)
	if err != nil {
		log.SetError(err)
		return lib.Fail(err)
	}
	log.Set("query", query)

	resp, err := c.client.R().
		SetContext(ctx).
		SetHeader(lib.HeaderAccept, lib.ContentJSON).
		SetMultipartField(fieldName, fileName, "", reader).
		SetMultipartFormData(formData).
		Post(c.gateway + "?" + query)
	if err != nil {
		log.SetError(err)
		return lib.Fail(err)
	}
	log.SetRespHeader(resp.Header())
	log.SetStatusCode(resp.StatusCode())
	log.SetRespBody(string(resp.Body()))
	if !resp.IsSuccess() {
		return lib.Fail(fmt.Errorf("HTTP Request Error, StatusCode = %d", resp.StatusCode()))
	}

	// 签名校验
	ret, err := c.verifyResp(action.RespKey(), resp.Body())
	if err != nil {
		log.SetError(err)
		return lib.Fail(err)
	}

	// JSON串，无需解密
	if strings.HasPrefix(ret.String(), "{") {
		if code := ret.Get("code").String(); code != CodeOK {
			return lib.Fail(fmt.Errorf("%s | %s (sub_code = %s, sub_msg = %s)", code, ret.Get("msg").String(), ret.Get("sub_code").String(), ret.Get("sub_msg").String()))
		}
		return ret, nil
	}

	// 非JSON串，需解密
	data, err := c.Decrypt(ret.String())
	if err != nil {
		log.SetError(err)
		return lib.Fail(err)
	}
	log.Set("decrypt", string(data))

	return gjson.ParseBytes(data), nil
}

func (c *Client) verifyResp(key string, body []byte) (gjson.Result, error) {
	if c.pubKey == nil {
		return lib.Fail(errors.New("public key is nil (forgotten configure?)"))
	}

	ret := gjson.ParseBytes(body)

	signByte, err := base64.StdEncoding.DecodeString(ret.Get("sign").String())
	if err != nil {
		return lib.Fail(err)
	}

	hash := crypto.SHA256
	if ret.Get("sign_type").String() == "RSA" {
		hash = crypto.SHA1
	}

	if errResp := ret.Get("error_response"); errResp.Exists() {
		if err = c.pubKey.Verify(hash, []byte(errResp.Raw), signByte); err != nil {
			return lib.Fail(err)
		}

		return lib.Fail(fmt.Errorf("%s | %s (sub_code = %s, sub_msg = %s)",
			errResp.Get("code").String(),
			errResp.Get("msg").String(),
			errResp.Get("sub_code").String(),
			errResp.Get("sub_msg").String(),
		))
	}

	resp := ret.Get(key)
	if err = c.pubKey.Verify(hash, []byte(resp.Raw), signByte); err != nil {
		return lib.Fail(err)
	}
	return resp, nil
}

// PageExecute 致敬官方SDK
func (c *Client) PageExecute(method string, options ...ActionOption) (string, error) {
	action := NewAction(method, options...)
	query, err := action.Encode(c)
	if err != nil {
		return "", err
	}
	return c.gateway + "?" + query, nil
}

// Encrypt 数据加密
func (c *Client) Encrypt(data string) (string, error) {
	key, err := base64.StdEncoding.DecodeString(c.aesKey)
	if err != nil {
		return "", fmt.Errorf("aes_key base64.decode error: %w", err)
	}
	ct, err := xcrypto.AESEncryptCBC(key, make([]byte, 16), []byte(data))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(ct.Bytes()), nil
}

// Decrypt 数据解密
func (c *Client) Decrypt(encryptData string) ([]byte, error) {
	key, err := base64.StdEncoding.DecodeString(c.aesKey)
	if err != nil {
		return nil, fmt.Errorf("aes_key base64.decode error: %w", err)
	}
	data, err := base64.StdEncoding.DecodeString(encryptData)
	if err != nil {
		return nil, fmt.Errorf("encrypt_data base64.decode error: %w", err)
	}
	return xcrypto.AESDecryptCBC(key, make([]byte, 16), data)
}

// DecodeEncryptData 解析加密数据，如：授权的用户信息和手机号
func (c *Client) DecodeEncryptData(hash crypto.Hash, data, sign string) ([]byte, error) {
	if c.pubKey == nil {
		return nil, errors.New("public key is nil (forgotten configure?)")
	}

	signByte, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return nil, fmt.Errorf("sign base64.decode error: %w", err)
	}
	if err = c.pubKey.Verify(hash, []byte(`"`+data+`"`), signByte); err != nil {
		return nil, fmt.Errorf("sign verified error: %w", err)
	}
	return c.Decrypt(data)
}

// VerifyNotify 验证回调通知表单数据
func (c *Client) VerifyNotify(form url.Values) (value.V, error) {
	if c.pubKey == nil {
		return nil, errors.New("public key is nil (forgotten configure?)")
	}

	sign, err := base64.StdEncoding.DecodeString(form.Get("sign"))
	if err != nil {
		return nil, err
	}

	v := value.V{}
	for key, vals := range form {
		if key == "sign_type" || key == "sign" || len(vals) == 0 {
			continue
		}
		v.Set(key, vals[0])
	}
	str := v.Encode("=", "&", value.WithEmptyMode(value.EmptyIgnore))

	hash := crypto.SHA256
	if form.Get("sign_type") == "RSA" {
		hash = crypto.SHA1
	}
	if err = c.pubKey.Verify(hash, []byte(str), sign); err != nil {
		return nil, err
	}
	return v, nil
}

// Option 自定义设置项
type Option func(c *Client)

// WithHttpClient 设置自定义 HTTP Client
func WithHttpClient(cli *http.Client) Option {
	return func(c *Client) {
		c.client = resty.NewWithClient(cli)
	}
}

// WithPrivateKey 设置商户RSA私钥
func WithPrivateKey(key *xcrypto.PrivateKey) Option {
	return func(c *Client) {
		c.prvKey = key
	}
}

// WithPublicKey 设置平台RSA公钥
func WithPublicKey(key *xcrypto.PublicKey) Option {
	return func(c *Client) {
		c.pubKey = key
	}
}

// WithLogger 设置日志记录
func WithLogger(fn func(ctx context.Context, err error, data map[string]string)) Option {
	return func(c *Client) {
		c.logger = fn
	}
}

// NewClient 生成支付宝客户端
func NewClient(appid, aesKey string, options ...Option) *Client {
	c := &Client{
		appid:   appid,
		aesKey:  aesKey,
		gateway: "https://openapi.alipay.com/gateway.do",
		client:  lib.NewClient(),
	}
	for _, f := range options {
		f(c)
	}
	return c
}

// NewSandbox 生成支付宝沙箱环境
func NewSandbox(appid, aesKey string, options ...Option) *Client {
	c := &Client{
		appid:   appid,
		aesKey:  aesKey,
		gateway: "https://openapi-sandbox.dl.alipaydev.com/gateway.do",
		client:  lib.NewClient(),
	}
	for _, f := range options {
		f(c)
	}
	return c
}
