package alipay

import (
	"crypto"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/yiigo/sdk-go/internal/value"
)

type Action struct {
	method  string
	params  V
	bizData X
	encrypt bool
}

// RespKey 返回「method」对应的「xxx_response」
func (a *Action) RespKey() string {
	return strings.ReplaceAll(a.method, ".", "_") + "_response"
}

// Encode 签名并生成请求Body
func (a *Action) Encode(c *Client) (string, error) {
	if c.prvKey == nil {
		return "", errors.New("private key is nil (forgotten configure?)")
	}

	v := make(V)

	v.Set("app_id", c.appid)
	v.Set("method", a.method)
	v.Set("format", "JSON")
	v.Set("charset", "utf-8")
	v.Set("sign_type", "RSA2")
	v.Set("timestamp", time.Now().In(time.Local).Format("2006-01-02 15:04:05"))
	v.Set("version", "1.0")

	for key, val := range a.params {
		v.Set(key, val)
	}

	if len(a.bizData) != 0 {
		bizByte, err := json.Marshal(a.bizData)
		if err != nil {
			return "", err
		}

		bizContent := string(bizByte)

		if a.encrypt {
			bizContent, err = c.Encrypt(bizContent)
			if err != nil {
				return "", err
			}
		}

		v.Set("biz_content", bizContent)
	}

	sign, err := c.prvKey.Sign(crypto.SHA256, []byte(v.Encode("=", "&", value.WithEmptyMode(value.EmptyIgnore))))
	if err != nil {
		return "", err
	}

	v.Set("sign", base64.StdEncoding.EncodeToString(sign))

	return v.Encode("=", "&", value.WithEmptyMode(value.EmptyIgnore), value.WithKVEscape()), nil
}

// ActionOption Action选项
type ActionOption func(a *Action)

// WithReturnURL 设置支付成功跳转URL(HTTP/HTTPS开头字符串)
func WithReturnURL(url string) ActionOption {
	return func(a *Action) {
		a.params.Set("return_url", url)
	}
}

// WithNotifyURL 设置异步回调通知URL
func WithNotifyURL(url string) ActionOption {
	return func(a *Action) {
		a.params.Set("notify_url", url)
	}
}

// WithAuthToken 用户授权令牌
func WithAuthToken(token string) ActionOption {
	return func(a *Action) {
		a.params.Set("auth_token", token)
	}
}

// WithOAuthCode 设置授权码(用授权码来换取授权令牌)
func WithOAuthCode(code string) ActionOption {
	return func(a *Action) {
		a.params.Set("grant_type", string(OAuthCode))
		a.params.Set("code", code)
	}
}

// WithRefreshToken 设置刷新令牌(用刷新令牌来换取一个新的授权令牌)
func WithRefreshToken(token string) ActionOption {
	return func(a *Action) {
		a.params.Set("grant_type", string(RefreshToken))
		a.params.Set("refresh_token", token)
	}
}

// WithAppAuthToken 设置第三方应用授权Token
func WithAppAuthToken(token string) ActionOption {
	return func(a *Action) {
		a.params.Set("app_auth_token", token)
	}
}

// WithScene 设置业务场景描述
func WithScene(scene string) ActionOption {
	return func(a *Action) {
		a.params.Set("scene", scene)
	}
}

// WithKVParam 设置其它非「biz_content」参数
func WithKVParam(k, v string) ActionOption {
	return func(a *Action) {
		a.params.Set(k, v)
	}
}

// WithBizContent 设置「biz_content」参数
func WithBizContent(data X) ActionOption {
	return func(a *Action) {
		a.bizData = data
	}
}

// WithEncrypt 设置请求加密
func WithEncrypt() ActionOption {
	return func(a *Action) {
		a.encrypt = true
		a.params.Set("encrypt_type", "AES")
	}
}

// NewAction 生成Action
func NewAction(method string, options ...ActionOption) *Action {
	action := &Action{
		method: method,
		params: make(V),
	}

	for _, f := range options {
		f(action)
	}

	return action
}

// --------------------------- V3 Option ---------------------------

type V3HeaderOption func(h http.Header)

func WithV3AppAuthToken(token string) V3HeaderOption {
	return func(h http.Header) {
		h.Set(HeaderAppAuthToken, token)
	}
}

func WithV3RootCertSN(sn string) V3HeaderOption {
	return func(h http.Header) {
		h.Set(HeaderRootCertSN, sn)
	}
}

func WithV3MethodOverride(method string) V3HeaderOption {
	return func(h http.Header) {
		h.Set(HeaderMethodOverride, method)
	}
}
