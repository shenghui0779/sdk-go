package oa

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/wx"
	"github.com/tidwall/gjson"
)

// TemplateInfo 模板信息
type TemplateInfo struct {
	TemplateID      string `json:"template_id"`      // 模板ID
	Title           string `json:"title"`            // 模板标题
	PrimaryIndustry string `json:"primary_industry"` // 模板所属行业的一级行业
	DeputyIndustry  string `json:"deputy_industry"`  // 模板所属行业的二级行业
	Content         string `json:"content"`          // 模板内容
	Example         string `json:"example"`          // 模板示例
}

// GetTemplateList 获取模板列表
func GetTemplateList(dest *[]*TemplateInfo) wx.Action {
	return wx.NewAction(TemplateListURL,
		wx.WithMethod(wx.MethodGet),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal([]byte(gjson.GetBytes(resp, "template_list").Raw), dest)
		}),
	)
}

// DeleteTemplate 删除模板
func DeleteTemplate(templateID string) wx.Action {
	return wx.NewAction(TemplateDeleteURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(wx.X{"template_id": templateID})
		}),
	)
}

// MessageBody 消息内容体
type MessageBody map[string]map[string]string

// MessageMinip 跳转小程序
type MessageMinip struct {
	AppID    string `json:"appid"`    // 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系，暂不支持小游戏）
	Pagepath string `json:"pagepath"` // 所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar），要求该小程序已发布，暂不支持小游戏
}

// TemplateMessage 公众号模板消息
type TemplateMessage struct {
	TemplateID  string        // 模板ID
	URL         string        // 模板跳转链接（海外帐号没有跳转能力）
	MiniProgram *MessageMinip // 跳转小程序
	Data        MessageBody   // 模板内容，格式形如：{"key1":{"value":"V","color":"#"},"key2":{"value": "V","color":"#"}}
}

// SendTemplateMessage 发送模板消息
func SendTemplateMessage(openID string, msg *TemplateMessage) wx.Action {
	return wx.NewAction(TemplateMessageSendURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			params := wx.X{
				"touser":      openID,
				"template_id": msg.TemplateID,
				"data":        msg.Data,
			}

			if msg.URL != "" {
				params["url"] = msg.URL
			}

			if msg.MiniProgram != nil {
				params["miniprogram"] = msg.MiniProgram
			}

			return json.Marshal(params)
		}),
	)
}

// SendSubscribeMessage 发送一次性订阅消息
func SendSubscribeMessage(openID, scene, title string, msg *TemplateMessage) wx.Action {
	return wx.NewAction(SubscribeMessageSendURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			params := wx.X{
				"scene":       scene,
				"title":       title,
				"touser":      openID,
				"template_id": msg.TemplateID,
				"data":        msg.Data,
			}

			if msg.URL != "" {
				params["url"] = msg.URL
			}

			if msg.MiniProgram != nil {
				params["miniprogram"] = msg.MiniProgram
			}

			return json.Marshal(params)
		}),
	)
}
