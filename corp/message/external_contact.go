package message

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsExternalContactMessageSend struct {
	ToExternalUser         []string            `json:"to_external_user,omitempty"`
	ToParentUserID         []string            `json:"to_parent_userid,omitempty"`
	ToStudentUserID        []string            `json:"to_student_userid,omitempty"`
	ToParty                []string            `json:"to_party,omitempty"`
	ToAll                  int                 `json:"toall,omitempty"`
	AgentID                int64               `json:"agentid"`
	MsgType                string              `json:"msgtype"`
	Text                   *TextMessage        `json:"text,omitempty"`
	Image                  *ImageMessage       `json:"image,omitempty"`
	Voice                  *VoiceMessage       `json:"voice,omitempty"`
	Video                  *VideoMessage       `json:"video,omitempty"`
	File                   *FileMessage        `json:"file,omitempty"`
	TextCard               *TextCardMessage    `json:"textcard,omitempty"`
	News                   *NewsMessage        `json:"news,omitempty"`
	MPNews                 *MPNewsMessage      `json:"mpnews,omitempty"`
	Markdown               *MarkdownMessage    `json:"markdown,omitempty"`
	MinipNotice            *MinipNoticeMessage `json:"miniprogram_notice,omitempty"`
	EnableIDTrans          int                 `json:"enable_id_trans,omitempty"`
	EnableDuplicateCheck   int                 `json:"enable_duplicate_check,omitempty"`
	DuplicateCheckInterval int                 `json:"duplicate_check_interval,omitempty"`
}

type ResultExternalContactMessageSend struct {
	InvalidExternalUser  []string `json:"invalid_external_user"`
	InvalidParentUserID  []string `json:"invalid_parent_userid"`
	InvalidStudentUserID []string `json:"invalid_student_userid"`
	InvalidParty         []string `json:"invalid_party"`
}

func SendExternalContactMessage(params *ParamsExternalContactMessageSend, result *ResultExternalContactMessageSend) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
