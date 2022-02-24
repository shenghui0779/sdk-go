package minip

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

func TestGetPubTemplateKeyWords(t *testing.T) {
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

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetPubTemplateKeyWords("99", result))

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

func TestListPubTemplateTitle(t *testing.T) {
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
            "categoryId": 616
        }
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatetitles?access_token=ACCESS_TOKEN&ids=2%2C616&limit=1&start=0", nil).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(ResultPubTemplTitles)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetPubTemplateTitles("2,616", 0, 1, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultPubTemplTitles{
		Count: 55,
		Data: []*PubTemplTitle{
			{
				TID:        99,
				Title:      "付款成功通知",
				Type:       2,
				CategoryID: 616,
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
            "content": "会议时间:{{date2.DATA}}\\n会议地点:{{thing1.DATA}}\\n",
            "example": "会议时间:2016年8月8日\\n会议地点:TIT会议室\\n",
            "type": 2
        },
        {
            "priTmplId": "cy_DfOZL7lypxHh3ja3DyAUbn1GYQRGwezuy5LBTFME",
            "title": "洗衣机故障提醒",
            "content": "完成时间:{{time1.DATA}}\\n所在位置:{{enum_string2.DATA}}\\n提示说明:{{enum_string3.DATA}}\\n",
            "example": "完成时间:2021年10月21日 12:00:00\\n所在位置:客厅\\n提示说明:设备发生故障，导致工作异常，请及时查看\\n",
			"type": 3,
            "keywordEnumValueList": [
                {
                    "enumValueList": [
                        "客厅",
                        "餐厅",
                        "厨房",
                        "卧室",
                        "主卧",
                        "次卧",
                        "客卧",
                        "父母房",
                        "儿童房",
                        "男孩房",
                        "女孩房",
                        "卫生间",
                        "主卧卫生间",
                        "公共卫生间",
                        "衣帽间",
                        "书房",
                        "游戏室",
                        "阳台",
                        "地下室",
                        "储物间",
                        "车库",
                        "保姆房",
                        "其他房间"
                    ],
                    "keywordCode": "enum_string2.DATA"
                },
                {
                    "enumValueList": [
                        "设备发生故障，导致工作异常，请及时查看"
                    ],
                    "keywordCode": "enum_string3.DATA"
                }
            ]
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
				PriTmplId: "9Aw5ZV1j9xdWTFEkqCpZ7mIBbSC34khK55OtzUPl0rU",
				Title:     "报名结果通知",
				Content:   "会议时间:{{date2.DATA}}\\n会议地点:{{thing1.DATA}}\\n",
				Example:   "会议时间:2016年8月8日\\n会议地点:TIT会议室\\n",
				Type:      2,
			},
			{
				PriTmplId: "cy_DfOZL7lypxHh3ja3DyAUbn1GYQRGwezuy5LBTFME",
				Title:     "洗衣机故障提醒",
				Content:   "完成时间:{{time1.DATA}}\\n所在位置:{{enum_string2.DATA}}\\n提示说明:{{enum_string3.DATA}}\\n",
				Example:   "完成时间:2021年10月21日 12:00:00\\n所在位置:客厅\\n提示说明:设备发生故障，导致工作异常，请及时查看\\n",
				Type:      3,
				KeywordEnumValueList: []*KeywordEnumValue{
					{
						EnumValueList: []string{"客厅", "餐厅", "厨房", "卧室", "主卧", "次卧", "客卧", "父母房", "儿童房", "男孩房", "女孩房", "卫生间", "主卧卫生间", "公共卫生间", "衣帽间", "书房", "游戏室", "阳台", "地下室", "储物间", "车库", "保姆房", "其他房间"},
						KeywordCode:   "enum_string2.DATA",
					},
					{
						EnumValueList: []string{"设备发生故障，导致工作异常，请及时查看"},
						KeywordCode:   "enum_string3.DATA",
					},
				},
			},
		},
	}, result)
}
