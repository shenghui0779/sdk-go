package pub

import (
	"encoding/json"
	"errors"
	"fmt"

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
	*WXPub
}

// Code2Token 获取公众号授权AccessToken
func (s *Sns) Code2Token(code string) (*AuthToken, error) {
	resp, err := s.Client.Get(fmt.Sprintf("%s?appid=%s&secret=%s&code=%s&grant_type=authorization_code", SnsCode2Token, s.AppID, s.AppSecret, code))

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
	url := fmt.Sprintf("%s=%s&openid=%s", SnsCheckAccessTokenURL, accessToken, openid)

	resp, err := s.Client.Get(url)

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
	resp, err := s.Client.Get(fmt.Sprintf("%s?appid=%s&grant_type=refresh_token&refresh_token=%s", SnsRefreshAccessTokenURL, s.AppID, refreshToken))

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
	resp, err := s.Client.Get(fmt.Sprintf("%s?access_token=%s&openid=%s&lang=zh_CN", SnsUserInfoURL, accessToken, openid))

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
