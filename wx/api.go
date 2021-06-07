package wx

import (
	"context"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
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
	// Write writes fields to multipart writer
	Write(w *multipart.Writer) error
}

type httpUpload struct {
	filefield string
	filename  string
	filebytes []byte
	metafield string
	metadata  string
	err       error
}

func (u *httpUpload) Write(w *multipart.Writer) error {
	if u.err != nil {
		return u.err
	}

	part, err := w.CreateFormFile(u.filefield, u.filename)

	if err != nil {
		return err
	}

	if _, err = part.Write(u.filebytes); err != nil {
		return err
	}

	// metadata
	if len(u.metafield) != 0 {
		if err = w.WriteField(u.metafield, u.metadata); err != nil {
			return err
		}
	}

	return nil
}

// UploadOption configures how we set up the upload from.
type UploadOption func(u *httpUpload)

// UploadByPath uploads by file path
func UploadByPath(path string) UploadOption {
	return func(u *httpUpload) {
		path, err := filepath.Abs(path)

		if err != nil {
			u.err = fmt.Errorf("read content by path error: %s", err.Error())

			return
		}

		b, err := ioutil.ReadFile(path)

		if err != nil {
			u.err = fmt.Errorf("read content by path error: %s", err.Error())

			return
		}

		u.filebytes = b
	}
}

// UploadByContent uploads by file content
func UploadByContent(content []byte) UploadOption {
	return func(u *httpUpload) {
		u.filebytes = content
	}
}

// UploadByResourceURL uploads file by resource url
func UploadByResourceURL(url string) UploadOption {
	return func(u *httpUpload) {
		resp, err := http.Get(url)

		if err != nil {
			u.err = fmt.Errorf("get content by resource url error: %s", err.Error())

			return
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			u.err = fmt.Errorf("error http code from resource url: %d", resp.StatusCode)

			return
		}

		b, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			u.err = fmt.Errorf("get content by resource url error: %s", err.Error())

			return
		}

		u.filebytes = b
	}
}

// WithMetaField specifies the metadata field to upload from.
func WithMetaField(name, value string) UploadOption {
	return func(u *httpUpload) {
		u.metafield = name
		u.metadata = value
	}
}

// NewUploadForm returns an upload form
func NewUploadForm(fieldname, filename string, options ...UploadOption) UploadForm {
	form := &httpUpload{
		filefield: fieldname,
		filename:  filename,
	}

	for _, f := range options {
		f(form)
	}

	return form
}

// Action is the interface that handle wechat api
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
	UploadForm() UploadForm

	// Decode decodes response
	Decode() func(resp []byte) error

	// TLS specifies the request with certificate
	TLS() bool
}

type wxapi struct {
	reqURL     string
	method     HTTPMethod
	query      url.Values
	wxml       func(appid, mchid, nonce string) (WXML, error)
	body       func() ([]byte, error)
	uploadForm UploadForm
	decode     func(resp []byte) error
	tls        bool
}

func (a *wxapi) URL(accessToken ...string) string {
	if len(accessToken) != 0 {
		a.query.Set("access_token", accessToken[0])
	}

	if len(a.query) == 0 {
		return a.reqURL
	}

	return fmt.Sprintf("%s?%s", a.reqURL, a.query.Encode())
}

func (a *wxapi) Method() HTTPMethod {
	return a.method
}

func (a *wxapi) WXML(appid, mchid, nonce string) (WXML, error) {
	if a.wxml == nil {
		return WXML{}, nil
	}

	return a.wxml(appid, mchid, nonce)
}

func (a *wxapi) Body() ([]byte, error) {
	if a.body == nil {
		return nil, nil
	}

	return a.body()
}

func (a *wxapi) UploadForm() UploadForm {
	if a.uploadForm == nil {
		return new(httpUpload)
	}

	return a.uploadForm
}

func (a *wxapi) Decode() func(resp []byte) error {
	return a.decode
}

func (a *wxapi) TLS() bool {
	return a.tls
}

// ActionOption configures how we set up the action
type ActionOption func(api *wxapi)

// WithMethod specifies the `method` to Action.
func WithMethod(method HTTPMethod) ActionOption {
	return func(api *wxapi) {
		api.method = method
	}
}

// WithQuery specifies the `query` to Action.
func WithQuery(key, value string) ActionOption {
	return func(api *wxapi) {
		api.query.Set(key, value)
	}
}

// WithBody specifies the `body` to Action.
func WithBody(f func() ([]byte, error)) ActionOption {
	return func(api *wxapi) {
		api.body = f
	}
}

// WithWXML specifies the `wxml` to Action.
func WithWXML(f func(appid, mchid, nonce string) (WXML, error)) ActionOption {
	return func(api *wxapi) {
		api.wxml = f
	}
}

// WithUploadForm specifies the `upload form` to Action.
func WithUploadForm(form UploadForm) ActionOption {
	return func(api *wxapi) {
		api.uploadForm = form
	}
}

// WithDecode specifies the `decode` to Action.
func WithDecode(f func(resp []byte) error) ActionOption {
	return func(api *wxapi) {
		api.decode = f
	}
}

// WithTLS specifies the `tls` to Action.
func WithTLS() ActionOption {
	return func(api *wxapi) {
		api.tls = true
	}
}

// NewAction returns a new action
func NewAction(reqURL string, options ...ActionOption) Action {
	api := &wxapi{
		reqURL: reqURL,
		query:  url.Values{},
	}

	for _, f := range options {
		f(api)
	}

	return api
}
