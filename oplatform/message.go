package oplatform

import (
	"encoding/json"
	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
	"github.com/shenghui0779/yiigo"
)

// SendKFTextMessage 发送客服文本消息（支持插入跳小程序的文字链）
func SendKFTextMessage(openID, accessToken, text string, kfAccount ...string) wx.Action {
	return wx.NewPostAction(urls.OffiaKFMessageSend,
		wx.WithQuery("access_token", accessToken),
		wx.WithBody(func() ([]byte, error) {
			data := yiigo.X{
				"touser":  openID,
				"msgtype": "text",
				"text": yiigo.X{
					"content": text,
				},
			}

			if len(kfAccount) != 0 {
				data["customservice"] = yiigo.X{
					"kf_account": kfAccount[0],
				}
			}

			return json.Marshal(data)
		}),
	)
}
