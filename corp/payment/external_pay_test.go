package payment

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/chenghonour/gochat/corp"
	"github.com/chenghonour/gochat/mock"
	"github.com/chenghonour/gochat/wx"
)

func TestAddMerchant(t *testing.T) {
	body := []byte(`{"mch_id":"12334","merchant_name":"xxx"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalpay/addmerchant?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", AddMerchant("12334", "xxx"))

	assert.Nil(t, err)
}

func TestGetMerchant(t *testing.T) {
	body := []byte(`{"mch_id":"12334"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"bind_status": 0,
	"mch_id": "12334",
	"merchant_name": "test",
	"allow_use_scope": {
		"user": [
			"zhangsan",
			"lisi"
		],
		"partyid": [
			1
		],
		"tagid": [
			1,
			2,
			3
		]
	}
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalpay/getmerchant?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultMerchantGet)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetMerchant("12334", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMerchantGet{
		BindStatus:   0,
		MchID:        "12334",
		MerchantName: "test",
		AllowUseScope: &AllowUseScope{
			User:    []string{"zhangsan", "lisi"},
			PartyID: []int64{1},
			TagID:   []int64{1, 2, 3},
		},
	}, result)
}

func TestDeleteMerchant(t *testing.T) {
	body := []byte(`{"mch_id":"12334"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalpay/delmerchant?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", DeleteMerchant("12334"))

	assert.Nil(t, err)
}

func TestSetMchUseScope(t *testing.T) {
	body := []byte(`{"mch_id":"12334","allow_use_scope":{"user":["zhangsan","lisi"],"partyid":[1],"tagid":[1,2,3]}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalpay/set_mch_use_scope?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	scope := &AllowUseScope{
		User:    []string{"zhangsan", "lisi"},
		PartyID: []int64{1},
		TagID:   []int64{1, 2, 3},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SetMchUseScope("12334", scope))

	assert.Nil(t, err)
}

func TestGetBillList(t *testing.T) {
	body := []byte(`{"begin_time":1605171726,"end_time":1605172726,"payee_userid":"zhangshan","cursor":"CURSOR","limit":10}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode":0,
	"errmsg":"ok",
	"next_cursor":"CURSOR",
	"bill_list":[
		{
			"transaction_id":"xxxxx",
			"trade_state":1,
			"pay_time":12345,
			"out_trade_no":"xxxx",
			"external_userid":"xxxx",
			"total_fee":100,
			"payee_userid":"zhangshan",
			"payment_type":1,
			"mch_id":"123454",
			"remark":"xxxx",
			"commodity_list":[
				{
					"description":"手机",
					"amount":1
				}
			],
			"total_refund_fee":100,
			"refund_list":[
				{
					"out_refund_no":"xx",
					"refund_userid":"xxx",
					"refund_comment":"xxx",
					"refund_reqtime":1605171790,
					"refund_status":1,
					"refund_fee":100
				}
			],
			"payer_info":{
				"name":"xxx",
				"phone":"xxx",
				"address":"xxx"
			}
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalpay/get_bill_list?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsBillList{
		BeginTime:   1605171726,
		EndTime:     1605172726,
		PayeeUserID: "zhangshan",
		Cursor:      "CURSOR",
		Limit:       10,
	}
	result := new(ResultBillList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetBillList(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultBillList{
		NextCursor: "CURSOR",
		BillList: []*BillInfo{
			{
				TransactionID:  "xxxxx",
				TradeState:     1,
				PayTime:        12345,
				OutTradeNO:     "xxxx",
				ExternalUserID: "xxxx",
				TotalFee:       100,
				PayeeUserID:    "zhangshan",
				PaymentType:    1,
				MchID:          "123454",
				Remark:         "xxxx",
				TotalRefundFee: 100,
				CommodityList: []*CommodityInfo{
					{
						Description: "手机",
						Amount:      1,
					},
				},
				RefundList: []*RefundInfo{
					{
						OutRefundNO:   "xx",
						RefundUserID:  "xxx",
						RefundComment: "xxx",
						RefundReqTime: 1605171790,
						RefundStatus:  1,
						RefundFee:     100,
					},
				},
				PayerInfo: &PayerInfo{
					Name:    "xxx",
					Phone:   "xxx",
					Address: "xxx",
				},
			},
		},
	}, result)
}
