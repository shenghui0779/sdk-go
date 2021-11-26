package message

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsAppchatCreate struct {
	Name     string   `json:"name"`
	Owner    string   `json:"owner"`
	UserList []string `json:"userlist"`
	ChatID   string   `json:"chatid,omitempty"`
}

type ResultAppchartCreate struct {
	ChatID string `json:"chatid"`
}

func CreateAppchat(params *ParamsAppchatCreate, result *ResultAppchartCreate) wx.Action {
	return wx.NewPostAction(urls.CorpAppchatCreate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsAppchatUpdate struct {
	ChartID     string   `json:"chartid"`
	Name        string   `json:"name,omitempty"`
	Owner       string   `json:"owner,omitempty"`
	AddUserList []string `json:"add_user_list,omitempty"`
	DelUserList []string `json:"del_user_list,omitempty"`
}

func UpdateAppchat(params *ParamsAppchatUpdate) wx.Action {
	return wx.NewPostAction(urls.CorpAppchatUpdate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ResultAppchatGet struct {
	ChatInfo *AppchatInfo `json:"chat_info"`
}

type AppchatInfo struct {
	ChatID   string   `json:"chatid"`
	Name     string   `json:"name"`
	Owner    string   `json:"owner"`
	UserList []string `json:"userlist"`
}

func GetAppchat(chatID string, result *ResultAppchatGet) wx.Action {
	return wx.NewGetAction(urls.CorpAppchatGet,
		wx.WithQuery("chatid", chatID),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsAppchatSend struct {
	ChatID   string           `json:"chatid"`
	MsgType  string           `json:"msgtype"`
	Text     *TextMessage     `json:"text,omitempty"`
	Image    *ImageMessage    `json:"image,omitempty"`
	Voice    *VoiceMessage    `json:"voice,omitempty"`
	Video    *VideoMessage    `json:"video,omitempty"`
	File     *FileMessage     `json:"file,omitempty"`
	TextCard *TextCardMessage `json:"textcard,omitempty"`
	News     *NewsMessage     `json:"news,omitempty"`
	MPNews   *MPNewsMessage   `json:"mpnews,omitempty"`
	Markdown *MarkdownMessage `json:"markdown,omitempty"`
	Safe     int              `json:"safe"`
}

func SendAppchat(params *ParamsAppchatSend) wx.Action {
	return wx.NewPostAction(urls.CorpAppchatSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}
