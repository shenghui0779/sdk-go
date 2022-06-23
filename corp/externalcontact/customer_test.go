package externalcontact

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/chenghonour/gochat/corp"
	"github.com/chenghonour/gochat/mock"
	"github.com/chenghonour/gochat/wx"
)

func TestList(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"external_userid": [
		"woAJ2GCAAAXtWyujaWJHDDGi0mACAAA",
		"wmqfasd1e1927831291723123109rAAA"
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/list?access_token=ACCESS_TOKEN&userid=USERID", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", List("USERID", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultList{
		ExternalUserID: []string{
			"woAJ2GCAAAXtWyujaWJHDDGi0mACAAA",
			"wmqfasd1e1927831291723123109rAAA",
		},
	}, result)
}

func TestGet(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"external_contact": {
		"external_userid": "woAJ2GCAAAXtWyujaWJHDDGi0mACHAAA",
		"name": "李四",
		"position": "Manager",
		"avatar": "http://p.qlogo.cn/bizmail/IcsdgagqefergqerhewSdage/0",
		"corp_name": "腾讯",
		"corp_full_name": "腾讯科技有限公司",
		"type": 2,
		"gender": 1,
		"unionid": "ozynqsulJFCZ2z1aYeS8h-nuasdAAA",
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
						"appid": "wx8bd80126147df384",
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
			"description": "对接采购事务",
			"createtime": 1525779812,
			"tags": [
				{
					"group_name": "标签分组名称",
					"tag_name": "标签名称",
					"tag_id": "etAJ2GCAAAXtWyujaWJHDDGi0mACHAAA",
					"type": 1
				},
				{
					"group_name": "标签分组名称",
					"tag_name": "标签名称",
					"type": 2
				},
				{
					"group_name": "标签分组名称",
					"tag_name": "标签名称",
					"tag_id": "stAJ2GCAAAXtWyujaWJHDDGi0mACHAAA",
					"type": 3
				}
			],
			"remark_corp_name": "腾讯科技",
			"remark_mobiles": [
				"13800000001",
				"13000000002"
			],
			"oper_userid": "rocky",
			"add_way": 1
		},
		{
			"userid": "tommy",
			"remark": "李总",
			"description": "采购问题咨询",
			"createtime": 1525881637,
			"state": "外联二维码1",
			"oper_userid": "woAJ2GCAAAXtWyujaWJHDDGi0mACHAAA",
			"add_way": 3
		}
	],
	"next_cursor": "NEXT_CURSOR"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get?access_token=ACCESS_TOKEN&cursor=CURSOR&external_userid=EXTERNAL_USERID", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultGet)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", Get("EXTERNAL_USERID", "CURSOR", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultGet{
		ExternalContact: &ExternalContact{
			ExternalUserID: "woAJ2GCAAAXtWyujaWJHDDGi0mACHAAA",
			Name:           "李四",
			Position:       "Manager",
			Avatar:         "http://p.qlogo.cn/bizmail/IcsdgagqefergqerhewSdage/0",
			CorpName:       "腾讯",
			CorpFullName:   "腾讯科技有限公司",
			Type:           2,
			Gender:         1,
			UnionID:        "ozynqsulJFCZ2z1aYeS8h-nuasdAAA",
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
							AppID:    "wx8bd80126147df384",
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
				Description:    "对接采购事务",
				CreateTime:     1525779812,
				RemarkCorpName: "腾讯科技",
				RemarkMobiles:  []string{"13800000001", "13000000002"},
				OperUserID:     "rocky",
				AddWay:         1,
				Tags: []*FollowTag{
					{
						GroupName: "标签分组名称",
						TagName:   "标签名称",
						TagID:     "etAJ2GCAAAXtWyujaWJHDDGi0mACHAAA",
						Type:      1,
					},
					{
						GroupName: "标签分组名称",
						TagName:   "标签名称",
						Type:      2,
					},
					{
						GroupName: "标签分组名称",
						TagName:   "标签名称",
						TagID:     "stAJ2GCAAAXtWyujaWJHDDGi0mACHAAA",
						Type:      3,
					},
				},
			},
			{
				UserID:      "tommy",
				Remark:      "李总",
				Description: "采购问题咨询",
				CreateTime:  1525881637,
				OperUserID:  "woAJ2GCAAAXtWyujaWJHDDGi0mACHAAA",
				AddWay:      3,
				State:       "外联二维码1",
			},
		},
		NextCursor: "NEXT_CURSOR",
	}, result)
}

func TestBatchGetByUser(t *testing.T) {
	body := []byte(`{"userid_list":["zhangsan","lisi"],"limit":100}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"external_contact_list": [
		{
			"external_contact": {
				"external_userid": "woAJ2GCAAAXtWyujaWJHDDGi0mACHAAA",
				"name": "李四",
				"position": "Manager",
				"avatar": "http://p.qlogo.cn/bizmail/IcsdgagqefergqerhewSdage/0",
				"corp_name": "腾讯",
				"corp_full_name": "腾讯科技有限公司",
				"type": 2,
				"gender": 1,
				"unionid": "ozynqsulJFCZ2z1aYeS8h-nuasdAAA",
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
								"appid": "wx8bd80126147df384",
								"pagepath": "/index",
								"title": "my miniprogram"
							}
						}
					]
				}
			},
			"follow_info": {
				"userid": "rocky",
				"remark": "李部长",
				"description": "对接采购事务",
				"createtime": 1525779812,
				"tag_id": [
					"etAJ2GCAAAXtWyujaWJHDDGi0mACHAAA"
				],
				"remark_corp_name": "腾讯科技",
				"remark_mobiles": [
					"13800000001",
					"13000000002"
				],
				"oper_userid": "rocky",
				"add_way": 1
			}
		},
		{
			"external_contact": {
				"external_userid": "woAJ2GCAAAXtWyujaWJHDDGi0mACHBBB",
				"name": "王五",
				"position": "Engineer",
				"avatar": "http://p.qlogo.cn/bizmail/IcsdgagqefergqerhewSdage/0",
				"corp_name": "腾讯",
				"corp_full_name": "腾讯科技有限公司",
				"type": 2,
				"gender": 1,
				"unionid": "ozynqsulJFCZ2asdaf8h-nuasdAAA"
			},
			"follow_info": {
				"userid": "lisi",
				"remark": "王助理",
				"description": "采购问题咨询",
				"createtime": 1525881637,
				"tag_id": [
					"etAJ2GCAAAXtWyujaWJHDDGi0mACHAAA",
					"stJHDDGi0mAGi0mACHBBByujaW"
				],
				"state": "外联二维码1",
				"oper_userid": "woAJ2GCAAAd1asdasdjO4wKmE8AabjBBB",
				"add_way": 3
			}
		}
	],
	"next_cursor": "r9FqSqsI8fgNbHLHE5QoCP50UIg2cFQbfma3l2QsmwI"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/batch/get_by_user?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	userIDs := []string{"zhangsan", "lisi"}

	result := new(ResultBatchGetByUser)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", BatchGetByUser(userIDs, "", 100, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultBatchGetByUser{
		ExternalContactList: []*CustomerBatchGetData{
			{
				ExternalContact: &ExternalContact{
					ExternalUserID: "woAJ2GCAAAXtWyujaWJHDDGi0mACHAAA",
					Name:           "李四",
					Position:       "Manager",
					Avatar:         "http://p.qlogo.cn/bizmail/IcsdgagqefergqerhewSdage/0",
					CorpName:       "腾讯",
					CorpFullName:   "腾讯科技有限公司",
					Type:           2,
					Gender:         1,
					UnionID:        "ozynqsulJFCZ2z1aYeS8h-nuasdAAA",
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
									AppID:    "wx8bd80126147df384",
									PagePath: "/index",
								},
							},
						},
					},
				},
				FollowInfo: &FollowInfo{
					UserID:         "rocky",
					Remark:         "李部长",
					Description:    "对接采购事务",
					CreateTime:     1525779812,
					TagID:          []string{"etAJ2GCAAAXtWyujaWJHDDGi0mACHAAA"},
					RemarkCorpName: "腾讯科技",
					RemarkMobiles:  []string{"13800000001", "13000000002"},
					OperUserID:     "rocky",
					AddWay:         1,
				},
			},
			{
				ExternalContact: &ExternalContact{
					ExternalUserID: "woAJ2GCAAAXtWyujaWJHDDGi0mACHBBB",
					Name:           "王五",
					Position:       "Engineer",
					Avatar:         "http://p.qlogo.cn/bizmail/IcsdgagqefergqerhewSdage/0",
					CorpName:       "腾讯",
					CorpFullName:   "腾讯科技有限公司",
					Type:           2,
					Gender:         1,
					UnionID:        "ozynqsulJFCZ2asdaf8h-nuasdAAA",
				},
				FollowInfo: &FollowInfo{
					UserID:      "lisi",
					Remark:      "王助理",
					Description: "采购问题咨询",
					CreateTime:  1525881637,
					TagID: []string{
						"etAJ2GCAAAXtWyujaWJHDDGi0mACHAAA",
						"stJHDDGi0mAGi0mACHBBByujaW",
					},
					State:      "外联二维码1",
					OperUserID: "woAJ2GCAAAd1asdasdjO4wKmE8AabjBBB",
					AddWay:     3,
				},
			},
		},
		NextCursor: "r9FqSqsI8fgNbHLHE5QoCP50UIg2cFQbfma3l2QsmwI",
	}, result)
}

func TestRemark(t *testing.T) {
	body := []byte(`{"userid":"zhangsan","external_userid":"woAJ2GCAAAd1asdasdjO4wKmE8Aabj9AAA","remark":"备注信息","description":"描述信息","remark_company":"腾讯科技","remark_mobiles":["13800000001","13800000002"],"remark_pic_mediaid":"MEDIAID"}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/remark?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsRemark{
		UserID:           "zhangsan",
		ExternalUserID:   "woAJ2GCAAAd1asdasdjO4wKmE8Aabj9AAA",
		Remark:           "备注信息",
		Description:      "描述信息",
		RemarkCompany:    "腾讯科技",
		RemarkMobiles:    []string{"13800000001", "13800000002"},
		RemarkPicMediaID: "MEDIAID",
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", Remark(params))

	assert.Nil(t, err)
}
