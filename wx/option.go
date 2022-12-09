package wx

// ActionOption configures how we set up the action
type ActionOption func(a *action)

// WithQuery sets query params for action.
func WithQuery(key, value string) ActionOption {
	return func(a *action) {
		a.query.Set(key, value)
	}
}

// WithBody sets post body for action.
func WithBody(f func() ([]byte, error)) ActionOption {
	return func(a *action) {
		a.body = f
	}
}

// WithWXML sets post with xml for action.
func WithWXML(f func(mchid, apikey, nonce string) (WXML, error)) ActionOption {
	return func(a *action) {
		a.wxml = f
	}
}

// WithUpload sets uploadform for action.
func WithUpload(f func() (UploadForm, error)) ActionOption {
	return func(a *action) {
		a.upload = true
		a.uploadform = f
	}
}

// WithDecode sets response decode for action.
func WithDecode(f func(b []byte) error) ActionOption {
	return func(a *action) {
		a.decode = f
	}
}

// WithTLS sets request with tls for action.
func WithTLS() ActionOption {
	return func(a *action) {
		a.tls = true
	}
}
