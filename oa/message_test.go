package oa

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/wx"
	"github.com/stretchr/testify/assert"
)

func TestGetTemplateList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

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

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/template/del_private_template?access_token=ACCESS_TOKEN", []byte(`{"template_id":"Dyvp3-Ff0cnail_CDSzk1fIc6-9lOkxsQE7exTJbwUE"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", DeleteTemplate("Dyvp3-Ff0cnail_CDSzk1fIc6-9lOkxsQE7exTJbwUE"))

	assert.Nil(t, err)
}

func TestSendTemplateMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=ACCESS_TOKEN", []byte(`{"data":{"first":{"color":"#173177","value":"恭喜你购买成功！"},"keyword1":{"color":"#173177","value":"巧克力"},"remark":{"color":"#173177","value":"欢迎再次购买！"}},"miniprogram":{"appid":"xiaochengxuappid12345","pagepath":"index?foo=bar"},"template_id":"ngqIpbwh8bUfcSsECmogfXcV14J0tQlEpBO27izEYtY","touser":"OPENID","url":"http://weixin.qq.com/download"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	message := &TemplateMessage{
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

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendTemplateMessage("OPENID", message))

	assert.Nil(t, err)
}

func TestSendSubscribeMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/template/subscribe?access_token=ACCESS_TOKEN", []byte(`{"data":{"content":{"color":"COLOR","value":"VALUE"}},"miniprogram":{"appid":"xiaochengxuappid12345","pagepath":"index?foo=bar"},"scene":"SCENE","template_id":"TEMPLATE_ID","title":"TITLE","touser":"OPENID","url":"URL"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	message := &TemplateMessage{
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

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendSubscribeMessage("OPENID", "SCENE", "TITLE", message))

	assert.Nil(t, err)
}
