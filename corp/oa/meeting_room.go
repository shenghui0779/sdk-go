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

// AddMeetingRoom 添加会议室
func AddMeetingRoom(params *ParamsMeetingRoomAdd, result *ResultMeetingRoomAdd) wx.Action {
	return wx.NewPostAction(urls.CorpOAMeetingRoomAdd,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
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

// ListMeetingRoom 查询会议室
func ListMeetingRoom(params *ParamsMeetingRoomList, result *ResultMeetingRoomList) wx.Action {
	return wx.NewPostAction(urls.CorpOAMeetingRoomList,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
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

// EditMeetingRoom 编辑会议室
func EditMeetingRoom(params *ParamsMeetingRoomEdit) wx.Action {
	return wx.NewPostAction(urls.CorpOAMeetingRoomEdit,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type ParamsMeetingRoomDelete struct {
	MeetingRoomID int64 `json:"meetingroom_id"`
}

// DeleteMeetingRoom 删除会议室
func DeleteMeetingRoom(meetingRoomID int64) wx.Action {
	params := &ParamsMeetingRoomDelete{
		MeetingRoomID: meetingRoomID,
	}

	return wx.NewPostAction(urls.CorpOAMeetingRoomDelete,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type ParamsMeetingRoomBookingInfo struct {
	MeetingRoomID int64  `json:"meetingroom_id"`
	StartTime     int64  `json:"start_time,omitempty"`
	EndTime       int64  `json:"end_time,omitempty"`
	City          string `json:"city,omitempty"`
	Building      string `json:"building,omitempty"`
	Floor         string `json:"floor,omitempty"`
}

type ResultMeetingRoomBookingInfo struct {
	BookingList []*MeetingRoomBookingInfo `json:"booking_list"`
}

type MeetingRoomBookingInfo struct {
	MeetingRoomID int64                         `json:"meetingroom_id"`
	Schedule      []*MeetingRoomBookingSchedule `json:"schedule"`
}

type MeetingRoomBookingSchedule struct {
	MeetingID  string `json:"meeting_id"`
	ScheduleID string `json:"schedule_id"`
	StartTime  int64  `json:"start_time"`
	EndTime    int64  `json:"end_time"`
	Booker     string `json:"booker"`
}

// GetMeetingRoomBookingInfo 查询会议室的预定信息
func GetMeetingRoomBookingInfo(params *ParamsMeetingRoomBookingInfo, result *ResultMeetingRoomBookingInfo) wx.Action {
	return wx.NewPostAction(urls.CorpOAGetMeetingRoomBookingInfo,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
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
	MeetingID  string `json:"meeting_id"`
	ScheduleID string `json:"schedule_id"`
}

// BookMeetingRoom 预定会议室
func BookMeetingRoom(params *ParamsMeetingRoomBook, result *ResultMeetingRoomBook) wx.Action {
	return wx.NewPostAction(urls.CorpOAMeetingRoomBook,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ParamsMeetingRoomBookCancel struct {
	MeetingID    string `json:"meeting_id"`
	KeepSchedule int    `json:"keep_schedule"`
}

// 取消预定会议室
func CancelBookMeetingRoom(meetingID string, keepSchedule int) wx.Action {
	params := &ParamsMeetingRoomBookCancel{
		MeetingID:    meetingID,
		KeepSchedule: keepSchedule,
	}

	return wx.NewPostAction(urls.CorpOAMeetingRoomCancelBook,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type ParamsMeetingRoomBookingInfoByMeetingID struct {
	MeetingRoomID int64  `json:"meetingroom_id"`
	MeetingID     string `json:"meeting_id"`
}

// GetMeetingRoomBookingInfoByID 根据会议ID查询会议室的预定信息
func GetMeetingRoomBookingInfoByID(meetingRoomID int64, meetingID string, result *MeetingRoomBookingInfo) wx.Action {
	params := &ParamsMeetingRoomBookingInfoByMeetingID{
		MeetingRoomID: meetingRoomID,
		MeetingID:     meetingID,
	}

	return wx.NewPostAction(urls.CorpOAGetMeetingRoomBookingInfoByMeetingID,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}
