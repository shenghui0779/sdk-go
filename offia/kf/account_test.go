package kf

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/chenghonour/gochat/mock"
	"github.com/chenghonour/gochat/offia"
	"github.com/chenghonour/gochat/wx"
)

func TestGetAccountList(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
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
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://api.weixin.qq.com/cgi-bin/customservice/getkflist?access_token=ACCESS_TOKEN", nil).Return(resp, nil)

	oa := offia.New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(ResultAccountList)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetAccountList(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultAccountList{
		KFList: []*Account{
			{
				KFID:         "1001",
				KFAccount:    "test1@test",
				KFNick:       "ntest1",
				KFHeadImgURL: "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl2iccsvYbHvnphkyGtnvjfUS8Ym0GSaLic0FD3vN0V8PILcibEGb2fPfEOmw/0",
				KFWX:         "kfwx1",
			},
			{
				KFID:             "1002",
				KFAccount:        "test2@test",
				KFNick:           "ntest2",
				KFHeadImgURL:     "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl2iccsvYbHvnphkyGtnvjfUS8Ym0GSaLic0FD3vN0V8PILcibEGb2fPfEOmw/0",
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
		Body: io.NopCloser(bytes.NewReader([]byte(`{
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
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://api.weixin.qq.com/cgi-bin/customservice/getonlinekflist?access_token=ACCESS_TOKEN", nil).Return(resp, nil)

	oa := offia.New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	result := new(ResultOnlineList)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetOnlineList(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultOnlineList{
		KFOnlineList: []*Online{
			{
				KFID:         "1001",
				KFAccount:    "test1@test",
				Status:       1,
				AcceptedCase: 1,
			},
			{
				KFID:         "1002",
				KFAccount:    "test2@test",
				Status:       1,
				AcceptedCase: 2,
			},
		},
	}, result)
}

func TestAddAccount(t *testing.T) {
	body := []byte(`{"kf_account":"test1@test","nickname":"客服1"}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/customservice/kfaccount/add?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := offia.New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", AddAccount("test1@test", "客服1"))

	assert.Nil(t, err)
}

func TestUpdateAccount(t *testing.T) {
	body := []byte(`{"kf_account":"test1@test","nickname":"客服1"}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/customservice/kfaccount/update?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := offia.New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UpdateAccount("test1@test", "客服1"))

	assert.Nil(t, err)
}

func TestInviteWorker(t *testing.T) {
	body := []byte(`{"kf_account":"test1@test","invite_wx":"test_kfwx"}`)

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.weixin.qq.com/customservice/kfaccount/inviteworker?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	oa := offia.New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", InviteWorker("test1@test", "test_kfwx"))

	assert.Nil(t, err)
}

func TestUploadAvatar(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/customservice/kfaccount/uploadheadimg?access_token=ACCESS_TOKEN&kf_account=ACCOUNT", gomock.AssignableToTypeOf(wx.NewUploadForm())).Return(resp, nil)

	oa := offia.New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UploadAvatar("ACCOUNT", "../../mock/test.jpg"))

	assert.Nil(t, err)
}

func TestDeleteAccount(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"errcode":0,"errmsg":"ok"}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://api.weixin.qq.com/customservice/kfaccount/del?access_token=ACCESS_TOKEN&kf_account=ACCOUNT", nil).Return(resp, nil)

	oa := offia.New("APPID", "APPSECRET")
	oa.SetClient(wx.WithHTTPClient(client))

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", DeleteAccount("ACCOUNT"))

	assert.Nil(t, err)
}
