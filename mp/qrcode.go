package mp

import (
	"bytes"
	"encoding/json"
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

// CreateQRCode 创建小程序二维码 - 数量有限
func CreateQRCode(path string, dest *QRCode, options ...QRCodeOption) wx.Action {
	return wx.NewOpenPostAPI(QRCodeCreateURL, url.Values{}, wx.NewPostBody(func() ([]byte, error) {
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

		bodyStr, err := MarshalWithNoEscapeHTML(params)

		if err != nil {
			return nil, err
		}

		return []byte(bodyStr), nil
	}), func(resp []byte) error {
		dest.Buffer = make([]byte, len(resp))
		copy(dest.Buffer, resp)

		return nil
	})
}

// GetQRCode 获取小程序二维码 - 数量有限
func GetQRCode(path string, dest *QRCode, options ...QRCodeOption) wx.Action {
	return wx.NewOpenPostAPI(QRCodeGetURL, url.Values{}, wx.NewPostBody(func() ([]byte, error) {
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

		bodyStr, err := MarshalWithNoEscapeHTML(params)

		if err != nil {
			return nil, err
		}

		return []byte(bodyStr), nil
	}), func(resp []byte) error {
		dest.Buffer = make([]byte, len(resp))
		copy(dest.Buffer, resp)

		return nil
	})
}

// GetUnlimitQRCode 获取小程序二维码 - 数量不限
func GetUnlimitQRCode(scene string, dest *QRCode, options ...QRCodeOption) wx.Action {
	return wx.NewOpenPostAPI(QRCodeGetUnlimitURL, url.Values{}, wx.NewPostBody(func() ([]byte, error) {
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

		bodyStr, err := MarshalWithNoEscapeHTML(params)

		if err != nil {
			return nil, err
		}

		return []byte(bodyStr), nil
	}), func(resp []byte) error {
		dest.Buffer = make([]byte, len(resp))
		copy(dest.Buffer, resp)

		return nil
	})
}

// MarshalWithNoEscapeHTML marshal with no escape HTML
func MarshalWithNoEscapeHTML(v interface{}) (string, error) {
	buf := wx.BufferPool.Get().(*bytes.Buffer)
	buf.Reset()

	defer wx.BufferPool.Put(buf)

	jsonEncoder := json.NewEncoder(buf)
	jsonEncoder.SetEscapeHTML(false)

	if err := jsonEncoder.Encode(v); err != nil {
		return "", err
	}

	return buf.String(), nil
}
