package school

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
	body := []byte(`{"name":"一年级","parentid":5,"id":2,"type":1,"register_year":2018,"standard_grade":1,"order":1,"department_admins":[{"userid":"zhangsan","type":4,"subject":"语文"},{"userid":"lisi","type":3,"subject":"数学"}]}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/school/department/create?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsDepartmentCreate{
		Name:          "一年级",
		ParentID:      5,
		ID:            2,
		Type:          1,
		RegisterYear:  2018,
		StandardGrade: 1,
		Order:         1,
		DepartmentAdmins: []*DepartmentAdminCreate{
			{
				UserID:  "zhangsan",
				Type:    4,
				Subject: "语文",
			},
			{
				UserID:  "lisi",
				Type:    3,
				Subject: "数学",
			},
		},
	}
	result := new(ResultDepartmentCreate)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", CreateDepartment(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultDepartmentCreate{
		ID: 2,
	}, result)
}

func TestUpdateDeparment(t *testing.T) {
	body := []byte(`{"name":"一年级","parentid":5,"id":2,"register_year":2018,"standard_grade":1,"order":1,"new_id":100,"department_admins":[{"op":0,"userid":"zhangsan","type":3,"subject":"语文"},{"op":1,"userid":"lisi","type":4,"subject":"数学"}]}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/school/department/update?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsDepartmentUpdate{
		Name:          "一年级",
		ParentID:      5,
		ID:            2,
		RegisterYear:  2018,
		StandardGrade: 1,
		Order:         1,
		NewID:         100,
		DepartmentAdmins: []*DepartmentAdminUpdate{
			{
				OP:      0,
				UserID:  "zhangsan",
				Type:    3,
				Subject: "语文",
			},
			{
				OP:      1,
				UserID:  "lisi",
				Type:    4,
				Subject: "数学",
			},
		},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UpdateDeparment(params))

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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/school/department/delete?access_token=ACCESS_TOKEN&id=1", nil).Return(resp, nil)

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
	"departments": [
		{
			"name": "一年级",
			"parentid": 1,
			"id": 2,
			"type": 2,
			"register_year": 2018,
			"standard_grade": 1,
			"order": 1,
			"department_admins": [
				{
					"userid": "zhangsan",
					"type": 1
				},
				{
					"userid": "lisi",
					"type": 2
				}
			],
			"is_graduated": 0
		},
		{
			"name": "一年级一班",
			"parentid": 1,
			"id": 3,
			"type": 1,
			"department_admins": [
				{
					"userid": "zhangsan",
					"type": 3,
					"subject": "语文"
				},
				{
					"userid": "lisi",
					"type": 4,
					"subject": "数学"
				}
			],
			"open_group_chat": 1,
			"group_chat_id": "group_chat_id"
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/school/department/list?access_token=ACCESS_TOKEN&id=1", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultDepartmentList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListDepartment(1, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultDepartmentList{
		Departments: []*Department{
			{
				Name:          "一年级",
				ParentID:      1,
				ID:            2,
				Type:          2,
				RegisterYear:  2018,
				StandardGrade: 1,
				Order:         1,
				DepartmentAdmins: []*DepartmentAdmin{
					{
						UserID: "zhangsan",
						Type:   1,
					},
					{
						UserID: "lisi",
						Type:   2,
					},
				},
			},
			{
				Name:          "一年级一班",
				ParentID:      1,
				ID:            3,
				Type:          1,
				OpenGroupChat: 1,
				GroupChatID:   "group_chat_id",
				DepartmentAdmins: []*DepartmentAdmin{
					{
						UserID:  "zhangsan",
						Type:    3,
						Subject: "语文",
					},
					{
						UserID:  "lisi",
						Type:    4,
						Subject: "数学",
					},
				},
			},
		},
	}, result)
}

func TestSetUpgradeInfo(t *testing.T) {
	body := []byte(`{"upgrade_time":1594090969,"upgrade_switch":2}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"next_upgrade_time": 1625587200,
	"errcode": 0,
	"errmsg": "ok"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/school/set_upgrade_info?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultUpgradeInfoSet)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SetUpgradeInfo(1594090969, 2, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultUpgradeInfoSet{
		NextUpgradeTime: 1625587200,
	}, result)
}
