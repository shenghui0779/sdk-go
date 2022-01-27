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

// UploadMedia 上传临时素材
func UploadMedia(mediaType MediaType, mediaPath string, result *ResultMediaUpload) wx.Action {
	_, filename := filepath.Split(mediaPath)

	return wx.NewPostAction(urls.OffiaMediaUpload,
		wx.WithQuery("type", string(mediaType)),
		wx.WithUpload(func() (yiigo.UploadForm, error) {
			path, err := filepath.Abs(filepath.Clean(mediaPath))

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
func UploadMediaByURL(mediaType MediaType, filename, url string, result *ResultMediaUpload) wx.Action {
	return wx.NewPostAction(urls.OffiaMediaUpload,
		wx.WithQuery("type", string(mediaType)),
		wx.WithUpload(func() (yiigo.UploadForm, error) {
			resp, err := yiigo.HTTPGet(context.Background(), url)

			if err != nil {
				return nil, err
			}

			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)

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

// ResultMaterialAdd 永久素材新增结果
type ResultMaterialAdd struct {
	MediaID string `json:"media_id"`
	URL     string `json:"url"`
}

// AddMaterial 新增其他类型永久素材（支持图片、音频、缩略图）
func AddMaterial(mediaType MediaType, mediaPath string, result *ResultMaterialAdd) wx.Action {
	_, filename := filepath.Split(mediaPath)

	return wx.NewPostAction(urls.OffiaMaterialAdd,
		wx.WithQuery("type", string(mediaType)),
		wx.WithUpload(func() (yiigo.UploadForm, error) {
			path, err := filepath.Abs(filepath.Clean(mediaPath))

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

// AddMaterialByURL 新增其他类型永久素材（支持图片、音频、缩略图）
func AddMaterialByURL(mediaType MediaType, filename, url string, result *ResultMaterialAdd) wx.Action {
	return wx.NewPostAction(urls.OffiaMaterialAdd,
		wx.WithQuery("type", string(mediaType)),
		wx.WithUpload(func() (yiigo.UploadForm, error) {
			resp, err := yiigo.HTTPGet(context.Background(), url)

			if err != nil {
				return nil, err
			}

			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)

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

// GetNewsMaterial 获取永久素材 - 图文
func GetNewsMaterial(mediaID string, result *ResultNewsMaterialGet) wx.Action {
	params := &ParamsMaterialGet{
		MediaID: mediaID,
	}

	return wx.NewPostAction(urls.OffiaMaterialGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ResultVideoMaterialGet struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DownURL     string `json:"down_url"`
}

// GetVideoMaterial 获取永久素材 - 视频消息
func GetVideoMaterial(mediaID string, result *ResultVideoMaterialGet) wx.Action {
	params := &ParamsMaterialGet{
		MediaID: mediaID,
	}

	return wx.NewPostAction(urls.OffiaMaterialGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ResultOtherMaterialGet struct {
	Buffer []byte
}

// GetOtherMaterial 获取永久素材 - 其他类型的素材消息（返回的直接为素材的内容）
func GetOtherMaterial(mediaID string, result *ResultOtherMaterialGet) wx.Action {
	params := &ParamsMaterialGet{
		MediaID: mediaID,
	}

	return wx.NewPostAction(urls.OffiaMaterialGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			result.Buffer = make([]byte, len(resp))
			copy(result.Buffer, resp)

			return nil
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

// UploadImg 上传图文消息内的图片（不受公众号的素材库中图片数量的100000个的限制，图片仅支持jpg/png格式，大小必须在1MB以下）
func UploadImg(imgPath string, result *ResultMaterialAdd) wx.Action {
	_, filename := filepath.Split(imgPath)

	return wx.NewPostAction(urls.OffiaNewsImgUpload,
		wx.WithUpload(func() (yiigo.UploadForm, error) {
			path, err := filepath.Abs(filepath.Clean(imgPath))

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

// UploadImgByURL 上传图文消息内的图片（不受公众号的素材库中图片数量的100000个的限制，图片仅支持jpg/png格式，大小必须在1MB以下）
func UploadImgByURL(filename, url string, result *ResultMaterialAdd) wx.Action {
	return wx.NewPostAction(urls.OffiaNewsImgUpload,
		wx.WithUpload(func() (yiigo.UploadForm, error) {
			resp, err := yiigo.HTTPGet(context.Background(), url)

			if err != nil {
				return nil, err
			}

			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)

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
	NeedOpenComment    int    `json:"need_open_comment,omitempty"`
	OnlyFansCanComment int    `json:"only_fans_can_comment,omitempty"`
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

type ParamsNewsUpdate struct {
	MediaID  string       `json:"media_id"`
	Index    string       `json:"index"`
	Articles *NewsArticle `json:"articles"`
}

// UpdateNews 编辑图文素材
func UpdateNews(params *ParamsNewsUpdate) wx.Action {
	return wx.NewPostAction(urls.OffiaNewsUpdate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ResultMaterialCount struct {
	VoiceCount int `json:"voice_count"`
	VideoCount int `json:"video_count"`
	ImageCount int `json:"image_count"`
	NewsCount  int `json:"news_count"`
}

func GetMaterialCount(result *ResultMaterialCount) wx.Action {
	return wx.NewGetAction(urls.OffiaMaterialCountGet,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
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

func ListMatertial(params *ParamsMaterialList, result *ResultMaterialList) wx.Action {
	return wx.NewPostAction(urls.OffiaMaterialBatchGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
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

func ListMaterialNews(offset, count int, result *ResultMaterialNewsList) wx.Action {
	params := &ParamsMaterialList{
		Type:   MediaType("news"),
		Offset: offset,
		Count:  count,
	}

	return wx.NewPostAction(urls.OffiaMaterialBatchGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
