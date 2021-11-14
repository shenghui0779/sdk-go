package minip

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type MsgTemplateValue struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}

// MsgTemplateData 消息模板内容
type MsgTemplateData map[string]*MsgTemplateValue

// MinipState 小程序类型
type MinipState string

const (
	MinipDeveloper MinipState = "developer" // 开发版
	MinipTrial     MinipState = "trial"     // 体验版
	MinipFormal    MinipState = "formal"    // 正式版
)

// KFMsgType 客服消息类型
type KFMsgType string

const (
	KFText  KFMsgType = "text"            // 文本消息
	KFImage KFMsgType = "image"           // 图片消息
	KFLink  KFMsgType = "link"            // 图文链接消息
	KFMinip KFMsgType = "miniprogrampage" // 小程序卡片消息
)

// TypingCmd 输入状态命令
type TypingCmd string

// 微信支持的输入状态命令
const (
	Typing       TypingCmd = "Typing"       // 正在输入
	CancelTyping TypingCmd = "CancelTyping" // 取消输入
)

// MsgMinip 跳转小程序
type MsgMinip struct {
	AppID    string `json:"appid"`    // 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系，暂不支持小游戏）
	Pagepath string `json:"pagepath"` // 所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar），要求该小程序已发布，暂不支持小游戏
}

// TemplateMsg 统一服务消息数据
type TemplateMsg struct {
	AppID      string          `json:"appid"`       // 公众号appid，要求与小程序有绑定且同主体
	TemplateID string          `json:"template_id"` // 模板ID
	URL        string          `json:"url"`         // 模板跳转链接（海外帐号没有跳转能力）
	Minip      *MsgMinip       `json:"miniprogram"` // 跳转小程序
	Data       MsgTemplateData `json:"data"`        // 模板内容，格式形如：{"key1": {"value": any}, "key2": {"value": any}}
}

// ParamsUniformMsg 统一服务消息参数
type ParamsUniformMsg struct {
	ToUser        string       `json:"touser"` // 用户openid，可以是小程序的openid，也可以是mp_template_msg.appid对应的公众号的openid
	MPTemplateMsg *TemplateMsg `json:"mp_template_msg"`
}

// Uniform 发送统一服务消息
func SendUniformMessage(params *ParamsUniformMsg) wx.Action {
	return wx.NewPostAction(urls.MinipUniformMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

// ParamsSubscribeMsg 订阅消息参数
type ParamsSubscribeMsg struct {
	ToUser     string          `json:"touser"`                      // 接收者（用户）的 openid
	TemplateID string          `json:"template_id"`                 // 所需下发的订阅模板ID
	Page       string          `json:"page,omitempty"`              // 点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转
	MinipState MinipState      `json:"miniprogram_state,omitempty"` // 跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版
	Lang       string          `json:"lang,omitempty"`              // 进入小程序查看”的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN
	Data       MsgTemplateData `json:"data"`                        // 模板内容，格式形如：{"key1": {"value": any}, "key2": {"value": any}}
}

// SendSubscribeMessage 发送订阅消息
func SendSubscribeMessage(params *ParamsSubscribeMsg) wx.Action {
	return wx.NewPostAction(urls.MinipSubscribeMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

// MsgText 消息文本
type MsgText struct {
	Content string `json:"content"`
}

// ParamsKFTextMsg 客服文本消息参数
type ParamsKFTextMsg struct {
	ToUser  string    `json:"touser"`
	MsgType KFMsgType `json:"msgtype"`
	Text    *MsgText  `json:"text"`
}

// SendKFTextMessage 发送客服文本消息（支持插入跳小程序的文字链）
func SendKFTextMessage(openid, text string) wx.Action {
	params := &ParamsKFTextMsg{
		ToUser:  openid,
		MsgType: KFText,
		Text: &MsgText{
			Content: text,
		},
	}

	return wx.NewPostAction(urls.MinipKFMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

// MsgImage 消息图片
type MsgImage struct {
	MediaID string `json:"media_id"`
}

// ParamsKFImageMsg 客服图片消息参数
type ParamsKFImageMsg struct {
	ToUser  string    `json:"touser"`
	MsgType KFMsgType `json:"msgtype"`
	Image   *MsgImage `json:"image"`
}

// SendKFImageMessage 发送客服图片消息（媒体ID，通过素材接口上传获得）
func SendKFImageMessage(openid, mediaID string) wx.Action {
	params := &ParamsKFImageMsg{
		ToUser:  openid,
		MsgType: KFImage,
		Image: &MsgImage{
			MediaID: mediaID,
		},
	}

	return wx.NewPostAction(urls.MinipKFMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

// MsgLink 消息图文链接消息
type MsgLink struct {
	Title       string `json:"title"`       // 消息标题
	Description string `json:"description"` // 图文链接消息
	URL         string `json:"url"`         // 图文链接消息被点击后跳转的链接
	ThumbURL    string `json:"thumb_url"`   // 图文链接消息的图片链接，支持 JPG、PNG 格式，较好的效果为大图 640 yiigo.X 320，小图 80 yiigo.X 80
}

// ParamsKFLinkMsg 客服图文链接消息参数
type ParamsKFLinkMsg struct {
	ToUser  string    `json:"touser"`
	MsgType KFMsgType `json:"msgtype"`
	Link    *MsgLink  `json:"link"`
}

// SendKFLinkMessage 发送客服图文链接消息
func SendKFLinkMessage(openid string, msg *MsgLink) wx.Action {
	params := &ParamsKFLinkMsg{
		ToUser:  openid,
		MsgType: KFLink,
		Link:    msg,
	}

	return wx.NewPostAction(urls.MinipKFMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

// MsgMinipPage 小程序卡片消息
type MsgMinipPage struct {
	Title        string `json:"title"`          // 消息标题
	Pagepath     string `json:"pagepath"`       // 小程序的页面路径，跟app.json对齐，支持参数，比如pages/index/index?foo=bar
	ThumbMediaID string `json:"thumb_media_id"` // 小程序消息卡片的封面， image 类型的 media_id，通过 新增素材接口 上传图片文件获得，建议大小为 520*416
}

// ParamsKFMinipMsg 客服小程序卡片消息
type ParamsKFMinipMsg struct {
	ToUser    string        `json:"touser"`
	MsgType   KFMsgType     `json:"msgtype"`
	MinipPage *MsgMinipPage `json:"miniprogrampage"`
}

// SendKFMinipMessage 发送客服小程序卡片消息
func SendKFMinipMessage(openid string, msg *MsgMinipPage) wx.Action {
	params := &ParamsKFMinipMsg{
		ToUser:    openid,
		MsgType:   KFMinip,
		MinipPage: msg,
	}

	return wx.NewPostAction(urls.MinipKFMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsKFTyping struct {
	ToUser  string    `json:"touser"`
	Command TypingCmd `json:"command"`
}

// SendKFTyping 下发当前输入状态（仅支持客服消息）
func SendKFTyping(openid string, cmd TypingCmd) wx.Action {
	params := &ParamsKFTyping{
		ToUser:  openid,
		Command: cmd,
	}

	return wx.NewPostAction(urls.MinipKFTypingSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}
