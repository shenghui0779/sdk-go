package oa

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

func TestGetTemplateDetail(t *testing.T) {
	body := []byte(`{"template_id":"ZLqk8pcsAoXZ1eYa6vpAgfX28MPdYU3ayMaSPHaaa"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"template_names": [
		{
			"text": "全字段",
			"lang": "zh_CN"
		}
	],
	"template_content": {
		"controls": [
			{
				"property": {
					"control": "Selector",
					"id": "Selector-15111111111",
					"title": [
						{
							"text": "单选控件",
							"lang": "zh_CN"
						}
					],
					"placeholder": [
						{
							"text": "这是单选控件的说明",
							"lang": "zh_CN"
						}
					],
					"require": 0,
					"un_print": 0
				},
				"config": {
					"selector": {
						"type": "single",
						"options": [
							{
								"key": "option-15111111111",
								"value": [
									{
										"text": "选项1",
										"lang": "zh_CN"
									}
								]
							},
							{
								"key": "option-15222222222",
								"value": [
									{
										"text": "选项2",
										"lang": "zh_CN"
									}
								]
							}
						]
					}
				}
			}
		]
	}
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/oa/gettemplatedetail?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultTemplateDetail)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetTemplateDetail("ZLqk8pcsAoXZ1eYa6vpAgfX28MPdYU3ayMaSPHaaa", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultTemplateDetail{
		TemplateNames: []*DisplayText{
			{
				Text: "全字段",
				Lang: "zh_CN",
			},
		},
		TemplateContent: &TemplateContent{
			Controls: []*TemplateControl{
				{
					Property: &ControlProperty{
						Control: ControlSelector,
						ID:      "Selector-15111111111",
						Title: []*DisplayText{
							{
								Text: "单选控件",
								Lang: "zh_CN",
							},
						},
						Placeholder: []*DisplayText{
							{
								Text: "这是单选控件的说明",
								Lang: "zh_CN",
							},
						},
						Require: 0,
						UnPrint: 0,
					},
					Config: &ControlConfig{
						Selector: &SelectorConfig{
							Type: "single",
							Options: []*SelectorOption{
								{
									Key: "option-15111111111",
									Value: []*DisplayText{
										{
											Text: "选项1",
											Lang: "zh_CN",
										},
									},
								},
								{
									Key: "option-15222222222",
									Value: []*DisplayText{
										{
											Text: "选项2",
											Lang: "zh_CN",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}, result)
}

func TestApplyEvent(t *testing.T) {
	body := []byte(`{"creator_userid":"WangXiaoMing","template_id":"3Tka1eD6v6JfzhDMqPd3aMkFdxqtJMc2ZRioeFXkaaa","use_template_approver":0,"choose_department":2,"approver":[{"attr":2,"userid":["WuJunJie","WangXiaoMing"]},{"attr":1,"userid":["LiuXiaoGang"]}],"notifyer":["WuJunJie","WangXiaoMing"],"notify_type":1,"apply_data":{"contents":[{"control":"Text","id":"Text-15111111111","value":{"text":"文本填写的内容"}}]},"summary_list":[{"summary_info":[{"text":"摘要第1行","lang":"zh_CN"}]},{"summary_info":[{"text":"摘要第2行","lang":"zh_CN"}]},{"summary_info":[{"text":"摘要第3行","lang":"zh_CN"}]}]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"sp_no": "202001010001"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/oa/applyevent?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsApplyEvent{
		CreatorUserID:       "WangXiaoMing",
		TemplateID:          "3Tka1eD6v6JfzhDMqPd3aMkFdxqtJMc2ZRioeFXkaaa",
		UseTemplateApprover: 0,
		ChooseDepartment:    2,
		Approver: []*Approver{
			{
				Attr:   2,
				UserID: []string{"WuJunJie", "WangXiaoMing"},
			},
			{
				Attr:   1,
				UserID: []string{"LiuXiaoGang"},
			},
		},
		Notifyer:   []string{"WuJunJie", "WangXiaoMing"},
		NotifyType: 1,
		ApplyData: &ApplyData{
			Contents: []*ApplyContent{
				{
					Control: ControlText,
					ID:      "Text-15111111111",
					Value: &ControlValue{
						Text: "文本填写的内容",
					},
				},
			},
		},
		SummaryList: []*ApplySummaryInfo{
			{
				SummaryInfo: []*DisplayText{
					{
						Text: "摘要第1行",
						Lang: "zh_CN",
					},
				},
			},
			{
				SummaryInfo: []*DisplayText{
					{
						Text: "摘要第2行",
						Lang: "zh_CN",
					},
				},
			},
			{
				SummaryInfo: []*DisplayText{
					{
						Text: "摘要第3行",
						Lang: "zh_CN",
					},
				},
			},
		},
	}

	result := new(ResultApplyEvent)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ApplyEvent(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultApplyEvent{
		SPNO: "202001010001",
	}, result)
}

func TestGetApprovalInfo(t *testing.T) {
	body := []byte(`{"starttime":"1569546000","endtime":"1569718800","cursor":0,"size":100,"filters":[{"key":"template_id","value":"ZLqk8pcsAoaXZ1eY56vpAgfX28MPdYU3ayMaSPHaaa"},{"key":"creator","value":"WuJunJie"},{"key":"department","value":"1"},{"key":"sp_status","value":"1"}]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"sp_no_list": [
		"201909270001",
		"201909270002",
		"201909270003"
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/oa/getapprovalinfo?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsApprovalInfo{
		StartTime: "1569546000",
		EndTime:   "1569718800",
		Cursor:    0,
		Size:      100,
		Filters: []*KeyValue{
			{
				Key:   "template_id",
				Value: "ZLqk8pcsAoaXZ1eY56vpAgfX28MPdYU3ayMaSPHaaa",
			},
			{
				Key:   "creator",
				Value: "WuJunJie",
			},
			{
				Key:   "department",
				Value: "1",
			},
			{
				Key:   "sp_status",
				Value: "1",
			},
		},
	}

	result := new(ResultApprovalInfo)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetApprovalInfo(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultApprovalInfo{
		SPNOList: []string{
			"201909270001",
			"201909270002",
			"201909270003",
		},
	}, result)
}

func TestGetApprovalDetail(t *testing.T) {
	body := []byte(`{"sp_no":"201909270001"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"info": {
		"sp_no": "201909270002",
		"sp_name": "全字段",
		"sp_status": 1,
		"template_id": "Bs5KJ2NT4ncf4ZygaE8MB3779yUW8nsMaJd3mmE9v",
		"apply_time": 1569584428,
		"applyer": {
			"userid": "WuJunJie",
			"partyid": "2"
		},
		"sp_record": [
			{
				"sp_status": 1,
				"approverattr": 1,
				"details": [
					{
						"approver": {
							"userid": "WuJunJie"
						},
						"speech": "",
						"sp_status": 1,
						"sptime": 0,
						"media_id": []
					},
					{
						"approver": {
							"userid": "WangXiaoMing"
						},
						"speech": "",
						"sp_status": 1,
						"sptime": 0,
						"media_id": []
					}
				]
			}
		],
		"notifyer": [
			{
				"userid": "LiuXiaoGang"
			}
		],
		"apply_data": {
			"contents": [
				{
					"control": "Text",
					"id": "Text-15111111111",
					"title": [
						{
							"text": "文本控件",
							"lang": "zh_CN"
						}
					],
					"value": {
						"text": "文本填写的内容",
						"tips": [],
						"members": [],
						"departments": [],
						"files": [],
						"children": [],
						"stat_field": []
					}
				}
			]
		},
		"comments": [
			{
				"commentUserInfo": {
					"userid": "WuJunJie"
				},
				"commenttime": 1569584111,
				"commentcontent": "这是备注信息",
				"commentid": "6741314136717778040",
				"media_id": [
					"WWCISP_Xa1dXIyC9VC2vGTXyBjUXh4GQ31G-a7jilEjFjkYBfncSJv0kM1cZAIXULWbbtosVqA7hprZIUkl4GP0DYZKDrIay9vCzeQelmmHiczwfn80v51EtuNouzBhUBTWo9oQIIzsSftjaVmd4EC_dj5-rayfDl6yIIRdoUs1V_Gz6Pi3yH37ELOgLNAPYUSJpA6V190Xunl7b0s5K5XC9c7eX5vlJek38rB_a2K-kMFMiM1mHDqnltoPa_NT9QynXuHi"
				]
			}
		]
	}
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/oa/getapprovaldetail?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultApprovalDetail)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetApprovalDetail("201909270001", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultApprovalDetail{
		Info: &ApprovalDetail{
			SPNO:       "201909270002",
			SPName:     "全字段",
			SPStatus:   1,
			TemplateID: "Bs5KJ2NT4ncf4ZygaE8MB3779yUW8nsMaJd3mmE9v",
			ApplyTime:  1569584428,
			Applyer: &Applyer{
				UserID:  "WuJunJie",
				PartyID: "2",
			},
			SPRecord: []*ApprovalSPRecord{
				{
					SPStatus:     1,
					ApproverAttr: 1,
					Details: []*ApprovalSPDetail{
						{
							Approver: &OAUser{
								UserID: "WuJunJie",
							},
							Speech:   "",
							SPStatus: 1,
							SPTime:   0,
							MediaID:  []string{},
						},
						{
							Approver: &OAUser{
								UserID: "WangXiaoMing",
							},
							Speech:   "",
							SPStatus: 1,
							SPTime:   0,
							MediaID:  []string{},
						},
					},
				},
			},
			Notifyer: []*OAUser{
				{
					UserID: "LiuXiaoGang",
				},
			},
			ApplyData: &ApplyData{
				Contents: []*ApplyContent{
					{
						Control: ControlText,
						ID:      "Text-15111111111",
						Title: []*DisplayText{
							{
								Text: "文本控件",
								Lang: "zh_CN",
							},
						},
						Value: &ControlValue{
							Text:        "文本填写的内容",
							Tips:        []interface{}{},
							Members:     []*ContactMember{},
							Departments: []*ContactDepartment{},
							Files:       []*FileValue{},
							Children:    []*TableValue{},
							StatField:   []interface{}{},
						},
					},
				},
			},
			Comments: []*ApprovalComment{
				{
					CommentUserInfo: &OAUser{
						UserID: "WuJunJie",
					},
					CommentTime:    1569584111,
					CommentContent: "这是备注信息",
					CommentID:      "6741314136717778040",
					MediaID:        []string{"WWCISP_Xa1dXIyC9VC2vGTXyBjUXh4GQ31G-a7jilEjFjkYBfncSJv0kM1cZAIXULWbbtosVqA7hprZIUkl4GP0DYZKDrIay9vCzeQelmmHiczwfn80v51EtuNouzBhUBTWo9oQIIzsSftjaVmd4EC_dj5-rayfDl6yIIRdoUs1V_Gz6Pi3yH37ELOgLNAPYUSJpA6V190Xunl7b0s5K5XC9c7eX5vlJek38rB_a2K-kMFMiM1mHDqnltoPa_NT9QynXuHi"},
				},
			},
		},
	}, result)
}

func TestGetOpenApprovalData(t *testing.T) {
	body := []byte(`{"thirdNo":"thirdNoxxx"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"data": {
		"ThirdNo": "thirdNoxxx",
		"OpenTemplateId": "1234567111",
		"OpenSpName": "付款",
		"OpenSpstatus": 1,
		"ApplyTime": 1527837645,
		"ApplyUsername": "jackiejjwu",
		"ApplyUserParty": "产品部",
		"ApplyUserImage": "http://www.qq.com/xxx.png",
		"ApplyUserId": "WuJunJie",
		"ApprovalNodes": {
			"ApprovalNode": [
				{
					"NodeStatus": 1,
					"NodeAttr": 1,
					"NodeType": 1,
					"Items": {
						"Item": [
							{
								"ItemName": "chauvetxiao",
								"ItemParty": "产品部",
								"ItemImage": "http://www.qq.com/xxx.png",
								"ItemUserId": "XiaoWen",
								"ItemStatus": 1,
								"ItemSpeech": "",
								"ItemOpTime": 0
							}
						]
					}
				}
			]
		},
		"NotifyNodes": {
			"NotifyNode": [
				{
					"ItemName": "jinhuiguo",
					"ItemParty": "行政部",
					"ItemImage": "http://www.qq.com/xxx.png",
					"ItemUserId": "GuoJinHui"
				}
			]
		},
		"ApproverStep": 0
	}
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/corp/getopenapprovaldata?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultOpenApprovalData)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetOpenApprovalData("thirdNoxxx", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultOpenApprovalData{
		Data: &OpenApprovalData{
			ThirdNO:        "thirdNoxxx",
			OpenTemplateID: "1234567111",
			OpenSPName:     "付款",
			OpenSPStatus:   1,
			ApplyTime:      1527837645,
			ApplyUsername:  "jackiejjwu",
			ApplyUserParty: "产品部",
			ApplyUserImage: "http://www.qq.com/xxx.png",
			ApplyUserID:    "WuJunJie",
			ApprovalNodes: &ApprovalNodes{
				ApprovalNode: []*ApprovalNode{
					{
						NodeStatus: 1,
						NodeAttr:   1,
						NodeType:   1,
						Items: &ApprovalNodeItems{
							Item: []*ApprovalNodeItem{
								{
									ItemName:   "chauvetxiao",
									ItemParty:  "产品部",
									ItemImage:  "http://www.qq.com/xxx.png",
									ItemUserID: "XiaoWen",
									ItemStatus: 1,
									ItemSpeech: "",
									ItemOPTime: 0,
								},
							},
						},
					},
				},
			},
			NotifyNodes: &ApprovalNotifyNodes{
				NotifyNode: []*ApprovalNotifyNode{
					{
						ItemName:   "jinhuiguo",
						ItemParty:  "行政部",
						ItemImage:  "http://www.qq.com/xxx.png",
						ItemUserID: "GuoJinHui",
					},
				},
			},
			ApproverStep: 0,
		},
	}, result)
}
