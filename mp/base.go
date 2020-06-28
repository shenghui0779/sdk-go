package mp

import "github.com/shenghui0779/gochat/utils"

// WXMP 微信小程序
type WXMP struct {
	AppID     string
	AppSecret string
	Client    *utils.HTTPClient
}

// Sns returns new sns
func (wx *WXMP) Sns(options ...utils.HTTPRequestOption) *Sns {
	return &Sns{
		mp:      wx,
		options: options,
	}
}

// CgiBin returns new cgi-bin
func (wx *WXMP) CgiBin(options ...utils.HTTPRequestOption) *CgiBin {
	return &CgiBin{
		mp:      wx,
		options: options,
	}
}

// BizDataCrypt returns new bizdatacrypt
func (wx *WXMP) BizDataCrypt(encryptedData, sessionKey, iv string) *BizDataCrypt {
	return &BizDataCrypt{
		mp:            wx,
		encryptedData: encryptedData,
		sessionKey:    sessionKey,
		iv:            iv,
	}
}

// Message returns new message
func (wx *WXMP) Message(options ...utils.HTTPRequestOption) *Message {
	return &Message{
		mp:      wx,
		options: options,
	}
}

// QRCode returns new qrcode
func (wx *WXMP) QRCode(options ...utils.HTTPRequestOption) *QRCode {
	return &QRCode{
		mp:      wx,
		options: options,
	}
}
