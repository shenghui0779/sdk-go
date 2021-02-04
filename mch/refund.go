package mch

import (
	"strconv"

	"github.com/shenghui0779/gochat/wx"
)

// RefundData 退款数据
type RefundData struct {
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

// RefundByTransactionID 根据微信订单号退款
func RefundByTransactionID(transactionID string, data *RefundData) wx.Action {
	return wx.NewAction(RefundApplyURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithTLS(),
		wx.WithWXML(func(appid, mchid, nonce string) (wx.WXML, error) {
			body := wx.WXML{
				"appid":          appid,
				"mch_id":         mchid,
				"nonce_str":      nonce,
				"transaction_id": transactionID,
				"out_refund_no":  data.OutRefundNO,
				"total_fee":      strconv.Itoa(data.TotalFee),
				"refund_fee":     strconv.Itoa(data.RefundFee),
				"sign_type":      SignMD5,
			}

			if data.RefundFeeType != "" {
				body["refund_fee_type"] = data.RefundFeeType
			}

			if data.RefundDesc != "" {
				body["refund_desc"] = data.RefundDesc
			}

			if data.RefundAccount != "" {
				body["refund_account"] = data.RefundAccount
			}

			if data.NotifyURL != "" {
				body["notify_url"] = data.NotifyURL
			}

			return body, nil
		}),
	)
}

// RefundByOutTradeNO 根据商户订单号退款
func RefundByOutTradeNO(outTradeNO string, data *RefundData) wx.Action {
	return wx.NewAction(RefundApplyURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithTLS(),
		wx.WithWXML(func(appid, mchid, nonce string) (wx.WXML, error) {
			body := wx.WXML{
				"appid":         appid,
				"mch_id":        mchid,
				"nonce_str":     nonce,
				"out_trade_no":  outTradeNO,
				"out_refund_no": data.OutRefundNO,
				"total_fee":     strconv.Itoa(data.TotalFee),
				"refund_fee":    strconv.Itoa(data.RefundFee),
				"sign_type":     SignMD5,
			}

			if data.RefundFeeType != "" {
				body["refund_fee_type"] = data.RefundFeeType
			}

			if data.RefundDesc != "" {
				body["refund_desc"] = data.RefundDesc
			}

			if data.RefundAccount != "" {
				body["refund_account"] = data.RefundAccount
			}

			if data.NotifyURL != "" {
				body["notify_url"] = data.NotifyURL
			}

			return body, nil
		}),
	)
}

// QueryRefundByRefundID 根据微信退款单号查询退款信息
func QueryRefundByRefundID(refundID string, offset ...int) wx.Action {
	return wx.NewAction(RefundQueryURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithWXML(func(appid, mchid, nonce string) (wx.WXML, error) {
			body := wx.WXML{
				"appid":     appid,
				"mch_id":    mchid,
				"refund_id": refundID,
				"nonce_str": nonce,
				"sign_type": SignMD5,
			}

			if len(offset) != 0 {
				body["offset"] = strconv.Itoa(offset[0])
			}

			return body, nil
		}),
	)
}

// QueryRefundByOutRefundNO 根据商户退款单号查询退款信息
func QueryRefundByOutRefundNO(outRefundNO string, offset ...int) wx.Action {
	return wx.NewAction(RefundQueryURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithWXML(func(appid, mchid, nonce string) (wx.WXML, error) {
			body := wx.WXML{
				"appid":         appid,
				"mch_id":        mchid,
				"out_refund_no": outRefundNO,
				"nonce_str":     nonce,
				"sign_type":     SignMD5,
			}

			if len(offset) != 0 {
				body["offset"] = strconv.Itoa(offset[0])
			}

			return body, nil
		}),
	)
}

// QueryRefundByTransactionID 根据微信订单号查询退款信息
func QueryRefundByTransactionID(transactionID string, offset ...int) wx.Action {
	return wx.NewAction(RefundQueryURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithWXML(func(appid, mchid, nonce string) (wx.WXML, error) {
			body := wx.WXML{
				"appid":          appid,
				"mch_id":         mchid,
				"transaction_id": transactionID,
				"nonce_str":      nonce,
				"sign_type":      SignMD5,
			}

			if len(offset) != 0 {
				body["offset"] = strconv.Itoa(offset[0])
			}

			return body, nil
		}),
	)
}

// QueryRefundByOutTradeNO 根据商户订单号查询退款信息
func QueryRefundByOutTradeNO(outTradeNO string, offset ...int) wx.Action {
	return wx.NewAction(RefundQueryURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithWXML(func(appid, mchid, nonce string) (wx.WXML, error) {
			body := wx.WXML{
				"appid":        appid,
				"mch_id":       mchid,
				"out_trade_no": outTradeNO,
				"nonce_str":    nonce,
				"sign_type":    SignMD5,
			}

			if len(offset) != 0 {
				body["offset"] = strconv.Itoa(offset[0])
			}

			return body, nil
		}),
	)
}
