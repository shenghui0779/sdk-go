package kf

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ErrServicer struct {
	UserID  string `json:"userid"`
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type ParamsServicerAdd struct {
	OpenKFID   string   `json:"open_kfid"`
	UserIDList []string `json:"userid_list"`
}

type ResultServicerAdd struct {
	ResultList []*ErrServicer `json:"result_list"`
}

func AddServicer(openKFID string, userIDs []string, result *ResultServicerAdd) wx.Action {
	params := &ParamsServicerAdd{
		OpenKFID:   openKFID,
		UserIDList: userIDs,
	}

	return wx.NewPostAction(urls.CorpKFServicerAdd,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsServicerDelete struct {
	OpenKFID   string   `json:"open_kfid"`
	UserIDList []string `json:"userid_list"`
}

type ResultServicerDelete struct {
	ResultList []*ErrServicer `json:"result_list"`
}

func DeleteServicer(openKFID string, userIDs []string, result *ResultServicerDelete) wx.Action {
	params := &ParamsServicerDelete{
		OpenKFID:   openKFID,
		UserIDList: userIDs,
	}

	return wx.NewPostAction(urls.CorpKFServicerDelete,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ServicerListData struct {
	UserID string `json:"userid"`
	Status int    `json:"status"`
}

type ResultServicerList struct {
	ServicerList []*ServicerListData `json:"servicer_list"`
}

func ListServicer(openKFID string, result *ResultServicerList) wx.Action {
	return wx.NewGetAction(urls.CorpKFServicerList,
		wx.WithQuery("open_kfid", openKFID),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsServiceState struct {
	OpenKFID       string `json:"open_kfid"`
	ExternalUserID string `json:"external_userid"`
}

type ResultServiceState struct {
	ServiceState   int    `json:"service_state"`
	ServicerUserID string `json:"servicer_userid"`
}

func GetServiceState(openKFID, externalUserID string, result *ResultServiceState) wx.Action {
	params := &ParamsServiceState{
		OpenKFID:       openKFID,
		ExternalUserID: externalUserID,
	}

	return wx.NewPostAction(urls.CorpKFServiceStateGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsServiceStateTransfer struct {
	OpenKFID       string `json:"open_kfid"`
	ExternalUserID string `json:"external_userid"`
	ServiceState   int    `json:"service_state"`
	ServicerUserID string `json:"servicer_userid"`
}

type ResultServiceStateTransfer struct {
	MsgCode string `json:"msg_code"`
}

func TransferServiceState(params *ParamsServiceStateTransfer, result *ResultServiceStateTransfer) wx.Action {
	return wx.NewPostAction(urls.CorpKFServiceStateTransfer,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
