package minip

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/shenghui0779/yiigo"

	"github.com/shenghui0779/gochat/urls"
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

	return wx.NewUploadAction(urls.MinipImageSecCheck,
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

type ParamsMediaCheckAsync struct {
	MediaType SecMediaType `json:"media_type"`
	MediaURL  string       `json:"media_url"`
}

// ResultMediaCheckAsync 异步校验结果
type ResultMediaCheckAsync struct {
	TraceID string // 任务id，用于匹配异步推送结果
}

// MediaCheckAsync 异步校验图片/音频是否含有违法违规内容
func MediaCheckAsync(params *ParamsMediaCheckAsync, result *ResultMediaCheckAsync) wx.Action {
	return wx.NewPostAction(urls.MinipMediaCheckAsync,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// MsgSecCheck 检查一段文本是否含有违法违规内容
func MsgSecCheck(content string) wx.Action {
	return wx.NewPostAction(urls.MinipMsgSecCheck,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(yiigo.X{
				"content": content,
			})
		}),
	)
}
