package minip

import (
	"encoding/json"
	"strconv"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsTemplAdd struct {
	TID       string `json:"tid"`
	KidList   []int  `json:"kidList"`
	SceneDesc string `json:"sceneDesc"`
}

type ResultTemplAdd struct {
	PriTmplID string `json:"priTmplId"`
}

// AddTemplate 订阅消息 - 组合模板并添加至帐号下的个人模板库
func AddTemplate(params *ParamsTemplAdd, result *ResultTemplAdd) wx.Action {
	return wx.NewPostAction(urls.MinipAddTemplate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsTemplDelete struct {
	PriTmplID string `json:"priTmplId"`
}

// DeleteTemplate 订阅消息 - 删除帐号下的个人模板
func DeleteTemplate(templID string) wx.Action {
	params := &ParamsTemplDelete{
		PriTmplID: templID,
	}

	return wx.NewPostAction(urls.MinipDeleteTemplate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type ResultCategoryGet struct {
	Data []*CategoryData `json:"data"`
}

type CategoryData struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// GetCategory 订阅消息 - 获取小程序账号的类目
func GetCategory(result *ResultCategoryGet) wx.Action {
	return wx.NewGetAction(urls.MinipGetCategory,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
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

// GetPubTemplateKeyWordsByID 订阅消息 - 获取模板标题下的关键词列表
func GetPubTemplateKeyWordsByID(tid string, result *ResultPubTemplKeywords) wx.Action {
	return wx.NewGetAction(urls.MinipGetetPubTemplateKeyWordsByID,
		wx.WithQuery("tid", tid),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ResultPubTemplTitleList struct {
	Count int              `json:"count"`
	Data  []*PubTemplTitle `json:"data"`
}

type PubTemplTitle struct {
	TID        int64  `json:"tid"`
	Title      string `json:"title"`
	Type       int    `json:"type"`
	CategoryID int64  `json:"categoryId"`
}

// ListPubTemplateTitle 订阅消息 - 获取帐号所属类目下的公共模板标题（多个类目id用逗号隔开）
func ListPubTemplateTitle(ids string, start, limit int, result *ResultPubTemplTitleList) wx.Action {
	return wx.NewGetAction(urls.MinipGetPubTemplateTitleList,
		wx.WithQuery("ids", ids),
		wx.WithQuery("start", strconv.Itoa(start)),
		wx.WithQuery("limit", strconv.Itoa(limit)),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ResultTemplList struct {
	Data []*TemplData `json:"data"`
}

type TemplData struct {
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

// ListTemplate 订阅消息 - 获取当前帐号下的个人模板列表
func ListTemplate(result *ResultTemplList) wx.Action {
	return wx.NewGetAction(urls.MinipGetTemplateList,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
