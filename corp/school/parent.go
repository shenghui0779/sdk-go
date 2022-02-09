package school

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsChild struct {
	StudentUserID string `json:"student_userid"`
	Relation      string `json:"relation"`
}

type ParamsParentCreate struct {
	ParentUserID string         `json:"parent_userid"`
	Mobile       string         `json:"mobile"`
	ToInvite     *bool          `json:"to_invite,omitempty"`
	Children     []*ParamsChild `json:"children"`
}

func CreateParent(params *ParamsParentCreate) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolParentCreate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsParentUpdate struct {
	ParentUserID    string         `json:"parent_userid"`
	NewParentUserID string         `json:"new_parent_userid,omitempty"`
	Mobile          string         `json:"mobile,omitempty"`
	Children        []*ParamsChild `json:"children,omitempty"`
}

func UpdateParent(params *ParamsParentUpdate) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolParentUpdate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

func DeleteParent(userID string) wx.Action {
	return wx.NewGetAction(urls.CorpSchoolParentDelete,
		wx.WithQuery("userid", userID),
	)
}

type ParentErrResult struct {
	ParentUserID string `json:"parent_userid"`
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
}

type ParamsParentBatchCreate struct {
	Parents []*ParamsParentCreate `json:"parents"`
}

type ResultParentBatchCreate struct {
	ResultList []*ParentErrResult `json:"result_list"`
}

func BatchCreateParent(params *ParamsParentBatchCreate, result *ResultParentBatchCreate) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolParentBatchCreate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsParentBatchUpdate struct {
	Parents []*ParamsParentUpdate `json:"parents"`
}

type ResultParentBatchUpdate struct {
	ResultList []*ParentErrResult `json:"result_list"`
}

func BatchUpdateParent(params *ParamsParentBatchUpdate, result *ResultParentBatchUpdate) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolParentBatchUpdate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsParentBatchDelete struct {
	UserIDList []string `json:"useridlist"`
}

type ResultParentBatchDelete struct {
	ResultList []*ParentErrResult `json:"result_list"`
}

func BatchDeleteParent(params *ParamsParentBatchDelete, result *ResultParentBatchDelete) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolParentBatchDelete,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
