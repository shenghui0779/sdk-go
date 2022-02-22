package user

import (
	"encoding/json"
	"strconv"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsTagCreate struct {
	TagName string `json:"tagname"`
}

type ResultTagCreate struct {
	TagID int64 `json:"tagid"`
}

// CreateTag 创建标签
func CreateTag(name string, result *ResultTagCreate) wx.Action {
	params := &ParamsTagCreate{
		TagName: name,
	}

	return wx.NewPostAction(urls.CorpUserTagCreate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsTagUpdate struct {
	TagID   int64  `json:"tagid"`
	TagName string `json:"tagname"`
}

func UpdateTag(tagID int64, tagName string) wx.Action {
	params := &ParamsTagUpdate{
		TagID:   tagID,
		TagName: tagName,
	}

	return wx.NewPostAction(urls.CorpUserTagUpdate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

func DeleteTag(tagID int64) wx.Action {
	return wx.NewGetAction(urls.CorpUserTagDelete,
		wx.WithQuery("tagid", strconv.FormatInt(tagID, 10)),
	)
}

type ResultTagUser struct {
	TagName   string     `json:"tagname"`
	UserList  []*TagUser `json:"userlist"`
	PartyList []int      `json:"partylist"`
}

type TagUser struct {
	UserID string `json:"userid"`
	Name   string `json:"name"`
}

func GetTagUser(tagID int64, result *ResultTagUser) wx.Action {
	return wx.NewGetAction(urls.CorpUserTagGetUser,
		wx.WithQuery("tagid", strconv.FormatInt(tagID, 10)),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsTagUserAdd struct {
	TagID     int64    `json:"tagid"`
	UserList  []string `json:"userlist"`
	PartyList []int64  `json:"partylist"`
}

type ResultTagUserAdd struct {
	InvalidList  string  `json:"invalidlist"`
	InvalidParty []int64 `json:"invalidparty"`
}

func AddTagUser(tagID int64, userIDs []string, partyIDs []int64, result *ResultTagUserAdd) wx.Action {
	params := &ParamsTagUserAdd{
		TagID:     tagID,
		UserList:  userIDs,
		PartyList: partyIDs,
	}

	return wx.NewPostAction(urls.CorpUserTagAddUser,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsTagUserDelete struct {
	TagID     int64    `json:"tagid"`
	UserList  []string `json:"userlist"`
	PartyList []int64  `json:"partylist"`
}

type ResultTagUserDelete struct {
	InvalidList  string  `json:"invalidlist"`
	InvalidParty []int64 `json:"invalidparty"`
}

func DeleteTagUser(tagID int64, userIDs []string, partyIDs []int64, result *ResultTagUserDelete) wx.Action {
	params := &ParamsTagUserDelete{
		TagID:     tagID,
		UserList:  userIDs,
		PartyList: partyIDs,
	}

	return wx.NewPostAction(urls.CorpUserTagDeleteUser,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ResultTagList struct {
	TagList []*Tag `json:"taglist"`
}

type Tag struct {
	TagID   int64  `json:"tagid"`
	TagName string `json:"tagname"`
}

func ListTag(result *ResultTagList) wx.Action {
	return wx.NewGetAction(urls.CorpUserTagList,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
