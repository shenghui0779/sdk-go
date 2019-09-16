package pub

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/iiinsomnia/gochat/consts"
	"github.com/iiinsomnia/gochat/utils"

	"github.com/tidwall/gjson"
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
	appid     string
	appsecret string
	client    *utils.HTTPClient
}

// GetAccessToken returns access_token
func (c *CgiBin) GetAccessToken() (*AccessToken, error) {
	resp, err := c.client.Get(fmt.Sprintf("%s?grant_type=client_credential&appid=%s&secret=%s", consts.PubCgiBinAccessTokenURL, c.appid, c.appsecret))

	if err != nil {
		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return nil, errors.New(r.Get("errmsg").String())
	}

	reply := new(AccessToken)

	if err := json.Unmarshal(resp, reply); err != nil {
		return nil, err
	}

	return reply, nil
}

// GetTicket returns jsapi ticket
func (c *CgiBin) GetTicket(accessToken string) (*JSAPITicket, error) {
	resp, err := c.client.Get(fmt.Sprintf("%s?access_token=%s&type=jsapi", consts.PubCgiBinTicketURL, accessToken))

	if err != nil {
		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return nil, errors.New(r.Get("errmsg").String())
	}

	reply := new(JSAPITicket)

	if err := json.Unmarshal(resp, reply); err != nil {
		return nil, err
	}

	return reply, nil
}
