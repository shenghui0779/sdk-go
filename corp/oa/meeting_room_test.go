package oa

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

func TestAddMeetingRoom(t *testing.T) {
	body := []byte(`{"name":"18F-会议室","capacity":10,"city":"深圳","building":"腾讯大厦","floor":"18F","equipment":[1,2,3],"coordinate":{"latitude":"22.540503","longitude":"113.934528"}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"meetingroom_id": 1
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/add?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsMeetingRoomAdd{
		Name:      "18F-会议室",
		Capacity:  10,
		City:      "深圳",
		Building:  "腾讯大厦",
		Floor:     "18F",
		Equipment: []int64{1, 2, 3},
		Coordinate: &Coordinate{
			Latitude:  "22.540503",
			Longitude: "113.934528",
		},
	}
	result := new(ResultMeetingRoomAdd)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", AddMeetingRoom(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMeetingRoomAdd{
		MeetingRoomID: 1,
	}, result)
}

func TestListMeetingRoom(t *testing.T) {
	body := []byte(`{"city":"深圳","building":"腾讯大厦","floor":"18F","equipment":[1,2,3]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"meetingroom_list": [
		{
			"meetingroom_id": 1,
			"name": "18F-会议室",
			"capacity": 10,
			"city": "深圳",
			"building": "腾讯大厦",
			"floor": "18F",
			"equipment": [
				1,
				2,
				3
			],
			"coordinate": {
				"latitude": "22.540503",
				"longitude": "113.934528"
			},
			"need_approval": 1
		},
		{
			"meetingroom_id": 2,
			"name": "19F-会议室",
			"capacity": 20,
			"city": "深圳",
			"building": "腾讯大厦",
			"floor": "19F",
			"equipment": [
				2,
				3
			],
			"coordinate": {
				"latitude": "22.540503",
				"longitude": "113.934528"
			}
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/list?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsMeetingRoomList{
		City:      "深圳",
		Building:  "腾讯大厦",
		Floor:     "18F",
		Equipment: []int64{1, 2, 3},
	}
	result := new(ResultMeetingRoomList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListMeetingRoom(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMeetingRoomList{
		MeetingRoomList: []*MeetingRoom{
			{
				MeetingRoomID: 1,
				Name:          "18F-会议室",
				Capacity:      10,
				City:          "深圳",
				Building:      "腾讯大厦",
				Floor:         "18F",
				Equipment:     []int64{1, 2, 3},
				Coordinate: &Coordinate{
					Latitude:  "22.540503",
					Longitude: "113.934528",
				},
				NeedApproval: 1,
			},
			{
				MeetingRoomID: 2,
				Name:          "19F-会议室",
				Capacity:      20,
				City:          "深圳",
				Building:      "腾讯大厦",
				Floor:         "19F",
				Equipment:     []int64{2, 3},
				Coordinate: &Coordinate{
					Latitude:  "22.540503",
					Longitude: "113.934528",
				},
			},
		},
	}, result)
}

func TestEditMeetingRoom(t *testing.T) {
	body := []byte(`{"meetingroom_id":2,"name":"18F-会议室","capacity":10,"city":"深圳","building":"腾讯大厦","floor":"18F","equipment":[1,2,3],"coordinate":{"latitude":"22.540503","longitude":"113.934528"}}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/edit?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsMeetingRoomEdit{
		MeetingRoomID: 2,
		Name:          "18F-会议室",
		Capacity:      10,
		City:          "深圳",
		Building:      "腾讯大厦",
		Floor:         "18F",
		Equipment:     []int64{1, 2, 3},
		Coordinate: &Coordinate{
			Latitude:  "22.540503",
			Longitude: "113.934528",
		},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", EditMeetingRoom(params))

	assert.Nil(t, err)
}

func TestDeleteMeetingRoom(t *testing.T) {
	body := []byte(`{"meetingroom_id":1}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/del?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", DeleteMeetingRoom(1))

	assert.Nil(t, err)
}

func TestGetMeetingRoomBookingInfo(t *testing.T) {
	body := []byte(`{"meetingroom_id":1,"start_time":1593532800,"end_time":1593619200,"city":"深圳","building":"腾讯大厦","floor":"18F"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"booking_list": [
		{
			"meetingroom_id": 1,
			"schedule": [
				{
					"meeting_id": "mtebsada6e027c123cbafAAA",
					"schedule_id": "17c7d2bd9f20d652840f72f59e796AAA",
					"start_time": 1593532800,
					"end_time": 1593662400,
					"booker": "zhangsan"
				}
			]
		},
		{
			"meetingroom_id": 2,
			"schedule": []
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/get_booking_info?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsMeetingRoomBookingInfo{
		MeetingRoomID: 1,
		StartTime:     1593532800,
		EndTime:       1593619200,
		City:          "深圳",
		Building:      "腾讯大厦",
		Floor:         "18F",
	}
	result := new(ResultMeetingRoomBookingInfo)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetMeetingRoomBookingInfo(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMeetingRoomBookingInfo{
		BookingList: []*MeetingRoomBookingInfo{
			{
				MeetingRoomID: 1,
				Schedule: []*MeetingRoomBookingSchedule{
					{
						MeetingID:  "mtebsada6e027c123cbafAAA",
						ScheduleID: "17c7d2bd9f20d652840f72f59e796AAA",
						StartTime:  1593532800,
						EndTime:    1593662400,
						Booker:     "zhangsan",
					},
				},
			},
			{
				MeetingRoomID: 2,
				Schedule:      []*MeetingRoomBookingSchedule{},
			},
		},
	}, result)
}

func TestBookMeetingRoom(t *testing.T) {
	body := []byte(`{"meetingroom_id":1,"subject":"周会","start_time":1593532800,"end_time":1593619200,"booker":"zhangsan","attendees":["lisi","wangwu"]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"meeting_id": "mtgsaseb6e027c123cbafAAA",
	"schedule_id": "17c7d2bd9f20d652840f72f59e796AAA"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/book?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsMeetingRoomBook{
		MeetingRoomID: 1,
		Subject:       "周会",
		StartTime:     1593532800,
		EndTime:       1593619200,
		Booker:        "zhangsan",
		Attendees:     []string{"lisi", "wangwu"},
	}
	result := new(ResultMeetingRoomBook)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", BookMeetingRoom(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMeetingRoomBook{
		MeetingID:  "mtgsaseb6e027c123cbafAAA",
		ScheduleID: "17c7d2bd9f20d652840f72f59e796AAA",
	}, result)
}

func TestCancelBookMeetingRoom(t *testing.T) {
	body := []byte(`{"meeting_id":"mt42b34949gsaseb6e027c123cbafAAA","keep_schedule":1}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/cancel_book?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", CancelBookMeetingRoom("mt42b34949gsaseb6e027c123cbafAAA", 1))

	assert.Nil(t, err)
}

func TestGetMeetingRoomBookingInfoByID(t *testing.T) {
	body := []byte(`{"meetingroom_id":1,"meeting_id":"mtebsada6e027c123cbafAAA"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"meetingroom_id": 1,
	"schedule": [
		{
			"meeting_id": "mtebsada6e027c123cbafAAA",
			"schedule_id": "17c7d2bd9f20d652840f72f59e796AAA",
			"start_time": 1593532800,
			"end_time": 1593662400,
			"booker": "zhangsan"
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/get_booking_info_by_meeting_id?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(MeetingRoomBookingInfo)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetMeetingRoomBookingInfoByID(1, "mtebsada6e027c123cbafAAA", result))

	assert.Nil(t, err)
	assert.Equal(t, &MeetingRoomBookingInfo{
		MeetingRoomID: 1,
		Schedule: []*MeetingRoomBookingSchedule{
			{
				MeetingID:  "mtebsada6e027c123cbafAAA",
				ScheduleID: "17c7d2bd9f20d652840f72f59e796AAA",
				StartTime:  1593532800,
				EndTime:    1593662400,
				Booker:     "zhangsan",
			},
		},
	}, result)
}
