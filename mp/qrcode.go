package mp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/iiinsomnia/gochat/utils"
	"github.com/tidwall/gjson"
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

// funcWXAQRCodeOption implements wxa_qrcode option
type funcWXAQRCodeOption struct {
	f func(options *qrcodeOptions)
}

func (fo *funcWXAQRCodeOption) apply(o *qrcodeOptions) {
	fo.f(o)
}

func newFuncWXAQRCodeOption(f func(options *qrcodeOptions)) *funcWXAQRCodeOption {
	return &funcWXAQRCodeOption{f: f}
}

// WithWXAQRPage specifies the `page` to wxa_qrcode.
func WithWXAQRPage(s string) QRCodeOption {
	return newFuncWXAQRCodeOption(func(o *qrcodeOptions) {
		o.page = s
	})
}

// WithWXAQRWidth specifies the `width` to wxa_qrcode.
func WithWXAQRWidth(w int) QRCodeOption {
	return newFuncWXAQRCodeOption(func(o *qrcodeOptions) {
		o.width = w
	})
}

// WithWXAQRAutoColor specifies the `auto_color` to wxa_qrcode.
func WithWXAQRAutoColor(b bool) QRCodeOption {
	return newFuncWXAQRCodeOption(func(o *qrcodeOptions) {
		o.autoColor = b
	})
}

// WithWXAQRLineColor specifies the `line_color` to wxa_qrcode.
func WithWXAQRLineColor(m map[string]int) QRCodeOption {
	return newFuncWXAQRCodeOption(func(o *qrcodeOptions) {
		o.lineColor = m
	})
}

// WithWXAQRIsHyaline specifies the `is_hyaline` to wxa_qrcode.
func WithWXAQRIsHyaline(b bool) QRCodeOption {
	return newFuncWXAQRCodeOption(func(o *qrcodeOptions) {
		o.isHyaline = b
	})
}

// QRCode 小程序二维码
type QRCode struct {
	client *utils.HTTPClient
}

// Create 数量有限
func (q *QRCode) Create(accessToken, path string, options ...QRCodeOption) ([]byte, error) {
	o := new(qrcodeOptions)

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

	resp, err := q.client.Post(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token=%s", accessToken), b, utils.WithRequestHeader("Content-Type", "application/json; charset=utf-8"))

	if err != nil {
		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return nil, errors.New(r.Get("errmsg").String())
	}

	return resp, nil
}

// Get 数量有限
func (q *QRCode) Get(accessToken, path string, options ...QRCodeOption) ([]byte, error) {
	o := new(qrcodeOptions)

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

	resp, err := q.client.Post(fmt.Sprintf("https://api.weixin.qq.com/wxa/getwxacode?access_token=%s", accessToken), b, utils.WithRequestHeader("Content-Type", "application/json; charset=utf-8"))

	if err != nil {
		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return nil, errors.New(r.Get("errmsg").String())
	}

	return resp, nil
}

// GetUnlimit 数量不限
func (q *QRCode) GetUnlimit(accessToken, scene string, options ...QRCodeOption) ([]byte, error) {
	o := new(qrcodeOptions)

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

	resp, err := q.client.Post(fmt.Sprintf("https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%s", accessToken), b, utils.WithRequestHeader("Content-Type", "application/json; charset=utf-8"))

	if err != nil {
		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return nil, errors.New(r.Get("errmsg").String())
	}

	return resp, nil
}

// MarshalWithNoEscapeHTML marshal with no escape HTML
func MarshalWithNoEscapeHTML(v interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)

	jsonEncoder := json.NewEncoder(buf)
	jsonEncoder.SetEscapeHTML(false)

	if err := jsonEncoder.Encode(v); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
