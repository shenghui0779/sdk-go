package mp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tidwall/gjson"
)

type wxaqrcodeOptions struct {
	page      string
	width     int
	autoColor bool
	lineColor map[string]int
	isHyaline bool
}

// WXAQRCodeOption configures how we set up the wxa_qrcode
type WXAQRCodeOption interface {
	apply(options *wxaqrcodeOptions)
}

// funcWXAQRCodeOption implements wxa_qrcode option
type funcWXAQRCodeOption struct {
	f func(options *wxaqrcodeOptions)
}

func (fo *funcWXAQRCodeOption) apply(o *wxaqrcodeOptions) {
	fo.f(o)
}

func newFuncWXAQRCodeOption(f func(options *wxaqrcodeOptions)) *funcWXAQRCodeOption {
	return &funcWXAQRCodeOption{f: f}
}

// WithWXAQRPage specifies the `page` to wxa_qrcode.
func WithWXAQRPage(s string) WXAQRCodeOption {
	return newFuncWXAQRCodeOption(func(o *wxaqrcodeOptions) {
		o.page = s
	})
}

// WithWXAQRWidth specifies the `width` to wxa_qrcode.
func WithWXAQRWidth(w int) WXAQRCodeOption {
	return newFuncWXAQRCodeOption(func(o *wxaqrcodeOptions) {
		o.width = w
	})
}

// WithWXAQRAutoColor specifies the `auto_color` to wxa_qrcode.
func WithWXAQRAutoColor(b bool) WXAQRCodeOption {
	return newFuncWXAQRCodeOption(func(o *wxaqrcodeOptions) {
		o.autoColor = b
	})
}

// WithWXAQRLineColor specifies the `line_color` to wxa_qrcode.
func WithWXAQRLineColor(m map[string]int) WXAQRCodeOption {
	return newFuncWXAQRCodeOption(func(o *wxaqrcodeOptions) {
		o.lineColor = m
	})
}

// WithWXAQRIsHyaline specifies the `is_hyaline` to wxa_qrcode.
func WithWXAQRIsHyaline(b bool) WXAQRCodeOption {
	return newFuncWXAQRCodeOption(func(o *wxaqrcodeOptions) {
		o.isHyaline = b
	})
}

// WXAQRCode 小程序二维码
type WXAQRCode struct {
	accessToken string
}

// CreateWXQRACode 数量有限
func (q *WXAQRCode) CreateWXAQRCode(path string, options ...WXAQRCodeOption) ([]byte, error) {
	o := new(wxaqrcodeOptions)

	if len(options) > 0 {
		for _, option := range options {
			option.apply(o)
		}
	}

	body := utils.X{"path": path}

	if o.width != 0 {
		body["width"] = o.width
	}

	b, err := MarshalWithNoEscapeHTML(body)

	if err != nil {
		return nil, err
	}

	resp, err := utils.HTTPPost(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token=%s", q.accessToken), b, utils.WithRequestHeader("Content-Type", "application/json; charset=utf-8"))

	if err != nil {
		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return nil, errors.New(r.Get("errmsg").String())
	}

	return resp, nil
}

// GetWXACode 数量有限
func (q *WXAQRCode) GetWXACode(path string, options ...WXAQRCodeOption) ([]byte, error) {
	o := new(wxaqrcodeOptions)

	if len(options) > 0 {
		for _, option := range options {
			option.apply(o)
		}
	}

	body := utils.X{"path": path}

	if o.width != 0 {
		body["width"] = o.width
	}

	if o.autoColor {
		body["auto_color"] = true
	}

	if len(o.lineColor) != 0 {
		body["line_color"] = o.lineColor
	}

	if o.isHyaline {
		body["is_hyaline"] = true
	}

	b, err := MarshalWithNoEscapeHTML(body)

	if err != nil {
		return nil, err
	}

	resp, err := utils.HTTPPost(fmt.Sprintf("https://api.weixin.qq.com/wxa/getwxacode?access_token=%s", q.accessToken), b, utils.WithRequestHeader("Content-Type", "application/json; charset=utf-8"))

	if err != nil {
		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return nil, errors.New(r.Get("errmsg").String())
	}

	return resp, nil
}

// GetWXACodeUnlimit 数量不限
func (q *WXAQRCode) GetWXACodeUnlimit(scene string, options ...WXAQRCodeOption) ([]byte, error) {
	o := new(wxaqrcodeOptions)

	if len(options) > 0 {
		for _, option := range options {
			option.apply(o)
		}
	}

	body := utils.X{"scene": scene}

	if o.page != "" {
		body["page"] = o.page
	}

	if o.width != 0 {
		body["width"] = o.width
	}

	if o.autoColor {
		body["auto_color"] = true
	}

	if len(o.lineColor) != 0 {
		body["line_color"] = o.lineColor
	}

	if o.isHyaline {
		body["is_hyaline"] = true
	}

	b, err := MarshalWithNoEscapeHTML(body)

	if err != nil {
		return nil, err
	}

	resp, err := utils.HTTPPost(fmt.Sprintf("https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%s", q.accessToken), b, utils.WithRequestHeader("Content-Type", "application/json; charset=utf-8"))

	if err != nil {
		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return nil, errors.New(r.Get("errmsg").String())
	}

	return resp, nil
}

// NewWXAQRCode ...
func NewWXAQRCode(accessToken string) *WXAQRCode {
	return &WXAQRCode{accessToken: accessToken}
}

// MarshalWithNoEscapeHTML ...
func MarshalWithNoEscapeHTML(v interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)

	jsonEncoder := json.NewEncoder(buf)
	jsonEncoder.SetEscapeHTML(false)

	if err := jsonEncoder.Encode(v); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
