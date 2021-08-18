package external_contact

import (
	"encoding/json"

	"github.com/shenghui0779/yiigo"
	"github.com/tidwall/gjson"

	"github.com/shenghui0779/gochat/wx"
)

type ExternalContactWayType int

const (
	ExternalContactWaySingle ExternalContactWayType = 1
	ExternalContactWayMulti  ExternalContactWayType = 2
)

type ExternalContactWayScene int

const (
	ExternalContactWayMinip  ExternalContactWayScene = 1
	ExternalContactWayQRCode ExternalContactWayScene = 2
)

func GetExternalContactFollowUserList(dest *[]string) wx.Action {
	return wx.NewGetAction(ExternalContactFollowUserListURL,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal([]byte(gjson.GetBytes(resp, "follow_user").Raw), dest)
		}),
	)
}

type ExternalContactWay struct {
	ConfigID      string                      `json:"config_id,omitempty"`
	Type          ExternalContactWayType      `json:"type,omitempty"`
	Scene         ExternalContactWayScene     `json:"scene,omitempty"`
	Style         int                         `json:"style,omitempty"`
	Remark        string                      `json:"remark,omitempty"`
	SkipVerify    bool                        `json:"skip_verify,omitempty"`
	State         string                      `json:"state,omitempty"`
	User          []string                    `json:"user,omitempty"`
	Party         []int64                     `json:"party,omitempty"`
	IsTemp        bool                        `json:"is_temp,omitempty"`
	ExpiresIn     int                         `json:"expires_in,omitempty"`
	ChatExpiresIn int                         `json:"chat_expires_in,omitempty"`
	UnionID       string                      `json:"unionid,omitempty"`
	Conclusions   *ExternalContactConclusions `json:"conclusions,omitempty"`
}

type ExternalContactConclusions struct {
	Text        *ExternalContactTextConclusion  `json:"text,omitempty"`
	Image       *ExternalContactImageConclusion `json:"image,omitempty"`
	Link        *ExternalContactLinkConclusion  `json:"link,omitempty"`
	MiniProgram *ExternalContactMinipConclusion `json:"mini_program,omitempty"`
}

type ExternalContactTextConclusion struct {
	Content string `json:"content"`
}

type ExternalContactImageConclusion struct {
	MediaID string `json:"media_id"`
}

type ExternalContactLinkConclusion struct {
	Title  string `json:"title"`
	PicURL string `json:"picurl"`
	Desc   string `json:"desc"`
	URL    string `json:"url"`
}

type ExternalContactMinipConclusion struct {
	Title      string `json:"title"`
	PicMediaID string `json:"pic_media_id"`
	AppID      string `json:"appid"`
	Page       string `json:"page"`
}

type ExternalContactWayAddResult struct {
	ConfigID string `json:"config_id"`
	QRCode   string `json:"qr_code"`
}

func AddExternalContactWay(dest *ExternalContactWayAddResult, data *ExternalContactWay) wx.Action {
	return wx.NewPostAction(ExternalContactWayAddURL,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(data)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

func GetExternalContactWay(dest *ExternalContactWay, configID string) wx.Action {
	return wx.NewPostAction(ExternalContactWayGetURL,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(yiigo.X{"config_id": configID})
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal([]byte(gjson.GetBytes(resp, "contact_way").Raw), dest)
		}),
	)
}

func UpdateExternalContactWay(data *ExternalContactWay) wx.Action {
	return wx.NewPostAction(ExternalContactWayUpdateURL,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(data)
		}),
	)
}

func DeleteExternalContactWay(configID string) wx.Action {
	return wx.NewPostAction(ExternalContactWayDeleteURL,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(yiigo.X{"config_id": configID})
		}),
	)
}

func CloseExternalContactTempChat(userID, externalUserID string) wx.Action {
	return wx.NewPostAction(ExternalContactCloseTempChatURL,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(yiigo.X{
				"userid":          userID,
				"external_userid": externalUserID,
			})
		}),
	)
}
