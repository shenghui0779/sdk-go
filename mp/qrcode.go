package mp

import (
	"github.com/shenghui0779/gochat/wx"
)

type qrcodeSettings struct {
	page      string
	width     int
	autoColor bool
	lineColor map[string]int
	isHyaline bool
}

// QRCodeOption configures how we set up the wxa_qrcode
type QRCodeOption func(s *qrcodeSettings)

// WithQRCodePage specifies the `page` to qrcode.
func WithQRCodePage(page string) QRCodeOption {
	return func(settings *qrcodeSettings) {
		settings.page = page
	}
}

// WithQRCodeWidth specifies the `width` to qrcode.
func WithQRCodeWidth(width int) QRCodeOption {
	return func(settings *qrcodeSettings) {
		settings.width = width
	}
}

// WithQRCodeAutoColor specifies the `auto_color` to qrcode.
func WithQRCodeAutoColor() QRCodeOption {
	return func(settings *qrcodeSettings) {
		settings.autoColor = true
	}
}

// WithQRCodeLineColor specifies the `line_color` to qrcode.
func WithQRCodeLineColor(r, g, b int) QRCodeOption {
	return func(settings *qrcodeSettings) {
		settings.lineColor = map[string]int{
			"r": r,
			"g": g,
			"b": b,
		}
	}
}

// WithQRCodeIsHyaline specifies the `is_hyaline` to qrcode.
func WithQRCodeIsHyaline() QRCodeOption {
	return func(settings *qrcodeSettings) {
		settings.isHyaline = true
	}
}

// QRCode 小程序二维码
type QRCode struct {
	Buffer []byte
}

// CreateQRCode 创建小程序二维码（数量有限）
func CreateQRCode(dest *QRCode, path string, options ...QRCodeOption) wx.Action {
	return wx.NewAction(QRCodeCreateURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			settings := new(qrcodeSettings)

			if len(options) != 0 {
				for _, f := range options {
					f(settings)
				}
			}

			params := wx.X{"path": path}

			if settings.width != 0 {
				params["width"] = settings.width
			}

			return wx.MarshalWithNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			dest.Buffer = make([]byte, len(resp))
			copy(dest.Buffer, resp)

			return nil
		}),
	)
}

// GetQRCode 获取小程序二维码（数量有限）
func GetQRCode(dest *QRCode, path string, options ...QRCodeOption) wx.Action {
	return wx.NewAction(QRCodeGetURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			settings := new(qrcodeSettings)

			if len(options) != 0 {
				for _, f := range options {
					f(settings)
				}
			}

			params := wx.X{"path": path}

			if settings.width != 0 {
				params["width"] = settings.width
			}

			if settings.autoColor {
				params["auto_color"] = true
			}

			if len(settings.lineColor) != 0 {
				params["line_color"] = settings.lineColor
			}

			if settings.isHyaline {
				params["is_hyaline"] = true
			}

			return wx.MarshalWithNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			dest.Buffer = make([]byte, len(resp))
			copy(dest.Buffer, resp)

			return nil
		}),
	)
}

// GetUnlimitQRCode 获取小程序二维码（数量不限）
func GetUnlimitQRCode(dest *QRCode, scene string, options ...QRCodeOption) wx.Action {
	return wx.NewAction(QRCodeGetUnlimitURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			settings := new(qrcodeSettings)

			if len(options) != 0 {
				for _, f := range options {
					f(settings)
				}
			}

			params := wx.X{"scene": scene}

			if settings.page != "" {
				params["page"] = settings.page
			}

			if settings.width != 0 {
				params["width"] = settings.width
			}

			if settings.autoColor {
				params["auto_color"] = true
			}

			if len(settings.lineColor) != 0 {
				params["line_color"] = settings.lineColor
			}

			if settings.isHyaline {
				params["is_hyaline"] = true
			}

			return wx.MarshalWithNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			dest.Buffer = make([]byte, len(resp))
			copy(dest.Buffer, resp)

			return nil
		}),
	)
}
