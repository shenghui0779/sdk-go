package linkedcorp

import (
	"encoding/json"
	"fmt"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
	"github.com/shenghui0779/yiigo"
)

type ParamsDepartmentList struct {
	LinkedID     string `json:"linked_id"`
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

func ListDeparment(params *ParamsDepartmentList, result *ResultDepartmentList) wx.Action {
	return wx.NewPostAction(urls.CorpLinkedcorpDepartmentList,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(yiigo.X{
				"department_id": fmt.Sprintf("%s/%s", params.LinkedID, params.DepartmentID),
			})
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
