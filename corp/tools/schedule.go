package tools

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type Schedule struct {
	ScheduleID  string              `json:"schedule_id"`
	Organizer   string              `json:"organizer"`
	Attendees   []*ScheduleAttendee `json:"attendees"`
	Summary     string              `json:"summary"`
	Description string              `json:"description"`
	Reminders   *ScheduleReminders  `json:"reminders"`
	Location    string              `json:"location"`
	CalID       string              `json:"cal_id"`
	StartTime   int64               `json:"start_time"`
	EndTime     int64               `json:"end_time"`
	Status      int                 `json:"status"`
}

type ScheduleAttendee struct {
	UserID         string `json:"user_id"`
	ResponseStatus int    `json:"response_status,omitempty"`
}

type ScheduleReminders struct {
	IsRemind              int                    `json:"is_remind,omitempty"`
	IsRepeat              int                    `json:"is_repeat,omitempty"`
	RemindBeforeEventSecs int                    `json:"remind_before_event_secs,omitempty"`
	RemindTimeDiffs       []int                  `json:"remind_time_diffs,omitempty"`
	RepeatType            int                    `json:"repeat_type,omitempty"`
	RepeatUntil           int64                  `json:"repeat_until,omitempty"`
	IsCustomRepeat        int                    `json:"is_custom_repeat,omitempty"`
	RepeatInterval        int                    `json:"repeat_interval,omitempty"`
	RepeatDayOfWeek       []int                  `json:"repeat_day_of_week,omitempty"`
	RepeatDayOfMonth      []int                  `json:"repeat_day_of_month,omitempty"`
	Timezone              int                    `json:"timezone,omitempty"`
	ExcludeTimeList       []*ScheduleExcludeTime `json:"exclude_time_list,omitempty"`
}

type ScheduleExcludeTime struct {
	StartTime int64 `json:"start_time"`
}

type ParamsScheduleAdd struct {
	Schedule *ScheduleAddData `json:"schedule"`
	AgentID  int64            `json:"agentid,omitempty"`
}

type ScheduleAddData struct {
	Organizer   string              `json:"organizer"`
	Attendees   []*ScheduleAttendee `json:"attendees,omitempty"`
	Summary     string              `json:"summary,omitempty"`
	Description string              `json:"description,omitempty"`
	Reminders   *ScheduleReminders  `json:"reminders,omitempty"`
	Location    string              `json:"location,omitempty"`
	CalID       string              `json:"cal_id,omitempty"`
	StartTime   int64               `json:"start_time,omitempty"`
	EndTime     int64               `json:"end_time,omitempty"`
}

type ResultScheduleAdd struct {
	ScheduleID string `json:"schedule_id"`
}

func AddSchedule(params *ParamsScheduleAdd, result *ResultScheduleAdd) wx.Action {
	return wx.NewPostAction(urls.CorpToolsScheduleAdd,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsScheduleUpdate struct {
	Schedule *ScheduleUpdateData `json:"schedule"`
}

type ScheduleUpdateData struct {
	Organizer   string              `json:"organizer"`
	ScheduleID  string              `json:"schedule_id"`
	Attendees   []*ScheduleAttendee `json:"attendees,omitempty"`
	Summary     string              `json:"summary,omitempty"`
	Description string              `json:"description,omitempty"`
	Reminders   *ScheduleReminders  `json:"reminders,omitempty"`
	Location    string              `json:"location,omitempty"`
	CalID       string              `json:"cal_id,omitempty"`
	StartTime   int64               `json:"start_time,omitempty"`
	EndTime     int64               `json:"end_time,omitempty"`
}

func UpdateSchedule(params *ParamsScheduleUpdate) wx.Action {
	return wx.NewPostAction(urls.CorpToolsScheduleUpdate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsScheduleGet struct {
	ScheduleIDList []string `json:"schedule_id_list"`
}

type ResultScheduleGet struct {
	ScheduleList []*Schedule `json:"schedule_list"`
}

func GetSchedule(params *ParamsScheduleGet, result *ResultScheduleGet) wx.Action {
	return wx.NewPostAction(urls.CorpToolsScheduleGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsScheduleDelete struct {
	ScheduleID string `json:"schedule_id"`
}

func DeleteSchedule(params *ParamsScheduleDelete) wx.Action {
	return wx.NewPostAction(urls.CorpToolsScheduleDelete,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsSchduleGetByCalendar struct {
	CalID  string `json:"cal_id"`
	Offset int    `json:"offset,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}

type ResultScheduleGetByCalendar struct {
	ScheduleList []*Schedule `json:"schedule_list"`
}

func GetSchduleByCalendar(params *ParamsSchduleGetByCalendar, result *ResultScheduleGetByCalendar) wx.Action {
	return wx.NewPostAction(urls.CorpToolsScheduleGetByCalendar,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
