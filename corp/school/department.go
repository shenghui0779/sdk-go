package school

import (
	"encoding/json"
	"strconv"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type DepartmentAdminCreate struct {
	UserID  string `json:"userid"`
	Type    int    `json:"type"`
	Subject string `json:"subject,omitempty"`
}

type ParamsDepartmentCreate struct {
	Name             string                   `json:"name,omitempty"`
	ParentID         int64                    `json:"parentid"`
	ID               int64                    `json:"id,omitempty"`
	RegisterYear     int                      `json:"register_year,omitempty"`
	StandardGrade    int                      `json:"standard_grade,omitempty"`
	Order            int                      `json:"order,omitempty"`
	DepartmentAdmins []*DepartmentAdminCreate `json:"department_admins,omitempty"`
}

type ResultDepartmentCreate struct {
	ID int64 `json:"id"`
}

func CreateDepartment(params *ParamsDepartmentCreate, result *ResultDepartmentCreate) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolDepartmentCreate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type DepartmentAdminUpdate struct {
	OP      int    `json:"op"`
	UserID  string `json:"userid"`
	Type    int    `json:"type,omitempty"`
	Subject string `json:"subject,omitempty"`
}

type ParamsDepartmentUpdate struct {
	Name             string                   `json:"name,omitempty"`
	ParentID         int64                    `json:"parentid,omitempty"`
	ID               int64                    `json:"id"`
	RegisterYear     int                      `json:"register_year,omitempty"`
	StandardGrade    int                      `json:"standard_grade,omitempty"`
	Order            int                      `json:"order,omitempty"`
	DepartmentAdmins []*DepartmentAdminUpdate `json:"department_admins,omitempty"`
}

func UpdateDeparment(params *ParamsDepartmentUpdate) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolDepartmentUpdate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

func DeleteDepartment(id int64) wx.Action {
	return wx.NewGetAction(urls.CorpSchoolDepartmentDelete,
		wx.WithQuery("id", strconv.FormatInt(id, 10)),
	)
}

type DepartmentAdmin struct {
	UserID  string `json:"userid"`
	Type    int    `json:"type"`
	Subject string `json:"subject"`
}

type Department struct {
	Name             string             `json:"name,omitempty"`
	ParentID         int64              `json:"parentid"`
	ID               int64              `json:"id,omitempty"`
	RegisterYear     int                `json:"register_year,omitempty"`
	StandardGrade    int                `json:"standard_grade,omitempty"`
	Order            int                `json:"order,omitempty"`
	IsGraduated      int                `json:"is_graduated"`
	OpenGroupChat    int                `json:"open_group_chat"`
	GroupChatID      string             `json:"group_chat_id"`
	DepartmentAdmins []*DepartmentAdmin `json:"department_admins,omitempty"`
}

type ResultDepartmentList struct {
	Departments []*Department `json:"departments"`
}

func ListDepartment(id int64, result *ResultDepartmentList) wx.Action {
	options := []wx.ActionOption{
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	}

	if id > 0 {
		options = append(options, wx.WithQuery("id", strconv.FormatInt(id, 10)))
	}

	return wx.NewGetAction(urls.CorpSchoolDepartmentList, options...)
}

type ParamsUpgradeInfoSet struct {
	UpgradeTime   int64 `json:"upgrade_time"`
	UpgradeSwitch int   `json:"upgrade_switch"`
}

type ResultUpgradeInfoSet struct {
	NextUpgradeTime int64 `json:"next_upgrade_time"`
}

func SetUpgradeInfo(params *ParamsUpgradeInfoSet, result *ResultUpgradeInfoSet) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolSetUpgradeInfo,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
