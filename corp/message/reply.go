package message

import (
	"encoding/xml"
	"time"

	"github.com/shenghui0779/gochat/event"
	"github.com/shenghui0779/gochat/wx"
)

type XMLText struct {
	Content wx.CDATA `xml:"Content"`
}

type XMLMedia struct {
	MediaID wx.CDATA `xml:"MediaId"`
}

type XMLVideo struct {
	MediaID     wx.CDATA `xml:"MediaId"`
	Title       wx.CDATA `xml:"Title"`
	Description wx.CDATA `xml:"Description"`
}

type XMLNews struct {
	Articles []*NewsArticle `xml:"Articles"`
}

type XMLNewsArticle struct {
	Title       wx.CDATA `xml:"Title"`
	Description wx.CDATA `xml:"Description"`
	URL         wx.CDATA `xml:"Url"`
	PicURL      wx.CDATA `xml:"PicUrl"`
}

// Reply 消息回复
type Reply struct {
	XMLName      xml.Name          `xml:"xml"`
	FromUserName wx.CDATA          `xml:"FromUserName"`
	ToUserName   wx.CDATA          `xml:"ToUserName"`
	CreateTime   int64             `xml:"CreateTime"`
	MsgType      wx.CDATA          `xml:"MsgType"`
	Content      wx.CDATA          `xml:"Content,omitempty"`
	Image        *XMLMedia         `xml:"Image,omitempty"`
	Voice        *XMLMedia         `xml:"Voice,omitempty"`
	Video        *XMLVideo         `xml:"Video,omitempty"`
	ArticleCount int               `xml:"ArticleCount,omitempty"`
	Articles     []*XMLNewsArticle `xml:"Articles>item,omitempty"`
}

func (r *Reply) Bytes(from, to string) ([]byte, error) {
	r.FromUserName = wx.CDATA(from)
	r.ToUserName = wx.CDATA(to)

	return xml.Marshal(r)
}

func ReplyText(content string) event.Reply {
	return &Reply{
		CreateTime: time.Now().Unix(),
		MsgType:    wx.CDATA(event.MsgText),
		Content:    wx.CDATA(content),
	}
}

func ReplyImage(mediaID string) event.Reply {
	return &Reply{
		CreateTime: time.Now().Unix(),
		MsgType:    wx.CDATA(event.MsgText),
		Image: &XMLMedia{
			MediaID: wx.CDATA(mediaID),
		},
	}
}

func ReplyVoice(mediaID string) event.Reply {
	return &Reply{
		CreateTime: time.Now().Unix(),
		MsgType:    wx.CDATA(event.MsgText),
		Voice: &XMLMedia{
			MediaID: wx.CDATA(mediaID),
		},
	}
}

func ReplyVideo(mediaID, title, description string) event.Reply {
	return &Reply{
		CreateTime: time.Now().Unix(),
		MsgType:    wx.CDATA(event.MsgText),
		Video: &XMLVideo{
			MediaID:     wx.CDATA(mediaID),
			Title:       wx.CDATA(title),
			Description: wx.CDATA(description),
		},
	}
}

func ReplyNews(articles ...*XMLNewsArticle) event.Reply {
	return &Reply{
		CreateTime:   time.Now().Unix(),
		MsgType:      wx.CDATA(event.MsgText),
		ArticleCount: len(articles),
		Articles:     articles,
	}
}
