package mp

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/wx"
)

// MediaType 素材类型
type MediaType string

// 微信支持的素材类型
var MediaImage MediaType = "image" // 图片

// MediaUploadResult  临时素材上传信息
type MediaUploadResult struct {
	Type      string `json:"type"`
	MediaID   string `json:"media_id"`
	CreatedAt int64  `json:"created_at"`
}

// UploadMedia 上传临时素材到微信服务器
func UploadMedia(dest *MediaUploadResult, mediaType MediaType, filename string) wx.Action {
	return wx.NewAction(MediaUploadURL,
		wx.WithMethod(wx.MethodUpload),
		wx.WithQuery("type", string(mediaType)),
		wx.WithUploadForm("media", filename, nil),
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
