package mp

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/tidwall/gjson"

	"github.com/iiinsomnia/gochat/utils"
)

// UploadData 媒体上传数据
type UploadData struct {
	MediaType string     // 文件类型，合法值：image
	FormData  url.Values // form-data 中媒体文件标识，有filename、filelength、content-type等信息
}

// Media 媒体
type Media struct {
	mp      *WXMP
	options []utils.HTTPRequestOption
}

// Upload 上传媒体
func (m *Media) Upload(data *UploadData, accessToken string) (string, error) {
	m.options = append(m.options, utils.WithRequestHeader("Content-Type", "application/x-www-form-urlencoded"))

	resp, err := m.mp.Client.Post(fmt.Sprintf("%s?access_token=%s&type=%s", MediaUploadURL, accessToken, data.MediaType), []byte(data.FormData.Encode()), m.options...)

	if err != nil {
		return "", err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return "", errors.New(r.Get("errmsg").String())
	}

	return r.Get("media_id").String(), nil
}

// Get 获取媒体
func (m *Media) Get(mediaID, accessToken string) ([]byte, error) {
	resp, err := m.mp.Client.Get(fmt.Sprintf("%s?access_token=%s&media_id=%s", MediaGetURL, accessToken, mediaID), m.options...)

	if err != nil {
		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return nil, errors.New(r.Get("errmsg").String())
	}

	return resp, nil
}
