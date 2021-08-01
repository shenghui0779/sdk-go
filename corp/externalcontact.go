package corp

import (
	"encoding/json"

	"github.com/tidwall/gjson"

	"github.com/shenghui0779/gochat/wx"
)

func GetExternalContactFollowUserList(dest *[]string) wx.Action {
	return wx.NewGetAction(ExternalContactFollowUserListURL,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal([]byte(gjson.GetBytes(resp, "follow_user").Raw), dest)
		}),
	)
}

func AddExternalContactWay() {

}
