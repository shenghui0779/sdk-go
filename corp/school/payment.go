package school

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type PaymentInfo struct {
	StudentUserID     string `json:"student_userid"`
	TradeState        int    `json:"trade_state"`
	TradeNO           string `json:"trade_no"`
	PayerParentUserID string `json:"payer_parent_userid"`
}

type ParamsPaymentGet struct {
	PaymentID string `json:"payment_id"`
}

type ResultPaymentGet struct {
	ProjectName   string         `json:"project_name"`
	Amount        int            `json:"amount"`
	PaymentResult []*PaymentInfo `json:"payment_result"`
}

func GetPaymentResult(paymentID string, result *ResultPaymentGet) wx.Action {
	params := &ParamsPaymentGet{
		PaymentID: paymentID,
	}

	return wx.NewPostAction(urls.CorpSchoolGetPaymentResult,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsTradeGet struct {
	PaymentID string `json:"payment_id"`
	TradeNO   string `json:"trade_no"`
}

type ResultTradeGet struct {
	TransactionID string `json:"transaction_id"`
	PayTime       int64  `json:"pay_time"`
}

func GetTrade(params *ParamsTradeGet, result *ResultTradeGet) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolGetTrade,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
