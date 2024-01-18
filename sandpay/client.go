package sandpay

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

	"github.com/shenghui0779/sdk-go/lib"
	libCrypto "github.com/shenghui0779/sdk-go/lib/crypto"
	libHttp "github.com/shenghui0779/sdk-go/lib/http"
)

// Client 杉德支付客户端
type Client struct {
	mchID   string
	prvKey  *libCrypto.PrivateKey
	pubKey  *libCrypto.PublicKey
	httpCli libHttp.HTTPClient
	logger  func(ctx context.Context, data map[string]string)
}

// MchID 返回商品ID
func (c *Client) MchID() string {
	return c.mchID
}

// Do 请求杉德API
func (c *Client) Do(ctx context.Context, reqURL string, form *Form) (*Form, error) {
	log := lib.NewReqLog(http.MethodPost, reqURL)
	defer log.Do(ctx, c.logger)

	body, err := form.URLEncode(c.mchID, c.prvKey)
	if err != nil {
		return nil, err
	}

	log.SetReqBody(body)

	resp, err := c.httpCli.Do(ctx, http.MethodPost, reqURL, []byte(body), libHttp.WithHeader(libHttp.HeaderContentType, libHttp.ContentForm))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	log.SetRespHeader(resp.Header)
	log.SetStatusCode(resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", resp.StatusCode)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.SetRespBody(string(b))

	query, err := url.QueryUnescape(string(b))
	if err != nil {
		return nil, err
	}

	v, err := url.ParseQuery(query)
	if err != nil {
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

// NewClient 生成杉德支付客户端
func NewClient(mchID string, options ...Option) *Client {
	c := &Client{
		mchID:   mchID,
		httpCli: libHttp.NewDefaultClient(),
	}

	for _, f := range options {
		f(c)
	}

	return c
}
