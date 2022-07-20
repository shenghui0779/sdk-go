package offia

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsPublishSubmit struct {
	MediaID string `json:"media_id"`
}

type ResultPublishSubmit struct {
	PublishID string `json:"publish_id"`
}

// SubmitPublish 发布能力 - 发布图文
// 开发者需要先将图文素材以草稿的形式保存（见“草稿箱/新建草稿”，如需从已保存的草稿中选择，见“草稿箱/获取草稿列表”），选择要发布的草稿 media_id 进行发布
func SubmitPublish(mediaID string, result *ResultPublishSubmit) wx.Action {
	params := &ParamsPublishSubmit{
		MediaID: mediaID,
	}

	return wx.NewPostAction(urls.OffiaPublishSubmit,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ParamsPublishGet struct {
	PublishID string `json:"publish_id"`
}

type ResultPublishGet struct {
	PublishID     string           `json:"publish_id"`
	PublishStatus int              `json:"publish_status"`
	ArticleID     string           `json:"article_id"`
	ArticleDetail *PublishArticles `json:"article_detail"`
	FailIDX       []int            `json:"fail_idx"`
}

type PublishArticles struct {
	Count int            `json:"count"`
	Item  []*PublishItem `json:"item"`
}

type PublishItem struct {
	IDX        int    `json:"idx"`
	ArticleURL string `json:"article_url"`
}

// GetPublish 发布能力 - 发布状态轮询
// 开发者可以尝试通过下面的发布状态轮询接口获知发布情况。
func GetPublish(publishID string, result *ResultPublishGet) wx.Action {
	params := &ParamsPublishGet{
		PublishID: publishID,
	}

	return wx.NewPostAction(urls.OffiaPublishGet,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ParamsPublishDelete struct {
	ArticleID string `json:"article_id"`
	Index     int    `json:"index"`
}

// DeletePublish 发布能力 - 删除发布
func DeletePublish(articleID string, index int) wx.Action {
	params := &ParamsPublishDelete{
		ArticleID: articleID,
		Index:     index,
	}

	return wx.NewPostAction(urls.OffiaPublishDelete,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type ParamsPublishArticle struct {
	ArticleID string `json:"article_id"`
}

type ResultPublishArticle struct {
	NewsItem []*PublishArticle `json:"news_item"`
}

type PublishArticle struct {
	Title              string `json:"title"`
	Author             string `json:"author"`
	Digest             string `json:"digest"`
	Content            string `json:"content"`
	ContentSourceURL   string `json:"content_source_url"`
	ThumbMediaID       string `json:"thumb_media_id"`
	ShowCoverPic       int    `json:"show_cover_pic"`
	NeedOpenComment    int    `json:"need_open_comment"`
	OnlyFansCanComment int    `json:"only_fans_can_comment"`
	URL                string `json:"url"`
	IsDelete           bool   `json:"is_delete"`
}

// GetPublishArticle 发布能力 - 通过 article_id 获取已发布文章
func GetPublishArticle(articleID string, result *ResultPublishArticle) wx.Action {
	params := &ParamsPublishArticle{
		ArticleID: articleID,
	}

	return wx.NewPostAction(urls.OffiaPublishGetArticle,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ParamsPublishBatchGet struct {
	Offset    int `json:"offset"`
	Count     int `json:"count"`
	NoContent int `json:"no_content,omitempty"`
}

type ResultPublishBatchGet struct {
	TotalCount int              `json:"total_count"`
	ItemCount  int              `json:"item_count"`
	Item       []*PublishDetail `json:"item"`
}

type PublishDetail struct {
	ArticleID  string          `json:"article_id"`
	Content    *PublishContent `json:"content"`
	UpdateTime int64           `json:"update_time"`
}

type PublishContent struct {
	NewsItem []*PublishArticle `json:"news_item"`
}

// BatchGetPublish 发布能力 - 获取成功发布列表
func BatchGetPublish(offset, count, nocontent int, result *ResultPublishBatchGet) wx.Action {
	params := &ParamsPublishBatchGet{
		Offset:    offset,
		Count:     count,
		NoContent: nocontent,
	}

	return wx.NewPostAction(urls.OffiaPublishBatchGet,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}
