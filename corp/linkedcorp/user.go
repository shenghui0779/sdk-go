package linkedcorp

import (
	"encoding/json"
	"fmt"

	"github.com/shenghui0779/yiigo"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type UserInfo struct {
	UserID     string   `json:"userid"`
	Name       string   `json:"name"`
	Department []string `json:"department"`
	Mobile     string   `json:"mobile"`
	Telephone  string   `json:"telephone"`
	EMail      string   `json:"email"`
	Position   string   `json:"position"`
	CorpID     string   `json:"corpid"`
	ExtAttr    *ExtAttr `json:"extattr"`
}

type ExtAttr struct {
	Attrs []*Attr `json:"attrs"`
}

type Attr struct {
	Type        int        `json:"type"`
	Name        string     `json:"name"`
	Text        *AttrText  `json:"text,omitempty"`
	Web         *AttrWeb   `json:"web,omitempty"`
	Miniprogram *AttrMinip `json:"miniprogram,omitempty"`
}

type AttrText struct {
	Value string `json:"value"`
}

type AttrWeb struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type AttrMinip struct {
	Title    string `json:"title"`
	AppID    string `json:"appid"`
	Pagepath string `json:"pagepath"`
}

type ParamsUserGet struct {
	CorpID string `json:"corp_id"`
	UserID string `json:"user_id"`
}

type ResultUserGet struct {
	UserInfo *UserInfo `json:"user_info"`
}

func GetUser(params *ParamsUserGet, result *ResultUserGet) wx.Action {
	return wx.NewPostAction(urls.CorpLinkedcorpUserGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(yiigo.X{
				"userid": fmt.Sprintf("%s/%s", params.CorpID, params.UserID),
			})
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsUserSimpleList struct {
	LinkedID     string `json:"linked_id"`
	DepartmentID string `json:"department_id"`
	FetchChild   bool   `json:"fetch_child"`
}

type UserSimpleListData struct {
	UserID    string   `json:"userid"`
	Name      string   `json:"name"`
	Deparment []string `json:"deparment"`
	CorpID    string   `json:"corpid"`
}

type ResultUserSimpleList struct {
	UserList []*UserSimpleListData `json:"userlist"`
}

func ListUserSimple(params *ParamsUserSimpleList, result *ResultUserSimpleList) wx.Action {
	body := yiigo.X{
		"department_id": fmt.Sprintf("%s/%s", params.LinkedID, params.DepartmentID),
	}

	if params.FetchChild {
		body["fetch_child"] = 1
	}

	return wx.NewPostAction(urls.CorpLinkedcorpUserSimpleList,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(body)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsUserList struct {
	LinkedID     string `json:"linked_id"`
	DepartmentID string `json:"department_id"`
	FetchChild   bool   `json:"fetch_child"`
}

type ResultUserList struct {
	UserList []*UserInfo `json:"userlist"`
}

func ListUser(params *ParamsUserList, result *ResultUserList) wx.Action {
	body := yiigo.X{
		"department_id": fmt.Sprintf("%s/%s", params.LinkedID, params.DepartmentID),
	}

	if params.FetchChild {
		body["fetch_child"] = 1
	}

	return wx.NewPostAction(urls.CorpLinkedcorpUserList,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(body)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
