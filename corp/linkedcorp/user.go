package linkedcorp

import (
	"encoding/json"
	"fmt"

	"github.com/shenghui0779/yiigo"
	"github.com/tidwall/gjson"

	"github.com/shenghui0779/gochat/corp/common"
	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type UserInfo struct {
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

func GetUser(dest *UserInfo, corpID, userID string) wx.Action {
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

func GetUserSimpleList(dest *[]*UserInfo, linkedID, departmentID string, fetchChild bool) wx.Action {
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

func GetUserList(dest *[]*UserInfo, linkedID, departmentID string, fetchChild bool) wx.Action {
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
