package ysepay

import (
	"context"
	"crypto"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/uuid"
	"github.com/tidwall/gjson"

	"github.com/shenghui0779/sdk-go/lib"
	libCrypto "github.com/shenghui0779/sdk-go/lib/crypto"
	libHttp "github.com/shenghui0779/sdk-go/lib/http"
	"github.com/shenghui0779/sdk-go/lib/value"
)

// Client 银盛支付客户端
type Client struct {
	host    string
	mchNO   string
	desKey  string
	prvKey  *libCrypto.PrivateKey
	pubKey  *libCrypto.PublicKey
	httpCli libHttp.HTTPClient
	logger  func(ctx context.Context, data map[string]string)
}

// MchNO 返回商户号
func (c *Client) MchNO() string {
	return c.mchNO
}

func (c *Client) url(api string) string {
	var builder strings.Builder

	builder.WriteString(c.host)
	builder.WriteString("/api/")
	builder.WriteString(api)

	return builder.String()
}

// Encrypt 敏感数据DES加密
func (c *Client) Encrypt(plain string) (string, error) {
	b, err := libCrypto.DESEncryptECB([]byte(c.desKey), []byte(plain))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}

// MustEncrypt 敏感数据DES加密；若发生错误，则Panic
func (c *Client) MustEncrypt(plain string) string {
	b, err := libCrypto.DESEncryptECB([]byte(c.desKey), []byte(plain))
	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(b)
}

// Decrypt 敏感数据DES解密
func (c *Client) Decrypt(cipher string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(cipher)
	if err != nil {
		return "", err
	}

	plain, err := libCrypto.DESEncryptECB([]byte(c.desKey), b)
	if err != nil {
		return "", err
	}

	return string(plain), nil
}

// PostForm 发送POST表单请求
func (c *Client) PostForm(ctx context.Context, api, serviceNO string, bizData value.V) (gjson.Result, error) {
	reqURL := c.url(api)

	log := lib.NewReqLog(http.MethodPost, reqURL)
	defer log.Do(ctx, c.logger)

	form, err := c.reqForm(uuid.NewString(), serviceNO, bizData)
	if err != nil {
		return fail(err)
	}

	log.SetReqBody(form)

	resp, err := c.httpCli.Do(ctx, http.MethodPost, reqURL, []byte(form), libHttp.WithHeader(libHttp.HeaderContentType, libHttp.ContentForm))
	if err != nil {
		return fail(err)
	}
	defer resp.Body.Close()

	log.SetRespHeader(resp.Header)
	log.SetStatusCode(resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		return fail(fmt.Errorf("HTTP Request Error, StatusCode = %d", resp.StatusCode))
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return fail(err)
	}

	log.SetRespBody(string(b))

	return c.verifyResp(b)
}

// reqForm 生成请求表单
func (c *Client) reqForm(reqID, serviceNO string, bizData value.V) (string, error) {
	if c.prvKey == nil {
		return "", errors.New("private key is nil (forgotten configure?)")
	}

	v := value.V{}

	v.Set("requestId", reqID)
	v.Set("srcMerchantNo", c.mchNO)
	v.Set("version", "v2.0.0")
	v.Set("charset", "UTF-8")
	v.Set("serviceNo", serviceNO)
	v.Set("signType", "RSA")

	if len(bizData) != 0 {
		bizByte, err := json.Marshal(bizData)
		if err != nil {
			return "", err
		}

		v.Set("bizReqJson", string(bizByte))
	}

	sign, err := c.prvKey.Sign(crypto.SHA1, []byte(v.Encode("=", "&", value.WithIgnoreKeys("sign"), value.WithEmptyMode(value.EmptyIgnore))))
	if err != nil {
		return "", err
	}

	v.Set("sign", base64.StdEncoding.EncodeToString(sign))

	return v.Encode("=", "&", value.WithEmptyMode(value.EmptyIgnore), value.WithKVEscape()), nil
}

func (c *Client) verifyResp(body []byte) (gjson.Result, error) {
	if c.pubKey == nil {
		return fail(errors.New("public key is nil (forgotten configure?)"))
	}

	ret := gjson.ParseBytes(body)

	sign, err := base64.StdEncoding.DecodeString(ret.Get("sign").String())
	if err != nil {
		return fail(err)
	}

	v := value.V{}

	v.Set("requestId", ret.Get("requestId").String())
	v.Set("code", ret.Get("code").String())
	v.Set("msg", ret.Get("msg").String())
	v.Set("bizResponseJson", ret.Get("bizResponseJson").String())

	err = c.pubKey.Verify(crypto.SHA1, []byte(v.Encode("=", "&", value.WithEmptyMode(value.EmptyIgnore))), sign)
	if err != nil {
		return fail(err)
	}

	if code := ret.Get("code").String(); code != SysOK {
		if code == SysAccepting {
			return fail(ErrSysAccepting)
		}

		return fail(fmt.Errorf("%s | %s", code, ret.Get("msg").String()))
	}

	return ret.Get("bizResponseJson"), nil
}

// VerifyNotify 解析并验证异步回调通知，返回BizJSON数据
func (c *Client) VerifyNotify(form url.Values) (gjson.Result, error) {
	if c.pubKey == nil {
		return fail(errors.New("public key is nil (forgotten configure?)"))
	}

	sign, err := base64.StdEncoding.DecodeString(form.Get("sign"))
	if err != nil {
		return fail(err)
	}

	v := value.V{}

	v.Set("requestId", form.Get("requestId"))
	v.Set("version", form.Get("version"))
	v.Set("charset", form.Get("charset"))
	v.Set("serviceNo", form.Get("serviceNo"))
	v.Set("signType", form.Get("signType"))
	v.Set("bizResponseJson", form.Get("bizResponseJson"))

	err = c.pubKey.Verify(crypto.SHA1, []byte(v.Encode("=", "&", value.WithEmptyMode(value.EmptyIgnore))), sign)
	if err != nil {
		return fail(err)
	}

	return gjson.Parse(form.Get("bizResponseJson")), nil
}

// Option 自定义设置项
type Option func(c *Client)

// WithHttpCli 设置自定义 HTTP Client
func WithHttpCli(cli *http.Client) Option {
	return func(c *Client) {
		c.httpCli = libHttp.NewHTTPClient(cli)
	}
}

// WithPrivateKey 设置商户RSA私钥
func WithPrivateKey(key *libCrypto.PrivateKey) Option {
	return func(c *Client) {
		c.prvKey = key
	}
}

// WithPublicKey 设置平台RSA公钥
func WithPublicKey(key *libCrypto.PublicKey) Option {
	return func(c *Client) {
		c.pubKey = key
	}
}

// WithLogger 设置日志记录
func WithLogger(f func(ctx context.Context, data map[string]string)) Option {
	return func(c *Client) {
		c.logger = f
	}
}

// NewClient 生成银盛支付客户端
func NewClient(mchNO, desKey string, options ...Option) *Client {
	c := &Client{
		host:    "https://eqt.ysepay.com",
		mchNO:   mchNO,
		desKey:  desKey,
		httpCli: libHttp.NewDefaultClient(),
	}

	for _, f := range options {
		f(c)
	}

	return c
}
