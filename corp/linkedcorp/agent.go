package linkedcorp

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ResultAgentPermList struct {
	UserIDs       []string `json:"userids"`
	DepartmentIDs []string `json:"department_ids"`
}

func ListAgentPerm(result *ResultAgentPermList) wx.Action {
	return wx.NewPostAction(urls.CorpLinkedcorpPermList,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
