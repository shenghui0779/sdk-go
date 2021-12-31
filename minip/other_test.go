package minip

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/mock"
	"github.com/shenghui0779/gochat/wx"
	"github.com/shenghui0779/yiigo"
	"github.com/stretchr/testify/assert"
)

func TestInvokeService(t *testing.T) {
	body := []byte(`{"service":"wx79ac3de8be320b71","api":"OcrAllInOne","data":{"data_type":3,"img_url":"http://mmbiz.qpic.cn/mmbiz_jpg/7UFjuNbYxibu66xSqsQqKcuoGBZM77HIyibdiczeWibdMeA2XMt5oibWVQMgDibriazJSOibLqZxcO6DVVcZMxDKgeAtbQ/0","ocr_type":1},"client_msg_id":"id123"}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"data": "{\"idcard_res\":{\"type\":0,\"name\":{\"text\":\"abc\",\"pos\"…0312500}}},\"image_width\":480,\"image_height\":304}}"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/wxa/servicemarket?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	mp := New("APPID", "APPSECRET")
	mp.SetClient(wx.WithHTTPClient(client))

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

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", InvokeService(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultServiceInvoke{
		Data: `{"idcard_res":{"type":0,"name":{"text":"abc","pos"…0312500}}},"image_width":480,"image_height":304}}`,
	}, result)
}

func TestSoterVerify(t *testing.T) {
	body := []byte(`{"openid":"$openid","json_string":"$resultJSON","json_signature":"$resultJSONSignature"}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"is_ok": true
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/soter/verify_signature?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	mp := New("APPID", "APPSECRET")
	mp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsSoterVerify{
		OpenID:        "$openid",
		JSONString:    "$resultJSON",
		JSONSignature: "$resultJSONSignature",
	}

	result := new(ResultSoterVerify)

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", SoterVerify(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultSoterVerify{
		IsOK: true,
	}, result)
}

func TestGetUserRiskRank(t *testing.T) {
	body := []byte(`{"appid":"APPID","openid":"OPENID","scene":1,"mobile_no":"12345678","client_ip":"******","email_address":"****@qq.com"}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"risk_rank": 0
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/wxa/getuserriskrank?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	mp := New("APPID", "APPSECRET")
	mp.SetClient(wx.WithHTTPClient(client))

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
