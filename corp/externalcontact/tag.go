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
	Order      uint32     `json:"order"`
	Deleted    bool       `json:"deleted"`
	Tag        []*CorpTag `json:"tag"`
}

type CorpTag struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CreateTime int64  `json:"create_time"`
	Order      uint32 `json:"order"`
	Deleted    bool   `json:"deleted"`
}

func ListCorpTag(params *ParamsCorpTagList, result *ResultCorpTagList) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactCorpTagList,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsCorpTag struct {
	Name  string `json:"name"`
	Order uint32 `json:"order,omitempty"`
}

type ParamsCorpTagAdd struct {
	GroupID   string           `json:"group_id,omitempty"`
	GroupName string           `json:"group_name,omitempty"`
	Order     uint32           `json:"order,omitempty"`
	Tag       []*ParamsCorpTag `json:"tag"`
	AgentID   int64            `json:"agentid,omitempty"`
}

type ResultCorpTagAdd struct {
	TagGroup *CorpTagGroup `json:"tag_group"`
}

func AddCorpTag(params *ParamsCorpTagAdd, result *ResultCorpTagAdd) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactCorpTagAdd,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsCorpTagEdit struct {
	ID      string `json:"id"`
	Name    string `json:"name,omitempty"`
	Order   uint32 `json:"order,omitempty"`
	AgentID int64  `json:"agentid,omitempty"`
}

func EditCorpTag(params *ParamsCorpTagEdit) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactCorpTagEdit,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsCorpTagDelete struct {
	TagID   []string `json:"tag_id"`
	GroupID []string `json:"group_id"`
	AgentID int64    `json:"agentid"`
}

func DeleteCorpTag(params *ParamsCorpTagDelete) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactCorpTagDelete,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
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

func ListStrategyTag(params *ParamsStrategyTagList, result *ResultStrategyTagList) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactStrategyTagList,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
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

func AddStrategyTag(params *ParamsStrategyTagAdd, result *ResultStrategyTagAdd) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactStrategyTagAdd,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsStrategyTagEdit struct {
	ID    string `json:"id"`
	Name  string `json:"name,omitempty"`
	Order uint32 `json:"order,omitempty"`
}

func EditStrategyTag(params *ParamsStrategyTagEdit) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactStrategyTagEdit,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsStrategyTagDelete struct {
	TagID   []string `json:"tag_id,omitempty"`
	GroupID []string `json:"group_id,omitempty"`
}

func DeleteStrategyTag(params *ParamsStrategyTagDelete) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactStrategyTagDelete,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsTagMark struct {
	UserID         string   `json:"userid"`
	ExternalUserID string   `json:"external_userid"`
	AddTag         []string `json:"add_tag,omitempty"`
	RemoveTag      []string `json:"remove_tag,omitempty"`
}

func MarkTag(params *ParamsTagMark) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactMarkTag,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}
