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

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/servicemarket?access_token=ACCESS_TOKEN", gomock.AssignableToTypeOf(postBody)).Return([]byte(`{
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

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", InvokeService(data, dest))

	assert.Nil(t, err)
	assert.Equal(t, `{"idcard_res":{"type":0,"name":{"text":"abc","pos"…0312500}}},"image_width":480,"image_height":304}}`, dest.Data)
}

func TestSoterVerify(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/soter/verify_signature?access_token=ACCESS_TOKEN", gomock.AssignableToTypeOf(postBody)).Return([]byte(`{
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

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SoterVerify(sign, dest))

	assert.Nil(t, err)
	assert.True(t, dest.OK)
}
