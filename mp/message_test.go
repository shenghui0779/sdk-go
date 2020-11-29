package mp

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/public"
	"github.com/stretchr/testify/assert"
)

func TestSendUniformMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := public.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/uniform_send?access_token=ACCESS_TOKEN", gomock.AssignableToTypeOf(postBody)).Return([]byte(`{"errcode": 0,"errmsg": "ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(SuperreSolutionResult)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SuperreSolutionByURL("ENCODE_URL", dest))

	assert.Nil(t, err)
	assert.Equal(t, &SuperreSolutionResult{
		MediaID: "6WXsIXkG7lXuDLspD9xfm5dsvHzb0EFl0li6ySxi92ap8Vl3zZoD9DpOyNudeJGB",
	}, dest)
}
