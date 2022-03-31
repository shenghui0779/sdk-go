package minip

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

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
func ImageSecCheck(imgPath string) wx.Action {
	_, filename := filepath.Split(imgPath)

	return wx.NewPostAction(urls.MinipImageSecCheck,
		wx.WithUpload(func() (wx.UploadForm, error) {
			path, err := filepath.Abs(filepath.Clean(imgPath))

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
	)
}

type ParamsMediaCheckAsync struct {
	MediaType SecMediaType `json:"media_type"`
	MediaURL  string       `json:"media_url"`
}

// ResultMediaCheckAsync 异步校验结果
type ResultMediaCheckAsync struct {
	TraceID string `json:"trace_id"` // 任务id，用于匹配异步推送结果
}

// MediaCheckAsync 异步校验图片/音频是否含有违法违规内容
func MediaCheckAsync(mediaType SecMediaType, mediaURL string, result *ResultMediaCheckAsync) wx.Action {
	params := &ParamsMediaCheckAsync{
		MediaType: mediaType,
		MediaURL:  mediaURL,
	}

	return wx.NewPostAction(urls.MinipMediaCheckAsync,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
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
			return wx.MarshalNoEscapeHTML(wx.M{
				"content": content,
			})
		}),
	)
}
