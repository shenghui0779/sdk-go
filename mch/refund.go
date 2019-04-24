package mch

import (
	"errors"
	"strconv"

	"github.com/iiinsomnia/gochat/utils"
)

// Refund 退款数据
type RefundData struct {
	OutRefundNO   string
	TotalFee      string
	RefundFee     string
	RefundFeeType string
	RefundDesc    string
	RefundAccount string
	NotifyURL     string
}

// Refund 退款操作
type Refund struct {
	appid     string
	mchid     string
	apikey    string
	sslClient *utils.HTTPClient
	result    utils.WXML
}

// RefundByTransactionID 根据微信订单号退款
func (r *Refund) RefundByTransactionID(transactionID string, data *RefundData) error {
	body := utils.WXML{
		"appid":           r.appid,
		"mch_id":          r.mchid,
		"nonce_str":       utils.NonceStr(),
		"sign_type":       "MD5",
		"transaction_id":  transactionID,
		"out_refund_no":   data.OutRefundNO,
		"total_fee":       data.TotalFee,
		"refund_fee":      data.RefundFee,
		"refund_fee_type": "CNY",
		"refund_desc":     data.RefundDesc,
		"notify_url":      data.NotifyURL,
	}

	return r.doSSL("https://api.mch.weixin.qq.com/secapi/pay/refund", body)
}

// RefundByOutTradeNO 根据微信订单号退款
func (r *Refund) RefundByOutTradeNO(outTradeNO string, data *RefundData) error {
	body := utils.WXML{
		"appid":           r.appid,
		"mch_id":          r.mchid,
		"nonce_str":       utils.NonceStr(),
		"sign_type":       "MD5",
		"out_trade_no":    outTradeNO,
		"out_refund_no":   data.OutRefundNO,
		"total_fee":       data.TotalFee,
		"refund_fee":      data.RefundFee,
		"refund_fee_type": "CNY",
		"refund_desc":     data.RefundDesc,
		"notify_url":      data.NotifyURL,
	}

	return r.doSSL("https://api.mch.weixin.qq.com/secapi/pay/refund", body)
}

// QueryByRefundID 根据微信退款单号查询
func (r *Refund) QueryByRefundID(refundID string, offset ...int) error {
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
func (r *Refund) QueryByOutRefundNO(outRefundNO string, offset ...int) error {
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
func (r *Refund) QueryByTransactionID(transactionID string, offset ...int) error {
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
func (r *Refund) QueryByOutTradeNO(outTradeNO string, offset ...int) error {
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

func (r *Refund) do(url string, body utils.WXML) error {
	body["sign"] = utils.PaySign(body, r.apikey)

	resp, err := utils.HTTPPostXML(url, body)

	if err != nil {
		return err
	}

	if resp["return_code"] != "SUCCESS" {
		return errors.New(resp["return_msg"])
	}

	if resp["result_code"] != "SUCCESS" {
		return errors.New(resp["err_code_des"])
	}

	r.result = resp

	return nil
}

func (r *Refund) doSSL(url string, body utils.WXML) error {
	body["sign"] = utils.PaySign(body, r.apikey)

	resp, err := r.sslClient.PostXML(url, body)

	if err != nil {
		return err
	}

	if resp["return_code"] != "SUCCESS" {
		return errors.New(resp["return_msg"])
	}

	if resp["result_code"] != "SUCCESS" {
		return errors.New(resp["err_code_des"])
	}

	r.result = resp

	return nil
}

// Result 退款结果
func (r *Refund) Result() utils.WXML {
	return r.result
}

// RefundID 退款ID
func (r *Refund) RefundID() string {
	refundID, ok := r.result["refund_id"]

	if !ok {
		return ""
	}

	return refundID
}
