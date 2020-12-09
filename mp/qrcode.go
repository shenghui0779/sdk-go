package mp

import (
	"net/url"

	"github.com/shenghui0779/gochat/wx"
)

type qrcodeOptions struct {
	page      string
	width     int
	autoColor bool
	lineColor map[string]int
	isHyaline bool
}

// QRCodeOption configures how we set up the wxa_qrcode
type QRCodeOption interface {
	apply(options *qrcodeOptions)
}

// funcQRCodeOption implements wxa_qrcode option
type funcQRCodeOption struct {
	f func(options *qrcodeOptions)
}

func (fo *funcQRCodeOption) apply(o *qrcodeOptions) {
	fo.f(o)
}

func newFuncQRCodeOption(f func(options *qrcodeOptions)) *funcQRCodeOption {
	return &funcQRCodeOption{f: f}
}

// WithQRCodePage specifies the `page` to qrcode.
func WithQRCodePage(s string) QRCodeOption {
	return newFuncQRCodeOption(func(o *qrcodeOptions) {
		o.page = s
	})
}

// WithQRCodeWidth specifies the `width` to qrcode.
func WithQRCodeWidth(w int) QRCodeOption {
	return newFuncQRCodeOption(func(o *qrcodeOptions) {
		o.width = w
	})
}

// WithQRCodeAutoColor specifies the `auto_color` to qrcode.
func WithQRCodeAutoColor() QRCodeOption {
	return newFuncQRCodeOption(func(o *qrcodeOptions) {
		o.autoColor = true
	})
}

// WithQRCodeLineColor specifies the `line_color` to qrcode.
func WithQRCodeLineColor(r, g, b int) QRCodeOption {
	return newFuncQRCodeOption(func(o *qrcodeOptions) {
		o.lineColor = map[string]int{
			"r": r,
			"g": g,
			"b": b,
		}
	})
}

// WithQRCodeIsHyaline specifies the `is_hyaline` to qrcode.
func WithQRCodeIsHyaline() QRCodeOption {
	return newFuncQRCodeOption(func(o *qrcodeOptions) {
		o.isHyaline = true
	})
}

// QRCode 小程序二维码
type QRCode struct {
	Buffer []byte
}

// CreateQRCode 创建小程序二维码（数量有限）
func CreateQRCode(dest *QRCode, path string, options ...QRCodeOption) wx.Action {
	return wx.NewPostAPI(QRCodeCreateURL, url.Values{}, func() ([]byte, error) {
		o := new(qrcodeOptions)

		if len(options) > 0 {
			for _, option := range options {
				option.apply(o)
			}
		}

		params := wx.X{"path": path}

		if o.width != 0 {
			params["width"] = o.width
		}

		return wx.MarshalWithNoEscapeHTML(params)
	}, func(resp []byte) error {
		dest.Buffer = make([]byte, len(resp))
		copy(dest.Buffer, resp)

		return nil
	})
}

// GetQRCode 获取小程序二维码（数量有限）
func GetQRCode(dest *QRCode, path string, options ...QRCodeOption) wx.Action {
	return wx.NewPostAPI(QRCodeGetURL, url.Values{}, func() ([]byte, error) {
		o := new(qrcodeOptions)

		if len(options) > 0 {
			for _, option := range options {
				option.apply(o)
			}
		}

		params := wx.X{"path": path}

		if o.width != 0 {
			params["width"] = o.width
		}

		if o.autoColor {
			params["auto_color"] = true
		}

		if len(o.lineColor) != 0 {
			params["line_color"] = o.lineColor
		}

		if o.isHyaline {
			params["is_hyaline"] = true
		}

		return wx.MarshalWithNoEscapeHTML(params)
	}, func(resp []byte) error {
		dest.Buffer = make([]byte, len(resp))
		copy(dest.Buffer, resp)

		return nil
	})
}

// GetUnlimitQRCode 获取小程序二维码（数量不限）
func GetUnlimitQRCode(dest *QRCode, scene string, options ...QRCodeOption) wx.Action {
	return wx.NewPostAPI(QRCodeGetUnlimitURL, url.Values{}, func() ([]byte, error) {
		o := new(qrcodeOptions)

		if len(options) > 0 {
			for _, option := range options {
				option.apply(o)
			}
		}

		params := wx.X{"scene": scene}

		if o.page != "" {
			params["page"] = o.page
		}

		if o.width != 0 {
			params["width"] = o.width
		}

		if o.autoColor {
			params["auto_color"] = true
		}

		if len(o.lineColor) != 0 {
			params["line_color"] = o.lineColor
		}

		if o.isHyaline {
			params["is_hyaline"] = true
		}

		return wx.MarshalWithNoEscapeHTML(params)
	}, func(resp []byte) error {
		dest.Buffer = make([]byte, len(resp))
		copy(dest.Buffer, resp)

		return nil
	})
}
