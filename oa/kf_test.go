package oa

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/wx"
	"github.com/stretchr/testify/assert"
)

func TestGetKFAccountList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/customservice/getkflist?access_token=ACCESS_TOKEN").Return([]byte(`{
		"kf_list" : [
		   {
			  "kf_account": "test1@test",
			  "kf_headimgurl": "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl2iccsvYbHvnphkyGtnvjfUS8Ym0GSaLic0FD3vN0V8PILcibEGb2fPfEOmw/0",
			  "kf_id": "1001",
			  "kf_nick": "ntest1",
			  "kf_wx": "kfwx1"
		   },
		   {
			  "kf_account": "test2@test",
			  "kf_headimgurl": "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl2iccsvYbHvnphkyGtnvjfUS8Ym0GSaLic0FD3vN0V8PILcibEGb2fPfEOmw/0",
			  "kf_id": "1002",
			  "kf_nick": "ntest2",
			  "invite_wx": "kfwx2",
			  "invite_expire_time": 123456789,
			  "invite_status": "waiting"
		   }
		]
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := make([]*KFAccount, 0)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetKFAccountList(&dest))

	assert.Nil(t, err)
	assert.Equal(t, []*KFAccount{
		{
			ID:         "1001",
			Account:    "test1@test",
			Nickname:   "ntest1",
			HeadImgURL: "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl2iccsvYbHvnphkyGtnvjfUS8Ym0GSaLic0FD3vN0V8PILcibEGb2fPfEOmw/0",
			Weixin:     "kfwx1",
		},
		{
			ID:               "1002",
			Account:          "test2@test",
			Nickname:         "ntest2",
			HeadImgURL:       "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl2iccsvYbHvnphkyGtnvjfUS8Ym0GSaLic0FD3vN0V8PILcibEGb2fPfEOmw/0",
			InviteWeixin:     "kfwx2",
			InviteExpireTime: 123456789,
			InviteStatus:     InviteWaiting,
		},
	}, dest)
}

func TestGetKFOnlineList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/customservice/getonlinekflist?access_token=ACCESS_TOKEN").Return([]byte(`{
		"kf_online_list": [
			{
				"kf_account": "test1@test",
				"status": 1,
				"kf_id": "1001",
				"accepted_case": 1
			},
			{
				"kf_account": "test2@test",
				"status": 1,
				"kf_id": "1002",
				"accepted_case": 2
			}
		]
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := make([]*KFOnline, 0)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetKFOnlineList(&dest))

	assert.Nil(t, err)
	assert.Equal(t, []*KFOnline{
		{
			ID:           "1001",
			Account:      "test1@test",
			Status:       1,
			AcceptedCase: 1,
		},
		{
			ID:           "1002",
			Account:      "test2@test",
			Status:       1,
			AcceptedCase: 2,
		},
	}, dest)
}

func TestAddKFAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/customservice/kfaccount/add?access_token=ACCESS_TOKEN", []byte(`{"kf_account":"test1@test","nickname":"客服1"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", AddKFAccount("test1@test", "客服1"))

	assert.Nil(t, err)
}

func TestUpdateKFAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/customservice/kfaccount/update?access_token=ACCESS_TOKEN", []byte(`{"kf_account":"test1@test","nickname":"客服1"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UpdateKFAccount("test1@test", "客服1"))

	assert.Nil(t, err)
}

func TestInviteKFWorker(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/customservice/kfaccount/inviteworker?access_token=ACCESS_TOKEN", []byte(`{"invite_wx":"test_kfwx","kf_account":"test1@test"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", InviteKFWorker("test1@test", "test_kfwx"))

	assert.Nil(t, err)
}

func TestUploadKFAvatar(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/customservice/kfaccount/uploadheadimg?access_token=ACCESS_TOKEN&kf_account=KFACCOUNT", wx.NewUploadForm("media", "test.jpg", nil)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UploadKFAvatar("KFACCOUNT", "test.jpg"))

	assert.Nil(t, err)
}

func TestDeleteKFAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/customservice/kfaccount/del?access_token=ACCESS_TOKEN&kf_account=KFACCOUNT").Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", DeleteKFAccount("KFACCOUNT"))

	assert.Nil(t, err)
}

func TestCreateKFSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/customservice/kfsession/create?access_token=ACCESS_TOKEN", []byte(`{"kf_account":"test1@test","openid":"OPENID"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", CreateKFSession("test1@test", "OPENID"))

	assert.Nil(t, err)
}

func TestCloseKFSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/customservice/kfsession/close?access_token=ACCESS_TOKEN", []byte(`{"kf_account":"test1@test","openid":"OPENID"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", CloseKFSession("test1@test", "OPENID"))

	assert.Nil(t, err)
}

func TestGetKFSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/customservice/kfsession/getsession?access_token=ACCESS_TOKEN&openid=OPENID").Return([]byte(`{
		"createtime": 123456789,
		"kf_account": "test1@test"
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(KFSession)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetKFSession(dest, "OPENID"))

	assert.Nil(t, err)
	assert.Equal(t, &KFSession{
		KFAccount:  "test1@test",
		CreateTime: 123456789,
	}, dest)
}

func TestGetKFSessionList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/customservice/kfsession/getsessionlist?access_token=ACCESS_TOKEN&kf_account=KFACCOUNT").Return([]byte(`{
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

	dest := make([]*KFSession, 0)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetKFSessionList(&dest, "KFACCOUNT"))

	assert.Nil(t, err)
	assert.Equal(t, []*KFSession{
		{
			OpenID:     "OPENID1",
			CreateTime: 123456789,
		},
		{
			OpenID:     "OPENID2",
			CreateTime: 123456790,
		},
	}, dest)
}

func TestGetKFWaitCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

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

	dest := new(KFWaitCase)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetKFWaitCase(dest))

	assert.Nil(t, err)
	assert.Equal(t, &KFWaitCase{
		Count: 150,
		List: []*KFSession{
			{
				OpenID:     "OPENID1",
				LatestTime: 123456789,
			},
			{
				OpenID:     "OPENID2",
				LatestTime: 123456790,
			},
		},
	}, dest)
}

func TestGetKFMsgRecordList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

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

	dest := new(KFMsgRecordList)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetKFMsgRecordList(dest, 1, 987654321, 987654321, 10000))

	assert.Nil(t, err)
	assert.Equal(t, &KFMsgRecordList{
		MsgID:  20165267,
		Number: 2,
		RecordList: []*KFMsgRecord{
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
	}, dest)
}
