package oa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/shenghui0779/gochat/helpers"
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

// UploadMedia 上传临时素材到微信服务器
func UploadMedia(mediaType MediaType, filename string, dest *MediaUploadResult) Action {
	return &WechatAPI{
		body: helpers.NewUploadBody("media", filename, func() ([]byte, error) {
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

// MaterialUploadResult 永久素材上传结果
type MaterialUploadResult struct {
	MediaID string `json:"media_id"`
	URL     string `json:"url"`
}

// MaterialArticle 文章素材
type MaterialArticle struct {
	Title              string `json:"title"`
	ThumbMediaID       string `json:"thumb_media_id"`
	Author             string `json:"author"`
	Digest             string `json:"digest"`
	ShowCoverPic       string `json:"show_cover_pic"`
	Content            string `json:"content"`
	ContentSourceURL   string `json:"content_source_url"`
	NeedOpenComment    int    `json:"need_open_comment"`
	OnlyFansCanComment int    `json:"only_fans_can_comment"`
}

// UploadMaterialNews 上传永久图文素材（公众号的素材库保存总数量有上限：图文消息素材、图片素材上限为100000，其他类型为1000）
func UploadMaterialNews(articles []*MaterialArticle, dest *MaterialUploadResult) Action {
	return &WechatAPI{
		body: helpers.NewPostBody(func() ([]byte, error) {
			return json.Marshal(map[string][]*MaterialArticle{
				"articles": articles,
			})
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("POST|%s?access_token=%s", MaterialNewsUploadURL, accessToken)
		},
		decode: func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		},
	}
}

// UploadMaterialImage 上传图文消息内的图片（不受公众号的素材库中图片数量的100000个的限制）
func UploadMaterialImage(filename string, dest *MaterialUploadResult) Action {
	return &WechatAPI{
		body: helpers.NewUploadBody("media", filename, func() ([]byte, error) {
			return ioutil.ReadFile(filename)
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("UPLOAD|%s?access_token=%s", MaterialImageUploadURL, accessToken)
		},
		decode: func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		},
	}
}
