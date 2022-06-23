package externalcontact

import (
	"encoding/json"

	"github.com/chenghonour/gochat/event"
	"github.com/chenghonour/gochat/urls"
	"github.com/chenghonour/gochat/wx"
)

type MomentText struct {
	Content string `json:"content,omitempty"`
}

type MomentImage struct {
	MediaID string `json:"media_id"`
}

type MomentVideo struct {
	MediaID      string `json:"media_id"`
	ThumbMediaID string `json:"thumb_media_id,omitempty"`
}

type MomentLink struct {
	Title   string `json:"title,omitempty"`
	URL     string `json:"url"`
	MediaID string `json:"media_id"`
}

type MomentLocation struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Name      string `json:"name"`
}

type MomentSenderList struct {
	UserList       []string `json:"user_list,omitempty"`
	DepartmentList []int64  `json:"department_list,omitempty"`
}

type MomentExternalContactList struct {
	TagList []string `json:"tag_list,omitempty"`
}

type MomentVisibleRange struct {
	SenderList          *MomentSenderList          `json:"sender_list,omitempty"`
	ExternalContactList *MomentExternalContactList `json:"external_contact_list,omitempty"`
}

type MomentAttachment struct {
	MsgType event.MsgType `json:"msgtype"`
	Image   *MomentImage  `json:"image,omitempty"`
	Video   *MomentVideo  `json:"video,omitempty"`
	Link    *MomentLink   `json:"link,omitempty"`
}

type ParamsMomentTaskAdd struct {
	Text         *MomentText         `json:"text,omitempty"`
	Attachments  []*MomentAttachment `json:"attachments,omitempty"`
	VisibleRange *MomentVisibleRange `json:"visible_range,omitempty"`
}

type ResultMomentTaskAdd struct {
	JobID string `json:"jobid"`
}

// AddMomentTask 创建客户朋友圈的发表任务
func AddMomentTask(params *ParamsMomentTaskAdd, result *ResultMomentTaskAdd) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactAddMomentTask,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type MomentTaskResult struct {
	ErrCode                    int                               `json:"errcode"`
	ErrMsg                     string                            `json:"errmsg"`
	MomentID                   string                            `json:"moment_id"`
	InvalidSenderList          *MomentInvalidSenderList          `json:"invalid_sender_list"`
	InvalidExternalContactList *MomentInvalidExternalContactList `json:"invalid_external_contact_list"`
}

type MomentInvalidSenderList struct {
	UserList       []string `json:"user_list"`
	DepartmentList []int64  `json:"department_list"`
}

type MomentInvalidExternalContactList struct {
	TagList []string `json:"tag_list"`
}

type ResultMomentTaskResult struct {
	Status int               `json:"status"`
	Type   string            `json:"type"`
	Result *MomentTaskResult `json:"result"`
}

// GetMomentTaskResult 获取客户朋友圈的任务创建结果
func GetMomentTaskResult(jobID string, result *ResultMomentTaskResult) wx.Action {
	return wx.NewGetAction(urls.CorpExternalContactGetMomentTaskResult,
		wx.WithQuery("jobid", jobID),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsMomentList struct {
	StartTime  int64  `json:"start_time"`
	EndTime    int64  `json:"end_time"`
	Creator    string `json:"creator,omitempty"`
	FilterType int    `json:"filter_type,omitempty"`
	Cursor     string `json:"cursor,omitempty"`
	Limit      int    `json:"limit,omitempty"`
}

type MomentListData struct {
	MomentID    string          `json:"moment_id"`
	Creator     string          `json:"creator"`
	CreateTime  string          `json:"create_time"`
	CreateType  int             `json:"create_type"`
	VisibleType int             `json:"visible_type"`
	Text        *MomentText     `json:"text"`
	Image       []*MomentImage  `json:"image"`
	Video       *MomentVideo    `json:"video"`
	Link        *MomentLink     `json:"link"`
	Location    *MomentLocation `json:"location"`
}

type ResultMomentList struct {
	NextCursor string            `json:"next_cursor"`
	MomentList []*MomentListData `json:"moment_list"`
}

// ListMoment 获取企业全部的发表列表
func ListMoment(params *ParamsMomentList, result *ResultMomentList) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactGetMomentList,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsMomentTaskGet struct {
	MomentID string `json:"moment_id"`
	Cursor   string `json:"cursor,omitempty"`
	Limit    int    `json:"limit,omitempty"`
}

type MomentTask struct {
	UserID        string `json:"userid"`
	PublishStatus int    `json:"publish_status"`
}

type ResultMomentTaskGet struct {
	NextCursor string        `json:"next_cursor"`
	TaskList   []*MomentTask `json:"task_list"`
}

// GetMomentTask 获取客户朋友圈企业发表的列表
func GetMomentTask(momentID, cursor string, limit int, result *ResultMomentTaskGet) wx.Action {
	params := &ParamsMomentTaskGet{
		MomentID: momentID,
		Cursor:   cursor,
		Limit:    limit,
	}

	return wx.NewPostAction(urls.CorpExternalContactGetMomentTask,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsMomentCustomerList struct {
	MomentID string `json:"moment_id"`
	UserID   string `json:"userid"`
	Cursor   string `json:"cursor,omitempty"`
	Limit    int    `json:"limit,omitempty"`
}

type MomentCustomer struct {
	UserID         string `json:"userid"`
	ExternalUserID string `json:"external_userid"`
}

type ResultMomentCustomerList struct {
	NextCursor   string            `json:"next_cursor"`
	CustomerList []*MomentCustomer `json:"customer_list"`
}

// ListMomentCustomer 获取客户朋友圈发表时选择的可见范围
func ListMomentCustomer(momentID, userID, cursor string, limit int, result *ResultMomentCustomerList) wx.Action {
	params := &ParamsMomentCustomerList{
		MomentID: momentID,
		UserID:   userID,
		Cursor:   cursor,
		Limit:    limit,
	}

	return wx.NewPostAction(urls.CorpExternalContactGetMomentCustomerList,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsMomentSendResult struct {
	MomentID string `json:"moment_id"`
	UserID   string `json:"userid"`
	Cursor   string `json:"cursor,omitempty"`
	Limit    int    `json:"limit,omitempty"`
}

type ResultMomentSendResult struct {
	NextCursor   string            `json:"next_cursor"`
	CustomerList []*MomentCustomer `json:"customer_list"`
}

// GetMomentSendResult 获取客户朋友圈发表后的可见客户列表
func GetMomentSendResult(momentID, userID, cursor string, limit int, result *ResultMomentSendResult) wx.Action {
	params := &ParamsMomentSendResult{
		MomentID: momentID,
		UserID:   userID,
		Cursor:   cursor,
		Limit:    limit,
	}

	return wx.NewPostAction(urls.CorpExternalContactGetMomentSentResult,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsMomentComments struct {
	MomentID string `json:"moment_id"`
	UserID   string `json:"userid"`
	Cursor   string `json:"cursor,omitempty"`
	Limit    int    `json:"limit,omitempty"`
}

type MomentCommentData struct {
	UserID         string `json:"userid"`
	ExternalUserID string `json:"external_userid"`
	CreateTime     int64  `json:"create_time"`
}

type MomentLikeData struct {
	UserID         string `json:"userid"`
	ExternalUserID string `json:"external_userid"`
	CreateTime     int64  `json:"create_time"`
}

type ResultMomentComments struct {
	CommentList []*MomentCommentData `json:"comment_list"`
	LikeList    []*MomentLikeData    `json:"like_list"`
}

// GetMomentComments 获取客户朋友圈的互动数据
func GetMomentComments(momentID, userID, cursor string, limit int, result *ResultMomentComments) wx.Action {
	params := &ParamsMomentComments{
		MomentID: momentID,
		UserID:   userID,
		Cursor:   cursor,
		Limit:    limit,
	}

	return wx.NewPostAction(urls.CorpExternalContactGetMomentComments,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
