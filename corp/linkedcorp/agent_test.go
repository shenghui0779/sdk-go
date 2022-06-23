package linkedcorp

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/chenghonour/gochat/corp"
	"github.com/chenghonour/gochat/mock"
	"github.com/chenghonour/gochat/wx"
)

func TestListAgentPerm(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "userids": [
        "CORPID/USERID"
    ],
    "department_ids": [
        "LINKEDID/DEPARTMENTID"
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/agent/get_perm_list?access_token=ACCESS_TOKEN", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultAgentPermList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListAgentPerm(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultAgentPermList{
		UserIDs:       []string{"CORPID/USERID"},
		DepartmentIDs: []string{"LINKEDID/DEPARTMENTID"},
	}, result)
}
