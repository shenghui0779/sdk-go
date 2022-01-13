package school

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsHealthReportStat struct {
	Date string `json:"date"`
}

type ResultHealthReportStat struct {
	PV int `json:"pv"`
	UV int `json:"uv"`
}

func GetHealthReportStat(params *ParamsHealthReportStat, result *ResultHealthReportStat) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolGetHealthReportStat,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsHealthReportJobIDs struct {
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
}

type ResultHealthReportJobIDs struct {
	Ending int      `json:"ending"`
	JobIDs []string `json:"jobids"`
}

func GetHealthReportJobIDs(params *ParamsHealthReportJobIDs, result *ResultHealthReportJobIDs) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolGetHealthReportJobIDs,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type HealthReportJobInfo struct {
	Title             string                    `json:"title"`
	Creator           string                    `json:"creator"`
	Type              int                       `json:"type"`
	ApplyRange        *HealthReportApplyRange   `json:"apply_range"`
	ReportTo          *HealthReportTo           `json:"report_to"`
	ReportType        int                       `json:"report_type"`
	SkipWeekend       int                       `json:"skip_weekend"`
	FinishCnt         int                       `json:"finish_cnt"`
	QuestionTemplates []*HealthQuestionTemplate `json:"question_templates"`
}

type HealthReportApplyRange struct {
	UserIDs  []string `json:"userids"`
	PartyIDs []string `json:"partyids"`
}

type HealthReportTo struct {
	UserIDs []string `json:"userids"`
}

type HealthQuestionTemplate struct {
	QuestionID   int64                   `json:"question_id"`
	Title        string                  `json:"title"`
	QuestionType int                     `json:"question_type"`
	IsRequired   int                     `json:"is_required"`    // 健康上报：任务详情返回
	IsMustFill   int                     `json:"is_must_fill"`   // 复学码：老师/学生健康信息返回
	IsNotDisplay int                     `json:"is_not_display"` // 复学码：老师/学生健康信息返回
	OptionList   []*HealthQuestionOption `json:"option_list"`
}

type HealthQuestionOption struct {
	OptionID   int64  `json:"option_id"`
	OptionText string `json:"option_text"`
}

type ParamsHealthReportJobInfo struct {
	JobID string `json:"jobid"`
	Date  string `json:"date"`
}

type ResultHealthReportJobInfo struct {
	JobInfo *HealthReportJobInfo `json:"job_info"`
}

func GetHealthReportJobInfo(params *ParamsHealthReportJobInfo, result *ResultHealthReportJobInfo) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolGetHealthReportJobInfo,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type HealthReportAnswer struct {
	IDType       int                  `json:"id_type"`
	UserID       string               `json:"userid"`
	ReportValues []*HealthReportValue `json:"report_values"`
}

type HealthReportValue struct {
	QuestionID   int64    `json:"question_id"`
	SingleChoice int      `json:"single_choice"`
	MultiChoice  []int    `json:"multi_choice"`
	Text         string   `json:"text"`
	FileID       []string `json:"fileid"`
}

type ParamsHealthReportAnswer struct {
	JobID  string `json:"jobid"`
	Date   string `json:"date"`
	Offset int    `json:"offset,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}

type ResultHealthReportAnswer struct {
	Answers []*HealthReportAnswer `json:"answers"`
}

func GetHealthReportAnswer(params *ParamsHealthReportAnswer, result *ResultHealthReportAnswer) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolGetHealthReportAnswer,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type HealthInfo struct {
	UserID             string               `json:"userid"`
	HealthQRCodeStatus int                  `json:"health_qrcode_status"`
	SelfSubmit         int                  `json:"self_submit"`
	ReportValues       []*HealthReportValue `json:"report_values"`
}

type ParamsTeacherHealthInfo struct {
	Date    string `json:"date"`
	NextKey string `json:"next_key"`
	Limit   int    `json:"limit"`
}

type ResultTeacherHealthInfo struct {
	HealthInfos       []*HealthInfo             `json:"health_infos"`
	QuestionTemplates []*HealthQuestionTemplate `json:"question_templates"`
	TemplateID        string                    `json:"template_id"`
	Ending            int                       `json:"ending"`
	NextKey           string                    `json:"next_key"`
}

func GetTeacherHealthInfo(params *ParamsTeacherHealthInfo, result *ResultHealthInfo) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolGetTeacherHealthInfo,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsHealthInfo struct {
	Date    string `json:"date"`
	NextKey string `json:"next_key"`
	Limit   int    `json:"limit"`
}

type ResultHealthInfo struct {
	HealthInfos       []*HealthInfo             `json:"health_infos"`
	QuestionTemplates []*HealthQuestionTemplate `json:"question_templates"`
	TemplateID        string                    `json:"template_id"`
	Ending            int                       `json:"ending"`
	NextKey           string                    `json:"next_key"`
}

func GetStudentHealthInfo(params *ParamsHealthInfo, result *ResultHealthInfo) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolGetStudentHealthInfo,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type HealthQRCode struct {
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
	UserID     string `json:"userid"`
	QRCodeData string `json:"qrcode_data"`
}

type ParamsHealthQRCode struct {
	Type    int      `json:"type"`
	UserIDs []string `json:"userids"`
}

type ResultHealthQRCode struct {
	ResultList []*HealthQRCode `json:"result_list"`
}

func GetHealthQRCode(params *ParamsHealthQRCode, result *ResultHealthQRCode) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolGetHealthQRCode,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
