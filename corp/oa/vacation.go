package oa

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ResultCorpVacationConf struct {
	Lists []*VacationConf `json:"lists"`
}

type VacationConf struct {
	ID             int64              `json:"id"`
	Name           string             `json:"name"`
	TimeAttr       int                `json:"time_attr"`
	DurationType   int                `json:"duration_type"`
	QuotaAttr      *VacationQuotaAttr `json:"quota_attr"`
	PerdayDuration int                `json:"perday_duration"`
}

type VacationQuotaAttr struct {
	Type              int   `json:"type"`
	AutoresetTime     int64 `json:"autoreset_time"`
	AutoresetDuration int   `json:"autoreset_duration"`
}

func GetCorpVacationConf(result *ResultCorpVacationConf) wx.Action {
	return wx.NewGetAction(urls.CorpOAGetVacationCorpConf,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsUserVacationQuota struct {
	UserID string `json:"userid"`
}

type ResultUserVacationQuota struct {
	Lists []*VacationQuota `json:"lists"`
}

type VacationQuota struct {
	ID             int64  `json:"id"`
	AssignDuration int    `json:"assignduration"`
	UsedDuration   int    `json:"usedduration"`
	LeftDuration   int    `json:"leftduration"`
	VacationName   string `json:"vacationname"`
}

func GetUserVacationQuota(userID string, result *ResultUserVacationQuota) wx.Action {
	params := &ParamsUserVacationQuota{
		UserID: userID,
	}

	return wx.NewPostAction(urls.CorpOAGetUserVacationQuota,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsOneUserQuotaSet struct {
	UserID       string `json:"userid"`
	VacationID   int64  `json:"vacation_id"`
	LeftDuration int    `json:"left_duration"`
	TimeAttr     int    `json:"time_attr"`
	Remarks      string `json:"remarks"`
}

func SetOneUserQuota(params *ParamsOneUserQuotaSet) wx.Action {
	return wx.NewPostAction(urls.CorpOASetOneUserVacationQuota,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}
