package mch

import (
	"errors"

	"github.com/iiinsomnia/gochat"
	"github.com/iiinsomnia/gochat/utils"
)

// OrderQuery 订单查询
type OrderQuery struct {
	TransactionID string
	OutTradeNO    string
	Channel       gochat.WXChannel
	wxResp        utils.WXML
}

// Do 订单查询
func (o *OrderQuery) Do() error {
	cfg := gochat.GetConfigWithChannel(o.Channel)

	body := utils.WXML{
		"appid":          cfg.AppID,
		"mch_id":         cfg.MchID,
		"nonce_str":      utils.NonceStr(),
		"sign_type":      "MD5",
		"transaction_id": o.TransactionID,
		"out_trade_no":   o.OutTradeNO,
	}

	body["sign"] = utils.PaySign(body, cfg.ApiKey)

	resp, err := utils.PostXML("https://api.mch.weixin.qq.com/pay/orderquery", body)

	if err != nil {
		return err
	}

	if resp["return_code"] != "SUCCESS" {
		return errors.New(resp["return_msg"])
	}

	if resp["result_code"] != "SUCCESS" {
		return errors.New(resp["err_code_des"])
	}

	o.wxResp = resp

	return nil
}

// Result 订单查询结果
func (o *OrderQuery) Result() utils.WXML {
	return o.wxResp
}
