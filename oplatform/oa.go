/*
@Time : 2021/8/17 5:34 下午
@Author : 21
@File : oa
@Software: GoLand
*/
package oplatform

import (
	"encoding/json"
	"fmt"
	"github.com/shenghui0779/gochat/wx"
)

// 关联小程序
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Official__Accounts/Mini_Program_Management_Permission.html
type WxopenWxamplink struct {
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
	Appid                  string `json:"appid"`
	NotifyUsers            string `json:"notify_users"`
	ShowProfile            string `json:"show_profile"`
}

func SetWxopenWxamplink(data *WxopenWxamplink) wx.Action {
	return wx.NewPostAction(fmt.Sprintf(WxopenWxamplinkUrl, data.AuthorizerRefreshToken),
		wx.WithBody(func() (bytes []byte, e error) {
			return json.Marshal(data)
		}),
		wx.WithDecode(func(resp []byte) error {

			return nil
		},
		))
}
