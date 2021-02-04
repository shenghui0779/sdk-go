package mp

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/wx"
	"github.com/stretchr/testify/assert"
)

func TestQRCodeOption(t *testing.T) {
	options := []QRCodeOption{
		WithQRCodeWidth(430),
		WithQRCodePage("pages/index"),
		WithQRCodeAutoColor(),
		WithQRCodeLineColor(1, 2, 3),
		WithQRCodeIsHyaline(),
	}

	settings := new(qrcodeSettings)

	for _, f := range options {
		f(settings)
	}

	assert.Equal(t, 430, settings.width)
	assert.Equal(t, "pages/index", settings.page)
	assert.True(t, settings.autoColor)
	assert.Equal(t, map[string]int{
		"r": 1,
		"g": 2,
		"b": 3,
	}, settings.lineColor)
	assert.True(t, settings.isHyaline)
}

func TestCreateQRCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token=ACCESS_TOKEN", []byte(`{"path":"page/index/index","width":430}`)).Return([]byte("BUFFER"), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(QRCode)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", CreateQRCode(dest, "page/index/index", WithQRCodeWidth(430)))

	assert.Nil(t, err)
	assert.Equal(t, "BUFFER", string(dest.Buffer))
}

func TestGetQRCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/getwxacode?access_token=ACCESS_TOKEN", []byte(`{"path":"page/index/index","width":430}`)).Return([]byte("BUFFER"), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(QRCode)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetQRCode(dest, "page/index/index", WithQRCodeWidth(430)))

	assert.Nil(t, err)
	assert.Equal(t, "BUFFER", string(dest.Buffer))
}

func TestGetUnlimitQRCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=ACCESS_TOKEN", []byte(`{"scene":"a=1"}`)).Return([]byte("BUFFER"), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(QRCode)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetUnlimitQRCode(dest, "a=1"))

	assert.Nil(t, err)
	assert.Equal(t, "BUFFER", string(dest.Buffer))
}
