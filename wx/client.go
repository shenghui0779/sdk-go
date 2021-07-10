package wx

import (
	"bytes"
	"context"
	"crypto/tls"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/shenghui0779/yiigo"
)

// HTTPMethod http request method
type HTTPMethod string

const (
	MethodGet    HTTPMethod = "GET"
	MethodPost   HTTPMethod = "POST"
	MethodUpload HTTPMethod = "UPLOAD"
)

// Client is the interface that do http request
type Client interface {
	// Get sends an HTTP get request
	Get(ctx context.Context, reqURL string, options ...yiigo.HTTPOption) ([]byte, error)

	// Post sends an HTTP post request
	Post(ctx context.Context, reqURL string, body []byte, options ...yiigo.HTTPOption) ([]byte, error)

	// PostXML sends an HTTP post request with xml
	PostXML(ctx context.Context, reqURL string, body WXML, options ...yiigo.HTTPOption) ([]byte, error)

	// Upload sends an HTTP post request for uploading media
	Upload(ctx context.Context, reqURL string, form yiigo.UploadForm, options ...yiigo.HTTPOption) ([]byte, error)
}

type apiclient struct {
	client yiigo.HTTPClient
}

// Get http get request
func (c *apiclient) Get(ctx context.Context, reqURL string, options ...yiigo.HTTPOption) ([]byte, error) {
	resp, err := c.client.Do(ctx, http.MethodGet, reqURL, nil, options...)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// Post http post request
func (c *apiclient) Post(ctx context.Context, reqURL string, body []byte, options ...yiigo.HTTPOption) ([]byte, error) {
	options = append(options, yiigo.WithHTTPHeader("Content-Type", "application/json; charset=utf-8"))

	resp, err := c.client.Do(ctx, http.MethodPost, reqURL, bytes.NewReader(body), options...)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// PostXML http xml post request
func (c *apiclient) PostXML(ctx context.Context, reqURL string, body WXML, options ...yiigo.HTTPOption) ([]byte, error) {
	xmlStr, err := FormatMap2XML(body)

	if err != nil {
		return nil, err
	}

	options = append(options, yiigo.WithHTTPHeader("Content-Type", "text/xml; charset=utf-8"))

	resp, err := c.client.Do(ctx, http.MethodPost, reqURL, strings.NewReader(xmlStr), options...)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// Upload http upload media
func (c *apiclient) Upload(ctx context.Context, reqURL string, form yiigo.UploadForm, options ...yiigo.HTTPOption) ([]byte, error) {
	resp, err := c.client.Upload(ctx, reqURL, form, options...)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// TLSOption configures how we set up the http client tls config
type TLSOption func(cfg *tls.Config)

// WithInsecureSkipVerify specifies the `InsecureSkipVerify` to http client tls config.
func WithInsecureSkipVerify() TLSOption {
	return func(cfg *tls.Config) {
		cfg.InsecureSkipVerify = true
	}
}

// WithTLSCertificates specifies the certificate to http client tls config.
func WithTLSCertificates(certs ...tls.Certificate) TLSOption {
	return func(cfg *tls.Config) {
		cfg.Certificates = certs
	}
}

// NewClient returns a new http client
func NewClient(options ...TLSOption) Client {
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

	if len(options) != 0 {
		t.TLSClientConfig = new(tls.Config)

		for _, f := range options {
			f(t.TLSClientConfig)
		}
	}

	return &apiclient{
		client: yiigo.NewHTTPClient(&http.Client{
			Transport: t,
		}),
	}
}
