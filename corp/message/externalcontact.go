package message

import (
	"encoding/json"

	"github.com/chenghonour/gochat/event"
	"github.com/chenghonour/gochat/urls"
	"github.com/chenghonour/gochat/wx"
)

type ExternalContactExtra struct {
	ToExternalUser         []string
	ToParentUserID         []string
	ToStudentUserID        []string
	ToParty                []string
	ToAll                  int
	EnableIDTrans          int
	EnableDuplicateCheck   int
	DuplicateCheckInterval int
}

type ExternalContactMsg struct {
	ToExternalUser         []string      `json:"to_external_user,omitempty"`
	ToParentUserID         []string      `json:"to_parent_userid,omitempty"`
	ToStudentUserID        []string      `json:"to_student_userid,omitempty"`
	ToParty                []string      `json:"to_party,omitempty"`
	ToAll                  int           `json:"toall,omitempty"`
	MsgType                event.MsgType `json:"msgtype"`
	AgentID                int64         `json:"agentid,omitempty"`
	Text                   *Text         `json:"text,omitempty"`
	Image                  *Media        `json:"image,omitempty"`
	Voice                  *Media        `json:"voice,omitempty"`
	Video                  *Video        `json:"video,omitempty"`
	File                   *Media        `json:"file,omitempty"`
	TextCard               *TextCard     `json:"textcard,omitempty"`
	News                   *News         `json:"news,omitempty"`
	MPNews                 *MPNews       `json:"mpnews,omitempty"`
	Miniprogram            *Miniprogram  `json:"miniprogram,omitempty"`
	EnableIDTrans          int           `json:"enable_id_trans,omitempty"`
	EnableDuplicateCheck   int           `json:"enable_duplicate_check,omitempty"`
	DuplicateCheckInterval int           `json:"duplicate_check_interval,omitempty"`
}

type ResultExternalContactSend struct {
	InvalidExternalUser  []string `json:"invalid_external_user"`
	InvalidParentUserID  []string `json:"invalid_parent_userid"`
	InvalidStudentUserID []string `json:"invalid_student_userid"`
	InvalidParty         []string `json:"invalid_party"`
}

// SendExternalContactText 发送「学校通知」（文本消息）
func SendExternalContactText(agentID int64, content string, extra *ExternalContactExtra, result *ResultExternalContactSend) wx.Action {
	msg := &ExternalContactMsg{
		MsgType: event.MsgText,
		AgentID: agentID,
		Text: &Text{
			Content: content,
		},
	}

	if extra != nil {
		msg.ToExternalUser = extra.ToExternalUser
		msg.ToParentUserID = extra.ToParentUserID
		msg.ToStudentUserID = extra.ToStudentUserID
		msg.ToParty = extra.ToParty
		msg.ToAll = extra.ToAll
		msg.EnableIDTrans = extra.EnableIDTrans
		msg.EnableDuplicateCheck = extra.EnableDuplicateCheck
		msg.DuplicateCheckInterval = extra.DuplicateCheckInterval
	}

	return wx.NewPostAction(urls.CorpExternalContactMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// SendExternalContactImage 发送「学校通知」（图片消息）
func SendExternalContactImage(agentID int64, mediaID string, extra *ExternalContactExtra, result *ResultExternalContactSend) wx.Action {
	msg := &ExternalContactMsg{
		MsgType: event.MsgImage,
		AgentID: agentID,
		Image: &Media{
			MediaID: mediaID,
		},
	}

	if extra != nil {
		msg.ToExternalUser = extra.ToExternalUser
		msg.ToParentUserID = extra.ToParentUserID
		msg.ToStudentUserID = extra.ToStudentUserID
		msg.ToParty = extra.ToParty
		msg.ToAll = extra.ToAll
		msg.EnableIDTrans = extra.EnableIDTrans
		msg.EnableDuplicateCheck = extra.EnableDuplicateCheck
		msg.DuplicateCheckInterval = extra.DuplicateCheckInterval
	}

	return wx.NewPostAction(urls.CorpExternalContactMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// SendExternalContactVoice 发送「学校通知」（语音消息）
func SendExternalContactVoice(agentID int64, mediaID string, extra *ExternalContactExtra, result *ResultExternalContactSend) wx.Action {
	msg := &ExternalContactMsg{
		MsgType: event.MsgVoice,
		AgentID: agentID,
		Voice: &Media{
			MediaID: mediaID,
		},
	}

	if extra != nil {
		msg.ToExternalUser = extra.ToExternalUser
		msg.ToParentUserID = extra.ToParentUserID
		msg.ToStudentUserID = extra.ToStudentUserID
		msg.ToParty = extra.ToParty
		msg.ToAll = extra.ToAll
		msg.EnableIDTrans = extra.EnableIDTrans
		msg.EnableDuplicateCheck = extra.EnableDuplicateCheck
		msg.DuplicateCheckInterval = extra.DuplicateCheckInterval
	}

	return wx.NewPostAction(urls.CorpExternalContactMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// SendExternalContactVideo 发送「学校通知」（视频消息）
func SendExternalContactVideo(agentID int64, video *Video, extra *ExternalContactExtra, result *ResultExternalContactSend) wx.Action {
	msg := &ExternalContactMsg{
		MsgType: event.MsgVideo,
		AgentID: agentID,
		Video:   video,
	}

	if extra != nil {
		msg.ToExternalUser = extra.ToExternalUser
		msg.ToParentUserID = extra.ToParentUserID
		msg.ToStudentUserID = extra.ToStudentUserID
		msg.ToParty = extra.ToParty
		msg.ToAll = extra.ToAll
		msg.EnableIDTrans = extra.EnableIDTrans
		msg.EnableDuplicateCheck = extra.EnableDuplicateCheck
		msg.DuplicateCheckInterval = extra.DuplicateCheckInterval
	}

	return wx.NewPostAction(urls.CorpExternalContactMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// SendExternalContactFile 发送「学校通知」（文件消息）
func SendExternalContactFile(agentID int64, mediaID string, extra *ExternalContactExtra, result *ResultExternalContactSend) wx.Action {
	msg := &ExternalContactMsg{
		MsgType: event.MsgFile,
		AgentID: agentID,
		File: &Media{
			MediaID: mediaID,
		},
	}

	if extra != nil {
		msg.ToExternalUser = extra.ToExternalUser
		msg.ToParentUserID = extra.ToParentUserID
		msg.ToStudentUserID = extra.ToStudentUserID
		msg.ToParty = extra.ToParty
		msg.ToAll = extra.ToAll
		msg.EnableIDTrans = extra.EnableIDTrans
		msg.EnableDuplicateCheck = extra.EnableDuplicateCheck
		msg.DuplicateCheckInterval = extra.DuplicateCheckInterval
	}

	return wx.NewPostAction(urls.CorpExternalContactMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// SendExternalContactNews 发送「学校通知」（图文消息）
func SendExternalContactNews(agentID int64, articles []*NewsArticle, extra *ExternalContactExtra, result *ResultExternalContactSend) wx.Action {
	msg := &ExternalContactMsg{
		MsgType: event.MsgNews,
		AgentID: agentID,
		News: &News{
			Articles: articles,
		},
	}

	if extra != nil {
		msg.ToExternalUser = extra.ToExternalUser
		msg.ToParentUserID = extra.ToParentUserID
		msg.ToStudentUserID = extra.ToStudentUserID
		msg.ToParty = extra.ToParty
		msg.ToAll = extra.ToAll
		msg.EnableIDTrans = extra.EnableIDTrans
		msg.EnableDuplicateCheck = extra.EnableDuplicateCheck
		msg.DuplicateCheckInterval = extra.DuplicateCheckInterval
	}

	return wx.NewPostAction(urls.CorpExternalContactMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// SendExternalContactMPNews 发送「学校通知」（图文消息 - mpnews）
func SendExternalContactMPNews(agentID int64, articles []*MPNewsArticle, extra *ExternalContactExtra, result *ResultExternalContactSend) wx.Action {
	msg := &ExternalContactMsg{
		MsgType: event.MsgMPNews,
		AgentID: agentID,
		MPNews: &MPNews{
			Articles: articles,
		},
	}

	if extra != nil {
		msg.ToExternalUser = extra.ToExternalUser
		msg.ToParentUserID = extra.ToParentUserID
		msg.ToStudentUserID = extra.ToStudentUserID
		msg.ToParty = extra.ToParty
		msg.ToAll = extra.ToAll
		msg.EnableIDTrans = extra.EnableIDTrans
		msg.EnableDuplicateCheck = extra.EnableDuplicateCheck
		msg.DuplicateCheckInterval = extra.DuplicateCheckInterval
	}

	return wx.NewPostAction(urls.CorpExternalContactMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// SendExternalContactMiniprogram 发送「学校通知」（小程序消息）
func SendExternalContactMiniprogram(agentID int64, minip *Miniprogram, extra *ExternalContactExtra, result *ResultExternalContactSend) wx.Action {
	msg := &ExternalContactMsg{
		MsgType:     event.MsgMinip,
		AgentID:     agentID,
		Miniprogram: minip,
	}

	if extra != nil {
		msg.ToExternalUser = extra.ToExternalUser
		msg.ToParentUserID = extra.ToParentUserID
		msg.ToStudentUserID = extra.ToStudentUserID
		msg.ToParty = extra.ToParty
		msg.ToAll = extra.ToAll
		msg.EnableIDTrans = extra.EnableIDTrans
		msg.EnableDuplicateCheck = extra.EnableDuplicateCheck
		msg.DuplicateCheckInterval = extra.DuplicateCheckInterval
	}

	return wx.NewPostAction(urls.CorpExternalContactMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
