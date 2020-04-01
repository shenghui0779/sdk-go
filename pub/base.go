package pub

import "github.com/iiinsomnia/gochat/utils"

// WXPub 微信公众号
type WXPub struct {
	AccountID      string
	AppID          string
	AppSecret      string
	SignToken      string
	EncodingAESKey string
	Client         *utils.HTTPClient
}

// Sns returns new sns
func (wx *WXPub) Sns(options ...utils.HTTPRequestOption) *Sns {
	return &Sns{
		pub:     wx,
		options: options,
	}
}

// CgiBin returns new cgi-bin
func (wx *WXPub) CgiBin(options ...utils.HTTPRequestOption) *CgiBin {
	return &CgiBin{
		pub:     wx,
		options: options,
	}
}

// MsgChiper returns new event msg crypt
func (wx *WXPub) EventMsgCrypt(cipherMsg string) *MsgCrypt {
	return &MsgCrypt{
		pub:        wx,
		cipherText: cipherMsg,
	}
}

// Menu returns new menu
func (wx *WXPub) Menu(options ...utils.HTTPRequestOption) *Menu {
	return &Menu{
		pub:     wx,
		options: options,
	}
}

// Subscriber returns new subscriber
func (wx *WXPub) Subscriber(options ...utils.HTTPRequestOption) *Subscriber {
	return &Subscriber{
		pub:     wx,
		options: options,
	}
}

// TplMsg returns new tpl msg
func (wx *WXPub) TplMsg(options ...utils.HTTPRequestOption) *TplMsg {
	return &TplMsg{
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
