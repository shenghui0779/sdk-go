package corpgroup

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/chenghonour/gochat/corp"
	"github.com/chenghonour/gochat/mock"
	"github.com/chenghonour/gochat/wx"
	"github.com/golang/mock/gomock"
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

func TestListChain(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"chains": [
		{
			"chain_id": "chainid1",
			"chain_name": "能源供应链"
		},
		{
			"chain_id": "chainid2",
			"chain_name": "原材料供应链"
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/corpgroup/corp/get_chain_list?access_token=ACCESS_TOKEN", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultChainList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListChain(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultChainList{
		Chains: []*Chain{
			{
				ChainID:   "chainid1",
				ChainName: "能源供应链",
			},
			{
				ChainID:   "chainid2",
				ChainName: "原材料供应链",
			},
		},
	}, result)
}

func TestGetChainGroup(t *testing.T) {
	body := []byte(`{"chain_id":"Chxxxxxx"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"groups": [
		{
			"groupid": 2,
			"group_name": "一级经销商",
			"parentid": 1,
			"order": 1
		},
		{
			"groupid": 3,
			"group_name": "二级经销商",
			"parentid": 2,
			"order": 3
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/corpgroup/corp/get_chain_group?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultChainGroup)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetChainGroup("Chxxxxxx", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultChainGroup{
		Groups: []*ChainGroup{
			{
				GroupID:   2,
				GroupName: "一级经销商",
				ParentID:  1,
				Order:     1,
			},
			{
				GroupID:   3,
				GroupName: "二级经销商",
				ParentID:  2,
				Order:     3,
			},
		},
	}, result)
}

func TestListChainCorpInfo(t *testing.T) {
	body := []byte(`{"chain_id":"Chxxxxxx","groupid":1,"fetch_child":1}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"group_corps": [
		{
			"groupid": 2,
			"corpid": "wwxxxx",
			"corp_name": "美馨粮油公司",
			"custom_id": "wof3du51quo5sl1is"
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/corpgroup/corp/get_chain_corpinfo_list?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultChainCorpInfoList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListChainCorpInfo("Chxxxxxx", 1, 1, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultChainCorpInfoList{
		GroupCorps: []*ChainGroupCorp{
			{
				GroupID:  2,
				CorpID:   "wwxxxx",
				CorpName: "美馨粮油公司",
				CustomID: "wof3du51quo5sl1is",
			},
		},
	}, result)
}

func TestUnionIDToExternalUserID(t *testing.T) {
	body := []byte(`{"unionid":"xxxxx","openid":"xxxxx","corpid":"xxxxx"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"external_userid_info": [
		{
			"corpid": "AAAAA",
			"external_userid": "BBBB"
		},
		{
			"corpid": "CCCCC",
			"external_userid": "DDDDD"
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/corpgroup/unionid_to_external_userid?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultUnionIDToExternalUserID)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UnionIDToExternalUserID("xxxxx", "xxxxx", "xxxxx", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultUnionIDToExternalUserID{
		ExternalUserIDInfo: []*ExternalUserIDInfo{
			{
				CorpID:         "AAAAA",
				ExternalUserID: "BBBB",
			},
			{
				CorpID:         "CCCCC",
				ExternalUserID: "DDDDD",
			},
		},
	}, result)
}
