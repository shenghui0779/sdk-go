package tools

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/chenghonour/gochat/corp"
	"github.com/chenghonour/gochat/mock"
	"github.com/chenghonour/gochat/wx"
)

func TestAddSchedule(t *testing.T) {
	body := []byte(`{"schedule":{"organizer":"userid1","start_time":1571274600,"end_time":1571320210,"attendees":[{"userid":"userid2"}],"summary":"需求评审会议","description":"2.0版本需求初步评审","reminders":{"is_remind":1,"remind_before_event_secs":3600,"is_repeat":1,"repeat_type":7,"repeat_until":1606976813,"is_custom_repeat":1,"repeat_interval":1,"repeat_day_of_week":[3,7],"repeat_day_of_month":[10,21],"timezone":8},"location":"广州国际媒体港10楼1005会议室","cal_id":"wcjgewCwAAqeJcPI1d8Pwbjt7nttzAAA"},"agentid":1000014}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "schedule_id": "17c7d2bd9f20d652840f72f59e796AAA"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/oa/schedule/add?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsScheduleAdd{
		Schedule: &ScheduleAddData{
			Organizer: "userid1",
			StartTime: 1571274600,
			EndTime:   1571320210,
			Attendees: []*ParamsScheduleAttendee{
				{
					UserID: "userid2",
				},
			},
			Summary:     "需求评审会议",
			Description: "2.0版本需求初步评审",
			Reminders: &ParamsScheduleReminders{
				IsRemind:              1,
				RemindBeforeEventSecs: 3600,
				IsRepeat:              1,
				RepeatType:            7,
				RepeatUntil:           1606976813,
				IsCustomRepeat:        1,
				RepeatInterval:        1,
				RepeatDayOfWeek:       []int{3, 7},
				RepeatDayOfMonth:      []int{10, 21},
				Timezone:              8,
			},
			Location: "广州国际媒体港10楼1005会议室",
			CalID:    "wcjgewCwAAqeJcPI1d8Pwbjt7nttzAAA",
		},
		AgentID: 1000014,
	}
	result := new(ResultScheduleAdd)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", AddSchedule(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultScheduleAdd{
		ScheduleID: "17c7d2bd9f20d652840f72f59e796AAA",
	}, result)
}

func TestUpdateSchedule(t *testing.T) {
	body := []byte(`{"schedule":{"organizer":"userid1","schedule_id":"17c7d2bd9f20d652840f72f59e796AAA","start_time":1571274600,"end_time":1571320210,"attendees":[{"userid":"userid2"}],"summary":"test_summary","description":"test_description","reminders":{"is_remind":1,"remind_before_event_secs":3600,"is_repeat":1,"repeat_type":7,"repeat_until":1606976813,"is_custom_repeat":1,"repeat_interval":1,"repeat_day_of_week":[3,7],"repeat_day_of_month":[10,21],"timezone":8},"location":"test_place","skip_attendees":false}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/oa/schedule/update?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsScheduleUpdate{
		Schedule: &ScheduleUpdateData{
			Organizer:  "userid1",
			ScheduleID: "17c7d2bd9f20d652840f72f59e796AAA",
			StartTime:  1571274600,
			EndTime:    1571320210,
			Attendees: []*ParamsScheduleAttendee{
				{
					UserID: "userid2",
				},
			},
			Summary:     "test_summary",
			Description: "test_description",
			Reminders: &ParamsScheduleReminders{
				IsRemind:              1,
				RemindBeforeEventSecs: 3600,
				IsRepeat:              1,
				RepeatType:            7,
				RepeatUntil:           1606976813,
				IsCustomRepeat:        1,
				RepeatInterval:        1,
				RepeatDayOfWeek:       []int{3, 7},
				RepeatDayOfMonth:      []int{10, 21},
				Timezone:              8,
			},
			Location:      "test_place",
			SkipAttendees: false,
		},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UpdateSchedule(params))

	assert.Nil(t, err)
}

func TestGetSchedule(t *testing.T) {
	body := []byte(`{"schedule_id_list":["17c7d2bd9f20d652840f72f59e796AAA"]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "schedule_list": [
        {
            "schedule_id": "17c7d2bd9f20d652840f72f59e796AAA",
            "organizer": "userid1",
            "attendees": [
                {
                    "userid": "userid2",
                    "response_status": 1
                }
            ],
            "summary": "test_summary",
            "description": "test_content",
            "reminders": {
                "is_remind": 1,
                "is_repeat": 1,
                "remind_before_event_secs": 3600,
                "remind_time_diffs": [
                    -3600
                ],
                "repeat_type": 7,
                "repeat_until": 1606976813,
                "is_custom_repeat": 1,
                "repeat_interval": 1,
                "repeat_day_of_week": [
                    3,
                    7
                ],
                "repeat_day_of_month": [
                    10,
                    21
                ],
                "timezone": 8,
                "exclude_time_list": [
                    {
                        "start_time": 1571361000
                    }
                ]
            },
            "location": "test_place",
            "cal_id": "wcjgewCwAAqeJcPI1d8Pwbjt7nttzAAA",
            "start_time": 1571274600,
            "end_time": 1571579410,
            "status": 1
        }
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/oa/schedule/get?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultScheduleGet)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetSchedule([]string{"17c7d2bd9f20d652840f72f59e796AAA"}, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultScheduleGet{
		ScheduleList: []*Schedule{
			{
				ScheduleID: "17c7d2bd9f20d652840f72f59e796AAA",
				Organizer:  "userid1",
				Attendees: []*ScheduleAttendee{
					{
						UserID:         "userid2",
						ResponseStatus: 1,
					},
				},
				Summary:     "test_summary",
				Description: "test_content",
				Reminders: &ScheduleReminders{
					IsRemind:              1,
					IsRepeat:              1,
					RemindBeforeEventSecs: 3600,
					RemindTimeDiffs:       []int{-3600},
					RepeatType:            7,
					RepeatUntil:           1606976813,
					IsCustomRepeat:        1,
					RepeatInterval:        1,
					RepeatDayOfWeek:       []int{3, 7},
					RepeatDayOfMonth:      []int{10, 21},
					Timezone:              8,
					ExcludeTimeList: []*ScheduleExcludeTime{
						{
							StartTime: 1571361000,
						},
					},
				},
				Location:  "test_place",
				CalID:     "wcjgewCwAAqeJcPI1d8Pwbjt7nttzAAA",
				StartTime: 1571274600,
				EndTime:   1571579410,
				Status:    1,
			},
		},
	}, result)
}

func TestDeleteSchedule(t *testing.T) {
	body := []byte(`{"schedule_id":"17c7d2bd9f20d652840f72f59e796AAA"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/oa/schedule/del?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", DeleteSchedule("17c7d2bd9f20d652840f72f59e796AAA"))

	assert.Nil(t, err)
}

func TestGetSchduleByCalendar(t *testing.T) {
	body := []byte(`{"cal_id":"wcjgewCwAAqeJcPI1d8Pwbjt7nttzAAA","offset":100,"limit":1000}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "schedule_list": [
        {
            "schedule_id": "17c7d2bd9f20d652840f72f59e796AAA",
            "sequence": 100,
            "attendees": [
                {
                    "userid": "userid1",
                    "response_status": 0
                }
            ],
            "summary": "test_summary",
            "description": "test_content",
            "reminders": {
                "is_remind": 1,
                "is_repeat": 1,
                "remind_before_event_secs": 3600,
                "repeat_type": 7,
                "repeat_until": 1606976813,
                "is_custom_repeat": 1,
                "repeat_interval": 1,
                "repeat_day_of_week": [
                    3,
                    7
                ],
                "repeat_day_of_month": [
                    10,
                    21
                ],
                "timezone": 8
            },
            "location": "test_place",
            "start_time": 1571274600,
            "end_time": 1571320210,
            "status": 1,
            "cal_id": "wcjgewCwAAqeJcPI1d8Pwbjt7nttzAAA"
        }
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/oa/schedule/get_by_calendar?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultScheduleGetByCalendar)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetScheduleByCalendar("wcjgewCwAAqeJcPI1d8Pwbjt7nttzAAA", 100, 1000, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultScheduleGetByCalendar{
		ScheduleList: []*Schedule{
			{
				ScheduleID: "17c7d2bd9f20d652840f72f59e796AAA",
				Attendees: []*ScheduleAttendee{
					{
						UserID:         "userid1",
						ResponseStatus: 0,
					},
				},
				Summary:     "test_summary",
				Description: "test_content",
				Reminders: &ScheduleReminders{
					IsRemind:              1,
					IsRepeat:              1,
					RemindBeforeEventSecs: 3600,
					RepeatType:            7,
					RepeatUntil:           1606976813,
					IsCustomRepeat:        1,
					RepeatInterval:        1,
					RepeatDayOfWeek:       []int{3, 7},
					RepeatDayOfMonth:      []int{10, 21},
					Timezone:              8,
				},
				Location:  "test_place",
				CalID:     "wcjgewCwAAqeJcPI1d8Pwbjt7nttzAAA",
				StartTime: 1571274600,
				EndTime:   1571320210,
				Status:    1,
			},
		},
	}, result)
}

func TestAddScheduleAttendee(t *testing.T) {
	body := []byte(`{"schedule_id":"17c7d2bd9f20d652840f72f59e796AAA","attendees":[{"userid":"userid2"}]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/oa/schedule/add_attendees?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", AddScheduleAttendee("17c7d2bd9f20d652840f72f59e796AAA", "userid2"))

	assert.Nil(t, err)
}

func TestDeleteScheduleAttendee(t *testing.T) {
	body := []byte(`{"schedule_id":"17c7d2bd9f20d652840f72f59e796AAA","attendees":[{"userid":"userid2"}]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/oa/schedule/del_attendees?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", DeleteScheduleAttendee("17c7d2bd9f20d652840f72f59e796AAA", "userid2"))

	assert.Nil(t, err)
}
