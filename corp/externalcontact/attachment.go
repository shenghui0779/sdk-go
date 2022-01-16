package externalcontact

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strconv"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
	"github.com/shenghui0779/yiigo"
)

type MediaType string

const (
	MediaImage MediaType = "image"
	MediaVideo MediaType = "video"
	MediaFile  MediaType = "file"
)

type AttachmentType int

const (
	AttachmentMoment       AttachmentType = 1 // 朋友圈
	AttachmentProductAlbum AttachmentType = 2 // 商品相册
)

type ResultAttachmentUpload struct {
	Type      MediaType `json:"type"`
	MediaID   string    `json:"media_id"`
	CreatedAt int64     `json:"created_at"`
}

type ParamsAttachmentUpload struct {
	MediaType      MediaType      `json:"media_type"`
	AttachmentType AttachmentType `json:"attachment_type"`
	Path           string         `json:"path"`
}

func UploadAttachment(params *ParamsAttachmentUpload, result *ResultAttachmentUpload) wx.Action {
	_, filename := filepath.Split(params.Path)

	return wx.NewPostAction(urls.CorpExternalContactUploadAttachment,
		wx.WithQuery("media_type", string(params.MediaType)),
		wx.WithQuery("attachment_type", strconv.Itoa(int(params.AttachmentType))),
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
	MediaType      MediaType      `json:"media_type"`
	AttachmentType AttachmentType `json:"attachment_type"`
	Filename       string         `json:"filename"`
	URL            string         `json:"url"`
}

// UploadAttachmentByURL 上传附件资源
func UploadAttachmentByURL(params *ParamsAttachmentUploadByURL, result *ResultAttachmentUpload) wx.Action {
	return wx.NewPostAction(urls.OffiaMediaUpload,
		wx.WithQuery("media_type", string(params.MediaType)),
		wx.WithQuery("attachment_type", strconv.Itoa(int(params.AttachmentType))),
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
