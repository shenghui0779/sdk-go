# 支付（普通商户直连模式）

```go
import (
    "github.com/shenghui0779/gochat"
    "github.com/shenghui0779/gochat/mch"
)
```

### 初始化商户实例

```go
wxpay := gochat.NewMch(appid, mchid, apikey)

// 涉及退款等，需要加载证书（三选一）
wxpay.LoadCertFromPemBlock(certBlock, keyBlock)
wxpay.LoadCertFromPemFile(certFile, keyFile)
wxpay.LoadCertFromP12File(path)
```

### 订单

```go
// 统一下单
wxpay.Do(ctx, mch.UnifyOrder(orderData))

// APP拉起支付
wxpay.APPAPI(prepayID)

// JSAPI拉起支付
wxpay.JSAPI(prepayID)

// 根据微信订单号查询
wxpay.Do(ctx, mch.QueryOrderByTransactionID(transactionID))

// 根据商户订单号查询
wxpay.Do(ctx, mch.QueryOrderByOutTradeNO(outTradeNO))

// 关闭订单
wxpay.Do(ctx, mch.CloseOrder(outTradeNO))
```

### 退款

```go
// 根据微信订单号退款
wxpay.Do(ctx, mch.RefundByTransactionID(transactionID, refundData))

// 根据商户订单号退款
wxpay.Do(ctx, mch.RefundByOutTradeNO(outTradeNO, refundData))

// 根据微信退款单号查询
wxpay.Do(ctx, mch.QueryRefundByRefundID(refundID))

// 根据商户退款单号查询
wxpay.Do(ctx, mch.QueryRefundByOutRefundNO(outRefundNO))

// 根据微信订单号查询
wxpay.Do(ctx, mch.QueryRefundByTransactionID(transactionID))

// 根据商户订单号查询
wxpay.Do(ctx, mch.QueryRefundByOutTradeNO(outTradeNO))
```

### 委托扣款

```go
// APP纯签约
wxpay.Do(ctx, mch.APPEntrust(contract))

// 公众号纯签约
wxpay.Do(ctx, mch.OAEntrust(contract))

// 小程序纯签约，返回小程序所需的 `extraData` 数据
wxpay.Do(ctx, mch.MPEntrust(contract))

// H5纯签约
wxpay.Do(ctx, mch.H5Entrust(contract))

// 支付中签约
wxpay.Do(ctx, mch.EntrustOrder(orderData))

// 根据微信返回的委托代扣协议id查询签约关系
wxpay.Do(ctx, mch.QueryContractByID(contractID))

// 根据签约协议号查询签约关系，需要商户平台配置的代扣模版id
wxpay.Do(ctx, mch.QueryContractByCode(planID, contractCode))

// 申请扣款
wxpay.Do(ctx, mch.PappayApply(pappayData))

// 根据微信订单号查询
wxpay.Do(ctx, mch.QueryPappayByTransactionID(transactionID))

// 根据商户订单号查询
wxpay.Do(ctx, mch.QueryPappayByOutTradeNO(outTradeNO))

// 根据微信返回的委托代扣协议id解约
wxpay.Do(ctx, mch.DeleteContractByID(contractID, remark))

// 根据签约协议号解约，需要商户平台配置的代扣模版id
wxpay.Do(ctx, mch.DeleteContractByID(planID, contractCode, remark))
```

### 企业付款

```go
// 付款到零钱
wxpay.Do(ctx, mch.TransferToBalance(balanceData))

// 付款到零钱订单查询
wxpay.Do(ctx, mch.QueryTransferBalanceOrder(partnerTradeNO))

// 付款到银行卡
wxpay.Do(ctx, mch.TransferToBankCard(bankCardData, pubKey))

// 付款到银行卡订单查询
wxpay.Do(ctx, mch.QueryTransferBankCardOrder(partnerTradeNO))
```

### 企业红包

```go
// 发放普通红包
wxpay.Do(ctx, mch.SendNormalRedpack(redpackData))

// 发放裂变红包
wxpay.Do(ctx, mch.SendGroupRedpack(redpackData))

// 发放小程序红包
wxpay.Do(ctx, mch.SendMinipRedpack(redpackData))

// 领取红包JSAPI
wxpay.MinipRedpackJSAPI(package)

// 查询红包记录
wxpay.Do(ctx, mch.QueryRedpackByBillNO(billNO))
```

### 回调通知

```go
// 签名验证
wxpay.VerifyWXMLResult(wxml)

// 退款信息解密
wxpay.DecryptWithAES256ECB(encrypt)
```

### 账单&评论

```go
// 下载交易账单
wxpay.DownloadBill(ctx, billDate, billType)

// 下载资金账单
wxpay.DownloadFundFlow(ctx, billDate, accountType)

// 拉取订单评价数据
wxpay.BatchQueryComment(ctx, beginTime, endTime, offset, limit)
```
