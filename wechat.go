package gochat

import (
	"github.com/iiinsomnia/gochat/mch"
	"github.com/iiinsomnia/gochat/mp"
	"github.com/iiinsomnia/gochat/pub"
	"github.com/iiinsomnia/gochat/utils"
)

// NewWXMch 微信商户
func NewWXMch(appid, mchid, apikey, sslCertFile, sslkeyFile string) (*mch.WXMch, error) {
	mch := &mch.WXMch{
		AppID:  appid,
		MchID:  mchid,
		ApiKey: apikey,
	}

	c, err := utils.NewHTTPClient(utils.WithHTTPSSLCertFile(sslCertFile, sslkeyFile))

	if err != nil {
		return nil, err
	}

	mch.SetSSLClient(c)

	return mch, nil
}

// NewWXPub 微信公众号
func NewWXPub(accountid, appid, appsecret, signToken, encodingAESKey string) *pub.WXPub {
	return &pub.WXPub{
		AccountID:      accountid,
		AppID:          appid,
		AppSecret:      appsecret,
		SignToken:      signToken,
		EncodingAESKey: encodingAESKey,
	}
}

// NewWXMP 微信小程序
func NewWXMP(appid, appsecret string) *mp.WXMP {
	return &mp.WXMP{
		AppID:     appid,
		AppSecret: appsecret,
	}
}
