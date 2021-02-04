package oa

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/wx"
)

// AuthToken 公众号网页授权Token
type AuthToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`
}

// AccessToken 公众号普通AccessToken
type AccessToken struct {
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

// CheckAuthToken 检验授权凭证（access_token）是否有效
func CheckAuthToken(openid string) wx.Action {
	return wx.NewAction(SnsCheckAccessTokenURL,
		wx.WithMethod(wx.MethodGet),
		wx.WithQuery("openid", openid),
	)
}

// AuthUser 授权用户信息
type AuthUser struct {
	OpenID     string   `json:"openid"`
	UnionID    string   `json:"unionid"`
	Nickname   string   `json:"nickname"`
	Sex        string   `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgURL string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
}

// GetAuthUser 获取授权用户信息（注意：使用网页授权的access_token）
func GetAuthUser(dest *AuthUser, openid string) wx.Action {
	return wx.NewAction(SnsUserInfoURL,
		wx.WithMethod(wx.MethodGet),
		wx.WithQuery("openid", openid),
		wx.WithQuery("lang", "zh_CN"),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// JS-SDK ticket 类型
type TicketType string

// 微信支持的 JS-SDK ticket
const (
	APITicket   TicketType = "wx_card"
	JSAPITicket TicketType = "jsapi"
)

// JSSDKSign JS-SDK签名
type JSSDKSign struct {
	Signature string `json:"signature"`
	Noncestr  string `json:"noncestr"`
	Timestamp int64  `json:"timestamp"`
}

// JSSDKTicket 公众号 JS-SDK ticket
type JSSDKTicket struct {
	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
}

// GetJSSDKTicket 获取 JS-SDK ticket (注意：使用普通access_token)
func GetJSSDKTicket(dest *JSSDKTicket, t TicketType) wx.Action {
	return wx.NewAction(CgiBinTicketURL,
		wx.WithMethod(wx.MethodGet),
		wx.WithQuery("type", string(t)),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}
