package gochat

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/iiinsomnia/yiigo"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
)

type SnsReply struct {
	SessionKey   string `json:"session_key"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
	OpenID       string `json:"openid"`
	UnionID      string `json:"unionid"`
	Scope        string `json:"scope"`
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
}

// Sns ...
type Sns struct {
	code      string
	appID     string
	appSecret string
	reply     *SnsReply
}

// Code2Session 获取小程序授权SessionKey
func (s *Sns) Code2Session() error {
	resp, err := yiigo.HTTPGet(fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", s.appID, s.appSecret, s.code))

	if err != nil {
		yiigo.Logger.Error("wx sns code2session error", zap.String("error", err.Error()), zap.ByteString("resp", resp))

		return err
	}

	yiigo.Logger.Info("wx sns code2session resp", zap.ByteString("resp", resp))

	reply := new(SnsReply)

	if err := json.Unmarshal(resp, reply); err != nil {
		yiigo.Logger.Error("unmarshal wx sns code2session resp error", zap.String("error", err.Error()))

		return err
	}

	if reply.ErrCode != 0 {
		yiigo.Logger.Error("wx sns code2session error", zap.Int("code", reply.ErrCode), zap.String("error", reply.ErrMsg))

		return errors.New(reply.ErrMsg)
	}

	s.reply = reply

	return nil
}

// Code2Token 获取公众号授权AccessToken
func (s *Sns) Code2Token() error {
	resp, err := yiigo.HTTPGet(fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", s.appID, s.appSecret, s.code))

	if err != nil {
		yiigo.Logger.Error("wx sns code2token error", zap.String("error", err.Error()))

		return err
	}

	yiigo.Logger.Info("wx sns code2token resp", zap.ByteString("resp", resp))

	reply := new(SnsReply)

	if err := json.Unmarshal(resp, reply); err != nil {
		yiigo.Logger.Error("unmarshal wx sns code2token resp error", zap.String("error", err.Error()))

		return err
	}

	if reply.ErrCode != 0 {
		yiigo.Logger.Error("wx sns code2token error", zap.Int("code", reply.ErrCode), zap.String("error", reply.ErrMsg))

		return errors.New(reply.ErrMsg)
	}

	s.reply = reply

	return nil
}

// SessionKey ...
func (s *Sns) SessionKey() string {
	return s.reply.SessionKey
}

// AccessToken ...
func (s *Sns) AccessToken() string {
	return s.reply.AccessToken
}

// RefreshToken ...
func (s *Sns) RefreshToken() string {
	return s.reply.RefreshToken
}

// ExpiresIn ...
func (s *Sns) ExpiresIn() int64 {
	return s.reply.ExpiresIn
}

// OpenID ...
func (s *Sns) OpenID() string {
	return s.reply.OpenID
}

// UnionID ...
func (s *Sns) UnionID() string {
	return s.reply.UnionID
}

// NewSns ...
func NewSns(code, appID, appSecret string) *Sns {
	return &Sns{
		code:      code,
		appID:     appID,
		appSecret: appSecret,
	}
}

// NewSnsWithChannel ...
func NewSnsWithChannel(code, channel string) *Sns {
	setting := GetSettingsWithChannel(channel)

	return &Sns{
		code:      code,
		appID:     setting.AppID,
		appSecret: setting.AppSecret,
	}
}

// CheckSnsAccessToken 校验授权AccessToken是否有效
func CheckSnsAccessToken(accessToken, openid string) bool {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/auth?access_token=%s&openid=%s", accessToken, openid)

	resp, err := yiigo.HTTPGet(url)

	if err != nil {
		yiigo.Logger.Error("wx check sns access_token error", zap.String("error", err.Error()))

		return false
	}

	yiigo.Logger.Info("wx check sns access_token resp", zap.ByteString("resp", resp))

	if gjson.GetBytes(resp, "errcode").Int() != 0 {
		return false
	}

	return true
}

// RefreshSnsToken ...
type RefreshSnsToken struct {
	refreshToken string
	channel      string
	reply        *SnsReply
}

// RefreshSnsAccessToken 刷新授权AccessToken
func (r *RefreshSnsToken) Do(refreshToken, channel WXChannel) error {
	cfg := GetConfigWithChannel(channel)

	resp, err := yiigo.HTTPGet(fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s", setting.AppID, refreshToken))

	if err != nil {
		yiigo.Logger.Error("wx refresh sns access_token error", zap.String("error", err.Error()))

		return err
	}

	yiigo.Logger.Info("wx refresh sns access_token resp", zap.ByteString("resp", resp))

	reply := new(SnsReply)

	if err := json.Unmarshal(resp, reply); err != nil {
		yiigo.Logger.Error("unmarshal wx refresh sns access_token resp error", zap.String("error", err.Error()))

		return err
	}

	if reply.ErrCode != 0 {
		yiigo.Logger.Error("wx refresh sns access_token error", zap.Int("code", reply.ErrCode), zap.String("error", reply.ErrMsg))

		return errors.New(reply.ErrMsg)
	}

	r.reply = reply

	return nil
}

// AccessToken ...
func (r *RefreshSnsToken) AccessToken() string {
	return r.reply.AccessToken
}

// RefreshToken ...
func (r *RefreshSnsToken) RefreshToken() string {
	return r.reply.RefreshToken
}

// ExpiresIn ...
func (r *RefreshSnsToken) ExpiresIn() int64 {
	return r.reply.ExpiresIn
}

// NewRefreshSnsToken ...
func NewRefreshSnsToken(refreshToken, channel string) *RefreshSnsToken {
	return &RefreshSnsToken{
		refreshToken: refreshToken,
		channel:      channel,
	}
}

// SnsUserReply 微信用户信息
type SnsUserReply struct {
	OpenID    string   `json:"openid"`
	UnionID   string   `json:"unionid"`
	Nickname  string   `json:"nickname"`
	Gender    int      `json:"sex"`
	Province  string   `json:"province"`
	City      string   `json:"city"`
	Country   string   `json:"country"`
	Avatar    string   `json:"headimgurl"`
	Privilege []string `json:"privilege"`
	ErrCode   int      `json:"errcode"`
	ErrMsg    string   `json:"errmsg"`
}

// SnsUser ...
type SnsUser struct {
	accessToken string
	openid      string
	reply       *SnsUserReply
}

// GetUserInfo 获取微信用户信息
func (s *SnsUser) GetUserInfo() error {
	resp, err := yiigo.HTTPGet(fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN", s.accessToken, s.openid))

	if err != nil {
		yiigo.Logger.Error("wx sns userinfo error", zap.String("error", err.Error()))

		return err
	}

	yiigo.Logger.Info("wx sns userinfo resp", zap.ByteString("resp", resp))

	reply := new(SnsUserReply)

	if err := json.Unmarshal(resp, reply); err != nil {
		yiigo.Logger.Error("unmarshal wx sns userinfo resp error", zap.String("error", err.Error()), zap.ByteString("resp", resp))

		return err
	}

	if reply.ErrCode != 0 {
		yiigo.Logger.Error("wx sns userinfo error", zap.Int("code", reply.ErrCode), zap.String("error", reply.ErrMsg))

		return errors.New(reply.ErrMsg)
	}

	s.reply = reply

	return nil
}

// OpenID get openid
func (s *SnsUser) OpenID() string {
	return s.reply.OpenID
}

// UnionID get unionid
func (s *SnsUser) UnionID() string {
	return s.reply.UnionID
}

// Nickname get nickname
func (s *SnsUser) Nickname() string {
	return s.reply.Nickname
}

// Avatar get avatar
func (s *SnsUser) Avatar() string {
	return s.reply.Avatar
}

func NewSnsUser(accessToken, openid string) *SnsUser {
	return &SnsUser{
		accessToken: accessToken,
		openid:      openid,
	}
}
