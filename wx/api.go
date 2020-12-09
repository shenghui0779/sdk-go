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
	MethodGet    = "GET"
	MethodPost   = "POST"
	MethodUpload = "UPLOAD"
)

// Client is the interface that do http request
type Client interface {
	Get(ctx context.Context, reqURL string, options ...HTTPOption) ([]byte, error)
	Post(ctx context.Context, reqURL string, body []byte, options ...HTTPOption) ([]byte, error)
	PostXML(ctx context.Context, reqURL string, body WXML, options ...HTTPOption) (WXML, error)
	Upload(ctx context.Context, reqURL string, form *UploadForm, options ...HTTPOption) ([]byte, error)
}

// UploadForm http upload form
type UploadForm struct {
	fieldname   string
	filename    string
	extraFields map[string]string
}

func (f *UploadForm) FieldName() string {
	return f.fieldname
}

func (f *UploadForm) FileName() string {
	return f.filename
}

func (f *UploadForm) ExtraFields() map[string]string {
	return f.extraFields
}

func (f *UploadForm) Buffer() ([]byte, error) {
	path, err := filepath.Abs(f.filename)

	if err != nil {
		return nil, err
	}

	return ioutil.ReadFile(path)
}

// NewUploadForm returns new uplod form
func NewUploadForm(fieldname, filename string, extraFields map[string]string) *UploadForm {
	return &UploadForm{
		fieldname:   fieldname,
		filename:    filename,
		extraFields: extraFields,
	}
}

// HTTPBody is a Body implementation
type HTTPBody struct {
	wxml       func(appid, mchid, nonce string) (WXML, error)
	bytes      func() ([]byte, error)
	uploadForm *UploadForm
}

func (h *HTTPBody) WXML(appid, mchid, nonce string) (WXML, error) {
	return h.wxml(appid, mchid, nonce)
}

func (h *HTTPBody) Bytes() ([]byte, error) {
	if h.bytes == nil {
		return nil, nil
	}

	return h.bytes()
}

func (h *HTTPBody) UploadForm() *UploadForm {
	return h.uploadForm
}

// Action is the interface that handle wechat api
type Action interface {
	URL(accessToken ...string) string
	Method() HTTPMethod
	Body() *HTTPBody
	Decode() func(resp []byte) error
	TLS() bool
}

// API is a Action implementation
type API struct {
	reqURL string
	method HTTPMethod
	query  url.Values
	body   *HTTPBody
	decode func(resp []byte) error
	tls    bool
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

func (a *API) Body() *HTTPBody {
	return a.body
}

func (a *API) Decode() func(resp []byte) error {
	return a.decode
}

func (a *API) TLS() bool {
	return a.tls
}

// NewAction returns a new action
func NewAction(reqURL string, method HTTPMethod, query url.Values, body *HTTPBody, decode func(resp []byte) error, tls bool) Action {
	return &API{
		reqURL: reqURL,
		method: method,
		query:  query,
		body:   body,
		decode: decode,
		tls:    tls,
	}
}

// NewMchAPI returns mch action
func NewMchAPI(reqURL string, f func(appid, mchid, nonce string) (WXML, error)) Action {
	return NewAction(reqURL, MethodPost, url.Values{}, &HTTPBody{wxml: f}, nil, false)
}

// NewMchTLSAPI return mch action
func NewMchTLSAPI(reqURL string, f func(appid, mchid, nonce string) (WXML, error)) Action {
	return NewAction(reqURL, MethodPost, url.Values{}, &HTTPBody{wxml: f}, nil, true)
}

// NewGetAPI returns get action
func NewGetAPI(reqURL string, query url.Values, decode func(resp []byte) error) Action {
	return NewAction(reqURL, MethodGet, query, nil, decode, false)
}

// NewPostAPI returns post action
func NewPostAPI(reqURL string, query url.Values, body func() ([]byte, error), decode func(resp []byte) error) Action {
	return NewAction(reqURL, MethodPost, query, &HTTPBody{bytes: body}, decode, false)
}

// NewUploadAPI returns upload action
func NewUploadAPI(reqURL string, query url.Values, form *UploadForm, decode func(resp []byte) error) Action {
	return NewAction(reqURL, MethodUpload, query, &HTTPBody{uploadForm: form}, decode, false)
}
