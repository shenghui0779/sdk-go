package externalcontact

import (
	"encoding/json"

	"github.com/chenghonour/gochat/event"
	"github.com/chenghonour/gochat/urls"
	"github.com/chenghonour/gochat/wx"
)

// ChatType 群发任务的类型
type ChatType string

const (
	ChatSingle ChatType = "single" // 发送给客户
	ChatGroup  ChatType = "group"  // 发送给客户群
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

type MsgAttachment struct {
	MsgType event.MsgType `json:"msgtype"`
	Image   *GroupImage   `json:"image,omitempty"`
	Link    *GroupLink    `json:"link,omitempty"`
	Minip   *GroupMinip   `json:"miniprogram,omitempty"`
	Video   *GroupVideo   `json:"video,omitempty"`
	File    *GroupFile    `json:"file,omitempty"`
}

type ParamsMsgTemplateAdd struct {
	ChatType       ChatType         `json:"chat_type,omitempty"`
	ExternalUserID []string         `json:"external_userid,omitempty"`
	Sender         string           `json:"sender,omitempty"`
	Text           *GroupText       `json:"text,omitempty"`
	Attachments    []*MsgAttachment `json:"attachments,omitempty"`
}

type ResultMsgTemplateAdd struct {
	FailList []string `json:"fail_list"`
	MsgID    string   `json:"msgid"`
}

// AddMsgTemplate 创建企业群发
func AddMsgTemplate(params *ParamsMsgTemplateAdd, result *ResultMsgTemplateAdd) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactAddMsgTemplate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type GroupMsg struct {
	MsgID       string           `json:"msgid"`
	Creator     string           `json:"creator"`
	CreateTime  string           `json:"create_time"`
	CreateType  int              `json:"create_type"`
	Text        *GroupText       `json:"text,omitempty"`
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

// ListGroupMsg 获取群发记录列表
func ListGroupMsg(params *ParamsGroupMsgList, result *ResultGroupMsgList) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactGetGroupMsgList,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsGroupMsgTask struct {
	MsgID  string `json:"msgid"`
	Limit  int    `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`
}

type ResultGroupMsgTask struct {
	NextCursor string          `json:"next_cursor"`
	TaskList   []*GroupMsgTask `json:"task_list"`
}

type GroupMsgTask struct {
	UserID   string `json:"userid"`
	Status   int    `json:"status"`
	SendTime int64  `json:"send_time"`
}

// GetGroupMsgTask 获取群发成员发送任务列表
func GetGroupMsgTask(params *ParamsGroupMsgTask, result *ResultGroupMsgTask) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactGetGroupMsgTask,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsGroupMsgSendResult struct {
	MsgID  string `json:"msgid"`
	UserID string `json:"userid"`
	Limit  int    `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`
}

type ResultGroupMsgSendResult struct {
	NextCursor string                `json:"next_cursor"`
	SendList   []*GroupMsgSendResult `json:"send_list"`
}

type GroupMsgSendResult struct {
	ExternalUserID string `json:"external_userid"`
	ChatID         string `json:"chat_id"`
	UserID         string `json:"userid"`
	Status         int    `json:"status"`
	SendTime       int64  `json:"send_time"`
}

// GetGroupMsgSendResult 获取企业群发成员执行结果
func GetGroupMsgSendResult(params *ParamsGroupMsgSendResult, result *ResultGroupMsgSendResult) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactGetGroupMsgSendResult,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsWelcomeMsgSend struct {
	WelcomeCode string           `json:"welcome_code"`
	Text        *GroupText       `json:"text,omitempty"`
	Attachments []*MsgAttachment `json:"attachments,omitempty"`
}

// SendWelcomeMsg 发送新客户欢迎语
func SendWelcomeMsg(params *ParamsWelcomeMsgSend) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactSendWelcomeMsg,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type ParamsGroupWelcomeTemplateAdd struct {
	Text    *GroupText  `json:"text,omitempty"`
	Image   *GroupImage `json:"image,omitempty"`
	Link    *GroupLink  `json:"link,omitempty"`
	Minip   *GroupMinip `json:"miniprogram,omitempty"`
	File    *GroupFile  `json:"file,omitempty"`
	Video   *GroupVideo `json:"video,omitempty"`
	AgentID int64       `json:"agentid,omitempty"`
	Notify  int         `json:"notify,omitempty"`
}

type ResultGroupWelcomeTemplateAdd struct {
	TemplateID string `json:"template_id"`
}

// AddGroupWelcomeTemplate 添加入群欢迎语素材
func AddGroupWelcomeTemplate(params *ParamsGroupWelcomeTemplateAdd, result *ResultGroupWelcomeTemplateAdd) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactGroupWelcomeTemplateAdd,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsGroupWelcomeTemplateEdit struct {
	TemplateID string      `json:"template_id"`
	Text       *GroupText  `json:"text,omitempty"`
	Image      *GroupImage `json:"image,omitempty"`
	Link       *GroupLink  `json:"link,omitempty"`
	Minip      *GroupMinip `json:"miniprogram,omitempty"`
	File       *GroupFile  `json:"file,omitempty"`
	Video      *GroupVideo `json:"video,omitempty"`
	AgentID    int64       `json:"agentid,omitempty"`
}

// EditGroupWelcomeTemplate 编辑入群欢迎语素材
func EditGroupWelcomeTemplate(params *ParamsGroupWelcomeTemplateEdit) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactGroupWelcomeTemplateEdit,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type ParamsGroupWelcomeTemplateGet struct {
	TemplateID string `json:"template_id"`
}

type ResultGroupWelcomeTemplateGet struct {
	Text  *GroupText  `json:"text"`
	Image *GroupImage `json:"image"`
	Link  *GroupLink  `json:"link"`
	Minip *GroupMinip `json:"miniprogram"`
	File  *GroupFile  `json:"file"`
	Video *GroupVideo `json:"video"`
}

// GetGroupWelcomeTemplate 获取入群欢迎语素材
func GetGroupWelcomeTemplate(templateID string, result *ResultGroupWelcomeTemplateGet) wx.Action {
	params := &ParamsGroupWelcomeTemplateGet{
		TemplateID: templateID,
	}

	return wx.NewPostAction(urls.CorpExternalContactGroupWelcomeTemplateGet,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
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

// DeleteGroupWelcomeTemplate 删除入群欢迎语素材
func DeleteGroupWelcomeTemplate(templateID string, agentID int64) wx.Action {
	params := &ParamsGroupWelcomeTemplateDelete{
		TemplateID: templateID,
		AgentID:    agentID,
	}

	return wx.NewPostAction(urls.CorpExternalContactGroupWelcomeTemplateDelete,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}
