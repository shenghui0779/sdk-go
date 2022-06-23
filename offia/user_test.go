package offia

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

func TestCreateTag(t *testing.T) {
	body := []byte(`{"tag":{"name":"广东"}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "tag": {
        "id": 134,
        "name": "广东"
    }
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/tags/create?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(ResultTagCreate)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", CreateTag("广东", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultTagCreate{
		Tag: &Tag{
			ID:   134,
			Name: "广东",
		},
	}, result)
}

func TestUpdateTag(t *testing.T) {
	body := []byte(`{"tag":{"id":134,"name":"广东人"}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/tags/update?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UpdateTag(134, "广东人"))

	assert.Nil(t, err)
}

func TestGetTags(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "tags": [
        {
            "id": 1,
            "name": "每天一罐可乐星人",
            "count": 0
        },
        {
            "id": 2,
            "name": "星标组",
            "count": 0
        },
        {
            "id": 127,
            "name": "广东",
            "count": 5
        }
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://api.weixin.qq.com/cgi-bin/tags/get?access_token=ACCESS_TOKEN", nil).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(ResultTagsGet)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetTags(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultTagsGet{
		Tags: []*Tag{
			{
				ID:    1,
				Name:  "每天一罐可乐星人",
				Count: 0,
			},
			{
				ID:    2,
				Name:  "星标组",
				Count: 0,
			},
			{
				ID:    127,
				Name:  "广东",
				Count: 5,
			},
		},
	}, result)
}

func TestDeleteTag(t *testing.T) {
	body := []byte(`{"tag":{"id":134}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/tags/delete?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", DeleteTag(134))

	assert.Nil(t, err)
}

func TestGetTagUsers(t *testing.T) {
	body := []byte(`{"tagid":134,"next_openid":""}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "count": 2,
    "data": {
        "openid": [
            "ocYxcuAEy30bX0NXmGn4ypqx3tI0",
            "ocYxcuBt0mRugKZ7tGAHPnUaOW7Y"
        ]
    },
    "next_openid": "ocYxcuBt0mRugKZ7tGAHPnUaOW7Y"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/user/tag/get?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(ResultTagUsers)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetTagUsers(134, "", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultTagUsers{
		Count: 2,
		Data: &TagUserData{
			OpenID: []string{"ocYxcuAEy30bX0NXmGn4ypqx3tI0", "ocYxcuBt0mRugKZ7tGAHPnUaOW7Y"},
		},
		NextOpenID: "ocYxcuBt0mRugKZ7tGAHPnUaOW7Y",
	}, result)
}

func TestBatchTaggingUsers(t *testing.T) {
	body := []byte(`{"tagid":134,"openid_list":["ocYxcuAEy30bX0NXmGn4ypqx3tI0","ocYxcuBt0mRugKZ7tGAHPnUaOW7Y"]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/tags/members/batchtagging?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", BatchTaggingUsers(134, "ocYxcuAEy30bX0NXmGn4ypqx3tI0", "ocYxcuBt0mRugKZ7tGAHPnUaOW7Y"))

	assert.Nil(t, err)
}

func TestBatchUnTaggingUsers(t *testing.T) {
	body := []byte(`{"tagid":134,"openid_list":["ocYxcuAEy30bX0NXmGn4ypqx3tI0","ocYxcuBt0mRugKZ7tGAHPnUaOW7Y"]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/tags/members/batchuntagging?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", BatchUnTaggingUsers(134, "ocYxcuAEy30bX0NXmGn4ypqx3tI0", "ocYxcuBt0mRugKZ7tGAHPnUaOW7Y"))

	assert.Nil(t, err)
}

func TestGetUserTags(t *testing.T) {
	body := []byte(`{"openid":"ocYxcuBt0mRugKZ7tGAHPnUaOW7Y"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "tagid_list": [
        134,
        2
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/tags/getidlist?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(ResultUserTags)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetUserTags("ocYxcuBt0mRugKZ7tGAHPnUaOW7Y", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultUserTags{
		TagIDList: []int64{134, 2},
	}, result)
}

func TestSetUserRemark(t *testing.T) {
	body := []byte(`{"openid":"oDF3iY9ffA-hqb2vVvbr7qxf6A0Q","remark":"pangzi"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/user/info/updateremark?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SetUserRemark("oDF3iY9ffA-hqb2vVvbr7qxf6A0Q", "pangzi"))

	assert.Nil(t, err)
}

func TestGetUserInfo(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"subscribe": 1,
	"openid": "o6_bmjrPTlm6_2sgVt7hMZOPfL2M",
	"nickname": "Band",
	"sex": 1,
	"language": "zh_CN",
	"city": "广州",
	"province": "广东",
	"country": "中国",
	"headimgurl": "http://thirdwx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/0",
	"subscribe_time": 1382694957,
	"unionid": "o6_bmasdasdsad6_2sgVt7hMZOPfL",
	"remark": "",
	"groupid": 0,
	"tagid_list": [128, 2],
	"subscribe_scene": "ADD_SCENE_QR_CODE",
	"qr_scene": 98765,
	"qr_scene_str": ""
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://api.weixin.qq.com/cgi-bin/user/info?access_token=ACCESS_TOKEN&lang=zh_CN&openid=OPENID", nil).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(UserInfo)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetUserInfo("OPENID", "zh_CN", result))

	assert.Nil(t, err)
	assert.Equal(t, &UserInfo{
		Subscribe:      1,
		OpenID:         "o6_bmjrPTlm6_2sgVt7hMZOPfL2M",
		NickName:       "Band",
		Sex:            1,
		Country:        "中国",
		City:           "广州",
		Province:       "广东",
		Language:       "zh_CN",
		HeadImgURL:     "http://thirdwx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/0",
		SubscribeTime:  1382694957,
		UnionID:        "o6_bmasdasdsad6_2sgVt7hMZOPfL",
		Remark:         "",
		GroupID:        0,
		TagidList:      []int64{128, 2},
		SubscribeScene: AddSceneQRCode,
		QRScene:        98765,
		QRSceneStr:     "",
	}, result)
}

func TestBatchGetUserInfo(t *testing.T) {
	body := []byte(`{"user_list":[{"openid":"otvxTs4dckWG7imySrJd6jSi0CWE","lang":"zh_CN"},{"openid":"otvxTs_JZ6SEiP0imdhpi50fuSZg","lang":"zh_CN"}]}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"user_info_list": [
		{
			"subscribe": 1,
			"openid": "otvxTs4dckWG7imySrJd6jSi0CWE",
			"nickname": "iWithery",
			"sex": 1,
			"language": "zh_CN",
			"city": "揭阳",
			"province": "广东",
			"country": "中国",
			"headimgurl": "http://thirdwx.qlogo.cn/mmopen/xbIQx1GRqdvyqkMMhEaGOX802l1CyqMJNgUzKP8MeAeHFicRDSnZH7FY4XB7p8XHXIf6uJA2SCunTPicGKezDC4saKISzRj3nz/0",
			"subscribe_time": 1434093047,
			"unionid": "oR5GjjgEhCMJFyzaVZdrxZ2zRRF4",
			"remark": "",
			"groupid": 0,
			"tagid_list": [128, 2],
			"subscribe_scene": "ADD_SCENE_QR_CODE",
			"qr_scene": 98765,
			"qr_scene_str": ""
		},
		{
			"subscribe": 0,
			"openid": "otvxTs_JZ6SEiP0imdhpi50fuSZg"
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/user/info/batchget?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	users := []*ParamsUserInfo{
		{
			OpenID: "otvxTs4dckWG7imySrJd6jSi0CWE",
			Lang:   "zh_CN",
		},
		{
			OpenID: "otvxTs_JZ6SEiP0imdhpi50fuSZg",
			Lang:   "zh_CN",
		},
	}

	result := new(ResultBatchUserInfo)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", BatchGetUserInfo(users, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultBatchUserInfo{
		UserInfoList: []*UserInfo{
			{
				Subscribe:      1,
				OpenID:         "otvxTs4dckWG7imySrJd6jSi0CWE",
				NickName:       "iWithery",
				Sex:            1,
				Country:        "中国",
				City:           "揭阳",
				Province:       "广东",
				Language:       "zh_CN",
				HeadImgURL:     "http://thirdwx.qlogo.cn/mmopen/xbIQx1GRqdvyqkMMhEaGOX802l1CyqMJNgUzKP8MeAeHFicRDSnZH7FY4XB7p8XHXIf6uJA2SCunTPicGKezDC4saKISzRj3nz/0",
				SubscribeTime:  1434093047,
				UnionID:        "oR5GjjgEhCMJFyzaVZdrxZ2zRRF4",
				Remark:         "",
				GroupID:        0,
				TagidList:      []int64{128, 2},
				SubscribeScene: AddSceneQRCode,
				QRScene:        98765,
				QRSceneStr:     "",
			},
			{
				Subscribe: 0,
				OpenID:    "otvxTs_JZ6SEiP0imdhpi50fuSZg",
			},
		},
	}, result)
}

func TestListUser(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"total": 2,
	"count": 2,
	"data": {
		"openid": ["OPENID1", "OPENID2"]
	},
	"next_openid": "NEXT_OPENID"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://api.weixin.qq.com/cgi-bin/user/get?access_token=ACCESS_TOKEN&next_openid=NEXT_OPENID", nil).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(ResultUserList)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", ListUser("NEXT_OPENID", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultUserList{
		Total: 2,
		Count: 2,
		Data: UserListData{
			OpenID: []string{"OPENID1", "OPENID2"},
		},
		NextOpenID: "NEXT_OPENID",
	}, result)
}

func TestListBlackUsers(t *testing.T) {
	body := []byte(`{"begin_openid":"OPENID1"}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"total": 3,
	"count": 3,
	"data": {
		"openid": [
			"OPENID1",
			"OPENID2",
			"OPENID10000"
			]
	},
	"next_openid": "OPENID10000"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/tags/members/getblacklist?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(ResultBlackList)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", ListBlackUsers("OPENID1", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultBlackList{
		Total: 3,
		Count: 3,
		Data: UserListData{
			OpenID: []string{"OPENID1", "OPENID2", "OPENID10000"},
		},
		NextOpenID: "OPENID10000",
	}, result)
}

func TestBatchBlackUsers(t *testing.T) {
	body := []byte(`{"openid_list":["OPENID1","OPENID2"]}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/tags/members/batchblacklist?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", BatchBlackUsers("OPENID1", "OPENID2"))

	assert.Nil(t, err)
}

func TestBatchUnBlackUsers(t *testing.T) {
	body := []byte(`{"openid_list":["OPENID1","OPENID2"]}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/cgi-bin/tags/members/batchunblacklist?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", BatchUnBlackUsers("OPENID1", "OPENID2"))

	assert.Nil(t, err)
}
