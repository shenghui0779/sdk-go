package oa

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsJournalRecordList struct {
	StartTime int64       `json:"starttime"`
	EndTime   int64       `json:"endtime"`
	Cursor    int         `json:"cursor"`
	Limit     int         `json:"limit"`
	Filters   []*KeyValue `json:"filters,omitempty"`
}

type ResultJournalRecordList struct {
	JournalUUIDList []string `json:"journaluuid_list"`
	NextCursor      int      `json:"next_cursor"`
	EndFlag         int      `json:"endflag"`
}

func ListJouralRecord(params *ParamsJournalRecordList, result *ResultJournalRecordList) wx.Action {
	return wx.NewPostAction(urls.CorpOAGetJournalRecordList,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsJournalRecordDetail struct {
	JournalUUID string `json:"journaluuid"`
}

type ResultJournalRecordDetail struct {
	Info *JournalRecordDetail `json:"info"`
}

type JournalRecordDetail struct {
	JournalUUID     string                  `json:"journal_uuid"`
	TemplateName    string                  `json:"template_name"`
	ReportTime      int64                   `json:"report_time"`
	Submitter       *OAUser                 `json:"submitter"`
	Receivers       []*OAUser               `json:"receivers"`
	ReadedReceivers []*OAUser               `json:"readed_receivers"`
	ApplyData       *ApplyData              `json:"apply_data"`
	Comments        []*JournalRecordComment `json:"comments"`
}

type JournalRecordComment struct {
	CommentID       int64   `json:"commentid"`
	ToCommentID     int64   `json:"tocommentid"`
	CommentUserInfo *OAUser `json:"comment_userinfo"`
	Content         string  `json:"content"`
	CommentTime     int64   `json:"comment_time"`
}

func GetJournalRecordDetail(journaluuid string, result *ResultJournalRecordDetail) wx.Action {
	params := &ParamsJournalRecordDetail{
		JournalUUID: journaluuid,
	}

	return wx.NewPostAction(urls.CorpOAGetJournalRecordDetail,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsJournalStatList struct {
	TemplateID string `json:"template_id"`
	StartTime  int64  `json:"starttime"`
	EndTime    int64  `json:"endtime"`
}

type ResultJournalStatList struct {
	StatList []*JournalStat `json:"stat_list"`
}

type JournalStat struct {
	TemplateID     string            `json:"template_id"`
	TemplateName   string            `json:"template_name"`
	ReportRange    *JournalRange     `json:"report_range"`
	WhiteRange     *JournalRange     `json:"white_range"`
	Receivers      *JournalReceivers `json:"receivers"`
	CycleBeginTime int64             `json:"cycle_begin_time"`
	CycleEndTime   int64             `json:"cycle_end_time"`
	StatBeginTime  int64             `json:"stat_begin_time"`
	StatEndTime    int64             `json:"stat_end_time"`
	ReportList     []*JournalReport  `json:"report_list"`
	UnReportList   []*JournalReport  `json:"unreport_list"`
	ReportType     int               `json:"report_type"`
}

type JournalRange struct {
	UserList  []*JournalUser  `json:"user_list"`
	PartyList []*JournalParty `json:"party_list"`
	TagList   []*JournalTag   `json:"tag_list"`
}

type JournalReceivers struct {
	UserList   []*JournalUser   `json:"user_list"`
	TagList    []*JournalTag    `json:"tag_list"`
	LeaderList []*JournalLeader `json:"leader_list"`
}

type JournalUser struct {
	UserID string `json:"userid"`
}

type JournalParty struct {
	OpenPartyID string `json:"open_partyid"`
}

type JournalTag struct {
	OpenTagID string `json:"open_tagid"`
}

type JournalLeader struct {
	Level int64 `json:"level"`
}

type JournalReport struct {
	User     *OAUser              `json:"user"`
	ItemList []*JournalReportItem `json:"itemlist"`
}

type JournalReportItem struct {
	JournalUUID string `json:"journaluuid"`
	ReportTime  int64  `json:"reporttime"`
	Flag        int    `json:"flag"`
}

func ListJournalStat(templateID string, starttime, endtime int64, result *ResultJournalStatList) wx.Action {
	params := &ParamsJournalStatList{
		TemplateID: templateID,
		StartTime:  starttime,
		EndTime:    endtime,
	}

	return wx.NewPostAction(urls.CorpOAGetJournalStatList,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
