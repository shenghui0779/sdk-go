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

type Text struct {
	Content string `json:"content"`
	MenuID  string `json:"menu_id,omitempty"`
}

type Image struct {
	MediaID string `json:"media_id"`
}

type Voice struct {
	MediaID string `json:"media_id"`
}

type Video struct {
	MediaID string `json:"media_id"`
}

type File struct {
	MediaID string `json:"media_id"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Name      string  `json:"name"`
	Address   string  `json:"address"`
}

type Link struct {
	Title        string `json:"title"`
	Desc         string `json:"desc"`
	URL          string `json:"url"`
	PicURL       string `json:"pic_url,omitempty"`
	ThumbMediaID string `json:"thumb_media_id,omitempty"`
}

type BusinessCard struct {
	UserID string `json:"userid"`
}

type Minip struct {
	Title        string `json:"title"`
	AppID        string `json:"appid"`
	PagePath     string `json:"pagepath"`
	ThumbMediaID string `json:"thumb_media_id"`
}

type Menu struct {
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

type Event struct {
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

type ParamsMsgSync struct {
	Cursor string `json:"cursor"`
	Token  string `json:"token"`
	Limit  int    `json:"limit"`
}

type ResultMsgSync struct {
	NextCursor string         `json:"next_cursor"`
	HasMore    int            `json:"has_more"`
	MsgList    []*MsgListData `json:"msg_list"`
}

type MsgListData struct {
	MsgID          string        `json:"msgid"`
	OpenKFID       string        `json:"open_kfid"`
	ExternalUserID string        `json:"external_userid"`
	SendTime       int64         `json:"send_time"`
	Origin         int           `json:"origin"`
	ServicerUserID string        `json:"servicer_userid"`
	MsgType        event.MsgType `json:"msg_type"`
	Text           *Text         `json:"text"`
	Image          *Image        `json:"image"`
	Voice          *Voice        `json:"voice"`
	Video          *Video        `json:"video"`
	File           *File         `json:"file"`
	Location       *Location     `json:"location"`
	Link           *Link         `json:"link"`
	BussinessCard  *BusinessCard `json:"bussiness_card"`
	Minip          *Minip        `json:"miniprogram"`
	Menu           *Menu         `json:"msgmenu"`
	Event          *Event        `json:"event"`
}

func SyncMsg(params *ParamsMsgSync, result *ResultMsgSync) wx.Action {
	return wx.NewPostAction(urls.CorpKFSyncMsg,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsMsgSend struct {
	ToUser   string        `json:"touser"`
	OpenKFID string        `json:"open_kfid"`
	MsgID    string        `json:"msgid,omitempty"`
	MsgType  event.MsgType `json:"msgtype"`
	Text     *Text         `json:"text,omitempty"`
	Image    *Image        `json:"image,omitempty"`
	Voice    *Voice        `json:"voice,omitempty"`
	Video    *Video        `json:"video,omitempty"`
	File     *File         `json:"file,omitempty"`
	Link     *Link         `json:"link,omitempty"`
	Minip    *Minip        `json:"miniprogram,omitempty"`
	Menu     *Menu         `json:"msgmenu,omitempty"`
	Location *Location     `json:"location,omitempty"`
}

type ResultMsgSend struct {
	MsgID string `json:"msgid"`
}

func SendTextMsg(toUser, openKFID, msgID string, text *Text, result *ResultMsgSend) wx.Action {
	params := &ParamsMsgSend{
		ToUser:   toUser,
		OpenKFID: openKFID,
		MsgID:    msgID,
		MsgType:  event.MsgText,
		Text:     text,
	}

	return wx.NewPostAction(urls.CorpKFSendMsg,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendImageMsg(toUser, openKFID, msgID string, image *Image, result *ResultMsgSend) wx.Action {
	params := &ParamsMsgSend{
		ToUser:   toUser,
		OpenKFID: openKFID,
		MsgID:    msgID,
		MsgType:  event.MsgImage,
		Image:    image,
	}

	return wx.NewPostAction(urls.CorpKFSendMsg,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendVoiceMsg(toUser, openKFID, msgID string, voice *Voice, result *ResultMsgSend) wx.Action {
	params := &ParamsMsgSend{
		ToUser:   toUser,
		OpenKFID: openKFID,
		MsgID:    msgID,
		MsgType:  event.MsgText,
		Voice:    voice,
	}

	return wx.NewPostAction(urls.CorpKFSendMsg,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendVideoMsg(toUser, openKFID, msgID string, video *Video, result *ResultMsgSend) wx.Action {
	params := &ParamsMsgSend{
		ToUser:   toUser,
		OpenKFID: openKFID,
		MsgID:    msgID,
		MsgType:  event.MsgText,
		Video:    video,
	}

	return wx.NewPostAction(urls.CorpKFSendMsg,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendFileMsg(toUser, openKFID, msgID string, file *File, result *ResultMsgSend) wx.Action {
	params := &ParamsMsgSend{
		ToUser:   toUser,
		OpenKFID: openKFID,
		MsgID:    msgID,
		MsgType:  event.MsgText,
		File:     file,
	}

	return wx.NewPostAction(urls.CorpKFSendMsg,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendLinkMsg(toUser, openKFID, msgID string, link *Link, result *ResultMsgSend) wx.Action {
	params := &ParamsMsgSend{
		ToUser:   toUser,
		OpenKFID: openKFID,
		MsgID:    msgID,
		MsgType:  event.MsgText,
		Link:     link,
	}

	return wx.NewPostAction(urls.CorpKFSendMsg,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendMinipMsg(toUser, openKFID, msgID string, minip *Minip, result *ResultMsgSend) wx.Action {
	params := &ParamsMsgSend{
		ToUser:   toUser,
		OpenKFID: openKFID,
		MsgID:    msgID,
		MsgType:  event.MsgText,
		Minip:    minip,
	}

	return wx.NewPostAction(urls.CorpKFSendMsg,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendMenuMsg(toUser, openKFID, msgID string, menu *Menu, result *ResultMsgSend) wx.Action {
	params := &ParamsMsgSend{
		ToUser:   toUser,
		OpenKFID: openKFID,
		MsgID:    msgID,
		MsgType:  event.MsgMsgMenu,
		Menu:     menu,
	}

	return wx.NewPostAction(urls.CorpKFSendMsg,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendLocationMsg(toUser, openKFID, msgID string, location *Location, result *ResultMsgSend) wx.Action {
	params := &ParamsMsgSend{
		ToUser:   toUser,
		OpenKFID: openKFID,
		MsgID:    msgID,
		MsgType:  event.MsgLocation,
		Location: location,
	}

	return wx.NewPostAction(urls.CorpKFSendMsg,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
