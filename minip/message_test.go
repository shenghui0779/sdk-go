package minip

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/wx"
)

func TestSendUniformMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/uniform_send?access_token=ACCESS_TOKEN", []byte(`{"touser":"OPENID","mp_template_msg":{"appid":"APPID","template_id":"TEMPLATE_ID","url":"http://weixin.qq.com/download","miniprogram":{"appid":"xiaochengxuappid12345","pagepath":"index?foo=bar"},"data":{"first":{"value":"恭喜你购买成功！","color":"#173177"},"keyword1":{"value":"巧克力","color":"#173177"},"keyword2":{"value":"39.8元","color":"#173177"},"keyword3":{"value":"2014年9月22日","color":"#173177"},"remark":{"value":"欢迎再次购买！","color":"#173177"}}}}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	params := &ParamsUniformMsg{
		ToUser: "OPENID",
		MPTemplateMsg: &MPTemplateMsg{
			AppID:      "APPID",
			TemplateID: "TEMPLATE_ID",
			URL:        "http://weixin.qq.com/download",
			Minip: &MsgMinip{
				AppID:    "xiaochengxuappid12345",
				Pagepath: "index?foo=bar",
			},
			Data: MsgTemplateData{
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
		},
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendUniformMessage(params))

	assert.Nil(t, err)
}

func TestSendSubscribeMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=ACCESS_TOKEN", []byte(`{"touser":"OPENID","template_id":"TEMPLATE_ID","page":"index","miniprogram_state":"developer","lang":"zh_CN","data":{"date01":{"value":"2015年01月05日"},"number01":{"value":"339208499"},"site01":{"value":"TIT创意园"},"site02":{"value":"广州市新港中路397号"}}}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	params := &ParamsSubscribeMsg{
		ToUser:     "OPENID",
		TemplateID: "TEMPLATE_ID",
		Page:       "index",
		MinipState: MinipDeveloper,
		Lang:       "zh_CN",
		Data: MsgTemplateData{
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

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendSubscribeMessage(params))

	assert.Nil(t, err)
}

func TestSendKFTextMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", []byte(`{"touser":"OPENID","msgtype":"text","text":{"content":"Hello World"}}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendKFTextMessage("OPENID", "Hello World"))

	assert.Nil(t, err)
}

func TestSendKFImageMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", []byte(`{"touser":"OPENID","msgtype":"image","image":{"media_id":"MEDIA_ID"}}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendKFImageMessage("OPENID", "MEDIA_ID"))

	assert.Nil(t, err)
}

func TestSendKFLinkMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", []byte(`{"touser":"OPENID","msgtype":"link","link":{"title":"Happy Day","description":"Is Really A Happy Day","url":"URL","thumb_url":"THUMB_URL"}}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	msg := &MsgLink{
		Title:       "Happy Day",
		Description: "Is Really A Happy Day",
		URL:         "URL",
		ThumbURL:    "THUMB_URL",
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendKFLinkMessage("OPENID", msg))

	assert.Nil(t, err)
}

func TestSendKFMinipMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", []byte(`{"touser":"OPENID","msgtype":"miniprogrampage","miniprogrampage":{"title":"title","pagepath":"pagepath","thumb_media_id":"thumb_media_id"}}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	msg := &MsgMinipPage{
		Title:        "title",
		Pagepath:     "pagepath",
		ThumbMediaID: "thumb_media_id",
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendKFMinipMessage("OPENID", msg))

	assert.Nil(t, err)
}

func TestSetTyping(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/custom/typing?access_token=ACCESS_TOKEN", []byte(`{"touser":"OPENID","command":"Typing"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendKFTyping("OPENID", Typing))

	assert.Nil(t, err)
}
