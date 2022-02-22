package corp

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/mock"
	"github.com/shenghui0779/gochat/wx"
	"github.com/stretchr/testify/assert"
)

func TestGetAPIDomainIP(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"ip_list": [
		"182.254.11.176",
		"182.254.78.66"
	],
	"errcode": 0,
	"errmsg": "ok"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/get_api_domain_ip?access_token=ACCESS_TOKEN", nil).Return(resp, nil)

	cp := New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultIP)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetAPIDomainIP(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultIP{
		IPList: []string{"182.254.11.176", "182.254.78.66"},
	}, result)
}

func TestGetCallbackIP(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"ip_list": [
		"101.226.103.*",
		"101.226.62.*"
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/getcallbackip?access_token=ACCESS_TOKEN", nil).Return(resp, nil)

	cp := New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultIP)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetCallbackIP(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultIP{
		IPList: []string{"101.226.103.*", "101.226.62.*"},
	}, result)
}

func TestGetOAuthUser(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"UserId":"USERID",
	"DeviceId":"DEVICEID"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?access_token=ACCESS_TOKEN&code=CODE", nil).Return(resp, nil)

	cp := New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultOAuthUser)
	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetOAuthUser("CODE", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultOAuthUser{
		UserID:   "USERID",
		DeviceID: "DEVICEID",
	}, result)
}

func TestUserAuthSucc(t *testing.T) {
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/user/authsucc?access_token=ACCESS_TOKEN&userid=USERID", nil).Return(resp, nil)

	cp := New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UserAuthSucc("USERID"))

	assert.Nil(t, err)
}
