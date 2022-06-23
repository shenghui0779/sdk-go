package school

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

func TestGetHealthReportStat(t *testing.T) {
	body := []byte(`{"date":"2020-03-27"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"pv": 100,
	"uv": 50
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/health/get_health_report_stat?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultHealthReportStat)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetHealthReportStat("2020-03-27", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultHealthReportStat{
		PV: 100,
		UV: 50,
	}, result)
}

func TestGetHealthReportJobIDs(t *testing.T) {
	body := []byte(`{"offset":1,"limit":100}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"ending": 1,
	"jobids": [
		"jobid1",
		"jobid2"
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/health/get_report_jobids?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultHealthReportJobIDs)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetHealthReportJobIDs(1, 100, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultHealthReportJobIDs{
		Ending: 1,
		JobIDs: []string{"jobid1", "jobid2"},
	}, result)
}

func TestGetHealthReportJobInfo(t *testing.T) {
	body := []byte(`{"jobid":"jobid1","date":"2020-03-27"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"job_info": {
		"title": "职工收集任务",
		"creator": "creator_userid",
		"type": 1,
		"apply_range": {
			"userids": [
				"userid1",
				"userid2"
			],
			"partyids": [
				1,
				2
			]
		},
		"report_to": {
			"userids": [
				"userid1",
				"userid2"
			]
		},
		"report_type": 1,
		"skip_weekend": 0,
		"finish_cnt": 10,
		"question_templates": [
			{
				"question_id": 1,
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
			},
			{
				"question_id": 2,
				"title": "常驻地址",
				"question_type": 1,
				"is_required": 0
			}
		]
	}
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/health/get_report_job_info?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultHealthReportJobInfo)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetHealthReportJobInfo("jobid1", "2020-03-27", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultHealthReportJobInfo{
		JobInfo: &HealthReportJobInfo{
			Title:   "职工收集任务",
			Creator: "creator_userid",
			Type:    1,
			ApplyRange: &HealthReportApplyRange{
				UserIDs:  []string{"userid1", "userid2"},
				PartyIDs: []int64{1, 2},
			},
			ReportTo: &HealthReportTo{
				UserIDs: []string{"userid1", "userid2"},
			},
			ReportType:  1,
			SkipWeekend: 0,
			FinishCnt:   10,
			QuestionTemplates: []*HealthQuestionTemplate{
				{
					QuestionID:   1,
					Title:        "请问你有任何身体不适吗？",
					QuestionType: 2,
					IsRequired:   1,
					OptionList: []*HealthQuestionOption{
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
				{
					QuestionID:   2,
					Title:        "常驻地址",
					QuestionType: 1,
					IsRequired:   0,
				},
			},
		},
	}, result)
}

func TestGetHealthReportAnswer(t *testing.T) {
	body := []byte(`{"jobid":"jobid1","date":"2020-03-27","offset":1,"limit":100}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"answers": [
		{
			"id_type": 1,
			"userid": "userid2",
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
				},
				{
					"question_id": 4,
					"fileid": [
						"XXXXXXX"
					]
				}
			]
		},
		{
			"id_type": 2,
			"student_userid": "student_userid1",
			"parent_userid": "parent_userid1",
			"report_time": 123456789,
			"report_values": [
				{
					"question_id": 1,
					"single_choice": 1
				},
				{
					"question_id": 2,
					"text": "广东省深圳市"
				},
				{
					"question_id": 3,
					"multi_choice": [
						1,
						2,
						3
					]
				},
				{
					"question_id": 4,
					"fileid": [
						"XXXXXXX"
					]
				}
			]
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/health/get_report_answer?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultHealthReportAnswer)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetHealthReportAnswer("jobid1", "2020-03-27", 1, 100, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultHealthReportAnswer{
		Answers: []*HealthReportAnswer{
			{
				IDType:     1,
				UserID:     "userid2",
				ReportTime: 123456789,
				ReportValues: []*HealthReportValue{
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
					{
						QuestionID: 4,
						FileID:     []string{"XXXXXXX"},
					},
				},
			},
			{
				IDType:        2,
				StudentUserID: "student_userid1",
				ParentUserID:  "parent_userid1",
				ReportTime:    123456789,
				ReportValues: []*HealthReportValue{
					{
						QuestionID:   1,
						SingleChoice: 1,
					},
					{
						QuestionID: 2,
						Text:       "广东省深圳市",
					},
					{
						QuestionID:  3,
						MultiChoice: []int{1, 2, 3},
					},
					{
						QuestionID: 4,
						FileID:     []string{"XXXXXXX"},
					},
				},
			},
		},
	}, result)
}

func TestGetTeacherCustomizeHealthInfo(t *testing.T) {
	body := []byte(`{"date":"2020-03-27","next_key":"NEXT_KEY","limit":100}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"health_infos": [
		{
			"userid": "zhangsan",
			"health_qrcode_status": 1,
			"self_submit": 2,
			"report_values": [
				{
					"question_id": 1,
					"single_chose": 1
				},
				{
					"question_id": 2,
					"text": "浑身难受"
				}
			]
		},
		{
			"userid": "lisi",
			"health_qrcode_status": 2,
			"self_submit": 1,
			"report_values": [
				{
					"question_id": 1,
					"single_chose": 2
				}
			]
		}
	],
	"question_templates": [
		{
			"question_id": 1,
			"title": "请问你有任何身体不适吗？",
			"question_type": 2,
			"is_must_fill": 1,
			"is_not_display": 2,
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
		},
		{
			"question_id": 2,
			"title": "具体哪里不适？（第一题为没有的可以不答）",
			"question_type": 1,
			"is_must_fill": 2,
			"is_not_display": 2
		}
	],
	"template_id": "XXXXXXXXXXXXXXXXX",
	"ending": 1,
	"next_key": "NEXT_KEY"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/school/user/get_teacher_customize_health_info?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultCustomizeHealthInfo)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetTeacherCustomizeHealthInfo("2020-03-27", "NEXT_KEY", 100, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultCustomizeHealthInfo{
		HealthInfos: []*CustomizeHealthInfo{
			{
				UserID:             "zhangsan",
				HealthQRCodeStatus: 1,
				SelfSubmit:         2,
				ReportValues: []*CustomizeHealthReportValue{
					{
						QuestionID:  1,
						SingleChose: 1,
					},
					{
						QuestionID: 2,
						Text:       "浑身难受",
					},
				},
			},
			{
				UserID:             "lisi",
				HealthQRCodeStatus: 2,
				SelfSubmit:         1,
				ReportValues: []*CustomizeHealthReportValue{
					{
						QuestionID:  1,
						SingleChose: 2,
					},
				},
			},
		},
		QuestionTemplates: []*CustomizeHealthQuestionTemplate{
			{
				QuestionID:   1,
				Title:        "请问你有任何身体不适吗？",
				QuestionType: 2,
				IsMustFill:   1,
				IsNotDisplay: 2,
				OptionList: []*HealthQuestionOption{
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
			{
				QuestionID:   2,
				Title:        "具体哪里不适？（第一题为没有的可以不答）",
				QuestionType: 1,
				IsMustFill:   2,
				IsNotDisplay: 2,
			},
		},
		TemplateID: "XXXXXXXXXXXXXXXXX",
		Ending:     1,
		NextKey:    "NEXT_KEY",
	}, result)
}

func TestGetStudentCustomizeHealthInfo(t *testing.T) {
	body := []byte(`{"date":"2020-03-27","next_key":"NEXT_KEY","limit":100}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"health_infos": [
		{
			"userid": "zhangsan",
			"health_qrcode_status": 1,
			"self_submit": 2,
			"report_values": [
				{
					"question_id": 1,
					"single_chose": 1
				},
				{
					"question_id": 2,
					"text": "浑身难受"
				}
			]
		},
		{
			"userid": "lisi",
			"health_qrcode_status": 2,
			"self_submit": 1,
			"report_values": [
				{
					"question_id": 1,
					"single_chose": 2
				}
			]
		}
	],
	"question_templates": [
		{
			"question_id": 1,
			"title": "请问你有任何身体不适吗？",
			"question_type": 2,
			"is_must_fill": 1,
			"is_not_display": 2,
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
		},
		{
			"question_id": 2,
			"title": "具体哪里不适？（第一题为没有的可以不答）",
			"question_type": 1,
			"is_must_fill": 2,
			"is_not_display": 2
		}
	],
	"template_id": "XXXXXXXXXXXXXXXXX",
	"ending": 1,
	"next_key": "NEXT_KEY"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/school/user/get_student_customize_health_info?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultCustomizeHealthInfo)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetStudentCustomizeHealthInfo("2020-03-27", "NEXT_KEY", 100, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultCustomizeHealthInfo{
		HealthInfos: []*CustomizeHealthInfo{
			{
				UserID:             "zhangsan",
				HealthQRCodeStatus: 1,
				SelfSubmit:         2,
				ReportValues: []*CustomizeHealthReportValue{
					{
						QuestionID:  1,
						SingleChose: 1,
					},
					{
						QuestionID: 2,
						Text:       "浑身难受",
					},
				},
			},
			{
				UserID:             "lisi",
				HealthQRCodeStatus: 2,
				SelfSubmit:         1,
				ReportValues: []*CustomizeHealthReportValue{
					{
						QuestionID:  1,
						SingleChose: 2,
					},
				},
			},
		},
		QuestionTemplates: []*CustomizeHealthQuestionTemplate{
			{
				QuestionID:   1,
				Title:        "请问你有任何身体不适吗？",
				QuestionType: 2,
				IsMustFill:   1,
				IsNotDisplay: 2,
				OptionList: []*HealthQuestionOption{
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
			{
				QuestionID:   2,
				Title:        "具体哪里不适？（第一题为没有的可以不答）",
				QuestionType: 1,
				IsMustFill:   2,
				IsNotDisplay: 2,
			},
		},
		TemplateID: "XXXXXXXXXXXXXXXXX",
		Ending:     1,
		NextKey:    "NEXT_KEY",
	}, result)
}

func TestGetHealthQRCode(t *testing.T) {
	body := []byte(`{"type":1,"userids":["userid1","userid2"]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"result_list": [
		{
			"errcode": 0,
			"errmsg": "ok",
			"userid": "userid1",
			"qrcode_data": "QRCODE_DATA"
		},
		{
			"errcode": 1,
			"errmsg": "record not found",
			"userid": "userid1"
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/school/user/get_health_qrcode?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	userIDs := []string{"userid1", "userid2"}

	result := new(ResultHealthQRCode)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetHealthQRCode(1, userIDs, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultHealthQRCode{
		ResultList: []*HealthQRCode{
			{
				ErrCode:    0,
				ErrMsg:     "ok",
				UserID:     "userid1",
				QRCodeData: "QRCODE_DATA",
			},
			{
				ErrCode: 1,
				ErrMsg:  "record not found",
				UserID:  "userid1",
			},
		},
	}, result)
}
