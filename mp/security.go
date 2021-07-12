package mp

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/shenghui0779/yiigo"
	"github.com/tidwall/gjson"

	"github.com/shenghui0779/gochat/wx"
)

// SecMediaType 检测的素材类型
type SecMediaType int

// 微信支持的素材类型
var (
	SecMediaAudio SecMediaType = 1 // 音频
	SecMediaImage SecMediaType = 2 // 图片
)

// ImageSecCheck 校验一张图片是否含有违法违规内容
func ImageSecCheck(path string) wx.Action {
	_, filename := filepath.Split(path)

	return wx.NewAction(ImageSecCheckURL,
		wx.WithMethod(wx.MethodUpload),
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
	)
}

// MediaSecAsyncResult 异步校验结果
type MediaSecAsyncResult struct {
	TraceID string // 任务id，用于匹配异步推送结果
}

// MediaSecCheckAsync 异步校验图片/音频是否含有违法违规内容
func MediaSecCheckAsync(dest *MediaSecAsyncResult, mediaType SecMediaType, mediaURL string) wx.Action {
	return wx.NewAction(MediaCheckAsyncURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(yiigo.X{
				"media_type": mediaType,
				"media_url":  mediaURL,
			})
		}),
		wx.WithDecode(func(resp []byte) error {
			dest.TraceID = gjson.GetBytes(resp, "trace_id").String()

			return nil
		}),
	)
}

// MsgSecCheck 检查一段文本是否含有违法违规内容
func MsgSecCheck(content string) wx.Action {
	return wx.NewAction(MsgSecCheckURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(yiigo.X{
				"content": content,
			})
		}),
	)
}
