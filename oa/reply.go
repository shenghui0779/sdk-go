package oa

import (
	"encoding/xml"
	"time"

	"github.com/shenghui0779/gochat/event"
	"github.com/shenghui0779/gochat/wx"
)

// ReplyHeader 公众号消息回复公共头
type ReplyHeader struct {
	XMLName      xml.Name `xml:"xml"`
	FromUserName wx.CDATA `xml:"FromUserName"`
	ToUserName   wx.CDATA `xml:"ToUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      wx.CDATA `xml:"MsgType"`
}

// TextReply 公众号文本回复
type TextReply struct {
	ReplyHeader
	Content wx.CDATA `xml:"Content"`
}

func (r *TextReply) Bytes(from, to string) ([]byte, error) {
	r.ReplyHeader.FromUserName = wx.CDATA(from)
	r.ReplyHeader.ToUserName = wx.CDATA(to)

	return xml.Marshal(r)
}

// ImageReply 公众号图片回复消息
type ImageReply struct {
	ReplyHeader
	Image struct {
		MediaID wx.CDATA `xml:"MediaId"` // 通过素材管理接口上传多媒体文件得到 MediaId
	} `xml:"Image"`
}

func (r *ImageReply) Bytes(from, to string) ([]byte, error) {
	r.ReplyHeader.FromUserName = wx.CDATA(from)
	r.ReplyHeader.ToUserName = wx.CDATA(to)

	return xml.Marshal(r)
}

// VoiceReply 公众号语音回复
type VoiceReply struct {
	ReplyHeader
	Voice struct {
		MediaID wx.CDATA `xml:"MediaId"` // 通过素材管理接口上传多媒体文件得到 MediaId
	} `xml:"Voice"`
}

func (r *VoiceReply) Bytes(from, to string) ([]byte, error) {
	r.ReplyHeader.FromUserName = wx.CDATA(from)
	r.ReplyHeader.ToUserName = wx.CDATA(to)

	return xml.Marshal(r)
}

// VideoReply 公众号视频回复
type VideoReply struct {
	ReplyHeader
	Video struct {
		MediaID     wx.CDATA `xml:"MediaId"`               // 通过素材管理接口上传多媒体文件得到 MediaId
		Title       wx.CDATA `xml:"Title,omitempty"`       // 视频消息的标题, 可以为空
		Description wx.CDATA `xml:"Description,omitempty"` // 视频消息的描述, 可以为空
	} `xml:"Video" json:"Video"`
}

func (r *VideoReply) Bytes(from, to string) ([]byte, error) {
	r.ReplyHeader.FromUserName = wx.CDATA(from)
	r.ReplyHeader.ToUserName = wx.CDATA(to)

	return xml.Marshal(r)
}

// MusicReply 公众号音乐回复
type MusicReply struct {
	ReplyHeader
	Music struct {
		Title        wx.CDATA `xml:"Title,omitempty"`       // 音乐标题
		Description  wx.CDATA `xml:"Description,omitempty"` // 音乐描述
		MusicURL     wx.CDATA `xml:"MusicUrl,omitempty"`    // 音乐链接
		HQMusicURL   wx.CDATA `xml:"HQMusicUrl,omitempty"`  // 高质量音乐链接, WIFI环境优先使用该链接播放音乐
		ThumbMediaID wx.CDATA `xml:"ThumbMediaId"`          // 通过素材管理接口上传多媒体文件得到 ThumbMediaId
	} `xml:"Music"`
}

func (r *MusicReply) Bytes(from, to string) ([]byte, error) {
	r.ReplyHeader.FromUserName = wx.CDATA(from)
	r.ReplyHeader.ToUserName = wx.CDATA(to)

	return xml.Marshal(r)
}

// NewsReply 公众号图文回复
type NewsReply struct {
	ReplyHeader
	ArticleCount int        `xml:"ArticleCount"`  // 图文消息个数, 限制为10条以内
	Articles     []*Article `xml:"Articles>item"` // 多条图文消息信息, 默认第一个item为大图, 注意, 如果图文数超过10, 则将会无响应
}

// Article 公众号图文
type Article struct {
	Title       wx.CDATA `xml:"Title"`       // 图文消息标题
	Description wx.CDATA `xml:"Description"` // 图文消息描述
	PicURL      wx.CDATA `xml:"PicUrl"`      // 图片链接, 支持JPG, PNG格式, 较好的效果为大图360*200, 小图200*200
	URL         wx.CDATA `xml:"Url"`         // 点击图文消息跳转链接
}

func (r *NewsReply) Bytes(from, to string) ([]byte, error) {
	r.ReplyHeader.FromUserName = wx.CDATA(from)
	r.ReplyHeader.ToUserName = wx.CDATA(to)

	return xml.Marshal(r)
}

// Transfer2KFReply 公众号消息转客服
type Transfer2KFReply struct {
	ReplyHeader
	TransInfo *TransInfo `xml:"TransInfo,omitempty"`
}

// TransInfo 转发客服账号
type TransInfo struct {
	KFAccount wx.CDATA `xml:"KfAccount"`
}

func (r *Transfer2KFReply) Bytes(from, to string) ([]byte, error) {
	r.ReplyHeader.FromUserName = wx.CDATA(from)
	r.ReplyHeader.ToUserName = wx.CDATA(to)

	return xml.Marshal(r)
}

// NewTextReply 回复文本消息
func NewTextReply(content string) event.Reply {
	r := &TextReply{
		ReplyHeader: ReplyHeader{
			CreateTime: time.Now().Unix(),
			MsgType:    wx.CDATA("text"),
		},
	}

	r.Content = wx.CDATA(content)

	return r
}

// NewImageReply 回复图片消息
func NewImageReply(mediaID string) event.Reply {
	r := &ImageReply{
		ReplyHeader: ReplyHeader{
			CreateTime: time.Now().Unix(),
			MsgType:    wx.CDATA("image"),
		},
	}

	r.Image.MediaID = wx.CDATA(mediaID)

	return r
}

// NewVoiceReply 回复语音消息
func NewVoiceReply(mediaID string) event.Reply {
	r := &VoiceReply{
		ReplyHeader: ReplyHeader{
			CreateTime: time.Now().Unix(),
			MsgType:    wx.CDATA("voice"),
		},
	}

	r.Voice.MediaID = wx.CDATA(mediaID)

	return r
}

// NewVideoReply 回复视频消息
func NewVideoReply(mediaID, title, desc string) event.Reply {
	r := &VideoReply{
		ReplyHeader: ReplyHeader{
			CreateTime: time.Now().Unix(),
			MsgType:    wx.CDATA("video"),
		},
	}

	r.Video.MediaID = wx.CDATA(mediaID)
	r.Video.Title = wx.CDATA(title)
	r.Video.Description = wx.CDATA(desc)

	return r
}

// NewMusicReply 回复音乐消息
func NewMusicReply(thumbMediaID, title, desc, musicURL, HQMusicURL string) event.Reply {
	r := &MusicReply{
		ReplyHeader: ReplyHeader{
			CreateTime: time.Now().Unix(),
			MsgType:    wx.CDATA("music"),
		},
	}

	r.Music.Title = wx.CDATA(title)
	r.Music.Description = wx.CDATA(desc)
	r.Music.MusicURL = wx.CDATA(musicURL)
	r.Music.HQMusicURL = wx.CDATA(HQMusicURL)
	r.Music.ThumbMediaID = wx.CDATA(thumbMediaID)

	return r
}

// NewNewsReply 回复图文消息
func NewNewsReply(count int, articles ...*Article) event.Reply {
	r := &NewsReply{
		ReplyHeader: ReplyHeader{
			CreateTime: time.Now().Unix(),
			MsgType:    wx.CDATA("news"),
		},
	}

	r.ArticleCount = count
	r.Articles = articles

	return r
}

// NewTransfer2KFReply 消息转发到客服
func NewTransfer2KFReply(kfAccount ...string) event.Reply {
	r := &Transfer2KFReply{
		ReplyHeader: ReplyHeader{
			CreateTime: time.Now().Unix(),
			MsgType:    wx.CDATA("transfer_customer_service"),
		},
	}

	if len(kfAccount) != 0 {
		r.TransInfo = &TransInfo{KFAccount: wx.CDATA(kfAccount[0])}
	}

	return r
}
