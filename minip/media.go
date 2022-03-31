package minip

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

// MediaType 素材类型
type MediaType string

// 微信支持的素材类型
const MediaImage MediaType = "image" // 图片

// ResultMediaUpload  临时素材上传结果
type ResultMediaUpload struct {
	Type      MediaType `json:"type"`
	MediaID   string    `json:"media_id"`
	CreatedAt int64     `json:"created_at"`
}

// UploadTempMedia 客服消息 - 上传临时素材到微信服务器
func UploadTempMedia(mediaType MediaType, mediaPath string, result *ResultMediaUpload) wx.Action {
	_, filename := filepath.Split(mediaPath)

	return wx.NewPostAction(urls.MinipMediaUpload,
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

// UploadTempMediaByURL 客服消息 - 上传临时素材到微信服务器
func UploadTempMediaByURL(mediaType MediaType, filename, url string, result *ResultMediaUpload) wx.Action {
	return wx.NewPostAction(urls.MinipMediaUpload,
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

// Media 临时素材
type Media struct {
	Buffer []byte
}

// GetTempMedia 客服消息 - 获取客服消息内的临时素材
func GetTempMedia(mediaID string, media *Media) wx.Action {
	return wx.NewGetAction(urls.MinipMediaGet,
		wx.WithQuery("media_id", mediaID),
		wx.WithDecode(func(resp []byte) error {
			media.Buffer = make([]byte, len(resp))

			copy(media.Buffer, resp)

			return nil
		}),
	)
}
