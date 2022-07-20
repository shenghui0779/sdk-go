package externalcontact

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsCorpTagList struct {
	TagID   []string `json:"tag_id,omitempty"`
	GroupID []string `json:"group_id,omitempty"`
}

type ResultCorpTagList struct {
	TagGroup []*CorpTagGroup `json:"tag_group"`
}

type CorpTagGroup struct {
	GroupID    string     `json:"group_id"`
	GroupName  string     `json:"group_name"`
	CreateTime int64      `json:"create_time"`
	Order      int        `json:"order"`
	Deleted    bool       `json:"deleted"`
	Tag        []*CorpTag `json:"tag"`
}

type CorpTag struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CreateTime int64  `json:"create_time"`
	Order      int    `json:"order"`
	Deleted    bool   `json:"deleted"`
}

// ListCorpTag 获取企业标签库
func ListCorpTag(tagIDs, groupIDs []string, result *ResultCorpTagList) wx.Action {
	params := &ParamsCorpTagList{
		TagID:   tagIDs,
		GroupID: groupIDs,
	}

	return wx.NewPostAction(urls.CorpExternalContactCorpTagList,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ParamsCorpTag struct {
	Name  string `json:"name"`
	Order int    `json:"order,omitempty"`
}

type ParamsCorpTagAdd struct {
	GroupID   string           `json:"group_id,omitempty"`
	GroupName string           `json:"group_name,omitempty"`
	Order     int              `json:"order,omitempty"`
	Tag       []*ParamsCorpTag `json:"tag"`
	AgentID   int64            `json:"agentid,omitempty"`
}

type ResultCorpTagAdd struct {
	TagGroup *CorpTagGroup `json:"tag_group"`
}

// AddCorpTag 添加企业客户标签
func AddCorpTag(params *ParamsCorpTagAdd, result *ResultCorpTagAdd) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactCorpTagAdd,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ParamsCorpTagEdit struct {
	ID      string `json:"id"`
	Name    string `json:"name,omitempty"`
	Order   int    `json:"order,omitempty"`
	AgentID int64  `json:"agentid,omitempty"`
}

// EditCorpTag 编辑企业客户标签
func EditCorpTag(params *ParamsCorpTagEdit) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactCorpTagEdit,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type ParamsCorpTagDelete struct {
	TagID   []string `json:"tag_id"`
	GroupID []string `json:"group_id"`
	AgentID int64    `json:"agentid"`
}

// DeleteCorpTag 删除企业客户标签
func DeleteCorpTag(tagIDs, groupIDs []string, agentID int64) wx.Action {
	params := &ParamsCorpTagDelete{
		TagID:   tagIDs,
		GroupID: groupIDs,
		AgentID: agentID,
	}

	return wx.NewPostAction(urls.CorpExternalContactCorpTagDelete,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type ParamsStrategyTagList struct {
	StrategyID int64    `json:"strategy_id,omitempty"`
	TagID      []string `json:"tag_id,omitempty"`
	GroupID    []string `json:"group_id,omitempty"`
}

type ResultStrategyTagList struct {
	TagGroup []*StrategyTagGroup `json:"tag_group"`
}

type StrategyTagGroup struct {
	GroupID    string         `json:"group_id"`
	GroupName  string         `json:"group_name"`
	CreateTime int64          `json:"create_time"`
	Order      uint32         `json:"order"`
	StrategyID int64          `json:"strategy_id"`
	Tag        []*StrategyTag `json:"tag"`
}

type StrategyTag struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CreateTime int64  `json:"create_time"`
	Order      uint32 `json:"order"`
}

// ListStrategyTag 获取指定规则组下的企业客户标签
func ListStrategyTag(strategyID int64, tagIDs, groupIDs []string, result *ResultStrategyTagList) wx.Action {
	params := &ParamsStrategyTagList{
		StrategyID: strategyID,
		TagID:      tagIDs,
		GroupID:    groupIDs,
	}

	return wx.NewPostAction(urls.CorpExternalContactStrategyTagList,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ParamsStrategyTagAdd struct {
	StrategyID int64                `json:"strategy_id,omitempty"`
	GroupID    string               `json:"group_id"`
	GroupName  string               `json:"group_name"`
	Order      uint32               `json:"order"`
	Tag        []*ParamsStrategyTag `json:"tag"`
}

type ParamsStrategyTag struct {
	Name  string `json:"name"`
	Order uint32 `json:"order,omitempty"`
}

type ResultStrategyTagAdd struct {
	TagGroup *StrategyTagGroup `json:"tag_group"`
}

// AddStrategyTag 为指定规则组创建企业客户标签
func AddStrategyTag(params *ParamsStrategyTagAdd, result *ResultStrategyTagAdd) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactStrategyTagAdd,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ParamsStrategyTagEdit struct {
	ID    string `json:"id"`
	Name  string `json:"name,omitempty"`
	Order uint32 `json:"order,omitempty"`
}

// EditStrategyTag 编辑指定规则组下的企业客户标签
func EditStrategyTag(params *ParamsStrategyTagEdit) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactStrategyTagEdit,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type ParamsStrategyTagDelete struct {
	TagID   []string `json:"tag_id,omitempty"`
	GroupID []string `json:"group_id,omitempty"`
}

// DeleteStrategyTag 删除指定规则组下的企业客户标签
func DeleteStrategyTag(tagIDs, groupIDs []string) wx.Action {
	params := &ParamsStrategyTagDelete{
		TagID:   tagIDs,
		GroupID: groupIDs,
	}

	return wx.NewPostAction(urls.CorpExternalContactStrategyTagDelete,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type ParamsTagMark struct {
	UserID         string   `json:"userid"`
	ExternalUserID string   `json:"external_userid"`
	AddTag         []string `json:"add_tag,omitempty"`
	RemoveTag      []string `json:"remove_tag,omitempty"`
}

// MarkTag 编辑客户企业标签
func MarkTag(params *ParamsTagMark) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactMarkTag,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}
