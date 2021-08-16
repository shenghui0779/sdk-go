/*
@Time : 2021/8/16 4:56 下午
@Author : 21
@File : oplatform
@Software: GoLand
*/
package oplatform

import (
	"github.com/shenghui0779/gochat/wx"
)

type Oplatform  struct {
	appid          string
	appsecret      string
	token          string
	encodingAESKey string
	componentVerifyTicket string //
	nonce          func(size uint) string
	client         wx.Client
}


// New returns new wechat mini program
func New(appid, appsecret string) *Oplatform {
	return &Oplatform{
		appid:     appid,
		appsecret: appsecret,
		nonce:     wx.Nonce,
		client:    wx.NewClient(wx.WithInsecureSkipVerify()),
	}
}


// SetServerConfig 设置服务器配置
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/component_verify_ticket.html
func (o *Oplatform) SetServerConfig(token, encodingAESKey ,componentVerifyTicket  string) {
	o.token = token
	o.encodingAESKey = encodingAESKey
	o.componentVerifyTicket = componentVerifyTicket
}

// AppID returns appid
func (o *Oplatform) AppID() string {
	return o.appid
}

// AppSecret returns app secret
func (o *Oplatform) AppSecret() string {
	return o.appsecret
}

// ComponentVerifyTicket returns app componentVerifyTicket
func (o *Oplatform)  ComponentVerifyTicket () string {
	return o.componentVerifyTicket
}




