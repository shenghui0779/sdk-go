package minip

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/mock"
	"github.com/shenghui0779/gochat/wx"
)

func TestGetPaidUnionIDByTransactionID(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"unionid": "oTmHYjg-tElZ68xxxxxxxxhy1Rgk"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://api.weixin.qq.com/wxa/getpaidunionid?access_token=ACCESS_TOKEN&openid=OPENID&transaction_id=TRANSACTION_ID", nil).Return(resp, nil)

	mp := New("APPID", "APPSECRET")
	mp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultPaidUnionID)

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", GetPaidUnionIDByTransactionID("OPENID", "TRANSACTION_ID", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultPaidUnionID{
		UnionID: "oTmHYjg-tElZ68xxxxxxxxhy1Rgk",
	}, result)
}

func TestGetPaidUnionIDByOutTradeNO(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"unionid": "oTmHYjg-tElZ68xxxxxxxxhy1Rgk"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://api.weixin.qq.com/wxa/getpaidunionid?access_token=ACCESS_TOKEN&mch_id=MCH_ID&openid=OPENID&out_trade_no=OUT_TRADE_NO", nil).Return(resp, nil)

	mp := New("APPID", "APPSECRET")
	mp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultPaidUnionID)

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", GetPaidUnionIDByOutTradeNO("OPENID", "MCH_ID", "OUT_TRADE_NO", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultPaidUnionID{
		UnionID: "oTmHYjg-tElZ68xxxxxxxxhy1Rgk",
	}, result)
}
