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

func TestSendNormalRedpack(t *testing.T) {
	body, err := wx.FormatMap2XMLForTest(wx.WXML{
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
	})

	assert.Nil(t, err)

	resp := []byte(`<xml>
	<return_code>SUCCESS</return_code>
	<result_code>SUCCESS</result_code>
	<wxappid>wx2421b1c4370ec43b</wxappid>
	<mch_id>10000100</mch_id>
	<mch_billno>0010010404201411170000046545</mch_billno>
	<re_openid>onqOjjmM1tad-3ROpncN-yUfa6uI</re_openid>
	<total_amount>1</total_amount>
</xml>`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendredpack", body).Return(resp, nil)

	mch := New("10000100", "192006250b4c09247ec02edce69f6a2d", WithNonce(func() string {
		return "50780e0cca98c8c8e814883e5caa672e"
	}), WithMockClient(client))

	r, err := mch.Do(context.TODO(), SendNormalRedpack("wx2421b1c4370ec43b", &ParamsRedpack{
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
	assert.Equal(t, wx.WXML{
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
	body, err := wx.FormatMap2XMLForTest(wx.WXML{
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
	})

	assert.Nil(t, err)

	resp := []byte(`<xml>
	<return_code>SUCCESS</return_code>
	<result_code>SUCCESS</result_code>
	<wxappid>wx2421b1c4370ec43b</wxappid>
	<mch_id>10000100</mch_id>
	<mch_billno>0010010404201411170000046545</mch_billno>
	<re_openid>onqOjjmM1tad-3ROpncN-yUfa6uI</re_openid>
	<total_amount>1</total_amount>
</xml>`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendgroupredpack", body).Return(resp, nil)

	mch := New("10000100", "192006250b4c09247ec02edce69f6a2d", WithNonce(func() string {
		return "50780e0cca98c8c8e814883e5caa672e"
	}), WithMockClient(client))

	r, err := mch.Do(context.TODO(), SendGroupRedpack("wx2421b1c4370ec43b", &ParamsRedpack{
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
	assert.Equal(t, wx.WXML{
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
	body, err := wx.FormatMap2XMLForTest(wx.WXML{
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
	})

	assert.Nil(t, err)

	resp := []byte(`<xml>
	<return_code>SUCCESS</return_code>
	<result_code>SUCCESS</result_code>
	<wxappid>wx2421b1c4370ec43b</wxappid>
	<mch_id>10000100</mch_id>
	<mch_billno>2334580734271081478888000026</mch_billno>
	<re_openid>oeDV3t7xy1IkfYFzgOsCZvdRzx3U</re_openid>
	<total_amount>10</total_amount>
	<send_listid>1000041701201609263000000204000</send_listid>
	<package>sendid=242e8abd163d300019b2cae74ba8e8c06e3f0e51ab84d16b3c80decd22a5b672&ver=8&sign=4110d649a5aef52dd6b95654ddf91ca7d5411ac159ace4e1a766b7d3967a1c3dfe1d256811445a4abda2d9cfa4a9b377a829258bd00d90313c6c346f2349fe5d&mchid=11475856&appid=wxd27ebc41b85ce36d</package>
</xml>`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendminiprogramhb", body).Return(resp, nil)

	mch := New("10000100", "192006250b4c09247ec02edce69f6a2d", WithNonce(func() string {
		return "50780e0cca98c8c8e814883e5caa672e"
	}), WithMockClient(client))

	r, err := mch.Do(context.TODO(), SendMinipRedpack("wx2421b1c4370ec43b", &ParamsRedpack{
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
	assert.Equal(t, wx.WXML{
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

func TestQueryRedpack(t *testing.T) {
	body, err := wx.FormatMap2XMLForTest(wx.WXML{
		"appid":      "wx2421b1c4370ec43b",
		"mch_id":     "10000100",
		"mch_billno": "9010080799701411170000046603",
		"bill_type":  "MCHT",
		"nonce_str":  "50780e0cca98c8c8e814883e5caa672e",
		"sign":       "B52930D6136EA0B5A40F5692EA47DE08",
	})

	assert.Nil(t, err)

	resp := []byte(`<xml>
	<return_code>SUCCESS</return_code>
	<result_code>SUCCESS</result_code>
	<mch_id>10000100</mch_id>
	<mch_billno>9010080799701411170000046603</mch_billno>
	<detail_id>10000417012016080830956240040</detail_id>
	<status>RECEIVED</status>
	<send_type>ACTIVITY</send_type>
	<hb_type>NORMAL</hb_type>
	<total_amount>100</total_amount>
	<total_num>1</total_num>
	<send_time>2016-08-08 21:49:22</send_time>
</xml>`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.mch.weixin.qq.com/mmpaymkttransfers/gethbinfo", body).Return(resp, nil)

	mch := New("10000100", "192006250b4c09247ec02edce69f6a2d", WithNonce(func() string {
		return "50780e0cca98c8c8e814883e5caa672e"
	}), WithMockClient(client))

	r, err := mch.Do(context.TODO(), QueryRedpack("wx2421b1c4370ec43b", "9010080799701411170000046603"))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
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
