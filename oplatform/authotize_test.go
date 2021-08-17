/*
@Time : 2021/8/17 11:18 上午
@Author : 21
@File : authotize_test
@Software: GoLand
*/
package oplatform

import (
	"context"
	"fmt"
	"github.com/tidwall/gjson"
	"testing"
)

func TestGetPreAuthCode(t *testing.T) {
	//var EventMsg = "QqzXZAu+dmwtinjQ6ilJd57JpD1XI7gbFU4IfnHc9vZj0l84ZjjGtIB9lZokMkm3xeGcVAoYPTOfNuz10Z6yaKJqzdLo5IFd7G72Jd3bAJladFdd2ZVh8RHIyFRsZ3Np1uIT6rhy89cypSo0txNLAQOJtBsYDG+WnSkD4IhQjM8CgmeF7K5ORWb66dRTFqaFfEbV157DbpJOhgqlLc+OrkqtjAVz2W+IMzHwJ8jvfka2+huvEWpudQ6TroXxArEPIWustZVDoTxkKVT+dJDvjovFym0wO/f4ludEtkcw8So1f9l4SYYle/SkItioLdlvR4kGxlpySTektweVLNKhYQHrGZATyTNH2TxJpRvsBNwdO0OkNddngDW08xAPhPc+3BORwvQZE3VRGSdAOpzYAniSCL9u8G+mAm8tLyqRtPdgMGjYIQtykTkHzn7OUO7JhsqYm5ez7OtOw0PTLe+TVA=="
	op := New("wxc83d3daa98fe100c","dd8c33e9d4634923f70a77ada891f4f8")
	op.SetServerConfig("womeibanfale","zhinengxiugainimenle00000000000000000000001","123123")
	err := op.Do(context.Background(),GetPreAuthCode(&PreAuthCode{
		ComponentAppid:       op.appid,
		ComponentAccessToken: "111111",
	}))
	fmt.Print(err)
}

func TestGetApiComponentToken(t *testing.T) {
	var  j = `{
  "authorization_info": {
    "authorizer_appid": "wxf8b4f85f3a794e77",
    "authorizer_access_token": "QXjUqNqfYVH0yBE1iI_7vuN_9gQbpjfK7hYwJ3P7xOa88a89-Aga5x1NMYJyB8G2yKt1KCl0nPC3W9GJzw0Zzq_dBxc8pxIGUNi_bFes0qM",
    "expires_in": 7200,
    "authorizer_refresh_token": "dTo-YCXPL4llX-u1W1pPpnp8Hgm4wpJtlR6iV0doKdY",
    "func_info": [
      {
        "funcscope_category": {
          "id": 1
        }
      },
      {
        "funcscope_category": {
          "id": 2
        }
      },
      {
        "funcscope_category": {
          "id": 3
        }
      }
    ]
  }
}`
  jsonStr := gjson.GetBytes([]byte(j),"authorization_info.authorizer_appid")
  fmt.Println(jsonStr)
}
