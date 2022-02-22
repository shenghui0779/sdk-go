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
	CreateTime     int64    `json:"createtime"`
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
	CreateTime     int64        `json:"createtime"`
	RemarkCorpName string       `json:"remark_corp_name"`
	RemarkMobiles  []string     `json:"remark_mobiles"`
	OperUserID     string       `json:"oper_userid"`
	AddWay         int          `json:"add_way"`
	State          string       `json:"state"`
	Tags           []*FollowTag `json:"tags"`
}

type FollowTag struct {
	GroupName string `json:"group_name"`
	TagName   string `json:"tag_name"`
	TagID     string `json:"tag_id"`
	Type      int    `json:"type"`
}

type ResultList struct {
	ExternalUserID []string `json:"external_userid"`
}

// List 获取客户列表
func List(userID string, result *ResultList) wx.Action {
	return wx.NewGetAction(urls.CorpExternalContactList,
		wx.WithQuery("userid", userID),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ResultGet struct {
	ExternalContact *ExternalContact `json:"external_contact"`
	FollowUser      []*FollowUser    `json:"follow_user"`
	NextCursor      string           `json:"next_cursor"`
}

// Get 获取客户详情
func Get(externalUserID, cursor string, result *ResultGet) wx.Action {
	options := []wx.ActionOption{
		wx.WithQuery("external_userid", externalUserID),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	}

	if len(cursor) != 0 {
		options = append(options, wx.WithQuery("cursor", cursor))
	}

	return wx.NewGetAction(urls.CorpExternalContactGet, options...)
}

type ParamsBatchGetByUser struct {
	UserIDList []string `json:"userid_list"`
	Cursor     string   `json:"cursor,omitempty"`
	Limit      int      `json:"limit,omitempty"`
}

type ResultBatchGetByUser struct {
	ExternalContactList []*CustomerBatchGetData `json:"external_contact_list"`
	NextCursor          string                  `json:"next_cursor"`
}

type CustomerBatchGetData struct {
	ExternalContact *ExternalContact `json:"external_contact"`
	FollowInfo      *FollowInfo      `json:"follow_info"`
}

// BatchGetByUser 批量获取客户详情
func BatchGetByUser(userIDs []string, cursor string, limit int, result *ResultBatchGetByUser) wx.Action {
	params := &ParamsBatchGetByUser{
		UserIDList: userIDs,
		Cursor:     cursor,
		Limit:      limit,
	}

	return wx.NewPostAction(urls.CorpExternalContactBatchGetByUser,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsRemark struct {
	UserID           string   `json:"userid"`
	ExternalUserID   string   `json:"external_userid"`
	Remark           string   `json:"remark"`
	Description      string   `json:"description"`
	RemarkCompany    string   `json:"remark_company"`
	RemarkMobiles    []string `json:"remark_mobiles"`
	RemarkPicMediaID string   `json:"remark_pic_mediaid"`
}

// Remark 修改客户备注信息
func Remark(params *ParamsRemark) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactRemark,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}
