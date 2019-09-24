package mch

import (
	"fmt"
	"sort"
	"strings"

	"github.com/iiinsomnia/gochat/utils"
)

// SignWithMD5 生成MD5签名
func SignWithMD5(m utils.WXML, apikey string) string {
	signature := utils.MD5(buildSignStr(m, apikey))

	return strings.ToUpper(signature)
}

// SignWithHMacSHA256 生成HMAC-SHA256签名
func SignWithHMacSHA256(m utils.WXML, apikey string) string {
	signature := utils.HMAC("sha256", buildSignStr(m, apikey), apikey)

	return strings.ToUpper(signature)
}

// Sign 生成签名
func buildSignStr(m utils.WXML, apikey string) string {
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

	kvs = append(kvs, fmt.Sprintf("key=%s", apikey))

	return strings.Join(kvs, "&")
}
