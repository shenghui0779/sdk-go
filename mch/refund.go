package mch

import "strconv"

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
func RefundByTransactionID(transactionID string, data *RefundData) Action {
	f := func(appid, mchid, apikey, nonce string) (WXML, error) {
		body := WXML{
			"appid":          appid,
			"mch_id":         mchid,
			"nonce_str":      nonce,
			"sign_type":      SignMD5,
			"transaction_id": transactionID,
			"out_refund_no":  data.OutRefundNO,
			"total_fee":      strconv.Itoa(data.TotalFee),
			"refund_fee":     strconv.Itoa(data.RefundFee),
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

		body["sign"] = SignWithMD5(body, apikey, true)

		return body, nil
	}

	return &WechatAPI{
		wxml: f,
		url:  RefundApplyURL,
		tls:  true,
	}
}

// RefundByOutTradeNO 根据商户订单号退款
func RefundByOutTradeNO(outTradeNO string, data *RefundData) Action {
	f := func(appid, mchid, apikey, nonce string) (WXML, error) {
		body := WXML{
			"appid":         appid,
			"mch_id":        mchid,
			"nonce_str":     nonce,
			"sign_type":     SignMD5,
			"out_trade_no":  outTradeNO,
			"out_refund_no": data.OutRefundNO,
			"total_fee":     strconv.Itoa(data.TotalFee),
			"refund_fee":    strconv.Itoa(data.RefundFee),
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

		body["sign"] = SignWithMD5(body, apikey, true)

		return body, nil
	}

	return &WechatAPI{
		wxml: f,
		url:  RefundApplyURL,
		tls:  true,
	}
}

// QueryRefundByRefundID 根据微信退款单号查询退款信息
func QueryRefundByRefundID(refundID string, offset ...int) Action {
	f := func(appid, mchid, apikey, nonce string) (WXML, error) {
		body := WXML{
			"appid":     appid,
			"mch_id":    mchid,
			"refund_id": refundID,
			"nonce_str": nonce,
			"sign_type": SignMD5,
		}

		if len(offset) > 0 {
			body["offset"] = strconv.Itoa(offset[0])
		}

		body["sign"] = SignWithMD5(body, apikey, true)

		return body, nil
	}

	return &WechatAPI{
		wxml: f,
		url:  RefundQueryURL,
	}
}

// QueryRefundByOutRefundNO 根据商户退款单号查询退款信息
func QueryRefundByOutRefundNO(outRefundNO string, offset ...int) Action {
	f := func(appid, mchid, apikey, nonce string) (WXML, error) {
		body := WXML{
			"appid":         appid,
			"mch_id":        mchid,
			"out_refund_no": outRefundNO,
			"nonce_str":     nonce,
			"sign_type":     SignMD5,
		}

		if len(offset) > 0 {
			body["offset"] = strconv.Itoa(offset[0])
		}

		body["sign"] = SignWithMD5(body, apikey, true)

		return body, nil
	}

	return &WechatAPI{
		wxml: f,
		url:  RefundQueryURL,
	}
}

// QueryRefundByTransactionID 根据微信订单号查询退款信息
func QueryRefundByTransactionID(transactionID string, offset ...int) Action {
	f := func(appid, mchid, apikey, nonce string) (WXML, error) {
		body := WXML{
			"appid":          appid,
			"mch_id":         mchid,
			"transaction_id": transactionID,
			"nonce_str":      nonce,
			"sign_type":      SignMD5,
		}

		if len(offset) > 0 {
			body["offset"] = strconv.Itoa(offset[0])
		}

		body["sign"] = SignWithMD5(body, apikey, true)

		return body, nil
	}

	return &WechatAPI{
		wxml: f,
		url:  RefundQueryURL,
	}
}

// QueryRefundByOutTradeNO 根据商户订单号查询退款信息
func QueryRefundByOutTradeNO(outTradeNO string, offset ...int) Action {
	f := func(appid, mchid, apikey, nonce string) (WXML, error) {
		body := WXML{
			"appid":        appid,
			"mch_id":       mchid,
			"out_trade_no": outTradeNO,
			"nonce_str":    nonce,
			"sign_type":    "MD5",
		}

		if len(offset) > 0 {
			body["offset"] = strconv.Itoa(offset[0])
		}

		body["sign"] = SignWithMD5(body, apikey, true)

		return body, nil
	}

	return &WechatAPI{
		wxml: f,
		url:  RefundQueryURL,
	}
}
