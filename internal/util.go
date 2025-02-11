package internal

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"

	"github.com/tidwall/gjson"
)

var Fail = func(err error) (gjson.Result, error) { return gjson.Result{}, err }

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
