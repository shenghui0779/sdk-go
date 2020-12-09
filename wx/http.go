package wx

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
	"strings"
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

// HTTPClient is a Client implementation for wechat http request
type HTTPClient struct {
	client  *http.Client
	timeout time.Duration
}

func (h *HTTPClient) do(ctx context.Context, req *http.Request, options ...HTTPOption) ([]byte, error) {
	o := &httpOptions{
		headers: make(map[string]string),
		timeout: h.timeout,
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

	resp, err := h.client.Do(req.WithContext(ctx))

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
func (h *HTTPClient) Get(ctx context.Context, url string, options ...HTTPOption) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	return h.do(ctx, req, options...)
}

// Post http post request
func (h *HTTPClient) Post(ctx context.Context, url string, body []byte, options ...HTTPOption) ([]byte, error) {
	options = append(options, WithHTTPHeader("Content-Type", "application/json; charset=utf-8"))

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))

	if err != nil {
		return nil, err
	}

	return h.do(ctx, req, options...)
}

// PostXML http xml post request
func (h *HTTPClient) PostXML(ctx context.Context, url string, body WXML, options ...HTTPOption) (WXML, error) {
	xmlStr, err := FormatMap2XML(body)

	if err != nil {
		return nil, err
	}

	options = append(options, WithHTTPHeader("Content-Type", "text/xml; charset=utf-8"))

	req, err := http.NewRequest("POST", url, strings.NewReader(xmlStr))

	if err != nil {
		return nil, err
	}

	resp, err := h.do(ctx, req, options...)

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
func (h *HTTPClient) Upload(ctx context.Context, url string, form *UploadForm, options ...HTTPOption) ([]byte, error) {
	media, err := form.Buffer()

	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(make([]byte, 0, 4<<10)) // 4kb
	w := multipart.NewWriter(buf)

	fw, err := w.CreateFormFile(form.FieldName(), form.FileName())

	if err != nil {
		return nil, err
	}

	if _, err = io.Copy(fw, bytes.NewReader(media)); err != nil {
		return nil, err
	}

	// add extra fields
	if extraFields := form.ExtraFields(); len(extraFields) != 0 {
		for k, v := range extraFields {
			if err = w.WriteField(k, v); err != nil {
				return nil, err
			}
		}
	}

	options = append(options, WithHTTPHeader("Content-Type", w.FormDataContentType()))

	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()

	req, err := http.NewRequest("POST", url, buf)

	if err != nil {
		return nil, err
	}

	return h.do(ctx, req, options...)
}

// NewHTTPClient returns a new http client
func NewHTTPClient(tlsCfg ...*tls.Config) *HTTPClient {
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

	return &HTTPClient{
		client: &http.Client{
			Transport: t,
		},
		timeout: defaultTimeout,
	}
}
