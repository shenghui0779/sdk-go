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

func TestAddCalendar(t *testing.T) {
	body := []byte(`{"calendar":{"organizer":"userid1","readonly":1,"set_as_default":1,"summary":"test_summary","color":"#FF3030","description":"test_describe","shares":[{"userid":"userid2"},{"userid":"userid3","readonly":1}]},"agentid":1000014}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "cal_id": "wcjgewCwAAqeJcPI1d8Pwbjt7nttzAAA"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/oa/calendar/add?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	readonly := new(int)
	*readonly = 1

	params := &ParamsCalendarAdd{
		Calendar: &CalendarAddData{
			Organizer:    "userid1",
			ReadOnly:     readonly,
			SetAsDefault: 1,
			Summary:      "test_summary",
			Color:        "#FF3030",
			Description:  "test_describe",
			Shares: []*CalendarShare{
				{
					UserID: "userid2",
				},
				{
					UserID:   "userid3",
					ReadOnly: readonly,
				},
			},
		},
		AgentID: 1000014,
	}
	result := new(ResultCalendarAdd)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", AddCalendar(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultCalendarAdd{
		CalID: "wcjgewCwAAqeJcPI1d8Pwbjt7nttzAAA",
	}, result)
}

func TestUpdateCalendar(t *testing.T) {
	body := []byte(`{"calendar":{"cal_id":"wcjgewCwAAqeJcPI1d8Pwbjt7nttzAAA","readonly":1,"summary":"test_summary","color":"#FF3030","description":"test_describe_1","shares":[{"userid":"userid1"},{"userid":"userid2","readonly":1}]}}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/oa/calendar/update?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	readonly := new(int)
	*readonly = 1

	params := &ParamsCalendarUpdate{
		Calendar: &CalendarUpdateData{
			CalID:       "wcjgewCwAAqeJcPI1d8Pwbjt7nttzAAA",
			ReadOnly:    readonly,
			Summary:     "test_summary",
			Color:       "#FF3030",
			Description: "test_describe_1",
			Shares: []*CalendarShare{
				{
					UserID: "userid1",
				},
				{
					UserID:   "userid2",
					ReadOnly: readonly,
				},
			},
		},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UpdateCalendar(params))

	assert.Nil(t, err)
}

func TestGetCalendar(t *testing.T) {
	body := []byte(`{"cal_id_list":["wcjgewCwAAqeJcPI1d8Pwbjt7nttzAAA"]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "calendar_list": [
        {
            "cal_id": "wcjgewCwAAqeJcPI1d8Pwbjt7nttzAAA",
            "organizer": "userid1",
            "readonly": 1,
            "summary": "test_summary",
            "color": "#FF3030",
            "description": "test_describe_1",
            "shares": [
                {
                    "userid": "userid2"
                },
                {
                    "userid": "userid1",
                    "readonly": 1
                }
            ]
        }
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/oa/calendar/get?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultCalendarGet)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetCalendar([]string{"wcjgewCwAAqeJcPI1d8Pwbjt7nttzAAA"}, result))

	readonly := new(int)
	*readonly = 1

	assert.Nil(t, err)
	assert.Equal(t, &ResultCalendarGet{
		CalendarList: []*Calendar{
			{
				CalID:       "wcjgewCwAAqeJcPI1d8Pwbjt7nttzAAA",
				Organizer:   "userid1",
				ReadOnly:    readonly,
				Summary:     "test_summary",
				Color:       "#FF3030",
				Description: "test_describe_1",
				Shares: []*CalendarShare{
					{
						UserID: "userid2",
					},
					{
						UserID:   "userid1",
						ReadOnly: readonly,
					},
				},
			},
		},
	}, result)
}

func TestDeleteCalendar(t *testing.T) {
	body := []byte(`{"cal_id":"wcjgewCwAAqeJcPI1d8Pwbjt7nttzAAA"}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/oa/calendar/del?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", DeleteCalendar("wcjgewCwAAqeJcPI1d8Pwbjt7nttzAAA"))

	assert.Nil(t, err)
}
