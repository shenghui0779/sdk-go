package mp

import (
	"encoding/json"
	"net/url"

	"github.com/shenghui0779/gochat/wx"
	"github.com/tidwall/gjson"
)

// SecMediaType 检测的素材类型
type SecMediaType int

// 微信支持的素材类型
var (
	SecMediaAudio SecMediaType = 1 // 音频
	SecMediaImage SecMediaType = 2 // 图片
)

// ImageSecCheck 校验一张图片是否含有违法违规内容
func ImageSecCheck(filename string) wx.Action {
	return wx.NewOpenUploadAPI(ImageSecCheckURL, url.Values{}, wx.NewUploadBody("media", filename, nil), nil)
}

// MediaSecAsyncResult 异步校验结果
type MediaSecAsyncResult struct {
	TraceID string // 任务id，用于匹配异步推送结果
}

// MediaSecCheckAsync 异步校验图片/音频是否含有违法违规内容
func MediaSecCheckAsync(dest *MediaSecAsyncResult, mediaType SecMediaType, mediaURL string) wx.Action {
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
