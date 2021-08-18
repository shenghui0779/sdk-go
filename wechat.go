package gochat

import (
	"github.com/shenghui0779/gochat/mch"
	"github.com/shenghui0779/gochat/minip"
	"github.com/shenghui0779/gochat/offia"
)

// NewMch 微信商户
func NewMch(appid, mchid, apikey string) *mch.Mch {
	return mch.New(appid, mchid, apikey)
}

// NewOffia 微信公众号
func NewOffia(appid, appsecret string) *offia.Offia {
	return offia.New(appid, appsecret)
}

// NewMinip 微信小程序
func NewMinip(appid, appsecret string) *minip.Minip {
	return minip.New(appid, appsecret)
}
