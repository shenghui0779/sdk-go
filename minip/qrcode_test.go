package minip

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/mock"
)

func TestCreateQRCode(t *testing.T) {
	body := []byte(`{"path":"page/index/index","width":430}`)

	resp := []byte("BUFFER")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	mp := New("APPID", "APPSECRET", WithMockClient(client))

	qrcode := new(QRCode)

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", CreateQRCode("page/index/index", 430, qrcode))

	assert.Nil(t, err)
	assert.Equal(t, "BUFFER", string(qrcode.Buffer))
}

func TestGetQRCode(t *testing.T) {
	body := []byte(`{"path":"page/index/index","width":430}`)

	resp := []byte("BUFFER")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/wxa/getwxacode?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	mp := New("APPID", "APPSECRET", WithMockClient(client))

	params := &ParamsQRCodeGet{
		Path:  "page/index/index",
		Width: 430,
	}
	qrcode := new(QRCode)

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", GetQRCode(params, qrcode))

	assert.Nil(t, err)
	assert.Equal(t, "BUFFER", string(qrcode.Buffer))
}

func TestGetUnlimitQRCode(t *testing.T) {
	body := []byte(`{"scene":"a=1"}`)

	resp := []byte("BUFFER")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	mp := New("APPID", "APPSECRET", WithMockClient(client))

	params := &ParamsQRCodeUnlimit{
		Scene: "a=1",
	}
	qrcode := new(QRCode)

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", GetUnlimitQRCode(params, qrcode))

	assert.Nil(t, err)
	assert.Equal(t, "BUFFER", string(qrcode.Buffer))
}
