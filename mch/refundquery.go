package mch

import (
	"errors"
	"strconv"

	"github.com/iiinsomnia/yiigo"
	"go.uber.org/zap"
	"meipian.cn/printapi/wechat"
	"meipian.cn/printapi/wechat/utils"
)

// RefundQuery 退款查询
type RefundQuery struct {
	RefundID    string
	OutRefundNO string
	Offset      int
	Channel     string
	wxResp      utils.WXML
}

// Do 退款查询
func (r *RefundQuery) Do() error {
	settings := wechat.GetSettingsWithChannel(r.Channel)

	body := utils.WXML{
		"appid":         settings.AppID,
		"mch_id":        settings.MchID,
		"nonce_str":     utils.NonceStr(),
		"sign_type":     "MD5",
		"refund_id":     r.RefundID,
		"out_refund_no": r.OutRefundNO,
	}

	if r.Offset > 0 {
		body["offset"] = strconv.Itoa(r.Offset)
	}

	body["sign"] = utils.PaySign(body, settings.ApiKey)

	resp, err := utils.PostXML(defaultClient, "https://api.mch.weixin.qq.com/pay/refundquery", body)

	if err != nil {
		return err
	}

	if resp["return_code"] != "SUCCESS" {
		yiigo.Logger.Error("do wx refundquery error", zap.String("refund_id", r.RefundID), zap.String("refund_no", r.OutRefundNO), zap.String("error", resp["return_msg"]))

		return errors.New(resp["return_msg"])
	}

	if resp["result_code"] != "SUCCESS" {
		yiigo.Logger.Error("do wx refundquery error", zap.String("refund_id", r.RefundID), zap.String("refund_no", r.OutRefundNO), zap.String("code", resp["err_code"]), zap.String("error", resp["err_code_des"]))

		return errors.New(resp["err_code_des"])
	}

	r.wxResp = resp

	return nil
}

// Result 退款查询结果
func (r *RefundQuery) Result() WXML {
	return r.wxResp
}
