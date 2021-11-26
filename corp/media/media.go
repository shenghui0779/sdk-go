package media

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/shenghui0779/yiigo"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type MediaType string

const (
	MediaImage MediaType = "image"
	MediaVoice MediaType = "voice"
	MediaVideo MediaType = "video"
	MediaFile  MediaType = "file"
)

type ParamsUpload struct {
	Type MediaType
	Path string
}

type ResultUpload struct {
	Type      MediaType `json:"type"`
	MediaID   string    `json:"media_id"`
	CreatedAt string    `json:"created_at"`
}

func Upload(params *ParamsUpload, result *ResultUpload) wx.Action {
	_, filename := filepath.Split(params.Path)

	return wx.NewPostAction(urls.CorpMediaUpload,
		wx.WithQuery("type", string(params.Type)),
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

type ParamsUploadByURL struct {
	Type     MediaType
	MediaURL string
	Filename string
}

// UploadByURL 上传临时素材
func UploadByURL(params *ParamsUploadByURL, result *ResultUpload) wx.Action {
	return wx.NewPostAction(urls.OffiaMediaUpload,
		wx.WithQuery("type", string(params.Type)),
		wx.WithUpload(func() (yiigo.UploadForm, error) {
			resp, err := yiigo.HTTPGet(context.Background(), params.MediaURL)

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

type ParamsUploadimg struct {
	Path string
}

type ResultUploadimg struct {
	URL string `json:"url"`
}

// Uploadimg 上传图片
func Uploadimg(params *ParamsUploadimg, result *ResultUploadimg) wx.Action {
	_, filename := filepath.Split(params.Path)

	return wx.NewPostAction(urls.CorpMediaUploadimg,
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

type ParamsUploadimgByURL struct {
	ImageURL string
	Filename string
}

// UploadimgByURL 上传图片
func UploadimgByURL(params *ParamsUploadimgByURL, result *ResultUploadimg) wx.Action {
	return wx.NewPostAction(urls.CorpMediaUploadimg,
		wx.WithUpload(func() (yiigo.UploadForm, error) {
			resp, err := yiigo.HTTPGet(context.Background(), params.ImageURL)

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
