package mch

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/wx"
	"github.com/stretchr/testify/assert"
)

func TestUnifyOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().PostXML(gomock.AssignableToTypeOf(context.TODO()), "https://api.mch.weixin.qq.com/pay/unifiedorder", wx.WXML{
		"appid":            "wx2421b1c4370ec43b",
		"mch_id":           "10000100",
		"nonce_str":        "1add1a30ac87aa2db72f57a2375d8fec",
		"trade_type":       "APP",
		"body":             "APP支付测试",
		"out_trade_no":     "1415659990",
		"total_fee":        "1",
		"fee_type":         "CNY",
		"spbill_create_ip": "14.23.150.211",
		"notify_url":       "http://wxpay.wxutil.com/pub_v2/pay/notify.v2.php",
		"attach":           "支付测试",
		"sign_type":        "MD5",
		"sign":             "7C07373FE5EAEDB936F3E454875C9462",
	}).Return(wx.WXML{
		"return_code": "SUCCESS",
		"return_msg":  "OK",
		"appid":       "wx2421b1c4370ec43b",
		"mch_id":      "10000100",
		"nonce_str":   "IITRi8Iabbblz1Jc",
		"sign":        "E515C9BE3D3129764915407267CA0243",
		"result_code": "SUCCESS",
		"prepay_id":   "wx201411101639507cbf6ffd8b0779950874",
		"trade_type":  "APP",
	}, nil)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "1add1a30ac87aa2db72f57a2375d8fec"
	}
	mch.client = client
	mch.tlsClient = client

	r, err := mch.Do(context.TODO(), UnifyOrder(&OrderData{
		OutTradeNO:     "1415659990",
		TotalFee:       1,
		SpbillCreateIP: "14.23.150.211",
		TradeType:      TradeAPP,
		Body:           "APP支付测试",
		NotifyURL:      "http://wxpay.wxutil.com/pub_v2/pay/notify.v2.php",
		Attach:         "支付测试",
	}))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
		"return_code": "SUCCESS",
		"return_msg":  "OK",
		"appid":       "wx2421b1c4370ec43b",
		"mch_id":      "10000100",
		"nonce_str":   "IITRi8Iabbblz1Jc",
		"sign":        "E515C9BE3D3129764915407267CA0243",
		"result_code": "SUCCESS",
		"prepay_id":   "wx201411101639507cbf6ffd8b0779950874",
		"trade_type":  "APP",
	}, r)
}

func TestQueryOrderByTransactionID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().PostXML(gomock.AssignableToTypeOf(context.TODO()), "https://api.mch.weixin.qq.com/pay/orderquery", wx.WXML{
		"appid":          "wx2421b1c4370ec43b",
		"mch_id":         "10000100",
		"transaction_id": "1008450740201411110005820873",
		"nonce_str":      "ec2316275641faa3aacf3cc599e8730f",
		"sign_type":      "MD5",
		"sign":           "CA9B10C422366B6647827F0E6C18A4D8",
	}).Return(wx.WXML{
		"return_code":    "SUCCESS",
		"return_msg":     "OK",
		"appid":          "wx2421b1c4370ec43b",
		"mch_id":         "10000100",
		"device_info":    "1000",
		"nonce_str":      "TN55wO9Pba5yENl8",
		"sign":           "07EACC03ED8DD7F1BAB6BBE1853EF998",
		"result_code":    "SUCCESS",
		"openid":         "oUpF8uN95-Ptaags6E_roPHg7AG0",
		"is_subscribe":   "Y",
		"trade_type":     "APP",
		"bank_type":      "CCB_DEBIT",
		"total_fee":      "1",
		"fee_type":       "CNY",
		"transaction_id": "1008450740201411110005820873",
		"out_trade_no":   "1415757673",
		"attach":         "订单额外描述",
		"time_end":       "20141111170043",
		"trade_state":    "SUCCESS",
	}, nil)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "ec2316275641faa3aacf3cc599e8730f"
	}
	mch.client = client
	mch.tlsClient = client

	r, err := mch.Do(context.TODO(), QueryOrderByTransactionID("1008450740201411110005820873"))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
		"return_code":    "SUCCESS",
		"return_msg":     "OK",
		"appid":          "wx2421b1c4370ec43b",
		"mch_id":         "10000100",
		"device_info":    "1000",
		"nonce_str":      "TN55wO9Pba5yENl8",
		"sign":           "07EACC03ED8DD7F1BAB6BBE1853EF998",
		"result_code":    "SUCCESS",
		"openid":         "oUpF8uN95-Ptaags6E_roPHg7AG0",
		"is_subscribe":   "Y",
		"trade_type":     "APP",
		"bank_type":      "CCB_DEBIT",
		"total_fee":      "1",
		"fee_type":       "CNY",
		"transaction_id": "1008450740201411110005820873",
		"out_trade_no":   "1415757673",
		"attach":         "订单额外描述",
		"time_end":       "20141111170043",
		"trade_state":    "SUCCESS",
	}, r)
}

func TestQueryOrderByOutTradeNO(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().PostXML(gomock.AssignableToTypeOf(context.TODO()), "https://api.mch.weixin.qq.com/pay/orderquery", wx.WXML{
		"appid":        "wx2421b1c4370ec43b",
		"mch_id":       "10000100",
		"out_trade_no": "1415757673",
		"nonce_str":    "ec2316275641faa3aacf3cc599e8730f",
		"sign_type":    "MD5",
		"sign":         "5F222EA3F23200DD4E86C4C42E96698D",
	}).Return(wx.WXML{
		"return_code":    "SUCCESS",
		"return_msg":     "OK",
		"appid":          "wx2421b1c4370ec43b",
		"mch_id":         "10000100",
		"device_info":    "1000",
		"nonce_str":      "TN55wO9Pba5yENl8",
		"sign":           "07EACC03ED8DD7F1BAB6BBE1853EF998",
		"result_code":    "SUCCESS",
		"openid":         "oUpF8uN95-Ptaags6E_roPHg7AG0",
		"is_subscribe":   "Y",
		"trade_type":     "APP",
		"bank_type":      "CCB_DEBIT",
		"total_fee":      "1",
		"fee_type":       "CNY",
		"transaction_id": "1008450740201411110005820873",
		"out_trade_no":   "1415757673",
		"attach":         "订单额外描述",
		"time_end":       "20141111170043",
		"trade_state":    "SUCCESS",
	}, nil)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "ec2316275641faa3aacf3cc599e8730f"
	}
	mch.client = client
	mch.tlsClient = client

	r, err := mch.Do(context.TODO(), QueryOrderByOutTradeNO("1415757673"))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
		"return_code":    "SUCCESS",
		"return_msg":     "OK",
		"appid":          "wx2421b1c4370ec43b",
		"mch_id":         "10000100",
		"device_info":    "1000",
		"nonce_str":      "TN55wO9Pba5yENl8",
		"sign":           "07EACC03ED8DD7F1BAB6BBE1853EF998",
		"result_code":    "SUCCESS",
		"openid":         "oUpF8uN95-Ptaags6E_roPHg7AG0",
		"is_subscribe":   "Y",
		"trade_type":     "APP",
		"bank_type":      "CCB_DEBIT",
		"total_fee":      "1",
		"fee_type":       "CNY",
		"transaction_id": "1008450740201411110005820873",
		"out_trade_no":   "1415757673",
		"attach":         "订单额外描述",
		"time_end":       "20141111170043",
		"trade_state":    "SUCCESS",
	}, r)
}

func TestCloseOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().PostXML(gomock.AssignableToTypeOf(context.TODO()), "https://api.mch.weixin.qq.com/pay/closeorder", wx.WXML{
		"appid":        "wx2421b1c4370ec43b",
		"mch_id":       "10000100",
		"out_trade_no": "1415983244",
		"nonce_str":    "4ca93f17ddf3443ceabf72f26d64fe0e",
		"sign_type":    "MD5",
		"sign":         "72D4DE9625257C606558F1027331C516",
	}).Return(wx.WXML{
		"return_code": "SUCCESS",
		"return_msg":  "OK",
		"appid":       "wx2421b1c4370ec43b",
		"mch_id":      "10000100",
		"nonce_str":   "BFK89FC6rxKCOjLX",
		"sign":        "808C1D11E84411F8DF1DF1ADC960B491",
		"result_code": "SUCCESS",
		"result_msg":  "OK",
	}, nil)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "4ca93f17ddf3443ceabf72f26d64fe0e"
	}
	mch.client = client
	mch.tlsClient = client

	r, err := mch.Do(context.TODO(), CloseOrder("1415983244"))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
		"return_code": "SUCCESS",
		"return_msg":  "OK",
		"appid":       "wx2421b1c4370ec43b",
		"mch_id":      "10000100",
		"nonce_str":   "BFK89FC6rxKCOjLX",
		"sign":        "808C1D11E84411F8DF1DF1ADC960B491",
		"result_code": "SUCCESS",
		"result_msg":  "OK",
	}, r)
}
