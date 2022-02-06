package report

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

func TestAddGrid(t *testing.T) {
	body := []byte(`{"grid_name":"grid_name","grid_parent_id":"grid_id","grid_admin":["zhangsan"],"grid_member":["lisi","invaliduser"]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"grid_id": "grid_id",
	"invalid_userids": [
		"invaliduser"
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/report/grid/add?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsGridAdd{
		GridName:     "grid_name",
		GridParentID: "grid_id",
		GridAdmin:    []string{"zhangsan"},
		GridMember:   []string{"lisi", "invaliduser"},
	}
	result := new(ResultGridAdd)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", AddGrid(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultGridAdd{
		GridID:         "grid_id",
		InvalidUserIDs: []string{"invaliduser"},
	}, result)
}

func TestUpdateGrid(t *testing.T) {
	body := []byte(`{"grid_id":"grid_id","grid_name":"grid_name","grid_parent_id":"grid_id","grid_admin":["zhangsan"],"grid_member":["lisi","invaliduser"]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invalid_userids": [
		"invaliduser"
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/report/grid/update?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsGridUpdate{
		GridID:       "grid_id",
		GridName:     "grid_name",
		GridParentID: "grid_id",
		GridAdmin:    []string{"zhangsan"},
		GridMember:   []string{"lisi", "invaliduser"},
	}
	result := new(ResultGridUpdate)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UpdateGrid(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultGridUpdate{
		InvalidUserIDs: []string{"invaliduser"},
	}, result)
}

func TestDeleteGrid(t *testing.T) {
	body := []byte(`{"grid_id":"grid_id"}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/report/grid/delete?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", DeleteGrid("grid_id"))

	assert.Nil(t, err)
}

func TestListGrid(t *testing.T) {
	body := []byte(`{"grid_id":"grid_id"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"grid_list": [
		{
			"grid_id": "grid_id",
			"grid_name": "grid_name",
			"grid_parent_id": "grid_id",
			"grid_admin": [
				"zhangsan"
			],
			"grid_member": [
				"lisi",
				"11111"
			]
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/report/grid/list?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultGridList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListGrid("grid_id", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultGridList{
		GridList: []*Grid{
			{
				GridID:       "grid_id",
				GridName:     "grid_name",
				GridParentID: "grid_id",
				GridAdmin:    []string{"zhangsan"},
				GridMember:   []string{"lisi", "11111"},
			},
		},
	}, result)
}

func TestGetUserGridInfo(t *testing.T) {
	body := []byte(`{"userid":"zhangsan"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"manage_grids": [
		{
			"grid_id": "grid_id1",
			"grid_name": "grid_name1"
		},
		{
			"grid_id": "grid_id2",
			"grid_name": "grid_name2"
		}
	],
	"joined_grids": [
		{
			"grid_id": "grid_id1",
			"grid_name": "grid_name1"
		},
		{
			"grid_id": "grid_id2",
			"grid_name": "grid_name2"
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/report/grid/get_user_grid_info?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultUserGridInfo)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetUserGridInfo("zhangsan", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultUserGridInfo{
		ManageGrids: []*UserGridInfo{
			{
				GridID:   "grid_id1",
				GridName: "grid_name1",
			},
			{
				GridID:   "grid_id2",
				GridName: "grid_name2",
			},
		},
		JoinedGrids: []*UserGridInfo{
			{
				GridID:   "grid_id1",
				GridName: "grid_name1",
			},
			{
				GridID:   "grid_id2",
				GridName: "grid_name2",
			},
		},
	}, result)
}
