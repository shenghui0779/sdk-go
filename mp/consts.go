package mp

// auth
const (
	AccessTokenURL  = "https://api.weixin.qq.com/cgi-bin/token"
	Code2SessionURL = "https://api.weixin.qq.com/sns/jscode2session"
	PaidUnionURL    = "https://api.weixin.qq.com/wxa/getpaidunionid"
)

// msg
const (
	UniformMessageSendURL   = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/uniform_send"
	SubscribeMessageSendURL = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send"
	TemplateMessageSendURL  = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send"
	KFMessageSendURL        = "https://api.weixin.qq.com/cgi-bin/message/custom/send"
	SetTypingURL            = "https://api.weixin.qq.com/cgi-bin/message/custom/typing"
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

// plugin
const (
	PluginManageURL    = "https://api.weixin.qq.com/wxa/plugin"
	PluginDevManageURL = "https://api.weixin.qq.com/wxa/devplugin"
)

// security
const (
	ImageSecCheckURL   = "https://api.weixin.qq.com/wxa/img_sec_check"
	MediaCheckAsyncURL = "https://api.weixin.qq.com/wxa/media_check_async"
	MsgSecCheckURL     = "https://api.weixin.qq.com/wxa/msg_sec_check"
)

// image
const (
	AICropURL          = "https://api.weixin.qq.com/cv/img/aicrop"
	ScanQRCodeURL      = "https://api.weixin.qq.com/cv/img/qrcode"
	SuperreSolutionURL = "https://api.weixin.qq.com/cv/img/superresolution"
)

// ocr
const (
	OCRIDCardURL          = "https://api.weixin.qq.com/cv/ocr/idcard"
	OCRBankCardURL        = "https://api.weixin.qq.com/cv/ocr/bankcard"
	OCRPlateNumberURL     = "https://api.weixin.qq.com/cv/ocr/platenum"
	OCRDriverLicenseURL   = "https://api.weixin.qq.com/cv/ocr/drivinglicense"
	OCRVehicleLicenseURL  = "https://api.weixin.qq.com/cv/ocr/driving"
	OCRBusinessLicenseURL = "https://api.weixin.qq.com/cv/ocr/bizlicense"
	OCRPrintedTextURL     = "https://api.weixin.qq.com/cv/ocr/comm"
)

// other
const (
	InvokeServiceURL = "https://api.weixin.qq.com/wxa/servicemarket"
	SoterVerifyURL   = "https://api.weixin.qq.com/cgi-bin/soter/verify_signature"
	UserRiskRankURL  = "https://api.weixin.qq.com/wxa/getuserriskrank"
)
