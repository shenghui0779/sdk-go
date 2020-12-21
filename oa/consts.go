package oa

// AuthScope 应用授权作用域
type AuthScope string

// 公众号支持的应用授权作用域
const (
	ScopeSnsapiBase AuthScope = "snsapi_base"     // 静默授权使用，不弹出授权页面，直接跳转，只能获取用户openid
	ScopeSnsapiUser AuthScope = "snsapi_userinfo" // 弹出授权页面，可通过openid拿到昵称、性别、所在地。并且，即使在未关注的情况下，只要用户授权，也能获取其信息
)

// oauth2
const AuthorizeURL = "https://open.weixin.qq.com/connect/oauth2/authorize"

// cgi-bin
const (
	CgiBinAccessTokenURL = "https://api.weixin.qq.com/cgi-bin/token"
	CgiBinTicketURL      = "https://api.weixin.qq.com/cgi-bin/ticket/getticket"
)

// menu
const (
	MenuCreateURL            = "https://api.weixin.qq.com/cgi-bin/menu/create"
	MenuAddConditionalURL    = "https://api.weixin.qq.com/cgi-bin/menu/addconditional"
	MenuTryMatchURL          = "https://api.weixin.qq.com/cgi-bin/menu/trymatch"
	MenuListURL              = "https://api.weixin.qq.com/cgi-bin/menu/get"
	MenuDeleteURL            = "https://api.weixin.qq.com/cgi-bin/menu/delete"
	MenuDeleteConditionalURL = "https://api.weixin.qq.com/cgi-bin/menu/delconditional"
)

// sns
const (
	SnsCode2TokenURL         = "https://api.weixin.qq.com/sns/oauth2/access_token"
	SnsCheckAccessTokenURL   = "https://api.weixin.qq.com/sns/auth"
	SnsRefreshAccessTokenURL = "https://api.weixin.qq.com/sns/oauth2/refresh_token"
	SnsUserInfoURL           = "https://api.weixin.qq.com/sns/userinfo"
)

// subscriber
const (
	SubscriberGetURL      = "https://api.weixin.qq.com/cgi-bin/user/info"
	SubscriberBatchGetURL = "https://api.weixin.qq.com/cgi-bin/user/info/batchget"
	SubscriberListURL     = "https://api.weixin.qq.com/cgi-bin/user/get"
	BlackListGetURL       = "https://api.weixin.qq.com/cgi-bin/tags/members/getblacklist"
	BatchBlackListURL     = "https://api.weixin.qq.com/cgi-bin/tags/members/batchblacklist"
	BatchUnBlackListURL   = "https://api.weixin.qq.com/cgi-bin/tags/members/batchunblacklist"
	UserRemarkSetURL      = "https://api.weixin.qq.com/cgi-bin/user/info/updateremark"
)

// message
const (
	TemplateListURL         = "https://api.weixin.qq.com/cgi-bin/template/get_all_private_template"
	TemplateDeleteURL       = "https://api.weixin.qq.com/cgi-bin/template/del_private_template"
	TemplateMessageSendURL  = "https://api.weixin.qq.com/cgi-bin/message/template/send"
	SubscribeMessageSendURL = "https://api.weixin.qq.com/cgi-bin/message/template/subscribe"
)

// popularize
const (
	QRCodeCreateURL     = "https://api.weixin.qq.com/cgi-bin/qrcode/create"
	QRCodeShowURL       = "https://mp.weixin.qq.com/cgi-bin/showqrcode"
	ShortURLGenerateURL = "https://api.weixin.qq.com/cgi-bin/shorturl"
)

// media
const (
	MediaUploadURL     = "https://api.weixin.qq.com/cgi-bin/media/upload"
	MediaGetURL        = "https://api.weixin.qq.com/cgi-bin/media/get"
	NewsAddURL         = "https://api.weixin.qq.com/cgi-bin/material/add_news"
	NewsImageUploadURL = "https://api.weixin.qq.com/cgi-bin/media/uploadimg"
	MaterialAddURL     = "https://api.weixin.qq.com/cgi-bin/material/add_material"
	MaterialDeleteURL  = "https://api.weixin.qq.com/cgi-bin/material/del_material"
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

// KF
const (
	KFAccountListURL   = "https://api.weixin.qq.com/cgi-bin/customservice/getkflist"
	KFOnlineListURL    = "https://api.weixin.qq.com/cgi-bin/customservice/getonlinekflist"
	KFAccountAddURL    = "https://api.weixin.qq.com/customservice/kfaccount/add"
	KFInviteURL        = "https://api.weixin.qq.com/customservice/kfaccount/inviteworker"
	KFAccountUpdateURL = "https://api.weixin.qq.com/customservice/kfaccount/update"
	KFAvatarUploadURL  = "https://api.weixin.qq.com/customservice/kfaccount/uploadheadimg"
	KFDeleteURL        = "https://api.weixin.qq.com/customservice/kfaccount/del"
	KFSessionCreateURL = "https://api.weixin.qq.com/customservice/kfsession/create"
	KFSessionCloseURL  = "https://api.weixin.qq.com/customservice/kfsession/close"
	KFSessionGetURL    = "https://api.weixin.qq.com/customservice/kfsession/getsession"
	KFSessionListURL   = "https://api.weixin.qq.com/customservice/kfsession/getsessionlist"
	KFWaitCaseURL      = "https://api.weixin.qq.com/customservice/kfsession/getwaitcase"
	KFMsgRecordListURL = "https://api.weixin.qq.com/customservice/msgrecord/getmsglist"
)
