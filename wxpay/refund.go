package wxpay

import (
	"errors"

	"github.com/iiinsomnia/yiigo"
	"go.uber.org/zap"
	"meipian.cn/printapi/wechat"
	"meipian.cn/printapi/wechat/utils"
)

// Refund 退款
type Refund struct {
	TransactionID string
	OutTradeNO    string
	OutRefundNO   string
	TotalFee      string
	RefundFee     string
	RefundFeeType string
	RefundDesc    string
	RefundAccount string
	NotifyURL     string
	Channel       string
	wxResp        utils.WXML
}

// Do 申请退款
func (r *Refund) Do() error {
	var client *yiigo.HTTPClient

	switch r.Channel {
	case wechat.WXAPP:
		client = tlsAPPClient
	case wechat.WXPub:
		client = tlsPubClient
	case wechat.WXH5:
		client = tlsH5Client
	case wechat.WXMP:
		client = tlsMiniClient
	}

	settings := wechat.GetSettingsWithChannel(r.Channel)

	body := utils.WXML{
		"appid":           settings.AppID,
		"mch_id":          settings.MchID,
		"nonce_str":       utils.NonceStr(),
		"sign_type":       "MD5",
		"transaction_id":  r.TransactionID,
		"out_trade_no":    r.OutTradeNO,
		"out_refund_no":   r.OutRefundNO,
		"total_fee":       r.TotalFee,
		"refund_fee":      r.RefundFee,
		"refund_fee_type": "CNY",
		"refund_desc":     r.RefundDesc,
		"notify_url":      r.NotifyURL,
	}

	body["sign"] = utils.PaySign(body, settings.ApiKey)

	resp, err := utils.PostXML(client, "https://api.mch.weixin.qq.com/secapi/pay/refund", body)

	if err != nil {
		return err
	}

	if resp["return_code"] != "SUCCESS" {
		yiigo.Logger.Error("do wx refund error", zap.String("transaction_id", r.TransactionID), zap.String("trade_no", r.OutTradeNO), zap.String("error", resp["return_msg"]))

		return errors.New(resp["return_msg"])
	}

	if resp["result_code"] != "SUCCESS" {
		yiigo.Logger.Error("do wx refund error", zap.String("transaction_id", r.TransactionID), zap.String("trade_no", r.OutTradeNO), zap.String("code", resp["err_code"]), zap.String("error", resp["err_code_des"]))

		return errors.New(resp["err_code_des"])
	}

	r.wxResp = resp

	return nil
}

// Result 退款结果
func (r *Refund) Result() utils.WXML {
	return r.wxResp
}

// RefundID 退款ID
func (r *Refund) RefundID() string {
	refundID, ok := r.wxResp["refund_id"]

	if !ok {
		return ""
	}

	return refundID
}
