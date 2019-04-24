package mch

import "github.com/iiinsomnia/gochat/utils"

// WXMch 微信商户
type WXMch struct {
	AppID     string
	MchID     string
	ApiKey    string
	sslClient *utils.HTTPClient
}

// SetSSLClient set wxmch ssl client
func (wx *WXMch) SetSSLClient(c *utils.HTTPClient) {
	wx.sslClient = c
}

// Order returns new order
func (wx *WXMch) Order() *Order {
	return &Order{
		appid:  wx.AppID,
		mchid:  wx.MchID,
		apikey: wx.ApiKey,
	}
}

// Refund returns new refund
func (wx *WXMch) Refund() *Refund {
	return &Refund{
		appid:     wx.AppID,
		mchid:     wx.MchID,
		apikey:    wx.ApiKey,
		sslClient: wx.sslClient,
	}
}
