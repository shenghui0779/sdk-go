package school

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsHealthReportStatGet struct {
	Date string `json:"date"`
}

type ResultHealthReportStatGet struct {
	PV int `json:"pv"`
	UV int `json:"uv"`
}

func GetHealthReportStat(params *ParamsHealthReportStatGet, result *ResultHealthReportStatGet) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolGetHealthReportStat,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsHealthReportJobIDsGet struct {
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
}

type ResultHealthReportJobIDsGet struct {
	Ending int      `json:"ending"`
	JobIDs []string `json:"jobids"`
}

func GetHealthReportJobIDs(params *ParamsHealthReportJobIDsGet, result *ResultHealthReportJobIDsGet) wx.Action {
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

type ParamsHealthReportJobInfoGet struct {
	JobID string `json:"jobid"`
	Date  string `json:"date"`
}

type ResultHealthReportJobInfoGet struct {
	JobInfo *HealthReportJobInfo `json:"job_info"`
}

func GetHealthReportJobInfo(params *ParamsHealthReportJobInfoGet, result *ResultHealthReportJobInfoGet) wx.Action {
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

type ParamsHealthReportAnswerGet struct {
	JobID  string `json:"jobid"`
	Date   string `json:"date"`
	Offset int    `json:"offset,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}

type ResultHealthReportAnswerGet struct {
	Answers []*HealthReportAnswer `json:"answers"`
}

func GetHealthReportAnswer(params *ParamsHealthReportAnswerGet, result *ResultHealthReportAnswerGet) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolGetHealthReportAnswer,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
