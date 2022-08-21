package wx

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

func TestMD5(t *testing.T) {
	assert.Equal(t, "483367436bc9a6c5256bfc29a24f955e", MD5("iiinsomnia"))
}

func TestSHA256(t *testing.T) {
	assert.Equal(t, "efed14231acf19fdca03adfac049171c109c922008e64dbaaf51a0c2cf11306b", SHA256("iiinsomnia"))
}

func TestHMacSHA256(t *testing.T) {
	assert.Equal(t, "8a50abfc64f67dac0f6aa8b6560d26517574ce30b5f3487a258ab04b30776db4", HMacSHA256("ILoveGochat", "iiinsomnia"))
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

	// 签名校验来自：[微信支付接口签名校验工具](https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=20_1)
	assert.Equal(t, "66724B3332E124BFC3D62A31A68F7887", SignMD5.Do("192006250b4c09247ec02edce69f6a2d", m, true))
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

	// 签名校验来自：[微信支付接口签名校验工具](https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=20_1)
	assert.Equal(t, "3B12F569A5714858F8251366BC3CBCDDBD249905CCA01D8F56D365EF1FC2CA5C", SignHMacSHA256.Do("192006250b4c09247ec02edce69f6a2d", m, true))
}

func TestUint32Bytes(t *testing.T) {
	i := uint32(250)
	b := EncodeUint32ToBytes(i)

	assert.Equal(t, i, DecodeBytesToUint32(b))
}

func TestMarshalNoEscapeHTML(t *testing.T) {
	b, err := MarshalNoEscapeHTML(M{
		"action":   "long2short",
		"long_url": "http://wap.koudaitong.com/v2/showcase/goods?alias=128wi9shh&spm=h56083&redirect_count=1",
	})

	assert.Nil(t, err)
	assert.Equal(t, `{"action":"long2short","long_url":"http://wap.koudaitong.com/v2/showcase/goods?alias=128wi9shh&spm=h56083&redirect_count=1"}`, string(b))
}
