package mch

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/helpers"
	"github.com/stretchr/testify/assert"
)

func TestSendNormalRedpack(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := helpers.NewMockHTTPClient(ctrl)

	client.EXPECT().PostXML(gomock.AssignableToTypeOf(context.TODO()), RedpackNormalURL, helpers.WXML{
		"wxappid":      "wx2421b1c4370ec43b",
		"mch_id":       "10000100",
		"mch_billno":   "0010010404201411170000046545",
		"send_name":    "send_name",
		"re_openid":    "onqOjjmM1tad-3ROpncN-yUfa6uI",
		"total_amount": "200",
		"total_num":    "1",
		"wishing":      "恭喜发财",
		"client_ip":    "127.0.0.1",
		"act_name":     "新年红包",
		"remark":       "新年红包",
		"scene_id":     "PRODUCT_2",
		"risk_info":    "posttime%3d123123412%26clientversion%3d234134%26mobile%3d122344545%26deviceid%3dIOS",
		"nonce_str":    "50780e0cca98c8c8e814883e5caa672e",
		"sign":         "CAE645705D54BA78424107C6048E45B8",
	}).Return(helpers.WXML{
		"return_code":  "SUCCESS",
		"result_code":  "SUCCESS",
		"wxappid":      "wx2421b1c4370ec43b",
		"mch_id":       "10000100",
		"mch_billno":   "0010010404201411170000046545",
		"re_openid":    "onqOjjmM1tad-3ROpncN-yUfa6uI",
		"total_amount": "1",
	}, nil)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "50780e0cca98c8c8e814883e5caa672e"
	}
	mch.client = client
	mch.tlsClient = client

	r, err := mch.Do(context.TODO(), SendNormalRedpack(&RedpackData{
		MchBillNO:   "0010010404201411170000046545",
		SendName:    "send_name",
		ReOpenID:    "onqOjjmM1tad-3ROpncN-yUfa6uI",
		TotalAmount: 200,
		TotalNum:    1,
		Wishing:     "恭喜发财",
		ClientIP:    "127.0.0.1",
		ActName:     "新年红包",
		Remark:      "新年红包",
		SceneID:     "PRODUCT_2",
		RiskInfo:    "posttime%3d123123412%26clientversion%3d234134%26mobile%3d122344545%26deviceid%3dIOS",
	}))

	assert.Nil(t, err)
	assert.Equal(t, helpers.WXML{
		"return_code":  "SUCCESS",
		"result_code":  "SUCCESS",
		"wxappid":      "wx2421b1c4370ec43b",
		"mch_id":       "10000100",
		"mch_billno":   "0010010404201411170000046545",
		"re_openid":    "onqOjjmM1tad-3ROpncN-yUfa6uI",
		"total_amount": "1",
	}, r)
}

func TestSendGroupRedpack(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := helpers.NewMockHTTPClient(ctrl)

	client.EXPECT().PostXML(gomock.AssignableToTypeOf(context.TODO()), RedpackGroupURL, helpers.WXML{
		"wxappid":      "wx2421b1c4370ec43b",
		"mch_id":       "10000100",
		"mch_billno":   "0010010404201411170000046545",
		"send_name":    "send_name",
		"re_openid":    "onqOjjmM1tad-3ROpncN-yUfa6uI",
		"total_amount": "600",
		"total_num":    "3",
		"amt_type":     "ALL_RAND",
		"wishing":      "恭喜发财",
		"act_name":     "新年红包",
		"remark":       "新年红包",
		"scene_id":     "PRODUCT_2",
		"risk_info":    "posttime%3d123123412%26clientversion%3d234134%26mobile%3d122344545%26deviceid%3dIOS",
		"nonce_str":    "50780e0cca98c8c8e814883e5caa672e",
		"sign":         "A7E8609BDC147326E8EE82BD031EBA3D",
	}).Return(helpers.WXML{
		"return_code":  "SUCCESS",
		"result_code":  "SUCCESS",
		"wxappid":      "wx2421b1c4370ec43b",
		"mch_id":       "10000100",
		"mch_billno":   "0010010404201411170000046545",
		"re_openid":    "onqOjjmM1tad-3ROpncN-yUfa6uI",
		"total_amount": "1",
	}, nil)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "50780e0cca98c8c8e814883e5caa672e"
	}
	mch.client = client
	mch.tlsClient = client

	r, err := mch.Do(context.TODO(), SendGroupRedpack(&RedpackData{
		MchBillNO:   "0010010404201411170000046545",
		SendName:    "send_name",
		ReOpenID:    "onqOjjmM1tad-3ROpncN-yUfa6uI",
		TotalAmount: 600,
		TotalNum:    3,
		Wishing:     "恭喜发财",
		ActName:     "新年红包",
		Remark:      "新年红包",
		SceneID:     "PRODUCT_2",
		RiskInfo:    "posttime%3d123123412%26clientversion%3d234134%26mobile%3d122344545%26deviceid%3dIOS",
	}))

	assert.Nil(t, err)
	assert.Equal(t, helpers.WXML{
		"return_code":  "SUCCESS",
		"result_code":  "SUCCESS",
		"wxappid":      "wx2421b1c4370ec43b",
		"mch_id":       "10000100",
		"mch_billno":   "0010010404201411170000046545",
		"re_openid":    "onqOjjmM1tad-3ROpncN-yUfa6uI",
		"total_amount": "1",
	}, r)
}

func TestSendMinipRedpack(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := helpers.NewMockHTTPClient(ctrl)

	client.EXPECT().PostXML(gomock.AssignableToTypeOf(context.TODO()), RedpackMinipURL, helpers.WXML{
		"wxappid":      "wx2421b1c4370ec43b",
		"mch_id":       "10000100",
		"mch_billno":   "2334580734271081478888000026",
		"send_name":    "miniprogramtest",
		"re_openid":    "oeDV3t7xy1IkfYFzgOsCZvdRjb45",
		"total_amount": "100",
		"total_num":    "1",
		"wishing":      "wishing",
		"act_name":     "act_name",
		"remark":       "remark",
		"notify_way":   "MINI_PROGRAM_JSAPI",
		"nonce_str":    "50780e0cca98c8c8e814883e5caa672e",
		"sign":         "A3F75B94BB93591BA9065556F70855FA",
	}).Return(helpers.WXML{
		"return_code":  "SUCCESS",
		"result_code":  "SUCCESS",
		"wxappid":      "wx2421b1c4370ec43b",
		"mch_id":       "10000100",
		"mch_billno":   "2334580734271081478888000026",
		"re_openid":    "oeDV3t7xy1IkfYFzgOsCZvdRzx3U",
		"total_amount": "10",
		"send_listid":  "1000041701201609263000000204000",
		"package":      "sendid=242e8abd163d300019b2cae74ba8e8c06e3f0e51ab84d16b3c80decd22a5b672&ver=8&sign=4110d649a5aef52dd6b95654ddf91ca7d5411ac159ace4e1a766b7d3967a1c3dfe1d256811445a4abda2d9cfa4a9b377a829258bd00d90313c6c346f2349fe5d&mchid=11475856&appid=wxd27ebc41b85ce36d",
	}, nil)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "50780e0cca98c8c8e814883e5caa672e"
	}
	mch.client = client
	mch.tlsClient = client

	r, err := mch.Do(context.TODO(), SendMinipRedpack(&RedpackData{
		MchBillNO:   "2334580734271081478888000026",
		SendName:    "miniprogramtest",
		ReOpenID:    "oeDV3t7xy1IkfYFzgOsCZvdRjb45",
		TotalAmount: 100,
		TotalNum:    1,
		Wishing:     "wishing",
		ActName:     "act_name",
		Remark:      "remark",
	}))

	assert.Nil(t, err)
	assert.Equal(t, helpers.WXML{
		"return_code":  "SUCCESS",
		"result_code":  "SUCCESS",
		"wxappid":      "wx2421b1c4370ec43b",
		"mch_id":       "10000100",
		"mch_billno":   "2334580734271081478888000026",
		"re_openid":    "oeDV3t7xy1IkfYFzgOsCZvdRzx3U",
		"total_amount": "10",
		"send_listid":  "1000041701201609263000000204000",
		"package":      "sendid=242e8abd163d300019b2cae74ba8e8c06e3f0e51ab84d16b3c80decd22a5b672&ver=8&sign=4110d649a5aef52dd6b95654ddf91ca7d5411ac159ace4e1a766b7d3967a1c3dfe1d256811445a4abda2d9cfa4a9b377a829258bd00d90313c6c346f2349fe5d&mchid=11475856&appid=wxd27ebc41b85ce36d",
	}, r)
}

func TestQueryRedpackByBillNO(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := helpers.NewMockHTTPClient(ctrl)

	client.EXPECT().PostXML(gomock.AssignableToTypeOf(context.TODO()), RedpackQueryURL, helpers.WXML{
		"appid":      "wx2421b1c4370ec43b",
		"mch_id":     "10000100",
		"mch_billno": "9010080799701411170000046603",
		"bill_type":  "MCHT",
		"nonce_str":  "50780e0cca98c8c8e814883e5caa672e",
		"sign":       "B52930D6136EA0B5A40F5692EA47DE08",
	}).Return(helpers.WXML{
		"return_code":  "SUCCESS",
		"result_code":  "SUCCESS",
		"mch_id":       "10000100",
		"mch_billno":   "9010080799701411170000046603",
		"detail_id":    "10000417012016080830956240040",
		"status":       "RECEIVED",
		"send_type":    "ACTIVITY",
		"hb_type":      "NORMAL",
		"total_amount": "100",
		"total_num":    "1",
		"send_time":    "2016-08-08 21:49:22",
	}, nil)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "50780e0cca98c8c8e814883e5caa672e"
	}
	mch.client = client
	mch.tlsClient = client

	r, err := mch.Do(context.TODO(), QueryRedpackByBillNO("9010080799701411170000046603"))

	assert.Nil(t, err)
	assert.Equal(t, helpers.WXML{
		"return_code":  "SUCCESS",
		"result_code":  "SUCCESS",
		"mch_id":       "10000100",
		"mch_billno":   "9010080799701411170000046603",
		"detail_id":    "10000417012016080830956240040",
		"status":       "RECEIVED",
		"send_type":    "ACTIVITY",
		"hb_type":      "NORMAL",
		"total_amount": "100",
		"total_num":    "1",
		"send_time":    "2016-08-08 21:49:22",
	}, r)
}
