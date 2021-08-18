package offia

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

	dest := new(MediaUploadResult)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UploadMedia(dest, MediaImage, "../test/test.jpg"))

	assert.Nil(t, err)
	assert.Equal(t, &MediaUploadResult{
		Type:      "image",
		MediaID:   "MEDIA_ID",
		CreatedAt: 1606717010,
	}, dest)
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

	dest := new(MediaUploadResult)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UploadMediaByURL(dest, MediaImage, "test.png", "https://golang.google.cn/doc/gopher/pkg.png"))

	assert.Nil(t, err)
	assert.Equal(t, &MediaUploadResult{
		Type:      "image",
		MediaID:   "MEDIA_ID",
		CreatedAt: 1606717010,
	}, dest)
}

func TestAddNews(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/material/add_news?access_token=ACCESS_TOKEN", []byte(`{"articles":[{"title":"TITLE","thumb_media_id":"THUMB_MEDIA_ID","author":"AUTHOR","digest":"DIGEST","show_cover_pic":1,"content":"CONTENT","content_source_url":"CONTENT_SOURCE_URL","need_open_comment":1,"only_fans_can_comment":1}]}`)).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"media_id": "MEDIA_ID"
	  }`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(MaterialAddResult)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", AddNews(dest, &NewsArticle{
		Title:              "TITLE",
		ThumbMediaID:       "THUMB_MEDIA_ID",
		Author:             "AUTHOR",
		Digest:             "DIGEST",
		ShowCoverPic:       1,
		Content:            "CONTENT",
		ContentSourceURL:   "CONTENT_SOURCE_URL",
		NeedOpenComment:    1,
		OnlyFansCanComment: 1,
	}))

	assert.Nil(t, err)
	assert.Equal(t, &MaterialAddResult{
		MediaID: "MEDIA_ID",
	}, dest)
}

func TestUploadNewsImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=ACCESS_TOKEN", gomock.AssignableToTypeOf(yiigo.NewUploadForm())).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"url": "URL"
	  }`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(MaterialAddResult)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UploadNewsImage(dest, "../test/test.jpg"))

	assert.Nil(t, err)
	assert.Equal(t, &MaterialAddResult{
		URL: "URL",
	}, dest)
}

func TestUploadNewsImageByURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=ACCESS_TOKEN", gomock.AssignableToTypeOf(yiigo.NewUploadForm())).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"url": "URL"
	  }`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(MaterialAddResult)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UploadNewsImageByURL(dest, "test.png", "https://golang.google.cn/doc/gopher/pkg.png"))

	assert.Nil(t, err)
	assert.Equal(t, &MaterialAddResult{
		URL: "URL",
	}, dest)
}

func TestAddMaterial(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=ACCESS_TOKEN&type=image", gomock.AssignableToTypeOf(yiigo.NewUploadForm())).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"media_id": "MEDIA_ID",
		"url": "URL"
	  }`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(MaterialAddResult)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", AddMaterial(dest, MediaImage, "../test/test.jpg"))

	assert.Nil(t, err)
	assert.Equal(t, &MaterialAddResult{
		MediaID: "MEDIA_ID",
		URL:     "URL",
	}, dest)
}

func TestAddMaterialByURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=ACCESS_TOKEN&type=image", gomock.AssignableToTypeOf(yiigo.NewUploadForm())).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"media_id": "MEDIA_ID",
		"url": "URL"
	  }`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(MaterialAddResult)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", AddMaterialByURL(dest, MediaImage, "test.png", "https://golang.google.cn/doc/gopher/pkg.png"))

	assert.Nil(t, err)
	assert.Equal(t, &MaterialAddResult{
		MediaID: "MEDIA_ID",
		URL:     "URL",
	}, dest)
}

func TestUploadVideo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=ACCESS_TOKEN&type=video", gomock.AssignableToTypeOf(yiigo.NewUploadForm())).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"media_id": "MEDIA_ID",
		"url": "URL"
	  }`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(MaterialAddResult)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UploadVideo(dest, "../test/test.mp4", "TITLE", "INTRODUCTION"))

	assert.Nil(t, err)
	assert.Equal(t, &MaterialAddResult{
		MediaID: "MEDIA_ID",
		URL:     "URL",
	}, dest)
}

func TestUploadVideoByURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=ACCESS_TOKEN&type=video", gomock.AssignableToTypeOf(yiigo.NewUploadForm())).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"media_id": "MEDIA_ID",
		"url": "URL"
	  }`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(MaterialAddResult)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UploadVideoByURL(dest, "test.mp4", "TITLE", "INTRODUCTION", "https://video.ivwen.com/users/4576112/46e9506e35534ddb961772727f32399d.mp4"))

	assert.Nil(t, err)
	assert.Equal(t, &MaterialAddResult{
		MediaID: "MEDIA_ID",
		URL:     "URL",
	}, dest)
}

func TestDeleteMaterial(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/material/del_material?access_token=ACCESS_TOKEN", []byte(`{"media_id":"MEDIA_ID"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", DeleteMaterial("MEDIA_ID"))

	assert.Nil(t, err)
}
