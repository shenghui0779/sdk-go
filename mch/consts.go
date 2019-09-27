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
	ResultSuccess = "SUCCESS"
	ResultFail    = "FAIL"
)

const (
	ContractAdd    = "ADD"
	ContractDelete = "DELETE"
)

const (
	TradeStateSuccess = "SUCCESS"    // 支付成功
	TradeStateRefund  = "REFUND"     // 转入退款
	TradeStateNotpay  = "NOTPAY"     // 未支付
	TradeStateClosed  = "CLOSED"     // 已关闭
	TradeStateRevoked = "REVOKED"    // 已撤销（刷卡支付）
	TradeStatePaying  = "USERPAYING" // 用户支付中
	TradeStateAccept  = "ACCEPT"     // 已接收，等待扣款
	TradeStateError   = "PAYERROR"   // 支付失败
	TradeStatePayFail = "PAY_FAIL"   // 支付失败(其他原因，如银行返回失败)
)

const (
	ContractEntrustUndo = "1" // 未签约
	ContractEntrustOK   = "0" // 已签约
)

const (
	ContractDeleteUndo     = "0" // 未解约
	ContractDeleteExpired  = "1" // 有效期过自动解约
	ContractDeleteUser     = "2" // 用户主动解约
	ContractDeleteAPI      = "3" // 商户API解约
	ContractDeletePlatform = "4" // 商户平台解约
	ContractDeleteLogout   = "5" // 注销
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
