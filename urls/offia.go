package urls

// oauth
const (
	Oauth2Authorize = "https://open.weixin.qq.com/connect/oauth2/authorize"
	QRCodeAuthorize = "https://open.work.weixin.qq.com/wwopen/sso/qrConnect"
)

// cgi-bin
const (
	OffiaCgiBinAccessToken = "https://api.weixin.qq.com/cgi-bin/token"
	OffiaCgiBinTicket      = "https://api.weixin.qq.com/cgi-bin/ticket/getticket"
)

// menu
const (
	OffiaMenuCreate            = "https://api.weixin.qq.com/cgi-bin/menu/create"
	OffiaMenuAddConditional    = "https://api.weixin.qq.com/cgi-bin/menu/addconditional"
	OffiaMenuTryMatch          = "https://api.weixin.qq.com/cgi-bin/menu/trymatch"
	OffiaMenuList              = "https://api.weixin.qq.com/cgi-bin/menu/get"
	OffiaMenuDelete            = "https://api.weixin.qq.com/cgi-bin/menu/delete"
	OffiaMenuDeleteConditional = "https://api.weixin.qq.com/cgi-bin/menu/delconditional"
)

// sns
const (
	OffiaSnsCode2Token         = "https://api.weixin.qq.com/sns/oauth2/access_token"
	OffiaSnsCheckAccessToken   = "https://api.weixin.qq.com/sns/auth"
	OffiaSnsRefreshAccessToken = "https://api.weixin.qq.com/sns/oauth2/refresh_token"
	OffiaSnsUserInfo           = "https://api.weixin.qq.com/sns/userinfo"
)

// subscriber
const (
	OffiaSubscriberGet      = "https://api.weixin.qq.com/cgi-bin/user/info"
	OffiaSubscriberBatchGet = "https://api.weixin.qq.com/cgi-bin/user/info/batchget"
	OffiaSubscriberList     = "https://api.weixin.qq.com/cgi-bin/user/get"
	OffiaBlackListGet       = "https://api.weixin.qq.com/cgi-bin/tags/members/getblacklist"
	OffiaBatchBlackList     = "https://api.weixin.qq.com/cgi-bin/tags/members/batchblacklist"
	OffiaBatchUnBlackList   = "https://api.weixin.qq.com/cgi-bin/tags/members/batchunblacklist"
	OffiaUserRemarkSet      = "https://api.weixin.qq.com/cgi-bin/user/info/updateremark"
)

// message
const (
	OffiaTemplateList         = "https://api.weixin.qq.com/cgi-bin/template/get_all_private_template"
	OffiaTemplateDelete       = "https://api.weixin.qq.com/cgi-bin/template/del_private_template"
	OffiaTemplateMessageSend  = "https://api.weixin.qq.com/cgi-bin/message/template/send"
	OffiaSubscribeMessageSend = "https://api.weixin.qq.com/cgi-bin/message/template/subscribe"
)

// popularize
const (
	OffiaQRCodeCreate     = "https://api.weixin.qq.com/cgi-bin/qrcode/create"
	OffiaQRCodeShow       = "https://mp.weixin.qq.com/cgi-bin/showqrcode"
	OffiaShortURLGenerate = "https://api.weixin.qq.com/cgi-bin/shorturl"
)

// media
const (
	OffiaMediaUpload     = "https://api.weixin.qq.com/cgi-bin/media/upload"
	OffiaMediaGet        = "https://api.weixin.qq.com/cgi-bin/media/get"
	OffiaNewsAdd         = "https://api.weixin.qq.com/cgi-bin/material/add_news"
	OffiaNewUpdate     = "https://api.weixin.qq.com/cgi-bin/material/update_news"
	OffiaNewsImageUpload = "https://api.weixin.qq.com/cgi-bin/media/uploadimg"
	OffiaMaterialAdd     = "https://api.weixin.qq.com/cgi-bin/material/add_material"
	OffiaMaterialDelete  = "https://api.weixin.qq.com/cgi-bin/material/del_material"
	OffiaMaterialGet     = "https://api.weixin.qq.com/cgi-bin/material/get_material"
)

// image
const (
	OffiaAICrop          = "https://api.weixin.qq.com/cv/img/aicrop"
	OffiaScanQRCode      = "https://api.weixin.qq.com/cv/img/qrcode"
	OffiaSuperreSolution = "https://api.weixin.qq.com/cv/img/superresolution"
)

// ocr
const (
	OffiaOCRIDCard          = "https://api.weixin.qq.com/cv/ocr/idcard"
	OffiaOCRBankCard        = "https://api.weixin.qq.com/cv/ocr/bankcard"
	OffiaOCRPlateNumber     = "https://api.weixin.qq.com/cv/ocr/platenum"
	OffiaOCRDriverLicense   = "https://api.weixin.qq.com/cv/ocr/drivinglicense"
	OffiaOCRVehicleLicense  = "https://api.weixin.qq.com/cv/ocr/driving"
	OffiaOCRBusinessLicense = "https://api.weixin.qq.com/cv/ocr/bizlicense"
	OffiaOCRPrintedText     = "https://api.weixin.qq.com/cv/ocr/comm"
)

// KF
const (
	OffiaKFAccountList   = "https://api.weixin.qq.com/cgi-bin/customservice/getkflist"
	OffiaKFOnlineList    = "https://api.weixin.qq.com/cgi-bin/customservice/getonlinekflist"
	OffiaKFAccountAdd    = "https://api.weixin.qq.com/customservice/kfaccount/add"
	OffiaKFInvite        = "https://api.weixin.qq.com/customservice/kfaccount/inviteworker"
	OffiaKFAccountUpdate = "https://api.weixin.qq.com/customservice/kfaccount/update"
	OffiaKFAvatarUpload  = "https://api.weixin.qq.com/customservice/kfaccount/uploadheadimg"
	OffiaKFDelete        = "https://api.weixin.qq.com/customservice/kfaccount/del"
	OffiaKFSessionCreate = "https://api.weixin.qq.com/customservice/kfsession/create"
	OffiaKFSessionClose  = "https://api.weixin.qq.com/customservice/kfsession/close"
	OffiaKFSessionGet    = "https://api.weixin.qq.com/customservice/kfsession/getsession"
	OffiaKFSessionList   = "https://api.weixin.qq.com/customservice/kfsession/getsessionlist"
	OffiaKFWaitCase      = "https://api.weixin.qq.com/customservice/kfsession/getwaitcase"
	OffiaKFMsgRecordList = "https://api.weixin.qq.com/customservice/msgrecord/getmsglist"
	OffiaKFMessageSend   = "https://api.weixin.qq.com/cgi-bin/message/custom/send"
	OffiaSetTyping       = "https://api.weixin.qq.com/cgi-bin/message/custom/typing"
)
