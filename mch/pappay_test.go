package mch

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/wx"
	"github.com/stretchr/testify/assert"
)

func TestAPPEntrust(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().PostXML(gomock.AssignableToTypeOf(context.TODO()), "https://api.mch.weixin.qq.com/papay/preentrustweb", wx.WXML{
		"appid":                    "wx2421b1c4370ec43b",
		"mch_id":                   "10000100",
		"plan_id":                  "12535",
		"contract_code":            "100000",
		"request_serial":           "1000",
		"contract_display_account": "微信代扣",
		"notify_url":               "https://weixin.qq.com",
		"version":                  "1.0",
		"timestamp":                "1414488825",
		"return_app":               "Y",
		"sign_type":                "MD5",
		"sign":                     "588134C9FA5B9D4E89E44FA303F6CB6F",
	}).Return([]byte(`<xml>
	<return_code>SUCCESS</return_code>
	<return_msg>OK</return_msg>
	<appid>wx2421b1c4370ec43b</appid>
	<mch_id>10000100</mch_id>
	<nonce_str>IITRi8Iabbblz1Jc</nonce_str>
	<sign>A07C2571BA6F4FBFDB82490C97776AB4</sign>
	<result_code>SUCCESS</result_code>
	<pre_entrustweb_id>5778aadY9nltAsZzXixCkFIGYnV2V</pre_entrustweb_id>
</xml>`), nil)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "IITRi8Iabbblz1Jc"
	}
	mch.client = client
	mch.tlsClient = client

	r, err := mch.Do(context.TODO(), APPEntrust(&Contract{
		PlanID:                 "12535",
		ContractCode:           "100000",
		RequestSerial:          1000,
		ContractDisplayAccount: "微信代扣",
		Timestamp:              1414488825,
		NotifyURL:              "https://weixin.qq.com",
		ReturnAPP:              true,
	}))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
		"return_code":       "SUCCESS",
		"return_msg":        "OK",
		"appid":             "wx2421b1c4370ec43b",
		"mch_id":            "10000100",
		"nonce_str":         "IITRi8Iabbblz1Jc",
		"sign":              "A07C2571BA6F4FBFDB82490C97776AB4",
		"result_code":       "SUCCESS",
		"pre_entrustweb_id": "5778aadY9nltAsZzXixCkFIGYnV2V",
	}, r)
}

func TestOAEntrust(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "IITRi8Iabbblz1Jc"
	}
	mch.client = client
	mch.tlsClient = client

	r, err := mch.Do(context.TODO(), OAEntrust(&Contract{
		PlanID:                 "106",
		ContractCode:           "122",
		RequestSerial:          123,
		ContractDisplayAccount: "name1",
		Timestamp:              1414488825,
		NotifyURL:              "www.qq.com/test/papay",
	}))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{"entrust_url": "https://api.mch.weixin.qq.com/papay/entrustweb?appid=wx2421b1c4370ec43b&contract_code=122&contract_display_account=name1&mch_id=10000100&notify_url=www.qq.com%2Ftest%2Fpapay&plan_id=106&request_serial=123&sign=48F3F8F08E560D736E8D0FEFACBB012E&sign_type=MD5&timestamp=1414488825&version=1.0"}, r)
}

func TestMPEntrust(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "IITRi8Iabbblz1Jc"
	}
	mch.client = client
	mch.tlsClient = client

	r, err := mch.Do(context.TODO(), MPEntrust(&Contract{
		PlanID:                 "106",
		ContractCode:           "122",
		RequestSerial:          123,
		ContractDisplayAccount: "张三",
		Timestamp:              1414488825,
		NotifyURL:              "https://www.qq.com/test/papay",
	}))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
		"appid":                    "wx2421b1c4370ec43b",
		"mch_id":                   "10000100",
		"plan_id":                  "106",
		"contract_code":            "122",
		"request_serial":           "123",
		"contract_display_account": "张三",
		"notify_url":               "https://www.qq.com/test/papay",
		"timestamp":                "1414488825",
		"sign_type":                "MD5",
		"sign":                     "E0EC5B06A03B55F2B1FC754AB04D8381",
	}, r)
}

func TestH5Entrust(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "IITRi8Iabbblz1Jc"
	}
	mch.client = client
	mch.tlsClient = client

	r, err := mch.Do(context.TODO(), H5Entrust(&Contract{
		PlanID:                 "106",
		ContractCode:           "122",
		RequestSerial:          123,
		ContractDisplayAccount: "name1",
		Timestamp:              1414488825,
		NotifyURL:              "www.qq.com/test/papay",
		SpbillCreateIP:         "12.1.1.12",
		ReturnAPPID:            "wxcbda96de0b165542",
	}))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{"entrust_url": "https://api.mch.weixin.qq.com/papay/h5entrustweb?appid=wx2421b1c4370ec43b&clientip=12.1.1.12&contract_code=122&contract_display_account=name1&mch_id=10000100&notify_url=www.qq.com%2Ftest%2Fpapay&plan_id=106&request_serial=123&return_appid=wxcbda96de0b165542&sign=CE76472E3C209CB2B3F6FC6A649B6849D4BCC78F4A4A820EEF4D5A55EF3F2660&sign_type=HMAC-SHA256&timestamp=1414488825&version=1.0"}, r)
}

func TestEntrustByOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().PostXML(gomock.AssignableToTypeOf(context.TODO()), "https://api.mch.weixin.qq.com/pay/contractorder", wx.WXML{
		"appid":                    "wx2421b1c4370ec43b",
		"mch_id":                   "10000100",
		"contract_mchid":           "10000100",
		"contract_appid":           "wx2421b1c4370ec43b",
		"out_trade_no":             "123456",
		"device_info":              "013467007045764",
		"body":                     "Ipad mini 16G 白色",
		"detail":                   "Ipad mini 16G 白色",
		"notify_url":               "https://weixin.qq.com",
		"total_fee":                "888",
		"fee_type":                 "CNY",
		"spbill_create_ip":         "123.12.12.123",
		"trade_type":               "JSAPI",
		"plan_id":                  "123",
		"contract_code":            "100001256",
		"request_serial":           "1000",
		"contract_display_account": "微信代扣",
		"contract_notify_url":      "https://yoursite.com",
		"nonce_str":                "5K8264ILTKCH16CQ2502SI8ZNMTM67VS",
		"sign_type":                "MD5",
		"sign":                     "B59547B883DA4E3E3483DFCD1FC0CB98",
	}).Return([]byte(`<xml>
	<return_code>SUCCESS</return_code>
	<result_code>SUCCESS</result_code>
	<appid>wx2421b1c4370ec43b</appid>
	<mch_id>10000100</mch_id>
	<nonce_str>IITRi8Iabbblz1Jc</nonce_str>
	<sign>27CB53BB1FD0528DB119910CB1A456E0</sign>
	<prepay_id>wx201410272009395522657a690389285100</prepay_id>
	<trade_type>JSAPI</trade_type>
	<code_url>weixin://wxpay/s/An4baqw</code_url>
	<plan_id>123</plan_id>
	<out_trade_no>123456</out_trade_no>
</xml>`), nil)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "5K8264ILTKCH16CQ2502SI8ZNMTM67VS"
	}
	mch.client = client
	mch.tlsClient = client

	r, err := mch.Do(context.TODO(), EntrustByOrder(&ContractOrder{
		OutTradeNO:             "123456",
		TotalFee:               888,
		SpbillCreateIP:         "123.12.12.123",
		TradeType:              TradeJSAPI,
		Body:                   "Ipad mini 16G 白色",
		PlanID:                 "123",
		ContractCode:           "100001256",
		RequestSerial:          1000,
		ContractDisplayAccount: "微信代扣",
		PaymentNotifyURL:       "https://weixin.qq.com",
		ContractNotifyURL:      "https://yoursite.com",
		DeviceInfo:             "013467007045764",
		Detail:                 "Ipad mini 16G 白色",
	}))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
		"return_code":  "SUCCESS",
		"result_code":  "SUCCESS",
		"appid":        "wx2421b1c4370ec43b",
		"mch_id":       "10000100",
		"nonce_str":    "IITRi8Iabbblz1Jc",
		"sign":         "27CB53BB1FD0528DB119910CB1A456E0",
		"prepay_id":    "wx201410272009395522657a690389285100",
		"trade_type":   "JSAPI",
		"code_url":     "weixin://wxpay/s/An4baqw",
		"plan_id":      "123",
		"out_trade_no": "123456",
	}, r)
}

func TestQueryContractByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().PostXML(gomock.AssignableToTypeOf(context.TODO()), "https://api.mch.weixin.qq.com/papay/querycontract", wx.WXML{
		"appid":       "wx2421b1c4370ec43b",
		"mch_id":      "10000100",
		"contract_id": "201509160000028648",
		"version":     "1.0",
		"sign_type":   "MD5",
		"sign":        "423CE9A1BA3B99329D7B9A4ACC8B57EA",
	}).Return([]byte(`<xml>
	<return_code>SUCCESS</return_code>
	<result_code>SUCCESS</result_code>
	<appid>wx2421b1c4370ec43b</appid>
	<mch_id>10000100</mch_id>
	<contract_id>201509160000028648</contract_id>
	<plan_id>123</plan_id>
	<openid>oHZx6uMbIG46UXQ3SKxVYEgw1LZs</openid>
	<request_serial>1000</request_serial>
	<contract_code>1023658866</contract_code>
	<contract_display_account>test</contract_display_account>
	<contract_state>1</contract_state>
	<contract_signed_time>1438141845</contract_signed_time>
	<contract_expired_time>1453953047</contract_expired_time>
	<contract_terminated_time>1438157486</contract_terminated_time>
	<contract_termination_mode>3</contract_termination_mode>
	<contract_termination_remark>delete ....</contract_termination_remark>
	<sign>35B3B9261A6A4E75BFB560FE0D6EA8CE</sign>
</xml>`), nil)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "IITRi8Iabbblz1Jc"
	}
	mch.client = client
	mch.tlsClient = client

	r, err := mch.Do(context.TODO(), QueryContractByID("201509160000028648"))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
		"return_code":                 "SUCCESS",
		"result_code":                 "SUCCESS",
		"appid":                       "wx2421b1c4370ec43b",
		"mch_id":                      "10000100",
		"contract_id":                 "201509160000028648",
		"plan_id":                     "123",
		"openid":                      "oHZx6uMbIG46UXQ3SKxVYEgw1LZs",
		"request_serial":              "1000",
		"contract_code":               "1023658866",
		"contract_display_account":    "test",
		"contract_state":              "1",
		"contract_signed_time":        "1438141845",
		"contract_expired_time":       "1453953047",
		"contract_terminated_time":    "1438157486",
		"contract_termination_mode":   "3",
		"contract_termination_remark": "delete ....",
		"sign":                        "35B3B9261A6A4E75BFB560FE0D6EA8CE",
	}, r)
}

func TestQueryContractByCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().PostXML(gomock.AssignableToTypeOf(context.TODO()), "https://api.mch.weixin.qq.com/papay/querycontract", wx.WXML{
		"appid":         "wx2421b1c4370ec43b",
		"mch_id":        "10000100",
		"plan_id":       "123",
		"contract_code": "1023658866",
		"version":       "1.0",
		"sign_type":     "MD5",
		"sign":          "8657D130240326FF8A9BA9FFD3A22C54",
	}).Return([]byte(`<xml>
	<return_code>SUCCESS</return_code>
	<result_code>SUCCESS</result_code>
	<appid>wx2421b1c4370ec43b</appid>
	<mch_id>10000100</mch_id>
	<contract_id>201509160000028648</contract_id>
	<plan_id>123</plan_id>
	<openid>oHZx6uMbIG46UXQ3SKxVYEgw1LZs</openid>
	<request_serial>1000</request_serial>
	<contract_code>1023658866</contract_code>
	<contract_display_account>test</contract_display_account>
	<contract_state>1</contract_state>
	<contract_signed_time>1438141845</contract_signed_time>
	<contract_expired_time>1453953047</contract_expired_time>
	<contract_terminated_time>1438157486</contract_terminated_time>
	<contract_termination_mode>3</contract_termination_mode>
	<contract_termination_remark>delete ....</contract_termination_remark>
	<sign>35B3B9261A6A4E75BFB560FE0D6EA8CE</sign>
</xml>`), nil)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "IITRi8Iabbblz1Jc"
	}
	mch.client = client
	mch.tlsClient = client

	r, err := mch.Do(context.TODO(), QueryContractByCode("123", "1023658866"))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
		"return_code":                 "SUCCESS",
		"result_code":                 "SUCCESS",
		"appid":                       "wx2421b1c4370ec43b",
		"mch_id":                      "10000100",
		"contract_id":                 "201509160000028648",
		"plan_id":                     "123",
		"openid":                      "oHZx6uMbIG46UXQ3SKxVYEgw1LZs",
		"request_serial":              "1000",
		"contract_code":               "1023658866",
		"contract_display_account":    "test",
		"contract_state":              "1",
		"contract_signed_time":        "1438141845",
		"contract_expired_time":       "1453953047",
		"contract_terminated_time":    "1438157486",
		"contract_termination_mode":   "3",
		"contract_termination_remark": "delete ....",
		"sign":                        "35B3B9261A6A4E75BFB560FE0D6EA8CE",
	}, r)
}

func TestPappayApply(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().PostXML(gomock.AssignableToTypeOf(context.TODO()), "https://api.mch.weixin.qq.com/pay/pappayapply", wx.WXML{
		"appid":            "wx2421b1c4370ec43b",
		"mch_id":           "10000100",
		"out_trade_no":     "217752501201407033233368018",
		"body":             "水电代扣",
		"total_fee":        "888",
		"fee_type":         "CNY",
		"spbill_create_ip": "8.8.8.8",
		"notify_url":       "http://yoursite.com/wxpay.html",
		"trade_type":       "PAP",
		"contract_id":      "Wx15463511252015071056489715",
		"nonce_str":        "5K8264ILTKCH16CQ2502SI8ZNMTM67VS",
		"sign_type":        "MD5",
		"sign":             "B93354AC5A479C50A85019E8B47C01FE",
	}).Return([]byte(`<xml>
	<return_code>SUCCESS</return_code>
	<return_msg>OK</return_msg>
	<appid>wx2421b1c4370ec43b</appid>
	<mch_id>10000100</mch_id>
	<nonce_str>IITRi8Iabbblz1Jc</nonce_str>
	<sign>1D001A3A187A984976FDB371813F898F</sign>
	<result_code>SUCCESS</result_code>
</xml>`), nil)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "5K8264ILTKCH16CQ2502SI8ZNMTM67VS"
	}
	mch.client = client
	mch.tlsClient = client

	r, err := mch.Do(context.TODO(), PappayApply(&PappayData{
		OutTradeNO:     "217752501201407033233368018",
		TotalFee:       888,
		SpbillCreateIP: "8.8.8.8",
		ContractID:     "Wx15463511252015071056489715",
		Body:           "水电代扣",
		NotifyURL:      "http://yoursite.com/wxpay.html",
	}))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
		"return_code": "SUCCESS",
		"return_msg":  "OK",
		"appid":       "wx2421b1c4370ec43b",
		"mch_id":      "10000100",
		"nonce_str":   "IITRi8Iabbblz1Jc",
		"sign":        "1D001A3A187A984976FDB371813F898F",
		"result_code": "SUCCESS",
	}, r)
}

func TestDeleteContractByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().PostXML(gomock.AssignableToTypeOf(context.TODO()), "https://api.mch.weixin.qq.com/papay/deletecontract", wx.WXML{
		"appid":                       "wx2421b1c4370ec43b",
		"mch_id":                      "10000100",
		"contract_id":                 "100005698",
		"contract_termination_remark": "原因",
		"version":                     "1.0",
		"sign_type":                   "MD5",
		"sign":                        "9A1F0B64467723BDF0D26CC39ED0B001",
	}).Return([]byte(`<xml>
	<return_code>SUCCESS</return_code>
	<result_code>SUCCESS</result_code>
	<appid>wx2421b1c4370ec43b</appid>
	<mch_id>10000100</mch_id>
	<contract_id>100005698</contract_id>
	<sign>D1F898877B9FC523A6F2FC993BE5B78F</sign>
</xml>`), nil)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "IITRi8Iabbblz1Jc"
	}
	mch.client = client
	mch.tlsClient = client

	r, err := mch.Do(context.TODO(), DeleteContractByID("100005698", "原因"))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
		"return_code": "SUCCESS",
		"result_code": "SUCCESS",
		"appid":       "wx2421b1c4370ec43b",
		"mch_id":      "10000100",
		"contract_id": "100005698",
		"sign":        "D1F898877B9FC523A6F2FC993BE5B78F",
	}, r)
}

func TestDeleteContractByCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().PostXML(gomock.AssignableToTypeOf(context.TODO()), "https://api.mch.weixin.qq.com/papay/deletecontract", wx.WXML{
		"appid":                       "wx2421b1c4370ec43b",
		"mch_id":                      "10000100",
		"plan_id":                     "12251",
		"contract_code":               "1234",
		"contract_termination_remark": "原因",
		"version":                     "1.0",
		"sign_type":                   "MD5",
		"sign":                        "9AA3B7174E2B71DA6268741A70E1DDDF",
	}).Return([]byte(`<xml>
	<return_code>SUCCESS</return_code>
	<result_code>SUCCESS</result_code>
	<appid>wx2421b1c4370ec43b</appid>
	<mch_id>10000100</mch_id>
	<contract_id>100005698</contract_id>
	<sign>D1F898877B9FC523A6F2FC993BE5B78F</sign>
</xml>`), nil)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "IITRi8Iabbblz1Jc"
	}
	mch.client = client
	mch.tlsClient = client

	r, err := mch.Do(context.TODO(), DeleteContractByCode("12251", "1234", "原因"))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
		"return_code": "SUCCESS",
		"result_code": "SUCCESS",
		"appid":       "wx2421b1c4370ec43b",
		"mch_id":      "10000100",
		"contract_id": "100005698",
		"sign":        "D1F898877B9FC523A6F2FC993BE5B78F",
	}, r)
}

func TestQueryPappayByTransactionID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().PostXML(gomock.AssignableToTypeOf(context.TODO()), "https://api.mch.weixin.qq.com/pay/paporderquery", wx.WXML{
		"appid":          "wx2421b1c4370ec43b",
		"mch_id":         "10000100",
		"transaction_id": "1008450740201411110005820873",
		"nonce_str":      "0b9f35f484df17a732e537c37708d1d0",
		"sign_type":      "MD5",
		"sign":           "264E5038F1CB9D66132E769ABB5B745C",
	}).Return([]byte(`<xml>
	<return_code>SUCCESS</return_code>
	<return_msg>OK</return_msg>
	<appid>wx2421b1c4370ec43b</appid>
	<mch_id>10000100</mch_id>
	<result_code>SUCCESS</result_code>
	<device_info>1000</device_info>
	<openid>oUpF8uN95-Ptaags6E_roPHg7AG0</openid>
	<is_subscribe>Y</is_subscribe>
	<trade_type>MICROPAY</trade_type>
	<bank_type>CCB_DEBIT</bank_type>
	<total_fee>1</total_fee>
	<fee_type>CNY</fee_type>
	<transaction_id>1008450740201411110005820873</transaction_id>
	<out_trade_no>1415757673</out_trade_no>
	<attach>订单额外描述</attach>
	<time_end>20141111170043</time_end>
	<trade_state>SUCCESS</trade_state>
	<nonce_str>TN55wO9Pba5yENl8</nonce_str>
	<sign>9C2A03FD2D080D1B9618946C73C7608D</sign>
</xml>`), nil)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "0b9f35f484df17a732e537c37708d1d0"
	}
	mch.client = client
	mch.tlsClient = client

	r, err := mch.Do(context.TODO(), QueryPappayByTransactionID("1008450740201411110005820873"))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
		"return_code":    "SUCCESS",
		"return_msg":     "OK",
		"appid":          "wx2421b1c4370ec43b",
		"mch_id":         "10000100",
		"result_code":    "SUCCESS",
		"device_info":    "1000",
		"openid":         "oUpF8uN95-Ptaags6E_roPHg7AG0",
		"is_subscribe":   "Y",
		"trade_type":     "MICROPAY",
		"bank_type":      "CCB_DEBIT",
		"total_fee":      "1",
		"fee_type":       "CNY",
		"transaction_id": "1008450740201411110005820873",
		"out_trade_no":   "1415757673",
		"attach":         "订单额外描述",
		"time_end":       "20141111170043",
		"trade_state":    "SUCCESS",
		"nonce_str":      "TN55wO9Pba5yENl8",
		"sign":           "9C2A03FD2D080D1B9618946C73C7608D",
	}, r)
}

func TestQueryPappayByOutTradeNO(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().PostXML(gomock.AssignableToTypeOf(context.TODO()), "https://api.mch.weixin.qq.com/pay/paporderquery", wx.WXML{
		"appid":        "wx2421b1c4370ec43b",
		"mch_id":       "10000100",
		"out_trade_no": "1415757673",
		"nonce_str":    "0b9f35f484df17a732e537c37708d1d0",
		"sign_type":    "MD5",
		"sign":         "5F14ED52C2F179580A1DED73268A1009",
	}).Return([]byte(`<xml>
	<return_code>SUCCESS</return_code>
	<return_msg>OK</return_msg>
	<appid>wx2421b1c4370ec43b</appid>
	<mch_id>10000100</mch_id>
	<result_code>SUCCESS</result_code>
	<device_info>1000</device_info>
	<openid>oUpF8uN95-Ptaags6E_roPHg7AG0</openid>
	<is_subscribe>Y</is_subscribe>
	<trade_type>MICROPAY</trade_type>
	<bank_type>CCB_DEBIT</bank_type>
	<total_fee>1</total_fee>
	<fee_type>CNY</fee_type>
	<transaction_id>1008450740201411110005820873</transaction_id>
	<out_trade_no>1415757673</out_trade_no>
	<attach>订单额外描述</attach>
	<time_end>20141111170043</time_end>
	<trade_state>SUCCESS</trade_state>
	<nonce_str>TN55wO9Pba5yENl8</nonce_str>
	<sign>9C2A03FD2D080D1B9618946C73C7608D</sign>
</xml>`), nil)

	mch := New("wx2421b1c4370ec43b", "10000100", "192006250b4c09247ec02edce69f6a2d")

	mch.nonce = func(size int) string {
		return "0b9f35f484df17a732e537c37708d1d0"
	}
	mch.client = client
	mch.tlsClient = client

	r, err := mch.Do(context.TODO(), QueryPappayByOutTradeNO("1415757673"))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
		"return_code":    "SUCCESS",
		"return_msg":     "OK",
		"appid":          "wx2421b1c4370ec43b",
		"mch_id":         "10000100",
		"result_code":    "SUCCESS",
		"device_info":    "1000",
		"openid":         "oUpF8uN95-Ptaags6E_roPHg7AG0",
		"is_subscribe":   "Y",
		"trade_type":     "MICROPAY",
		"bank_type":      "CCB_DEBIT",
		"total_fee":      "1",
		"fee_type":       "CNY",
		"transaction_id": "1008450740201411110005820873",
		"out_trade_no":   "1415757673",
		"attach":         "订单额外描述",
		"time_end":       "20141111170043",
		"trade_state":    "SUCCESS",
		"nonce_str":      "TN55wO9Pba5yENl8",
		"sign":           "9C2A03FD2D080D1B9618946C73C7608D",
	}, r)
}
