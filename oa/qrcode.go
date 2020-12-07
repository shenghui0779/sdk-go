package oa

import (
	"encoding/json"
	"net/url"

	"github.com/shenghui0779/gochat/wx"
)

// QRCode 二维码获取信息
type QRCode struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds int64  `json:"expire_seconds"`
	URL           string `json:"url"`
}

// CreateTempQRCode 创建临时二维码
func CreateTempQRCode(dest *QRCode, senceID int, expireSeconds ...int) wx.Action {
	return wx.NewOpenPostAPI(QRCodeCreateURL, url.Values{}, wx.NewPostBody(func() ([]byte, error) {
		params := wx.X{
			"action_name": "QR_SCENE",
			"action_info": map[string]map[string]int{
				"scene": {"scene_id": senceID},
			},
		}

		if len(expireSeconds) != 0 {
			params["expire_seconds"] = expireSeconds[0]
		}

		return json.Marshal(params)
	}), func(resp []byte) error {
		return json.Unmarshal(resp, dest)
	})
}

// CreatePermQRCode 创建永久二维码
func CreatePermQRCode(dest *QRCode, senceID int, expireSeconds ...int) wx.Action {
	return wx.NewOpenPostAPI(QRCodeCreateURL, url.Values{}, wx.NewPostBody(func() ([]byte, error) {
		params := wx.X{
			"action_name": "QR_LIMIT_SCENE",
			"action_info": map[string]map[string]int{
				"scene": {"scene_id": senceID},
			},
		}

		if len(expireSeconds) != 0 {
			params["expire_seconds"] = expireSeconds[0]
		}

		return json.Marshal(params)
	}), func(resp []byte) error {
		return json.Unmarshal(resp, dest)
	})
}
