package oa

import (
	"encoding/json"
	"net/url"

	"github.com/shenghui0779/gochat/public"
)

// AuthToken 公众号网页授权Token
type AuthToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`
}

// AuthUser 授权微信用户信息
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

// AccessToken 公众号普通AccessToken
type AccessToken struct {
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

// JSAPITicket 公众号 jsapi ticket
type JSAPITicket struct {
	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
}

// CheckAuthToken 校验网页授权AccessToken是否有效
func CheckAuthToken(openid string) public.Action {
	query := url.Values{}

	query.Set("openid", openid)

	return public.NewOpenGetAPI(SnsCheckAccessTokenURL, query, nil)
}

// GetAuthUser 获取授权微信用户信息（注意：使用网页授权的access_token）
func GetAuthUser(openid string, dest *AuthUser) public.Action {
	query := url.Values{}

	query.Set("openid", openid)
	query.Set("lang", "zh_CN")

	return public.NewOpenGetAPI(SnsUserInfoURL, query, func(resp []byte) error {
		return json.Unmarshal(resp, dest)
	})
}

// GetJSAPITicket 获取 jsapi ticket (注意：使用普通access_token)
func GetJSAPITicket(dest *JSAPITicket) public.Action {
	query := url.Values{}

	query.Set("type", "jsapi")

	return public.NewOpenGetAPI(CgiBinTicketURL, query, func(resp []byte) error {
		return json.Unmarshal(resp, dest)
	})
}
