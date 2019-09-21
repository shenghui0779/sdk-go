package mch

import (
	"fmt"
	"strconv"
	"time"

	"github.com/iiinsomnia/gochat/utils"
)

// WXMch 微信商户
type WXMch struct {
	AppID     string
	MchID     string
	ApiKey    string
	client    *utils.HTTPClient
	sslClient *utils.HTTPClient
}

// SetHTTPClient set wxmch http client
func (wx *WXMch) SetHTTPClient(c *utils.HTTPClient) {
	wx.client = c
}

// SetSSLClient set wxmch http ssl client
func (wx *WXMch) SetSSLClient(c *utils.HTTPClient) {
	wx.sslClient = c
}

// Order returns new order
func (wx *WXMch) Order() *Order {
	order := new(Order)

	order.appid = wx.AppID
	order.mchid = wx.MchID
	order.apikey = wx.ApiKey
	order.client = wx.client

	return order
}

// Refund returns new refund
func (wx *WXMch) Refund() *Refund {
	refund := new(Refund)

	refund.appid = wx.AppID
	refund.mchid = wx.MchID
	refund.apikey = wx.ApiKey
	refund.client = wx.client
	refund.sslClient = wx.sslClient

	return refund
}

// Pappay returns new pappay
func (wx *WXMch) Pappay() *Pappay {
	pappay := new(Pappay)

	pappay.appid = wx.AppID
	pappay.mchid = wx.MchID
	pappay.apikey = wx.ApiKey
	pappay.client = wx.client

	return pappay
}

// APPAPI 用于APP拉起支付
func (wx *WXMch) APPAPI(prepayID string) utils.WXML {
	ch := utils.WXML{
		"appid":     wx.AppID,
		"partnerid": wx.MchID,
		"prepayid":  prepayID,
		"package":   "Sign=WXPay",
		"noncestr":  utils.NonceStr(),
		"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
	}

	ch["sign"] = SignWithMD5(ch, wx.ApiKey)

	return ch
}

// JSAPI 用于JS拉起支付
func (wx *WXMch) JSAPI(prepayID string) utils.WXML {
	ch := utils.WXML{
		"appId":     wx.AppID,
		"nonceStr":  utils.NonceStr(),
		"package":   fmt.Sprintf("prepay_id=%s", prepayID),
		"signType":  SignMD5,
		"timeStamp": strconv.FormatInt(time.Now().Unix(), 10),
	}

	ch["paySign"] = SignWithMD5(ch, wx.ApiKey)

	return ch
}
