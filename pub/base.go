package pub

import (
	"time"

	"github.com/iiinsomnia/gochat/utils"
	"meipian.cn/printapi/wechat"
)

type WXPub struct {
	AccountID      string
	AppID          string
	AppSecret      string
	SignToken      string
	EncodingAESKey string
}

// UseTextReplyMsg returns a new wxpub text reply msg
func (wx *WXPub) UseTextReplyMsg(openid, content string) *TextReplyMsg {
	m := &TextReplyMsg{
		ReplyMsgHeader: ReplyMsgHeader{
			ToUserName:   utils.CDATA(openid),
			FromUserName: utils.CDATA(wx.AccountID),
			CreateTime:   time.Now().Unix(),
			MsgType:      utils.CDATA("text"),
		},
	}

	m.Content = utils.CDATA(content)

	return m
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
