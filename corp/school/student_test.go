package school

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/corp"
	"github.com/shenghui0779/gochat/mock"
)

func TestCreateStudent(t *testing.T) {
	body := []byte(`{"student_userid":"zhangsan","name":"张三","department":[1,2]}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/school/user/create_student?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	params := &ParamsStudentCreate{
		StudentUserID: "zhangsan",
		Name:          "张三",
		Department:    []int64{1, 2},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", CreateStudent(params))

	assert.Nil(t, err)
}

func TestUpdateStudent(t *testing.T) {
	body := []byte(`{"student_userid":"zhangsan","new_student_userid":"NEW_ID","name":"张三","department":[1,2]}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/school/user/update_student?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	params := &ParamsStudentUpdate{
		StudentUserID:    "zhangsan",
		NewStudentUserID: "NEW_ID",
		Name:             "张三",
		Department:       []int64{1, 2},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UpdateStudent(params))

	assert.Nil(t, err)
}

func TestDeleteStudent(t *testing.T) {
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/school/user/delete_student?access_token=ACCESS_TOKEN&userid=USERID", nil).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", DeleteStudent("USERID"))

	assert.Nil(t, err)
}

func TestBatchCreateStudent(t *testing.T) {
	body := []byte(`{"students":[{"student_userid":"zhangsan","name":"张三","department":[1,2]},{"student_userid":"lisi","name":"李四","department":[3,4]}]}`)
	resp := []byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "result_list": [
        {
            "student_userid": "zhangsan",
            "errcode": 1,
            "errmsg": "invalid student_userid: zhangsan"
        }
    ]
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/school/user/batch_create_student?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	params := &ParamsStudentBatchCreate{
		Students: []*ParamsStudentCreate{
			{
				StudentUserID: "zhangsan",
				Name:          "张三",
				Department:    []int64{1, 2},
			},
			{
				StudentUserID: "lisi",
				Name:          "李四",
				Department:    []int64{3, 4},
			},
		},
	}
	result := new(ResultStudentBatchCreate)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", BatchCreateStudent(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultStudentBatchCreate{
		ResultList: []*StudentErrRet{
			{
				StudentUserID: "zhangsan",
				ErrCode:       1,
				ErrMsg:        "invalid student_userid: zhangsan",
			},
		},
	}, result)
}

func TestBatchUpdateStudent(t *testing.T) {
	body := []byte(`{"students":[{"student_userid":"zhangsan","new_student_userid":"zhangsan_new","name":"张三","department":[1,2]},{"student_userid":"lisi","name":"李四","department":[3,4]}]}`)
	resp := []byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "result_list": [
        {
            "student_userid": "zhangsan",
            "errcode": 1,
            "errmsg": "invalid student_userid: zhangsan"
        }
    ]
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/school/user/batch_update_student?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	params := &ParamsStudentBatchUpdate{
		Students: []*ParamsStudentUpdate{
			{
				StudentUserID:    "zhangsan",
				NewStudentUserID: "zhangsan_new",
				Name:             "张三",
				Department:       []int64{1, 2},
			},
			{
				StudentUserID: "lisi",
				Name:          "李四",
				Department:    []int64{3, 4},
			},
		},
	}
	result := new(ResultStudentBatchUpdate)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", BatchUpdateStudent(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultStudentBatchUpdate{
		ResultList: []*StudentErrRet{
			{
				StudentUserID: "zhangsan",
				ErrCode:       1,
				ErrMsg:        "invalid student_userid: zhangsan",
			},
		},
	}, result)
}

func TestBatchDeleteStudent(t *testing.T) {
	body := []byte(`{"useridlist":["zhangsan","lisi"]}`)
	resp := []byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "result_list": [
        {
            "student_userid": "lisi",
            "errcode": 1111,
            "errmsg": "userid not found"
        }
    ]
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/school/user/batch_delete_student?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	userIDs := []string{"zhangsan", "lisi"}

	result := new(ResultStudentBatchDelete)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", BatchDeleteStudent(userIDs, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultStudentBatchDelete{
		ResultList: []*StudentErrRet{
			{
				StudentUserID: "lisi",
				ErrCode:       1111,
				ErrMsg:        "userid not found",
			},
		},
	}, result)
}
