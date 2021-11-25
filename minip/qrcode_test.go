package minip

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/mock"
	"github.com/stretchr/testify/assert"
)

func TestCreateQRCode(t *testing.T) {
	body := []byte(`{"path":"page/index/index","width":430}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte("BUFFER"))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	mp := New("APPID", "APPSECRET")
	mp.SetClient(client)

	params := &ParamsQRCodeCreate{
		Path:  "page/index/index",
		Width: 430,
	}
	qrcode := new(QRCode)

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", CreateQRCode(params, qrcode))

	assert.Nil(t, err)
	assert.Equal(t, "BUFFER", string(qrcode.Buffer))
}

func TestGetQRCode(t *testing.T) {
	body := []byte(`{"path":"page/index/index","width":430}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte("BUFFER"))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/wxa/getwxacode?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	mp := New("APPID", "APPSECRET")
	mp.SetClient(client)

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

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte("BUFFER"))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	mp := New("APPID", "APPSECRET")
	mp.SetClient(client)

	params := &ParamsQRCodeUnlimit{
		Scene: "a=1",
	}
	qrcode := new(QRCode)

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", GetUnlimitQRCode(params, qrcode))

	assert.Nil(t, err)
	assert.Equal(t, "BUFFER", string(qrcode.Buffer))
}
