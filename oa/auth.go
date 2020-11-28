package oa

import (
	"encoding/json"
	"fmt"

	"github.com/shenghui0779/gochat/internal"
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
func CheckAuthToken(openid string) internal.Action {
	return &WechatAPI{
		url: func(accessToken string) string {
			return fmt.Sprintf("GET|%s?access_token=%s&openid=%s", SnsCheckAccessTokenURL, accessToken, openid)
		},
	}
}

// GetAuthUser 获取授权微信用户信息（注意：使用网页授权的access_token）
func GetAuthUser(openid string, dest *AuthUser) internal.Action {
	return &WechatAPI{
		url: func(accessToken string) string {
			return fmt.Sprintf("GET|%s?access_token=%s&openid=%s&lang=zh_CN", SnsUserInfoURL, accessToken, openid)
		},
		decode: func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		},
	}
}

// GetJSAPITicket 获取 jsapi ticket (注意：使用普通access_token)
func GetJSAPITicket(dest *JSAPITicket) internal.Action {
	return &WechatAPI{
		url: func(accessToken string) string {
			return fmt.Sprintf("GET|%s?access_token=%s&type=jsapi", CgiBinTicketURL, accessToken)
		},
		decode: func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		},
	}
}
