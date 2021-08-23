package kf

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/event"
	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type MenuType string

const (
	MenuClick MenuType = "click"
	MenuView  MenuType = "view"
	MenuMinip MenuType = "miniprogram"
)

type ClickMenu struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

type ViewMenu struct {
	URL     string `json:"url"`
	Content string `json:"content"`
}

type MinipMenu struct {
	AppID    string `json:"appid"`
	PagePath string `json:"pagepath"`
	Content  string `json:"content"`
}

type TextMessage struct {
	Content string `json:"content"`
	MenuID  string `json:"menu_id,omitempty"`
}

type ImageMessage struct {
	MediaID string `json:"media_id"`
}

type VoiceMessage struct {
	MediaID string `json:"media_id"`
}

type VideoMessage struct {
	MediaID string `json:"media_id"`
}

type FileMessage struct {
	MediaID string `json:"media_id"`
}

type LocationMessage struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Name      string  `json:"name"`
	Address   string  `json:"address"`
}

type LinkMessage struct {
	Title        string `json:"title"`
	Desc         string `json:"desc"`
	URL          string `json:"url"`
	PicURL       string `json:"pic_url,omitempty"`
	ThumbMediaID string `json:"thumb_media_id,omitempty"`
}

type BusinessCardMessage struct {
	UserID string `json:"userid"`
}

type MinipMessage struct {
	Title        string `json:"title"`
	AppID        string `json:"appid"`
	PagePath     string `json:"pagepath"`
	ThumbMediaID string `json:"thumb_media_id"`
}

type MenuMessage struct {
	HeadContent string      `json:"head_content"`
	TailContent string      `json:"tail_content"`
	List        []*MenuItem `json:"list"`
}

type MenuItem struct {
	Type  MenuType   `json:"type"`
	Click *ClickMenu `json:"click,omitempty"`
	View  *ViewMenu  `json:"view,omitempty"`
	Minip *MinipMenu `json:"miniprogram,omitempty"`
}

type EventMessage struct {
	EventType         event.EventType `json:"event_type"`
	OpenKFID          string          `json:"open_kfid,omitempty"`
	ExternalUserID    string          `json:"external_userid,omitempty"`
	FailMsgID         string          `json:"fail_msgid,omitempty"`
	FailType          int             `json:"fail_type,omitempty"`
	ServicerUserID    string          `json:"servicer_userid,omitempty"`
	Status            int             `json:"status,omitempty"`
	ChangeType        int             `json:"change_type,omitempty"`
	OldServicerUserID string          `json:"old_servicer_userid,omitempty"`
	NewServicerUserID string          `json:"new_servicer_userid,omitempty"`
}

type ParamsSyncMsg struct {
	Cursor string `json:"cursor"`
	Token  string `json:"token"`
	Limit  int    `json:"limit"`
}

type ResultSyncMsg struct {
	NextCursor string         `json:"next_cursor"`
	HasMore    int            `json:"has_more"`
	MsgList    []*MsgListItem `json:"msg_list"`
}

type MsgListItem struct {
	MsgID          string               `json:"msgid"`
	OpenKFID       string               `json:"open_kfid"`
	ExternalUserID string               `json:"external_userid"`
	SendTime       int64                `json:"send_time"`
	Origin         int                  `json:"origin"`
	ServicerUserID string               `json:"servicer_userid"`
	MsgType        event.MsgType        `json:"msg_type"`
	Text           *TextMessage         `json:"text"`
	Image          *ImageMessage        `json:"image"`
	Voice          *VoiceMessage        `json:"voice"`
	Video          *VideoMessage        `json:"video"`
	File           *FileMessage         `json:"file"`
	Location       *LocationMessage     `json:"location"`
	Link           *LinkMessage         `json:"link"`
	BussinessCard  *BusinessCardMessage `json:"bussiness_card"`
	Minip          *MinipMessage        `json:"miniprogram"`
	Menu           *MenuMessage         `json:"menu"`
	Event          *EventMessage        `json:"event"`
}

func SyncMsg(params *ParamsSyncMsg, result *ResultSyncMsg) wx.Action {
	return wx.NewPostAction(urls.CorpKFSyncMsg,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsSendMsg struct {
	ToUser   string           `json:"touser"`
	OpenKFID string           `json:"open_kfid"`
	MsgID    string           `json:"msgid,omitempty"`
	MsgType  event.MsgType    `json:"msgtype"`
	Text     *TextMessage     `json:"text,omitempty"`
	Image    *ImageMessage    `json:"image,omitempty"`
	Voice    *VoiceMessage    `json:"voice,omitempty"`
	Video    *VideoMessage    `json:"video,omitempty"`
	File     *FileMessage     `json:"file,omitempty"`
	Link     *LinkMessage     `json:"link,omitempty"`
	Minip    *MinipMessage    `json:"miniprogram,omitempty"`
	Menu     *MenuMessage     `json:"menu,omitempty"`
	Location *LocationMessage `json:"location,omitempty"`
}

type ResultSendMsg struct {
	MsgID string `json:"msgid"`
}

func SendMsg(params *ParamsSendMsg, result ResultSendMsg) wx.Action {
	return wx.NewPostAction(urls.CorpKFSendMsg,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
