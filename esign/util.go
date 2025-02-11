package esign

import (
	"crypto/md5"
	"encoding/base64"
	"io"
	"os"

	"github.com/yiigo/sdk-go/internal/value"
)

type V = value.V

const (
	HeaderContentMD5           = "Content-MD5"
	HeaderTSignOpenAppID       = "X-Tsign-Open-App-Id"
	HeaderTSignOpenAuthMode    = "X-Tsign-Open-Auth-Mode"
	HeaderTSignOpenCaTimestamp = "X-Tsign-Open-Ca-Timestamp"
	HeaderTSignOpenCaSignature = "X-Tsign-Open-Ca-Signature"
	HeaderTSignOpenTimestamp   = "X-Tsign-Open-TIMESTAMP"
	HeaderTSignOpenSignature   = "X-Tsign-Open-SIGNATURE"
)

const (
	AcceptAll    = "*/*"
	AuthModeSign = "Signature"
)

// ContentMD5 计算内容MD5值
func ContentMD5(b []byte) string {
	h := md5.New()
	h.Write(b)

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// FileMD5 计算文件MD5值
func FileMD5(filename string) (string, int64) {
	f, err := os.Open(filename)
	if err != nil {
		return err.Error(), -1
	}
	defer f.Close()

	h := md5.New()
	n, err := io.Copy(h, f)
	if err != nil {
		return err.Error(), -1
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), n
}
