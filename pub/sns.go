package pub

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/iiinsomnia/gochat/utils"
	"github.com/tidwall/gjson"
)

// AuthToken 公众号授权Token
type AuthToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`
}

// User 微信用户信息
type User struct {
	OpenID    string   `json:"openid"`
	UnionID   string   `json:"unionid"`
	Nickname  string   `json:"nickname"`
	Gender    int      `json:"sex"`
	Province  string   `json:"province"`
	City      string   `json:"city"`
	Country   string   `json:"country"`
	Avatar    string   `json:"headimgurl"`
	Privilege []string `json:"privilege"`
}

// Sns sns
type Sns struct {
	appid     string
	appsecret string
}

// Code2Token 获取公众号授权AccessToken
func (s *Sns) Code2Token(code string) (*AuthToken, error) {
	resp, err := utils.HTTPGet(fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", s.appid, s.appsecret, code))

	if err != nil {
		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return nil, errors.New(r.Get("errmsg").String())
	}

	reply := new(AuthToken)

	if err := json.Unmarshal(resp, reply); err != nil {
		return nil, err
	}

	return reply, nil
}

// CheckAccessToken 校验授权AccessToken是否有效
func (s *Sns) CheckAccessToken(accessToken, openid string) bool {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/auth?access_token=%s&openid=%s", accessToken, openid)

	resp, err := utils.HTTPGet(url)

	if err != nil {
		return false
	}

	if gjson.GetBytes(resp, "errcode").Int() != 0 {
		return false
	}

	return true
}

// RefreshAccessToken 刷新授权AccessToken
func (s *Sns) RefreshAccessToken(refreshToken string) (*AuthToken, error) {
	resp, err := utils.HTTPGet(fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s", s.appid, refreshToken))

	if err != nil {
		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return nil, errors.New(r.Get("errmsg").String())
	}

	reply := new(AuthToken)

	if err := json.Unmarshal(resp, reply); err != nil {
		return nil, err
	}

	return reply, nil
}

// GetUserInfo 获取微信用户信息
func (s *Sns) GetUserInfo(accessToken, openid string) (*User, error) {
	resp, err := utils.HTTPGet(fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN", accessToken, openid))

	if err != nil {
		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return nil, errors.New(r.Get("errmsg").String())
	}

	user := new(User)

	if err := json.Unmarshal(resp, user); err != nil {
		return nil, err
	}

	return user, nil
}
