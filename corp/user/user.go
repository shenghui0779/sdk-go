package user

import (
	"encoding/json"
	"strconv"

	"github.com/shenghui0779/gochat/corp/common"
	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsUserCreate struct {
	UserID           string                  `json:"user_id"`
	Name             string                  `json:"name"`
	Alias            string                  `json:"alias,omitempty"`
	Mobile           string                  `json:"mobile,omitempty"`
	Department       []int64                 `json:"department"`
	Order            []int64                 `json:"order,omitempty"`
	Position         string                  `json:"position,omitempty"`
	Gender           int                     `json:"gender,omitempty"`
	Email            string                  `json:"email,omitempty"`
	IsLeaderInDept   []int64                 `json:"is_leader_in_dept,omitempty"`
	Enable           int                     `json:"enable,omitempty"`
	AvatarMediaID    string                  `json:"avatar_mediaid,omitempty"`
	Telephone        string                  `json:"telephone,omitempty"`
	Address          string                  `json:"address,omitempty"`
	MainDepartment   int64                   `json:"main_department,omitempty"`
	ExtAttr          *common.ExtAttr         `json:"extattr,omitempty"`
	ToInvite         bool                    `json:"to_invite,omitempty"`
	ExternalPosition string                  `json:"external_position,omitempty"`
	ExternalProfile  *common.ExternalProfile `json:"external_profile,omitempty"`
}

// CreateUser 创建成员
func CreateUser(params *ParamsUserCreate) wx.Action {
	return wx.NewPostAction(urls.CorpUserCreate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsUserUpdate struct {
	UserID           string                  `json:"user_id"`
	Name             string                  `json:"name,omitempty"`
	Alias            string                  `json:"alias,omitempty"`
	Mobile           string                  `json:"mobile,omitempty"`
	Department       []int64                 `json:"department"`
	Order            []int64                 `json:"order,omitempty"`
	Position         string                  `json:"position,omitempty"`
	Gender           int                     `json:"gender,omitempty"`
	Email            string                  `json:"email,omitempty"`
	IsLeaderInDept   []int64                 `json:"is_leader_in_dept,omitempty"`
	Enable           int                     `json:"enable,omitempty"`
	AvatarMediaID    string                  `json:"avatar_mediaid,omitempty"`
	Telephone        string                  `json:"telephone,omitempty"`
	Address          string                  `json:"address,omitempty"`
	MainDepartment   int64                   `json:"main_department,omitempty"`
	ExtAttr          *common.ExtAttr         `json:"extattr,omitempty"`
	ToInvite         bool                    `json:"to_invite,omitempty"`
	ExternalPosition string                  `json:"external_position,omitempty"`
	ExternalProfile  *common.ExternalProfile `json:"external_profile,omitempty"`
}

// UpdateUser 更新成员
func UpdateUser(params *ParamsUserUpdate) wx.Action {
	return wx.NewPostAction(urls.CorpUserUpdate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type UserInfo struct {
	UserID           string                  `json:"user_id"`
	Name             string                  `json:"name"`
	Alias            string                  `json:"alias,omitempty"`
	Mobile           string                  `json:"mobile,omitempty"`
	Department       []int64                 `json:"department"`
	Order            []int64                 `json:"order,omitempty"`
	Position         string                  `json:"position,omitempty"`
	Gender           int                     `json:"gender,omitempty"`
	Email            string                  `json:"email,omitempty"`
	IsLeaderInDept   []int64                 `json:"is_leader_in_dept,omitempty"`
	Enable           int                     `json:"enable,omitempty"`
	Avatar           string                  `json:"avatar,omitempty"`
	ThumbAvatar      string                  `json:"thumb_avatar,omitempty"`
	Telephone        string                  `json:"telephone,omitempty"`
	Address          string                  `json:"address,omitempty"`
	OpenUserID       string                  `json:"open_userid,omitempty"`
	MainDepartment   int64                   `json:"main_department,omitempty"`
	ExtAttr          *common.ExtAttr         `json:"extattr,omitempty"`
	Status           int                     `json:"status,omitempty"`
	QRCode           string                  `json:"qr_code,omitempty"`
	ExternalPosition string                  `json:"external_position,omitempty"`
	ExternalProfile  *common.ExternalProfile `json:"external_profile,omitempty"`
}

// GetUser 读取成员
func GetUser(userID string, result *UserInfo) wx.Action {
	return wx.NewGetAction(urls.CorpUserGet,
		wx.WithQuery("userid", userID),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// DeleteUser 删除成员
func DeleteUser(userID string) wx.Action {
	return wx.NewGetAction(urls.CorpUserDelete,
		wx.WithQuery("userid", userID),
	)
}

type ParamsUserBatchDelete struct {
	UserIDList []string `json:"useridlist"`
}

// BatchDeleteUser 批量删除成员
func BatchDeleteUser(userids ...string) wx.Action {
	params := &ParamsUserBatchDelete{
		UserIDList: userids,
	}

	return wx.NewPostAction(urls.CorpUserBatchDelete,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type UserSimple struct {
	UserID     string  `json:"userid"`
	Name       string  `json:"name"`
	Department []int64 `json:"department"`
	OpenUserID string  `json:"open_userid"`
}

type ParamsUserSimpleList struct {
	DepartmentID int64 `json:"department_id"`
	FetchChild   int   `json:"fetch_child"`
}

type ResultUserSimipleList struct {
	UserList []*UserSimple `json:"userlist"`
}

// GetUserSimpleList 获取部门成员
func GetUserSimpleList(params *ParamsUserSimpleList, result *ResultUserSimipleList) wx.Action {
	return wx.NewGetAction(urls.CorpUserSimpleList,
		wx.WithQuery("department_id", strconv.FormatInt(params.DepartmentID, 10)),
		wx.WithQuery("fetch_child", strconv.Itoa(params.FetchChild)),
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
	UserList []*UserSimple `json:"userlist"`
}

// GetUserList 获取部门成员详情
func GetUserList(params *ParamsUserList, result *ResultUserList) wx.Action {
	return wx.NewGetAction(urls.CorpUserList,
		wx.WithQuery("department_id", strconv.FormatInt(params.DepartmentID, 10)),
		wx.WithQuery("fetch_child", strconv.Itoa(params.FetchChild)),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsOpenIDConvert struct {
	UserID string `json:"userid"`
}

type ResultOpenIDConvert struct {
	OpenID string `json:"openid"`
}

// ConvertToOpenID userid转openid
func ConvertToOpenID(userID string, result *ResultOpenIDConvert) wx.Action {
	params := &ParamsOpenIDConvert{
		UserID: userID,
	}

	return wx.NewPostAction(urls.CorpConvertToOpenID,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsUserIDConvert struct {
	OpenID string `json:"openid"`
}

type ResultUserIDConvert struct {
	UserID string `json:"userid"`
}

// ConvertToUserID openid转userid
func ConvertToUserID(openid string, result *ResultUserIDConvert) wx.Action {
	params := &ParamsUserIDConvert{
		OpenID: openid,
	}

	return wx.NewPostAction(urls.CorpConvertToUserID,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsBatchInvite struct {
	User  []string `json:"user,omitempty"`  // 成员ID列表，最多支持1000个
	Party []string `json:"party,omitempty"` // 部门ID列表，最多支持100个
	Tag   []string `json:"tag,omitempty"`   // 标签ID列表，最多支持100个
}

type ResultBatchInvite struct {
	InvalidUser  []string `json:"invaliduser"`
	InvalidParth []string `json:"invalidparth"`
	InvalidTag   []string `json:"invalidtag"`
}

// BatchInvite 邀请成员
func BatchInvite(params *ParamsBatchInvite, result *ResultBatchInvite) wx.Action {
	return wx.NewPostAction(urls.CorpBatchInvite,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type JoinQRCode struct {
	URL string `json:"join_qrcode"` // 二维码链接，有效期7天
}

// GetJoinQRCode 获取加入企业二维码
// 尺寸：
// 1 - 171 x 171
// 2 - 399 x 399
// 3 - 741 x 741
// 4 - 2052 x 2052
func GetJoinQRCode(size int, result *JoinQRCode) wx.Action {
	options := []wx.ActionOption{
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	}

	if size != 0 {
		options = append(options, wx.WithQuery("size_type", strconv.Itoa(size)))
	}

	return wx.NewGetAction(urls.CorpJoinQRCode, options...)
}

type ParamsActiveStat struct {
	Date string `json:"date"`
}

type ResultActiveStat struct {
	ActiveCnt int `json:"active_cnt"`
}

// GetUserActiveStat 获取企业活跃成员数
func GetUserActiveStat(date string, result *ResultActiveStat) wx.Action {
	params := &ParamsActiveStat{
		Date: date,
	}

	return wx.NewPostAction(urls.CorpUserActiveStat,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
