package wx

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/url"
	"path/filepath"
)

// HTTPMethod http request method
type HTTPMethod string

const (
	MethodGet    HTTPMethod = "GET"
	MethodPost   HTTPMethod = "POST"
	MethodUpload HTTPMethod = "UPLOAD"
)

// HTTPClient is the interface that do http request
type HTTPClient interface {
	// Get sends an HTTP get request
	Get(ctx context.Context, reqURL string, options ...HTTPOption) ([]byte, error)
	// Post sends an HTTP post request
	Post(ctx context.Context, reqURL string, body []byte, options ...HTTPOption) ([]byte, error)
	// Upload sends an HTTP post request with xml
	PostXML(ctx context.Context, reqURL string, body WXML, options ...HTTPOption) ([]byte, error)
	// Upload sends an HTTP post request for uploading media
	Upload(ctx context.Context, reqURL string, form UploadForm, options ...HTTPOption) ([]byte, error)
}

// UploadForm is the interface for http upload
type UploadForm interface {
	// FieldName returns field name for upload
	FieldName() string

	// FileName returns filename for upload
	FileName() string

	// ExtraFields returns extra fields for upload
	ExtraFields() map[string]string

	// Buffer returns the buffer of media
	Buffer() ([]byte, error)
}

type httpUploadForm struct {
	fieldname   string
	filename    string
	extraFields map[string]string
}

func (f *httpUploadForm) FieldName() string {
	return f.fieldname
}

func (f *httpUploadForm) FileName() string {
	return f.filename
}

func (f *httpUploadForm) ExtraFields() map[string]string {
	return f.extraFields
}

func (f *httpUploadForm) Buffer() ([]byte, error) {
	path, err := filepath.Abs(f.filename)

	if err != nil {
		return nil, err
	}

	return ioutil.ReadFile(path)
}

// NewUploadForm returns new upload form
func NewUploadForm(fieldname, filename string, extraFields map[string]string) UploadForm {
	return &httpUploadForm{
		fieldname:   fieldname,
		filename:    filename,
		extraFields: extraFields,
	}
}

// Action is the interface that handle wechat api
type Action interface {
	URL(accessToken ...string) string
	Method() HTTPMethod
	WXML(appid, mchid, nonce string) (WXML, error)
	Body() ([]byte, error)
	UploadForm() UploadForm
	Decode() func(resp []byte) error
	TLS() bool
}

// API is a Action implementation
type API struct {
	reqURL     string
	method     HTTPMethod
	query      url.Values
	wxml       func(appid, mchid, nonce string) (WXML, error)
	body       func() ([]byte, error)
	uploadForm UploadForm
	decode     func(resp []byte) error
	tls        bool
}

func (a *API) URL(accessToken ...string) string {
	if len(accessToken) != 0 {
		a.query.Set("access_token", accessToken[0])
	}

	if len(a.query) == 0 {
		return a.reqURL
	}

	return fmt.Sprintf("%s?%s", a.reqURL, a.query.Encode())
}

func (a *API) Method() HTTPMethod {
	return a.method
}

func (a *API) WXML(appid, mchid, nonce string) (WXML, error) {
	if a.wxml == nil {
		return WXML{}, nil
	}

	return a.wxml(appid, mchid, nonce)
}

func (a *API) Body() ([]byte, error) {
	if a.body == nil {
		return nil, nil
	}

	return a.body()
}

func (a *API) UploadForm() UploadForm {
	if a.uploadForm == nil {
		return new(httpUploadForm)
	}

	return a.uploadForm
}

func (a *API) Decode() func(resp []byte) error {
	return a.decode
}

func (a *API) TLS() bool {
	return a.tls
}

// APIOption configures how we set up the wechat API
type APIOption func(api *API)

// WithMethod specifies the `method` to API.
func WithMethod(method HTTPMethod) APIOption {
	return func(api *API) {
		api.method = method
	}
}

// WithQuery specifies the `query` to API.
func WithQuery(key, value string) APIOption {
	return func(api *API) {
		api.query.Set(key, value)
	}
}

// WithBody specifies the `body` to API.
func WithBody(f func() ([]byte, error)) APIOption {
	return func(api *API) {
		api.body = f
	}
}

// WithWXML specifies the `wxml` to API.
func WithWXML(f func(appid, mchid, nonce string) (WXML, error)) APIOption {
	return func(api *API) {
		api.wxml = f
	}
}

// WithUploadForm specifies the `upload form` to API.
func WithUploadForm(fieldname, filename string, extraFields map[string]string) APIOption {
	return func(api *API) {
		api.uploadForm = NewUploadForm(fieldname, filename, extraFields)
	}
}

// WithDecode specifies the `decode` to API.
func WithDecode(f func(resp []byte) error) APIOption {
	return func(api *API) {
		api.decode = f
	}
}

// WithTLS specifies the `tls` to API.
func WithTLS() APIOption {
	return func(api *API) {
		api.tls = true
	}
}

// NewAPI returns a new action
func NewAPI(reqURL string, options ...APIOption) Action {
	api := &API{
		reqURL: reqURL,
		query:  url.Values{},
	}

	for _, f := range options {
		f(api)
	}

	return api
}
