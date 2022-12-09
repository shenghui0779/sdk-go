package minip

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/mock"
	"github.com/shenghui0779/gochat/wx"
)

func TestImageSecCheck(t *testing.T) {
	resp := []byte(`{"errcode":0,"errmsg":"ok"}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/img_sec_check?access_token=ACCESS_TOKEN", gomock.AssignableToTypeOf(wx.NewUploadForm())).Return(resp, nil)

	mp := New("APPID", "APPSECRET", WithMockClient(client))

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", ImageSecCheck("../mock/test.jpg"))

	assert.Nil(t, err)
}

func TestMediaCheckAsync(t *testing.T) {
	body := []byte(`{"media_type":2,"media_url":"https://developers.weixin.qq.com/miniprogram/assets/images/head_global_z_@all.png","version":2,"scene":1,"openid":"OPENID"}`)

	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"trace_id": "967e945cd8a3e458f3c74dcb886068e9"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/wxa/media_check_async?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	mp := New("APPID", "APPSECRET", WithMockClient(client))

	params := &ParamsMediaCheckAsync{
		MediaType: SecMediaImage,
		MediaURL:  "https://developers.weixin.qq.com/miniprogram/assets/images/head_global_z_@all.png",
		Version:   2,
		Scene:     SecSceneDoc,
		OpenID:    "OPENID",
	}

	result := new(ResultMediaCheckAsync)

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", MediaCheckAsync(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMediaCheckAsync{
		TraceID: "967e945cd8a3e458f3c74dcb886068e9",
	}, result)
}

func TestMsgSecCheck(t *testing.T) {
	body := []byte(`{"content":"hello world!","version":2,"scene":1,"openid":"OPENID"}`)

	resp := []byte(`{
	"errcode":0,
	"errmsg":"ok",
	"result":{
		"suggest":"risky",
		"label":20001
	},
	"detail":[
		{
			"strategy":"content_model",
			"errcode":0,
			"suggest":"risky",
			"label":20006,
			"prob":90
		},
		{
			"strategy":"keyword",
			"errcode":0,
			"suggest":"pass",
			"label":20006,
			"keyword":"命中的关键词1"
		},
		{
			"strategy":"keyword",
			"errcode":0,
			"suggest":"risky",
			"label":20006,
			"keyword":"命中的关键词2"
		}
	],
	"trace_id":"60ae120f-371d5872-7941a05b"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/wxa/msg_sec_check?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	mp := New("APPID", "APPSECRET", WithMockClient(client))

	params := &ParamsMsgCheck{
		Content: "hello world!",
		Version: 2,
		Scene:   SecSceneDoc,
		OpenID:  "OPENID",
	}

	result := new(ResultMsgCheck)

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", MsgSecCheck(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMsgCheck{
		TraceID: "60ae120f-371d5872-7941a05b",
		Result: &MsgCheckRet{
			Suggest: "risky",
			Label:   20001,
		},
		Detail: []*MsgCheckItem{
			{
				Strategy: "content_model",
				ErrCode:  0,
				Suggest:  "risky",
				Label:    20006,
				Prob:     90,
			},
			{
				Strategy: "keyword",
				ErrCode:  0,
				Suggest:  "pass",
				Label:    20006,
				Keyword:  "命中的关键词1",
			},
			{
				Strategy: "keyword",
				ErrCode:  0,
				Suggest:  "risky",
				Label:    20006,
				Keyword:  "命中的关键词2",
			},
		},
	}, result)
}

func TestGetUserRiskRank(t *testing.T) {
	body := []byte(`{"appid":"APPID","openid":"OPENID","scene":1,"mobile_no":"12345678","client_ip":"******","email_address":"****@qq.com"}`)

	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"risk_rank": 0
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/wxa/getuserriskrank?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	mp := New("APPID", "APPSECRET", WithMockClient(client))

	params := &ParamsUserRisk{
		AppID:        "APPID",
		OpenID:       "OPENID",
		Scene:        RiskCheat,
		MobileNO:     "12345678",
		ClientIP:     "******",
		EmailAddress: "****@qq.com",
	}

	result := new(ResultUserRisk)

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", GetUserRiskRank(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultUserRisk{
		RiskRank: 0,
	}, result)
}
