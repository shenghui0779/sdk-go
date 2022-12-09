package linkedcorp

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/corp"
	"github.com/shenghui0779/gochat/mock"
)

func TestListAgentPerm(t *testing.T) {
	resp := []byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "userids": [
        "CORPID/USERID"
    ],
    "department_ids": [
        "LINKEDID/DEPARTMENTID"
    ]
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/agent/get_perm_list?access_token=ACCESS_TOKEN", nil).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultAgentPermList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListAgentPerm(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultAgentPermList{
		UserIDs:       []string{"CORPID/USERID"},
		DepartmentIDs: []string{"LINKEDID/DEPARTMENTID"},
	}, result)
}
