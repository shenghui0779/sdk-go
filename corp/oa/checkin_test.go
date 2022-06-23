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

func TestGetCorpCheckinOption(t *testing.T) {
	body := []byte(`{}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"group": [
		{
			"grouptype": 1,
			"groupid": 69,
			"groupname": "打卡规则1",
			"checkindate": [
				{
					"workdays": [
						1,
						2,
						3,
						4,
						5
					],
					"checkintime": [
						{
							"work_sec": 36000,
							"off_work_sec": 43200,
							"remind_work_sec": 35400,
							"remind_off_work_sec": 43200
						},
						{
							"work_sec": 50400,
							"off_work_sec": 72000,
							"remind_work_sec": 49800,
							"remind_off_work_sec": 72000
						}
					],
					"noneed_offwork": true,
					"limit_aheadtime": 10800000,
					"flex_on_duty_time": 0,
					"flex_off_duty_time": 0
				}
			],
			"spe_workdays": [
				{
					"timestamp": 1512144000,
					"notes": "必须打卡的日期",
					"checkintime": [
						{
							"work_sec": 32400,
							"off_work_sec": 61200,
							"remind_work_sec": 31800,
							"remind_off_work_sec": 61200
						}
					]
				}
			],
			"spe_offdays": [
				{
					"timestamp": 1512057600,
					"notes": "不需要打卡的日期",
					"checkintime": []
				}
			],
			"sync_holidays": true,
			"need_photo": true,
			"wifimac_infos": [
				{
					"wifiname": "Tencent-WiFi-1",
					"wifimac": "c0:7b:bc:37:f8:d3"
				},
				{
					"wifiname": "Tencent-WiFi-2",
					"wifimac": "70:10:5c:7d:f6:d5"
				}
			],
			"note_can_use_local_pic": false,
			"allow_checkin_offworkday": true,
			"allow_apply_offworkday": true,
			"loc_infos": [
				{
					"lat": 30547030,
					"lng": 104062890,
					"loc_title": "腾讯成都大厦",
					"loc_detail": "四川省成都市武侯区高新南区天府三街",
					"distance": 300
				},
				{
					"lat": 23097490,
					"lng": 113323750,
					"loc_title": "T.I.T创意园",
					"loc_detail": "广东省广州市海珠区新港中路397号",
					"distance": 300
				}
			],
			"range": {
				"partyid": [],
				"userid": [
					"icef",
					"LiJingZhong"
				],
				"tagid": [
					2
				]
			},
			"create_time": 1606204343,
			"white_users": [
				"canno"
			],
			"type": 0,
			"reporterinfo": {
				"reporters": [
					{
						"userid": "brant"
					}
				],
				"updatetime": 1606305508
			},
			"ot_info": {
				"type": 2,
				"allow_ot_workingday": true,
				"allow_ot_nonworkingday": false,
				"otcheckinfo": {
					"ot_workingday_time_start": 1800,
					"ot_workingday_time_min": 1800,
					"ot_workingday_time_max": 14400,
					"ot_nonworkingday_time_min": 1800,
					"ot_nonworkingday_time_max": 14400,
					"ot_workingday_restinfo": {
						"type": 2,
						"fix_time_rule": {
							"fix_time_begin_sec": 43200,
							"fix_time_end_sec": 46800
						},
						"cal_ottime_rule": {
							"items": [
								{
									"ot_time": 18000,
									"rest_time": 3600
								}
							]
						}
					},
					"ot_nonworkingday_restinfo": {
						"type": 2,
						"fix_time_rule": {
							"fix_time_begin_sec": 43200,
							"fix_time_end_sec": 46800
						},
						"cal_ottime_rule": {
							"items": [
								{
									"ot_time": 18000,
									"rest_time": 3600
								}
							]
						}
					},
					"ot_nonworkingday_spanday_time": 0
				},
				"uptime": 1606275664,
				"otapplyinfo": {
					"allow_ot_workingday": true,
					"allow_ot_nonworkingday": true,
					"uptime": 1606275664,
					"ot_workingday_restinfo": {
						"type": 2,
						"fix_time_rule": {
							"fix_time_begin_sec": 43200,
							"fix_time_end_sec": 46800
						},
						"cal_ottime_rule": {
							"items": [
								{
									"ot_time": 18000,
									"rest_time": 3600
								}
							]
						}
					},
					"ot_nonworkingday_restinfo": {
						"type": 2,
						"fix_time_rule": {
							"fix_time_begin_sec": 43200,
							"fix_time_end_sec": 46800
						},
						"cal_ottime_rule": {
							"items": [
								{
									"ot_time": 18000,
									"rest_time": 3600
								}
							]
						}
					},
					"ot_nonworkingday_spanday_time": 0
				}
			},
			"allow_apply_bk_cnt": -1,
			"option_out_range": 0,
			"create_userid": "gaogao",
			"use_face_detect": false,
			"allow_apply_bk_day_limit": -1,
			"update_userid": "sandy",
			"schedulelist": [
				{
					"schedule_id": 221,
					"schedule_name": "2",
					"time_section": [
						{
							"time_id": 1,
							"work_sec": 32400,
							"off_work_sec": 61200,
							"remind_work_sec": 31800,
							"remind_off_work_sec": 61200,
							"rest_begin_time": 43200,
							"rest_end_time": 46800,
							"allow_rest": false
						}
					],
					"limit_aheadtime": 14400000,
					"noneed_offwork": false,
					"limit_offtime": 14400,
					"flex_on_duty_time": 0,
					"flex_off_duty_time": 0,
					"allow_flex": false,
					"late_rule": {
						"allow_offwork_after_time": false,
						"timerules": [
							{
								"offwork_after_time": 3600,
								"onwork_flex_time": 3600
							}
						]
					},
					"max_allow_arrive_early": 0,
					"max_allow_arrive_late": 0
				}
			],
			"offwork_interval_time": 300
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/checkin/getcorpcheckinoption?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultCorpCheckinOption)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetCorpCheckinOption(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultCorpCheckinOption{
		Group: []*CorpCheckinOption{
			{
				GroupType: 1,
				GroupID:   69,
				GroupName: "打卡规则1",
				CheckinDate: []*CheckinDate{
					{
						Workdays: []int{1, 2, 3, 4, 5},
						CheckinTime: []*CheckinTime{
							{
								WorkSec:          36000,
								OffWorkSec:       43200,
								RemindWorkSec:    35400,
								RemindOffWorkSec: 43200,
							},
							{
								WorkSec:          50400,
								OffWorkSec:       72000,
								RemindWorkSec:    49800,
								RemindOffWorkSec: 72000,
							},
						},
						NoNeedOffWork:   true,
						LimitAheadTime:  10800000,
						FlexOnDutyTime:  0,
						FlexOffDutyTime: 0,
					},
				},
				SpeWorkdays: []*CheckinSpeDay{
					{
						Timestamp: 1512144000,
						Notes:     "必须打卡的日期",
						CheckinTime: []*CheckinTime{
							{
								WorkSec:          32400,
								OffWorkSec:       61200,
								RemindWorkSec:    31800,
								RemindOffWorkSec: 61200,
							},
						},
					},
				},
				SpeOffDays: []*CheckinSpeDay{
					{
						Timestamp:   1512057600,
						Notes:       "不需要打卡的日期",
						CheckinTime: []*CheckinTime{},
					},
				},
				SyncHolidays: true,
				NeedPhoto:    true,
				WifiMacInfos: []*CheckinWifiMac{
					{
						WifiName: "Tencent-WiFi-1",
						WifiMac:  "c0:7b:bc:37:f8:d3",
					},
					{
						WifiName: "Tencent-WiFi-2",
						WifiMac:  "70:10:5c:7d:f6:d5",
					},
				},
				NoteCanUseLocalPic:     false,
				AllowCheckinOffWorkday: true,
				AllowApplyOffWorkday:   true,
				LocInfos: []*CheckinLoc{
					{
						Lat:       30547030,
						Lng:       104062890,
						LocTitle:  "腾讯成都大厦",
						LocDetail: "四川省成都市武侯区高新南区天府三街",
						Distance:  300,
					},
					{
						Lat:       23097490,
						Lng:       113323750,
						LocTitle:  "T.I.T创意园",
						LocDetail: "广东省广州市海珠区新港中路397号",
						Distance:  300,
					},
				},
				Range: &CheckinRange{
					PartyID: []string{},
					UserID:  []string{"icef", "LiJingZhong"},
					TagID:   []int64{2},
				},
				CreateTime: 1606204343,
				WhiteUsers: []string{"canno"},
				Type:       0,
				ReporterInfo: &CheckinReporter{
					Reporters: []*OAUser{
						{
							UserID: "brant",
						},
					},
					UpdateTime: 1606305508,
				},
				OtInfo: &CheckinOt{
					Type:                 2,
					AllowOtWorkingDay:    true,
					AllowOtNonWorkingDay: false,
					OtCheckInfo: &CheckinOtCheck{
						OtWorkingDayTimeStart:  1800,
						OtWorkingDayTimeMin:    1800,
						OtWorkingDayTimeMax:    14400,
						OtNonWorkingDayTimeMin: 1800,
						OtNonWorkingDayTimeMax: 14400,
						OtWorkingDayRestInfo: &CheckinOtRest{
							Type: 2,
							FixTimeRule: &CheckinFixTimeRule{
								FixTimeBeginSec: 43200,
								FixTimeEndSec:   46800,
							},
							CalOtTimeRule: &CheckinCalOtTimeRule{
								Items: []*CheckinOtTimeRule{
									{
										OtTime:   18000,
										RestTime: 3600,
									},
								},
							},
						},
						OtNonWorkingDayRestInfo: &CheckinOtRest{
							Type: 2,
							FixTimeRule: &CheckinFixTimeRule{
								FixTimeBeginSec: 43200,
								FixTimeEndSec:   46800,
							},
							CalOtTimeRule: &CheckinCalOtTimeRule{
								Items: []*CheckinOtTimeRule{
									{
										OtTime:   18000,
										RestTime: 3600,
									},
								},
							},
						},
						OtNonWorkingDaySpandayTime: 0,
					},
					UpTime: 1606275664,
					OtApplyInfo: &CheckinOtApply{
						AllowOtWorkingDay:    true,
						AllowOtNonWorkingDay: true,
						UpTime:               1606275664,
						OtWorkingDayRestInfo: &CheckinOtRest{
							Type: 2,
							FixTimeRule: &CheckinFixTimeRule{
								FixTimeBeginSec: 43200,
								FixTimeEndSec:   46800,
							},
							CalOtTimeRule: &CheckinCalOtTimeRule{
								Items: []*CheckinOtTimeRule{
									{
										OtTime:   18000,
										RestTime: 3600,
									},
								},
							},
						},
						OtNonWorkingDayRestInfo: &CheckinOtRest{
							Type: 2,
							FixTimeRule: &CheckinFixTimeRule{
								FixTimeBeginSec: 43200,
								FixTimeEndSec:   46800,
							},
							CalOtTimeRule: &CheckinCalOtTimeRule{
								Items: []*CheckinOtTimeRule{
									{
										OtTime:   18000,
										RestTime: 3600,
									},
								},
							},
						},
						OtNonWorkingDaySpandayTime: 0,
					},
				},
				AllowApplyBKCnt:      -1,
				OptionOutRange:       0,
				CreateUserID:         "gaogao",
				UseFaceDetect:        false,
				AllowApplyBKDayLimit: -1,
				UpdateUserID:         "sandy",
				ScheduleList: []*OptionSchedule{
					{
						ScheduleID:   221,
						ScheduleName: "2",
						TimeSection: []*CheckinTimeSection{
							{
								TimeID:           1,
								WorkSec:          32400,
								OffWorkSec:       61200,
								RemindWorkSec:    31800,
								RemindOffWorkSec: 61200,
								RestBeginTime:    43200,
								RestEndTime:      46800,
								AllowRest:        false,
							},
						},
						LimitAheadTime:  14400000,
						NoNeedOffWork:   false,
						LimitOffTime:    14400,
						FlexOnDutyTime:  0,
						FlexOffDutyTime: 0,
						AllowFlex:       false,
						LateRule: &CheckinLateRule{
							AllowOffWorkAfterTime: false,
							TimeRules: []*CheckinTimeRule{
								{
									OffWorkAfterTime: 3600,
									OnWorkFlexTime:   3600,
								},
							},
						},
						MaxAllowArriveEarly: 0,
						MaxAllowArriveLate:  0,
					},
				},
				OffWorkIntervalTime: 300,
			},
		},
	}, result)
}

func TestGetCheckinOption(t *testing.T) {
	body := []byte(`{"datetime":1511971200,"useridlist":["james","paul"]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"info": [
		{
			"userid": "james",
			"group": {
				"grouptype": 1,
				"groupid": 69,
				"groupname": "打卡规则1",
				"checkindate": [
					{
						"workdays": [
							1,
							2,
							3,
							4,
							5
						],
						"checkintime": [
							{
								"work_sec": 36000,
								"off_work_sec": 43200,
								"remind_work_sec": 35400,
								"remind_off_work_sec": 43200
							},
							{
								"work_sec": 50400,
								"off_work_sec": 72000,
								"remind_work_sec": 49800,
								"remind_off_work_sec": 72000
							}
						],
						"flex_time": 300000,
						"noneed_offwork": true,
						"limit_aheadtime": 10800000,
						"flex_on_duty_time": 0,
						"flex_off_duty_time": 0
					}
				],
				"spe_workdays": [
					{
						"timestamp": 1512144000,
						"notes": "必须打卡的日期",
						"checkintime": [
							{
								"work_sec": 32400,
								"off_work_sec": 61200,
								"remind_work_sec": 31800,
								"remind_off_work_sec": 61200
							}
						]
					}
				],
				"spe_offdays": [
					{
						"timestamp": 1512057600,
						"notes": "不需要打卡的日期",
						"checkintime": []
					}
				],
				"sync_holidays": true,
				"need_photo": true,
				"wifimac_infos": [
					{
						"wifiname": "Tencent-WiFi-1",
						"wifimac": "c0:7b:bc:37:f8:d3"
					},
					{
						"wifiname": "Tencent-WiFi-2",
						"wifimac": "70:10:5c:7d:f6:d5"
					}
				],
				"note_can_use_local_pic": false,
				"allow_checkin_offworkday": true,
				"allow_apply_offworkday": true,
				"loc_infos": [
					{
						"lat": 30547030,
						"lng": 104062890,
						"loc_title": "腾讯成都大厦",
						"loc_detail": "四川省成都市武侯区高新南区天府三街",
						"distance": 300
					},
					{
						"lat": 23097490,
						"lng": 113323750,
						"loc_title": "T.I.T创意园",
						"loc_detail": "广东省广州市海珠区新港中路397号",
						"distance": 300
					}
				],
				"schedulelist": [
					{
						"schedule_id": 221,
						"schedule_name": "2",
						"time_section": [
							{
								"time_id": 1,
								"work_sec": 32400,
								"off_work_sec": 61200,
								"remind_work_sec": 31800,
								"remind_off_work_sec": 61200,
								"rest_begin_time": 43200,
								"rest_end_time": 46800,
								"allow_rest": false
							}
						],
						"limit_aheadtime": 14400000,
						"noneed_offwork": false,
						"limit_offtime": 14400,
						"flex_on_duty_time": 0,
						"flex_off_duty_time": 0,
						"allow_flex": false,
						"late_rule": {
							"allow_offwork_after_time": false,
							"timerules": [
								{
									"offwork_after_time": 3600,
									"onwork_flex_time": 3600
								}
							]
						},
						"max_allow_arrive_early": 0,
						"max_allow_arrive_late": 0
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/checkin/getcheckinoption?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	userIDs := []string{"james", "paul"}

	result := new(ResultCheckinOption)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetCheckinOption(1511971200, userIDs, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultCheckinOption{
		Info: []*CheckinOption{
			{
				UserID: "james",
				Group: &CheckinGroup{
					GroupType: 1,
					GroupID:   69,
					GroupName: "打卡规则1",
					CheckinDate: []*CheckinDate{
						{
							Workdays: []int{1, 2, 3, 4, 5},
							CheckinTime: []*CheckinTime{
								{
									WorkSec:          36000,
									OffWorkSec:       43200,
									RemindWorkSec:    35400,
									RemindOffWorkSec: 43200,
								},
								{
									WorkSec:          50400,
									OffWorkSec:       72000,
									RemindWorkSec:    49800,
									RemindOffWorkSec: 72000,
								},
							},
							FlexTime:        300000,
							NoNeedOffWork:   true,
							LimitAheadTime:  10800000,
							FlexOnDutyTime:  0,
							FlexOffDutyTime: 0,
						},
					},
					SpeWorkdays: []*CheckinSpeDay{
						{
							Timestamp: 1512144000,
							Notes:     "必须打卡的日期",
							CheckinTime: []*CheckinTime{
								{
									WorkSec:          32400,
									OffWorkSec:       61200,
									RemindWorkSec:    31800,
									RemindOffWorkSec: 61200,
								},
							},
						},
					},
					SpeOffDays: []*CheckinSpeDay{
						{
							Timestamp:   1512057600,
							Notes:       "不需要打卡的日期",
							CheckinTime: []*CheckinTime{},
						},
					},
					SyncHolidays: true,
					NeedPhoto:    true,
					WifiMacInfos: []*CheckinWifiMac{
						{
							WifiName: "Tencent-WiFi-1",
							WifiMac:  "c0:7b:bc:37:f8:d3",
						},
						{
							WifiName: "Tencent-WiFi-2",
							WifiMac:  "70:10:5c:7d:f6:d5",
						},
					},
					NoteCanUseLocalPic:     false,
					AllowCheckinOffWorkday: true,
					AllowApplyOffWorkday:   true,
					LocInfos: []*CheckinLoc{
						{
							Lat:       30547030,
							Lng:       104062890,
							LocTitle:  "腾讯成都大厦",
							LocDetail: "四川省成都市武侯区高新南区天府三街",
							Distance:  300,
						},
						{
							Lat:       23097490,
							Lng:       113323750,
							LocTitle:  "T.I.T创意园",
							LocDetail: "广东省广州市海珠区新港中路397号",
							Distance:  300,
						},
					},
					ScheduleList: []*OptionSchedule{
						{
							ScheduleID:   221,
							ScheduleName: "2",
							TimeSection: []*CheckinTimeSection{
								{
									TimeID:           1,
									WorkSec:          32400,
									OffWorkSec:       61200,
									RemindWorkSec:    31800,
									RemindOffWorkSec: 61200,
									RestBeginTime:    43200,
									RestEndTime:      46800,
									AllowRest:        false,
								},
							},
							LimitAheadTime:  14400000,
							NoNeedOffWork:   false,
							LimitOffTime:    14400,
							FlexOnDutyTime:  0,
							FlexOffDutyTime: 0,
							AllowFlex:       false,
							LateRule: &CheckinLateRule{
								AllowOffWorkAfterTime: false,
								TimeRules: []*CheckinTimeRule{
									{
										OffWorkAfterTime: 3600,
										OnWorkFlexTime:   3600,
									},
								},
							},
							MaxAllowArriveEarly: 0,
							MaxAllowArriveLate:  0,
						},
					},
				},
			},
		},
	}, result)
}

func TestGetCheckinData(t *testing.T) {
	body := []byte(`{"opencheckindatatype":3,"starttime":1492617600,"endtime":1492790400,"useridlist":["james","paul"]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
"errcode": 0,
"errmsg": "ok",
"checkindata": [
	{
		"userid": "james",
		"groupid": 1,
		"groupname": "打卡一组",
		"checkin_type": "上班打卡",
		"checkin_time": 1492617610,
		"exception_type": "地点异常",
		"location_title": "依澜府",
		"location_detail": "四川省成都市武侯区益州大道中段784号附近",
		"wifiname": "办公一区",
		"wifimac": "3c:46:d8:0c:7a:70",
		"notes": "路上堵车，迟到了5分钟",
		"mediaids": [
			"WWCISP_G8PYgRaOVHjXWUWFqchpBqqqUpGj0OyR9z6WTwhnMZGCPHxyviVstiv_2fTG8YOJq8L8zJT2T2OvTebANV-2MQ"
		],
		"sch_checkin_time": 1492617610,
		"schedule_id": 0,
		"timeline_id": 2
	},
	{
		"userid": "paul",
		"groupid": 2,
		"groupname": "打卡二组",
		"checkin_type": "外出打卡",
		"checkin_time": 1492617620,
		"exception_type": "时间异常",
		"location_title": "重庆出口加工区",
		"location_detail": "重庆市渝北区金渝大道101号金渝大道",
		"wifiname": "办公室二区",
		"wifimac": "3c:46:d8:0c:7a:71",
		"notes": "",
		"mediaids": [
			"WWCISP_G8PYgRaOVHjXWUWFqchpBqqqUpGj0OyR9z6WTwhnMZGCPHxyviVstiv_2fTG8YOJq8L8zJT2T2OvTebANV-2MQ"
		],
		"lat": 30547645,
		"lng": 104063236,
		"deviceid": "E5FA89F6-3926-4972-BE4F-4A7ACF4701E2",
		"sch_checkin_time": 1492617610,
		"schedule_id": 3,
		"timeline_id": 1
	}
]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/checkin/getcheckindata?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	userIDs := []string{"james", "paul"}

	result := new(ResultCheckinData)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetCheckinData(3, 1492617600, 1492790400, userIDs, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultCheckinData{
		CheckinData: []*CheckinData{
			{
				UserID:         "james",
				GroupID:        1,
				GroupName:      "打卡一组",
				CheckinType:    "上班打卡",
				CheckinTime:    1492617610,
				ExceptionType:  "地点异常",
				LocationTitle:  "依澜府",
				LocationDetail: "四川省成都市武侯区益州大道中段784号附近",
				WifiName:       "办公一区",
				WifiMac:        "3c:46:d8:0c:7a:70",
				Notes:          "路上堵车，迟到了5分钟",
				MediaIDs:       []string{"WWCISP_G8PYgRaOVHjXWUWFqchpBqqqUpGj0OyR9z6WTwhnMZGCPHxyviVstiv_2fTG8YOJq8L8zJT2T2OvTebANV-2MQ"},
				SchCheckinTime: 1492617610,
				ScheduleID:     0,
				TimelineID:     2,
			},
			{
				UserID:         "paul",
				GroupID:        2,
				GroupName:      "打卡二组",
				CheckinType:    "外出打卡",
				CheckinTime:    1492617620,
				ExceptionType:  "时间异常",
				LocationTitle:  "重庆出口加工区",
				LocationDetail: "重庆市渝北区金渝大道101号金渝大道",
				WifiName:       "办公室二区",
				WifiMac:        "3c:46:d8:0c:7a:71",
				Notes:          "",
				MediaIDs:       []string{"WWCISP_G8PYgRaOVHjXWUWFqchpBqqqUpGj0OyR9z6WTwhnMZGCPHxyviVstiv_2fTG8YOJq8L8zJT2T2OvTebANV-2MQ"},
				Lat:            30547645,
				Lng:            104063236,
				DeviceID:       "E5FA89F6-3926-4972-BE4F-4A7ACF4701E2",
				SchCheckinTime: 1492617610,
				ScheduleID:     3,
				TimelineID:     1,
			},
		},
	}, result)
}

func TestGetCheckinDayData(t *testing.T) {
	body := []byte(`{"starttime":1599062400,"endtime":1599062400,"useridlist":["ZhangSan"]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"datas": [
		{
			"base_info": {
				"date": 1599062400,
				"record_type": 1,
				"name": "张三",
				"name_ex": "Three Zhang",
				"departs_name": "有家企业/realempty;有家企业;有家企业/部门A4",
				"acctid": "ZhangSan",
				"rule_info": {
					"groupid": 10,
					"groupname": "规则测试",
					"scheduleid": 0,
					"schedulename": "",
					"checkintime": [
						{
							"work_sec": 38760,
							"off_work_sec": 38880
						}
					]
				},
				"day_type": 0
			},
			"summary_info": {
				"checkin_count": 2,
				"regular_work_sec": 31,
				"standard_work_sec": 120,
				"earliest_time": 38827,
				"lastest_time": 38858
			},
			"holiday_infos": [
				{
					"sp_description": {
						"data": [
							{
								"text": "09/03 10:00~09/03 10:01",
								"lang": "zh_CN"
							}
						]
					},
					"sp_number": "202009030002",
					"sp_title": {
						"data": [
							{
								"text": "请假0.1小时",
								"lang": "zh_CN"
							}
						]
					}
				},
				{
					"sp_description": {
						"data": [
							{
								"text": "08/25 14:37~09/10 14:37",
								"lang": "zh_CN"
							}
						]
					},
					"sp_number": "202008270004",
					"sp_title": {
						"data": [
							{
								"text": "加班17.0小时",
								"lang": "zh_CN"
							}
						]
					}
				}
			],
			"exception_infos": [
				{
					"count": 1,
					"duration": 60,
					"exception": 1
				},
				{
					"count": 1,
					"duration": 60,
					"exception": 2
				}
			],
			"ot_info": {
				"ot_status": 1,
				"ot_duration": 3600,
				"exception_duration": []
			},
			"sp_items": [
				{
					"count": 1,
					"duration": 360,
					"time_type": 0,
					"type": 1,
					"vacation_id": 2,
					"name": "年假"
				},
				{
					"count": 0,
					"duration": 0,
					"time_type": 0,
					"type": 100,
					"vacation_id": 0,
					"name": "外勤次数"
				}
			]
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/checkin/getcheckin_daydata?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	userIDs := []string{"ZhangSan"}

	result := new(ResultCheckinDayData)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetCheckinDayData(1599062400, 1599062400, userIDs, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultCheckinDayData{
		Datas: []*CheckinDayData{
			{
				BaseInfo: &CheckinDayBase{
					Date:        1599062400,
					RecordType:  1,
					Name:        "张三",
					NameEX:      "Three Zhang",
					DepartsName: "有家企业/realempty;有家企业;有家企业/部门A4",
					AcctID:      "ZhangSan",
					RuleInfo: &CheckinDayRule{
						GroupID:      10,
						GroupName:    "规则测试",
						ScheduleID:   0,
						ScheduleName: "",
						CheckinTime: []*CheckinTime{
							{
								WorkSec:    38760,
								OffWorkSec: 38880,
							},
						},
						DayType: 0,
					},
				},
				SummaryInfo: &CheckinDaySummary{
					CheckinCount:    2,
					RegularWorkSec:  31,
					StandardWorkSec: 120,
					EarliestTime:    38827,
					LastestTime:     38858,
				},
				HolidayInfos: []*CheckinHoliday{
					{
						SPDescription: &CheckinSPText{
							Data: []*DisplayText{
								{
									Text: "09/03 10:00~09/03 10:01",
									Lang: "zh_CN",
								},
							},
						},
						SPNumber: "202009030002",
						SPTitle: &CheckinSPText{
							Data: []*DisplayText{
								{
									Text: "请假0.1小时",
									Lang: "zh_CN",
								},
							},
						},
					},
					{
						SPDescription: &CheckinSPText{
							Data: []*DisplayText{
								{
									Text: "08/25 14:37~09/10 14:37",
									Lang: "zh_CN",
								},
							},
						},
						SPNumber: "202008270004",
						SPTitle: &CheckinSPText{
							Data: []*DisplayText{
								{
									Text: "加班17.0小时",
									Lang: "zh_CN",
								},
							},
						},
					},
				},
				ExceptionInfos: []*CheckinException{
					{
						Count:     1,
						Duration:  60,
						Exception: 1,
					},
					{
						Count:     1,
						Duration:  60,
						Exception: 2,
					},
				},
				OtInfo: &CheckinDayOt{
					OtStatus:          1,
					OtDuration:        3600,
					ExceptionDuration: []int{},
				},
				SPItems: []*CheckinSPItem{
					{
						Count:      1,
						Duration:   360,
						TimeType:   0,
						Type:       1,
						VacationID: 2,
						Name:       "年假",
					},
					{
						Count:      0,
						Duration:   0,
						TimeType:   0,
						Type:       100,
						VacationID: 0,
						Name:       "外勤次数",
					},
				},
			},
		},
	}, result)
}

func TestGetCheckinMonthData(t *testing.T) {
	body := []byte(`{"starttime":1599062400,"endtime":1599408000,"useridlist":["ZhangSan"]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"datas": [
		{
			"base_info": {
				"record_type": 1,
				"name": "张三",
				"name_ex": "Three Zhang",
				"departs_name": "有家企业/realempty;有家企业;有家企业/部门A4",
				"acctid": "ZhangSan",
				"rule_info": {
					"groupid": 10,
					"groupname": "规则测试"
				}
			},
			"summary_info": {
				"except_days": 3,
				"regular_work_sec": 31,
				"standard_work_sec": 29040,
				"work_days": 3
			},
			"exception_infos": [
				{
					"count": 2,
					"duration": 28920,
					"exception": 4
				},
				{
					"count": 1,
					"duration": 60,
					"exception": 1
				},
				{
					"count": 1,
					"duration": 60,
					"exception": 2
				}
			],
			"sp_items": [
				{
					"count": 0,
					"duration": 0,
					"time_type": 0,
					"type": 100,
					"vacation_id": 0,
					"name": "外勤次数"
				},
				{
					"count": 1,
					"duration": 0,
					"time_type": 0,
					"type": 1,
					"vacation_id": 2,
					"name": "年假"
				}
			],
			"overwork_info": {
				"workday_over_sec": 10800
			}
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/checkin/getcheckin_monthdata?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	userIDs := []string{"ZhangSan"}

	result := new(ResultCheckinMonthData)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetCheckinMonthData(1599062400, 1599408000, userIDs, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultCheckinMonthData{
		Datas: []*CheckinMonthData{
			{
				BaseInfo: &CheckinMonthBase{
					RecordType:  1,
					Name:        "张三",
					NameEX:      "Three Zhang",
					DepartsName: "有家企业/realempty;有家企业;有家企业/部门A4",
					AcctID:      "ZhangSan",
					RuleInfo: &CheckinMonthRule{
						GroupID:   10,
						GroupName: "规则测试",
					},
				},
				SummaryInfo: &CheckinMonthSummary{
					WorkDays:        3,
					ExceptDays:      3,
					RegularWorkSec:  31,
					StandardWorkSec: 29040,
				},
				ExceptionInfos: []*CheckinException{
					{
						Count:     2,
						Duration:  28920,
						Exception: 4,
					},
					{
						Count:     1,
						Duration:  60,
						Exception: 1,
					},
					{
						Count:     1,
						Duration:  60,
						Exception: 2,
					},
				},
				SPItems: []*CheckinSPItem{
					{
						Count:      0,
						Duration:   0,
						TimeType:   0,
						Type:       100,
						VacationID: 0,
						Name:       "外勤次数",
					},
					{
						Count:      1,
						Duration:   0,
						TimeType:   0,
						Type:       1,
						VacationID: 2,
						Name:       "年假",
					},
				},
				OverworkInfo: &CheckinOverwork{
					WorkdayOverSec: 10800,
				},
			},
		},
	}, result)
}

func TestGetCheckinScheduleList(t *testing.T) {
	body := []byte(`{"starttime":1492617600,"endtime":1492790400,"useridlist":["james","paul"]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"schedule_list": [
		{
			"userid": "james",
			"yearmonth": 202011,
			"groupid": 11,
			"groupname": "排班",
			"schedule": {
				"scheduleList": [
					{
						"day": 25,
						"schedule_info": {
							"schedule_id": 229,
							"schedule_name": "早班",
							"time_section": [
								{
									"id": 1,
									"work_sec": 32400,
									"off_work_sec": 43200,
									"remind_work_sec": 32400,
									"remind_off_work_sec": 43200
								}
							]
						}
					},
					{
						"day": 26,
						"schedule_info": {
							"schedule_id": 171,
							"schedule_name": "晚班",
							"time_section": [
								{
									"id": 2,
									"work_sec": 64800,
									"off_work_sec": 79200,
									"remind_work_sec": 64800,
									"remind_off_work_sec": 79200
								}
							]
						}
					},
					{
						"day": 30,
						"schedule_info": {
							"schedule_id": 0,
							"schedule_name": "休息",
							"time_section": []
						}
					}
				]
			}
		}
	],
	"errcode": 0,
	"errmsg": "ok"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/checkin/getcheckinschedulist?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	userIDs := []string{"james", "paul"}

	result := new(ResultCheckinScheduleListGet)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetCheckinScheduleList(1492617600, 1492790400, userIDs, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultCheckinScheduleListGet{
		ScheduleList: []*CheckinSchedule{
			{
				UserID:    "james",
				YearMonth: 202011,
				GroupID:   11,
				GroupName: "排班",
				Schedule: &Schedule{
					ScheduleList: []*ScheduleData{
						{
							Day: 25,
							ScheduleInfo: &ScheduleInfo{
								ScheduleID:   229,
								ScheduleName: "早班",
								TimeSection: []*ScheduleTimeSection{
									{
										ID:               1,
										WorkSec:          32400,
										OffWorkSec:       43200,
										RemindWorkSec:    32400,
										RemindOffWorkSec: 43200,
									},
								},
							},
						},
						{
							Day: 26,
							ScheduleInfo: &ScheduleInfo{
								ScheduleID:   171,
								ScheduleName: "晚班",
								TimeSection: []*ScheduleTimeSection{
									{
										ID:               2,
										WorkSec:          64800,
										OffWorkSec:       79200,
										RemindWorkSec:    64800,
										RemindOffWorkSec: 79200,
									},
								},
							},
						},
						{
							Day: 30,
							ScheduleInfo: &ScheduleInfo{
								ScheduleID:   0,
								ScheduleName: "休息",
								TimeSection:  []*ScheduleTimeSection{},
							},
						},
					},
				},
			},
		},
	}, result)
}

func TestSetCheckinScheduleList(t *testing.T) {
	body := []byte(`{"groupid":226,"yearmonth":202012,"items":[{"userid":"james","day":5,"schedule_id":234}]}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/checkin/setcheckinschedulist?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	items := []*ScheduleListItem{
		{
			UserID:     "james",
			Day:        5,
			ScheduleID: 234,
		},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SetCheckinScheduleList(226, 202012, items...))

	assert.Nil(t, err)
}

func TestAddCheckinUserFace(t *testing.T) {
	body := []byte(`{"userid":"james","userface":"PLACE_HOLDER"}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/checkin/addcheckinuserface?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", AddCheckinUserFace("james", "PLACE_HOLDER"))

	assert.Nil(t, err)
}

func TestGetHardwareCheckinData(t *testing.T) {
	body := []byte(`{"filter_type":1,"starttime":1492617600,"endtime":1492790400,"useridlist":["james","paul"]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"checkindata": [
		{
			"userid": "james",
			"checkin_time": 1492617610,
			"device_sn": "xxxxx",
			"device_name": "xxxx门店"
		},
		{
			"userid": "paul",
			"checkin_time": 1492617620,
			"device_sn": "yyyy",
			"device_name": "yyyy门店"
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/hardware/get_hardware_checkin_data?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	userIDs := []string{"james", "paul"}

	result := new(ResultHardwareCheckinData)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetHardwareCheckinData(1, 1492617600, 1492790400, userIDs, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultHardwareCheckinData{
		CheckinData: []*HardwareCheckinData{
			{
				UserID:      "james",
				CheckinTime: 1492617610,
				DeviceSN:    "xxxxx",
				DeviceName:  "xxxx门店",
			},
			{
				UserID:      "paul",
				CheckinTime: 1492617620,
				DeviceSN:    "yyyy",
				DeviceName:  "yyyy门店",
			},
		},
	}, result)
}
