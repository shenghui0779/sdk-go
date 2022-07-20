package offia

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

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

// ResultMediaUpload 临时素材上传结果
type ResultMediaUpload struct {
	Type      string `json:"type"`
	MediaID   string `json:"media_id"`
	CreatedAt int64  `json:"created_at"`
}

// UploadMedia 素材管理 - 上传临时素材
func UploadMedia(mediaType MediaType, mediaPath string, result *ResultMediaUpload) wx.Action {
	_, filename := filepath.Split(mediaPath)

	return wx.NewPostAction(urls.OffiaMediaUpload,
		wx.WithQuery("type", string(mediaType)),
		wx.WithUpload(func() (wx.UploadForm, error) {
			path, err := filepath.Abs(filepath.Clean(mediaPath))

			if err != nil {
				return nil, err
			}

			return wx.NewUploadForm(
				wx.WithFormFile("media", filename, func(w io.Writer) error {
					f, err := os.Open(path)

					if err != nil {
						return err
					}

					defer f.Close()

					if _, err = io.Copy(w, f); err != nil {
						return err
					}

					return nil
				}),
			), nil
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// UploadMediaByURL 素材管理 - 上传临时素材
func UploadMediaByURL(mediaType MediaType, filename, url string, result *ResultMediaUpload) wx.Action {
	return wx.NewPostAction(urls.OffiaMediaUpload,
		wx.WithQuery("type", string(mediaType)),
		wx.WithUpload(func() (wx.UploadForm, error) {
			return wx.NewUploadForm(
				wx.WithFormFile("media", filename, func(w io.Writer) error {
					resp, err := wx.HTTPGet(context.Background(), url)

					if err != nil {
						return err
					}

					defer resp.Body.Close()

					if _, err = io.Copy(w, resp.Body); err != nil {
						return err
					}

					return nil
				}),
			), nil
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// ResultMaterialAdd 永久素材新增结果
type ResultMaterialAdd struct {
	MediaID string `json:"media_id"`
	URL     string `json:"url"`
}

// AddMaterial 素材管理 - 新增其他类型永久素材（支持图片、音频、缩略图）
func AddMaterial(mediaType MediaType, mediaPath string, result *ResultMaterialAdd) wx.Action {
	_, filename := filepath.Split(mediaPath)

	return wx.NewPostAction(urls.OffiaMaterialAdd,
		wx.WithQuery("type", string(mediaType)),
		wx.WithUpload(func() (wx.UploadForm, error) {
			path, err := filepath.Abs(filepath.Clean(mediaPath))

			if err != nil {
				return nil, err
			}

			return wx.NewUploadForm(
				wx.WithFormFile("media", filename, func(w io.Writer) error {
					f, err := os.Open(path)

					if err != nil {
						return err
					}

					defer f.Close()

					if _, err = io.Copy(w, f); err != nil {
						return err
					}

					return nil
				}),
			), nil
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// AddMaterialByURL 素材管理 - 新增其他类型永久素材（支持图片、音频、缩略图）
func AddMaterialByURL(mediaType MediaType, filename, url string, result *ResultMaterialAdd) wx.Action {
	return wx.NewPostAction(urls.OffiaMaterialAdd,
		wx.WithQuery("type", string(mediaType)),
		wx.WithUpload(func() (wx.UploadForm, error) {
			return wx.NewUploadForm(
				wx.WithFormFile("media", filename, func(w io.Writer) error {
					resp, err := wx.HTTPGet(context.Background(), url)

					if err != nil {
						return err
					}

					defer resp.Body.Close()

					if _, err = io.Copy(w, resp.Body); err != nil {
						return err
					}

					return nil
				}),
			), nil
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ParamsMaterialGet struct {
	MediaID string `json:"media_id"`
}

type ResultNewsMaterialGet struct {
	NewsItem []*MaterialNewsItem `json:"news_item"`
}

type MaterialNewsItem struct {
	Title            string `json:"title"`
	ThumbMediaID     string `json:"thumb_media_id"`
	ShowCoverPic     int    `json:"show_cover_pic"`
	Author           string `json:"author"`
	Digest           string `json:"digest"`
	Content          string `json:"content"`
	URL              string `json:"url"`
	ContentSourceURL string `json:"content_source_url"`
}

// GetNewsMaterial 素材管理 - 获取永久素材（图文）
func GetNewsMaterial(mediaID string, result *ResultNewsMaterialGet) wx.Action {
	params := &ParamsMaterialGet{
		MediaID: mediaID,
	}

	return wx.NewPostAction(urls.OffiaMaterialGet,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ResultVideoMaterialGet struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DownURL     string `json:"down_url"`
}

// GetVideoMaterial 素材管理 - 获取永久素材（视频消息）
func GetVideoMaterial(mediaID string, result *ResultVideoMaterialGet) wx.Action {
	params := &ParamsMaterialGet{
		MediaID: mediaID,
	}

	return wx.NewPostAction(urls.OffiaMaterialGet,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ResultOtherMaterialGet struct {
	Buffer []byte
}

// GetOtherMaterial 素材管理 - 获取永久素材（其他类型的素材消息，返回的直接为素材的内容）
func GetOtherMaterial(mediaID string, result *ResultOtherMaterialGet) wx.Action {
	params := &ParamsMaterialGet{
		MediaID: mediaID,
	}

	return wx.NewPostAction(urls.OffiaMaterialGet,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			result.Buffer = make([]byte, len(b))
			copy(result.Buffer, b)

			return nil
		}),
	)
}

// DeleteMaterial 素材管理 - 删除永久素材
func DeleteMaterial(mediaID string) wx.Action {
	return wx.NewPostAction(urls.OffiaMaterialDelete,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(wx.M{"media_id": mediaID})
		}),
	)
}

// UploadImg 素材管理 - 上传图文消息内的图片（不受公众号的素材库中图片数量的100000个的限制，图片仅支持jpg/png格式，大小必须在1MB以下）
func UploadImg(imgPath string, result *ResultMaterialAdd) wx.Action {
	_, filename := filepath.Split(imgPath)

	return wx.NewPostAction(urls.OffiaNewsImgUpload,
		wx.WithUpload(func() (wx.UploadForm, error) {
			path, err := filepath.Abs(filepath.Clean(imgPath))

			if err != nil {
				return nil, err
			}

			return wx.NewUploadForm(
				wx.WithFormFile("media", filename, func(w io.Writer) error {
					f, err := os.Open(path)

					if err != nil {
						return err
					}

					defer f.Close()

					if _, err = io.Copy(w, f); err != nil {
						return err
					}

					return nil
				}),
			), nil
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// UploadImgByURL 素材管理 - 上传图文消息内的图片（不受公众号的素材库中图片数量的100000个的限制，图片仅支持jpg/png格式，大小必须在1MB以下）
func UploadImgByURL(filename, url string, result *ResultMaterialAdd) wx.Action {
	return wx.NewPostAction(urls.OffiaNewsImgUpload,
		wx.WithUpload(func() (wx.UploadForm, error) {
			return wx.NewUploadForm(
				wx.WithFormFile("media", filename, func(w io.Writer) error {
					resp, err := wx.HTTPGet(context.Background(), url)

					if err != nil {
						return err
					}

					defer resp.Body.Close()

					if _, err = io.Copy(w, resp.Body); err != nil {
						return err
					}

					return nil
				}),
			), nil
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// UploadVideo 素材管理 - 上传视频永久素材
func UploadVideo(videoPath, title, description string, result *ResultMaterialAdd) wx.Action {
	_, filename := filepath.Split(videoPath)

	return wx.NewPostAction(urls.OffiaMaterialAdd,
		wx.WithQuery("type", string(MediaVideo)),
		wx.WithUpload(func() (wx.UploadForm, error) {
			path, err := filepath.Abs(filepath.Clean(videoPath))

			if err != nil {
				return nil, err
			}

			return wx.NewUploadForm(
				wx.WithFormFile("media", filename, func(w io.Writer) error {
					f, err := os.Open(path)

					if err != nil {
						return err
					}

					defer f.Close()

					if _, err = io.Copy(w, f); err != nil {
						return err
					}

					return nil
				}),
				wx.WithFormField("description", fmt.Sprintf(`{"title":"%s", "introduction":"%s"}`, title, description)),
			), nil
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// UploadVideoByURL 素材管理 - 上传视频永久素材
func UploadVideoByURL(filename, videoURL, title, description string, result *ResultMaterialAdd) wx.Action {
	return wx.NewPostAction(urls.OffiaMaterialAdd,
		wx.WithQuery("type", string(MediaVideo)),
		wx.WithUpload(func() (wx.UploadForm, error) {
			return wx.NewUploadForm(
				wx.WithFormFile("media", filename, func(w io.Writer) error {
					resp, err := wx.HTTPGet(context.Background(), videoURL)

					if err != nil {
						return err
					}

					defer resp.Body.Close()

					if _, err = io.Copy(w, resp.Body); err != nil {
						return err
					}

					return nil
				}),
				wx.WithFormField("description", fmt.Sprintf(`{"title":"%s", "introduction":"%s"}`, title, description)),
			), nil
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
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
	NeedOpenComment    int    `json:"need_open_comment,omitempty"`
	OnlyFansCanComment int    `json:"only_fans_can_comment,omitempty"`
}

type ParamsNewsAdd struct {
	Articles []*NewsArticle `json:"articles"`
}

// AddNews 素材管理 - 新增永久图文素材（公众号的素材库保存总数量有上限：图文消息素材、图片素材上限为100000，其他类型为1000）
func AddNews(articles []*NewsArticle, result *ResultMaterialAdd) wx.Action {
	params := &ParamsNewsAdd{
		Articles: articles,
	}

	return wx.NewPostAction(urls.OffiaNewsAdd,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ParamsNewsUpdate struct {
	MediaID  string       `json:"media_id"`
	Index    string       `json:"index"`
	Articles *NewsArticle `json:"articles"`
}

// UpdateNews 素材管理 - 编辑永久图文素材
func UpdateNews(mediaID, index string, article *NewsArticle) wx.Action {
	params := &ParamsNewsUpdate{
		MediaID:  mediaID,
		Index:    index,
		Articles: article,
	}

	return wx.NewPostAction(urls.OffiaNewsUpdate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type ResultMaterialCount struct {
	VoiceCount int `json:"voice_count"`
	VideoCount int `json:"video_count"`
	ImageCount int `json:"image_count"`
	NewsCount  int `json:"news_count"`
}

// GetMaterialCount 素材管理 - 获取素材总数
func GetMaterialCount(result *ResultMaterialCount) wx.Action {
	return wx.NewGetAction(urls.OffiaMaterialCountGet,
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ParamsMaterialList struct {
	Type   MediaType `json:"type"`
	Offset int       `json:"offset"`
	Count  int       `json:"count"`
}

type ResultMaterialList struct {
	TotalCount int                 `json:"total_count"`
	ItemCount  int                 `json:"item_count"`
	Item       []*MaterialListItem `json:"item"`
}

type MaterialListItem struct {
	MediaID    string `json:"media_id"`
	Name       string `json:"name"`
	UpdateTime int64  `json:"update_time"`
	URL        string `json:"url"`
}

// ListMatertial 素材管理 - 获取素材列表（其他类型：图片、语音、视频）
func ListMatertial(mediaType MediaType, offset, count int, result *ResultMaterialList) wx.Action {
	params := &ParamsMaterialList{
		Type:   mediaType,
		Offset: offset,
		Count:  count,
	}

	return wx.NewPostAction(urls.OffiaMaterialBatchGet,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ResultMaterialNewsList struct {
	TotalCount int                     `json:"total_count"`
	ItemCount  int                     `json:"item_count"`
	Item       []*MaterialNewsListItem `json:"item"`
}

type MaterialNewsListItem struct {
	MediaID    string                   `json:"media_id"`
	UpdateTime int64                    `json:"update_time"`
	Content    *MaterialNewsListContent `json:"content"`
}

type MaterialNewsListContent struct {
	NewsItem []*MaterialNewsItem `json:"news_item"`
}

// ListMaterialNews 素材管理 - 获取素材列表（永久图文消息）
func ListMaterialNews(offset, count int, result *ResultMaterialNewsList) wx.Action {
	params := &ParamsMaterialList{
		Type:   MediaType("news"),
		Offset: offset,
		Count:  count,
	}

	return wx.NewPostAction(urls.OffiaMaterialBatchGet,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}
