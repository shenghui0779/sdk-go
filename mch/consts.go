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
	NotFound      = "NOT_FOUND"   // 数据不存在
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
	CouponTypeCash   = "CASH"    // 充值代金券
	CouponTypeNoCash = "NO_CASH" // 非充值优惠券
)

const (
	RefundStatusSuccess    = "SUCCESS"     // 退款成功
	RefundStatusClosed     = "REFUNDCLOSE" // 退款关闭
	RefundStatusProcessing = "PROCESSING"  // 退款处理中
	RefundStatusChange     = "CHANGE"      // 退款异常
)

const (
	RefundChannelOriginal      = "ORIGINAL"       // 原路退款
	RefundChannelBalance       = "BALANCE"        // 退回到余额
	RefundChannelOtherBalance  = "OTHER_BALANCE"  // 原账户异常退到其他余额账户
	RefundChannelOtherBankCard = "OTHER_BANKCARD" // 原银行卡异常退到其他银行卡
)

const (
	OrderNotExist  = "ORDERNOTEXIST"  // 订单不存在
	RefundNotExist = "REFUNDNOTEXIST" // 退款不存在
)

const (
	ContractOAEntrust = "offical_accounts_entrust"
	ContractMPEntrust = "mini_program_entrust"
	ContractH5Entrust = "h5_entrust"
)

const (
	ContractAdd    = "ADD"    // 签约
	ContractDelete = "DELETE" // 解约
)

const (
	ContractEntrustUndo       = "1" // 未签约
	ContractEntrustOK         = "0" // 已签约
	ContractEntrustProcessing = "9" // 签约进行中
)

const (
	ContractDeleteUndo     = "0" // 未解约
	ContractDeleteExpired  = "1" // 有效期过自动解约
	ContractDeleteUser     = "2" // 用户主动解约
	ContractDeleteAPI      = "3" // 商户API解约
	ContractDeletePlatform = "4" // 商户平台解约
	ContractDeleteLogout   = "5" // 注销
	ContractDeleteContact  = "7" // 用户联系客服发起的解约
)

const (
	TransferNoCheck    = "NO_CHECK"    // 不校验真实姓名
	TransferForceCheck = "FORCE_CHECK" // 强校验真实姓名
)

const (
	TransferStatusProcessing = "PROCESSING" // 处理中
	TransferStatusSuccess    = "SUCCESS"    // 转账成功
	TransferStatusFailed     = "FAILED"     // 转账失败
	TransferStatusBankFail   = "BANK_FAIL"  // 银行退票
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

const (
	RedpackStatusSending   = "SENDING"   // 发放中
	RedpackStatusSent      = "SENT"      // 已发放待领取
	RedpackStatusFailed    = "FAILED"    // 发放失败
	RedpackStatusReceived  = "RECEIVED"  // 已领取
	RedpackStatusRefunding = "RFUND_ING" // 退款中
	RedpackStatusRefund    = "REFUND"    // 已退款
)

const (
	RedpackTypeNormal = "NORMAL" // 普通红包
	RedpackTypeGroup  = "GROUP"  // 裂变红包
)

const (
	RedpackSendTypeAPI      = "API"      // 通过API接口发放
	RedpackSendTypeUpload   = "UPLOAD"   // 通过上传文件方式发放
	RedpackSendTypeActivity = "ACTIVITY" // 通过活动方式发放
)

// 账单类型
const (
	BillTypeAll            = "ALL"             // 当日所有订单信息（不含充值退款订单）
	BillTypeSuccess        = "SUCCESS"         // 当日成功支付的订单（不含充值退款订单）
	BillTypeRefund         = "REFUND"          // 当日退款订单（不含充值退款订单）
	BillTypeRechargeRefund = "RECHARGE_REFUND" // 当日充值退款订单
)

// 资金账户类型
const (
	AccountTypeBasic     = "Basic"     // 基本账户
	AccountTypeOperation = "Operation" // 运营账户
	AccountTypeFees      = "Fees"      // 手续费账户
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
	PappayOAEntrustURL      = "https://api.mch.weixin.qq.com/papay/entrustweb"     // 公众号纯签约
	PappayH5EntrustURL      = "https://api.mch.weixin.qq.com/papay/h5entrustweb"   // H5纯签约
	PappayContractOrderURL  = "https://api.mch.weixin.qq.com/pay/contractorder"    // 支付中签约
	PappayContractQueryURL  = "https://api.mch.weixin.qq.com/papay/querycontract"  // 签约查询
	PappayContractDeleteURL = "https://api.mch.weixin.qq.com/papay/deletecontract" // 申请解约
	PappayApplyURL          = "https://api.mch.weixin.qq.com/pay/pappayapply"      // 申请扣款
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
	RedpackNormalURL = "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendredpack"       // 普通红包
	RedpackGroupURL  = "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendgroupredpack"  // 裂变红包
	RedpackMinipURL  = "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendminiprogramhb" // 小程序红包
	RedpackQueryURL  = "https://api.mch.weixin.qq.com/mmpaymkttransfers/gethbinfo"         // 红包查询
)

// URL - other
const (
	DownloadBillURL      = "https://api.mch.weixin.qq.com/pay/downloadbill"                // 下载交易账单
	DownloadFundFlowURL  = "https://api.mch.weixin.qq.com/pay/downloadfundflow"            // 下载资金账单
	BatchQueryCommentURL = "https://api.mch.weixin.qq.com/billcommentsp/batchquerycomment" // 拉取订单评价数据
)
