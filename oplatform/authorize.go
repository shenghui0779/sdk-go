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

// 构建移动端授权链接 https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Before_Develop/Authorization_Process_Technical_Description.html
type SafeBindComponent struct {
	ComponentAppid string `json:"component_appid"`
	PreAuthCode    string `json:"pre_auth_code"`
	RedirectUri    string `json:"redirect_uri"`
	AuthType       int    `json:"auth_type"`
	BizAppId       string `json:"biz_app_id"`
}

// 使用授权码获取授权信息
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/authorization_info.html#%E8%AF%B7%E6%B1%82%E5%9C%B0%E5%9D%80
type ComponentApiQueryAuth struct {
	ComponentAccessToken string             `json:"component_access_token"`
	ComponentAppid       string             `json:"component_appid"`
	AuthorizationCode    string             `json:"authorization_code"`
	AuthorizationInfo    *AuthorizationInfo `json:"authorization_info"`
}

// TODO 授权之后的用户 信息 取消 func_info 暂时没时间补充
type AuthorizationInfo struct {
	AuthorizerAppid        string `json:"authorizer_appid"`
	AuthorizerAccessToken  string `json:"authorizer_access_token"`
	ExpiresIn              int64    `json:"expires_in"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
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
func GetPreAuthCode(data *PreAuthCode) wx.Action {
	return wx.NewPostAction(fmt.Sprintf(ComponentApiCreatePreAuthCode, data.ComponentAccessToken),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(data)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, data)
		}),
	)
}

// 授权码获取授权信息
func GetComponentApiQueryAuth(data *ComponentApiQueryAuth) wx.Action {
	return wx.NewPostAction(fmt.Sprintf(ComponentApiQueryAuthUrl, data.ComponentAccessToken),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(data)
		}),
		wx.WithDecode(func(resp []byte) error {
			data.AuthorizationInfo.AuthorizerAppid =  gjson.GetBytes(resp, "authorization_info.authorizer_appid").String()
			data.AuthorizationInfo.AuthorizerAccessToken =  gjson.GetBytes(resp, "authorization_info.authorizer_access_token").String()
			data.AuthorizationInfo.ExpiresIn =  gjson.GetBytes(resp, "authorization_info.authorizer_access_token").Int()
			data.AuthorizationInfo.AuthorizerRefreshToken =  gjson.GetBytes(resp, "authorization_info.authorizer_refresh_token").String()
			return nil
		}),
	)
}
