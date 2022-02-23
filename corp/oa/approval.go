package oa

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type TemplateContent struct {
	Controls []*TemplateControl `json:"controls"`
}

type TemplateControl struct {
	Property *ControlProperty `json:"property"`
	Config   *ControlConfig   `json:"config"`
}

type ParamsTemplateDetail struct {
	TemplateID string `json:"template_id"`
}

type ResultTemplateDetail struct {
	TemplateNames   []*DisplayText   `json:"template_names"`
	TemplateContent *TemplateContent `json:"template_content"`
}

// GetTemplateDetail 获取审批模板详情
func GetTemplateDetail(templateID string, result *ResultTemplateDetail) wx.Action {
	params := &ParamsTemplateDetail{
		TemplateID: templateID,
	}

	return wx.NewPostAction(urls.CorpOAGetTemplateDetail,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type Approver struct {
	Attr   int      `json:"attr"`
	UserID []string `json:"userid"`
}

type ApplySummaryInfo struct {
	SummaryInfo []*DisplayText `json:"summary_info"`
}

type ParamsApplyEvent struct {
	CreatorUserID       string              `json:"creator_userid"`
	TemplateID          string              `json:"template_id"`
	UseTemplateApprover int                 `json:"use_template_approver"`
	ChooseDepartment    int64               `json:"choose_department,omitempty"`
	Approver            []*Approver         `json:"approver"`
	Notifyer            []string            `json:"notifyer"`
	NotifyType          int                 `json:"notify_type"`
	ApplyData           *ApplyData          `json:"apply_data"`
	SummaryList         []*ApplySummaryInfo `json:"summary_list"`
}

type ResultApplyEvent struct {
	SPNO string `json:"sp_no"`
}

// ApplyEvent 提交审批申请
func ApplyEvent(params *ParamsApplyEvent, result *ResultApplyEvent) wx.Action {
	return wx.NewPostAction(urls.CorpOAApplyEvent,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsApprovalInfo struct {
	StartTime string      `json:"starttime"`
	EndTime   string      `json:"endtime"`
	Cursor    int         `json:"cursor"`
	Size      int         `json:"size"`
	Filters   []*KeyValue `json:"filters,omitempty"`
}

type ResultApprovalInfo struct {
	SPNOList []string `json:"sp_no_list"`
}

// GetApprovalInfo 批量获取审批单号
func GetApprovalInfo(params *ParamsApprovalInfo, result *ResultApprovalInfo) wx.Action {
	return wx.NewPostAction(urls.CorpOAGetApprovalInfo,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsApprovalDetail struct {
	SPNO string `json:"sp_no"`
}

type ResultApprovalDetail struct {
	Info *ApprovalDetail `json:"info"`
}

type ApprovalDetail struct {
	SPNO       string              `json:"sp_no"`
	SPName     string              `json:"sp_name"`
	SPStatus   int                 `json:"sp_status"`
	TemplateID string              `json:"template_id"`
	ApplyTime  int64               `json:"apply_time"`
	Applyer    *Applyer            `json:"applyer"`
	SPRecord   []*ApprovalSPRecord `json:"sp_record"`
	Notifyer   []*OAUser           `json:"notifyer"`
	ApplyData  *ApplyData          `json:"apply_data"`
	Comments   []*ApprovalComment  `json:"comments"`
}

type Applyer struct {
	UserID  string `json:"userid"`
	PartyID string `json:"partyid"`
}

type ApprovalSPRecord struct {
	SPStatus     int                 `json:"sp_status"`
	ApproverAttr int                 `json:"approverattr"`
	Details      []*ApprovalSPDetail `json:"details"`
}

type ApprovalSPDetail struct {
	Approver *OAUser  `json:"approver"`
	Speech   string   `json:"speech"`
	SPStatus int      `json:"sp_status"`
	SPTime   int64    `json:"sptime"`
	MediaID  []string `json:"media_id"`
}

type ApprovalComment struct {
	CommentUserInfo *OAUser  `json:"commentUserInfo"`
	CommentTime     int64    `json:"commenttime"`
	CommentContent  string   `json:"commentcontent"`
	CommentID       string   `json:"commentid"`
	MediaID         []string `json:"media_id"`
}

// GetApprovalDetail 获取审批申请详情
func GetApprovalDetail(spno string, result *ResultApprovalDetail) wx.Action {
	params := &ParamsApprovalDetail{
		SPNO: spno,
	}

	return wx.NewPostAction(urls.CorpOAGetApprovalDetail,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsOpenApprovalData struct {
	ThirdNO string `json:"thirdNo"`
}

type ResultOpenApprovalData struct {
	Data *OpenApprovalData `json:"data"`
}

type ApprovalNodes struct {
	ApprovalNode []*ApprovalNode `json:"ApprovalNode"`
}

type ApprovalNode struct {
	NodeStatus int                `json:"NodeStatus"`
	NodeAttr   int                `json:"NodeAttr"`
	NodeType   int                `json:"NodeType"`
	Items      *ApprovalNodeItems `json:"Items"`
}

type ApprovalNodeItems struct {
	Item []*ApprovalNodeItem `json:"Item"`
}

type ApprovalNodeItem struct {
	ItemName   string `json:"ItemName"`
	ItemParty  string `json:"ItemParty"`
	ItemImage  string `json:"ItemImage"`
	ItemUserID string `json:"ItemUserId"`
	ItemStatus int    `json:"ItemStatus"`
	ItemSpeech string `json:"ItemSpeech"`
	ItemOPTime int64  `json:"ItemOPTime"`
}

type ApprovalNotifyNodes struct {
	NotifyNode []*ApprovalNotifyNode `json:"NotifyNode"`
}

type ApprovalNotifyNode struct {
	ItemName   string `json:"ItemName"`
	ItemParty  string `json:"ItemParty"`
	ItemImage  string `json:"ItemImage"`
	ItemUserID string `json:"ItemUserId"`
}

type OpenApprovalData struct {
	ThirdNO        string               `json:"ThirdNo"`
	OpenTemplateID string               `json:"OpenTemplateId"`
	OpenSPName     string               `json:"OpenSpName"`
	OpenSPStatus   int                  `json:"OpenSpstatus"`
	ApplyTime      int64                `json:"ApplyTime"`
	ApplyUsername  string               `json:"ApplyUsername"`
	ApplyUserParty string               `json:"ApplyUserParty"`
	ApplyUserImage string               `json:"ApplyUserImage"`
	ApplyUserID    string               `json:"ApplyUserId"`
	ApprovalNodes  *ApprovalNodes       `json:"ApprovalNodes"`
	NotifyNodes    *ApprovalNotifyNodes `json:"NotifyNodes"`
	ApproverStep   int                  `json:"ApproverStep"`
}

// GetOpenApprovalData 查询自建应用审批单当前状态
func GetOpenApprovalData(thirdNO string, result *ResultOpenApprovalData) wx.Action {
	params := &ParamsOpenApprovalData{
		ThirdNO: thirdNO,
	}

	return wx.NewPostAction(urls.CorpOAGetOpenApprovalData,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
