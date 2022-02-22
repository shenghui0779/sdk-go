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

type WorkbenchKeyData struct {
	Items []*KeyDataItem `json:"items"`
}

type KeyDataItem struct {
	Key      string `json:"key"`
	Data     string `json:"data"`
	JumpURL  string `json:"jump_url"`
	PagePath string `json:"pagepath"`
}

type WorkbenchImage struct {
	URL      string `json:"url"`
	JumpURL  string `json:"jump_url"`
	PagePath string `json:"pagepath"`
}

type WorkbenchList struct {
	Items []*ListItem `json:"items"`
}

type ListItem struct {
	Title    string `json:"title"`
	JumpURL  string `json:"jump_url"`
	PagePath string `json:"pagepath"`
}

type WorkbenchWebView struct {
	URL      string `json:"url"`
	JumpURL  string `json:"jump_url"`
	PagePath string `json:"pagepath"`
}

type ParamsWorkbenchTemplateSet struct {
	AgentID         int64             `json:"agentid"`
	Type            TplType           `json:"type"`
	KeyData         *WorkbenchKeyData `json:"keydata,omitempty"`
	Image           *WorkbenchImage   `json:"image,omitempty"`
	List            *WorkbenchList    `json:"list,omitempty"`
	WebView         *WorkbenchWebView `json:"webview,omitempty"`
	ReplaceUserData bool              `json:"replace_user_data"`
}

// SetWorkbenchNormalTemplate 设置应用在工作台展示的模版（从自定义模式切换为普通宫格或者列表展示模式）
func SetWorkbenchNormalTemplate(agentID int64, replaceUserData bool) wx.Action {
	params := &ParamsWorkbenchTemplateSet{
		AgentID:         agentID,
		Type:            TplNormal,
		ReplaceUserData: replaceUserData,
	}

	return wx.NewPostAction(urls.CorpSetWorkbenchTemplate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

// SetWorkbenchKeyDataTemplate 设置应用在工作台展示的模版（关键数据型）
func SetWorkbenchKeyDataTemplate(agentID int64, keydata *WorkbenchKeyData, replaceUserData bool) wx.Action {
	params := &ParamsWorkbenchTemplateSet{
		AgentID:         agentID,
		Type:            TplKeyData,
		KeyData:         keydata,
		ReplaceUserData: replaceUserData,
	}

	return wx.NewPostAction(urls.CorpSetWorkbenchTemplate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

// SetWorkbenchImageTemplate 设置应用在工作台展示的模版（图片型）
func SetWorkbenchImageTemplate(agentID int64, image *WorkbenchImage, replaceUserData bool) wx.Action {
	params := &ParamsWorkbenchTemplateSet{
		AgentID:         agentID,
		Type:            TplImage,
		Image:           image,
		ReplaceUserData: replaceUserData,
	}

	return wx.NewPostAction(urls.CorpSetWorkbenchTemplate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

// SetWorkbenchListTemplate 设置应用在工作台展示的模版（列表型）
func SetWorkbenchListTemplate(agentID int64, list *WorkbenchList, replaceUserData bool) wx.Action {
	params := &ParamsWorkbenchTemplateSet{
		AgentID:         agentID,
		Type:            TplList,
		List:            list,
		ReplaceUserData: replaceUserData,
	}
	return wx.NewPostAction(urls.CorpSetWorkbenchTemplate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

// SetWorkbenchWebViewTemplate 设置应用在工作台展示的模版（webview型）
func SetWorkbenchWebViewTemplate(agentID int64, webview *WorkbenchWebView, replaceUserData bool) wx.Action {
	params := &ParamsWorkbenchTemplateSet{
		AgentID:         agentID,
		Type:            TplWebView,
		WebView:         webview,
		ReplaceUserData: replaceUserData,
	}
	return wx.NewPostAction(urls.CorpSetWorkbenchTemplate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type ParamsWorkbenchTemplateGet struct {
	AgentID int64 `json:"agentid"`
}

type ResultWorkbenchTemplateGet struct {
	Type            TplType           `json:"type"`
	KeyData         *WorkbenchKeyData `json:"keydata"`
	Image           *WorkbenchImage   `json:"image"`
	List            *WorkbenchList    `json:"list"`
	WebView         *WorkbenchWebView `json:"webview"`
	ReplaceUserData bool              `json:"replace_user_data"`
}

// GetWorkbenchTemplate 获取应用在工作台展示的模版
func GetWorkbenchTemplate(agentID int64, result *ResultWorkbenchTemplateGet) wx.Action {
	params := &ParamsWorkbenchTemplateGet{
		AgentID: agentID,
	}

	return wx.NewPostAction(urls.CorpGetWorkbenchTemplate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsWorkbenchDataSet struct {
	AgentID int64             `json:"agentid"`
	UserID  string            `json:"userid"`
	Type    TplType           `json:"type"`
	KeyData *WorkbenchKeyData `json:"keydata,omitempty"`
	Image   *WorkbenchImage   `json:"image,omitempty"`
	List    *WorkbenchList    `json:"list,omitempty"`
	WebView *WorkbenchWebView `json:"webview,omitempty"`
}

// SetWorkbenchKeyData 设置应用在用户工作台展示的数据（关键数据型）
func SetWorkbenchKeyData(agentID int64, userID string, keydata *WorkbenchKeyData) wx.Action {
	params := &ParamsWorkbenchDataSet{
		AgentID: agentID,
		UserID:  userID,
		Type:    TplKeyData,
		KeyData: keydata,
	}

	return wx.NewPostAction(urls.CorpSetWorkbenchData,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

// SetWorkbenchImageData 设置应用在用户工作台展示的数据（图片型）
func SetWorkbenchImageData(agentID int64, userID string, image *WorkbenchImage) wx.Action {
	params := &ParamsWorkbenchDataSet{
		AgentID: agentID,
		UserID:  userID,
		Type:    TplImage,
		Image:   image,
	}

	return wx.NewPostAction(urls.CorpSetWorkbenchData,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

// SetWorkbenchListData 设置应用在用户工作台展示的数据（列表型）
func SetWorkbenchListData(agentID int64, userID string, list *WorkbenchList) wx.Action {
	params := &ParamsWorkbenchDataSet{
		AgentID: agentID,
		UserID:  userID,
		Type:    TplList,
		List:    list,
	}

	return wx.NewPostAction(urls.CorpSetWorkbenchData,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

// SetWorkbenchWebViewData 设置应用在用户工作台展示的数据（webview型）
func SetWorkbenchWebViewData(agentID int64, userID string, webview *WorkbenchWebView) wx.Action {
	params := &ParamsWorkbenchDataSet{
		AgentID: agentID,
		UserID:  userID,
		Type:    TplWebView,
		WebView: webview,
	}

	return wx.NewPostAction(urls.CorpSetWorkbenchData,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}
