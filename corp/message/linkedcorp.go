package message

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsLinkedcorpMessageSend struct {
	ToUser      []string            `json:"touser,omitempty"`
	ToParty     []string            `json:"toparty,omitempty"`
	ToTag       []string            `json:"totag,omitempty"`
	ToAll       int                 `json:"toall,omitempty"`
	AgentID     int64               `json:"agentid"`
	MsgType     string              `json:"msgtype"`
	Text        *TextMessage        `json:"text,omitempty"`
	Image       *ImageMessage       `json:"image,omitempty"`
	Voice       *VoiceMessage       `json:"voice,omitempty"`
	Video       *VideoMessage       `json:"video,omitempty"`
	File        *FileMessage        `json:"file,omitempty"`
	TextCard    *TextCardMessage    `json:"textcard,omitempty"`
	News        *NewsMessage        `json:"news,omitempty"`
	MPNews      *MPNewsMessage      `json:"mpnews,omitempty"`
	Markdown    *MarkdownMessage    `json:"markdown,omitempty"`
	MinipNotice *MinipNoticeMessage `json:"miniprogram_notice,omitempty"`
}

type ResultLinkedcorpMessageSend struct {
	InvalidUser  []string `json:"invaliduser"`
	InvalidParty []string `json:"invalidparty"`
	InvalidTag   []string `json:"invalidtag"`
}

func LinkedcorpMessageSend(params *ParamsLinkedcorpMessageSend, result *ResultLinkedcorpMessageSend) wx.Action {
	return wx.NewPostAction(urls.CorpLinkedcorpMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
