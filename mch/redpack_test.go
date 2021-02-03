package mch

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/wx"
	"github.com/stretchr/testify/assert"
)

func TestSendNormalRedpack(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().PostXML(gomock.AssignableToTypeOf(context.TODO()), "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendredpack", wx.WXML{
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
		"sign_type":    "MD5",
		"sign":         "C9BB9D2CBE57D6E3A28BD220AFA2248D",
	}).Return([]byte(`<xml>
	<return_code>SUCCESS</return_code>
	<result_code>SUCCESS</result_code>
	<wxappid>wx2421b1c4370ec43b</wxappid>
	<mch_id>10000100</mch_id>
	<mch_billno>0010010404201411170000046545</mch_billno>
	<re_openid>onqOjjmM1tad-3ROpncN-yUfa6uI</re_openid>
	<total_amount>1</total_amount>
</xml>`), nil)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "50780e0cca98c8c8e814883e5caa672e"
	}

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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().PostXML(gomock.AssignableToTypeOf(context.TODO()), "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendgroupredpack", wx.WXML{
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
		"sign_type":    "MD5",
		"sign":         "07A8148D88B056AE56BFBFCC8CBC0401",
	}).Return([]byte(`<xml>
	<return_code>SUCCESS</return_code>
	<result_code>SUCCESS</result_code>
	<wxappid>wx2421b1c4370ec43b</wxappid>
	<mch_id>10000100</mch_id>
	<mch_billno>0010010404201411170000046545</mch_billno>
	<re_openid>onqOjjmM1tad-3ROpncN-yUfa6uI</re_openid>
	<total_amount>1</total_amount>
</xml>`), nil)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "50780e0cca98c8c8e814883e5caa672e"
	}

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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().PostXML(gomock.AssignableToTypeOf(context.TODO()), "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendminiprogramhb", wx.WXML{
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
		"sign_type":    "MD5",
		"sign":         "68D051FA341FE68A671439BF28980CE6",
	}).Return([]byte(`<xml>
	<return_code>SUCCESS</return_code>
	<result_code>SUCCESS</result_code>
	<wxappid>wx2421b1c4370ec43b</wxappid>
	<mch_id>10000100</mch_id>
	<mch_billno>2334580734271081478888000026</mch_billno>
	<re_openid>oeDV3t7xy1IkfYFzgOsCZvdRzx3U</re_openid>
	<total_amount>10</total_amount>
	<send_listid>1000041701201609263000000204000</send_listid>
	<package>sendid=242e8abd163d300019b2cae74ba8e8c06e3f0e51ab84d16b3c80decd22a5b672&ver=8&sign=4110d649a5aef52dd6b95654ddf91ca7d5411ac159ace4e1a766b7d3967a1c3dfe1d256811445a4abda2d9cfa4a9b377a829258bd00d90313c6c346f2349fe5d&mchid=11475856&appid=wxd27ebc41b85ce36d</package>
</xml>`), nil)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "50780e0cca98c8c8e814883e5caa672e"
	}

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

func TestQueryRedpackByBillNO(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().PostXML(gomock.AssignableToTypeOf(context.TODO()), "https://api.mch.weixin.qq.com/mmpaymkttransfers/gethbinfo", wx.WXML{
		"appid":      "wx2421b1c4370ec43b",
		"mch_id":     "10000100",
		"mch_billno": "9010080799701411170000046603",
		"bill_type":  "MCHT",
		"nonce_str":  "50780e0cca98c8c8e814883e5caa672e",
		"sign_type":  "MD5",
		"sign":       "231F70D63D64EB36C1BE83E7E598B280",
	}).Return([]byte(`<xml>
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
</xml>`), nil)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "50780e0cca98c8c8e814883e5caa672e"
	}

	mch.tlsClient = client

	r, err := mch.Do(context.TODO(), QueryRedpackByBillNO("9010080799701411170000046603"))

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
