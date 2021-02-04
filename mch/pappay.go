package mch

import (
	"strconv"

	"github.com/shenghui0779/gochat/wx"
)

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

// PappayData 扣款数据
type PappayData struct {
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

// APPEntrust APP纯签约
func APPEntrust(c *Contract) wx.Action {
	return wx.NewAction(PappayAPPEntrustURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithWXML(func(appid, mchid, nonce string) (wx.WXML, error) {
			body := wx.WXML{
				"appid":                    appid,
				"mch_id":                   mchid,
				"plan_id":                  c.PlanID,
				"contract_code":            c.ContractCode,
				"request_serial":           strconv.FormatInt(c.RequestSerial, 10),
				"contract_display_account": c.ContractDisplayAccount,
				"version":                  "1.0",
				"timestamp":                strconv.FormatInt(c.Timestamp, 10),
				"notify_url":               c.NotifyURL,
				"sign_type":                SignMD5,
			}

			if c.ReturnAPP {
				body["return_app"] = "Y"
			}

			return body, nil
		}),
	)
}

// OAEntrust 公众号纯签约
func OAEntrust(c *Contract) wx.Action {
	return wx.NewAction(ContractOAEntrust,
		wx.WithWXML(func(appid, mchid, nonce string) (wx.WXML, error) {
			body := wx.WXML{
				"appid":                    appid,
				"mch_id":                   mchid,
				"plan_id":                  c.PlanID,
				"contract_code":            c.ContractCode,
				"request_serial":           strconv.FormatInt(c.RequestSerial, 10),
				"contract_display_account": c.ContractDisplayAccount,
				"version":                  "1.0",
				"timestamp":                strconv.FormatInt(c.Timestamp, 10),
				"notify_url":               c.NotifyURL,
				"sign_type":                SignMD5,
			}

			if c.ReturnWeb {
				body["return_web"] = "1"
			}

			return body, nil
		}),
	)
}

// MPEntrust 小程序纯签约，返回小程序所需的 extraData 数据
func MPEntrust(c *Contract) wx.Action {
	return wx.NewAction(ContractMPEntrust,
		wx.WithWXML(func(appid, mchid, nonce string) (wx.WXML, error) {
			extraData := wx.WXML{
				"appid":                    appid,
				"mch_id":                   mchid,
				"plan_id":                  c.PlanID,
				"contract_code":            c.ContractCode,
				"request_serial":           strconv.FormatInt(c.RequestSerial, 10),
				"contract_display_account": c.ContractDisplayAccount,
				"timestamp":                strconv.FormatInt(c.Timestamp, 10),
				"notify_url":               c.NotifyURL,
				"sign_type":                SignMD5,
			}

			if c.OuterID != 0 {
				extraData["outerid"] = strconv.FormatInt(c.OuterID, 10)
			}

			return extraData, nil
		}),
	)
}

// H5Entrust H5纯签约
func H5Entrust(c *Contract) wx.Action {
	return wx.NewAction(ContractH5Entrust,
		wx.WithWXML(func(appid, mchid, nonce string) (wx.WXML, error) {
			body := wx.WXML{
				"appid":                    appid,
				"mch_id":                   mchid,
				"plan_id":                  c.PlanID,
				"contract_code":            c.ContractCode,
				"request_serial":           strconv.FormatInt(c.RequestSerial, 10),
				"contract_display_account": c.ContractDisplayAccount,
				"version":                  "1.0",
				"timestamp":                strconv.FormatInt(c.Timestamp, 10),
				"clientip":                 c.SpbillCreateIP,
				"notify_url":               c.NotifyURL,
				"sign_type":                SignHMacSHA256,
			}

			if c.ReturnAPPID != "" {
				body["return_appid"] = c.ReturnAPPID
			}

			return body, nil
		}),
	)
}

// EntrustByOrder 支付中签约
func EntrustByOrder(order *ContractOrder) wx.Action {
	return wx.NewAction(PappayContractOrderURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithWXML(func(appid, mchid, nonce string) (wx.WXML, error) {
			body := wx.WXML{
				"appid":                    appid,
				"mch_id":                   mchid,
				"contract_appid":           appid,
				"contract_mchid":           mchid,
				"nonce_str":                nonce,
				"fee_type":                 "CNY",
				"trade_type":               order.TradeType,
				"body":                     order.Body,
				"out_trade_no":             order.OutTradeNO,
				"total_fee":                strconv.Itoa(order.TotalFee),
				"spbill_create_ip":         order.SpbillCreateIP,
				"plan_id":                  order.PlanID,
				"contract_code":            order.ContractCode,
				"request_serial":           strconv.FormatInt(order.RequestSerial, 10),
				"contract_display_account": order.ContractDisplayAccount,
				"notify_url":               order.PaymentNotifyURL,
				"contract_notify_url":      order.ContractNotifyURL,
				"sign_type":                SignMD5,
			}

			if order.DeviceInfo != "" {
				body["device_info"] = order.DeviceInfo
			}

			if order.Detail != "" {
				body["detail"] = order.Detail
			}

			if order.Attach != "" {
				body["attach"] = order.Attach
			}

			if order.FeeType != "" {
				body["fee_type"] = order.FeeType
			}

			if order.TimeStart != "" {
				body["time_start"] = order.TimeStart
			}

			if order.TimeExpire != "" {
				body["time_expire"] = order.TimeExpire
			}

			if order.GoodsTag != "" {
				body["goods_tag"] = order.GoodsTag
			}

			if order.ProductID != "" {
				body["product_id"] = order.ProductID
			}

			if order.LimitPay != "" {
				body["limit_pay"] = order.LimitPay
			}

			if order.OpenID != "" {
				body["openid"] = order.OpenID
			}

			return body, nil
		}),
	)
}

// QueryContractByID 根据微信返回的委托代扣协议id查询签约关系
func QueryContractByID(contractID string) wx.Action {
	return wx.NewAction(PappayContractQueryURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithWXML(func(appid, mchid, nonce string) (wx.WXML, error) {
			return wx.WXML{
				"appid":       appid,
				"mch_id":      mchid,
				"contract_id": contractID,
				"version":     "1.0",
				"sign_type":   SignMD5,
			}, nil
		}),
	)
}

// QueryContractByCode 根据签约协议号查询签约关系，需要商户平台配置的代扣模版id
func QueryContractByCode(planID, contractCode string) wx.Action {
	return wx.NewAction(PappayContractQueryURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithWXML(func(appid, mchid, nonce string) (wx.WXML, error) {
			return wx.WXML{
				"appid":         appid,
				"mch_id":        mchid,
				"plan_id":       planID,
				"contract_code": contractCode,
				"version":       "1.0",
				"sign_type":     SignMD5,
			}, nil
		}),
	)
}

// PappayApply 申请扣款
func PappayApply(data *PappayData) wx.Action {
	return wx.NewAction(PappayApplyURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithWXML(func(appid, mchid, nonce string) (wx.WXML, error) {
			body := wx.WXML{
				"appid":            appid,
				"mch_id":           mchid,
				"nonce_str":        nonce,
				"fee_type":         "CNY",
				"trade_type":       TradePAP,
				"notify_url":       data.NotifyURL,
				"body":             data.Body,
				"out_trade_no":     data.OutTradeNO,
				"total_fee":        strconv.Itoa(data.TotalFee),
				"contract_id":      data.ContractID,
				"spbill_create_ip": data.SpbillCreateIP,
				"sign_type":        SignMD5,
			}

			if data.Detail != "" {
				body["detail"] = data.Detail
			}

			if data.Attach != "" {
				body["attach"] = data.Attach
			}

			if data.FeeType != "" {
				body["fee_type"] = data.FeeType
			}

			if data.GoodsTag != "" {
				body["goods_tag"] = data.GoodsTag
			}

			if data.Receipt {
				body["receipt"] = "Y"
			}

			return body, nil
		}),
	)
}

// DeleteContractByID 根据微信返回的委托代扣协议id解约
func DeleteContractByID(contractID, remark string) wx.Action {
	return wx.NewAction(PappayContractDeleteURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithWXML(func(appid, mchid, nonce string) (wx.WXML, error) {
			return wx.WXML{
				"appid":                       appid,
				"mch_id":                      mchid,
				"contract_id":                 contractID,
				"version":                     "1.0",
				"contract_termination_remark": remark,
				"sign_type":                   SignMD5,
			}, nil
		}),
	)
}

// DeleteContractByCode 根据签约协议号解约，需要商户平台配置的代扣模版id
func DeleteContractByCode(planID, contractCode, remark string) wx.Action {
	return wx.NewAction(PappayContractDeleteURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithWXML(func(appid, mchid, nonce string) (wx.WXML, error) {
			return wx.WXML{
				"appid":                       appid,
				"mch_id":                      mchid,
				"plan_id":                     planID,
				"contract_code":               contractCode,
				"version":                     "1.0",
				"contract_termination_remark": remark,
				"sign_type":                   SignMD5,
			}, nil
		}),
	)
}

// QueryPappayByTransactionID 根据微信订单号查询扣款信息
func QueryPappayByTransactionID(transactionID string) wx.Action {
	return wx.NewAction(PappayOrderQueryURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithWXML(func(appid, mchid, nonce string) (wx.WXML, error) {
			return wx.WXML{
				"appid":          appid,
				"mch_id":         mchid,
				"transaction_id": transactionID,
				"nonce_str":      nonce,
				"sign_type":      SignMD5,
			}, nil
		}),
	)
}

// QueryPappayByOutTradeNO 根据商户订单号查询扣款信息
func QueryPappayByOutTradeNO(outTradeNO string) wx.Action {
	return wx.NewAction(PappayOrderQueryURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithWXML(func(appid, mchid, nonce string) (wx.WXML, error) {
			return wx.WXML{
				"appid":        appid,
				"mch_id":       mchid,
				"out_trade_no": outTradeNO,
				"nonce_str":    nonce,
				"sign_type":    SignMD5,
			}, nil
		}),
	)
}
