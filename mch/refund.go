package mch

import (
	"strconv"

	"github.com/shenghui0779/gochat/utils"
)

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

// Refund 退款操作
type Refund struct {
	mch     *WXMch
	options []utils.RequestOption
}

// RefundByTransactionID 根据微信订单号退款
func (r *Refund) RefundByTransactionID(transactionID string, data *RefundData) (utils.WXML, error) {
	body := utils.WXML{
		"appid":          r.mch.appid,
		"mch_id":         r.mch.mchid,
		"nonce_str":      utils.NonceStr(),
		"sign_type":      SignMD5,
		"transaction_id": transactionID,
		"out_refund_no":  data.OutRefundNO,
		"total_fee":      strconv.Itoa(data.TotalFee),
		"refund_fee":     strconv.Itoa(data.RefundFee),
	}

	if data.RefundFeeType != "" {
		body["refund_fee_type"] = data.RefundFeeType
	}

	if data.RefundDesc != "" {
		body["refund_desc"] = data.RefundDesc
	}

	if data.RefundAccount != "" {
		body["refund_account"] = data.RefundAccount
	}

	if data.NotifyURL != "" {
		body["notify_url"] = data.NotifyURL
	}

	return r.mch.tlsPost(RefundApplyURL, body, r.options...)
}

// RefundByOutTradeNO 根据商户订单号退款
func (r *Refund) RefundByOutTradeNO(outTradeNO string, data *RefundData) (utils.WXML, error) {
	body := utils.WXML{
		"appid":         r.mch.appid,
		"mch_id":        r.mch.mchid,
		"nonce_str":     utils.NonceStr(),
		"sign_type":     SignMD5,
		"out_trade_no":  outTradeNO,
		"out_refund_no": data.OutRefundNO,
		"total_fee":     strconv.Itoa(data.TotalFee),
		"refund_fee":    strconv.Itoa(data.RefundFee),
	}

	if data.RefundFeeType != "" {
		body["refund_fee_type"] = data.RefundFeeType
	}

	if data.RefundDesc != "" {
		body["refund_desc"] = data.RefundDesc
	}

	if data.RefundAccount != "" {
		body["refund_account"] = data.RefundAccount
	}

	if data.NotifyURL != "" {
		body["notify_url"] = data.NotifyURL
	}

	return r.mch.tlsPost(RefundApplyURL, body, r.options...)
}

// QueryByRefundID 根据微信退款单号查询
func (r *Refund) QueryByRefundID(refundID string, offset ...int) (utils.WXML, error) {
	body := utils.WXML{
		"appid":     r.mch.appid,
		"mch_id":    r.mch.mchid,
		"refund_id": refundID,
		"nonce_str": utils.NonceStr(),
		"sign_type": SignMD5,
	}

	if len(offset) > 0 {
		body["offset"] = strconv.Itoa(offset[0])
	}

	return r.mch.post(RefundQueryURL, body, r.options...)
}

// QueryByOutRefundNO 根据商户退款单号查询
func (r *Refund) QueryByOutRefundNO(outRefundNO string, offset ...int) (utils.WXML, error) {
	body := utils.WXML{
		"appid":         r.mch.appid,
		"mch_id":        r.mch.mchid,
		"out_refund_no": outRefundNO,
		"nonce_str":     utils.NonceStr(),
		"sign_type":     SignMD5,
	}

	if len(offset) > 0 {
		body["offset"] = strconv.Itoa(offset[0])
	}

	return r.mch.post(RefundQueryURL, body, r.options...)
}

// QueryByTransactionID 根据微信订单号查询
func (r *Refund) QueryByTransactionID(transactionID string, offset ...int) (utils.WXML, error) {
	body := utils.WXML{
		"appid":          r.mch.appid,
		"mch_id":         r.mch.mchid,
		"transaction_id": transactionID,
		"nonce_str":      utils.NonceStr(),
		"sign_type":      SignMD5,
	}

	if len(offset) > 0 {
		body["offset"] = strconv.Itoa(offset[0])
	}

	return r.mch.post(RefundQueryURL, body, r.options...)
}

// QueryByOutTradeNO 根据商户订单号查询
func (r *Refund) QueryByOutTradeNO(outTradeNO string, offset ...int) (utils.WXML, error) {
	body := utils.WXML{
		"appid":        r.mch.appid,
		"mch_id":       r.mch.mchid,
		"out_trade_no": outTradeNO,
		"nonce_str":    utils.NonceStr(),
		"sign_type":    "MD5",
	}

	if len(offset) > 0 {
		body["offset"] = strconv.Itoa(offset[0])
	}

	return r.mch.post(RefundQueryURL, body, r.options...)
}
