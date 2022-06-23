package externalcontact

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

func TestListCustomerStrategy(t *testing.T) {
	body := []byte(`{"cursor":"CURSOR","limit":1000}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"strategy": [
		{
			"strategy_id": 1
		},
		{
			"strategy_id": 2
		}
	],
	"next_cursor": "NEXT_CURSOR"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_strategy/list?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultCustomerStrategyList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListCustomerStrategy("CURSOR", 1000, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultCustomerStrategyList{
		Strategy: []*CustomerStrategyListData{
			{
				StrategyID: 1,
			},
			{
				StrategyID: 2,
			},
		},
		NextCursor: "NEXT_CURSOR",
	}, result)
}

func TestGetCustomerStrategy(t *testing.T) {
	body := []byte(`{"strategy_id":1}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"strategy": {
		"strategy_id": 1,
		"parent_id": 0,
		"strategy_name": "NAME",
		"create_time": 1557838797,
		"admin_list": [
			"zhangsan",
			"lisi"
		],
		"privilege": {
			"view_customer_list": true,
			"view_customer_data": true,
			"view_room_list": true,
			"contact_me": true,
			"join_room": true,
			"share_customer": false,
			"oper_resign_customer": true,
			"oper_resign_group": true,
			"send_customer_msg": true,
			"edit_welcome_msg": true,
			"view_behavior_data": true,
			"view_room_data": true,
			"send_group_msg": true,
			"room_deduplication": true,
			"rapid_reply": true,
			"onjob_customer_transfer": true,
			"edit_anti_spam_rule": true,
			"export_customer_list": true,
			"export_customer_data": true,
			"export_customer_group_list": true,
			"manage_customer_tag": true
		}
	}
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_strategy/get?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultCustomerStrategyGet)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetCustomerStrategy(1, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultCustomerStrategyGet{
		Strategy: &CustomerStrategy{
			StrategyID:   1,
			ParentID:     0,
			StrategyName: "NAME",
			CreateTime:   1557838797,
			AdminList:    []string{"zhangsan", "lisi"},
			Privilege: &CustomerStrategyPrivilege{
				ViewCustomerList:        true,
				ViewCustomerData:        true,
				ViewRoomList:            true,
				ContactMe:               true,
				JoinRoom:                true,
				ShareCustomer:           false,
				OperResignCustomer:      true,
				OperResignGroup:         true,
				SendCustomerMsg:         true,
				EditWelcomeMsg:          true,
				ViewBehaviorData:        true,
				ViewRoomData:            true,
				SendGroupMsg:            true,
				RoomDeduplication:       true,
				RapidReply:              true,
				OnjobCustomerTransfer:   true,
				EditAntiSpamRule:        true,
				ExportCustomerList:      true,
				ExportCustomerData:      true,
				ExportCustomerGroupList: true,
				ManageCustomerTag:       true,
			},
		},
	}, result)
}

func TestCreateCustomerStrategy(t *testing.T) {
	body := []byte(`{"strategy_name":"NAME","admin_list":["zhangsan","lisi"],"privilege":{"view_customer_list":true,"view_customer_data":true,"view_room_list":true,"contact_me":true,"join_room":true,"share_customer":false,"oper_resign_customer":true,"oper_resign_group":true,"send_customer_msg":true,"edit_welcome_msg":true,"view_behavior_data":true,"view_room_data":true,"send_group_msg":true,"room_deduplication":true,"rapid_reply":true,"onjob_customer_transfer":true,"edit_anti_spam_rule":true,"export_customer_list":true,"export_customer_data":true,"export_customer_group_list":true,"manage_customer_tag":true},"range":[{"type":1,"userid":"zhangsan"},{"type":2,"partyid":1}]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"strategy_id": 1
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_strategy/create?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsCustomerStrategyCreate{
		StrategyName: "NAME",
		AdminList:    []string{"zhangsan", "lisi"},
		Privilege: &CustomerStrategyPrivilege{
			ViewCustomerList:        true,
			ViewCustomerData:        true,
			ViewRoomList:            true,
			ContactMe:               true,
			JoinRoom:                true,
			ShareCustomer:           false,
			OperResignCustomer:      true,
			OperResignGroup:         true,
			SendCustomerMsg:         true,
			EditWelcomeMsg:          true,
			ViewBehaviorData:        true,
			ViewRoomData:            true,
			SendGroupMsg:            true,
			RoomDeduplication:       true,
			RapidReply:              true,
			OnjobCustomerTransfer:   true,
			EditAntiSpamRule:        true,
			ExportCustomerList:      true,
			ExportCustomerData:      true,
			ExportCustomerGroupList: true,
			ManageCustomerTag:       true,
		},
		Range: []*CustomerStrategyRange{
			{
				Type:   1,
				UserID: "zhangsan",
			},
			{
				Type:    2,
				PartyID: 1,
			},
		},
	}

	result := new(ResultCustomerStrategyCreate)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", CreateCustomerStrategy(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultCustomerStrategyCreate{
		StrategyID: 1,
	}, result)
}

func TestGetCustomerStrategyRange(t *testing.T) {
	body := []byte(`{"strategy_id":1,"cursor":"CURSOR","limit":1000}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"range": [
		{
			"type": 1,
			"userid": "zhangsan"
		},
		{
			"type": 2,
			"partyid": 1
		}
	],
	"next_cursor": "NEXT_CURSOR"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_strategy/get_range?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultCustomerStrategyRange)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetCustomerStrategyRange(1, "CURSOR", 1000, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultCustomerStrategyRange{
		Range: []*CustomerStrategyRange{
			{
				Type:   1,
				UserID: "zhangsan",
			},
			{
				Type:    2,
				PartyID: 1,
			},
		},
	}, result)
}

func TestEditCustomerStrategy(t *testing.T) {
	body := []byte(`{"strategy_id":1,"strategy_name":"NAME","admin_list":["zhangsan","lisi"],"privilege":{"view_customer_list":true,"view_customer_data":true,"view_room_list":true,"contact_me":true,"join_room":true,"share_customer":false,"oper_resign_customer":true,"oper_resign_group":true,"send_customer_msg":true,"edit_welcome_msg":true,"view_behavior_data":true,"view_room_data":true,"send_group_msg":true,"room_deduplication":true,"rapid_reply":true,"onjob_customer_transfer":true,"edit_anti_spam_rule":true,"export_customer_list":true,"export_customer_data":true,"export_customer_group_list":true,"manage_customer_tag":true},"range_add":[{"type":1,"userid":"zhangsan"},{"type":2,"partyid":1}],"range_del":[{"type":1,"userid":"lisi"},{"type":2,"partyid":2}]}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_strategy/edit?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsCustomerStrategyEdit{
		StrategyID:   1,
		StrategyName: "NAME",
		AdminList:    []string{"zhangsan", "lisi"},
		Privilege: &CustomerStrategyPrivilege{
			ViewCustomerList:        true,
			ViewCustomerData:        true,
			ViewRoomList:            true,
			ContactMe:               true,
			JoinRoom:                true,
			ShareCustomer:           false,
			OperResignCustomer:      true,
			OperResignGroup:         true,
			SendCustomerMsg:         true,
			EditWelcomeMsg:          true,
			ViewBehaviorData:        true,
			ViewRoomData:            true,
			SendGroupMsg:            true,
			RoomDeduplication:       true,
			RapidReply:              true,
			OnjobCustomerTransfer:   true,
			EditAntiSpamRule:        true,
			ExportCustomerList:      true,
			ExportCustomerData:      true,
			ExportCustomerGroupList: true,
			ManageCustomerTag:       true,
		},
		RangeAdd: []*CustomerStrategyRange{
			{
				Type:   1,
				UserID: "zhangsan",
			},
			{
				Type:    2,
				PartyID: 1,
			},
		},
		RangeDel: []*CustomerStrategyRange{
			{
				Type:   1,
				UserID: "lisi",
			},
			{
				Type:    2,
				PartyID: 2,
			},
		},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", EditCustomerStrategy(params))

	assert.Nil(t, err)
}

func TestDeleteCustomerStrategy(t *testing.T) {
	body := []byte(`{"strategy_id":1}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_strategy/del?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", DeleteCustomerStrategy(1))

	assert.Nil(t, err)
}

func TestListMomentStrategy(t *testing.T) {
	body := []byte(`{"cursor":"CURSOR","limit":1000}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"strategy": [
		{
			"strategy_id": 1
		},
		{
			"strategy_id": 2
		}
	],
	"next_cursor": "NEXT_CURSOR"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/moment_strategy/list?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultMomentStrategyList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListMomentStrategy("CURSOR", 1000, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMomentStrategyList{
		Strategy: []*MomentStrategyListData{
			{
				StrategyID: 1,
			},
			{
				StrategyID: 2,
			},
		},
		NextCursor: "NEXT_CURSOR",
	}, result)
}

func TestGetMomentStrategy(t *testing.T) {
	body := []byte(`{"strategy_id":1}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"strategy": {
		"strategy_id": 1,
		"parent_id": 0,
		"strategy_name": "NAME",
		"create_time": 1557838797,
		"admin_list": [
			"zhangsan",
			"lisi"
		],
		"privilege": {
			"view_moment_list": true,
			"send_moment": true,
			"manage_moment_cover_and_sign": true
		}
	}
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/moment_strategy/get?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultMomentStrategyGet)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetMomentStrategy(1, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMomentStrategyGet{
		Strategy: &MomentStrategy{
			StrategyID:   1,
			ParentID:     0,
			StrategyName: "NAME",
			CreateTime:   1557838797,
			AdminList:    []string{"zhangsan", "lisi"},
			Privilege: &MomentStrategyPrivilege{
				ViewMomentList:           true,
				SendMoment:               true,
				ManageMomentCoverAndSign: true,
			},
		},
	}, result)
}

func TestGetMomentStrategyRange(t *testing.T) {
	body := []byte(`{"strategy_id":1,"cursor":"CURSOR","limit":1000}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"range": [
		{
			"type": 1,
			"userid": "zhangsan"
		},
		{
			"type": 2,
			"partyid": 1
		}
	],
	"next_cursor": "NEXT_CURSOR"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/moment_strategy/get_range?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultMomentStrategyRange)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetMomentStrategyRange(1, "CURSOR", 1000, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMomentStrategyRange{
		Range: []*MomentStrategyRange{
			{
				Type:   1,
				UserID: "zhangsan",
			},
			{
				Type:    2,
				PartyID: 1,
			},
		},
	}, result)
}

func TestCreateMomentStrategy(t *testing.T) {
	body := []byte(`{"strategy_name":"NAME","admin_list":["zhangsan","lisi"],"privilege":{"send_moment":true,"view_moment_list":true,"manage_moment_cover_and_sign":true},"range":[{"type":1,"userid":"zhangsan"},{"type":2,"partyid":1}]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"strategy_id": 1
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/moment_strategy/create?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsMomentStrategyCreate{
		StrategyName: "NAME",
		AdminList:    []string{"zhangsan", "lisi"},
		Privilege: &MomentStrategyPrivilege{
			SendMoment:               true,
			ViewMomentList:           true,
			ManageMomentCoverAndSign: true,
		},
		Range: []*MomentStrategyRange{
			{
				Type:   1,
				UserID: "zhangsan",
			},
			{
				Type:    2,
				PartyID: 1,
			},
		},
	}

	result := new(ResultMomentStrategyCreate)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", CreateMomentStrategy(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMomentStrategyCreate{
		StrategyID: 1,
	}, result)
}

func TestEditMomentStrategy(t *testing.T) {
	body := []byte(`{"strategy_id":1,"strategy_name":"NAME","admin_list":["zhangsan","lisi"],"privilege":{"send_moment":true,"view_moment_list":true,"manage_moment_cover_and_sign":true},"range_add":[{"type":1,"userid":"zhangsan"},{"type":2,"partyid":1}],"range_del":[{"type":1,"userid":"lisi"},{"type":2,"partyid":2}]}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/moment_strategy/edit?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsMomentStrategyEdit{
		StrategyID:   1,
		StrategyName: "NAME",
		AdminList:    []string{"zhangsan", "lisi"},
		Privilege: &MomentStrategyPrivilege{
			SendMoment:               true,
			ViewMomentList:           true,
			ManageMomentCoverAndSign: true,
		},
		RangeAdd: []*MomentStrategyRange{
			{
				Type:   1,
				UserID: "zhangsan",
			},
			{
				Type:    2,
				PartyID: 1,
			},
		},
		RangeDel: []*MomentStrategyRange{
			{
				Type:   1,
				UserID: "lisi",
			},
			{
				Type:    2,
				PartyID: 2,
			},
		},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", EditMomentStrategy(params))

	assert.Nil(t, err)
}

func TestDeleteMomentStrategy(t *testing.T) {
	body := []byte(`{"strategy_id":1}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/moment_strategy/del?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", DeleteMomentStrategy(1))

	assert.Nil(t, err)
}
