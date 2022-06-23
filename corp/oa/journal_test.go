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

func TestListJouralRecord(t *testing.T) {
	body := []byte(`{"starttime":1606230000,"endtime":1606361304,"cursor":0,"limit":10,"filters":[{"key":"creator","value":"kele"},{"key":"department","value":"1"},{"key":"template_id","value":"3TmALk1ogfgKiQE3e3jRwnTUhMTh8vca1N8zUVNUx"}]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"journaluuid_list": [
		"41eJejN57EJNzr8HrZfmKyCN7xwKw1qRxCZUxCVuo9fsWVMSKac6nk4q8rARTDaVNdg",
		"41eJejN57EJNzr8HrZfmKy7rmnZS5HGzpqUefyqCRhjdY9GWQQ6gcaNfaW6GPAdG5cg",
		"41eJejN57EJNzr8HrZfmKy2mkwnjMJPgE6UZfqnW5qMeZ1ag3qr1Amb98DbtVH89VJx",
		"41eJejN57EJNzr8HrZfmKyGXVp9cRByeSREpFtReMKpuAPYZYiCU4em8JKJNmCBYmxg",
		"41eJejN57EJNzr8HrZfmKy3NphvW9E8bYRTAMWcwo9oPhVEFv9cE2jUry8ZNsZYjuUx",
		"41eJejN57EJNzr8HrZfmKyDqJCnct6mYayM4tiEXGmoYmfUp1nDdNQSyxemtBHZa3ss",
		"41eJejN57EJNzr8HrZfmKyHr64ZdZa6JHYztDaS6hCmPMKtBN3YvD1FSFmauNU36Wxd",
		"41eJejN57EJNzr8HrZfmKyChHx58aDhGrvN7yKywBJs33yzUyqUF11sdBFcUBou2NQx",
		"41eJejN57EJNzr8HrZfmKy4w4AtPJyxQoGWmv7hnrZYwmdWVJQEhvgxT5mjEbC1xP43",
		"41eJejN57EJNzr8HrZfmKyFcSr1RLmAoBS7fnwiFcQJuVQfYZwcork67DZ36YFijmR2"
	],
	"next_cursor": 34,
	"endflag": 0
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/oa/journal/get_record_list?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsJournalRecordList{
		StartTime: 1606230000,
		EndTime:   1606361304,
		Cursor:    0,
		Limit:     10,
		Filters: []*KeyValue{
			{
				Key:   "creator",
				Value: "kele",
			},
			{
				Key:   "department",
				Value: "1",
			},
			{
				Key:   "template_id",
				Value: "3TmALk1ogfgKiQE3e3jRwnTUhMTh8vca1N8zUVNUx",
			},
		},
	}
	result := new(ResultJournalRecordList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListJouralRecord(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultJournalRecordList{
		JournalUUIDList: []string{
			"41eJejN57EJNzr8HrZfmKyCN7xwKw1qRxCZUxCVuo9fsWVMSKac6nk4q8rARTDaVNdg",
			"41eJejN57EJNzr8HrZfmKy7rmnZS5HGzpqUefyqCRhjdY9GWQQ6gcaNfaW6GPAdG5cg",
			"41eJejN57EJNzr8HrZfmKy2mkwnjMJPgE6UZfqnW5qMeZ1ag3qr1Amb98DbtVH89VJx",
			"41eJejN57EJNzr8HrZfmKyGXVp9cRByeSREpFtReMKpuAPYZYiCU4em8JKJNmCBYmxg",
			"41eJejN57EJNzr8HrZfmKy3NphvW9E8bYRTAMWcwo9oPhVEFv9cE2jUry8ZNsZYjuUx",
			"41eJejN57EJNzr8HrZfmKyDqJCnct6mYayM4tiEXGmoYmfUp1nDdNQSyxemtBHZa3ss",
			"41eJejN57EJNzr8HrZfmKyHr64ZdZa6JHYztDaS6hCmPMKtBN3YvD1FSFmauNU36Wxd",
			"41eJejN57EJNzr8HrZfmKyChHx58aDhGrvN7yKywBJs33yzUyqUF11sdBFcUBou2NQx",
			"41eJejN57EJNzr8HrZfmKy4w4AtPJyxQoGWmv7hnrZYwmdWVJQEhvgxT5mjEbC1xP43",
			"41eJejN57EJNzr8HrZfmKyFcSr1RLmAoBS7fnwiFcQJuVQfYZwcork67DZ36YFijmR2",
		},
		NextCursor: 34,
		EndFlag:    0,
	}, result)
}

func TestGetJournalRecordDetail(t *testing.T) {
	body := []byte(`{"journaluuid":"41eJejN57EJNzr8HrZfmKyCN7xwKw1qRxCZUxCVuo9fsWVMSKac6nk4q8rARTDaVNdx"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"info": {
		"journal_uuid": "41eJejN57EJNzr8HrZfmKyJZ6E3W9NQbr94x6QEA6MwvK2sVqFQNWy4BaF4Ptyzk26",
		"template_name": "今日工作汇报",
		"report_time": 1606365591,
		"submitter": {
			"userid": "LiQiJun"
		},
		"receivers": [
			{
				"userid": "LiQiJun"
			}
		],
		"readed_receivers": [
			{
				"userid": "LiQiJun"
			}
		],
		"apply_data": {
			"contents": [
				{
					"control": "Text",
					"id": "Text-1606365477123",
					"title": [
						{
							"text": "工作内容",
							"lang": "zh_CN"
						}
					],
					"value": {
						"text": "今日暂无工作",
						"tips": [],
						"members": [],
						"departments": [],
						"files": [],
						"children": [],
						"stat_field": [],
						"sum_field": [],
						"related_approval": [],
						"students": [],
						"classes": []
					}
				}
			]
		},
		"comments": [
			{
				"commentid": 6899287783354824502,
				"tocommentid": 0,
				"comment_userinfo": {
					"userid": "LiYiBo"
				},
				"content": "加油",
				"comment_time": 1606365615
			}
		]
	}
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/oa/journal/get_record_detail?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultJournalRecordDetail)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetJournalRecordDetail("41eJejN57EJNzr8HrZfmKyCN7xwKw1qRxCZUxCVuo9fsWVMSKac6nk4q8rARTDaVNdx", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultJournalRecordDetail{
		Info: &JournalRecordDetail{
			JournalUUID:  "41eJejN57EJNzr8HrZfmKyJZ6E3W9NQbr94x6QEA6MwvK2sVqFQNWy4BaF4Ptyzk26",
			TemplateName: "今日工作汇报",
			ReportTime:   1606365591,
			Submitter: &OAUser{
				UserID: "LiQiJun",
			},
			Receivers: []*OAUser{
				{
					UserID: "LiQiJun",
				},
			},
			ReadedReceivers: []*OAUser{
				{
					UserID: "LiQiJun",
				},
			},
			ApplyData: &ApplyData{
				Contents: []*ApplyContent{
					{
						Control: "Text",
						ID:      "Text-1606365477123",
						Title: []*DisplayText{
							{
								Text: "工作内容",
								Lang: "zh_CN",
							},
						},
						Value: &ControlValue{
							Text:            "今日暂无工作",
							Tips:            []interface{}{},
							Members:         []*ContactMember{},
							Departments:     []*ContactDepartment{},
							Files:           []*FileValue{},
							Children:        []*TableValue{},
							StatField:       []interface{}{},
							SumField:        []interface{}{},
							Students:        []*SchoolContactStudent{},
							Classes:         []*SchoolContactClass{},
							RelatedApproval: []*RelatedApprovalValue{},
						},
					},
				},
			},
			Comments: []*JournalRecordComment{
				{
					CommentID:   6899287783354824502,
					ToCommentID: 0,
					CommentUserInfo: &OAUser{
						UserID: "LiYiBo",
					},
					Content:     "加油",
					CommentTime: 1606365615,
				},
			},
		},
	}, result)
}

func TestListJournalStat(t *testing.T) {
	body := []byte(`{"template_id":"3TmALk1ogfgKiQE3e3jRwnTUhMTh8vca1N8zUVNUx","starttime":1604160000,"endtime":1606363092}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"stat_list": [
		{
			"template_id": "3TmALk1ogfgKiQE3e3jRwnTUhMTh8vca1N8zUVNU",
			"template_name": "日报",
			"report_range": {
				"user_list": [
					{
						"userid": "user1"
					}
				],
				"party_list": [
					{
						"open_partyid": "1"
					}
				],
				"tag_list": []
			},
			"white_range": {
				"user_list": [],
				"party_list": [],
				"tag_list": []
			},
			"receivers": {
				"user_list": [
					{
						"userid": "user3"
					}
				],
				"tag_list": [],
				"leader_list": []
			},
			"cycle_begin_time": 1606147200,
			"cycle_end_time": 1606233600,
			"stat_begin_time": 1606147200,
			"stat_end_time": 1606230000,
			"report_list": [
				{
					"user": {
						"userid": "user2"
					},
					"itemlist": [
						{
							"journaluuid": "4U9abSUrpY78VNxeNNv3J5TW5e9VLj8cDymH9py1Efpuj5X8QCDQx3stKr69pia3UL8auRjrCMsiRjgzL8mvKnff",
							"reporttime": 1606218548,
							"flag": 0
						}
					]
				}
			],
			"unreport_list": [
				{
					"user": {
						"userid": "user1"
					},
					"itemlist": [
						{
							"journaluuid": "",
							"reporttime": 1606147200,
							"flag": 0
						}
					]
				},
				{
					"user": {
						"userid": "user3"
					},
					"itemlist": [
						{
							"journaluuid": "",
							"reporttime": 1606147200,
							"flag": 0
						}
					]
				}
			],
			"report_type": 2
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/oa/journal/get_stat_list?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultJournalStatList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListJournalStat("3TmALk1ogfgKiQE3e3jRwnTUhMTh8vca1N8zUVNUx", 1604160000, 1606363092, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultJournalStatList{
		StatList: []*JournalStat{
			{
				TemplateID:   "3TmALk1ogfgKiQE3e3jRwnTUhMTh8vca1N8zUVNU",
				TemplateName: "日报",
				ReportRange: &JournalRange{
					UserList: []*JournalUser{
						{
							UserID: "user1",
						},
					},
					PartyList: []*JournalParty{
						{
							OpenPartyID: "1",
						},
					},
					TagList: []*JournalTag{},
				},
				WhiteRange: &JournalRange{
					UserList:  []*JournalUser{},
					PartyList: []*JournalParty{},
					TagList:   []*JournalTag{},
				},
				Receivers: &JournalReceivers{
					UserList: []*JournalUser{
						{
							UserID: "user3",
						},
					},
					TagList:    []*JournalTag{},
					LeaderList: []*JournalLeader{},
				},
				CycleBeginTime: 1606147200,
				CycleEndTime:   1606233600,
				StatBeginTime:  1606147200,
				StatEndTime:    1606230000,
				ReportList: []*JournalReport{
					{
						User: &OAUser{
							UserID: "user2",
						},
						ItemList: []*JournalReportItem{
							{
								JournalUUID: "4U9abSUrpY78VNxeNNv3J5TW5e9VLj8cDymH9py1Efpuj5X8QCDQx3stKr69pia3UL8auRjrCMsiRjgzL8mvKnff",
								ReportTime:  1606218548,
								Flag:        0,
							},
						},
					},
				},
				UnReportList: []*JournalReport{
					{
						User: &OAUser{
							UserID: "user1",
						},
						ItemList: []*JournalReportItem{
							{
								JournalUUID: "",
								ReportTime:  1606147200,
								Flag:        0,
							},
						},
					},
					{
						User: &OAUser{
							UserID: "user3",
						},
						ItemList: []*JournalReportItem{
							{
								JournalUUID: "",
								ReportTime:  1606147200,
								Flag:        0,
							},
						},
					},
				},
				ReportType: 2,
			},
		},
	}, result)
}
