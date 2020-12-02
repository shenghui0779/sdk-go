package mp

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/wx"
	"github.com/stretchr/testify/assert"
)

func TestCode2Session(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=APPSECRET&js_code=JSCODE&grant_type=authorization_code").Return([]byte(`{
		"openid": "OPENID",
		"session_key": "SESSION_KEY",
		"unionid": "UNIONID",
		"errcode": 0,
		"errmsg": "ok"
	}`), nil)

	mp := New("APPID", "APPSECRET")
	mp.client = client

	authSession, err := mp.Code2Session(context.TODO(), "JSCODE")

	assert.Nil(t, err)
	assert.Equal(t, &AuthSession{
		SessionKey: "SESSION_KEY",
		OpenID:     "OPENID",
		UnionID:    "UNIONID",
	}, authSession)
}

func TestAccessToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/token?appid=APPID&secret=APPSECRET&grant_type=client_credential").Return([]byte(`{
		"access_token": "ACCESS_TOKEN",
		"expires_in": 7200,
		"errcode": 0,
		"errmsg": "ok"
	}`), nil)

	mp := New("APPID", "APPSECRET")
	mp.client = client

	accessToken, err := mp.AccessToken(context.TODO())

	assert.Nil(t, err)
	assert.Equal(t, &AccessToken{
		Token:     "ACCESS_TOKEN",
		ExpiresIn: 7200,
	}, accessToken)
}

var postBody wx.Body

func TestMain(m *testing.M) {
	postBody = wx.NewPostBody(func() ([]byte, error) {
		return nil, nil
	})

	m.Run()
}
