package minip

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

// AddSubscribeTemplate 订阅消息 - 组合模板并添加至帐号下的个人模板库
func AddSubscribeTemplate(params *ParamsSubscribeTemplAdd, result *ResultSubscribeTemplAdd) wx.Action {
	return wx.NewPostAction(urls.MinipSubscribeAddTemplate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ParamsSubscribeTemplDelete struct {
	PriTmplID string `json:"priTmplId"`
}

// DeleteSubscribeTemplate 订阅消息 - 删除帐号下的个人模板
func DeleteSubscribeTemplate(priTmplID string) wx.Action {
	params := &ParamsSubscribeTemplDelete{
		PriTmplID: priTmplID,
	}

	return wx.NewPostAction(urls.MinipSubscribeDeleteTemplate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type ResultSubscribeCategory struct {
	Data []*SubscribeCategory `json:"data"`
}

type SubscribeCategory struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// GetSubscribeCategory 订阅消息 - 获取小程序账号的类目
func GetSubscribeCategory(result *ResultSubscribeCategory) wx.Action {
	return wx.NewGetAction(urls.MinipSubscribeGetCategory,
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ResultPubTemplKeywords struct {
	Data []*PubTemplKeywords `json:"data"`
}

type PubTemplKeywords struct {
	KID     int64  `json:"kid"`
	Name    string `json:"name"`
	Example string `json:"example"`
	Rule    string `json:"rule"`
}

// GetPubTemplateKeyWords 订阅消息 - 获取模板标题下的关键词列表
func GetPubTemplateKeyWords(tid string, result *ResultPubTemplKeywords) wx.Action {
	return wx.NewGetAction(urls.MinipSubscribeGetPubTemplateKeyWords,
		wx.WithQuery("tid", tid),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ResultPubTemplTitles struct {
	Count int              `json:"count"`
	Data  []*PubTemplTitle `json:"data"`
}

type PubTemplTitle struct {
	TID        int64  `json:"tid"`
	Title      string `json:"title"`
	Type       int    `json:"type"`
	CategoryID int64  `json:"categoryId"`
}

// GetPubTemplateTitles 订阅消息 - 获取帐号所属类目下的公共模板标题（多个类目id用逗号隔开）
func GetPubTemplateTitles(ids string, start, limit int, result *ResultPubTemplTitles) wx.Action {
	return wx.NewGetAction(urls.MinipSubscribeGetPubTemplateTitles,
		wx.WithQuery("ids", ids),
		wx.WithQuery("start", strconv.Itoa(start)),
		wx.WithQuery("limit", strconv.Itoa(limit)),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ResultSubscribeTemplList struct {
	Data []*SubscribeTemplInfo `json:"data"`
}

type SubscribeTemplInfo struct {
	PriTmplId            string              `json:"priTmplId"`
	Title                string              `json:"title"`
	Content              string              `json:"content"`
	Example              string              `json:"example"`
	Type                 int                 `json:"type"`
	KeywordEnumValueList []*KeywordEnumValue `json:"keywordEnumValueList"`
}

type KeywordEnumValue struct {
	EnumValueList []string `json:"enumValueList"`
	KeywordCode   string   `json:"keywordCode"`
}

// ListSubscribeTemplate 订阅消息 - 获取当前帐号下的个人模板列表
func ListSubscribeTemplate(result *ResultSubscribeTemplList) wx.Action {
	return wx.NewGetAction(urls.MinipSubscribeGetTemplateList,
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}
