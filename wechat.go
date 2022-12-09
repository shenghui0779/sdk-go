package gochat

import (
	"github.com/shenghui0779/gochat/corp"
	"github.com/shenghui0779/gochat/mch"
	"github.com/shenghui0779/gochat/minip"
	"github.com/shenghui0779/gochat/offia"
)

// NewMch 微信商户
func NewMch(mchid, apikey string, options ...mch.Option) *mch.Mch {
	return mch.New(mchid, apikey, options...)
}

// NewOffia 微信公众号
func NewOffia(appid, appsecret string, options ...offia.Option) *offia.Offia {
	return offia.New(appid, appsecret, options...)
}

// NewMinip 微信小程序
func NewMinip(appid, appsecret string, options ...minip.Option) *minip.Minip {
	return minip.New(appid, appsecret, options...)
}

// NewCorp 企业微信
func NewCorp(corpid string, options ...corp.Option) *corp.Corp {
	return corp.New(corpid, options...)
}
