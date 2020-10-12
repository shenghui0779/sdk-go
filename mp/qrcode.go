package mp

import (
	"encoding/json"
	"fmt"

	"github.com/shenghui0779/gochat/utils"
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
	mp      *WXMP
	options []utils.RequestOption
}

// Create 数量有限
func (q *QRCode) Create(accessToken, path string, options ...QRCodeOption) ([]byte, error) {
	o := new(qrcodeOptions)

	if len(options) > 0 {
		for _, option := range options {
			option.apply(o)
		}
	}

	params := utils.X{"path": path}

	if o.width != 0 {
		params["width"] = o.width
	}

	bodyStr, err := MarshalWithNoEscapeHTML(params)

	if err != nil {
		return nil, err
	}

	return q.mp.post(fmt.Sprintf("%s?access_token=%s", QRCodeCreateURL, accessToken), []byte(bodyStr), q.options...)
}

// Get 数量有限
func (q *QRCode) Get(accessToken, path string, options ...QRCodeOption) ([]byte, error) {
	o := new(qrcodeOptions)

	if len(options) > 0 {
		for _, option := range options {
			option.apply(o)
		}
	}

	params := utils.X{"path": path}

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

	return q.mp.post(fmt.Sprintf("%s?access_token=%s", QRCodeGetURL, accessToken), []byte(bodyStr), q.options...)
}

// GetUnlimit 数量不限
func (q *QRCode) GetUnlimit(accessToken, scene string, options ...QRCodeOption) ([]byte, error) {
	o := new(qrcodeOptions)

	if len(options) > 0 {
		for _, option := range options {
			option.apply(o)
		}
	}

	params := utils.X{"scene": scene}

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

	return q.mp.post(fmt.Sprintf("%s?access_token=%s", QRCodeGetUnlimitURL, accessToken), []byte(bodyStr), q.options...)
}

// MarshalWithNoEscapeHTML marshal with no escape HTML
func MarshalWithNoEscapeHTML(v interface{}) (string, error) {
	buf := utils.BufPool.Get()
	defer utils.BufPool.Put(buf)

	jsonEncoder := json.NewEncoder(buf)
	jsonEncoder.SetEscapeHTML(false)

	if err := jsonEncoder.Encode(v); err != nil {
		return "", err
	}

	return buf.String(), nil
}
