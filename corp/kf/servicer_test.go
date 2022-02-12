package kf

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

func TestAddServicer(t *testing.T) {
	body := []byte(`{"open_kfid":"kfxxxxxxxxxxxxxx","userid_list":["zhangsan","lisi"]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "result_list": [
        {
            "userid": "zhangsan",
            "errcode": 0,
            "errmsg": "success"
        },
        {
            "userid": "lisi",
            "errcode": 0,
            "errmsg": "ignored"
        }
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/kf/servicer/add?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	userIDs := []string{"zhangsan", "lisi"}

	result := new(ResultServicerAdd)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", AddServicer("kfxxxxxxxxxxxxxx", userIDs, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultServicerAdd{
		ResultList: []*ErrServicer{
			{
				UserID:  "zhangsan",
				ErrCode: 0,
				ErrMsg:  "success",
			},
			{
				UserID:  "lisi",
				ErrCode: 0,
				ErrMsg:  "ignored",
			},
		},
	}, result)
}

func TestDeleteServicer(t *testing.T) {
	body := []byte(`{"open_kfid":"kfxxxxxxxxxxxxxx","userid_list":["zhangsan","lisi"]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "result_list": [
        {
            "userid": "zhangsan",
            "errcode": 0,
            "errmsg": "success"
        },
        {
            "userid": "lisi",
            "errcode": 0,
            "errmsg": "ignored"
        }
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/kf/servicer/del?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	userIDs := []string{"zhangsan", "lisi"}

	result := new(ResultServicerDelete)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", DeleteServicer("kfxxxxxxxxxxxxxx", userIDs, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultServicerDelete{
		ResultList: []*ErrServicer{
			{
				UserID:  "zhangsan",
				ErrCode: 0,
				ErrMsg:  "success",
			},
			{
				UserID:  "lisi",
				ErrCode: 0,
				ErrMsg:  "ignored",
			},
		},
	}, result)
}

func TestListServicer(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "servicer_list": [
        {
            "userid": "zhangsan",
            "status": 0
        },
        {
            "userid": "lisi",
            "status": 1
        }
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/kf/servicer/list?access_token=ACCESS_TOKEN&open_kfid=XXX", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultServicerList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListServicer("XXX", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultServicerList{
		ServicerList: []*ServicerListData{
			{
				UserID: "zhangsan",
				Status: 0,
			},
			{
				UserID: "lisi",
				Status: 1,
			},
		},
	}, result)
}

func TestGetServiceState(t *testing.T) {
	body := []byte(`{"open_kfid":"wkxxxxxxxxxxxxxxxxxx","external_userid":"wmxxxxxxxxxxxxxxxxxx"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "service_state": 3,
    "servicer_userid": "zhangsan"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/kf/service_state/get?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultServiceState)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetServiceState("wkxxxxxxxxxxxxxxxxxx", "wmxxxxxxxxxxxxxxxxxx", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultServiceState{
		ServiceState:   3,
		ServicerUserID: "zhangsan",
	}, result)
}

func TestTransferServiceState(t *testing.T) {
	body := []byte(`{"open_kfid":"wkxxxxxxxxxxxxxxxxxx","external_userid":"wmxxxxxxxxxxxxxxxxxx","service_state":3,"servicer_userid":"zhangsan"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "msg_code": "MSG_CODE"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/kf/service_state/trans?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsServiceStateTransfer{
		OpenKFID:       "wkxxxxxxxxxxxxxxxxxx",
		ExternalUserID: "wmxxxxxxxxxxxxxxxxxx",
		ServiceState:   3,
		ServicerUserID: "zhangsan",
	}

	result := new(ResultServiceStateTransfer)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", TransferServiceState(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultServiceStateTransfer{
		MsgCode: "MSG_CODE",
	}, result)
}
