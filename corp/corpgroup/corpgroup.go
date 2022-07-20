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
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
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
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
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
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ResultChainList struct {
	Chains []*Chain `json:"chains"`
}

type Chain struct {
	ChainID   string `json:"chain_id"`
	ChainName string `json:"chain_name"`
}

// ListChain 获取上下游列表
func ListChain(result *ResultChainList) wx.Action {
	return wx.NewGetAction(urls.CorpGroupGetChainList,
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ParamsChainGroup struct {
	ChainID string `json:"chain_id"`
}

type ResultChainGroup struct {
	Groups []*ChainGroup `json:"groups"`
}

type ChainGroup struct {
	GroupID   int64  `json:"groupid"`
	GroupName string `json:"group_name"`
	ParentID  int64  `json:"parentid"`
	Order     int    `json:"order"`
}

// GetChainGroup 获取上下游通讯录分组
func GetChainGroup(chainID string, result *ResultChainGroup) wx.Action {
	params := &ParamsChainGroup{
		ChainID: chainID,
	}

	return wx.NewPostAction(urls.CorpGroupGetChainGroup,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ParamsChainCorpInfoList struct {
	ChainID    string `json:"chain_id"`
	GroupID    int64  `json:"groupid,omitempty"`
	FetchChild int    `json:"fetch_child,omitempty"`
}

type ResultChainCorpInfoList struct {
	GroupCorps []*ChainGroupCorp `json:"group_corps"`
}

type ChainGroupCorp struct {
	GroupID  int64  `json:"groupid"`
	CorpID   string `json:"corpid"`
	CorpName string `json:"corp_name"`
	CustomID string `json:"custom_id"`
}

// ListChainCorpInfo 获取企业上下游通讯录分组下的企业详情列表
func ListChainCorpInfo(chainID string, groupID int64, fetchChild int, result *ResultChainCorpInfoList) wx.Action {
	params := &ParamsChainCorpInfoList{
		ChainID:    chainID,
		GroupID:    groupID,
		FetchChild: fetchChild,
	}

	return wx.NewPostAction(urls.CorpGroupGetChainCorpInfoList,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ParamsUnionIDToExternalUserID struct {
	UnionID string `json:"unionid"`
	OpenID  string `json:"openid"`
	CorpID  string `json:"corpid,omitempty"`
}

type ResultUnionIDToExternalUserID struct {
	ExternalUserIDInfo []*ExternalUserIDInfo `json:"external_userid_info"`
}

type ExternalUserIDInfo struct {
	CorpID         string `json:"corpid"`
	ExternalUserID string `json:"external_userid"`
}

// UnionIDToExternalUserID 上下游企业应用获取微信用户的external_userid
func UnionIDToExternalUserID(unionid, openid, corpid string, result *ResultUnionIDToExternalUserID) wx.Action {
	params := &ParamsUnionIDToExternalUserID{
		UnionID: unionid,
		OpenID:  openid,
		CorpID:  corpid,
	}

	return wx.NewPostAction(urls.CorpGroupUnionIDToExternalUserID,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}
