package event

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/shenghui0779/gochat/wx"
)

// Reply 消息回复
type Reply interface {
	Bytes(from, to string) ([]byte, error)
}

// ReplyMessage  回复消息
type ReplyMessage struct {
	XMLName      xml.Name `xml:"xml"`
	Encrypt      wx.CDATA `xml:"Encrypt"`
	MsgSignature wx.CDATA `xml:"MsgSignature"`
	TimeStamp    int64    `xml:"TimeStamp"`
	Nonce        wx.CDATA `xml:"Nonce"`
}

// BuildReply 生成回复消息
func BuildReply(token, nonce, encryptMsg string) *ReplyMessage {
	now := time.Now().Unix()

	signItems := []string{token, strconv.FormatInt(now, 10), nonce, encryptMsg}

	sort.Strings(signItems)

	h := sha1.New()
	h.Write([]byte(strings.Join(signItems, "")))

	return &ReplyMessage{
		Encrypt:      wx.CDATA(encryptMsg),
		MsgSignature: wx.CDATA(hex.EncodeToString(h.Sum(nil))),
		TimeStamp:    now,
		Nonce:        wx.CDATA(nonce),
	}
}
