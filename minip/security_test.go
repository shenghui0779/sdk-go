package minip

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/yiigo"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/mock"
	"github.com/shenghui0779/gochat/wx"
)

func TestImageSecCheck(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/img_sec_check?access_token=ACCESS_TOKEN", gomock.AssignableToTypeOf(yiigo.NewUploadForm())).Return(resp, nil)

	mp := New("APPID", "APPSECRET")
	mp.SetClient(wx.WithHTTPClient(client))

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", ImageSecCheck("../mock/test.jpg"))

	assert.Nil(t, err)
}

func TestMediaCheckAsync(t *testing.T) {
	body := []byte(`{"media_type":2,"media_url":"https://developers.weixin.qq.com/miniprogram/assets/images/head_global_z_@all.png"}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"trace_id": "967e945cd8a3e458f3c74dcb886068e9"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/wxa/media_check_async?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	mp := New("APPID", "APPSECRET")
	mp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultMediaCheckAsync)

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", MediaCheckAsync(SecMediaImage, "https://developers.weixin.qq.com/miniprogram/assets/images/head_global_z_@all.png", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMediaCheckAsync{
		TraceID: "967e945cd8a3e458f3c74dcb886068e9",
	}, result)
}

func TestMsgSecCheck(t *testing.T) {
	body := []byte(`{"content":"hello world!"}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/wxa/msg_sec_check?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	mp := New("APPID", "APPSECRET")
	mp.SetClient(wx.WithHTTPClient(client))

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", MsgSecCheck("hello world!"))

	assert.Nil(t, err)
}
