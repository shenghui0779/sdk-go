package oa

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/corp"
	"github.com/shenghui0779/gochat/mock"
)

func TestCallPstncc(t *testing.T) {
	body := []byte(`{"callee_userid":["james","paul"]}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"states": [
		{
			"code": 0,
			"callid": "6-20190510201844181887818-4d0251082406000-out",
			"userid": "james"
		},
		{
			"code": 0,
			"callid": "6-20190510201844181887818-4d025109f806000-out",
			"userid": "paul"
		}
	]
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/pstncc/call?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultPstnccCall)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", CallPstncc([]string{"james", "paul"}, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultPstnccCall{
		States: []*PstnccState{
			{
				Code:   0,
				CallID: "6-20190510201844181887818-4d0251082406000-out",
				UserID: "james",
			},
			{
				Code:   0,
				CallID: "6-20190510201844181887818-4d025109f806000-out",
				UserID: "paul",
			},
		},
	}, result)
}

func TestGetPstnccStates(t *testing.T) {
	body := []byte(`{"callee_userid":"james","callid":"6-20190510201844181887818-4d0251082406000-out"}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"istalked": 1,
	"calltime": 1557306531,
	"talktime": 2,
	"reason": 0
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/pstncc/getstates?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultPstnccStates)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetPstnccStates("james", "6-20190510201844181887818-4d0251082406000-out", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultPstnccStates{
		IsTalked: 1,
		CallTime: 1557306531,
		TalkTime: 2,
		Reason:   0,
	}, result)
}
