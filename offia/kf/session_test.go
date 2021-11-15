package kf

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/wx"
	"github.com/stretchr/testify/assert"
)

func TestCreateSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/customservice/kfsession/create?access_token=ACCESS_TOKEN", []byte(`{"kf_account":"test1@test","openid":"OPENID"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", CreateSession("test1@test", "OPENID"))

	assert.Nil(t, err)
}

func TestCloseSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/customservice/kfsession/close?access_token=ACCESS_TOKEN", []byte(`{"kf_account":"test1@test","openid":"OPENID"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", CloseSession("test1@test", "OPENID"))

	assert.Nil(t, err)
}

func TestGetSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/customservice/kfsession/getsession?access_token=ACCESS_TOKEN&openid=OPENID").Return([]byte(`{
		"createtime": 123456789,
		"kf_account": "test1@test"
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	result := new(Session)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetSession(result, "OPENID"))

	assert.Nil(t, err)
	assert.Equal(t, &Session{
		Account:    "test1@test",
		CreateTime: 123456789,
	}, result)
}

func TestGetSessionList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/customservice/kfsession/getsessionlist?access_token=ACCESS_TOKEN&kf_account=ACCOUNT").Return([]byte(`{
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
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	result := make([]*Session, 0)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetSessionList(&result, "ACCOUNT"))

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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/customservice/kfsession/getwaitcase?access_token=ACCESS_TOKEN").Return([]byte(`{
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
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/customservice/msgrecord/getmsglist?access_token=ACCESS_TOKEN", []byte(`{"endtime":987654321,"msgid":1,"number":10000,"starttime":987654321}`)).Return([]byte(`{
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
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	result := new(MsgRecordList)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetMsgRecordList(result, 1, 987654321, 987654321, 10000))

	assert.Nil(t, err)
	assert.Equal(t, &MsgRecordList{
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
