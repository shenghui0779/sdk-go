package offia

import (
	"encoding/json"
	"strconv"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsSubscribeTemplAdd struct {
	TID       string `json:"tid"`
	KidList   []int  `json:"kidList"`
	SceneDesc string `json:"sceneDesc"`
}

type ResultSubscribeTemplAdd struct {
	PriTmplID string `json:"priTmplId"`
}

// AddSubscribeTemplate 订阅通知 - 选用模板
func AddSubscribeTemplate(params *ParamsSubscribeTemplAdd, result *ResultSubscribeTemplAdd) wx.Action {
	return wx.NewPostAction(urls.OffiaSubscribeAddTemplate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsSubscribeTemplDelete struct {
	PriTmplID string `json:"priTmplId"`
}

// DeleteSubscribeTemplate 订阅通知 - 删除模板
func DeleteSubscribeTemplate(priTmplID string) wx.Action {
	params := &ParamsSubscribeTemplDelete{
		PriTmplID: priTmplID,
	}

	return wx.NewPostAction(urls.OffiaSubscribeDeleteTemplate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type SubscribeCategory struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type ResultSubscribeCategory struct {
	Data []*SubscribeCategory `json:"data"`
}

// GetSubscribeCategory 订阅通知 - 获取公众号类目
func GetSubscribeCategory(result *ResultSubscribeCategory) wx.Action {
	return wx.NewGetAction(urls.OffiaSubscribeGetCategory,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type PubTemplKeywords struct {
	KID     int64  `json:"kid"`
	Name    string `json:"name"`
	Example string `json:"example"`
	Rule    string `json:"rule"`
}

type ResultPubTemplKeywords struct {
	Data []*PubTemplKeywords `json:"data"`
}

// GetPubTemplateKeywords 订阅通知 - 获取模板中的关键词
func GetPubTemplateKeywords(tid string, result *ResultPubTemplKeywords) wx.Action {
	return wx.NewGetAction(urls.OffiaSubscribeGetPubTemplateKeywords,
		wx.WithQuery("tid", tid),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type PubTemplTitle struct {
	TID        int64  `json:"tid"`
	Title      string `json:"title"`
	Type       int    `json:"type"`
	CategoryID string `json:"categoryId"`
}

type ResultPubTemplTitles struct {
	Count int              `json:"count"`
	Data  []*PubTemplTitle `json:"data"`
}

// GetPubTemplateTitles 订阅通知 - 获取类目下的公共模板
func GetPubTemplateTitles(ids string, start, limit int, result *ResultPubTemplTitles) wx.Action {
	return wx.NewGetAction(urls.OffiaSubscribeGetPubTemplateTitles,
		wx.WithQuery("ids", ids),
		wx.WithQuery("start", strconv.Itoa(start)),
		wx.WithQuery("limit", strconv.Itoa(limit)),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ResultSubscribeTemplList struct {
	Data []*SubscribeTemplInfo `json:"data"`
}

type SubscribeTemplInfo struct {
	PriTmplID string `json:"priTmplId"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Example   string `json:"example"`
	Type      int    `json:"type"`
}

// ListSubscribeTemplate 订阅通知 - 获取私有模板列表
func ListSubscribeTemplate(result *ResultSubscribeTemplList) wx.Action {
	return wx.NewGetAction(urls.OffiaSubscribeGetTemplateList,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type SubscribeMsg struct {
	ToUser     string       `json:"touser"`
	TemplateID string       `json:"template_id"`
	Page       string       `json:"page,omitempty"`
	Minip      *MsgMinip    `json:"miniprogram,omitempty"`
	Data       MsgTemplData `json:"data"`
}

// SendSubscribePageMsg 订阅通知 - 发送订阅通知（网页跳转）
func SendSubscribePageMsg(templateID, openid, page string, data MsgTemplData) wx.Action {
	msg := &SubscribeMsg{
		ToUser:     openid,
		TemplateID: templateID,
		Page:       page,
		Data:       data,
	}

	return wx.NewPostAction(urls.OffiaSubscribeMsgBizSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(msg)
		}),
	)
}

// SendSubscribeMinipMsg 订阅通知 - 发送订阅通知（小程序跳转）
func SendSubscribeMinipMsg(templateID, openid string, minip *MsgMinip, data MsgTemplData) wx.Action {
	msg := &SubscribeMsg{
		ToUser:     openid,
		TemplateID: templateID,
		Minip:      minip,
		Data:       data,
	}

	return wx.NewPostAction(urls.OffiaSubscribeMsgBizSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(msg)
		}),
	)
}
