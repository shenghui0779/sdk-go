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
	Fee                   int                   `json:"fee"`
	Title                 string                `json:"title"`
	BillingTime           int64                 `json:"billing_time"`
	BillingNO             string                `json:"billing_no"`
	BillingCode           string                `json:"billing_code"`
	Tax                   int                   `json:"tax"`
	FeeWithoutTax         int                   `json:"fee_without_tax"`
	Detail                string                `json:"detail"`
	PdfURL                string                `json:"pdf_url"`
	TriPdfUrl             string                `json:"tri_pdf_url"`
	CheckCode             string                `json:"check_code"`
	BuyerNumber           string                `json:"buyer_number"`
	BuyerAddressAndPhone  string                `json:"buyer_address_and_phone"`
	BuyerBankAccount      string                `json:"buyer_bank_account"`
	SellerNumber          string                `json:"seller_number"`
	SellerAddressAndPhone string                `json:"seller_address_and_phone"`
	SellerBankAccount     string                `json:"seller_bank_account"`
	Remarks               string                `json:"remarks"`
	Cashier               string                `json:"cashier"`
	Maker                 string                `json:"maker"`
	ReimburseStatus       ReimburseStatus       `json:"reimburse_status"`
	OrderID               string                `json:"order_id"`
	Info                  []*InvoiceProductInfo `json:"info"`
}

type InvoiceProductInfo struct {
	Name  string `json:"name"`
	Num   int    `json:"num"`
	Unit  string `json:"unit"`
	Fee   int    `json:"fee"`
	Price int    `json:"price"`
}

type ParamsInvoiceInfo struct {
	CardID      string `json:"card_id"`
	EncryptCode string `json:"encrypt_code"`
}

type ResultInvoiceInfo struct {
	CardID    string           `json:"card_id"`
	BeginTime int64            `json:"begin_time"`
	EndTime   int64            `json:"end_time"`
	OpenID    string           `json:"openid"`
	Type      string           `json:"type"`
	Payee     string           `json:"payee"`
	Detail    string           `json:"detail"`
	UserInfo  *InvoiceUserInfo `json:"user_info"`
}

func GetInvoiceInfo(params *ParamsInvoiceInfo, result *ResultInvoiceInfo) wx.Action {
	return wx.NewPostAction(urls.CorpInvoiceGetInfo,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsInvoiceStatusUpdate struct {
	CardID          string          `json:"card_id"`
	EncryptCode     string          `json:"encrypt_code"`
	ReimburseStatus ReimburseStatus `json:"reimburse_status"`
}

func UpdateInvoiceStatus(params *ParamsInvoiceStatusUpdate) wx.Action {
	return wx.NewPostAction(urls.CorpInvoiceUpdateStatus,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsInvoiceBatchInfo struct {
	ItemList []*ParamsInvoiceInfo `json:"item_list"`
}

type ResultInvoiceBatchInfo struct {
	ItemList []*ResultInvoiceInfo `json:"item_list"`
}

func BatchGetInvoiceInfo(params *ParamsInvoiceBatchInfo, result *ResultInvoiceBatchInfo) wx.Action {
	return wx.NewPostAction(urls.CorpInvoiceBatchGetInfo,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsInvoiceStatusBatchUpdate struct {
	OpenID          string               `json:"openid"`
	ReimburseStatus ReimburseStatus      `json:"reimburse_status"`
	InvoiceList     []*ParamsInvoiceInfo `json:"invoice_list"`
}

func BatchUpdateInvoiceStatus(params *ParamsInvoiceStatusBatchUpdate) wx.Action {
	return wx.NewPostAction(urls.CorpInvoiceBatchUpdateStatus,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}
