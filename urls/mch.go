package urls

const MchRSAPublicKey = "https://fraud.mch.weixin.qq.com/risk/getpublickey"

// order
const (
	MchOrderUnify = "https://api.mch.weixin.qq.com/pay/unifiedorder" // 统一下单
	MchOrderQuery = "https://api.mch.weixin.qq.com/pay/orderquery"   // 订单查询
	MchOrderClose = "https://api.mch.weixin.qq.com/pay/closeorder"   // 订单关闭
)

// refund
const (
	MchRefundApply = "https://api.mch.weixin.qq.com/secapi/pay/refund" // 申请退款
	MchRefundQuery = "https://api.mch.weixin.qq.com/pay/refundquery"   // 退款查询
)

// pappay
const (
	MchPappayAPPEntrust     = "https://api.mch.weixin.qq.com/papay/preentrustweb"  // APP纯签约
	MchPappayOAEntrust      = "https://api.mch.weixin.qq.com/papay/entrustweb"     // 公众号纯签约
	MchPappayH5Entrust      = "https://api.mch.weixin.qq.com/papay/h5entrustweb"   // H5纯签约
	MchPappayContractOrder  = "https://api.mch.weixin.qq.com/pay/contractorder"    // 支付中签约
	MchPappayContractQuery  = "https://api.mch.weixin.qq.com/papay/querycontract"  // 签约查询
	MchPappayContractDelete = "https://api.mch.weixin.qq.com/papay/deletecontract" // 申请解约
	MchPappayApply          = "https://api.mch.weixin.qq.com/pay/pappayapply"      // 申请扣款
	MchPappayOrderQuery     = "https://api.mch.weixin.qq.com/pay/paporderquery"    // 扣款查询
)

// transfer
const (
	MchTransferToBalance          = "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers"             // 企业付款到零钱
	MchTransferBalanceOrderQuery  = "https://api.mch.weixin.qq.com/mmpaymkttransfers/gettransferinfo"                 // 企业付款到零钱订单查询
	MchTransferToBankCard         = "https://api.mch.weixin.qq.com/mmpaysptrans/pay_bank"                             // 企业付款到银行卡
	MchTransferBankCardOrderQuery = "https://api.mch.weixin.qq.com/mmpaysptrans/query_bank"                           // 企业付款到银行卡订单查询
	MchTransferToPocket           = "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/paywwsptrans2pocket"   // 企业向员工付款
	MchTransferPocketOrderQuery   = "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/querywwsptrans2pocket" // 企业向员工付款订单查询
)

// redpack
const (
	MchRedpackNormal    = "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendredpack"        // 普通红包
	MchRedpackGroup     = "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendgroupredpack"   // 裂变红包
	MchRedpackMinip     = "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendminiprogramhb"  // 小程序红包
	MchRedpackQuery     = "https://api.mch.weixin.qq.com/mmpaymkttransfers/gethbinfo"          // 红包查询
	MchRedpackCorp      = "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendworkwxredpack"  // 企业红包
	MchRedpackCorpQuery = "https://api.mch.weixin.qq.com/mmpaymkttransfers/queryworkwxredpack" // 企业红包查询
)

// other
const (
	MchDownloadBill      = "https://api.mch.weixin.qq.com/pay/downloadbill"                // 下载交易账单
	MchDownloadFundFlow  = "https://api.mch.weixin.qq.com/pay/downloadfundflow"            // 下载资金账单
	MchBatchQueryComment = "https://api.mch.weixin.qq.com/billcommentsp/batchquerycomment" // 拉取订单评价数据
)
