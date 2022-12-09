package mch

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/mock"
	"github.com/shenghui0779/gochat/wx"
)

func TestUnifyOrder(t *testing.T) {
	body, err := wx.FormatMap2XMLForTest(wx.WXML{
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
		"sign":             "845E712D712B283EEB45448079C12D41",
	})

	assert.Nil(t, err)

	resp := []byte(`<xml>
	<return_code>SUCCESS</return_code>
	<return_msg>OK</return_msg>
	<appid>wx2421b1c4370ec43b</appid>
	<mch_id>10000100</mch_id>
	<nonce_str>IITRi8Iabbblz1Jc</nonce_str>
	<sign>E515C9BE3D3129764915407267CA0243</sign>
	<result_code>SUCCESS</result_code>
	<prepay_id>wx201411101639507cbf6ffd8b0779950874</prepay_id>
	<trade_type>APP</trade_type>
</xml>`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.mch.weixin.qq.com/pay/unifiedorder", body).Return(resp, nil)

	mch := New("10000100", "192006250b4c09247ec02edce69f6a2d", WithNonce(func() string {
		return "1add1a30ac87aa2db72f57a2375d8fec"
	}), WithMockClient(client))

	r, err := mch.Do(context.TODO(), UnifyOrder("wx2421b1c4370ec43b", &ParamsUnifyOrder{
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
	body, err := wx.FormatMap2XMLForTest(wx.WXML{
		"appid":          "wx2421b1c4370ec43b",
		"mch_id":         "10000100",
		"transaction_id": "1008450740201411110005820873",
		"nonce_str":      "ec2316275641faa3aacf3cc599e8730f",
		"sign":           "3714E8E1F1327F1FD47627DE033745C2",
	})

	assert.Nil(t, err)

	resp := []byte(`<xml>
	<return_code>SUCCESS</return_code>
	<return_msg>OK</return_msg>
	<appid>wx2421b1c4370ec43b</appid>
	<mch_id>10000100</mch_id>
	<device_info>1000</device_info>
	<nonce_str>TN55wO9Pba5yENl8</nonce_str>
	<sign>07EACC03ED8DD7F1BAB6BBE1853EF998</sign>
	<result_code>SUCCESS</result_code>
	<openid>oUpF8uN95-Ptaags6E_roPHg7AG0</openid>
	<is_subscribe>Y</is_subscribe>
	<trade_type>APP</trade_type>
	<bank_type>CCB_DEBIT</bank_type>
	<total_fee>1</total_fee>
	<fee_type>CNY</fee_type>
	<transaction_id>1008450740201411110005820873</transaction_id>
	<out_trade_no>1415757673</out_trade_no>
	<attach>订单额外描述</attach>
	<time_end>20141111170043</time_end>
	<trade_state>SUCCESS</trade_state>
</xml>`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.mch.weixin.qq.com/pay/orderquery", body).Return(resp, nil)

	mch := New("10000100", "192006250b4c09247ec02edce69f6a2d", WithNonce(func() string {
		return "ec2316275641faa3aacf3cc599e8730f"
	}), WithMockClient(client))

	r, err := mch.Do(context.TODO(), QueryOrderByTransactionID("wx2421b1c4370ec43b", "1008450740201411110005820873"))

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
	body, err := wx.FormatMap2XMLForTest(wx.WXML{
		"appid":        "wx2421b1c4370ec43b",
		"mch_id":       "10000100",
		"out_trade_no": "1415757673",
		"nonce_str":    "ec2316275641faa3aacf3cc599e8730f",
		"sign":         "B077F50A14E8A088ACAD5A2318F7207E",
	})

	assert.Nil(t, err)

	resp := []byte(`<xml>
	<return_code>SUCCESS</return_code>
	<return_msg>OK</return_msg>
	<appid>wx2421b1c4370ec43b</appid>
	<mch_id>10000100</mch_id>
	<device_info>1000</device_info>
	<nonce_str>TN55wO9Pba5yENl8</nonce_str>
	<sign>07EACC03ED8DD7F1BAB6BBE1853EF998</sign>
	<result_code>SUCCESS</result_code>
	<openid>oUpF8uN95-Ptaags6E_roPHg7AG0</openid>
	<is_subscribe>Y</is_subscribe>
	<trade_type>APP</trade_type>
	<bank_type>CCB_DEBIT</bank_type>
	<total_fee>1</total_fee>
	<fee_type>CNY</fee_type>
	<transaction_id>1008450740201411110005820873</transaction_id>
	<out_trade_no>1415757673</out_trade_no>
	<attach>订单额外描述</attach>
	<time_end>20141111170043</time_end>
	<trade_state>SUCCESS</trade_state>
</xml>`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.mch.weixin.qq.com/pay/orderquery", body).Return(resp, nil)

	mch := New("10000100", "192006250b4c09247ec02edce69f6a2d", WithNonce(func() string {
		return "ec2316275641faa3aacf3cc599e8730f"
	}), WithMockClient(client))

	r, err := mch.Do(context.TODO(), QueryOrderByOutTradeNO("wx2421b1c4370ec43b", "1415757673"))

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
	body, err := wx.FormatMap2XMLForTest(wx.WXML{
		"appid":        "wx2421b1c4370ec43b",
		"mch_id":       "10000100",
		"out_trade_no": "1415983244",
		"nonce_str":    "4ca93f17ddf3443ceabf72f26d64fe0e",
		"sign":         "8D2363E7A0948301A065F369DEB9011B",
	})

	assert.Nil(t, err)

	resp := []byte(`<xml>
	<return_code>SUCCESS</return_code>
	<return_msg>OK</return_msg>
	<appid>wx2421b1c4370ec43b</appid>
	<mch_id>10000100</mch_id>
	<nonce_str>BFK89FC6rxKCOjLX</nonce_str>
	<sign>808C1D11E84411F8DF1DF1ADC960B491</sign>
	<result_code>SUCCESS</result_code>
	<result_msg>OK</result_msg>
</xml>`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.mch.weixin.qq.com/pay/closeorder", body).Return(resp, nil)

	mch := New("10000100", "192006250b4c09247ec02edce69f6a2d", WithNonce(func() string {
		return "4ca93f17ddf3443ceabf72f26d64fe0e"
	}), WithMockClient(client))

	r, err := mch.Do(context.TODO(), CloseOrder("wx2421b1c4370ec43b", "1415983244"))

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
