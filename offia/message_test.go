package offia

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/wx"
)

func TestGetAllPrivateTemplate(t *testing.T) {
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

	result := new(ResultAllPrivateTemplate)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetAllPrivateTemplate(result))

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
	}, result)
}

func TestDelPrivateTemplate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/template/del_private_template?access_token=ACCESS_TOKEN", []byte(`{"template_id":"Dyvp3-Ff0cnail_CDSzk1fIc6-9lOkxsQE7exTJbwUE"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", DelPrivateTemplate("Dyvp3-Ff0cnail_CDSzk1fIc6-9lOkxsQE7exTJbwUE"))

	assert.Nil(t, err)
}

func TestSendTemplateMsg(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=ACCESS_TOKEN", []byte(`{"data":{"first":{"color":"#173177","value":"恭喜你购买成功！"},"keyword1":{"color":"#173177","value":"巧克力"},"remark":{"color":"#173177","value":"欢迎再次购买！"}},"miniprogram":{"appid":"xiaochengxuappid12345","pagepath":"index?foo=bar"},"template_id":"ngqIpbwh8bUfcSsECmogfXcV14J0tQlEpBO27izEYtY","touser":"OPENID","url":"http://weixin.qq.com/download"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	params := &ParamsTemplateMsg{
		ToUser:     "OPENID",
		TemplateID: "ngqIpbwh8bUfcSsECmogfXcV14J0tQlEpBO27izEYtY",
		URL:        "http://weixin.qq.com/download",
		Minip: &MsgMinip{
			AppID:    "xiaochengxuappid12345",
			Pagepath: "index?foo=bar",
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
			"remark": {
				Value: "欢迎再次购买！",
				Color: "#173177",
			},
		},
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendTemplateMsg(params))

	assert.Nil(t, err)
}

func TestSendSubscribeTemplateMsg(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/template/subscribe?access_token=ACCESS_TOKEN", []byte(`{"data":{"content":{"color":"COLOR","value":"VALUE"}},"miniprogram":{"appid":"xiaochengxuappid12345","pagepath":"index?foo=bar"},"scene":"SCENE","template_id":"TEMPLATE_ID","title":"TITLE","touser":"OPENID","url":"URL"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	params := &ParamsTemplateMsgSubscribe{
		ToUser:     "OPENID",
		Scene:      "SCENE",
		Title:      "TITLE",
		TemplateID: "TEMPLATE_ID",
		URL:        "URL",
		Minip: &MsgMinip{
			AppID:    "xiaochengxuappid12345",
			Pagepath: "index?foo=bar",
		},
		Data: MsgTemplData{
			"content": {
				Value: "VALUE",
				Color: "COLOR",
			},
		},
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendSubscribeTemplateMsg(params))

	assert.Nil(t, err)
}
