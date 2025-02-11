package sandpay

import (
	"context"
	"crypto"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-resty/resty/v2"

	"github.com/yiigo/sdk-go/internal"
	"github.com/yiigo/sdk-go/internal/xcrypto"
)

// Client 杉德支付客户端
type Client struct {
	mchID  string
	prvKey *xcrypto.PrivateKey
	pubKey *xcrypto.PublicKey
	client *resty.Client
	logger func(ctx context.Context, err error, data map[string]string)
}

// MchID 返回商品ID
func (c *Client) MchID() string {
	return c.mchID
}

// Do 请求杉德API
func (c *Client) Do(ctx context.Context, reqURL string, form *Form) (*Form, error) {
	log := internal.NewReqLog(http.MethodPost, reqURL)
	defer log.Do(ctx, c.logger)

	body, err := form.URLEncode(c.mchID, c.prvKey)
	if err != nil {
		log.SetError(err)
		return nil, err
	}
	log.SetReqBody(body)

	resp, err := c.client.R().
		SetContext(ctx).
		SetHeader(internal.HeaderContentType, internal.ContentForm).
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

	query, err := url.QueryUnescape(string(resp.Body()))
	if err != nil {
		log.SetError(err)
		return nil, err
	}
	v, err := url.ParseQuery(query)
	if err != nil {
		log.SetError(err)
		return nil, err
	}
	return c.Verify(v)
}

// Verify 验证并解析杉德API结果或回调通知
func (c *Client) Verify(form url.Values) (*Form, error) {
	if c.pubKey == nil {
		return nil, errors.New("public key is nil (forgotten configure?)")
	}

	sign, err := base64.StdEncoding.DecodeString(strings.Replace(form.Get("sign"), " ", "+", -1))
	if err != nil {
		return nil, err
	}
	if err = c.pubKey.Verify(crypto.SHA1, []byte(form.Get("data")), sign); err != nil {
		return nil, err
	}

	ret := new(Form)
	if err := json.Unmarshal([]byte(form.Get("data")), ret); err != nil {
		return nil, err
	}
	return ret, nil
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

// NewClient 生成杉德支付客户端
func NewClient(mchID string, options ...Option) *Client {
	c := &Client{
		mchID:  mchID,
		client: internal.NewClient(),
	}
	for _, f := range options {
		f(c)
	}
	return c
}
