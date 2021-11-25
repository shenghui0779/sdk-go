package offia

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/shenghui0779/yiigo"

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

type ParamsMediaUpload struct {
	MediaType MediaType `json:"media_type"`
	Path      string    `json:"path"`
}

// UploadMedia 上传临时素材
func UploadMedia(params *ParamsMediaUpload, result *ResultMediaUpload) wx.Action {
	_, filename := filepath.Split(params.Path)

	return wx.NewPostAction(urls.OffiaMediaUpload,
		wx.WithQuery("type", string(params.MediaType)),
		wx.WithUpload(func() (yiigo.UploadForm, error) {
			path, err := filepath.Abs(filepath.Clean(params.Path))

			if err != nil {
				return nil, err
			}

			body, err := ioutil.ReadFile(path)

			if err != nil {
				return nil, err
			}

			return yiigo.NewUploadForm(
				yiigo.WithFileField("media", filename, body),
			), nil
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsMediaUploadByURL struct {
	MediaType MediaType
	Filename  string
	URL       string
}

// UploadMediaByURL 上传临时素材
func UploadMediaByURL(params *ParamsMediaUploadByURL, result *ResultMediaUpload) wx.Action {
	return wx.NewPostAction(urls.OffiaMediaUpload,
		wx.WithQuery("type", string(params.MediaType)),
		wx.WithUpload(func() (yiigo.UploadForm, error) {
			resp, err := yiigo.HTTPGet(context.Background(), params.URL)

			if err != nil {
				return nil, err
			}

			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				return nil, err
			}

			return yiigo.NewUploadForm(
				yiigo.WithFileField("media", params.Filename, body),
			), nil
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// ResultMaterialAdd 永久素材新增结果
type ResultMaterialAdd struct {
	MediaID string `json:"media_id"`
	URL     string `json:"url"`
}

type ParamsMaterialAdd struct {
	MediaType MediaType `json:"media_type"`
	Path      string    `json:"path"`
}

// AddMaterial 新增其他类型永久素材（支持图片、音频、缩略图）
func AddMaterial(params *ParamsMaterialAdd, result *ResultMaterialAdd) wx.Action {
	_, filename := filepath.Split(params.Path)

	return wx.NewPostAction(urls.OffiaMaterialAdd,
		wx.WithQuery("type", string(params.MediaType)),
		wx.WithUpload(func() (yiigo.UploadForm, error) {
			path, err := filepath.Abs(filepath.Clean(params.Path))

			if err != nil {
				return nil, err
			}

			body, err := ioutil.ReadFile(path)

			if err != nil {
				return nil, err
			}

			return yiigo.NewUploadForm(
				yiigo.WithFileField("media", filename, body),
			), nil
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsMaterialAddByURL struct {
	MediaType MediaType `json:"media_type"`
	Filename  string    `json:"filename"`
	URL       string    `json:"url"`
}

// AddMaterialByURL 新增其他类型永久素材（支持图片、音频、缩略图）
func AddMaterialByURL(params *ParamsMaterialAddByURL, result *ResultMaterialAdd) wx.Action {
	return wx.NewPostAction(urls.OffiaMaterialAdd,
		wx.WithQuery("type", string(params.MediaType)),
		wx.WithUpload(func() (yiigo.UploadForm, error) {
			resp, err := yiigo.HTTPGet(context.Background(), params.URL)

			if err != nil {
				return nil, err
			}

			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				return nil, err
			}

			return yiigo.NewUploadForm(
				yiigo.WithFileField("media", params.Filename, body),
			), nil
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
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

// UploadImage 上传图文消息内的图片（不受公众号的素材库中图片数量的100000个的限制，图片仅支持jpg/png格式，大小必须在1MB以下）
func UploadImage(path string, result *ResultMaterialAdd) wx.Action {
	_, filename := filepath.Split(path)

	return wx.NewPostAction(urls.OffiaNewsImageUpload,
		wx.WithUpload(func() (yiigo.UploadForm, error) {
			path, err := filepath.Abs(filepath.Clean(path))

			if err != nil {
				return nil, err
			}

			body, err := ioutil.ReadFile(path)

			if err != nil {
				return nil, err
			}

			return yiigo.NewUploadForm(
				yiigo.WithFileField("media", filename, body),
			), nil
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsImageUploadByURL struct {
	Filename string `json:"filename"`
	URL      string `json:"url"`
}

// UploadImageByURL 上传图文消息内的图片（不受公众号的素材库中图片数量的100000个的限制，图片仅支持jpg/png格式，大小必须在1MB以下）
func UploadImageByURL(params *ParamsImageUploadByURL, result *ResultMaterialAdd) wx.Action {
	return wx.NewPostAction(urls.OffiaNewsImageUpload,
		wx.WithUpload(func() (yiigo.UploadForm, error) {
			resp, err := yiigo.HTTPGet(context.Background(), params.URL)

			if err != nil {
				return nil, err
			}

			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				return nil, err
			}

			return yiigo.NewUploadForm(
				yiigo.WithFileField("media", params.Filename, body),
			), nil
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsVideoUpload struct {
	Path        string `json:"path"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// UploadVideo 上传视频永久素材
func UploadVideo(params *ParamsVideoUpload, result *ResultMaterialAdd) wx.Action {
	_, filename := filepath.Split(params.Path)

	return wx.NewPostAction(urls.OffiaMaterialAdd,
		wx.WithQuery("type", string(MediaVideo)),
		wx.WithUpload(func() (yiigo.UploadForm, error) {
			path, err := filepath.Abs(filepath.Clean(params.Path))

			if err != nil {
				return nil, err
			}

			body, err := ioutil.ReadFile(path)

			if err != nil {
				return nil, err
			}

			return yiigo.NewUploadForm(
				yiigo.WithFileField("media", filename, body),
				yiigo.WithFormField("description", fmt.Sprintf(`{"title":"%s", "introduction":"%s"}`, params.Title, params.Description)),
			), nil
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsVideoUploadByURL struct {
	Filename    string `json:"filename"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// UploadVideoByURL 上传视频永久素材
func UploadVideoByURL(params *ParamsVideoUploadByURL, result *ResultMaterialAdd) wx.Action {
	return wx.NewPostAction(urls.OffiaMaterialAdd,
		wx.WithQuery("type", string(MediaVideo)),
		wx.WithUpload(func() (yiigo.UploadForm, error) {
			resp, err := yiigo.HTTPGet(context.Background(), params.URL)

			if err != nil {
				return nil, err
			}

			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				return nil, err
			}

			return yiigo.NewUploadForm(
				yiigo.WithFileField("media", params.Filename, body),
				yiigo.WithFormField("description", fmt.Sprintf(`{"title":"%s", "introduction":"%s"}`, params.Title, params.Description)),
			), nil
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
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
	NeedOpenComment    int    `json:"need_open_comment"`
	OnlyFansCanComment int    `json:"only_fans_can_comment"`
}

type ParamsNewsAdd struct {
	Articles []*NewsArticle `json:"articles"`
}

// AddNews 新增永久图文素材（公众号的素材库保存总数量有上限：图文消息素材、图片素材上限为100000，其他类型为1000）
func AddNews(params *ParamsNewsAdd, result *ResultMaterialAdd) wx.Action {
	return wx.NewPostAction(urls.OffiaNewsAdd,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
