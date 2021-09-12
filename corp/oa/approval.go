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

type ControlProperty struct {
	Control     string         `json:"control"`
	ID          string         `json:"id"`
	Title       []*DisplayText `json:"title"`
	Placeholder []*DisplayText `json:"placeholder"`
	Require     int            `json:"require"`
	UnPrint     int            `json:"un_print"`
}

type ControlConfig struct {
	Date         *ControlDate       `json:"date,omitempty"`
	Selector     *ControlSelector   `json:"selector,omitempty"`
	Table        *ControlTable      `json:"table,omitempty"`
	Attendance   *ControlAttendance `json:"attendance,omitempty"`
	VacationList *ControlVacation   `json:"vacation_list,omitempty"`
}

type ControlDate struct {
	Type string `json:"type"`
}

type ControlSelector struct {
	Type    string            `json:"type"`
	ExpType int               `json:"exp_type"`
	Options []*SelectorOption `json:"options"`
}

type SelectorOption struct {
	Key   string         `json:"key"`
	Value []*DisplayText `json:"value"`
}

type ControlContact struct {
	Type string `json:"type"`
	Mode string `json:"mode"`
}

type ControlTable struct {
	Childern  []*TableChild `json:"childern"`
	StatField interface{}   `json:"stat_field"`
}

type TableChild struct {
	Property *ControlProperty
}

type ControlAttendance struct {
	Type      int                  `json:"type"`
	DateRange *AttendanceDateRange `json:"date_range"`
}

type AttendanceDateRange struct {
	Type string `json:"type"`
}

type ControlVacation struct {
	Item []*VacationItem `json:"item"`
}

type VacationItem struct {
	ID   int64          `json:"id"`
	Name []*DisplayText `json:"name"`
}

type ParamsTemplateDetailGet struct {
	TemplateID string `json:"template_id"`
}

type ResultTemplateDetailGet struct {
	TemplateNames   []*DisplayText   `json:"template_names"`
	TemplateContent *TemplateContent `json:"template_content"`
}

func GetTemplateDetail(params *ParamsTemplateDetailGet, result *ResultTemplateDetailGet) wx.Action {
	return wx.NewPostAction(urls.CorpOAGetTemplateDetail,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
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

type ApplyData struct {
	Contents []*ApplyContent `json:"contents"`
}

type ApplyContent struct {
	Control string       `json:"control"`
	ID      string       `json:"id"`
	Value   *DisplayText `json:"value"`
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
	SpNo string `json:"sp_no"`
}

func ApplyEvent(params *ParamsApplyEvent, result *ResultApplyEvent) wx.Action {
	return wx.NewPostAction(urls.CorpOAApplyEvent,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
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

type NotifyNodes struct {
	NotifyNode []*NotifyNode `json:"NotifyNode"`
}

type NotifyNode struct {
	ItemName   string `json:"ItemName"`
	ItemParty  string `json:"ItemParty"`
	ItemImage  string `json:"ItemImage"`
	ItemUserID string `json:"ItemUserId"`
}

type OpenApprovalData struct {
	ThirdNO        string         `json:"ThirdNo"`
	OpenTemplateID string         `json:"OpenTemplateId"`
	OpenSPName     string         `json:"OpenSpName"`
	OpenSPStatus   string         `json:"OpenSpstatus"`
	ApplyTime      int64          `json:"ApplyTime"`
	ApplyUsername  string         `json:"ApplyUsername"`
	ApplyUserParty string         `json:"ApplyUserParty"`
	ApplyUserImage string         `json:"ApplyUserImage"`
	ApplyUserID    string         `json:"ApplyUserId"`
	ApprovalNodes  *ApprovalNodes `json:"ApprovalNodes"`
	NotifyNodes    *NotifyNodes   `json:"NotifyNodes"`
	ApproverStep   int            `json:"ApproverStep"`
}

type ParamsOpenApprovalDataGet struct {
	ThirdNO string `json:"thirdNo"`
}

type ResultOpenApprovalDataGet struct {
	Data *OpenApprovalData `json:"data"`
}

func GetOpenApprovalData(params *ParamsOpenApprovalDataGet, result *ResultOpenApprovalDataGet) wx.Action {
	return wx.NewPostAction(urls.CorpOAGetOpenApprovalData,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
