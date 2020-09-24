package mp

import (
	"fmt"
	"net/url"

	"github.com/shenghui0779/gochat/utils"
	"github.com/tidwall/gjson"
)

// UploadData 媒体上传数据
type UploadData struct {
	MediaType string     // 文件类型，合法值：image
	FormData  url.Values // form-data 中媒体文件标识，有filename、filelength、content-type等信息
}

// Media 媒体
type Media struct {
	mp      *WXMP
	options []utils.RequestOption
}

// Upload 上传媒体
func (m *Media) Upload(data *UploadData, accessToken string) (string, error) {
	b, err := m.mp.upload(fmt.Sprintf("%s?access_token=%s&type=%s", MediaUploadURL, accessToken, data.MediaType), []byte(data.FormData.Encode()), m.options...)

	if err != nil {
		return "", err
	}

	return gjson.GetBytes(b, "media_id").String(), nil
}

// Get 获取媒体
func (m *Media) Get(mediaID, accessToken string) ([]byte, error) {
	return m.mp.get(fmt.Sprintf("%s?access_token=%s&media_id=%s", MediaGetURL, accessToken, mediaID), m.options...)
}
