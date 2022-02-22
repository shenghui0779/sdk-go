package kf

import (
	"github.com/shenghui0779/gochat/event"
	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

// MsgText 消息文本
type MsgText struct {
	Content string `json:"content"`
}

// MsgMedia 消息媒体（图片等）
type MsgMedia struct {
	MediaID string `json:"media_id"`
}

// MsgVideo 消息视频
type MsgVideo struct {
	MediaID      string `json:"media_id"`       // 视频消息（点击跳转到图文消息页）的媒体ID
	ThumbMediaID string `json:"thumb_media_id"` // 缩略图的媒体ID
	Title        string `json:"title"`          // 视频消息的标题
	Description  string `json:"description"`    // 视频消息的描述
}

// MsgMusic 消息音乐
type MsgMusic struct {
	Title        string `json:"title"`          // 音乐消息的标题
	Description  string `json:"description"`    // 音乐消息的描述
	MusicURL     string `json:"musicurl"`       // 音乐链接
	HQMusicURL   string `json:"hqmusicurl"`     // 高品质音乐链接，wifi环境优先使用该链接播放音乐
	ThumbMediaID string `json:"thumb_media_id"` // 缩略图的媒体ID
}

// MsgArticle 消息图文
type MsgArticle struct {
	Title       string `json:"title"`       // 图文消息的标题
	Description string `json:"description"` // 图文消息的描述
	URL         string `json:"url"`         // 图文消息被点击后跳转的链接
	PicURL      string `json:"picurl"`      // 图文消息的图片链接，支持JPG、PNG格式，较好的效果为大图640*320，小图80*80
}

// MsgNews 消息图文
type MsgNews struct {
	Articles []*MsgArticle `json:"articles"`
}

// MsgMPNewsArticle 消息图文
type MsgMPNewsArticle struct {
	ArticleID string `json:"article_id"`
}

// MenuOption 消息菜单选项
type MenuOption struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

// MsgMenu 消息菜单
type MsgMenu struct {
	HeadContent string        `json:"head_content"`
	TailContent string        `json:"tail_content"`
	List        []*MenuOption `json:"list"`
}

// MsgCard 消息卡券
type MsgCard struct {
	CardID string `json:"card_id"`
}

// MsgMinipPage 小程序卡片
type MsgMinipPage struct {
	Title        string `json:"title"`          // 消息标题
	AppID        string `json:"appid"`          // 小程序的appid，要求小程序的appid需要与公众号有关联关系
	Pagepath     string `json:"pagepath"`       // 小程序的页面路径，跟app.json对齐，支持参数，比如pages/index/index?foo=bar
	ThumbMediaID string `json:"thumb_media_id"` // 小程序卡片图片的媒体ID，小程序卡片图片建议大小为520*416
}

// MsgKF 消息客服
type MsgKF struct {
	KFAccount string `json:"kf_account"`
}

// ParamsMessage 客服消息参数
type ParamsMessage struct {
	ToUser        string            `json:"touser"`
	MsgType       event.MsgType     `json:"msgtype"`
	Text          *MsgText          `json:"text,omitempty"`
	Image         *MsgMedia         `json:"image,omitempty"`
	Voice         *MsgMedia         `json:"voice,omitempty"`
	Video         *MsgVideo         `json:"video,omitempty"`
	Music         *MsgMusic         `json:"music,omitempty"`
	News          *MsgNews          `json:"news,omitempty"`
	MPNews        *MsgMedia         `json:"mpnews,omitempty"`
	MPNewsArticle *MsgMPNewsArticle `json:"mpnewsarticle,omitempty"`
	Menu          *MsgMenu          `json:"msgmenu,omitempty"`
	Card          *MsgCard          `json:"wxcard,omitempty"`
	MinipPage     *MsgMinipPage     `json:"miniprogrampage,omitempty"`
	CustomService *MsgKF            `json:"customservice,omitempty"`
}

// SendTextMsg 发送客服文本消息（支持插入跳小程序的文字链）
func SendTextMsg(openID, content string, kfAccount ...string) wx.Action {
	params := &ParamsMessage{
		ToUser:  openID,
		MsgType: event.MsgText,
		Text: &MsgText{
			Content: content,
		},
	}

	if len(kfAccount) > 0 {
		params.CustomService = &MsgKF{
			KFAccount: kfAccount[0],
		}
	}

	return wx.NewPostAction(urls.OffiaKFMsgSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

// SendImageMsg 发送客服图片消息（媒体ID，通过素材接口上传获得）
func SendImageMsg(openID, mediaID string, kfAccount ...string) wx.Action {
	params := &ParamsMessage{
		ToUser:  openID,
		MsgType: event.MsgImage,
		Image: &MsgMedia{
			MediaID: mediaID,
		},
	}

	if len(kfAccount) > 0 {
		params.CustomService = &MsgKF{
			KFAccount: kfAccount[0],
		}
	}

	return wx.NewPostAction(urls.OffiaKFMsgSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

// SendVoiceMsg 发送客服语音消息（媒体ID，通过素材接口上传获得）
func SendVoiceMsg(openID, mediaID string, kfAccount ...string) wx.Action {
	params := &ParamsMessage{
		ToUser:  openID,
		MsgType: event.MsgVoice,
		Voice: &MsgMedia{
			MediaID: mediaID,
		},
	}

	if len(kfAccount) > 0 {
		params.CustomService = &MsgKF{
			KFAccount: kfAccount[0],
		}
	}

	return wx.NewPostAction(urls.OffiaKFMsgSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

// SendVideoMsg 发送客服视频消息（媒体ID，通过素材接口上传获得）
func SendVideoMsg(openID string, video *MsgVideo, kfAccount ...string) wx.Action {
	params := &ParamsMessage{
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
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

// SendMusicMsg 发送客服音乐消息
func SendMusicMsg(openID string, music *MsgMusic, kfAccount ...string) wx.Action {
	params := &ParamsMessage{
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
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

// SendNewsMsg 发送客服图文消息（点击跳转到外链；图文消息条数限制在1条以内，注意，如果图文数超过1，则将会返回错误码45008）
func SendNewsMsg(openID string, articles []*MsgArticle, kfAccount ...string) wx.Action {
	params := &ParamsMessage{
		ToUser:  openID,
		MsgType: event.MsgNews,
		News: &MsgNews{
			Articles: articles,
		},
	}

	if len(kfAccount) > 0 {
		params.CustomService = &MsgKF{
			KFAccount: kfAccount[0],
		}
	}

	return wx.NewPostAction(urls.OffiaKFMsgSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

// SendMPNewsMsg 发送图文消息（点击跳转到图文消息页面；图文消息条数限制在1条以内，注意，如果图文数超过1，则将会返回错误码45008）
func SendMPNewsMsg(openID, mediaID string, kfAccount ...string) wx.Action {
	params := &ParamsMessage{
		ToUser:  openID,
		MsgType: event.MsgMPNews,
		MPNews: &MsgMedia{
			MediaID: mediaID,
		},
	}

	if len(kfAccount) > 0 {
		params.CustomService = &MsgKF{
			KFAccount: kfAccount[0],
		}
	}

	return wx.NewPostAction(urls.OffiaKFMsgSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

// SendMPNewsArticleMsg 发送图文消息（点击跳转到图文消息页面）使用通过 “发布” 系列接口得到的 article_id
// 注意: 草稿接口灰度完成后，将不再支持此前客服接口中带 media_id 的 mpnews 类型的图文消息
func SendMPNewsArticleMsg(openID, articleID string, kfAccount ...string) wx.Action {
	params := &ParamsMessage{
		ToUser:  openID,
		MsgType: event.MsgMPNewsArticle,
		MPNewsArticle: &MsgMPNewsArticle{
			ArticleID: articleID,
		},
	}

	if len(kfAccount) > 0 {
		params.CustomService = &MsgKF{
			KFAccount: kfAccount[0],
		}
	}

	return wx.NewPostAction(urls.OffiaKFMsgSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

// SendMenuMsg 发送客服菜单消息
func SendMenuMsg(openID string, menu *MsgMenu, kfAccount ...string) wx.Action {
	params := &ParamsMessage{
		ToUser:  openID,
		MsgType: event.MsgMsgMenu,
		Menu:    menu,
	}

	if len(kfAccount) > 0 {
		params.CustomService = &MsgKF{
			KFAccount: kfAccount[0],
		}
	}

	return wx.NewPostAction(urls.OffiaKFMsgSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

// SendWXCardMsg 发送客服卡券消息（特别注意：客服消息接口投放卡券仅支持非自定义Code码和导入code模式的卡券的卡券）
func SendWXCardMsg(openID, cardID string, kfAccount ...string) wx.Action {
	params := &ParamsMessage{
		ToUser:  openID,
		MsgType: event.MsgCard,
		Card: &MsgCard{
			CardID: cardID,
		},
	}

	if len(kfAccount) > 0 {
		params.CustomService = &MsgKF{
			KFAccount: kfAccount[0],
		}
	}

	return wx.NewPostAction(urls.OffiaKFMsgSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

// SendMinipPageMsg 发送客服小程序卡片消息
func SendMinipPageMsg(openID string, minipPage *MsgMinipPage, kfAccount ...string) wx.Action {
	params := &ParamsMessage{
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
			return wx.MarshalNoEscapeHTML(params)
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
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}
