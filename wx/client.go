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
	"go.uber.org/zap"
)

// Client is the interface that do http request
type Client interface {
	// Post sends an HTTP post request
	Do(ctx context.Context, method, reqURL string, body []byte, options ...yiigo.HTTPOption) ([]byte, error)

	// Post sends an HTTP post request
	DoXML(ctx context.Context, method, reqURL string, body WXML, options ...yiigo.HTTPOption) ([]byte, error)

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

func (c *wxclient) DoXML(ctx context.Context, method, reqURL string, body WXML, options ...yiigo.HTTPOption) ([]byte, error) {
	logData := &LogData{
		URL:    reqURL,
		Method: method,
	}

	now := time.Now().Local()

	defer func() {
		logData.Duration = time.Since(now)
		c.logger.Log(ctx, logData)
	}()

	reqBody, err := FormatMap2XML(body)

	if err != nil {
		logData.Error = err

		return nil, err
	}

	logData.Body = reqBody

	options = append(options, yiigo.WithHTTPHeader("Content-Type", "text/xml; charset=utf-8"))

	resp, err := c.client.Do(ctx, method, reqURL, reqBody, options...)

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

type Logger interface {
	Log(ctx context.Context, data *LogData)
}

type LogData struct {
	URL        string        `json:"url"`
	Method     string        `json:"method"`
	Body       []byte        `json:"body"`
	StatusCode int           `json:"status_code"`
	Response   []byte        `json:"response"`
	Duration   time.Duration `json:"duration"`
	Error      error         `json:"error"`
}

type wxlogger struct{}

func (l *wxlogger) Log(ctx context.Context, data *LogData) {
	fields := make([]zap.Field, 0, 7)

	fields = append(fields,
		zap.String("method", data.Method),
		zap.String("url", data.URL),
		zap.ByteString("body", data.Body),
		zap.ByteString("response", data.Response),
		zap.Int("status", data.StatusCode),
		zap.String("duration", data.Duration.String()),
	)

	if data.Error != nil {
		fields = append(fields, zap.Error(data.Error))

		yiigo.Logger().Error("[gochat] action do error", fields...)

		return
	}

	yiigo.Logger().Info("[gochat] action do info", fields...)
}

// DefaultLogger returns default logger
func DefaultLogger() Logger {
	return new(wxlogger)
}
