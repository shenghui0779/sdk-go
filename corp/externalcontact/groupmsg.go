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

type MsgText struct {
	Content string `json:"content,omitempty"`
}

type MsgImage struct {
	MediaID string `json:"media_id,omitempty"`
	PicURL  string `json:"pic_url,omitempty"`
}

type MsgLink struct {
	Title  string `json:"title"`
	PicURL string `json:"picurl,omitempty"`
	Desc   string `json:"desc,omitempty"`
	URL    string `json:"url"`
}

type MsgMinip struct {
	Title      string `json:"title"`
	PicMediaID string `json:"pic_media_id"`
	AppID      string `json:"appid"`
	Page       string `json:"page"`
}

type MsgVideo struct {
	MediaID string `json:"media_id"`
}

type MsgFile struct {
	MediaID string `json:"media_id"`
}

type MsgAttachment struct {
	MsgType MediaType `json:"msg_type"`
	Image   *MsgImage `json:"image,omitempty"`
	Link    *MsgLink  `json:"link,omitempty"`
	Minip   *MsgMinip `json:"miniprogram,omitempty"`
	Video   *MsgVideo `json:"video,omitempty"`
	File    *MsgFile  `json:"file,omitempty"`
}

type ParamsMsgTemplateAdd struct {
	ChatType       ChatType         `json:"chat_type,omitempty"`
	ExternalUserID []string         `json:"external_userid,omitempty"`
	Sender         string           `json:"sender,omitempty"`
	Text           *MsgText         `json:"text,omitempty"`
	Attachments    []*MsgAttachment `json:"attachments,omitempty"`
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

type GroupMsg struct {
	MsgID       string           `json:"msg_id"`
	Creator     string           `json:"creator"`
	CreateTime  int64            `json:"create_time"`
	CreateType  int              `json:"create_type"`
	Text        *MsgText         `json:"text,omitempty"`
	Attachments []*MsgAttachment `json:"attachments,omitempty"`
}

type ParamsGroupMsgList struct {
	ChatType   ChatType `json:"chat_type"`
	StartTime  int64    `json:"start_time"`
	EndTime    int64    `json:"end_time"`
	Creator    string   `json:"creator,omitempty"`
	FilterType int      `json:"filter_type,omitempty"`
	Limit      int      `json:"limit,omitempty"`
	Cursor     string   `json:"cursor,omitempty"`
}

type ResultGroupMsgList struct {
	NextCursor   string      `json:"next_cursor"`
	GroupMsgList []*GroupMsg `json:"group_msg_list"`
}

func ListGroupMsg(params *ParamsGroupMsgList, result *ResultGroupMsgList) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactGetGroupMsgList,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsGroupMsgTaskGet struct {
	MsgID  string `json:"msg_id"`
	Limit  int    `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`
}

type GroupMsgTaskListData struct {
	UserID   string `json:"userid"`
	Status   int    `json:"status"`
	SendTime int64  `json:"send_time"`
}

type ResultGroupMsgTaskGet struct {
	NextCursor string                  `json:"next_cursor"`
	TaskList   []*GroupMsgTaskListData `json:"task_list"`
}

func GetGroupMsgTask(params *ParamsGroupMsgTaskGet, result *ResultGroupMsgTaskGet) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactGetGroupMsgTask,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsGroupMsgSendResultGet struct {
	MsgID  string `json:"msgid"`
	UserID string `json:"userid"`
	Limit  int    `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`
}

type GroupMsgSendResult struct {
	ExternalUserID string `json:"external_userid"`
	ChatID         string `json:"chat_id"`
	UserID         string `json:"userid"`
	Status         int    `json:"status"`
	SendTime       int64  `json:"send_time"`
}

type ResultGroupMsgSendResultGet struct {
	NextCursor string                `json:"next_cursor"`
	SendList   []*GroupMsgSendResult `json:"send_list"`
}

func GetGroupMsgSendResult(params *ParamsGroupMsgSendResultGet, result *ResultGroupMsgSendResultGet) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactGetGroupMsgSendResult,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsWelcomeMsgSend struct {
	WelcomeCode string           `json:"welcome_code"`
	Text        *MsgText         `json:"text,omitempty"`
	Attachments []*MsgAttachment `json:"attachments,omitempty"`
}

func SendWelcomeMsg(params *ParamsWelcomeMsgSend) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactSendWelcomeMsg,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsGroupWelcomeTemplateAdd struct {
	Text    *MsgText  `json:"text,omitempty"`
	Image   *MsgImage `json:"image,omitempty"`
	Link    *MsgLink  `json:"link,omitempty"`
	Minip   *MsgMinip `json:"miniprogram,omitempty"`
	File    *MsgFile  `json:"file,omitempty"`
	Video   *MsgVideo `json:"video,omitempty"`
	AgentID int64     `json:"agent_id,omitempty"`
	Notify  int       `json:"notify,omitempty"`
}

type ResultGroupWelcomeTemplateAdd struct {
	TemplateID string `json:"template_id"`
}

func AddGroupWelcomeTemplate(params *ParamsGroupWelcomeTemplateAdd, result *ResultGroupWelcomeTemplateAdd) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactGroupWelcomeTemplateAdd,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsGroupWelcomeTemplateEdit struct {
	TemplateID string    `json:"template_id"`
	Text       *MsgText  `json:"text,omitempty"`
	Image      *MsgImage `json:"image,omitempty"`
	Link       *MsgLink  `json:"link,omitempty"`
	Minip      *MsgMinip `json:"miniprogram,omitempty"`
	File       *MsgFile  `json:"file,omitempty"`
	Video      *MsgVideo `json:"video,omitempty"`
	AgentID    int64     `json:"agent_id,omitempty"`
}

func EditGroupWelcomeTemplate(params *ParamsGroupWelcomeTemplateEdit) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactGroupWelcomeTemplateEdit,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsGroupWelcomeTemplateGet struct {
	TemplateID string `json:"template_id"`
}

type ResultGroupWelcomeTemplateGet struct {
	Text  *MsgText  `json:"text"`
	Image *MsgImage `json:"image"`
	Link  *MsgLink  `json:"link"`
	Minip *MsgMinip `json:"miniprogram"`
	File  *MsgFile  `json:"file"`
	Video *MsgVideo `json:"video"`
}

func GetGroupWelcomeTemplate(params *ParamsGroupWelcomeTemplateGet, result *ResultGroupWelcomeTemplateGet) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactGroupWelcomeTemplateGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsGroupWelcomeTemplateDelete struct {
	TemplateID string `json:"template_id"`
	AgentID    int64  `json:"agentid,omitempty"`
}

func DeleteGroupWelcomeTemplate(params *ParamsGroupWelcomeTemplateDelete) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactGroupWelcomeTemplateDelete,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}
