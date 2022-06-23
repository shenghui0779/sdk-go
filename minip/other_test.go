package minip

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/chenghonour/gochat/mock"
	"github.com/chenghonour/gochat/wx"
)

func TestInvokeService(t *testing.T) {
	body := []byte(`{"service":"wx79ac3de8be320b71","api":"OcrAllInOne","data":{"data_type":3,"img_url":"http://mmbiz.qpic.cn/mmbiz_jpg/7UFjuNbYxibu66xSqsQqKcuoGBZM77HIyibdiczeWibdMeA2XMt5oibWVQMgDibriazJSOibLqZxcO6DVVcZMxDKgeAtbQ/0","ocr_type":1},"client_msg_id":"id123"}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"data": "{\"idcard_res\":{\"type\":0,\"name\":{\"text\":\"abc\",\"pos\"…0312500}}},\"image_width\":480,\"image_height\":304}}"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/wxa/servicemarket?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	mp := New("APPID", "APPSECRET")
	mp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsServiceInvoke{
		Service: "wx79ac3de8be320b71",
		API:     "OcrAllInOne",
		Data: wx.M{
			"data_type": 3,
			"img_url":   "http://mmbiz.qpic.cn/mmbiz_jpg/7UFjuNbYxibu66xSqsQqKcuoGBZM77HIyibdiczeWibdMeA2XMt5oibWVQMgDibriazJSOibLqZxcO6DVVcZMxDKgeAtbQ/0",
			"ocr_type":  1,
		},
		ClientMsgID: "id123",
	}

	result := new(ResultServiceInvoke)

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", InvokeService(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultServiceInvoke{
		Data: `{"idcard_res":{"type":0,"name":{"text":"abc","pos"…0312500}}},"image_width":480,"image_height":304}}`,
	}, result)
}

func TestSoterVerify(t *testing.T) {
	body := []byte(`{"openid":"$openid","json_string":"$resultJSON","json_signature":"$resultJSONSignature"}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"is_ok": true
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/soter/verify_signature?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	mp := New("APPID", "APPSECRET")
	mp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultSoterVerify)

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", SoterVerify("$openid", "$resultJSON", "$resultJSONSignature", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultSoterVerify{
		IsOK: true,
	}, result)
}

func TestGenerateShortLink(t *testing.T) {
	body := []byte(`{"page_url":"/pages/publishHomework/publishHomework?query1=q1","page_title":"Homework title","is_permanent":false}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "link": "Short Link"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/wxa/genwxashortlink?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(ResultShortLink)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GenerateShortLink("/pages/publishHomework/publishHomework?query1=q1", "Homework title", false, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultShortLink{
		Link: "Short Link",
	}, result)
}

func TestGetUserRiskRank(t *testing.T) {
	body := []byte(`{"appid":"APPID","openid":"OPENID","scene":1,"mobile_no":"12345678","client_ip":"******","email_address":"****@qq.com"}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"risk_rank": 0
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/wxa/getuserriskrank?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	mp := New("APPID", "APPSECRET")
	mp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsUserRisk{
		AppID:        "APPID",
		OpenID:       "OPENID",
		Scene:        RiskCheat,
		MobileNO:     "12345678",
		ClientIP:     "******",
		EmailAddress: "****@qq.com",
	}

	result := new(ResultUserRisk)

	err := mp.Do(context.TODO(), "ACCESS_TOKEN", GetUserRiskRank(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultUserRisk{
		RiskRank: 0,
	}, result)
}

func TestGenerateScheme(t *testing.T) {
	body := []byte(`{"jump_wxa":{"path":"/pages/publishHomework/publishHomework"},"is_expire":true,"expire_time":1606737600}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"openlink": "Scheme"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/wxa/generatescheme?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	params := &ParamsSchemeGenerate{
		JumpWxa: &SchemeJumpWxa{
			Path: "/pages/publishHomework/publishHomework",
		},
		IsExpire:   true,
		ExpireTime: 1606737600,
	}
	result := new(ResultSchemeGenerate)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GenerateScheme(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultSchemeGenerate{
		OpenLink: "Scheme",
	}, result)
}

func TestQueryScheme(t *testing.T) {
	body := []byte(`{"scheme":"weixin://dl/business/?t=XTSkBZlzqmn"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"scheme_info": {
		"appid": "appid",
		"path": "",
		"query": "",
		"create_time": 611928113,
		"expire_time": 0,
		"env_version": "release"
	},
	"scheme_quota": {
		"long_time_used": 100,
		"long_time_limit": 100000
	}
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/wxa/queryscheme?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(ResultSchemeQuery)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", QueryScheme("weixin://dl/business/?t=XTSkBZlzqmn", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultSchemeQuery{
		SchemeInfo: &SchemeInfo{
			AppID:      "appid",
			Path:       "",
			Query:      "",
			CreateTime: 611928113,
			ExpireTime: 0,
			EnvVersion: EnvRelease,
		},
		SchemeQuota: &SchemeQuota{
			LongTimeUsed:  100,
			LongTimeLimit: 100000,
		},
	}, result)
}

func TestGenerateURLLink(t *testing.T) {
	body := []byte(`{"path":"/pages/publishHomework/publishHomework","is_expire":true,"expire_type":1,"expire_interval":1,"env_version":"release","cloud_base":{"env":"xxx","domain":"xxx.xx","path":"/jump-wxa.html","query":"a=1&b=2"}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"url_link": "URL Link"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/wxa/generate_urllink?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	params := &ParamsURLLinkGenerate{
		Path:           "/pages/publishHomework/publishHomework",
		IsExpire:       true,
		ExpireType:     1,
		ExpireInterval: 1,
		EnvVersion:     EnvRelease,
		CloudBase: &CloudBase{
			Env:    "xxx",
			Domain: "xxx.xx",
			Path:   "/jump-wxa.html",
			Query:  "a=1&b=2",
		},
	}
	result := new(ResultURLLinkGenerate)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GenerateURLLink(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultURLLinkGenerate{
		URLLink: "URL Link",
	}, result)
}

func TestQueryURLLink(t *testing.T) {
	body := []byte(`{"url_link":"https://wxaurl.cn/BQZRrcFCPvg"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"url_link_info": {
		"appid": "appid",
		"path": "",
		"query": "",
		"create_time": 611928113,
		"expire_time": 0,
		"env_version": "release",
		"cloud_base": {
			"env": "",
			"doamin": "",
			"path": "",
			"query": "",
			"resource_appid": ""
		}
	},
	"url_link_quota": {
		"long_time_used": 100,
		"long_time_limit": 100000
	}
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/wxa/query_urllink?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(ResultURLLinkQuery)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", QueryURLLink("https://wxaurl.cn/BQZRrcFCPvg", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultURLLinkQuery{
		URLLinkInfo: &URLLinkInfo{
			AppID:      "appid",
			Path:       "",
			Query:      "",
			CreateTime: 611928113,
			ExpireTime: 0,
			EnvVersion: EnvRelease,
			CloudBase: &CloudBase{
				Env:           "",
				Domain:        "",
				Path:          "",
				Query:         "",
				ResourceAppID: "",
			},
		},
		URLLinkQuota: &URLLinkQuota{
			LongTimeUsed:  100,
			LongTimeLimit: 100000,
		},
	}, result)
}
