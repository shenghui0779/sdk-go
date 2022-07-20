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

// CreateParent 创建家长
func CreateParent(params *ParamsParentCreate) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolParentCreate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type ParamsParentUpdate struct {
	ParentUserID    string         `json:"parent_userid"`
	NewParentUserID string         `json:"new_parent_userid,omitempty"`
	Mobile          string         `json:"mobile,omitempty"`
	Children        []*ParamsChild `json:"children,omitempty"`
}

// UpdateParent 更新家长
func UpdateParent(params *ParamsParentUpdate) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolParentUpdate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

// DeleteParent 删除家长
func DeleteParent(userID string) wx.Action {
	return wx.NewGetAction(urls.CorpSchoolParentDelete,
		wx.WithQuery("userid", userID),
	)
}

type ParentErrRet struct {
	ParentUserID string `json:"parent_userid"`
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
}

type ParamsParentBatchCreate struct {
	Parents []*ParamsParentCreate `json:"parents"`
}

type ResultParentBatchCreate struct {
	ResultList []*ParentErrRet `json:"result_list"`
}

// BatchCreateParent 批量创建家长
func BatchCreateParent(params *ParamsParentBatchCreate, result *ResultParentBatchCreate) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolParentBatchCreate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ParamsParentBatchUpdate struct {
	Parents []*ParamsParentUpdate `json:"parents"`
}

type ResultParentBatchUpdate struct {
	ResultList []*ParentErrRet `json:"result_list"`
}

// BatchUpdateParent 批量更新家长
func BatchUpdateParent(params *ParamsParentBatchUpdate, result *ResultParentBatchUpdate) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolParentBatchUpdate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ParamsParentBatchDelete struct {
	UserIDList []string `json:"useridlist"`
}

type ResultParentBatchDelete struct {
	ResultList []*ParentErrRet `json:"result_list"`
}

// BatchDeleteParent 批量删除家长
func BatchDeleteParent(userIDs []string, result *ResultParentBatchDelete) wx.Action {
	params := &ParamsParentBatchDelete{
		UserIDList: userIDs,
	}

	return wx.NewPostAction(urls.CorpSchoolParentBatchDelete,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}
