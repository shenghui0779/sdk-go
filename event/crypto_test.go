package event

import (
	"encoding/base64"
	"testing"

	"github.com/shenghui0779/gochat/wx"
	"github.com/stretchr/testify/assert"
)

func TestCrypto(t *testing.T) {
	appid := "wx1def0e9e5891b338"
	encodingAESKey := "jxAko083VoJ3lcPXJWzcGJ0M1tFVLgdD6qAq57GJY1U"

	cb, err := Encrypt(appid, encodingAESKey, "343a802b6073aae5", []byte("<xml><ToUserName><![CDATA[gh_3ad31c0ba9b5]]></ToUserName><FromUserName><![CDATA[oB4tA6ANthOfuQ5XSlkdPsWOVUsY]]></FromUserName><CreateTime>1606902602</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[ILoveGochat]]></Content><MsgId>10086</MsgId></xml>"))

	assert.Nil(t, err)

	pb, err := Decrypt(appid, encodingAESKey, base64.StdEncoding.EncodeToString(cb))

	assert.Nil(t, err)

	msg, err := wx.ParseXML2Map(pb)

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
		"ToUserName":   "gh_3ad31c0ba9b5",
		"FromUserName": "oB4tA6ANthOfuQ5XSlkdPsWOVUsY",
		"CreateTime":   "1606902602",
		"MsgType":      "text",
		"MsgId":        "10086",
		"Content":      "ILoveGochat",
	}, msg)
}
