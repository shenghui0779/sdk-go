package pub

// WXPub 微信公众号
type WXPub struct {
	AccountID      string
	AppID          string
	AppSecret      string
	SignToken      string
	EncodingAESKey string
}

// Sns returns new sns
func (wx *WXPub) Sns() *Sns {
	return &Sns{
		appid:     wx.AppID,
		appsecret: wx.AppSecret,
	}
}

// CgiBin returns new cgi-bin
func (wx *WXPub) CgiBin() *CgiBin {
	return &CgiBin{
		appid:     wx.AppID,
		appsecret: wx.AppSecret,
	}
}

// MsgChiper returns new msg chiper
func (wx *WXPub) MsgChiper() *MsgChiper {
	return &MsgChiper{
		appid:          wx.AppID,
		encodingAESKey: wx.EncodingAESKey,
	}
}

func (wx *WXPub) Menu(accessToken string) *Menu {
	return &Menu{accessToken: accessToken}
}

func (wx *WXPub) Subsciber(accessToken string) *Subscriber {
	return &Subscriber{accessToken: accessToken}
}

// Reply returns new reply
func (wx *WXPub) Reply() *Reply {
	return &Reply{
		accountID:      wx.AccountID,
		appid:          wx.AppID,
		signToken:      wx.SignToken,
		encodingAESKey: wx.EncodingAESKey,
	}
}
