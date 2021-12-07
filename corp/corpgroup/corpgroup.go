package corpgroup

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type CorpInfo struct {
	CorpID   string `json:"corpid"`
	CorpName string `json:"corp_name"`
	AgentID  int64  `json:"agentid"`
}

type ParamsAppShareInfoList struct {
	AgentID int64 `json:"agentid"`
}

type ResultAppShareInfoList struct {
	CorpList []*CorpInfo `json:"corp_list"`
}

func ListAppShareInfo(params *ParamsAppShareInfoList, result *ResultAppShareInfoList) wx.Action {
	return wx.NewPostAction(urls.CorpGroupListAppShareInfo,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsTokenGet struct {
	CorpID  string `json:"corpid"`
	AgentID int64  `json:"agentid"`
}

type ResultTokenGet struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func GetToken(params *ParamsTokenGet, result *ResultTokenGet) wx.Action {
	return wx.NewPostAction(urls.CorpGroupGetToken,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsMinipSessionTransfer struct {
	UserID     string `json:"userid"`
	SessionKey string `json:"session_key"`
}

type ResultMinipSessionTransfer struct {
	UserID     string `json:"userid"`
	SessionKey string `json:"session_key"`
}

func TransferMinipSession(params *ParamsMinipSessionTransfer, result *ResultMinipSessionTransfer) wx.Action {
	return wx.NewPostAction(urls.CorpGroupMinipTransferSession,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
