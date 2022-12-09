package externalcontact

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/corp"
	"github.com/shenghui0779/gochat/mock"
)

func TestListGroupChat(t *testing.T) {
	body := []byte(`{"owner_filter":{"userid_list":["abel"]},"cursor":"r9FqSqsI8fgNbHLHE5QoCP50UIg2cFQbfma3l2QsmwI","limit":10}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"group_chat_list": [
		{
			"chat_id": "wrOgQhDgAAMYQiS5ol9G7gK9JVAAAA",
			"status": 0
		},
		{
			"chat_id": "wrOgQhDgAAcwMTB7YmDkbeBsAAAA",
			"status": 0
		}
	],
	"next_cursor": "tJzlB9tdqfh-g7i_J-ehOz_TWcd7dSKa39_AqCIeMFw"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/list?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	params := &ParamsGroupChatList{
		OwnerFilter: &GroupChatOwnerFilter{
			UserIDList: []string{"abel"},
		},
		Cursor: "r9FqSqsI8fgNbHLHE5QoCP50UIg2cFQbfma3l2QsmwI",
		Limit:  10,
	}

	result := new(ResultGroupChatList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListGroupChat(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultGroupChatList{
		GroupChatList: []*GroupChatListData{
			{
				ChatID: "wrOgQhDgAAMYQiS5ol9G7gK9JVAAAA",
				Status: 0,
			},
			{
				ChatID: "wrOgQhDgAAcwMTB7YmDkbeBsAAAA",
				Status: 0,
			},
		},
		NextCursor: "tJzlB9tdqfh-g7i_J-ehOz_TWcd7dSKa39_AqCIeMFw",
	}, result)
}

func TestGetGroupChat(t *testing.T) {
	body := []byte(`{"chat_id":"wrOgQhDgAAMYQiS5ol9G7gK9JVAAAA","need_name":1}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"group_chat": {
		"chat_id": "wrOgQhDgAAMYQiS5ol9G7gK9JVAAAA",
		"name": "销售客服群",
		"owner": "ZhuShengBen",
		"create_time": 1572505490,
		"notice": "文明沟通，拒绝脏话",
		"member_list": [
			{
				"userid": "abel",
				"type": 1,
				"join_time": 1572505491,
				"join_scene": 1,
				"invitor": {
					"userid": "jack"
				},
				"group_nickname": "客服小张",
				"name": "张三丰"
			},
			{
				"userid": "wmOgQhDgAAuXFJGwbve4g4iXknfOAAAA",
				"type": 2,
				"unionid": "ozynqsulJFCZ2z1aYeS8h-nuasdAAA",
				"join_time": 1572505491,
				"join_scene": 1,
				"group_nickname": "顾客老王",
				"name": "王语嫣"
			}
		],
		"admin_list": [
			{
				"userid": "sam"
			},
			{
				"userid": "pony"
			}
		]
	}
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/get?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultGroupChatGet)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetGroupChat("wrOgQhDgAAMYQiS5ol9G7gK9JVAAAA", 1, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultGroupChatGet{
		GroupChat: &GroupChat{
			ChatID:     "wrOgQhDgAAMYQiS5ol9G7gK9JVAAAA",
			Name:       "销售客服群",
			Owner:      "ZhuShengBen",
			CreateTime: 1572505490,
			Notice:     "文明沟通，拒绝脏话",
			MemberList: []*GroupChatMember{
				{
					UserID:    "abel",
					Type:      1,
					JoinTime:  1572505491,
					JoinScene: 1,
					Invitor: &GroupChatInvitor{
						UserID: "jack",
					},
					GroupNickname: "客服小张",
					Name:          "张三丰",
				},
				{
					UserID:        "wmOgQhDgAAuXFJGwbve4g4iXknfOAAAA",
					Type:          2,
					Unionid:       "ozynqsulJFCZ2z1aYeS8h-nuasdAAA",
					JoinTime:      1572505491,
					JoinScene:     1,
					GroupNickname: "顾客老王",
					Name:          "王语嫣",
				},
			},
			AdminList: []*GroupChatAdmin{
				{
					UserID: "sam",
				},
				{
					UserID: "pony",
				},
			},
		},
	}, result)
}

func TestOpenGIDToChatID(t *testing.T) {
	body := []byte(`{"opengid":"oAAAAAAA"}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"chat_id": "ooAAAAAAAAAAA"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/opengid_to_chatid?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultOpenGIDToChatID)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", OpenGIDToChatID("oAAAAAAA", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultOpenGIDToChatID{
		ChatID: "ooAAAAAAAAAAA",
	}, result)
}
