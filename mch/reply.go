package mch

import "encoding/xml"

// CDATA XML CDATA section which is defined as blocks of text that are not parsed by the parser, but are otherwise recognized as markup.
type CDATA string

// MarshalXML encodes the receiver as zero or more XML elements.
func (c CDATA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(c)}, start)
}

// Reply 回复支付结果
type Reply struct {
	XMLName    xml.Name `xml:"xml"`
	ReturnCode CDATA    `xml:"return_code"`
	ReturnMsg  CDATA    `xml:"return_msg"`
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
		ReturnMsg:  CDATA(msg),
	}
}
