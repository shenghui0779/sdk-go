package mch

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/iiinsomnia/gochat/utils"
)

// UnifiedOrder 统一下单数据
type UnifiedOrder struct {
	// 必填参数
	OutTradeNO     string // 商户系统内部的订单号，32个字符内、可包含字母，其他说明见商户订单号
	TotalFee       int    // 订单总金额，单位为分，详见支付金额
	SpbillCreateIP string // APP和网页支付提交用户端ip，Native支付填调用微信支付API的机器IP
	TradeType      string // 取值如下：JSAPI，NATIVE，APP，MWEB，详细说明见参数规定
	Body           string // 商品或支付单简要描述
	NotifyURL      string // 接收微信支付异步通知回调地址，通知url必须为直接可访问的url，不能携带参数
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
	Receipt    bool   // 是否在支付成功消息和支付详情页中出现开票入口，注：需要在微信支付商户平台或微信公众平台开通电子发票功能
	SceneInfo  string // 该字段用于上报支付的场景信息
}

// Order 订单操作
type Order struct {
	appid  string
	mchid  string
	apikey string
	client *utils.HTTPClient
}

// Unify 统一下单
func (o *Order) Unify(order *UnifiedOrder) (utils.WXML, error) {
	body := utils.WXML{
		"appid":            o.appid,
		"mch_id":           o.mchid,
		"nonce_str":        utils.NonceStr(),
		"sign_type":        SignMD5,
		"fee_type":         "CNY",
		"trade_type":       order.TradeType,
		"notify_url":       order.NotifyURL,
		"body":             order.Body,
		"out_trade_no":     order.OutTradeNO,
		"total_fee":        strconv.Itoa(order.TotalFee),
		"spbill_create_ip": order.SpbillCreateIP,
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

	if order.Receipt {
		body["receipt"] = "Y"
	}

	if order.SceneInfo != "" {
		body["scene_info"] = order.SceneInfo
	}

	return o.do(OrderUnifyURL, body)
}

// QueryByTransactionID 根据微信订单号查询
func (o *Order) QueryByTransactionID(transactionID string) (utils.WXML, error) {
	body := utils.WXML{
		"appid":          o.appid,
		"mch_id":         o.mchid,
		"transaction_id": transactionID,
		"nonce_str":      utils.NonceStr(),
		"sign_type":      SignMD5,
	}

	return o.do(OrderQueryURL, body)
}

// QueryByOutTradeNO 根据商户订单号查询
func (o *Order) QueryByOutTradeNO(outTradeNO string) (utils.WXML, error) {
	body := utils.WXML{
		"appid":        o.appid,
		"mch_id":       o.mchid,
		"out_trade_no": outTradeNO,
		"nonce_str":    utils.NonceStr(),
		"sign_type":    SignMD5,
	}

	return o.do(OrderQueryURL, body)
}

// Close 关闭订单【注意：订单生成后不能马上调用关单接口，最短调用时间间隔为5分钟。】
func (o *Order) Close(outTradeNO string) (utils.WXML, error) {
	body := utils.WXML{
		"appid":        o.appid,
		"mch_id":       o.mchid,
		"out_trade_no": outTradeNO,
		"nonce_str":    utils.NonceStr(),
		"sign_type":    SignMD5,
	}

	return o.do(OrderCloseURL, body)
}

func (o *Order) do(url string, body utils.WXML) (utils.WXML, error) {
	body["sign"] = SignWithMD5(body, o.apikey)

	resp, err := o.client.PostXML(url, body)

	if err != nil {
		return nil, err
	}

	if resp["return_code"] != ReplySuccess {
		return nil, errors.New(resp["return_msg"])
	}

	if resp["result_code"] != ReplySuccess {
		return nil, errors.New(resp["err_code_des"])
	}

	if signature := SignWithMD5(resp, o.apikey); signature != resp["sign"] {
		return nil, fmt.Errorf("order resp signature verified failed, want: %s, got: %s", signature, resp["sign"])
	}

	if resp["appid"] != o.appid {
		return nil, fmt.Errorf("order resp appid mismatch, want: %s, got: %s", o.appid, resp["appid"])
	}

	if resp["mch_id"] != o.mchid {
		return nil, fmt.Errorf("order resp mchid mismatch, want: %s, got: %s", o.mchid, resp["mch_id"])
	}

	return resp, nil
}
