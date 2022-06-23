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

func TestListDeparment(t *testing.T) {
	body := []byte(`{"department_id":"LINKEDID/DEPARTMENTID"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "department_list": [
        {
            "department_id": "1",
            "department_name": "测试部门1",
            "parentid": "0",
            "order": 100000000
        },
        {
            "department_id": "2",
            "department_name": "测试部门2",
            "parentid": "1",
            "order": 99999999
        }
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/department/list?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultDepartmentList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListDeparment("LINKEDID", "DEPARTMENTID", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultDepartmentList{
		DepartmentList: []*DepartmentListData{
			{
				DepartmentID:   "1",
				DepartmentName: "测试部门1",
				ParentID:       "0",
				Order:          100000000,
			},
			{
				DepartmentID:   "2",
				DepartmentName: "测试部门2",
				ParentID:       "1",
				Order:          99999999,
			},
		},
	}, result)
}
