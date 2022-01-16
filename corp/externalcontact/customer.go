package externalcontact

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ExternalContact struct {
	ExternalUserID  string           `json:"external_userid"`
	Name            string           `json:"name"`
	Position        string           `json:"position"`
	Avatar          string           `json:"avatar"`
	CorpName        string           `json:"corp_name"`
	CorpFullName    string           `json:"corp_full_name"`
	Type            int              `json:"type"`
	Gender          int              `json:"gender"`
	UnionID         string           `json:"unionid"`
	ExternalProfile *ExternalProfile `json:"external_profile"`
}

type ExternalProfile struct {
	ExternalCorpName string          `json:"external_corp_name"`
	WechatChannels   *WechatChannels `json:"wechat_channels"`
	ExternalAttr     []*Attr         `json:"external_attr"`
}

type WechatChannels struct {
	Nickname string `json:"nickname"`
	Status   int    `json:"status"`
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
	Title string `json:"title"`
	URL   string `json:"url"`
}

type AttrMinip struct {
	Title    string `json:"title"`
	AppID    string `json:"appid"`
	Pagepath string `json:"pagepath"`
}

type FollowInfo struct {
	UserID         string   `json:"userid"`
	Remark         string   `json:"remark"`
	Description    string   `json:"description"`
	CreateTime     int64    `json:"create_time"`
	TagID          []string `json:"tag_id"`
	RemarkCorpName string   `json:"remark_corp_name"`
	RemarkMobiles  []string `json:"remark_mobiles"`
	State          string   `json:"state"`
	OperUserID     string   `json:"oper_userid"`
	AddWay         int      `json:"add_way"`
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
	GroupName string
	TagName   string
	TagID     string
	Type      int
}

type ResultCustomerList struct {
	ExternalUserID []string `json:"external_userid"`
}

func ListCustomer(userID string, result *ResultCustomerList) wx.Action {
	return wx.NewGetAction(urls.CorpExternalContactCustomerList,
		wx.WithQuery("userid", userID),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsCustomerGet struct {
	ExternalUserID string `json:"external_userid"`
	Cursor         string `json:"cursor"`
}

type ResultCustomerGet struct {
	ExternalContact *ExternalContact `json:"external_contact"`
	FollowUser      []*FollowUser    `json:"follow_user"`
	NextCursor      string           `json:"next_cursor"`
}

func GetCustomer(params *ParamsCustomerGet, result *ResultCustomerGet) wx.Action {
	options := []wx.ActionOption{
		wx.WithQuery("external_userid", params.ExternalUserID),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	}

	if len(params.Cursor) != 0 {
		options = append(options, wx.WithQuery("cursor", params.Cursor))
	}

	return wx.NewGetAction(urls.CorpExternalContactCustomerGet, options...)
}

type ParamsCustomerBatchGetByUser struct {
	UserIDList []string `json:"userid_list"`
	Cursor     string   `json:"cursor,omitempty"`
	Limit      int      `json:"limit,omitempty"`
}

type ResultCustomerBatchGetByUser struct {
	ExternalContactList []*CustomerBatchGetData `json:"external_contact_list"`
	NextCursor          string                  `json:"next_cursor"`
}

type CustomerBatchGetData struct {
	ExternalContact *ExternalContact `json:"external_contact"`
	FollowInfo      *FollowInfo      `json:"follow_info"`
}

func BatchGetCustomerByUser(params *ParamsCustomerBatchGetByUser, result *ResultCustomerBatchGetByUser) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactCustomerBatchGetByUser,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsCustomerRemark struct {
	UserID           string   `json:"userid"`
	ExternalUserID   string   `json:"external_userid"`
	Remark           string   `json:"remark"`
	Description      string   `json:"description"`
	RemarkCompany    string   `json:"remark_company"`
	RemarkMobiles    []string `json:"remark_mobiles"`
	RemarkPicMediaID string   `json:"remark_pic_mediaid"`
}

func RemarkCustomer(params *ParamsCustomerRemark) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactCustomerRemark,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}
