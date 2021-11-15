package kf

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/wx"
	"github.com/stretchr/testify/assert"
)

func TestSendTextMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", []byte(`{"customservice":{"kf_account":"test1@kftest"},"msgtype":"text","text":{"content":"Hello World"},"touser":"OPENID"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendTextMessage("OPENID", "Hello World", "test1@kftest"))

	assert.Nil(t, err)
}

func TestSendImageMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", []byte(`{"customservice":{"kf_account":"test1@kftest"},"image":{"media_id":"MEDIA_ID"},"msgtype":"image","touser":"OPENID"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendImageMessage("OPENID", "MEDIA_ID", "test1@kftest"))

	assert.Nil(t, err)
}

func TestSendVoiceMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", []byte(`{"customservice":{"kf_account":"test1@kftest"},"msgtype":"voice","touser":"OPENID","voice":{"media_id":"MEDIA_ID"}}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendVoiceMessage("OPENID", "MEDIA_ID", "test1@kftest"))

	assert.Nil(t, err)
}

func TestSendVideoMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", []byte(`{"customservice":{"kf_account":"test1@kftest"},"msgtype":"video","touser":"OPENID","video":{"media_id":"MEDIA_ID","thumb_media_id":"THUMB_MEDIA_ID","title":"TITLE","description":"DESCRIPTION"}}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	msg := &VideoMessage{
		MediaID:      "MEDIA_ID",
		ThumbMediaID: "THUMB_MEDIA_ID",
		Title:        "TITLE",
		Description:  "DESCRIPTION",
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendVideoMessage("OPENID", msg, "test1@kftest"))

	assert.Nil(t, err)
}

func TestSendMusicMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", []byte(`{"customservice":{"kf_account":"test1@kftest"},"msgtype":"music","music":{"title":"MUSIC_TITLE","description":"MUSIC_DESCRIPTION","musicurl":"MUSIC_URL","hqmusicurl":"HQ_MUSIC_URL","thumb_media_id":"THUMB_MEDIA_ID"},"touser":"OPENID"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	msg := &MusicMessage{
		Title:        "MUSIC_TITLE",
		Description:  "MUSIC_DESCRIPTION",
		MusicURL:     "MUSIC_URL",
		HQMusicURL:   "HQ_MUSIC_URL",
		ThumbMediaID: "THUMB_MEDIA_ID",
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendMusicMessage("OPENID", msg, "test1@kftest"))

	assert.Nil(t, err)
}

func TestSendNewsMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", []byte(`{"customservice":{"kf_account":"test1@kftest"},"msgtype":"news","news":{"articles":[{"title":"Happy Day","description":"Is Really A Happy Day","url":"URL","picurl":"PIC_URL"}]},"touser":"OPENID"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	articles := []*Article{
		{
			Title:       "Happy Day",
			Description: "Is Really A Happy Day",
			URL:         "URL",
			PicURL:      "PIC_URL",
		},
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendNewsMessage("OPENID", articles, "test1@kftest"))

	assert.Nil(t, err)
}

func TestSendMPNewsMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", []byte(`{"customservice":{"kf_account":"test1@kftest"},"mpnews":{"media_id":"MEDIA_ID"},"msgtype":"mpnews","touser":"OPENID"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendMPNewsMessage("OPENID", "MEDIA_ID", "test1@kftest"))

	assert.Nil(t, err)
}

func TestSendMenuMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", []byte(`{"customservice":{"kf_account":"test1@kftest"},"msgmenu":{"head_content":"您对本次服务是否满意呢? ","tail_content":"欢迎再次光临","list":[{"id":"101","content":"满意"},{"id":"102","content":"不满意"}]},"msgtype":"msgmenu","touser":"OPENID"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	msg := &MenuMessage{
		HeadContent: "您对本次服务是否满意呢? ",
		TailContent: "欢迎再次光临",
		List: []*MenuOption{
			{
				ID:      "101",
				Content: "满意",
			},
			{
				ID:      "102",
				Content: "不满意",
			},
		},
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendMenuMessage("OPENID", msg, "test1@kftest"))

	assert.Nil(t, err)
}

func TestSendCardMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", []byte(`{"customservice":{"kf_account":"test1@kftest"},"msgtype":"wxcard","touser":"OPENID","wxcard":{"card_id":"123dsdajkasd231jhksad"}}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendCardMessage("OPENID", "123dsdajkasd231jhksad", "test1@kftest"))

	assert.Nil(t, err)
}

func TestSendMinipMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", []byte(`{"miniprogrampage":{"title":"title","appid":"appid","pagepath":"pagepath","thumb_media_id":"thumb_media_id"},"msgtype":"miniprogrampage","touser":"OPENID"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	msg := &MinipMessage{
		Title:        "title",
		AppID:        "appid",
		Pagepath:     "pagepath",
		ThumbMediaID: "thumb_media_id",
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendMinipMessage("OPENID", msg))

	assert.Nil(t, err)
}

func TestSetTyping(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/custom/typing?access_token=ACCESS_TOKEN", []byte(`{"command":"Typing","touser":"OPENID"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SetTyping("OPENID", Typing))

	assert.Nil(t, err)
}
