package linkedcorp

import (
	"encoding/json"
	"fmt"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
	"github.com/shenghui0779/yiigo"
	"github.com/tidwall/gjson"
)

type Department struct {
	DepartmentID   string `json:"department_id"`
	DepartmentName string `json:"department_name"`
	ParentID       string `json:"parentid"`
	Order          int    `json:"order"`
}

func GetDeparmentList(dest *[]*Department, linkedID, departmentID string) wx.Action {
	return wx.NewPostAction(urls.CorpLinkedcorpDepartmentList,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(yiigo.X{
				"department_id": fmt.Sprintf("%s/%s", linkedID, departmentID),
			})
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal([]byte(gjson.GetBytes(resp, "department_list").Raw), dest)
		}),
	)
}
