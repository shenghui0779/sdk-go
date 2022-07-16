package mch

import (
	"strconv"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

// ParamsRefund 退款参数
type ParamsRefund struct {
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

// RefundByTransactionID 根据微信订单号退款（需要证书）
// 注意：一笔退款失败后重新提交，请不要更换退款单号，请使用原商户退款单号。
func RefundByTransactionID(appid, transactionID string, params *ParamsRefund) wx.Action {
	return wx.NewPostAction(urls.MchRefundApply,
		wx.WithTLS(),
		wx.WithWXML(func(mchid, apikey, nonce string) (wx.WXML, error) {
			m := wx.WXML{
				"appid":          appid,
				"mch_id":         mchid,
				"nonce_str":      nonce,
				"transaction_id": transactionID,
				"out_refund_no":  params.OutRefundNO,
				"total_fee":      strconv.Itoa(params.TotalFee),
				"refund_fee":     strconv.Itoa(params.RefundFee),
			}

			if params.RefundFeeType != "" {
				m["refund_fee_type"] = params.RefundFeeType
			}

			if params.RefundDesc != "" {
				m["refund_desc"] = params.RefundDesc
			}

			if params.RefundAccount != "" {
				m["refund_account"] = params.RefundAccount
			}

			if params.NotifyURL != "" {
				m["notify_url"] = params.NotifyURL
			}

			// 签名
			m["sign"] = wx.SignMD5.Do(apikey, m, true)

			return m, nil
		}),
	)
}

// RefundByOutTradeNO 根据商户订单号退款（需要证书）
// 注意：一笔退款失败后重新提交，请不要更换退款单号，请使用原商户退款单号。
func RefundByOutTradeNO(appid, outTradeNO string, params *ParamsRefund) wx.Action {
	return wx.NewPostAction(urls.MchRefundApply,
		wx.WithTLS(),
		wx.WithWXML(func(mchid, apikey, nonce string) (wx.WXML, error) {
			m := wx.WXML{
				"appid":         appid,
				"mch_id":        mchid,
				"nonce_str":     nonce,
				"out_trade_no":  outTradeNO,
				"out_refund_no": params.OutRefundNO,
				"total_fee":     strconv.Itoa(params.TotalFee),
				"refund_fee":    strconv.Itoa(params.RefundFee),
			}

			if params.RefundFeeType != "" {
				m["refund_fee_type"] = params.RefundFeeType
			}

			if params.RefundDesc != "" {
				m["refund_desc"] = params.RefundDesc
			}

			if params.RefundAccount != "" {
				m["refund_account"] = params.RefundAccount
			}

			if params.NotifyURL != "" {
				m["notify_url"] = params.NotifyURL
			}

			// 签名
			m["sign"] = wx.SignMD5.Do(apikey, m, true)

			return m, nil
		}),
	)
}

// QueryRefundByRefundID 根据微信退款单号查询退款信息
func QueryRefundByRefundID(appid, refundID string, offset ...int) wx.Action {
	return wx.NewPostAction(urls.MchRefundQuery,
		wx.WithWXML(func(mchid, apikey, nonce string) (wx.WXML, error) {
			m := wx.WXML{
				"appid":     appid,
				"mch_id":    mchid,
				"refund_id": refundID,
				"nonce_str": nonce,
			}

			if len(offset) != 0 {
				m["offset"] = strconv.Itoa(offset[0])
			}

			// 签名
			m["sign"] = wx.SignMD5.Do(apikey, m, true)

			return m, nil
		}),
	)
}

// QueryRefundByOutRefundNO 根据商户退款单号查询退款信息
func QueryRefundByOutRefundNO(appid, outRefundNO string, offset ...int) wx.Action {
	return wx.NewPostAction(urls.MchRefundQuery,
		wx.WithWXML(func(mchid, apikey, nonce string) (wx.WXML, error) {
			m := wx.WXML{
				"appid":         appid,
				"mch_id":        mchid,
				"out_refund_no": outRefundNO,
				"nonce_str":     nonce,
			}

			if len(offset) != 0 {
				m["offset"] = strconv.Itoa(offset[0])
			}

			// 签名
			m["sign"] = wx.SignMD5.Do(apikey, m, true)

			return m, nil
		}),
	)
}

// QueryRefundByTransactionID 根据微信订单号查询退款信息
func QueryRefundByTransactionID(appid, transactionID string, offset ...int) wx.Action {
	return wx.NewPostAction(urls.MchRefundQuery,
		wx.WithWXML(func(mchid, apikey, nonce string) (wx.WXML, error) {
			m := wx.WXML{
				"appid":          appid,
				"mch_id":         mchid,
				"transaction_id": transactionID,
				"nonce_str":      nonce,
			}

			if len(offset) != 0 {
				m["offset"] = strconv.Itoa(offset[0])
			}

			// 签名
			m["sign"] = wx.SignMD5.Do(apikey, m, true)

			return m, nil
		}),
	)
}

// QueryRefundByOutTradeNO 根据商户订单号查询退款信息
func QueryRefundByOutTradeNO(appid, outTradeNO string, offset ...int) wx.Action {
	return wx.NewPostAction(urls.MchRefundQuery,
		wx.WithWXML(func(mchid, apikey, nonce string) (wx.WXML, error) {
			m := wx.WXML{
				"appid":        appid,
				"mch_id":       mchid,
				"out_trade_no": outTradeNO,
				"nonce_str":    nonce,
			}

			if len(offset) != 0 {
				m["offset"] = strconv.Itoa(offset[0])
			}

			// 签名
			m["sign"] = wx.SignMD5.Do(apikey, m, true)

			return m, nil
		}),
	)
}
