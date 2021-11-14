package offia

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

// KFMsgType 客服消息类型
type KFMsgType string

const (
	KFMsgText          KFMsgType = "text"            // 文本消息
	KFMsgImage         KFMsgType = "image"           // 图片消息
	KFMsgVoice         KFMsgType = "voice"           // 语音消息
	KFMsgVideo         KFMsgType = "video"           // 视频消息
	KFMsgMusic         KFMsgType = "music"           // 音乐消息
	KFMsgNews          KFMsgType = "news"            // 图文消息
	KFMsgMPNews        KFMsgType = "mpnews"          // 图文消息
	KFMsgMPNewsArticle KFMsgType = "mpnewsarticle"   // 图文消息
	KFMsgMenu          KFMsgType = "msgmenu"         // 菜单消息
	KFMsgCard          KFMsgType = "wxcard"          // 卡券
	KFMsgMinipPage     KFMsgType = "miniprogrampage" // 小程序卡片
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

type KFText struct {
	Content string `json:"content"`
}

type KFMedia struct {
	MediaID string `json:"media_id"`
}

type KFVideo struct {
	MediaID      string `json:"media_id"`       // 视频消息（点击跳转到图文消息页）的媒体ID
	ThumbMediaID string `json:"thumb_media_id"` // 缩略图的媒体ID
	Title        string `json:"title"`          // 视频消息的标题
	Description  string `json:"description"`    // 视频消息的描述
}

type KFMusic struct {
	Title        string `json:"title"`          // 音乐消息的标题
	Description  string `json:"description"`    // 音乐消息的描述
	MusicURL     string `json:"musicurl"`       // 音乐链接
	HQMusicURL   string `json:"hqmusicurl"`     // 高品质音乐链接，wifi环境优先使用该链接播放音乐
	ThumbMediaID string `json:"thumb_media_id"` // 缩略图的媒体ID
}

type KFArticle struct {
	Title       string `json:"title"`       // 图文消息的标题
	Description string `json:"description"` // 图文消息的描述
	URL         string `json:"url"`         // 图文消息被点击后跳转的链接
	PicURL      string `json:"picurl"`      // 图文消息的图片链接，支持JPG、PNG格式，较好的效果为大图640*320，小图80*80
}

type KFMPNewsArticle struct {
	ArticleID string `json:"article_id"`
}

type KFMenuOption struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

type KFMenu struct {
	HeadContent string          `json:"head_content"`
	TailContent string          `json:"tail_content"`
	List        []*KFMenuOption `json:"list"`
}

type KFCard struct {
	CardID string `json:"card_id"`
}

type KFMinipPage struct {
	Title        string `json:"title"`          // 消息标题
	AppID        string `json:"appid"`          // 小程序的appid，要求小程序的appid需要与公众号有关联关系
	Pagepath     string `json:"pagepath"`       // 小程序的页面路径，跟app.json对齐，支持参数，比如pages/index/index?foo=bar
	ThumbMediaID string `json:"thumb_media_id"` // 小程序卡片图片的媒体ID，小程序卡片图片建议大小为520*416
}

type MsgKF struct {
	KFAccount string `json:"kf_account"`
}

// ParamsKFMsg 客服消息参数
type ParamsKFMsg struct {
	ToUser        string           `json:"touser"`
	MsgType       KFMsgType        `json:"msgtype"`
	Text          *KFText          `json:"text,omitempty"`
	Image         *KFMedia         `json:"image,omitempty"`
	Voice         *KFMedia         `json:"voice,omitempty"`
	Video         *KFVideo         `json:"video,omitempty"`
	Music         *KFMusic         `json:"music,omitempty"`
	News          []*KFArticle     `json:"news,omitempty"`
	MPNews        *KFMedia         `json:"mpnews,omitempty"`
	MPNewsArticle *KFMPNewsArticle `json:"mpnewsarticle,omitempty"`
	Menu          *KFMenu          `json:"msgmenu,omitempty"`
	Card          *KFCard          `json:"wxcard,omitempty"`
	MinipPage     *KFMinipPage     `json:"miniprogrampage,omitempty"`
	CustomService *MsgKF           `json:"customservice,omitempty"`
}

// SendKFTextMsg 发送客服文本消息（支持插入跳小程序的文字链）
func SendKFTextMsg(openID string, text *KFText, kfAccount ...string) wx.Action {
	params := &ParamsKFMsg{
		ToUser:  openID,
		MsgType: KFMsgText,
		Text:    text,
	}

	if len(kfAccount) > 0 {
		params.CustomService = &MsgKF{
			KFAccount: kfAccount[0],
		}
	}

	return wx.NewPostAction(urls.OffiaKFMsgSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

// SendKFImageMsg 发送客服图片消息（媒体ID，通过素材接口上传获得）
func SendKFImageMsg(openID string, image *KFMedia, kfAccount ...string) wx.Action {
	params := &ParamsKFMsg{
		ToUser:  openID,
		MsgType: KFMsgImage,
		Image:   image,
	}

	if len(kfAccount) > 0 {
		params.CustomService = &MsgKF{
			KFAccount: kfAccount[0],
		}
	}

	return wx.NewPostAction(urls.OffiaKFMsgSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

// SendKFVoiceMsg 发送客服语音消息（媒体ID，通过素材接口上传获得）
func SendKFVoiceMsg(openID string, voice *KFMedia, kfAccount ...string) wx.Action {
	params := &ParamsKFMsg{
		ToUser:  openID,
		MsgType: KFMsgVoice,
		Voice:   voice,
	}

	if len(kfAccount) > 0 {
		params.CustomService = &MsgKF{
			KFAccount: kfAccount[0],
		}
	}

	return wx.NewPostAction(urls.OffiaKFMsgSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

// SendKFVideoMsg 发送客服视频消息（媒体ID，通过素材接口上传获得）
func SendKFVideoMsg(openID string, video *KFVideo, kfAccount ...string) wx.Action {
	params := &ParamsKFMsg{
		ToUser:  openID,
		MsgType: KFMsgVideo,
		Video:   video,
	}

	if len(kfAccount) > 0 {
		params.CustomService = &MsgKF{
			KFAccount: kfAccount[0],
		}
	}

	return wx.NewPostAction(urls.OffiaKFMsgSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

// SendKFMusicMsg 发送客服音乐消息
func SendKFMusicMsg(openID string, music *KFMusic, kfAccount ...string) wx.Action {
	params := &ParamsKFMsg{
		ToUser:  openID,
		MsgType: KFMsgMusic,
		Music:   music,
	}

	if len(kfAccount) > 0 {
		params.CustomService = &MsgKF{
			KFAccount: kfAccount[0],
		}
	}

	return wx.NewPostAction(urls.OffiaKFMsgSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

// SendKFNewsMsg 发送客服图文消息（点击跳转到外链；图文消息条数限制在1条以内，注意，如果图文数超过1，则将会返回错误码45008）
func SendKFNewsMsg(openID string, news []*KFArticle, kfAccount ...string) wx.Action {
	params := &ParamsKFMsg{
		ToUser:  openID,
		MsgType: KFMsgNews,
		News:    news,
	}

	if len(kfAccount) > 0 {
		params.CustomService = &MsgKF{
			KFAccount: kfAccount[0],
		}
	}

	return wx.NewPostAction(urls.OffiaKFMsgSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

// SendKFMPNewsMsg 发送图文消息（点击跳转到图文消息页面；图文消息条数限制在1条以内，注意，如果图文数超过1，则将会返回错误码45008）
func SendKFMPNewsMsg(openID string, news *KFMedia, kfAccount ...string) wx.Action {
	params := &ParamsKFMsg{
		ToUser:  openID,
		MsgType: KFMsgMPNews,
		MPNews:  news,
	}

	if len(kfAccount) > 0 {
		params.CustomService = &MsgKF{
			KFAccount: kfAccount[0],
		}
	}

	return wx.NewPostAction(urls.OffiaKFMsgSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

// SendKFMPNewsArticleMsg 发送图文消息（点击跳转到图文消息页面）使用通过 “发布” 系列接口得到的 article_id
// 注意: 草稿接口灰度完成后，将不再支持此前客服接口中带 media_id 的 mpnews 类型的图文消息
func SendKFMPNewsArticleMsg(openID string, article *KFMPNewsArticle, kfAccount ...string) wx.Action {
	params := &ParamsKFMsg{
		ToUser:        openID,
		MsgType:       KFMsgMPNewsArticle,
		MPNewsArticle: article,
	}

	if len(kfAccount) > 0 {
		params.CustomService = &MsgKF{
			KFAccount: kfAccount[0],
		}
	}

	return wx.NewPostAction(urls.OffiaKFMsgSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

// SendKFMenuMsg 发送客服菜单消息
func SendKFMenuMsg(openID string, menu *KFMenu, kfAccount ...string) wx.Action {
	params := &ParamsKFMsg{
		ToUser:  openID,
		MsgType: KFMsgMenu,
		Menu:    menu,
	}

	if len(kfAccount) > 0 {
		params.CustomService = &MsgKF{
			KFAccount: kfAccount[0],
		}
	}

	return wx.NewPostAction(urls.OffiaKFMsgSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

// SendKFCardMsg 发送客服卡券消息（特别注意：客服消息接口投放卡券仅支持非自定义Code码和导入code模式的卡券的卡券）
func SendKFCardMsg(openID string, card *KFCard, kfAccount ...string) wx.Action {
	params := &ParamsKFMsg{
		ToUser:  openID,
		MsgType: KFMsgCard,
		Card:    card,
	}

	if len(kfAccount) > 0 {
		params.CustomService = &MsgKF{
			KFAccount: kfAccount[0],
		}
	}

	return wx.NewPostAction(urls.OffiaKFMsgSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

// SendKFMinipMsg 发送客服小程序卡片消息
func SendKFMinipMsg(openID string, minipPage *KFMinipPage, kfAccount ...string) wx.Action {
	params := &ParamsKFMsg{
		ToUser:    openID,
		MsgType:   KFMsgMinipPage,
		MinipPage: minipPage,
	}

	if len(kfAccount) > 0 {
		params.CustomService = &MsgKF{
			KFAccount: kfAccount[0],
		}
	}

	return wx.NewPostAction(urls.OffiaKFMsgSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

// TypeCmd 输入状态命令
type TypeCmd string

// 微信支持的输入状态命令
const (
	Typing       TypeCmd = "Typing"       // 正在输入
	CancelTyping TypeCmd = "CancelTyping" // 取消输入
)

type ParamsKFTyping struct {
	ToUser  string  `json:"touser"`
	Command TypeCmd `json:"command"`
}

// SendKFTyping 下发当前输入状态（仅支持客服消息）
func SendKFTyping(openID string, cmd TypeCmd) wx.Action {
	params := &ParamsKFTyping{
		ToUser:  openID,
		Command: cmd,
	}

	return wx.NewPostAction(urls.OffiaSetTyping,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}
