package corp

import (
	"encoding/json"
	"strconv"

	"github.com/tidwall/gjson"

	"github.com/shenghui0779/gochat/wx"
)

type Department struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name"`
	NameEN   string `json:"name_en,omitempty"`
	ParentID int    `json:"parentid,omitempty"`
	Order    int    `json:"order,omitempty"`
}

// DepartmentCreate 创建部门
func DepartmentCreate(data *Department) wx.Action {
	return wx.NewPostAction(DepartmentCreateURL,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(data)
		}),
		wx.WithDecode(func(resp []byte) error {
			data.ID = int(gjson.GetBytes(resp, "id").Int())

			return nil
		}),
	)
}

// DepartmentUpdate 更新部门
func DepartmentUpdate(data *Department) wx.Action {
	return wx.NewPostAction(DepartmentUpdateURL,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(data)
		}),
	)
}

// DepartmentDelete 删除部门
func DepartmentDelete(id int) wx.Action {
	return wx.NewGetAction(DepartmentDeleteURL,
		wx.WithQuery("id", strconv.Itoa(id)),
	)
}

// DepartmentList 获取部门列表
func DepartmentList(dest *[]*Department, id ...int) wx.Action {
	options := []wx.ActionOption{
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal([]byte(gjson.GetBytes(resp, "department").Raw), dest)
		}),
	}

	if len(id) != 0 {
		options = append(options, wx.WithQuery("id", strconv.Itoa(id[0])))
	}

	return wx.NewGetAction(DepartmentListURL, options...)
}
