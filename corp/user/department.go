package user

import (
	"encoding/json"
	"strconv"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsDepartmentCreate struct {
	Name     string `json:"name"`
	NameEN   string `json:"name_en,omitempty"`
	ParentID int64  `json:"parentid"`
	Order    int64  `json:"order,omitempty"`
}

type ResultDepartmentCreate struct {
	ID int64 `json:"id"`
}

// CreateDepartment 创建部门
func CreateDepartment(params *ParamsDepartmentCreate, result *ResultDepartmentCreate) wx.Action {
	return wx.NewPostAction(urls.CorpUserDepartmentCreate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsDepartmentUpdate struct {
	ID       int64  `json:"id"`
	Name     string `json:"name,omitempty"`
	NameEN   string `json:"name_en,omitempty"`
	ParentID int64  `json:"parentid,omitempty"`
	Order    int64  `json:"order,omitempty"`
}

// UpdateDepartment 更新部门
func UpdateDepartment(params *ParamsDepartmentUpdate) wx.Action {
	return wx.NewPostAction(urls.CorpUserDepartmentUpdate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

// DeleteDepartment 删除部门
func DeleteDepartment(id int64) wx.Action {
	return wx.NewGetAction(urls.CorpUserDepartmentDelete,
		wx.WithQuery("id", strconv.FormatInt(id, 10)),
	)
}

type Department struct {
	ID               int64    `json:"id"`
	Name             string   `json:"name"`
	NameEN           string   `json:"name_en"`
	DepartmentLeader []string `json:"department_leader"`
	ParentID         int64    `json:"parentid"`
	Order            int64    `json:"order"`
}

type ResultDepartmentList struct {
	Department []*Department `json:"department"`
}

// ListDepartment 获取部门列表
func ListDepartment(id int64, result *ResultDepartmentList) wx.Action {
	options := []wx.ActionOption{
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	}

	if id > 0 {
		options = append(options, wx.WithQuery("id", strconv.FormatInt(id, 10)))
	}

	return wx.NewGetAction(urls.CorpUserDepartmentList, options...)
}

type SimpleDepartment struct {
	ID       int64 `json:"id"`
	ParentID int64 `json:"parentid"`
	Order    int64 `json:"order"`
}

type ResultSimpleDepartmentList struct {
	DepartmentID []*SimpleDepartment `json:"department_id"`
}

func ListSimpleDepartment(id int64, result *ResultSimpleDepartmentList) wx.Action {
	options := []wx.ActionOption{
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	}

	if id > 0 {
		options = append(options, wx.WithQuery("id", strconv.FormatInt(id, 10)))
	}

	return wx.NewGetAction(urls.CorpUserDepartmentSimpleList, options...)
}

type ResultDepartmentGet struct {
	Department *Department `json:"department"`
}

func GetDepartment(id int64, result *ResultDepartmentGet) wx.Action {
	options := []wx.ActionOption{
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	}

	if id > 0 {
		options = append(options, wx.WithQuery("id", strconv.FormatInt(id, 10)))
	}

	return wx.NewGetAction(urls.CorpUserDepartmentGet, options...)
}
