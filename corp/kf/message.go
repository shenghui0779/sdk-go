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

type Text struct {
	Content string `json:"content"`
	MenuID  string `json:"menu_id,omitempty"`
}

type Media struct {
	MediaID string `json:"media_id"`
}

type Location struct {
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
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

type Event struct {
	EventType         event.EventType `json:"event_type"`
	OpenKFID          string          `json:"open_kfid,omitempty"`
	ExternalUserID    string          `json:"external_userid,omitempty"`
	Scene             string          `json:"scene,omitempty"`
	SceneParam        string          `json:"scene_param,omitempty"`
	WelcomeCode       string          `json:"welcome_code,omitempty"`
	FailMsgID         string          `json:"fail_msgid,omitempty"`
	FailType          int             `json:"fail_type,omitempty"`
	ServicerUserID    string          `json:"servicer_userid,omitempty"`
	Status            int             `json:"status,omitempty"`
	ChangeType        int             `json:"change_type,omitempty"`
	OldServicerUserID string          `json:"old_servicer_userid,omitempty"`
	NewServicerUserID string          `json:"new_servicer_userid,omitempty"`
	MsgCode           string          `json:"msg_code,omitempty"`
}

type ParamsMsgSync struct {
	Cursor      string `json:"cursor,omitempty"`
	Token       string `json:"token,omitempty"`
	Limit       int    `json:"limit,omitempty"`
	VoiceFormat int    `json:"voice_format,omitempty"`
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
	MsgType        event.MsgType `json:"msgtype"`
	Text           *Text         `json:"text"`
	Image          *Media        `json:"image"`
	Voice          *Media        `json:"voice"`
	Video          *Media        `json:"video"`
	File           *Media        `json:"file"`
	Location       *Location     `json:"location"`
	Link           *Link         `json:"link"`
	BusinessCard   *BusinessCard `json:"business_card"`
	Minip          *Minip        `json:"miniprogram"`
	Menu           *Menu         `json:"msgmenu"`
	Event          *Event        `json:"event"`
}

// SyncMsg 读取消息
func SyncMsg(params *ParamsMsgSync, result *ResultMsgSync) wx.Action {
	return wx.NewPostAction(urls.CorpKFSyncMsg,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ParamsMsgSend struct {
	ToUser   string        `json:"touser"`
	OpenKFID string        `json:"open_kfid"`
	MsgType  event.MsgType `json:"msgtype"`
	Text     *Text         `json:"text,omitempty"`
	Image    *Media        `json:"image,omitempty"`
	Voice    *Media        `json:"voice,omitempty"`
	Video    *Media        `json:"video,omitempty"`
	File     *Media        `json:"file,omitempty"`
	Link     *Link         `json:"link,omitempty"`
	Minip    *Minip        `json:"miniprogram,omitempty"`
	Menu     *Menu         `json:"msgmenu,omitempty"`
	Location *Location     `json:"location,omitempty"`
}

type ResultMsgSend struct {
	MsgID string `json:"msgid"`
}

// SendTextMsg 发送文本消息
func SendTextMsg(toUser, openKFID, content string, result *ResultMsgSend) wx.Action {
	params := &ParamsMsgSend{
		ToUser:   toUser,
		OpenKFID: openKFID,
		MsgType:  event.MsgText,
		Text: &Text{
			Content: content,
		},
	}

	return wx.NewPostAction(urls.CorpKFSendMsg,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// SendImageMsg 发送图片消息
func SendImageMsg(toUser, openKFID, mediaID string, result *ResultMsgSend) wx.Action {
	params := &ParamsMsgSend{
		ToUser:   toUser,
		OpenKFID: openKFID,
		MsgType:  event.MsgImage,
		Image: &Media{
			MediaID: mediaID,
		},
	}

	return wx.NewPostAction(urls.CorpKFSendMsg,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// SendVoiceMsg 发送语音消息
func SendVoiceMsg(toUser, openKFID, mediaID string, result *ResultMsgSend) wx.Action {
	params := &ParamsMsgSend{
		ToUser:   toUser,
		OpenKFID: openKFID,
		MsgType:  event.MsgVoice,
		Voice: &Media{
			MediaID: mediaID,
		},
	}

	return wx.NewPostAction(urls.CorpKFSendMsg,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// SendVideoMsg 发送视频消息
func SendVideoMsg(toUser, openKFID, mediaID string, result *ResultMsgSend) wx.Action {
	params := &ParamsMsgSend{
		ToUser:   toUser,
		OpenKFID: openKFID,
		MsgType:  event.MsgVideo,
		Video: &Media{
			MediaID: mediaID,
		},
	}

	return wx.NewPostAction(urls.CorpKFSendMsg,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// SendFileMsg 发送文件消息
func SendFileMsg(toUser, openKFID, mediaID string, result *ResultMsgSend) wx.Action {
	params := &ParamsMsgSend{
		ToUser:   toUser,
		OpenKFID: openKFID,
		MsgType:  event.MsgFile,
		File: &Media{
			MediaID: mediaID,
		},
	}

	return wx.NewPostAction(urls.CorpKFSendMsg,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// SendLinkMsg 发送图文链接消息
func SendLinkMsg(toUser, openKFID string, link *Link, result *ResultMsgSend) wx.Action {
	params := &ParamsMsgSend{
		ToUser:   toUser,
		OpenKFID: openKFID,
		MsgType:  event.MsgLink,
		Link:     link,
	}

	return wx.NewPostAction(urls.CorpKFSendMsg,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// SendMinipMsg 发送小程序消息
func SendMinipMsg(toUser, openKFID string, minip *Minip, result *ResultMsgSend) wx.Action {
	params := &ParamsMsgSend{
		ToUser:   toUser,
		OpenKFID: openKFID,
		MsgType:  event.MsgMinip,
		Minip:    minip,
	}

	return wx.NewPostAction(urls.CorpKFSendMsg,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// SendMenuMsg 发送菜单消息
func SendMenuMsg(toUser, openKFID string, menu *Menu, result *ResultMsgSend) wx.Action {
	params := &ParamsMsgSend{
		ToUser:   toUser,
		OpenKFID: openKFID,
		MsgType:  event.MsgMsgMenu,
		Menu:     menu,
	}

	return wx.NewPostAction(urls.CorpKFSendMsg,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// SendLocationMsg 发送地理位置消息
func SendLocationMsg(toUser, openKFID string, location *Location, result *ResultMsgSend) wx.Action {
	params := &ParamsMsgSend{
		ToUser:   toUser,
		OpenKFID: openKFID,
		MsgType:  event.MsgLocation,
		Location: location,
	}

	return wx.NewPostAction(urls.CorpKFSendMsg,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ParamsMsgOnEvent struct {
	Code    string        `json:"code"`
	MsgType event.MsgType `json:"msgtype"`
	Text    *Text         `json:"text,omitempty"`
	Menu    *Menu         `json:"msgmenu,omitempty"`
}

// SendTextMsgOnEvent 发送欢迎语等事件响应消息（文本消息）
func SendTextMsgOnEvent(code, content string, result *ResultMsgSend) wx.Action {
	params := &ParamsMsgOnEvent{
		Code:    code,
		MsgType: event.MsgText,
		Text: &Text{
			Content: content,
		},
	}

	return wx.NewPostAction(urls.CorpKFSendMsgOnEvent,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// SendMenuMsgOnEvent 发送欢迎语等事件响应消息（菜单消息）
func SendMenuMsgOnEvent(code string, menu *Menu, result *ResultMsgSend) wx.Action {
	params := &ParamsMsgOnEvent{
		Code:    code,
		MsgType: event.MsgMsgMenu,
		Menu:    menu,
	}

	return wx.NewPostAction(urls.CorpKFSendMsgOnEvent,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}
