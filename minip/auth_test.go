package minip

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/wx"
)

func TestGetPaidUnionIDByTransactionID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/getpaidunionid?access_token=ACCESS_TOKEN&openid=OPENID&transaction_id=TRANSACTION_ID").Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"unionid": "oTmHYjg-tElZ68xxxxxxxxhy1Rgk"
	}`), nil)

	mp := New("APPID", "APPSECRET")
	mp.client = client

	result := new(ResultPaidUnionID)

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", GetPaidUnionIDByTransactionID("OPENID", "TRANSACTION_ID", result))

	assert.Nil(t, err)
	assert.Equal(t, "oTmHYjg-tElZ68xxxxxxxxhy1Rgk", result.UnionID)
}

func TestGetPaidUnionIDByOutTradeNO(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/getpaidunionid?access_token=ACCESS_TOKEN&mch_id=MCH_ID&openid=OPENID&out_trade_no=OUT_TRADE_NO").Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"unionid": "oTmHYjg-tElZ68xxxxxxxxhy1Rgk"
	}`), nil)

	mp := New("APPID", "APPSECRET")
	mp.client = client

	result := new(ResultPaidUnionID)

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", GetPaidUnionIDByOutTradeNO("OPENID", "MCH_ID", "OUT_TRADE_NO", result))

	assert.Nil(t, err)
	assert.Equal(t, "oTmHYjg-tElZ68xxxxxxxxhy1Rgk", result.UnionID)
}
