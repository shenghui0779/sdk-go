package offia

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplyText(t *testing.T) {
	b, err := ReplyText("this is a test").Bytes("fromUser", "toUser")

	assert.Nil(t, err)

	expected := `<xml><FromUserName><![CDATA[fromUser]]></FromUserName><ToUserName><![CDATA[toUser]]></ToUserName><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[this is a test]]></Content></xml>`

	assert.Equal(t, expected, string(b))
}

func TestReplyImage(t *testing.T) {
	b, err := ReplyImage("media_id").Bytes("fromUser", "toUser")

	assert.Nil(t, err)

	expected := `<xml><FromUserName><![CDATA[fromUser]]></FromUserName><ToUserName><![CDATA[toUser]]></ToUserName><MsgType><![CDATA[image]]></MsgType><Image><MediaId><![CDATA[media_id]]></MediaId></Image></xml>`

	assert.Equal(t, expected, string(b))
}

func TestReplyVoice(t *testing.T) {
	b, err := ReplyVoice("media_id").Bytes("fromUser", "toUser")

	assert.Nil(t, err)

	expected := `<xml><FromUserName><![CDATA[fromUser]]></FromUserName><ToUserName><![CDATA[toUser]]></ToUserName><MsgType><![CDATA[voice]]></MsgType><Voice><MediaId><![CDATA[media_id]]></MediaId></Voice></xml>`

	assert.Equal(t, expected, string(b))
}

func TestReplyVideo(t *testing.T) {
	b, err := ReplyVideo("media_id", "title", "description").Bytes("fromUser", "toUser")

	assert.Nil(t, err)

	expected := `<xml><FromUserName><![CDATA[fromUser]]></FromUserName><ToUserName><![CDATA[toUser]]></ToUserName><MsgType><![CDATA[video]]></MsgType><Video><MediaId><![CDATA[media_id]]></MediaId><Title><![CDATA[title]]></Title><Description><![CDATA[description]]></Description></Video></xml>`

	assert.Equal(t, expected, string(b))
}

func TestReplyMusic(t *testing.T) {
	music := &XMLMusic{
		Title:        "TITLE",
		Description:  "DESCRIPTION",
		MusicURL:     "MUSIC_Url",
		HQMusicURL:   "HQ_MUSIC_Url",
		ThumbMediaID: "media_id",
	}

	b, err := ReplyMusic(music).Bytes("fromUser", "toUser")

	assert.Nil(t, err)

	expected := `<xml><FromUserName><![CDATA[fromUser]]></FromUserName><ToUserName><![CDATA[toUser]]></ToUserName><MsgType><![CDATA[music]]></MsgType><Music><Title><![CDATA[TITLE]]></Title><Description><![CDATA[DESCRIPTION]]></Description><MusicUrl><![CDATA[MUSIC_Url]]></MusicUrl><HQMusicUrl><![CDATA[HQ_MUSIC_Url]]></HQMusicUrl><ThumbMediaId><![CDATA[media_id]]></ThumbMediaId></Music></xml>`

	assert.Equal(t, expected, string(b))
}

func TestReplyNews(t *testing.T) {
	articles := []*XMLNewsArticle{
		{
			Title:       "title1",
			Description: "description1",
			URL:         "url",
			PicURL:      "picurl",
		},
	}

	b, err := ReplyNews(articles...).Bytes("fromUser", "toUser")

	assert.Nil(t, err)

	expected := `<xml><FromUserName><![CDATA[fromUser]]></FromUserName><ToUserName><![CDATA[toUser]]></ToUserName><MsgType><![CDATA[news]]></MsgType><ArticleCount>1</ArticleCount><Articles><item><Title><![CDATA[title1]]></Title><Description><![CDATA[description1]]></Description><Url><![CDATA[url]]></Url><PicUrl><![CDATA[picurl]]></PicUrl></item></Articles></xml>`

	assert.Equal(t, expected, string(b))
}

func TestTransferToKF(t *testing.T) {
	b, err := TransferToKF("test1@test").Bytes("fromUser", "toUser")

	assert.Nil(t, err)

	expected := `<xml><FromUserName><![CDATA[fromUser]]></FromUserName><ToUserName><![CDATA[toUser]]></ToUserName><MsgType><![CDATA[transfer_customer_service]]></MsgType><TransInfo><KfAccount><![CDATA[test1@test]]></KfAccount></TransInfo></xml>`

	assert.Equal(t, expected, string(b))
}
