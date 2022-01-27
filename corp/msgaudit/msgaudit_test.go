package msgaudit

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

func TestListPermitUser(t *testing.T) {
	body := []byte(`{"type":1}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"ids": [
		"userid_111",
		"userid_222",
		"userid_333"
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/msgaudit/get_permit_user_list?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultPermitUserList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListPermitUser(1, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultPermitUserList{
		IDs: []string{"userid_111", "userid_222", "userid_333"},
	}, result)
}

func TestCheckSingleAgree(t *testing.T) {
	body := []byte(`{"info":[{"userid":"XuJinSheng","exteranalopenid":"wmeDKaCQAAGd9oGiQWxVsAKwV2HxNAAA"},{"userid":"XuJinSheng","exteranalopenid":"wmeDKaCQAAIQ_p7ACn_jpLVBJSGocAAA"},{"userid":"XuJinSheng","exteranalopenid":"wmeDKaCQAAPE_p7ABnxkpLBBJSGocAAA"}]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"agreeinfo": [
		{
			"status_change_time": 1562766651,
			"userid": "XuJinSheng",
			"exteranalopenid": "wmeDKaCPAAGdvxciQWxVsAKwV2HxNAAA",
			"agree_status": "Agree"
		},
		{
			"status_change_time": 1562766651,
			"userid": "XuJinSheng",
			"exteranalopenid": "wmeDKaCQAAIQ_p7ACnxksfeBJSGocAAA",
			"agree_status": "Disagree"
		},
		{
			"status_change_time": 1562766651,
			"userid": "XuJinSheng",
			"exteranalopenid": "wmeDKaCwAAIQ_p7ACnxckLBBJSGocAAA",
			"agree_status": "Agree"
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/msgaudit/check_single_agree?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsSingleAgreeCheck{
		Info: []*SingleCheckInfo{
			{
				UserID:         "XuJinSheng",
				ExternalOpenID: "wmeDKaCQAAGd9oGiQWxVsAKwV2HxNAAA",
			},
			{
				UserID:         "XuJinSheng",
				ExternalOpenID: "wmeDKaCQAAIQ_p7ACn_jpLVBJSGocAAA",
			},
			{
				UserID:         "XuJinSheng",
				ExternalOpenID: "wmeDKaCQAAPE_p7ABnxkpLBBJSGocAAA",
			},
		},
	}

	result := new(ResultSingleAgreeCheck)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", CheckSingleAgree(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultSingleAgreeCheck{
		AgreeInfo: []*SingleAgreeInfo{
			{
				StatusChangeTime: 1562766651,
				UserID:           "XuJinSheng",
				ExternalOpenID:   "wmeDKaCPAAGdvxciQWxVsAKwV2HxNAAA",
				AgreeStatus:      "Agree",
			},
			{
				StatusChangeTime: 1562766651,
				UserID:           "XuJinSheng",
				ExternalOpenID:   "wmeDKaCQAAIQ_p7ACnxksfeBJSGocAAA",
				AgreeStatus:      "Disagree",
			},
			{
				StatusChangeTime: 1562766651,
				UserID:           "XuJinSheng",
				ExternalOpenID:   "wmeDKaCwAAIQ_p7ACnxckLBBJSGocAAA",
				AgreeStatus:      "Agree",
			},
		},
	}, result)
}

func TestCheckRoomAgree(t *testing.T) {
	body := []byte(`{"roomid":"wrjc7bDwAASxc8tZvBErFE02BtPWyAAA"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"agreeinfo": [
		{
			"status_change_time": 1562766651,
			"exteranalopenid": "wmeDKaCQAAGdtHdiQWxVadfwV2HxNAAA",
			"agree_status": "Agree"
		},
		{
			"status_change_time": 1562766651,
			"exteranalopenid": "wmeDKaCQAAIQ_p9ACyiopLBBJSGocAAA",
			"agree_status": "Disagree"
		},
		{
			"status_change_time": 1562766651,
			"exteranalopenid": "wmeDKaCQAAIQ_p9ACnxacyBBJSGocAAA",
			"agree_status": "Agree"
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/msgaudit/check_room_agree?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultRoomAgreeCheck)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", CheckRoomAgree("wrjc7bDwAASxc8tZvBErFE02BtPWyAAA", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultRoomAgreeCheck{
		AgreeInfo: []*RoomAgreeInfo{
			{
				StatusChangeTime: 1562766651,
				ExternalOpenID:   "wmeDKaCQAAGdtHdiQWxVadfwV2HxNAAA",
				AgreeStatus:      "Agree",
			},
			{
				StatusChangeTime: 1562766651,
				ExternalOpenID:   "wmeDKaCQAAIQ_p9ACyiopLBBJSGocAAA",
				AgreeStatus:      "Disagree",
			},
			{
				StatusChangeTime: 1562766651,
				ExternalOpenID:   "wmeDKaCQAAIQ_p9ACnxacyBBJSGocAAA",
				AgreeStatus:      "Agree",
			},
		},
	}, result)
}

func TestGetGroupChat(t *testing.T) {
	body := []byte(`{"roomid":"wrNplhCgAAIVZohLe57zKnvIV7xBKrig"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"roomname": "蓦然回首",
	"creator": "ZhangWenChao",
	"room_create_time": 1592361604,
	"notice": "",
	"members": [
		{
			"memberid": "ZhangWenChao",
			"jointime": 1592361605
		},
		{
			"memberid": "xujinsheng",
			"jointime": 1592377076
		}
	],
	"errcode": 0,
	"errmsg": "ok"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/msgaudit/groupchat/get?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultGroupChatGet)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetGroupChat("wrNplhCgAAIVZohLe57zKnvIV7xBKrig", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultGroupChatGet{
		RoomName:       "蓦然回首",
		Creator:        "ZhangWenChao",
		RoomCreateTime: 1592361604,
		Notice:         "",
		Members: []*GroupMember{
			{
				MemberID: "ZhangWenChao",
				JoinTime: 1592361605,
			},
			{
				MemberID: "xujinsheng",
				JoinTime: 1592377076,
			},
		},
	}, result)
}
