package linkedcorp

import (
	"encoding/json"
	"fmt"

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
	Name        string     `json:"name"`
	Type        int        `json:"type"`
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
	UserID string `json:"userid"`
}

type ResultUserGet struct {
	UserInfo *UserInfo `json:"user_info"`
}

func GetUser(corpID, userID string, result *ResultUserGet) wx.Action {
	params := &ParamsUserGet{
		UserID: fmt.Sprintf("%s/%s", corpID, userID),
	}

	return wx.NewPostAction(urls.CorpLinkedcorpUserGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsSimpleUserList struct {
	DepartmentID string `json:"department_id"`
	FetchChild   bool   `json:"fetch_child"`
}

type SimpleUser struct {
	UserID     string   `json:"userid"`
	Name       string   `json:"name"`
	Department []string `json:"department"`
	CorpID     string   `json:"corpid"`
}

type ResultSimpleUserList struct {
	UserList []*SimpleUser `json:"userlist"`
}

func ListSimpleUser(linkedID, departmentID string, fetchChild bool, result *ResultSimpleUserList) wx.Action {
	params := &ParamsSimpleUserList{
		DepartmentID: fmt.Sprintf("%s/%s", linkedID, departmentID),
		FetchChild:   fetchChild,
	}

	return wx.NewPostAction(urls.CorpLinkedcorpUserSimpleList,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsUserList struct {
	DepartmentID string `json:"department_id"`
	FetchChild   bool   `json:"fetch_child"`
}

type ResultUserList struct {
	UserList []*UserInfo `json:"userlist"`
}

func ListUser(linkedID, departmentID string, fetchChild bool, result *ResultUserList) wx.Action {
	params := &ParamsUserList{
		DepartmentID: fmt.Sprintf("%s/%s", linkedID, departmentID),
		FetchChild:   fetchChild,
	}

	return wx.NewPostAction(urls.CorpLinkedcorpUserList,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
