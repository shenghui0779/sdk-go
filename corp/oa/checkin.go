package oa

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type CorpCheckinOption struct {
	GroupType              int               `json:"grouptype"`
	GroupID                int64             `json:"groupid"`
	GroupName              string            `json:"groupname"`
	CheckinDate            []*CheckinDate    `json:"checkindate"`
	SpeWorkdays            []*CheckinSpeDay  `json:"spe_workdays"`
	SpeOffDays             []*CheckinSpeDay  `json:"spe_offdays"`
	SyncHolidays           bool              `json:"sync_holidays"`
	NeedPhoto              bool              `json:"need_photo"`
	WifiMacInfos           []*CheckinWifiMac `json:"wifimac_infos"`
	NoteCanUseLocalPic     bool              `json:"note_can_use_local_pic"`
	AllowCheckinOffWorkday bool              `json:"allow_checkin_offworkday"`
	AllowApplyOffWorkday   bool              `json:"allow_apply_offworkday"`
	LocInfos               []*CheckinLoc     `json:"loc_infos"`
	Range                  *CheckinRange     `json:"range"`
	CreateTime             int64             `json:"create_time"`
	WhiteUsers             []string          `json:"white_users"`
	Type                   int               `json:"type"`
	ReporterInfo           *CheckinReporter  `json:"reporterinfo"`
	OtInfo                 *CheckinOt        `json:"ot_info"`
	AllowApplyBKCnt        int               `json:"allow_apply_bk_cnt"`
	OptionOutRange         int               `json:"option_out_range"`
	CreateUserID           string            `json:"create_userid"`
	UseFaceDetect          bool              `json:"use_face_detect"`
	AllowApplyBKDayLimit   int               `json:"allow_apply_bk_day_limit"`
	UpdateUserID           string            `json:"update_userid"`
	ScheduleList           []*OptionSchedule `json:"schedulelist"`
	OffWorkIntervalTime    int               `json:"offwork_interval_time"`
}

type CheckinOption struct {
	UserID string        `json:"userid"`
	Group  *CheckinGroup `json:"group"`
}

type CheckinGroup struct {
	GroupType              int               `json:"grouptype"`
	GroupID                int64             `json:"groupid"`
	GroupName              string            `json:"groupname"`
	CheckinDate            []*CheckinDate    `json:"checkindate"`
	SpeWorkdays            []*CheckinSpeDay  `json:"spe_workdays"`
	SpeOffDays             []*CheckinSpeDay  `json:"spe_offdays"`
	SyncHolidays           bool              `json:"sync_holidays"`
	NeedPhoto              bool              `json:"need_photo"`
	WifiMacInfos           []*CheckinWifiMac `json:"wifimac_infos"`
	NoteCanUseLocalPic     bool              `json:"note_can_use_local_pic"`
	AllowCheckinOffWorkday bool              `json:"allow_checkin_offworkday"`
	AllowApplyOffWorkday   bool              `json:"allow_apply_offworkday"`
	LocInfos               []*CheckinLoc     `json:"loc_infos"`
	ScheduleList           []*OptionSchedule `json:"schedulelist"`
}

type CheckinDate struct {
	Workdays        []int          `json:"workdays"`
	CheckinTime     []*CheckinTime `json:"checkintime"`
	FlexTime        int            `json:"flex_time"`
	NoNeedOffWork   bool           `json:"noneed_offwork"`
	LimitAheadTime  int            `json:"limit_aheadtime"`
	FlexOnDutyTime  int            `json:"flex_on_duty_time"`
	FlexOffDutyTime int            `json:"flex_off_duty_time"`
}

type CheckinSpeDay struct {
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

type CheckinWifiMac struct {
	WifiName string `json:"wifiname"`
	WifiMac  string `json:"wifimac"`
}

type CheckinLoc struct {
	Lat       int64  `json:"lat"`
	Lng       int64  `json:"lng"`
	LocTitle  string `json:"loc_title"`
	LocDetail string `json:"loc_detail"`
	Distance  int    `json:"distance"`
}

type CheckinRange struct {
	PartyID []string `json:"partyid"`
	UserID  []string `json:"userid"`
	TagID   []int64  `json:"tagid"`
}

type CheckinReporter struct {
	Reporters  []*OAUser `json:"reporters"`
	UpdateTime int64     `json:"updatetime"`
}

type CheckinOt struct {
	Type                 int             `json:"type"`
	AllowOtWorkingDay    bool            `json:"allow_ot_workingday"`
	AllowOtNonWorkingDay bool            `json:"allow_ot_nonworkingday"`
	OtCheckInfo          *CheckinOtCheck `json:"otcheckinfo"`
	UpTime               int64           `json:"uptime"`
	OtApplyInfo          *CheckinOtApply `json:"otapplyinfo"`
}

type CheckinOtCheck struct {
	OtWorkingDayTimeStart      int            `json:"ot_workingday_time_start"`
	OtWorkingDayTimeMin        int            `json:"ot_workingday_time_min"`
	OtWorkingDayTimeMax        int            `json:"ot_workingday_time_max"`
	OtNonWorkingDayTimeMin     int            `json:"ot_nonworkingday_time_min"`
	OtNonWorkingDayTimeMax     int            `json:"ot_nonworkingday_time_max"`
	OtWorkingDayRestInfo       *CheckinOtRest `json:"ot_workingday_restinfo"`
	OtNonWorkingDayRestInfo    *CheckinOtRest `json:"ot_nonworkingday_restinfo"`
	OtNonWorkingDaySpandayTime int            `json:"ot_nonworkingday_spanday_time"`
}

type CheckinOtApply struct {
	AllowOtWorkingDay          bool           `json:"allow_ot_workingday"`
	AllowOtNonWorkingDay       bool           `json:"allow_ot_nonworkingday"`
	UpTime                     int64          `json:"uptime"`
	OtWorkingDayRestInfo       *CheckinOtRest `json:"ot_workingday_restinfo"`
	OtNonWorkingDayRestInfo    *CheckinOtRest `json:"ot_nonworkingday_restinfo"`
	OtNonWorkingDaySpandayTime int            `json:"ot_nonworkingday_spanday_time"`
}

type CheckinOtRest struct {
	Type          int                   `json:"type"`
	FixTimeRule   *CheckinFixTimeRule   `json:"fix_time_rule"`
	CalOtTimeRule *CheckinCalOtTimeRule `json:"cal_ottime_rule"`
}

type CheckinFixTimeRule struct {
	FixTimeBeginSec int `json:"fix_time_begin_sec"`
	FixTimeEndSec   int `json:"fix_time_end_sec"`
}

type CheckinCalOtTimeRule struct {
	Items []*CheckinOtTimeRule `json:"items"`
}

type CheckinOtTimeRule struct {
	OtTime   int `json:"ot_time"`
	RestTime int `json:"rest_time"`
}

type OptionSchedule struct {
	ScheduleID          int64                 `json:"schedule_id"`
	ScheduleName        string                `json:"schedule_name"`
	TimeSection         []*CheckinTimeSection `json:"time_section"`
	LimitAheadTime      int                   `json:"limit_aheadtime"`
	NoNeedOffWork       bool                  `json:"noneed_offwork"`
	LimitOffTime        int                   `json:"limit_offtime"`
	FlexOnDutyTime      int                   `json:"flex_on_duty_time"`
	FlexOffDutyTime     int                   `json:"flex_off_duty_time"`
	AllowFlex           bool                  `json:"allow_flex"`
	LateRule            *CheckinLateRule      `json:"late_rule"`
	MaxAllowArriveEarly int                   `json:"max_allow_arrive_early"`
	MaxAllowArriveLate  int                   `json:"max_allow_arrive_late"`
}

type CheckinTimeSection struct {
	TimeID           int64 `json:"time_id"`
	WorkSec          int   `json:"work_sec"`
	OffWorkSec       int   `json:"off_work_sec"`
	RemindWorkSec    int   `json:"remind_work_sec"`
	RemindOffWorkSec int   `json:"remind_off_work_sec"`
	RestBeginTime    int   `json:"rest_begin_time"`
	RestEndTime      int   `json:"rest_end_time"`
	AllowRest        bool  `json:"allow_rest"`
}

type CheckinLateRule struct {
	AllowOffWorkAfterTime bool               `json:"allow_offwork_after_time"`
	TimeRules             []*CheckinTimeRule `json:"timerules"`
}

type CheckinTimeRule struct {
	OffWorkAfterTime int `json:"offwork_after_time"`
	OnWorkFlexTime   int `json:"onwork_flex_time"`
}

type ResultCorpCheckinOption struct {
	Group []*CorpCheckinOption `json:"group"`
}

func GetCorpCheckinOption(result *ResultCorpCheckinOption) wx.Action {
	return wx.NewPostAction(urls.CorpOAGetCorpCheckinOption,
		wx.WithBody(func() ([]byte, error) {
			return []byte("{}"), nil
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsCheckinOption struct {
	DateTime   int64    `json:"datetime"`
	UserIDList []string `json:"useridlist"`
}

type ResultCheckinOption struct {
	Info []*CheckinOption `json:"info"`
}

func GetCheckinOption(datetime int64, userIDs []string, result *ResultCheckinOption) wx.Action {
	params := &ParamsCheckinOption{
		DateTime:   datetime,
		UserIDList: userIDs,
	}

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
	Lat            int64    `json:"lat"`
	Lng            int64    `json:"lng"`
	DeviceID       string   `json:"deviceid"`
	SchCheckinTime int64    `json:"sch_checkin_time"`
	ScheduleID     int64    `json:"schedule_id"`
	TimelineID     int64    `json:"timeline_id"`
}

type ParamsCheckinData struct {
	OpenCheckinDataType int      `json:"opencheckindatatype"`
	StartTime           int64    `json:"starttime"`
	EndTime             int64    `json:"endtime"`
	UserIDList          []string `json:"useridlist"`
}

type ResultCheckinData struct {
	CheckinData []*CheckinData `json:"checkindata"`
}

func GetCheckinData(dataType int, starttime, endtime int64, userIDs []string, result *ResultCheckinData) wx.Action {
	params := &ParamsCheckinData{
		OpenCheckinDataType: dataType,
		StartTime:           starttime,
		EndTime:             endtime,
		UserIDList:          userIDs,
	}

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
	BaseInfo       *CheckinDayBase     `json:"base_info"`
	SummaryInfo    *CheckinDaySummary  `json:"summary_info"`
	HolidayInfos   []*CheckinHoliday   `json:"holiday_infos"`
	ExceptionInfos []*CheckinException `json:"exception_infos"`
	OtInfo         *CheckinDayOt       `json:"ot_info"`
	SPItems        []*CheckinSPItem    `json:"sp_items"`
}

type CheckinMonthData struct {
	BaseInfo       *CheckinMonthBase    `json:"base_info"`
	SummaryInfo    *CheckinMonthSummary `json:"summary_info"`
	ExceptionInfos []*CheckinException  `json:"exception_infos"`
	SPItems        []*CheckinSPItem     `json:"sp_items"`
	OverworkInfo   *CheckinOverwork     `json:"overwork_info"`
}

type CheckinDayBase struct {
	Date        int64           `json:"date"`
	RecordType  int             `json:"record_type"`
	Name        string          `json:"name"`
	NameEX      string          `json:"name_ex"`
	DepartsName string          `json:"departs_name"`
	AcctID      string          `json:"acctid"`
	RuleInfo    *CheckinDayRule `json:"rule_info"`
}

type CheckinMonthBase struct {
	RecordType  int               `json:"record_type"`
	Name        string            `json:"name"`
	NameEX      string            `json:"name_ex"`
	DepartsName string            `json:"departs_name"`
	AcctID      string            `json:"acctid"`
	RuleInfo    *CheckinMonthRule `json:"rule_info"`
}

type CheckinDayRule struct {
	GroupID      int64          `json:"groupid"`
	GroupName    string         `json:"groupname"`
	ScheduleID   int64          `json:"scheduleid"`
	ScheduleName string         `json:"schedulename"`
	CheckinTime  []*CheckinTime `json:"checkintime"`
	DayType      int            `json:"day_type"`
}

type CheckinMonthRule struct {
	GroupID   int64  `json:"groupid"`
	GroupName string `json:"groupname"`
}

type CheckinDaySummary struct {
	CheckinCount    int `json:"checkin_count"`
	RegularWorkSec  int `json:"regular_work_sec"`
	StandardWorkSec int `json:"standard_work_sec"`
	EarliestTime    int `json:"earliest_time"`
	LastestTime     int `json:"lastest_time"`
}

type CheckinMonthSummary struct {
	WorkDays        int `json:"work_days"`
	ExceptDays      int `json:"except_days"`
	RegularWorkSec  int `json:"regular_work_sec"`
	StandardWorkSec int `json:"standard_work_sec"`
}

type CheckinHoliday struct {
	SPDescription *CheckinSPText `json:"sp_description"`
	SPNumber      string         `json:"sp_number"`
	SPTitle       *CheckinSPText `json:"sp_title"`
}

type CheckinSPText struct {
	Data []*DisplayText `json:"data"`
}

type CheckinException struct {
	Count     int `json:"count"`
	Duration  int `json:"duration"`
	Exception int `json:"exception"`
}

type CheckinDayOt struct {
	OtStatus          int   `json:"ot_status"`
	OtDuration        int   `json:"ot_duration"`
	ExceptionDuration []int `json:"exception_duration"`
}

type CheckinSPItem struct {
	Count      int    `json:"count"`
	Duration   int    `json:"duration"`
	TimeType   int    `json:"time_type"`
	Type       int    `json:"type"`
	VacationID int    `json:"vacation_id"`
	Name       string `json:"name"`
}

type CheckinOverwork struct {
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

func GetCheckinDayData(starttime, endtime int64, userIDs []string, result *ResultCheckinDayData) wx.Action {
	params := &ParamsCheckinDayData{
		StartTime:  starttime,
		EndTime:    endtime,
		UserIDList: userIDs,
	}

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

func GetCheckinMonthData(starttime, endtime int64, userIDs []string, result *ResultCheckinMonthData) wx.Action {
	params := &ParamsCheckinMonthData{
		StartTime:  starttime,
		EndTime:    endtime,
		UserIDList: userIDs,
	}

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
	ScheduleList []*CheckinSchedule `json:"schedule_list"`
}

type CheckinSchedule struct {
	UserID    string    `json:"userid"`
	YearMonth int       `json:"yearmonth"`
	GroupID   int64     `json:"groupid"`
	GroupName string    `json:"groupname"`
	Schedule  *Schedule `json:"schedule"`
}

type Schedule struct {
	ScheduleList []*ScheduleData `json:"scheduleList"`
}

type ScheduleData struct {
	Day          int           `json:"day"`
	ScheduleInfo *ScheduleInfo `json:"schedule_info"`
}

type ScheduleInfo struct {
	ScheduleID   int64                  `json:"schedule_id"`
	ScheduleName string                 `json:"schedule_name"`
	TimeSection  []*ScheduleTimeSection `json:"time_section"`
}

type ScheduleTimeSection struct {
	ID               int64 `json:"id"`
	WorkSec          int   `json:"work_sec"`
	OffWorkSec       int   `json:"off_work_sec"`
	RemindWorkSec    int   `json:"remind_work_sec"`
	RemindOffWorkSec int   `json:"remind_off_work_sec"`
}

func GetCheckinScheduleList(starttime, endtime int64, userIDs []string, result *ResultCheckinScheduleListGet) wx.Action {
	params := &ParamsCheckinScheduleListGet{
		StartTime:  starttime,
		EndTime:    endtime,
		UserIDList: userIDs,
	}

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
	GroupID   int64               `json:"groupid"`
	YearMonth int                 `json:"yearmonth"`
	Items     []*ScheduleListItem `json:"items"`
}

type ScheduleListItem struct {
	UserID     string `json:"userid"`
	Day        int    `json:"day"`
	ScheduleID int64  `json:"schedule_id"`
}

func SetCheckinScheduleList(groupID int64, yearmonth int, items ...*ScheduleListItem) wx.Action {
	params := &ParamsCheckinScheduleListSet{
		GroupID:   groupID,
		YearMonth: yearmonth,
		Items:     items,
	}

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

func AddCheckinUserFace(userID, userFace string) wx.Action {
	params := &ParamsCheckinUserFaceAdd{
		UserID:   userID,
		UserFace: userFace,
	}

	return wx.NewPostAction(urls.CorpOAAddCheckinUserFace,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsHardwareCheckinData struct {
	FilterType int      `json:"filter_type"`
	StartTime  int64    `json:"starttime"`
	EndTime    int64    `json:"endtime"`
	UserIDList []string `json:"useridlist"`
}

type ResultHardwareCheckinData struct {
	CheckinData []*HardwareCheckinData
}

type HardwareCheckinData struct {
	UserID      string `json:"userid"`
	CheckinTime int64  `json:"checkin_time"`
	DeviceSN    string `json:"device_sn"`
	DeviceName  string `json:"device_name"`
}

func GetHardwareCheckinData(filterType int, starttime, endtime int64, userIDs []string, result *ResultHardwareCheckinData) wx.Action {
	params := &ParamsHardwareCheckinData{
		FilterType: filterType,
		StartTime:  starttime,
		EndTime:    endtime,
		UserIDList: userIDs,
	}

	return wx.NewPostAction(urls.CorpOAGetHardwareCheckinData,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
