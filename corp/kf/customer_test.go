package kf

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/corp"
	"github.com/shenghui0779/gochat/mock"
)

func TestBatchGetCustomer(t *testing.T) {
	body := []byte(`{"external_userid_list":["wmxxxxxxxxxxxxxxxxxxxxxx","zhangsan"]}`)
	resp := []byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "customer_list": [
        {
            "external_userid": "wmxxxxxxxxxxxxxxxxxxxxxx",
            "nickname": "张三",
            "avatar": "http://xxxxx",
            "gender": 1,
            "unionid": "oxasdaosaosdasdasdasd"
        }
    ],
    "invalid_external_userid": [
        "zhangsan"
    ]
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/kf/customer/batchget?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	externalUserIDs := []string{"wmxxxxxxxxxxxxxxxxxxxxxx", "zhangsan"}

	result := new(ResultCustomerBatchGet)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", BatchGetCustomer(externalUserIDs, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultCustomerBatchGet{
		CustomerList: []*Customer{
			{
				ExternalUserID: "wmxxxxxxxxxxxxxxxxxxxxxx",
				Nickname:       "张三",
				Avatar:         "http://xxxxx",
				Gender:         1,
				UnionID:        "oxasdaosaosdasdasdasd",
			},
		},
		InvalidExternalUserID: []string{"zhangsan"},
	}, result)
}

func TestGetUpgradeServiceConfig(t *testing.T) {
	resp := []byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "member_range": {
        "userid_list": [
            "zhangsan",
            "lisi"
        ],
        "department_id_list": [
            2,
            3
        ]
    },
    "groupchat_range": {
        "chat_id_list": [
            "wraaaaaaaaaaaaaaaa",
            "wrbbbbbbbbbbbbbbb"
        ]
    }
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/kf/customer/get_upgrade_service_config?access_token=ACCESS_TOKEN", nil).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultServiceUpgradeConfig)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetUpgradeServiceConfig(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultServiceUpgradeConfig{
		MemberRange: &MemberRange{
			UserIDList:       []string{"zhangsan", "lisi"},
			DepartmentIDList: []int64{2, 3},
		},
		GroupChatRange: &GroupChatRange{
			ChatIDList: []string{"wraaaaaaaaaaaaaaaa", "wrbbbbbbbbbbbbbbb"},
		},
	}, result)
}

func TestUpgradeMemberService(t *testing.T) {
	body := []byte(`{"open_kfid":"kfxxxxxxxxxxxxxx","external_userid":"wmxxxxxxxxxxxxxxxxxx","type":1,"member":{"userid":"zhangsan","wording":"你好，我是你的专属服务专员zhangsan"}}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/kf/customer/upgrade_service?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	member := &Member{
		UserID:  "zhangsan",
		Wording: "你好，我是你的专属服务专员zhangsan",
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UpgradeMemberService("kfxxxxxxxxxxxxxx", "wmxxxxxxxxxxxxxxxxxx", member))

	assert.Nil(t, err)
}

func TestUpgradeGroupChatService(t *testing.T) {
	body := []byte(`{"open_kfid":"kfxxxxxxxxxxxxxx","external_userid":"wmxxxxxxxxxxxxxxxxxx","type":2,"groupchat":{"chat_id":"wraaaaaaaaaaaaaaaa","wording":"欢迎加入你的专属服务群"}}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/kf/customer/upgrade_service?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	groupChat := &GroupChat{
		ChatID:  "wraaaaaaaaaaaaaaaa",
		Wording: "欢迎加入你的专属服务群",
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UpgradeGroupChatService("kfxxxxxxxxxxxxxx", "wmxxxxxxxxxxxxxxxxxx", groupChat))

	assert.Nil(t, err)
}

func TestCancelUpgradeService(t *testing.T) {
	body := []byte(`{"open_kfid":"kfxxxxxxxxxxxxxx","external_userid":"wmxxxxxxxxxxxxxxxxxx"}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/kf/customer/cancel_upgrade_service?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", CancelUpgradeService("kfxxxxxxxxxxxxxx", "wmxxxxxxxxxxxxxxxxxx"))

	assert.Nil(t, err)
}
