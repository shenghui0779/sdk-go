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
func (wx *WXPub) Sns() *Sns {
	return &Sns{wx}
}

// CgiBin returns new cgi-bin
func (wx *WXPub) CgiBin() *CgiBin {
	return &CgiBin{wx}
}

// MsgChiper returns new msg chiper
func (wx *WXPub) MsgChiper() *MsgChiper {
	return &MsgChiper{wx}
}

// Menu returns new menu
func (wx *WXPub) Menu() *Menu {
	return &Menu{wx}
}

// Subscriber returns new subscriber
func (wx *WXPub) Subscriber() *Subscriber {
	return &Subscriber{wx}
}

// TplMsg returns new tpl msg
func (wx *WXPub) TplMsg() *TplMsg {
	return &TplMsg{wx}
}

// Reply returns new reply
func (wx *WXPub) Reply() *Reply {
	return &Reply{WXPub: wx}
}
