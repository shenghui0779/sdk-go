package mch

import (
	"errors"
	"fmt"
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
	appid     string
	mchid     string
	apikey    string
	client    *utils.HTTPClient
	sslClient *utils.HTTPClient
}

// RefundByTransactionID 根据微信订单号退款
func (r *Refund) RefundByTransactionID(transactionID string, data *RefundData) (utils.WXML, error) {
	body := utils.WXML{
		"appid":          r.appid,
		"mch_id":         r.mchid,
		"nonce_str":      utils.NonceStr(),
		"sign_type":      "MD5",
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

	return r.doSSL("https://api.mch.weixin.qq.com/secapi/pay/refund", body)
}

// RefundByOutTradeNO 根据微信订单号退款
func (r *Refund) RefundByOutTradeNO(outTradeNO string, data *RefundData) (utils.WXML, error) {
	body := utils.WXML{
		"appid":         r.appid,
		"mch_id":        r.mchid,
		"nonce_str":     utils.NonceStr(),
		"sign_type":     "MD5",
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

	return r.doSSL("https://api.mch.weixin.qq.com/secapi/pay/refund", body)
}

// QueryByRefundID 根据微信退款单号查询
func (r *Refund) QueryByRefundID(refundID string, offset ...int) (utils.WXML, error) {
	body := utils.WXML{
		"appid":     r.appid,
		"mch_id":    r.mchid,
		"refund_id": refundID,
		"nonce_str": utils.NonceStr(),
		"sign_type": "MD5",
	}

	if len(offset) > 0 {
		body["offset"] = strconv.Itoa(offset[0])
	}

	return r.do("https://api.mch.weixin.qq.com/pay/refundquery", body)
}

// QueryByOutRefundNO 根据商户退款单号查询
func (r *Refund) QueryByOutRefundNO(outRefundNO string, offset ...int) (utils.WXML, error) {
	body := utils.WXML{
		"appid":         r.appid,
		"mch_id":        r.mchid,
		"out_refund_no": outRefundNO,
		"nonce_str":     utils.NonceStr(),
		"sign_type":     "MD5",
	}

	if len(offset) > 0 {
		body["offset"] = strconv.Itoa(offset[0])
	}

	return r.do("https://api.mch.weixin.qq.com/pay/refundquery", body)
}

// QueryByTransactionID 根据微信订单号查询
func (r *Refund) QueryByTransactionID(transactionID string, offset ...int) (utils.WXML, error) {
	body := utils.WXML{
		"appid":          r.appid,
		"mch_id":         r.mchid,
		"transaction_id": transactionID,
		"nonce_str":      utils.NonceStr(),
		"sign_type":      "MD5",
	}

	if len(offset) > 0 {
		body["offset"] = strconv.Itoa(offset[0])
	}

	return r.do("https://api.mch.weixin.qq.com/pay/refundquery", body)
}

// QueryByOutTradeNO 根据商户订单号查询
func (r *Refund) QueryByOutTradeNO(outTradeNO string, offset ...int) (utils.WXML, error) {
	body := utils.WXML{
		"appid":        r.appid,
		"mch_id":       r.mchid,
		"out_trade_no": outTradeNO,
		"nonce_str":    utils.NonceStr(),
		"sign_type":    "MD5",
	}

	if len(offset) > 0 {
		body["offset"] = strconv.Itoa(offset[0])
	}

	return r.do("https://api.mch.weixin.qq.com/pay/refundquery", body)
}

func (r *Refund) do(url string, body utils.WXML) (utils.WXML, error) {
	body["sign"] = Sign(body, r.apikey)

	resp, err := r.client.PostXML(url, body)

	if err != nil {
		return nil, err
	}

	if resp["return_code"] != "SUCCESS" {
		return nil, errors.New(resp["return_msg"])
	}

	if resp["result_code"] != "SUCCESS" {
		return nil, errors.New(resp["err_code_des"])
	}

	if signature := Sign(resp, r.apikey); signature != resp["sign"] {
		return nil, fmt.Errorf("refund resp signature verified failed, want: %s, got: %s", signature, resp["sign"])
	}

	if resp["appid"] != r.appid {
		return nil, fmt.Errorf("refund resp appid mismatch, want: %s, got: %s", r.appid, resp["appid"])
	}

	if resp["mch_id"] != r.mchid {
		return nil, fmt.Errorf("refund resp mchid mismatch, want: %s, got: %s", r.mchid, resp["mch_id"])
	}

	return resp, nil
}

func (r *Refund) doSSL(url string, body utils.WXML) (utils.WXML, error) {
	body["sign"] = Sign(body, r.apikey)

	resp, err := r.sslClient.PostXML(url, body)

	if err != nil {
		return nil, err
	}

	if resp["return_code"] != "SUCCESS" {
		return nil, errors.New(resp["return_msg"])
	}

	if resp["result_code"] != "SUCCESS" {
		return nil, errors.New(resp["err_code_des"])
	}

	if signature := Sign(resp, r.apikey); signature != resp["sign"] {
		return nil, fmt.Errorf("refund resp signature verified failed, want: %s, got: %s", signature, resp["sign"])
	}

	if resp["appid"] != r.appid {
		return nil, fmt.Errorf("refund resp appid mismatch, want: %s, got: %s", r.appid, resp["appid"])
	}

	if resp["mch_id"] != r.mchid {
		return nil, fmt.Errorf("refund resp mchid mismatch, want: %s, got: %s", r.mchid, resp["mch_id"])
	}

	return resp, nil
}
