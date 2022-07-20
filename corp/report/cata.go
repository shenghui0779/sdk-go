package report

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsGridCataAdd struct {
	CategoryName     string `json:"category_name"`
	Level            int    `json:"level"`
	ParentCategoryID string `json:"parent_category_id,omitempty"`
}

type ResultGridCataAdd struct {
	CategoryID string `json:"category_id"`
}

// AddGridCata 添加事件类别
func AddGridCata(params *ParamsGridCataAdd, result *ResultGridCataAdd) wx.Action {
	return wx.NewPostAction(urls.CorpReportGridCataAdd,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ParamsGridCataUpdate struct {
	CategoryID       string `json:"category_id"`
	CategoryName     string `json:"category_name"`
	Level            int    `json:"level"`
	ParentCategoryID string `json:"parent_category_id,omitempty"`
}

// UpdateGridCata 修改事件类别
func UpdateGridCata(params *ParamsGridCataUpdate) wx.Action {
	return wx.NewPostAction(urls.CorpReportGridCataUpdate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type ParamsGridCataDelete struct {
	CategoryID string `json:"category_id"`
}

// DeleteGridCata 删除事件类别
func DeleteGridCata(categoryID string) wx.Action {
	params := &ParamsGridCataDelete{
		CategoryID: categoryID,
	}

	return wx.NewPostAction(urls.CorpReportGridCataDelete,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type ResultGridCataList struct {
	CategoryList []*GridCategory `json:"category_list"`
}

type GridCategory struct {
	CategoryID       string `json:"category_id"`
	CategoryName     string `json:"category_name"`
	Level            int    `json:"level"`
	ParentCategoryID string `json:"parent_category_id"`
}

// ListGridCata 获取事件类别列表
func ListGridCata(result *ResultGridCataList) wx.Action {
	return wx.NewPostAction(urls.CorpReportGridCataList,
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}
