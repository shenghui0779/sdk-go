package offia

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/mock"
)

func TestCheckOAuthToken(t *testing.T) {
	resp := []byte(`{"errcode":0,"errmsg":"ok"}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://api.weixin.qq.com/sns/auth?access_token=ACCESS_TOKEN&openid=OPENID", nil).Return(resp, nil)

	oa := New("APPID", "APPSECRET", WithMockClient(client))

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", CheckOAuthToken("OPENID"))

	assert.Nil(t, err)
}

func TestGetOAuthUser(t *testing.T) {
	resp := []byte(`{
	"openid": "OPENID",
	"nickname": "NICKNAME",
	"sex": 1,
	"province": "PROVINCE",
	"city": "CITY",
	"country": "COUNTRY",
	"headimgurl": "https://thirdwx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/46",
	"privilege": ["PRIVILEGE1", "PRIVILEGE2"],
	"unionid": "o6_bmasdasdsad6_2sgVt7hMZOPfL"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://api.weixin.qq.com/sns/userinfo?access_token=ACCESS_TOKEN&lang=zh_CN&openid=OPENID", nil).Return(resp, nil)

	oa := New("APPID", "APPSECRET", WithMockClient(client))

	result := new(ResultOAuthUser)
	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetOAuthUser("OPENID", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultOAuthUser{
		OpenID:     "OPENID",
		UnionID:    "o6_bmasdasdsad6_2sgVt7hMZOPfL",
		Nickname:   "NICKNAME",
		Sex:        1,
		Province:   "PROVINCE",
		City:       "CITY",
		Country:    "COUNTRY",
		HeadImgURL: "https://thirdwx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/46",
		Privilege:  []string{"PRIVILEGE1", "PRIVILEGE2"},
	}, result)
}

func TestGetApiTicket(t *testing.T) {
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"ticket": "bxLdikRXVbTPdHSM05e5u5sUoXNKd8-41ZO3MhKoyN5OfkWITDGgnr2fwJ0m9E8NYzWKVZvdVtaUgWvsdshFKA",
	"expires_in": 7200
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=ACCESS_TOKEN&type=jsapi", nil).Return(resp, nil)

	oa := New("APPID", "APPSECRET", WithMockClient(client))

	result := new(ResultApiTicket)
	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetApiTicket(JSAPITicket, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultApiTicket{
		Ticket:    "bxLdikRXVbTPdHSM05e5u5sUoXNKd8-41ZO3MhKoyN5OfkWITDGgnr2fwJ0m9E8NYzWKVZvdVtaUgWvsdshFKA",
		ExpiresIn: 7200,
	}, result)
}
