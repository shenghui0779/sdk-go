package mch

import (
	"net/url"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

// ShortURL 转换短链接
// 该接口主要用于Native支付模式一中的二维码链接转成短链接(weixin://wxpay/s/XXXXXX)，减小二维码数据量，提升扫描速度和精确度。
// 【注意】二维码链接无需URLEncode，因为：签名需用原串
func ShortURL(appid, longURL string) wx.Action {
	return wx.NewPostAction(urls.MchToolsShortURL,
		wx.WithWXML(func(mchid, apikey, nonce string) (wx.WXML, error) {
			m := wx.WXML{
				"appid":     appid,
				"mch_id":    mchid,
				"long_url":  longURL,
				"nonce_str": nonce,
			}

			// 签名用原串
			m["sign"] = wx.SignMD5.Do(apikey, m, true)

			// 传输需URLencode
			m["long_url"] = url.QueryEscape(longURL)

			return m, nil
		}),
	)
}

// AuthCodeToOpenID 付款码查询openid
// 通过付款码查询公众号Openid，调用查询后，该付款码只能由此商户号发起扣款，直至付款码更新。
func AuthCodeToOpenID(appid, authCode string) wx.Action {
	return wx.NewPostAction(urls.MchAuthCodeToOpenID,
		wx.WithWXML(func(mchid, apikey, nonce string) (wx.WXML, error) {
			m := wx.WXML{
				"appid":     appid,
				"mch_id":    mchid,
				"auth_code": authCode,
				"nonce_str": nonce,
			}

			// 签名
			m["sign"] = wx.SignMD5.Do(apikey, m, true)

			return m, nil
		}),
	)
}

// RSAPublicKey 获取RSA加密公钥（需要证书）
func RSAPublicKey() wx.Action {
	return wx.NewPostAction(urls.MchRSAPublicKey,
		wx.WithTLS(),
		wx.WithWXML(func(mchid, apikey, nonce string) (wx.WXML, error) {
			m := wx.WXML{
				"mch_id":    mchid,
				"nonce_str": nonce,
				"sign_type": "MD5",
			}

			// 签名
			m["sign"] = wx.SignMD5.Do(apikey, m, true)

			return m, nil
		}),
	)
}
