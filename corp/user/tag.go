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
func CreateTag(params *ParamsTagCreate, result *ResultTagCreate) wx.Action {
	return wx.NewPostAction(urls.CorpTagCreate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
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

func UpdateTag(params *ParamsTagUpdate) wx.Action {
	return wx.NewPostAction(urls.CorpTagUpdate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

func DeleteTag(tagID int64) wx.Action {
	return wx.NewGetAction(urls.CorpTagDelete,
		wx.WithQuery("tagid", strconv.FormatInt(tagID, 10)),
	)
}

type Tag struct {
	TagName   string     `json:"tagname"`
	UserList  []*TagUser `json:"userlist"`
	PartyList []int      `json:"partylist"`
}

type TagUser struct {
	UserID string `json:"userid"`
	Name   string `json:"name"`
}

func GetTag(tagID int64, result *Tag) wx.Action {
	return wx.NewGetAction(urls.CorpTagGet,
		wx.WithQuery("tagid", strconv.FormatInt(tagID, 10)),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ResultTagList struct {
	TagList []*Tag `json:"taglist"`
}

func ListTag(result *ResultTagList) wx.Action {
	return wx.NewGetAction(urls.CorpTagList,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
