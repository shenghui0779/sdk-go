package wx

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/url"
	"path/filepath"
)

// WXML deal with xml for wechat
type WXML map[string]string

// HTTPMethod http request method
type HTTPMethod string

const (
	MethodGet    = "GET"
	MethodPost   = "POST"
	MethodUpload = "UPLOAD"
)

// Client is the interface that do http request
type Client interface {
	// Get do http get request
	Get(ctx context.Context, reqURL string, options ...HTTPOption) ([]byte, error)
	// Post do http post request
	Post(ctx context.Context, reqURL string, body Body, options ...HTTPOption) ([]byte, error)
	// PostXML do http xml request
	PostXML(ctx context.Context, reqURL string, body WXML, options ...HTTPOption) (WXML, error)
	// Upload do http upload request
	Upload(ctx context.Context, reqURL string, body Body, options ...HTTPOption) ([]byte, error)
}

// Body is the interface that defines request body
type Body interface {
	// FieldName returns field name for upload request
	FieldName() string
	// FieldName returns file name for upload request
	FileName() string
	// Description returns file description for upload request
	Description() X
	// Bytes returns body for post request
	Bytes() func() ([]byte, error)
}

// Action is the interface that handle wechat api
type Action interface {
	// URL returns url for request
	URL() func(accessToken ...string) string
	// Method returns request method
	Method() HTTPMethod
	// WXML returns body for xml request
	WXML() func(appid, mchid, apikey, nonce string) (WXML, error)
	// Body returns body for post and upload request
	Body() Body
	// Decode decodes http response
	Decode() func(resp []byte) error
	// TLS specifies use tls client
	TLS() bool
}

// HTTPBody is a Body implementation for http request body
type HTTPBody struct {
	fieldname   string
	filename    string
	description X
	bytes       func() ([]byte, error)
}

func (h *HTTPBody) FieldName() string {
	return h.fieldname
}

func (h *HTTPBody) FileName() string {
	return h.filename
}

func (h *HTTPBody) Description() X {
	return h.description
}

func (h *HTTPBody) Bytes() func() ([]byte, error) {
	return h.bytes
}

// NewPostBody returns post body
func NewPostBody(f func() ([]byte, error)) *HTTPBody {
	return &HTTPBody{bytes: f}
}

// NewUploadBody returns upload body
func NewUploadBody(fieldname, filename string, description X) *HTTPBody {
	return &HTTPBody{
		fieldname:   fieldname,
		filename:    filename,
		description: description,
		bytes: func() ([]byte, error) {
			path, err := filepath.Abs(fieldname)

			if err != nil {
				return nil, err
			}

			return ioutil.ReadFile(path)
		},
	}
}

// API is a Action implementation for handle wechat api
type API struct {
	reqURL func(accessToken ...string) string
	method HTTPMethod
	wxml   func(appid, mchid, apikey, nonce string) (WXML, error)
	body   Body
	decode func(resp []byte) error
	tls    bool
}

func (a *API) URL() func(accessToken ...string) string {
	return a.reqURL
}

func (a *API) Method() HTTPMethod {
	return a.method
}

func (a *API) WXML() func(appid, mchid, apikey, nonce string) (WXML, error) {
	return a.wxml
}

func (a *API) Body() Body {
	return a.body
}

func (a *API) Decode() func(resp []byte) error {
	return a.decode
}

func (a *API) TLS() bool {
	return a.tls
}

// NewMchAPI returns mch api
func NewMchAPI(reqURL string, wxml func(appid, mchid, apikey, nonce string) (WXML, error), tls bool) *API {
	return &API{
		reqURL: func(accessToken ...string) string {
			return reqURL
		},
		method: MethodPost,
		wxml:   wxml,
		tls:    tls,
	}
}

// NewOpenAPI returns open api
func NewOpenAPI(reqURL string, method HTTPMethod, query url.Values, body Body, decode func(resp []byte) error) *API {
	return &API{
		reqURL: func(accessToken ...string) string {
			if len(accessToken) != 0 {
				query.Set("access_token", accessToken[0])
			}

			return fmt.Sprintf("%s?%s", reqURL, query.Encode())
		},
		method: method,
		body:   body,
		decode: decode,
	}
}

// NewOpenGetAPI returns open get api
func NewOpenGetAPI(reqURL string, query url.Values, decode func(resp []byte) error) *API {
	return NewOpenAPI(reqURL, MethodGet, query, nil, decode)
}

// NewOpenPostAPI returns open post api
func NewOpenPostAPI(reqURL string, query url.Values, body Body, decode func(resp []byte) error) *API {
	return NewOpenAPI(reqURL, MethodPost, query, body, decode)
}

// NewOpenUploadAPI returns open upload api
func NewOpenUploadAPI(reqURL string, query url.Values, body Body, decode func(resp []byte) error) *API {
	return NewOpenAPI(reqURL, MethodUpload, query, body, decode)
}
