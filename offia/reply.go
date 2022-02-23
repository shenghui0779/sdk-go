package offia

import (
	"encoding/xml"
	"time"

	"github.com/shenghui0779/gochat/event"
	"github.com/shenghui0779/gochat/wx"
)

type XMLText struct {
	Content wx.CDATA `xml:"Content,omitempty"`
}

type XMLMedia struct {
	MediaID wx.CDATA `xml:"MediaId,omitempty"`
}

type XMLVideo struct {
	MediaID     wx.CDATA `xml:"MediaId,omitempty"`
	Title       wx.CDATA `xml:"Title,omitempty"`
	Description wx.CDATA `xml:"Description,omitempty"`
}

type XMLMusic struct {
	Title        wx.CDATA `xml:"Title,omitempty"`
	Description  wx.CDATA `xml:"Description,omitempty"`
	MusicURL     wx.CDATA `xml:"MusicUrl,omitempty"`
	HQMusicURL   wx.CDATA `xml:"HQMusicUrl,omitempty"`
	ThumbMediaID wx.CDATA `xml:"ThumbMediaId,omitempty"`
}

type XMLNews struct {
	Articles []*XMLNewsArticle `xml:"item,omitempty"`
}

type XMLNewsArticle struct {
	Title       wx.CDATA `xml:"Title,omitempty"`
	Description wx.CDATA `xml:"Description,omitempty"`
	URL         wx.CDATA `xml:"Url,omitempty"`
	PicURL      wx.CDATA `xml:"PicUrl,omitempty"`
}

type XMLTransInfo struct {
	KFAccount wx.CDATA `xml:"KfAccount,omitempty"`
}

// Reply 消息回复
type Reply struct {
	XMLName      xml.Name      `xml:"xml"`
	FromUserName wx.CDATA      `xml:"FromUserName,omitempty"`
	ToUserName   wx.CDATA      `xml:"ToUserName,omitempty"`
	CreateTime   int64         `xml:"CreateTime,omitempty"`
	MsgType      wx.CDATA      `xml:"MsgType,omitempty"`
	Content      wx.CDATA      `xml:"Content,omitempty"`
	Image        *XMLMedia     `xml:"Image,omitempty"`
	Voice        *XMLMedia     `xml:"Voice,omitempty"`
	Video        *XMLVideo     `xml:"Video,omitempty"`
	Music        *XMLMusic     `xml:"Music,omitempty"`
	ArticleCount int           `xml:"ArticleCount,omitempty"`
	Articles     *XMLNews      `xml:"Articles,omitempty"`
	TransInfo    *XMLTransInfo `xml:"TransInfo,omitempty"`
}

func (r *Reply) Bytes(from, to string) ([]byte, error) {
	r.FromUserName = wx.CDATA(from)
	r.ToUserName = wx.CDATA(to)
	r.CreateTime = time.Now().Unix() // 执行 testing 前，请注释掉

	return xml.Marshal(r)
}

// ReplyText 回复文本消息
func ReplyText(content string) event.Reply {
	return &Reply{
		MsgType: wx.CDATA(event.MsgText),
		Content: wx.CDATA(content),
	}
}

// ReplyImage 回复图片消息
func ReplyImage(mediaID string) event.Reply {
	return &Reply{
		MsgType: wx.CDATA(event.MsgImage),
		Image: &XMLMedia{
			MediaID: wx.CDATA(mediaID),
		},
	}
}

// ReplyVoice 回复语音消息
func ReplyVoice(mediaID string) event.Reply {
	return &Reply{
		MsgType: wx.CDATA(event.MsgVoice),
		Voice: &XMLMedia{
			MediaID: wx.CDATA(mediaID),
		},
	}
}

// ReplyVideo 回复视频消息
func ReplyVideo(mediaID, title, description string) event.Reply {
	return &Reply{
		MsgType: wx.CDATA(event.MsgVideo),
		Video: &XMLVideo{
			MediaID:     wx.CDATA(mediaID),
			Title:       wx.CDATA(title),
			Description: wx.CDATA(description),
		},
	}
}

// ReplyMusic 回复音乐消息
func ReplyMusic(music *XMLMusic) event.Reply {
	return &Reply{
		MsgType: wx.CDATA(event.MsgMusic),
		Music:   music,
	}
}

// ReplyNews 回复图文消息
func ReplyNews(articles ...*XMLNewsArticle) event.Reply {
	return &Reply{
		MsgType:      wx.CDATA(event.MsgNews),
		ArticleCount: len(articles),
		Articles: &XMLNews{
			Articles: articles,
		},
	}
}

// TransferToKF 消息转发到客服
func TransferToKF(kfAccount ...string) event.Reply {
	reply := &Reply{
		MsgType: wx.CDATA(event.MsgTransferToKF),
	}

	if len(kfAccount) != 0 {
		reply.TransInfo = &XMLTransInfo{KFAccount: wx.CDATA(kfAccount[0])}
	}

	return reply
}
