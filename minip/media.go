package minip

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/shenghui0779/yiigo"

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

type ParamsMediaUpload struct {
	MediaType MediaType `json:"media_type"`
	Path      string    `json:"path"`
}

// UploadMedia 上传临时素材到微信服务器
func UploadMedia(params *ParamsMediaUpload, result *ResultMediaUpload) wx.Action {
	_, filename := filepath.Split(params.Path)

	return wx.NewUploadAction(urls.MinipMediaUpload,
		wx.WithQuery("type", string(params.MediaType)),
		wx.WithUploadField(&wx.UploadField{
			FileField: "media",
			Filename:  filename,
		}),
		wx.WithBody(func() ([]byte, error) {
			path, err := filepath.Abs(filepath.Clean(params.Path))

			if err != nil {
				return nil, err
			}

			return ioutil.ReadFile(path)
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

// UploadMediaByURL 上传临时素材到微信服务器
func UploadMediaByURL(params *ParamsMediaUploadByURL, result *ResultMediaUpload) wx.Action {
	return wx.NewUploadAction(urls.MinipMediaUpload,
		wx.WithQuery("type", string(params.MediaType)),
		wx.WithUploadField(&wx.UploadField{
			FileField: "media",
			Filename:  params.Filename,
		}),
		wx.WithBody(func() ([]byte, error) {
			resp, err := yiigo.HTTPGet(context.TODO(), params.URL)

			if err != nil {
				return nil, err
			}

			defer resp.Body.Close()

			return ioutil.ReadAll(resp.Body)
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

// GetMedia 获取客服消息内的临时素材
func GetMedia(mediaID string, media *Media) wx.Action {
	return wx.NewGetAction(urls.MinipMediaGet,
		wx.WithQuery("media_id", mediaID),
		wx.WithDecode(func(resp []byte) error {
			media.Buffer = make([]byte, len(resp))

			copy(media.Buffer, resp)

			return nil
		}),
	)
}
