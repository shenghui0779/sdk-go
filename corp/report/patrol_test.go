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

func TestGetPatrolGridInfo(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"grid_list": [
		{
			"grid_id": "grid_id",
			"grid_name": "grid_name",
			"grid_admin": [
				"zhangsan",
				"lisi"
			]
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/report/patrol/get_grid_info?access_token=ACCESS_TOKEN", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultPatrolGridInfo)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetPatrolGridInfo(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultPatrolGridInfo{
		GridList: []*PatrolGrid{
			{
				GridID:    "grid_id",
				GridName:  "grid_name",
				GridAdmin: []string{"zhangsan", "lisi"},
			},
		},
	}, result)
}

func TestGetPatrolCorpStatus(t *testing.T) {
	body := []byte(`{"grid_id":"grid_id"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"processing": 1,
	"added_today": 1,
	"solved_today": 1,
	"total_case": 1,
	"to_be_assigned": 1,
	"total_solved": 1
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/report/patrol/get_corp_status?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultPatrolCorpStatus)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetPatrolCorpStatus("grid_id", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultPatrolCorpStatus{
		Processing:   1,
		AddedToday:   2,
		SolvedToday:  3,
		TotalCase:    4,
		ToBeAssigned: 5,
		TotalSolved:  6,
	}, result)
}

func TestGetPatrolUserStatus(t *testing.T) {
	body := []byte(`{"userid":"zhangsan"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"processing": 1,
	"added_today": 1,
	"solved_today": 1
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/report/patrol/get_user_status?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultPatrolUserStatus)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetPatrolUserStatus("zhangsan", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultPatrolUserStatus{
		Processing:  1,
		AddedToday:  1,
		SolvedToday: 1,
	}, result)
}

func TestGetPatrolCategoryStatistic(t *testing.T) {
	body := []byte(`{"category_id":"category_id"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"dashboard_list": [
		{
			"category_id": "category_id",
			"category_name": "category name",
			"category_level": 1,
			"total_case": 100,
			"total_solved": 100,
			"category_type": 1
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/report/patrol/category_statistic?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultPatrolCategoryStatistic)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetPatrolCategoryStatistic("category_id", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultPatrolCategoryStatistic{
		DashboardList: []*PatrolCategoryStatistic{
			{
				CategoryID:    "category_id",
				CategoryName:  "category name",
				CategoryLevel: 1,
				CategoryType:  100,
				TotalCase:     100,
				TotalSolved:   1,
			},
		},
	}, result)
}

func TestListPatrolOrder(t *testing.T) {
	body := []byte(`{"begin_create_time":12345678,"begin_modify_time":12345678,"cursor":"cursor","limit":20}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"next_cursor": "next_cursor",
	"order_list": [
		{
			"order_id": "order_id",
			"desc": "test",
			"urge_type": 1,
			"case_name": "测试事件",
			"grid_name": "测试网格",
			"grid_id": "grid_id",
			"create_time": 12345678,
			"image_urls": [
				"https://image1.qq.com",
				"https://image2.qq.com"
			],
			"video_media_ids": [
				"mediaid1",
				"mediaid2"
			],
			"location": {
				"name": "测试小区",
				"address": "实例小区，不真实存在，经纬度无意义",
				"latitude": 0,
				"longitude": 0
			},
			"processor_userids": [
				"zhangsan",
				"lisi"
			],
			"process_list": [
				{
					"process_type": 1,
					"solve_userid": "zhangsan",
					"process_desc": "第一个流程",
					"status": 1,
					"solved_time": 123456789,
					"image_urls": [
						"https://image1.qq.com",
						"https://image2.qq.com"
					],
					"video_media_ids": [
						"mediaid1",
						"mediaid2"
					]
				}
			]
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/report/patrol/get_order_list?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsPatrolOrderList{
		BeginCreateTime: 12345678,
		BeginModifyTime: 12345678,
		Cursor:          "cursor",
		Limit:           20,
	}
	result := new(ResultPatrolOrderList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListPatrolOrder(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultPatrolOrderList{
		NextCursor: "next_cursor",
		OrderList: []*PatrolOrder{
			{
				OrderID:       "order_id",
				Desc:          "test",
				UrgeType:      1,
				CaseName:      "测试事件",
				GridName:      "测试网格",
				GridID:        "grid_id",
				CreateTime:    12345678,
				ImageURLs:     []string{"https://image1.qq.com", "https://image2.qq.com"},
				VideoMediaIDs: []string{"mediaid1", "mediaid2"},
				Location: &Location{
					Name:      "测试小区",
					Address:   "实例小区，不真实存在，经纬度无意义",
					Latitude:  0,
					Longitude: 0,
				},
				ProcessorUserIDs: []string{"zhangsan", "lisi"},
				ProcessList: []*Process{
					{
						ProcessType:   1,
						SolveUserID:   "zhangsan",
						ProcessDesc:   "第一个流程",
						Status:        1,
						SolvedTime:    123456789,
						ImageURLs:     []string{"https://image1.qq.com", "https://image2.qq.com"},
						VideoMediaIDs: []string{"mediaid1", "mediaid2"},
					},
				},
			},
		},
	}, result)
}

func TestGetPatrolOrderInfo(t *testing.T) {
	body := []byte(`{"order_id":"order_id"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"order_info": {
		"order_id": "order_id",
		"desc": "test",
		"urge_type": 1,
		"case_name": "测试事件",
		"grid_name": "测试网格",
		"grid_id": "grid_id",
		"create_time": 12345678,
		"image_urls": [
			"https://image1.qq.com",
			"https://image2.qq.com"
		],
		"video_media_ids": [
			"mediaid1",
			"mediaid2"
		],
		"location": {
			"name": "测试小区",
			"address": "实例小区，不真实存在，经纬度无意义",
			"latitude": 0,
			"longitude": 0
		},
		"processor_userids": [
			"zhangsan",
			"lisi"
		],
		"process_list": [
			{
				"process_type": 1,
				"solve_userid": "zhangsan",
				"process_desc": "第一个流程",
				"status": 1,
				"solved_time": 123456789,
				"image_urls": [
					"https://image1.qq.com",
					"https://image2.qq.com"
				],
				"video_media_ids": [
					"mediaid1",
					"mediaid2"
				]
			}
		]
	}
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/report/patrol/get_order_info?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultPatrolOrderInfo)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetPatrolOrderInfo("order_id", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultPatrolOrderInfo{
		OrderInfo: &PatrolOrder{
			OrderID:       "order_id",
			Desc:          "test",
			UrgeType:      1,
			CaseName:      "测试事件",
			GridName:      "测试网格",
			GridID:        "grid_id",
			CreateTime:    12345678,
			ImageURLs:     []string{"https://image1.qq.com", "https://image2.qq.com"},
			VideoMediaIDs: []string{"mediaid1", "mediaid2"},
			Location: &Location{
				Name:      "测试小区",
				Address:   "实例小区，不真实存在，经纬度无意义",
				Latitude:  0,
				Longitude: 0,
			},
			ProcessorUserIDs: []string{},
			ProcessList: []*Process{
				{
					ProcessType:   1,
					SolveUserID:   "zhangsan",
					ProcessDesc:   "第一个流程",
					Status:        1,
					SolvedTime:    123456789,
					ImageURLs:     []string{"https://image1.qq.com", "https://image2.qq.com"},
					VideoMediaIDs: []string{"mediaid1", "mediaid2"},
				},
			},
		},
	}, result)
}
