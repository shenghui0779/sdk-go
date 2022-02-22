package linkedcorp

import (
	"encoding/json"
	"fmt"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsDepartmentList struct {
	DepartmentID string `json:"department_id"`
}

type DepartmentListData struct {
	DepartmentID   string `json:"department_id"`
	DepartmentName string `json:"department_name"`
	ParentID       string `json:"parentid"`
	Order          int    `json:"order"`
}

type ResultDepartmentList struct {
	DepartmentList []*DepartmentListData `json:"department_list"`
}

// ListDeparment 获取互联企业部门列表
func ListDeparment(linkedID, departmentID string, result *ResultDepartmentList) wx.Action {
	params := &ParamsDepartmentList{
		DepartmentID: fmt.Sprintf("%s/%s", linkedID, departmentID),
	}

	return wx.NewPostAction(urls.CorpLinkedcorpDepartmentList,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
