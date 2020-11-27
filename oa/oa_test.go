package oa

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/helpers"
	"github.com/stretchr/testify/assert"
)

func TestCode2AuthToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := helpers.NewMockHTTPClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/sns/oauth2/access_token?appid=wxa06e66cf23dc4370&secret=1208c7f9e08b4edd26fd86406a5b30aa&code=CODE&grant_type=authorization_code").Return([]byte(`{
		"access_token": "ACCESS_TOKEN",
		"expires_in": 7200,
		"refresh_token": "REFRESH_TOKEN",
		"openid": "OPENID",
		"scope": "SCOPE"
	}`), nil)

	oa := New("wxa06e66cf23dc4370", "1208c7f9e08b4edd26fd86406a5b30aa")
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

	client := helpers.NewMockHTTPClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=wxa06e66cf23dc4370&grant_type=refresh_token&refresh_token=REFRESH_TOKEN").Return([]byte(`{
		"access_token": "ACCESS_TOKEN",
		"expires_in": 7200,
		"refresh_token": "REFRESH_TOKEN",
		"openid": "OPENID",
		"scope": "SCOPE"
	}`), nil)

	oa := New("wxa06e66cf23dc4370", "1208c7f9e08b4edd26fd86406a5b30aa")
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

	client := helpers.NewMockHTTPClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wxa06e66cf23dc4370&secret=1208c7f9e08b4edd26fd86406a5b30aa").Return([]byte(`{
		"access_token": "39_VzXkFDAJsEVTWbXUZDU3NqHtP6mzcAA7RJvcy1o9e-7fdJ-UuxPYLdBFMiGhpdoeKqVWMGqBe8ldUrMasRv1z_T8RmHKDiybC29wZ_vexHlyQ5YDGb33rff1mBNpOLM9f5nv7oag8UYBSc79ASMcAAADVP",
		"expires_in": 7200
	}`), nil)

	oa := New("wxa06e66cf23dc4370", "1208c7f9e08b4edd26fd86406a5b30aa")
	oa.client = client

	accessToken, err := oa.AccessToken(context.TODO())

	assert.Nil(t, err)
	assert.Equal(t, &AccessToken{
		Token:     "39_VzXkFDAJsEVTWbXUZDU3NqHtP6mzcAA7RJvcy1o9e-7fdJ-UuxPYLdBFMiGhpdoeKqVWMGqBe8ldUrMasRv1z_T8RmHKDiybC29wZ_vexHlyQ5YDGb33rff1mBNpOLM9f5nv7oag8UYBSc79ASMcAAADVP",
		ExpiresIn: 7200,
	}, accessToken)
}

var postBody helpers.HTTPBody

func TestMain(m *testing.M) {
	postBody = helpers.NewPostBody(func() ([]byte, error) {
		return nil, nil
	})

	m.Run()
}
