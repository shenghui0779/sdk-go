package mp

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/shenghui0779/yiigo"

	"github.com/shenghui0779/gochat/wx"
)

// MediaType 素材类型
type MediaType string

// 微信支持的素材类型
const MediaImage MediaType = "image" // 图片

// MediaUploadResult  临时素材上传信息
type MediaUploadResult struct {
	Type      string `json:"type"`
	MediaID   string `json:"media_id"`
	CreatedAt int64  `json:"created_at"`
}

// UploadMedia 上传临时素材到微信服务器
func UploadMedia(dest *MediaUploadResult, mediaType MediaType, path string) wx.Action {
	_, filename := filepath.Split(path)

	return wx.NewAction(MediaUploadURL,
		wx.WithMethod(wx.MethodUpload),
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

// UploadMediaByURL 上传临时素材到微信服务器
func UploadMediaByURL(dest *MediaUploadResult, mediaType MediaType, filename, resourceURL string) wx.Action {
	return wx.NewAction(MediaUploadURL,
		wx.WithMethod(wx.MethodUpload),
		wx.WithQuery("type", string(mediaType)),
		wx.WithUploadField(&wx.UploadField{
			FileField: "media",
			Filename:  filename,
		}),
		wx.WithBody(func() ([]byte, error) {
			resp, err := yiigo.HTTPGet(context.TODO(), resourceURL)
			// resp, err := http.Get(resourceURL)

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

// Media 临时素材
type Media struct {
	Buffer []byte
}

// GetMedia 获取客服消息内的临时素材
func GetMedia(dest *Media, mediaID string) wx.Action {
	return wx.NewAction(MediaGetURL,
		wx.WithMethod(wx.MethodGet),
		wx.WithQuery("media_id", mediaID),
		wx.WithDecode(func(resp []byte) error {
			dest.Buffer = make([]byte, len(resp))

			copy(dest.Buffer, resp)

			return nil
		}),
	)
}
