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

func TestGetUserBehaviorData(t *testing.T) {
	body := []byte(`{"userid":["zhangsan","lisi"],"partyid":[1001,1002],"start_time":1536508800,"end_time":1536595200}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "behavior_data": [
        {
            "stat_time": 1536508800,
            "chat_cnt": 100,
            "message_cnt": 80,
            "reply_percentage": 60.25,
            "avg_reply_time": 1,
            "negative_feedback_cnt": 0,
            "new_apply_cnt": 6,
            "new_contact_cnt": 5
        },
        {
            "stat_time": 1536595200,
            "chat_cnt": 20,
            "message_cnt": 40,
            "reply_percentage": 100,
            "avg_reply_time": 1,
            "negative_feedback_cnt": 0,
            "new_apply_cnt": 6,
            "new_contact_cnt": 5
        }
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_user_behavior_data?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsUserBehaviorData{
		UserID:    []string{"zhangsan", "lisi"},
		PartyID:   []int64{1001, 1002},
		StartTime: 1536508800,
		EndTime:   1536595200,
	}

	result := new(ResultUserBehaviorData)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetUserBehaviorData(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultUserBehaviorData{
		BehaviorData: []*UserBehaviorData{
			{
				StatTime:            1536508800,
				ChatCnt:             100,
				MessageCnt:          80,
				ReplyPercentage:     60.25,
				AvgReplyTime:        1,
				NegativeFeedbackCnt: 0,
				NewApplyCnt:         6,
				NewContactCnt:       5,
			},
			{
				StatTime:            1536595200,
				ChatCnt:             20,
				MessageCnt:          40,
				ReplyPercentage:     100,
				AvgReplyTime:        1,
				NegativeFeedbackCnt: 0,
				NewApplyCnt:         6,
				NewContactCnt:       5,
			},
		},
	}, result)
}

func TestGetGroupChatStatistic(t *testing.T) {
	body := []byte(`{"day_begin_time":1600272000,"day_end_time":1600444800,"owner_filter":{"userid_list":["zhangsan"]},"order_by":2,"limit":1000}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "total": 2,
    "next_offset": 2,
    "items": [
        {
            "owner": "zhangsan",
            "data": {
                "new_chat_cnt": 2,
                "chat_total": 2,
                "chat_has_msg": 0,
                "new_member_cnt": 0,
                "member_total": 6,
                "member_has_msg": 0,
                "msg_total": 0,
                "migrate_trainee_chat_cnt": 3
            }
        },
        {
            "owner": "lisi",
            "data": {
                "new_chat_cnt": 1,
                "chat_total": 3,
                "chat_has_msg": 2,
                "new_member_cnt": 0,
                "member_total": 6,
                "member_has_msg": 0,
                "msg_total": 0,
                "migrate_trainee_chat_cnt": 3
            }
        }
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/statistic?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsGroupChatStatistic{
		DayBeginTime: 1600272000,
		DayEndTime:   1600444800,
		OwnerFilter: &GroupChatOwnerFilter{
			UserIDList: []string{"zhangsan"},
		},
		OrderBy: 2,
		Limit:   1000,
	}

	result := new(ResultGroupChatStatistic)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetGroupChatStatistic(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultGroupChatStatistic{
		Total:      2,
		NextOffset: 2,
		Items: []*GroupChatStatisticItem{
			{
				Owner: "zhangsan",
				Data: &GroupChatStatisticData{
					NewChatCnt:            2,
					ChatTotal:             2,
					MemberTotal:           6,
					MigrateTraineeChatCnt: 3,
				},
			},
			{
				Owner: "lisi",
				Data: &GroupChatStatisticData{
					NewChatCnt:            1,
					ChatTotal:             3,
					ChatHasMsg:            2,
					MemberTotal:           6,
					MigrateTraineeChatCnt: 3,
				},
			},
		},
	}, result)
}

func TestGetGroupChatStatisticByDay(t *testing.T) {
	body := []byte(`{"day_begin_time":1600272000,"day_end_time":1600358400,"owner_filter":{"userid_list":["zhangsan"]}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "items": [
        {
            "stat_time": 1600272000,
            "data": {
                "new_chat_cnt": 2,
                "chat_total": 2,
                "chat_has_msg": 0,
                "new_member_cnt": 0,
                "member_total": 6,
                "member_has_msg": 0,
                "msg_total": 0,
                "migrate_trainee_chat_cnt": 3
            }
        },
        {
            "stat_time": 1600358400,
            "data": {
                "new_chat_cnt": 2,
                "chat_total": 2,
                "chat_has_msg": 0,
                "new_member_cnt": 0,
                "member_total": 6,
                "member_has_msg": 0,
                "msg_total": 0,
                "migrate_trainee_chat_cnt": 3
            }
        }
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/statistic_group_by_day?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsGroupChatStatisticByDay{
		DayBeginTime: 1600272000,
		DayEndTime:   1600358400,
		OwnerFilter: &GroupChatOwnerFilter{
			UserIDList: []string{"zhangsan"},
		},
	}

	result := new(ResultGroupChatStatisticByDay)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetGroupChatStatisticByDay(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultGroupChatStatisticByDay{
		Items: []*GroupChatStatisticByDayItem{
			{
				StatTime: 1600272000,
				Data: &GroupChatStatisticData{
					NewChatCnt:            2,
					ChatTotal:             2,
					MemberTotal:           6,
					MigrateTraineeChatCnt: 3,
				},
			},
			{
				StatTime: 1600358400,
				Data: &GroupChatStatisticData{
					NewChatCnt:            2,
					ChatTotal:             2,
					MemberTotal:           6,
					MigrateTraineeChatCnt: 3,
				},
			},
		},
	}, result)
}
