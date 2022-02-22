package school

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ResultSubscribeQRCode struct {
	QRCodeBig    string `json:"qrcode_big"`
	QRCodeMiddle string `json:"qrcode_middle"`
	QRCodeThumb  string `json:"qrcode_thumb"`
}

func GetSubscribeQRCode(result *ResultSubscribeQRCode) wx.Action {
	return wx.NewGetAction(urls.CorpSchoolGetSubscribeQRCode,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsSubscribeModeSet struct {
	SubscribeMode int `json:"subscribe_mode"`
}

func SetSubscribeMode(mode int) wx.Action {
	params := &ParamsSubscribeModeSet{
		SubscribeMode: mode,
	}

	return wx.NewPostAction(urls.CorpSchoolSetSubscribeMode,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type ResultSubscribeModeGet struct {
	SubscribeMode int `json:"subscribe_mode"`
}

func GetSubscribeMode(result *ResultSubscribeModeGet) wx.Action {
	return wx.NewGetAction(urls.CorpSchoolGetSubscribeMode,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ExternalContact struct {
	ExternalUserID  string           `json:"external_userid"`
	Name            string           `json:"name"`
	ForeignKey      string           `json:"foreign_key"`
	Position        string           `json:"position"`
	Avatar          string           `json:"avatar"`
	CorpName        string           `json:"corp_name"`
	CorpFullName    string           `json:"corp_full_name"`
	Type            int              `json:"type"`
	Gender          int              `json:"gender"`
	UnionID         string           `json:"unionid"`
	IsSubscribe     int              `json:"is_subscribe"`
	SubscriberInfo  *SubscriberInfo  `json:"subscriber_info"`
	ExternalProfile *ExternalProfile `json:"external_profile"`
}

type SubscriberInfo struct {
	TagID         []string `json:"tag_id"`
	RemarkMobiles []string `json:"remark_mobiles"`
	Remark        string   `json:"remark"`
}

type ExternalProfile struct {
	ExternalAttr []*Attr `json:"external_attr"`
}

type Attr struct {
	Type        int        `json:"type"`
	Name        string     `json:"name"`
	Text        *AttrText  `json:"text"`
	Web         *AttrWeb   `json:"web"`
	Miniprogram *AttrMinip `json:"miniprogram"`
}

type AttrText struct {
	Value string `json:"value"`
}

type AttrWeb struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type AttrMinip struct {
	Title    string `json:"title"`
	AppID    string `json:"appid"`
	Pagepath string `json:"pagepath"`
}

type FollowUser struct {
	UserID         string       `json:"userid"`
	Remark         string       `json:"remark"`
	Description    string       `json:"description"`
	CreateTime     int64        `json:"createtime"`
	RemarkCorpName string       `json:"remark_corp_name"`
	RemarkMobiles  []string     `json:"remark_mobiles"`
	State          string       `json:"state"`
	Tags           []*FollowTag `json:"tags"`
}

type FollowTag struct {
	GroupName string `json:"group_name"`
	TagName   string `json:"tag_name"`
	TagID     string `json:"tag_id"`
	Type      int    `json:"type"`
}

type ResultExternalContact struct {
	ExternalContact *ExternalContact `json:"external_contact"`
	FollowUser      []*FollowUser    `json:"follow_user"`
}

func GetExternalContact(externalUserID string, result *ResultExternalContact) wx.Action {
	return wx.NewGetAction(urls.CorpExternalContactGet,
		wx.WithQuery("external_userid", externalUserID),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsOpenIDConvert struct {
	ExternalUserID string `json:"external_userid"`
}

type ResultOpenIDConvert struct {
	OpenID string `json:"openid"`
}

// ConvertToOpenID 外部联系人openid转换
func ConvertToOpenID(userID string, result *ResultOpenIDConvert) wx.Action {
	params := &ParamsOpenIDConvert{
		ExternalUserID: userID,
	}

	return wx.NewPostAction(urls.CorpExternalContactConvertToOpenID,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
