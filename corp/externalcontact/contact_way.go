package externalcontact

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ContactType int

const (
	ContactSingle ContactType = 1
	ContactMulti  ContactType = 2
)

type ContactScene int

const (
	ContactMinip  ContactScene = 1
	ContactQRCode ContactScene = 2
)

type ResultFollowUserList struct {
	FollowUser []string `json:"follow_user"`
}

func ListFollowUser(result *ResultFollowUserList) wx.Action {
	return wx.NewGetAction(urls.CorpExternalContactFollowUserList,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ContactWay struct {
	ConfigID      string       `json:"config_id"`
	Type          ContactType  `json:"type"`
	Scene         ContactScene `json:"scene"`
	Style         int          `json:"style"`
	Remark        string       `json:"remark"`
	SkipVerify    bool         `json:"skip_verify"`
	State         string       `json:"state"`
	User          []string     `json:"user"`
	Party         []int64      `json:"party"`
	IsTemp        bool         `json:"is_temp"`
	ExpiresIn     int          `json:"expires_in"`
	ChatExpiresIn int          `json:"chat_expires_in"`
	UnionID       string       `json:"unionid"`
	Conclusions   *Conclusions `json:"conclusion"`
}

type Conclusions struct {
	Text        *TextConclusion  `json:"text,omitempty"`
	Image       *ImageConclusion `json:"image,omitempty"`
	Link        *LinkConclusion  `json:"link,omitempty"`
	MiniProgram *MinipConclusion `json:"mini_program,omitempty"`
}

type TextConclusion struct {
	Content string `json:"content"`
}

type ImageConclusion struct {
	MediaID string `json:"media_id"`
}

type LinkConclusion struct {
	Title  string `json:"title"`
	PicURL string `json:"picurl"`
	Desc   string `json:"desc"`
	URL    string `json:"url"`
}

type MinipConclusion struct {
	Title      string `json:"title"`
	PicMediaID string `json:"pic_media_id"`
	AppID      string `json:"appid"`
	Page       string `json:"page"`
}

type ParamsContactWayAdd struct {
	Type          ContactType  `json:"type,omitempty"`
	Scene         ContactScene `json:"scene,omitempty"`
	Style         int          `json:"style,omitempty"`
	Remark        string       `json:"remark,omitempty"`
	SkipVerify    bool         `json:"skip_verify,omitempty"`
	State         string       `json:"state,omitempty"`
	User          []string     `json:"user,omitempty"`
	Party         []int64      `json:"party,omitempty"`
	IsTemp        bool         `json:"is_temp,omitempty"`
	ExpiresIn     int          `json:"expires_in,omitempty"`
	ChatExpiresIn int          `json:"chat_expires_in,omitempty"`
	UnionID       string       `json:"unionid,omitempty"`
	Conclusions   *Conclusions `json:"conclusions,omitempty"`
}

type ResultContactWayAdd struct {
	ConfigID string `json:"config_id"`
	QRCode   string `json:"qr_code"`
}

func AddContactWay(params *ParamsContactWayAdd, result *ResultContactWayAdd) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactWayAdd,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsContactWayUpdate struct {
	ConfigID      string       `json:"config_id"`
	Type          ContactType  `json:"type,omitempty"`
	Scene         ContactScene `json:"scene,omitempty"`
	Style         int          `json:"style,omitempty"`
	Remark        string       `json:"remark,omitempty"`
	SkipVerify    bool         `json:"skip_verify,omitempty"`
	State         string       `json:"state,omitempty"`
	User          []string     `json:"user,omitempty"`
	Party         []int64      `json:"party,omitempty"`
	IsTemp        bool         `json:"is_temp,omitempty"`
	ExpiresIn     int          `json:"expires_in,omitempty"`
	ChatExpiresIn int          `json:"chat_expires_in,omitempty"`
	UnionID       string       `json:"unionid,omitempty"`
	Conclusions   *Conclusions `json:"conclusions,omitempty"`
}

func UpdateContactWay(params *ParamsContactWayUpdate) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactWayUpdate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsContactWayGet struct {
	ConfigID string `json:"config_id"`
}

type ResultContactWayGet struct {
	ContactWay *ContactWay `json:"contact_way"`
}

func GetContactWay(params *ParamsContactWayGet, result *ResultContactWayGet) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactWayGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsContactWayList struct {
	StartTime int64  `json:"start_time,omitempty"`
	EntTime   int64  `json:"ent_time,omitempty"`
	Cursor    string `json:"cursor,omitempty"`
	Limit     int    `json:"limit,omitempty"`
}

type ResultContactWayList struct {
	ContactWay []*ContactWayListData `json:"contact_way"`
	NextCursor string                `json:"next_cursor"`
}

type ContactWayListData struct {
	ConfigID string `json:"config_id"`
}

func ListContactWay(params *ParamsContactWayList, result *ResultContactWayList) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactWayList,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsContactWayDelete struct {
	ConfigID string `json:"config_id"`
}

func DeleteContactWay(params *ParamsContactWayDelete) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactWayDelete,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsTempChatClose struct {
	UserID         string `json:"userid"`
	ExternalUserID string `json:"external_userid"`
}

func CloseTempChat(params *ParamsTempChatClose) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactCloseTempChat,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}
