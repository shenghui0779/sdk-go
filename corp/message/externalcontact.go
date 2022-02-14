package message

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/event"
	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ExternalContactMsg struct {
	ToExternalUser         []string      `json:"to_external_user,omitempty"`
	ToParentUserID         []string      `json:"to_parent_userid,omitempty"`
	ToStudentUserID        []string      `json:"to_student_userid,omitempty"`
	ToParty                []string      `json:"to_party,omitempty"`
	ToAll                  int           `json:"toall,omitempty"`
	AgentID                int64         `json:"agentid"`
	MsgType                event.MsgType `json:"msgtype"`
	Text                   *Text         `json:"text,omitempty"`
	Image                  *Media        `json:"image,omitempty"`
	Voice                  *Media        `json:"voice,omitempty"`
	Video                  *Video        `json:"video,omitempty"`
	File                   *Media        `json:"file,omitempty"`
	TextCard               *TextCard     `json:"textcard,omitempty"`
	News                   *News         `json:"news,omitempty"`
	MPNews                 *MPNews       `json:"mpnews,omitempty"`
	Markdown               *Text         `json:"markdown,omitempty"`
	MinipNotice            *MinipNotice  `json:"miniprogram_notice,omitempty"`
	EnableIDTrans          int           `json:"enable_id_trans,omitempty"`
	EnableDuplicateCheck   int           `json:"enable_duplicate_check,omitempty"`
	DuplicateCheckInterval int           `json:"duplicate_check_interval,omitempty"`
}

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

type ResultExternalContactSend struct {
	InvalidExternalUser  []string `json:"invalid_external_user"`
	InvalidParentUserID  []string `json:"invalid_parent_userid"`
	InvalidStudentUserID []string `json:"invalid_student_userid"`
	InvalidParty         []string `json:"invalid_party"`
}

func SendExternalContactText(agentID int64, content string, extra *ExternalContactExtra, result *ResultExternalContactSend) wx.Action {
	msg := &ExternalContactMsg{
		AgentID: agentID,
		MsgType: event.MsgText,
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
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendExternalContactImage(agentID int64, mediaID string, extra *ExternalContactExtra, result *ResultExternalContactSend) wx.Action {
	msg := &ExternalContactMsg{
		AgentID: agentID,
		MsgType: event.MsgImage,
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
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendExternalContactVoice(agentID int64, mediaID string, extra *ExternalContactExtra, result *ResultExternalContactSend) wx.Action {
	msg := &ExternalContactMsg{
		AgentID: agentID,
		MsgType: event.MsgVoice,
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
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendExternalContactVideo(agentID int64, video *Video, extra *ExternalContactExtra, result *ResultExternalContactSend) wx.Action {
	msg := &ExternalContactMsg{
		AgentID: agentID,
		MsgType: event.MsgVideo,
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
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendExternalContactFile(agentID int64, mediaID string, extra *ExternalContactExtra, result *ResultExternalContactSend) wx.Action {
	msg := &ExternalContactMsg{
		AgentID: agentID,
		MsgType: event.MsgFile,
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
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendExternalContactTextCard(agentID int64, card *TextCard, extra *ExternalContactExtra, result *ResultExternalContactSend) wx.Action {
	msg := &ExternalContactMsg{
		AgentID:  agentID,
		MsgType:  event.MsgTextCard,
		TextCard: card,
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
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendExternalContactNews(agentID int64, articles []*NewsArticle, extra *ExternalContactExtra, result *ResultExternalContactSend) wx.Action {
	msg := &ExternalContactMsg{
		AgentID: agentID,
		MsgType: event.MsgNews,
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
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendExternalContactMPNews(agentID int64, articles []*MPNewsArticle, extra *ExternalContactExtra, result *ResultExternalContactSend) wx.Action {
	msg := &ExternalContactMsg{
		AgentID: agentID,
		MsgType: event.MsgMPNews,
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
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendExternalContactMarkdown(agentID int64, content string, extra *ExternalContactExtra, result *ResultExternalContactSend) wx.Action {
	msg := &ExternalContactMsg{
		AgentID: agentID,
		MsgType: event.MsgMarkdown,
		Markdown: &Text{
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
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendExternalContactMinipNotice(agentID int64, content string, extra *ExternalContactExtra, result *ResultExternalContactSend) wx.Action {
	msg := &ExternalContactMsg{
		AgentID: agentID,
		MsgType: event.MsgMinipNotice,
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
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
