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

func TestTransferCustomer(t *testing.T) {
	body := []byte(`{"handover_userid":"zhangsan","takeover_userid":"lisi","external_userid":["woAJ2GCAAAXtWyujaWJHDDGi0mACAAAA","woAJ2GCAAAXtWyujaWJHDDGi0mACBBBB"],"transfer_success_msg":"您好，您的服务已升级，后续将由我的同事李四@腾讯接替我的工作，继续为您服务。"}`)
	resp := []byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "customer": [
        {
            "external_userid": "woAJ2GCAAAXtWyujaWJHDDGi0mACAAAA",
            "errcode": 40096
        },
        {
            "external_userid": "woAJ2GCAAAXtWyujaWJHDDGi0mACBBBB",
            "errcode": 0
        }
    ]
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/transfer_customer?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	params := &ParamsCustomerTranster{
		HandoverUserID:     "zhangsan",
		TakeoverUserID:     "lisi",
		ExternalUserID:     []string{"woAJ2GCAAAXtWyujaWJHDDGi0mACAAAA", "woAJ2GCAAAXtWyujaWJHDDGi0mACBBBB"},
		TransferSuccessMsg: "您好，您的服务已升级，后续将由我的同事李四@腾讯接替我的工作，继续为您服务。",
	}

	result := new(ResultCustomerTransfer)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", TransferCustomer(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultCustomerTransfer{
		Customer: []*ErrCustomerTransfer{
			{
				ExternalUserID: "woAJ2GCAAAXtWyujaWJHDDGi0mACAAAA",
				ErrCode:        40096,
			},
			{
				ExternalUserID: "woAJ2GCAAAXtWyujaWJHDDGi0mACBBBB",
				ErrCode:        0,
			},
		},
	}, result)
}

func TestGetTransferResult(t *testing.T) {
	body := []byte(`{"handover_userid":"zhangsan","takeover_userid":"lisi","cursor":"CURSOR"}`)
	resp := []byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "customer": [
        {
            "external_userid": "woAJ2GCAAAXtWyujaWJHDDGi0mACCCC",
            "status": 1,
            "takeover_time": 1588262400
        },
        {
            "external_userid": "woAJ2GCAAAXtWyujaWJHDDGi0mACBBBB",
            "status": 2,
            "takeover_time": 1588482400
        },
        {
            "external_userid": "woAJ2GCAAAXtWyujaWJHDDGi0mACAAAA",
            "status": 3,
            "takeover_time": 0
        }
    ],
    "next_cursor": "NEXT_CURSOR"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/transfer_result?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultTransferRet)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetTransferResult("zhangsan", "lisi", "CURSOR", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultTransferRet{
		Customer: []*TransferRet{
			{
				ExternalUserID: "woAJ2GCAAAXtWyujaWJHDDGi0mACCCC",
				Status:         1,
				TakeoverTime:   1588262400,
			},
			{
				ExternalUserID: "woAJ2GCAAAXtWyujaWJHDDGi0mACBBBB",
				Status:         2,
				TakeoverTime:   1588482400,
			},
			{
				ExternalUserID: "woAJ2GCAAAXtWyujaWJHDDGi0mACAAAA",
				Status:         3,
				TakeoverTime:   0,
			},
		},
		NextCursor: "NEXT_CURSOR",
	}, result)
}

func TestListUnassigned(t *testing.T) {
	body := []byte(`{"page_size":100}`)
	resp := []byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "info": [
        {
            "handover_userid": "zhangsan",
            "external_userid": "woAJ2GCAAAd4uL12hdfsdasassdDmAAAAA",
            "dimission_time": 1550838571
        },
        {
            "handover_userid": "lisi",
            "external_userid": "wmAJ2GCAAAzLTI123ghsdfoGZNqqAAAA",
            "dimission_time": 1550661468
        }
    ],
    "is_last": false,
    "next_cursor": "aSfwejksvhToiMMfFeIGZZ"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_unassigned_list?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultUnassignedList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListUnassigned(0, 100, "", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultUnassignedList{
		Info: []*UnassignedInfo{
			{
				HandoverUserID: "zhangsan",
				ExternalUserID: "woAJ2GCAAAd4uL12hdfsdasassdDmAAAAA",
				DimissionTime:  1550838571,
			},
			{
				HandoverUserID: "lisi",
				ExternalUserID: "wmAJ2GCAAAzLTI123ghsdfoGZNqqAAAA",
				DimissionTime:  1550661468,
			},
		},
		IsLast:     false,
		NextCursor: "aSfwejksvhToiMMfFeIGZZ",
	}, result)
}

func TestTransferResignedCustomer(t *testing.T) {
	body := []byte(`{"handover_userid":"zhangsan","takeover_userid":"lisi","external_userid":["woAJ2GCAAAXtWyujaWJHDDGi0mACBBBB","woAJ2GCAAAXtWyujaWJHDDGi0mACAAAA"]}`)
	resp := []byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "customer": [
        {
            "external_userid": "woAJ2GCAAAXtWyujaWJHDDGi0mACBBBB",
            "errcode": 0
        },
        {
            "external_userid": "woAJ2GCAAAXtWyujaWJHDDGi0mACAAAA",
            "errcode": 40096
        }
    ]
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/resigned/transfer_customer?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	externalUserIDs := []string{"woAJ2GCAAAXtWyujaWJHDDGi0mACBBBB", "woAJ2GCAAAXtWyujaWJHDDGi0mACAAAA"}

	result := new(ResultResignedCustomerTransfer)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", TransferResignedCustomer("zhangsan", "lisi", externalUserIDs, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultResignedCustomerTransfer{
		Customer: []*ErrCustomerTransfer{
			{
				ExternalUserID: "woAJ2GCAAAXtWyujaWJHDDGi0mACBBBB",
				ErrCode:        0,
			},
			{
				ExternalUserID: "woAJ2GCAAAXtWyujaWJHDDGi0mACAAAA",
				ErrCode:        40096,
			},
		},
	}, result)
}

func TestGetResignedTransferResult(t *testing.T) {
	body := []byte(`{"handover_userid":"zhangsan","takeover_userid":"lisi","cursor":"CURSOR"}`)
	resp := []byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "customer": [
        {
            "external_userid": "woAJ2GCAAAXtWyujaWJHDDGi0mACCCC",
            "status": 1,
            "takeover_time": 1588262400
        },
        {
            "external_userid": "woAJ2GCAAAXtWyujaWJHDDGi0mACBBBB",
            "status": 2,
            "takeover_time": 1588482400
        },
        {
            "external_userid": "woAJ2GCAAAXtWyujaWJHDDGi0mACAAAA",
            "status": 3,
            "takeover_time": 0
        }
    ],
    "next_cursor": "NEXT_CURSOR"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/resigned/transfer_result?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultResignedTransferRet)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetResignedTransferResult("zhangsan", "lisi", "CURSOR", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultResignedTransferRet{
		Customer: []*ResignedTransferRet{
			{
				ExternalUserID: "woAJ2GCAAAXtWyujaWJHDDGi0mACCCC",
				Status:         1,
				TakeoverTime:   1588262400,
			},
			{
				ExternalUserID: "woAJ2GCAAAXtWyujaWJHDDGi0mACBBBB",
				Status:         2,
				TakeoverTime:   1588482400,
			},
			{
				ExternalUserID: "woAJ2GCAAAXtWyujaWJHDDGi0mACAAAA",
				Status:         3,
				TakeoverTime:   0,
			},
		},
		NextCursor: "NEXT_CURSOR",
	}, result)
}

func TestTransferGroupChat(t *testing.T) {
	body := []byte(`{"chat_id_list":["wrOgQhDgAAcwMTB7YmDkbeBsgT_AAAA","wrOgQhDgAAMYQiS5ol9G7gK9JVQUAAAA"],"new_owner":"zhangsan"}`)
	resp := []byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "failed_chat_list": [
        {
            "chat_id": "wrOgQhDgAAcwMTB7YmDkbeBsgT_KAAAA",
            "errcode": 90500,
            "errmsg": "the owner of this chat is not resigned"
        }
    ]
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/transfer?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultGroupChatTransfer)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", TransferGroupChat([]string{"wrOgQhDgAAcwMTB7YmDkbeBsgT_AAAA", "wrOgQhDgAAMYQiS5ol9G7gK9JVQUAAAA"}, "zhangsan", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultGroupChatTransfer{
		FailedChatList: []*ErrGroupChatTransfer{
			{
				ChatID:  "wrOgQhDgAAcwMTB7YmDkbeBsgT_KAAAA",
				ErrCode: 90500,
				ErrMsg:  "the owner of this chat is not resigned",
			},
		},
	}, result)
}
