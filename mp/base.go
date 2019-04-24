package mp

type WXMP struct {
	AppID     string
	AppSecret string
}

func (wx *WXMP) UseSns() *Sns {
	return &Sns{
		appID:     wx.AppID,
		appSecret: wx.AppSecret,
	}
}

func (wx *WXMP) UseBizDataCrypt() *BizDataCrypt {
	return &BizDataCrypt{
		appID: wx.AppID,
	}
}

func (wx *WXMP) UseCgiBin() *CgiBin {
	return &CgiBin{
		appID:     wx.AppID,
		appSecret: wx.AppSecret,
	}
}
