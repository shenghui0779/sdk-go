package mp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/shenghui0779/gochat/internal"
)

// MediaType 素材类型
type MediaType string

// 微信支持的素材类型
var MediaImage MediaType = "image" // 图片

// MediaUploadInfo 临时素材上传信息
type MediaUploadInfo struct {
	Type      string `json:"type"`
	MediaID   string `json:"media_id"`
	CreatedAt int64  `json:"created_at"`
}

// UploadMedia 上传临时素材到微信服务器
func UploadMedia(mediaType MediaType, filename string, dest *MediaUploadInfo) internal.Action {
	return &WechatAPI{
		body: internal.NewUploadBody("media", filename, func() ([]byte, error) {
			return ioutil.ReadFile(filename)
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("UPLOAD|%s?access_token=%s&type=%s", MediaUploadURL, accessToken, mediaType)
		},
		decode: func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		},
	}
}

// Media 临时素材
type Media struct {
	Buffer []byte
}

// GetMedia 获取客服消息内的临时素材
func GetMedia(mediaID string, dest *Media) internal.Action {
	return &WechatAPI{
		url: func(accessToken string) string {
			return fmt.Sprintf("GET|%s?access_token=%s&media_id=%s", MediaGetURL, accessToken, mediaID)
		},
		decode: func(resp []byte) error {
			dest.Buffer = make([]byte, len(resp))

			copy(dest.Buffer, resp)

			return nil
		},
	}
}
