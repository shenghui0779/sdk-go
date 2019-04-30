package mch

import (
	"fmt"
	"sort"
	"strings"

	"github.com/iiinsomnia/gochat/utils"
)

// Sign 生成签名
func Sign(m utils.WXML, apikey string) string {
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

	signature := utils.MD5(strings.Join(kvs, "&"))

	return strings.ToUpper(signature)
}
