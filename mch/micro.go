package mch

import (
	"strconv"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsMicroPay struct {
	// 必填参数
	OutTradeNO     string // 商户系统内部的订单号，32个字符内、可包含字母，其他说明见商户订单号
	TotalFee       int    // 订单总金额，单位为分，详见支付金额
	SpbillCreateIP string // APP和网页支付提交用户端ip，Native支付填调用微信支付API的机器IP
	AuthCode       string // 扫码支付付款码，设备读取用户微信中的条码或者二维码信息（用户付款码规则：18位纯数字，前缀以10、11、12、13、14、15开头）
	Body           string // 商品或支付单简要描述
	// 选填参数
	DeviceInfo    string // 终端设备号(门店号或收银设备ID)，注意：PC网页或公众号内支付请传"WEB"
	Detail        string // 商品名称明细列表
	Attach        string // 附加数据，在查询API和支付通知中原样返回，该字段主要用于商户携带订单的自定义数据
	FeeType       string // 符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	TimeStart     string // 订单生成时间，格式为yyyyMMddHHmmss，如：2009年12月25日9点10分10秒 表示为：20091225091010
	TimeExpire    string // 订单失效时间，格式为yyyyMMddHHmmss，如：2009年12月27日9点10分10秒 表示为：20091227091010
	GoodsTag      string // 商品标记，代金券或立减优惠功能的参数，说明详见代金券或立减优惠
	LimitPay      string // no_credit--指定不能使用信用卡支付
	Receipt       bool   // 是否在支付成功消息和支付详情页中出现开票入口，注：需要在微信支付商户平台或微信公众平台开通电子发票功能
	ProfitSharing bool   // 是否需要分账：Y-是；N-否；默认不分账
	SceneInfo     string // 该字段用于上报支付的场景信息
}

// MicroPay 扫码支付
// 【提醒1】提交支付请求后微信会同步返回支付结果。当返回结果为“系统错误”时，商户系统等待5秒后调用「查询订单API」，查询支付实际交易结果；当返回结果为“USERPAYING”时，商户系统可设置间隔时间(建议10秒)重新查询支付结果，直到支付成功或超时(建议45秒)；
// 【提醒2】在调用查询接口返回后，如果交易状况不明晰，请调用「撤销订单API」，此时如果交易失败则关闭订单，该单不能再支付成功；如果交易成功，则将扣款退回到用户账户。当撤销无返回或错误时，请再次调用。注意：请勿扣款后立即调用「撤销订单API」，建议至少15秒后再调用。
func MicroPay(appid string, params *ParamsMicroPay, options ...SLOption) wx.Action {
	return wx.NewPostAction(urls.MchMicroPay,
		wx.WithWXML(func(mchid, apikey, nonce string) (wx.WXML, error) {
			m := wx.WXML{
				"appid":            appid,
				"mch_id":           mchid,
				"nonce_str":        nonce,
				"auth_code":        params.AuthCode,
				"fee_type":         "CNY",
				"body":             params.Body,
				"out_trade_no":     params.OutTradeNO,
				"total_fee":        strconv.Itoa(params.TotalFee),
				"spbill_create_ip": params.SpbillCreateIP,
			}

			for _, f := range options {
				f(m)
			}

			if len(params.DeviceInfo) != 0 {
				m["device_info"] = params.DeviceInfo
			}

			if len(params.Detail) != 0 {
				m["detail"] = params.Detail
			}

			if len(params.Attach) != 0 {
				m["attach"] = params.Attach
			}

			if len(params.FeeType) != 0 {
				m["fee_type"] = params.FeeType
			}

			if len(params.TimeStart) != 0 {
				m["time_start"] = params.TimeStart
			}

			if len(params.TimeExpire) != 0 {
				m["time_expire"] = params.TimeExpire
			}

			if len(params.GoodsTag) != 0 {
				m["goods_tag"] = params.GoodsTag
			}

			if len(params.LimitPay) != 0 {
				m["limit_pay"] = params.LimitPay
			}

			if params.Receipt {
				m["receipt"] = "Y"
			}

			if params.ProfitSharing {
				m["profit_sharing"] = "Y"
			}

			if len(params.SceneInfo) != 0 {
				m["scene_info"] = params.SceneInfo
			}

			// 签名
			m["sign"] = wx.SignMD5.Do(apikey, m, true)

			return m, nil
		}))
}

// ReverseByTransactionID 撤销订单
// 支付交易返回失败或支付系统超时，调用该接口撤销交易。如果此订单用户支付失败，微信支付系统会将此订单关闭；如果用户支付成功，微信支付系统会将此订单资金退还给用户。
// 【注意】7天以内的交易单可调用撤销，其他正常支付的单如需实现相同功能请调用申请退款API。提交支付交易后调用「查询订单API」，没有明确的支付结果再调用「撤销订单API」。
func ReverseByTransactionID(appid, transactionID string, options ...SLOption) wx.Action {
	return wx.NewPostAction(urls.MchOrderReverse,
		wx.WithTLS(),
		wx.WithWXML(func(mchid, apikey, nonce string) (wx.WXML, error) {
			m := wx.WXML{
				"appid":          appid,
				"mch_id":         mchid,
				"transaction_id": transactionID,
				"nonce_str":      nonce,
			}

			for _, f := range options {
				f(m)
			}

			// 签名
			m["sign"] = wx.SignMD5.Do(apikey, m, true)

			return m, nil
		}),
	)
}

// ReverseByOutTradeNO 撤销订单
// 支付交易返回失败或支付系统超时，调用该接口撤销交易。如果此订单用户支付失败，微信支付系统会将此订单关闭；如果用户支付成功，微信支付系统会将此订单资金退还给用户。
// 【注意】7天以内的交易单可调用撤销，其他正常支付的单如需实现相同功能请调用申请退款API。提交支付交易后调用「查询订单API」，没有明确的支付结果再调用「撤销订单API」。
func ReverseByOutTradeNO(appid, outTradeNO string, options ...SLOption) wx.Action {
	return wx.NewPostAction(urls.MchOrderReverse,
		wx.WithTLS(),
		wx.WithWXML(func(mchid, apikey, nonce string) (wx.WXML, error) {
			m := wx.WXML{
				"appid":        appid,
				"mch_id":       mchid,
				"out_trade_no": outTradeNO,
				"nonce_str":    nonce,
			}

			for _, f := range options {
				f(m)
			}

			// 签名
			m["sign"] = wx.SignMD5.Do(apikey, m, true)

			return m, nil
		}),
	)
}
