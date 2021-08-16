/*
@Time : 2021/8/16 5:05 下午
@Author : 21
@File : authorize
@Software: GoLand
*/
package oplatform

import (
	"encoding/json"
	"fmt"
	"github.com/shenghui0779/gochat/wx"
	"github.com/tidwall/gjson"
)

// 获取令牌 https://api.weixin.qq.com/cgi-bin/component/api_component_token
type ComponentAccessToken struct {
	ComponentAppid        string `json:"component_appid"`
	ComponentAppsecret    string `json:"component_appsecret"`
	ComponentVerifyTicket string `json:"component_verify_ticket"`
	// 返回值
	ComponentAccessToken string `json:"component_access_token"`
	ExpiresIn            int    `json:"expires_in"`
}

// 获取预授权码 https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/pre_auth_code.html
type PreAuthCode struct {
	ComponentAppid       string `json:"component_appid"`
	ComponentAccessToken string `json:"component_access_token"`
	PreAuthCode          string `json:"pre_auth_code"`
	ExpiresIn            int    `json:"expires_in"`
}

// 获取令牌
func GetApiComponentToken(data *ComponentAccessToken) wx.Action {
	return wx.NewPostAction(ComponentApiComponentTokenUrl,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(data)
		}),
		wx.WithDecode(func(resp []byte) error {
			data.ComponentAccessToken = gjson.GetBytes(resp, "component_access_token").String()
			data.ExpiresIn = int(gjson.GetBytes(resp, "expires_in").Int())
			return nil
		}),
	)
}

// 获取预授权码
func GetPreAuthCode(data *PreAuthCode) wx.Action  {
	return wx.NewGetAction(fmt.Sprintf(ComponentApiCreatePreAuthCode, data.ComponentAccessToken),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, data)
		}),
	)
}
