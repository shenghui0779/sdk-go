package wx

import (
	"fmt"
	"net/http"
	"net/url"
)

// Action is the interface that handle wechat api
type Action interface {
	// Method returns action method
	Method() string

	// URL returns request url
	URL(accessToken ...string) string

	// WXML returns body for xml request
	WXML(mchid, apikey, nonce string) (WXML, error)

	// Body returns body for post request
	Body() ([]byte, error)

	// UploadForm returns form for uploading media
	UploadForm() (UploadForm, error)

	// Decode decodes response
	Decode(b []byte) error

	// IsUpload specifies the request does upload
	IsUpload() bool

	// TLS specifies the request with certificate
	IsTLS() bool
}

type action struct {
	method     string
	reqURL     string
	query      url.Values
	wxml       func(mchid, apikey, nonce string) (WXML, error)
	body       func() ([]byte, error)
	uploadform func() (UploadForm, error)
	decode     func(b []byte) error
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

func (a *action) WXML(mchid, apikey, nonce string) (WXML, error) {
	if a.wxml != nil {
		return a.wxml(mchid, apikey, nonce)
	}

	return WXML{}, nil
}

func (a *action) Body() ([]byte, error) {
	if a.body != nil {
		return a.body()
	}

	return nil, nil
}

func (a *action) UploadForm() (UploadForm, error) {
	if a.uploadform != nil {
		return a.uploadform()
	}

	return NewUploadForm(), nil
}

func (a *action) Decode(b []byte) error {
	if a.decode != nil {
		return a.decode(b)
	}

	return nil
}

func (a *action) IsUpload() bool {
	return a.upload
}

func (a *action) IsTLS() bool {
	return a.tls
}

// ActionOption configures how we set up the action
type ActionOption func(a *action)

// WithQuery sets query params for action.
func WithQuery(key, value string) ActionOption {
	return func(a *action) {
		a.query.Set(key, value)
	}
}

// WithBody sets post body for action.
func WithBody(f func() ([]byte, error)) ActionOption {
	return func(a *action) {
		a.body = f
	}
}

// WithWXML sets post with xml for action.
func WithWXML(f func(mchid, apikey, nonce string) (WXML, error)) ActionOption {
	return func(a *action) {
		a.wxml = f
	}
}

// WithUpload sets uploadform for action.
func WithUpload(f func() (UploadForm, error)) ActionOption {
	return func(a *action) {
		a.upload = true
		a.uploadform = f
	}
}

// WithDecode sets response decode for action.
func WithDecode(f func(b []byte) error) ActionOption {
	return func(a *action) {
		a.decode = f
	}
}

// WithTLS sets request with tls for action.
func WithTLS() ActionOption {
	return func(a *action) {
		a.tls = true
	}
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
