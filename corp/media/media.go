package media

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"os"
	"path/filepath"

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
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// UploadByURL 上传临时素材
func UploadByURL(mediaType MediaType, filename, url string, result *ResultUpload) wx.Action {
	return wx.NewPostAction(urls.CorpMediaUpload,
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
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// UploadByByte 上传临时素材
func UploadByByte(mediaType MediaType, filename string, fileByte []byte, result *ResultUpload) wx.Action {
	return wx.NewPostAction(urls.CorpMediaUpload,
		wx.WithQuery("type", string(mediaType)),
		wx.WithUpload(func() (wx.UploadForm, error) {
			return wx.NewUploadForm(
				wx.WithFormFile("media", filename, func(w io.Writer) error {
					f := bytes.NewBuffer(fileByte)
					if _, err := io.Copy(w, f); err != nil {
						return err
					}

					return nil
				}),
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
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// UploadImgByURL 上传图片
func UploadImgByURL(filename, url string, result *ResultUploadImg) wx.Action {
	return wx.NewPostAction(urls.CorpMediaUploadImg,
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
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
