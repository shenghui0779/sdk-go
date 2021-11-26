package wx

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/shenghui0779/yiigo"
)

// Client is the interface that do http request
type Client interface {
	// Post sends an HTTP post request
	Do(ctx context.Context, method, reqURL string, body []byte, options ...yiigo.HTTPOption) ([]byte, error)

	// Upload sends an HTTP post request for uploading media
	Upload(ctx context.Context, reqURL string, form yiigo.UploadForm, options ...yiigo.HTTPOption) ([]byte, error)

	// SetHTTPClient set http client
	SetHTTPClient(c yiigo.HTTPClient)

	// SetLogger set logger
	SetLogger(l Logger)
}

type wxclient struct {
	client yiigo.HTTPClient
	logger Logger
}

func (c *wxclient) Do(ctx context.Context, method, reqURL string, body []byte, options ...yiigo.HTTPOption) ([]byte, error) {
	logData := &LogData{
		URL:    reqURL,
		Method: method,
		Body:   body,
	}

	now := time.Now().Local()

	defer func() {
		logData.Duration = time.Since(now)
		c.logger.Log(ctx, logData)
	}()

	resp, err := c.client.Do(ctx, method, reqURL, body, options...)

	if err != nil {
		logData.Error = err

		return nil, err
	}

	defer resp.Body.Close()

	logData.StatusCode = resp.StatusCode

	if resp.StatusCode >= http.StatusBadRequest {
		io.Copy(ioutil.Discard, resp.Body)

		return nil, fmt.Errorf("unexpected status %d", resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		logData.Error = err

		return nil, err
	}

	logData.Response = b

	return b, nil
}

func (c *wxclient) Upload(ctx context.Context, reqURL string, form yiigo.UploadForm, options ...yiigo.HTTPOption) ([]byte, error) {
	logData := &LogData{
		URL:    reqURL,
		Method: http.MethodPost,
	}

	now := time.Now().Local()

	defer func() {
		logData.Duration = time.Since(now)
		c.logger.Log(ctx, logData)
	}()

	resp, err := c.client.Upload(ctx, reqURL, form, options...)

	if err != nil {
		logData.Error = err

		return nil, err
	}

	defer resp.Body.Close()

	logData.StatusCode = resp.StatusCode

	if resp.StatusCode >= http.StatusBadRequest {
		io.Copy(ioutil.Discard, resp.Body)

		return nil, fmt.Errorf("unexpected status %d", resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		logData.Error = err

		return nil, err
	}

	logData.Response = b

	return b, nil
}

func (c *wxclient) SetHTTPClient(client yiigo.HTTPClient) {
	c.client = client
}

func (c *wxclient) SetLogger(l Logger) {
	c.logger = l
}

// DefaultClient returns a new default wechat client
func DefaultClient(certs ...tls.Certificate) Client {
	tlscfg := &tls.Config{
		InsecureSkipVerify: true,
	}

	if len(certs) != 0 {
		tlscfg.Certificates = certs
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 60 * time.Second,
			}).DialContext,
			TLSClientConfig:       tlscfg,
			MaxIdleConns:          0,
			MaxIdleConnsPerHost:   1000,
			MaxConnsPerHost:       1000,
			IdleConnTimeout:       60 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}

	return &wxclient{
		client: yiigo.NewHTTPClient(client),
		logger: DefaultLogger(),
	}
}
