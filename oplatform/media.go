/*
@Time : 2021/8/23 10:32 上午
@Author : 21
@File : media
@Software: GoLand
*/
package oplatform

import (
	"encoding/json"
	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"path/filepath"
)

//素材操作

// 新增临时素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/New_temporary_materials.html
type MediaUpload struct {
	// 参数token
	AccessToken string `json:"access_token"`
	Type        string `json:"type"`
	// 返回参数
	MaterialAddResult *MaterialAddResult
}

// 永久图文素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html
type MediaUploadImg struct {
	// 参数token
	AccessToken string `json:"access_token"`

	Type string `json:"type"`
	// 返回参数
	MaterialAddResult *MaterialAddResult
}
type MaterialAddResult struct {
	MediaID   string `json:"media_id"`
	URL       string `json:"url"`
	Type      string `json:"type"`
	CreatedAt int    `json:"created_at"`
}

// 永久图文素材
func MaterialAddNewsImage(dest *MediaUploadImg, path string) wx.Action {
	_, filename := filepath.Split(path)

	return wx.NewUploadAction(urls.OaAddMaterial,
		wx.WithQuery("type", string(dest.Type)),
		wx.WithQuery("access_token", string(dest.AccessToken)),
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
			dest.MaterialAddResult = &MaterialAddResult{
				MediaID: "",
				URL:     "",
			}
			dest.MaterialAddResult.URL = gjson.GetBytes(resp, "url").String()

			return nil
		}),
	)
}

// UploadMedia 上传临时素材
func UploadMedia(dest *MediaUpload, path string) wx.Action {
	_, filename := filepath.Split(path)
	return wx.NewUploadAction(urls.OaMediaUpload,
		wx.WithQuery("type", string(dest.Type)),
		wx.WithQuery("access_token", string(dest.AccessToken)),
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
			dest.MaterialAddResult = &MaterialAddResult{}
			return json.Unmarshal(resp, dest.MaterialAddResult)
		}),
	)
}
