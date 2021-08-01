package corp

import (
	"encoding/json"
	"strconv"

	"github.com/tidwall/gjson"

	"github.com/shenghui0779/gochat/wx"
)

type Department struct {
	ID       int64  `json:"id,omitempty"`
	Name     string `json:"name"`
	NameEN   string `json:"name_en,omitempty"`
	ParentID int64  `json:"parentid,omitempty"`
	Order    int64  `json:"order,omitempty"`
}

// CreateDepartment 创建部门
func CreateDepartment(data *Department) wx.Action {
	return wx.NewPostAction(DepartmentCreateURL,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(data)
		}),
		wx.WithDecode(func(resp []byte) error {
			data.ID = gjson.GetBytes(resp, "id").Int()

			return nil
		}),
	)
}

// UpdateDepartment 更新部门
func UpdateDepartment(data *Department) wx.Action {
	return wx.NewPostAction(DepartmentUpdateURL,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(data)
		}),
	)
}

// DeleteDepartment 删除部门
func DeleteDepartment(id int64) wx.Action {
	return wx.NewGetAction(DepartmentDeleteURL,
		wx.WithQuery("id", strconv.FormatInt(id, 10)),
	)
}

// GetDepartmentList 获取部门列表
func GetDepartmentList(dest *[]*Department, id ...int64) wx.Action {
	options := []wx.ActionOption{
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal([]byte(gjson.GetBytes(resp, "department").Raw), dest)
		}),
	}

	if len(id) != 0 {
		options = append(options, wx.WithQuery("id", strconv.FormatInt(id[0], 10)))
	}

	return wx.NewGetAction(DepartmentListURL, options...)
}
