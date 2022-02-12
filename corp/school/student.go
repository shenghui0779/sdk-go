package school

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsStudentCreate struct {
	StudentUserID string  `json:"student_userid"`
	Name          string  `json:"name"`
	Department    []int64 `json:"department"`
}

func CreateStudent(params *ParamsStudentCreate) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolStudentCreate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsStudentUpdate struct {
	StudentUserID    string  `json:"student_userid"`
	NewStudentUserID string  `json:"new_student_userid,omitempty"`
	Name             string  `json:"name,omitempty"`
	Department       []int64 `json:"department,omitempty"`
}

func UpdateStudent(params *ParamsStudentUpdate) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolStudentUpdate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

func DeleteStudent(userID string) wx.Action {
	return wx.NewGetAction(urls.CorpSchoolStudentDelete,
		wx.WithQuery("userid", userID),
	)
}

type StudentErrResult struct {
	StudentUserID string `json:"student_userid"`
	ErrCode       int    `json:"errcode"`
	ErrMsg        string `json:"errmsg"`
}

type ParamsStudentBatchCreate struct {
	Students []*ParamsStudentCreate `json:"students"`
}

type ResultStudentBatchCreate struct {
	ResultList []*StudentErrResult `json:"result_list"`
}

func BatchCreateStudent(params *ParamsStudentBatchCreate, result *ResultStudentBatchCreate) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolStudentBatchCreate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsStudentBatchUpdate struct {
	Students []*ParamsStudentUpdate `json:"students"`
}

type ResultStudentBatchUpdate struct {
	ResultList []*StudentErrResult `json:"result_list"`
}

func BatchUpdateStudent(params *ParamsStudentBatchUpdate, result *ResultStudentBatchUpdate) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolStudentBatchUpdate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsStudentBatchDelete struct {
	UserIDList []string `json:"useridlist"`
}

type ResultStudentBatchDelete struct {
	ResultList []*StudentErrResult `json:"result_list"`
}

func BatchDeleteStudent(userIDs []string, result *ResultStudentBatchDelete) wx.Action {
	params := &ParamsStudentBatchDelete{
		UserIDList: userIDs,
	}

	return wx.NewPostAction(urls.CorpSchoolStudentBatchDelete,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
