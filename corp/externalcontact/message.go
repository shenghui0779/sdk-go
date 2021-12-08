package externalcontact

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ChatType string

const (
	ChatSingle ChatType = "single"
	ChatGroup  ChatType = "group"
)

type GroupMsgType string

const (
	GroupMsgImage GroupMsgType = "image"
	GroupMsgLink  GroupMsgType = "link"
	GroupMsgMinip GroupMsgType = "miniprogram"
	GroupMsgVideo GroupMsgType = "video"
	GroupMsgFile  GroupMsgType = "file"
)

type GroupText struct {
	Content string `json:"content,omitempty"`
}

type GroupImage struct {
	MediaID string `json:"media_id,omitempty"`
	PicURL  string `json:"pic_url,omitempty"`
}

type GroupLink struct {
	Title  string `json:"title"`
	PicURL string `json:"picurl,omitempty"`
	Desc   string `json:"desc,omitempty"`
	URL    string `json:"url"`
}

type GroupMinip struct {
	Title      string `json:"title"`
	PicMediaID string `json:"pic_media_id"`
	AppID      string `json:"appid"`
	Page       string `json:"page"`
}

type GroupVideo struct {
	MediaID string `json:"media_id"`
}

type GroupFile struct {
	MediaID string `json:"media_id"`
}

type GroupAttachment struct {
	MsgType GroupMsgType `json:"msg_type"`
	Image   *GroupImage  `json:"image,omitempty"`
	Link    *GroupLink   `json:"link,omitempty"`
	Minip   *GroupMinip  `json:"miniprogram,omitempty"`
	Video   *GroupVideo  `json:"video,omitempty"`
	File    *GroupFile   `json:"file,omitempty"`
}

type ParamsMsgTemplateAdd struct {
	ChatType       ChatType           `json:"chat_type,omitempty"`
	ExternalUserID []string           `json:"external_userid,omitempty"`
	Sender         string             `json:"sender,omitempty"`
	Text           *GroupText         `json:"text,omitempty"`
	Attachments    []*GroupAttachment `json:"attachments,omitempty"`
}

type ResultMsgTemplateAdd struct {
	FailList []string `json:"fail_list"`
	MsgID    string   `json:"msgid"`
}

func AddMsgTemplate(params *ParamsMsgTemplateAdd, result *ResultMsgTemplateAdd) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactAddMsgTemplate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
