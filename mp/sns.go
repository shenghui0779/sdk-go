package mp

import (
	"encoding/json"
	"fmt"

	"github.com/shenghui0779/gochat/utils"
)

// AuthSession 小程序授权Session
type AuthSession struct {
	SessionKey string `json:"session_key"`
	OpenID     string `json:"openid"`
	UnionID    string `json:"unionid"`
}

// Sns sns
type Sns struct {
	mp      *WXMP
	options []utils.RequestOption
}

// Code2Session 获取小程序授权SessionKey
func (s *Sns) Code2Session(code string) (*AuthSession, error) {
	b, err := s.mp.get(fmt.Sprintf("%s?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", Code2SessionURL, s.mp.appid, s.mp.appsecret, code), s.options...)

	if err != nil {
		return nil, err
	}

	resp := new(AuthSession)

	if err := json.Unmarshal(b, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
