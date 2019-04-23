package wxpay

import (
	"encoding/xml"

	"meipian.cn/printapi/wechat/utils"
)

// Reply 回复支付结果
type Reply struct {
	XMLName    xml.Name    `xml:"xml"`
	ReturnCode utils.CDATA `xml:"return_code"`
	ReturnMsg  utils.CDATA `xml:"return_msg"`
}

// ReplyOK ...
func ReplyOK() *Reply {
	return &Reply{
		ReturnCode: "SUCCESS",
		ReturnMsg:  "OK",
	}
}

// ReplyFail ...
func ReplyFail(msg string) *Reply {
	return &Reply{
		ReturnCode: "FAIL",
		ReturnMsg:  utils.CDATA(msg),
	}
}
