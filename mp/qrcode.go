package mp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/iiinsomnia/gochat/consts"
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

	resp, err := q.client.Post(fmt.Sprintf("%s?access_token=%s", consts.MPQRCodeCreateURL, accessToken), b, utils.WithRequestHeader("Content-Type", "application/json; charset=utf-8"))

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

	resp, err := q.client.Post(fmt.Sprintf("%s?access_token=%s", consts.MPQRCodeGetURL, accessToken), b, utils.WithRequestHeader("Content-Type", "application/json; charset=utf-8"))

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

	resp, err := q.client.Post(fmt.Sprintf("%s?access_token=%s", consts.MPQRCodeGetUnlimitURL, accessToken), b, utils.WithRequestHeader("Content-Type", "application/json; charset=utf-8"))

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
