package offia

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/chenghonour/gochat/mock"
	"github.com/chenghonour/gochat/wx"
)

func TestUploadMedia(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"type": "image",
	"media_id": "MEDIA_ID",
	"created_at": 1606717010
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=image", gomock.AssignableToTypeOf(wx.NewUploadForm())).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(ResultMediaUpload)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UploadMedia(MediaImage, "../mock/test.jpg", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMediaUpload{
		Type:      "image",
		MediaID:   "MEDIA_ID",
		CreatedAt: 1606717010,
	}, result)
}

func TestUploadMediaByURL(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"type": "image",
	"media_id": "MEDIA_ID",
	"created_at": 1606717010
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=image", gomock.AssignableToTypeOf(wx.NewUploadForm())).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(ResultMediaUpload)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UploadMediaByURL(MediaImage, "test.png", "https://golang.google.cn/doc/gopher/pkg.png", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMediaUpload{
		Type:      "image",
		MediaID:   "MEDIA_ID",
		CreatedAt: 1606717010,
	}, result)
}

func TestAddMaterial(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"media_id": "MEDIA_ID",
	"url": "URL"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=ACCESS_TOKEN&type=image", gomock.AssignableToTypeOf(wx.NewUploadForm())).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(ResultMaterialAdd)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", AddMaterial(MediaImage, "../mock/test.jpg", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMaterialAdd{
		MediaID: "MEDIA_ID",
		URL:     "URL",
	}, result)
}

func TestAddMaterialByURL(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"media_id": "MEDIA_ID",
	"url": "URL"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=ACCESS_TOKEN&type=image", gomock.AssignableToTypeOf(wx.NewUploadForm())).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(ResultMaterialAdd)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", AddMaterialByURL(MediaImage, "test.png", "https://golang.google.cn/doc/gopher/pkg.png", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMaterialAdd{
		MediaID: "MEDIA_ID",
		URL:     "URL",
	}, result)
}

func TestGetNewsMaterial(t *testing.T) {
	body := []byte(`{"media_id":"MEDIA_ID"}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"news_item": [
		{
			"title": "TITLE",
			"thumb_media_id": "THUMB_MEDIA_ID",
			"show_cover_pic": 1,
			"author": "AUTHOR",
			"digest": "DIGEST",
			"content": "CONTENT",
			"url": "URL",
			"content_source_url": "CONTENT_SOURCE_URL"
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/material/get_material?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(ResultNewsMaterialGet)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetNewsMaterial("MEDIA_ID", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultNewsMaterialGet{
		NewsItem: []*MaterialNewsItem{
			{
				Title:            "TITLE",
				ThumbMediaID:     "THUMB_MEDIA_ID",
				ShowCoverPic:     1,
				Author:           "AUTHOR",
				Digest:           "DIGEST",
				Content:          "CONTENT",
				URL:              "URL",
				ContentSourceURL: "CONTENT_SOURCE_URL",
			},
		},
	}, result)
}

func TestGetVideoMaterial(t *testing.T) {
	body := []byte(`{"media_id":"MEDIA_ID"}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"title": "TITLE",
	"description": "DESCRIPTION",
	"down_url": "DOWN_URL"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/material/get_material?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(ResultVideoMaterialGet)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetVideoMaterial("MEDIA_ID", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultVideoMaterialGet{
		Title:       "TITLE",
		Description: "DESCRIPTION",
		DownURL:     "DOWN_URL",
	}, result)
}

func TestDeleteMaterial(t *testing.T) {
	body := []byte(`{"media_id":"MEDIA_ID"}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/material/del_material?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", DeleteMaterial("MEDIA_ID"))

	assert.Nil(t, err)
}

func TestUploadImg(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"url": "URL"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=ACCESS_TOKEN", gomock.AssignableToTypeOf(wx.NewUploadForm())).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(ResultMaterialAdd)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UploadImg("../mock/test.jpg", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMaterialAdd{
		URL: "URL",
	}, result)
}

func TestUploadImgByURL(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"url": "URL"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=ACCESS_TOKEN", gomock.AssignableToTypeOf(wx.NewUploadForm())).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(ResultMaterialAdd)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UploadImgByURL("test.png", "https://golang.google.cn/doc/gopher/pkg.png", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMaterialAdd{
		URL: "URL",
	}, result)
}

func TestUploadVideo(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"media_id": "MEDIA_ID",
	"url": "URL"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=ACCESS_TOKEN&type=video", gomock.AssignableToTypeOf(wx.NewUploadForm())).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(ResultMaterialAdd)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UploadVideo("../mock/test.mp4", "TITLE", "INTRODUCTION", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMaterialAdd{
		MediaID: "MEDIA_ID",
		URL:     "URL",
	}, result)
}

func TestUploadVideoByURL(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"media_id": "MEDIA_ID",
	"url": "URL"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=ACCESS_TOKEN&type=video", gomock.AssignableToTypeOf(wx.NewUploadForm())).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(ResultMaterialAdd)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UploadVideoByURL("test.mp4", "https://video.ivwen.com/users/4576112/46e9506e35534ddb961772727f32399d.mp4", "TITLE", "INTRODUCTION", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMaterialAdd{
		MediaID: "MEDIA_ID",
		URL:     "URL",
	}, result)
}

func TestAddNews(t *testing.T) {
	body := []byte(`{"articles":[{"title":"TITLE","thumb_media_id":"THUMB_MEDIA_ID","author":"AUTHOR","digest":"DIGEST","show_cover_pic":1,"content":"CONTENT","content_source_url":"CONTENT_SOURCE_URL","need_open_comment":1,"only_fans_can_comment":1}]}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"media_id": "MEDIA_ID"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/material/add_news?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	articles := []*NewsArticle{
		{
			Title:              "TITLE",
			ThumbMediaID:       "THUMB_MEDIA_ID",
			Author:             "AUTHOR",
			Digest:             "DIGEST",
			ShowCoverPic:       1,
			Content:            "CONTENT",
			ContentSourceURL:   "CONTENT_SOURCE_URL",
			NeedOpenComment:    1,
			OnlyFansCanComment: 1,
		},
	}

	result := new(ResultMaterialAdd)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", AddNews(articles, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMaterialAdd{
		MediaID: "MEDIA_ID",
	}, result)
}

func TestUpdateNews(t *testing.T) {
	body := []byte(`{"media_id":"MEDIA_ID","index":"INDEX","articles":{"title":"TITLE","thumb_media_id":"THUMB_MEDIA_ID","author":"AUTHOR","digest":"DIGEST","show_cover_pic":1,"content":"CONTENT","content_source_url":"CONTENT_SOURCE_URL"}}`)

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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/material/update_news?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	article := &NewsArticle{
		Title:            "TITLE",
		ThumbMediaID:     "THUMB_MEDIA_ID",
		Author:           "AUTHOR",
		Digest:           "DIGEST",
		ShowCoverPic:     1,
		Content:          "CONTENT",
		ContentSourceURL: "CONTENT_SOURCE_URL",
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UpdateNews("MEDIA_ID", "INDEX", article))

	assert.Nil(t, err)
}

func TestGetMaterialCount(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"voice_count": 1,
	"video_count": 2,
	"image_count": 3,
	"news_count": 4
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://api.weixin.qq.com/cgi-bin/material/get_materialcount?access_token=ACCESS_TOKEN", nil).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(ResultMaterialCount)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetMaterialCount(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMaterialCount{
		VoiceCount: 1,
		VideoCount: 2,
		ImageCount: 3,
		NewsCount:  4,
	}, result)
}

func TestListMaterial(t *testing.T) {
	body := []byte(`{"type":"image","offset":0,"count":10}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"total_count": 10,
	"item_count": 1,
	"item": [
		{
			"media_id": "MEDIA_ID",
			"name": "NAME",
			"update_time": 1643266823,
			"url": "URL"
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(ResultMaterialList)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", ListMatertial(MediaImage, 0, 10, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMaterialList{
		TotalCount: 10,
		ItemCount:  1,
		Item: []*MaterialListItem{
			{
				MediaID:    "MEDIA_ID",
				Name:       "NAME",
				UpdateTime: 1643266823,
				URL:        "URL",
			},
		},
	}, result)
}

func TestListMaterialNews(t *testing.T) {
	body := []byte(`{"type":"news","offset":0,"count":10}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"total_count": 10,
	"item_count": 1,
	"item": [
		{
			"media_id": "MEDIA_ID",
			"update_time": 1643266823,
			"content": {
				"news_item": [
					{
						"title": "TITLE",
						"thumb_media_id": "THUMB_MEDIA_ID",
						"show_cover_pic": 1,
						"author": "AUTHOR",
						"digest": "DIGEST",
						"content": "CONTENT",
						"url": "URL",
						"content_source_url": "CONTETN_SOURCE_URL"
					}
				]
			}
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(ResultMaterialNewsList)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", ListMaterialNews(0, 10, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMaterialNewsList{
		TotalCount: 10,
		ItemCount:  1,
		Item: []*MaterialNewsListItem{
			{
				MediaID:    "MEDIA_ID",
				UpdateTime: 1643266823,
				Content: &MaterialNewsListContent{
					NewsItem: []*MaterialNewsItem{
						{
							Title:            "TITLE",
							ThumbMediaID:     "THUMB_MEDIA_ID",
							ShowCoverPic:     1,
							Author:           "AUTHOR",
							Digest:           "DIGEST",
							Content:          "CONTENT",
							URL:              "URL",
							ContentSourceURL: "CONTETN_SOURCE_URL",
						},
					},
				},
			},
		},
	}, result)
}
