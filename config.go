package gochat

import "sync"

type WXChannel string // 渠道

const (
	WXAPP WXChannel = "wxapp" // 微信APP渠道名称
	WXPub WXChannel = "wxpub" // 微信公众号渠道名称
	WXH5  WXChannel = "wxh5"  // 微信H5渠道名称
	WXMP  WXChannel = "wxmp"  // 微信小程序渠道名称
)

// WXConfig 微信配置
type WXConfig struct {
	AccountID      string
	AppID          string
	AppSecret      string
	MchID          string
	ApiKey         string
	TradeType      string
	SignToken      string
	EncodingAESKey string
	SSLCertFile    string
	SSLKeyFile     string
}

var wxmap sync.Map

// SetWXAPPConfig 设置微信APP配置
func SetWXAPPConfig(appid, appsecret, mchid, apikey, sslCertFile, sslkeyFile string) {
	wxmap.Store(WXAPP, &WXConfig{
		AppID:       appid,
		AppSecret:   appsecret,
		MchID:       mchid,
		ApiKey:      apikey,
		TradeType:   "APP",
		SSLCertFile: sslCertFile,
		SSLKeyFile:  sslkeyFile,
	})
}

// SetWXPubConfig 设置微信公众号配置
func SetWXPubConfig(accountid, appid, appsecret, mchid, apikey, signToken, encodingAESKey, sslCertFile, sslkeyFile string) {
	wxmap.Store(WXPub, &WXConfig{
		AccountID:      accountid,
		AppID:          appid,
		AppSecret:      appsecret,
		MchID:          mchid,
		ApiKey:         apikey,
		TradeType:      "JSAPI",
		SignToken:      signToken,
		EncodingAESKey: encodingAESKey,
		SSLCertFile:    sslCertFile,
		SSLKeyFile:     sslkeyFile,
	})
}

// SetWXMPConfig 设置微信小程序配置
func SetWXMPConfig(appid, appsecret, mchid, apikey, sslCertFile, sslkeyFile string) {
	wxmap.Store(WXMP, &WXConfig{
		AppID:       appid,
		AppSecret:   appsecret,
		MchID:       mchid,
		ApiKey:      apikey,
		TradeType:   "JSAPI",
		SSLCertFile: sslCertFile,
		SSLKeyFile:  sslkeyFile,
	})
}

// SetWXH5Config 设置H5配置
func SetWXH5Config(appid, appsecret, mchid, apikey, sslCertFile, sslkeyFile string) {
	wxmap.Store(WXH5, &WXConfig{
		AppID:       appid,
		AppSecret:   appsecret,
		MchID:       mchid,
		ApiKey:      apikey,
		TradeType:   "MWEB",
		SSLCertFile: sslCertFile,
		SSLKeyFile:  sslkeyFile,
	})
}

// GetConfigWithChannel returns wx config with channel
func GetConfigWithChannel(channel WXChannel) *WXConfig {
	v, ok := wxmap.Load(channel)

	if !ok {
		return new(WXConfig)
	}

	return v.(*WXConfig)
}
