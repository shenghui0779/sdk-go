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

func TestGetUser(t *testing.T) {
	resp := []byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "user_type": 1,
    "student": {
        "student_userid": "zhangsan",
        "name": "张三",
        "department": [
            1,
            2
        ],
        "parents": [
            {
                "parent_userid": "zhangsan_parent1",
                "relation": "爸爸",
                "mobile": "18000000000",
                "is_subscribe": 1,
                "external_userid": "xxxxx"
            },
            {
                "parent_userid": "zhangsan_parent2",
                "relation": "妈妈",
                "mobile": "18000000001",
                "is_subscribe": 0
            }
        ]
    },
    "parent": {
        "parent_userid": "zhangsan_parent2",
        "mobile": "18000000003",
        "is_subscribe": 1,
        "external_userid": "xxxxx",
        "children": [
            {
                "student_userid": "zhangsan",
                "relation": "妈妈"
            },
            {
                "student_userid": "lisi",
                "relation": "伯母"
            }
        ]
    }
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/school/user/get?access_token=ACCESS_TOKEN&userid=USERID", nil).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultUserGet)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetUser("USERID", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultUserGet{
		UserType: 1,
		Student: &Student{
			StudentUserID: "zhangsan",
			Name:          "张三",
			Department:    []int64{1, 2},
			Parents: []*Parent{
				{
					ParentUserID:   "zhangsan_parent1",
					Relation:       "爸爸",
					Mobile:         "18000000000",
					IsSubscribe:    1,
					ExternalUserID: "xxxxx",
				},
				{
					ParentUserID: "zhangsan_parent2",
					Relation:     "妈妈",
					Mobile:       "18000000001",
					IsSubscribe:  0,
				},
			},
		},
		Parent: &Parent{
			ParentUserID:   "zhangsan_parent2",
			Mobile:         "18000000003",
			IsSubscribe:    1,
			ExternalUserID: "xxxxx",
			Children: []*Child{
				{
					StudentUserID: "zhangsan",
					Relation:      "妈妈",
				},
				{
					StudentUserID: "lisi",
					Relation:      "伯母",
				},
			},
		},
	}, result)
}

func TestListUser(t *testing.T) {
	resp := []byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "students": [
        {
            "student_userid": "zhangsan",
            "name": "张三",
            "department": [
                1,
                2
            ],
            "parents": [
                {
                    "parent_userid": "zhangsan_parent1",
                    "relation": "爸爸",
                    "mobile": "18000000001",
                    "is_subscribe": 1,
                    "external_userid": "xxx"
                },
                {
                    "parent_userid": "zhangsan_parent2",
                    "relation": "妈妈",
                    "mobile": "18000000002",
                    "is_subscribe": 0
                }
            ]
        },
        {
            "student_userid": "lisi",
            "name": "李四",
            "department": [
                4,
                5
            ],
            "parents": [
                {
                    "parent_userid": "lisi_parent1",
                    "relation": "爷爷",
                    "mobile": "18000000003",
                    "is_subscribe": 1,
                    "external_userid": "xxx"
                },
                {
                    "parent_userid": "lisi_parent2",
                    "relation": "妈妈",
                    "mobile": "18000000004",
                    "is_subscribe": 1,
                    "external_userid": "xxx"
                }
            ]
        }
    ]
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/school/user/list?access_token=ACCESS_TOKEN&department_id=1&fetch_child=0", nil).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultUserList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListUser(1, 0, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultUserList{
		Students: []*Student{
			{
				StudentUserID: "zhangsan",
				Name:          "张三",
				Department:    []int64{1, 2},
				Parents: []*Parent{
					{
						ParentUserID:   "zhangsan_parent1",
						Relation:       "爸爸",
						Mobile:         "18000000001",
						IsSubscribe:    1,
						ExternalUserID: "xxx",
					},
					{
						ParentUserID: "zhangsan_parent2",
						Relation:     "妈妈",
						Mobile:       "18000000002",
						IsSubscribe:  0,
					},
				},
			},
			{
				StudentUserID: "lisi",
				Name:          "李四",
				Department:    []int64{4, 5},
				Parents: []*Parent{
					{
						ParentUserID:   "lisi_parent1",
						Relation:       "爷爷",
						Mobile:         "18000000003",
						IsSubscribe:    1,
						ExternalUserID: "xxx",
					},
					{
						ParentUserID:   "lisi_parent2",
						Relation:       "妈妈",
						Mobile:         "18000000004",
						IsSubscribe:    1,
						ExternalUserID: "xxx",
					},
				},
			},
		},
	}, result)
}

func TestSetArchSyncMode(t *testing.T) {
	body := []byte(`{"arch_sync_mode":1}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/school/set_arch_sync_mode?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SetArchSyncMode(1))

	assert.Nil(t, err)
}

func TestListParent(t *testing.T) {
	resp := []byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "parents": [
        {
            "parent_userid": "zhangsan_parent",
            "mobile": "18900000000",
            "is_subscribe": 1,
            "external_userid": "xxx",
            "children": [
                {
                    "student_userid": "zhangsan",
                    "relation": "爸爸",
                    "name": "张三"
                }
            ]
        },
        {
            "parent_userid": "lisi_parent",
            "mobile": "18900000001",
            "is_subscribe": 0,
            "children": [
                {
                    "student_userid": "lisi",
                    "relation": "妈妈",
                    "name": "李四"
                }
            ]
        }
    ]
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/school/user/list_parent?access_token=ACCESS_TOKEN&department_id=1", nil).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultParentList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListParent(1, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultParentList{
		Parents: []*Parent{
			{
				ParentUserID:   "zhangsan_parent",
				Mobile:         "18900000000",
				IsSubscribe:    1,
				ExternalUserID: "xxx",
				Children: []*Child{
					{
						StudentUserID: "zhangsan",
						Relation:      "爸爸",
						Name:          "张三",
					},
				},
			},
			{
				ParentUserID: "lisi_parent",
				Mobile:       "18900000001",
				IsSubscribe:  0,
				Children: []*Child{
					{
						StudentUserID: "lisi",
						Relation:      "妈妈",
						Name:          "李四",
					},
				},
			},
		},
	}, result)
}
