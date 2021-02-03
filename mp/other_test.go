package mp

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/wx"
	"github.com/stretchr/testify/assert"
)

func TestInvokeService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/servicemarket?access_token=ACCESS_TOKEN", []byte(`{"service":"wx79ac3de8be320b71","api":"OcrAllInOne","data":{"data_type":3,"img_url":"http://mmbiz.qpic.cn/mmbiz_jpg/7UFjuNbYxibu66xSqsQqKcuoGBZM77HIyibdiczeWibdMeA2XMt5oibWVQMgDibriazJSOibLqZxcO6DVVcZMxDKgeAtbQ/0","ocr_type":1},"client_msg_id":"id123"}`)).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"data": "{\"idcard_res\":{\"type\":0,\"name\":{\"text\":\"abc\",\"pos\"…0312500}}},\"image_width\":480,\"image_height\":304}}",
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	data := &InvokeData{
		Service: "wx79ac3de8be320b71",
		API:     "OcrAllInOne",
		Data: wx.X{
			"img_url":   "http://mmbiz.qpic.cn/mmbiz_jpg/7UFjuNbYxibu66xSqsQqKcuoGBZM77HIyibdiczeWibdMeA2XMt5oibWVQMgDibriazJSOibLqZxcO6DVVcZMxDKgeAtbQ/0",
			"data_type": 3,
			"ocr_type":  1,
		},
		ClientMsgID: "id123",
	}

	dest := new(InvokeResult)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", InvokeService(dest, data))

	assert.Nil(t, err)
	assert.Equal(t, `{"idcard_res":{"type":0,"name":{"text":"abc","pos"…0312500}}},"image_width":480,"image_height":304}}`, dest.Data)
}

func TestSoterVerify(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/soter/verify_signature?access_token=ACCESS_TOKEN", []byte(`{"openid":"$openid","json_string":"$resultJSON","json_signature":"$resultJSONSignature"}`)).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"is_ok": true
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	sign := &SoterSignature{
		OpenID:        "$openid",
		JSONString:    "$resultJSON",
		JSONSignature: "$resultJSONSignature",
	}

	dest := new(SoterVerifyResult)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SoterVerify(dest, sign))

	assert.Nil(t, err)
	assert.True(t, dest.OK)
}

func TestGetUserRiskRank(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/getuserriskrank?access_token=ACCESS_TOKEN", []byte(`{"appid":"APPID","openid":"OPENID","scene":1,"mobile_no":"12345678","client_ip":"******","email_address":"****@qq.com"}`)).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"risk_rank": 0
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	data := &UserRiskData{
		AppID:        "APPID",
		OpenID:       "OPENID",
		Scene:        RiskCheat,
		MobileNO:     "12345678",
		ClientIP:     "******",
		EmailAddress: "****@qq.com",
	}

	dest := new(UserRiskResult)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetUserRiskRank(dest, data))

	assert.Nil(t, err)
	assert.Equal(t, 0, dest.RiskRank)
}
