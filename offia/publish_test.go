package offia

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/mock"
)

func TestSubmitPublish(t *testing.T) {
	body := []byte(`{"media_id":"MEDIA_ID"}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"publish_id": "100000001"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/freepublish/submit?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET", WithMockClient(client))

	result := new(ResultPublishSubmit)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SubmitPublish("MEDIA_ID", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultPublishSubmit{
		PublishID: "100000001",
	}, result)
}

func TestGetPublish(t *testing.T) {
	body := []byte(`{"publish_id":"100000001"}`)
	resp := []byte(`{
	"publish_id": "100000001",
	"publish_status": 1,
	"article_id": "ARTICLE_ID",
	"article_detail": {
		"count": 1,
		"item": [
			{
				"idx": 1,
				"article_url": "ARTICLE_URL"
			}
		]
	},
	"fail_idx": [
		1,
		2
	]
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/freepublish/get?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET", WithMockClient(client))

	result := new(ResultPublishGet)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetPublish("100000001", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultPublishGet{
		PublishID:     "100000001",
		PublishStatus: 1,
		ArticleID:     "ARTICLE_ID",
		ArticleDetail: &PublishArticles{
			Count: 1,
			Item: []*PublishItem{
				{
					IDX:        1,
					ArticleURL: "ARTICLE_URL",
				},
			},
		},
		FailIDX: []int{1, 2},
	}, result)
}

func TestDeletePublish(t *testing.T) {
	body := []byte(`{"article_id":"ARTICLE_ID","index":1}`)
	resp := []byte(`{"errcode":0,"errmsg":"ok"}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/freepublish/delete?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET", WithMockClient(client))

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", DeletePublish("ARTICLE_ID", 1))

	assert.Nil(t, err)
}

func TestGetPublishArticle(t *testing.T) {
	body := []byte(`{"article_id":"ARTICLE_ID"}`)
	resp := []byte(`{
	"news_item": [
		{
			"title": "TITLE",
			"author": "AUTHOR",
			"digest": "DIGEST",
			"content": "CONTENT",
			"content_source_url": "CONTENT_SOURCE_URL",
			"thumb_media_id": "THUMB_MEDIA_ID",
			"show_cover_pic": 1,
			"need_open_comment": 0,
			"only_fans_can_comment": 0,
			"url": "URL",
			"is_deleted": false
		}
	]
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/freepublish/getarticle?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET", WithMockClient(client))

	result := new(ResultPublishArticle)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetPublishArticle("ARTICLE_ID", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultPublishArticle{
		NewsItem: []*PublishArticle{
			{
				Title:              "TITLE",
				Author:             "AUTHOR",
				Digest:             "DIGEST",
				Content:            "CONTENT",
				ContentSourceURL:   "CONTENT_SOURCE_URL",
				ThumbMediaID:       "THUMB_MEDIA_ID",
				ShowCoverPic:       1,
				NeedOpenComment:    0,
				OnlyFansCanComment: 0,
				URL:                "URL",
				IsDelete:           false,
			},
		},
	}, result)
}

func TestBatchGetPublish(t *testing.T) {
	body := []byte(`{"offset":1,"count":10,"no_content":1}`)
	resp := []byte(`{
	"total_count": 10,
	"item_count": 1,
	"item": [
		{
			"article_id": "ARTICLE_ID",
			"content": {
				"news_item": [
					{
						"title": "TITLE",
						"author": "AUTHOR",
						"digest": "DIGEST",
						"content": "CONTENT",
						"content_source_url": "CONTENT_SOURCE_URL",
						"thumb_media_id": "THUMB_MEDIA_ID",
						"show_cover_pic": 1,
						"need_open_comment": 0,
						"only_fans_can_comment": 0,
						"url": "URL",
						"is_deleted": false
					}
				]
			},
			"update_time": 1645685624
		}
	]
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/freepublish/batchget?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET", WithMockClient(client))

	result := new(ResultPublishBatchGet)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", BatchGetPublish(1, 10, 1, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultPublishBatchGet{
		TotalCount: 10,
		ItemCount:  1,
		Item: []*PublishDetail{
			{
				ArticleID: "ARTICLE_ID",
				Content: &PublishContent{
					NewsItem: []*PublishArticle{
						{
							Title:              "TITLE",
							Author:             "AUTHOR",
							Digest:             "DIGEST",
							Content:            "CONTENT",
							ContentSourceURL:   "CONTENT_SOURCE_URL",
							ThumbMediaID:       "THUMB_MEDIA_ID",
							ShowCoverPic:       1,
							NeedOpenComment:    0,
							OnlyFansCanComment: 0,
							URL:                "URL",
							IsDelete:           false,
						},
					},
				},
				UpdateTime: 1645685624,
			},
		},
	}, result)
}
