package offia

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/mock"
)

func TestAddDraft(t *testing.T) {
	body := []byte(`{"articles":[{"title":"TITLE","author":"AUTHOR","digest":"DIGEST","content":"CONTENT","content_source_url":"CONTENT_SOURCE_URL","thumb_media_id":"THUMB_MEDIA_ID"}]}`)
	resp := []byte(`{"media_id":"MEDIA_ID"}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/draft/add?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")

	params := &ParamsDraftAdd{
		Articles: []*DraftArticle{
			{
				Title:            "TITLE",
				Author:           "AUTHOR",
				Digest:           "DIGEST",
				Content:          "CONTENT",
				ContentSourceURL: "CONTENT_SOURCE_URL",
				ThumbMediaID:     "THUMB_MEDIA_ID",
			},
		},
	}
	result := new(ResultDraftAdd)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", AddDraft(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultDraftAdd{
		MediaID: "MEDIA_ID",
	}, result)
}

func TestUpdateDraft(t *testing.T) {
	body := []byte(`{"media_id":"MEDIA_ID","index":1,"articles":{"title":"TITLE","author":"AUTHOR","digest":"DIGEST","content":"CONTENT","content_source_url":"CONTENT_SOURCE_URL","thumb_media_id":"THUMB_MEDIA_ID","show_cover_pic":1}}`)
	resp := []byte(`{"errcode":0,"errmsg":"ok"}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/draft/update?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")

	article := &DraftArticle{
		Title:            "TITLE",
		Author:           "AUTHOR",
		Digest:           "DIGEST",
		Content:          "CONTENT",
		ContentSourceURL: "CONTENT_SOURCE_URL",
		ThumbMediaID:     "THUMB_MEDIA_ID",
		ShowCoverPic:     1,
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UpdateDraft("MEDIA_ID", 1, article))

	assert.Nil(t, err)
}

func TestGetDraft(t *testing.T) {
	body := []byte(`{"media_id":"MEDIA_ID"}`)
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
			"url": "URL"
		}
	]
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/draft/get?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")

	result := new(ResultDraftGet)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetDraft("MEDIA_ID", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultDraftGet{
		NewsItem: []*DraftArticle{
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
			},
		},
	}, result)
}

func TestDeleteDraft(t *testing.T) {
	body := []byte(`{"media_id":"MEDIA_ID"}`)
	resp := []byte(`{"errcode":0,"errmsg":"ok"}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/draft/delete?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", DeleteDraft("MEDIA_ID"))

	assert.Nil(t, err)
}

func TestGetDraftCount(t *testing.T) {
	resp := []byte(`{"total_count":10}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://api.weixin.qq.com/cgi-bin/draft/count?access_token=ACCESS_TOKEN", nil).Return(resp, nil)

	oa := New("APPID", "APPSECRET")

	result := new(ResultDraftCount)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetDraftCount(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultDraftCount{
		TotalCount: 10,
	}, result)
}

func TestBatchGetDraft(t *testing.T) {
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
						"url": "URL"
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/draft/batchget?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")

	result := new(ResultDraftBatchGet)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", BatchGetDraft(1, 10, 1, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultDraftBatchGet{
		TotalCount: 10,
		ItemCount:  1,
		Item: []*DraftDetail{
			{
				ArticleID: "ARTICLE_ID",
				Content: &DraftContent{
					NewsItem: []*DraftArticle{
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
						},
					},
				},
				UpdateTime: 1645685624,
			},
		},
	}, result)
}
