package wx

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
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

type wxclient struct {
	client   yiigo.HTTPClient
	insecure bool
	certs    []tls.Certificate
	logger   Logger
}

// Get http get request
func (c *wxclient) Get(ctx context.Context, reqURL string, options ...yiigo.HTTPOption) ([]byte, error) {
	logData := &LogData{
		URL:    reqURL,
		Method: http.MethodGet,
	}

	now := time.Now().Local()

	defer func() {
		logData.Duration = time.Since(now)
		c.logger.Log(ctx, logData)
	}()

	resp, err := c.client.Do(ctx, http.MethodGet, reqURL, nil, options...)

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

// Post http post request
func (c *wxclient) Post(ctx context.Context, reqURL string, body []byte, options ...yiigo.HTTPOption) ([]byte, error) {
	logData := &LogData{
		URL:    reqURL,
		Method: http.MethodPost,
		Body:   body,
	}

	now := time.Now().Local()

	defer func() {
		logData.Duration = time.Since(now)
		c.logger.Log(ctx, logData)
	}()

	options = append(options, yiigo.WithHTTPHeader("Content-Type", "application/json; charset=utf-8"))

	resp, err := c.client.Do(ctx, http.MethodPost, reqURL, bytes.NewReader(body), options...)

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

// PostXML http xml post request
func (c *wxclient) PostXML(ctx context.Context, reqURL string, body WXML, options ...yiigo.HTTPOption) ([]byte, error) {
	logData := &LogData{
		URL:    reqURL,
		Method: http.MethodPost,
	}

	now := time.Now().Local()

	defer func() {
		logData.Duration = time.Since(now)
		c.logger.Log(ctx, logData)
	}()

	xmlStr, err := FormatMap2XML(body)

	if err != nil {
		logData.Error = err

		return nil, err
	}

	logData.Body = []byte(xmlStr)

	options = append(options, yiigo.WithHTTPHeader("Content-Type", "text/xml; charset=utf-8"))

	resp, err := c.client.Do(ctx, http.MethodPost, reqURL, strings.NewReader(xmlStr), options...)

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

// Upload http upload media
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

// NewClient returns a new http client
func NewClient(options ...ClientOption) Client {
	client := &wxclient{
		logger: new(wxlogger),
	}

	for _, f := range options {
		f(client)
	}

	t := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 60 * time.Second,
		}).DialContext,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		MaxIdleConns:          0,
		MaxIdleConnsPerHost:   1000,
		MaxConnsPerHost:       1000,
		IdleConnTimeout:       60 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	if client.insecure || len(client.certs) != 0 {
		t.TLSClientConfig = &tls.Config{
			Certificates:       client.certs,
			InsecureSkipVerify: client.insecure,
		}
	}

	return &wxclient{
		client: yiigo.NewHTTPClient(&http.Client{
			Transport: t,
		}),
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
	fields := make([]zap.Field, 0, 5)

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
	}

	if data.Error != nil {
		yiigo.Logger().Error("[gochat] action do error", fields...)
	}

	yiigo.Logger().Info("[gochat] action do info", fields...)
}
