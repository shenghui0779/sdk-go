package mp

import (
	"encoding/json"
	"fmt"

	"github.com/shenghui0779/gochat/utils"
)

// AccessToken wxmp access_token
type AccessToken struct {
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

// CgiBin cgi-bin
type CgiBin struct {
	mp      *WXMP
	options []utils.RequestOption
}

// GetAccessToken returns access_token
func (p *CgiBin) GetAccessToken() (*AccessToken, error) {
	b, err := p.mp.get(fmt.Sprintf("%s?grant_type=client_credential&appid=%s&secret=%s", AccessTokenURL, p.mp.appid, p.mp.appsecret), p.options...)

	if err != nil {
		return nil, err
	}

	resp := new(AccessToken)

	if err := json.Unmarshal(b, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
