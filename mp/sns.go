package mp

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/iiinsomnia/gochat/utils"

	"github.com/tidwall/gjson"
)

// AuthSession 小程序授权Session
type AuthSession struct {
	SessionKey string `json:"session_key"`
	OpenID     string `json:"openid"`
	UnionID    string `json:"unionid"`
}

// Sns sns
type Sns struct {
	appid     string
	appsecret string
	client    *utils.HTTPClient
}

// Code2Session 获取小程序授权SessionKey
func (s *Sns) Code2Session(code string) (*AuthSession, error) {
	resp, err := s.client.Get(fmt.Sprintf("%s?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", Code2SessionURL, s.appid, s.appsecret, code))

	if err != nil {
		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return nil, errors.New(r.Get("errmsg").String())
	}

	reply := new(AuthSession)

	if err := json.Unmarshal(resp, reply); err != nil {
		return nil, err
	}

	return reply, nil
}
