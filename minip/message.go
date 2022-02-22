package minip

import (
	"github.com/shenghui0779/gochat/event"
	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type MsgTemplValue struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}

// MsgTemplData 消息模板内容
type MsgTemplData map[string]*MsgTemplValue

// MinipState 小程序类型
type MinipState string

const (
	MinipDeveloper MinipState = "developer" // 开发版
	MinipTrial     MinipState = "trial"     // 体验版
	MinipFormal    MinipState = "formal"    // 正式版
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
	AppID      string       `json:"appid"`       // 公众号appid，要求与小程序有绑定且同主体
	TemplateID string       `json:"template_id"` // 模板ID
	URL        string       `json:"url"`         // 模板跳转链接（海外帐号没有跳转能力）
	Minip      *MsgMinip    `json:"miniprogram"` // 跳转小程序
	Data       MsgTemplData `json:"data"`        // 模板内容，格式形如：{"key1": {"value": any}, "key2": {"value": any}}
}

// UniformMsg 统一服务消息参数
type UniformMsg struct {
	ToUser        string       `json:"touser"` // 用户openid，可以是小程序的openid，也可以是mp_template_msg.appid对应的公众号的openid
	MPTemplateMsg *TemplateMsg `json:"mp_template_msg"`
}

// SendUniformMsg 统一服务消息 - 发送统一服务消息
func SendUniformMsg(touser string, msg *TemplateMsg) wx.Action {
	params := &UniformMsg{
		ToUser:        touser,
		MPTemplateMsg: msg,
	}

	return wx.NewPostAction(urls.MinipUniformMsgSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

// SubscribeMsg 订阅消息参数
type SubscribeMsg struct {
	ToUser     string       `json:"touser"`                      // 接收者（用户）的 openid
	TemplateID string       `json:"template_id"`                 // 所需下发的订阅模板ID
	Page       string       `json:"page,omitempty"`              // 点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转
	MinipState MinipState   `json:"miniprogram_state,omitempty"` // 跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版
	Lang       string       `json:"lang,omitempty"`              // 进入小程序查看”的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN
	Data       MsgTemplData `json:"data"`                        // 模板内容，格式形如：{"key1": {"value": any}, "key2": {"value": any}}
}

// SendSubscribeMsg 订阅消息 - 发送订阅消息
func SendSubscribeMsg(msg *SubscribeMsg) wx.Action {
	return wx.NewPostAction(urls.MinipSubscribeMsgSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(msg)
		}),
	)
}

// KFText 客服消息文本
type KFText struct {
	Content string `json:"content"`
}

// KFMedia 客服消息媒体（图片等）
type KFMedia struct {
	MediaID string `json:"media_id"`
}

// KFLink 客服消息链接
type KFLink struct {
	Title       string `json:"title"`       // 消息标题
	Description string `json:"description"` // 图文链接消息
	URL         string `json:"url"`         // 图文链接消息被点击后跳转的链接
	ThumbURL    string `json:"thumb_url"`   // 图文链接消息的图片链接，支持 JPG、PNG 格式，较好的效果为大图 640 yiigo.X 320，小图 80 yiigo.X 80
}

// KFMinipPage 客服小程序卡片
type KFMinipPage struct {
	Title        string `json:"title"`          // 消息标题
	Pagepath     string `json:"pagepath"`       // 小程序的页面路径，跟app.json对齐，支持参数，比如pages/index/index?foo=bar
	ThumbMediaID string `json:"thumb_media_id"` // 小程序消息卡片的封面， image 类型的 media_id，通过 新增素材接口 上传图片文件获得，建议大小为 520*416
}

type KFMsg struct {
	ToUser    string        `json:"touser"`
	MsgType   event.MsgType `json:"msgtype"`
	Text      *KFText       `json:"text,omitempty"`
	Image     *KFMedia      `json:"image,omitempty"`
	Link      *KFLink       `json:"link,omitempty"`
	MinipPage *KFMinipPage  `json:"miniprogrampage,omitempty"`
}

// SendKFTextMsg 客服消息 - 发送客服文本消息（支持插入跳小程序的文字链）
func SendKFTextMsg(openid, content string) wx.Action {
	msg := &KFMsg{
		ToUser:  openid,
		MsgType: event.MsgText,
		Text: &KFText{
			Content: content,
		},
	}

	return wx.NewPostAction(urls.MinipKFMsgSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(msg)
		}),
	)
}

// SendKFImageMsg 客服消息 - 发送客服图片消息（媒体ID，通过素材接口上传获得）
func SendKFImageMsg(openid, mediaID string) wx.Action {
	msg := &KFMsg{
		ToUser:  openid,
		MsgType: event.MsgImage,
		Image: &KFMedia{
			MediaID: mediaID,
		},
	}

	return wx.NewPostAction(urls.MinipKFMsgSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(msg)
		}),
	)
}

// SendKFLinkMsg 客服消息 - 发送客服图文链接消息
func SendKFLinkMsg(openid string, link *KFLink) wx.Action {
	msg := &KFMsg{
		ToUser:  openid,
		MsgType: event.MsgLink,
		Link:    link,
	}

	return wx.NewPostAction(urls.MinipKFMsgSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(msg)
		}),
	)
}

// SendKFMinipMsg 客服消息 - 发送客服小程序卡片消息
func SendKFMinipMsg(openid string, minipPage *KFMinipPage) wx.Action {
	msg := &KFMsg{
		ToUser:    openid,
		MsgType:   event.MsgMinipPage,
		MinipPage: minipPage,
	}

	return wx.NewPostAction(urls.MinipKFMsgSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(msg)
		}),
	)
}

type KFTyping struct {
	ToUser  string    `json:"touser"`
	Command TypingCmd `json:"command"`
}

// SendKFTyping 客服消息 - 下发当前输入状态（仅支持客服消息）
func SendKFTyping(openid string, cmd TypingCmd) wx.Action {
	typing := &KFTyping{
		ToUser:  openid,
		Command: cmd,
	}

	return wx.NewPostAction(urls.MinipKFTypingSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(typing)
		}),
	)
}
