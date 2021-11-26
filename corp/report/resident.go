package report

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ResidentGrid struct {
	GridID    string   `json:"grid_id"`
	GridName  string   `json:"grid_name"`
	GridAdmin []string `json:"grid_admin"`
}

type ResultResidentGridInfo struct {
	GridList []*ResidentGrid `json:"grid_list"`
}

func GetResidentGridInfo(result *ResultResidentGridInfo) wx.Action {
	return wx.NewGetAction(urls.CorpReportGetResidentGridInfo,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsResidentCorpStatus struct {
	GridID string `json:"grid_id"`
}

type ResultResidentCorpStatus struct {
	Processing    int `json:"processing"`
	AddedToday    int `json:"added_today"`
	SolvedToday   int `json:"solved_today"`
	Pending       int `json:"pending"`
	TotalCase     int `json:"total_case"`
	TotalAccepted int `json:"total_accepted"`
	TotalSolved   int `json:"total_solved"`
}

func GetResidentCorpStatus(params *ParamsResidentCorpStatus, result *ResultResidentCorpStatus) wx.Action {
	return wx.NewPostAction(urls.CorpReportGetResidentCorpStatus,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsResidentUserStatus struct {
	UserID string `json:"userid"`
}

type ResultResidentUserStatus struct {
	Processing  int `json:"processing"`
	AddedToday  int `json:"added_today"`
	SolvedToday int `json:"solved_today"`
	Pending     int `json:"pending"`
}

func GetResidentUserStatus(params *ParamsResidentUserStatus, result *ResultResidentUserStatus) wx.Action {
	return wx.NewPostAction(urls.CorpReportGetResidentUserStatus,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ResidentCategoryStatistic struct {
	CategoryID    string `json:"category_id"`
	CategoryName  string `json:"category_name"`
	CategoryLevel int    `json:"category_level"`
	CategoryType  int    `json:"category_type"`
	TotalCase     int    `json:"total_case"`
	TotalSolved   int    `json:"total_solved"`
}

type ParamsResidentCategoryStatistic struct {
	CategoryID string `json:"category_id"`
}

type ResultResidentCategoryStatistic struct {
	DashboardList []*ResidentCategoryStatistic `json:"dashboard_list"`
}

func GetResidentCategoryStatistic(params *ParamsResidentCategoryStatistic, result *ResultResidentCategoryStatistic) wx.Action {
	return wx.NewPostAction(urls.CorpReportResidentCategoryStatistic,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ResidentOrder struct {
	OrderID          string     `json:"order_id"`
	Desc             string     `json:"desc"`
	UrgeType         int        `json:"urge_type"`
	CaseName         string     `json:"case_name"`
	GridName         string     `json:"grid_name"`
	GridID           string     `json:"grid_id"`
	ReporterName     string     `json:"reporter_name"`
	ReporterMobile   string     `json:"reporter_mobile"`
	UnionID          string     `json:"unionid"`
	CreateTime       int64      `json:"create_time"`
	ImageURLs        []string   `json:"image_urls"`
	VideoMediaIDs    []string   `json:"video_media_ids"`
	Location         *Location  `json:"location"`
	ProcessorUserIDs []string   `json:"processor_userids"`
	ProcessList      []*Process `json:"process_list"`
}

type ParamsResidentOrderList struct {
	BeginCreateTime int64  `json:"begin_create_time,omitempty"`
	BeginModifyTime int64  `json:"begin_modify_time,omitempty"`
	Cursor          string `json:"cursor,omitempty"`
	Limit           int    `json:"limit,omitempty"`
}

type ResultResidentOrderList struct {
	NextCursor string           `json:"next_cursor"`
	OrderList  []*ResidentOrder `json:"order_list"`
}

func GetResidentOrderList(params *ParamsResidentOrderList, result *ResultResidentOrderList) wx.Action {
	return wx.NewPostAction(urls.CorpReportGetResidentOrderList,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsResidentOrderInfo struct {
	OrderID string `json:"order_id"`
}

type ResultResidentOrderInfo struct {
	OrderInfo *ResidentOrder `json:"order_info"`
}

func GetResidentOrderInfo(params *ParamsResidentOrderInfo, result *ResultResidentOrderInfo) wx.Action {
	return wx.NewPostAction(urls.CorpReportGetResidentOrderInfo,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
