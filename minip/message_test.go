package minip

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/chenghonour/gochat/mock"
	"github.com/chenghonour/gochat/wx"
)

func TestSendUniformMessage(t *testing.T) {
	body := []byte(`{"touser":"OPENID","mp_template_msg":{"appid":"APPID","template_id":"TEMPLATE_ID","url":"http://weixin.qq.com/download","miniprogram":{"appid":"xiaochengxuappid12345","pagepath":"index?foo=bar"},"data":{"first":{"value":"恭喜你购买成功！","color":"#173177"},"keyword1":{"value":"巧克力","color":"#173177"},"keyword2":{"value":"39.8元","color":"#173177"},"keyword3":{"value":"2014年9月22日","color":"#173177"},"remark":{"value":"欢迎再次购买！","color":"#173177"}}}}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/uniform_send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	mp := New("APPID", "APPSECRET")
	mp.SetClient(wx.WithHTTPClient(client))

	msg := &TemplateMsg{
		AppID:      "APPID",
		TemplateID: "TEMPLATE_ID",
		URL:        "http://weixin.qq.com/download",
		Minip: &MsgMinip{
			AppID:    "xiaochengxuappid12345",
			PagePath: "index?foo=bar",
		},
		Data: MsgTemplData{
			"first": {
				Value: "恭喜你购买成功！",
				Color: "#173177",
			},
			"keyword1": {
				Value: "巧克力",
				Color: "#173177",
			},
			"keyword2": {
				Value: "39.8元",
				Color: "#173177",
			},
			"keyword3": {
				Value: "2014年9月22日",
				Color: "#173177",
			},
			"remark": {
				Value: "欢迎再次购买！",
				Color: "#173177",
			},
		},
	}

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", SendUniformMsg("OPENID", msg))

	assert.Nil(t, err)
}

func TestSendSubscribeMessage(t *testing.T) {
	body := []byte(`{"touser":"OPENID","template_id":"TEMPLATE_ID","page":"index","miniprogram_state":"developer","lang":"zh_CN","data":{"date01":{"value":"2015年01月05日"},"number01":{"value":"339208499"},"site01":{"value":"TIT创意园"},"site02":{"value":"广州市新港中路397号"}}}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	mp := New("APPID", "APPSECRET")
	mp.SetClient(wx.WithHTTPClient(client))

	msg := &SubscribeMsg{
		ToUser:     "OPENID",
		TemplateID: "TEMPLATE_ID",
		Page:       "index",
		MinipState: MinipDeveloper,
		Lang:       "zh_CN",
		Data: MsgTemplData{
			"date01": {
				Value: "2015年01月05日",
			},
			"number01": {
				Value: "339208499",
			},
			"site01": {
				Value: "TIT创意园",
			},
			"site02": {
				Value: "广州市新港中路397号",
			},
		},
	}

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", SendSubscribeMsg(msg))

	assert.Nil(t, err)
}

func TestSendKFTextMessage(t *testing.T) {
	body := []byte(`{"touser":"OPENID","msgtype":"text","text":{"content":"Hello World"}}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	mp := New("APPID", "APPSECRET")
	mp.SetClient(wx.WithHTTPClient(client))

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", SendKFTextMsg("OPENID", "Hello World"))

	assert.Nil(t, err)
}

func TestSendKFImageMessage(t *testing.T) {
	body := []byte(`{"touser":"OPENID","msgtype":"image","image":{"media_id":"MEDIA_ID"}}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	mp := New("APPID", "APPSECRET")
	mp.SetClient(wx.WithHTTPClient(client))

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", SendKFImageMsg("OPENID", "MEDIA_ID"))

	assert.Nil(t, err)
}

func TestSendKFLinkMessage(t *testing.T) {
	body := []byte(`{"touser":"OPENID","msgtype":"link","link":{"title":"Happy Day","description":"Is Really A Happy Day","url":"URL","thumb_url":"THUMB_URL"}}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	mp := New("APPID", "APPSECRET")
	mp.SetClient(wx.WithHTTPClient(client))

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", SendKFLinkMsg("OPENID", &KFLink{
		Title:       "Happy Day",
		Description: "Is Really A Happy Day",
		URL:         "URL",
		ThumbURL:    "THUMB_URL",
	}))

	assert.Nil(t, err)
}

func TestSendKFMinipMessage(t *testing.T) {
	body := []byte(`{"touser":"OPENID","msgtype":"miniprogrampage","miniprogrampage":{"title":"title","pagepath":"pagepath","thumb_media_id":"thumb_media_id"}}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	mp := New("APPID", "APPSECRET")
	mp.SetClient(wx.WithHTTPClient(client))

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", SendKFMinipMsg("OPENID", &KFMinipPage{
		Title:        "title",
		PagePath:     "pagepath",
		ThumbMediaID: "thumb_media_id",
	}))

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

	mp := New("APPID", "APPSECRET")
	mp.SetClient(wx.WithHTTPClient(client))

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", SendKFTyping("OPENID", Typing))

	assert.Nil(t, err)
}
