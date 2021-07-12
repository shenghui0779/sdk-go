package corp

import (
	"encoding/json"

	"github.com/shenghui0779/yiigo"

	"github.com/shenghui0779/gochat/wx"
)

// Gender 性别
type Gender string

const (
	GenderUnknown Gender = "0" // 未知
	GenderMale    Gender = "1" // 男性
	GenderFemale  Gender = "2" // 女性
)

type UserAttrType int

const (
	UserTextAttr  UserAttrType = 0
	UserWebAttr   UserAttrType = 1
	UserMinipAttr UserAttrType = 2
)

type UserAttr struct {
	Type        UserAttrType      `json:"type"`
	Name        string            `json:"name"`
	Text        map[string]string `json:"text,omitempty"`
	Web         map[string]string `json:"web,omitempty"`
	Miniprogram map[string]string `json:"miniprogram,omitempty"`
}

type UserExtAttr struct {
	Attrs []*UserAttr `json:"attrs"`
}

type UserExternalProfile struct {
	ExternalCorpName string      `json:"external_corp_name"`
	ExternalAttr     []*UserAttr `json:"external_attr"`
}

type User struct {
	UserID           string               `json:"user_id"`
	Name             string               `json:"name"`
	Alias            string               `json:"alias,omitempty"`
	Mobile           string               `json:"mobile,omitempty"`
	Department       []int                `json:"department"`
	Order            []int                `json:"order,omitempty"`
	Position         string               `json:"position,omitempty"`
	Gender           Gender               `json:"gender,omitempty"`
	Email            string               `json:"email,omitempty"`
	IsLeaderInDept   []int                `json:"is_leader_in_dept,omitempty"`
	Enable           int                  `json:"enable,omitempty"`
	AvatarMediaID    string               `json:"avatar_mediaid,omitempty"` // 仅创建/更新可见
	Avatar           string               `json:"avatar,omitempty"`         // 仅详情可见
	ThumbAvatar      string               `json:"thumb_avatar,omitempty"`   // 仅详情可见
	Telephone        string               `json:"telephone,omitempty"`
	Address          string               `json:"address,omitempty"`
	OpenUserID       string               `json:"open_userid,omitempty"` // 仅详情可见
	MainDepartment   int                  `json:"main_department,omitempty"`
	ExtAttr          *UserExtAttr         `json:"extattr,omitempty"`
	ToInvite         bool                 `json:"to_invite,omitempty"` // 仅创建/更新可见
	Status           int                  `json:"status,omitempty"`    // 仅详情可见
	QRCode           string               `json:"qr_code,omitempty"`   // 仅详情可见
	ExternalPosition string               `json:"external_position,omitempty"`
	ExternalProfile  *UserExternalProfile `json:"external_profile,omitempty"`
}

func CreateUser(data *User) wx.Action {
	return wx.NewAction(UserCreateURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(data)
		}),
	)
}

func GetUser(dest *User, userID string) wx.Action {
	return wx.NewAction(UserGetURL,
		wx.WithMethod(wx.MethodGet),
		wx.WithQuery("userid", userID),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

func UpdateUser(data *User) wx.Action {
	return wx.NewAction(UserUpdateURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(data)
		}),
	)
}

func DeleteUser(userID string) wx.Action {
	return wx.NewAction(UserDeleteURL,
		wx.WithMethod(wx.MethodGet),
		wx.WithQuery("userid", userID),
	)
}

func BatchDeleteUser(userIDs ...string) wx.Action {
	return wx.NewAction(UserBatchDeleteURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(yiigo.X{
				"useridlist": userIDs,
			})
		}),
	)
}
