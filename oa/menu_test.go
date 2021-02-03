package oa

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/wx"
	"github.com/stretchr/testify/assert"
)

func TestCreateMenu(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN", []byte(`{"button":[{"type":"click","name":"今日歌曲","key":"V1001_TODAY_MUSIC"},{"name":"菜单","sub_button":[{"type":"view","name":"搜索","url":"http://www.soso.com/"},{"type":"miniprogram","name":"wxa","url":"http://mp.weixin.qq.com","appid":"wx286b93c14bbf93aa","pagepath":"pages/lunar/index"},{"type":"click","name":"赞一下我们","key":"V1001_GOOD"}]}]}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	action := CreateMenu(
		ClickButton("今日歌曲", "V1001_TODAY_MUSIC"),
		GroupButton("菜单", ViewButton("搜索", "http://www.soso.com/"), MinipButton("wxa", "wx286b93c14bbf93aa", "pages/lunar/index", "http://mp.weixin.qq.com"), ClickButton("赞一下我们", "V1001_GOOD")),
	)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", action)

	assert.Nil(t, err)
}

func TestCreateConditionalMenu(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/menu/addconditional?access_token=ACCESS_TOKEN", []byte(`{"button":[{"type":"click","name":"今日歌曲","key":"V1001_TODAY_MUSIC"},{"name":"菜单","sub_button":[{"type":"view","name":"搜索","url":"http://www.soso.com/"},{"type":"miniprogram","name":"wxa","url":"http://mp.weixin.qq.com","appid":"wx286b93c14bbf93aa","pagepath":"pages/lunar/index"},{"type":"click","name":"赞一下我们","key":"V1001_GOOD"}]}],"matchrule":{"tag_id":"2","sex":"1","country":"中国","province":"广东","city":"广州","client_platform_type":"2","language":"zh_CN"}}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	matchRule := &MenuMatchRule{
		TagID:              "2",
		Sex:                "1",
		Country:            "中国",
		Province:           "广东",
		City:               "广州",
		ClientPlatformType: "2",
		Language:           "zh_CN",
	}

	action := CreateConditionalMenu(matchRule,
		ClickButton("今日歌曲", "V1001_TODAY_MUSIC"),
		GroupButton("菜单", ViewButton("搜索", "http://www.soso.com/"), MinipButton("wxa", "wx286b93c14bbf93aa", "pages/lunar/index", "http://mp.weixin.qq.com"), ClickButton("赞一下我们", "V1001_GOOD")),
	)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", action)

	assert.Nil(t, err)
}

func TestTryMatchMenu(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/menu/trymatch?access_token=ACCESS_TOKEN", []byte(`{"user_id":"weixin"}`)).Return([]byte(`{
		"button": [
			{
				"type": "view",
				"name": "tx",
				"url": "http://www.qq.com/",
				"sub_button": []
			}
		]
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := make([]*MenuButton, 0)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", TryMatchMenu(&dest, "weixin"))

	assert.Nil(t, err)
	assert.Equal(t, []*MenuButton{
		{
			Type:      ButtonView,
			Name:      "tx",
			URL:       "http://www.qq.com/",
			SubButton: []*MenuButton{},
		},
	}, dest)
}

func TestGetMenu(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/menu/get?access_token=ACCESS_TOKEN").Return([]byte(`{
		"menu": {
			"button": [
				{
					"type": "click",
					"name": "今日歌曲",
					"key": "V1001_TODAY_MUSIC",
					"sub_button": []
				}
			],
			"menuid": 208396938
		},
		"conditionalmenu": [
			{
				"button": [
					{
						"type": "click",
						"name": "今日歌曲",
						"key": "V1001_TODAY_MUSIC",
						"sub_button": []
					},
					{
						"name": "菜单",
						"sub_button": [
							{
								"type": "view",
								"name": "搜索",
								"url": "http://www.soso.com/",
								"sub_button": []
							},
							{
								"type": "view",
								"name": "视频",
								"url": "http://v.qq.com/",
								"sub_button": []
							},
							{
								"type": "click",
								"name": "赞一下我们",
								"key": "V1001_GOOD",
								"sub_button": []
							}
						]
					}
				],
				"matchrule": {
					"tag_id": "2",
					"sex": "1",
					"country": "中国",
					"province": "广东",
					"city": "广州",
					"client_platform_type": "2"
				},
				"menuid": 208396993
			}
		]
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(MenuInfo)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetMenu(dest))

	assert.Nil(t, err)
	assert.Equal(t, &MenuInfo{
		Menu: Menu{
			Button: []*MenuButton{
				{
					Type:      ButtonClick,
					Name:      "今日歌曲",
					Key:       "V1001_TODAY_MUSIC",
					URL:       "",
					AppID:     "",
					Pagepath:  "",
					MediaID:   "",
					SubButton: []*MenuButton{},
				},
			},
			MenuID: 208396938,
		},
		ConditionalMenu: []*ConditionalMenu{
			{
				Button: []*MenuButton{
					{
						Type:      ButtonClick,
						Name:      "今日歌曲",
						Key:       "V1001_TODAY_MUSIC",
						SubButton: []*MenuButton{},
					},
					{
						Name: "菜单",
						SubButton: []*MenuButton{
							{
								Type:      ButtonView,
								Name:      "搜索",
								URL:       "http://www.soso.com/",
								SubButton: []*MenuButton{},
							},
							{
								Type:      ButtonView,
								Name:      "视频",
								URL:       "http://v.qq.com/",
								SubButton: []*MenuButton{},
							},
							{
								Type:      ButtonClick,
								Name:      "赞一下我们",
								Key:       "V1001_GOOD",
								SubButton: []*MenuButton{},
							},
						},
					},
				},
				MatchRule: MenuMatchRule{
					TagID:              "2",
					Sex:                "1",
					Country:            "中国",
					Province:           "广东",
					City:               "广州",
					ClientPlatformType: "2",
					Language:           "",
				},
				MenuID: 208396993,
			},
		},
	}, dest)
}

func TestDeleteMenu(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=ACCESS_TOKEN").Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", DeleteMenu())

	assert.Nil(t, err)
}

func TestDeleteConditionalMenu(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/menu/delconditional?access_token=ACCESS_TOKEN", []byte(`{"menuid":"208379533"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", DeleteConditionalMenu("208379533"))

	assert.Nil(t, err)
}
