package offia

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/mock"
)

func TestCreateMenu(t *testing.T) {
	body := []byte(`{"button":[{"type":"click","name":"今日歌曲","key":"V1001_TODAY_MUSIC"},{"name":"菜单","sub_button":[{"type":"view","name":"搜索","url":"http://www.soso.com/"},{"type":"miniprogram","name":"wxa","url":"http://mp.weixin.qq.com","appid":"wx286b93c14bbf93aa","pagepath":"pages/lunar/index"},{"type":"click","name":"赞一下我们","key":"V1001_GOOD"}]}]}`)

	resp := []byte(`{"errcode":0,"errmsg":"ok"}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")

	buttons := []*MenuButton{
		ClickButton("今日歌曲", "V1001_TODAY_MUSIC"),
		GroupButton("菜单",
			ViewButton("搜索", "http://www.soso.com/"),
			MinipButton("wxa", "wx286b93c14bbf93aa", "pages/lunar/index", "http://mp.weixin.qq.com"),
			ClickButton("赞一下我们", "V1001_GOOD"),
		),
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", CreateMenu(buttons...))

	assert.Nil(t, err)
}

func TestCreateConditionalMenu(t *testing.T) {
	body := []byte(`{"button":[{"type":"click","name":"今日歌曲","key":"V1001_TODAY_MUSIC"},{"name":"菜单","sub_button":[{"type":"view","name":"搜索","url":"http://www.soso.com/"},{"type":"miniprogram","name":"wxa","url":"http://mp.weixin.qq.com","appid":"wx286b93c14bbf93aa","pagepath":"pages/lunar/index"},{"type":"click","name":"赞一下我们","key":"V1001_GOOD"}]}],"matchrule":{"tag_id":"2","client_platform_type":"2"}}`)

	resp := []byte(`{"errcode":0,"errmsg":"ok"}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/menu/addconditional?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")

	rule := &MenuMatchRule{
		TagID:              "2",
		ClientPlatformType: "2",
	}
	buttons := []*MenuButton{
		ClickButton("今日歌曲", "V1001_TODAY_MUSIC"),
		GroupButton("菜单",
			ViewButton("搜索", "http://www.soso.com/"),
			MinipButton("wxa", "wx286b93c14bbf93aa", "pages/lunar/index", "http://mp.weixin.qq.com"),
			ClickButton("赞一下我们", "V1001_GOOD"),
		),
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", CreateConditionalMenu(rule, buttons...))

	assert.Nil(t, err)
}

func TestTryMatchMenu(t *testing.T) {
	body := []byte(`{"user_id":"weixin"}`)

	resp := []byte(`{
	"button": [
		{
			"type": "view",
			"name": "tx",
			"url": "http://www.qq.com/",
			"sub_button": []
		}
	]
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/menu/trymatch?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")

	result := new(ResultMenuMatch)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", TryMatchMenu("weixin", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMenuMatch{
		Button: []*MenuButton{
			{
				Type:      ButtonView,
				Name:      "tx",
				URL:       "http://www.qq.com/",
				SubButton: []*MenuButton{},
			},
		},
	}, result)
}

func TestGetMenu(t *testing.T) {
	resp := []byte(`{
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
				"client_platform_type": "2"
			},
			"menuid": 208396993
		}
	]
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://api.weixin.qq.com/cgi-bin/menu/get?access_token=ACCESS_TOKEN", nil).Return(resp, nil)

	oa := New("APPID", "APPSECRET")

	result := new(ResultMenuGet)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetMenu(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMenuGet{
		Menu: Menu{
			Button: []*MenuButton{
				{
					Type:      ButtonClick,
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
					ClientPlatformType: "2",
				},
				MenuID: 208396993,
			},
		},
	}, result)
}

func TestDeleteMenu(t *testing.T) {
	resp := []byte(`{"errcode":0,"errmsg":"ok"}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=ACCESS_TOKEN", nil).Return(resp, nil)

	oa := New("APPID", "APPSECRET")

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", DeleteMenu())

	assert.Nil(t, err)
}

func TestDeleteConditionalMenu(t *testing.T) {
	body := []byte(`{"menuid":"208379533"}`)

	resp := []byte(`{"errcode":0,"errmsg":"ok"}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/menu/delconditional?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", DeleteConditionalMenu("208379533"))

	assert.Nil(t, err)
}
