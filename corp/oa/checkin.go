package oa

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type DisplayText struct {
	Text string `json:"text"`
	Lang string `json:"lang"`
}

type CorpCheckinOption struct {
	GroupType              int               `json:"grouptype"`
	GroupID                int64             `json:"groupid"`
	GroupName              string            `json:"groupname"`
	CheckinDate            *CheckinDate      `json:"checkindate"`
	SpeWorkdays            []*SpeDay         `json:"spe_workdays"`
	SpeOffDays             []*SpeDay         `json:"spe_offdays"`
	SyncHolidays           bool              `json:"sync_holidays"`
	NeedPhoto              bool              `json:"need_photo"`
	WifiMacInfos           []*WifiMacInfo    `json:"wifimac_infos"`
	NoteCanUseLocalPic     bool              `json:"note_can_use_local_pic"`
	AllowCheckinOffWorkday bool              `json:"allow_checkin_offworkday"`
	AllowApplyOffWorkday   bool              `json:"allow_apply_offworkday"`
	LocInfos               []*LocInfo        `json:"loc_infos"`
	Range                  *Range            `json:"range"`
	CreateTime             int64             `json:"create_time"`
	WhiteUsers             []string          `json:"white_users"`
	Type                   int               `json:"type"`
	ReporterInfo           *ReporterInfo     `json:"reporterinfo"`
	OtInfo                 *OptionOtInfo     `json:"ot_info"`
	AllowApplyBKCnt        int               `json:"allow_apply_bk_cnt"`
	AllowApplyBKDayLimit   int               `json:"allow_apply_bk_day_limit"`
	OptionOutRange         int               `json:"option_out_range"`
	CreateUserID           string            `json:"create_userid"`
	UpdateUserID           string            `json:"update_userid"`
	UseFaceDetect          bool              `json:"use_face_detect"`
	ScheduleList           []*OptionSchedule `json:"schedulelist"`
	OffWorkIntervalTime    int               `json:"off_work_interval_time"`
}

type CheckinInfo struct {
	UserID string           `json:"userid"`
	Group  []*CheckinOption `json:"group"`
}

type CheckinOption struct {
	GroupType              int               `json:"grouptype"`
	GroupID                int64             `json:"groupid"`
	GroupName              string            `json:"groupname"`
	CheckinDate            *CheckinDate      `json:"checkindate"`
	SpeWorkdays            []*SpeDay         `json:"spe_workdays"`
	SpeOffDays             []*SpeDay         `json:"spe_offdays"`
	SyncHolidays           bool              `json:"sync_holidays"`
	NeedPhoto              bool              `json:"need_photo"`
	WifiMacInfos           []*WifiMacInfo    `json:"wifimac_infos"`
	NoteCanUseLocalPic     bool              `json:"note_can_use_local_pic"`
	AllowCheckinOffWorkday bool              `json:"allow_checkin_offworkday"`
	AllowApplyOffWorkday   bool              `json:"allow_apply_offworkday"`
	LocInfos               []*LocInfo        `json:"loc_infos"`
	ScheduleList           []*OptionSchedule `json:"schedulelist"`
}

type CheckinDate struct {
	Workdays        []int          `json:"workdays"`
	CheckinTime     []*CheckinTime `json:"checkintime"`
	NoNeedOffWork   bool           `json:"noneed_offwork"`
	LimitAheadTime  int            `json:"limit_aheadtime"`
	FlexOnDutyTime  int            `json:"flex_on_duty_time"`
	FlexOffDutyTime int            `json:"flex_off_duty_time"`
}

type SpeDay struct {
	Timestamp   int64          `json:"timestamp"`
	Notes       string         `json:"notes"`
	CheckinTime []*CheckinTime `json:"checkintime"`
}

type CheckinTime struct {
	WorkSec          int `json:"work_sec"`
	OffWorkSec       int `json:"off_work_sec"`
	RemindWorkSec    int `json:"remind_work_sec"`
	RemindOffWorkSec int `json:"remind_off_work_sec"`
}

type WifiMacInfo struct {
	WifiName string `json:"wifiname"`
	WifiMac  string `json:"wifimac"`
}

type LocInfo struct {
	Lat       int64  `json:"lat"`
	Lng       int64  `json:"lng"`
	LocTitle  string `json:"loc_title"`
	LocDetail string `json:"loc_detail"`
	Distance  int    `json:"distance"`
}

type Range struct {
	PartyID []string `json:"partyid"`
	UserID  []string `json:"userid"`
	TagID   []int64  `json:"tagid"`
}

type ReporterInfo struct {
	Reporters  []*Reporter `json:"reporters"`
	UpdateTime int64       `json:"updatetime"`
}

type Reporter struct {
	UserID string `json:"userid"`
}

type OptionOtInfo struct {
	Type                 int          `json:"type"`
	AllowOtWorkingDay    bool         `json:"allow_ot_workingday"`
	AllowOtNonWorkingDay bool         `json:"allow_ot_nonworkingday"`
	OtCheckInfo          *OtCheckInfo `json:"otcheckinfo"`
	OtApplyInfo          *OtApplyInfo `json:"otapplyinfo"`
	UpTime               int64        `json:"uptime"`
}

type OtCheckInfo struct {
	OtWorkingDayTimeStart      int         `json:"ot_workingday_time_start"`
	OtWorkingDayTimeMin        int         `json:"ot_workingday_time_min"`
	OtWorkingDayTimeMax        int         `json:"ot_workingday_time_max"`
	OtNonWorkingDayTimeMin     int         `json:"ot_nonworkingday_time_min"`
	OtNonWorkingDayTimeMax     int         `json:"ot_nonworkingday_time_max"`
	OtWorkingDayRestInfo       *OtRestInfo `json:"ot_workingday_restinfo"`
	OtNonWorkingDayRestInfo    *OtRestInfo `json:"ot_nonworkingday_restinfo"`
	OtNonWorkingDaySpandayTime int         `json:"ot_nonworkingday_spanday_time"`
}

type OtApplyInfo struct {
	AllowOtWorkingDay          bool        `json:"allow_ot_workingday"`
	AllowOtNonWorkingDay       bool        `json:"allow_ot_nonworkingday"`
	UpTime                     int64       `json:"uptime"`
	OtWorkingDayRestInfo       *OtRestInfo `json:"ot_workingday_restinfo"`
	OtNonWorkingDayRestInfo    *OtRestInfo `json:"ot_nonworkingday_restinfo"`
	OtNonWorkingDaySpandayTime int         `json:"ot_nonworkingday_spanday_time"`
}

type OtRestInfo struct {
	Type          int            `json:"type"`
	FixTimeRule   *FixTimeRule   `json:"fix_time_rule"`
	CalOtTimeRule *CalOtTimeRule `json:"cal_ottime_rule"`
}

type FixTimeRule struct {
	FixTimeBeginSec int `json:"fix_time_begin_sec"`
	FixTimeEndSec   int `json:"fix_time_end_sec"`
}

type CalOtTimeRule struct {
	Items []*OtTimeRuleItem `json:"items"`
}

type OtTimeRuleItem struct {
	OtTime   int `json:"ot_time"`
	RestTime int `json:"rest_time"`
}

type OptionSchedule struct {
	ScheduleID          int64          `json:"schedule_id"`
	ScheduleName        string         `json:"schedule_name"`
	TimeSection         []*TimeSection `json:"time_section"`
	LimitAheadTime      int            `json:"limit_aheadtime"`
	NoNeedOffWork       bool           `json:"noneed_offwork"`
	LimitOffTime        int            `json:"limit_off_time"`
	FlexOnDutyTime      int            `json:"flex_on_duty_time"`
	FlexOffDutyTime     int            `json:"flex_off_duty_time"`
	AllowFlex           bool           `json:"allow_flex"`
	LateRule            *LateRule      `json:"late_rule"`
	MaxAllowArriveEarly int            `json:"max_allow_arrive_early"`
	MaxAllowArriveLate  int            `json:"max_allow_arrive_late"`
}

type TimeSection struct {
	TimeID           int64 `json:"time_id"`
	WorkSec          int   `json:"work_sec"`
	OffWorkSec       int   `json:"off_work_sec"`
	RemindWorkSec    int   `json:"remind_work_sec"`
	RemindOffWorkSec int   `json:"remind_off_work_sec"`
	RestBeginTime    int   `json:"rest_begin_time"`
	RestEndTime      int   `json:"rest_end_time"`
	AllowRest        bool  `json:"allow_rest"`
}

type LateRule struct {
	AllowOffWorkAfterTime bool        `json:"allow_offwork_after_time"`
	TimeRules             []*TimeRule `json:"timerules"`
}

type TimeRule struct {
	OffWorkAfterTime int `json:"offwork_after_time"`
	OnWorkFlexTime   int `json:"onwork_flex_time"`
}

type ResultCorpCheckinOption struct {
	Group []*CorpCheckinOption `json:"group"`
}

func GetCorpCheckinOption(result *ResultCorpCheckinOption) wx.Action {
	return wx.NewPostAction(urls.CorpOAGetCorpCheckinOption,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsCheckinOptionGet struct {
	DateTime   int64    `json:"datetime"`
	UserIDList []string `json:"useridlist"`
}

type ResultCheckinOptionGet struct {
	Info *CheckinInfo `json:"info"`
}

func GetCheckinOption(params *ParamsCheckinOptionGet, result *ResultCheckinOptionGet) wx.Action {
	return wx.NewPostAction(urls.CorpOAGetCheckinOption,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type CheckinData struct {
	UserID         string   `json:"userid"`
	GroupID        int64    `json:"groupid"`
	GroupName      string   `json:"groupname"`
	CheckinType    string   `json:"checkin_type"`
	CheckinTime    int64    `json:"checkin_time"`
	ExceptionType  string   `json:"exception_type"`
	LocationTitle  string   `json:"location_title"`
	LocationDetail string   `json:"location_detail"`
	WifiName       string   `json:"wifiname"`
	WifiMac        string   `json:"wifimac"`
	Notes          string   `json:"notes"`
	MediaIDs       []string `json:"mediaids"`
	SchCheckinTime int64    `json:"sch_checkin_time"`
	ScheduleID     int64    `json:"schedule_id"`
	TimelineID     int64    `json:"timeline_id"`
}

type ParamsCheckinDataGet struct {
	OpenCheckinDataType int      `json:"opencheckindatatype"`
	StartTime           int64    `json:"starttime"`
	EndTime             int64    `json:"endtime"`
	UserIDList          []string `json:"useridlist"`
}

type ResultCheckinDataGet struct {
	CheckinData []*CheckinData `json:"checkindata"`
}

func GetCheckinData(params *ParamsCheckinDataGet, result *ResultCheckinDataGet) wx.Action {
	return wx.NewPostAction(urls.CorpOAGetCheckinData,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type CheckinDayData struct {
	BaseInfo       *DayBaseInfo     `json:"base_info"`
	SummaryInfo    *DaySummaryInfo  `json:"summary_info"`
	HolidayInfos   []*HolidayInfo   `json:"holiday_infos"`
	ExceptionInfos []*ExceptionInfo `json:"exception_infos"`
	OtInfo         *DayOtInfo       `json:"ot_info"`
	SPItems        []*SPItem        `json:"sp_items"`
}

type CheckinMonthData struct {
	BaseInfo       *MonthBaseInfo    `json:"base_info"`
	SummaryInfo    *MonthSummaryInfo `json:"summary_info"`
	ExceptionInfos []*ExceptionInfo  `json:"exception_infos"`
	SPItems        []*SPItem         `json:"sp_items"`
	OverworkInfo   *OverworkInfo     `json:"overwork_info"`
}

type DayBaseInfo struct {
	Date        int64        `json:"date"`
	RecordType  int          `json:"record_type"`
	Name        string       `json:"name"`
	NameEX      string       `json:"name_ex"`
	DepartsName string       `json:"departs_name"`
	AcctID      string       `json:"acctid"`
	RuleInfo    *DayRuleInfo `json:"rule_info"`
}

type MonthBaseInfo struct {
	RecordType  int            `json:"record_type"`
	Name        string         `json:"name"`
	NameEX      string         `json:"name_ex"`
	DepartsName string         `json:"departs_name"`
	AcctID      string         `json:"acctid"`
	RuleInfo    *MonthRuleInfo `json:"rule_info"`
}

type DayRuleInfo struct {
	GroupID      int64          `json:"groupid"`
	GroupName    string         `json:"groupname"`
	ScheduleID   int64          `json:"scheduleid"`
	ScheduleName string         `json:"schedulename"`
	CheckinTime  []*CheckinTime `json:"checkintime"`
	DayType      int            `json:"day_type"`
}

type MonthRuleInfo struct {
	GroupID   int64  `json:"groupid"`
	GroupName string `json:"groupname"`
}

type DaySummaryInfo struct {
	CheckinCount    int `json:"checkin_count"`
	RegularWorkSec  int `json:"regular_work_sec"`
	StandardWorkSec int `json:"standard_work_sec"`
	EarliestTime    int `json:"earliest_time"`
	LastestTime     int `json:"lastest_time"`
}

type MonthSummaryInfo struct {
	WorkDays        int `json:"work_days"`
	ExceptDays      int `json:"except_days"`
	RegularWorkSec  int `json:"regular_work_sec"`
	StandardWorkSec int `json:"standard_work_sec"`
}

type HolidayInfo struct {
	SPDescription *SPDescription `json:"sp_description"`
	SPNumber      string         `json:"sp_number"`
	SPTitle       *SPTitle       `json:"sp_title"`
}

type SPDescription struct {
	Data []*DisplayText `json:"data"`
}

type SPTitle struct {
	Data []*DisplayText `json:"data"`
}

type ExceptionInfo struct {
	Count     int `json:"count"`
	Duration  int `json:"duration"`
	Exception int `json:"exception"`
}

type DayOtInfo struct {
	OtStatus          int `json:"ot_status"`
	OtDuration        int `json:"ot_duration"`
	ExceptionDuration int `json:"exception_duration"`
}

type SPItem struct {
	Count      int    `json:"count"`
	Duration   int    `json:"duration"`
	TimeType   int    `json:"time_type"`
	Type       int    `json:"type"`
	VacationID int    `json:"vacation_id"`
	Name       string `json:"name"`
}

type OverworkInfo struct {
	WorkdayOverSec int `json:"workday_over_sec"`
}

type ParamsCheckinDayData struct {
	StartTime  int64    `json:"starttime"`
	EndTime    int64    `json:"endtime"`
	UserIDList []string `json:"useridlist"`
}

type ResultCheckinDayData struct {
	Datas []*CheckinDayData `json:"datas"`
}

func GetCheckinDayData(params *ParamsCheckinDayData, result *ResultCheckinDayData) wx.Action {
	return wx.NewPostAction(urls.CorpOAGetCheckinDayData,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsCheckinMonthData struct {
	StartTime  int64    `json:"starttime"`
	EndTime    int64    `json:"endtime"`
	UserIDList []string `json:"useridlist"`
}

type ResultCheckinMonthData struct {
	Datas []*CheckinMonthData `json:"datas"`
}

func GetCheckinMonthData(params *ParamsCheckinMonthData, result *ResultCheckinMonthData) wx.Action {
	return wx.NewPostAction(urls.CorpOAGetCheckinMonthData,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsCheckinScheduleListGet struct {
	StartTime  int64    `json:"starttime"`
	EndTime    int64    `json:"endtime"`
	UserIDList []string `json:"useridlist"`
}

type ResultCheckinScheduleListGet struct {
	ScheduleList []*UserCheckinSchedule `json:"schedule_list"`
}

type UserCheckinSchedule struct {
	UserID    string           `json:"userid"`
	YearMonth int              `json:"yearmonth"`
	GroupID   int64            `json:"groupid"`
	GroupName string           `json:"groupname"`
	Schedule  *CheckinSchedule `json:"schedule"`
}

type CheckinSchedule struct {
	ScheduleList []*Schedule `json:"schedule_list"`
}

type Schedule struct {
	Day          int           `json:"day"`
	ScheduleInfo *ScheduleInfo `json:"schedule_info"`
}

type ScheduleInfo struct {
	ScheduleID   int64                `json:"schedule_id"`
	ScheduleName string               `json:"schedule_name"`
	TimeSection  *ScheduleTimeSection `json:"time_section"`
}

type ScheduleTimeSection struct {
	ID               int64 `json:"id"`
	WorkSec          int   `json:"work_sec"`
	OffWorkSec       int   `json:"off_work_sec"`
	RemindWorkSec    int   `json:"remind_work_sec"`
	RemindOffWorkSec int   `json:"remind_off_work_sec"`
}

func GetCheckinScheduleList(params *ParamsCheckinScheduleListGet, result *ResultCheckinScheduleListGet) wx.Action {
	return wx.NewPostAction(urls.CorpOAGetCheckinScheduleList,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsCheckinScheduleListSet struct {
	GroupID   int64           `json:"groupid"`
	YearMonth int             `json:"yearmonth"`
	Items     []*ScheduleItem `json:"items"`
}

type ScheduleItem struct {
	UserID     string `json:"userid"`
	Day        int    `json:"day"`
	ScheduleID int64  `json:"schedule_id"`
}

func SetCheckinScheduleList(params *ParamsCheckinScheduleListSet) wx.Action {
	return wx.NewPostAction(urls.CorpOASetCheckinScheduleList,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsCheckinUserFaceAdd struct {
	UserID   string `json:"userid"`
	UserFace string `json:"userface"`
}

func AddCheckinUserFace(params *ParamsCheckinUserFaceAdd) wx.Action {
	return wx.NewPostAction(urls.CorpOAAddCheckinUserFace,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}
