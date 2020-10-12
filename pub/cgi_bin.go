package pub

import (
	"encoding/json"
	"fmt"

	"github.com/shenghui0779/gochat/utils"
)

// AccessToken wxpub access_token
type AccessToken struct {
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

// JSAPITicket wxpub js api ticket
type JSAPITicket struct {
	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
}

// CgiBin cgi-bin
type CgiBin struct {
	pub     *WXPub
	options []utils.RequestOption
}

// GetAccessToken returns access_token
func (c *CgiBin) GetAccessToken() (*AccessToken, error) {
	b, err := c.pub.get(fmt.Sprintf("%s?grant_type=client_credential&appid=%s&secret=%s", CgiBinAccessTokenURL, c.pub.appid, c.pub.appsecret), c.options...)

	if err != nil {
		return nil, err
	}

	resp := new(AccessToken)

	if err := json.Unmarshal(b, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// GetTicket returns jsapi ticket
func (c *CgiBin) GetTicket(accessToken string) (*JSAPITicket, error) {
	b, err := c.pub.get(fmt.Sprintf("%s?access_token=%s&type=jsapi", CgiBinTicketURL, accessToken), c.options...)

	if err != nil {
		return nil, err
	}

	resp := new(JSAPITicket)

	if err := json.Unmarshal(b, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
