package corp

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/wx"
)

// AccessToken 企业微信AccessToken
type AccessToken struct {
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

type APIDomainIP struct {
	IPList []string `json:"ip_list"`
}

func GetAPIDomainIP(dest *APIDomainIP) wx.Action {
	return wx.NewGetAction(APIDomainIPURL,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}
