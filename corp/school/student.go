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

// CreateStudent 创建学生
func CreateStudent(params *ParamsStudentCreate) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolStudentCreate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type ParamsStudentUpdate struct {
	StudentUserID    string  `json:"student_userid"`
	NewStudentUserID string  `json:"new_student_userid,omitempty"`
	Name             string  `json:"name,omitempty"`
	Department       []int64 `json:"department,omitempty"`
}

// UpdateStudent 更新学生
func UpdateStudent(params *ParamsStudentUpdate) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolStudentUpdate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

// DeleteStudent 删除学生
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

// BatchCreateStudent 批量创建学生
func BatchCreateStudent(params *ParamsStudentBatchCreate, result *ResultStudentBatchCreate) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolStudentBatchCreate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
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

// BatchUpdateStudent 批量更新学生
func BatchUpdateStudent(params *ParamsStudentBatchUpdate, result *ResultStudentBatchUpdate) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolStudentBatchUpdate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
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

// BatchDeleteStudent 批量删除学生
func BatchDeleteStudent(userIDs []string, result *ResultStudentBatchDelete) wx.Action {
	params := &ParamsStudentBatchDelete{
		UserIDList: userIDs,
	}

	return wx.NewPostAction(urls.CorpSchoolStudentBatchDelete,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
