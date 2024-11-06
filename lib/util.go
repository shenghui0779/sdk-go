package lib

import (
	"bytes"
	"crypto/rand"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/tidwall/gjson"
)

// X 类型别名
type X map[string]any

var Fail = func(err error) (gjson.Result, error) { return gjson.Result{}, err }

var DefaultHttpClient = &http.Client{
	Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 60 * time.Second,
		}).DialContext,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		MaxIdleConns:          0,
		MaxIdleConnsPerHost:   1000,
		MaxConnsPerHost:       1000,
		IdleConnTimeout:       60 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	},
}

// Nonce 生成指定长度的随机串 (最好是偶数)
func Nonce(size uint) string {
	nonce := make([]byte, size/2)
	_, _ = io.ReadFull(rand.Reader, nonce)
	return hex.EncodeToString(nonce)
}

// NonceByte 生成指定长度的随机字节
func NonceByte(size uint) []byte {
	nonce := make([]byte, size)
	_, _ = io.ReadFull(rand.Reader, nonce)
	return nonce
}

// MarshalNoEscapeHTML 不带HTML转义的JSON序列化
func MarshalNoEscapeHTML(v interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)

	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(v); err != nil {
		return nil, err
	}

	b := buf.Bytes()
	// 去掉 go std 给末尾加的 '\n'
	// @see https://github.com/golang/go/issues/7767
	if l := len(b); l != 0 && b[l-1] == '\n' {
		b = b[:l-1]
	}
	return b, nil
}
