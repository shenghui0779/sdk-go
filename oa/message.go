package oa

import (
	"encoding/json"
	"fmt"

	"github.com/shenghui0779/gochat/helpers"
)

// MessageBody 消息内容体
type MessageBody map[string]map[string]string

// TemplateMessage 公众号模板消息
type TemplateMessage struct {
	OpenID      string      // 接收者（用户）的 openid
	TemplateID  string      // 模板ID
	RedirectURL string      // 模板跳转链接（海外帐号没有跳转能力）
	MPAppID     string      // 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系，暂不支持小游戏）
	MPPagePath  string      // 所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar），要求该小程序已发布，暂不支持小游戏
	Data        MessageBody // 模板内容，格式形如：{"key1": {"value": any}, "key2": {"value": any}}
	Color       string      // 模板内容字体颜色，不填默认为黑色
}

// SendTemplateMessage 发送模板消息
func SendTemplateMessage(msg *TemplateMessage) Action {
	return &WechatAPI{
		body: helpers.NewPostBody(func() ([]byte, error) {
			params := helpers.X{
				"touser":      msg.OpenID,
				"template_id": msg.TemplateID,
				"data":        msg.Data,
			}

			if msg.RedirectURL != "" {
				params["url"] = msg.RedirectURL
			}

			if msg.MPAppID != "" {
				params["miniprogram"] = map[string]string{
					"appid":    msg.MPAppID,
					"pagepath": msg.MPPagePath,
				}
			}

			return json.Marshal(params)
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("POST|%s?access_token=%s", TemplateMessageSendURL, accessToken)
		},
	}
}
