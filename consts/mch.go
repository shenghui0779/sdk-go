package consts

// 交易类型
const (
	MchTradeAPP    = "APP"
	MchTradeJSAPI  = "JSAPI"
	MchTradeMP     = "JSAPI"
	MchTradeH5     = "MWEB"
	MchTradeNative = "NATIVE"
)

// 签名类型
const (
	MchSignMD5        = "MD5"
	MchSignHMacSHA256 = "HMAC-SHA256"
)

// 返回结果
const (
	MchReplySuccess = "SUCCESS"
	MchReplyFailed  = "FAIL"
)

// URL - order
const (
	MchOrderUnifyURL = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	MchOrderQueryURL = "https://api.mch.weixin.qq.com/pay/orderquery"
	MchOrderCloseURL = "https://api.mch.weixin.qq.com/pay/closeorder"
)

// URL - refund
const (
	MchRefundApplyURL = "https://api.mch.weixin.qq.com/secapi/pay/refund"
	MchRefundQueryURL = "https://api.mch.weixin.qq.com/pay/refundquery"
)

// URL - pappay
const (
	MchPappayAPPEntrustURL     = "https://api.mch.weixin.qq.com/papay/preentrustweb"
	MchPappayPubEntrustURL     = "https://api.mch.weixin.qq.com/papay/entrustweb"
	MchPappayH5EntrustURL      = "https://api.mch.weixin.qq.com/papay/h5entrustweb"
	MchPappayContractQueryURL  = "https://api.mch.weixin.qq.com/papay/querycontract"
	MchPappayContractDeleteURL = "https://api.mch.weixin.qq.com/papay/deletecontract"
	MchPappayPayApplyURL       = "https://api.mch.weixin.qq.com/pay/pappayapply"
	MchPappayOrderQueryURL     = "https://api.mch.weixin.qq.com/pay/paporderquery"
)
