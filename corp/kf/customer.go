package kf

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsCustomerBatchGet struct {
	ExternalUseridList []string `json:"external_userid_list"`
}

type ResultCustomerBatchGet struct {
	CustomerList          []*Customer `json:"customer_list"`
	InvalidExternalUserID []string    `json:"invalid_external_userid"`
}

type Customer struct {
	ExternalUserID string `json:"external_userid"`
	Nickname       string `json:"nickname"`
	Avatar         string `json:"avatar"`
	Gender         int    `json:"gender"`
	UnionID        string `json:"unionid"`
}

// BatchGetCustomer 获取客户基础信息
func BatchGetCustomer(externalUserIDs []string, result *ResultCustomerBatchGet) wx.Action {
	params := &ParamsCustomerBatchGet{
		ExternalUseridList: externalUserIDs,
	}

	return wx.NewPostAction(urls.CorpKFCustomerBatchGet,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ResultServiceUpgradeConfig struct {
	MemberRange    *MemberRange    `json:"member_range"`
	GroupChatRange *GroupChatRange `json:"groupchat_range"`
}

type MemberRange struct {
	UserIDList       []string `json:"userid_list"`
	DepartmentIDList []int64  `json:"department_id_list"`
}

type GroupChatRange struct {
	ChatIDList []string `json:"chat_id_list"`
}

// GetUpgradeServiceConfig 获取配置的专员与客户群
func GetUpgradeServiceConfig(result *ResultServiceUpgradeConfig) wx.Action {
	return wx.NewGetAction(urls.CorpKFGetUpgradeServiceConfig,
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ParamsServiceUpgrade struct {
	OpenKFID       string     `json:"open_kfid"`
	ExternalUserID string     `json:"external_userid"`
	Type           int        `json:"type"` // 升级到专员服务还是客户群服务：1 - 专员服务；2 - 客户群服务
	Member         *Member    `json:"member,omitempty"`
	GroupChat      *GroupChat `json:"groupchat,omitempty"`
}

type Member struct {
	UserID  string `json:"userid"`
	Wording string `json:"wording"`
}

type GroupChat struct {
	ChatID  string `json:"chat_id"`
	Wording string `json:"wording"`
}

// UpgradeMemberService 为客户升级为专员服务
func UpgradeMemberService(openKFID, externalUserID string, member *Member) wx.Action {
	params := &ParamsServiceUpgrade{
		OpenKFID:       openKFID,
		ExternalUserID: externalUserID,
		Type:           1,
		Member:         member,
	}

	return wx.NewPostAction(urls.CorpKFUpgradeService,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

// UpgradeGroupChatService 为客户升级为客户群服务
func UpgradeGroupChatService(openKFID, externalUserID string, groupChat *GroupChat) wx.Action {
	params := &ParamsServiceUpgrade{
		OpenKFID:       openKFID,
		ExternalUserID: externalUserID,
		Type:           2,
		GroupChat:      groupChat,
	}

	return wx.NewPostAction(urls.CorpKFUpgradeService,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type ParamsServiceUpgradeCancel struct {
	OpenKFID       string `json:"open_kfid"`
	ExternalUserID string `json:"external_userid"`
}

// CancelUpgradeService 为客户取消推荐
func CancelUpgradeService(openKFID, externalUserID string) wx.Action {
	params := &ParamsServiceUpgradeCancel{
		OpenKFID:       openKFID,
		ExternalUserID: externalUserID,
	}

	return wx.NewPostAction(urls.CorpKFCancelUpgradeService,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}
