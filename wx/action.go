package wx

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/shenghui0779/yiigo"
)

// Action is the interface that handle wechat api
type Action interface {
	// Method returns action method
	Method() string

	// URL returns request url
	URL(accessToken ...string) string

	// WXML returns body for xml request
	WXML(appid, mchid, nonce string) (WXML, error)

	// Body returns body for post request
	Body() ([]byte, error)

	// UploadForm returns form for uploading media
	UploadForm() (yiigo.UploadForm, error)

	// Decode decodes response
	Decode(resp []byte) error

	// IsUpload specifies the request does upload
	IsUpload() bool

	// TLS specifies the request with certificate
	IsTLS() bool
}

type UploadField struct {
	FileField string
	Filename  string
	MetaField string
	Metadata  string
}

type action struct {
	method     string
	reqURL     string
	query      url.Values
	wxml       func(appid, mchid, nonce string) (WXML, error)
	body       func() ([]byte, error)
	uploadform func() (yiigo.UploadForm, error)
	decode     func(resp []byte) error
	upload     bool
	tls        bool
}

func (a *action) Method() string {
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
	if a.wxml != nil {
		return a.wxml(appid, mchid, nonce)
	}

	return WXML{}, nil
}

func (a *action) Body() ([]byte, error) {
	if a.body != nil {
		return a.body()
	}

	return nil, nil
}

func (a *action) UploadForm() (yiigo.UploadForm, error) {
	if a.uploadform != nil {
		return a.uploadform()
	}

	return yiigo.NewUploadForm(), nil
}

// func (a *action) UploadForm() (yiigo.UploadForm, error) {
// 	if a.uploadfield == nil {
// 		return yiigo.NewUploadForm(), nil
// 	}

// 	body, err := a.Body()

// 	if err != nil {
// 		return nil, err
// 	}

// 	fields := []yiigo.UploadField{
// 		yiigo.WithFileField(a.uploadfield.FileField, a.uploadfield.Filename, body),
// 	}

// 	if len(a.uploadfield.MetaField) != 0 {
// 		fields = append(fields, yiigo.WithFormField(a.uploadfield.MetaField, a.uploadfield.Metadata))
// 	}

// 	return yiigo.NewUploadForm(fields...), nil
// }

func (a *action) Decode(resp []byte) error {
	if a.decode == nil {
		return a.decode(resp)
	}

	return nil
}

func (a *action) IsUpload() bool {
	return a.upload
}

func (a *action) IsTLS() bool {
	return a.tls
}

// NewAction returns a new action
func NewAction(method string, reqURL string, options ...ActionOption) Action {
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
	return NewAction(http.MethodGet, reqURL, options...)
}

// NewPostAction returns a new action with POST method
func NewPostAction(reqURL string, options ...ActionOption) Action {
	return NewAction(http.MethodPost, reqURL, options...)
}
