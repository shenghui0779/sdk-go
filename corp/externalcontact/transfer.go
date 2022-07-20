package externalcontact

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsCustomerTranster struct {
	HandoverUserID     string   `json:"handover_userid"`
	TakeoverUserID     string   `json:"takeover_userid"`
	ExternalUserID     []string `json:"external_userid"`
	TransferSuccessMsg string   `json:"transfer_success_msg,omitempty"`
}

type ErrCustomerTransfer struct {
	ExternalUserID string `json:"external_userid"`
	ErrCode        int    `json:"errcode"`
}

type ResultCustomerTransfer struct {
	Customer []*ErrCustomerTransfer `json:"customer"`
}

// TransferCustomer 分配在职成员的客户
func TransferCustomer(params *ParamsCustomerTranster, result *ResultCustomerTransfer) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactTransferCustomer,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ParamsTransferRet struct {
	HandoverUserID string `json:"handover_userid"`
	TakeoverUserID string `json:"takeover_userid"`
	Cursor         string `json:"cursor,omitempty"`
}

type ResultTransferRet struct {
	Customer   []*TransferRet `json:"customer"`
	NextCursor string         `json:"next_cursor"`
}

type TransferRet struct {
	ExternalUserID string `json:"external_userid"`
	Status         int    `json:"status"`
	TakeoverTime   int64  `json:"takeover_time"`
}

// GetTransferResult 查询客户接替状态
func GetTransferResult(handoverUserID, takeoverUserID, cursor string, result *ResultTransferRet) wx.Action {
	params := &ParamsTransferRet{
		HandoverUserID: handoverUserID,
		TakeoverUserID: takeoverUserID,
		Cursor:         cursor,
	}

	return wx.NewPostAction(urls.CorpExternalContactTransferResult,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ParamsUnassignedList struct {
	PageID   int    `json:"page_id,omitempty"`
	Cursor   string `json:"cursor,omitempty"`
	PageSize int    `json:"page_size,omitempty"`
}

type UnassignedInfo struct {
	HandoverUserID string `json:"handover_userid"`
	ExternalUserID string `json:"external_userid"`
	DimissionTime  int64  `json:"dimission_time"`
}

type ResultUnassignedList struct {
	Info       []*UnassignedInfo `json:"info"`
	IsLast     bool              `json:"is_last"`
	NextCursor string            `json:"next_cursor"`
}

// ListUnassigned 获取待分配的离职成员列表
func ListUnassigned(pageID, pageSize int, cursor string, result *ResultUnassignedList) wx.Action {
	params := &ParamsUnassignedList{
		PageID:   pageID,
		Cursor:   cursor,
		PageSize: pageSize,
	}

	return wx.NewPostAction(urls.CorpExternalContactGetUnassignedList,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ParamsResignedCustomerTransfer struct {
	HandoverUserID string   `json:"handover_userid"`
	TakeoverUserID string   `json:"takeover_userid"`
	ExternalUserID []string `json:"external_userid"`
}

type ResultResignedCustomerTransfer struct {
	Customer []*ErrCustomerTransfer `json:"customer"`
}

// TransferResignedCustomer 分配离职成员的客户
func TransferResignedCustomer(handoverUserID, takeoverUserID string, externalUserIDs []string, result *ResultResignedCustomerTransfer) wx.Action {
	params := &ParamsResignedCustomerTransfer{
		HandoverUserID: handoverUserID,
		TakeoverUserID: takeoverUserID,
		ExternalUserID: externalUserIDs,
	}

	return wx.NewPostAction(urls.CorpExternalContactTransferResignedCustomer,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ParamsResignedTransferRet struct {
	HandoverUserID string `json:"handover_userid"`
	TakeoverUserID string `json:"takeover_userid"`
	Cursor         string `json:"cursor,omitempty"`
}

type ResultResignedTransferRet struct {
	Customer   []*ResignedTransferRet `json:"customer"`
	NextCursor string                 `json:"next_cursor"`
}

type ResignedTransferRet struct {
	ExternalUserID string `json:"external_userid"`
	Status         int    `json:"status"`
	TakeoverTime   int64  `json:"takeover_time"`
}

// GetResignedTransferResult 查询客户接替状态
func GetResignedTransferResult(handoverUserID, takeoverUserID, cursor string, result *ResultResignedTransferRet) wx.Action {
	params := &ParamsResignedTransferRet{
		HandoverUserID: handoverUserID,
		TakeoverUserID: takeoverUserID,
		Cursor:         cursor,
	}

	return wx.NewPostAction(urls.CorpExternalContactResignedTransferResult,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ParamsGroupChatTransfer struct {
	ChatIDList []string `json:"chat_id_list"`
	NewOwner   string   `json:"new_owner"`
}

type ErrGroupChatTransfer struct {
	ChatID  string `json:"chat_id"`
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type ResultGroupChatTransfer struct {
	FailedChatList []*ErrGroupChatTransfer `json:"failed_chat_list"`
}

// TransferGroupChat 分配离职成员的客户群
func TransferGroupChat(chatIDs []string, newOwner string, result *ResultGroupChatTransfer) wx.Action {
	params := &ParamsGroupChatTransfer{
		ChatIDList: chatIDs,
		NewOwner:   newOwner,
	}

	return wx.NewPostAction(urls.CorpExternalContactGroupChatTranster,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}
