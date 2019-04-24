package mp

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/iiinsomnia/gochat/utils"
)

// Sns ...
type Sns struct {
	appid     string
	appsecret string
	reply     *snsReply
}

type snsReply struct {
	SessionKey string `json:"session_key"`
	OpenID     string `json:"openid"`
	UnionID    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// Code2Session 获取小程序授权SessionKey
func (s *Sns) Code2Session(code string) error {
	resp, err := utils.HTTPGet(fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", s.appid, s.appsecret, code))

	if err != nil {
		return err
	}

	reply := new(snsReply)

	if err := json.Unmarshal(resp, reply); err != nil {
		return err
	}

	if reply.ErrCode != 0 {
		return errors.New(reply.ErrMsg)
	}

	s.reply = reply

	return nil
}

// SessionKey ...
func (s *Sns) SessionKey() string {
	return s.reply.SessionKey
}

// OpenID ...
func (s *Sns) OpenID() string {
	return s.reply.OpenID
}

// UnionID ...
func (s *Sns) UnionID() string {
	return s.reply.UnionID
}
