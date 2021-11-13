package minip

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/wx"
)

func TestCreateQRCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token=ACCESS_TOKEN", []byte(`{"path":"page/index/index","width":430}`)).Return([]byte("BUFFER"), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	params := &ParamsQRCodeCreate{
		Path:  "page/index/index",
		Width: 430,
	}
	qrcode := new(QRCode)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", CreateQRCode(params, qrcode))

	assert.Nil(t, err)
	assert.Equal(t, "BUFFER", string(qrcode.Buffer))
}

func TestGetQRCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/getwxacode?access_token=ACCESS_TOKEN", []byte(`{"path":"page/index/index","width":430}`)).Return([]byte("BUFFER"), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	params := &ParamsQRCodeGet{
		Path:  "page/index/index",
		Width: 430,
	}
	qrcode := new(QRCode)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetQRCode(params, qrcode))

	assert.Nil(t, err)
	assert.Equal(t, "BUFFER", string(qrcode.Buffer))
}

func TestGetUnlimitQRCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=ACCESS_TOKEN", []byte(`{"scene":"a=1"}`)).Return([]byte("BUFFER"), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	params := &ParamsQRCodeUnlimit{
		Scene: "a=1",
		Width: 430,
	}
	qrcode := new(QRCode)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetUnlimitQRCode(params, qrcode))

	assert.Nil(t, err)
	assert.Equal(t, "BUFFER", string(qrcode.Buffer))
}
