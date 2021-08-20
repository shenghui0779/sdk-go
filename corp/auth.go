package corp

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

// AccessToken 企业微信AccessToken
type AccessToken struct {
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

type APIDomainIP struct {
	IPList []string `json:"ip_list"`
}

func GetAPIDomainIP(dest *APIDomainIP) wx.Action {
	return wx.NewGetAction(urls.CorpAddrBookAPIDomainIP,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

type UserInfo struct {
	UserID         string `json:"UserId"`
	OpenID         string `json:"OpenId"`
	DeviceID       string `json:"DeviceId"`
	ExternalUserID string `json:"external_userid"`
}

// GetUserInfo 获取访问用户身份
func GetUserInfo(dest *UserInfo, code string) wx.Action {
	return wx.NewGetAction(urls.CorpAddrBookUserInfo,
		wx.WithQuery("code", code),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// UserAuthSucc 二次验证
func UserAuthSucc(userID string) wx.Action {
	return wx.NewGetAction(urls.CorpAddrBookUserAuthSucc,
		wx.WithQuery("userid", userID),
	)
}
