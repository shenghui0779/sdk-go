package corpgroup

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

func TestListAppShareInfo(t *testing.T) {
	body := []byte(`{"agentid":1111,"business_type":1,"corpid":"wwcorp","limit":100,"cursor":"xxxxxx"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"ending": 0,
	"corp_list": [
		{
			"corpid": "wwcorpid1",
			"corp_name": "测试企业1",
			"agentid": 1111
		},
		{
			"corpid": "wwcorpid2",
			"corp_name": "测试企业2",
			"agentid": 1112
		}
	],
	"next_cursor": "next_cursor1111"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/corpgroup/corp/list_app_share_info?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsAppShareInfoList{
		AgentID:      1111,
		BusinessType: 1,
		CorpID:       "wwcorp",
		Limit:        100,
		Cursor:       "xxxxxx",
	}

	result := new(ResultAppShareInfoList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListAppShareInfo(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultAppShareInfoList{
		Ending: 0,
		CorpList: []*CorpInfo{
			{
				CorpID:   "wwcorpid1",
				CorpName: "测试企业1",
				AgentID:  1111,
			},
			{
				CorpID:   "wwcorpid2",
				CorpName: "测试企业2",
				AgentID:  1112,
			},
		},
		NextCursor: "next_cursor1111",
	}, result)
}

func TestGetCorpAccessToken(t *testing.T) {
	body := []byte(`{"corpid":"wwabc","business_type":1,"agentid":1111}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"access_token": "accesstoken000001",
	"expires_in": 7200
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/corpgroup/corp/gettoken?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsCorpAccessToken{
		CorpID:       "wwabc",
		BusinessType: 1,
		AgentID:      1111,
	}

	result := new(ResultCorpAccessToken)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetCorpAccessToken(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultCorpAccessToken{
		AccessToken: "accesstoken000001",
		ExpiresIn:   7200,
	}, result)
}

func TestTransferMinipSession(t *testing.T) {
	body := []byte(`{"userid":"wmAoNVCwAAUrSqEqz7oQpEIEMVWDrPeg","session_key":"n8cnNEoyW1pxSRz6/Lwjwg=="}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"userid": "abcdef",
	"session_key": "DGAuy2KVaGcnsUrXk8ERgw==",
	"errcode": 0,
	"errmsg": "ok"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/miniprogram/transfer_session?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsMinipSessionTransfer{
		UserID:     "wmAoNVCwAAUrSqEqz7oQpEIEMVWDrPeg",
		SessionKey: "n8cnNEoyW1pxSRz6/Lwjwg==",
	}

	result := new(ResultMinipSessionTransfer)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", TransferMinipSession(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMinipSessionTransfer{
		UserID:     "abcdef",
		SessionKey: "DGAuy2KVaGcnsUrXk8ERgw==",
	}, result)
}
