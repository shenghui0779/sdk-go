package mp

// WXMP 微信小程序
type WXMP struct {
	AppID     string
	AppSecret string
}

// Sns returns new sns
func (wx *WXMP) Sns() *Sns {
	return &Sns{
		appid:     wx.AppID,
		appsecret: wx.AppSecret,
	}
}

// CgiBin returns new cgi-bin
func (wx *WXMP) CgiBin() *CgiBin {
	return &CgiBin{
		appid:     wx.AppID,
		appsecret: wx.AppSecret,
	}
}

// BizDataCrypt returns new bizdatacrypt
func (wx *WXMP) BizDataCrypt() *BizDataCrypt {
	return &BizDataCrypt{
		appid: wx.AppID,
	}
}
