package offia

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type DraftArticle struct {
	Title              string `json:"title"`
	Author             string `json:"author,omitempty"`
	Digest             string `json:"digest,omitempty"`
	Content            string `json:"content"`
	ContentSourceURL   string `json:"content_source_url,omitempty"`
	ThumbMediaID       string `json:"thumb_media_id"`
	ShowCoverPic       int    `json:"show_cover_pic,omitempty"`
	NeedOpenComment    int    `json:"need_open_comment,omitempty"`
	OnlyFansCanComment int    `json:"only_fans_can_comment,omitempty"`
	URL                string `json:"url,omitempty"`
}

type ParamsDraftAdd struct {
	Articles []*DraftArticle `json:"articles"`
}

type ResultDraftAdd struct {
	MediaID string `json:"media_id"`
}

// AddDraft 草稿箱 - 新建草稿
func AddDraft(params *ParamsDraftAdd, result *ResultDraftAdd) wx.Action {
	return wx.NewPostAction(urls.OffiaDraftAdd,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ParamsDraftUpdate struct {
	MediaID  string        `json:"media_id"`
	Index    int           `json:"index"`
	Articles *DraftArticle `json:"articles"`
}

// UpdateDraft 草稿箱 - 修改草稿
func UpdateDraft(mediaID string, index int, article *DraftArticle) wx.Action {
	params := &ParamsDraftUpdate{
		MediaID:  mediaID,
		Index:    index,
		Articles: article,
	}

	return wx.NewPostAction(urls.OffiaDraftUpdate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type ParamsDraftGet struct {
	MediaID string `json:"media_id"`
}

type ResultDraftGet struct {
	NewsItem []*DraftArticle `json:"news_item"`
}

// GetDraft 草稿箱 - 获取草稿
func GetDraft(mediaID string, result *ResultDraftGet) wx.Action {
	params := &ParamsDraftGet{
		MediaID: mediaID,
	}

	return wx.NewPostAction(urls.OffiaDraftGet,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ParamsDraftDelete struct {
	MediaID string `json:"media_id"`
}

// DeleteDraft 草稿箱 - 删除草稿
func DeleteDraft(mediaID string) wx.Action {
	params := &ParamsDraftDelete{
		MediaID: mediaID,
	}

	return wx.NewPostAction(urls.OffiaDraftDelete,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type ResultDraftCount struct {
	TotalCount int `json:"total_count"`
}

// GetDraftCount 草稿箱 - 获取草稿总数
func GetDraftCount(result *ResultDraftCount) wx.Action {
	return wx.NewGetAction(urls.OffiaDraftCount,
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ParamsDraftBatchGet struct {
	Offset    int `json:"offset"`
	Count     int `json:"count"`
	NoContent int `json:"no_content,omitempty"`
}

type ResultDraftBatchGet struct {
	TotalCount int            `json:"total_count"`
	ItemCount  int            `json:"item_count"`
	Item       []*DraftDetail `json:"item"`
}

type DraftDetail struct {
	ArticleID  string        `json:"article_id"`
	Content    *DraftContent `json:"content"`
	UpdateTime int64         `json:"update_time"`
}

type DraftContent struct {
	NewsItem []*DraftArticle `json:"news_item"`
}

// BatchGetDraft 草稿箱 - 获取草稿列表
func BatchGetDraft(offset, count, nocontent int, result *ResultDraftBatchGet) wx.Action {
	params := &ParamsDraftBatchGet{
		Offset:    offset,
		Count:     count,
		NoContent: nocontent,
	}

	return wx.NewPostAction(urls.OffiaDraftBatchGet,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}
