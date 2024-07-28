package wechat

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/sdk-go/lib/value"
)

func TestXML(t *testing.T) {
	m := value.V{
		"appid":     "wx2421b1c4370ec43b",
		"partnerid": "10000100",
		"prepayid":  "WX1217752501201407033233368018",
		"package":   "Sign=WXPay",
		"noncestr":  "5K8264ILTKCH16CQ2502SI8ZNMTM67VS",
		"timestamp": "1514363815",
	}
	x, err := ValueToXML(m)
	assert.Nil(t, err)

	r, err := XMLToValue(x)
	assert.Nil(t, err)
	assert.Equal(t, m, r)
}
