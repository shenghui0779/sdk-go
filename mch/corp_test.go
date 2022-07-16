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

func TestSendCorpRedpack(t *testing.T) {
	body, err := wx.FormatMap2XMLForTest(wx.WXML{
		"wxappid":                "wx8888888888888888",
		"mch_id":                 "10000098",
		"mch_billno":             "123456",
		"re_openid":              "oxTWIuGaIt6gTKsQRLau2M0yL16E",
		"total_amount":           "1000",
		"wishing":                "感谢您参加猜灯谜活动，祝您元宵节快乐！",
		"act_name":               "猜灯谜抢红包活动",
		"remark":                 "猜越多得越多，快来抢！",
		"sender_name":            "XX活动",
		"sender_header_media_id": "1G6nrLmr5EC3MMb_-zK1dDdzmd0p7cNliYu9V5w7o8K0",
		"nonce_str":              "5K8264ILTKCH16CQ2502SI8ZNMTM67VS",
		"workwx_sign":            "FCB8624E2CED048CA15E2FF2B87C91A6",
		"sign":                   "F49B5E6435F626E21918C3CBE0EFC416",
	})

	assert.Nil(t, err)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`<xml>
	<return_code><![CDATA[SUCCESS]]></return_code>
	<return_msg><![CDATA[ok]]></return_msg>
	<sign><![CDATA[3894524D14BA9FE46E7E58DF3514EF12]]></sign>
	<result_code><![CDATA[SUCCESS]]></result_code>
	<mch_billno><![CDATA[123456]]></mch_billno>
	<mch_id><![CDATA[10000098]]></mch_id>
	<wxappid><![CDATA[wx8888888888888888]]></wxappid>
	<re_openid><![CDATA[oxTWIuGaIt6gTKsQRLau2M0yL16E]]></re_openid>
	<total_amount><![CDATA[1000]]></total_amount>
	<send_listid><![CDATA[235785324578098]]></send_listid>
	<sender_name><![CDATA[XX活动]]></sender_name>
	<sender_header_media_id><![CDATA[1G6nrLmr5EC3MMb_-zK1dDdzmd0p7cNliYu9V5w7o8K0]]></sender_header_media_id>
</xml>`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendworkwxredpack", body).Return(resp, nil)

	mch := New("10000098", "192006250b4c09247ec02edce69f6a2d", p12cert)

	mch.nonce = func() string {
		return "5K8264ILTKCH16CQ2502SI8ZNMTM67VS"
	}

	mch.SetTLSClient(wx.WithHTTPClient(client))

	r, err := mch.Do(context.TODO(), SendCorpRedpack("wx8888888888888888", "192006250b4c09247ec02edce69f6a2d", &ParamsCorpRedpack{
		MchBillNO:           "123456",
		ReOpenID:            "oxTWIuGaIt6gTKsQRLau2M0yL16E",
		TotalAmount:         1000,
		Wishing:             "感谢您参加猜灯谜活动，祝您元宵节快乐！",
		ActName:             "猜灯谜抢红包活动",
		Remark:              "猜越多得越多，快来抢！",
		SenderName:          "XX活动",
		SenderHeaderMediaID: "1G6nrLmr5EC3MMb_-zK1dDdzmd0p7cNliYu9V5w7o8K0",
	}))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
		"return_code":            "SUCCESS",
		"return_msg":             "ok",
		"result_code":            "SUCCESS",
		"wxappid":                "wx8888888888888888",
		"mch_id":                 "10000098",
		"sign":                   "3894524D14BA9FE46E7E58DF3514EF12",
		"mch_billno":             "123456",
		"re_openid":              "oxTWIuGaIt6gTKsQRLau2M0yL16E",
		"total_amount":           "1000",
		"send_listid":            "235785324578098",
		"sender_name":            "XX活动",
		"sender_header_media_id": "1G6nrLmr5EC3MMb_-zK1dDdzmd0p7cNliYu9V5w7o8K0",
	}, r)
}

func TestQueryCorpRedpack(t *testing.T) {
	body, err := wx.FormatMap2XMLForTest(wx.WXML{
		"appid":      "wx8888888888888888",
		"mch_id":     "10000098",
		"mch_billno": "123456",
		"nonce_str":  "5K8264ILTKCH16CQ2502SI8ZNMTM67VS",
		"sign":       "124D7B0A14B3D98F109DF3D560190EF0",
	})

	assert.Nil(t, err)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`<xml>
	<return_code><![CDATA[SUCCESS]]></return_code>
	<return_msg><![CDATA[ok]]></return_msg>
	<sign><![CDATA[645086515F23B2113649B68D00D131E4]]></sign>
	<result_code><![CDATA[SUCCESS]]></result_code>
	<mch_billno><![CDATA[123456]]></mch_billno>
	<mch_id><![CDATA[10000098]]></mch_id>
	<detail_id><![CDATA[43235678654322356]]></detail_id>
	<status><![CDATA[RECEIVED]]></status>
	<send_type><![CDATA[API]]></send_type>
	<total_amount><![CDATA[5000]]></total_amount>
	<reason><![CDATA[余额不足]]></reason>
	<send_time><![CDATA[2017-07-20 22:45:12]]></send_time>
	<wishing><![CDATA[新年快乐]]></wishing>
	<remark><![CDATA[新年红包]]></remark>
	<act_name><![CDATA[新年红包]]></act_name>
	<openid><![CDATA[ohO4GtzOAAYMp2yapORH3dQB3W18]]></openid>
	<amount><![CDATA[100]]></amount>
	<rcv_time><![CDATA[2017-07-20 22:46:59]]></rcv_time>
	<sender_name><![CDATA[XX活动]]></sender_name>
	<sender_header_media_id><![CDATA[1G6nrLmr5EC3MMb_-zK1dDdzmd0p7cNliYu9V5w7o8K0]]></sender_header_media_id>
</xml>
`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.mch.weixin.qq.com/mmpaymkttransfers/queryworkwxredpack", body).Return(resp, nil)

	mch := New("10000098", "192006250b4c09247ec02edce69f6a2d", p12cert)

	mch.nonce = func() string {
		return "5K8264ILTKCH16CQ2502SI8ZNMTM67VS"
	}

	mch.SetTLSClient(wx.WithHTTPClient(client))

	r, err := mch.Do(context.TODO(), QueryCorpRedpack("wx8888888888888888", "123456"))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
		"return_code":            "SUCCESS",
		"return_msg":             "ok",
		"result_code":            "SUCCESS",
		"mch_id":                 "10000098",
		"mch_billno":             "123456",
		"sign":                   "645086515F23B2113649B68D00D131E4",
		"detail_id":              "43235678654322356",
		"status":                 "RECEIVED",
		"send_type":              "API",
		"total_amount":           "5000",
		"reason":                 "余额不足",
		"send_time":              "2017-07-20 22:45:12",
		"wishing":                "新年快乐",
		"remark":                 "新年红包",
		"act_name":               "新年红包",
		"openid":                 "ohO4GtzOAAYMp2yapORH3dQB3W18",
		"amount":                 "100",
		"rcv_time":               "2017-07-20 22:46:59",
		"sender_name":            "XX活动",
		"sender_header_media_id": "1G6nrLmr5EC3MMb_-zK1dDdzmd0p7cNliYu9V5w7o8K0",
	}, r)
}

func TestTransferToPocket(t *testing.T) {
	body, err := wx.FormatMap2XMLForTest(wx.WXML{
		"appid":            "wxe062425f740c8888",
		"mch_id":           "1900000109",
		"device_info":      "013467007045764",
		"nonce_str":        "3PG2J4ILTKCH16CQ2502SI8ZNMTM67VS",
		"partner_trade_no": "100000982017072019616",
		"openid":           "ohO4Gt7wVPxIT1A9GjFaMYMiZY1s",
		"check_name":       "NO_CHECK",
		"re_user_name":     "张三",
		"amount":           "100",
		"desc":             "六月出差报销费用",
		"spbill_create_ip": "10.2.3.10",
		"workwx_sign":      "FB5C2BD04CB52ED2176F5F883F9EA81F",
		"ww_msg_type":      "NORMAL_MSG",
		"act_name":         "示例项目",
		"sign":             "BC0B057C9F1FA1599509BA627281322F",
	})

	assert.Nil(t, err)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`<xml>
	<return_code><![CDATA[SUCCESS]]></return_code>
	<return_msg><![CDATA[ok]]></return_msg>
	<appid><![CDATA[wxec38b8ff840b8888]]></appid>
	<mch_id><![CDATA[1900000109]]></mch_id>
	<device_info><![CDATA[013467007045764]]></device_info>
	<nonce_str><![CDATA[lxuDzMnRjpcXzxLx0q]]></nonce_str>
	<result_code><![CDATA[SUCCESS]]></result_code>
	<partner_trade_no><![CDATA[100000982017072019616]]></partner_trade_no>
	<payment_no><![CDATA[1000018301201505190181489473]]></payment_no>
	<payment_time><![CDATA[2017-07-20 22:05:59]]></payment_time>
	<sign>2D3F4D8C4D68DA71684846F835552300</sign>
</xml>`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/paywwsptrans2pocket", body).Return(resp, nil)

	mch := New("1900000109", "192006250b4c09247ec02edce69f6a2d", p12cert)

	mch.nonce = func() string {
		return "3PG2J4ILTKCH16CQ2502SI8ZNMTM67VS"
	}

	mch.SetTLSClient(wx.WithHTTPClient(client))

	r, err := mch.Do(context.TODO(), TransferToPocket("wxe062425f740c8888", "192006250b4c09247ec02edce69f6a2d", &ParamsCorpTransfer{
		PartnerTradeNO: "100000982017072019616",
		OpenID:         "ohO4Gt7wVPxIT1A9GjFaMYMiZY1s",
		CheckName:      "NO_CHECK",
		Amount:         100,
		Desc:           "六月出差报销费用",
		SpbillCreateIP: "10.2.3.10",
		WWMsgType:      "NORMAL_MSG",
		ActName:        "示例项目",
		ReUserName:     "张三",
		DeviceInfo:     "013467007045764",
	}))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
		"return_code":      "SUCCESS",
		"return_msg":       "ok",
		"appid":            "wxec38b8ff840b8888",
		"mch_id":           "1900000109",
		"device_info":      "013467007045764",
		"nonce_str":        "lxuDzMnRjpcXzxLx0q",
		"result_code":      "SUCCESS",
		"partner_trade_no": "100000982017072019616",
		"payment_no":       "1000018301201505190181489473",
		"payment_time":     "2017-07-20 22:05:59",
		"sign":             "2D3F4D8C4D68DA71684846F835552300",
	}, r)
}

func TestQueryTransferPocket(t *testing.T) {
	body, err := wx.FormatMap2XMLForTest(wx.WXML{
		"appid":            "wxe062425f740c8888",
		"mch_id":           "10000097",
		"partner_trade_no": "0010010404201411170000046545",
		"nonce_str":        "50780e0cca98c8c8e814883e5caa672e",
		"sign":             "AAE662C19C1B5ED4FB273D53F73D07DF",
	})

	assert.Nil(t, err)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`<xml>
	<return_code><![CDATA[SUCCESS]]></return_code>
	<return_msg><![CDATA[获取成功]]></return_msg>
	<result_code><![CDATA[SUCCESS]]></result_code>
	<mch_id>10000097</mch_id>
	<appid><![CDATA[wxe062425f740c30d8]]></appid>
	<detail_id><![CDATA[1000000000201503283103439304]]></detail_id>
	<partner_trade_no><![CDATA[0010010404201411170000046545]]></partner_trade_no>
	<status><![CDATA[SUCCESS]]></status>
	<payment_amount>100</payment_amount>
	<openid><![CDATA[oxTWIuGaIt6gTKsQRLau2M0yL16E]]></openid>
	<transfer_time><![CDATA[2017-07-22 20:10:15]]></transfer_time>
	<transfer_name><![CDATA[测试]]></transfer_name>
	<desc><![CDATA[付款测试]]></desc>
	<sign>A5CC7D1D0431EE5E54B67B7F6AF84220</sign>
</xml>`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/querywwsptrans2pocket", body).Return(resp, nil)

	mch := New("10000097", "192006250b4c09247ec02edce69f6a2d", p12cert)

	mch.nonce = func() string {
		return "50780e0cca98c8c8e814883e5caa672e"
	}

	mch.SetTLSClient(wx.WithHTTPClient(client))

	r, err := mch.Do(context.TODO(), QueryTransferPocket("wxe062425f740c8888", "0010010404201411170000046545"))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
		"return_code":      "SUCCESS",
		"return_msg":       "获取成功",
		"result_code":      "SUCCESS",
		"mch_id":           "10000097",
		"appid":            "wxe062425f740c30d8",
		"detail_id":        "1000000000201503283103439304",
		"partner_trade_no": "0010010404201411170000046545",
		"status":           "SUCCESS",
		"payment_amount":   "100",
		"openid":           "oxTWIuGaIt6gTKsQRLau2M0yL16E",
		"transfer_name":    "测试",
		"transfer_time":    "2017-07-22 20:10:15",
		"desc":             "付款测试",
		"sign":             "A5CC7D1D0431EE5E54B67B7F6AF84220",
	}, r)
}
