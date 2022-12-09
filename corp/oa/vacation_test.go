package oa

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/corp"
	"github.com/shenghui0779/gochat/mock"
)

func TestGetCorpVacationConf(t *testing.T) {
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"lists": [
		{
			"id": 1,
			"name": "年假",
			"time_attr": 0,
			"duration_type": 0,
			"quota_attr": {
				"type": 2,
				"autoreset_time": 0,
				"autoreset_duration": 0
			},
			"perday_duration": 86400
		},
		{
			"id": 2,
			"name": "事假",
			"time_attr": 0,
			"duration_type": 0,
			"quota_attr": {
				"type": 2,
				"autoreset_time": 0,
				"autoreset_duration": 0
			},
			"perday_duration": 86400
		},
		{
			"id": 3,
			"name": "病假",
			"time_attr": 0,
			"duration_type": 0,
			"quota_attr": {
				"type": 2,
				"autoreset_time": 0,
				"autoreset_duration": 0
			},
			"perday_duration": 86400
		}
	]
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/oa/vacation/getcorpconf?access_token=ACCESS_TOKEN", nil).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultCorpVacationConf)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetCorpVacationConf(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultCorpVacationConf{
		Lists: []*VacationConf{
			{
				ID:           1,
				Name:         "年假",
				TimeAttr:     0,
				DurationType: 0,
				QuotaAttr: &VacationQuotaAttr{
					Type:              2,
					AutoresetTime:     0,
					AutoresetDuration: 0,
				},
				PerdayDuration: 86400,
			},
			{
				ID:           2,
				Name:         "事假",
				TimeAttr:     0,
				DurationType: 0,
				QuotaAttr: &VacationQuotaAttr{
					Type:              2,
					AutoresetTime:     0,
					AutoresetDuration: 0,
				},
				PerdayDuration: 86400,
			},
			{
				ID:           3,
				Name:         "病假",
				TimeAttr:     0,
				DurationType: 0,
				QuotaAttr: &VacationQuotaAttr{
					Type:              2,
					AutoresetTime:     0,
					AutoresetDuration: 0,
				},
				PerdayDuration: 86400,
			},
		},
	}, result)
}

func TestGetUserVacationQuota(t *testing.T) {
	body := []byte(`{"userid":"ZhangSan"}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"lists": [
		{
			"id": 1,
			"assignduration": 0,
			"usedduration": 0,
			"leftduration": 604800,
			"vacationname": "年假"
		},
		{
			"id": 2,
			"assignduration": 0,
			"usedduration": 0,
			"leftduration": 1296000,
			"vacationname": "事假"
		},
		{
			"id": 3,
			"assignduration": 0,
			"usedduration": 0,
			"leftduration": 0,
			"vacationname": "病假"
		}
	]
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/oa/vacation/getuservacationquota?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	result := new(ResultUserVacationQuota)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetUserVacationQuota("ZhangSan", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultUserVacationQuota{
		Lists: []*VacationQuota{
			{
				ID:             1,
				AssignDuration: 0,
				UsedDuration:   0,
				LeftDuration:   604800,
				VacationName:   "年假",
			},
			{
				ID:             2,
				AssignDuration: 0,
				UsedDuration:   0,
				LeftDuration:   1296000,
				VacationName:   "事假",
			},
			{
				ID:             3,
				AssignDuration: 0,
				UsedDuration:   0,
				LeftDuration:   0,
				VacationName:   "病假",
			},
		},
	}, result)
}

func TestSetOneUserQuota(t *testing.T) {
	body := []byte(`{"userid":"ZhangSan","vacation_id":1,"leftduration":604800,"time_attr":1,"remarks":"PLACE_HOLDER"}`)
	resp := []byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/oa/vacation/setoneuserquota?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID", corp.WithMockClient(client))

	params := &ParamsOneUserQuotaSet{
		UserID:       "ZhangSan",
		VacationID:   1,
		LeftDuration: 604800,
		TimeAttr:     1,
		Remarks:      "PLACE_HOLDER",
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SetOneUserQuota(params))

	assert.Nil(t, err)
}
