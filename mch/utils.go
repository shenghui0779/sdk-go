package mch

import (
	"fmt"
	"net/url"
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

func buildSignStr(m utils.WXML, apikey string) string {
	query := url.Values{}

	for k, v := range m {
		if k == "sign" || v == "" {
			continue
		}

		query.Add(k, v)
	}

	return fmt.Sprintf("%s&key=%s", query.Encode(), apikey)
}
