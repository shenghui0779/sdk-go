package minip

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/yiigo"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/wx"
)

func TestInvokeService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/servicemarket?access_token=ACCESS_TOKEN", []byte(`{"service":"wx79ac3de8be320b71","api":"OcrAllInOne","data":{"data_type":3,"img_url":"http://mmbiz.qpic.cn/mmbiz_jpg/7UFjuNbYxibu66xSqsQqKcuoGBZM77HIyibdiczeWibdMeA2XMt5oibWVQMgDibriazJSOibLqZxcO6DVVcZMxDKgeAtbQ/0","ocr_type":1},"client_msg_id":"id123"}`)).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"data": "{\"idcard_res\":{\"type\":0,\"name\":{\"text\":\"abc\",\"pos\"…0312500}}},\"image_width\":480,\"image_height\":304}}"
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	params := &ParamsServiceInvoke{
		Service: "wx79ac3de8be320b71",
		API:     "OcrAllInOne",
		Data: yiigo.X{
			"data_type": 3,
			"img_url":   "http://mmbiz.qpic.cn/mmbiz_jpg/7UFjuNbYxibu66xSqsQqKcuoGBZM77HIyibdiczeWibdMeA2XMt5oibWVQMgDibriazJSOibLqZxcO6DVVcZMxDKgeAtbQ/0",
			"ocr_type":  1,
		},
		ClientMsgID: "id123",
	}

	result := new(ResultServiceInvoke)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", InvokeService(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultServiceInvoke{
		Data: `{"idcard_res":{"type":0,"name":{"text":"abc","pos"…0312500}}},"image_width":480,"image_height":304}}`,
	}, result)
}

func TestSoterVerify(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/soter/verify_signature?access_token=ACCESS_TOKEN", []byte(`{"openid":"$openid","json_string":"$resultJSON","json_signature":"$resultJSONSignature"}`)).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"is_ok": true
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	params := &ParamsSoterVerify{
		OpenID:        "$openid",
		JSONString:    "$resultJSON",
		JSONSignature: "$resultJSONSignature",
	}

	result := new(ResultSoterVerify)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SoterVerify(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultSoterVerify{
		IsOK: true,
	}, result)
}

func TestGetUserRiskRank(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/getuserriskrank?access_token=ACCESS_TOKEN", []byte(`{"appid":"APPID","openid":"OPENID","scene":1,"mobile_no":"12345678","client_ip":"******","email_address":"****@qq.com"}`)).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"risk_rank": 0
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	params := &ParamsUserRisk{
		AppID:        "APPID",
		OpenID:       "OPENID",
		Scene:        RiskCheat,
		MobileNO:     "12345678",
		ClientIP:     "******",
		EmailAddress: "****@qq.com",
	}

	result := new(ResultUserRisk)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetUserRiskRank(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultUserRisk{
		RiskRank: 0,
	}, result)
}
