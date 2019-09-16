package consts

// cgi-bin
const (
	PubCgiBinAccessTokenURL = "https://api.weixin.qq.com/cgi-bin/token"
	PubCgiBinTicketURL      = "https://api.weixin.qq.com/cgi-bin/ticket/getticket"
)

// menu
const (
	PubMenuCreateURL            = "https://api.weixin.qq.com/cgi-bin/menu/create"
	PubMenuAddConditionalURL    = "https://api.weixin.qq.com/cgi-bin/menu/addconditional"
	PubMenuListURL              = "https://api.weixin.qq.com/cgi-bin/menu/get"
	PubMenuDeleteURL            = "https://api.weixin.qq.com/cgi-bin/menu/delete"
	PubMenuDeleteConditionalURL = "https://api.weixin.qq.com/cgi-bin/menu/delconditional"
)

// sns
const (
	PubSnsCode2Token            = "https://api.weixin.qq.com/sns/oauth2/access_token"
	PubSnsCheckAccessTokenURL   = "https://api.weixin.qq.com/sns/auth"
	PubSnsRefreshAccessTokenURL = "https://api.weixin.qq.com/sns/oauth2/refresh_token"
	PubSnsUserInfoURL           = "https://api.weixin.qq.com/sns/userinfo"
)

// subscriber
const (
	PubSubscriberGetURL      = "https://api.weixin.qq.com/cgi-bin/user/info"
	PubSubscriberBatchGetURL = "https://api.weixin.qq.com/cgi-bin/user/info/batchget"
	PubSubscriberListURL     = "https://api.weixin.qq.com/cgi-bin/user/get"
)

// tpl-msg
const (
	PubTplMsgSendURL = "https://api.weixin.qq.com/cgi-bin/message/template/send"
)
