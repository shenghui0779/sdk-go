package addr_book

import (
	"encoding/json"
	"strconv"

	"github.com/tidwall/gjson"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type Tag struct {
	ID   int64  `json:"tagid"`
	Name string `json:"tagname"`
}

// CreateTag 创建标签
func CreateTag(data *Tag) wx.Action {
	return wx.NewPostAction(urls.CorpAddrBookTagCreate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(data)
		}),
		wx.WithDecode(func(resp []byte) error {
			data.ID = gjson.GetBytes(resp, "tagid").Int()

			return nil
		}),
	)
}

func UpdateTag(data *Tag) wx.Action {
	return wx.NewPostAction(urls.CorpAddrBookTagUpdate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(data)
		}),
	)
}

func DeleteTag(tagID int64) wx.Action {
	return wx.NewGetAction(urls.CorpAddrBookTagDelete,
		wx.WithQuery("tagid", strconv.FormatInt(tagID, 10)),
	)
}

type TagUser struct {
	UserID string `json:"userid"`
	Name   string `json:"name"`
}

type TagSpec struct {
	TagName   string     `json:"tagname"`
	UserList  []*TagUser `json:"userlist"`
	PartyList []int      `json:"partylist"`
}

func GetTag(dest *TagSpec, tagID int64) wx.Action {
	return wx.NewGetAction(urls.CorpAddrBookTagGet,
		wx.WithQuery("tagid", strconv.FormatInt(tagID, 10)),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

func GetTagList(dest *[]*Tag) wx.Action {
	return wx.NewGetAction(urls.CorpAddrBookTagList,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal([]byte(gjson.GetBytes(resp, "taglist").Raw), dest)
		}),
	)
}
