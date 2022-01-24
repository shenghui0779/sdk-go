package kf

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsAccountAdd struct {
	Name    string `json:"name"`
	MediaID string `json:"media_id"`
}

type ResultAccountAdd struct {
	OpenKFID string `json:"open_kfid"`
}

func AddAccount(params *ParamsAccountAdd, result *ResultAccountAdd) wx.Action {
	return wx.NewPostAction(urls.CorpKFAccountAdd,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsAccountDelete struct {
	OpenKFID string `json:"open_kfid"`
}

func DeleteAccount(openKFID string) wx.Action {
	params := &ParamsAccountDelete{
		OpenKFID: openKFID,
	}

	return wx.NewPostAction(urls.CorpKFAccountDelete,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsAccountUpdate struct {
	OpenKFID string `json:"open_kfid"`
	Name     string `json:"name"`
	MediaID  string `json:"media_id"`
}

func UpdateAccount(params *ParamsAccountUpdate) wx.Action {
	return wx.NewPostAction(urls.CorpKFAccountUpdate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ResultAccountList struct {
	AccountList []*AccountListData `json:"account_list"`
}

type AccountListData struct {
	OpenKFID string `json:"open_kfid"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
}

func ListAccount(result *ResultAccountList) wx.Action {
	return wx.NewGetAction(urls.CorpKFAccountList,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsContactWayAdd struct {
	OpenKFID string `json:"open_kfid"`
	Scene    string `json:"scene"`
}

type ResultContactWayAdd struct {
	URL string `json:"url"`
}

func AddContactWay(params *ParamsContactWayAdd, result *ResultContactWayAdd) wx.Action {
	return wx.NewPostAction(urls.CorpKFAddContactWay,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
