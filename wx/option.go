package wx

// ClientOption configures how we set up the wechat client.
type ClientOption func(c *wxclient)

// WithHTTPClient sets http client for wechat client.
func WithHTTPClient(client HTTPClient) ClientOption {
	return func(c *wxclient) {
		c.client = client
	}
}

// WithLogger sets logger for wechat client.
func WithLogger(logger Logger) ClientOption {
	return func(c *wxclient) {
		c.logger = logger
	}
}

// WithDebug sets debug mode for wechat client.
func WithDebug() ClientOption {
	return func(c *wxclient) {
		c.debug = true
	}
}

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
func WithWXML(f func(mchid, nonce string) (WXML, error)) ActionOption {
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
func WithDecode(f func(resp []byte) error) ActionOption {
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
