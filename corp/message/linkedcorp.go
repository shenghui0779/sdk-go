package message

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/event"
	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
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
	AgentID     int64         `json:"agentid"`
	MsgType     event.MsgType `json:"msgtype"`
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

func SendLinkedcorpText(agentID int64, content string, extra *LinkedcorpExtra, result *ResultLinkedcorpSend) wx.Action {
	msg := &LinkedcorpMsg{
		AgentID: agentID,
		MsgType: event.MsgText,
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
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendLinkedcorpImage(agentID int64, mediaID string, extra *LinkedcorpExtra, result *ResultLinkedcorpSend) wx.Action {
	msg := &LinkedcorpMsg{
		AgentID: agentID,
		MsgType: event.MsgImage,
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
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendLinkedcorpVoice(agentID int64, mediaID string, extra *LinkedcorpExtra, result *ResultLinkedcorpSend) wx.Action {
	msg := &LinkedcorpMsg{
		AgentID: agentID,
		MsgType: event.MsgVoice,
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
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendLinkedcorpVideo(agentID int64, video *Video, extra *LinkedcorpExtra, result *ResultLinkedcorpSend) wx.Action {
	msg := &LinkedcorpMsg{
		AgentID: agentID,
		MsgType: event.MsgVideo,
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
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendLinkedcorpFile(agentID int64, mediaID string, extra *LinkedcorpExtra, result *ResultLinkedcorpSend) wx.Action {
	msg := &LinkedcorpMsg{
		AgentID: agentID,
		MsgType: event.MsgFile,
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
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendLinkedcorpTextCard(agentID int64, card *TextCard, extra *LinkedcorpExtra, result *ResultLinkedcorpSend) wx.Action {
	msg := &LinkedcorpMsg{
		AgentID:  agentID,
		MsgType:  event.MsgTextCard,
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
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendLinkedcorpNews(agentID int64, articles []*NewsArticle, extra *LinkedcorpExtra, result *ResultLinkedcorpSend) wx.Action {
	msg := &LinkedcorpMsg{
		AgentID: agentID,
		MsgType: event.MsgNews,
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
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendLinkedcorpMPNews(agentID int64, articles []*MPNewsArticle, extra *LinkedcorpExtra, result *ResultLinkedcorpSend) wx.Action {
	msg := &LinkedcorpMsg{
		AgentID: agentID,
		MsgType: event.MsgMPNews,
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
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendLinkedcorpMarkdown(agentID int64, content string, extra *LinkedcorpExtra, result *ResultLinkedcorpSend) wx.Action {
	msg := &LinkedcorpMsg{
		AgentID: agentID,
		MsgType: event.MsgMarkdown,
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
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendLinkedcorpMinipNotice(agentID int64, notice *MinipNotice, extra *LinkedcorpExtra, result *ResultLinkedcorpSend) wx.Action {
	msg := &LinkedcorpMsg{
		AgentID:     agentID,
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
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
