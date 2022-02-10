package tools

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type DialRecord struct {
	CallTime      int64         `json:"call_time"`
	TotalDuration int           `json:"total_duration"`
	CallType      int           `json:"call_type"`
	Caller        *DialCaller   `json:"caller"`
	Callee        []*DialCallee `json:"callee"`
}

type DialCaller struct {
	UserID   string `json:"userid"`
	Duration int    `json:"duration"`
}

type DialCallee struct {
	UserID   string `json:"userid"`
	Phone    string `json:"phone"`
	Duration int    `json:"duration"`
}

type ParamsDialRecord struct {
	StartTime int64 `json:"start_time,omitempty"`
	EndTime   int64 `json:"end_time,omitempty"`
	Offset    int   `json:"offset,omitempty"`
	Limit     int   `json:"limit,omitempty"`
}

type ResultDialRecord struct {
	Record []*DialRecord `json:"record"`
}

func GetDialRecord(params *ParamsDialRecord, result *ResultDialRecord) wx.Action {
	return wx.NewPostAction(urls.CorpToolsDialRecordGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
