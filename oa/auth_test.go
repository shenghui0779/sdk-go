package oa

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/helpers"
	"github.com/stretchr/testify/assert"
)

func TestCheckAuthToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := helpers.NewMockHTTPClient(ctrl)

	client.EXPECT().Get(context.TODO(), "https://api.weixin.qq.com/sns/auth?access_token=ACCESS_TOKEN&openid=OPENID").Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("wxa06e66cf23dc4370", "1208c7f9e08b4edd26fd86406a5b30aa")
	oa.client = client

	err := oa.Do("ACCESS_TOKEN", CheckAuthToken("OPENID"))

	assert.Nil(t, err)
}

func TestGetAuthUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := NewMockHTTPClient(ctrl)

	client.EXPECT().Get("https://api.weixin.qq.com/sns/userinfo?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN").Return([]byte(`{
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

	oa := New("wxa06e66cf23dc4370", "1208c7f9e08b4edd26fd86406a5b30aa")
	oa.client = client

	receiver := new(AuthUser)
	err := oa.Do("ACCESS_TOKEN", GetAuthUser("OPENID", receiver))

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
	}, receiver)
}
