package oa

import (
	"encoding/json"

	"github.com/tidwall/gjson"

	"github.com/shenghui0779/gochat/wx"
)

// TemplateInfo 模板信息
type TemplateInfo struct {
	TemplateID      string `json:"template_id"`      // 模板ID
	Title           string `json:"title"`            // 模板标题
	PrimaryIndustry string `json:"primary_industry"` // 模板所属行业的一级行业
	DeputyIndustry  string `json:"deputy_industry"`  // 模板所属行业的二级行业
	Content         string `json:"content"`          // 模板内容
	Example         string `json:"example"`          // 模板示例
}

// GetTemplateList 获取模板列表
func GetTemplateList(dest *[]*TemplateInfo) wx.Action {
	return wx.NewAction(TemplateListURL,
		wx.WithMethod(wx.MethodGet),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal([]byte(gjson.GetBytes(resp, "template_list").Raw), dest)
		}),
	)
}

// DeleteTemplate 删除模板
func DeleteTemplate(templateID string) wx.Action {
	return wx.NewAction(TemplateDeleteURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(wx.X{"template_id": templateID})
		}),
	)
}

// MessageBody 消息内容体
type MessageBody map[string]map[string]string

// MessageMinip 跳转小程序
type MessageMinip struct {
	AppID    string `json:"appid"`    // 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系，暂不支持小游戏）
	Pagepath string `json:"pagepath"` // 所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar），要求该小程序已发布，暂不支持小游戏
}

// TemplateMessage 公众号模板消息
type TemplateMessage struct {
	TemplateID  string        // 模板ID
	URL         string        // 模板跳转链接（海外帐号没有跳转能力）
	MiniProgram *MessageMinip // 跳转小程序
	Data        MessageBody   // 模板内容，格式形如：{"key1":{"value":"V","color":"#"},"key2":{"value": "V","color":"#"}}
}

// SendTemplateMessage 发送模板消息
func SendTemplateMessage(openID string, msg *TemplateMessage) wx.Action {
	return wx.NewAction(TemplateMessageSendURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			params := wx.X{
				"touser":      openID,
				"template_id": msg.TemplateID,
				"data":        msg.Data,
			}

			if msg.URL != "" {
				params["url"] = msg.URL
			}

			if msg.MiniProgram != nil {
				params["miniprogram"] = msg.MiniProgram
			}

			return json.Marshal(params)
		}),
	)
}

// SendSubscribeMessage 发送一次性订阅消息
func SendSubscribeMessage(openID, scene, title string, msg *TemplateMessage) wx.Action {
	return wx.NewAction(SubscribeMessageSendURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			params := wx.X{
				"scene":       scene,
				"title":       title,
				"touser":      openID,
				"template_id": msg.TemplateID,
				"data":        msg.Data,
			}

			if msg.URL != "" {
				params["url"] = msg.URL
			}

			if msg.MiniProgram != nil {
				params["miniprogram"] = msg.MiniProgram
			}

			return json.Marshal(params)
		}),
	)
}

// SendKFTextMessage 发送客服文本消息（支持插入跳小程序的文字链）
func SendKFTextMessage(openID, text string, kfAccount ...string) wx.Action {
	return wx.NewAction(KFMessageSendURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			data := wx.X{
				"touser":  openID,
				"msgtype": "text",
				"text": wx.X{
					"content": text,
				},
			}

			if len(kfAccount) != 0 {
				data["customservice"] = wx.X{
					"kf_account": kfAccount[0],
				}
			}

			return json.Marshal(data)
		}),
	)
}

// SendKFImageMessage 发送客服图片消息（媒体ID，通过素材接口上传获得）
func SendKFImageMessage(openID, mediaID string, kfAccount ...string) wx.Action {
	return wx.NewAction(KFMessageSendURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			data := wx.X{
				"touser":  openID,
				"msgtype": "image",
				"image": wx.X{
					"media_id": mediaID,
				},
			}

			if len(kfAccount) != 0 {
				data["customservice"] = wx.X{
					"kf_account": kfAccount[0],
				}
			}

			return json.Marshal(data)
		}),
	)
}

// SendKFVoiceMessage 发送客服语音消息（媒体ID，通过素材接口上传获得）
func SendKFVoiceMessage(openID, mediaID string, kfAccount ...string) wx.Action {
	return wx.NewAction(KFMessageSendURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			data := wx.X{
				"touser":  openID,
				"msgtype": "voice",
				"voice": wx.X{
					"media_id": mediaID,
				},
			}

			if len(kfAccount) != 0 {
				data["customservice"] = wx.X{
					"kf_account": kfAccount[0],
				}
			}

			return json.Marshal(data)
		}),
	)
}

// KFVideoMessage 客服视频消息
type KFVideoMessage struct {
	MediaID      string `json:"media_id"`       // 视频消息（点击跳转到图文消息页）的媒体ID
	ThumbMediaID string `json:"thumb_media_id"` // 缩略图的媒体ID
	Title        string `json:"title"`          // 视频消息的标题
	Description  string `json:"description"`    // 视频消息的描述
}

// SendKFVideoMessage 发送客服视频消息（媒体ID，通过素材接口上传获得）
func SendKFVideoMessage(openID string, msg *KFVideoMessage, kfAccount ...string) wx.Action {
	return wx.NewAction(KFMessageSendURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			data := wx.X{
				"touser":  openID,
				"msgtype": "video",
				"video":   msg,
			}

			if len(kfAccount) != 0 {
				data["customservice"] = wx.X{
					"kf_account": kfAccount[0],
				}
			}

			return json.Marshal(data)
		}),
	)
}

// KFMusicMessage 客服音乐消息
type KFMusicMessage struct {
	Title        string `json:"title"`          // 音乐消息的标题
	Description  string `json:"description"`    // 音乐消息的描述
	MusicURL     string `json:"musicurl"`       // 音乐链接
	HQMusicURL   string `json:"hqmusicurl"`     // 高品质音乐链接，wifi环境优先使用该链接播放音乐
	ThumbMediaID string `json:"thumb_media_id"` // 缩略图的媒体ID
}

// SendKFMusicMessage 发送客服音乐消息
func SendKFMusicMessage(openID string, msg *KFMusicMessage, kfAccount ...string) wx.Action {
	return wx.NewAction(KFMessageSendURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			data := wx.X{
				"touser":  openID,
				"msgtype": "music",
				"music":   msg,
			}

			if len(kfAccount) != 0 {
				data["customservice"] = wx.X{
					"kf_account": kfAccount[0],
				}
			}

			return json.Marshal(data)
		}),
	)
}

// KFArticle 客服图文
type KFArticle struct {
	Title       string `json:"title"`       // 图文消息的标题
	Description string `json:"description"` // 图文消息的描述
	URL         string `json:"url"`         // 图文消息被点击后跳转的链接
	PicURL      string `json:"picurl"`      // 图文消息的图片链接，支持JPG、PNG格式，较好的效果为大图640*320，小图80*80
}

// SendKFNewsMessage 发送客服图文消息（点击跳转到外链；图文消息条数限制在1条以内，注意，如果图文数超过1，则将会返回错误码45008）
func SendKFNewsMessage(openID string, articles []*KFArticle, kfAccount ...string) wx.Action {
	return wx.NewAction(KFMessageSendURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			data := wx.X{
				"touser":  openID,
				"msgtype": "news",
				"news": wx.X{
					"articles": articles,
				},
			}

			if len(kfAccount) != 0 {
				data["customservice"] = wx.X{
					"kf_account": kfAccount[0],
				}
			}

			return json.Marshal(data)
		}),
	)
}

// SendKFMPNewsMessage 发送图文消息（点击跳转到图文消息页面；图文消息条数限制在1条以内，注意，如果图文数超过1，则将会返回错误码45008）
func SendKFMPNewsMessage(openID, mediaID string, kfAccount ...string) wx.Action {
	return wx.NewAction(KFMessageSendURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			data := wx.X{
				"touser":  openID,
				"msgtype": "mpnews",
				"mpnews": wx.X{
					"media_id": mediaID,
				},
			}

			if len(kfAccount) != 0 {
				data["customservice"] = wx.X{
					"kf_account": kfAccount[0],
				}
			}

			return json.Marshal(data)
		}),
	)
}

// KFMenuMessage 客服菜单消息
type KFMenuMessage struct {
	HeadContent string          `json:"head_content"`
	TailContent string          `json:"tail_content"`
	List        []*KFMenuOption `json:"list"`
}

// KFMenuOption 客服菜单选项
type KFMenuOption struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

// SendKFMenuMessage 发送客服菜单消息
func SendKFMenuMessage(openID string, msg *KFMenuMessage, kfAccount ...string) wx.Action {
	return wx.NewAction(KFMessageSendURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			data := wx.X{
				"touser":  openID,
				"msgtype": "msgmenu",
				"msgmenu": msg,
			}

			if len(kfAccount) != 0 {
				data["customservice"] = wx.X{
					"kf_account": kfAccount[0],
				}
			}

			return json.Marshal(data)
		}),
	)
}

// SendKFCardMessage 发送客服卡券消息（特别注意：客服消息接口投放卡券仅支持非自定义Code码和导入code模式的卡券的卡券）
func SendKFCardMessage(openID, cardID string, kfAccount ...string) wx.Action {
	return wx.NewAction(KFMessageSendURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			data := wx.X{
				"touser":  openID,
				"msgtype": "wxcard",
				"wxcard": wx.X{
					"card_id": cardID,
				},
			}

			if len(kfAccount) != 0 {
				data["customservice"] = wx.X{
					"kf_account": kfAccount[0],
				}
			}

			return json.Marshal(data)
		}),
	)
}

// KFMinipMessage 客服小程序卡片消息
type KFMinipMessage struct {
	Title        string `json:"title"`          // 消息标题
	AppID        string `json:"appid"`          // 小程序的appid，要求小程序的appid需要与公众号有关联关系
	Pagepath     string `json:"pagepath"`       // 小程序的页面路径，跟app.json对齐，支持参数，比如pages/index/index?foo=bar
	ThumbMediaID string `json:"thumb_media_id"` // 小程序卡片图片的媒体ID，小程序卡片图片建议大小为520*416
}

// SendKFMinipMessage 发送客服小程序卡片消息
func SendKFMinipMessage(openID string, msg *KFMinipMessage, kfAccount ...string) wx.Action {
	return wx.NewAction(KFMessageSendURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			data := wx.X{
				"touser":          openID,
				"msgtype":         "miniprogrampage",
				"miniprogrampage": msg,
			}

			if len(kfAccount) != 0 {
				data["customservice"] = wx.X{
					"kf_account": kfAccount[0],
				}
			}

			return json.Marshal(data)
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
