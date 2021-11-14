package offia

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

// Sex 性别
type Sex int

const (
	SexUnknown Sex = 0 // 未知
	SexMale    Sex = 1 // 男性
	SexFemale  Sex = 2 // 女性
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
	return wx.NewGetAction(urls.OffiaSnsCheckAccessToken,
		wx.WithQuery("openid", openid),
	)
}

// ResultAuthInfo 授权用户信息
type ResultAuthInfo struct {
	OpenID     string   `json:"openid"`
	UnionID    string   `json:"unionid"`
	Nickname   string   `json:"nickname"`
	Sex        Sex      `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgURL string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
}

// GetAuthInfo 获取授权用户信息（注意：使用网页授权的access_token）
func GetAuthInfo(openid string, result *ResultAuthInfo) wx.Action {
	return wx.NewGetAction(urls.OffiaSnsUserInfo,
		wx.WithQuery("openid", openid),
		wx.WithQuery("lang", "zh_CN"),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
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

// ResultJSSDKTicket 公众号 JS-SDK ticket
type ResultJSSDKTicket struct {
	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
}

// GetJSSDKTicket 获取 JS-SDK ticket (注意：使用普通access_token)
func GetJSSDKTicket(t TicketType, result *ResultJSSDKTicket) wx.Action {
	return wx.NewGetAction(urls.OffiaCgiBinTicket,
		wx.WithQuery("type", string(t)),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
