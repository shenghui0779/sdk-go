package kf

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/mock"
	"github.com/shenghui0779/gochat/offia"
	"github.com/stretchr/testify/assert"
)

func TestCreateSession(t *testing.T) {
	body := []byte(`{"kf_account":"test1@test","openid":"OPENID"}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/customservice/kfsession/create?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := offia.New("APPID", "APPSECRET")
	oa.SetClient(client)

	params := &ParamsSessionCreate{
		Account: "test1@test",
		OpenID:  "OPENID",
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", CreateSession(params))

	assert.Nil(t, err)
}

func TestCloseSession(t *testing.T) {
	body := []byte(`{"kf_account":"test1@test","openid":"OPENID"}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://api.weixin.qq.com/customservice/kfsession/close?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := offia.New("APPID", "APPSECRET")
	oa.SetClient(client)

	params := &ParamsSessionClose{
		Account: "test1@test",
		OpenID:  "OPENID",
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", CloseSession(params))

	assert.Nil(t, err)
}

func TestGetSession(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"createtime": 123456789,
	"kf_account": "test1@test"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://api.weixin.qq.com/customservice/kfsession/getsession?access_token=ACCESS_TOKEN&openid=OPENID", nil).Return(resp, nil)

	oa := offia.New("APPID", "APPSECRET")
	oa.SetClient(client)

	result := new(Session)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetSession("OPENID", result))

	assert.Nil(t, err)
	assert.Equal(t, &Session{
		Account:    "test1@test",
		CreateTime: 123456789,
	}, result)
}

func TestGetSessionList(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"sessionlist": [
		{
			"createtime": 123456789,
			"openid": "OPENID1"
		},
		{
			"createtime": 123456790,
			"openid": "OPENID2"
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://api.weixin.qq.com/customservice/kfsession/getsessionlist?access_token=ACCESS_TOKEN&kf_account=ACCOUNT", nil).Return(resp, nil)

	oa := offia.New("APPID", "APPSECRET")
	oa.SetClient(client)

	result := new(ResultSessionList)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetSessionList("ACCOUNT", result))

	assert.Nil(t, err)
	assert.Equal(t, []*Session{
		{
			OpenID:     "OPENID1",
			CreateTime: 123456789,
		},
		{
			OpenID:     "OPENID2",
			CreateTime: 123456790,
		},
	}, result)
}

func TestGetWaitCase(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"count": 150,
	"waitcaselist": [
		{
				"latest_time": 123456789,
				"openid": "OPENID1"
		},
		{
				"latest_time": 123456790,
				"openid": "OPENID2"
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://api.weixin.qq.com/customservice/kfsession/getwaitcase?access_token=ACCESS_TOKEN", nil).Return(resp, nil)

	oa := offia.New("APPID", "APPSECRET")
	oa.SetClient(client)

	result := new(WaitCase)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetWaitCase(result))

	assert.Nil(t, err)
	assert.Equal(t, &WaitCase{
		Count: 150,
		List: []*Session{
			{
				OpenID:     "OPENID1",
				LatestTime: 123456789,
			},
			{
				OpenID:     "OPENID2",
				LatestTime: 123456790,
			},
		},
	}, result)
}

func TestGetMsgRecordList(t *testing.T) {
	body := []byte(`{"endtime":987654321,"msgid":1,"number":10000,"starttime":987654321}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"recordlist": [
		{
			"openid": "oDF3iY9WMaswOPWjCIp_f3Bnpljk",
			"opercode": 2002,
			"text":"您好，客服test1为您服务。",
			"time":1400563710,
			"worker":  "test1@test"
		},
		{
			"openid":  "oDF3iY9WMaswOPWjCIp_f3Bnpljk",
			"opercode": 2003,
			"text": "你好，有什么事情？",
			"time": 1400563731,
			"worker": "test1@test"
		}
	],
	"number": 2,
	"msgid": 20165267
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/customservice/msgrecord/getmsglist?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := offia.New("APPID", "APPSECRET")
	oa.SetClient(client)

	params := &ParamsMsgRecordList{
		MsgID:     1,
		StartTime: 987654321,
		EndTime:   987654321,
		Number:    10000,
	}
	result := new(ResultMsgRecordList)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetMsgRecordList(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMsgRecordList{
		MsgID:  20165267,
		Number: 2,
		RecordList: []*MsgRecord{
			{
				Worker:   "test1@test",
				OpenID:   "oDF3iY9WMaswOPWjCIp_f3Bnpljk",
				OperCode: 2002,
				Text:     "您好，客服test1为您服务。",
				Time:     1400563710,
			},
			{
				Worker:   "test1@test",
				OpenID:   "oDF3iY9WMaswOPWjCIp_f3Bnpljk",
				OperCode: 2003,
				Text:     "你好，有什么事情？",
				Time:     1400563731,
			},
		},
	}, result)
}
