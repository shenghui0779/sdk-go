package externalcontact

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
	"github.com/shenghui0779/yiigo"
)

type MediaType string

const (
	MediaImage MediaType = "image"
	MediaLink  MediaType = "link"
	MediaMinip MediaType = "miniprogram"
	MediaVideo MediaType = "video"
	MediaFile  MediaType = "file"
)

type ResultAttachmentUpload struct {
	Type      MediaType `json:"type"`
	MediaID   string    `json:"media_id"`
	CreatedAt int64     `json:"created_at"`
}

type ParamsAttachmentUpload struct {
	MediaType      MediaType
	AttachmentType int
	Path           string
}

func UploadAttachment(params *ParamsAttachmentUpload, result *ResultAttachmentUpload) wx.Action {
	_, filename := filepath.Split(params.Path)

	return wx.NewPostAction(urls.OffiaMediaUpload,
		wx.WithQuery("type", string(params.MediaType)),
		wx.WithUpload(func() (yiigo.UploadForm, error) {
			path, err := filepath.Abs(filepath.Clean(params.Path))

			if err != nil {
				return nil, err
			}

			body, err := ioutil.ReadFile(path)

			if err != nil {
				return nil, err
			}

			return yiigo.NewUploadForm(
				yiigo.WithFileField("media", filename, body),
			), nil
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsAttachmentUploadByURL struct {
	MediaType      MediaType
	AttachmentType int
	Filename       string
	URL            string
}

// UploadAttachmentByURL 上传附件资源
func UploadAttachmentByURL(params *ParamsAttachmentUploadByURL, result *ResultAttachmentUpload) wx.Action {
	return wx.NewPostAction(urls.OffiaMediaUpload,
		wx.WithQuery("type", string(params.MediaType)),
		wx.WithUpload(func() (yiigo.UploadForm, error) {
			resp, err := yiigo.HTTPGet(context.Background(), params.URL)

			if err != nil {
				return nil, err
			}

			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				return nil, err
			}

			return yiigo.NewUploadForm(
				yiigo.WithFileField("media", params.Filename, body),
			), nil
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
