package pub

import (
	"encoding/base64"
	"encoding/xml"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/iiinsomnia/gochat/utils"
)

// DefaultReply 公众号默认回复
const DefaultReply = "success"

// ReplyHeader 公众号消息回复公共头
type ReplyHeader struct {
	ToUserName   utils.CDATA `xml:"ToUserName"`
	FromUserName utils.CDATA `xml:"FromUserName"`
	CreateTime   int64       `xml:"CreateTime"`
	MsgType      utils.CDATA `xml:"MsgType"`
}

// TextReply 公众号文本回复
type TextReply struct {
	XMLName xml.Name `xml:"xml"`
	ReplyHeader
	Content utils.CDATA `xml:"Content"`
}

// ImageReply 公众号图片回复
type ImageReply struct {
	XMLName xml.Name `xml:"xml"`
	ReplyHeader
	Image struct {
		MediaID utils.CDATA `xml:"MediaId"` // 通过素材管理接口上传多媒体文件得到 MediaId
	} `xml:"Image"`
}

// VoiceReply 公众号语音回复
type VoiceReply struct {
	XMLName xml.Name `xml:"xml"`
	ReplyHeader
	Voice struct {
		MediaID utils.CDATA `xml:"MediaId"` // 通过素材管理接口上传多媒体文件得到 MediaId
	} `xml:"Voice"`
}

// Video 公众号视频回复
type VideoReply struct {
	XMLName xml.Name `xml:"xml"`
	ReplyHeader
	Video struct {
		MediaID     utils.CDATA `xml:"MediaId"`               // 通过素材管理接口上传多媒体文件得到 MediaId
		Title       utils.CDATA `xml:"Title,omitempty"`       // 视频消息的标题, 可以为空
		Description utils.CDATA `xml:"Description,omitempty"` // 视频消息的描述, 可以为空
	} `xml:"Video" json:"Video"`
}

// Music 公众号音乐回复
type MusicReply struct {
	XMLName xml.Name `xml:"xml"`
	ReplyHeader
	Music struct {
		Title        utils.CDATA `xml:"Title,omitempty"`       // 音乐标题
		Description  utils.CDATA `xml:"Description,omitempty"` // 音乐描述
		MusicURL     utils.CDATA `xml:"MusicUrl,omitempty"`    // 音乐链接
		HQMusicURL   utils.CDATA `xml:"HQMusicUrl,omitempty"`  // 高质量音乐链接, WIFI环境优先使用该链接播放音乐
		ThumbMediaID utils.CDATA `xml:"ThumbMediaId"`          // 通过素材管理接口上传多媒体文件得到 ThumbMediaId
	} `xml:"Music"`
}

// Article 公众号图文
type Article struct {
	Title       utils.CDATA `xml:"Title"`       // 图文消息标题
	Description utils.CDATA `xml:"Description"` // 图文消息描述
	PicURL      utils.CDATA `xml:"PicUrl"`      // 图片链接, 支持JPG, PNG格式, 较好的效果为大图360*200, 小图200*200
	URL         utils.CDATA `xml:"Url"`         // 点击图文消息跳转链接
}

// Articles 公众号图文回复
type ArticlesReply struct {
	XMLName xml.Name `xml:"xml"`
	ReplyHeader
	ArticleCount int        `xml:"ArticleCount"`  // 图文消息个数, 限制为10条以内
	Articles     []*Article `xml:"Articles>item"` // 多条图文消息信息, 默认第一个item为大图, 注意, 如果图文数超过10, 则将会无响应
}

// TransInfo 转发客服账号
type TransInfo struct {
	KfAccount utils.CDATA `xml:"KfAccount"`
}

// Transfer2KF 公众号消息转客服
type Transfer2KF struct {
	XMLName xml.Name `xml:"xml"`
	ReplyHeader
	TransInfo *TransInfo `xml:"TransInfo,omitempty"`
}

// ReplyMsg 公众号回复消息
type ReplyMsg struct {
	XMLName      xml.Name    `xml:"xml"`
	Encrypt      utils.CDATA `xml:"Encrypt"`
	MsgSignature utils.CDATA `xml:"MsgSignature"`
	TimeStamp    int64       `xml:"TimeStamp"`
	Nonce        utils.CDATA `xml:"Nonce"`
}

// Reply 公众号回复
type Reply struct {
	*WXPub
	msg *ReplyMsg
}

func (r *Reply) encrypt(data []byte) ([]byte, error) {
	key, err := base64.StdEncoding.DecodeString(r.EncodingAESKey + "=")

	if err != nil {
		return nil, err
	}

	contentLen := len(data)
	appidOffset := 20 + contentLen

	plainText := make([]byte, appidOffset+len(r.AppID))

	copy(plainText[:16], utils.RandomStr(16))
	copy(plainText[16:20], utils.EncodeUint32ToBytes(uint32(contentLen)))
	copy(plainText[20:], data)
	copy(plainText[appidOffset:], r.AppID)

	cipherText, err := utils.AESCBCEncrypt(plainText, key)

	if err != nil {
		return nil, err
	}

	return cipherText, nil
}

func (r *Reply) build(encrypt string) *ReplyMsg {
	now := time.Now().Unix()
	nonce := utils.NonceStr()

	signItems := []string{r.SignToken, strconv.FormatInt(now, 10), nonce, encrypt}

	sort.Strings(signItems)

	msg := &ReplyMsg{
		Encrypt:      utils.CDATA(encrypt),
		MsgSignature: utils.CDATA(utils.SHA1(strings.Join(signItems, ""))),
		TimeStamp:    now,
		Nonce:        utils.CDATA(nonce),
	}

	return msg
}

// Text build wxpub text reply msg
func (r *Reply) Text(openid, content string) (*ReplyMsg, error) {
	m := &TextReply{
		ReplyHeader: ReplyHeader{
			ToUserName:   utils.CDATA(openid),
			FromUserName: utils.CDATA(r.AccountID),
			CreateTime:   time.Now().Unix(),
			MsgType:      utils.CDATA("text"),
		},
	}

	m.Content = utils.CDATA(content)

	b, err := xml.Marshal(m)

	if err != nil {
		return nil, err
	}

	// 消息加密
	cipherText, err := r.encrypt(b)

	if err != nil {
		return nil, err
	}

	return r.build(base64.StdEncoding.EncodeToString(cipherText)), nil
}

// Image build wxpub image reply msg
func (r *Reply) Image(openid, mediaID string) (*ReplyMsg, error) {
	m := &ImageReply{
		ReplyHeader: ReplyHeader{
			ToUserName:   utils.CDATA(openid),
			FromUserName: utils.CDATA(r.AccountID),
			CreateTime:   time.Now().Unix(),
			MsgType:      utils.CDATA("text"),
		},
	}

	m.Image.MediaID = utils.CDATA(mediaID)

	b, err := xml.Marshal(m)

	if err != nil {
		return nil, err
	}

	// 消息加密
	cipherText, err := r.encrypt(b)

	if err != nil {
		return nil, err
	}

	return r.build(base64.StdEncoding.EncodeToString(cipherText)), nil
}

// Voice build wxpub voice reply msg
func (r *Reply) Voice(openid, mediaID string) (*ReplyMsg, error) {
	m := &VoiceReply{
		ReplyHeader: ReplyHeader{
			ToUserName:   utils.CDATA(openid),
			FromUserName: utils.CDATA(r.AccountID),
			CreateTime:   time.Now().Unix(),
			MsgType:      utils.CDATA("text"),
		},
	}

	m.Voice.MediaID = utils.CDATA(mediaID)

	b, err := xml.Marshal(m)

	if err != nil {
		return nil, err
	}

	// 消息加密
	cipherText, err := r.encrypt(b)

	if err != nil {
		return nil, err
	}

	return r.build(base64.StdEncoding.EncodeToString(cipherText)), nil
}

// Video build wxpub video reply msg
func (r *Reply) Video(openid, mediaID, title, desc string) (*ReplyMsg, error) {
	m := &VideoReply{
		ReplyHeader: ReplyHeader{
			ToUserName:   utils.CDATA(openid),
			FromUserName: utils.CDATA(r.AccountID),
			CreateTime:   time.Now().Unix(),
			MsgType:      utils.CDATA("text"),
		},
	}

	m.Video.MediaID = utils.CDATA(mediaID)
	m.Video.Title = utils.CDATA(title)
	m.Video.Description = utils.CDATA(desc)

	b, err := xml.Marshal(m)

	if err != nil {
		return nil, err
	}

	// 消息加密
	cipherText, err := r.encrypt(b)

	if err != nil {
		return nil, err
	}

	return r.build(base64.StdEncoding.EncodeToString(cipherText)), nil
}

// Music build wxpub music reply msg
func (r *Reply) Music(openid, mediaID, title, desc, url, HQUrl string) (*ReplyMsg, error) {
	m := &MusicReply{
		ReplyHeader: ReplyHeader{
			ToUserName:   utils.CDATA(openid),
			FromUserName: utils.CDATA(r.AccountID),
			CreateTime:   time.Now().Unix(),
			MsgType:      utils.CDATA("text"),
		},
	}

	m.Music.Title = utils.CDATA(title)
	m.Music.Description = utils.CDATA(desc)
	m.Music.MusicURL = utils.CDATA(url)
	m.Music.HQMusicURL = utils.CDATA(HQUrl)
	m.Music.ThumbMediaID = utils.CDATA(mediaID)

	b, err := xml.Marshal(m)

	if err != nil {
		return nil, err
	}

	// 消息加密
	cipherText, err := r.encrypt(b)

	if err != nil {
		return nil, err
	}

	return r.build(base64.StdEncoding.EncodeToString(cipherText)), nil
}

// Articles build wxpub articles reply msg
func (r *Reply) Articles(openid string, count int, articles ...*Article) (*ReplyMsg, error) {
	m := &ArticlesReply{
		ReplyHeader: ReplyHeader{
			ToUserName:   utils.CDATA(openid),
			FromUserName: utils.CDATA(r.AccountID),
			CreateTime:   time.Now().Unix(),
			MsgType:      utils.CDATA("text"),
		},
	}

	m.ArticleCount = count
	m.Articles = articles

	b, err := xml.Marshal(m)

	if err != nil {
		return nil, err
	}

	// 消息加密
	cipherText, err := r.encrypt(b)

	if err != nil {
		return nil, err
	}

	return r.build(base64.StdEncoding.EncodeToString(cipherText)), nil
}

// Transfer2KF transfer msg to wxpub kf
func (r *Reply) Transfer2KF(openid string, kfAccount ...string) (*ReplyMsg, error) {
	m := &Transfer2KF{
		ReplyHeader: ReplyHeader{
			ToUserName:   utils.CDATA(openid),
			FromUserName: utils.CDATA(r.AccountID),
			CreateTime:   time.Now().Unix(),
			MsgType:      utils.CDATA("transfer_customer_service"),
		},
	}

	if len(kfAccount) > 0 {
		m.TransInfo = &TransInfo{KfAccount: utils.CDATA(kfAccount[0])}
	}

	b, err := xml.Marshal(m)

	if err != nil {
		return nil, err
	}

	// 消息加密
	cipherText, err := r.encrypt(b)

	if err != nil {
		return nil, err
	}

	return r.build(base64.StdEncoding.EncodeToString(cipherText)), nil
}
