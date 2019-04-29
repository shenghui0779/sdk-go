package mp

import "github.com/iiinsomnia/gochat/utils"

// WXMP 微信小程序
type WXMP struct {
	AppID     string
	AppSecret string
	client    *utils.HTTPClient
}

// SetHTTPClient set wxmp http client
func (wx *WXMP) SetHTTPClient(c *utils.HTTPClient) {
	wx.client = c
}

// Sns returns new sns
func (wx *WXMP) Sns() *Sns {
	sns := new(Sns)

	sns.appid = wx.AppID
	sns.appsecret = wx.AppSecret
	sns.client = wx.client

	return sns
}

// CgiBin returns new cgi-bin
func (wx *WXMP) CgiBin() *CgiBin {
	cgibin := new(CgiBin)

	cgibin.appid = wx.AppID
	cgibin.appsecret = wx.AppSecret
	cgibin.client = wx.client

	return cgibin
}

// BizDataCrypt returns new bizdatacrypt
func (wx *WXMP) BizDataCrypt() *BizDataCrypt {
	bizcrpyt := new(BizDataCrypt)

	bizcrpyt.appid = wx.AppID

	return bizcrpyt
}

// QRCode returns new qrcode
func (wx *WXMP) QRCode() *QRCode {
	qrcode := new(QRCode)

	qrcode.client = wx.client

	return qrcode
}
