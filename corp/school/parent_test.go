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

func TestCreateParent(t *testing.T) {
	body := []byte(`{"parent_userid":"zhangsan_parent_userid","mobile":"10000000000","to_invite":false,"children":[{"student_userid":"zhangsan","relation":"爸爸"},{"student_userid":"lisi","relation":"伯父"}]}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/school/user/create_parent?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsParentCreate{
		ParentUserID: "zhangsan_parent_userid",
		Mobile:       "10000000000",
		ToInvite:     new(bool),
		Children: []*ParamsChild{
			{
				StudentUserID: "zhangsan",
				Relation:      "爸爸",
			},
			{
				StudentUserID: "lisi",
				Relation:      "伯父",
			},
		},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", CreateParent(params))

	assert.Nil(t, err)
}

func TestUpdateParent(t *testing.T) {
	body := []byte(`{"parent_userid":"zhangsan_parent_userid","new_parent_userid":"NEW_ID","mobile":"18000000000","children":[{"student_userid":"zhangsan","relation":"爸爸"},{"student_userid":"lisi","relation":"伯父"}]}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/school/user/update_parent?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsParentUpdate{
		ParentUserID:    "zhangsan_parent_userid",
		NewParentUserID: "NEW_ID",
		Mobile:          "18000000000",
		Children: []*ParamsChild{
			{
				StudentUserID: "zhangsan",
				Relation:      "爸爸",
			},
			{
				StudentUserID: "lisi",
				Relation:      "伯父",
			},
		},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UpdateParent(params))

	assert.Nil(t, err)
}

func TestDeleteParent(t *testing.T) {
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/school/user/delete_parent?access_token=ACCESS_TOKEN&userid=USERID", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", DeleteParent("USERID"))

	assert.Nil(t, err)
}

func TestBatchCreateParent(t *testing.T) {
	body := []byte(`{"parents":[{"parent_userid":"zhangsan_parent_userid","mobile":"18000000000","to_invite":false,"children":[{"student_userid":"zhangsan","relation":"爸爸"},{"student_userid":"lisi","relation":"伯父"}]},{"parent_userid":"lisi_parent_userid","mobile":"18000000001","children":[{"student_userid":"lisi","relation":"爸爸"},{"student_userid":"zhangsan","relation":"伯父"}]}]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "result_list": [
        {
            "parent_userid": "lisi_parent_userid",
            "errcode": 1,
            "errmsg": "invalid parent_userid: lisi_parent_userid"
        }
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/school/user/batch_create_parent?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsParentBatchCreate{
		Parents: []*ParamsParentCreate{
			{
				ParentUserID: "zhangsan_parent_userid",
				Mobile:       "18000000000",
				ToInvite:     new(bool),
				Children: []*ParamsChild{
					{
						StudentUserID: "zhangsan",
						Relation:      "爸爸",
					},
					{
						StudentUserID: "lisi",
						Relation:      "伯父",
					},
				},
			},
			{
				ParentUserID: "lisi_parent_userid",
				Mobile:       "18000000001",
				Children: []*ParamsChild{
					{
						StudentUserID: "lisi",
						Relation:      "爸爸",
					},
					{
						StudentUserID: "zhangsan",
						Relation:      "伯父",
					},
				},
			},
		},
	}
	result := new(ResultParentBatchCreate)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", BatchCreateParent(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultParentBatchCreate{
		ResultList: []*ParentErrResult{
			{
				ParentUserID: "lisi_parent_userid",
				ErrCode:      1,
				ErrMsg:       "invalid parent_userid: lisi_parent_userid",
			},
		},
	}, result)
}

func TestBatchUpdateParent(t *testing.T) {
	body := []byte(`{"parents":[{"parent_userid":"zhangsan_baba","new_parent_userid":"zhangsan_baba_new","mobile":"10000000000","children":[{"student_userid":"zhangsan","relation":"爸爸"}]},{"parent_userid":"lisi_mama","mobile":"10000000001","children":[{"student_userid":"lisi","relation":"妈妈"}]}]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "result_list": [
        {
            "parent_userid": "lisi_parent_userid",
            "errcode": 1,
            "errmsg": "invalid parent_userid: lisi_parent_userid"
        }
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/school/user/batch_update_parent?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsParentBatchUpdate{
		Parents: []*ParamsParentUpdate{
			{
				ParentUserID:    "zhangsan_baba",
				NewParentUserID: "zhangsan_baba_new",
				Mobile:          "10000000000",
				Children: []*ParamsChild{
					{
						StudentUserID: "zhangsan",
						Relation:      "爸爸",
					},
				},
			},
			{
				ParentUserID: "lisi_mama",
				Mobile:       "10000000001",
				Children: []*ParamsChild{
					{
						StudentUserID: "lisi",
						Relation:      "妈妈",
					},
				},
			},
		},
	}
	result := new(ResultParentBatchUpdate)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", BatchUpdateParent(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultParentBatchUpdate{
		ResultList: []*ParentErrResult{
			{
				ParentUserID: "lisi_parent_userid",
				ErrCode:      1,
				ErrMsg:       "invalid parent_userid: lisi_parent_userid",
			},
		},
	}, result)
}

func TestBatchDeleteParent(t *testing.T) {
	body := []byte(`{"useridlist":["zhangsan","lisi"]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"result_list": [
		{
			"parent_userid": "lisi",
			"errcode": 1111,
			"errmsg": "userid not found"
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/school/user/batch_delete_parent?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	userIDs := []string{"zhangsan", "lisi"}

	result := new(ResultParentBatchDelete)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", BatchDeleteParent(userIDs, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultParentBatchDelete{
		ResultList: []*ParentErrResult{
			{
				ParentUserID: "lisi",
				ErrCode:      1111,
				ErrMsg:       "userid not found",
			},
		},
	}, result)
}
