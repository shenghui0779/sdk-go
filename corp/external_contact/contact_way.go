package external_contact

import (
	"encoding/json"

	"github.com/shenghui0779/yiigo"
	"github.com/tidwall/gjson"

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

func GetFollowUserList(dest *[]string) wx.Action {
	return wx.NewGetAction(urls.CorpExternalContactFollowUserList,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal([]byte(gjson.GetBytes(resp, "follow_user").Raw), dest)
		}),
	)
}

type ContactWay struct {
	ConfigID      string       `json:"config_id,omitempty"`
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

type ContactWayAddResult struct {
	ConfigID string `json:"config_id"`
	QRCode   string `json:"qr_code"`
}

func AddContactWay(dest *ContactWayAddResult, data *ContactWay) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactWayAdd,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(data)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

func GetContactWay(dest *ContactWay, configID string) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactWayGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(yiigo.X{"config_id": configID})
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal([]byte(gjson.GetBytes(resp, "contact_way").Raw), dest)
		}),
	)
}

func UpdateContactWay(data *ContactWay) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactWayUpdate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(data)
		}),
	)
}

func DeleteContactWay(configID string) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactWayDelete,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(yiigo.X{"config_id": configID})
		}),
	)
}

func CloseTempChat(userID, externalUserID string) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactCloseTempChat,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(yiigo.X{
				"userid":          userID,
				"external_userid": externalUserID,
			})
		}),
	)
}
