package oa

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/internal"
	"github.com/stretchr/testify/assert"
)

func TestCheckAuthToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := internal.NewMockClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/sns/auth?access_token=ACCESS_TOKEN&openid=OPENID").Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", CheckAuthToken("OPENID"))

	assert.Nil(t, err)
}

func TestGetAuthUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := internal.NewMockClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/sns/userinfo?access_token=ACCESS_TOKEN&lang=zh_CN&openid=OPENID").Return([]byte(`{
		"openid": "OPENID",
		"nickname": "NICKNAME",
		"sex": "1",
		"province": "PROVINCE",
		"city": "CITY",
		"country": "COUNTRY",
		"headimgurl": "https://thirdwx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/46",
		"privilege": ["PRIVILEGE1", "PRIVILEGE2"],
		"unionid": "o6_bmasdasdsad6_2sgVt7hMZOPfL"
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(AuthUser)
	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetAuthUser("OPENID", dest))

	assert.Nil(t, err)
	assert.Equal(t, &AuthUser{
		OpenID:     "OPENID",
		UnionID:    "o6_bmasdasdsad6_2sgVt7hMZOPfL",
		Nickname:   "NICKNAME",
		Sex:        "1",
		Province:   "PROVINCE",
		City:       "CITY",
		Country:    "COUNTRY",
		HeadImgURL: "https://thirdwx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/46",
		Privilege:  []string{"PRIVILEGE1", "PRIVILEGE2"},
	}, dest)
}
