package oa

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateMenu(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := NewMockHTTPClient(ctrl)

	body := []byte(`{"button":[{"key":"V1001_TODAY_MUSIC","name":"今日歌曲","type":"click"},{"name":"菜单","sub_button":[{"name":"搜索","type":"view","url":"http://www.soso.com/"},{"appid":"wx286b93c14bbf93aa","name":"wxa","pagepath":"pages/lunar/index","type":"miniprogram","url":"http://mp.weixin.qq.com"},{"key":"V1001_GOOD","name":"赞一下我们","type":"click"}]}]}`)
	client.EXPECT().Post("https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN", body, gomock.AssignableToTypeOf(WithHTTPHeader("Content-Type", "application/json; charset=utf-8"))).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("wxa06e66cf23dc4370", "1208c7f9e08b4edd26fd86406a5b30aa")
	oa.client = client

	action := CreateMenu(
		ClickButton("今日歌曲", "V1001_TODAY_MUSIC"),
		GroupButton("菜单", ViewButton("搜索", "http://www.soso.com/"), MPButton("wxa", "wx286b93c14bbf93aa", "pages/lunar/index", "http://mp.weixin.qq.com"), ClickButton("赞一下我们", "V1001_GOOD")),
	)

	err := oa.Do("ACCESS_TOKEN", action)

	assert.Nil(t, err)
}

func TestCreateConditionalMenu(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := NewMockHTTPClient(ctrl)

	body := []byte(`{"button":[{"key":"V1001_TODAY_MUSIC","name":"今日歌曲","type":"click"},{"name":"菜单","sub_button":[{"name":"搜索","type":"view","url":"http://www.soso.com/"},{"appid":"wx286b93c14bbf93aa","name":"wxa","pagepath":"pages/lunar/index","type":"miniprogram","url":"http://mp.weixin.qq.com"},{"key":"V1001_GOOD","name":"赞一下我们","type":"click"}]}],"matchrule":{"tag_id":"2","sex":"1","country":"中国","province":"广东","city":"广州","client_platform_type":"2","language":"zh_CN"}}`)
	client.EXPECT().Post("https://api.weixin.qq.com/cgi-bin/menu/addconditional?access_token=ACCESS_TOKEN", body, gomock.AssignableToTypeOf(WithHTTPHeader("Content-Type", "application/json; charset=utf-8"))).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("wxa06e66cf23dc4370", "1208c7f9e08b4edd26fd86406a5b30aa")
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
		GroupButton("菜单", ViewButton("搜索", "http://www.soso.com/"), MPButton("wxa", "wx286b93c14bbf93aa", "pages/lunar/index", "http://mp.weixin.qq.com"), ClickButton("赞一下我们", "V1001_GOOD")),
	)

	err := oa.Do("ACCESS_TOKEN", action)

	assert.Nil(t, err)
}

func TestGetMenu(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := NewMockHTTPClient(ctrl)

	client.EXPECT().Get("https://api.weixin.qq.com/cgi-bin/menu/get?access_token=ACCESS_TOKEN").Return([]byte(`{
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

	oa := New("wxa06e66cf23dc4370", "1208c7f9e08b4edd26fd86406a5b30aa")
	oa.client = client

	receiver := new(MenuInfo)

	err := oa.Do("ACCESS_TOKEN", GetMenu(receiver))

	assert.Nil(t, err)
	assert.Equal(t, &MenuInfo{
		DefaultMenu: &DefaultMenu{
			Button: []*MenuButton{
				{
					Type:      "click",
					Name:      "今日歌曲",
					Key:       "V1001_TODAY_MUSIC",
					URL:       "",
					AppID:     "",
					PagePath:  "",
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
						Type:      "click",
						Name:      "今日歌曲",
						Key:       "V1001_TODAY_MUSIC",
						URL:       "",
						AppID:     "",
						PagePath:  "",
						MediaID:   "",
						SubButton: []*MenuButton{},
					},
					{
						Type:     "",
						Name:     "菜单",
						Key:      "",
						URL:      "",
						AppID:    "",
						PagePath: "",
						MediaID:  "",
						SubButton: []*MenuButton{
							{
								Type:      "view",
								Name:      "搜索",
								Key:       "",
								URL:       "http://www.soso.com/",
								AppID:     "",
								PagePath:  "",
								MediaID:   "",
								SubButton: []*MenuButton{},
							},
							{
								Type:      "view",
								Name:      "视频",
								Key:       "",
								URL:       "http://v.qq.com/",
								AppID:     "",
								PagePath:  "",
								MediaID:   "",
								SubButton: []*MenuButton{},
							},
							{
								Type:      "click",
								Name:      "赞一下我们",
								Key:       "V1001_GOOD",
								URL:       "",
								AppID:     "",
								PagePath:  "",
								MediaID:   "",
								SubButton: []*MenuButton{},
							},
						},
					},
				},
				MatchRule: &MenuMatchRule{
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
	}, receiver)
}

func TestDeleteMenu(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := NewMockHTTPClient(ctrl)

	client.EXPECT().Get("https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=ACCESS_TOKEN").Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("wxa06e66cf23dc4370", "1208c7f9e08b4edd26fd86406a5b30aa")
	oa.client = client

	err := oa.Do("ACCESS_TOKEN", DeleteMenu())

	assert.Nil(t, err)
}

func TestDeleteConditionalMenu(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := NewMockHTTPClient(ctrl)

	body := []byte(`{"menuid":"208379533"}`)
	client.EXPECT().Post("https://api.weixin.qq.com/cgi-bin/menu/delconditional?access_token=ACCESS_TOKEN", body, gomock.AssignableToTypeOf(WithHTTPHeader("Content-Type", "application/json; charset=utf-8"))).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("wxa06e66cf23dc4370", "1208c7f9e08b4edd26fd86406a5b30aa")
	oa.client = client

	err := oa.Do("ACCESS_TOKEN", DeleteConditionalMenu("208379533"))

	assert.Nil(t, err)
}
