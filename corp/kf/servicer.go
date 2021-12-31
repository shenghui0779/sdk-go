package kf

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsServicerAdd struct {
	OpenKFID   string   `json:"open_kfid"`
	UserIDList []string `json:"userid_list"`
}

type ResultServicerAdd struct {
	ResultList []*ServicerAddItem `json:"result_list"`
}

type ServicerAddItem struct {
	UserID  string `json:"userid"`
	ErrCode string `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func AddServicer(params *ParamsServicerAdd, result *ResultServicerAdd) wx.Action {
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

type ErrServicerDelete struct {
	UserID  string `json:"userid"`
	ErrCode string `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type ResultServicerDelete struct {
	ResultList []*ErrServicerDelete `json:"result_list"`
}

func DeleteServicer(params *ParamsServicerDelete, result *ResultServicerDelete) wx.Action {
	return wx.NewPostAction(urls.CorpKFServicerDelete,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsServicerList struct {
	OpenKFID string `json:"open_kfid"`
}

type ServicerListData struct {
	UserID string `json:"userid"`
	Status int    `json:"status"`
}

type ResultServicerList struct {
	ServicerList []*ServicerListData `json:"servicer_list"`
}

func ListServicer(params *ParamsServicerList, result *ResultServicerList) wx.Action {
	return wx.NewGetAction(urls.CorpKFServicerList,
		wx.WithQuery("open_kfid", params.OpenKFID),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
