# 支付

## 使用

```go
import "github.com/shenghui0779/wechat_pay"
```

### 初始化商户实例

```go
mch := gochat.NewMch(appid, mchid, apikey)

// 加载证书
if err := mch.LoadCertFromPemBlock(certBlock, keyBlock); err != nil {
    // 错误处理...
}

// mch.LoadCertFromPemFile(certFile, keyFile)
// mch.LoadCertFromP12File(path)
```

### 统一下单

- 下单参数

```go
// OrderData 统一下单数据
type OrderData struct {
	// 必填参数
	OutTradeNO     string // 商户系统内部的订单号，32个字符内、可包含字母，其他说明见商户订单号
	TotalFee       int    // 订单总金额，单位为分，详见支付金额
	SpbillCreateIP string // APP和网页支付提交用户端ip，Native支付填调用微信支付API的机器IP
	TradeType      string // 取值如下：JSAPI，NATIVE，APP，MWEB，详细说明见参数规定
	Body           string // 商品或支付单简要描述
	NotifyURL      string // 接收微信支付异步通知回调地址，通知url必须为直接可访问的url，不能携带参数
	// 选填参数
	DeviceInfo string // 终端设备号(门店号或收银设备ID)，注意：PC网页或公众号内支付请传"WEB"
	Detail     string // 商品名称明细列表
	Attach     string // 附加数据，在查询API和支付通知中原样返回，该字段主要用于商户携带订单的自定义数据
	FeeType    string // 符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	TimeStart  string // 订单生成时间，格式为yyyyMMddHHmmss，如：2009年12月25日9点10分10秒 表示为：20091225091010
	TimeExpire string // 订单失效时间，格式为yyyyMMddHHmmss，如：2009年12月27日9点10分10秒 表示为：20091227091010
	GoodsTag   string // 商品标记，代金券或立减优惠功能的参数，说明详见代金券或立减优惠
	ProductID  string // trade_type=NATIVE，此参数必传。此id为二维码中包含的商品ID，商户自行定义
	LimitPay   string // no_credit--指定不能使用信用卡支付
	OpenID     string // trade_type=JSAPI，此参数必传，用户在商户appid下的唯一标识
	Receipt    bool   // 是否在支付成功消息和支付详情页中出现开票入口，注：需要在微信支付商户平台或微信公众平台开通电子发票功能
	SceneInfo  string // 该字段用于上报支付的场景信息
}
```

- 统一下单

```go
r, err := mch.Do(ctx, UnifyOrder(orderData))

if err != nil {
    // 错误处理...
}

if r["result_code"] != mch.ResultSuccess {
    // 失败处理...
}

// 成功逻辑处理...
```

- 拉起支付

```go
prepayID := r["prepay_id"]

// APP支付
appapi := mch.APPAPI(prepayID, time.Now().Unix())

// JSAPI支付
jspai := mch.JSAPI(prepayID, time.Now().Unix())
```

### 订单查询

- 根据微信订单号查询

```go
r, err := mch.Do(ctx, QueryOrderByTransactionID(transactionID))
```

- 根据商户订单号查询

```go
r, err := mch.Do(ctx, QueryOrderByOutTradeNO(outTradeNO))
```

### 关闭订单

```go
r, err := mch.Do(ctx, CloseOrder(outTradeNO))
```

### 申请退款

- 退款参数

```go
// RefundData 退款数据
type RefundData struct {
	// 必填参数
	OutRefundNO string // 商户系统内部的退款单号，商户系统内部唯一，同一退款单号多次请求只退一笔
	TotalFee    int    // 订单总金额，单位为分，只能为整数，详见支付金额
	RefundFee   int    // 退款总金额，订单总金额，单位为分，只能为整数，详见支付金额
	// 选填参数
	RefundFeeType string // 货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	RefundDesc    string // 若商户传入，会在下发给用户的退款消息中体现退款原因
	RefundAccount string // 退款资金来源，仅针对老资金流商户使用
	NotifyURL     string // 异步接收微信支付退款结果通知的回调地址，通知URL必须为外网可访问的url，不允许带参数
}
```

- 根据微信订单号退款

```go
r, err := mch.Do(ctx, RefundByTransactionID(transactionID, refundData))

if err != nil {
    // 错误处理...
}

if r["result_code"] != mch.ResultSuccess {
    // 失败处理...
}

// 成功逻辑处理...
```

- 根据商户订单号退款

```go
r, err := mch.Do(ctx, RefundByOutTradeNO(outTradeNO, refundData))

if err != nil {
    // 错误处理...
}

if r["result_code"] != mch.ResultSuccess {
    // 失败处理...
}

// 成功逻辑处理...
```

### 查询退款

- 根据微信退款单号查询

```go
r, err := mch.Do(ctx, QueryRefundByRefundID(refundID))
```

- 根据商户退款单号查询

```go
r, err := mch.Do(ctx, QueryRefundByOutRefundNO(outRefundNO))
```

- 根据微信订单号查询

```go
r, err := mch.Do(ctx, QueryRefundByTransactionID(transactionID))
```

- 根据商户订单号查询

```go
r, err := mch.Do(ctx, QueryRefundByOutTradeNO(outTradeNO))
```

### 委托扣款

#### 纯签约

- 签约参数

```go
// Contract 微信纯签约协议
type Contract struct {
	// 必填字段
	PlanID                 string // 协议模板id，设置路径见开发步骤
	ContractCode           string // 商户侧的签约协议号，由商户生成
	RequestSerial          int64  // 商户请求签约时的序列号，要求唯一性，纯数字, 范围不能超过Int64的范围
	ContractDisplayAccount string // 签约用户的名称，用于页面展示，参数值不支持UTF8非3字节编码的字符，如表情符号，故请勿使用微信昵称
	SpbillCreateIP         string // 用户客户端的真实IP地址，H5签约必填
	Timestamp              int64  // 系统当前时间戳，10位
	NotifyURL              string // 用于接收签约成功消息的回调通知地址，对notify_url参数值需进行encode处理,注意是对参数值进行encode
	// 选填字段
	ReturnAPP   bool   // APP签约选填，签约后是否返回app，注：签约参数appid必须为发起签约的app所有，且在微信开放平台注册过
	ReturnWeb   bool   // 公众号签约选填，签约后是否返回签约页面的referrer url, 不填或获取不到referrer则不返回; 跳转referrer url时会自动带上参数from_wxpay=1
	OuterID     int64  // 小程序签约选填，用户在商户侧的标识
	ReturnAPPID string // H5签约选填，商户具有指定返回app的权限时，签约成功将返回appid指定的app应用，如不填且签约发起时的浏览器UA可被微信识别，则跳转到浏览器，否则留在微信
}
```

- APP纯签约

```go
r, err := mch.Do(ctx, APPEntrust(contract))
```

- 公众号纯签约

```go
r := mch.Do(ctx, OAEntrust(contract))
```

- 小程序纯签约，返回小程序所需的 `extraData` 数据

```go
r := mch.Do(ctx, MPEntrust(contract))
```

- H5纯签约

```go
r := mch.Do(ctx, H5Entrust(contract))
```

#### 支付中签约

- 支付并签约参数

```go
// ContractOrder 支付并签约
type ContractOrder struct {
	// 必填参数
	OutTradeNO             string // 商户系统内部的订单号，32个字符内、可包含字母，其他说明见商户订单号
	TotalFee               int    // 订单总金额，单位为分，详见支付金额
	SpbillCreateIP         string // APP和网页支付提交用户端ip，Native支付填调用微信支付API的机器IP
	TradeType              string // 取值如下：JSAPI，NATIVE，APP，MWEB，详细说明见参数规定
	Body                   string // 商品或支付单简要描述
	PlanID                 string // 协议模板id，设置路径见开发步骤
	ContractCode           string // 商户侧的签约协议号，由商户生成
	RequestSerial          int64  // 商户请求签约时的序列号，要求唯一性，纯数字, 范围不能超过Int64的范围
	ContractDisplayAccount string // 签约用户的名称，用于页面展示，参数值不支持UTF8非3字节编码的字符，如表情符号，故请勿使用微信昵称
	PaymentNotifyURL       string // 接收微信支付异步通知回调地址，通知url必须为直接可访问的url，不能携带参数
	ContractNotifyURL      string // 签约信息回调通知的url
	// 选填参数
	DeviceInfo string // 终端设备号(门店号或收银设备ID)，注意：PC网页或公众号内支付请传"WEB"
	Detail     string // 商品名称明细列表
	Attach     string // 附加数据，在查询API和支付通知中原样返回，该字段主要用于商户携带订单的自定义数据
	FeeType    string // 符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	TimeStart  string // 订单生成时间，格式为yyyyMMddHHmmss，如：2009年12月25日9点10分10秒 表示为：20091225091010
	TimeExpire string // 订单失效时间，格式为yyyyMMddHHmmss，如：2009年12月27日9点10分10秒 表示为：20091227091010
	GoodsTag   string // 商品标记，代金券或立减优惠功能的参数，说明详见代金券或立减优惠
	ProductID  string // trade_type=NATIVE，此参数必传。此id为二维码中包含的商品ID，商户自行定义
	LimitPay   string // no_credit--指定不能使用信用卡支付
	OpenID     string // trade_type=JSAPI，此参数必传，用户在商户appid下的唯一标识
}
```

- 下单

```go
r, err := mch.Do(ctx, EntrustOrder(orderData))

if err != nil {
    // 错误处理...
}

if r["result_code"] != mch.ResultSuccess {
    // 失败处理...
}

// 成功逻辑处理...
```

#### 签约查询

- 根据微信返回的委托代扣协议id查询签约关系

```go
r, err := mch.Do(ctx, QueryContractByID(contractID))
```

- 根据签约协议号查询签约关系，需要商户平台配置的代扣模版id

```go
r, err := mch.Do(ctx, QueryContractByCode(planID, contractCode))
```

#### 申请扣款

- 扣款参数

```go
// PappayApplyData 扣款申请数据
type PappayApplyData struct {
	// 必填参数
	OutTradeNO     string // 商户系统内部的订单号，32个字符内、可包含字母，其他说明见商户订单号
	TotalFee       int    // 订单总金额，单位为分，详见支付金额
	SpbillCreateIP string // APP和网页支付提交用户端ip，Native支付填调用微信支付API的机器IP
	ContractID     string // 签约成功后，微信返回的委托代扣协议id
	Body           string // 商品或支付单简要描述
	NotifyURL      string // 接收微信支付异步通知回调地址，通知url必须为直接可访问的url，不能携带参数
	// 选填参数
	Detail   string // 商品名称明细列表
	Attach   string // 附加数据，在查询API和支付通知中原样返回，该字段主要用于商户携带订单的自定义数据
	FeeType  string // 符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	GoodsTag string // 商品标记，代金券或立减优惠功能的参数，说明详见代金券或立减优惠
	Receipt  bool   // 是否在支付成功消息和支付详情页中出现开票入口，注：需要在微信支付商户平台或微信公众平台开通电子发票功能
}
```

- 申请扣款

```go
r, err := mch.Do(ctx, PappayApply(applyData))

if err != nil {
    // 错误处理...
}

if r["result_code"] != mch.ResultSuccess {
    // 失败处理...
}

// 成功逻辑处理...
```

#### 扣款订单查询

- 根据微信订单号查询

```go
r, err := mch.Do(ctx, QueryPappayByTransactionID(transactionID))
```

- 根据商户订单号查询

```go
r, err := mch.Do(ctx, QueryPappayByOutTradeNO(outTradeNO))
```

#### 解约

- 根据微信返回的委托代扣协议id解约

```go
r, err := mch.Do(ctx, DeleteContractByID(contractID, remark))
```

- 根据签约协议号解约，需要商户平台配置的代扣模版id

```go
r, err := mch.Do(ctx, DeleteContractByID(planID, contractCode, remark))
```

### 企业支付

#### 付款到零钱

- 付款参数

```go
// TransferBalanceData 付款到零钱数据
type TransferBalanceData struct {
	// 必填参数
	PartnerTradeNO string // 商户订单号，需保持唯一性 (只能是字母或者数字，不能包含有其它字符)
	OpenID         string // 商户appid下，某用户的openid
	CheckName      string // NO_CHECK：不校验真实姓名；FORCE_CHECK：强校验真实姓名
	Amount         int    // 企业付款金额，单位：分
	Desc           string // 企业付款备注，必填。注意：备注中的敏感词会被转成字符*
	// 选填参数
	ReUserName     string // 收款用户真实姓名。如果check_name设置为FORCE_CHECK，则必填用户真实姓名
	DeviceInfo     string // 微信支付分配的终端设备号
	SpbillCreateIP string // 该IP同在商户平台设置的IP白名单中的IP没有关联，该IP可传用户端或者服务端的IP
}
```

- 付款

```go
r, err := mch.Do(ctx, TransferToBalance(balanceData))

if err != nil {
    // 错误处理...
}

if r["result_code"] != mch.ResultSuccess {
    // 失败处理...
}

// 成功逻辑处理...
```

- 付款单查询

```go
r, err := mch.Do(ctx, QueryTransferBalanceOrder(partnerTradeNO))
```

#### 付款到银行卡

- 付款参数

```go
// TransferBankCardData 付款到银行卡数据
type TransferBankCardData struct {
	// 必填参数
	PartnerTradeNO string // 商户订单号，需保持唯一（只允许数字[0~9]或字母[A~Z]和[a~z]，最短8位，最长32位）
	EncBankNO      string // 收款方银行卡号（采用标准RSA算法，公钥由微信侧提供）
	EncTrueName    string // 收款方用户名（采用标准RSA算法，公钥由微信侧提供）
	BankCode       string // 银行卡所在开户行编号，参考：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_4
	Amount         int    // 付款金额：RMB分（支付总额，不含手续费）注：大于0的整数
	// 选填参数
	Desc string // 企业付款到银行卡付款说明，即订单备注（UTF8编码，允许100个字符以内）
}
```

- 付款

```go
r, err := mch.Do(ctx, TransferToBankCard(bankCardData, pubKey))

if err != nil {
    // 错误处理...
}

if r["result_code"] != mch.ResultSuccess {
    // 失败处理...
}

// 成功逻辑处理...
```

- 付款单查询

```go
r, err := mch.Do(ctx, QueryTransferBankCardOrder(partnerTradeNO))
```

### 企业红包

- 红包参数

```go
// RedpackData 红包发放数据
type RedpackData struct {
	// 必填参数
	MchBillNO   string // 商户订单号（每个订单号必须唯一。取值范围：0~9，a~z，A~Z）接口根据商户订单号支持重入，如出现超时可再调用
	SendName    string // 红包发送者名称；注意：敏感词会被转义成字符*
	ReOpenID    string // 接受红包的用户openid
	TotalAmount int    // 付款金额，单位：分
	TotalNum    int    // 红包发放总人数
	Wishing     string // 红包祝福语；注意：敏感词会被转义成字符*
	ClientIP    string // 调用接口的机器Ip地址
	ActName     string // 活动名称；注意：敏感词会被转义成字符*
	Remark      string // 备注信息
	// 选填参数
	AmtType   string // 红包金额设置方式，适用于裂变红包，ALL_RAND — 全部随机,商户指定总金额和红包发放总人数，由微信支付随机计算出各红包金额
	NotifyWay string // 通过JSAPI方式领取红包，小程序红包固定传 MINI_PROGRAM_JSAPI
	SceneID   string // 发放红包使用场景，红包金额大于200或者小于1元时必传
	RiskInfo  string // 活动信息，urlencode(posttime=xx&mobile=xx&deviceid=xx。posttime：用户操作的时间戳；mobile：业务系统账号的手机号，国家代码-手机号，不需要+号；deviceid：MAC地址或者设备唯一标识；clientversion：用户操作的客户端版本
}
```

- 发放普通红包

```go
r, err := mch.Do(ctx, SendNormalRedpack(redpackData))

if err != nil {
    // 错误处理...
}

if r["result_code"] != mch.ResultSuccess {
    // 失败处理...
}

// 成功逻辑处理...
```

- 发放裂变红包

```go
r, err := mch.Do(ctx, SendGroupRedpack(redpackData))

if err != nil {
    // 错误处理...
}

if r["result_code"] != mch.ResultSuccess {
    // 失败处理...
}
```

- 发放小程序红包

```go
r, err := mch.Do(ctx, SendMinipRedpack(redpackData))

if err != nil {
    // 错误处理...
}

if r["result_code"] != mch.ResultSuccess {
    // 失败处理...
}

// 领取红包
jsapi := mch.APPAPI(r["package"], time.Now().Unix())
```

- 查询红包记录

```go
r, err := mch.Do(ctx, QueryRedpackByBillNO(billNO))
```

