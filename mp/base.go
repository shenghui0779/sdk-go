package mp

import (
	"errors"

	"github.com/shenghui0779/gochat/utils"
	"github.com/tidwall/gjson"
)

// WXMP 微信小程序
type WXMP struct {
	appid     string
	appsecret string
	client    *utils.WXClient
}

func New(appid, appsecret string) *WXMP {
	return &WXMP{
		appid:     appid,
		appsecret: appsecret,
		client:    utils.NewWXClient(),
	}
}

// Sns returns new sns
func (wx *WXMP) Sns(options ...utils.RequestOption) *Sns {
	return &Sns{
		mp:      wx,
		options: options,
	}
}

// CgiBin returns new cgi-bin
func (wx *WXMP) CgiBin(options ...utils.RequestOption) *CgiBin {
	return &CgiBin{
		mp:      wx,
		options: options,
	}
}

// BizDataCrypto returns new bizdata crypto
func (wx *WXMP) BizDataCrypto(encryptedData, sessionKey, iv string) *BizDataCrypto {
	return &BizDataCrypto{
		mp:            wx,
		encryptedData: encryptedData,
		sessionKey:    sessionKey,
		iv:            iv,
	}
}

// Message returns new message
func (wx *WXMP) Message(options ...utils.RequestOption) *Message {
	return &Message{
		mp:      wx,
		options: options,
	}
}

// QRCode returns new qrcode
func (wx *WXMP) QRCode(options ...utils.RequestOption) *QRCode {
	return &QRCode{
		mp:      wx,
		options: options,
	}
}

func (wx *WXMP) get(reqURL string, options ...utils.RequestOption) ([]byte, error) {
	resp, err := wx.client.Get(reqURL, options...)

	if err != nil {
		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return nil, errors.New(r.Get("errmsg").String())
	}

	return resp, nil
}

func (wx *WXMP) post(reqURL string, body []byte, options ...utils.RequestOption) ([]byte, error) {
	options = append(options, utils.WithHeader("Content-Type", "application/json; charset=utf-8"))

	resp, err := wx.client.Post(reqURL, body, options...)

	if err != nil {
		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return nil, errors.New(r.Get("errmsg").String())
	}

	return resp, nil
}

func (wx *WXMP) upload(reqURL string, body []byte, options ...utils.RequestOption) ([]byte, error) {
	options = append(options, utils.WithHeader("Content-Type", "application/x-www-form-urlencoded"))

	resp, err := wx.client.Post(reqURL, body, options...)

	if err != nil {
		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return nil, errors.New(r.Get("errmsg").String())
	}

	return resp, nil
}
