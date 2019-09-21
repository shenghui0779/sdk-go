package mch

// 交易类型
const (
	TradeAPP    = "APP"
	TradeJSAPI  = "JSAPI"
	TradeMWEB   = "MWEB"
	TradeNative = "NATIVE"
	TradePAP    = "PAP"
)

// 签名类型
const (
	SignMD5        = "MD5"
	SignHMacSHA256 = "HMAC-SHA256"
)

// 返回结果
const (
	ReplySuccess = "SUCCESS"
	ReplyFailed  = "FAIL"
)

// URL - order
const (
	OrderUnifyURL = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	OrderQueryURL = "https://api.mch.weixin.qq.com/pay/orderquery"
	OrderCloseURL = "https://api.mch.weixin.qq.com/pay/closeorder"
)

// URL - refund
const (
	RefundApplyURL = "https://api.mch.weixin.qq.com/secapi/pay/refund"
	RefundQueryURL = "https://api.mch.weixin.qq.com/pay/refundquery"
)

// URL - pappay
const (
	PappayAPPEntrustURL     = "https://api.mch.weixin.qq.com/papay/preentrustweb"
	PappayPubEntrustURL     = "https://api.mch.weixin.qq.com/papay/entrustweb"
	PappayH5EntrustURL      = "https://api.mch.weixin.qq.com/papay/h5entrustweb"
	PappayContractOrderURL  = "https://api.mch.weixin.qq.com/pay/contractorder"
	PappayContractQueryURL  = "https://api.mch.weixin.qq.com/papay/querycontract"
	PappayContractDeleteURL = "https://api.mch.weixin.qq.com/papay/deletecontract"
	PappayPayApplyURL       = "https://api.mch.weixin.qq.com/pay/pappayapply"
	PappayOrderQueryURL     = "https://api.mch.weixin.qq.com/pay/paporderquery"
)
