package offia

import (
	"encoding/json"

	"github.com/shenghui0779/yiigo"
	"github.com/tidwall/gjson"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

// QRCode 二维码获取信息
type QRCode struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds int64  `json:"expire_seconds"`
	URL           string `json:"url"`
}

// CreateTempQRCode 创建临时二维码（expireSeconds：二维码有效时间，最大不超过2592000秒（即30天），不填，则默认有效期为30秒。）
func CreateTempQRCode(dest *QRCode, senceID int, expireSeconds ...int) wx.Action {
	return wx.NewPostAction(urls.OAQRCodeCreate,
		wx.WithBody(func() ([]byte, error) {
			params := yiigo.X{
				"action_name": "QR_SCENE",
				"action_info": map[string]map[string]int{
					"scene": {"scene_id": senceID},
				},
			}

			if len(expireSeconds) != 0 {
				params["expire_seconds"] = expireSeconds[0]
			}

			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// CreatePermQRCode 创建永久二维码（expireSeconds：二维码有效时间，最大不超过2592000秒（即30天），不填，则默认有效期为30秒。）
func CreatePermQRCode(dest *QRCode, senceID int) wx.Action {
	return wx.NewPostAction(urls.OAQRCodeCreate,
		wx.WithBody(func() ([]byte, error) {
			params := yiigo.X{
				"action_name": "QR_LIMIT_SCENE",
				"action_info": map[string]map[string]int{
					"scene": {"scene_id": senceID},
				},
			}

			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// ShortURL 短链接
type ShortURL struct {
	URL string
}

// Long2ShortURL 长链接转短链接（长链接支持http://、https://、weixin://wxpay格式的url）
func Long2ShortURL(dest *ShortURL, longURL string) wx.Action {
	return wx.NewPostAction(urls.OAShortURLGenerate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalWithNoEscapeHTML(yiigo.X{
				"action":   "long2short",
				"long_url": longURL,
			})
		}),
		wx.WithDecode(func(resp []byte) error {
			dest.URL = gjson.GetBytes(resp, "short_url").String()

			return nil
		}),
	)
}
