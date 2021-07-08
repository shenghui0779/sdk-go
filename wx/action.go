package wx

import (
	"fmt"
	"net/url"

	"github.com/shenghui0779/yiigo"
)

// Action is the interface that handle wechat a
type Action interface {
	// URL returns request url
	URL(accessToken ...string) string

	// Method returns request method
	Method() HTTPMethod

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
	reqURL      string
	method      HTTPMethod
	query       url.Values
	wxml        func(appid, mchid, nonce string) (WXML, error)
	body        func() ([]byte, error)
	uploadfield *UploadField
	decode      func(resp []byte) error
	tls         bool
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

func (a *action) Method() HTTPMethod {
	return a.method
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

	fields := []yiigo.UploadOption{
		yiigo.WithFieldField(a.uploadfield.FileField, a.uploadfield.Filename, body),
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

// ActionOption configures how we set up the action
type ActionOption func(a *action)

// WithMethod specifies the `method` to Action.
func WithMethod(method HTTPMethod) ActionOption {
	return func(a *action) {
		a.method = method
	}
}

// WithQuery specifies the `query` to Action.
func WithQuery(key, value string) ActionOption {
	return func(a *action) {
		a.query.Set(key, value)
	}
}

// WithBody specifies the `body` to Action.
func WithBody(f func() ([]byte, error)) ActionOption {
	return func(a *action) {
		a.body = f
	}
}

// WithWXML specifies the `wxml` to Action.
func WithWXML(f func(appid, mchid, nonce string) (WXML, error)) ActionOption {
	return func(a *action) {
		a.wxml = f
	}
}

// WithUploadField specifies the `upload field` to Action.
func WithUploadField(field *UploadField) ActionOption {
	return func(a *action) {
		a.uploadfield = field
	}
}

// WithDecode specifies the `decode` to Action.
func WithDecode(f func(resp []byte) error) ActionOption {
	return func(a *action) {
		a.decode = f
	}
}

// WithTLS specifies the `tls` to Action.
func WithTLS() ActionOption {
	return func(a *action) {
		a.tls = true
	}
}

// NewAction returns a new action
func NewAction(reqURL string, options ...ActionOption) Action {
	a := &action{
		reqURL: reqURL,
		query:  url.Values{},
	}

	for _, f := range options {
		f(a)
	}

	return a
}
