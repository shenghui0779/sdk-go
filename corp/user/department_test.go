package user

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/corp"
	"github.com/shenghui0779/gochat/mock"
	"github.com/shenghui0779/gochat/wx"
)

func TestCreateDepartment(t *testing.T) {
	body := []byte(`{"name":"广州研发中心","name_en":"RDGZ","parentid":1,"order":1}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "created",
    "id": 2
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/department/create?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsDepartmentCreate{
		Name:     "广州研发中心",
		NameEN:   "RDGZ",
		ParentID: 1,
		Order:    1,
	}
	result := new(ResultDepartmentCreate)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", CreateDepartment(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultDepartmentCreate{
		ID: 2,
	}, result)
}

func TestUpdateDepartment(t *testing.T) {
	body := []byte(`{"id":2,"name":"广州研发中心","name_en":"RDGZ","parentid":1,"order":1}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/department/update?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsDepartmentUpdate{
		ID:       2,
		Name:     "广州研发中心",
		NameEN:   "RDGZ",
		ParentID: 1,
		Order:    1,
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UpdateDepartment(params))

	assert.Nil(t, err)
}

func TestDeleteDepartment(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/department/delete?access_token=ACCESS_TOKEN&id=1", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", DeleteDepartment(1))

	assert.Nil(t, err)
}

func TestListDepartment(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "department": [
        {
            "id": 2,
            "name": "广州研发中心",
            "name_en": "RDGZ",
            "department_leader": [
                "zhangsan",
                "lisi"
            ],
            "parentid": 1,
            "order": 10
        },
        {
            "id": 3,
            "name": "邮箱产品部",
            "name_en": "mail",
            "department_leader": [
                "lisi",
                "wangwu"
            ],
            "parentid": 2,
            "order": 40
        }
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/department/list?access_token=ACCESS_TOKEN&id=1", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultDepartmentList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListDepartment(1, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultDepartmentList{
		Department: []*Department{
			{
				ID:               2,
				Name:             "广州研发中心",
				NameEN:           "RDGZ",
				DepartmentLeader: []string{"zhangsan", "lisi"},
				ParentID:         1,
				Order:            10,
			},
			{
				ID:               3,
				Name:             "邮箱产品部",
				NameEN:           "mail",
				DepartmentLeader: []string{"lisi", "wangwu"},
				ParentID:         2,
				Order:            40,
			},
		},
	}, result)
}

func TestListSimpleDepartment(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "department_id": [
        {
            "id": 2,
            "parentid": 1,
            "order": 10
        },
        {
            "id": 3,
            "parentid": 2,
            "order": 40
        }
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/department/simplelist?access_token=ACCESS_TOKEN&id=1", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultSimpleDepartmentList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListSimpleDepartment(1, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultSimpleDepartmentList{
		DepartmentID: []*SimpleDepartment{
			{
				ID:       2,
				ParentID: 1,
				Order:    10,
			},
			{
				ID:       3,
				ParentID: 2,
				Order:    40,
			},
		},
	}, result)
}

func TestGetDepartment(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "department": {
        "id": 2,
        "name": "广州研发中心",
        "name_en": "RDGZ",
        "department_leader": [
            "zhangsan",
            "lisi"
        ],
        "parentid": 1,
        "order": 10
    }
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/department/get?access_token=ACCESS_TOKEN&id=1", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultDepartmentGet)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetDepartment(1, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultDepartmentGet{
		Department: &Department{
			ID:               2,
			Name:             "广州研发中心",
			NameEN:           "RDGZ",
			DepartmentLeader: []string{"zhangsan", "lisi"},
			ParentID:         1,
			Order:            10,
		},
	}, result)
}
