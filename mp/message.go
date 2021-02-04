package mp

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/wx"
)

// MessageBody 消息内容体
type MessageBody map[string]map[string]string

// MessageMinip 跳转小程序
type MessageMinip struct {
	AppID    string `json:"appid"`    // 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系，暂不支持小游戏）
	Pagepath string `json:"pagepath"` // 所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar），要求该小程序已发布，暂不支持小游戏
}

// UniformMessage 统一服务消息
type UniformMessage struct {
	MPTemplateMessage *TemplateMessage   // 小程序模板消息相关的信息，可以参考小程序模板消息接口; 有此节点则优先发送小程序模板消息
	OATemplateMessage *OATemplateMessage // 公众号模板消息相关的信息，可以参考公众号模板消息接口；有此节点并且没有 MPTemplateMessage 节点时，发送公众号模板消息
}

// SubscribeMessage 小程序订阅消息
type SubscribeMessage struct {
	TemplateID string      // 所需下发的订阅模板ID
	Page       string      // 点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转
	Data       MessageBody // 模板内容，格式形如：{"key1": {"value": any}, "key2": {"value": any}}
	MinipState string      // 跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版
	Lang       string      // 进入小程序查看”的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN
}

// TemplateMessage 小程序模板消息
type TemplateMessage struct {
	TemplateID      string      // 所需下发的模板消息的id
	Page            string      // 点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转
	FormID          string      // 表单提交场景下，为 submit 事件带上的 formId；支付场景下，为本次支付的 prepay_id
	Data            MessageBody // 模板内容，格式形如：{"key1": {"value": any}, "key2": {"value": any}}，不填则下发空模板
	EmphasisKeyword string      // 模板需要放大的关键词，不填则默认无放大，如："keyword1.DATA"
}

// OATemplateMessage 公众号模板消息
type OATemplateMessage struct {
	AppID       string        // 公众号appid，要求与小程序有绑定且同主体
	TemplateID  string        // 模板ID
	RedirectURL string        // 模板跳转链接（海外帐号没有跳转能力）
	MiniProgram *MessageMinip // 跳转小程序
	Data        MessageBody   // 模板内容，格式形如：{"key1": {"value": any}, "key2": {"value": any}}
}

// Uniform 发送统一服务消息
func SendUniformMessage(openID string, msg *UniformMessage) wx.Action {
	return wx.NewAction(UniformMessageSendURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			params := wx.X{
				"touser": openID,
			}

			// 小程序模板消息
			if msg.MPTemplateMessage != nil {
				tplMsg := wx.X{
					"template_id": msg.MPTemplateMessage.TemplateID,
					"form_id":     msg.MPTemplateMessage.FormID,
				}

				if msg.MPTemplateMessage.Page != "" {
					tplMsg["page"] = msg.MPTemplateMessage.Page
				}

				if msg.MPTemplateMessage.Data != nil {
					tplMsg["data"] = msg.MPTemplateMessage.Data
				}

				if msg.MPTemplateMessage.EmphasisKeyword != "" {
					tplMsg["emphasis_keyword"] = msg.MPTemplateMessage.EmphasisKeyword
				}

				params["weapp_template_msg"] = tplMsg
			}

			// 公众号模板消息
			if msg.OATemplateMessage != nil {
				tplMsg := wx.X{
					"appid":       msg.OATemplateMessage.AppID,
					"template_id": msg.OATemplateMessage.TemplateID,
					"data":        msg.OATemplateMessage.Data,
				}

				if msg.OATemplateMessage.RedirectURL != "" {
					tplMsg["url"] = msg.OATemplateMessage.RedirectURL
				}

				// 公众号模板消息所要跳转的小程序，小程序的必须与公众号具有绑定关系
				if msg.OATemplateMessage.MiniProgram != nil {
					tplMsg["miniprogram"] = msg.OATemplateMessage.MiniProgram
				}

				params["mp_template_msg"] = tplMsg
			}

			return json.Marshal(params)
		}),
	)
}

// SendSubscribeMessage 发送订阅消息
func SendSubscribeMessage(openID string, msg *SubscribeMessage) wx.Action {
	return wx.NewAction(SubscribeMessageSendURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			params := wx.X{
				"touser":      openID,
				"template_id": msg.TemplateID,
				"data":        msg.Data,
			}

			if msg.Page != "" {
				params["page"] = msg.Page
			}

			if msg.MinipState != "" {
				params["miniprogram_state"] = msg.MinipState
			}

			if msg.Lang != "" {
				params["lang"] = msg.Lang
			}

			return json.Marshal(params)
		}),
	)
}

// SendTemplateMessage 发送模板消息（已废弃，请使用订阅消息）
func SendTemplateMessage(openID string, msg *TemplateMessage) wx.Action {
	return wx.NewAction(TemplateMessageSendURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			params := wx.X{
				"touser":      openID,
				"template_id": msg.TemplateID,
				"form_id":     msg.FormID,
			}

			if msg.Page != "" {
				params["page"] = msg.Page
			}

			if msg.Data != nil {
				params["data"] = msg.Data
			}

			if msg.EmphasisKeyword != "" {
				params["emphasis_keyword"] = msg.EmphasisKeyword
			}

			return json.Marshal(params)
		}),
	)
}

// KFTextMessage 客服文本消息
type KFTextMessage struct {
	Content string `json:"content"` // 文本消息内容
}

// SendKFTextMessage 发送客服文本消息
func SendKFTextMessage(openID string, msg *KFTextMessage) wx.Action {
	return wx.NewAction(KFMessageSendURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(wx.X{
				"touser":  openID,
				"msgtype": "text",
				"text":    msg,
			})
		}),
	)
}

// KFImageMessage 客服图片消息
type KFImageMessage struct {
	MediaID string `json:"media_id"` // 发送的图片的媒体ID，通过 新增素材接口 上传图片文件获得
}

// SendKFImageMessage 发送客服图片消息
func SendKFImageMessage(openID string, msg *KFImageMessage) wx.Action {
	return wx.NewAction(KFMessageSendURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(wx.X{
				"touser":  openID,
				"msgtype": "image",
				"image":   msg,
			})
		}),
	)
}

// KFLinkMessage 客服图文链接消息
type KFLinkMessage struct {
	Title       string `json:"title"`       // 消息标题
	Description string `json:"description"` // 图文链接消息
	RedirectURL string `json:"url"`         // 图文链接消息被点击后跳转的链接
	ThumbURL    string `json:"thumb_url"`   // 图文链接消息的图片链接，支持 JPG、PNG 格式，较好的效果为大图 640 wx.X 320，小图 80 wx.X 80
}

// SendKFLinkMessage 发送客服图文链接消息
func SendKFLinkMessage(openID string, msg *KFLinkMessage) wx.Action {
	return wx.NewAction(KFMessageSendURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(wx.X{
				"touser":  openID,
				"msgtype": "link",
				"link":    msg,
			})
		}),
	)
}

// KFMinipMessage 客服小程序卡片消息
type KFMinipMessage struct {
	Title        string `json:"title"`          // 消息标题
	Pagepath     string `json:"pagepath"`       // 小程序的页面路径，跟app.json对齐，支持参数，比如pages/index/index?foo=bar
	ThumbMediaID string `json:"thumb_media_id"` // 小程序消息卡片的封面， image 类型的 media_id，通过 新增素材接口 上传图片文件获得，建议大小为 520*416
}

// SendKFMinipMessage 发送客服小程序卡片消息
func SendKFMinipMessage(openID string, msg *KFMinipMessage) wx.Action {
	return wx.NewAction(KFMessageSendURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(wx.X{
				"touser":          openID,
				"msgtype":         "miniprogrampage",
				"miniprogrampage": msg,
			})
		}),
	)
}

// TypeCommand 输入状态命令
type TypeCommand string

// 微信支持的输入状态命令
const (
	Typing       TypeCommand = "Typing"       // 正在输入
	CancelTyping TypeCommand = "CancelTyping" // 取消输入
)

// SetTyping 下发当前输入状态（仅支持客服消息）
func SetTyping(openID string, cmd TypeCommand) wx.Action {
	return wx.NewAction(SetTypingURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(wx.X{
				"touser":  openID,
				"command": cmd,
			})
		}),
	)
}
