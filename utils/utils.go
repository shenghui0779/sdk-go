package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/iiinsomnia/gochat"
)

// ErrIllegaAppID appid illegal
var ErrIllegaAppID = errors.New("appid is not match")

// NonceStr 随机字符串
func NonceStr() string {
	now := time.Now()

	return strconv.FormatInt(now.Unix()+int64(now.Nanosecond()), 10)
}

// MD5 calculate the md5 hash of a string.
func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil))
}

// SHA1 calculate the sha1 hash of a string.
func SHA1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil))
}

// PaySign 生成签名
func PaySign(m gochat.WXML, apiKey string) string {
	l := len(m)

	ks := make([]string, 0, l)
	kvs := make([]string, 0, l)

	for k := range m {
		if k == "sign" {
			continue
		}

		ks = append(ks, k)
	}

	sort.Strings(ks)

	for _, k := range ks {
		if v, ok := m[k]; ok && v != "" {
			kvs = append(kvs, fmt.Sprintf("%s=%s", k, v))
		}
	}

	kvs = append(kvs, fmt.Sprintf("key=%s", apiKey))

	signature := yiigo.MD5(strings.Join(kvs, "&"))

	return strings.ToUpper(signature)
}

// EncodeUint32ToBytes 把整数 uint32 格式化成 4 字节的网络字节序
func EncodeUint32ToBytes(i uint32) []byte {
	b := make([]byte, 4)

	b[0] = byte(i >> 24)
	b[1] = byte(i >> 16)
	b[2] = byte(i >> 8)
	b[3] = byte(i)

	return b
}

// DecodeBytesToUint32 从 4 字节的网络字节序里解析出整数 uint32
func DecodeBytesToUint32(b []byte) uint32 {
	if len(b) != 4 {
		return 0
	}

	return uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
}

// RandomStr 随机字符串
func RandomStr(n int) string {
	salt := make([]string, 0, n)

	pattern := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz"
	l := len(pattern)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		p := rand.Intn(l)
		salt = append(salt, string(pattern[p]))
	}

	return strings.Join(salt, "")
}
