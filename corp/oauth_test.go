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
