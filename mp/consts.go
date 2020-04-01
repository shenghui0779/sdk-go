package mp

// cgi-bin
const (
	AccessTokenURL = "https://api.weixin.qq.com/cgi-bin/token"
)

// sns
const (
	Code2SessionURL = "https://api.weixin.qq.com/sns/jscode2session"
)

// msg
const (
	UniformMsgSendURL         = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/uniform_send"
	SubscribeMsgSendURL       = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send"
	TplMsgSendURL             = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send"
	CustomerServiceMsgSendURL = "https://api.weixin.qq.com/cgi-bin/message/custom/send"
	SetTypingURL              = "https://api.weixin.qq.com/cgi-bin/message/custom/typing"
)

// qrcode
const (
	QRCodeCreateURL     = "https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode"
	QRCodeGetURL        = "https://api.weixin.qq.com/wxa/getwxacode"
	QRCodeGetUnlimitURL = "https://api.weixin.qq.com/wxa/getwxacodeunlimit"
)

// media
const (
	MediaUploadURL = "https://api.weixin.qq.com/cgi-bin/media/upload"
	MediaGetURL    = "https://api.weixin.qq.com/cgi-bin/media/get"
)
