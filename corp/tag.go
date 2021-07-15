package corp

import (
	"encoding/json"

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
