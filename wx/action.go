package wx

import (
	"fmt"
	"net/url"

	"github.com/shenghui0779/yiigo"
)

// ActionMethod http request method
type ActionMethod string

const (
	MethodGet    ActionMethod = "GET"
	MethodPost   ActionMethod = "POST"
	MethodUpload ActionMethod = "UPLOAD"
	MethodNone   ActionMethod = "NONE"
)

// Action is the interface that handle wechat api
type Action interface {
	// URL returns request url
	URL(accessToken ...string) string

	// Method returns action method
	Method() ActionMethod

	// WXML returns body for xml request
	WXML(appid, mchid, nonce string) (WXML, error)

	// Body returns body for post request
	Body() ([]byte, error)

	// UploadForm returns form for uploading media
	UploadForm() (yiigo.UploadForm, error)

	// Decode decodes response
	Decode() func(resp []byte) error

	// TLS specifies the request with certificate
	TLS() bool
}

type UploadField struct {
	FileField string
	Filename  string
	MetaField string
	Metadata  string
}

type action struct {
	method      ActionMethod
	reqURL      string
	query       url.Values
	wxml        func(appid, mchid, nonce string) (WXML, error)
	body        func() ([]byte, error)
	uploadfield *UploadField
	decode      func(resp []byte) error
	tls         bool
}

func (a *action) Method() ActionMethod {
	return a.method
}

func (a *action) URL(accessToken ...string) string {
	if len(accessToken) != 0 {
		a.query.Set("access_token", accessToken[0])
	}

	if len(a.query) == 0 {
		return a.reqURL
	}

	return fmt.Sprintf("%s?%s", a.reqURL, a.query.Encode())
}

func (a *action) WXML(appid, mchid, nonce string) (WXML, error) {
	if a.wxml == nil {
		return WXML{}, nil
	}

	return a.wxml(appid, mchid, nonce)
}

func (a *action) Body() ([]byte, error) {
	if a.body == nil {
		return nil, nil
	}

	return a.body()
}

func (a *action) UploadForm() (yiigo.UploadForm, error) {
	if a.uploadfield == nil {
		return yiigo.NewUploadForm(), nil
	}

	body, err := a.Body()

	if err != nil {
		return nil, err
	}

	fields := []yiigo.UploadField{
		yiigo.WithFileField(a.uploadfield.FileField, a.uploadfield.Filename, body),
	}

	if len(a.uploadfield.MetaField) != 0 {
		fields = append(fields, yiigo.WithFormField(a.uploadfield.MetaField, a.uploadfield.Metadata))
	}

	return yiigo.NewUploadForm(fields...), nil
}

func (a *action) Decode() func(resp []byte) error {
	return a.decode
}

func (a *action) TLS() bool {
	return a.tls
}

// NewAction returns a new action
func NewAction(method ActionMethod, reqURL string, options ...ActionOption) Action {
	a := &action{
		method: method,
		reqURL: reqURL,
		query:  url.Values{},
	}

	for _, f := range options {
		f(a)
	}

	return a
}

// NewGetAction returns a new action with GET method
func NewGetAction(reqURL string, options ...ActionOption) Action {
	return NewAction(MethodGet, reqURL, options...)
}

// NewPostAction returns a new action with POST method
func NewPostAction(reqURL string, options ...ActionOption) Action {
	return NewAction(MethodPost, reqURL, options...)
}

// NewUploadAction returns a new action with UPLOAD method
func NewUploadAction(reqURL string, options ...ActionOption) Action {
	return NewAction(MethodUpload, reqURL, options...)
}
