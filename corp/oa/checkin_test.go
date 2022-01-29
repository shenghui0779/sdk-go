package oa

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/corp"
	"github.com/shenghui0779/gochat/mock"
	"github.com/shenghui0779/gochat/wx"
	"github.com/stretchr/testify/assert"
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
				CheckinDate: &CheckinDate{
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
						Lng:       1133237500,
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
						"flex_off_duty_time": 0,
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
				"groupname": "打卡规则1",
				"need_photo": true,
				"wifimac_infos": [
					{
						"wifiname": "Tencent-WiFi-1",
						"wifimac": "c0:7b:bc:37:f8:d3",
					},
					{
						"wifiname": "Tencent-WiFi-2",
						"wifimac": "70:10:5c:7d:f6:d5",
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

	params := &ParamsCheckinOption{
		DateTime:   1511971200,
		UserIDList: []string{"james", "paul"},
	}
	result := new(ResultCheckinOption)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetCheckinOption(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultCheckinOption{
		Info: &CheckinInfo{
			UserID: "",
			Group: []*CheckinOption{
				{
					GroupType: 0,
					GroupID:   0,
					GroupName: "",
					CheckinDate: &CheckinDate{
						Workdays: []int{},
						CheckinTime: []*CheckinTime{
							{
								WorkSec:          0,
								OffWorkSec:       0,
								RemindWorkSec:    0,
								RemindOffWorkSec: 0,
							},
						},
						NoNeedOffWork:   false,
						LimitAheadTime:  0,
						FlexOnDutyTime:  0,
						FlexOffDutyTime: 0,
					},
					SpeWorkdays: []*CheckinSpeDay{
						{
							Timestamp: 0,
							Notes:     "",
							CheckinTime: []*CheckinTime{
								{
									WorkSec:          0,
									OffWorkSec:       0,
									RemindWorkSec:    0,
									RemindOffWorkSec: 0,
								},
							},
						},
					},
					SpeOffDays: []*CheckinSpeDay{
						{
							Timestamp: 0,
							Notes:     "",
							CheckinTime: []*CheckinTime{
								{
									WorkSec:          0,
									OffWorkSec:       0,
									RemindWorkSec:    0,
									RemindOffWorkSec: 0,
								},
							},
						},
					},
					SyncHolidays: false,
					NeedPhoto:    false,
					WifiMacInfos: []*CheckinWifiMac{
						{
							WifiName: "",
							WifiMac:  "",
						},
					},
					NoteCanUseLocalPic:     false,
					AllowCheckinOffWorkday: false,
					AllowApplyOffWorkday:   false,
					LocInfos: []*CheckinLoc{
						{
							Lat:       0,
							Lng:       0,
							LocTitle:  "",
							LocDetail: "",
							Distance:  0,
						},
					},
					ScheduleList: []*OptionSchedule{
						{
							ScheduleID:   0,
							ScheduleName: "",
							TimeSection: []*CheckinTimeSection{
								{
									TimeID:           0,
									WorkSec:          0,
									OffWorkSec:       0,
									RemindWorkSec:    0,
									RemindOffWorkSec: 0,
									RestBeginTime:    0,
									RestEndTime:      0,
									AllowRest:        false,
								},
							},
							LimitAheadTime:  0,
							NoNeedOffWork:   false,
							LimitOffTime:    0,
							FlexOnDutyTime:  0,
							FlexOffDutyTime: 0,
							AllowFlex:       false,
							LateRule: &CheckinLateRule{
								AllowOffWorkAfterTime: false,
								TimeRules: []*CheckinTimeRule{
									{
										OffWorkAfterTime: 0,
										OnWorkFlexTime:   0,
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
	body := []byte(``)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/user/authsucc?access_token=ACCESS_TOKEN&userid=USERID", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN")

	assert.Nil(t, err)
}

func TestGetCheckinDayData(t *testing.T) {
	body := []byte(``)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/user/authsucc?access_token=ACCESS_TOKEN&userid=USERID", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN")

	assert.Nil(t, err)
}

func TestGetCheckinMonthData(t *testing.T) {
	body := []byte(``)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/user/authsucc?access_token=ACCESS_TOKEN&userid=USERID", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN")

	assert.Nil(t, err)
}

func TestGetCheckinScheduleList(t *testing.T) {
	body := []byte(``)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/user/authsucc?access_token=ACCESS_TOKEN&userid=USERID", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN")

	assert.Nil(t, err)
}

func TestSetCheckinScheduleList(t *testing.T) {
	body := []byte(``)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/user/authsucc?access_token=ACCESS_TOKEN&userid=USERID", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN")

	assert.Nil(t, err)
}

func TestAddCheckinUserFace(t *testing.T) {
	body := []byte(``)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/user/authsucc?access_token=ACCESS_TOKEN&userid=USERID", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN")

	assert.Nil(t, err)
}
