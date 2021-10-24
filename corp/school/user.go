package school

import (
	"encoding/json"
	"strconv"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type Student struct {
	StudentUserID string    `json:"student_userid"`
	Name          string    `json:"name"`
	Department    []int64   `json:"department"`
	Parents       []*Parent `json:"parents"`
}

type ResultUserGet struct {
	UserType int      `json:"user_type"`
	Student  *Student `json:"student"`
	Parent   *Parent  `json:"parent"`
}

func GetUser(userID string, result *ResultUserGet) wx.Action {
	return wx.NewGetAction(urls.CorpSchoolUserGet,
		wx.WithQuery("userid", userID),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsUserList struct {
	DepartmentID int64 `json:"department_id"`
	FetchChild   int   `json:"fetch_child"`
}

type ResultUserList struct {
	Students []*Student `json:"students"`
}

func ListUser(params *ParamsUserList, result *ResultUserList) wx.Action {
	return wx.NewGetAction(urls.CorpSchoolUserList,
		wx.WithQuery("department_id", strconv.FormatInt(params.DepartmentID, 10)),
		wx.WithQuery("fetch_child", strconv.Itoa(params.FetchChild)),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsArchSyncModeSet struct {
	ArchSyncMode int `json:"arch_sync_mode"`
}

func SetArchSyncMode(params *ParamsArchSyncModeSet) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolSetArchSyncMode,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ResultParentList struct {
	Parents []*Parent `json:"parents"`
}

func ListParent(departmentID int64, result *ResultParentList) wx.Action {
	return wx.NewGetAction(urls.CorpSchoolParentList,
		wx.WithQuery("department_id", strconv.FormatInt(departmentID, 10)),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
