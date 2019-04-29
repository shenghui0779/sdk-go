package mch

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/iiinsomnia/gochat/utils"
)

// OrderData 统一下单数据
type OrderData struct {
	TradeType  string
	OutTradeNO string
	TotalFee   int
	ClientIP   string
	OpenID     string
	NotifyURL  string
	Body       string
}

// Order 订单操作
type Order struct {
	appid  string
	mchid  string
	apikey string
	client *utils.HTTPClient
	result utils.WXML
}

// Unify 统一下单
func (o *Order) Unify(data *OrderData) error {
	body := utils.WXML{
		"appid":            o.appid,
		"mch_id":           o.mchid,
		"nonce_str":        utils.NonceStr(),
		"sign_type":        "MD5",
		"fee_type":         "CNY",
		"trade_type":       data.TradeType,
		"notify_url":       data.NotifyURL,
		"body":             data.Body,
		"out_trade_no":     data.OutTradeNO,
		"total_fee":        strconv.Itoa(data.TotalFee),
		"spbill_create_ip": data.ClientIP,
		"openid":           data.OpenID,
	}

	return o.do("https://api.mch.weixin.qq.com/pay/unifiedorder", body)
}

// QueryByTransactionID 根据微信订单号查询
func (o *Order) QueryByTransactionID(transactionID string) error {
	body := utils.WXML{
		"appid":          o.appid,
		"mch_id":         o.mchid,
		"transaction_id": transactionID,
		"nonce_str":      utils.NonceStr(),
		"sign_type":      "MD5",
	}

	return o.do("https://api.mch.weixin.qq.com/pay/orderquery", body)
}

// QueryByOutTradeNO 根据商户订单号查询
func (o *Order) QueryByOutTradeNO(outTradeNO string) error {
	body := utils.WXML{
		"appid":        o.appid,
		"mch_id":       o.mchid,
		"out_trade_no": outTradeNO,
		"nonce_str":    utils.NonceStr(),
		"sign_type":    "MD5",
	}

	return o.do("https://api.mch.weixin.qq.com/pay/orderquery", body)
}

// Close 关闭订单【注意：订单生成后不能马上调用关单接口，最短调用时间间隔为5分钟。】
func (o *Order) Close(outTradeNO string) error {
	body := utils.WXML{
		"appid":        o.appid,
		"mch_id":       o.mchid,
		"out_trade_no": outTradeNO,
		"nonce_str":    utils.NonceStr(),
		"sign_type":    "MD5",
	}

	return o.do("https://api.mch.weixin.qq.com/pay/closeorder", body)
}

func (o *Order) do(url string, body utils.WXML) error {
	body["sign"] = utils.PaySign(body, o.apikey)

	resp, err := o.client.PostXML(url, body)

	if err != nil {
		return err
	}

	if resp["return_code"] != "SUCCESS" {
		return errors.New(resp["return_msg"])
	}

	if resp["result_code"] != "SUCCESS" {
		return errors.New(resp["err_code_des"])
	}

	o.result = resp

	return nil
}

// Result 订单操作结果
func (o *Order) Result() utils.WXML {
	return o.result
}

// PrepayID 预付款ID
func (o *Order) PrepayID() string {
	prepayID, ok := o.result["prepay_id"]

	if !ok {
		return ""
	}

	return prepayID
}

// JSAPI 用于JS拉起支付
func (o *Order) JSAPI() utils.WXML {
	ch := utils.WXML{
		"appId":     o.appid,
		"nonceStr":  utils.NonceStr(),
		"package":   fmt.Sprintf("prepay_id=%s", o.PrepayID()),
		"signType":  "MD5",
		"timeStamp": strconv.FormatInt(time.Now().Unix(), 10),
	}

	ch["paySign"] = utils.PaySign(ch, o.apikey)

	return ch
}

// APPAPI 用于APP拉起支付
func (o *Order) APPAPI() utils.WXML {
	ch := utils.WXML{
		"appid":     o.appid,
		"partnerid": o.mchid,
		"prepayid":  o.PrepayID(),
		"package":   "Sign=WXPay",
		"noncestr":  utils.NonceStr(),
		"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
	}

	ch["sign"] = utils.PaySign(ch, o.apikey)

	return ch
}
