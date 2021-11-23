package wx

import "github.com/shenghui0779/yiigo"

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

// WithUpload specifies the `uploadform` to Action.
func WithUpload(f func() (yiigo.UploadForm, error)) ActionOption {
	return func(a *action) {
		a.upload = true
		a.uploadform = f
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
