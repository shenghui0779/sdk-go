package addr_book

import (
	"encoding/json"
	"fmt"

	"github.com/shenghui0779/yiigo"
	"github.com/tidwall/gjson"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type LinkedCorpAttrType int

const (
	LinkedCorpAttrText LinkedCorpAttrType = 0
	LinkedCorpAttrWeb  LinkedCorpAttrType = 1
)

type LinkedCorpPermList struct {
	UserIDs       []string `json:"userids"`
	DepartmentIDs []string `json:"department_ids"`
}

func GetLinkedCorpPermList(dest *LinkedCorpPermList) wx.Action {
	return wx.NewPostAction(urls.CorpLinkedCorpPermList,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

type LinkedCorpUser struct {
	UserID     string                 `json:"userid"`
	Name       string                 `json:"name"`
	Department []string               `json:"department"`
	Mobile     string                 `json:"mobile"`
	Telephone  string                 `json:"telephone"`
	EMail      string                 `json:"email"`
	Position   string                 `json:"position"`
	CorpID     string                 `json:"corpid"`
	ExtAttr    *LinkedCorpUserExtAttr `json:"extattr"`
}

type LinkedCorpUserExtAttr struct {
	Attrs []*LinkedCorpUserAttr `json:"attrs"`
}

type LinkedCorpUserAttr struct {
	Type  LinkedCorpAttrType `json:"type" `
	Name  string             `json:"name"`
	Value string             `json:"value"`
	Text  map[string]string  `json:"text"`
	Web   map[string]string  `json:"web"`
}

func GetLinkedCorpUser(dest *LinkedCorpUser, corpID, userID string) wx.Action {
	return wx.NewPostAction(urls.CorpLinkedCorpUserGet,
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

func GetLinkedCorpUserSimpleList(dest *[]*LinkedCorpUser, linkedID, departmentID string, fetchChild bool) wx.Action {
	child := 0

	if fetchChild {
		child = 1
	}

	return wx.NewPostAction(urls.CorpLinkedCorpUserSimpleList,
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

func GetLinkedCorpUserList(dest *[]*LinkedCorpUser, linkedID, departmentID string, fetchChild bool) wx.Action {
	child := 0

	if fetchChild {
		child = 1
	}

	return wx.NewPostAction(urls.CorpLinkedCorpUserList,
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

type LinkedCorpDepartment struct {
	DepartmentID   string `json:"department_id"`
	DepartmentName string `json:"department_name"`
	ParentID       string `json:"parentid"`
	Order          int    `json:"order"`
}

func GetLinkedCorpDeparmentList(dest *[]*LinkedCorpDepartment, linkedID, departmentID string) wx.Action {
	return wx.NewPostAction(urls.CorpLinkedCorpDepartmentList,
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
