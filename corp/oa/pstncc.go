package oa

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsPstnccCall struct {
	CalleeUserID []string `json:"callee_userid"`
}

type ResultPstnccCall struct {
	States []*PstnccState `json:"states"`
}

type PstnccState struct {
	Code   int    `json:"code"`
	CallID string `json:"callid"`
	UserID string `json:"userid"`
}

// CallPstncc 发起语音电话
func CallPstncc(calleeUserIDs []string, result *ResultPstnccCall) wx.Action {
	params := &ParamsPstnccCall{
		CalleeUserID: calleeUserIDs,
	}

	return wx.NewPostAction(urls.CorpOACallPstncc,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsPstnccStates struct {
	CalleeUserID string `json:"callee_userid"`
	CallID       string `json:"callid"`
}

type ResultPstnccStates struct {
	IsTalked int   `json:"istalked"`
	CallTime int64 `json:"calltime"`
	TalkTime int   `json:"talktime"`
	Reason   int   `json:"reason"`
}

// GetPstnccStates 获取接听状态
func GetPstnccStates(calleeUserID, callID string, result *ResultPstnccStates) wx.Action {
	params := &ParamsPstnccStates{
		CalleeUserID: calleeUserID,
		CallID:       callID,
	}

	return wx.NewPostAction(urls.CorpOAGetPstnccStates,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
