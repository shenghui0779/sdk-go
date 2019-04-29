package pub

import "github.com/iiinsomnia/gochat/utils"

// WXPub 微信公众号
type WXPub struct {
	AccountID      string
	AppID          string
	AppSecret      string
	SignToken      string
	EncodingAESKey string
	client         *utils.HTTPClient
}

// SetHTTPClient set wxpub http client
func (wx *WXPub) SetHTTPClient(c *utils.HTTPClient) {
	wx.client = c
}

// Sns returns new sns
func (wx *WXPub) Sns() *Sns {
	sns := new(Sns)

	sns.appid = wx.AppID
	sns.appsecret = wx.AppSecret
	sns.client = wx.client

	return sns
}

// CgiBin returns new cgi-bin
func (wx *WXPub) CgiBin() *CgiBin {
	cgibin := new(CgiBin)

	cgibin.appid = wx.AppID
	cgibin.appsecret = wx.AppSecret
	cgibin.client = wx.client

	return cgibin
}

// MsgChiper returns new msg chiper
func (wx *WXPub) MsgChiper() *MsgChiper {
	chiper := new(MsgChiper)

	chiper.appid = wx.AppID
	chiper.encodingAESKey = wx.EncodingAESKey

	return chiper
}

// Menu returns new menu
func (wx *WXPub) Menu() *Menu {
	menu := new(Menu)

	menu.client = wx.client

	return menu
}

// Subscriber returns new subscriber
func (wx *WXPub) Subscriber() *Subscriber {
	subscriber := new(Subscriber)

	subscriber.client = wx.client

	return subscriber
}

// TplMsg returns new tpl msg
func (wx *WXPub) TplMsg() *TplMsg {
	msg := new(TplMsg)

	msg.client = wx.client

	return msg
}

// Reply returns new reply
func (wx *WXPub) Reply() *Reply {
	reply := new(Reply)

	reply.accountID = wx.AccountID
	reply.appid = wx.AppID
	reply.signToken = wx.SignToken
	reply.encodingAESKey = wx.EncodingAESKey

	return reply
}
