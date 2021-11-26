package invoice

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

// ReimburseStatus 发票报销状态
type ReimburseStatus string

const (
	InvioceReimburseInit    ReimburseStatus = "INVOICE_REIMBURSE_INIT"    // 发票初始状态，未锁定
	InvioceReimburseLock    ReimburseStatus = "INVOICE_REIMBURSE_LOCK"    // 发票已锁定，无法重复提交报销
	InvioceReimburseClosure ReimburseStatus = "INVOICE_REIMBURSE_CLOSURE" // 发票已核销，从用户卡包中移除
)

type InvoiceUserInfo struct {
	Fee             int                   `json:"fee"`
	Title           string                `json:"title"`
	BillingTime     int64                 `json:"billing_time"`
	BillingNO       string                `json:"billing_no"`
	BillingCode     string                `json:"billing_code"`
	FeeWithoutTax   int                   `json:"fee_without_tax"`
	Tax             int                   `json:"tax"`
	Detail          string                `json:"detail"`
	PdfURL          string                `json:"pdf_url"`
	ReimburseStatus ReimburseStatus       `json:"reimburse_status"`
	CheckCode       string                `json:"check_code"`
	Info            []*InvoiceProductInfo `json:"info"`
}

type InvoiceProductInfo struct {
	Name  string `json:"name"`
	Num   int    `json:"num"`
	Unit  string `json:"unit"`
	Fee   int    `json:"fee"`
	Price int    `json:"price"`
}

type ParamsInvoiceInfoGet struct {
	CardID      string `json:"card_id"`
	EncryptCode string `json:"encrypt_code"`
}

type ResultInvoiceInfoGet struct {
	CardID    string           `json:"card_id"`
	BeginTime int64            `json:"begin_time"`
	EndTime   int64            `json:"end_time"`
	OpenID    string           `json:"open_id"`
	Type      string           `json:"type"`
	Payee     string           `json:"payee"`
	Detail    string           `json:"detail"`
	UserInfo  *InvoiceUserInfo `json:"user_info"`
}

func GetInvoiceInfo(params *ParamsInvoiceInfoGet, result *ResultInvoiceInfoGet) wx.Action {
	return wx.NewPostAction(urls.CorpInvoiceGetInfo,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsInvoiceReimburseStatusUpdate struct {
	CardID          string          `json:"card_id"`
	EncryptCode     string          `json:"encrypt_code"`
	ReimburseStatus ReimburseStatus `json:"reimburse_status"`
}

func UpdateInvoiceReimburseStatus(params *ParamsInvoiceReimburseStatusUpdate) wx.Action {
	return wx.NewPostAction(urls.CorpInvoiceUpdateReimburseStatus,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}
