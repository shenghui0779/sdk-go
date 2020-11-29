package mp

import (
	"encoding/json"
	"net/url"

	"github.com/shenghui0779/gochat/public"
)

// MessageBody 消息内容体
type MessageBody map[string]map[string]string

// UniformMessage 统一服务消息
type UniformMessage struct {
	OpenID             string             // 接收者（用户）的 openid，可以是小程序的openid，也可以是公众号的openid
	PubAppID           string             // 公众号appid，要求与小程序有绑定且同主体
	MPTemplateMessage  *TemplateMessage   // 小程序模板消息相关的信息，可以参考小程序模板消息接口; 有此节点则优先发送小程序模板消息
	PubTemplateMessage *OATemplateMessage // 公众号模板消息相关的信息，可以参考公众号模板消息接口；有此节点并且没有 MPTemplateMessage 节点时，发送公众号模板消息
}

// SubscribeMessage 小程序订阅消息
type SubscribeMessage struct {
	OpenID     string      // 接收者（用户）的 openid
	TemplateID string      // 所需下发的订阅模板ID
	PagePath   string      // 点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转
	Data       MessageBody // 模板内容，格式形如：{"key1": {"value": any}, "key2": {"value": any}}
	MPState    string      // 跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版
	Lang       string      // 进入小程序查看”的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN
}

// TemplateMessage 小程序模板消息
type TemplateMessage struct {
	OpenID          string      // 接收者（用户）的 openid
	TemplateID      string      // 所需下发的模板消息的id
	Page            string      // 点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转
	FormID          string      // 表单提交场景下，为 submit 事件带上的 formId；支付场景下，为本次支付的 prepay_id
	Data            MessageBody // 模板内容，格式形如：{"key1": {"value": any}, "key2": {"value": any}}，不填则下发空模板
	EmphasisKeyword string      // 模板需要放大的关键词，不填则默认无放大，如："keyword1.DATA"
}

// OATemplateMessage 公众号模板消息
type OATemplateMessage struct {
	OpenID      string      // 接收者（用户）的 openid
	TemplateID  string      // 模板ID
	RedirectURL string      // 模板跳转链接（海外帐号没有跳转能力）
	MPAppID     string      // 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系，暂不支持小游戏）
	MPPagePath  string      // 所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar），要求该小程序已发布，暂不支持小游戏
	Data        MessageBody // 模板内容，格式形如：{"key1": {"value": any}, "key2": {"value": any}}
}

// CustomerServiceMessage 小程序客服消息
type CustomerServiceMessage struct {
	OpenID      string        // 接收者（用户）的 openid
	MessageType string        // 消息类型：text|image|link|miniprogrampage
	Text        *TextMessage  // 文本消息
	Image       *ImageMessage // 图文链接
	Link        *LinkMessage  // 图文链接
	Page        *PageMessage  // 小程序卡片
}

// TextMessage 文本消息
type TextMessage struct {
	Content string // 文本消息内容
}

// ImageMessage 图片消息
type ImageMessage struct {
	MediaID string // 发送的图片的媒体ID，通过 新增素材接口 上传图片文件获得
}

// LinkMessage 图文链接
type LinkMessage struct {
	Title       string // 消息标题
	Description string // 图文链接消息
	RedirectURL string // 图文链接消息被点击后跳转的链接
	ThumbURL    string // 图文链接消息的图片链接，支持 JPG、PNG 格式，较好的效果为大图 640 public.X 320，小图 80 public.X 80
}

// PageMessage 小程序卡片
type PageMessage struct {
	Title        string // 消息标题
	Path         string // 小程序的页面路径，跟app.json对齐，支持参数，比如pages/index/index?foo=bar
	ThumbMediaID string // 小程序消息卡片的封面， image 类型的 media_id，通过 新增素材接口 上传图片文件获得，建议大小为 520*416
}

// TypingMessage 输入状态消息
type TypingMessage struct {
	OpenID  string // 接收者（用户）的 openid
	Command string // 命令：Typing|CancelTyping
}

// Uniform 发送统一服务消息
func SendUniformMessage(msg *UniformMessage) public.Action {
	return public.NewOpenPostAPI(UniformMessageSendURL, url.Values{}, public.NewPostBody(func() ([]byte, error) {
		params := public.X{
			"touser": msg.OpenID,
		}

		// 小程序模板消息
		if msg.MPTemplateMessage != nil {
			tplMsg := public.X{
				"template_id": msg.MPTemplateMessage.TemplateID,
				"form_id":     msg.MPTemplateMessage.FormID,
			}

			if msg.MPTemplateMessage.Page != "" {
				params["page"] = msg.MPTemplateMessage.Page
			}

			if msg.MPTemplateMessage.Data != nil {
				params["data"] = msg.MPTemplateMessage.Data
			}

			if msg.MPTemplateMessage.EmphasisKeyword != "" {
				params["emphasis_keyword"] = msg.MPTemplateMessage.EmphasisKeyword
			}

			params["weapp_template_msg"] = tplMsg
		}

		// 公众号模板消息
		if msg.PubTemplateMessage != nil {
			tplMsg := public.X{
				"appid":       msg.PubAppID,
				"template_id": msg.PubTemplateMessage.TemplateID,
				"data":        msg.PubTemplateMessage.Data,
			}

			if msg.PubTemplateMessage.RedirectURL != "" {
				tplMsg["url"] = msg.PubTemplateMessage.RedirectURL
			}

			// 公众号模板消息所要跳转的小程序，小程序的必须与公众号具有绑定关系
			if msg.PubTemplateMessage.MPAppID != "" {
				tplMsg["miniprogram"] = map[string]string{
					"appid":    msg.PubTemplateMessage.MPAppID,
					"pagepath": msg.PubTemplateMessage.MPPagePath,
				}
			}

			params["mp_template_msg"] = tplMsg
		}

		return json.Marshal(params)
	}), nil)
}

// SendSubscribeMessage 发送订阅消息
func SendSubscribeMessage(msg *SubscribeMessage) public.Action {
	return public.NewOpenPostAPI(SubscribeMessageSendURL, url.Values{}, public.NewPostBody(func() ([]byte, error) {
		params := public.X{
			"touser":      msg.OpenID,
			"template_id": msg.TemplateID,
			"data":        msg.Data,
		}

		if msg.PagePath != "" {
			params["page"] = msg.PagePath
		}

		if msg.MPState != "" {
			params["miniprogram_state"] = msg.MPState
		}

		if msg.Lang != "" {
			params["lang"] = msg.Lang
		}

		return json.Marshal(params)
	}), nil)
}

// SendTemplateMessage 发送模板消息
func SendTemplateMessage(msg *TemplateMessage) public.Action {
	return public.NewOpenPostAPI(TemplateMessageSendURL, url.Values{}, public.NewPostBody(func() ([]byte, error) {
		params := public.X{
			"touser":      msg.OpenID,
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
	}), nil)
}

// SendCustomerServiceMessage 发送客服消息
func SendCustomerServiceMessage(msg *CustomerServiceMessage) public.Action {
	return public.NewOpenPostAPI(CustomerServiceMessageSendURL, url.Values{}, public.NewPostBody(func() ([]byte, error) {
		params := public.X{
			"touser":  msg.OpenID,
			"msgtype": msg.MessageType,
		}

		if msg.Text != nil {
			params["text"] = msg.Text
		}

		if msg.Image != nil {
			params["image"] = msg.Image
		}

		if msg.Link != nil {
			params["link"] = msg.Link
		}

		if msg.Page != nil {
			params["miniprogrampage"] = msg.Page
		}

		return json.Marshal(params)
	}), nil)
}

// SetTyping 下发当前输入状态，仅支持客服消息
func SetTyping(msg *TypingMessage) public.Action {
	return public.NewOpenPostAPI(SetTypingURL, url.Values{}, public.NewPostBody(func() ([]byte, error) {
		params := public.X{
			"touser":  msg.OpenID,
			"command": msg.Command,
		}

		return json.Marshal(params)
	}), nil)
}
