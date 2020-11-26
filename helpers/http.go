package helpers

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"time"
)

// defaultTimeout default http request timeout
const defaultTimeout = 10 * time.Second

// httpOptions http request options
type httpOptions struct {
	headers map[string]string
	cookies []*http.Cookie
	close   bool
	timeout time.Duration
}

// HTTPOption configures how we set up the http request
type HTTPOption interface {
	apply(*httpOptions)
}

// funcHTTPOption implements request option
type funcHTTPOption struct {
	f func(*httpOptions)
}

func (fo *funcHTTPOption) apply(o *httpOptions) {
	fo.f(o)
}

func newFuncHTTPOption(f func(*httpOptions)) *funcHTTPOption {
	return &funcHTTPOption{f: f}
}

// WithHTTPHeader specifies the headers to http request.
func WithHTTPHeader(key, value string) HTTPOption {
	return newFuncHTTPOption(func(o *httpOptions) {
		o.headers[key] = value
	})
}

// WithHTTPCookies specifies the cookies to http request.
func WithHTTPCookies(cookies ...*http.Cookie) HTTPOption {
	return newFuncHTTPOption(func(o *httpOptions) {
		o.cookies = cookies
	})
}

// WithHTTPClose specifies close the connection after
// replying to this request (for servers) or after sending this
// request and reading its response (for clients).
func WithHTTPClose() HTTPOption {
	return newFuncHTTPOption(func(o *httpOptions) {
		o.close = true
	})
}

// WithHTTPTimeout specifies the timeout to http request.
func WithHTTPTimeout(d time.Duration) HTTPOption {
	return newFuncHTTPOption(func(o *httpOptions) {
		o.timeout = d
	})
}

// HTTPBody is the interface that defines http body
type HTTPBody interface {
	FieldName() string
	FileName() string
	Bytes() func() ([]byte, error)
}

// PostBody is a HTTPBody implementation for wechat http request body
type PostBody struct {
	fieldname string
	filename  string
	bytes     func() ([]byte, error)
}

// FieldName returns multipart fieldname
func (b *PostBody) FieldName() string {
	return b.fieldname
}

// FileName returns multipart filename
func (b *PostBody) FileName() string {
	return b.filename
}

// Bytes returns bytes closure
func (b *PostBody) Bytes() func() ([]byte, error) {
	return b.bytes
}

// NewPostBody returns post body
func NewPostBody(f func() ([]byte, error)) HTTPBody {
	return &PostBody{bytes: f}
}

// NewUploadBody returns upload body
func NewUploadBody(fieldname, filename string, f func() ([]byte, error)) HTTPBody {
	return &PostBody{
		fieldname: fieldname,
		filename:  fieldname,
		bytes:     f,
	}
}

// HTTPClient is the interface that wraps http request
type HTTPClient interface {
	Get(ctx context.Context, url string, options ...HTTPOption) ([]byte, error)
	Post(ctx context.Context, url string, body HTTPBody, options ...HTTPOption) ([]byte, error)
	PostXML(ctx context.Context, url string, body WXML, options ...HTTPOption) (WXML, error)
	Upload(ctx context.Context, url string, body HTTPBody, options ...HTTPOption) ([]byte, error)
}

// WXClient is a HTTPClient implementation for wechat http request
type WXClient struct {
	client  *http.Client
	timeout time.Duration
}

func (c *WXClient) do(ctx context.Context, req *http.Request, options ...HTTPOption) ([]byte, error) {
	o := &httpOptions{
		headers: make(map[string]string),
		timeout: c.timeout,
	}

	if len(options) > 0 {
		for _, option := range options {
			option.apply(o)
		}
	}

	// headers
	if len(o.headers) > 0 {
		for k, v := range o.headers {
			req.Header.Set(k, v)
		}
	}

	// cookies
	if len(o.cookies) > 0 {
		for _, v := range o.cookies {
			req.AddCookie(v)
		}
	}

	if o.close {
		req.Close = true
	}

	// timeout
	ctx, cancel := context.WithTimeout(ctx, o.timeout)

	defer cancel()

	resp, err := c.client.Do(req.WithContext(ctx))

	if err != nil {
		// If the context has been canceled, the context's error is probably more useful.
		select {
		case <-ctx.Done():
			err = ctx.Err()
		default:
		}

		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		io.Copy(ioutil.Discard, resp.Body)

		return nil, fmt.Errorf("error http code: %d", resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return b, nil
}

// Get http get request
func (c *WXClient) Get(ctx context.Context, url string, options ...HTTPOption) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	return c.do(ctx, req, options...)
}

// Post http post request
func (c *WXClient) Post(ctx context.Context, url string, body HTTPBody, options ...HTTPOption) ([]byte, error) {
	var (
		b   []byte
		err error
	)

	if f := body.Bytes(); f != nil {
		b, err = f()

		if err != nil {
			return nil, err
		}
	}

	options = append(options, WithHTTPHeader("Content-Type", "application/json; charset=utf-8"))

	req, err := http.NewRequest("POST", url, bytes.NewReader(b))

	if err != nil {
		return nil, err
	}

	return c.do(ctx, req, options...)
}

// PostXML http xml post request
func (c *WXClient) PostXML(ctx context.Context, url string, body WXML, options ...HTTPOption) (WXML, error) {
	xmlStr, err := FormatMap2XML(body)

	options = append(options, WithHTTPHeader("Content-Type", "text/xml; charset=utf-8"))

	req, err := http.NewRequest("POST", url, bytes.NewReader([]byte(xmlStr)))

	if err != nil {
		return nil, err
	}

	resp, err := c.do(ctx, req, options...)

	if err != nil {
		return nil, err
	}

	wxml, err := ParseXML2Map(resp)

	if err != nil {
		return nil, err
	}

	return wxml, nil
}

// Upload http upload media
func (c *WXClient) Upload(ctx context.Context, url string, body HTTPBody, options ...HTTPOption) ([]byte, error) {
	media, err := body.Bytes()()

	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	mw := multipart.NewWriter(buf)

	fw, err := mw.CreateFormFile(body.FieldName(), body.FileName())

	if err != nil {
		return nil, err
	}

	if _, err = io.Copy(fw, bytes.NewReader(media)); err != nil {
		return nil, err
	}

	options = append(options, WithHTTPHeader("Content-Type", mw.FormDataContentType()))

	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	mw.Close()

	req, err := http.NewRequest("POST", url, bytes.NewReader([]byte(buf.String())))

	if err != nil {
		return nil, err
	}

	return c.do(ctx, req, options...)
}

// NewHTTPClient returns a new http client
func NewHTTPClient(tlsCfg ...*tls.Config) HTTPClient {
	t := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 60 * time.Second,
		}).DialContext,
		MaxIdleConns:          0,
		MaxIdleConnsPerHost:   1000,
		MaxConnsPerHost:       1000,
		IdleConnTimeout:       60 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	if len(tlsCfg) != 0 {
		t.TLSClientConfig = tlsCfg[0]
	}

	return &WXClient{
		client: &http.Client{
			Transport: t,
		},
		timeout: defaultTimeout,
	}
}
