package report

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsGridAdd struct {
	GridName     string   `json:"grid_name"`
	GridParentID string   `json:"grid_parent_id"`
	GridAdmin    []string `json:"grid_admin"`
	GridMember   []string `json:"grid_member,omitempty"`
}

type ResultGridAdd struct {
	GridID         string   `json:"grid_id"`
	InvalidUserIDs []string `json:"invalid_userids"`
}

// AddGrid 添加网格
func AddGrid(params *ParamsGridAdd, result *ResultGridAdd) wx.Action {
	return wx.NewPostAction(urls.CorpReportGridAdd,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsGridUpdate struct {
	GridID       string   `json:"grid_id"`
	GridName     string   `json:"grid_name"`
	GridParentID string   `json:"grid_parent_id"`
	GridAdmin    []string `json:"grid_admin"`
	GridMember   []string `json:"grid_member,omitempty"`
}

type ResultGridUpdate struct {
	InvalidUserIDs []string `json:"invalid_userids"`
}

// UpdateGrid 编辑网格
func UpdateGrid(params *ParamsGridUpdate, result *ResultGridUpdate) wx.Action {
	return wx.NewPostAction(urls.CorpReportGridUpdate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsGridDelete struct {
	GridID string `json:"grid_id"`
}

// DeleteGrid 删除网格
func DeleteGrid(gridID string) wx.Action {
	params := &ParamsGridDelete{
		GridID: gridID,
	}

	return wx.NewPostAction(urls.CorpReportGridDelete,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type Grid struct {
	GridID       string   `json:"grid_id"`
	GridName     string   `json:"grid_name"`
	GridParentID string   `json:"grid_parent_id"`
	GridAdmin    []string `json:"grid_admin"`
	GridMember   []string `json:"grid_member"`
}

type ParamsGridList struct {
	GridID string `json:"grid_id"`
}

type ResultGridList struct {
	GridList []*Grid `json:"grid_list"`
}

// ListGrid 获取网格列表
func ListGrid(gridID string, result *ResultGridList) wx.Action {
	params := &ParamsGridList{
		GridID: gridID,
	}

	return wx.NewPostAction(urls.CorpReportGridList,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsUserGridInfo struct {
	UserID string `json:"userid"`
}

type ResultUserGridInfo struct {
	ManageGrids []*UserGridInfo `json:"manage_grids"`
	JoinedGrids []*UserGridInfo `json:"joined_grids"`
}

type UserGridInfo struct {
	GridID   string `json:"grid_id"`
	GridName string `json:"grid_name"`
}

// GetUserGridInfo 获取用户负责及参与的网格列表
func GetUserGridInfo(userID string, result *ResultUserGridInfo) wx.Action {
	params := &ParamsUserGridInfo{
		UserID: userID,
	}

	return wx.NewPostAction(urls.CorpReportGetUserGridInfo,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
