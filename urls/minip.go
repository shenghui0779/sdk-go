package consts

// auth
const (
	MinipAccessTokenURL  = "https://api.weixin.qq.com/cgi-bin/token"
	MinipCode2SessionURL = "https://api.weixin.qq.com/sns/jscode2session"
	MinipPaidUnionURL    = "https://api.weixin.qq.com/wxa/getpaidunionid"
)

// msg
const (
	MinipUniformMessageSendURL   = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/uniform_send"
	MinipSubscribeMessageSendURL = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send"
	MinipTemplateMessageSendURL  = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send"
	MinipKFMessageSendURL        = "https://api.weixin.qq.com/cgi-bin/message/custom/send"
	MinipSetTypingURL            = "https://api.weixin.qq.com/cgi-bin/message/custom/typing"
)

// qrcode
const (
	MinipQRCodeCreateURL     = "https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode"
	MinipQRCodeGetURL        = "https://api.weixin.qq.com/wxa/getwxacode"
	MinipQRCodeGetUnlimitURL = "https://api.weixin.qq.com/wxa/getwxacodeunlimit"
)

// media
const (
	MinipMediaUploadURL = "https://api.weixin.qq.com/cgi-bin/media/upload"
	MinipMediaGetURL    = "https://api.weixin.qq.com/cgi-bin/media/get"
)

// plugin
const (
	MinipPluginManageURL    = "https://api.weixin.qq.com/wxa/plugin"
	MinipPluginDevManageURL = "https://api.weixin.qq.com/wxa/devplugin"
)

// security
const (
	MinipImageSecCheckURL   = "https://api.weixin.qq.com/wxa/img_sec_check"
	MinipMediaCheckAsyncURL = "https://api.weixin.qq.com/wxa/media_check_async"
	MinipMsgSecCheckURL     = "https://api.weixin.qq.com/wxa/msg_sec_check"
)

// image
const (
	MinipAICropURL          = "https://api.weixin.qq.com/cv/img/aicrop"
	MinipScanQRCodeURL      = "https://api.weixin.qq.com/cv/img/qrcode"
	MinipSuperreSolutionURL = "https://api.weixin.qq.com/cv/img/superresolution"
)

// ocr
const (
	MinipOCRIDCardURL          = "https://api.weixin.qq.com/cv/ocr/idcard"
	MinipOCRBankCardURL        = "https://api.weixin.qq.com/cv/ocr/bankcard"
	MinipOCRPlateNumberURL     = "https://api.weixin.qq.com/cv/ocr/platenum"
	MinipOCRDriverLicenseURL   = "https://api.weixin.qq.com/cv/ocr/drivinglicense"
	MinipOCRVehicleLicenseURL  = "https://api.weixin.qq.com/cv/ocr/driving"
	MinipOCRBusinessLicenseURL = "https://api.weixin.qq.com/cv/ocr/bizlicense"
	MinipOCRPrintedTextURL     = "https://api.weixin.qq.com/cv/ocr/comm"
)

// other
const (
	MinipInvokeServiceURL = "https://api.weixin.qq.com/wxa/servicemarket"
	MinipSoterVerifyURL   = "https://api.weixin.qq.com/cgi-bin/soter/verify_signature"
	MinipUserRiskRankURL  = "https://api.weixin.qq.com/wxa/getuserriskrank"
)
