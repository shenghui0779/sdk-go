package minip

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/yiigo"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/wx"
)

func TestUploadMedia(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=image", gomock.AssignableToTypeOf(yiigo.NewUploadForm())).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"type": "image",
		"media_id": "MEDIA_ID",
		"created_at": 1606717010
	  }`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	params := &ParamsMediaUpload{
		MediaType: MediaImage,
		Path:      "../test/test.jpg",
	}
	result := new(ResultMediaUpload)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UploadMedia(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMediaUpload{
		Type:      "image",
		MediaID:   "MEDIA_ID",
		CreatedAt: 1606717010,
	}, result)
}

func TestUploadMediaByURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=image", gomock.AssignableToTypeOf(yiigo.NewUploadForm())).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"type": "image",
		"media_id": "MEDIA_ID",
		"created_at": 1606717010
	  }`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	params := &ParamsMediaUploadByURL{
		MediaType: MediaImage,
		Filename:  "test.png",
		URL:       "https://golang.google.cn/doc/gopher/pkg.png",
	}
	result := new(ResultMediaUpload)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UploadMediaByURL(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMediaUpload{
		Type:      "image",
		MediaID:   "MEDIA_ID",
		CreatedAt: 1606717010,
	}, result)
}

func TestGetMedia(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/media/get?access_token=ACCESS_TOKEN&media_id=MEDIA_ID").Return([]byte("BUFFER"), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	result := new(Media)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetMedia("MEDIA_ID", result))

	assert.Nil(t, err)
	assert.Equal(t, "BUFFER", string(result.Buffer))
}
