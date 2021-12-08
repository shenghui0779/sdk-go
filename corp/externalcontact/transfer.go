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

func TransferCustomer(params *ParamsCustomerTranster, result *ResultCustomerTransfer) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactTransferCustomer,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsTransferResultGet struct {
	HandoverUserID string `json:"handover_userid"`
	TakeoverUserID string `json:"takeover_userid"`
	Cursor         string `json:"cursor,omitempty"`
}

type TransferResult struct {
	ExternalUserID string `json:"external_userid"`
	Status         int    `json:"status"`
	TakeoverTime   int64  `json:"takeover_time"`
}

type ResultTransferResultGet struct {
	Customer   []*TransferResult `json:"customer"`
	NextCursor string            `json:"next_cursor"`
}

func GetTransferResult(params *ParamsTransferResultGet, result *ResultTransferResultGet) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactTransferResult,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
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

func ListUnassigned(params *ParamsUnassignedList, result *ResultUnassignedList) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactGetUnassignedList,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
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

func TransferGroupChat(params *ParamsGroupChatTransfer, result *ResultGroupChatTransfer) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactGroupChatTranster,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
