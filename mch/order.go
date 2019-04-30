package mch

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/iiinsomnia/gochat/utils"
)

// UnifiedOrder 统一下单数据
type UnifiedOrder struct {
	// 必填参数
	OutTradeNO     string // 商户系统内部的订单号，32个字符内、可包含字母，其他说明见商户订单号
	TotalFee       int    // 订单总金额，单位为分，详见支付金额
	SpbillCreateIP string // APP和网页支付提交用户端ip，Native支付填调用微信支付API的机器IP
	TradeType      string // 取值如下：JSAPI，NATIVE，APP，详细说明见参数规定
	Body           string // 商品或支付单简要描述
	NotifyURL      string // 接收微信支付异步通知回调地址，通知url必须为直接可访问的url，不能携带参数
	// 选填参数
	DeviceInfo string // 终端设备号(门店号或收银设备ID)，注意：PC网页或公众号内支付请传"WEB"
	Detail     string // 商品名称明细列表
	Attach     string // 附加数据，在查询API和支付通知中原样返回，该字段主要用于商户携带订单的自定义数据
	FeeType    string // 符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	TimeStart  string // 订单生成时间，格式为yyyyMMddHHmmss，如：2009年12月25日9点10分10秒 表示为：20091225091010。其他详见时间规则
	TimeExpire string // 订单失效时间，格式为yyyyMMddHHmmss，如：2009年12月27日9点10分10秒 表示为：20091227091010。其他详见时间规则
	GoodsTag   string // 商品标记，代金券或立减优惠功能的参数，说明详见代金券或立减优惠
	ProductID  string // trade_type=NATIVE，此参数必传。此id为二维码中包含的商品ID，商户自行定义
	LimitPay   string // no_credit--指定不能使用信用卡支付
	OpenID     string // trade_type=JSAPI，此参数必传，用户在商户appid下的唯一标识
	SubOpenID  string // trade_type=JSAPI，此参数必传，用户在子商户appid下的唯一标识。openid和sub_openid可以选传其中之一，如果选择传sub_openid,则必须传sub_appid
	SceneInfo  string // 该字段用于上报支付的场景信息，针对H5支付有以下三种场景,请根据对应场景上报,H5支付不建议在APP端使用，针对场景1、2请接入APP支付，不然可能会出现兼容性问题
}

// Order 订单操作
type Order struct {
	appid  string
	mchid  string
	apikey string
	client *utils.HTTPClient
}

// Unify 统一下单
func (o *Order) Unify(data *UnifiedOrder) (utils.WXML, error) {
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
		"spbill_create_ip": data.SpbillCreateIP,
	}

	if data.DeviceInfo != "" {
		body["device_info"] = data.DeviceInfo
	}

	if data.Detail != "" {
		body["detail"] = data.Detail
	}

	if data.Attach != "" {
		body["attach"] = data.Attach
	}

	if data.FeeType != "" {
		body["fee_type"] = data.FeeType
	}

	if data.TimeStart != "" {
		body["time_start"] = data.TimeStart
	}

	if data.TimeExpire != "" {
		body["time_expire"] = data.TimeExpire
	}

	if data.GoodsTag != "" {
		body["goods_tag"] = data.GoodsTag
	}

	if data.ProductID != "" {
		body["product_id"] = data.ProductID
	}

	if data.LimitPay != "" {
		body["limit_pay"] = data.LimitPay
	}

	if data.OpenID != "" {
		body["openid"] = data.OpenID
	}

	if data.SubOpenID != "" {
		body["sub_openid"] = data.SubOpenID
	}

	if data.SceneInfo != "" {
		body["scene_info"] = data.SceneInfo
	}

	return o.do("https://api.mch.weixin.qq.com/pay/unifiedorder", body)
}

// QueryByTransactionID 根据微信订单号查询
func (o *Order) QueryByTransactionID(transactionID string) (utils.WXML, error) {
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
func (o *Order) QueryByOutTradeNO(outTradeNO string) (utils.WXML, error) {
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
func (o *Order) Close(outTradeNO string) (utils.WXML, error) {
	body := utils.WXML{
		"appid":        o.appid,
		"mch_id":       o.mchid,
		"out_trade_no": outTradeNO,
		"nonce_str":    utils.NonceStr(),
		"sign_type":    "MD5",
	}

	return o.do("https://api.mch.weixin.qq.com/pay/closeorder", body)
}

func (o *Order) do(url string, body utils.WXML) (utils.WXML, error) {
	body["sign"] = Sign(body, o.apikey)

	resp, err := o.client.PostXML(url, body)

	if err != nil {
		return nil, err
	}

	if resp["return_code"] != "SUCCESS" {
		return nil, errors.New(resp["return_msg"])
	}

	if resp["result_code"] != "SUCCESS" {
		return nil, errors.New(resp["err_code_des"])
	}

	if signature := Sign(resp, o.apikey); signature != resp["sign"] {
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

// JSAPI 用于JS拉起支付
func (o *Order) JSAPI(prepayID string) utils.WXML {
	ch := utils.WXML{
		"appId":     o.appid,
		"nonceStr":  utils.NonceStr(),
		"package":   fmt.Sprintf("prepay_id=%s", prepayID),
		"signType":  "MD5",
		"timeStamp": strconv.FormatInt(time.Now().Unix(), 10),
	}

	ch["paySign"] = Sign(ch, o.apikey)

	return ch
}

// APPAPI 用于APP拉起支付
func (o *Order) APPAPI(prepayID string) utils.WXML {
	ch := utils.WXML{
		"appid":     o.appid,
		"partnerid": o.mchid,
		"prepayid":  prepayID,
		"package":   "Sign=WXPay",
		"noncestr":  utils.NonceStr(),
		"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
	}

	ch["sign"] = Sign(ch, o.apikey)

	return ch
}
