package school

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ResultSubscribeQRCode struct {
	QRCodeBig    string `json:"qrcode_big"`
	QRCodeMiddle string `json:"qrcode_middle"`
	QRCodeThumb  string `json:"qrcode_thumb"`
}

func GetSubscribeQRCode(result *ResultSubscribeQRCode) wx.Action {
	return wx.NewGetAction(urls.CorpSchoolGetSubscribeQRCode,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsSubscribeModeSet struct {
	SubscribeMode int `json:"subscribe_mode"`
}

func SetSubscribeMode(mode int) wx.Action {
	params := &ParamsSubscribeModeSet{
		SubscribeMode: mode,
	}

	return wx.NewPostAction(urls.CorpSchoolSetSubscribeMode,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ResultSubscribeModeGet struct {
	SubscribeMode int `json:"subscribe_mode"`
}

func GetSubscribeMode(result *ResultSubscribeModeGet) wx.Action {
	return wx.NewGetAction(urls.CorpSchoolGetSubscribeMode,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
