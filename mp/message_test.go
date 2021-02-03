package mp

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/wx"
	"github.com/stretchr/testify/assert"
)

func TestSendUniformMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/uniform_send?access_token=ACCESS_TOKEN", []byte(`{"mp_template_msg":{"appid":"APPID","data":{"first":{"color":"#173177","value":"恭喜你购买成功！"},"keyword1":{"color":"#173177","value":"巧克力"},"keyword2":{"color":"#173177","value":"39.8元"},"keyword3":{"color":"#173177","value":"2014年9月22日"},"remark":{"color":"#173177","value":"欢迎再次购买！"}},"miniprogram":{"appid":"xiaochengxuappid12345","pagepath":"index?foo=bar"},"template_id":"TEMPLATE_ID","url":"http://weixin.qq.com/download"},"touser":"OPENID","weapp_template_msg":{"data":{"keyword1":{"value":"339208499"},"keyword2":{"value":"2015年01月05日 12:30"},"keyword3":{"value":"腾讯微信总部"},"keyword4":{"value":"广州市海珠区新港中路397号"}},"emphasis_keyword":"keyword1.DATA","form_id":"FORMID","page":"page/page/index","template_id":"TEMPLATE_ID"}}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	message := &UniformMessage{
		MPTemplateMessage: &TemplateMessage{
			TemplateID: "TEMPLATE_ID",
			Page:       "page/page/index",
			FormID:     "FORMID",
			Data: MessageBody{
				"keyword1": {
					"value": "339208499",
				},
				"keyword2": {
					"value": "2015年01月05日 12:30",
				},
				"keyword3": {
					"value": "腾讯微信总部",
				},
				"keyword4": {
					"value": "广州市海珠区新港中路397号",
				},
			},
			EmphasisKeyword: "keyword1.DATA",
		},
		OATemplateMessage: &OATemplateMessage{
			AppID:       "APPID",
			TemplateID:  "TEMPLATE_ID",
			RedirectURL: "http://weixin.qq.com/download",
			MiniProgram: &MessageMinip{
				AppID:    "xiaochengxuappid12345",
				Pagepath: "index?foo=bar",
			},
			Data: MessageBody{
				"first": {
					"value": "恭喜你购买成功！",
					"color": "#173177",
				},
				"keyword1": {
					"value": "巧克力",
					"color": "#173177",
				},
				"keyword2": {
					"value": "39.8元",
					"color": "#173177",
				},
				"keyword3": {
					"value": "2014年9月22日",
					"color": "#173177",
				},
				"remark": {
					"value": "欢迎再次购买！",
					"color": "#173177",
				},
			},
		},
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendUniformMessage("OPENID", message))

	assert.Nil(t, err)
}

func TestSendSubscribeMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=ACCESS_TOKEN", []byte(`{"data":{"date01":{"value":"2015年01月05日"},"number01":{"value":"339208499"},"site01":{"value":"TIT创意园"},"site02":{"value":"广州市新港中路397号"}},"lang":"zh_CN","miniprogram_state":"developer","page":"index","template_id":"TEMPLATE_ID","touser":"OPENID"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	message := &SubscribeMessage{
		TemplateID: "TEMPLATE_ID",
		Page:       "index",
		Data: MessageBody{
			"number01": {
				"value": "339208499",
			},
			"date01": {
				"value": "2015年01月05日",
			},
			"site01": {
				"value": "TIT创意园",
			},
			"site02": {
				"value": "广州市新港中路397号",
			},
		},
		MinipState: "developer",
		Lang:       "zh_CN",
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendSubscribeMessage("OPENID", message))

	assert.Nil(t, err)
}

func TestSendTemplateMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token=ACCESS_TOKEN", []byte(`{"data":{"keyword1":{"value":"339208499"},"keyword2":{"value":"2015年01月05日 12:30"},"keyword3":{"value":"腾讯微信总部"},"keyword4":{"value":"广州市海珠区新港中路397号"}},"emphasis_keyword":"keyword1.DATA","form_id":"FORMID","page":"index","template_id":"TEMPLATE_ID","touser":"OPENID"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	message := &TemplateMessage{
		TemplateID: "TEMPLATE_ID",
		Page:       "index",
		FormID:     "FORMID",
		Data: MessageBody{
			"keyword1": {
				"value": "339208499",
			},
			"keyword2": {
				"value": "2015年01月05日 12:30",
			},
			"keyword3": {
				"value": "腾讯微信总部",
			},
			"keyword4": {
				"value": "广州市海珠区新港中路397号",
			},
		},
		EmphasisKeyword: "keyword1.DATA",
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendTemplateMessage("OPENID", message))

	assert.Nil(t, err)
}

func TestSendKFTextMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", []byte(`{"msgtype":"text","text":{"content":"Hello World"},"touser":"OPENID"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	message := &KFTextMessage{
		Content: "Hello World",
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendKFTextMessage("OPENID", message))

	assert.Nil(t, err)
}

func TestSendKFImageMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", []byte(`{"image":{"media_id":"MEDIA_ID"},"msgtype":"image","touser":"OPENID"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	message := &KFImageMessage{
		MediaID: "MEDIA_ID",
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendKFImageMessage("OPENID", message))

	assert.Nil(t, err)
}

func TestSendKFLinkMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", []byte(`{"link":{"title":"Happy Day","description":"Is Really A Happy Day","url":"URL","thumb_url":"THUMB_URL"},"msgtype":"link","touser":"OPENID"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	message := &KFLinkMessage{
		Title:       "Happy Day",
		Description: "Is Really A Happy Day",
		RedirectURL: "URL",
		ThumbURL:    "THUMB_URL",
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendKFLinkMessage("OPENID", message))

	assert.Nil(t, err)
}

func TestSendKFMinipMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", []byte(`{"miniprogrampage":{"title":"title","pagepath":"pagepath","thumb_media_id":"thumb_media_id"},"msgtype":"miniprogrampage","touser":"OPENID"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	message := &KFMinipMessage{
		Title:        "title",
		Pagepath:     "pagepath",
		ThumbMediaID: "thumb_media_id",
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendKFMinipMessage("OPENID", message))

	assert.Nil(t, err)
}

func TestSetTyping(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/custom/typing?access_token=ACCESS_TOKEN", []byte(`{"command":"Typing","touser":"OPENID"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SetTyping("OPENID", Typing))

	assert.Nil(t, err)
}
