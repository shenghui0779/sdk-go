package offia

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/wx"
)

func TestGetTemplateList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/template/get_all_private_template?access_token=ACCESS_TOKEN").Return([]byte(`{
		"template_list": [{
			"template_id": "iPk5sOIt5X_flOVKn5GrTFpncEYTojx6ddbt8WYoV5s",
			"title": "领取奖金提醒",
			"primary_industry": "IT科技",
			"deputy_industry": "互联网|电子商务",
			"content": "{ {result.DATA} }\n\n领奖金额:{ {withdrawMoney.DATA} }\n领奖  时间:    { {withdrawTime.DATA} }\n银行信息:{ {cardInfo.DATA} }\n到账时间:  { {arrivedTime.DATA} }\n{ {remark.DATA} }",
			"example": "您已提交领奖申请\n\n领奖金额：xxxx元\n领奖时间：2013-10-10 12:22:22\n银行信息：xx银行(尾号xxxx)\n到账时间：预计xxxxxxx\n\n预计将于xxxx到达您的银行卡"
		}]
   	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := make([]*TemplateInfo, 0)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetTemplateList(&dest))

	assert.Nil(t, err)
	assert.Equal(t, []*TemplateInfo{
		{
			TemplateID:      "iPk5sOIt5X_flOVKn5GrTFpncEYTojx6ddbt8WYoV5s",
			Title:           "领取奖金提醒",
			PrimaryIndustry: "IT科技",
			DeputyIndustry:  "互联网|电子商务",
			Content:         "{ {result.DATA} }\n\n领奖金额:{ {withdrawMoney.DATA} }\n领奖  时间:    { {withdrawTime.DATA} }\n银行信息:{ {cardInfo.DATA} }\n到账时间:  { {arrivedTime.DATA} }\n{ {remark.DATA} }",
			Example:         "您已提交领奖申请\n\n领奖金额：xxxx元\n领奖时间：2013-10-10 12:22:22\n银行信息：xx银行(尾号xxxx)\n到账时间：预计xxxxxxx\n\n预计将于xxxx到达您的银行卡",
		},
	}, dest)
}

func TestDeleteTemplate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/template/del_private_template?access_token=ACCESS_TOKEN", []byte(`{"template_id":"Dyvp3-Ff0cnail_CDSzk1fIc6-9lOkxsQE7exTJbwUE"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", DeleteTemplate("Dyvp3-Ff0cnail_CDSzk1fIc6-9lOkxsQE7exTJbwUE"))

	assert.Nil(t, err)
}

func TestSendTemplateMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=ACCESS_TOKEN", []byte(`{"data":{"first":{"color":"#173177","value":"恭喜你购买成功！"},"keyword1":{"color":"#173177","value":"巧克力"},"remark":{"color":"#173177","value":"欢迎再次购买！"}},"miniprogram":{"appid":"xiaochengxuappid12345","pagepath":"index?foo=bar"},"template_id":"ngqIpbwh8bUfcSsECmogfXcV14J0tQlEpBO27izEYtY","touser":"OPENID","url":"http://weixin.qq.com/download"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	msg := &TemplateMessage{
		TemplateID: "ngqIpbwh8bUfcSsECmogfXcV14J0tQlEpBO27izEYtY",
		URL:        "http://weixin.qq.com/download",
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
			"remark": {
				"value": "欢迎再次购买！",
				"color": "#173177",
			},
		},
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendTemplateMessage("OPENID", msg))

	assert.Nil(t, err)
}

func TestSendSubscribeMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/template/subscribe?access_token=ACCESS_TOKEN", []byte(`{"data":{"content":{"color":"COLOR","value":"VALUE"}},"miniprogram":{"appid":"xiaochengxuappid12345","pagepath":"index?foo=bar"},"scene":"SCENE","template_id":"TEMPLATE_ID","title":"TITLE","touser":"OPENID","url":"URL"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	msg := &TemplateMessage{
		TemplateID: "TEMPLATE_ID",
		URL:        "URL",
		MiniProgram: &MessageMinip{
			AppID:    "xiaochengxuappid12345",
			Pagepath: "index?foo=bar",
		},
		Data: MessageBody{
			"content": {
				"value": "VALUE",
				"color": "COLOR",
			},
		},
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendSubscribeMessage("OPENID", "SCENE", "TITLE", msg))

	assert.Nil(t, err)
}

func TestSendKFTextMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", []byte(`{"customservice":{"kf_account":"test1@kftest"},"msgtype":"text","text":{"content":"Hello World"},"touser":"OPENID"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendKFTextMessage("OPENID", "Hello World", "test1@kftest"))

	assert.Nil(t, err)
}

func TestSendKFImageMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", []byte(`{"customservice":{"kf_account":"test1@kftest"},"image":{"media_id":"MEDIA_ID"},"msgtype":"image","touser":"OPENID"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendKFImageMessage("OPENID", "MEDIA_ID", "test1@kftest"))

	assert.Nil(t, err)
}

func TestSendKFVoiceMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", []byte(`{"customservice":{"kf_account":"test1@kftest"},"msgtype":"voice","touser":"OPENID","voice":{"media_id":"MEDIA_ID"}}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendKFVoiceMessage("OPENID", "MEDIA_ID", "test1@kftest"))

	assert.Nil(t, err)
}

func TestSendKFVideoMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", []byte(`{"customservice":{"kf_account":"test1@kftest"},"msgtype":"video","touser":"OPENID","video":{"media_id":"MEDIA_ID","thumb_media_id":"THUMB_MEDIA_ID","title":"TITLE","description":"DESCRIPTION"}}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	msg := &KFVideoMessage{
		MediaID:      "MEDIA_ID",
		ThumbMediaID: "THUMB_MEDIA_ID",
		Title:        "TITLE",
		Description:  "DESCRIPTION",
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendKFVideoMessage("OPENID", msg, "test1@kftest"))

	assert.Nil(t, err)
}

func TestSendKFMusicMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", []byte(`{"customservice":{"kf_account":"test1@kftest"},"msgtype":"music","music":{"title":"MUSIC_TITLE","description":"MUSIC_DESCRIPTION","musicurl":"MUSIC_URL","hqmusicurl":"HQ_MUSIC_URL","thumb_media_id":"THUMB_MEDIA_ID"},"touser":"OPENID"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	msg := &KFMusicMessage{
		Title:        "MUSIC_TITLE",
		Description:  "MUSIC_DESCRIPTION",
		MusicURL:     "MUSIC_URL",
		HQMusicURL:   "HQ_MUSIC_URL",
		ThumbMediaID: "THUMB_MEDIA_ID",
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendKFMusicMessage("OPENID", msg, "test1@kftest"))

	assert.Nil(t, err)
}

func TestSendKFNewsMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", []byte(`{"customservice":{"kf_account":"test1@kftest"},"msgtype":"news","news":{"articles":[{"title":"Happy Day","description":"Is Really A Happy Day","url":"URL","picurl":"PIC_URL"}]},"touser":"OPENID"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	articles := []*KFArticle{
		{
			Title:       "Happy Day",
			Description: "Is Really A Happy Day",
			URL:         "URL",
			PicURL:      "PIC_URL",
		},
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendKFNewsMessage("OPENID", articles, "test1@kftest"))

	assert.Nil(t, err)
}

func TestSendKFMPNewsMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", []byte(`{"customservice":{"kf_account":"test1@kftest"},"mpnews":{"media_id":"MEDIA_ID"},"msgtype":"mpnews","touser":"OPENID"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendKFMPNewsMessage("OPENID", "MEDIA_ID", "test1@kftest"))

	assert.Nil(t, err)
}

func TestSendKFMenuMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", []byte(`{"customservice":{"kf_account":"test1@kftest"},"msgmenu":{"head_content":"您对本次服务是否满意呢? ","tail_content":"欢迎再次光临","list":[{"id":"101","content":"满意"},{"id":"102","content":"不满意"}]},"msgtype":"msgmenu","touser":"OPENID"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	msg := &KFMenuMessage{
		HeadContent: "您对本次服务是否满意呢? ",
		TailContent: "欢迎再次光临",
		List: []*KFMenuOption{
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

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendKFMenuMessage("OPENID", msg, "test1@kftest"))

	assert.Nil(t, err)
}

func TestSendKFCardMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", []byte(`{"customservice":{"kf_account":"test1@kftest"},"msgtype":"wxcard","touser":"OPENID","wxcard":{"card_id":"123dsdajkasd231jhksad"}}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendKFCardMessage("OPENID", "123dsdajkasd231jhksad", "test1@kftest"))

	assert.Nil(t, err)
}

func TestSendKFMinipMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN", []byte(`{"miniprogrampage":{"title":"title","appid":"appid","pagepath":"pagepath","thumb_media_id":"thumb_media_id"},"msgtype":"miniprogrampage","touser":"OPENID"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	msg := &KFMinipMessage{
		Title:        "title",
		AppID:        "appid",
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

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/custom/typing?access_token=ACCESS_TOKEN", []byte(`{"command":"Typing","touser":"OPENID"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SetTyping("OPENID", Typing))

	assert.Nil(t, err)
}
