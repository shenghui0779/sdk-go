package mch

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/helpers"
	"github.com/stretchr/testify/assert"
)

func TestRefundByTransactionID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := helpers.NewMockHTTPClient(ctrl)

	client.EXPECT().PostXML(gomock.AssignableToTypeOf(context.TODO()), "https://api.mch.weixin.qq.com/secapi/pay/refund", helpers.WXML{
		"appid":          "wx2421b1c4370ec43b",
		"mch_id":         "10000100",
		"out_refund_no":  "1415701182",
		"total_fee":      "1",
		"refund_fee":     "1",
		"transaction_id": "4008450740201411110005820873",
		"nonce_str":      "6cefdb308e1e2e8aabd48cf79e546a02",
		"sign_type":      "MD5",
		"sign":           "29261AD6EC439F4286BF2F959EBC699D",
	}).Return(helpers.WXML{
		"return_code":    "SUCCESS",
		"return_msg":     "OK",
		"appid":          "wx2421b1c4370ec43b",
		"mch_id":         "10000100",
		"nonce_str":      "NfsMFbUFpdbEhPXP",
		"sign":           "DF0FE19C59F29CA163DDEC52CD1346A9",
		"result_code":    "SUCCESS",
		"transaction_id": "4008450740201411110005820873",
		"out_trade_no":   "1415757673",
		"out_refund_no":  "1415701182",
		"refund_id":      "2008450740201411110000174436",
		"refund_fee":     "1",
	}, nil)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "6cefdb308e1e2e8aabd48cf79e546a02"
	}

	mch.tlsClient = client

	r, err := mch.Do(context.TODO(), RefundByTransactionID("4008450740201411110005820873", &RefundData{
		OutRefundNO: "1415701182",
		TotalFee:    1,
		RefundFee:   1,
	}))

	assert.Nil(t, err)
	assert.Equal(t, helpers.WXML{
		"return_code":    "SUCCESS",
		"return_msg":     "OK",
		"appid":          "wx2421b1c4370ec43b",
		"mch_id":         "10000100",
		"nonce_str":      "NfsMFbUFpdbEhPXP",
		"sign":           "DF0FE19C59F29CA163DDEC52CD1346A9",
		"result_code":    "SUCCESS",
		"transaction_id": "4008450740201411110005820873",
		"out_trade_no":   "1415757673",
		"out_refund_no":  "1415701182",
		"refund_id":      "2008450740201411110000174436",
		"refund_fee":     "1",
	}, r)
}

func TestRefundByOutTradeNO(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := helpers.NewMockHTTPClient(ctrl)

	client.EXPECT().PostXML(gomock.AssignableToTypeOf(context.TODO()), "https://api.mch.weixin.qq.com/secapi/pay/refund", helpers.WXML{
		"appid":         "wx2421b1c4370ec43b",
		"mch_id":        "10000100",
		"out_refund_no": "1415701182",
		"total_fee":     "1",
		"refund_fee":    "1",
		"out_trade_no":  "1415757673",
		"nonce_str":     "6cefdb308e1e2e8aabd48cf79e546a02",
		"sign_type":     "MD5",
		"sign":          "D5E6945E988003E6462ACFF8D7B2DA75",
	}).Return(helpers.WXML{
		"return_code":    "SUCCESS",
		"return_msg":     "OK",
		"appid":          "wx2421b1c4370ec43b",
		"mch_id":         "10000100",
		"nonce_str":      "NfsMFbUFpdbEhPXP",
		"sign":           "DF0FE19C59F29CA163DDEC52CD1346A9",
		"result_code":    "SUCCESS",
		"transaction_id": "4008450740201411110005820873",
		"out_trade_no":   "1415757673",
		"out_refund_no":  "1415701182",
		"refund_id":      "2008450740201411110000174436",
		"refund_fee":     "1",
	}, nil)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "6cefdb308e1e2e8aabd48cf79e546a02"
	}

	mch.tlsClient = client

	r, err := mch.Do(context.TODO(), RefundByOutTradeNO("1415757673", &RefundData{
		OutRefundNO: "1415701182",
		TotalFee:    1,
		RefundFee:   1,
	}))

	assert.Nil(t, err)
	assert.Equal(t, helpers.WXML{
		"return_code":    "SUCCESS",
		"return_msg":     "OK",
		"appid":          "wx2421b1c4370ec43b",
		"mch_id":         "10000100",
		"nonce_str":      "NfsMFbUFpdbEhPXP",
		"sign":           "DF0FE19C59F29CA163DDEC52CD1346A9",
		"result_code":    "SUCCESS",
		"transaction_id": "4008450740201411110005820873",
		"out_trade_no":   "1415757673",
		"out_refund_no":  "1415701182",
		"refund_id":      "2008450740201411110000174436",
		"refund_fee":     "1",
	}, r)
}

func TestQueryRefundByRefundID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := helpers.NewMockHTTPClient(ctrl)

	client.EXPECT().PostXML(gomock.AssignableToTypeOf(context.TODO()), "https://api.mch.weixin.qq.com/pay/refundquery", helpers.WXML{
		"appid":     "wx2421b1c4370ec43b",
		"mch_id":    "10000100",
		"refund_id": "2008450740201411110000174436",
		"nonce_str": "0b9f35f484df17a732e537c37708d1d0",
		"sign_type": "MD5",
		"sign":      "8086A266B3C667377A3AE64E3F547B91",
	}).Return(helpers.WXML{
		"return_code":     "SUCCESS",
		"return_msg":      "OK",
		"appid":           "wx2421b1c4370ec43b",
		"mch_id":          "10000100",
		"nonce_str":       "TeqClE3i0mvn3DrK",
		"sign":            "68D267B5AEA32EAB799174129F6131EE",
		"result_code":     "SUCCESS",
		"out_refund_no_0": "1415701182",
		"out_trade_no":    "1415757673",
		"refund_count":    "1",
		"refund_fee_0":    "1",
		"refund_id_0":     "2008450740201411110000174436",
		"refund_status_0": "PROCESSING",
		"transaction_id":  "1008450740201411110005820873",
	}, nil)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "0b9f35f484df17a732e537c37708d1d0"
	}

	mch.client = client

	r, err := mch.Do(context.TODO(), QueryRefundByRefundID("2008450740201411110000174436"))

	assert.Nil(t, err)
	assert.Equal(t, helpers.WXML{
		"return_code":     "SUCCESS",
		"return_msg":      "OK",
		"appid":           "wx2421b1c4370ec43b",
		"mch_id":          "10000100",
		"nonce_str":       "TeqClE3i0mvn3DrK",
		"sign":            "68D267B5AEA32EAB799174129F6131EE",
		"result_code":     "SUCCESS",
		"out_refund_no_0": "1415701182",
		"out_trade_no":    "1415757673",
		"refund_count":    "1",
		"refund_fee_0":    "1",
		"refund_id_0":     "2008450740201411110000174436",
		"refund_status_0": "PROCESSING",
		"transaction_id":  "1008450740201411110005820873",
	}, r)
}

func TestQueryRefundByOutRefundNO(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := helpers.NewMockHTTPClient(ctrl)

	client.EXPECT().PostXML(gomock.AssignableToTypeOf(context.TODO()), "https://api.mch.weixin.qq.com/pay/refundquery", helpers.WXML{
		"appid":         "wx2421b1c4370ec43b",
		"mch_id":        "10000100",
		"out_refund_no": "1415701182",
		"nonce_str":     "0b9f35f484df17a732e537c37708d1d0",
		"sign_type":     "MD5",
		"sign":          "46F57A796BFF54295FB163CA68CB439D",
	}).Return(helpers.WXML{
		"return_code":     "SUCCESS",
		"return_msg":      "OK",
		"appid":           "wx2421b1c4370ec43b",
		"mch_id":          "10000100",
		"nonce_str":       "TeqClE3i0mvn3DrK",
		"sign":            "68D267B5AEA32EAB799174129F6131EE",
		"result_code":     "SUCCESS",
		"out_refund_no_0": "1415701182",
		"out_trade_no":    "1415757673",
		"refund_count":    "1",
		"refund_fee_0":    "1",
		"refund_id_0":     "2008450740201411110000174436",
		"refund_status_0": "PROCESSING",
		"transaction_id":  "1008450740201411110005820873",
	}, nil)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "0b9f35f484df17a732e537c37708d1d0"
	}

	mch.client = client

	r, err := mch.Do(context.TODO(), QueryRefundByOutRefundNO("1415701182"))

	assert.Nil(t, err)
	assert.Equal(t, helpers.WXML{
		"return_code":     "SUCCESS",
		"return_msg":      "OK",
		"appid":           "wx2421b1c4370ec43b",
		"mch_id":          "10000100",
		"nonce_str":       "TeqClE3i0mvn3DrK",
		"sign":            "68D267B5AEA32EAB799174129F6131EE",
		"result_code":     "SUCCESS",
		"out_refund_no_0": "1415701182",
		"out_trade_no":    "1415757673",
		"refund_count":    "1",
		"refund_fee_0":    "1",
		"refund_id_0":     "2008450740201411110000174436",
		"refund_status_0": "PROCESSING",
		"transaction_id":  "1008450740201411110005820873",
	}, r)
}

func TestQueryRefundByTransactionID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := helpers.NewMockHTTPClient(ctrl)

	client.EXPECT().PostXML(gomock.AssignableToTypeOf(context.TODO()), "https://api.mch.weixin.qq.com/pay/refundquery", helpers.WXML{
		"appid":          "wx2421b1c4370ec43b",
		"mch_id":         "10000100",
		"transaction_id": "1008450740201411110005820873",
		"nonce_str":      "0b9f35f484df17a732e537c37708d1d0",
		"sign_type":      "MD5",
		"sign":           "264E5038F1CB9D66132E769ABB5B745C",
	}).Return(helpers.WXML{
		"return_code":     "SUCCESS",
		"return_msg":      "OK",
		"appid":           "wx2421b1c4370ec43b",
		"mch_id":          "10000100",
		"nonce_str":       "TeqClE3i0mvn3DrK",
		"sign":            "68D267B5AEA32EAB799174129F6131EE",
		"result_code":     "SUCCESS",
		"out_refund_no_0": "1415701182",
		"out_trade_no":    "1415757673",
		"refund_count":    "1",
		"refund_fee_0":    "1",
		"refund_id_0":     "2008450740201411110000174436",
		"refund_status_0": "PROCESSING",
		"transaction_id":  "1008450740201411110005820873",
	}, nil)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "0b9f35f484df17a732e537c37708d1d0"
	}

	mch.client = client

	r, err := mch.Do(context.TODO(), QueryRefundByTransactionID("1008450740201411110005820873"))

	assert.Nil(t, err)
	assert.Equal(t, helpers.WXML{
		"return_code":     "SUCCESS",
		"return_msg":      "OK",
		"appid":           "wx2421b1c4370ec43b",
		"mch_id":          "10000100",
		"nonce_str":       "TeqClE3i0mvn3DrK",
		"sign":            "68D267B5AEA32EAB799174129F6131EE",
		"result_code":     "SUCCESS",
		"out_refund_no_0": "1415701182",
		"out_trade_no":    "1415757673",
		"refund_count":    "1",
		"refund_fee_0":    "1",
		"refund_id_0":     "2008450740201411110000174436",
		"refund_status_0": "PROCESSING",
		"transaction_id":  "1008450740201411110005820873",
	}, r)
}

func TestQueryRefundByOutTradeNO(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := helpers.NewMockHTTPClient(ctrl)

	client.EXPECT().PostXML(gomock.AssignableToTypeOf(context.TODO()), "https://api.mch.weixin.qq.com/pay/refundquery", helpers.WXML{
		"appid":        "wx2421b1c4370ec43b",
		"mch_id":       "10000100",
		"out_trade_no": "1415757673",
		"nonce_str":    "0b9f35f484df17a732e537c37708d1d0",
		"sign_type":    "MD5",
		"sign":         "5F14ED52C2F179580A1DED73268A1009",
	}).Return(helpers.WXML{
		"return_code":     "SUCCESS",
		"return_msg":      "OK",
		"appid":           "wx2421b1c4370ec43b",
		"mch_id":          "10000100",
		"nonce_str":       "TeqClE3i0mvn3DrK",
		"sign":            "68D267B5AEA32EAB799174129F6131EE",
		"result_code":     "SUCCESS",
		"out_refund_no_0": "1415701182",
		"out_trade_no":    "1415757673",
		"refund_count":    "1",
		"refund_fee_0":    "1",
		"refund_id_0":     "2008450740201411110000174436",
		"refund_status_0": "PROCESSING",
		"transaction_id":  "1008450740201411110005820873",
	}, nil)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "0b9f35f484df17a732e537c37708d1d0"
	}

	mch.client = client

	r, err := mch.Do(context.TODO(), QueryRefundByOutTradeNO("1415757673"))

	assert.Nil(t, err)
	assert.Equal(t, helpers.WXML{
		"return_code":     "SUCCESS",
		"return_msg":      "OK",
		"appid":           "wx2421b1c4370ec43b",
		"mch_id":          "10000100",
		"nonce_str":       "TeqClE3i0mvn3DrK",
		"sign":            "68D267B5AEA32EAB799174129F6131EE",
		"result_code":     "SUCCESS",
		"out_refund_no_0": "1415701182",
		"out_trade_no":    "1415757673",
		"refund_count":    "1",
		"refund_fee_0":    "1",
		"refund_id_0":     "2008450740201411110000174436",
		"refund_status_0": "PROCESSING",
		"transaction_id":  "1008450740201411110005820873",
	}, r)
}
