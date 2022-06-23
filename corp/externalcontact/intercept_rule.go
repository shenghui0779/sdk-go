package externalcontact

import (
	"encoding/json"

	"github.com/chenghonour/gochat/urls"
	"github.com/chenghonour/gochat/wx"
)

type ExtraRule struct {
	SemanticsList []int `json:"semantics_list,omitempty"`
}

type RuleApplicableRange struct {
	UserList       []string `json:"user_list,omitempty"`
	DepartmentList []int64  `json:"department_list,omitempty"`
}

type InterceptRule struct {
	RuleID          string               `json:"rule_id"`
	RuleName        string               `json:"rule_name"`
	WordList        []string             `json:"word_list"`
	ExtraRule       *ExtraRule           `json:"extra_rule"`
	InterceptType   int                  `json:"intercept_type"`
	ApplicableRange *RuleApplicableRange `json:"applicable_range"`
}

type ParamsInterceptRuleAdd struct {
	RuleName        string               `json:"rule_name"`
	WordList        []string             `json:"word_list"`
	SemanticsList   []int                `json:"semantics_list,omitempty"`
	InterceptType   int                  `json:"intercept_type"`
	ApplicableRange *RuleApplicableRange `json:"applicable_range"`
}

type ResultInterceptRuleAdd struct {
	RuleID string `json:"rule_id"`
}

// AddInterceptRule 新建敏感词规则
func AddInterceptRule(params *ParamsInterceptRuleAdd, result *ResultInterceptRuleAdd) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactInterceptRuleAdd,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsInterceptRuleUpdate struct {
	RuleID                string               `json:"rule_id"`
	RuleName              string               `json:"rule_name,omitempty"`
	WordList              []string             `json:"word_list,omitempty"`
	ExtraRule             *ExtraRule           `json:"extra_rule,omitempty"`
	InterceptType         int                  `json:"intercept_type,omitempty"`
	AddApplicableRange    *RuleApplicableRange `json:"add_applicable_range,omitempty"`
	RemoveApplicableRange *RuleApplicableRange `json:"remove_applicable_range,omitempty"`
}

// UpdateInterceptRule 修改敏感词规则
func UpdateInterceptRule(params *ParamsInterceptRuleUpdate) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactInterceptRuleUpdate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type RuleListData struct {
	RuleID     string `json:"rule_id"`
	RuleName   string `json:"rule_name"`
	CreateTime int64  `json:"create_time"`
}

type ResultInterceptRuleList struct {
	RuleList []*RuleListData `json:"rule_list"`
}

// ListInterceptRule 获取敏感词规则列表
func ListInterceptRule(result *ResultInterceptRuleList) wx.Action {
	return wx.NewGetAction(urls.CorpExternalContactInterceptRuleList,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsInterceptRuleGet struct {
	RuleID string `json:"rule_id"`
}

type ResultInterceptRuleGet struct {
	Rule *InterceptRule `json:"rule"`
}

// GetInterceptRule 获取敏感词规则详情
func GetInterceptRule(ruleID string, result *ResultInterceptRuleGet) wx.Action {
	params := &ParamsInterceptRuleGet{
		RuleID: ruleID,
	}

	return wx.NewPostAction(urls.CorpExternalContactInterceptRuleGet,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsInterceptRuleDelete struct {
	RuleID string `json:"rule_id"`
}

// DeleteInterceptRule 删除敏感词规则
func DeleteInterceptRule(ruleID string) wx.Action {
	params := &ParamsInterceptRuleDelete{
		RuleID: ruleID,
	}

	return wx.NewPostAction(urls.CorpExternalContactInterceptRuleDelete,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}
