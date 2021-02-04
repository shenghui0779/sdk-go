package oa

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"

	"github.com/shenghui0779/gochat/wx"
)

// MediaType 素材类型
type MediaType string

// 微信支持的素材类型
var (
	MediaImage MediaType = "image" // 图片
	MediaVoice MediaType = "voice" // 音频
	MediaVideo MediaType = "video" // 视频
	MediaThumb MediaType = "thumb" // 缩略图
)

// MediaUploadResult 临时素材上传结果
type MediaUploadResult struct {
	Type      string `json:"type"`
	MediaID   string `json:"media_id"`
	CreatedAt int64  `json:"created_at"`
}

// UploadMedia 上传临时素材
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

// MaterialAddResult 永久素材新增结果
type MaterialAddResult struct {
	MediaID string `json:"media_id"`
	URL     string `json:"url"`
}

// NewsArticle 文章素材
type NewsArticle struct {
	Title              string `json:"title"`
	ThumbMediaID       string `json:"thumb_media_id"`
	Author             string `json:"author"`
	Digest             string `json:"digest"`
	ShowCoverPic       int    `json:"show_cover_pic"`
	Content            string `json:"content"`
	ContentSourceURL   string `json:"content_source_url"`
	NeedOpenComment    int    `json:"need_open_comment"`
	OnlyFansCanComment int    `json:"only_fans_can_comment"`
}

// AddNews 新增永久图文素材（公众号的素材库保存总数量有上限：图文消息素材、图片素材上限为100000，其他类型为1000）
func AddNews(dest *MaterialAddResult, articles ...*NewsArticle) wx.Action {
	return wx.NewAction(NewsAddURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(wx.X{"articles": articles})
		}),
		wx.WithDecode(func(resp []byte) error {
			dest.MediaID = gjson.GetBytes(resp, "media_id").String()

			return nil
		}),
	)
}

// UploadNewsImage 上传图文消息内的图片（不受公众号的素材库中图片数量的100000个的限制，图片仅支持jpg/png格式，大小必须在1MB以下）
func UploadNewsImage(dest *MaterialAddResult, filename string) wx.Action {
	return wx.NewAction(NewsImageUploadURL,
		wx.WithMethod(wx.MethodUpload),
		wx.WithUploadForm("media", filename, nil),
		wx.WithDecode(func(resp []byte) error {
			dest.URL = gjson.GetBytes(resp, "url").String()

			return nil
		}),
	)
}

// AddMaterial 新增其他类型永久素材（支持图片、音频、缩略图）
func AddMaterial(dest *MaterialAddResult, mediaType MediaType, filename string) wx.Action {
	return wx.NewAction(MaterialAddURL,
		wx.WithMethod(wx.MethodUpload),
		wx.WithQuery("type", string(mediaType)),
		wx.WithUploadForm("media", filename, nil),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// UploadVideo 上传视频永久素材
func UploadVideo(dest *MaterialAddResult, filename, title, introduction string) wx.Action {
	return wx.NewAction(MaterialAddURL,
		wx.WithMethod(wx.MethodUpload),
		wx.WithQuery("type", string(MediaVideo)),
		wx.WithUploadForm("media", filename, map[string]string{
			"description": fmt.Sprintf(`{"title":"%s", "introduction":"%s"}`, title, introduction),
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// DeleteMaterial 删除永久素材
func DeleteMaterial(mediaID string) wx.Action {
	return wx.NewAction(MaterialDeleteURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(wx.X{"media_id": mediaID})
		}),
	)
}
