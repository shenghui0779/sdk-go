package mp

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/shenghui0779/gochat/internal"
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

// WithQRCodePage specifies the `page` to wxa_qrcode.
func WithQRCodePage(s string) QRCodeOption {
	return newFuncQRCodeOption(func(o *qrcodeOptions) {
		o.page = s
	})
}

// WithQRCodeWidth specifies the `width` to wxa_qrcode.
func WithQRCodeWidth(w int) QRCodeOption {
	return newFuncQRCodeOption(func(o *qrcodeOptions) {
		o.width = w
	})
}

// WithQRCodeAutoColor specifies the `auto_color` to wxa_qrcode.
func WithQRCodeAutoColor(b bool) QRCodeOption {
	return newFuncQRCodeOption(func(o *qrcodeOptions) {
		o.autoColor = b
	})
}

// WithQRCodeLineColor specifies the `line_color` to wxa_qrcode.
func WithQRCodeLineColor(m map[string]int) QRCodeOption {
	return newFuncQRCodeOption(func(o *qrcodeOptions) {
		o.lineColor = m
	})
}

// WithQRCodeIsHyaline specifies the `is_hyaline` to wxa_qrcode.
func WithQRCodeIsHyaline(b bool) QRCodeOption {
	return newFuncQRCodeOption(func(o *qrcodeOptions) {
		o.isHyaline = b
	})
}

// QRCode 小程序二维码
type QRCode struct {
	Buffer []byte
}

// CreateQRCode 创建小程序二维码 - 数量有限
func CreateQRCode(path string, dest *QRCode, options ...QRCodeOption) internal.Action {
	return &WechatAPI{
		body: internal.NewPostBody(func() ([]byte, error) {
			o := new(qrcodeOptions)

			if len(options) > 0 {
				for _, option := range options {
					option.apply(o)
				}
			}

			params := internal.X{"path": path}

			if o.width != 0 {
				params["width"] = o.width
			}

			bodyStr, err := MarshalWithNoEscapeHTML(params)

			if err != nil {
				return nil, err
			}

			return []byte(bodyStr), nil
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("POST|%s?access_token=%s", QRCodeCreateURL, accessToken)
		},
		decode: func(resp []byte) error {
			dest.Buffer = make([]byte, len(resp))
			copy(dest.Buffer, resp)

			return nil
		},
	}
}

// GetQRCode 获取小程序二维码 - 数量有限
func GetQRCode(path string, dest *QRCode, options ...QRCodeOption) internal.Action {
	return &WechatAPI{
		body: internal.NewPostBody(func() ([]byte, error) {
			o := new(qrcodeOptions)

			if len(options) > 0 {
				for _, option := range options {
					option.apply(o)
				}
			}

			params := internal.X{"path": path}

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
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("POST|%s?access_token=%s", QRCodeGetURL, accessToken)
		},
		decode: func(resp []byte) error {
			dest.Buffer = make([]byte, len(resp))
			copy(dest.Buffer, resp)

			return nil
		},
	}
}

// GetUnlimitQRCode 获取小程序二维码 - 数量不限
func GetUnlimitQRCode(scene string, dest *QRCode, options ...QRCodeOption) internal.Action {
	return &WechatAPI{
		body: internal.NewPostBody(func() ([]byte, error) {
			o := new(qrcodeOptions)

			if len(options) > 0 {
				for _, option := range options {
					option.apply(o)
				}
			}

			params := internal.X{"scene": scene}

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
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("POST|%s?access_token=%s", QRCodeGetUnlimitURL, accessToken)
		},
		decode: func(resp []byte) error {
			dest.Buffer = make([]byte, len(resp))
			copy(dest.Buffer, resp)

			return nil
		},
	}
}

// MarshalWithNoEscapeHTML marshal with no escape HTML
func MarshalWithNoEscapeHTML(v interface{}) (string, error) {
	buf := internal.BufferPool.Get().(*bytes.Buffer)
	buf.Reset()

	defer internal.BufferPool.Put(buf)

	jsonEncoder := json.NewEncoder(buf)
	jsonEncoder.SetEscapeHTML(false)

	if err := jsonEncoder.Encode(v); err != nil {
		return "", err
	}

	return buf.String(), nil
}
