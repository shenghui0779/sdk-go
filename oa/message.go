package oa

import (
	"encoding/json"
	"net/url"

	"github.com/shenghui0779/gochat/wx"
	"github.com/tidwall/gjson"
)

// TemplateInfo 模板信息
type TemplateInfo struct {
	TemplateID      string `json:"template_id"`
	Title           string `json:"title"`
	PrimaryIndustry string `json:"primary_industry"`
	DeputyIndustry  string `json:"deputy_industry"`
	Content         string `json:"content"`
	Example         string `json:"example"`
}

// GetTemplateList 获取模板列表
func GetTemplateList(dest *[]TemplateInfo) wx.Action {
	return wx.NewOpenGetAPI(TemplateListURL, url.Values{}, func(resp []byte) error {
		r := gjson.GetBytes(resp, "template_list")

		return json.Unmarshal([]byte(r.Raw), dest)
	})
}

// DeleteTemplate 删除模板
func DeleteTemplate(templateID string) wx.Action {
	return wx.NewOpenPostAPI(TemplateDeleteURL, url.Values{}, wx.NewPostBody(func() ([]byte, error) {
		return json.Marshal(wx.X{"template_id": templateID})
	}), nil)
}

// MessageBody 消息内容体
type MessageBody map[string]map[string]string

// TemplateMessage 公众号模板消息
type TemplateMessage struct {
	TemplateID  string      // 模板ID
	RedirectURL string      // 模板跳转链接（海外帐号没有跳转能力）
	MinipAppID  string      // 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系，暂不支持小游戏）
	MinipPage   string      // 所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar），要求该小程序已发布，暂不支持小游戏
	Data        MessageBody // 模板内容，格式形如：{"key1":{"value":"V","color":"#"},"key2":{"value": "V","color":"#"}}
}

// SendTemplateMessage 发送模板消息
func SendTemplateMessage(openID string, msg *TemplateMessage) wx.Action {
	return wx.NewOpenPostAPI(TemplateMessageSendURL, url.Values{}, wx.NewPostBody(func() ([]byte, error) {
		params := wx.X{
			"touser":      openID,
			"template_id": msg.TemplateID,
			"data":        msg.Data,
		}

		if msg.RedirectURL != "" {
			params["url"] = msg.RedirectURL
		}

		if msg.MinipAppID != "" {
			params["miniprogram"] = map[string]string{
				"appid":    msg.MinipAppID,
				"pagepath": msg.MinipPage,
			}
		}

		return json.Marshal(params)
	}), nil)
}

// SendSubscribeMessage 发送订阅消息
func SendSubscribeMessage(openID, scene, title string, msg *TemplateMessage) wx.Action {
	return wx.NewOpenPostAPI(SubscribeMessageSendURL, url.Values{}, wx.NewPostBody(func() ([]byte, error) {
		params := wx.X{
			"scene":       scene,
			"title":       title,
			"touser":      openID,
			"template_id": msg.TemplateID,
			"data":        msg.Data,
		}

		if msg.RedirectURL != "" {
			params["url"] = msg.RedirectURL
		}

		if msg.MinipAppID != "" {
			params["miniprogram"] = map[string]string{
				"appid":    msg.MinipAppID,
				"pagepath": msg.MinipPage,
			}
		}

		return json.Marshal(params)
	}), nil)
}
