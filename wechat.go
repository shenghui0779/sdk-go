package gochat

import (
	"github.com/shenghui0779/gochat/mch"
	"github.com/shenghui0779/gochat/mp"
	"github.com/shenghui0779/gochat/pub"
)

// NewWXMch 微信商户
func NewWXMch(appid, mchid, apikey string, tlsInsecureSkipVerify bool) *mch.WXMch {
	return mch.New(appid, mchid, apikey, tlsInsecureSkipVerify)
}

// NewWXPub 微信公众号
func NewWXPub(accountid, appid, appsecret, signToken, encodingAESKey string) *pub.WXPub {
	return pub.New(accountid, appid, appsecret, signToken, encodingAESKey)
}

// NewWXMP 微信小程序
func NewWXMP(appid, appsecret string) *mp.WXMP {
	return mp.New(appid, appsecret)
}
