package report

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type Location struct {
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Process struct {
	ProcessType   int      `json:"process_type"`
	SolveUserID   string   `json:"solve_userid"`
	ProcessDesc   string   `json:"process_desc"`
	Status        int      `json:"status"`
	SolvedTime    int64    `json:"solved_time"`
	ImageURLs     []string `json:"image_urls"`
	VideoMediaIDs []string `json:"video_media_ids"`
}

type PatrolGrid struct {
	GridID    string   `json:"grid_id"`
	GridName  string   `json:"grid_name"`
	GridAdmin []string `json:"grid_admin"`
}

type ResultPatrolGridInfo struct {
	GridList []*PatrolGrid `json:"grid_list"`
}

func GetPatrolGridInfo(result *ResultPatrolGridInfo) wx.Action {
	return wx.NewGetAction(urls.CorpReportGetPatrolGridInfo,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsPatrolCorpStatus struct {
	GridID string `json:"grid_id"`
}

type ResultPatrolCorpStatus struct {
	Processing   int `json:"processing"`
	AddedToday   int `json:"added_today"`
	SolvedToday  int `json:"solved_today"`
	TotalCase    int `json:"total_case"`
	ToBeAssigned int `json:"to_be_assigned"`
	TotalSolved  int `json:"total_solved"`
}

func GetPatrolCorpStatus(gridID string, result *ResultPatrolCorpStatus) wx.Action {
	params := &ParamsPatrolCorpStatus{
		GridID: gridID,
	}

	return wx.NewPostAction(urls.CorpReportGetPatrolCorpStatus,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsPatrolUserStatus struct {
	UserID string `json:"userid"`
}

type ResultPatrolUserStatus struct {
	Processing  int `json:"processing"`
	AddedToday  int `json:"added_today"`
	SolvedToday int `json:"solved_today"`
}

func GetPatrolUserStatus(userID string, result *ResultPatrolUserStatus) wx.Action {
	params := &ParamsPatrolUserStatus{
		UserID: userID,
	}

	return wx.NewPostAction(urls.CorpReportGetPatrolUserStatus,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsPatrolCategoryStatistic struct {
	CategoryID string `json:"category_id"`
}

type ResultPatrolCategoryStatistic struct {
	DashboardList []*PatrolCategoryStatistic `json:"dashboard_list"`
}

type PatrolCategoryStatistic struct {
	CategoryID    string `json:"category_id"`
	CategoryName  string `json:"category_name"`
	CategoryLevel int    `json:"category_level"`
	CategoryType  int    `json:"category_type"`
	TotalCase     int    `json:"total_case"`
	TotalSolved   int    `json:"total_solved"`
}

func GetPatrolCategoryStatistic(categoryID string, result *ResultPatrolCategoryStatistic) wx.Action {
	params := &ParamsPatrolCategoryStatistic{
		CategoryID: categoryID,
	}

	return wx.NewPostAction(urls.CorpReportPatrolCategoryStatistic,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type PatrolOrder struct {
	OrderID          string     `json:"order_id"`
	Desc             string     `json:"desc"`
	UrgeType         int        `json:"urge_type"`
	CaseName         string     `json:"case_name"`
	GridName         string     `json:"grid_name"`
	GridID           string     `json:"grid_id"`
	CreateTime       int64      `json:"create_time"`
	ImageURLs        []string   `json:"image_urls"`
	VideoMediaIDs    []string   `json:"video_media_ids"`
	Location         *Location  `json:"location"`
	ProcessorUserIDs []string   `json:"processor_userids"`
	ProcessList      []*Process `json:"process_list"`
}

type ParamsPatrolOrderList struct {
	BeginCreateTime int64  `json:"begin_create_time,omitempty"`
	BeginModifyTime int64  `json:"begin_modify_time,omitempty"`
	Cursor          string `json:"cursor,omitempty"`
	Limit           int    `json:"limit,omitempty"`
}

type ResultPatrolOrderList struct {
	NextCursor string         `json:"next_cursor"`
	OrderList  []*PatrolOrder `json:"order_list"`
}

func ListPatrolOrder(beginCreateTime, beginModifyTime int64, cursor string, limit int, result *ResultPatrolOrderList) wx.Action {
	params := &ParamsPatrolOrderList{
		BeginCreateTime: beginCreateTime,
		BeginModifyTime: beginModifyTime,
		Cursor:          cursor,
		Limit:           limit,
	}

	return wx.NewPostAction(urls.CorpReportGetPatrolOrderList,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsPatrolOrderInfo struct {
	OrderID string `json:"order_id"`
}

type ResultPatrolOrderInfo struct {
	OrderInfo *PatrolOrder `json:"order_info"`
}

func GetPatrolOrderInfo(orderID string, result *ResultPatrolOrderInfo) wx.Action {
	params := &ParamsPatrolOrderInfo{
		OrderID: orderID,
	}

	return wx.NewPostAction(urls.CorpReportGetPatrolOrderInfo,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
