package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWXML(t *testing.T) {
	m := WXML{
		"appid":     "wx2421b1c4370ec43b",
		"partnerid": "10000100",
		"prepayid":  "WX1217752501201407033233368018",
		"package":   "Sign=WXPay",
		"noncestr":  "5K8264ILTKCH16CQ2502SI8ZNMTM67VS",
		"timestamp": "1514363815",
	}

	x, err := FormatMap2XML(m)

	assert.Nil(t, err)

	r, err := ParseXML2Map([]byte(x))

	assert.Nil(t, err)
	assert.Equal(t, m, r)
}

func TestSignWithMD5(t *testing.T) {
	m := WXML{
		"appid":     "wx2421b1c4370ec43b",
		"partnerid": "10000100",
		"prepayid":  "WX1217752501201407033233368018",
		"package":   "Sign=WXPay",
		"noncestr":  "5K8264ILTKCH16CQ2502SI8ZNMTM67VS",
		"timestamp": "1514363815",
	}

	sign := SignWithMD5(m, "192006250b4c09247ec02edce69f6a2d", true)

	// 签名校验来自：[微信支付接口签名校验工具](https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=20_1)
	assert.Equal(t, "66724B3332E124BFC3D62A31A68F7887", sign)
}

func TestSignWithHMacSHA256(t *testing.T) {
	m := WXML{
		"appid":     "wx2421b1c4370ec43b",
		"partnerid": "10000100",
		"prepayid":  "WX1217752501201407033233368018",
		"package":   "Sign=WXPay",
		"noncestr":  "5K8264ILTKCH16CQ2502SI8ZNMTM67VS",
		"timestamp": "1514363815",
	}

	sign := SignWithHMacSHA256(m, "192006250b4c09247ec02edce69f6a2d", true)

	// 签名校验来自：[微信支付接口签名校验工具](https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=20_1)
	assert.Equal(t, "3B12F569A5714858F8251366BC3CBCDDBD249905CCA01D8F56D365EF1FC2CA5C", sign)
}
