package pub

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
	SnsCode2Token            = "https://api.weixin.qq.com/sns/oauth2/access_token"
	SnsCheckAccessTokenURL   = "https://api.weixin.qq.com/sns/auth"
	SnsRefreshAccessTokenURL = "https://api.weixin.qq.com/sns/oauth2/refresh_token"
	SnsUserInfoURL           = "https://api.weixin.qq.com/sns/userinfo"
)

// subscriber
const (
	SubscriberGetURL      = "https://api.weixin.qq.com/cgi-bin/user/info"
	SubscriberBatchGetURL = "https://api.weixin.qq.com/cgi-bin/user/info/batchget"
	SubscriberListURL     = "https://api.weixin.qq.com/cgi-bin/user/get"
)

// msg
const (
	TplMsgSendURL = "https://api.weixin.qq.com/cgi-bin/message/template/send"
)
