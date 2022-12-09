package school

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/corp"
	"github.com/shenghui0779/gochat/mock"
)

func TestGetPaymentResult(t *testing.T) {
	body := []byte(`{"payment_id":"xxxx"}`)
	resp := []byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "project_name": "学费",
    "amount": 998,
    "payment_result": [
        {
            "student_userid": "xxxx",
            "trade_state": 1,
            "trade_no": "xxxxx",
            "payer_parent_userid": "zhangshan"
        }
    ]
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/school/get_payment_result?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultPaymentGet)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetPaymentResult("xxxx", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultPaymentGet{
		ProjectName: "学费",
		Amount:      998,
		PaymentResult: []*PaymentInfo{
			{
				StudentUserID:     "xxxx",
				TradeState:        1,
				TradeNO:           "xxxxx",
				PayerParentUserID: "zhangshan",
			},
		},
	}, result)
}

func TestGetTrade(t *testing.T) {
	body := []byte(`{"payment_id":"xxxx","trade_no":"xxxx"}`)
	resp := []byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "transaction_id": "xxxxx",
    "pay_time": 12345
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/school/get_trade?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultTradeGet)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetTrade("xxxx", "xxxx", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultTradeGet{
		TransactionID: "xxxxx",
		PayTime:       12345,
	}, result)
}
