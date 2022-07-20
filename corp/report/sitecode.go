package report

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsSiteCodeList struct {
	Cursor string `json:"cursor,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}

type ResultSiteCodeList struct {
	SiteCodeInfos []*SiteCodeInfo `json:"site_code_infos"`
	NextCursor    string          `json:"next_cursor"`
}

type SiteCodeInfo struct {
	ID        string   `json:"id"`
	Type      string   `json:"type"`
	Area      string   `json:"area"`
	Address   string   `json:"address"`
	Name      string   `json:"name"`
	Admin     []string `json:"admin"`
	QRCodeURL string   `json:"qr_code_url"`
}

// ListSiteCode 获取场所码列表
func ListSiteCode(cursor string, limit int, result *ResultSiteCodeList) wx.Action {
	params := &ParamsSiteCodeList{
		Cursor: cursor,
		Limit:  limit,
	}

	return wx.NewPostAction(urls.CorpReportGetSiteCodeList,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ParamsSiteCodeReportInfo struct {
	SiteID string `json:"siteid"`
}

type ResultSiteCodeReportInfo struct {
	QuestionTemplates []*QuestionTemplate `json:"question_templates"`
}

type QuestionTemplate struct {
	QuestionID   int64             `json:"question_id"`
	Title        string            `json:"title"`
	QuestionType int               `json:"question_type"`
	IsRequired   int               `json:"is_required"`
	OptionList   []*QuestionOption `json:"option_list"`
}

type QuestionOption struct {
	OptionID   int64  `json:"option_id"`
	OptionText string `json:"option_text"`
}

// GetSiteCodeReportInfo 获取场所码上报问卷
func GetSiteCodeReportInfo(siteID string, result *ResultSiteCodeReportInfo) wx.Action {
	params := &ParamsSiteCodeReportInfo{
		SiteID: siteID,
	}

	return wx.NewPostAction(urls.CorpReportGetSiteCodeReportInfo,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ParamsSiteCodeReportAnswer struct {
	SiteID string `json:"siteid"`
	Date   string `json:"date"`
	Cursor string `json:"cursor,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}

type ResultSiteCodeReportAnswer struct {
	Answers    []*QuestionAnswer `json:"answers"`
	NextCursor string            `json:"next_cursor"`
	HasMore    int               `json:"has_more"`
}

type QuestionAnswer struct {
	ReportTime   int64          `json:"report_time"`
	ReportValues []*ReportValue `json:"report_values"`
}

type ReportValue struct {
	QuestionID   int64  `json:"question_id"`
	SingleChoice int    `json:"single_choice"`
	MultiChoice  []int  `json:"multi_choice"`
	Text         string `json:"text"`
}

// GetSiteCodeReportAnswer 获取用户填写答案
func GetSiteCodeReportAnswer(siteID, date, cursor string, limit int, result *ResultSiteCodeReportAnswer) wx.Action {
	params := &ParamsSiteCodeReportAnswer{
		SiteID: siteID,
		Date:   date,
		Cursor: cursor,
		Limit:  limit,
	}

	return wx.NewPostAction(urls.CorpReportGetSiteCodeReportAnswer,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}
