package oa

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/public"
	"github.com/stretchr/testify/assert"
)

func TestCode2AuthToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := public.NewMockClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/sns/oauth2/access_token?appid=APPID&secret=APPSECRET&code=CODE&grant_type=authorization_code").Return([]byte(`{
		"access_token": "ACCESS_TOKEN",
		"expires_in": 7200,
		"refresh_token": "REFRESH_TOKEN",
		"openid": "OPENID",
		"scope": "SCOPE"
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	authToken, err := oa.Code2AuthToken(context.TODO(), "CODE")

	assert.Nil(t, err)
	assert.Equal(t, &AuthToken{
		AccessToken:  "ACCESS_TOKEN",
		RefreshToken: "REFRESH_TOKEN",
		ExpiresIn:    7200,
		OpenID:       "OPENID",
		Scope:        "SCOPE",
	}, authToken)
}

func TestRefreshAuthToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := public.NewMockClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=APPID&grant_type=refresh_token&refresh_token=REFRESH_TOKEN").Return([]byte(`{
		"access_token": "ACCESS_TOKEN",
		"expires_in": 7200,
		"refresh_token": "REFRESH_TOKEN",
		"openid": "OPENID",
		"scope": "SCOPE"
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	authToken, err := oa.RefreshAuthToken(context.TODO(), "REFRESH_TOKEN")

	assert.Nil(t, err)
	assert.Equal(t, &AuthToken{
		AccessToken:  "ACCESS_TOKEN",
		RefreshToken: "REFRESH_TOKEN",
		ExpiresIn:    7200,
		OpenID:       "OPENID",
		Scope:        "SCOPE",
	}, authToken)
}

func TestAccessToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := public.NewMockClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET").Return([]byte(`{
		"access_token": "39_VzXkFDAJsEVTWbXUZDU3NqHtP6mzcAA7RJvcy1o9e-7fdJ-UuxPYLdBFMiGhpdoeKqVWMGqBe8ldUrMasRv1z_T8RmHKDiybC29wZ_vexHlyQ5YDGb33rff1mBNpOLM9f5nv7oag8UYBSc79ASMcAAADVP",
		"expires_in": 7200
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	accessToken, err := oa.AccessToken(context.TODO())

	assert.Nil(t, err)
	assert.Equal(t, &AccessToken{
		Token:     "39_VzXkFDAJsEVTWbXUZDU3NqHtP6mzcAA7RJvcy1o9e-7fdJ-UuxPYLdBFMiGhpdoeKqVWMGqBe8ldUrMasRv1z_T8RmHKDiybC29wZ_vexHlyQ5YDGb33rff1mBNpOLM9f5nv7oag8UYBSc79ASMcAAADVP",
		ExpiresIn: 7200,
	}, accessToken)
}

var postBody public.Body

func TestMain(m *testing.M) {
	postBody = public.NewPostBody(func() ([]byte, error) {
		return nil, nil
	})

	m.Run()
}
