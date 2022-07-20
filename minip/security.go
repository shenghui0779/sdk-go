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
const (
	SecMediaAudio SecMediaType = 1 // 音频
	SecMediaImage SecMediaType = 2 // 图片
)

// SecCheckScene 检测场景
type SecCheckScene int

const (
	SecSceneDoc     SecCheckScene = 1 // 资料
	SecSceneComment SecCheckScene = 2 // 评论
	SecSceneForum   SecCheckScene = 3 // 论坛
	SecSceneLog     SecCheckScene = 4 // 社交日志
)

// SecCheckSuggest 建议
type SecCheckSuggest string

const (
	SecSuggestRisky  SecCheckSuggest = "risky"
	SecSuggestPass   SecCheckSuggest = "pass"
	SecSuggestReview SecCheckSuggest = "review"
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
	MediaType SecMediaType  `json:"media_type"`
	MediaURL  string        `json:"media_url"`
	Version   int           `json:"version"` // 接口版本号，2.0版本为固定值2
	Scene     SecCheckScene `json:"scene"`   // 场景枚举值
	OpenID    string        `json:"openid"`  // 用户的openid（用户需在近两小时访问过小程序）
}

// ResultMediaCheckAsync 异步校验结果
type ResultMediaCheckAsync struct {
	TraceID string `json:"trace_id"` // 任务id，用于匹配异步推送结果
}

// MediaCheckAsync 异步校验图片/音频是否含有违法违规内容
func MediaCheckAsync(params *ParamsMediaCheckAsync, result *ResultMediaCheckAsync) wx.Action {
	return wx.NewPostAction(urls.MinipMediaCheckAsync,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ParamsMsgCheck struct {
	Content   string        `json:"content"`             // 需检测的文本内容，文本字数的上限为2500字，需使用UTF-8编码
	Version   int           `json:"version"`             // 接口版本号，2.0版本为固定值2
	Scene     SecCheckScene `json:"scene"`               // 场景枚举值
	OpenID    string        `json:"openid"`              // 用户的openid（用户需在近两小时访问过小程序）
	Title     string        `json:"title,omitempty"`     // 文本标题，需使用UTF-8编码
	Nickname  string        `json:"nickname,omitempty"`  // 用户昵称，需使用UTF-8编码
	Signature string        `json:"signature,omitempty"` // 个性签名，该参数仅在资料类场景有效(scene=1)，需使用UTF-8编码
}

type ResultMsgCheck struct {
	TraceID string          `json:"trace_id"`
	Result  *MsgCheckRet    `json:"result"`
	Detail  []*MsgCheckItem `json:"detail"`
}

type MsgCheckRet struct {
	Suggest string `json:"suggest"`
	Label   int    `json:"label"`
}

type MsgCheckItem struct {
	Strategy string `json:"strategy"`
	ErrCode  int    `json:"errcode"`
	Suggest  string `json:"suggest"`
	Label    int    `json:"label"`
	Keyword  string `json:"keyword"`
	Prob     int    `json:"prob"`
}

// MsgSecCheck 检查一段文本是否含有违法违规内容
func MsgSecCheck(params *ParamsMsgCheck, result *ResultMsgCheck) wx.Action {
	return wx.NewPostAction(urls.MinipMsgSecCheck,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// RiskScene 风控场景
type RiskScene int

// 微信支持的风控场景值
const (
	RiskRegister RiskScene = 0 // 注册
	RiskCheat    RiskScene = 1 // 营销作弊
)

// ParamsUserRisk 用户风控参数
type ParamsUserRisk struct {
	AppID        string    `json:"appid"`                   // 小程序appid
	OpenID       string    `json:"openid"`                  // 用户的openid
	Scene        RiskScene `json:"scene"`                   // 场景值，0:注册，1:营销作弊
	MobileNO     string    `json:"mobile_no,omitempty"`     // 用户手机号
	ClientIP     string    `json:"client_ip"`               // 用户访问源ip
	EmailAddress string    `json:"email_address,omitempty"` // 用户邮箱地址
	ExtendedInfo string    `json:"extended_info,omitempty"` // 额外补充信息
	IsTest       bool      `json:"is_test,omitempty"`       // false：正式调用，true：测试调用
}

// ResultUserRisk 用户风控结果
type ResultUserRisk struct {
	RiskRank int `json:"risk_rank"`
}

// GetUserRiskRank 安全风控 - 获取用户的安全等级（无需用户授权）
func GetUserRiskRank(params *ParamsUserRisk, result *ResultUserRisk) wx.Action {
	return wx.NewPostAction(urls.MinipUserRiskRank,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}
