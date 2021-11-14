package offia

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsIndustrySet struct {
	IndustryID1 string `json:"industry_id1"`
	IndustryID2 string `json:"industry_id2"`
}

func SetIndustry(params *ParamsIndustrySet) wx.Action {
	return wx.NewPostAction(urls.OffiaSetIndustry,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type IndustryInfo struct {
	FirstClass  string `json:"first_class"`
	SecondClass string `json:"second_class"`
}

type ResultIndustryGet struct {
	PrimaryIndustry   *IndustryInfo `json:"primary_industry"`
	SecondaryIndustry *IndustryInfo `json:"secondary_industry"`
}

func GetIndustry(result *ResultIndustryGet) wx.Action {
	return wx.NewGetAction(urls.OffiaGetIndustry,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsTemplateAdd struct {
	TemplateIDShort string `json:"template_id_short"`
}

type ResultTemplateAdd struct {
	TemplateID string `json:"template_id"`
}

func AddTemplate(params *ParamsTemplateAdd, result *ResultTemplateAdd) wx.Action {
	return wx.NewPostAction(urls.OffiaTemplateAdd,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// TemplateInfo 模板信息
type TemplateInfo struct {
	TemplateID      string `json:"template_id"`      // 模板ID
	Title           string `json:"title"`            // 模板标题
	PrimaryIndustry string `json:"primary_industry"` // 模板所属行业的一级行业
	DeputyIndustry  string `json:"deputy_industry"`  // 模板所属行业的二级行业
	Content         string `json:"content"`          // 模板内容
	Example         string `json:"example"`          // 模板示例
}

type ResultAllPrivateTemplate struct {
	TemplateList []*TemplateInfo `json:"template_list"`
}

// GetAllPrivateTemplate 获取模板列表
func GetAllPrivateTemplate(result *ResultAllPrivateTemplate) wx.Action {
	return wx.NewGetAction(urls.OffiaTemplateList,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsPrivateTemplateDel struct {
	TemplateID string `json:"template_id"`
}

// DelPrivateTemplate 删除模板
func DelPrivateTemplate(params *ParamsPrivateTemplateDel) wx.Action {
	return wx.NewPostAction(urls.OffiaTemplateDelete,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type MsgTemplateValue struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}

// MsgTemplateData 消息模板内容
type MsgTemplateData map[string]*MsgTemplateValue

// MsgMinip 跳转小程序
type MsgMinip struct {
	AppID    string `json:"appid"`              // 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系，暂不支持小游戏）
	Pagepath string `json:"pagepath,omitempty"` // 所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar），要求该小程序已发布，暂不支持小游戏
}

// TemplateMsg 公众号模板消息
type ParamsTemplateMsg struct {
	ToUser     string          `json:"touser"`                // 接收者openid
	TemplateID string          `json:"template_id"`           // 模板ID
	URL        string          `json:"url,omitempty"`         // 模板跳转链接（海外帐号没有跳转能力）
	Minip      *MsgMinip       `json:"miniprogram,omitempty"` // 跳小程序所需数据，不需跳小程序可不用传该数据
	Data       MsgTemplateData `json:"data"`                  // 模板内容，格式形如：{"key1":{"value":"V","color":"#"},"key2":{"value": "V","color":"#"}}
}

// SendTemplateMsg 发送模板消息
func SendTemplateMsg(params *ParamsTemplateMsg) wx.Action {
	return wx.NewPostAction(urls.OffiaTemplateMsgSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsTemplateMsgSubscribe struct {
	ToUser     string          `json:"touser"`                // 接收者openid
	Scene      string          `json:"scene"`                 // 订阅场景值
	Title      string          `json:"title"`                 // 消息标题，15字以内
	TemplateID string          `json:"template_id"`           // 模板ID
	URL        string          `json:"url,omitempty"`         // 点击消息跳转的链接，需要有ICP备案
	Minip      *MsgMinip       `json:"miniprogram,omitempty"` // 跳小程序所需数据，不需跳小程序可不用传该数据
	Data       MsgTemplateData `json:"data"`                  // 消息正文，value为消息内容文本（200字以内），没有固定格式，可用\n换行，color为整段消息内容的字体颜色（目前仅支持整段消息为一种颜色）
}

// SubscribeTemplateMsg 推送订阅模板消息给到授权微信用户
func SubscribeTemplateMsg(params *ParamsTemplateMsgSubscribe) wx.Action {
	return wx.NewPostAction(urls.OffiaSubscribeMsgSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}
