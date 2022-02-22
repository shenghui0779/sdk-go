package corp

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

// AuthScope 应用授权作用域
type AuthScope string

// 企业微信支持的应用授权作用域
const (
	ScopeSnsapiBase AuthScope = "snsapi_base" // 企业自建应用固定填写：snsapi_base
)

// AccessToken 企业微信AccessToken
type AccessToken struct {
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

type ResultIP struct {
	IPList []string `json:"ip_list"`
}

func GetAPIDomainIP(result *ResultIP) wx.Action {
	return wx.NewGetAction(urls.CorpCgiBinAPIDomainIP,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func GetCallbackIP(result *ResultIP) wx.Action {
	return wx.NewGetAction(urls.CorpCgiBinCallbackIP,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ResultOAuthUser struct {
	UserID         string `json:"UserId"`
	OpenID         string `json:"OpenId"`
	DeviceID       string `json:"DeviceId"`
	ExternalUserID string `json:"external_userid"`
}

// GetOAuthUser 获取访问用户身份
func GetOAuthUser(code string, result *ResultOAuthUser) wx.Action {
	return wx.NewGetAction(urls.CorpCgiBinUserInfo,
		wx.WithQuery("code", code),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// UserAuthSucc 二次验证
func UserAuthSucc(userID string) wx.Action {
	return wx.NewGetAction(urls.CorpCgiBinUserAuthSucc,
		wx.WithQuery("userid", userID),
	)
}
