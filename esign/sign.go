package esign

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/url"
	"strings"

	"github.com/shenghui0779/sdk-go/lib/value"
)

type signfields struct {
	accept   string
	contMD5  string
	contType string
	date     string
	headers  value.V
	params   value.V
}

// SignOption 签名选项
type SignOption func(sf *signfields)

// WithSignAccept 设置Accept
func WithSignAccept(v string) SignOption {
	return func(sf *signfields) {
		sf.accept = v
	}
}

// WithSignContMD5 设置ContentMD5
func WithSignContMD5(v string) SignOption {
	return func(sf *signfields) {
		sf.contMD5 = v
	}
}

// WithSignContType 设置ContentType
func WithSignContType(v string) SignOption {
	return func(sf *signfields) {
		sf.contType = v
	}
}

// WithSignDate 设置Date
func WithSignDate(v string) SignOption {
	return func(sf *signfields) {
		sf.date = v
	}
}

// WithSignHeader 设置Header
func WithSignHeader(k, v string) SignOption {
	return func(sf *signfields) {
		sf.headers.Set(k, v)
	}
}

// WithSignParam 以 K-V 设置参数
func WithSignParam(k, v string) SignOption {
	return func(sf *signfields) {
		sf.params.Set(k, v)
	}
}

// WithSignValues 以 url.Values 设置参数
func WithSignValues(v url.Values) SignOption {
	return func(sf *signfields) {
		for key, vals := range v {
			if len(vals) != 0 {
				sf.params.Set(key, vals[0])
			} else {
				sf.params.Set(key, "")
			}
		}
	}
}

// Signer 签名器
type Signer struct {
	str string
}

// Do 生成签名
func (s *Signer) Do(secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(s.str))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// String 返回签名字符串
func (s *Signer) String() string {
	return s.str
}

// NewSigner 返回新的签名器
func NewSigner(method, path string, options ...SignOption) *Signer {
	fields := &signfields{
		accept:  "*/*",
		headers: make(value.V),
		params:  make(value.V),
	}

	for _, f := range options {
		f(fields)
	}

	var buf strings.Builder

	buf.WriteString(method)
	buf.WriteString("\n")
	buf.WriteString(fields.accept)
	buf.WriteString("\n")
	buf.WriteString(fields.contMD5)
	buf.WriteString("\n")
	buf.WriteString(fields.contType)
	buf.WriteString("\n")
	buf.WriteString(fields.date)
	buf.WriteString("\n")

	if len(fields.headers) != 0 {
		buf.WriteString(fields.headers.Encode(":", "\n"))
		buf.WriteString("\n")
	}

	buf.WriteString(path)

	if len(fields.params) != 0 {
		buf.WriteString("?")
		buf.WriteString(fields.params.Encode("=", "&", value.WithEmptyMode(value.EmptyOnlyKey)))
	}

	return &Signer{str: buf.String()}
}
