package gochat

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/iiinsomnia/yiigo"
	"go.uber.org/zap"
)

// CgiBin ...
type CgiBin struct {
	appID     string
	appSecret string
	reply     *CgiBinReply
}

// CgiBinReply ...
type CgiBinReply struct {
	AccessToken string `json:"access_token"`
	Ticket      string `json:"ticket"`
	ExpiresIn   int64  `json:"expires_in"`
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

// GetAccessToken 获取普通AccessToken
func (p *CgiBin) GetAccessToken() error {
	resp, err := yiigo.HTTPGet(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", p.appID, p.appSecret))

	if err != nil {
		yiigo.Logger.Error("get wx cgi-bin access_token error", zap.String("error", err.Error()))

		return err
	}

	yiigo.Logger.Info("get wx cgi-bin access_token resp", zap.ByteString("resp", resp))

	reply := new(CgiBinReply)

	if err := json.Unmarshal(resp, reply); err != nil {
		yiigo.Logger.Error("unmarshal wx cgi-bin access_token resp error", zap.String("error", err.Error()))

		return err
	}

	if reply.ErrCode != 0 {
		yiigo.Logger.Error("get wx cgi-bin access_token error", zap.Int("code", reply.ErrCode), zap.String("error", reply.ErrMsg))

		return errors.New(reply.ErrMsg)
	}

	p.reply = reply

	return nil
}

// GetTicket 获取 JSAPI ticket
func (p *CgiBin) GetTicket(accessToken string) error {
	resp, err := yiigo.HTTPGet(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi", accessToken))

	if err != nil {
		yiigo.Logger.Error("get wx cgi-bin js_ticket error", zap.String("error", err.Error()))

		return err
	}

	yiigo.Logger.Info("get wx cgi-bin js_ticket resp", zap.ByteString("resp", resp))

	reply := new(CgiBinReply)

	if err := json.Unmarshal(resp, reply); err != nil {
		yiigo.Logger.Error("unmarshal wx js_ticket resp error", zap.String("error", err.Error()))

		return err
	}

	if reply.ErrCode != 0 {
		yiigo.Logger.Error("get wx cgi-bin js_ticket error", zap.Int("code", reply.ErrCode), zap.String("error", reply.ErrMsg))

		return errors.New(reply.ErrMsg)
	}

	p.reply = reply

	return nil
}

// AccessToken ...
func (p *CgiBin) AccessToken() string {
	return p.reply.AccessToken
}

// Ticket ...
func (p *CgiBin) Ticket() string {
	return p.reply.Ticket
}

// ExpiresIn ...
func (p *CgiBin) ExpiresIn() int64 {
	return p.reply.ExpiresIn
}

// NewCgiBin ...
func NewCgiBin(appID, appSecret string) *CgiBin {
	return &CgiBin{
		appID:     appID,
		appSecret: appSecret,
	}
}

// NewCgiBinWithChannel ...
func NewCgiBinWithChannel(channel string) *CgiBin {
	setting := GetSettingsWithChannel(channel)

	return &CgiBin{
		appID:     setting.AppID,
		appSecret: setting.AppSecret,
	}
}
