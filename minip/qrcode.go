package minip

import (
	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type EnvVersion string

const (
	EnvRelease EnvVersion = "release"
	EnvTrial   EnvVersion = "trial"
	EnvDevelop EnvVersion = "develop"
)

// QRCode 小程序二维码
type QRCode struct {
	Buffer []byte
}

type ParamsQRCodeCreate struct {
	Path  string `json:"path"`            // 扫码进入的小程序页面路径，最大长度 128 字节，不能为空；对于小游戏，可以只传入 query 部分，来实现传参效果，如：传入 "?foo=bar"，即可在 wx.getLaunchOptionsSync 接口中的 query 参数获取到 {foo:"bar"}。
	Width int    `json:"width,omitempty"` // 二维码的宽度，单位 px。最小 280px，最大 1280px
}

// CreateQRCode 创建小程序二维码（数量有限）
func CreateQRCode(params *ParamsQRCodeCreate, qrcode *QRCode) wx.Action {
	return wx.NewPostAction(urls.MinipQRCodeCreate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalWithNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			qrcode.Buffer = make([]byte, len(resp))
			copy(qrcode.Buffer, resp)

			return nil
		}),
	)
}

// RGB rgb color
type RGB struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}

type ParamsQRCodeGet struct {
	Path      string `json:"path"`                 // 扫码进入的小程序页面路径，最大长度 128 字节，不能为空；对于小游戏，可以只传入 query 部分，来实现传参效果，如：传入 "?foo=bar"，即可在 wx.getLaunchOptionsSync 接口中的 query 参数获取到 {foo:"bar"}。
	Width     int    `json:"width,omitempty"`      // 二维码的宽度，单位 px。最小 280px，最大 1280px
	AutoColor bool   `json:"auto_color,omitempty"` // 自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调
	LineColor *RGB   `json:"line_color,omitempty"` // auto_color 为 false 时生效，使用 rgb 设置颜色，十进制表示，默认：{"r":0,"g":0,"b":0}
	IsHyaline bool   `json:"is_hyaline,omitempty"` // 是否需要透明底色，为 true 时，生成透明底色的小程序码
}

// GetQRCode 获取小程序二维码（数量有限）
func GetQRCode(params *ParamsQRCodeGet, qrcode *QRCode) wx.Action {
	return wx.NewPostAction(urls.MinipQRCodeGet,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalWithNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			qrcode.Buffer = make([]byte, len(resp))
			copy(qrcode.Buffer, resp)

			return nil
		}),
	)
}

type ParamsQRCodeUnlimit struct {
	Scene      string     `json:"scene"`                 // 最大32个可见字符，只支持数字，大小写英文以及部分特殊字符：!#$&'()*+,/:;=?@-._~，其它字符请自行编码为合法字符（因不支持%，中文无法使用 urlencode 处理，请使用其他编码方式）
	Page       string     `json:"page,omitempty"`        // 页面 page，例如 pages/index/index，根路径前不要填加 /，不能携带参数（参数请放在scene字段里），如果不填写这个字段，默认跳主页面
	CheckPath  *bool      `json:"check_path,omitempty"`  // 检查page 是否存在，为 true 时 page 必须是已经发布的小程序存在的页面（否则报错）；为 false 时允许小程序未发布或者 page 不存在， 但page 有数量上限（60000个）请勿滥用
	EnvVersion EnvVersion `json:"env_version,omitempty"` // 要打开的小程序版本。正式版为 "release"，体验版为 "trial"，开发版为 "develop"
	Width      int        `json:"width,omitempty"`       // 二维码的宽度，单位 px。最小 280px，最大 1280px
	AutoColor  bool       `json:"auto_color,omitempty"`  // 自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调
	LineColor  *RGB       `json:"line_color,omitempty"`  // auto_color 为 false 时生效，使用 rgb 设置颜色，十进制表示，默认：{"r":0,"g":0,"b":0}
	IsHyaline  bool       `json:"is_hyaline,omitempty"`  // 是否需要透明底色，为 true 时，生成透明底色的小程序码
}

// GetUnlimitQRCode 获取小程序二维码（数量不限）
func GetUnlimitQRCode(params *ParamsQRCodeUnlimit, qrcode *QRCode) wx.Action {
	return wx.NewPostAction(urls.MinipQRCodeGetUnlimit,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalWithNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			qrcode.Buffer = make([]byte, len(resp))
			copy(qrcode.Buffer, resp)

			return nil
		}),
	)
}
