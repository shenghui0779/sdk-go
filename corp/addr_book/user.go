package addr_book

import (
	"encoding/json"
	"strconv"

	"github.com/shenghui0779/yiigo"
	"github.com/tidwall/gjson"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

// Gender 性别
type Gender string

const (
	GenderUnknown Gender = "0" // 未知
	GenderMale    Gender = "1" // 男性
	GenderFemale  Gender = "2" // 女性
)

type AttrType int

const (
	TextAttr  AttrType = 0
	WebAttr   AttrType = 1
	MinipAttr AttrType = 2
)

type Attr struct {
	Type        AttrType          `json:"type"`
	Name        string            `json:"name"`
	Text        map[string]string `json:"text,omitempty"`
	Web         map[string]string `json:"web,omitempty"`
	Miniprogram map[string]string `json:"miniprogram,omitempty"`
}

type ExtAttr struct {
	Attrs []*Attr `json:"attrs"`
}

type ExternalProfile struct {
	ExternalCorpName string  `json:"external_corp_name"`
	WechatChannels   string  `json:"wechat_channels"`
	ExternalAttr     []*Attr `json:"external_attr"`
}

type User struct {
	UserID           string           `json:"user_id"`
	Name             string           `json:"name"`
	Alias            string           `json:"alias,omitempty"`
	Mobile           string           `json:"mobile,omitempty"`
	Department       []int64          `json:"department"`
	Order            []int64          `json:"order,omitempty"`
	Position         string           `json:"position,omitempty"`
	Gender           Gender           `json:"gender,omitempty"`
	Email            string           `json:"email,omitempty"`
	IsLeaderInDept   []int64          `json:"is_leader_in_dept,omitempty"`
	Enable           int              `json:"enable,omitempty"`
	AvatarMediaID    string           `json:"avatar_mediaid,omitempty"` // 仅创建/更新可见
	Avatar           string           `json:"avatar,omitempty"`         // 仅详情可见
	ThumbAvatar      string           `json:"thumb_avatar,omitempty"`   // 仅详情可见
	Telephone        string           `json:"telephone,omitempty"`
	Address          string           `json:"address,omitempty"`
	OpenUserID       string           `json:"open_userid,omitempty"` // 仅详情可见
	MainDepartment   int64            `json:"main_department,omitempty"`
	ExtAttr          *ExtAttr         `json:"extattr,omitempty"`
	ToInvite         bool             `json:"to_invite,omitempty"` // 仅创建/更新可见
	Status           int              `json:"status,omitempty"`    // 仅详情可见
	QRCode           string           `json:"qr_code,omitempty"`   // 仅详情可见
	ExternalPosition string           `json:"external_position,omitempty"`
	ExternalProfile  *ExternalProfile `json:"external_profile,omitempty"`
}

// CreateUser 创建成员
func CreateUser(data *User) wx.Action {
	return wx.NewPostAction(urls.CorpAddrBookUserCreate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(data)
		}),
	)
}

// GetUser 读取成员
func GetUser(dest *User, userID string) wx.Action {
	return wx.NewGetAction(urls.CorpAddrBookUserGet,
		wx.WithQuery("userid", userID),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// UpdateUser 更新成员
func UpdateUser(data *User) wx.Action {
	return wx.NewPostAction(urls.CorpAddrBookUserUpdate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(data)
		}),
	)
}

// DeleteUser 删除成员
func DeleteUser(userID string) wx.Action {
	return wx.NewGetAction(urls.CorpAddrBookUserDelete,
		wx.WithQuery("userid", userID),
	)
}

// BatchDeleteUser 批量删除成员
func BatchDeleteUser(userIDs ...string) wx.Action {
	return wx.NewPostAction(urls.CorpAddrBookUserBatchDelete,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(yiigo.X{
				"useridlist": userIDs,
			})
		}),
	)
}

// GetUserSimpleList 获取部门成员
func GetUserSimpleList(dest *[]*User, departmentID int64, fetchChild bool) wx.Action {
	child := 0

	if fetchChild {
		child = 1
	}

	return wx.NewGetAction(urls.CorpAddrBookUserSimpleList,
		wx.WithQuery("department_id", strconv.FormatInt(departmentID, 10)),
		wx.WithQuery("fetch_child", strconv.Itoa(child)),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal([]byte(gjson.GetBytes(resp, "userlist").Raw), dest)
		}),
	)
}

// UserList 获取部门成员详情
func UserList(dest *[]*User, departmentID int64, fetchChild bool) wx.Action {
	child := 0

	if fetchChild {
		child = 1
	}

	return wx.NewGetAction(urls.CorpAddrBookUserList,
		wx.WithQuery("department_id", strconv.FormatInt(departmentID, 10)),
		wx.WithQuery("fetch_child", strconv.Itoa(child)),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal([]byte(gjson.GetBytes(resp, "userlist").Raw), dest)
		}),
	)
}

type ConvertResult struct {
	UserID string `json:"userid"`
	OpenID string `json:"openid"`
}

// Convert2OpenID userid转openid
func Convert2OpenID(dest *ConvertResult, userID string) wx.Action {
	return wx.NewPostAction(urls.CorpAddrBookConvert2OpenID,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(yiigo.X{"userid": userID})
		}),
		wx.WithDecode(func(resp []byte) error {
			dest.UserID = userID
			dest.OpenID = gjson.GetBytes(resp, "openid").String()

			return nil
		}),
	)
}

// Convert2UserID openid转userid
func Convert2UserID(dest *ConvertResult, openid string) wx.Action {
	return wx.NewPostAction(urls.CorpAddrBookConvert2UserID,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(yiigo.X{"openid": openid})
		}),
		wx.WithDecode(func(resp []byte) error {
			dest.UserID = gjson.GetBytes(resp, "userid").String()
			dest.OpenID = openid

			return nil
		}),
	)
}

type ParamsInvite struct {
	User  []string `json:"user,omitempty"`  // 成员ID列表，最多支持1000个
	Party []string `json:"party,omitempty"` // 部门ID列表，最多支持100个
	Tag   []string `json:"tag,omitempty"`   // 标签ID列表，最多支持100个
}

type InviteResult struct {
	InvalidUser  []string `json:"invaliduser"`
	InvalidParth []string `json:"invalidparth"`
	InvalidTag   []string `json:"invalidtag"`
}

// BatchInvite 邀请成员
func BatchInvite(dest *InviteResult, params *ParamsInvite) wx.Action {
	return wx.NewPostAction(urls.CorpAddrBookBatchInvite,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
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
func GetJoinQRCode(dest *JoinQRCode, size ...int) wx.Action {
	options := []wx.ActionOption{
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	}

	if len(size) != 0 {
		options = append(options, wx.WithQuery("size_type", strconv.Itoa(size[0])))
	}

	return wx.NewGetAction(urls.CorpAddrBookJoinQRCode, options...)
}

type ActiveStat struct {
	Count int `json:"active_cnt"`
}

// GetActiveStat 获取企业活跃成员数
func GetActiveStat(dest *ActiveStat, date string) wx.Action {
	return wx.NewPostAction(urls.CorpAddrBookActiveStat,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(yiigo.X{
				"date": date,
			})
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}
