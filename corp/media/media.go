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

type ResultUpload struct {
	Type      MediaType `json:"type"`
	MediaID   string    `json:"media_id"`
	CreatedAt string    `json:"created_at"`
}

// Upload 上传临时素材
func Upload(mediaType MediaType, mediaPath string, result *ResultUpload) wx.Action {
	_, filename := filepath.Split(mediaPath)

	return wx.NewPostAction(urls.CorpMediaUpload,
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

// UploadByURL 上传临时素材
func UploadByURL(mediaType MediaType, filename, url string, result *ResultUpload) wx.Action {
	return wx.NewPostAction(urls.CorpMediaUpload,
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

type ResultUploadImg struct {
	URL string `json:"url"`
}

// UploadImg 上传图片
func UploadImg(imgPath string, result *ResultUploadImg) wx.Action {
	_, filename := filepath.Split(imgPath)

	return wx.NewPostAction(urls.CorpMediaUploadImg,
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

// UploadImgByURL 上传图片
func UploadImgByURL(filename, url string, result *ResultUploadImg) wx.Action {
	return wx.NewPostAction(urls.CorpMediaUploadImg,
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
