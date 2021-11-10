package offia

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/shenghui0779/yiigo"
	"github.com/tidwall/gjson"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

// MediaType 素材类型
type MediaType string

// 微信支持的素材类型
const (
	MediaImage MediaType = "image" // 图片
	MediaVoice MediaType = "voice" // 音频
	MediaVideo MediaType = "video" // 视频
	MediaThumb MediaType = "thumb" // 缩略图
)

// MediaUploadResult 临时素材上传结果
type MediaUploadResult struct {
	Type      string `json:"type"`
	MediaID   string `json:"media_id"`
	CreatedAt int64  `json:"created_at"`
}

// UploadMedia 上传临时素材
func UploadMedia(dest *MediaUploadResult, mediaType MediaType, path string) wx.Action {
	_, filename := filepath.Split(path)

	return wx.NewUploadAction(urls.OffiaMediaUpload,
		wx.WithQuery("type", string(mediaType)),
		wx.WithUploadField(&wx.UploadField{
			FileField: "media",
			Filename:  filename,
		}),
		wx.WithBody(func() ([]byte, error) {
			path, err := filepath.Abs(filepath.Clean(path))

			if err != nil {
				return nil, err
			}

			return ioutil.ReadFile(path)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// UploadMediaByURL 上传临时素材
func UploadMediaByURL(dest *MediaUploadResult, mediaType MediaType, filename, resourceURL string) wx.Action {
	return wx.NewUploadAction(urls.OffiaMediaUpload,
		wx.WithQuery("type", string(mediaType)),
		wx.WithUploadField(&wx.UploadField{
			FileField: "media",
			Filename:  filename,
		}),
		wx.WithBody(func() ([]byte, error) {
			resp, err := yiigo.HTTPGet(context.Background(), resourceURL)

			if err != nil {
				return nil, err
			}

			defer resp.Body.Close()

			return ioutil.ReadAll(resp.Body)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// MaterialAddResult 永久素材新增结果
type MaterialAddResult struct {
	MediaID string `json:"media_id"`
	URL     string `json:"url"`
}

// NewsArticle 文章素材
type NewsArticle struct {
	Title              string `json:"title"`
	ThumbMediaID       string `json:"thumb_media_id"`
	Author             string `json:"author"`
	Digest             string `json:"digest"`
	ShowCoverPic       int    `json:"show_cover_pic"`
	Content            string `json:"content"`
	ContentSourceURL   string `json:"content_source_url"`
	NeedOpenComment    int    `json:"need_open_comment"`
	OnlyFansCanComment int    `json:"only_fans_can_comment"`
}

// AddNews 新增永久图文素材（公众号的素材库保存总数量有上限：图文消息素材、图片素材上限为100000，其他类型为1000）
func AddNews(dest *MaterialAddResult, articles ...*NewsArticle) wx.Action {
	return wx.NewPostAction(urls.OffiaNewsAdd,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(yiigo.X{"articles": articles})
		}),
		wx.WithDecode(func(resp []byte) error {
			dest.MediaID = gjson.GetBytes(resp, "media_id").String()

			return nil
		}),
	)
}

type UpdateArticles struct {
	MediaID  string   `json:"media_id"`
	Index    string   `json:"index"`
	Articles Articles `json:"articles"`
}

type Articles struct {
	Title            string `json:"title"`
	ThumbMediaID     string `json:"thumb_media_id"`
	Author           string `json:"author"`
	Digest           string `json:"digest"`
	ShowCoverPic     int    `json:"show_cover_pic"`
	Content          string `json:"content"`
	ContentSourceURL string `json:"content_source_url"`
}

// UpdateNews 编辑图文素材
func UpdateNews(articles *UpdateArticles) wx.Action {
	return wx.NewPostAction(urls.OffiaNewUpdate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(&articles)
		}),
	)
}

//GetArticle 永久图文素材
type GetArticle struct {
	NewsItem []*NewsItem `json:"news_item"`
}
type NewsItem struct {
	Title            string `json:"title"`
	ThumbMediaID     string `json:"thumb_media_id"`
	ShowCoverPic     int    `json:"show_cover_pic"`
	Author           string `json:"author"`
	Digest           string `json:"digest"`
	Content          string `json:"content"`
	URL              string `json:"url"`
	ContentSourceURL string `json:"content_source_url"`
}

// GetNews 获取图文素材信息
func GetNews(dest *GetArticle, mediaId string) wx.Action {
	return wx.NewPostAction(urls.OffiaMaterialGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(yiigo.X{"media_id": mediaId})
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, &dest)
		}),
	)
}

// UploadNewsImage 上传图文消息内的图片（不受公众号的素材库中图片数量的100000个的限制，图片仅支持jpg/png格式，大小必须在1MB以下）
func UploadNewsImage(dest *MaterialAddResult, path string) wx.Action {
	_, filename := filepath.Split(path)

	return wx.NewUploadAction(urls.OffiaNewsImageUpload,
		wx.WithUploadField(&wx.UploadField{
			FileField: "media",
			Filename:  filename,
		}),
		wx.WithBody(func() ([]byte, error) {
			path, err := filepath.Abs(filepath.Clean(path))

			if err != nil {
				return nil, err
			}

			return ioutil.ReadFile(path)
		}),
		wx.WithDecode(func(resp []byte) error {
			dest.URL = gjson.GetBytes(resp, "url").String()

			return nil
		}),
	)
}

// UploadNewsImageByURL 上传图文消息内的图片（不受公众号的素材库中图片数量的100000个的限制，图片仅支持jpg/png格式，大小必须在1MB以下）
func UploadNewsImageByURL(dest *MaterialAddResult, filename, resourceURL string) wx.Action {
	return wx.NewUploadAction(urls.OffiaNewsImageUpload,
		wx.WithUploadField(&wx.UploadField{
			FileField: "media",
			Filename:  filename,
		}),
		wx.WithBody(func() ([]byte, error) {
			resp, err := yiigo.HTTPGet(context.Background(), resourceURL)

			if err != nil {
				return nil, err
			}

			defer resp.Body.Close()

			return ioutil.ReadAll(resp.Body)
		}),
		wx.WithDecode(func(resp []byte) error {
			dest.URL = gjson.GetBytes(resp, "url").String()

			return nil
		}),
	)
}

// AddMaterial 新增其他类型永久素材（支持图片、音频、缩略图）
func AddMaterial(dest *MaterialAddResult, mediaType MediaType, path string) wx.Action {
	_, filename := filepath.Split(path)

	return wx.NewUploadAction(urls.OffiaMaterialAdd,
		wx.WithQuery("type", string(mediaType)),
		wx.WithUploadField(&wx.UploadField{
			FileField: "media",
			Filename:  filename,
		}),
		wx.WithBody(func() ([]byte, error) {
			path, err := filepath.Abs(filepath.Clean(path))

			if err != nil {
				return nil, err
			}

			return ioutil.ReadFile(path)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// AddMaterialByURL 新增其他类型永久素材（支持图片、音频、缩略图）
func AddMaterialByURL(dest *MaterialAddResult, mediaType MediaType, filename, resourceURL string) wx.Action {
	return wx.NewUploadAction(urls.OffiaMaterialAdd,
		wx.WithQuery("type", string(mediaType)),
		wx.WithUploadField(&wx.UploadField{
			FileField: "media",
			Filename:  filename,
		}),
		wx.WithBody(func() ([]byte, error) {
			resp, err := yiigo.HTTPGet(context.Background(), resourceURL)

			if err != nil {
				return nil, err
			}

			defer resp.Body.Close()

			return ioutil.ReadAll(resp.Body)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// UploadVideo 上传视频永久素材
func UploadVideo(dest *MaterialAddResult, path, title, introduction string) wx.Action {
	_, filename := filepath.Split(path)

	return wx.NewUploadAction(urls.OffiaMaterialAdd,
		wx.WithQuery("type", string(MediaVideo)),
		wx.WithUploadField(&wx.UploadField{
			FileField: "media",
			Filename:  filename,
			MetaField: "description",
			Metadata:  fmt.Sprintf(`{"title":"%s", "introduction":"%s"}`, title, introduction),
		}),
		wx.WithBody(func() ([]byte, error) {
			path, err := filepath.Abs(filepath.Clean(path))

			if err != nil {
				return nil, err
			}

			return ioutil.ReadFile(path)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// UploadVideoByURL 上传视频永久素材
func UploadVideoByURL(dest *MaterialAddResult, filename, title, introduction, resourceURL string) wx.Action {
	return wx.NewUploadAction(urls.OffiaMaterialAdd,
		wx.WithQuery("type", string(MediaVideo)),
		wx.WithUploadField(&wx.UploadField{
			FileField: "media",
			Filename:  filename,
			MetaField: "description",
			Metadata:  fmt.Sprintf(`{"title":"%s", "introduction":"%s"}`, title, introduction),
		}),
		wx.WithBody(func() ([]byte, error) {
			resp, err := yiigo.HTTPGet(context.Background(), resourceURL)

			if err != nil {
				return nil, err
			}

			defer resp.Body.Close()

			return ioutil.ReadAll(resp.Body)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// DeleteMaterial 删除永久素材
func DeleteMaterial(mediaID string) wx.Action {
	return wx.NewPostAction(urls.OffiaMaterialDelete,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(yiigo.X{"media_id": mediaID})
		}),
	)
}
