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

// AddAccount 添加客服帐号
func AddAccount(name, mediaID string, result *ResultAccountAdd) wx.Action {
	params := &ParamsAccountAdd{
		Name:    name,
		MediaID: mediaID,
	}

	return wx.NewPostAction(urls.CorpKFAccountAdd,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsAccountDelete struct {
	OpenKFID string `json:"open_kfid"`
}

// DeleteAccount 删除客服帐号
func DeleteAccount(openKFID string) wx.Action {
	params := &ParamsAccountDelete{
		OpenKFID: openKFID,
	}

	return wx.NewPostAction(urls.CorpKFAccountDelete,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type ParamsAccountUpdate struct {
	OpenKFID string `json:"open_kfid"`
	Name     string `json:"name"`
	MediaID  string `json:"media_id"`
}

// UpdateAccount 修改客服帐号
func UpdateAccount(openKFID, name, mediaID string) wx.Action {
	params := &ParamsAccountUpdate{
		OpenKFID: openKFID,
		Name:     name,
		MediaID:  mediaID,
	}

	return wx.NewPostAction(urls.CorpKFAccountUpdate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
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

// ListAccount 获取客服帐号列表
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

// AddContactWay 获取客服帐号链接
func AddContactWay(openKFID, scene string, result *ResultContactWayAdd) wx.Action {
	params := &ParamsContactWayAdd{
		OpenKFID: openKFID,
		Scene:    scene,
	}

	return wx.NewPostAction(urls.CorpKFAddContactWay,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
