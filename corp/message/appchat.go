package message

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/event"
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

type AppchatMsg struct {
	ChatID   string        `json:"chatid"`
	MsgType  event.MsgType `json:"msgtype"`
	Text     *Text         `json:"text,omitempty"`
	Image    *Media        `json:"image,omitempty"`
	Voice    *Media        `json:"voice,omitempty"`
	Video    *Video        `json:"video,omitempty"`
	File     *Media        `json:"file,omitempty"`
	TextCard *TextCard     `json:"textcard,omitempty"`
	News     *News         `json:"news,omitempty"`
	MPNews   *MPNews       `json:"mpnews,omitempty"`
	Markdown *Text         `json:"markdown,omitempty"`
	Safe     int           `json:"safe"`
}

func SendAppchatText(chatID string, content string, safe int) wx.Action {
	msg := &AppchatMsg{
		ChatID:  chatID,
		MsgType: event.MsgText,
		Text: &Text{
			Content: content,
		},
		Safe: safe,
	}
	return wx.NewPostAction(urls.CorpAppchatSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(msg)
		}),
	)
}

func SendAppchatImage(chatID string, mediaID string, safe int) wx.Action {
	msg := &AppchatMsg{
		ChatID:  chatID,
		MsgType: event.MsgImage,
		Image: &Media{
			MediaID: mediaID,
		},
		Safe: safe,
	}
	return wx.NewPostAction(urls.CorpAppchatSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(msg)
		}),
	)
}

func SendAppchatVoice(chatID string, mediaID string, safe int) wx.Action {
	msg := &AppchatMsg{
		ChatID:  chatID,
		MsgType: event.MsgVoice,
		Voice: &Media{
			MediaID: mediaID,
		},
		Safe: safe,
	}
	return wx.NewPostAction(urls.CorpAppchatSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(msg)
		}),
	)
}

func SendAppchatVideo(chatID string, video *Video, safe int) wx.Action {
	msg := &AppchatMsg{
		ChatID:  chatID,
		MsgType: event.MsgVideo,
		Video:   video,
		Safe:    safe,
	}
	return wx.NewPostAction(urls.CorpAppchatSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(msg)
		}),
	)
}

func SendAppchatFile(chatID string, mediaID string, safe int) wx.Action {
	msg := &AppchatMsg{
		ChatID:  chatID,
		MsgType: event.MsgFile,
		File: &Media{
			MediaID: mediaID,
		},
		Safe: safe,
	}
	return wx.NewPostAction(urls.CorpAppchatSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(msg)
		}),
	)
}

func SendAppchatTextCard(chatID string, card *TextCard, safe int) wx.Action {
	msg := &AppchatMsg{
		ChatID:   chatID,
		MsgType:  event.MsgTextCard,
		TextCard: card,
		Safe:     safe,
	}
	return wx.NewPostAction(urls.CorpAppchatSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(msg)
		}),
	)
}

func SendAppchatNews(chatID string, articles []*NewsArticle, safe int) wx.Action {
	msg := &AppchatMsg{
		ChatID:  chatID,
		MsgType: event.MsgNews,
		News: &News{
			Articles: articles,
		},
		Safe: safe,
	}
	return wx.NewPostAction(urls.CorpAppchatSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(msg)
		}),
	)
}

func SendAppchatMPNews(chatID string, articles []*MPNewsArticle, safe int) wx.Action {
	msg := &AppchatMsg{
		ChatID:  chatID,
		MsgType: event.MsgMPNews,
		MPNews: &MPNews{
			Articles: articles,
		},
		Safe: safe,
	}
	return wx.NewPostAction(urls.CorpAppchatSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(msg)
		}),
	)
}

func SendAppchatMarkdown(chatID string, content string, safe int) wx.Action {
	msg := &AppchatMsg{
		ChatID:  chatID,
		MsgType: event.MsgMarkdown,
		Markdown: &Text{
			Content: content,
		},
		Safe: safe,
	}
	return wx.NewPostAction(urls.CorpAppchatSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(msg)
		}),
	)
}
