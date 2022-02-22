package agent

import (
	"encoding/json"
	"strconv"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ResultAgentGet struct {
	AgentID                 int64          `json:"agentid"`
	Name                    string         `json:"name"`
	SquareLogoURL           string         `json:"square_logo_url"`
	Description             string         `json:"description"`
	Close                   int            `json:"close"`
	RedirectDomain          string         `json:"redirect_domain"`
	ReportLocationFlag      int            `json:"report_location_flag"`
	ISReportenter           int            `json:"is_reportenter"`
	HomeURL                 string         `json:"home_url"`
	CustomizedPublishStatus int            `json:"customized_publish_status"`
	AllowUserInfos          AllowUserInfos `json:"allow_userinfos"`
	AllowPartys             AllowPartys    `json:"allow_partys"`
	AllowTags               AllowTags      `json:"allow_tags"`
}

type AllowUserInfos struct {
	User []AllowUser `json:"user"`
}

type AllowUser struct {
	UserID string `json:"userid"`
}

type AllowPartys struct {
	PartyID []int64 `json:"partyid"`
}

type AllowTags struct {
	TagID []int64 `json:"tagid"`
}

// GetAgent 获取指定的应用详情
func GetAgent(agentID int64, result *ResultAgentGet) wx.Action {
	return wx.NewGetAction(urls.CorpAgentGet,
		wx.WithQuery("agentid", strconv.FormatInt(agentID, 10)),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ResultAgentList struct {
	AgentList []*AgentListData `json:"agentlist"`
}

type AgentListData struct {
	AgentID       int64  `json:"agentid"`
	Name          string `json:"name"`
	SquareLogoURL string `json:"square_logo_url"`
}

// ListAgent 获取access_token对应的应用列表
func ListAgent(result *ResultAgentList) wx.Action {
	return wx.NewGetAction(urls.CorpAgentList,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsAgentSet struct {
	AgentID            int64  `json:"agentid"`
	ReportLocationFlag int    `json:"report_location_flag"`
	LogoMediaID        string `json:"logo_mediaid"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	RedirectDomain     string `json:"redirect_domain"`
	ISReportEnter      int    `json:"isreportenter"`
	HomeURL            string `json:"home_url"`
}

// SetAgent 设置应用
func SetAgent(params *ParamsAgentSet) wx.Action {
	return wx.NewPostAction(urls.CorpAgentSet,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}
