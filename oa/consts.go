package oa

// cgi-bin
const (
	CgiBinAccessTokenURL = "https://api.weixin.qq.com/cgi-bin/token"
	CgiBinTicketURL      = "https://api.weixin.qq.com/cgi-bin/ticket/getticket"
)

// menu
const (
	MenuCreateURL            = "https://api.weixin.qq.com/cgi-bin/menu/create"
	MenuAddConditionalURL    = "https://api.weixin.qq.com/cgi-bin/menu/addconditional"
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
	TemplateMessageSendURL = "https://api.weixin.qq.com/cgi-bin/message/template/send"
)

// media
const (
	MediaUploadURL         = "https://api.weixin.qq.com/cgi-bin/media/upload"
	MediaGetURL            = "https://api.weixin.qq.com/cgi-bin/media/get"
	MaterialNewsUploadURL  = "https://api.weixin.qq.com/cgi-bin/material/add_news"
	MaterialImageUploadURL = "https://api.weixin.qq.com/cgi-bin/media/uploadimg"
)
