package school

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/corp"
	"github.com/shenghui0779/gochat/mock"
)

func TestGetSubscribeQRCode(t *testing.T) {
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"qrcode_big": "http://p.qpic.cn/wwhead/XXXX",
	"qrcode_middle": "http://p.qpic.cn/wwhead/XXXX",
	"qrcode_thumb": "http://p.qpic.cn/wwhead/XXXX"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_subscribe_qr_code?access_token=ACCESS_TOKEN", nil).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultSubscribeQRCode)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetSubscribeQRCode(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultSubscribeQRCode{
		QRCodeBig:    "http://p.qpic.cn/wwhead/XXXX",
		QRCodeMiddle: "http://p.qpic.cn/wwhead/XXXX",
		QRCodeThumb:  "http://p.qpic.cn/wwhead/XXXX",
	}, result)
}

func TestSetSubscribeMode(t *testing.T) {
	body := []byte(`{"subscribe_mode":1}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/set_subscribe_mode?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SetSubscribeMode(1))

	assert.Nil(t, err)
}

func TestGetSubscribeMode(t *testing.T) {
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"subscribe_mode": 1
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_subscribe_mode?access_token=ACCESS_TOKEN", nil).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultSubscribeModeGet)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetSubscribeMode(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultSubscribeModeGet{
		SubscribeMode: 1,
	}, result)
}

func TestGetExternalContact(t *testing.T) {
	resp := []byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "external_contact": {
        "external_userid": "woAAAA",
        "name": "李四",
        "foreign_key": "lisi",
        "position": "Mangaer",
        "avatar": "http://p.qlogo.cn/bizmail/IcsdgagqefergqerhewSdage/0",
        "corp_name": "腾讯",
        "corp_full_name": "腾讯科技有限公司",
        "type": 2,
        "gender": 1,
        "unionid": "unAAAAA",
        "is_subscribe": 1,
        "subscriber_info": {
            "tag_id": [
                "TAG_ID1",
                "TAG_ID2"
            ],
            "remark_mobiles": [
                "10000000000",
                "10000000001"
            ],
            "remark": "李小明-爸爸"
        },
        "external_profile": {
            "external_attr": [
                {
                    "type": 0,
                    "name": "文本名称",
                    "text": {
                        "value": "文本"
                    }
                },
                {
                    "type": 1,
                    "name": "网页名称",
                    "web": {
                        "url": "http://www.test.com",
                        "title": "标题"
                    }
                },
                {
                    "type": 2,
                    "name": "测试app",
                    "miniprogram": {
                        "appid": "wxAAAAA",
                        "pagepath": "/index",
                        "title": "my miniprogram"
                    }
                }
            ]
        }
    },
    "follow_user": [
        {
            "userid": "rocky",
            "remark": "李部长",
            "description": "对接采购事物",
            "createtime": 1525779812,
            "tags": [
                {
                    "group_name": "标签分组名称",
                    "tag_name": "标签名称",
                    "type": 1
                }
            ],
            "remark_corp_name": "腾讯科技",
            "remark_mobiles": [
                "10000000003",
                "10000000004"
            ]
        },
        {
            "userid": "tommy",
            "remark": "李总",
            "description": "采购问题咨询",
            "createtime": 1525881637,
            "state": "外联二维码1"
        }
    ]
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get?access_token=ACCESS_TOKEN&external_userid=EXTERNAL_USERID", nil).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultExternalContact)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetExternalContact("EXTERNAL_USERID", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultExternalContact{
		ExternalContact: &ExternalContact{
			ExternalUserID: "woAAAA",
			Name:           "李四",
			ForeignKey:     "lisi",
			Position:       "Mangaer",
			Avatar:         "http://p.qlogo.cn/bizmail/IcsdgagqefergqerhewSdage/0",
			CorpName:       "腾讯",
			CorpFullName:   "腾讯科技有限公司",
			Type:           2,
			Gender:         1,
			UnionID:        "unAAAAA",
			IsSubscribe:    1,
			SubscriberInfo: &SubscriberInfo{
				TagID:         []string{"TAG_ID1", "TAG_ID2"},
				RemarkMobiles: []string{"10000000000", "10000000001"},
				Remark:        "李小明-爸爸",
			},
			ExternalProfile: &ExternalProfile{
				ExternalAttr: []*Attr{
					{
						Type: 0,
						Name: "文本名称",
						Text: &AttrText{
							Value: "文本",
						},
					},
					{
						Type: 1,
						Name: "网页名称",
						Web: &AttrWeb{
							Title: "标题",
							URL:   "http://www.test.com",
						},
					},
					{
						Type: 2,
						Name: "测试app",
						Miniprogram: &AttrMinip{
							Title:    "my miniprogram",
							AppID:    "wxAAAAA",
							PagePath: "/index",
						},
					},
				},
			},
		},
		FollowUser: []*FollowUser{
			{
				UserID:         "rocky",
				Remark:         "李部长",
				Description:    "对接采购事物",
				CreateTime:     1525779812,
				RemarkCorpName: "腾讯科技",
				RemarkMobiles:  []string{"10000000003", "10000000004"},
				Tags: []*FollowTag{
					{
						GroupName: "标签分组名称",
						TagName:   "标签名称",
						Type:      1,
					},
				},
			},
			{
				UserID:      "tommy",
				Remark:      "李总",
				Description: "采购问题咨询",
				CreateTime:  1525881637,
				State:       "外联二维码1",
			},
		},
	}, result)
}

func TestConvertToOpenID(t *testing.T) {
	body := []byte(`{"external_userid":"wmAAAAAAA"}`)
	resp := []byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "openid": "ooAAAAAAAAAAA"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/convert_to_openid?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultOpenIDConvert)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ConvertToOpenID("wmAAAAAAA", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultOpenIDConvert{
		OpenID: "ooAAAAAAAAAAA",
	}, result)
}

func TestGetAgentAllowScope(t *testing.T) {
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"allow_scope": {
		"students": [
			{
				"userid": "student1"
			},
			{
				"userid": "student2"
			}
		],
		"departments": [
			1,
			2
		]
	}
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/school/agent/get_allow_scope?access_token=ACCESS_TOKEN&agentid=1", nil).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultAgentAllowScope)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetAgentAllowScope(1, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultAgentAllowScope{
		AllowScope: &AgentAllowScope{
			Students: []*AgentAllowStudent{
				{
					UserID: "student1",
				},
				{
					UserID: "student2",
				},
			},
			Departments: []int64{1, 2},
		},
	}, result)
}
