package agent

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type TplType string

const (
	TplNormal  TplType = "normal"
	TplKeyData TplType = "keydata"
	TplImage   TplType = "image"
	TplList    TplType = "list"
	TplWebView TplType = "webview"
)

type KeyDataTpl struct {
	Key      string `json:"key"`
	Data     string `json:"data"`
	JumpURL  string `json:"jump_url"`
	PagePath string `json:"pagepath"`
}

type ImageTpl struct {
	URL      string `json:"url"`
	JumpURL  string `json:"jump_url"`
	PagePath string `json:"pagepath"`
}

type ListTpl struct {
	Title    string `json:"title"`
	JumpURL  string `json:"jump_url"`
	PagePath string `json:"pagepath"`
}

type WebViewTpl struct {
	URL      string `json:"url"`
	JumpURL  string `json:"jump_url"`
	PagePath string `json:"pagepath"`
}

type ParamsSetWorkbenchTemplate struct {
	AgentID         string      `json:"agentid"`
	Type            TplType     `json:"type"`
	KeyData         *KeyDataTpl `json:"keydata,omitempty"`
	Image           *ImageTpl   `json:"image,omitempty"`
	List            *ListTpl    `json:"list,omitempty"`
	WebView         *WebViewTpl `json:"webview,omitempty"`
	ReplaceUserData bool        `json:"replace_user_data"`
}

func SetWorkbenchTemplate(params *ParamsSetWorkbenchTemplate) wx.Action {
	return wx.NewPostAction(urls.CorpGetWorkbenchTemplate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsGetWorkbenchTemplate struct {
	AgentID string `json:"agentid"`
}

type ResultGetWorkbenchTemplate struct {
	Type            TplType     `json:"type"`
	KeyData         *KeyDataTpl `json:"keydata,omitempty"`
	Image           *ImageTpl   `json:"image,omitempty"`
	List            *ListTpl    `json:"list,omitempty"`
	WebView         *WebViewTpl `json:"webview,omitempty"`
	ReplaceUserData bool        `json:"replace_user_data"`
}

func GetWorkbenchTemplate(params *ParamsGetWorkbenchTemplate, result *ResultGetWorkbenchTemplate) wx.Action {
	return wx.NewPostAction(urls.CorpGetWorkbenchTemplate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsSetWorkbenchData struct {
	AgentID string      `json:"agentid"`
	Type    TplType     `json:"type"`
	UserID  string      `json:"userid"`
	KeyData *KeyDataTpl `json:"keydata,omitempty"`
	Image   *ImageTpl   `json:"image,omitempty"`
	List    *ListTpl    `json:"list,omitempty"`
	WebView *WebViewTpl `json:"webview,omitempty"`
}

func SetWorkbenchData(params *ParamsSetWorkbenchData) wx.Action {
	return wx.NewPostAction(urls.CorpSetWorkbenchData,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}
