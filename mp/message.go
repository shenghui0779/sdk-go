package mp

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tidwall/gjson"

	"github.com/shenghui0779/gochat/pub"
	"github.com/shenghui0779/gochat/utils"
)

// MsgBody 消息内容体
type MsgBody map[string]map[string]string

// UniformMsg 统一服务消息
type UniformMsg struct {
	OpenID    string      // 接收者（用户）的 openid，可以是小程序的openid，也可以是公众号的openid
	PubAppID  string      // 公众号appid，要求与小程序有绑定且同主体
	MPTplMsg  *TplMsg     // 小程序模板消息相关的信息，可以参考小程序模板消息接口; 有此节点则优先发送小程序模板消息
	PubTplMsg *pub.TplMsg // 公众号模板消息相关的信息，可以参考公众号模板消息接口；有此节点并且没有 MPTplMsg 节点时，发送公众号模板消息
}

// SubscribeMsg 小程序订阅消息
type SubscribeMsg struct {
	OpenID   string  // 接收者（用户）的 openid
	TplID    string  // 所需下发的订阅模板ID
	PagePath string  // 点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转
	Data     MsgBody // 模板内容，格式形如：{"key1": {"value": any}, "key2": {"value": any}}
	MPState  string  // 跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版
	Lang     string  // 进入小程序查看”的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN
}

// TplMsg 小程序模板消息
type TplMsg struct {
	OpenID          string  // 接收者（用户）的 openid
	TplID           string  // 所需下发的模板消息的id
	Page            string  // 点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转
	FormID          string  // 表单提交场景下，为 submit 事件带上的 formId；支付场景下，为本次支付的 prepay_id
	Data            MsgBody // 模板内容，格式形如：{"key1": {"value": any}, "key2": {"value": any}}，不填则下发空模板
	EmphasisKeyword string  // 模板需要放大的关键词，不填则默认无放大，如："keyword1.DATA"
}

// CustomerServiceMsg 小程序客服消息
type CustomerServiceMsg struct {
	OpenID  string    // 接收者（用户）的 openid
	MsgType string    // 消息类型：text|image|link|miniprogrampage
	Text    *TextMsg  // 文本消息
	Image   *ImageMsg // 图文链接
	Link    *LinkMsg  // 图文链接
	Page    *PageMsg  // 小程序卡片
}

// TextMsg 文本消息
type TextMsg struct {
	Content string // 文本消息内容
}

// ImageMsg 图片消息
type ImageMsg struct {
	MediaID string // 发送的图片的媒体ID，通过 新增素材接口 上传图片文件获得
}

// LinkMsg 图文链接
type LinkMsg struct {
	Title       string // 消息标题
	Description string // 图文链接消息
	RedirectURL string // 图文链接消息被点击后跳转的链接
	ThumbURL    string // 图文链接消息的图片链接，支持 JPG、PNG 格式，较好的效果为大图 640 X 320，小图 80 X 80
}

// PageMsg 小程序卡片
type PageMsg struct {
	Title        string // 消息标题
	Path         string // 小程序的页面路径，跟app.json对齐，支持参数，比如pages/index/index?foo=bar
	ThumbMediaID string // 小程序消息卡片的封面， image 类型的 media_id，通过 新增素材接口 上传图片文件获得，建议大小为 520*416
}

// TypingMsg 输入状态消息
type TypingMsg struct {
	OpenID  string // 接收者（用户）的 openid
	Command string // 命令：Typing|CancelTyping
}

// Message 小程序消息
type Message struct {
	mp      *WXMP
	options []utils.HTTPRequestOption
}

// Uniform 发送统一服务消息
func (m *Message) Uniform(msg *UniformMsg, accessToken string) error {
	body := utils.X{
		"touser": msg.OpenID,
	}

	// 小程序模板消息
	if msg.MPTplMsg != nil {
		tplMsg := utils.X{
			"template_id": msg.MPTplMsg.TplID,
			"form_id":     msg.MPTplMsg.FormID,
		}

		if msg.MPTplMsg.Page != "" {
			body["page"] = msg.MPTplMsg.Page
		}

		if msg.MPTplMsg.Data != nil {
			body["data"] = msg.MPTplMsg.Data
		}

		if msg.MPTplMsg.EmphasisKeyword != "" {
			body["emphasis_keyword"] = msg.MPTplMsg.EmphasisKeyword
		}

		body["weapp_template_msg"] = tplMsg
	}

	// 公众号模板消息
	if msg.PubTplMsg != nil {
		tplMsg := utils.X{
			"appid":       msg.PubAppID,
			"template_id": msg.PubTplMsg.TplID,
			"data":        msg.PubTplMsg.Data,
		}

		if msg.PubTplMsg.RedirectURL != "" {
			tplMsg["url"] = msg.PubTplMsg.RedirectURL
		}

		// 公众号模板消息所要跳转的小程序，小程序的必须与公众号具有绑定关系
		if msg.PubTplMsg.MPAppID != "" {
			tplMsg["miniprogram"] = map[string]string{
				"appid":    msg.PubTplMsg.MPAppID,
				"pagepath": msg.PubTplMsg.MPPagePath,
			}
		}

		body["mp_template_msg"] = tplMsg
	}

	b, err := json.Marshal(body)

	if err != nil {
		return err
	}

	m.options = append(m.options, utils.WithRequestHeader("Content-Type", "application/json; charset=utf-8"))

	resp, err := m.mp.Client.Post(fmt.Sprintf("%s?access_token=%s", UniformMsgSendURL, accessToken), b, m.options...)

	if err != nil {
		return err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return errors.New(r.Get("errmsg").String())
	}

	return nil
}

// Subscribe 发送订阅消息
func (m *Message) Subscribe(msg *SubscribeMsg, accessToken string) error {
	body := utils.X{
		"touser":      msg.OpenID,
		"template_id": msg.TplID,
		"data":        msg.Data,
	}

	if msg.PagePath != "" {
		body["page"] = msg.PagePath
	}

	if msg.MPState != "" {
		body["miniprogram_state"] = msg.MPState
	}

	if msg.Lang != "" {
		body["lang"] = msg.Lang
	}

	b, err := json.Marshal(body)

	if err != nil {
		return err
	}

	m.options = append(m.options, utils.WithRequestHeader("Content-Type", "application/json; charset=utf-8"))

	resp, err := m.mp.Client.Post(fmt.Sprintf("%s?access_token=%s", SubscribeMsgSendURL, accessToken), b, m.options...)

	if err != nil {
		return err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return errors.New(r.Get("errmsg").String())
	}

	return nil
}

// Template 发送模板消息
func (m *Message) Template(msg *TplMsg, accessToken string) error {
	body := utils.X{
		"touser":      msg.OpenID,
		"template_id": msg.TplID,
		"form_id":     msg.FormID,
	}

	if msg.Page != "" {
		body["page"] = msg.Page
	}

	if msg.Data != nil {
		body["data"] = msg.Data
	}

	if msg.EmphasisKeyword != "" {
		body["emphasis_keyword"] = msg.EmphasisKeyword
	}

	b, err := json.Marshal(body)

	if err != nil {
		return err
	}

	m.options = append(m.options, utils.WithRequestHeader("Content-Type", "application/json; charset=utf-8"))

	resp, err := m.mp.Client.Post(fmt.Sprintf("%s?access_token=%s", TplMsgSendURL, accessToken), b, m.options...)

	if err != nil {
		return err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return errors.New(r.Get("errmsg").String())
	}

	return nil
}

// CustomerService 发送客服消息
func (m *Message) CustomerService(msg *CustomerServiceMsg, accessToken string) error {
	body := utils.X{
		"touser":  msg.OpenID,
		"msgtype": msg.MsgType,
	}

	if msg.Text != nil {
		body["text"] = msg.Text
	}

	if msg.Image != nil {
		body["image"] = msg.Image
	}

	if msg.Link != nil {
		body["link"] = msg.Link
	}

	if msg.Page != nil {
		body["miniprogrampage"] = msg.Page
	}

	b, err := json.Marshal(body)

	if err != nil {
		return err
	}

	m.options = append(m.options, utils.WithRequestHeader("Content-Type", "application/json; charset=utf-8"))

	resp, err := m.mp.Client.Post(fmt.Sprintf("%s?access_token=%s", CustomerServiceMsgSendURL, accessToken), b, m.options...)

	if err != nil {
		return err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return errors.New(r.Get("errmsg").String())
	}

	return nil
}

// SetTyping 下发当前输入状态，仅支持客服消息
func (m *Message) SetTyping(msg *TypingMsg, accessToken string) error {
	body := utils.X{
		"touser":  msg.OpenID,
		"command": msg.Command,
	}

	b, err := json.Marshal(body)

	if err != nil {
		return err
	}

	m.options = append(m.options, utils.WithRequestHeader("Content-Type", "application/json; charset=utf-8"))

	resp, err := m.mp.Client.Post(fmt.Sprintf("%s?access_token=%s", SetTypingURL, accessToken), b, m.options...)

	if err != nil {
		return err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return errors.New(r.Get("errmsg").String())
	}

	return nil
}
