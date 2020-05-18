package mch

import (
	"errors"
	"strconv"

	"github.com/iiinsomnia/gochat/utils"
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
	options []utils.HTTPRequestOption
}

// RefundByTransactionID 根据微信订单号退款
func (r *Refund) RefundByTransactionID(transactionID string, data *RefundData) (utils.WXML, error) {
	body := utils.WXML{
		"appid":          r.mch.AppID,
		"mch_id":         r.mch.MchID,
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

	body["sign"] = SignWithMD5(body, r.mch.ApiKey)

	resp, err := r.mch.SSLClient.PostXML(RefundApplyURL, body, r.options...)

	if err != nil {
		return nil, err
	}

	if resp["return_code"] != ResultSuccess {
		return nil, errors.New(resp["return_msg"])
	}

	if err := r.mch.VerifyWXReply(resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// RefundByOutTradeNO 根据商户订单号退款
func (r *Refund) RefundByOutTradeNO(outTradeNO string, data *RefundData) (utils.WXML, error) {
	body := utils.WXML{
		"appid":         r.mch.AppID,
		"mch_id":        r.mch.MchID,
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

	body["sign"] = SignWithMD5(body, r.mch.ApiKey)

	resp, err := r.mch.SSLClient.PostXML(RefundApplyURL, body, r.options...)

	if err != nil {
		return nil, err
	}

	if resp["return_code"] != ResultSuccess {
		return nil, errors.New(resp["return_msg"])
	}

	if err := r.mch.VerifyWXReply(resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// QueryByRefundID 根据微信退款单号查询
func (r *Refund) QueryByRefundID(refundID string, offset ...int) (utils.WXML, error) {
	body := utils.WXML{
		"appid":     r.mch.AppID,
		"mch_id":    r.mch.MchID,
		"refund_id": refundID,
		"nonce_str": utils.NonceStr(),
		"sign_type": SignMD5,
	}

	if len(offset) > 0 {
		body["offset"] = strconv.Itoa(offset[0])
	}

	body["sign"] = SignWithMD5(body, r.mch.ApiKey)

	resp, err := r.mch.Client.PostXML(RefundQueryURL, body, r.options...)

	if err != nil {
		return nil, err
	}

	if resp["return_code"] != ResultSuccess {
		return nil, errors.New(resp["return_msg"])
	}

	if err := r.mch.VerifyWXReply(resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// QueryByOutRefundNO 根据商户退款单号查询
func (r *Refund) QueryByOutRefundNO(outRefundNO string, offset ...int) (utils.WXML, error) {
	body := utils.WXML{
		"appid":         r.mch.AppID,
		"mch_id":        r.mch.MchID,
		"out_refund_no": outRefundNO,
		"nonce_str":     utils.NonceStr(),
		"sign_type":     SignMD5,
	}

	if len(offset) > 0 {
		body["offset"] = strconv.Itoa(offset[0])
	}

	body["sign"] = SignWithMD5(body, r.mch.ApiKey)

	resp, err := r.mch.Client.PostXML(RefundQueryURL, body, r.options...)

	if err != nil {
		return nil, err
	}

	if resp["return_code"] != ResultSuccess {
		return nil, errors.New(resp["return_msg"])
	}

	if err := r.mch.VerifyWXReply(resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// QueryByTransactionID 根据微信订单号查询
func (r *Refund) QueryByTransactionID(transactionID string, offset ...int) (utils.WXML, error) {
	body := utils.WXML{
		"appid":          r.mch.AppID,
		"mch_id":         r.mch.MchID,
		"transaction_id": transactionID,
		"nonce_str":      utils.NonceStr(),
		"sign_type":      SignMD5,
	}

	if len(offset) > 0 {
		body["offset"] = strconv.Itoa(offset[0])
	}

	body["sign"] = SignWithMD5(body, r.mch.ApiKey)

	resp, err := r.mch.Client.PostXML(RefundQueryURL, body, r.options...)

	if err != nil {
		return nil, err
	}

	if resp["return_code"] != ResultSuccess {
		return nil, errors.New(resp["return_msg"])
	}

	if err := r.mch.VerifyWXReply(resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// QueryByOutTradeNO 根据商户订单号查询
func (r *Refund) QueryByOutTradeNO(outTradeNO string, offset ...int) (utils.WXML, error) {
	body := utils.WXML{
		"appid":        r.mch.AppID,
		"mch_id":       r.mch.MchID,
		"out_trade_no": outTradeNO,
		"nonce_str":    utils.NonceStr(),
		"sign_type":    "MD5",
	}

	if len(offset) > 0 {
		body["offset"] = strconv.Itoa(offset[0])
	}

	body["sign"] = SignWithMD5(body, r.mch.ApiKey)

	resp, err := r.mch.Client.PostXML(RefundQueryURL, body, r.options...)

	if err != nil {
		return nil, err
	}

	if resp["return_code"] != ResultSuccess {
		return nil, errors.New(resp["return_msg"])
	}

	if err := r.mch.VerifyWXReply(resp); err != nil {
		return nil, err
	}

	return resp, nil
}
