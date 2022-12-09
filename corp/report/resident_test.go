package report

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/corp"
	"github.com/shenghui0779/gochat/mock"
)

func TestGetResidentGridInfo(t *testing.T) {
	resp := []byte(`{
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
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/report/resident/get_grid_info?access_token=ACCESS_TOKEN", nil).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultResidentGridInfo)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetResidentGridInfo(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultResidentGridInfo{
		GridList: []*ResidentGrid{
			{
				GridID:    "grid_id",
				GridName:  "grid_name",
				GridAdmin: []string{"zhangsan", "lisi"},
			},
		},
	}, result)
}

func TestGetResidentCorpStatus(t *testing.T) {
	body := []byte(`{"grid_id":"grid_id"}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"processing": 1,
	"added_today": 1,
	"solved_today": 1,
	"pending": 10,
	"total_case": 1,
	"total_accepted": 1,
	"total_solved": 1
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/report/resident/get_corp_status?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultResidentCorpStatus)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetResidentCorpStatus("grid_id", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultResidentCorpStatus{
		Processing:    1,
		AddedToday:    1,
		SolvedToday:   1,
		Pending:       10,
		TotalCase:     1,
		TotalAccepted: 1,
		TotalSolved:   1,
	}, result)
}

func TestGetResidentUserStatus(t *testing.T) {
	body := []byte(`{"userid":"zhangsan"}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"processing": 1,
	"added_today": 1,
	"solved_today": 1,
	"pending": 1
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/report/resident/get_user_status?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultResidentUserStatus)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetResidentUserStatus("zhangsan", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultResidentUserStatus{
		Processing:  1,
		AddedToday:  1,
		SolvedToday: 1,
		Pending:     1,
	}, result)
}

func TestGetResidentCategoryStatistic(t *testing.T) {
	body := []byte(`{"category_id":"category_id"}`)
	resp := []byte(`{
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
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/report/resident/category_statistic?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultResidentCategoryStatistic)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetResidentCategoryStatistic("category_id", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultResidentCategoryStatistic{
		DashboardList: []*ResidentCategoryStatistic{
			{
				CategoryID:    "category_id",
				CategoryName:  "category name",
				CategoryLevel: 1,
				CategoryType:  1,
				TotalCase:     100,
				TotalSolved:   100,
			},
		},
	}, result)
}

func TestListResidentOrder(t *testing.T) {
	body := []byte(`{"begin_create_time":12345678,"begin_modify_time":12345678,"cursor":"cursor","limit":20}`)
	resp := []byte(`{
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
			"reporter_name": "上报人名称",
			"reporter_mobile": "上报人手机",
			"unionid": "上报人unionid",
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
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/report/resident/get_order_list?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultResidentOrderList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListResidentOrder(12345678, 12345678, "cursor", 20, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultResidentOrderList{
		NextCursor: "next_cursor",
		OrderList: []*ResidentOrder{
			{
				OrderID:        "order_id",
				Desc:           "test",
				UrgeType:       1,
				CaseName:       "测试事件",
				GridName:       "测试网格",
				GridID:         "grid_id",
				ReporterName:   "上报人名称",
				ReporterMobile: "上报人手机",
				UnionID:        "上报人unionid",
				CreateTime:     12345678,
				ImageURLs:      []string{"https://image1.qq.com", "https://image2.qq.com"},
				VideoMediaIDs:  []string{"mediaid1", "mediaid2"},
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

func TestGetResidentOrderInfo(t *testing.T) {
	body := []byte(`{"order_id":"order_id"}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"order_info": {
		"order_id": "order_id",
		"desc": "test",
		"urge_type": 1,
		"case_name": "测试事件",
		"grid_name": "测试网格",
		"grid_id": "grid_id",
		"reporter_name": "上报人名称",
		"reporter_mobile": "上报人手机",
		"unionid": "上报人unionid",
		"image_urls": [
			"https://image1.qq.com",
			"https://image2.qq.com"
		],
		"video_media_ids": [
			"mediaid1",
			"mediaid2"
		],
		"create_time": 12345678,
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
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/report/resident/get_order_info?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultResidentOrderInfo)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetResidentOrderInfo("order_id", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultResidentOrderInfo{
		OrderInfo: &ResidentOrder{
			OrderID:        "order_id",
			Desc:           "test",
			UrgeType:       1,
			CaseName:       "测试事件",
			GridName:       "测试网格",
			GridID:         "grid_id",
			ReporterName:   "上报人名称",
			ReporterMobile: "上报人手机",
			UnionID:        "上报人unionid",
			CreateTime:     12345678,
			ImageURLs:      []string{"https://image1.qq.com", "https://image2.qq.com"},
			VideoMediaIDs:  []string{"mediaid1", "mediaid2"},
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
	}, result)
}
