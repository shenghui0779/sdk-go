package mp

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/wx"
	"github.com/stretchr/testify/assert"
)

func TestImageSecCheck(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/img_sec_check?access_token=ACCESS_TOKEN", wx.NewUploadForm("media", "test.jpg", nil)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", ImageSecCheck("test.jpg"))

	assert.Nil(t, err)
}

func TestMediaCheckAsync(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/media_check_async?access_token=ACCESS_TOKEN", []byte(`{"media_type":2,"media_url":"https://developers.weixin.qq.com/miniprogram/assets/images/head_global_z_@all.png"}`)).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"trace_id": "967e945cd8a3e458f3c74dcb886068e9"
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(MediaSecAsyncResult)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", MediaSecCheckAsync(dest, SecMediaImage, "https://developers.weixin.qq.com/miniprogram/assets/images/head_global_z_@all.png"))

	assert.Nil(t, err)
	assert.Equal(t, "967e945cd8a3e458f3c74dcb886068e9", dest.TraceID)
}

func TestMsgSecCheck(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/msg_sec_check?access_token=ACCESS_TOKEN", []byte(`{"content":"hello world!"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", MsgSecCheck("hello world!"))

	assert.Nil(t, err)
}
