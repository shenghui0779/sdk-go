package gochat

import (
	"github.com/shenghui0779/gochat/mch"
	"github.com/shenghui0779/gochat/mp"
	"github.com/shenghui0779/gochat/oa"
)

// NewMch 微信商户
func NewMch(appid, mchid, apikey string) *mch.Mch {
	return mch.New(appid, mchid, apikey)
}

// NewPub 微信公众号
func NewOA(appid, appsecret string) *oa.OA {
	return oa.New(appid, appsecret)
}

// NewMP 微信小程序
func NewMP(appid, appsecret string) *mp.MP {
	return mp.New(appid, appsecret)
}
