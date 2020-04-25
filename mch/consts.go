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
	SystemError   = "SYSTEMERROR" // 系统繁忙，请稍后再试
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
	TransferNoCheck    = "NO_CHECK"    // 不校验真实姓名
	TransferForceCheck = "FORCE_CHECK" // 强校验真实姓名
)

const (
	RedpackScene1 = "PRODUCT_1" // 商品促销
	RedpackScene2 = "PRODUCT_2" // 抽奖
	RedpackScene3 = "PRODUCT_3" // 虚拟物品兑奖
	RedpackScene4 = "PRODUCT_4" // 企业内部福利
	RedpackScene5 = "PRODUCT_5" // 渠道分润
	RedpackScene6 = "PRODUCT_6" // 保险回馈
	RedpackScene7 = "PRODUCT_7" // 彩票派奖
	RedpackScene8 = "PRODUCT_8" // 税务刮奖
)

const RSAPublicKeyURL = "https://fraud.mch.weixin.qq.com/risk/getpublickey"

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

// URL - transfer
const (
	TransferToBalanceURL          = "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers" // 企业付款到零钱
	TransferBalanceOrderQueryURL  = "https://api.mch.weixin.qq.com/mmpaymkttransfers/gettransferinfo"     // 企业付款到零钱订单查询
	TransferToBankCardURL         = "https://api.mch.weixin.qq.com/mmpaysptrans/pay_bank"                 // 企业付款到银行卡
	TransferBankCardOrderQueryURL = "https://api.mch.weixin.qq.com/mmpaysptrans/query_bank"               // 企业付款到银行卡订单查询
)

// URL - redpack

const (
	RedpackNormalURL = "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendredpack"
	RedpackGroupURL  = "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendgroupredpack"
	RedpackMinipURL  = "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendminiprogramhb"
	RedpackQueryURL  = "https://api.mch.weixin.qq.com/mmpaymkttransfers/gethbinfo"
)
