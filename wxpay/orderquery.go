package wxpay

import (
	"errors"

	"github.com/iiinsomnia/yiigo"
	"go.uber.org/zap"
	"meipian.cn/printapi/wechat"
	"meipian.cn/printapi/wechat/utils"
)

// OrderQuery 订单查询
type OrderQuery struct {
	TransactionID string
	OutTradeNO    string
	Channel       string
	wxResp        utils.WXML
}

// Do 订单查询
func (o *OrderQuery) Do() error {
	settings := wechat.GetSettingsWithChannel(o.Channel)

	body := utils.WXML{
		"appid":          settings.AppID,
		"mch_id":         settings.MchID,
		"nonce_str":      utils.NonceStr(),
		"sign_type":      "MD5",
		"transaction_id": o.TransactionID,
		"out_trade_no":   o.OutTradeNO,
	}

	body["sign"] = utils.PaySign(body, settings.ApiKey)

	resp, err := utils.PostXML(defaultClient, "https://api.mch.weixin.qq.com/pay/orderquery", body)

	if err != nil {
		return err
	}

	if resp["return_code"] != "SUCCESS" {
		yiigo.Logger.Error("do wx orderquery error", zap.String("transaction_id", o.TransactionID), zap.String("trade_no", o.OutTradeNO), zap.String("error", resp["return_msg"]))

		return errors.New(resp["return_msg"])
	}

	if resp["result_code"] != "SUCCESS" {
		yiigo.Logger.Error("do wx orderquery error", zap.String("transaction_id", o.TransactionID), zap.String("trade_no", o.OutTradeNO), zap.String("code", resp["err_code"]), zap.String("error", resp["err_code_des"]))

		return errors.New(resp["err_code_des"])
	}

	o.wxResp = resp

	return nil
}

// Result 订单查询结果
func (o *OrderQuery) Result() utils.WXML {
	return o.wxResp
}
