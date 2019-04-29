package gochat

import (
	"github.com/iiinsomnia/gochat/mch"
	"github.com/iiinsomnia/gochat/mp"
	"github.com/iiinsomnia/gochat/pub"
	"github.com/iiinsomnia/gochat/utils"
)

// NewWXMch 微信商户
func NewWXMch(appid, mchid, apikey, sslCertFile, sslkeyFile string) *mch.WXMch {
	wxmch := &mch.WXMch{
		AppID:  appid,
		MchID:  mchid,
		ApiKey: apikey,
	}

	wxmch.SetHTTPClient(utils.DefaultHTTPClient)

	// SSL Client
	c, err := utils.NewHTTPClient(utils.WithHTTPSSLCertFile(sslCertFile, sslkeyFile))

	if err != nil {
		wxmch.SetSSLClient(utils.DefaultHTTPClient)
	} else {
		wxmch.SetSSLClient(c)
	}

	return wxmch
}

// NewWXPub 微信公众号
func NewWXPub(accountid, appid, appsecret, signToken, encodingAESKey string) *pub.WXPub {
	wxpub := &pub.WXPub{
		AccountID:      accountid,
		AppID:          appid,
		AppSecret:      appsecret,
		SignToken:      signToken,
		EncodingAESKey: encodingAESKey,
	}

	wxpub.SetHTTPClient(utils.DefaultHTTPClient)

	return wxpub
}

// NewWXMP 微信小程序
func NewWXMP(appid, appsecret string) *mp.WXMP {
	wxmp := &mp.WXMP{
		AppID:     appid,
		AppSecret: appsecret,
	}

	wxmp.SetHTTPClient(utils.DefaultHTTPClient)

	return wxmp
}
