package externalcontact

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/corp"
	"github.com/shenghui0779/gochat/mock"
	"github.com/shenghui0779/gochat/wx"
	"github.com/stretchr/testify/assert"
)

func TestAddInterceptRule(t *testing.T) {
	body := []byte(`{"rule_name":"rulename","word_list":["敏感词1","敏感词2"],"semantics_list":[1,2,3],"intercept_type":1,"applicable_range":{"user_list":["zhangshan"],"department_list":[2,3]}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"rule_id": "xxx"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_intercept_rule?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsInterceptRuleAdd{
		RuleName:      "rulename",
		WordList:      []string{"敏感词1", "敏感词2"},
		SemanticsList: []int{1, 2, 3},
		InterceptType: 1,
		ApplicableRange: &RuleApplicableRange{
			UserList:      []string{"zhangshan"},
			DeparmentList: []int64{2, 3},
		},
	}

	result := new(ResultInterceptRuleAdd)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", AddInterceptRule(params, result))

	assert.Nil(t, err)
	assert.Equal(t, ResultInterceptRuleAdd{
		RuleID: "xxx",
	}, result)
}

func TestUpdateInterceptRule(t *testing.T) {
	body := []byte(`{"rule_id":"xxxx","rule_name":"rulename","word_list":["敏感词1","敏感词2"],"extra_rule":{"semantics_list":[1,2,3]},"intercept_type":1,"add_applicable_range":{"user_list":["zhangshan"],"department_list":[2,3]},"remove_applicable_range":{"user_list":["zhangshan"],"department_list":[2,3]}}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/update_intercept_rule?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsInterceptRuleUpdate{
		RuleID:   "xxxx",
		RuleName: "rulename",
		WordList: []string{"敏感词1", "敏感词2"},
		ExtraRule: &ExtraRule{
			SemanticsList: []int{1, 2, 3},
		},
		InterceptType: 1,
		AddApplicableRange: &RuleApplicableRange{
			UserList:      []string{"zhangshan"},
			DeparmentList: []int64{2, 3},
		},
		RemoveApplicableRange: &RuleApplicableRange{
			UserList:      []string{"zhangshan"},
			DeparmentList: []int64{2, 3},
		},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UpdateInterceptRule(params))

	assert.Nil(t, err)
}

func TestListInterceptRule(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"rule_list": [
		{
			"rule_id": "xxxx",
			"rule_name": "rulename",
			"create_time": 1600000000
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_intercept_rule_list?access_token=ACCESS_TOKEN", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultInterceptRuleList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListInterceptRule(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultInterceptRuleList{
		RuleList: []*RuleListData{
			{
				RuleID:     "xxxx",
				RuleName:   "rulename",
				CreateTime: 1600000000,
			},
		},
	}, result)
}

func TestGetInterceptRule(t *testing.T) {
	body := []byte(`{"rule_id":"xxx"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"rule": {
		"rule_id": "xxx",
		"rule_name": "rulename",
		"word_list": [
			"敏感词1",
			"敏感词2"
		],
		"extra_rule": {
			"semantics_list": [
				1,
				2,
				3
			]
		},
		"intercept_type": 1,
		"applicable_range": {
			"user_list": [
				"zhangshan"
			],
			"department_list": [
				2,
				3
			]
		}
	}
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_intercept_rule?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultInterceptRuleGet)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetInterceptRule("xxx", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultInterceptRuleGet{
		Rule: &InterceptRule{
			RuleID:   "xxx",
			RuleName: "rulename",
			WordList: []string{"敏感词1", "敏感词2"},
			ExtraRule: &ExtraRule{
				SemanticsList: []int{1, 2, 3},
			},
			InterceptType: 1,
			ApplicableRange: &RuleApplicableRange{
				UserList:      []string{"zhangshan"},
				DeparmentList: []int64{2, 3},
			},
		},
	}, result)
}

func TestDeleteInterceptRule(t *testing.T) {
	body := []byte(`{"rule_id":"xxx"}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/del_intercept_rule?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", DeleteInterceptRule("xxx"))

	assert.Nil(t, err)
}
