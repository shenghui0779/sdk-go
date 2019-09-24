package mp

import "github.com/iiinsomnia/gochat/utils"

// WXMP 微信小程序
type WXMP struct {
	AppID     string
	AppSecret string
	Client    *utils.HTTPClient
}

// Sns returns new sns
func (wx *WXMP) Sns() *Sns {
	return &Sns{wx}
}

// CgiBin returns new cgi-bin
func (wx *WXMP) CgiBin() *CgiBin {
	return &CgiBin{wx}
}

// BizDataCrypt returns new bizdatacrypt
func (wx *WXMP) BizDataCrypt() *BizDataCrypt {
	return &BizDataCrypt{WXMP: wx}
}

// QRCode returns new qrcode
func (wx *WXMP) QRCode() *QRCode {
	return &QRCode{wx}
}
