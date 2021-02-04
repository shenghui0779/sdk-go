package mch

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/wx"
	"github.com/stretchr/testify/assert"
)

func TestTransferToBalance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().PostXML(gomock.AssignableToTypeOf(context.TODO()), "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers", wx.WXML{
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
	}).Return([]byte(`<xml>
	<return_code>SUCCESS</return_code>
	<mch_appid>wx2421b1c4370ec43b</mch_appid>
	<mchid>10000100</mchid>
	<nonce_str>lxuDzMnRjpcXzxLx0q</nonce_str>
	<result_code>SUCCESS</result_code>
	<partner_trade_no>10013574201505191526582441</partner_trade_no>
	<payment_no>1000018301201505190181489473</payment_no>
	<payment_time>2015-05-19 15:26:59</payment_time>
</xml>`), nil)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "3PG2J4ILTKCH16CQ2502SI8ZNMTM67VS"
	}

	mch.tlsClient = client

	r, err := mch.Do(context.TODO(), TransferToBalance(&TransferBalanceData{
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

func TestQueryTransferBalanceOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().PostXML(gomock.AssignableToTypeOf(context.TODO()), "https://api.mch.weixin.qq.com/mmpaymkttransfers/gettransferinfo", wx.WXML{
		"appid":            "wx2421b1c4370ec43b",
		"mch_id":           "10000100",
		"partner_trade_no": "1000005901201407261446939628",
		"nonce_str":        "50780e0cca98c8c8e814883e5caa672e",
		"sign_type":        "MD5",
		"sign":             "DF0024F9502E233115C0198912B4EB5D",
	}).Return([]byte(`<xml>
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
</xml>`), nil)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "50780e0cca98c8c8e814883e5caa672e"
	}

	mch.tlsClient = client

	r, err := mch.Do(context.TODO(), QueryTransferBalanceOrder("1000005901201407261446939628"))

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

// RSA每次加密结果不同，导致签名会变化
// func TestTransferToBankCard(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	client := wx.NewMockHTTPClient(ctrl)

// 	client.EXPECT().PostXML(gomock.AssignableToTypeOf(context.TODO()), "https://api.mch.weixin.qq.com/mmpaysptrans/pay_bank", wx.WXML{
// 		"mch_id":           "10000100",
// 		"partner_trade_no": "1212121221278",
// 		"amount":           "500",
// 		"bank_code":        "1002",
// 		"enc_bank_no":      "en4Y1l7D0dK+cRuLLDquRuswp9bsuB6MQke+bn0S+MF9sDKIDp4Tkiml9v90uSQof3nIaOZ/q1UTFFV7/bvrkQc6+PKxbx/Y9YcdmrUAS2HCB7uFRVmsu4xBtbDzAR0wnnTuUcr6DJz/HxgE9EUpXyhHUpNgXB4/GOxgJA5uBimBKA6z46AmGxLcgOkvOU9bo9+hgYDCrOOEwRiN1XC18llAsqjZPAJqkZibv9cEZ5zvmrT8zRBoi+L1N9ZUGuxvq1GpbsBOFE0PP4IFP60R216pz9/nhFBKi3rF0ohF3mnjBmycOVaOK0xm8lcEQQEV+94/4bqnIJOSg8UmHrArRQ==",
// 		"enc_true_name":    "ABpj6B97My6jKc2TwbkXM/W55LmlxmldJHhKr3n2cr36UeQCGOKlc3Cc1sQytng4hKrDd+qrXT3fmoRvxc10mnViGKdwq1G6XAmGYMMs2Pm0edzqWicrTi8/dcXoVaxLj4ZwCBm+8OtCpJefxGi9xZjpnXpUvEa2hzlPbghFNoPMHIOdECwzvYMqAM2OoRwqicTZgroRS0jI88NhM5UTn00ZwFSoN3VeFkkDSeKXZ25232l51WjBqyg6JLRGltPtiKwaNhCd5cxkPrCJrMJAzJ8PVQmBrEfRnyHDJiYGIQZ1bGoB9eKTN/+cjcGWuxyXDrpdIc0DJzCy/5Yswrv+qg==",
// 		"desc":             "test",
// 		"nonce_str":        "50780e0cca98c8c8e814883e5caa672e",
// 		"sign_type":        "MD5",
// 		"sign":             "93FD9CF5C2D3F2D6016A168F69D221D5",
// 	}).Return(wx.WXML{
// 		"return_code":      "SUCCESS",
// 		"mch_id":           "10000100",
// 		"nonce_str":        "50780e0cca98c8c8e814883e5caa672e",
// 		"result_code":      "SUCCESS",
// 		"partner_trade_no": "1212121221278",
// 		"amount":           "500",
// 		"payment_no":       "10000600500852017030900000020006012",
// 		"cmms_amt":         "0",
// 	}, nil)

// 	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

// 	mch.nonce = func(size int) string {
// 		return "50780e0cca98c8c8e814883e5caa672e"
// 	}

// 	mch.tlsClient = client

// 	r, err := mch.Do(context.TODO(), TransferToBankCard(&TransferBankCardData{
// 		PartnerTradeNO: "1212121221278",
// 		EncBankNO:      "6221882600114166800",
// 		EncTrueName:    "张三",
// 		BankCode:       "1002",
// 		Amount:         500,
// 		Desc:           "test",
// 	}, publicKey))

// 	assert.Nil(t, err)
// 	assert.Equal(t, wx.WXML{
// 		"return_code":      "SUCCESS",
// 		"mch_id":           "10000100",
// 		"nonce_str":        "50780e0cca98c8c8e814883e5caa672e",
// 		"result_code":      "SUCCESS",
// 		"partner_trade_no": "1212121221278",
// 		"amount":           "500",
// 		"payment_no":       "10000600500852017030900000020006012",
// 		"cmms_amt":         "0",
// 	}, r)
// }

func TestQueryTransferBankCardOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().PostXML(gomock.AssignableToTypeOf(context.TODO()), "https://api.mch.weixin.qq.com/mmpaysptrans/query_bank", wx.WXML{
		"mch_id":           "10000100",
		"partner_trade_no": "1212121221278",
		"nonce_str":        "50780e0cca98c8c8e814883e5caa672e",
		"sign_type":        "MD5",
		"sign":             "F5F586AE6B1BDB6756D2B1AD0A01BADA",
	}).Return([]byte(`<xml>
	<return_code>SUCCESS</return_code>
	<mch_id>10000100</mch_id>
	<result_code>SUCCESS</result_code>
	<partner_trade_no>1212121221278</partner_trade_no>
	<payment_no>10000600500852017030900000020006012</payment_no>
	<bank_no_md5>2260AB5EF3D290E28EFD3F74FF7A29A0</bank_no_md5>
	<true_name_md5>7F25B325D37790764ABA55DAD8D09B76</true_name_md5>
	<amount>500</amount>
	<status>PROCESSING</status>
	<cmms_amt>0</cmms_amt>
	<create_time>2017-03-09 15:04:04</create_time>
	<reason>福利测试</reason>
</xml>`), nil)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "50780e0cca98c8c8e814883e5caa672e"
	}

	mch.tlsClient = client

	r, err := mch.Do(context.TODO(), QueryTransferBankCardOrder("1212121221278"))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
		"return_code":      "SUCCESS",
		"mch_id":           "10000100",
		"result_code":      "SUCCESS",
		"partner_trade_no": "1212121221278",
		"payment_no":       "10000600500852017030900000020006012",
		"bank_no_md5":      "2260AB5EF3D290E28EFD3F74FF7A29A0",
		"true_name_md5":    "7F25B325D37790764ABA55DAD8D09B76",
		"amount":           "500",
		"status":           "PROCESSING",
		"cmms_amt":         "0",
		"create_time":      "2017-03-09 15:04:04",
		"reason":           "福利测试",
	}, r)
}

func TestRSAPublicKey(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().PostXML(gomock.AssignableToTypeOf(context.TODO()), "https://fraud.mch.weixin.qq.com/risk/getpublickey", wx.WXML{
		"mch_id":    "10000100",
		"nonce_str": "50780e0cca98c8c8e814883e5caa672e",
		"sign_type": "MD5",
		"sign":      "CA227C435D88EE017A9457B657FCA515",
	}).Return([]byte(`<xml>
	<return_code>SUCCESS</return_code>
	<mch_id>10000100</mch_id>
	<result_code>SUCCESS</result_code>
	<pub_key>-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAl1c+37GJSFSqbuHJ/wge
LzxLp7C2GYrjzVAnEF3xgjJVTltkQzdu3u+fcB3c/dgHX/Zdv5fqVoOqvoOMk4N4
zdGeaxN+Cm19c1gsxigNJDtm6Qno1s1T/qPph/zRArylM0N9Z3vWVEq4xI4B4NXk
6IoK/bXc1dwQe5UBzIZyzU5aWfqmTQilWEs7mqro43LTFkhN05QjC7IUFvWEhh6T
wvGYLBSAn+oNw/uSAu6B3c6dh+pslgORCzrIRs68GWsARGZkI/lmOJWEgzQ9KC7b
yHVqEnDDaWQFyQpq30JdP6YTXR/xlKyo8f1DingoSDXAhKMGRKaT4oIFkE6OA3jt
DQIDAQAB
-----END PUBLIC KEY-----</pub_key>
</xml>`), nil)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "50780e0cca98c8c8e814883e5caa672e"
	}

	mch.tlsClient = client

	r, err := mch.Do(context.TODO(), RSAPublicKey())

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
		"return_code": "SUCCESS",
		"mch_id":      "10000100",
		"result_code": "SUCCESS",
		"pub_key": `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAl1c+37GJSFSqbuHJ/wge
LzxLp7C2GYrjzVAnEF3xgjJVTltkQzdu3u+fcB3c/dgHX/Zdv5fqVoOqvoOMk4N4
zdGeaxN+Cm19c1gsxigNJDtm6Qno1s1T/qPph/zRArylM0N9Z3vWVEq4xI4B4NXk
6IoK/bXc1dwQe5UBzIZyzU5aWfqmTQilWEs7mqro43LTFkhN05QjC7IUFvWEhh6T
wvGYLBSAn+oNw/uSAu6B3c6dh+pslgORCzrIRs68GWsARGZkI/lmOJWEgzQ9KC7b
yHVqEnDDaWQFyQpq30JdP6YTXR/xlKyo8f1DingoSDXAhKMGRKaT4oIFkE6OA3jt
DQIDAQAB
-----END PUBLIC KEY-----`,
	}, r)
}
