package tools

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

func TestCreateLiving(t *testing.T) {
	body := []byte(`{"anchor_userid":"zhangsan","theme":"theme","living_start":1600000000,"living_duration":3600,"description":"test description","type":4,"agentid":1000014,"remind_time":60,"activity_cover_mediaid":"MEDIA_ID","activity_share_mediaid":"MEDIA_ID","activity_detail":{"description":"活动描述，非活动类型的直播不用传","image_list":["MEDIA_ID_1","MEDIA_ID_2"]}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "livingid": "XXXXXXXXX"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/living/create?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsLivingCreate{
		AnchorUserID:         "zhangsan",
		Theme:                "theme",
		LivingStart:          1600000000,
		LivingDuration:       3600,
		Description:          "test description",
		Type:                 4,
		AgentID:              1000014,
		RemindTime:           60,
		ActivityCoverMediaID: "MEDIA_ID",
		ActivityShareMediaID: "MEDIA_ID",
		ActivityDetail: &LivingActivity{
			Description: "活动描述，非活动类型的直播不用传",
			ImageList:   []string{"MEDIA_ID_1", "MEDIA_ID_2"},
		},
	}
	result := new(ResultLivingCreate)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", CreateLiving(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultLivingCreate{
		LivingID: "XXXXXXXXX",
	}, result)
}

func TestModifyLiving(t *testing.T) {
	body := []byte(`{"livingid":"XXXXXXXXX","theme":"theme","living_start":1600100000,"living_duration":3600,"description":"test description","type":1,"remind_time":60}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/living/modify?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsLivingModify{
		LivingID:       "XXXXXXXXX",
		Theme:          "theme",
		LivingStart:    1600100000,
		LivingDuration: 3600,
		Description:    "test description",
		Type:           1,
		RemindTime:     60,
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ModifyLiving(params))

	assert.Nil(t, err)
}

func TestCancelLiving(t *testing.T) {
	body := []byte(`{"livingid":"XXXXXXXXX"}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/living/cancel?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", CancelLiving("XXXXXXXXX"))

	assert.Nil(t, err)
}

func TestDeleteLivingReplayData(t *testing.T) {
	body := []byte(`{"livingid":"XXXXXXXXX"}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/living/delete_replay_data?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", DeleteLivingReplayData("XXXXXXXXX"))

	assert.Nil(t, err)
}

func TestGetLivingCode(t *testing.T) {
	body := []byte(`{"livingid":"XXXXXXXXX","openid":"abcopenid"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "living_code": "abcdef"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/living/get_living_code?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultLivingCode)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetLivingCode("XXXXXXXXX", "abcopenid", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultLivingCode{
		LivingCode: "abcdef",
	}, result)
}

func TestGetUserAllLivingID(t *testing.T) {
	body := []byte(`{"userid":"USERID","cursor":"NEXT_KEY","limit":20}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "next_cursor": "next_cursor",
    "livingid_list": [
        "livingid1",
        "livingid2"
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/living/get_user_all_livingid?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultUserAllLivingID)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetUserAllLivingID("USERID", "NEXT_KEY", 20, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultUserAllLivingID{
		NextCursor:   "next_cursor",
		LivingIDList: []string{"livingid1", "livingid2"},
	}, result)
}

func TestGetLivingInfo(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "living_info": {
        "theme": "直角三角形讲解",
        "living_start": 1586405229,
        "living_duration": 1800,
        "status": 3,
        "reserve_start": 1586405239,
        "reserve_living_duration": 1600,
        "description": "小学数学精选课程",
        "anchor_userid": "zhangsan",
        "main_department": 1,
        "viewer_num": 100,
        "comment_num": 110,
        "mic_num": 120,
        "open_replay": 1,
        "replay_status": 2,
        "type": 0,
        "push_stream_url": "https://www.qq.test.com",
        "online_count": 1,
        "subscribe_count": 1
    }
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/living/get_living_info?access_token=ACCESS_TOKEN&livingid=LIVINGID", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultLivingInfo)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetLivingInfo("LIVINGID", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultLivingInfo{
		LivingInfo: &LivingInfo{
			Theme:                 "直角三角形讲解",
			LivingStart:           1586405229,
			LivingDuration:        1800,
			Status:                3,
			ReserveStart:          1586405239,
			ReserveLivingDuration: 1600,
			Description:           "小学数学精选课程",
			AnchorUserID:          "zhangsan",
			MainDepartment:        1,
			ViewerNum:             100,
			CommentNum:            110,
			MicNum:                120,
			OpenReplay:            1,
			ReplayStatus:          2,
			Type:                  0,
			PushStreamURL:         "https://www.qq.test.com",
			OnlineCount:           1,
			SubscribeCount:        1,
		},
	}, result)
}

func TestGetLivingWatchStat(t *testing.T) {
	body := []byte(`{"livingid":"livingid1","next_key":"NEXT_KEY"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "ending": 1,
    "next_key": "NEXT_KEY",
    "stat_info": {
        "users": [
            {
                "userid": "userid",
                "watch_time": 30,
                "is_comment": 1,
                "is_mic": 1
            }
        ],
        "external_users": [
            {
                "external_userid": "external_userid1",
                "type": 1,
                "name": "user name",
                "watch_time": 30,
                "is_comment": 1,
                "is_mic": 1
            },
            {
                "external_userid": "external_userid2",
                "type": 2,
                "name": "user_name",
                "watch_time": 30,
                "is_comment": 1,
                "is_mic": 1
            }
        ]
    }
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/living/get_watch_stat?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultLivingWatchStat)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetLivingWatchStat("livingid1", "NEXT_KEY", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultLivingWatchStat{
		Ending:  1,
		NextKey: "NEXT_KEY",
		StatInfo: &LivingStatInfo{
			Users: []*LivingUser{
				{
					UserID:    "userid",
					WatchTime: 30,
					IsComment: 1,
					IsMic:     1,
				},
			},
			ExternalUsers: []*LivingExternalUser{
				{
					ExternalUserID: "external_userid1",
					Type:           1,
					Name:           "user name",
					WatchTime:      30,
					IsComment:      1,
					IsMic:          1,
				},
				{
					ExternalUserID: "external_userid2",
					Type:           2,
					Name:           "user_name",
					WatchTime:      30,
					IsComment:      1,
					IsMic:          1,
				},
			},
		},
	}, result)
}

func TestGetLivingShareInfo(t *testing.T) {
	body := []byte(`{"ww_share_code":"CODE"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "livingid": "livingid",
    "viewer_userid": "viewer_userid",
    "viewer_external_userid": "viewer_external_userid",
    "invitor_userid": "invitor_userid",
    "invitor_external_userid": "invitor_external_userid"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/living/get_living_share_info?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultLivingShareInfo)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetLivingShareInfo("CODE", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultLivingShareInfo{
		LivingID:              "livingid",
		ViewerUserID:          "viewer_userid",
		ViewerExternalUserID:  "viewer_external_userid",
		InvitorUserID:         "invitor_userid",
		InvitorExternalUserID: "invitor_external_userid",
	}, result)
}
