package pub

import (
	"errors"

	"github.com/shenghui0779/gochat/utils"
	"github.com/tidwall/gjson"
)

// WXPub 微信公众号
type WXPub struct {
	accountid      string
	appid          string
	appsecret      string
	signToken      string
	encodingAESKey string
	client         *utils.WXClient
}

func New(accountid, appid, appsecret, signToken, encodingAESKey string) *WXPub {
	return &WXPub{
		accountid:      accountid,
		appid:          appid,
		appsecret:      appsecret,
		signToken:      signToken,
		encodingAESKey: encodingAESKey,
		client:         utils.NewWXClient(),
	}
}

// Sns returns new sns
func (wx *WXPub) Sns(options ...utils.RequestOption) *Sns {
	return &Sns{
		pub:     wx,
		options: options,
	}
}

// CgiBin returns new cgi-bin
func (wx *WXPub) CgiBin(options ...utils.RequestOption) *CgiBin {
	return &CgiBin{
		pub:     wx,
		options: options,
	}
}

// MsgCrypt returns new msg crypt
func (wx *WXPub) MsgCrypt(cipherMsg string) *MsgCrypt {
	return &MsgCrypt{
		pub:        wx,
		cipherText: cipherMsg,
	}
}

// Menu returns new menu
func (wx *WXPub) Menu(options ...utils.RequestOption) *Menu {
	return &Menu{
		pub:     wx,
		options: options,
	}
}

// Subscriber returns new subscriber
func (wx *WXPub) Subscriber(options ...utils.RequestOption) *Subscriber {
	return &Subscriber{
		pub:     wx,
		options: options,
	}
}

// Message returns new tpl msg
func (wx *WXPub) Message(options ...utils.RequestOption) *Message {
	return &Message{
		pub:     wx,
		options: options,
	}
}

// Reply returns new reply
func (wx *WXPub) Reply(openid string) *Reply {
	return &Reply{
		pub:    wx,
		openid: openid,
	}
}

func (wx *WXPub) get(reqURL string, options ...utils.RequestOption) ([]byte, error) {
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

func (wx *WXPub) post(reqURL string, body []byte, options ...utils.RequestOption) ([]byte, error) {
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
