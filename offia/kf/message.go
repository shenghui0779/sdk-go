package kf

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/event"
	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

// Text 消息文本
type Text struct {
	Content string `json:"content"`
}

// Media 消息媒体（图片等）
type Media struct {
	MediaID string `json:"media_id"`
}

// Video 消息视频
type Video struct {
	MediaID      string `json:"media_id"`       // 视频消息（点击跳转到图文消息页）的媒体ID
	ThumbMediaID string `json:"thumb_media_id"` // 缩略图的媒体ID
	Title        string `json:"title"`          // 视频消息的标题
	Description  string `json:"description"`    // 视频消息的描述
}

// Music 消息音乐
type Music struct {
	Title        string `json:"title"`          // 音乐消息的标题
	Description  string `json:"description"`    // 音乐消息的描述
	MusicURL     string `json:"musicurl"`       // 音乐链接
	HQMusicURL   string `json:"hqmusicurl"`     // 高品质音乐链接，wifi环境优先使用该链接播放音乐
	ThumbMediaID string `json:"thumb_media_id"` // 缩略图的媒体ID
}

// Article 消息图文
type Article struct {
	Title       string `json:"title"`       // 图文消息的标题
	Description string `json:"description"` // 图文消息的描述
	URL         string `json:"url"`         // 图文消息被点击后跳转的链接
	PicURL      string `json:"picurl"`      // 图文消息的图片链接，支持JPG、PNG格式，较好的效果为大图640*320，小图80*80
}

// News 消息图文
type News struct {
	Articles []*Article `json:"articles"`
}

// MPNewsArticle 消息图文
type MPNewsArticle struct {
	ArticleID string `json:"article_id"`
}

// MenuOption 消息菜单选项
type MenuOption struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

// Menu 消息菜单
type Menu struct {
	HeadContent string        `json:"head_content"`
	TailContent string        `json:"tail_content"`
	List        []*MenuOption `json:"list"`
}

// Card 消息卡券
type Card struct {
	CardID string `json:"card_id"`
}

// MinipPage 小程序卡片
type MinipPage struct {
	Title        string `json:"title"`          // 消息标题
	AppID        string `json:"appid"`          // 小程序的appid，要求小程序的appid需要与公众号有关联关系
	Pagepath     string `json:"pagepath"`       // 小程序的页面路径，跟app.json对齐，支持参数，比如pages/index/index?foo=bar
	ThumbMediaID string `json:"thumb_media_id"` // 小程序卡片图片的媒体ID，小程序卡片图片建议大小为520*416
}

// MsgKF 消息客服
type MsgKF struct {
	KFAccount string `json:"kf_account"`
}

// ParamsMsgSend 客服消息参数
type ParamsMsgSend struct {
	ToUser        string         `json:"touser"`
	MsgType       event.MsgType  `json:"msgtype"`
	Text          *Text          `json:"text,omitempty"`
	Image         *Media         `json:"image,omitempty"`
	Voice         *Media         `json:"voice,omitempty"`
	Video         *Video         `json:"video,omitempty"`
	Music         *Music         `json:"music,omitempty"`
	News          *News          `json:"news,omitempty"`
	MPNews        *Media         `json:"mpnews,omitempty"`
	MPNewsArticle *MPNewsArticle `json:"mpnewsarticle,omitempty"`
	Menu          *Menu          `json:"msgmenu,omitempty"`
	Card          *Card          `json:"wxcard,omitempty"`
	MinipPage     *MinipPage     `json:"miniprogrampage,omitempty"`
	CustomService *MsgKF         `json:"customservice,omitempty"`
}

// SendTextMsg 发送客服文本消息（支持插入跳小程序的文字链）
func SendTextMsg(openID string, text *Text, kfAccount ...string) wx.Action {
	params := &ParamsMsgSend{
		ToUser:  openID,
		MsgType: event.MsgText,
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

// SendImageMsg 发送客服图片消息（媒体ID，通过素材接口上传获得）
func SendImageMsg(openID string, image *Media, kfAccount ...string) wx.Action {
	params := &ParamsMsgSend{
		ToUser:  openID,
		MsgType: event.MsgImage,
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

// SendVoiceMsg 发送客服语音消息（媒体ID，通过素材接口上传获得）
func SendVoiceMsg(openID string, voice *Media, kfAccount ...string) wx.Action {
	params := &ParamsMsgSend{
		ToUser:  openID,
		MsgType: event.MsgVoice,
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

// SendVideoMsg 发送客服视频消息（媒体ID，通过素材接口上传获得）
func SendVideoMsg(openID string, video *Video, kfAccount ...string) wx.Action {
	params := &ParamsMsgSend{
		ToUser:  openID,
		MsgType: event.MsgVideo,
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

// SendMusicMsg 发送客服音乐消息
func SendMusicMsg(openID string, music *Music, kfAccount ...string) wx.Action {
	params := &ParamsMsgSend{
		ToUser:  openID,
		MsgType: event.MsgMusic,
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

// SendNewsMsg 发送客服图文消息（点击跳转到外链；图文消息条数限制在1条以内，注意，如果图文数超过1，则将会返回错误码45008）
func SendNewsMsg(openID string, news *News, kfAccount ...string) wx.Action {
	params := &ParamsMsgSend{
		ToUser:  openID,
		MsgType: event.MsgNews,
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

// SendMPNewsMsg 发送图文消息（点击跳转到图文消息页面；图文消息条数限制在1条以内，注意，如果图文数超过1，则将会返回错误码45008）
func SendMPNewsMsg(openID string, news *Media, kfAccount ...string) wx.Action {
	params := &ParamsMsgSend{
		ToUser:  openID,
		MsgType: event.MsgMPNews,
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

// SendMPNewsArticleMsg 发送图文消息（点击跳转到图文消息页面）使用通过 “发布” 系列接口得到的 article_id
// 注意: 草稿接口灰度完成后，将不再支持此前客服接口中带 media_id 的 mpnews 类型的图文消息
func SendMPNewsArticleMsg(openID string, article *MPNewsArticle, kfAccount ...string) wx.Action {
	params := &ParamsMsgSend{
		ToUser:        openID,
		MsgType:       event.MsgMPNewsArticle,
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

// SendMenuMsg 发送客服菜单消息
func SendMenuMsg(openID string, menu *Menu, kfAccount ...string) wx.Action {
	params := &ParamsMsgSend{
		ToUser:  openID,
		MsgType: event.MsgMenu,
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

// SendCardMsg 发送客服卡券消息（特别注意：客服消息接口投放卡券仅支持非自定义Code码和导入code模式的卡券的卡券）
func SendCardMsg(openID string, card *Card, kfAccount ...string) wx.Action {
	params := &ParamsMsgSend{
		ToUser:  openID,
		MsgType: event.MsgCard,
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

// SendMinipMsg 发送客服小程序卡片消息
func SendMinipMsg(openID string, minipPage *MinipPage, kfAccount ...string) wx.Action {
	params := &ParamsMsgSend{
		ToUser:    openID,
		MsgType:   event.MsgMinipPage,
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

type ParamsTyping struct {
	ToUser  string  `json:"touser"`
	Command TypeCmd `json:"command"`
}

// SendTyping 下发当前输入状态（仅支持客服消息）
func SendTyping(openID string, cmd TypeCmd) wx.Action {
	params := &ParamsTyping{
		ToUser:  openID,
		Command: cmd,
	}

	return wx.NewPostAction(urls.OffiaSetTyping,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}
