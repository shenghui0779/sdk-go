package mch

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/mock"
	"github.com/shenghui0779/gochat/wx"
	"github.com/stretchr/testify/assert"
)

func TestMicroPay(t *testing.T) {
	body, err := wx.FormatMap2XMLForTest(wx.WXML{
		"appid":            "wx2421b1c4370ec43b",
		"mch_id":           "10000100",
		"nonce_str":        "8aaee146b1dee7cec9100add9b96cbe2",
		"auth_code":        "120269300684844649",
		"device_info":      "1000",
		"body":             "付款码支付测试",
		"out_trade_no":     "1415757673",
		"total_fee":        "1",
		"fee_type":         "CNY",
		"spbill_create_ip": "14.17.22.52",
		"attach":           "订单额外描述",
		"sign":             "660FEFCB2805EFE3639A2C98A571214C",
	})

	assert.Nil(t, err)

	resp := []byte(`<xml>
	<return_code>SUCCESS</return_code>
	<return_msg>OK</return_msg>
	<appid>wx2421b1c4370ec43b</appid>
	<mch_id>10000100</mch_id>
	<device_info>1000</device_info>
	<nonce_str>GOp3TRyMXzbMlkun</nonce_str>
	<sign>1C3E24AACED0148F116C1122EDF2A26F</sign>
	<result_code>SUCCESS</result_code>
	<openid>oUpF8uN95-Ptaags6E_roPHg7AG0</openid>
	<is_subscribe>Y</is_subscribe>
	<trade_type>MICROPAY</trade_type>
	<bank_type>CCB_DEBIT</bank_type>
	<total_fee>1</total_fee>
	<coupon_fee>0</coupon_fee>
	<fee_type>CNY</fee_type>
	<transaction_id>1008450740201411110005820873</transaction_id>
	<out_trade_no>1415757673</out_trade_no>
	<attach>订单额外描述</attach>
	<time_end>20141111170043</time_end>
</xml>`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.mch.weixin.qq.com/pay/micropay", body).Return(resp, nil)

	mch := New("10000100", "192006250b4c09247ec02edce69f6a2d", WithNonce(func() string {
		return "8aaee146b1dee7cec9100add9b96cbe2"
	}), WithMockClient(client))

	r, err := mch.Do(context.TODO(), MicroPay("wx2421b1c4370ec43b", &ParamsMicroPay{
		OutTradeNO:     "1415757673",
		TotalFee:       1,
		SpbillCreateIP: "14.17.22.52",
		AuthCode:       "120269300684844649",
		Body:           "付款码支付测试",
		DeviceInfo:     "1000",
		Attach:         "订单额外描述",
	}))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
		"return_code":    "SUCCESS",
		"return_msg":     "OK",
		"appid":          "wx2421b1c4370ec43b",
		"mch_id":         "10000100",
		"device_info":    "1000",
		"nonce_str":      "GOp3TRyMXzbMlkun",
		"sign":           "1C3E24AACED0148F116C1122EDF2A26F",
		"result_code":    "SUCCESS",
		"openid":         "oUpF8uN95-Ptaags6E_roPHg7AG0",
		"is_subscribe":   "Y",
		"trade_type":     "MICROPAY",
		"bank_type":      "CCB_DEBIT",
		"total_fee":      "1",
		"coupon_fee":     "0",
		"fee_type":       "CNY",
		"transaction_id": "1008450740201411110005820873",
		"out_trade_no":   "1415757673",
		"attach":         "订单额外描述",
		"time_end":       "20141111170043",
	}, r)
}

func TestReverseByTransactionID(t *testing.T) {
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
	<nonce_str>o5bAKF3o2ypC8hwa</nonce_str>
	<sign>9403693F38A33633FEE2CCF726E016E8</sign>
	<result_code>SUCCESS</result_code>
	<recall>N</recall>
</xml>`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.mch.weixin.qq.com/secapi/pay/reverse", body).Return(resp, nil)

	mch := New("10000100", "192006250b4c09247ec02edce69f6a2d", WithNonce(func() string {
		return "ec2316275641faa3aacf3cc599e8730f"
	}), WithMockClient(client))

	r, err := mch.Do(context.TODO(), ReverseByTransactionID("wx2421b1c4370ec43b", "1008450740201411110005820873"))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
		"return_code": "SUCCESS",
		"return_msg":  "OK",
		"appid":       "wx2421b1c4370ec43b",
		"mch_id":      "10000100",
		"nonce_str":   "o5bAKF3o2ypC8hwa",
		"sign":        "9403693F38A33633FEE2CCF726E016E8",
		"result_code": "SUCCESS",
		"recall":      "N",
	}, r)
}

func TestReverseByOutTradeNO(t *testing.T) {
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
	<nonce_str>o5bAKF3o2ypC8hwa</nonce_str>
	<sign>9403693F38A33633FEE2CCF726E016E8</sign>
	<result_code>SUCCESS</result_code>
	<recall>N</recall>
</xml>`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.mch.weixin.qq.com/secapi/pay/reverse", body).Return(resp, nil)

	mch := New("10000100", "192006250b4c09247ec02edce69f6a2d", WithNonce(func() string {
		return "ec2316275641faa3aacf3cc599e8730f"
	}), WithMockClient(client))

	r, err := mch.Do(context.TODO(), ReverseByOutTradeNO("wx2421b1c4370ec43b", "1415757673"))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
		"return_code": "SUCCESS",
		"return_msg":  "OK",
		"appid":       "wx2421b1c4370ec43b",
		"mch_id":      "10000100",
		"nonce_str":   "o5bAKF3o2ypC8hwa",
		"sign":        "9403693F38A33633FEE2CCF726E016E8",
		"result_code": "SUCCESS",
		"recall":      "N",
	}, r)
}
