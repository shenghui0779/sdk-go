package mp

import (
	"encoding/json"
	"net/url"

	"github.com/shenghui0779/gochat/wx"
	"github.com/tidwall/gjson"
)

// SecCheckMediaType 检测的素材类型
type SecCheckMediaType int

// 微信支持的素材类型
var (
	SecCheckMediaAudio SecCheckMediaType = 1 // 音频
	SecCheckMediaImage SecCheckMediaType = 2 // 图片
)

// ImageSecCheck 校验一张图片是否含有违法违规内容
func ImageSecCheck(filename string) wx.Action {
	return wx.NewOpenUploadAPI(ImageSecCheckURL, url.Values{}, wx.NewUploadBody("media", filename, nil), nil)
}

// MediaCheckAsyncResult 异步校验结果
type MediaCheckAsyncResult struct {
	TraceID string // 任务id，用于匹配异步推送结果
}

// MediaCheckAsync 异步校验图片/音频是否含有违法违规内容
func MediaCheckAsync(mediaType SecCheckMediaType, mediaURL string, dest *MediaCheckAsyncResult) wx.Action {
	return wx.NewOpenPostAPI(MediaCheckAsyncURL, url.Values{}, wx.NewPostBody(func() ([]byte, error) {
		return json.Marshal(wx.X{
			"media_type": mediaType,
			"media_url":  mediaURL,
		})
	}), func(resp []byte) error {
		dest.TraceID = gjson.GetBytes(resp, "trace_id").String()

		return nil
	})
}

// MsgSecCheck 检查一段文本是否含有违法违规内容
func MsgSecCheck(content string) wx.Action {
	return wx.NewOpenPostAPI(MsgSecCheckURL, url.Values{}, wx.NewPostBody(func() ([]byte, error) {
		return json.Marshal(wx.X{
			"content": content,
		})
	}), nil)
}
