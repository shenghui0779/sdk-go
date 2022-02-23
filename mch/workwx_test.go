package mch

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/mock"
	"github.com/shenghui0779/gochat/wx"
	"github.com/stretchr/testify/assert"
)

func TestSendWorkWXRedpack(t *testing.T) {
	body, err := wx.FormatMap2XML(wx.WXML{
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
	})

	assert.Nil(t, err)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`<xml>
	<return_code>SUCCESS</return_code>
	<result_code>SUCCESS</result_code>
	<wxappid>wx2421b1c4370ec43b</wxappid>
	<mch_id>10000100</mch_id>
	<mch_billno>0010010404201411170000046545</mch_billno>
	<re_openid>onqOjjmM1tad-3ROpncN-yUfa6uI</re_openid>
	<total_amount>1</total_amount>
</xml>`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendredpack", body).Return(resp, nil)

	mch := New("10000100", "192006250b4c09247ec02edce69f6a2d", p12cert)

	mch.nonce = func() string {
		return "50780e0cca98c8c8e814883e5caa672e"
	}

	mch.SetTLSClient(wx.WithHTTPClient(client))

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

func TestQueryWorkWXRedpack(t *testing.T) {
	body, err := wx.FormatMap2XML(wx.WXML{
		"appid":      "wx2421b1c4370ec43b",
		"mch_id":     "10000100",
		"mch_billno": "9010080799701411170000046603",
		"bill_type":  "MCHT",
		"nonce_str":  "50780e0cca98c8c8e814883e5caa672e",
		"sign_type":  "MD5",
		"sign":       "231F70D63D64EB36C1BE83E7E598B280",
	})

	assert.Nil(t, err)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`<xml>
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
</xml>`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.mch.weixin.qq.com/mmpaymkttransfers/gethbinfo", body).Return(resp, nil)

	mch := New("10000100", "192006250b4c09247ec02edce69f6a2d", p12cert)

	mch.nonce = func() string {
		return "50780e0cca98c8c8e814883e5caa672e"
	}

	mch.SetTLSClient(wx.WithHTTPClient(client))

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

func TestTransferToPocket(t *testing.T) {
	body, err := wx.FormatMap2XML(wx.WXML{
		"mch_appid":        "wx2421b1c4370ec43b",
		"mchid":            "10000100",
		"partner_trade_no": "100000982014120919616",
		"openid":           "ohO4Gt7wVPxIT1A9GjFaMYMiZY1s",
		"check_name":       "FORCE_CHECK",
		"re_user_name":     "张三",
		"amount":           "100",
		"desc":             "节日快乐!",
		"spbill_create_ip": "10.2.3.10",
		"nonce_str":        "3PG2J4ILTKCH16CQ2502SI8ZNMTM67VS",
		"sign_type":        "MD5",
		"sign":             "97CD9C3C88B189B60C230677CE0FC3BB",
	})

	assert.Nil(t, err)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`<xml>
	<return_code>SUCCESS</return_code>
	<mch_appid>wx2421b1c4370ec43b</mch_appid>
	<mchid>10000100</mchid>
	<nonce_str>lxuDzMnRjpcXzxLx0q</nonce_str>
	<result_code>SUCCESS</result_code>
	<partner_trade_no>10013574201505191526582441</partner_trade_no>
	<payment_no>1000018301201505190181489473</payment_no>
	<payment_time>2015-05-19 15:26:59</payment_time>
</xml>`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers", body).Return(resp, nil)

	mch := New("10000100", "192006250b4c09247ec02edce69f6a2d", p12cert)

	mch.nonce = func() string {
		return "3PG2J4ILTKCH16CQ2502SI8ZNMTM67VS"
	}

	mch.SetTLSClient(wx.WithHTTPClient(client))

	r, err := mch.Do(context.TODO(), TransferToBalance("wx2421b1c4370ec43b", &ParamsTransferBalance{
		PartnerTradeNO: "100000982014120919616",
		OpenID:         "ohO4Gt7wVPxIT1A9GjFaMYMiZY1s",
		CheckName:      "FORCE_CHECK",
		Amount:         100,
		Desc:           "节日快乐!",
		ReUserName:     "张三",
		SpbillCreateIP: "10.2.3.10",
	}))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
		"return_code":      "SUCCESS",
		"mch_appid":        "wx2421b1c4370ec43b",
		"mchid":            "10000100",
		"nonce_str":        "lxuDzMnRjpcXzxLx0q",
		"result_code":      "SUCCESS",
		"partner_trade_no": "10013574201505191526582441",
		"payment_no":       "1000018301201505190181489473",
		"payment_time":     "2015-05-19 15:26:59",
	}, r)
}

func TestQueryTransferPocket(t *testing.T) {
	body, err := wx.FormatMap2XML(wx.WXML{
		"appid":            "wx2421b1c4370ec43b",
		"mch_id":           "10000100",
		"partner_trade_no": "1000005901201407261446939628",
		"nonce_str":        "50780e0cca98c8c8e814883e5caa672e",
		"sign_type":        "MD5",
		"sign":             "DF0024F9502E233115C0198912B4EB5D",
	})

	assert.Nil(t, err)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`<xml>
	<return_code>SUCCESS</return_code>
	<appid>wx2421b1c4370ec43b</appid>
	<mch_id>10000100</mch_id>
	<result_code>SUCCESS</result_code>
	<detail_id>1000000000201503283103439304</detail_id>
	<partner_trade_no>1000005901201407261446939628</partner_trade_no>
	<status>SUCCESS</status>
	<payment_amount>650</payment_amount>
	<openid>oxTWIuGaIt6gTKsQRLau2M0yL16E</openid>
	<transfer_name>测试</transfer_name>
	<transfer_time>2015-04-21 20:00:00</transfer_time>
	<desc>福利测试</desc>
</xml>`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.mch.weixin.qq.com/mmpaymkttransfers/gettransferinfo", body).Return(resp, nil)

	mch := New("10000100", "192006250b4c09247ec02edce69f6a2d", p12cert)

	mch.nonce = func() string {
		return "50780e0cca98c8c8e814883e5caa672e"
	}

	mch.SetTLSClient(wx.WithHTTPClient(client))

	r, err := mch.Do(context.TODO(), QueryTransferBalance("wx2421b1c4370ec43b", "1000005901201407261446939628"))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
		"return_code":      "SUCCESS",
		"appid":            "wx2421b1c4370ec43b",
		"mch_id":           "10000100",
		"result_code":      "SUCCESS",
		"detail_id":        "1000000000201503283103439304",
		"partner_trade_no": "1000005901201407261446939628",
		"status":           "SUCCESS",
		"payment_amount":   "650",
		"openid":           "oxTWIuGaIt6gTKsQRLau2M0yL16E",
		"transfer_name":    "测试",
		"transfer_time":    "2015-04-21 20:00:00",
		"desc":             "福利测试",
	}, r)
}
