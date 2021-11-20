package externalcontact

import (
	"encoding/json"

	"github.com/tidwall/gjson"

	"github.com/shenghui0779/gochat/corp/common"
	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ExternalType int

const (
	ExternalWechat ExternalType = 1 // 外部联系人是微信用户
	ExternalCorp   ExternalType = 2 // 外部联系人是企业微信用户
)

type FollowTagType int

const (
	TagCorpSetting   FollowTagType = 1 // 企业设置
	TagUserCustom    FollowTagType = 2 // 用户自定义
	TagStrategyGroup FollowTagType = 3 // 规则组标签（仅系统应用返回）
)

type Customer struct {
	ExternalContact *ExternalContact `json:"external_contact"`
	FollowUser      []*FollowUser    `json:"follow_user"`
	NextCursor      string           `json:"next_cursor"`
}

type ExternalContact struct {
	ExternalUserID  string                  `json:"external_userid"`
	Name            string                  `json:"name"`
	Position        string                  `json:"position"`
	Avatar          string                  `json:"avatar"`
	CorpName        string                  `json:"corp_name"`
	CorpFullName    string                  `json:"corp_full_name"`
	Type            ExternalType            `json:"type"`
	Gender          int                     `json:"gender"`
	UnionID         string                  `json:"unionid"`
	ExternalProfile *common.ExternalProfile `json:"external_profile"`
}

type FollowUser struct {
	UserID         string       `json:"userid"`
	Remark         string       `json:"remark"`
	Description    string       `json:"description"`
	CreateTime     int64        `json:"create_time"`
	RemarkCorpName string       `json:"remark_corp_name"`
	RemarkMobiles  []string     `json:"remark_mobiles"`
	OperUserID     string       `json:"oper_userid"`
	AddWay         int          `json:"add_way"`
	State          string       `json:"state"`
	Tags           []*FollowTag `json:"tags"`
}

type FollowTag struct {
	TagID     string
	TagName   string
	Type      FollowTagType
	GroupName string
}

func GetList(dest *[]string, userID string) wx.Action {
	return wx.NewGetAction(urls.CorpExternalContactList,
		wx.WithQuery("userid", userID),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal([]byte(gjson.GetBytes(resp, "external_userid").Raw), dest)
		}),
	)
}

func Get() {

}
