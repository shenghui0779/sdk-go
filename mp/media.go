package mp

import (
	"encoding/json"
	"net/url"

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
	query := url.Values{}

	query.Set("type", string(mediaType))

	return wx.NewUploadAPI(MediaUploadURL, query, wx.NewUploadForm("media", filename, nil), func(resp []byte) error {
		return json.Unmarshal(resp, dest)
	})
}

// Media 临时素材
type Media struct {
	Buffer []byte
}

// GetMedia 获取客服消息内的临时素材
func GetMedia(dest *Media, mediaID string) wx.Action {
	query := url.Values{}

	query.Set("media_id", mediaID)

	return wx.NewGetAPI(MediaGetURL, query, func(resp []byte) error {
		dest.Buffer = make([]byte, len(resp))

		copy(dest.Buffer, resp)

		return nil
	})
}
