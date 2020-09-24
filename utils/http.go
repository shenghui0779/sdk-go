package utils

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"
)

// defaultTimeout default http request timeout
const defaultTimeout = 10 * time.Second

// WXML deal with xml for wechat
type WXML map[string]string

// tlsOptions https tls options
type tlsOptions struct {
	rootCAs            [][]byte
	certificates       []tls.Certificate
	insecureSkipVerify bool
}

// TLSOption configures how we set up the https transport
type TLSOption interface {
	apply(*tlsOptions)
}

// funcTLSOption implements tls option
type funcTLSOption struct {
	f func(*tlsOptions)
}

func (fo *funcTLSOption) apply(o *tlsOptions) {
	fo.f(o)
}

func newFuncTLSOption(f func(*tlsOptions)) *funcTLSOption {
	return &funcTLSOption{f: f}
}

// WithRootCA specifies the `RootCAs` to https transport.
func WithRootCA(crt []byte) TLSOption {
	return newFuncTLSOption(func(o *tlsOptions) {
		o.rootCAs = append(o.rootCAs, crt)
	})
}

// WithInsecureSkipVerify specifies the `Certificates` to https transport.
func WithCertificates(certs ...tls.Certificate) TLSOption {
	return newFuncTLSOption(func(o *tlsOptions) {
		o.certificates = certs
	})
}

// WithInsecureSkipVerify specifies the `InsecureSkipVerify` to https transport.
func WithInsecureSkipVerify(b bool) TLSOption {
	return newFuncTLSOption(func(o *tlsOptions) {
		o.insecureSkipVerify = b
	})
}

// requestOptions http request options
type requestOptions struct {
	headers map[string]string
	cookies []*http.Cookie
	close   bool
	timeout time.Duration
}

// RequestOption configures how we set up the http request
type RequestOption interface {
	apply(*requestOptions)
}

// funcRequestOption implements request option
type funcRequestOption struct {
	f func(*requestOptions)
}

func (fo *funcRequestOption) apply(o *requestOptions) {
	fo.f(o)
}

func newFuncRequestOption(f func(*requestOptions)) *funcRequestOption {
	return &funcRequestOption{f: f}
}

// WithHeader specifies the headers to http request.
func WithHeader(key, value string) RequestOption {
	return newFuncRequestOption(func(o *requestOptions) {
		o.headers[key] = value
	})
}

// WithCookies specifies the cookies to http request.
func WithCookies(cookies ...*http.Cookie) RequestOption {
	return newFuncRequestOption(func(o *requestOptions) {
		o.cookies = cookies
	})
}

// WithClose specifies close the connection after
// replying to this request (for servers) or after sending this
// request and reading its response (for clients).
func WithClose(b bool) RequestOption {
	return newFuncRequestOption(func(o *requestOptions) {
		o.close = b
	})
}

// WithTimeout specifies the timeout to http request.
func WithTimeout(d time.Duration) RequestOption {
	return newFuncRequestOption(func(o *requestOptions) {
		o.timeout = d
	})
}

// WXClient http client
type WXClient struct {
	client  *http.Client
	timeout time.Duration
}

// Get http get request
func (c *WXClient) Get(reqURL string, options ...RequestOption) ([]byte, error) {
	o := &requestOptions{
		headers: make(map[string]string),
		timeout: c.timeout,
	}

	if len(options) > 0 {
		for _, option := range options {
			option.apply(o)
		}
	}

	req, err := http.NewRequest("GET", reqURL, nil)

	if err != nil {
		return nil, err
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

	ctx, cancel := context.WithTimeout(req.Context(), o.timeout)

	defer cancel()

	resp, err := c.client.Do(req.WithContext(ctx))

	if err != nil {
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

// Post http post request
func (c *WXClient) Post(reqURL string, body []byte, options ...RequestOption) ([]byte, error) {
	o := &requestOptions{
		headers: make(map[string]string),
		timeout: c.timeout,
	}

	if len(options) > 0 {
		for _, option := range options {
			option.apply(o)
		}
	}

	req, err := http.NewRequest("POST", reqURL, bytes.NewReader(body))

	if err != nil {
		return nil, err
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

	ctx, cancel := context.WithTimeout(req.Context(), o.timeout)

	defer cancel()

	resp, err := c.client.Do(req.WithContext(ctx))

	if err != nil {
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

// GetXML http xml get request
func (c *WXClient) GetXML(reqURL string, body WXML, options ...RequestOption) (WXML, error) {
	query := url.Values{}

	for k, v := range body {
		query.Add(k, v)
	}

	resp, err := c.Get(fmt.Sprintf("%s?%s", reqURL, query.Encode()), options...)

	if err != nil {
		return nil, err
	}

	wxml, err := ParseXML2Map(resp)

	if err != nil {
		return nil, err
	}

	return wxml, nil
}

// PostXML http xml post request
func (c *WXClient) PostXML(reqURL string, body WXML, options ...RequestOption) (WXML, error) {
	xmlStr, err := FormatMap2XML(body)

	options = append(options, WithHeader("Content-Type", "text/xml; charset=utf-8"))

	resp, err := c.Post(reqURL, []byte(xmlStr), options...)

	if err != nil {
		return nil, err
	}

	wxml, err := ParseXML2Map(resp)

	if err != nil {
		return nil, err
	}

	return wxml, nil
}

// NewWXClient returns a new http client
func NewWXClient(options ...TLSOption) *WXClient {
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

	if len(options) > 0 {
		o := &tlsOptions{
			rootCAs: make([][]byte, 0),
		}

		for _, option := range options {
			option.apply(o)
		}

		tlsCfg := new(tls.Config)

		if len(o.rootCAs) > 0 {
			pool := x509.NewCertPool()

			for _, b := range o.rootCAs {
				pool.AppendCertsFromPEM(b)
			}

			tlsCfg.RootCAs = pool
		} else {
			if o.insecureSkipVerify {
				tlsCfg.InsecureSkipVerify = true
			}
		}

		if len(o.certificates) > 0 {
			tlsCfg.Certificates = o.certificates
		}

		t.TLSClientConfig = tlsCfg
	}

	return &WXClient{
		client: &http.Client{
			Transport: t,
		},
		timeout: defaultTimeout,
	}
}

// FormatMap2XML format map to xml
func FormatMap2XML(m WXML) (string, error) {
	buf := BufPool.Get()
	defer BufPool.Put(buf)

	if _, err := io.WriteString(buf, "<xml>"); err != nil {
		return "", err
	}

	for k, v := range m {
		if _, err := io.WriteString(buf, fmt.Sprintf("<%s>", k)); err != nil {
			return "", err
		}

		if err := xml.EscapeText(buf, []byte(v)); err != nil {
			return "", err
		}

		if _, err := io.WriteString(buf, fmt.Sprintf("</%s>", k)); err != nil {
			return "", err
		}
	}

	if _, err := io.WriteString(buf, "</xml>"); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// ParseXML2Map parse xml to map
func ParseXML2Map(b []byte) (WXML, error) {
	m := make(WXML)

	xmlReader := bytes.NewReader(b)

	var (
		d     = xml.NewDecoder(xmlReader)
		tk    xml.Token
		depth = 0 // current xml.Token depth
		key   string
		buf   bytes.Buffer
		err   error
	)

	for {
		tk, err = d.Token()

		if err != nil {
			if err == io.EOF {
				return m, nil
			}

			return nil, err
		}

		switch v := tk.(type) {
		case xml.StartElement:
			depth++

			switch depth {
			case 2:
				key = v.Name.Local
				buf.Reset()
			case 3:
				if err = d.Skip(); err != nil {
					return nil, err
				}

				depth--
				key = "" // key == "" indicates that the node with depth==2 has children
			}
		case xml.CharData:
			if depth == 2 && key != "" {
				buf.Write(v)
			}
		case xml.EndElement:
			if depth == 2 && key != "" {
				m[key] = buf.String()
			}

			depth--
		}
	}
}
