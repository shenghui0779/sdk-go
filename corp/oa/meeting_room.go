package oa

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type MeetingRoom struct {
	MeetingRoomID int64       `json:"meetingroom_id"`
	Name          string      `json:"name"`
	Capacity      int         `json:"capacity"`
	City          string      `json:"city"`
	Building      string      `json:"building"`
	Floor         string      `json:"floor"`
	Equipment     []int64     `json:"equipment"`
	Coordinate    *Coordinate `json:"coordinate"`
	NeedApproval  int         `json:"need_approval"`
}

type Coordinate struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type ParamsMeetingRoomAdd struct {
	Name       string      `json:"name"`
	Capacity   int         `json:"capacity"`
	City       string      `json:"city,omitempty"`
	Building   string      `json:"building,omitempty"`
	Floor      string      `json:"floor,omitempty"`
	Equipment  []int64     `json:"equipment,omitempty"`
	Coordinate *Coordinate `json:"coordinate,omitempty"`
}

type ResultMeetingRoomAdd struct {
	MeetingRoomID int64 `json:"meetingroom_id"`
}

func AddMeetingRoom(params *ParamsMeetingRoomAdd, result *ResultMeetingRoomAdd) wx.Action {
	return wx.NewPostAction(urls.CorpOAAddMeetingRoom,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsMeetingRoomList struct {
	City      string  `json:"city,omitempty"`
	Building  string  `json:"building,omitempty"`
	Floor     string  `json:"floor,omitempty"`
	Equipment []int64 `json:"equipment,omitempty"`
}

type ResultMeetingRoomList struct {
	MeetingRoomList []*MeetingRoom `json:"meetingroom_list"`
}

func MeetingRoomList(params *ParamsMeetingRoomList, result *ResultMeetingRoomList) wx.Action {
	return wx.NewPostAction(urls.CorpOAMeetingRoomList,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsMeetingRoomEdit struct {
	MeetingRoomID int64       `json:"meetingroom_id"`
	Name          string      `json:"name,omitempty"`
	Capacity      int         `json:"capacity,omitempty"`
	City          string      `json:"city,omitempty"`
	Building      string      `json:"building,omitempty"`
	Floor         string      `json:"floor,omitempty"`
	Equipment     []int64     `json:"equipment,omitempty"`
	Coordinate    *Coordinate `json:"coordinate,omitempty"`
}

func EditMeetingRoom(params *ParamsMeetingRoomEdit) wx.Action {
	return wx.NewPostAction(urls.CorpOAEditMeetingRoom,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsMeetingRoomDel struct {
	MeetingRoomID int64 `json:"meetingroom_id"`
}

func DelMeetingRoom(params *ParamsMeetingRoomDel) wx.Action {
	return wx.NewPostAction(urls.CorpOADelMeetingRoom,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type MeetingRoomBookingInfo struct {
	MeetingRoomID int64                         `json:"meetingroom_id"`
	Schedule      []*MeetingRoomBookingSchedule `json:"schedule"`
}

type MeetingRoomBookingSchedule struct {
	MeetingID  int64  `json:"meeting_id"`
	ScheduleID int64  `json:"schedule_id"`
	StartTime  int64  `json:"start_time"`
	EndTime    int64  `json:"end_time"`
	Booker     string `json:"booker"`
}

type ParamsMeetingRoomBookingInfoGet struct {
	MeetingRoomID int64  `json:"meetingroom_id"`
	StartTime     int64  `json:"start_time,omitempty"`
	EndTime       int64  `json:"end_time,omitempty"`
	City          string `json:"city,omitempty"`
	Building      string `json:"building,omitempty"`
	Floor         string `json:"floor,omitempty"`
}

type ResultMeetingRoomBookingInfoGet struct {
	BookingList []*MeetingRoomBookingInfo `json:"booking_list"`
}

func GetMeetingRoomBookingInfo(params *ParamsMeetingRoomBookingInfoGet, result *ResultMeetingRoomBookingInfoGet) wx.Action {
	return wx.NewPostAction(urls.CorpOAGetMeetingRoomBookingInfo,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsMeetingRoomBook struct {
	MeetingRoomID int64    `json:"meetingroom_id"`
	Subject       string   `json:"subject,omitempty"`
	StartTime     int64    `json:"start_time"`
	EndTime       int64    `json:"end_time"`
	Booker        string   `json:"booker"`
	Attendees     []string `json:"attendees,omitempty"`
}

type ResultMeetingRoomBook struct {
	MeetingID  int64 `json:"meeting_id"`
	ScheduleID int64 `json:"schedule_id"`
}

func BookMeetingRoom(params *ParamsMeetingRoomBook, result *ResultMeetingRoomBook) wx.Action {
	return wx.NewPostAction(urls.CorpOABookMeetingRoom,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
