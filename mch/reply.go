package mch

import (
	"encoding/xml"

	"github.com/iiinsomnia/gochat/utils"
)

// Reply 回复支付结果
type Reply struct {
	XMLName    xml.Name    `xml:"xml"`
	ReturnCode utils.CDATA `xml:"return_code"`
	ReturnMsg  utils.CDATA `xml:"return_msg"`
}

// ReplyOK 回复成功
func ReplyOK() *Reply {
	return &Reply{
		ReturnCode: ResultSuccess,
		ReturnMsg:  "OK",
	}
}

// ReplyFail 回复失败
func ReplyFail(msg string) *Reply {
	return &Reply{
		ReturnCode: ResultFail,
		ReturnMsg:  utils.CDATA(msg),
	}
}
