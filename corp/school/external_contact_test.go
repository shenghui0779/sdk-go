package school

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/corp"
	"github.com/shenghui0779/gochat/mock"
	"github.com/shenghui0779/gochat/wx"
)

func TestGetSubscribeQRCode(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"qrcode_big": "http://p.qpic.cn/wwhead/XXXX",
	"qrcode_middle": "http://p.qpic.cn/wwhead/XXXX",
	"qrcode_thumb": "http://p.qpic.cn/wwhead/XXXX"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_subscribe_qr_code?access_token=ACCESS_TOKEN", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultSubscribeQRCode)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetSubscribeQRCode(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultSubscribeQRCode{
		QRCodeBig:    "http://p.qpic.cn/wwhead/XXXX",
		QRCodeMiddle: "http://p.qpic.cn/wwhead/XXXX",
		QRCodeThumb:  "http://p.qpic.cn/wwhead/XXXX",
	}, result)
}

func TestSetSubscribeMode(t *testing.T) {
	body := []byte(`{"subscribe_mode":1}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/set_subscribe_mode?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SetSubscribeMode(1))

	assert.Nil(t, err)
}

func TestGetSubscribeMode(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"subscribe_mode": 1
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_subscribe_mode?access_token=ACCESS_TOKEN", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultSubscribeModeGet)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetSubscribeMode(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultSubscribeModeGet{
		SubscribeMode: 1,
	}, result)
}
