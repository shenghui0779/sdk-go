package tools

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsAttendees struct {
	UserID         []string `json:"userid,omitempty"`
	ExternalUserID []string `json:"external_userid,omitempty"`
	DeviceSN       []string `json:"device_sn,omitempty"`
}

type ParamsMeetingCreate struct {
	CreatorUserID   string           `json:"creator_userid"`
	Title           string           `json:"title"`
	MeetingStart    int64            `json:"meeting_start"`
	MeetingDuration int              `json:"meeting_duration"`
	Description     string           `json:"description,omitempty"`
	Type            int              `json:"type"`
	RemindTime      int              `json:"remind_time,omitempty"`
	AgentID         int64            `json:"agentid,omitempty"`
	Attendees       *ParamsAttendees `json:"attendees,omitempty"`
}

type ResultMeetingCreate struct {
	MeetingID string `json:"meetingid"`
}

func CreateMeeting(params *ParamsMeetingCreate, result *ResultMeetingCreate) wx.Action {
	return wx.NewPostAction(urls.CorpToolsMeetingCreate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsMeetingUpdate struct {
	MeetingID       string           `json:"meetingid"`
	Title           string           `json:"title,omitempty"`
	MeetingStart    int64            `json:"meeting_start,omitempty"`
	MeetingDuration int              `json:"meeting_duration,omitempty"`
	Description     string           `json:"description,omitempty"`
	Type            int              `json:"type,omitempty"`
	RemindTime      int              `json:"remind_time,omitempty"`
	AgentID         int64            `json:"agentid,omitempty"`
	Attendees       *ParamsAttendees `json:"attendees,omitempty"`
}

func UpdateMeeting(params *ParamsMeetingUpdate) wx.Action {
	return wx.NewPostAction(urls.CorpToolsMeetingUpdate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsMeetingCancel struct {
	MeetingID string `json:"meetingid"`
}

func CancelMeeting(params *ParamsMeetingCancel) wx.Action {
	return wx.NewPostAction(urls.CorpToolsMeetingCancel,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsUserMeetingIDGet struct {
	UserID    string `json:"userid"`
	Cursor    string `json:"cursor,omitempty"`
	BeginTime int64  `json:"begin_time,omitempty"`
	EndTime   int64  `json:"end_time,omitempty"`
	Limit     int    `json:"limit,omitempty"`
}

type ResultUserMeetingIDGet struct {
	MeetingIDList []string `json:"meetingid_list"`
}

func GetUserMeetingID(params *ParamsUserMeetingIDGet, result *ResultUserMeetingIDGet) wx.Action {
	return wx.NewPostAction(urls.CorpToolsMeetingGetUserMeetingID,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsMeetingInfoGet struct {
	MeetingID string `json:"meetingid"`
}

type ResultMeetingInfoGet struct {
	CreatorUserID          string            `json:"creator_userid"`
	Title                  string            `json:"title"`
	ReserveMeetingStart    int64             `json:"reserve_meeting_start"`
	ReserveMeetingDuration int               `json:"reserve_meeting_duration"`
	MeetingStart           int64             `json:"meeting_start"`
	MeetingDuration        int               `json:"meeting_duration"`
	Description            string            `json:"description"`
	MainDepartment         int64             `json:"main_department"`
	Type                   int               `json:"type"`
	Status                 int               `json:"status"`
	RemindTime             int               `json:"remind_time"`
	Attendees              *MeetingAttendees `json:"attendees"`
}

type MeetingAttendees struct {
	Member       []*MeetingMember       `json:"member"`
	ExternalUser []*MeetingExternalUser `json:"external_user"`
	Device       []*MeetingDevice       `json:"device"`
}

type MeetingMember struct {
	UserID string `json:"userid"`
	Status int    `json:"status"`
}

type MeetingExternalUser struct {
	ExternalUserID string `json:"external_userid"`
	Status         int    `json:"status"`
}

type MeetingDevice struct {
	DeviceSN string `json:"device_sn"`
	Status   int    `json:"status"`
}

func GetMeetingInfo(params *ParamsMeetingInfoGet, result *ResultMeetingInfoGet) wx.Action {
	return wx.NewPostAction(urls.CorpToolsMeetingGetInfo,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
