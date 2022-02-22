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
	AgentID      int64  `json:"agentid"`
	BusinessType int    `json:"business_type,omitempty"`
	CorpID       string `json:"corpid,omitempty"`
	Limit        int    `json:"limit,omitempty"`
	Cursor       string `json:"cursor,omitempty"`
}

type ResultAppShareInfoList struct {
	Ending     int         `json:"ending"`
	CorpList   []*CorpInfo `json:"corp_list"`
	NextCursor string      `json:"next_cursor"`
}

// ListAppShareInfo 获取应用共享信息
func ListAppShareInfo(params *ParamsAppShareInfoList, result *ResultAppShareInfoList) wx.Action {
	return wx.NewPostAction(urls.CorpGroupListAppShareInfo,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsCorpAccessToken struct {
	CorpID       string `json:"corpid"`
	BusinessType int    `json:"business_type,omitempty"`
	AgentID      int64  `json:"agentid"`
}

type ResultCorpAccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// GetCorpAccessToken 获取下级/下游企业的access_token
func GetCorpAccessToken(params *ParamsCorpAccessToken, result *ResultCorpAccessToken) wx.Action {
	return wx.NewPostAction(urls.CorpGroupGetAccessToken,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
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

// TransferMinipSession 获取下级/下游企业小程序session
func TransferMinipSession(params *ParamsMinipSessionTransfer, result *ResultMinipSessionTransfer) wx.Action {
	return wx.NewPostAction(urls.CorpGroupMinipTransferSession,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
