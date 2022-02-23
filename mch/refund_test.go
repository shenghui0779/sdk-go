package mch

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/mock"
	"github.com/shenghui0779/gochat/wx"
)

func TestRefundByTransactionID(t *testing.T) {
	body, err := wx.FormatMap2XML(wx.WXML{
		"appid":          "wx2421b1c4370ec43b",
		"mch_id":         "10000100",
		"out_refund_no":  "1415701182",
		"total_fee":      "1",
		"refund_fee":     "1",
		"transaction_id": "4008450740201411110005820873",
		"nonce_str":      "6cefdb308e1e2e8aabd48cf79e546a02",
		"sign_type":      "MD5",
		"sign":           "29261AD6EC439F4286BF2F959EBC699D",
	})

	assert.Nil(t, err)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`<xml>
	<return_code>SUCCESS</return_code>
	<return_msg>OK</return_msg>
	<appid>wx2421b1c4370ec43b</appid>
	<mch_id>10000100</mch_id>
	<nonce_str>NfsMFbUFpdbEhPXP</nonce_str>
	<sign>DF0FE19C59F29CA163DDEC52CD1346A9</sign>
	<result_code>SUCCESS</result_code>
	<transaction_id>4008450740201411110005820873</transaction_id>
	<out_trade_no>1415757673</out_trade_no>
	<out_refund_no>1415701182</out_refund_no>
	<refund_id>2008450740201411110000174436</refund_id>
	<refund_fee>1</refund_fee>
</xml>`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.mch.weixin.qq.com/secapi/pay/refund", body).Return(resp, nil)

	mch := New("10000100", "192006250b4c09247ec02edce69f6a2d", p12cert)

	mch.nonce = func() string {
		return "6cefdb308e1e2e8aabd48cf79e546a02"
	}

	mch.SetTLSClient(wx.WithHTTPClient(client))

	r, err := mch.Do(context.TODO(), RefundByTransactionID("wx2421b1c4370ec43b", "4008450740201411110005820873", &ParamsRefund{
		OutRefundNO: "1415701182",
		TotalFee:    1,
		RefundFee:   1,
	}))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
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
	body, err := wx.FormatMap2XML(wx.WXML{
		"appid":         "wx2421b1c4370ec43b",
		"mch_id":        "10000100",
		"out_refund_no": "1415701182",
		"total_fee":     "1",
		"refund_fee":    "1",
		"out_trade_no":  "1415757673",
		"nonce_str":     "6cefdb308e1e2e8aabd48cf79e546a02",
		"sign_type":     "MD5",
		"sign":          "D5E6945E988003E6462ACFF8D7B2DA75",
	})

	assert.Nil(t, err)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`<xml>
	<return_code>SUCCESS</return_code>
	<return_msg>OK</return_msg>
	<appid>wx2421b1c4370ec43b</appid>
	<mch_id>10000100</mch_id>
	<nonce_str>NfsMFbUFpdbEhPXP</nonce_str>
	<sign>DF0FE19C59F29CA163DDEC52CD1346A9</sign>
	<result_code>SUCCESS</result_code>
	<transaction_id>4008450740201411110005820873</transaction_id>
	<out_trade_no>1415757673</out_trade_no>
	<out_refund_no>1415701182</out_refund_no>
	<refund_id>2008450740201411110000174436</refund_id>
	<refund_fee>1</refund_fee>
</xml>`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.mch.weixin.qq.com/secapi/pay/refund", body).Return(resp, nil)

	mch := New("10000100", "192006250b4c09247ec02edce69f6a2d", p12cert)

	mch.nonce = func() string {
		return "6cefdb308e1e2e8aabd48cf79e546a02"
	}

	mch.SetTLSClient(wx.WithHTTPClient(client))

	r, err := mch.Do(context.TODO(), RefundByOutTradeNO("wx2421b1c4370ec43b", "1415757673", &ParamsRefund{
		OutRefundNO: "1415701182",
		TotalFee:    1,
		RefundFee:   1,
	}))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
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
	body, err := wx.FormatMap2XML(wx.WXML{
		"appid":     "wx2421b1c4370ec43b",
		"mch_id":    "10000100",
		"refund_id": "2008450740201411110000174436",
		"nonce_str": "0b9f35f484df17a732e537c37708d1d0",
		"sign_type": "MD5",
		"sign":      "8086A266B3C667377A3AE64E3F547B91",
	})

	assert.Nil(t, err)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`<xml>
	<return_code>SUCCESS</return_code>
	<return_msg>OK</return_msg>
	<appid>wx2421b1c4370ec43b</appid>
	<mch_id>10000100</mch_id>
	<nonce_str>TeqClE3i0mvn3DrK</nonce_str>
	<sign>68D267B5AEA32EAB799174129F6131EE</sign>
	<result_code>SUCCESS</result_code>
	<out_refund_no_0>1415701182</out_refund_no_0>
	<out_trade_no>1415757673</out_trade_no>
	<refund_count>1</refund_count>
	<refund_fee_0>1</refund_fee_0>
	<refund_id_0>2008450740201411110000174436</refund_id_0>
	<refund_status_0>PROCESSING</refund_status_0>
	<transaction_id>1008450740201411110005820873</transaction_id>
</xml>`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.mch.weixin.qq.com/pay/refundquery", body).Return(resp, nil)

	mch := New("10000100", "192006250b4c09247ec02edce69f6a2d", p12cert)

	mch.nonce = func() string {
		return "0b9f35f484df17a732e537c37708d1d0"
	}

	mch.SetClient(wx.WithHTTPClient(client))

	r, err := mch.Do(context.TODO(), QueryRefundByRefundID("wx2421b1c4370ec43b", "2008450740201411110000174436"))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
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
	body, err := wx.FormatMap2XML(wx.WXML{
		"appid":         "wx2421b1c4370ec43b",
		"mch_id":        "10000100",
		"out_refund_no": "1415701182",
		"nonce_str":     "0b9f35f484df17a732e537c37708d1d0",
		"sign_type":     "MD5",
		"sign":          "46F57A796BFF54295FB163CA68CB439D",
	})

	assert.Nil(t, err)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`<xml>
	<return_code>SUCCESS</return_code>
	<return_msg>OK</return_msg>
	<appid>wx2421b1c4370ec43b</appid>
	<mch_id>10000100</mch_id>
	<nonce_str>TeqClE3i0mvn3DrK</nonce_str>
	<sign>68D267B5AEA32EAB799174129F6131EE</sign>
	<result_code>SUCCESS</result_code>
	<out_refund_no_0>1415701182</out_refund_no_0>
	<out_trade_no>1415757673</out_trade_no>
	<refund_count>1</refund_count>
	<refund_fee_0>1</refund_fee_0>
	<refund_id_0>2008450740201411110000174436</refund_id_0>
	<refund_status_0>PROCESSING</refund_status_0>
	<transaction_id>1008450740201411110005820873</transaction_id>
</xml>`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.mch.weixin.qq.com/pay/refundquery", body).Return(resp, nil)

	mch := New("10000100", "192006250b4c09247ec02edce69f6a2d", p12cert)

	mch.nonce = func() string {
		return "0b9f35f484df17a732e537c37708d1d0"
	}

	mch.SetClient(wx.WithHTTPClient(client))

	r, err := mch.Do(context.TODO(), QueryRefundByOutRefundNO("wx2421b1c4370ec43b", "1415701182"))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
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
	body, err := wx.FormatMap2XML(wx.WXML{
		"appid":          "wx2421b1c4370ec43b",
		"mch_id":         "10000100",
		"transaction_id": "1008450740201411110005820873",
		"nonce_str":      "0b9f35f484df17a732e537c37708d1d0",
		"sign_type":      "MD5",
		"sign":           "264E5038F1CB9D66132E769ABB5B745C",
	})

	assert.Nil(t, err)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`<xml>
	<return_code>SUCCESS</return_code>
	<return_msg>OK</return_msg>
	<appid>wx2421b1c4370ec43b</appid>
	<mch_id>10000100</mch_id>
	<nonce_str>TeqClE3i0mvn3DrK</nonce_str>
	<sign>68D267B5AEA32EAB799174129F6131EE</sign>
	<result_code>SUCCESS</result_code>
	<out_refund_no_0>1415701182</out_refund_no_0>
	<out_trade_no>1415757673</out_trade_no>
	<refund_count>1</refund_count>
	<refund_fee_0>1</refund_fee_0>
	<refund_id_0>2008450740201411110000174436</refund_id_0>
	<refund_status_0>PROCESSING</refund_status_0>
	<transaction_id>1008450740201411110005820873</transaction_id>
</xml>`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.mch.weixin.qq.com/pay/refundquery", body).Return(resp, nil)

	mch := New("10000100", "192006250b4c09247ec02edce69f6a2d", p12cert)

	mch.nonce = func() string {
		return "0b9f35f484df17a732e537c37708d1d0"
	}

	mch.SetClient(wx.WithHTTPClient(client))

	r, err := mch.Do(context.TODO(), QueryRefundByTransactionID("wx2421b1c4370ec43b", "1008450740201411110005820873"))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
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
	body, err := wx.FormatMap2XML(wx.WXML{
		"appid":        "wx2421b1c4370ec43b",
		"mch_id":       "10000100",
		"out_trade_no": "1415757673",
		"nonce_str":    "0b9f35f484df17a732e537c37708d1d0",
		"sign_type":    "MD5",
		"sign":         "5F14ED52C2F179580A1DED73268A1009",
	})

	assert.Nil(t, err)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`<xml>
	<return_code>SUCCESS</return_code>
	<return_msg>OK</return_msg>
	<appid>wx2421b1c4370ec43b</appid>
	<mch_id>10000100</mch_id>
	<nonce_str>TeqClE3i0mvn3DrK</nonce_str>
	<sign>68D267B5AEA32EAB799174129F6131EE</sign>
	<result_code>SUCCESS</result_code>
	<out_refund_no_0>1415701182</out_refund_no_0>
	<out_trade_no>1415757673</out_trade_no>
	<refund_count>1</refund_count>
	<refund_fee_0>1</refund_fee_0>
	<refund_id_0>2008450740201411110000174436</refund_id_0>
	<refund_status_0>PROCESSING</refund_status_0>
	<transaction_id>1008450740201411110005820873</transaction_id>
</xml>`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.mch.weixin.qq.com/pay/refundquery", body).Return(resp, nil)

	mch := New("10000100", "192006250b4c09247ec02edce69f6a2d", p12cert)

	mch.nonce = func() string {
		return "0b9f35f484df17a732e537c37708d1d0"
	}

	mch.SetClient(wx.WithHTTPClient(client))

	r, err := mch.Do(context.TODO(), QueryRefundByOutTradeNO("wx2421b1c4370ec43b", "1415757673"))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
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
