package oa

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/wx"
	"github.com/stretchr/testify/assert"
)

func TestCreateTempQRCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=ACCESS_TOKEN", []byte(`{"action_info":{"scene":{"scene_id":123}},"action_name":"QR_SCENE","expire_seconds":60}`)).Return([]byte(`{
		"ticket": "gQH47joAAAAAAAAAASxodHRwOi8vd2VpeGluLnFxLmNvbS9xL2taZ2Z3TVRtNzJXV1Brb3ZhYmJJAAIEZ23sUwMEmm3sUw==",
		"expire_seconds": 60,
		"url": "http://weixin.qq.com/q/kZgfwMTm72WWPkovabbI"
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(QRCode)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", CreateTempQRCode(dest, 123, 60))

	assert.Nil(t, err)
	assert.Equal(t, &QRCode{
		Ticket:        "gQH47joAAAAAAAAAASxodHRwOi8vd2VpeGluLnFxLmNvbS9xL2taZ2Z3TVRtNzJXV1Brb3ZhYmJJAAIEZ23sUwMEmm3sUw==",
		ExpireSeconds: 60,
		URL:           "http://weixin.qq.com/q/kZgfwMTm72WWPkovabbI",
	}, dest)
}

func TestCreatePermQRCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=ACCESS_TOKEN", []byte(`{"action_info":{"scene":{"scene_id":123}},"action_name":"QR_LIMIT_SCENE"}`)).Return([]byte(`{
		"ticket": "gQH47joAAAAAAAAAASxodHRwOi8vd2VpeGluLnFxLmNvbS9xL2taZ2Z3TVRtNzJXV1Brb3ZhYmJJAAIEZ23sUwMEmm3sUw==",
		"expire_seconds": 60,
		"url": "http://weixin.qq.com/q/kZgfwMTm72WWPkovabbI"
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(QRCode)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", CreatePermQRCode(dest, 123))

	assert.Nil(t, err)
	assert.Equal(t, &QRCode{
		Ticket:        "gQH47joAAAAAAAAAASxodHRwOi8vd2VpeGluLnFxLmNvbS9xL2taZ2Z3TVRtNzJXV1Brb3ZhYmJJAAIEZ23sUwMEmm3sUw==",
		ExpireSeconds: 60,
		URL:           "http://weixin.qq.com/q/kZgfwMTm72WWPkovabbI",
	}, dest)
}

func TestLong2ShortURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/shorturl?access_token=ACCESS_TOKEN", []byte(`{"action":"long2short","long_url":"http://wap.koudaitong.com/v2/showcase/goods?alias=128wi9shh&spm=h56083&redirect_count=1"}`)).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"short_url": "http:\/\/w.url.cn\/s\/AvCo6Ih"
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(ShortURL)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", Long2ShortURL(dest, "http://wap.koudaitong.com/v2/showcase/goods?alias=128wi9shh&spm=h56083&redirect_count=1"))

	assert.Nil(t, err)
	assert.Equal(t, "http://w.url.cn/s/AvCo6Ih", dest.URL)
}
