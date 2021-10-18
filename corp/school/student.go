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

func DeleteStudent(userID string) wx.Action {
	return wx.NewGetAction(urls.CorpSchoolStudentDelete,
		wx.WithQuery("userid", userID),
	)
}

type ParamsStudentUpdate struct {
	StudentUserID    string  `json:"student_userid"`
	NewStudentUserID string  `json:"new_student_userid,omitempty"`
	Name             string  `json:"name,omitempty"`
	Department       []int64 `json:"department,omitempty"`
}

type StudentErrResult struct {
	StudentUserID string `json:"student_userid"`
	ErrCode       int    `json:"errcode"`
	ErrMsg        string `json:"errmsg"`
}
