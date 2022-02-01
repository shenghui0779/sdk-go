package payment

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsMerchantAdd struct {
	MchID        string `json:"mch_id"`
	MerchantName string `json:"merchant_name"`
}

func AddMerchant(mchID, mchName string) wx.Action {
	params := &ParamsMerchantAdd{
		MchID:        mchID,
		MerchantName: mchName,
	}

	return wx.NewPostAction(urls.CorpPaymentMerchantAdd,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type AllowUseScope struct {
	User    []string `json:"user,omitempty"`
	PartyID []int64  `json:"partyid,omitempty"`
	TagID   []int64  `json:"tagid,omitempty"`
}

type ParamsMerchantGet struct {
	MchID string `json:"mch_id"`
}

type ResultMerchantGet struct {
	BindStatus    int            `json:"bind_status"`
	MchID         string         `json:"mch_id"`
	MerchantName  string         `json:"merchant_name"`
	AllowUseScope *AllowUseScope `json:"allow_use_scope"`
}

func GetMerchant(mchID string, result *ResultMerchantGet) wx.Action {
	params := &ParamsMerchantGet{
		MchID: mchID,
	}

	return wx.NewPostAction(urls.CorpPaymentMerchantGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsMerchantDelete struct {
	MchID string `json:"mch_id"`
}

func DeleteMerchant(mchID string) wx.Action {
	params := &ParamsMerchantDelete{
		MchID: mchID,
	}

	return wx.NewPostAction(urls.CorpPaymentMerchantDelete,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsMchUseScopeSet struct {
	MchID         string         `json:"mch_id"`
	AllowUseScope *AllowUseScope `json:"allow_use_scope"`
}

func SetMchUseScope(mchID string, scope *AllowUseScope) wx.Action {
	params := &ParamsMchUseScopeSet{
		MchID:         mchID,
		AllowUseScope: scope,
	}

	return wx.NewPostAction(urls.CorpPaymentMchUseScopeSet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type BillInfo struct {
	TransactionID  string           `json:"transaction_id"`
	TradeState     int              `json:"trade_state"`
	PayTime        int64            `json:"pay_time"`
	OutTradeNO     string           `json:"out_trade_no"`
	ExternalUserID string           `json:"external_userid"`
	TotalFee       int              `json:"total_fee"`
	PayeeUserID    string           `json:"payee_userid"`
	PaymentType    int              `json:"payment_type"`
	MchID          string           `json:"mch_id"`
	Remark         string           `json:"remark"`
	TotalRefundFee int              `json:"total_refund_fee"`
	CommodityList  []*CommodityInfo `json:"commodity_list"`
	RefundList     []*RefundInfo    `json:"refund_list"`
	PayerInfo      *PayerInfo       `json:"payer_info"`
}

type CommodityInfo struct {
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}

type RefundInfo struct {
	OutRefundNO   string `json:"out_refund_no"`
	RefundUserID  string `json:"refund_userid"`
	RefundComment string `json:"refund_comment"`
	RefundReqTime int64  `json:"refund_reqtime"`
	RefundStatus  int    `json:"refund_status"`
	RefundFee     int    `json:"refund_fee"`
}

type PayerInfo struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type ParamsBillList struct {
	BeginTime   int64  `json:"begin_time"`
	EndTime     int64  `json:"end_time"`
	PayeeUserID string `json:"payee_userid,omitempty"`
	Cursor      string `json:"cursor,omitempty"`
	Limit       int    `json:"limit,omitempty"`
}

type ResultBillList struct {
	NextCursor string      `json:"next_cursor"`
	BillList   []*BillInfo `json:"bill_list"`
}

func GetBillList(params *ParamsBillList, result *ResultBillList) wx.Action {
	return wx.NewPostAction(urls.CorpPaymentBillListGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
