package externalcontact

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strconv"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
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

// UploadAttachment 上传附件资源
func UploadAttachment(mediaType MediaType, attachmentType AttachmentType, attachmentPath string, result *ResultAttachmentUpload) wx.Action {
	_, filename := filepath.Split(attachmentPath)

	return wx.NewPostAction(urls.CorpExternalContactUploadAttachment,
		wx.WithQuery("media_type", string(mediaType)),
		wx.WithQuery("attachment_type", strconv.Itoa(int(attachmentType))),
		wx.WithUpload(func() (wx.UploadForm, error) {
			path, err := filepath.Abs(filepath.Clean(attachmentPath))

			if err != nil {
				return nil, err
			}

			body, err := ioutil.ReadFile(path)

			if err != nil {
				return nil, err
			}

			return wx.NewUploadForm(
				wx.WithFileField("media", filename, body),
			), nil
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// UploadAttachmentByURL 上传附件资源
func UploadAttachmentByURL(mediaType MediaType, attachmentType AttachmentType, filename, url string, result *ResultAttachmentUpload) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactUploadAttachment,
		wx.WithQuery("media_type", string(mediaType)),
		wx.WithQuery("attachment_type", strconv.Itoa(int(attachmentType))),
		wx.WithUpload(func() (wx.UploadForm, error) {
			resp, err := wx.HTTPGet(context.Background(), url)

			if err != nil {
				return nil, err
			}

			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				return nil, err
			}

			return wx.NewUploadForm(
				wx.WithFileField("media", filename, body),
			), nil
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
