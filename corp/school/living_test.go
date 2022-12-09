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

func TestGetUserAllLivingID(t *testing.T) {
	body := []byte(`{"userid":"USERID","cursor":"NEXT_KEY","limit":20}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"next_cursor": "next_cursor",
	"livingid_list": [
		"livingid1",
		"livingid2"
	]
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/living/get_user_all_livingid?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultUserAllLivingID)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetUserAllLivingID("USERID", "NEXT_KEY", 20, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultUserAllLivingID{
		NextCursor:   "next_cursor",
		LivingIDList: []string{"livingid1", "livingid2"},
	}, result)
}

func TestGetLivingInfo(t *testing.T) {
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"living_info": {
		"theme": "直角三角形讲解",
		"living_start": 1586405229,
		"living_duration": 1800,
		"anchor_userid": "zhangsan",
		"living_range": {
			"partyids": [
				1,
				4,
				9
			],
			"group_names": [
				"group_name1",
				"group_name2"
			]
		},
		"viewer_num": 100,
		"comment_num": 110,
		"open_replay": 1,
		"push_stream_url": "https://www.qq.test.com"
	}
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/school/living/get_living_info?access_token=ACCESS_TOKEN&livingid=LIVINGID", nil).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultLivingInfo)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetLivingInfo("LIVINGID", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultLivingInfo{
		LivingInfo: &LivingInfo{
			Theme:          "直角三角形讲解",
			LivingStart:    1586405229,
			LivingDuration: 1800,
			AnchorUserID:   "zhangsan",
			LivingRange: &LivingRange{
				PartyIDs:   []int64{1, 4, 9},
				GroupNames: []string{"group_name1", "group_name2"},
			},
			ViewerNum:     100,
			CommentNum:    110,
			OpenReplay:    1,
			PushStreamURL: "https://www.qq.test.com",
		},
	}, result)
}

func TestGetLivingWatchStat(t *testing.T) {
	body := []byte(`{"livingid":"livingid1","next_key":"NEXT_KEY"}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"ending": 1,
	"next_key": "NEXT_KEY",
	"stat_infoes": {
		"students": [
			{
				"student_userid": "zhansan_child",
				"parent_userid": "zhangsan",
				"partyids": [
					10,
					11
				],
				"watch_time": 30,
				"enter_time": 1586433904,
				"leave_time": 1586434000,
				"is_comment": 1
			},
			{
				"student_userid": "lisi_child",
				"parent_userid": "lisi",
				"partyids": [
					10,
					11
				],
				"watch_time": 30,
				"enter_time": 1586433904,
				"leave_time": 1586434000,
				"is_comment": 0
			}
		],
		"visitors": [
			{
				"nickname": "wx_nickname1",
				"watch_time": 30,
				"enter_time": 1586433904,
				"leave_time": 1586434000,
				"is_comment": 1
			},
			{
				"nickname": "wx_nickname2",
				"watch_time": 30,
				"enter_time": 1586433904,
				"leave_time": 1586434000,
				"is_comment": 0
			}
		]
	}
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/school/living/get_watch_stat?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultLivingWatchStat)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetLivingWatchStat("livingid1", "NEXT_KEY", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultLivingWatchStat{
		Ending:  1,
		NextKey: "NEXT_KEY",
		StatInfoes: &LivingWatchStatInfo{
			Students: []*LivingWatchStudent{
				{
					StudentUserID: "zhansan_child",
					ParentUserID:  "zhangsan",
					PartyIDs:      []int64{10, 11},
					WatchTime:     30,
					EnterTime:     1586433904,
					LeaveTime:     1586434000,
					IsComment:     1,
				},
				{
					StudentUserID: "lisi_child",
					ParentUserID:  "lisi",
					PartyIDs:      []int64{10, 11},
					WatchTime:     30,
					EnterTime:     1586433904,
					LeaveTime:     1586434000,
					IsComment:     0,
				},
			},
			Visitors: []*LivingVisitor{
				{
					Nickname:  "wx_nickname1",
					WatchTime: 30,
					EnterTime: 1586433904,
					LeaveTime: 1586434000,
					IsComment: 1,
				},
				{
					Nickname:  "wx_nickname2",
					WatchTime: 30,
					EnterTime: 1586433904,
					LeaveTime: 1586434000,
					IsComment: 0,
				},
			},
		},
	}, result)
}

func TestGetLivingUnwatchStat(t *testing.T) {
	body := []byte(`{"livingid":"livingid1","next_key":"NEXT_KEY"}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"ending": 1,
	"next_key": "NEXT_KEY",
	"stat_info": {
		"students": [
			{
				"student_userid": "zhansan_child",
				"parent_userid": "zhangsan",
				"partyids": [
					10,
					11
				]
			},
			{
				"student_userid": "lisi_child",
				"parent_userid": "lisi",
				"partyids": [
					5
				]
			}
		]
	}
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/school/living/get_unwatch_stat?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultLivingUnwatchStat)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetLivingUnwatchStat("livingid1", "NEXT_KEY", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultLivingUnwatchStat{
		Ending:  1,
		NextKey: "NEXT_KEY",
		StatInfo: &LivingUnwatchStatInfo{
			Students: []*LivingUnwatchStudent{
				{
					StudentUserID: "zhansan_child",
					ParentUserID:  "zhangsan",
					PartyIDs:      []int64{10, 11},
				},
				{
					StudentUserID: "lisi_child",
					ParentUserID:  "lisi",
					PartyIDs:      []int64{5},
				},
			},
		},
	}, result)
}

func TestDeleteLivingReplayData(t *testing.T) {
	body := []byte(`{"livingid":"XXXXXXXXX"}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/living/delete_replay_data?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", DeleteLivingReplayData("XXXXXXXXX"))

	assert.Nil(t, err)
}
