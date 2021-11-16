package wx

import (
	"crypto/tls"
	"encoding/pem"
	"io/ioutil"
	"path/filepath"

	"golang.org/x/crypto/pkcs12"
)

func P12FileToCert(path, password string) (tls.Certificate, error) {
	fail := func(err error) (tls.Certificate, error) { return tls.Certificate{}, err }

	certPath, err := filepath.Abs(filepath.Clean(path))

	if err != nil {
		return fail(err)
	}

	p12, err := ioutil.ReadFile(certPath)

	if err != nil {
		return fail(err)
	}

	return pkcs12ToPem(p12, password)
}

func pkcs12ToPem(p12 []byte, password string) (tls.Certificate, error) {
	blocks, err := pkcs12.ToPEM(p12, password)

	if err != nil {
		return tls.Certificate{}, err
	}

	pemData := make([]byte, 0)

	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}

	// then use PEM data for tls to construct tls certificate:
	return tls.X509KeyPair(pemData, pemData)
}

// ActionOption configures how we set up the action
type ActionOption func(a *action)

// WithQuery specifies the `query` to Action.
func WithQuery(key, value string) ActionOption {
	return func(a *action) {
		a.query.Set(key, value)
	}
}

// WithBody specifies the `body` to Action.
func WithBody(f func() ([]byte, error)) ActionOption {
	return func(a *action) {
		a.body = f
	}
}

// WithWXML specifies the `wxml` to Action.
func WithWXML(f func(appid, mchid, nonce string) (WXML, error)) ActionOption {
	return func(a *action) {
		a.wxml = f
	}
}

// WithUploadField specifies the `upload field` to Action.
func WithUploadField(field *UploadField) ActionOption {
	return func(a *action) {
		a.uploadfield = field
	}
}

// WithDecode specifies the `decode` to Action.
func WithDecode(f func(resp []byte) error) ActionOption {
	return func(a *action) {
		a.decode = f
	}
}

// WithTLS specifies the `tls` to Action.
func WithTLS() ActionOption {
	return func(a *action) {
		a.tls = true
	}
}
