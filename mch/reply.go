package mch

import (
	"encoding/xml"

	"github.com/shenghui0779/gochat/wx"
)

// Reply 回复支付结果
type Reply struct {
	XMLName    xml.Name `xml:"xml"`
	ReturnCode wx.CDATA `xml:"return_code"`
	ReturnMsg  wx.CDATA `xml:"return_msg"`
}

// ReplyOK 回复成功
func ReplyOK() *Reply {
	return &Reply{
		ReturnCode: wx.CDATA(ResultSuccess),
		ReturnMsg:  wx.CDATA("OK"),
	}
}

// ReplyFail 回复失败
func ReplyFail(msg string) *Reply {
	return &Reply{
		ReturnCode: wx.CDATA(ResultFail),
		ReturnMsg:  wx.CDATA(msg),
	}
}
