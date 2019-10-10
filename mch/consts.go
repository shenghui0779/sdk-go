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
	ResultNull    = "RESULT NULL" // 查询结果为空
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
	RefundSuccess    = "SUCCESS"     // 退款成功
	RefundClosed     = "REFUNDCLOSE" // 退款关闭
	RefundProcessing = "PROCESSING"  // 退款处理中
	RefundChange     = "CHANGE"      // 退款异常
)

const (
	OrderNotExist  = "ORDERNOTEXIST"  // 订单不存在
	RefundNotExist = "REFUNDNOTEXIST" // 退款不存在
)

const (
	ContractAdd    = "ADD"    // 签约
	ContractDelete = "DELETE" // 解约
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

const (
	PappayErrAccount             = "ACCOUNTERROR"          // 用户帐号注销、银行卡异常或注销
	PappaytErrContractNotExist   = "CONTRACT_NOT_EXIST"    // 用户签约协议已过期或已解约
	PappayErrRuleLimit           = "RULELIMIT"             // 用户支付银行卡限额不足
	PappayErrBank                = "BANKERROR"             // 用户支付银行暂时无法提供服务
	PappayErrNotEnough           = "NOTENOUGH"             // 用户余额不足
	PappayErrUserAccountAbnormal = "USER_ACCOUNT_ABNORMAL" // 扣款用户的微信账号异常导致
	PappayErrUserNotExist        = "USER_NOT_EXIST"        // 扣款用户的微信账号已注销
)

// URL - order
const (
	OrderUnifyURL = "https://api.mch.weixin.qq.com/pay/unifiedorder" // 统一下单
	OrderQueryURL = "https://api.mch.weixin.qq.com/pay/orderquery"   // 订单查询
	OrderCloseURL = "https://api.mch.weixin.qq.com/pay/closeorder"   // 订单关闭
)

// URL - refund
const (
	RefundApplyURL = "https://api.mch.weixin.qq.com/secapi/pay/refund" // 申请退款
	RefundQueryURL = "https://api.mch.weixin.qq.com/pay/refundquery"   // 退款查询
)

// URL - pappay
const (
	PappayAPPEntrustURL     = "https://api.mch.weixin.qq.com/papay/preentrustweb"  // APP纯签约
	PappayPubEntrustURL     = "https://api.mch.weixin.qq.com/papay/entrustweb"     // 公众号纯签约
	PappayH5EntrustURL      = "https://api.mch.weixin.qq.com/papay/h5entrustweb"   // H5纯签约
	PappayContractOrderURL  = "https://api.mch.weixin.qq.com/pay/contractorder"    // 支付中签约
	PappayContractQueryURL  = "https://api.mch.weixin.qq.com/papay/querycontract"  // 签约查询
	PappayContractDeleteURL = "https://api.mch.weixin.qq.com/papay/deletecontract" // 申请解约
	PappayPayApplyURL       = "https://api.mch.weixin.qq.com/pay/pappayapply"      // 申请扣款
	PappayOrderQueryURL     = "https://api.mch.weixin.qq.com/pay/paporderquery"    // 扣款查询
)
