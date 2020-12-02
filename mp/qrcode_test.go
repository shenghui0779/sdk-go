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

	o := new(qrcodeOptions)

	for _, option := range options {
		option.apply(o)
	}

	assert.Equal(t, 430, o.width)
	assert.Equal(t, "pages/index", o.page)
	assert.True(t, o.autoColor)
	assert.Equal(t, map[string]int{
		"r": 1,
		"g": 2,
		"b": 3,
	}, o.lineColor)
	assert.True(t, o.isHyaline)
}

func TestCreateQRCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token=ACCESS_TOKEN", gomock.AssignableToTypeOf(postBody)).Return([]byte("BUFFER"), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(QRCode)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", CreateQRCode("page/index/index", dest))

	assert.Nil(t, err)
	assert.Equal(t, "BUFFER", string(dest.Buffer))
}

func TestGetQRCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/getwxacode?access_token=ACCESS_TOKEN", gomock.AssignableToTypeOf(postBody)).Return([]byte("BUFFER"), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(QRCode)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetQRCode("page/index/index", dest))

	assert.Nil(t, err)
	assert.Equal(t, "BUFFER", string(dest.Buffer))
}

func TestGetUnlimitQRCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=ACCESS_TOKEN", gomock.AssignableToTypeOf(postBody)).Return([]byte("BUFFER"), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(QRCode)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetUnlimitQRCode("a=1", dest))

	assert.Nil(t, err)
	assert.Equal(t, "BUFFER", string(dest.Buffer))
}
