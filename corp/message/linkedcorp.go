package message

import (
	"encoding/json"

	"github.com/chenghonour/gochat/event"
	"github.com/chenghonour/gochat/urls"
	"github.com/chenghonour/gochat/wx"
)

type LinkedcorpExtra struct {
	ToUser  []string
	ToParty []string
	ToTag   []string
	ToAll   int
}

type LinkedcorpMsg struct {
	ToUser      []string      `json:"touser,omitempty"`
	ToParty     []string      `json:"toparty,omitempty"`
	ToTag       []string      `json:"totag,omitempty"`
	ToAll       int           `json:"toall,omitempty"`
	MsgType     event.MsgType `json:"msgtype"`
	AgentID     int64         `json:"agentid,omitempty"`
	Text        *Text         `json:"text,omitempty"`
	Image       *Media        `json:"image,omitempty"`
	Voice       *Media        `json:"voice,omitempty"`
	Video       *Video        `json:"video,omitempty"`
	File        *Media        `json:"file,omitempty"`
	TextCard    *TextCard     `json:"textcard,omitempty"`
	News        *News         `json:"news,omitempty"`
	MPNews      *MPNews       `json:"mpnews,omitempty"`
	Markdown    *Text         `json:"markdown,omitempty"`
	MinipNotice *MinipNotice  `json:"miniprogram_notice,omitempty"`
}

type ResultLinkedcorpSend struct {
	InvalidUser  []string `json:"invaliduser"`
	InvalidParty []string `json:"invalidparty"`
	InvalidTag   []string `json:"invalidtag"`
}

// SendLinkedcorpText 发送企业互联消息（文本消息）
func SendLinkedcorpText(agentID int64, content string, extra *LinkedcorpExtra, result *ResultLinkedcorpSend) wx.Action {
	msg := &LinkedcorpMsg{
		MsgType: event.MsgText,
		AgentID: agentID,
		Text: &Text{
			Content: content,
		},
	}

	if extra != nil {
		msg.ToUser = extra.ToUser
		msg.ToParty = extra.ToParty
		msg.ToTag = extra.ToTag
		msg.ToAll = extra.ToAll
	}

	return wx.NewPostAction(urls.CorpLinkedcorpMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// SendLinkedcorpImage 发送企业互联消息（图片消息）
func SendLinkedcorpImage(agentID int64, mediaID string, extra *LinkedcorpExtra, result *ResultLinkedcorpSend) wx.Action {
	msg := &LinkedcorpMsg{
		MsgType: event.MsgImage,
		AgentID: agentID,
		Image: &Media{
			MediaID: mediaID,
		},
	}

	if extra != nil {
		msg.ToUser = extra.ToUser
		msg.ToParty = extra.ToParty
		msg.ToTag = extra.ToTag
		msg.ToAll = extra.ToAll
	}

	return wx.NewPostAction(urls.CorpLinkedcorpMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// SendLinkedcorpVoice 发送企业互联消息（语音消息）
func SendLinkedcorpVoice(agentID int64, mediaID string, extra *LinkedcorpExtra, result *ResultLinkedcorpSend) wx.Action {
	msg := &LinkedcorpMsg{
		MsgType: event.MsgVoice,
		AgentID: agentID,
		Voice: &Media{
			MediaID: mediaID,
		},
	}

	if extra != nil {
		msg.ToUser = extra.ToUser
		msg.ToParty = extra.ToParty
		msg.ToTag = extra.ToTag
		msg.ToAll = extra.ToAll
	}

	return wx.NewPostAction(urls.CorpLinkedcorpMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// SendLinkedcorpVideo 发送企业互联消息（视频消息）
func SendLinkedcorpVideo(agentID int64, video *Video, extra *LinkedcorpExtra, result *ResultLinkedcorpSend) wx.Action {
	msg := &LinkedcorpMsg{
		MsgType: event.MsgVideo,
		AgentID: agentID,
		Video:   video,
	}

	if extra != nil {
		msg.ToUser = extra.ToUser
		msg.ToParty = extra.ToParty
		msg.ToTag = extra.ToTag
		msg.ToAll = extra.ToAll
	}

	return wx.NewPostAction(urls.CorpLinkedcorpMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// SendLinkedcorpFile 发送企业互联消息（文件消息）
func SendLinkedcorpFile(agentID int64, mediaID string, extra *LinkedcorpExtra, result *ResultLinkedcorpSend) wx.Action {
	msg := &LinkedcorpMsg{
		MsgType: event.MsgFile,
		AgentID: agentID,
		File: &Media{
			MediaID: mediaID,
		},
	}

	if extra != nil {
		msg.ToUser = extra.ToUser
		msg.ToParty = extra.ToParty
		msg.ToTag = extra.ToTag
		msg.ToAll = extra.ToAll
	}

	return wx.NewPostAction(urls.CorpLinkedcorpMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// SendLinkedcorpTextCard 发送企业互联消息（文本卡片消息）
func SendLinkedcorpTextCard(agentID int64, card *TextCard, extra *LinkedcorpExtra, result *ResultLinkedcorpSend) wx.Action {
	msg := &LinkedcorpMsg{
		MsgType:  event.MsgTextCard,
		AgentID:  agentID,
		TextCard: card,
	}

	if extra != nil {
		msg.ToUser = extra.ToUser
		msg.ToParty = extra.ToParty
		msg.ToTag = extra.ToTag
		msg.ToAll = extra.ToAll
	}

	return wx.NewPostAction(urls.CorpLinkedcorpMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// SendLinkedcorpNews 发送企业互联消息（图文消息）
func SendLinkedcorpNews(agentID int64, articles []*NewsArticle, extra *LinkedcorpExtra, result *ResultLinkedcorpSend) wx.Action {
	msg := &LinkedcorpMsg{
		MsgType: event.MsgNews,
		AgentID: agentID,
		News: &News{
			Articles: articles,
		},
	}

	if extra != nil {
		msg.ToUser = extra.ToUser
		msg.ToParty = extra.ToParty
		msg.ToTag = extra.ToTag
		msg.ToAll = extra.ToAll
	}

	return wx.NewPostAction(urls.CorpLinkedcorpMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// SendLinkedcorpMPNews 发送企业互联消息（图文消息 - mpnews）
func SendLinkedcorpMPNews(agentID int64, articles []*MPNewsArticle, extra *LinkedcorpExtra, result *ResultLinkedcorpSend) wx.Action {
	msg := &LinkedcorpMsg{
		MsgType: event.MsgMPNews,
		AgentID: agentID,
		MPNews: &MPNews{
			Articles: articles,
		},
	}

	if extra != nil {
		msg.ToUser = extra.ToUser
		msg.ToParty = extra.ToParty
		msg.ToTag = extra.ToTag
		msg.ToAll = extra.ToAll
	}

	return wx.NewPostAction(urls.CorpLinkedcorpMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// SendLinkedcorpMarkdown 发送企业互联消息（markdown消息）
func SendLinkedcorpMarkdown(agentID int64, content string, extra *LinkedcorpExtra, result *ResultLinkedcorpSend) wx.Action {
	msg := &LinkedcorpMsg{
		MsgType: event.MsgMarkdown,
		AgentID: agentID,
		Markdown: &Text{
			Content: content,
		},
	}

	if extra != nil {
		msg.ToUser = extra.ToUser
		msg.ToParty = extra.ToParty
		msg.ToTag = extra.ToTag
		msg.ToAll = extra.ToAll
	}

	return wx.NewPostAction(urls.CorpLinkedcorpMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// SendLinkedcorpMinipNotice 发送企业互联消息（小程序通知消息）
func SendLinkedcorpMinipNotice(notice *MinipNotice, extra *LinkedcorpExtra, result *ResultLinkedcorpSend) wx.Action {
	msg := &LinkedcorpMsg{
		MsgType:     event.MsgMinipNotice,
		MinipNotice: notice,
	}

	if extra != nil {
		msg.ToUser = extra.ToUser
		msg.ToParty = extra.ToParty
		msg.ToTag = extra.ToTag
		msg.ToAll = extra.ToAll
	}

	return wx.NewPostAction(urls.CorpLinkedcorpMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
