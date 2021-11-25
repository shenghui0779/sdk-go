package kf

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/mock"
	"github.com/shenghui0779/gochat/offia"
	"github.com/shenghui0779/yiigo"
	"github.com/stretchr/testify/assert"
)

func TestGetAccountList(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader()),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/customservice/getkflist?access_token=ACCESS_TOKEN").Return([]byte(`{
		"kf_list" : [
		   {
			  "kf_account": "test1@test",
			  "kf_headimgurl": "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl2iccsvYbHvnphkyGtnvjfUS8Ym0GSaLic0FD3vN0V8PILcibEGb2fPfEOmw/0",
			  "kf_id": "1001",
			  "kf_nick": "ntest1",
			  "kf_wx": "kfwx1"
		   },
		   {
			  "kf_account": "test2@test",
			  "kf_headimgurl": "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl2iccsvYbHvnphkyGtnvjfUS8Ym0GSaLic0FD3vN0V8PILcibEGb2fPfEOmw/0",
			  "kf_id": "1002",
			  "kf_nick": "ntest2",
			  "invite_wx": "kfwx2",
			  "invite_expire_time": 123456789,
			  "invite_status": "waiting"
		   }
		]
	}`), nil)

	oa := offia.New("APPID", "APPSECRET")
	oa.SetClient(client)

	result := new(ResultAccountList)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetAccountList(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultAccountList{
		KFList: []*Account{
			{
				ID:         "1001",
				Account:    "test1@test",
				Nickname:   "ntest1",
				HeadImgURL: "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl2iccsvYbHvnphkyGtnvjfUS8Ym0GSaLic0FD3vN0V8PILcibEGb2fPfEOmw/0",
				Weixin:     "kfwx1",
			},
			{
				ID:               "1002",
				Account:          "test2@test",
				Nickname:         "ntest2",
				HeadImgURL:       "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl2iccsvYbHvnphkyGtnvjfUS8Ym0GSaLic0FD3vN0V8PILcibEGb2fPfEOmw/0",
				InviteWeixin:     "kfwx2",
				InviteExpireTime: 123456789,
				InviteStatus:     InviteWaiting,
			},
		},
	}, result)
}

func TestGetOnlineList(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader()),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/customservice/getonlinekflist?access_token=ACCESS_TOKEN").Return([]byte(`{
		"kf_online_list": [
			{
				"kf_account": "test1@test",
				"status": 1,
				"kf_id": "1001",
				"accepted_case": 1
			},
			{
				"kf_account": "test2@test",
				"status": 1,
				"kf_id": "1002",
				"accepted_case": 2
			}
		]
	}`), nil)

	oa := offia.New("APPID", "APPSECRET")
	oa.SetClient(client)

	result := new(ResultOnlineList)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetOnlineList(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultOnlineList{
		KFOnlineList: []*Online{
			{
				ID:           "1001",
				Account:      "test1@test",
				Status:       1,
				AcceptedCase: 1,
			},
			{
				ID:           "1002",
				Account:      "test2@test",
				Status:       1,
				AcceptedCase: 2,
			},
		},
	}, result)
}

func TestAddAccount(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader()),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/customservice/kfaccount/add?access_token=ACCESS_TOKEN", []byte(`{"kf_account":"test1@test","nickname":"客服1"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := offia.New("APPID", "APPSECRET")
	oa.SetClient(client)

	params := &ParamsAccountAdd{
		Account:  "test1@test",
		Nickname: "客服1",
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", AddAccount(params))

	assert.Nil(t, err)
}

func TestUpdateAccount(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader()),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/customservice/kfaccount/update?access_token=ACCESS_TOKEN", []byte(`{"kf_account":"test1@test","nickname":"客服1"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := offia.New("APPID", "APPSECRET")
	oa.SetClient(client)

	params := &ParamsAccountUpdate{
		Account:  "test1@test",
		Nickname: "客服1",
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UpdateAccount(params))

	assert.Nil(t, err)
}

func TestInviteWorker(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader()),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/customservice/kfaccount/inviteworker?access_token=ACCESS_TOKEN", []byte(`{"kf_account":"test1@test","invite_wx":"test_kfwx"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := offia.New("APPID", "APPSECRET")
	oa.SetClient(client)

	params := &ParamsWorkerInvite{
		KFAccount: "test1@test",
		InviteWX:  "test_kfwx",
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", InviteWorker(params))

	assert.Nil(t, err)
}

func TestUploadAvatar(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader()),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/customservice/kfaccount/uploadheadimg?access_token=ACCESS_TOKEN&kf_account=ACCOUNT", gomock.AssignableToTypeOf(yiigo.NewUploadForm())).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := offia.New("APPID", "APPSECRET")
	oa.SetClient(client)

	params := &ParamsAvatarUpload{
		KFAccount: "ACCOUNT",
		Path:      "../../mock/test.jpg",
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UploadAvatar(params))

	assert.Nil(t, err)
}

func TestDeleteAccount(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader()),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/customservice/kfaccount/del?access_token=ACCESS_TOKEN&kf_account=ACCOUNT").Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := offia.New("APPID", "APPSECRET")
	oa.SetClient(client)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", DeleteAccount("ACCOUNT"))

	assert.Nil(t, err)
}
