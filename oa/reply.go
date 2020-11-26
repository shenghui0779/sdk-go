package oa

import (
	"encoding/xml"
	"time"

	"github.com/shenghui0779/gochat/helpers"
)

// DefaultReply 公众号默认回复
const DefaultReply = "success"

type Reply interface {
	Bytes(from, to string) ([]byte, error)
}

// ReplyHeader 公众号消息回复公共头
type ReplyHeader struct {
	XMLName      xml.Name      `xml:"xml"`
	FromUserName helpers.CDATA `xml:"FromUserName"`
	ToUserName   helpers.CDATA `xml:"ToUserName"`
	CreateTime   int64         `xml:"CreateTime"`
	MsgType      helpers.CDATA `xml:"MsgType"`
}

// TextReply 公众号文本回复
type TextReply struct {
	ReplyHeader
	Content helpers.CDATA `xml:"Content"`
}

func (r *TextReply) Bytes(from, to string) ([]byte, error) {
	r.ReplyHeader.FromUserName = helpers.CDATA(from)
	r.ReplyHeader.ToUserName = helpers.CDATA(to)

	return xml.Marshal(r)
}

// ImageReply 公众号图片回复消息
type ImageReply struct {
	ReplyHeader
	Image struct {
		MediaID helpers.CDATA `xml:"MediaId"` // 通过素材管理接口上传多媒体文件得到 MediaId
	} `xml:"Image"`
}

func (r *ImageReply) Bytes(from, to string) ([]byte, error) {
	r.ReplyHeader.FromUserName = helpers.CDATA(from)
	r.ReplyHeader.ToUserName = helpers.CDATA(to)

	return xml.Marshal(r)
}

// VoiceReply 公众号语音回复
type VoiceReply struct {
	ReplyHeader
	Voice struct {
		MediaID helpers.CDATA `xml:"MediaId"` // 通过素材管理接口上传多媒体文件得到 MediaId
	} `xml:"Voice"`
}

func (r *VoiceReply) Bytes(from, to string) ([]byte, error) {
	r.ReplyHeader.FromUserName = helpers.CDATA(from)
	r.ReplyHeader.ToUserName = helpers.CDATA(to)

	return xml.Marshal(r)
}

// VideoReply 公众号视频回复
type VideoReply struct {
	ReplyHeader
	Video struct {
		MediaID     helpers.CDATA `xml:"MediaId"`               // 通过素材管理接口上传多媒体文件得到 MediaId
		Title       helpers.CDATA `xml:"Title,omitempty"`       // 视频消息的标题, 可以为空
		Description helpers.CDATA `xml:"Description,omitempty"` // 视频消息的描述, 可以为空
	} `xml:"Video" json:"Video"`
}

func (r *VideoReply) Bytes(from, to string) ([]byte, error) {
	r.ReplyHeader.FromUserName = helpers.CDATA(from)
	r.ReplyHeader.ToUserName = helpers.CDATA(to)

	return xml.Marshal(r)
}

// MusicReply 公众号音乐回复
type MusicReply struct {
	ReplyHeader
	Music struct {
		Title        helpers.CDATA `xml:"Title,omitempty"`       // 音乐标题
		Description  helpers.CDATA `xml:"Description,omitempty"` // 音乐描述
		MusicURL     helpers.CDATA `xml:"MusicUrl,omitempty"`    // 音乐链接
		HQMusicURL   helpers.CDATA `xml:"HQMusicUrl,omitempty"`  // 高质量音乐链接, WIFI环境优先使用该链接播放音乐
		ThumbMediaID helpers.CDATA `xml:"ThumbMediaId"`          // 通过素材管理接口上传多媒体文件得到 ThumbMediaId
	} `xml:"Music"`
}

func (r *MusicReply) Bytes(from, to string) ([]byte, error) {
	r.ReplyHeader.FromUserName = helpers.CDATA(from)
	r.ReplyHeader.ToUserName = helpers.CDATA(to)

	return xml.Marshal(r)
}

// NewsReply 公众号图文回复
type NewsReply struct {
	ReplyHeader
	ArticleCount int       `xml:"ArticleCount"`  // 图文消息个数, 限制为10条以内
	Articles     []Article `xml:"Articles>item"` // 多条图文消息信息, 默认第一个item为大图, 注意, 如果图文数超过10, 则将会无响应
}

// Article 公众号图文
type Article struct {
	Title       helpers.CDATA `xml:"Title"`       // 图文消息标题
	Description helpers.CDATA `xml:"Description"` // 图文消息描述
	PicURL      helpers.CDATA `xml:"PicUrl"`      // 图片链接, 支持JPG, PNG格式, 较好的效果为大图360*200, 小图200*200
	URL         helpers.CDATA `xml:"Url"`         // 点击图文消息跳转链接
}

func (r *NewsReply) Bytes(from, to string) ([]byte, error) {
	r.ReplyHeader.FromUserName = helpers.CDATA(from)
	r.ReplyHeader.ToUserName = helpers.CDATA(to)

	return xml.Marshal(r)
}

// Transfer2KFReply 公众号消息转客服
type Transfer2KFReply struct {
	ReplyHeader
	TransInfo *TransInfo `xml:"TransInfo,omitempty"`
}

// TransInfo 转发客服账号
type TransInfo struct {
	KfAccount helpers.CDATA `xml:"KfAccount"`
}

func (r *Transfer2KFReply) Bytes(from, to string) ([]byte, error) {
	r.ReplyHeader.FromUserName = helpers.CDATA(from)
	r.ReplyHeader.ToUserName = helpers.CDATA(to)

	return xml.Marshal(r)
}

// NewTextReply returns text reply
func NewTextReply(content string) Reply {
	r := &TextReply{
		ReplyHeader: ReplyHeader{
			CreateTime: time.Now().Unix(),
			MsgType:    helpers.CDATA("text"),
		},
	}

	r.Content = helpers.CDATA(content)

	return r
}

// NewImageReply returns image reply
func NewImageReply(mediaID string) Reply {
	r := &ImageReply{
		ReplyHeader: ReplyHeader{
			CreateTime: time.Now().Unix(),
			MsgType:    helpers.CDATA("image"),
		},
	}

	r.Image.MediaID = helpers.CDATA(mediaID)

	return r
}

// NewVoiceReply returns voice reply
func NewVoiceReply(mediaID string) Reply {
	r := &VoiceReply{
		ReplyHeader: ReplyHeader{
			CreateTime: time.Now().Unix(),
			MsgType:    helpers.CDATA("voice"),
		},
	}

	r.Voice.MediaID = helpers.CDATA(mediaID)

	return r
}

// NewVideoReply returns video reply
func NewVideoReply(mediaID, title, desc string) Reply {
	r := &VideoReply{
		ReplyHeader: ReplyHeader{
			CreateTime: time.Now().Unix(),
			MsgType:    helpers.CDATA("video"),
		},
	}

	r.Video.MediaID = helpers.CDATA(mediaID)
	r.Video.Title = helpers.CDATA(title)
	r.Video.Description = helpers.CDATA(desc)

	return r
}

// NewMusicReply returns music reply
func NewMusicReply(mediaID, title, desc, url, HQUrl string) Reply {
	r := &MusicReply{
		ReplyHeader: ReplyHeader{
			CreateTime: time.Now().Unix(),
			MsgType:    helpers.CDATA("music"),
		},
	}

	r.Music.Title = helpers.CDATA(title)
	r.Music.Description = helpers.CDATA(desc)
	r.Music.MusicURL = helpers.CDATA(url)
	r.Music.HQMusicURL = helpers.CDATA(HQUrl)
	r.Music.ThumbMediaID = helpers.CDATA(mediaID)

	return r
}

// NewArticleReply returns article reply
func NewArticleReply(count int, articles ...Article) Reply {
	r := &NewsReply{
		ReplyHeader: ReplyHeader{
			CreateTime: time.Now().Unix(),
			MsgType:    helpers.CDATA("news"),
		},
	}

	r.ArticleCount = count
	r.Articles = articles

	return r
}

// NewTransfer2KFReply returns transfer to kf reply
func NewTransfer2KFReply(kfAccount ...string) Reply {
	r := &Transfer2KFReply{
		ReplyHeader: ReplyHeader{
			CreateTime: time.Now().Unix(),
			MsgType:    helpers.CDATA("transfer_customer_service"),
		},
	}

	if len(kfAccount) > 0 {
		r.TransInfo = &TransInfo{KfAccount: helpers.CDATA(kfAccount[0])}
	}

	return r
}

// ReplyMessage 公众号回复
type ReplyMessage struct {
	XMLName      xml.Name      `xml:"xml"`
	Encrypt      helpers.CDATA `xml:"Encrypt"`
	MsgSignature helpers.CDATA `xml:"MsgSignature"`
	TimeStamp    int64         `xml:"TimeStamp"`
	Nonce        helpers.CDATA `xml:"Nonce"`
}
