package gochat

import (
	"crypto/tls"

	"github.com/shenghui0779/gochat/corp"
	"github.com/shenghui0779/gochat/mch"
	"github.com/shenghui0779/gochat/minip"
	"github.com/shenghui0779/gochat/offia"
)

// NewMch 微信商户
func NewMch(mchid, apikey string, certs ...tls.Certificate) *mch.Mch {
	return mch.New(mchid, apikey, certs...)
}

// NewOffia 微信公众号
func NewOffia(appid, appsecret string) *offia.Offia {
	return offia.New(appid, appsecret)
}

// NewMinip 微信小程序
func NewMinip(appid, appsecret string) *minip.Minip {
	return minip.New(appid, appsecret)
}

// NewCorp 企业微信
func NewCorp(corpid string) *corp.Corp {
	return corp.New(corpid)
}
