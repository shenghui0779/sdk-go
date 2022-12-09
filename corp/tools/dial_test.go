package tools

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/corp"
	"github.com/shenghui0779/gochat/mock"
)

func TestGetDialRecord(t *testing.T) {
	body := []byte(`{"start_time":1536508800,"end_time":1536940800,"offset":1,"limit":100}`)
	resp := []byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "record": [
        {
            "call_time": 1536508800,
            "total_duration": 10,
            "call_type": 1,
            "caller": {
                "userid": "tony",
                "duration": 10
            },
            "callee": [
                {
                    "phone": "138000800",
                    "duration": 10
                }
            ]
        },
        {
            "call_time": 1536940800,
            "total_duration": 20,
            "call_type": 2,
            "caller": {
                "userid": "tony",
                "duration": 10
            },
            "callee": [
                {
                    "phone": "138000800",
                    "duration": 5
                },
                {
                    "userid": "tom",
                    "duration": 5
                }
            ]
        }
    ]
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/dial/get_dial_record?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultDialRecord)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetDialRecord(1536508800, 1536940800, 1, 100, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultDialRecord{
		Record: []*DialRecord{
			{
				CallTime:      1536508800,
				TotalDuration: 10,
				CallType:      1,
				Caller: &DialCaller{
					UserID:   "tony",
					Duration: 10,
				},
				Callee: []*DialCallee{
					{
						Phone:    "138000800",
						Duration: 10,
					},
				},
			},
			{
				CallTime:      1536940800,
				TotalDuration: 20,
				CallType:      2,
				Caller: &DialCaller{
					UserID:   "tony",
					Duration: 10,
				},
				Callee: []*DialCallee{
					{
						Phone:    "138000800",
						Duration: 5,
					},
					{
						UserID:   "tom",
						Duration: 5,
					},
				},
			},
		},
	}, result)
}
