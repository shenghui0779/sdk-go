package addr_book

import (
	"encoding/json"
	"fmt"

	"github.com/shenghui0779/yiigo"
	"github.com/tidwall/gjson"

	"github.com/shenghui0779/gochat/corp/common"
	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type LinkedcorpPermList struct {
	UserIDs       []string `json:"userids"`
	DepartmentIDs []string `json:"department_ids"`
}

func GetLinkedcorpPermList(dest *LinkedcorpPermList) wx.Action {
	return wx.NewPostAction(urls.CorpLinkedcorpPermList,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

type LinkedcorpUser struct {
	UserID     string          `json:"userid"`
	Name       string          `json:"name"`
	Department []string        `json:"department"`
	Mobile     string          `json:"mobile"`
	Telephone  string          `json:"telephone"`
	EMail      string          `json:"email"`
	Position   string          `json:"position"`
	CorpID     string          `json:"corpid"`
	ExtAttr    *common.ExtAttr `json:"extattr"`
}

func GetLinkedcorpUser(dest *LinkedcorpUser, corpID, userID string) wx.Action {
	return wx.NewPostAction(urls.CorpLinkedcorpUserGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(yiigo.X{
				"userid": fmt.Sprintf("%s/%s", corpID, userID),
			})
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal([]byte(gjson.GetBytes(resp, "user_info").Raw), dest)
		}),
	)
}

func GetLinkedcorpUserSimpleList(dest *[]*LinkedcorpUser, linkedID, departmentID string, fetchChild bool) wx.Action {
	child := 0

	if fetchChild {
		child = 1
	}

	return wx.NewPostAction(urls.CorpLinkedcorpUserSimpleList,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(yiigo.X{
				"department_id": fmt.Sprintf("%s/%s", linkedID, departmentID),
				"fetch_child":   child,
			})
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal([]byte(gjson.GetBytes(resp, "userlist").Raw), dest)
		}),
	)
}

func GetLinkedcorpUserList(dest *[]*LinkedcorpUser, linkedID, departmentID string, fetchChild bool) wx.Action {
	child := 0

	if fetchChild {
		child = 1
	}

	return wx.NewPostAction(urls.CorpLinkedcorpUserList,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(yiigo.X{
				"department_id": fmt.Sprintf("%s/%s", linkedID, departmentID),
				"fetch_child":   child,
			})
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal([]byte(gjson.GetBytes(resp, "userlist").Raw), dest)
		}),
	)
}

type LinkedcorpDepartment struct {
	DepartmentID   string `json:"department_id"`
	DepartmentName string `json:"department_name"`
	ParentID       string `json:"parentid"`
	Order          int    `json:"order"`
}

func GetLinkedcorpDeparmentList(dest *[]*LinkedcorpDepartment, linkedID, departmentID string) wx.Action {
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
