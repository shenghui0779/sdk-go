package user

import (
	"encoding/json"
	"strconv"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ExternalProfile struct {
	ExternalCorpName string          `json:"external_corp_name"`
	WechatChannels   *WechatChannels `json:"wechat_channels"`
	ExternalAttr     []*Attr         `json:"external_attr"`
}

type WechatChannels struct {
	Nickname string `json:"nickname,omitempty"`
	Status   int    `json:"status,omitempty"`
}

type ExtAttr struct {
	Attrs []*Attr `json:"attrs"`
}

type Attr struct {
	Type        int        `json:"type"`
	Name        string     `json:"name"`
	Text        *AttrText  `json:"text,omitempty"`
	Web         *AttrWeb   `json:"web,omitempty"`
	Miniprogram *AttrMinip `json:"miniprogram,omitempty"`
}

type AttrText struct {
	Value string `json:"value"`
}

type AttrWeb struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

type AttrMinip struct {
	AppID    string `json:"appid"`
	Pagepath string `json:"pagepath"`
	Title    string `json:"title"`
}

type ParamsUserCreate struct {
	UserID           string           `json:"userid"`
	Name             string           `json:"name"`
	Alias            string           `json:"alias,omitempty"`
	Mobile           string           `json:"mobile,omitempty"`
	Department       []int64          `json:"department"`
	Order            []int            `json:"order,omitempty"`
	Position         string           `json:"position,omitempty"`
	Gender           string           `json:"gender,omitempty"`
	Email            string           `json:"email,omitempty"`
	BizMail          string           `json:"biz_mail,omitempty"`
	IsLeaderInDept   []int            `json:"is_leader_in_dept,omitempty"`
	DirectLeader     []string         `json:"direct_leader,omitempty"`
	Enable           int              `json:"enable,omitempty"`
	AvatarMediaID    string           `json:"avatar_mediaid,omitempty"`
	Telephone        string           `json:"telephone,omitempty"`
	Address          string           `json:"address,omitempty"`
	MainDepartment   int64            `json:"main_department,omitempty"`
	ExtAttr          *ExtAttr         `json:"extattr,omitempty"`
	ToInvite         *bool            `json:"to_invite,omitempty"`
	ExternalPosition string           `json:"external_position,omitempty"`
	ExternalProfile  *ExternalProfile `json:"external_profile,omitempty"`
}

// CreateUser 创建成员
func CreateUser(params *ParamsUserCreate) wx.Action {
	return wx.NewPostAction(urls.CorpUserCreate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type ParamsUserUpdate struct {
	UserID           string           `json:"userid"`
	Name             string           `json:"name,omitempty"`
	Alias            string           `json:"alias,omitempty"`
	Mobile           string           `json:"mobile,omitempty"`
	Department       []int64          `json:"department"`
	Order            []int            `json:"order,omitempty"`
	Position         string           `json:"position,omitempty"`
	Gender           string           `json:"gender,omitempty"`
	Email            string           `json:"email,omitempty"`
	BizMail          string           `json:"biz_mail,omitempty"`
	IsLeaderInDept   []int            `json:"is_leader_in_dept,omitempty"`
	DirectLeader     []string         `json:"direct_leader,omitempty"`
	Enable           int              `json:"enable,omitempty"`
	AvatarMediaID    string           `json:"avatar_mediaid,omitempty"`
	Telephone        string           `json:"telephone,omitempty"`
	Address          string           `json:"address,omitempty"`
	MainDepartment   int64            `json:"main_department,omitempty"`
	ExtAttr          *ExtAttr         `json:"extattr,omitempty"`
	ExternalPosition string           `json:"external_position,omitempty"`
	ExternalProfile  *ExternalProfile `json:"external_profile,omitempty"`
}

// UpdateUser 更新成员
func UpdateUser(params *ParamsUserUpdate) wx.Action {
	return wx.NewPostAction(urls.CorpUserUpdate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type User struct {
	UserID           string           `json:"userid"`
	Name             string           `json:"name"`
	Alias            string           `json:"alias"`
	Mobile           string           `json:"mobile"`
	Department       []int64          `json:"department"`
	Order            []int            `json:"order"`
	Position         string           `json:"position"`
	Gender           string           `json:"gender"`
	Email            string           `json:"email"`
	BizMail          string           `json:"biz_mail"`
	IsLeaderInDept   []int            `json:"is_leader_in_dept"`
	DirectLeader     []string         `json:"direct_leader,omitempty"`
	Avatar           string           `json:"avatar"`
	ThumbAvatar      string           `json:"thumb_avatar"`
	Telephone        string           `json:"telephone"`
	Address          string           `json:"address"`
	OpenUserID       string           `json:"open_userid"`
	MainDepartment   int64            `json:"main_department"`
	ExtAttr          *ExtAttr         `json:"extattr"`
	Status           int              `json:"status"`
	QRCode           string           `json:"qr_code"`
	ExternalPosition string           `json:"external_position"`
	ExternalProfile  *ExternalProfile `json:"external_profile"`
}

// GetUser 读取成员
func GetUser(userID string, result *User) wx.Action {
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
func BatchDeleteUser(userIDs ...string) wx.Action {
	params := &ParamsUserBatchDelete{
		UserIDList: userIDs,
	}

	return wx.NewPostAction(urls.CorpUserBatchDelete,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type SimpleUser struct {
	UserID     string  `json:"userid"`
	Name       string  `json:"name"`
	Department []int64 `json:"department"`
	OpenUserID string  `json:"open_userid"`
}

type ResultSimipleUserList struct {
	UserList []*SimpleUser `json:"userlist"`
}

// ListSimpleUser 获取部门成员
func ListSimpleUser(departmentID int64, fetchChild int, result *ResultSimipleUserList) wx.Action {
	return wx.NewGetAction(urls.CorpUserSimpleList,
		wx.WithQuery("department_id", strconv.FormatInt(departmentID, 10)),
		wx.WithQuery("fetch_child", strconv.Itoa(fetchChild)),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ResultUserList struct {
	UserList []*User `json:"userlist"`
}

// ListUser 获取部门成员详情
func ListUser(departmentID int64, fetchChild int, result *ResultUserList) wx.Action {
	return wx.NewGetAction(urls.CorpUserList,
		wx.WithQuery("department_id", strconv.FormatInt(departmentID, 10)),
		wx.WithQuery("fetch_child", strconv.Itoa(fetchChild)),
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

	return wx.NewPostAction(urls.CorpUserConvertToOpenID,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
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

	return wx.NewPostAction(urls.CorpUserConvertToUserID,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsBatchInvite struct {
	User  []string `json:"user,omitempty"`  // 成员ID列表，最多支持1000个
	Party []int64  `json:"party,omitempty"` // 部门ID列表，最多支持100个
	Tag   []int64  `json:"tag,omitempty"`   // 标签ID列表，最多支持100个
}

type ResultBatchInvite struct {
	InvalidUser  []string `json:"invaliduser"`
	InvalidParty []int64  `json:"invalidparty"`
	InvalidTag   []int64  `json:"invalidtag"`
}

// BatchInvite 邀请成员
func BatchInvite(userIDs []string, partyIDs, tagIDs []int64, result *ResultBatchInvite) wx.Action {
	params := &ParamsBatchInvite{
		User:  userIDs,
		Party: partyIDs,
		Tag:   tagIDs,
	}

	return wx.NewPostAction(urls.CorpUserBatchInvite,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
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

	return wx.NewGetAction(urls.CorpUserJoinQRCode, options...)
}

type ParamsActiveStat struct {
	Date string `json:"date"`
}

type ResultActiveStat struct {
	ActiveCnt int `json:"active_cnt"`
}

// GetActiveStat 获取企业活跃成员数
func GetActiveStat(date string, result *ResultActiveStat) wx.Action {
	params := &ParamsActiveStat{
		Date: date,
	}

	return wx.NewPostAction(urls.CorpUserActiveStat,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsUserID struct {
	Mobile string `json:"mobile"`
}

type ResultUserID struct {
	UserID string `json:"userid"`
}

// GetUserID 通过手机号获取其所对应的userid
func GetUserID(mobile string, result *ResultUserID) wx.Action {
	params := &ParamsUserID{
		Mobile: mobile,
	}

	return wx.NewPostAction(urls.CorpUserGetUserID,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
