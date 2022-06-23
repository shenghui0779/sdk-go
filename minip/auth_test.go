package minip

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/chenghonour/gochat/mock"
	"github.com/chenghonour/gochat/wx"
)

func TestGetPhoneNumber(t *testing.T) {
	body := []byte(`{"code":"e31968a7f94cc5ee25fafc2aef2773f0bb8c3937b22520eb8ee345274d00c144"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"phone_info": {
		"phoneNumber": "xxxxxx",
		"purePhoneNumber": "xxxxxx",
		"countryCode": "86",
		"watermark": {
			"timestamp": 1637744274,
			"appid": "xxxx"
		}
	}
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	mp := New("APPID", "APPSECRET")
	mp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultPhoneNumber)

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", GetPhoneNumber("e31968a7f94cc5ee25fafc2aef2773f0bb8c3937b22520eb8ee345274d00c144", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultPhoneNumber{
		PhoneInfo: &PhoneInfo{
			PhoneNumber:     "xxxxxx",
			PurePhoneNumber: "xxxxxx",
			CountryCode:     "86",
			Watermark: Watermark{
				Timestamp: 1637744274,
				AppID:     "xxxx",
			},
		},
	}, result)
}

func TestCheckEncryptedData(t *testing.T) {
	body := []byte(`{"encrypted_msg_hash":"657edd868c9715a9bebe42b833269a557a48498785397a796f1568c29a200b2c"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"vaild": true,
	"create_time": 1629121902
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/wxa/business/checkencryptedmsg?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	mp := New("APPID", "APPSECRET")
	mp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultEncryptedDataCheck)

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", CheckEncryptedData("hsSuSUsePBqSQw2rYMtf9Nvha603xX8f2BMQBcYRoJiMNwOqt/UEhrqekebG5ar0LFNAm5MD4Uz6zorRwiXJwbySJ/FEJHav4NsobBIU1PwdjbJWVQLFy7+YFkHB32OnQXWMh6ugW7Dyk2KS5BXp1f5lniKPp1KNLyNLlFlNZ2mgJCJmWvHj5AI7BLpWwoRvqRyZvVXo+9FsWqvBdxmAPA==", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultEncryptedDataCheck{
		Valid:      true,
		CreateTime: 1629121902,
	}, result)
}

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
