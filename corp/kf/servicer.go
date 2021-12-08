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

func ServicerAdd(params *ParamsServicerAdd, result *ResultServicerAdd) wx.Action {
	return wx.NewPostAction(urls.CorpKFServicerAdd,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsServicerDel struct {
	OpenKFID   string   `json:"open_kfid"`
	UserIDList []string `json:"userid_list"`
}

type ResultServicerDel struct {
	ResultList []*ServicerDelItem `json:"result_list"`
}

type ServicerDelItem struct {
	UserID  string `json:"userid"`
	ErrCode string `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func ServicerDel(params *ParamsServicerDel, result *ResultServicerDel) wx.Action {
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

type ResultServicerList struct {
	ServicerList []*ServicerListItem `json:"servicer_list"`
}

type ServicerListItem struct {
	UserID string `json:"userid"`
	Status int    `json:"status"`
}

func ServicerList(params *ParamsServicerList, result *ResultServicerList) wx.Action {
	return wx.NewGetAction(urls.CorpKFServicerList,
		wx.WithQuery("open_kfid", params.OpenKFID),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
