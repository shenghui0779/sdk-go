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

// SetIndustry 基础消息能力 - 模板消息 - 设置所属行业
func SetIndustry(id1, id2 string) wx.Action {
	params := &ParamsIndustrySet{
		IndustryID1: id1,
		IndustryID2: id2,
	}

	return wx.NewPostAction(urls.OffiaSetIndustry,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
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

// GetIndustry 基础消息能力 - 模板消息 - 获取设置的行业信息
func GetIndustry(result *ResultIndustryGet) wx.Action {
	return wx.NewGetAction(urls.OffiaGetIndustry,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsTemplAdd struct {
	TemplateIDShort string `json:"template_id_short"`
}

type ResultTemplAdd struct {
	TemplateID string `json:"template_id"`
}

// AddTemplate 基础消息能力 - 模板消息 - 获得模板ID
func AddTemplate(templIDShort string, result *ResultTemplAdd) wx.Action {
	params := &ParamsTemplAdd{
		TemplateIDShort: templIDShort,
	}

	return wx.NewPostAction(urls.OffiaTemplateAdd,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
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

// GetAllPrivateTemplate 基础消息能力 - 模板消息 - 获取模板列表
func GetAllPrivateTemplate(result *ResultAllPrivateTemplate) wx.Action {
	return wx.NewGetAction(urls.OffiaGetAllPrivateTemplate,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsPrivateTemplateDel struct {
	TemplateID string `json:"template_id"`
}

// DelPrivateTemplate 基础消息能力 - 模板消息 - 删除模板
func DelPrivateTemplate(templID string) wx.Action {
	params := &ParamsPrivateTemplateDel{
		TemplateID: templID,
	}

	return wx.NewPostAction(urls.OffiaDelPrivateTemplate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type MsgTemplValue struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}

// MsgTemplData 消息模板内容
type MsgTemplData map[string]*MsgTemplValue

// MsgMinip 跳转小程序
type MsgMinip struct {
	AppID    string `json:"appid"`              // 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系，暂不支持小游戏）
	PagePath string `json:"pagepath,omitempty"` // 所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar），要求该小程序已发布，暂不支持小游戏
}

type TemplateMsg struct {
	ToUser     string       `json:"touser"`                // 接收者openid
	TemplateID string       `json:"template_id"`           // 模板ID
	URL        string       `json:"url,omitempty"`         // 模板跳转链接（海外帐号没有跳转能力）
	Minip      *MsgMinip    `json:"miniprogram,omitempty"` // 跳小程序所需数据，不需跳小程序可不用传该数据
	Data       MsgTemplData `json:"data"`                  // 模板内容，格式形如：{"key1":{"value":"V","color":"#"},"key2":{"value": "V","color":"#"}}
}

// SendTemplateMsg 基础消息能力 - 模板消息 - 发送模板消息
func SendTemplateMsg(msg *TemplateMsg) wx.Action {
	return wx.NewPostAction(urls.OffiaTemplateMsgSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(msg)
		}),
	)
}

type TemplateSubscribeMsg struct {
	ToUser     string       `json:"touser"`                // 接收者openid
	Scene      string       `json:"scene"`                 // 订阅场景值
	Title      string       `json:"title"`                 // 消息标题，15字以内
	TemplateID string       `json:"template_id"`           // 模板ID
	URL        string       `json:"url,omitempty"`         // 点击消息跳转的链接，需要有ICP备案
	Minip      *MsgMinip    `json:"miniprogram,omitempty"` // 跳小程序所需数据，不需跳小程序可不用传该数据
	Data       MsgTemplData `json:"data"`                  // 消息正文，value为消息内容文本（200字以内），没有固定格式，可用\n换行，color为整段消息内容的字体颜色（目前仅支持整段消息为一种颜色）
}

// SendSubscribeTemplateMsg 基础消息能力 - 公众号一次性订阅消息
func SendSubscribeTemplateMsg(msg *TemplateSubscribeMsg) wx.Action {
	return wx.NewPostAction(urls.OffiaSubscribeTemplateMsgSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(msg)
		}),
	)
}
