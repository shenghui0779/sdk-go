package offia

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

func TestSetIndustry(t *testing.T) {
	body := []byte(`{"industry_id1":"1","industry_id2":"4"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/template/api_set_industry?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SetIndustry("1", "4"))

	assert.Nil(t, err)
}

func TestGetIndustry(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "primary_industry": {
        "first_class": "运输与仓储",
        "second_class": "快递"
    },
    "secondary_industry": {
        "first_class": "IT科技",
        "second_class": "互联网|电子商务"
    }
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://api.weixin.qq.com/cgi-bin/template/get_industry?access_token=ACCESS_TOKEN", nil).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(ResultIndustryGet)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetIndustry(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultIndustryGet{
		PrimaryIndustry: &IndustryInfo{
			FirstClass:  "运输与仓储",
			SecondClass: "快递",
		},
		SecondaryIndustry: &IndustryInfo{
			FirstClass:  "IT科技",
			SecondClass: "互联网|电子商务",
		},
	}, result)
}

func TestAddTemplate(t *testing.T) {
	body := []byte(`{"template_id_short":"TM00015"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "template_id": "Doclyl5uP7Aciu-qZ7mJNPtWkbkYnWBWVja26EGbNyk"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/template/api_add_template?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(ResultTemplAdd)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", AddTemplate("TM00015", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultTemplAdd{
		TemplateID: "Doclyl5uP7Aciu-qZ7mJNPtWkbkYnWBWVja26EGbNyk",
	}, result)
}

func TestGetAllPrivateTemplate(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"template_list": [{
		"template_id": "iPk5sOIt5X_flOVKn5GrTFpncEYTojx6ddbt8WYoV5s",
		"title": "领取奖金提醒",
		"primary_industry": "IT科技",
		"deputy_industry": "互联网|电子商务",
		"content": "{ {result.DATA} }\n\n领奖金额:{ {withdrawMoney.DATA} }\n领奖  时间:    { {withdrawTime.DATA} }\n银行信息:{ {cardInfo.DATA} }\n到账时间:  { {arrivedTime.DATA} }\n{ {remark.DATA} }",
		"example": "您已提交领奖申请\n\n领奖金额：xxxx元\n领奖时间：2013-10-10 12:22:22\n银行信息：xx银行(尾号xxxx)\n到账时间：预计xxxxxxx\n\n预计将于xxxx到达您的银行卡"
	}]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://api.weixin.qq.com/cgi-bin/template/get_all_private_template?access_token=ACCESS_TOKEN", nil).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(ResultAllPrivateTemplate)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetAllPrivateTemplate(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultAllPrivateTemplate{
		TemplateList: []*TemplateInfo{
			{
				TemplateID:      "iPk5sOIt5X_flOVKn5GrTFpncEYTojx6ddbt8WYoV5s",
				Title:           "领取奖金提醒",
				PrimaryIndustry: "IT科技",
				DeputyIndustry:  "互联网|电子商务",
				Content:         "{ {result.DATA} }\n\n领奖金额:{ {withdrawMoney.DATA} }\n领奖  时间:    { {withdrawTime.DATA} }\n银行信息:{ {cardInfo.DATA} }\n到账时间:  { {arrivedTime.DATA} }\n{ {remark.DATA} }",
				Example:         "您已提交领奖申请\n\n领奖金额：xxxx元\n领奖时间：2013-10-10 12:22:22\n银行信息：xx银行(尾号xxxx)\n到账时间：预计xxxxxxx\n\n预计将于xxxx到达您的银行卡",
			},
		},
	}, result)
}

func TestDelPrivateTemplate(t *testing.T) {
	body := []byte(`{"template_id":"Dyvp3-Ff0cnail_CDSzk1fIc6-9lOkxsQE7exTJbwUE"}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/template/del_private_template?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", DelPrivateTemplate("Dyvp3-Ff0cnail_CDSzk1fIc6-9lOkxsQE7exTJbwUE"))

	assert.Nil(t, err)
}

func TestSendTemplateMsg(t *testing.T) {
	body := []byte(`{"touser":"OPENID","template_id":"ngqIpbwh8bUfcSsECmogfXcV14J0tQlEpBO27izEYtY","url":"http://weixin.qq.com/download","miniprogram":{"appid":"xiaochengxuappid12345","pagepath":"index?foo=bar"},"data":{"first":{"value":"恭喜你购买成功！","color":"#173177"},"keyword1":{"value":"巧克力","color":"#173177"},"remark":{"value":"欢迎再次购买！","color":"#173177"}}}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	params := &TemplateMsg{
		ToUser:     "OPENID",
		TemplateID: "ngqIpbwh8bUfcSsECmogfXcV14J0tQlEpBO27izEYtY",
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
			"remark": {
				Value: "欢迎再次购买！",
				Color: "#173177",
			},
		},
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendTemplateMsg(params))

	assert.Nil(t, err)
}

func TestSubscribeTemplate(t *testing.T) {
	body := []byte(`{"touser":"OPENID","scene":"SCENE","title":"TITLE","template_id":"TEMPLATE_ID","url":"URL","miniprogram":{"appid":"xiaochengxuappid12345","pagepath":"index?foo=bar"},"data":{"content":{"value":"VALUE","color":"COLOR"}}}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/message/template/subscribe?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	params := &ParamsTemplateSubscribe{
		ToUser:     "OPENID",
		Scene:      "SCENE",
		Title:      "TITLE",
		TemplateID: "TEMPLATE_ID",
		URL:        "URL",
		Minip: &MsgMinip{
			AppID:    "xiaochengxuappid12345",
			PagePath: "index?foo=bar",
		},
		Data: MsgTemplData{
			"content": {
				Value: "VALUE",
				Color: "COLOR",
			},
		},
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SubscribeTemplate(params))

	assert.Nil(t, err)
}
