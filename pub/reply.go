package pub

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/iiinsomnia/yiigo"
	"go.uber.org/zap"
	"meipian.cn/printapi/wechat"
	"meipian.cn/printapi/wechat/utils"
)

// DefaultReply 公众号默认自动回复
const DefaultReply = "success"

// ReplyMsg 公众号回复消息
type ReplyMsg interface {
	// Encrypt 消息加密
	Encrypt() (string, error)
}

// ReplyMsgHeader 公众号消息回复公共头
type ReplyMsgHeader struct {
	ToUserName   utils.CDATA `xml:"ToUserName"`
	FromUserName utils.CDATA `xml:"FromUserName"`
	CreateTime   int64       `xml:"CreateTime"`
	MsgType      utils.CDATA `xml:"MsgType"`
}

// TextReplyMsg 公众号文本回复消息
type TextReplyMsg struct {
	XMLName xml.Name `xml:"xml"`
	ReplyMsgHeader
	Content utils.CDATA `xml:"Content"`
}

// Encrypt 消息加密
func (m *TextReplyMsg) Encrypt() (string, error) {
	b, err := xml.Marshal(m)

	if err != nil {
		yiigo.Logger.Error("marshal reply wxpub text error", zap.String("error", err.Error()))

		return "", err
	}

	cipherText, err := encryptReply(b)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// NewTextReplyMsg returns a new wxpub text reply msg
func NewTextReplyMsg(openid, content string) *TextReplyMsg {
	settings := wechat.GetSettingsWithChannel(wechat.WXPub)

	m := &TextReplyMsg{
		ReplyMsgHeader: ReplyMsgHeader{
			ToUserName:   utils.CDATA(openid),
			FromUserName: utils.CDATA(settings.AccountID),
			CreateTime:   time.Now().Unix(),
			MsgType:      utils.CDATA("text"),
		},
	}

	m.Content = utils.CDATA(content)

	return m
}

// ImageReplyMsg 公众号图片回复消息
type ImageReplyMsg struct {
	XMLName xml.Name `xml:"xml"`
	ReplyMsgHeader
	Image struct {
		MediaID utils.CDATA `xml:"MediaId"` // 通过素材管理接口上传多媒体文件得到 MediaId
	} `xml:"Image"`
}

// Encrypt 消息加密
func (m *ImageReplyMsg) Encrypt() (string, error) {
	b, err := xml.Marshal(m)

	if err != nil {
		yiigo.Logger.Error("marshal reply wxpub image error", zap.String("error", err.Error()))

		return "", err
	}

	cipherText, err := encryptReply(b)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// NewImageReplyMsg returns a new wxpub image reply msg
func NewImageReplyMsg(openid, mediaID string) *ImageReplyMsg {
	settings := wechat.GetSettingsWithChannel(wechat.WXPub)

	m := &ImageReplyMsg{
		ReplyMsgHeader: ReplyMsgHeader{
			ToUserName:   utils.CDATA(openid),
			FromUserName: utils.CDATA(settings.AccountID),
			CreateTime:   time.Now().Unix(),
			MsgType:      utils.CDATA("text"),
		},
	}

	m.Image.MediaID = utils.CDATA(mediaID)

	return m
}

// VoiceReplyMsg 公众号语音回复消息
type VoiceReplyMsg struct {
	XMLName xml.Name `xml:"xml"`
	ReplyMsgHeader
	Voice struct {
		MediaID utils.CDATA `xml:"MediaId"` // 通过素材管理接口上传多媒体文件得到 MediaId
	} `xml:"Voice"`
}

// Encrypt 消息加密
func (m *VoiceReplyMsg) Encrypt() (string, error) {
	b, err := xml.Marshal(m)

	if err != nil {
		yiigo.Logger.Error("marshal reply wxpub voice error", zap.String("error", err.Error()))

		return "", err
	}

	cipherText, err := encryptReply(b)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// NewVoiceReplyMsg returns a new wxpub voice reply msg
func NewVoiceReplyMsg(openid, mediaID string) *VoiceReplyMsg {
	settings := wechat.GetSettingsWithChannel(wechat.WXPub)

	m := &VoiceReplyMsg{
		ReplyMsgHeader: ReplyMsgHeader{
			ToUserName:   utils.CDATA(openid),
			FromUserName: utils.CDATA(settings.AccountID),
			CreateTime:   time.Now().Unix(),
			MsgType:      utils.CDATA("text"),
		},
	}

	m.Voice.MediaID = utils.CDATA(mediaID)

	return m
}

// Video 公众号视频回复消息
type VideoReplyMsg struct {
	XMLName xml.Name `xml:"xml"`
	ReplyMsgHeader
	Video struct {
		MediaID     utils.CDATA `xml:"MediaId"`               // 通过素材管理接口上传多媒体文件得到 MediaId
		Title       utils.CDATA `xml:"Title,omitempty"`       // 视频消息的标题, 可以为空
		Description utils.CDATA `xml:"Description,omitempty"` // 视频消息的描述, 可以为空
	} `xml:"Video" json:"Video"`
}

// Encrypt 消息加密
func (m *VideoReplyMsg) Encrypt() (string, error) {
	b, err := xml.Marshal(m)

	if err != nil {
		yiigo.Logger.Error("marshal reply wxpub video error", zap.String("error", err.Error()))

		return "", err
	}

	cipherText, err := encryptReply(b)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// NewVideoReplyMsg returns a new wxpub video reply msg
func NewVideoReplyMsg(openid, mediaID, title, desc string) *VideoReplyMsg {
	settings := wechat.GetSettingsWithChannel(wechat.WXPub)

	m := &VideoReplyMsg{
		ReplyMsgHeader: ReplyMsgHeader{
			ToUserName:   utils.CDATA(openid),
			FromUserName: utils.CDATA(settings.AccountID),
			CreateTime:   time.Now().Unix(),
			MsgType:      utils.CDATA("text"),
		},
	}

	m.Video.MediaID = utils.CDATA(mediaID)
	m.Video.Title = utils.CDATA(title)
	m.Video.Description = utils.CDATA(desc)

	return m
}

// Music 公众号音乐回复消息
type MusicReplyMsg struct {
	XMLName xml.Name `xml:"xml"`
	ReplyMsgHeader
	Music struct {
		Title        utils.CDATA `xml:"Title,omitempty"`       // 音乐标题
		Description  utils.CDATA `xml:"Description,omitempty"` // 音乐描述
		MusicURL     utils.CDATA `xml:"MusicUrl,omitempty"`    // 音乐链接
		HQMusicURL   utils.CDATA `xml:"HQMusicUrl,omitempty"`  // 高质量音乐链接, WIFI环境优先使用该链接播放音乐
		ThumbMediaID utils.CDATA `xml:"ThumbMediaId"`          // 通过素材管理接口上传多媒体文件得到 ThumbMediaId
	} `xml:"Music"`
}

// Encrypt 消息加密
func (m *MusicReplyMsg) Encrypt() (string, error) {
	b, err := xml.Marshal(m)

	if err != nil {
		yiigo.Logger.Error("marshal reply wxpub music error", zap.String("error", err.Error()))

		return "", err
	}

	cipherText, err := encryptReply(b)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// NewMusicReplyMsg returns a new wxpub music reply msg
func NewMusicReplyMsg(openid, mediaID, title, desc, url, HQUrl string) *MusicReplyMsg {
	settings := wechat.GetSettingsWithChannel(wechat.WXPub)

	m := &MusicReplyMsg{
		ReplyMsgHeader: ReplyMsgHeader{
			ToUserName:   utils.CDATA(openid),
			FromUserName: utils.CDATA(settings.AccountID),
			CreateTime:   time.Now().Unix(),
			MsgType:      utils.CDATA("text"),
		},
	}

	m.Music.Title = utils.CDATA(title)
	m.Music.Description = utils.CDATA(desc)
	m.Music.MusicURL = utils.CDATA(url)
	m.Music.HQMusicURL = utils.CDATA(HQUrl)
	m.Music.ThumbMediaID = utils.CDATA(mediaID)

	return m
}

// Article 公众号图文
type Article struct {
	Title       utils.CDATA `xml:"Title"`       // 图文消息标题
	Description utils.CDATA `xml:"Description"` // 图文消息描述
	PicURL      utils.CDATA `xml:"PicUrl"`      // 图片链接, 支持JPG, PNG格式, 较好的效果为大图360*200, 小图200*200
	URL         utils.CDATA `xml:"Url"`         // 点击图文消息跳转链接
}

// Articles 公众号图文回复消息
type ArticlesReplyMsg struct {
	XMLName xml.Name `xml:"xml"`
	ReplyMsgHeader
	ArticleCount int        `xml:"ArticleCount"`  // 图文消息个数, 限制为10条以内
	Articles     []*Article `xml:"Articles>item"` // 多条图文消息信息, 默认第一个item为大图, 注意, 如果图文数超过10, 则将会无响应
}

// Encrypt 消息加密
func (m *ArticlesReplyMsg) Encrypt() (string, error) {
	b, err := xml.Marshal(m)

	if err != nil {
		yiigo.Logger.Error("marshal reply wxpub articles error", zap.String("error", err.Error()))

		return "", err
	}

	cipherText, err := encryptReply(b)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// NewArticlesReplyMsg returns a new wxpub articles reply msg
func NewArticlesReplyMsg(openid string, count int, articles ...*Article) *ArticlesReplyMsg {
	settings := wechat.GetSettingsWithChannel(wechat.WXPub)

	m := &ArticlesReplyMsg{
		ReplyMsgHeader: ReplyMsgHeader{
			ToUserName:   utils.CDATA(openid),
			FromUserName: utils.CDATA(settings.AccountID),
			CreateTime:   time.Now().Unix(),
			MsgType:      utils.CDATA("text"),
		},
	}

	m.ArticleCount = count
	m.Articles = articles

	return m
}

// TransInfo 转发客服账号
type TransInfo struct {
	KfAccount utils.CDATA `xml:"KfAccount"`
}

// Transfer2KF 公众号消息转客服
type Transfer2KF struct {
	XMLName xml.Name `xml:"xml"`
	ReplyMsgHeader
	TransInfo *TransInfo `xml:"TransInfo,omitempty"`
}

// Encrypt 消息加密
func (m *Transfer2KF) Encrypt() (string, error) {
	b, err := xml.Marshal(m)

	if err != nil {
		yiigo.Logger.Error("marshal reply wxpub transfer to kf error", zap.String("error", err.Error()))

		return "", err
	}

	cipherText, err := encryptReply(b)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// NewTransfer2KF transfer msg to wxpub kf
func NewTransfer2KF(openid string, kfAccount ...string) *Transfer2KF {
	settings := wechat.GetSettingsWithChannel(wechat.WXPub)

	m := &Transfer2KF{
		ReplyMsgHeader: ReplyMsgHeader{
			ToUserName:   utils.CDATA(openid),
			FromUserName: utils.CDATA(settings.AccountID),
			CreateTime:   time.Now().Unix(),
			MsgType:      utils.CDATA("transfer_customer_service"),
		},
	}

	if len(kfAccount) > 0 {
		m.TransInfo = &TransInfo{KfAccount: utils.CDATA(kfAccount[0])}
	}

	return m
}

func encryptReply(data []byte) ([]byte, error) {
	settings := wechat.GetSettingsWithChannel(wechat.WXPub)

	key, err := base64.StdEncoding.DecodeString(settings.EncodingAESKey + "=")

	if err != nil {
		yiigo.Logger.Error("base64.decode wxpub EncodingAESKey error", zap.String("error", err.Error()))

		return nil, err
	}

	contentLen := len(data)
	appidOffset := 20 + contentLen

	plainText := make([]byte, appidOffset+len(settings.AppID))

	copy(plainText[:16], utils.RandomStr(16))
	copy(plainText[16:20], utils.EncodeUint32ToBytes(uint32(contentLen)))
	copy(plainText[20:], data)
	copy(plainText[appidOffset:], settings.AppID)

	cipherText, err := utils.AESCBCEncrypt(plainText, key)

	if err != nil {
		yiigo.Logger.Error("encrypt wxpub reply error", zap.String("error", err.Error()), zap.ByteString("plain_text", plainText))

		return nil, err
	}

	return cipherText, nil
}

// Reply 公众号回复
type Reply struct {
	XMLName      xml.Name    `xml:"xml"`
	Encrypt      utils.CDATA `xml:"Encrypt"`
	MsgSignature utils.CDATA `xml:"MsgSignature"`
	TimeStamp    int64       `xml:"TimeStamp"`
	Nonce        utils.CDATA `xml:"Nonce"`
}

// NewReply returns a new wxpub reply
func NewReply(msg ReplyMsg) (*Reply, error) {
	encrypt, err := msg.Encrypt()

	if err != nil {
		return nil, err
	}

	settings := wechat.GetSettingsWithChannel(wechat.WXPub)

	now := time.Now().Unix()
	nonce := utils.NonceStr()

	signItems := []string{settings.SignToken, strconv.FormatInt(now, 10), nonce, encrypt}

	sort.Strings(signItems)

	h := sha1.New()
	h.Write([]byte(strings.Join(signItems, "")))

	reply := &Reply{
		Encrypt:      utils.CDATA(encrypt),
		MsgSignature: utils.CDATA(hex.EncodeToString(h.Sum(nil))),
		TimeStamp:    now,
		Nonce:        utils.CDATA(nonce),
	}

	return reply, nil
}
