package user

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/corp"
	"github.com/shenghui0779/gochat/mock"
	"github.com/shenghui0779/gochat/wx"
)

func TestCreateUser(t *testing.T) {
	body := []byte(`{"userid":"zhangsan","name":"张三","alias":"jackzhang","mobile":"+86 13800000000","department":[1,2],"order":[10,40],"position":"产品经理","gender":"1","email":"zhangsan@gzdev.com","biz_mail":"zhangsan@qyycs2.wecom.work","is_leader_in_dept":[1,0],"direct_leader":["lisi","wangwu"],"enable":1,"avatar_mediaid":"2-G6nrLmr5EC3MNb_-zL1dDdzkd0p7cNliYu9V5w7o8K0","telephone":"020-123456","address":"广州市海珠区新港中路","main_department":1,"extattr":{"attrs":[{"type":0,"name":"文本名称","text":{"value":"文本"}},{"type":1,"name":"网页名称","web":{"url":"http://www.test.com","title":"标题"}}]},"to_invite":true,"external_position":"高级产品经理","external_profile":{"external_corp_name":"企业简称","wechat_channels":{"nickname":"视频号名称"},"external_attr":[{"type":0,"name":"文本名称","text":{"value":"文本"}},{"type":1,"name":"网页名称","web":{"url":"http://www.test.com","title":"标题"}},{"type":2,"name":"测试app","miniprogram":{"appid":"wx8bd8012614784fake","pagepath":"/index","title":"my miniprogram"}}]}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/user/create?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	toInvite := new(bool)
	*toInvite = true

	params := &ParamsUserCreate{
		UserID:         "zhangsan",
		Name:           "张三",
		Alias:          "jackzhang",
		Mobile:         "+86 13800000000",
		Department:     []int64{1, 2},
		Order:          []int{10, 40},
		Position:       "产品经理",
		Gender:         "1",
		Email:          "zhangsan@gzdev.com",
		BizMail:        "zhangsan@qyycs2.wecom.work",
		IsLeaderInDept: []int{1, 0},
		DirectLeader:   []string{"lisi", "wangwu"},
		Enable:         1,
		AvatarMediaID:  "2-G6nrLmr5EC3MNb_-zL1dDdzkd0p7cNliYu9V5w7o8K0",
		Telephone:      "020-123456",
		Address:        "广州市海珠区新港中路",
		MainDepartment: 1,
		ExtAttr: &ExtAttr{
			Attrs: []*Attr{
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
						URL:   "http://www.test.com",
						Title: "标题",
					},
				},
			},
		},
		ToInvite:         toInvite,
		ExternalPosition: "高级产品经理",
		ExternalProfile: &ExternalProfile{
			ExternalCorpName: "企业简称",
			WechatChannels: &WechatChannels{
				Nickname: "视频号名称",
			},
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
						URL:   "http://www.test.com",
						Title: "标题",
					},
				},
				{
					Type: 2,
					Name: "测试app",
					Miniprogram: &AttrMinip{
						AppID:    "wx8bd8012614784fake",
						Pagepath: "/index",
						Title:    "my miniprogram",
					},
				},
			},
		},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", CreateUser(params))

	assert.Nil(t, err)
}

func TestUpdateUser(t *testing.T) {
	body := []byte(`{"userid":"zhangsan","name":"李四","alias":"jackzhang","mobile":"13800000000","department":[1],"order":[10],"position":"后台工程师","gender":"1","email":"zhangsan@gzdev.com","biz_mail":"zhangsan@qyycs2.wecom.work","is_leader_in_dept":[1],"direct_leader":["lisi","wangwu"],"enable":1,"avatar_mediaid":"2-G6nrLmr5EC3MNb_-zL1dDdzkd0p7cNliYu9V5w7o8K0","telephone":"020-123456","address":"广州市海珠区新港中路","main_department":1,"extattr":{"attrs":[{"type":0,"name":"文本名称","text":{"value":"文本"}},{"type":1,"name":"网页名称","web":{"url":"http://www.test.com","title":"标题"}}]},"external_position":"工程师","external_profile":{"external_corp_name":"企业简称","wechat_channels":{"nickname":"视频号名称"},"external_attr":[{"type":0,"name":"文本名称","text":{"value":"文本"}},{"type":1,"name":"网页名称","web":{"url":"http://www.test.com","title":"标题"}},{"type":2,"name":"测试app","miniprogram":{"appid":"wx8bd80126147dFAKE","pagepath":"/index","title":"my miniprogram"}}]}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/user/update?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsUserUpdate{
		UserID:         "zhangsan",
		Name:           "李四",
		Alias:          "jackzhang",
		Mobile:         "13800000000",
		Department:     []int64{1},
		Order:          []int{10},
		Position:       "后台工程师",
		Gender:         "1",
		Email:          "zhangsan@gzdev.com",
		BizMail:        "zhangsan@qyycs2.wecom.work",
		IsLeaderInDept: []int{1},
		DirectLeader:   []string{"lisi", "wangwu"},
		Enable:         1,
		AvatarMediaID:  "2-G6nrLmr5EC3MNb_-zL1dDdzkd0p7cNliYu9V5w7o8K0",
		Telephone:      "020-123456",
		Address:        "广州市海珠区新港中路",
		MainDepartment: 1,
		ExtAttr: &ExtAttr{
			Attrs: []*Attr{
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
						URL:   "http://www.test.com",
						Title: "标题",
					},
				},
			},
		},
		ExternalPosition: "工程师",
		ExternalProfile: &ExternalProfile{
			ExternalCorpName: "企业简称",
			WechatChannels: &WechatChannels{
				Nickname: "视频号名称",
			},
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
						URL:   "http://www.test.com",
						Title: "标题",
					},
				},
				{
					Type: 2,
					Name: "测试app",
					Miniprogram: &AttrMinip{
						AppID:    "wx8bd80126147dFAKE",
						Pagepath: "/index",
						Title:    "my miniprogram",
					},
				},
			},
		},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UpdateUser(params))

	assert.Nil(t, err)
}

func TestGetUser(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "userid": "zhangsan",
    "name": "张三",
    "department": [
        1,
        2
    ],
    "order": [
        1,
        2
    ],
    "position": "后台工程师",
    "mobile": "13800000000",
    "gender": "1",
    "email": "zhangsan@gzdev.com",
    "biz_mail": "zhangsan@qyycs2.wecom.work",
    "is_leader_in_dept": [
        1,
        0
    ],
    "direct_leader": [
        "lisi",
        "wangwu"
    ],
    "avatar": "http://wx.qlogo.cn/mmopen/ajNVdqHZLLA3WJ6DSZUfiakYe37PKnQhBIeOQBO4czqrnZDS79FH5Wm5m4X69TBicnHFlhiafvDwklOpZeXYQQ2icg/0",
    "thumb_avatar": "http://wx.qlogo.cn/mmopen/ajNVdqHZLLA3WJ6DSZUfiakYe37PKnQhBIeOQBO4czqrnZDS79FH5Wm5m4X69TBicnHFlhiafvDwklOpZeXYQQ2icg/100",
    "telephone": "020-123456",
    "alias": "jackzhang",
    "address": "广州市海珠区新港中路",
    "open_userid": "xxxxxx",
    "main_department": 1,
    "extattr": {
        "attrs": [
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
            }
        ]
    },
    "status": 1,
    "qr_code": "https://open.work.weixin.qq.com/wwopen/userQRCode?vcode=xxx",
    "external_position": "产品经理",
    "external_profile": {
        "external_corp_name": "企业简称",
        "wechat_channels": {
            "nickname": "视频号名称",
            "status": 1
        },
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
                    "appid": "wx8bd80126147dFAKE",
                    "pagepath": "/index",
                    "title": "my miniprogram"
                }
            }
        ]
    }
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token=ACCESS_TOKEN&userid=USERID", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(User)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetUser("USERID", result))

	assert.Nil(t, err)
	assert.Equal(t, &User{
		UserID:         "zhangsan",
		Name:           "张三",
		Alias:          "jackzhang",
		Mobile:         "13800000000",
		Department:     []int64{1, 2},
		Order:          []int{1, 2},
		Position:       "后台工程师",
		Gender:         "1",
		Email:          "zhangsan@gzdev.com",
		BizMail:        "zhangsan@qyycs2.wecom.work",
		IsLeaderInDept: []int{1, 0},
		DirectLeader:   []string{"lisi", "wangwu"},
		Avatar:         "http://wx.qlogo.cn/mmopen/ajNVdqHZLLA3WJ6DSZUfiakYe37PKnQhBIeOQBO4czqrnZDS79FH5Wm5m4X69TBicnHFlhiafvDwklOpZeXYQQ2icg/0",
		ThumbAvatar:    "http://wx.qlogo.cn/mmopen/ajNVdqHZLLA3WJ6DSZUfiakYe37PKnQhBIeOQBO4czqrnZDS79FH5Wm5m4X69TBicnHFlhiafvDwklOpZeXYQQ2icg/100",
		Telephone:      "020-123456",
		Address:        "广州市海珠区新港中路",
		OpenUserID:     "xxxxxx",
		MainDepartment: 1,
		ExtAttr: &ExtAttr{
			Attrs: []*Attr{
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
						URL:   "http://www.test.com",
						Title: "标题",
					},
				},
			},
		},
		Status:           1,
		QRCode:           "https://open.work.weixin.qq.com/wwopen/userQRCode?vcode=xxx",
		ExternalPosition: "产品经理",
		ExternalProfile: &ExternalProfile{
			ExternalCorpName: "企业简称",
			WechatChannels: &WechatChannels{
				Nickname: "视频号名称",
				Status:   1,
			},
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
						URL:   "http://www.test.com",
						Title: "标题",
					},
				},
				{
					Type: 2,
					Name: "测试app",
					Miniprogram: &AttrMinip{
						AppID:    "wx8bd80126147dFAKE",
						Pagepath: "/index",
						Title:    "my miniprogram",
					},
				},
			},
		},
	}, result)
}

func TestDeleteUser(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/user/delete?access_token=ACCESS_TOKEN&userid=USERID", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", DeleteUser("USERID"))

	assert.Nil(t, err)
}

func TestBatchDeleteUser(t *testing.T) {
	body := []byte(`{"useridlist":["zhangsan","lisi"]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/user/batchdelete?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", BatchDeleteUser("zhangsan", "lisi"))

	assert.Nil(t, err)
}

func TestListSimpleUser(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "userlist": [
        {
            "userid": "zhangsan",
            "name": "张三",
            "department": [
                1,
                2
            ],
            "open_userid": "xxxxxx"
        }
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/user/simplelist?access_token=ACCESS_TOKEN&department_id=1&fetch_child=0", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultSimipleUserList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListSimpleUser(1, 0, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultSimipleUserList{
		UserList: []*SimpleUser{
			{
				UserID:     "zhangsan",
				Name:       "张三",
				Department: []int64{1, 2},
				OpenUserID: "xxxxxx",
			},
		},
	}, result)
}

func TestListUser(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "userlist": [
        {
            "userid": "zhangsan",
            "name": "李四",
            "department": [
                1,
                2
            ],
            "order": [
                1,
                2
            ],
            "position": "后台工程师",
            "mobile": "13800000000",
            "gender": "1",
            "email": "zhangsan@gzdev.com",
            "biz_mail": "zhangsan@qyycs2.wecom.work",
            "is_leader_in_dept": [
                1,
                0
            ],
            "direct_leader": [
                "lisi",
                "wangwu"
            ],
            "avatar": "http://wx.qlogo.cn/mmopen/ajNVdqHZLLA3WJ6DSZUfiakYe37PKnQhBIeOQBO4czqrnZDS79FH5Wm5m4X69TBicnHFlhiafvDwklOpZeXYQQ2icg/0",
            "thumb_avatar": "http://wx.qlogo.cn/mmopen/ajNVdqHZLLA3WJ6DSZUfiakYe37PKnQhBIeOQBO4czqrnZDS79FH5Wm5m4X69TBicnHFlhiafvDwklOpZeXYQQ2icg/100",
            "telephone": "020-123456",
            "alias": "jackzhang",
            "status": 1,
            "address": "广州市海珠区新港中路",
            "hide_mobile": 0,
            "english_name": "jacky",
            "open_userid": "xxxxxx",
            "main_department": 1,
            "extattr": {
                "attrs": [
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
                    }
                ]
            },
            "qr_code": "https://open.work.weixin.qq.com/wwopen/userQRCode?vcode=xxx",
            "external_position": "产品经理",
            "external_profile": {
                "external_corp_name": "企业简称",
                "wechat_channels": {
                    "nickname": "视频号名称",
                    "status": 1
                },
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
                            "appid": "wx8bd80126147dFAKE",
                            "pagepath": "/index",
                            "title": "miniprogram"
                        }
                    }
                ]
            }
        }
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/user/list?access_token=ACCESS_TOKEN&department_id=1&fetch_child=0", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultUserList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListUser(1, 0, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultUserList{
		UserList: []*User{
			{
				UserID:         "zhangsan",
				Name:           "李四",
				Alias:          "jackzhang",
				Mobile:         "13800000000",
				Department:     []int64{1, 2},
				Order:          []int{1, 2},
				Position:       "后台工程师",
				Gender:         "1",
				Email:          "zhangsan@gzdev.com",
				BizMail:        "zhangsan@qyycs2.wecom.work",
				IsLeaderInDept: []int{1, 0},
				DirectLeader:   []string{"lisi", "wangwu"},
				Avatar:         "http://wx.qlogo.cn/mmopen/ajNVdqHZLLA3WJ6DSZUfiakYe37PKnQhBIeOQBO4czqrnZDS79FH5Wm5m4X69TBicnHFlhiafvDwklOpZeXYQQ2icg/0",
				ThumbAvatar:    "http://wx.qlogo.cn/mmopen/ajNVdqHZLLA3WJ6DSZUfiakYe37PKnQhBIeOQBO4czqrnZDS79FH5Wm5m4X69TBicnHFlhiafvDwklOpZeXYQQ2icg/100",
				Telephone:      "020-123456",
				Address:        "广州市海珠区新港中路",
				OpenUserID:     "xxxxxx",
				MainDepartment: 1,
				ExtAttr: &ExtAttr{
					Attrs: []*Attr{
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
								URL:   "http://www.test.com",
								Title: "标题",
							},
						},
					},
				},
				Status:           1,
				QRCode:           "https://open.work.weixin.qq.com/wwopen/userQRCode?vcode=xxx",
				ExternalPosition: "产品经理",
				ExternalProfile: &ExternalProfile{
					ExternalCorpName: "企业简称",
					WechatChannels: &WechatChannels{
						Nickname: "视频号名称",
						Status:   1,
					},
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
								URL:   "http://www.test.com",
								Title: "标题",
							},
						},
						{
							Type: 2,
							Name: "测试app",
							Miniprogram: &AttrMinip{
								AppID:    "wx8bd80126147dFAKE",
								Pagepath: "/index",
								Title:    "miniprogram",
							},
						},
					},
				},
			},
		},
	}, result)
}

func TestConvertToOpenID(t *testing.T) {
	body := []byte(`{"userid":"zhangsan"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "openid": "oDjGHs-1yCnGrRovBj2yHij5JAAA"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/user/convert_to_openid?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultOpenIDConvert)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ConvertToOpenID("zhangsan", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultOpenIDConvert{
		OpenID: "oDjGHs-1yCnGrRovBj2yHij5JAAA",
	}, result)
}

func TestConvertToUserID(t *testing.T) {
	body := []byte(`{"openid":"oDjGHs-1yCnGrRovBj2yHij5JAAA"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "userid": "zhangsan"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/user/convert_to_userid?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultUserIDConvert)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ConvertToUserID("oDjGHs-1yCnGrRovBj2yHij5JAAA", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultUserIDConvert{
		UserID: "zhangsan",
	}, result)
}

func TestBatchInvite(t *testing.T) {
	body := []byte(`{"user":["UserID1","UserID2","UserID3"],"party":[1,2],"tag":[1,2]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "invaliduser": [
        "UserID1",
        "UserID2"
    ],
    "invalidparty": [
        1,
        2
    ],
    "invalidtag": [
        1,
        2
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/batch/invite?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	userIDs := []string{"UserID1", "UserID2", "UserID3"}
	partyIDs := []int64{1, 2}
	tagIDs := []int64{1, 2}

	result := new(ResultBatchInvite)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", BatchInvite(userIDs, partyIDs, tagIDs, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultBatchInvite{
		InvalidUser:  []string{"UserID1", "UserID2"},
		InvalidParty: []int64{1, 2},
		InvalidTag:   []int64{1, 2},
	}, result)
}

func TestGetJoinQRCode(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "join_qrcode": "https://work.weixin.qq.com/wework_admin/genqrcode?action=join&vcode=3db1fab03118ae2aa1544cb9abe84&r=hb_share_api_mjoin&qr_size=3"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/corp/get_join_qrcode?access_token=ACCESS_TOKEN&size_type=1", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(JoinQRCode)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetJoinQRCode(1, result))

	assert.Nil(t, err)
	assert.Equal(t, &JoinQRCode{
		URL: "https://work.weixin.qq.com/wework_admin/genqrcode?action=join&vcode=3db1fab03118ae2aa1544cb9abe84&r=hb_share_api_mjoin&qr_size=3",
	}, result)
}

func TestGetActiveStat(t *testing.T) {
	body := []byte(`{"date":"2020-03-27"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "active_cnt": 100
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/user/get_active_stat?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultActiveStat)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetActiveStat("2020-03-27", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultActiveStat{
		ActiveCnt: 100,
	}, result)
}

func TestGetUserID(t *testing.T) {
	body := []byte(`{"mobile":"13430388888"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "userid": "zhangsan"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/user/getuserid?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultUserID)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetUserID("13430388888", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultUserID{
		UserID: "zhangsan",
	}, result)
}
