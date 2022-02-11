package offia

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

// OAuthToken 公众号网页授权Token
type OAuthToken struct {
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

// CheckOAuthToken 检验授权凭证（access_token）是否有效
func CheckOAuthToken(openid string) wx.Action {
	return wx.NewGetAction(urls.OffiaSnsCheckAccessToken,
		wx.WithQuery("openid", openid),
	)
}

// ResultOAuthUser 授权用户信息
type ResultOAuthUser struct {
	OpenID     string   `json:"openid"`
	UnionID    string   `json:"unionid"`
	Nickname   string   `json:"nickname"`
	Sex        int      `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgURL string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
}

// GetOAuthUser 获取授权用户信息（注意：使用网页授权的access_token）
func GetOAuthUser(openid string, result *ResultOAuthUser) wx.Action {
	return wx.NewGetAction(urls.OffiaSnsUserInfo,
		wx.WithQuery("openid", openid),
		wx.WithQuery("lang", "zh_CN"),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// TicketType JS-SDK ticket 类型
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

// ResultApiTicket 公众号 api ticket
type ResultApiTicket struct {
	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
}

// GetApiTicket 获取 JS-SDK ticket (注意：使用普通access_token)
func GetApiTicket(ticketType TicketType, result *ResultApiTicket) wx.Action {
	return wx.NewGetAction(urls.OffiaCgiBinTicket,
		wx.WithQuery("type", string(ticketType)),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
