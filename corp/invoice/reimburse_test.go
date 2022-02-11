package invoice

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/corp"
	"github.com/shenghui0779/gochat/mock"
	"github.com/shenghui0779/gochat/wx"
)

func TestGetInvoiceInfo(t *testing.T) {
	body := []byte(`{"card_id":"CARDID","encrypt_code":"ENCRYPTCODE"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "card_id": "CARDID",
    "begin_time": 1469084420,
    "end_time": 2100236420,
    "openid": "oxfdsfsafaf",
    "type": "增值税电子普通发票",
    "payee": "深圳市XX有限公司",
    "detail": "可在公司企业微信内报销使用",
    "user_info": {
        "fee": 4000,
        "title": "XX有限公司",
        "billing_time": 1478620800,
        "billing_no": "00000001",
        "billing_code": "000000000001",
        "info": [
            {
                "name": "NAME",
                "num": 10,
                "unit": "吨",
                "fee": 10,
                "price": 4
            }
        ],
        "fee_without_tax": 2345,
        "tax": 123,
        "detail": "项目",
        "pdf_url": "pdf_url",
        "reimburse_status": "INVOICE_REIMBURSE_INIT",
        "check_code": "check_code"
    }
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/card/invoice/reimburse/getinvoiceinfo?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultInvoiceInfo)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetInvoiceInfo("CARDID", "ENCRYPTCODE", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultInvoiceInfo{
		CardID:    "CARDID",
		BeginTime: 1469084420,
		EndTime:   2100236420,
		OpenID:    "oxfdsfsafaf",
		Type:      "增值税电子普通发票",
		Payee:     "深圳市XX有限公司",
		Detail:    "可在公司企业微信内报销使用",
		UserInfo: &InvoiceUserInfo{
			Fee:             4000,
			Title:           "XX有限公司",
			BillingTime:     1478620800,
			BillingNO:       "00000001",
			BillingCode:     "000000000001",
			FeeWithoutTax:   2345,
			Tax:             123,
			Detail:          "项目",
			PdfURL:          "pdf_url",
			ReimburseStatus: "INVOICE_REIMBURSE_INIT",
			CheckCode:       "check_code",
			Info: []*InvoiceProductInfo{
				{
					Name:  "NAME",
					Num:   10,
					Unit:  "吨",
					Fee:   10,
					Price: 4,
				},
			},
		},
	}, result)
}

func TestUpdateInvoiceStatus(t *testing.T) {
	body := []byte(`{"card_id":"CARDID","encrypt_code":"ENCRYPTCODE","reimburse_status":"INVOICE_REIMBURSE_INIT"}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/card/invoice/reimburse/updateinvoicestatus?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UpdateInvoiceStatus("CARDID", "ENCRYPTCODE", InvioceReimburseInit))

	assert.Nil(t, err)
}

func TestBatchGetInvoiceInfo(t *testing.T) {
	body := []byte(`{"item_list":[{"card_id":"CARDID1","encrypt_code":"ENCRYPTCODE1"},{"card_id":"CARDID2","encrypt_code":"ENCRYPTCODE2"}]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "item_list": [
        {
            "card_id": "CARDID",
			"begin_time": 1469084420,
			"end_time": 2100236420,
            "openid": "ofdsfasfsafsafasf",
            "type": "增值税电子普通发票",
            "payee": "深圳市XX有限公司",
            "detail": "可在公司企业微信内报销使用",
            "user_info": {
                "fee": 4000,
                "title": "XX有限公司",
                "billing_time": 1478620800,
                "billing_no": "00000001",
                "billing_code": "000000000001",
                "info": [
                    {
                        "name": "NAME",
                        "num": 10,
                        "unit": "吨",
                        "fee": 10,
                        "price": 4
                    }
                ],
                "fee_without_tax": 2345,
                "tax": 123,
                "detail": "项目",
                "pdf_url": "pdf_url",
                "reimburse_status": "INVOICE_REIMBURSE_INIT",
                "order_id": "12345678",
                "check_code": "CHECKCODE",
                "buyer_number": "BUYERNUMER"
            }
        }
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/card/invoice/reimburse/getinvoiceinfobatch?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	invoices := []*ParamsInvoice{
		{
			CardID:      "CARDID1",
			EncryptCode: "ENCRYPTCODE1",
		},
		{
			CardID:      "CARDID2",
			EncryptCode: "ENCRYPTCODE2",
		},
	}

	result := new(ResultInvoiceBatchInfo)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", BatchGetInvoiceInfo(invoices, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultInvoiceBatchInfo{
		ItemList: []*ResultInvoiceInfo{
			{
				CardID:    "CARDID",
				BeginTime: 1469084420,
				EndTime:   2100236420,
				OpenID:    "ofdsfasfsafsafasf",
				Type:      "增值税电子普通发票",
				Payee:     "深圳市XX有限公司",
				Detail:    "可在公司企业微信内报销使用",
				UserInfo: &InvoiceUserInfo{
					Fee:             4000,
					Title:           "XX有限公司",
					BillingTime:     1478620800,
					BillingNO:       "00000001",
					BillingCode:     "000000000001",
					FeeWithoutTax:   2345,
					Tax:             123,
					Detail:          "项目",
					PdfURL:          "pdf_url",
					ReimburseStatus: "INVOICE_REIMBURSE_INIT",
					CheckCode:       "CHECKCODE",
					BuyerNumber:     "BUYERNUMER",
					OrderID:         "12345678",
					Info: []*InvoiceProductInfo{
						{
							Name:  "NAME",
							Num:   10,
							Unit:  "吨",
							Fee:   10,
							Price: 4,
						},
					},
				},
			},
		},
	}, result)
}

func TestBatchUpdateInvoiceStatus(t *testing.T) {
	body := []byte(`{"openid":"OPENID","reimburse_status":"INVOICE_REIMBURSE_INIT","invoice_list":[{"card_id":"cardid_1","encrypt_code":"encrypt_code_1"},{"card_id":"cardid_2","encrypt_code":"encrypt_code_2"}]}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/card/invoice/reimburse/updatestatusbatch?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	invoices := []*ParamsInvoice{
		{
			CardID:      "cardid_1",
			EncryptCode: "encrypt_code_1",
		},
		{
			CardID:      "cardid_2",
			EncryptCode: "encrypt_code_2",
		},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", BatchUpdateInvoiceStatus("OPENID", "INVOICE_REIMBURSE_INIT", invoices...))

	assert.Nil(t, err)
}
