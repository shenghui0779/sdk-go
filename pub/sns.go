package pub

import (
	"encoding/json"
	"fmt"

	"github.com/shenghui0779/gochat/utils"
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
	pub     *WXPub
	options []utils.RequestOption
}

// Code2Token 获取公众号授权AccessToken
func (s *Sns) Code2Token(code string) (*AuthToken, error) {
	b, err := s.pub.get(fmt.Sprintf("%s?appid=%s&secret=%s&code=%s&grant_type=authorization_code", SnsCode2Token, s.pub.appid, s.pub.appsecret, code), s.options...)

	if err != nil {
		return nil, err
	}

	resp := new(AuthToken)

	if err := json.Unmarshal(b, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// CheckAccessToken 校验授权AccessToken是否有效
func (s *Sns) CheckAccessToken(accessToken, openid string) error {
	_, err := s.pub.get(fmt.Sprintf("%s=%s&openid=%s", SnsCheckAccessTokenURL, accessToken, openid), s.options...)

	return err
}

// RefreshAccessToken 刷新授权AccessToken
func (s *Sns) RefreshAccessToken(refreshToken string) (*AuthToken, error) {
	b, err := s.pub.get(fmt.Sprintf("%s?appid=%s&grant_type=refresh_token&refresh_token=%s", SnsRefreshAccessTokenURL, s.pub.appid, refreshToken), s.options...)

	if err != nil {
		return nil, err
	}

	resp := new(AuthToken)

	if err := json.Unmarshal(b, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// GetUserInfo 获取微信用户信息
func (s *Sns) GetUserInfo(accessToken, openid string) (*User, error) {
	b, err := s.pub.get(fmt.Sprintf("%s?access_token=%s&openid=%s&lang=zh_CN", SnsUserInfoURL, accessToken, openid), s.options...)

	if err != nil {
		return nil, err
	}

	user := new(User)

	if err := json.Unmarshal(b, user); err != nil {
		return nil, err
	}

	return user, nil
}
