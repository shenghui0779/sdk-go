package corp

import (
	"encoding/json"
	"strconv"

	"github.com/tidwall/gjson"

	"github.com/shenghui0779/gochat/wx"
)

type Tag struct {
	ID   int    `json:"tagid"`
	Name string `json:"tagname"`
}

// TagCreate 创建标签
func TagCreate(data *Tag) wx.Action {
	return wx.NewPostAction(TagCreateURL,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(data)
		}),
		wx.WithDecode(func(resp []byte) error {
			data.ID = int(gjson.GetBytes(resp, "tagid").Int())

			return nil
		}),
	)
}

func TagUpdate(data *Tag) wx.Action {
	return wx.NewPostAction(TagUpdateURL,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(data)
		}),
	)
}

func TagDelete(tagID int) wx.Action {
	return wx.NewGetAction(TagDeleteURL,
		wx.WithQuery("tagid", strconv.Itoa(tagID)),
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

func TagGet(dest *TagSpec, tagID int) wx.Action {
	return wx.NewGetAction(TagGetURL,
		wx.WithQuery("tagid", strconv.Itoa(tagID)),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

func TagList(dest *[]*Tag) wx.Action {
	return wx.NewGetAction(TagListURL,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal([]byte(gjson.GetBytes(resp, "taglist").Raw), dest)
		}),
	)
}
