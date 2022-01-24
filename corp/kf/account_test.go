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

func TestAddAccount(t *testing.T) {
	body := []byte(`{"name":"新建的客服帐号","media_id":"294DpAog3YA5b9rTK4PjjfRfYLO0L5qpDHAJIzhhQ2jAEWjb9i661Q4lk8oFnPtmj"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "open_kfid": "wkAJ2GCAAAZSfhHCt7IFSvLKtMPxyJTw"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/kf/account/add?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsAccountAdd{
		Name:    "新建的客服帐号",
		MediaID: "294DpAog3YA5b9rTK4PjjfRfYLO0L5qpDHAJIzhhQ2jAEWjb9i661Q4lk8oFnPtmj",
	}

	result := new(ResultAccountAdd)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", AddAccount(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultAccountAdd{
		OpenKFID: "wkAJ2GCAAAZSfhHCt7IFSvLKtMPxyJTw",
	}, result)
}

func TestDeleteAccount(t *testing.T) {
	body := []byte(`{"open_kfid":"wkAJ2GCAAAZSfhHCt7IFSvLKtMPxyJTw"}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/kf/account/del?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", DeleteAccount("wkAJ2GCAAAZSfhHCt7IFSvLKtMPxyJTw"))

	assert.Nil(t, err)
}

func TestUpdateAccount(t *testing.T) {
	body := []byte(`{"open_kfid":"wkAJ2GCAAAZSfhHCt7IFSvLKtMPxyJTw","name":"修改客服名","media_id":"294DpAog3YA5b9rTK4PjjfRfYLO0L5qpDHAJIzhhQ2jAEWjb9i661Q4lk8oFnPtmj"}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/kf/account/update?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsAccountUpdate{
		OpenKFID: "wkAJ2GCAAAZSfhHCt7IFSvLKtMPxyJTw",
		Name:     "修改客服名",
		MediaID:  "294DpAog3YA5b9rTK4PjjfRfYLO0L5qpDHAJIzhhQ2jAEWjb9i661Q4lk8oFnPtmj",
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UpdateAccount(params))

	assert.Nil(t, err)
}

func TestListAccount(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "account_list": [
        {
            "open_kfid": "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
            "name": "咨询客服",
            "avatar": "https://wework.qpic.cn/wwhead/duc2TvpEgSSjibPZlNR6chpx9W3dtd9Ogp8XEmSNKGa6uufMWn2239HUPuwIFoYYZ7Ph580FPvo8/0"
        }
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/kf/account/list?access_token=ACCESS_TOKEN", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultAccountList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListAccount(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultAccountList{
		AccountList: []*AccountListData{
			{
				OpenKFID: "wkAJ2GCAAASSm4_FhToWMFea0xAFfd3Q",
				Name:     "咨询客服",
				Avatar:   "https://wework.qpic.cn/wwhead/duc2TvpEgSSjibPZlNR6chpx9W3dtd9Ogp8XEmSNKGa6uufMWn2239HUPuwIFoYYZ7Ph580FPvo8/0",
			},
		},
	}, result)
}

func TestAddContactWay(t *testing.T) {
	body := []byte(`{"open_kfid":"OPEN_KFID","scene":"12345"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "url": "https://work.weixin.qq.com/kf/kfcbf8f8d07ac7215f?enc_scene=ENCGFSDF567DF"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/kf/add_contact_way?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsContactWayAdd{
		OpenKFID: "OPEN_KFID",
		Scene:    "12345",
	}

	result := new(ResultContactWayAdd)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", AddContactWay(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultContactWayAdd{
		URL: "https://work.weixin.qq.com/kf/kfcbf8f8d07ac7215f?enc_scene=ENCGFSDF567DF",
	}, result)
}
