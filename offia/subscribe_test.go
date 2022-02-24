package offia

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/mock"
	"github.com/shenghui0779/gochat/wx"
)

func TestAddSubscribeTemplate(t *testing.T) {
	body := []byte(`{"tid":"401","kidList":[1,2],"sceneDesc":"测试数据"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errmsg": "ok",
	"errcode": 0,
	"priTmplId": "9Aw5ZV1j9xdWTFEkqCpZ7jWySL7aGN6rQom4gXINfJs"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/wxaapi/newtmpl/addtemplate?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	params := &ParamsSubscribeTemplAdd{
		TID:       "401",
		KidList:   []int{1, 2},
		SceneDesc: "测试数据",
	}
	result := new(ResultSubscribeTemplAdd)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", AddSubscribeTemplate(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultSubscribeTemplAdd{
		PriTmplID: "9Aw5ZV1j9xdWTFEkqCpZ7jWySL7aGN6rQom4gXINfJs",
	}, result)
}

func TestDeleteSubscribeTemplate(t *testing.T) {
	body := []byte(`{"priTmplId":"wDYzYZVxobJivW9oMpSCpuvACOfJXQIoKUm0PY397Tc"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/wxaapi/newtmpl/deltemplate?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", DeleteSubscribeTemplate("wDYzYZVxobJivW9oMpSCpuvACOfJXQIoKUm0PY397Tc"))

	assert.Nil(t, err)
}

func TestGetSubscribeCategory(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"data": [
		{
			"id": 616,
			"name": "公交"
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://api.weixin.qq.com/wxaapi/newtmpl/getcategory?access_token=ACCESS_TOKEN", nil).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(ResultSubscribeCategory)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetSubscribeCategory(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultSubscribeCategory{
		Data: []*SubscribeCategory{
			{
				ID:   616,
				Name: "公交",
			},
		},
	}, result)
}

func TestGetPubTemplateKeywords(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"data": [
		{
			"kid": 1,
			"name": "物品名称",
			"example": "名称",
			"rule": "thing"
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatekeywords?access_token=ACCESS_TOKEN&tid=99", nil).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(ResultPubTemplKeywords)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetPubTemplateKeywords("99", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultPubTemplKeywords{
		Data: []*PubTemplKeywords{
			{
				KID:     1,
				Name:    "物品名称",
				Example: "名称",
				Rule:    "thing",
			},
		},
	}, result)
}

func TestGetPubTemplateTitles(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"count": 55,
	"data": [
		{
			"tid": 99,
			"title": "付款成功通知",
			"type": 2,
			"categoryId": "616"
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatetitles?access_token=ACCESS_TOKEN&ids=616&limit=1&start=0", nil).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(ResultPubTemplTitles)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetPubTemplateTitles("616", 0, 1, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultPubTemplTitles{
		Count: 55,
		Data: []*PubTemplTitle{
			{
				TID:        99,
				Title:      "付款成功通知",
				Type:       2,
				CategoryID: "616",
			},
		},
	}, result)
}

func TestListSubscribeTemplate(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"data": [
		{
			"priTmplId": "9Aw5ZV1j9xdWTFEkqCpZ7mIBbSC34khK55OtzUPl0rU",
			"title": "报名结果通知",
			"content": "会议时间:{{date2.DATA}}\n会议地点:{{thing1.DATA}}\n",
			"example": "会议时间:2016年8月8日\n会议地点:TIT会议室\n",
			"type": 2
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://api.weixin.qq.com/wxaapi/newtmpl/gettemplate?access_token=ACCESS_TOKEN", nil).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(ResultSubscribeTemplList)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", ListSubscribeTemplate(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultSubscribeTemplList{
		Data: []*SubscribeTemplInfo{
			{
				PriTmplID: "9Aw5ZV1j9xdWTFEkqCpZ7mIBbSC34khK55OtzUPl0rU",
				Title:     "报名结果通知",
				Content:   "会议时间:{{date2.DATA}}\n会议地点:{{thing1.DATA}}\n",
				Example:   "会议时间:2016年8月8日\n会议地点:TIT会议室\n",
				Type:      2,
			},
		},
	}, result)
}

func TestSendSubscribePageMsg(t *testing.T) {
	body := []byte(`{"touser":"OPENID","template_id":"TEMPLATEID","page":"mp.weixin.qq.com","data":{"name1":{"value":"广州腾讯科技有限公司"},"thing8":{"value":"广州腾讯科技有限公司"},"time7":{"value":"2019年8月8日"}}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/message/subscribe/bizsend?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	data := MsgTemplData{
		"name1": &MsgTemplValue{
			Value: "广州腾讯科技有限公司",
		},
		"thing8": &MsgTemplValue{
			Value: "广州腾讯科技有限公司",
		},
		"time7": &MsgTemplValue{
			Value: "2019年8月8日",
		},
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendSubscribePageMsg("TEMPLATEID", "OPENID", "mp.weixin.qq.com", data))

	assert.Nil(t, err)
}

func TestSendSubscribeMinipMsg(t *testing.T) {
	body := []byte(`{"touser":"OPENID","template_id":"TEMPLATEID","miniprogram":{"appid":"APPID","pagepath":"index?foo=bar"},"data":{"name1":{"value":"广州腾讯科技有限公司"},"thing8":{"value":"广州腾讯科技有限公司"},"time7":{"value":"2019年8月8日"}}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/message/subscribe/bizsend?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	minip := &MsgMinip{
		AppID:    "APPID",
		PagePath: "index?foo=bar",
	}

	data := MsgTemplData{
		"name1": &MsgTemplValue{
			Value: "广州腾讯科技有限公司",
		},
		"thing8": &MsgTemplValue{
			Value: "广州腾讯科技有限公司",
		},
		"time7": &MsgTemplValue{
			Value: "2019年8月8日",
		},
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SendSubscribeMinipMsg("TEMPLATEID", "OPENID", minip, data))

	assert.Nil(t, err)
}
