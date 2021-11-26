package kf

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/mock"
	"github.com/shenghui0779/gochat/offia"
	"github.com/stretchr/testify/assert"
)

func TestSendTextMsg(t *testing.T) {
	body := []byte(`{"touser":"OPENID","msgtype":"text","text":{"content":"Hello World"},"customservice":{"kf_account":"test1@kftest"}}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := offia.New("APPID", "APPSECRET")
	oa.SetClient(client)

	text := &Text{
		Content: "Hello World",
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendTextMsg("OPENID", text, "test1@kftest"))

	assert.Nil(t, err)
}

func TestSendImageMsg(t *testing.T) {
	body := []byte(`{"touser":"OPENID","msgtype":"image","image":{"media_id":"MEDIA_ID"},"customservice":{"kf_account":"test1@kftest"}}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := offia.New("APPID", "APPSECRET")
	oa.SetClient(client)

	media := &Media{
		MediaID: "MEDIA_ID",
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendImageMsg("OPENID", media, "test1@kftest"))

	assert.Nil(t, err)
}

func TestSendVoiceMsg(t *testing.T) {
	body := []byte(`{"touser":"OPENID","msgtype":"voice","voice":{"media_id":"MEDIA_ID"},"customservice":{"kf_account":"test1@kftest"}}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := offia.New("APPID", "APPSECRET")
	oa.SetClient(client)

	media := &Media{
		MediaID: "MEDIA_ID",
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendVoiceMsg("OPENID", media, "test1@kftest"))

	assert.Nil(t, err)
}

func TestSendVideoMsg(t *testing.T) {
	body := []byte(`{"touser":"OPENID","msgtype":"video","video":{"media_id":"MEDIA_ID","thumb_media_id":"THUMB_MEDIA_ID","title":"TITLE","description":"DESCRIPTION"},"customservice":{"kf_account":"test1@kftest"}}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := offia.New("APPID", "APPSECRET")
	oa.SetClient(client)

	video := &Video{
		MediaID:      "MEDIA_ID",
		ThumbMediaID: "THUMB_MEDIA_ID",
		Title:        "TITLE",
		Description:  "DESCRIPTION",
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendVideoMsg("OPENID", video, "test1@kftest"))

	assert.Nil(t, err)
}

func TestSendMusicMsg(t *testing.T) {
	body := []byte(`{"touser":"OPENID","msgtype":"music","music":{"title":"MUSIC_TITLE","description":"MUSIC_DESCRIPTION","musicurl":"MUSIC_URL","hqmusicurl":"HQ_MUSIC_URL","thumb_media_id":"THUMB_MEDIA_ID"},"customservice":{"kf_account":"test1@kftest"}}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := offia.New("APPID", "APPSECRET")
	oa.SetClient(client)

	music := &Music{
		Title:        "MUSIC_TITLE",
		Description:  "MUSIC_DESCRIPTION",
		MusicURL:     "MUSIC_URL",
		HQMusicURL:   "HQ_MUSIC_URL",
		ThumbMediaID: "THUMB_MEDIA_ID",
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendMusicMsg("OPENID", music, "test1@kftest"))

	assert.Nil(t, err)
}

func TestSendNewsMsg(t *testing.T) {
	body := []byte(`{"touser":"OPENID","msgtype":"news","news":{"articles":[{"title":"Happy Day","description":"Is Really A Happy Day","url":"URL","picurl":"PIC_URL"}]},"customservice":{"kf_account":"test1@kftest"}}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := offia.New("APPID", "APPSECRET")
	oa.SetClient(client)

	news := &News{
		Articles: []*Article{
			{
				Title:       "Happy Day",
				Description: "Is Really A Happy Day",
				URL:         "URL",
				PicURL:      "PIC_URL",
			},
		},
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendNewsMsg("OPENID", news, "test1@kftest"))

	assert.Nil(t, err)
}

func TestSendMPNewsMsg(t *testing.T) {
	body := []byte(`{"touser":"OPENID","msgtype":"mpnews","mpnews":{"media_id":"MEDIA_ID"},"customservice":{"kf_account":"test1@kftest"}}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := offia.New("APPID", "APPSECRET")
	oa.SetClient(client)

	media := &Media{
		MediaID: "MEDIA_ID",
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendMPNewsMsg("OPENID", media, "test1@kftest"))

	assert.Nil(t, err)
}

func TestSendMenuMsg(t *testing.T) {
	body := []byte(`{"touser":"OPENID","msgtype":"msgmenu","msgmenu":{"head_content":"您对本次服务是否满意呢? ","tail_content":"欢迎再次光临","list":[{"id":"101","content":"满意"},{"id":"102","content":"不满意"}]},"customservice":{"kf_account":"test1@kftest"}}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := offia.New("APPID", "APPSECRET")
	oa.SetClient(client)

	menu := &Menu{
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

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendMenuMsg("OPENID", menu, "test1@kftest"))

	assert.Nil(t, err)
}

func TestSendCardMsg(t *testing.T) {
	body := []byte(`{"touser":"OPENID","msgtype":"wxcard","wxcard":{"card_id":"123dsdajkasd231jhksad"},"customservice":{"kf_account":"test1@kftest"}}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := offia.New("APPID", "APPSECRET")
	oa.SetClient(client)

	card := &Card{
		CardID: "123dsdajkasd231jhksad",
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendCardMsg("OPENID", card, "test1@kftest"))

	assert.Nil(t, err)
}

func TestSendMinipMsg(t *testing.T) {
	body := []byte(`{"touser":"OPENID","msgtype":"miniprogrampage","miniprogrampage":{"title":"title","appid":"appid","pagepath":"pagepath","thumb_media_id":"thumb_media_id"}}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := offia.New("APPID", "APPSECRET")
	oa.SetClient(client)

	msg := &MinipPage{
		Title:        "title",
		AppID:        "appid",
		Pagepath:     "pagepath",
		ThumbMediaID: "thumb_media_id",
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendMinipMsg("OPENID", msg))

	assert.Nil(t, err)
}

func TestSetTyping(t *testing.T) {
	body := []byte(`{"touser":"OPENID","command":"Typing"}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/message/custom/typing?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := offia.New("APPID", "APPSECRET")
	oa.SetClient(client)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendTyping("OPENID", Typing))

	assert.Nil(t, err)
}
