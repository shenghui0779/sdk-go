package externalcontact

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/chenghonour/gochat/urls"
	"github.com/chenghonour/gochat/wx"
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

			return wx.NewUploadForm(
				wx.WithFormFile("media", filename, func(w io.Writer) error {
					f, err := os.Open(path)

					if err != nil {
						return err
					}

					defer f.Close()

					if _, err = io.Copy(w, f); err != nil {
						return err
					}

					return nil
				}),
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
			return wx.NewUploadForm(
				wx.WithFormFile("media", filename, func(w io.Writer) error {
					resp, err := wx.HTTPGet(context.Background(), url)

					if err != nil {
						return err
					}

					defer resp.Body.Close()

					if _, err = io.Copy(w, resp.Body); err != nil {
						return err
					}

					return nil
				}),
			), nil
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
