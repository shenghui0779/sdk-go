package urls

// auth
const (
	MinipAccessToken        = "https://api.weixin.qq.com/cgi-bin/token"
	MinipCode2Session       = "https://api.weixin.qq.com/sns/jscode2session"
	MinipEncryptedDataCheck = "https://api.weixin.qq.com/wxa/business/checkencryptedmsg"
	MinipPaidUnion          = "https://api.weixin.qq.com/wxa/getpaidunionid"
)

// message
const (
	MinipUniformMsgSend   = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/uniform_send"
	MinipSubscribeMsgSend = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send"
	MinipKFMsgSend        = "https://api.weixin.qq.com/cgi-bin/message/custom/send"
	MinipKFTypingSend     = "https://api.weixin.qq.com/cgi-bin/message/custom/typing"
)

// qrcode
const (
	MinipQRCodeCreate     = "https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode"
	MinipQRCodeGet        = "https://api.weixin.qq.com/wxa/getwxacode"
	MinipQRCodeGetUnlimit = "https://api.weixin.qq.com/wxa/getwxacodeunlimit"
)

// media
const (
	MinipMediaUpload = "https://api.weixin.qq.com/cgi-bin/media/upload"
	MinipMediaGet    = "https://api.weixin.qq.com/cgi-bin/media/get"
)

// plugin
const (
	MinipPluginManage    = "https://api.weixin.qq.com/wxa/plugin"
	MinipPluginDevManage = "https://api.weixin.qq.com/wxa/devplugin"
)

// security
const (
	MinipImageSecCheck   = "https://api.weixin.qq.com/wxa/img_sec_check"
	MinipMediaCheckAsync = "https://api.weixin.qq.com/wxa/media_check_async"
	MinipMsgSecCheck     = "https://api.weixin.qq.com/wxa/msg_sec_check"
)

// image
const (
	MinipAICrop          = "https://api.weixin.qq.com/cv/img/aicrop"
	MinipScanQRCode      = "https://api.weixin.qq.com/cv/img/qrcode"
	MinipSuperreSolution = "https://api.weixin.qq.com/cv/img/superresolution"
)

// ocr
const (
	MinipOCRIDCard          = "https://api.weixin.qq.com/cv/ocr/idcard"
	MinipOCRBankCard        = "https://api.weixin.qq.com/cv/ocr/bankcard"
	MinipOCRPlateNumber     = "https://api.weixin.qq.com/cv/ocr/platenum"
	MinipOCRDriverLicense   = "https://api.weixin.qq.com/cv/ocr/drivinglicense"
	MinipOCRVehicleLicense  = "https://api.weixin.qq.com/cv/ocr/driving"
	MinipOCRBusinessLicense = "https://api.weixin.qq.com/cv/ocr/bizlicense"
	MinipOCRComm            = "https://api.weixin.qq.com/cv/ocr/comm"
)

// other
const (
	MinipInvokeService = "https://api.weixin.qq.com/wxa/servicemarket"
	MinipSoterVerify   = "https://api.weixin.qq.com/cgi-bin/soter/verify_signature"
	MinipShortLink     = "https://api.weixin.qq.com/wxa/genwxashortlink"
	MinipUserRiskRank  = "https://api.weixin.qq.com/wxa/getuserriskrank"
)

// subscribe
const (
	MinipAddTemplate                  = "https://api.weixin.qq.com/wxaapi/newtmpl/addtemplate"
	MinipDeleteTemplate               = "https://api.weixin.qq.com/wxaapi/newtmpl/deltemplate"
	MinipGetCategory                  = "https://api.weixin.qq.com/wxaapi/newtmpl/getcategory"
	MinipGetetPubTemplateKeyWordsByID = "https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatekeywords"
	MinipGetPubTemplateTitleList      = "https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatetitles"
	MinipGetTemplateList              = "https://api.weixin.qq.com/wxaapi/newtmpl/gettemplate"
)
