package externalcontact

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type MomentMsgType string

const (
	MomentMsgImage MomentMsgType = "image"
	MomentMsgVideo MomentMsgType = "video"
	MomentMsgLink  MomentMsgType = "link"
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
	UserList      []string `json:"user_list,omitempty"`
	DeparmentList []int64  `json:"deparment_list,omitempty"`
}

type MomentExternalContactList struct {
	TagList []string `json:"tag_list,omitempty"`
}

type MomentVisibleRange struct {
	SenderList          *MomentSenderList          `json:"sender_list,omitempty"`
	ExternalContactList *MomentExternalContactList `json:"external_contact_list,omitempty"`
}

type MomentAttachment struct {
	MsgType MomentMsgType `json:"msg_type"`
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

func AddMomentTask(params *ParamsMomentTaskAdd, result *ResultMomentTaskAdd) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactAddMomentTask,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type MomentInvalidSenderList struct {
	UserList      []string `json:"user_list"`
	DeparmentList []int64  `json:"deparment_list"`
}

type MomentInvalidExternalContactList struct {
	TagList []string `json:"tag_list"`
}

type MomentTaskResult struct {
	ErrCode                    int                               `json:"errcode"`
	ErrMsg                     string                            `json:"errmsg"`
	MomentID                   string                            `json:"moment_id"`
	InvalidSenderList          *MomentInvalidSenderList          `json:"invalid_sender_list"`
	InvalidExternalContactList *MomentInvalidExternalContactList `json:"invalid_external_contact_list"`
}

type ResultMomentTaskResultGet struct {
	Status int               `json:"status"`
	Type   string            `json:"type"`
	Result *MomentTaskResult `json:"result"`
}

func GetMomentTaskResult(jobID string, result *ResultMomentTaskResultGet) wx.Action {
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
	CreateTime  int64           `json:"create_time"`
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

func ListMoment(params *ParamsMomentList, result *ResultMomentList) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactGetMomentList,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
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

type MomentTaskListData struct {
	UserID        string `json:"userid"`
	PublishStatus int    `json:"publish_status"`
}

type ResultMomentTaskGet struct {
	NextCursor string                `json:"next_cursor"`
	TaskList   []*MomentTaskListData `json:"task_list"`
}

func GetMomentTask(params *ParamsMomentTaskGet, result *ResultMomentTaskGet) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactGetMomentTask,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsMomentCustomerList struct {
	MomentID string `json:"moment_id"`
	UserID   string `json:"user_id"`
	Cursor   string `json:"cursor,omitempty"`
	Limit    int    `json:"limit,omitempty"`
}

type MomentCustomerListData struct {
	UserID         string `json:"userid"`
	ExternalUserID string `json:"external_userid"`
}

type ResultMomentCustomerList struct {
	NextCursor   string                    `json:"next_cursor"`
	CustomerList []*MomentCustomerListData `json:"customer_list"`
}

func ListMomentCustomer(params *ParamsMomentCustomerList, result *ResultMomentCustomerList) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactGetMomentCustomerList,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsMomentSendResultGet struct {
	MomentID string `json:"moment_id"`
	UserID   string `json:"user_id"`
	Cursor   string `json:"cursor,omitempty"`
	Limit    int    `json:"limit,omitempty"`
}

type MomentSendCustomer struct {
	ExternalUserID string `json:"external_userid"`
}

type ResultMomentSendResultGet struct {
	NextCursor   string                `json:"next_cursor"`
	CustomerList []*MomentSendCustomer `json:"customer_list"`
}

func GetMomentSendResult(params *ParamsMomentSendResultGet, result *ResultMomentSendResultGet) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactGetMomentSentResult,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsMomentCommentsGet struct {
	MomentID string `json:"moment_id"`
	UserID   string `json:"user_id"`
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

type ResultMomentCommentsGet struct {
	CommentList []*MomentCommentData `json:"comment_list"`
	LikeList    []*MomentLikeData    `json:"like_list"`
}

func GetMomentComments(params *ParamsMomentCommentsGet, result *ResultMomentCommentsGet) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactGetMomentComments,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
