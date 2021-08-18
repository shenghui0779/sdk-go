package urls

// oauth
const OAAuthorize = "https://open.weixin.qq.com/connect/oauth2/authorize"

// cgi-bin
const (
	OACgiBinAccessToken = "https://api.weixin.qq.com/cgi-bin/token"
	OACgiBinTicket      = "https://api.weixin.qq.com/cgi-bin/ticket/getticket"
)

// menu
const (
	OAMenuCreate            = "https://api.weixin.qq.com/cgi-bin/menu/create"
	OAMenuAddConditional    = "https://api.weixin.qq.com/cgi-bin/menu/addconditional"
	OAMenuTryMatch          = "https://api.weixin.qq.com/cgi-bin/menu/trymatch"
	OAMenuList              = "https://api.weixin.qq.com/cgi-bin/menu/get"
	OAMenuDelete            = "https://api.weixin.qq.com/cgi-bin/menu/delete"
	OAMenuDeleteConditional = "https://api.weixin.qq.com/cgi-bin/menu/delconditional"
)

// sns
const (
	OASnsCode2Token         = "https://api.weixin.qq.com/sns/oauth2/access_token"
	OASnsCheckAccessToken   = "https://api.weixin.qq.com/sns/auth"
	OASnsRefreshAccessToken = "https://api.weixin.qq.com/sns/oauth2/refresh_token"
	OASnsUserInfo           = "https://api.weixin.qq.com/sns/userinfo"
)

// subscriber
const (
	OASubscriberGet      = "https://api.weixin.qq.com/cgi-bin/user/info"
	OASubscriberBatchGet = "https://api.weixin.qq.com/cgi-bin/user/info/batchget"
	OASubscriberList     = "https://api.weixin.qq.com/cgi-bin/user/get"
	OABlackListGet       = "https://api.weixin.qq.com/cgi-bin/tags/members/getblacklist"
	OABatchBlackList     = "https://api.weixin.qq.com/cgi-bin/tags/members/batchblacklist"
	OABatchUnBlackList   = "https://api.weixin.qq.com/cgi-bin/tags/members/batchunblacklist"
	OAUserRemarkSet      = "https://api.weixin.qq.com/cgi-bin/user/info/updateremark"
)

// message
const (
	OATemplateList         = "https://api.weixin.qq.com/cgi-bin/template/get_all_private_template"
	OATemplateDelete       = "https://api.weixin.qq.com/cgi-bin/template/del_private_template"
	OATemplateMessageSend  = "https://api.weixin.qq.com/cgi-bin/message/template/send"
	OASubscribeMessageSend = "https://api.weixin.qq.com/cgi-bin/message/template/subscribe"
)

// popularize
const (
	OAQRCodeCreate     = "https://api.weixin.qq.com/cgi-bin/qrcode/create"
	OAQRCodeShow       = "https://mp.weixin.qq.com/cgi-bin/showqrcode"
	OAShortURLGenerate = "https://api.weixin.qq.com/cgi-bin/shorturl"
)

// media
const (
	OAMediaUpload     = "https://api.weixin.qq.com/cgi-bin/media/upload"
	OAMediaGet        = "https://api.weixin.qq.com/cgi-bin/media/get"
	OANewsAdd         = "https://api.weixin.qq.com/cgi-bin/material/add_news"
	OANewsImageUpload = "https://api.weixin.qq.com/cgi-bin/media/uploadimg"
	OAMaterialAdd     = "https://api.weixin.qq.com/cgi-bin/material/add_material"
	OAMaterialDelete  = "https://api.weixin.qq.com/cgi-bin/material/del_material"
)

// image
const (
	OAAICrop          = "https://api.weixin.qq.com/cv/img/aicrop"
	OAScanQRCode      = "https://api.weixin.qq.com/cv/img/qrcode"
	OASuperreSolution = "https://api.weixin.qq.com/cv/img/superresolution"
)

// ocr
const (
	OAOCRIDCard          = "https://api.weixin.qq.com/cv/ocr/idcard"
	OAOCRBankCard        = "https://api.weixin.qq.com/cv/ocr/bankcard"
	OAOCRPlateNumber     = "https://api.weixin.qq.com/cv/ocr/platenum"
	OAOCRDriverLicense   = "https://api.weixin.qq.com/cv/ocr/drivinglicense"
	OAOCRVehicleLicense  = "https://api.weixin.qq.com/cv/ocr/driving"
	OAOCRBusinessLicense = "https://api.weixin.qq.com/cv/ocr/bizlicense"
	OAOCRPrintedText     = "https://api.weixin.qq.com/cv/ocr/comm"
)

// KF
const (
	OAKFAccountList   = "https://api.weixin.qq.com/cgi-bin/customservice/getkflist"
	OAKFOnlineList    = "https://api.weixin.qq.com/cgi-bin/customservice/getonlinekflist"
	OAKFAccountAdd    = "https://api.weixin.qq.com/customservice/kfaccount/add"
	OAKFInvite        = "https://api.weixin.qq.com/customservice/kfaccount/inviteworker"
	OAKFAccountUpdate = "https://api.weixin.qq.com/customservice/kfaccount/update"
	OAKFAvatarUpload  = "https://api.weixin.qq.com/customservice/kfaccount/uploadheadimg"
	OAKFDelete        = "https://api.weixin.qq.com/customservice/kfaccount/del"
	OAKFSessionCreate = "https://api.weixin.qq.com/customservice/kfsession/create"
	OAKFSessionClose  = "https://api.weixin.qq.com/customservice/kfsession/close"
	OAKFSessionGet    = "https://api.weixin.qq.com/customservice/kfsession/getsession"
	OAKFSessionList   = "https://api.weixin.qq.com/customservice/kfsession/getsessionlist"
	OAKFWaitCase      = "https://api.weixin.qq.com/customservice/kfsession/getwaitcase"
	OAKFMsgRecordList = "https://api.weixin.qq.com/customservice/msgrecord/getmsglist"
	OAKFMessageSend   = "https://api.weixin.qq.com/cgi-bin/message/custom/send"
	OASetTyping       = "https://api.weixin.qq.com/cgi-bin/message/custom/typing"
)
