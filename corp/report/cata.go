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

func AddGridCata(params *ParamsGridCataAdd, result *ResultGridCataAdd) wx.Action {
	return wx.NewPostAction(urls.CorpReportGridCataAdd,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsGridCataUpdate struct {
	CategoryID       string `json:"category_id"`
	CategoryName     string `json:"category_name"`
	Level            int    `json:"level"`
	ParentCategoryID string `json:"parent_category_id,omitempty"`
}

func UpdateGridCata(params *ParamsGridCataUpdate) wx.Action {
	return wx.NewPostAction(urls.CorpReportGridCataUpdate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsGridCataDelete struct {
	CategoryID string `json:"category_id"`
}

func DeleteGridCata(params *ParamsGridCataDelete) wx.Action {
	return wx.NewPostAction(urls.CorpReportGridCataDelete,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type GridCata struct {
	CataID       string `json:"cata_id"`
	CataName     string `json:"cata_name"`
	LevelID      int    `json:"level_id"`
	ParentCataID string `json:"parent_cata_id"`
}

type ResultGridCataList struct {
	CataList []*GridCata `json:"cata_list"`
}

func ListGridCata(result *ResultGridCataList) wx.Action {
	return wx.NewPostAction(urls.CorpReportGridCataList,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
