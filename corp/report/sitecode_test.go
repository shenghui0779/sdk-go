package report

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/corp"
	"github.com/shenghui0779/gochat/mock"
)

func TestListSiteCode(t *testing.T) {
	body := []byte(`{"cursor":"CURSOR","limit":100}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"site_code_infos": [
		{
			"id": "siteid",
			"type": "商场超市",
			"area": "广东省/广州市/荔湾区",
			"address": "广州市广州大饭店",
			"name": "广州市广州大饭店",
			"admin": [
				"zhangsan",
				"lisi"
			],
			"qr_code_url": "https://www.abc"
		}
	],
	"next_cursor": "NEXT_CURSOR"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/report/sitecode/list?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultSiteCodeList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListSiteCode("CURSOR", 100, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultSiteCodeList{
		SiteCodeInfos: []*SiteCodeInfo{
			{
				ID:        "siteid",
				Type:      "商场超市",
				Area:      "广东省/广州市/荔湾区",
				Address:   "广州市广州大饭店",
				Name:      "广州市广州大饭店",
				Admin:     []string{"zhangsan", "lisi"},
				QRCodeURL: "https://www.abc",
			},
		},
		NextCursor: "NEXT_CURSOR",
	}, result)
}

func TestGetSiteCodeReportInfo(t *testing.T) {
	body := []byte(`{"siteid":"xxxx"}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"question_templates": [
		{
			"question_id": 1,
			"title": "常驻地址",
			"question_type": 1,
			"is_required": 0
		},
		{
			"question_id": 2,
			"title": "请问你有任何身体不适吗？",
			"question_type": 2,
			"is_required": 1,
			"option_list": [
				{
					"option_id": 1,
					"option_text": "有"
				},
				{
					"option_id": 2,
					"option_text": "没有"
				}
			]
		}
	]
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/report/sitecode/get_site_report_info?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultSiteCodeReportInfo)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetSiteCodeReportInfo("xxxx", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultSiteCodeReportInfo{
		QuestionTemplates: []*QuestionTemplate{
			{
				QuestionID:   1,
				Title:        "常驻地址",
				QuestionType: 1,
				IsRequired:   0,
			},
			{
				QuestionID:   2,
				Title:        "请问你有任何身体不适吗？",
				QuestionType: 2,
				IsRequired:   1,
				OptionList: []*QuestionOption{
					{
						OptionID:   1,
						OptionText: "有",
					},
					{
						OptionID:   2,
						OptionText: "没有",
					},
				},
			},
		},
	}, result)
}

func TestGetSiteCodeReportAnswer(t *testing.T) {
	body := []byte(`{"siteid":"siteid","date":"2020-03-27","cursor":"cursor","limit":100}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"answers": [
		{
			"report_time": 123456789,
			"report_values": [
				{
					"question_id": 1,
					"single_choice": 2
				},
				{
					"question_id": 2,
					"text": "广东省广州市"
				},
				{
					"question_id": 3,
					"multi_choice": [
						1,
						3
					]
				}
			]
		}
	],
	"next_cursor": "NEXT_CURSOR",
	"has_more": 0
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/report/sitecode/get_report_answer?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultSiteCodeReportAnswer)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetSiteCodeReportAnswer("siteid", "2020-03-27", "cursor", 100, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultSiteCodeReportAnswer{
		Answers: []*QuestionAnswer{
			{
				ReportTime: 123456789,
				ReportValues: []*ReportValue{
					{
						QuestionID:   1,
						SingleChoice: 2,
					},
					{
						QuestionID: 2,
						Text:       "广东省广州市",
					},
					{
						QuestionID:  3,
						MultiChoice: []int{1, 3},
					},
				},
			},
		},
		NextCursor: "NEXT_CURSOR",
		HasMore:    0,
	}, result)
}
