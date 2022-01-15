package agent

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

func TestGetAgent(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"agentid": 1000005,
	"name": "HR助手",
	"square_logo_url": "https://p.qlogo.cn/bizmail/FicwmI50icF8GH9ib7rUAYR5kicLTgP265naVFQKnleqSlRhiaBx7QA9u7Q/0",
	"description": "HR服务与员工自助平台",
	"allow_userinfos": {
		"user": [
			{
				"userid": "zhangshan"
			},
			{
				"userid": "lisi"
			}
		]
	},
	"allow_partys": {
		"partyid": [
			1
		]
	},
	"allow_tags": {
		"tagid": [
			1,
			2,
			3
		]
	},
	"close": 0,
	"redirect_domain": "open.work.weixin.qq.com",
	"report_location_flag": 0,
	"isreportenter": 0,
	"home_url": "https://open.work.weixin.qq.com",
	"customized_publish_status": 1
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/agent/get?access_token=ACCESS_TOKEN&agentid=1000005", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultAgentGet)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetAgent(1000005, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultAgentGet{
		AgentID:                 1000005,
		Name:                    "HR助手",
		SquareLogoURL:           "https://p.qlogo.cn/bizmail/FicwmI50icF8GH9ib7rUAYR5kicLTgP265naVFQKnleqSlRhiaBx7QA9u7Q/0",
		Description:             "HR服务与员工自助平台",
		Close:                   0,
		RedirectDomain:          "open.work.weixin.qq.com",
		ReportLocationFlag:      0,
		ISReportenter:           0,
		HomeURL:                 "https://open.work.weixin.qq.com",
		CustomizedPublishStatus: 1,
		AllowUserInfos: AllowUserInfos{
			User: []AllowUser{
				{
					UserID: "zhangshan",
				},
				{
					UserID: "lisi",
				},
			},
		},
		AllowPartys: AllowPartys{
			PartyID: []int64{1},
		},
		AllowTags: AllowTags{
			TagID: []int64{1, 2, 3},
		},
	}, result)
}

func TestListAgent(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"agentlist": [
		{
			"agentid": 1000005,
			"name": "HR助手",
			"square_logo_url": "https://p.qlogo.cn/bizmail/FicwmI50icF8GH9ib7rUAYR5kicLTgP265naVFQKnleqSlRhiaBx7QA9u7Q/0"
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/agent/list?access_token=ACCESS_TOKEN", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultAgentList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListAgent(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultAgentList{
		AgentList: []*AgentListData{
			{
				AgentID:       1000005,
				Name:          "HR助手",
				SquareLogoURL: "https://p.qlogo.cn/bizmail/FicwmI50icF8GH9ib7rUAYR5kicLTgP265naVFQKnleqSlRhiaBx7QA9u7Q/0",
			},
		},
	}, result)
}

func TestSetAgent(t *testing.T) {
	body := []byte(`{"agentid":1000005,"report_location_flag":0,"logo_mediaid":"j5Y8X5yocspvBHcgXMSS6z1Cn9RQKREEJr4ecgLHi4YHOYP-plvom-yD9zNI0vEl","name":"财经助手","description":"内部财经服务平台","redirect_domain":"open.work.weixin.qq.com","isreportenter":0,"home_url":"https://open.work.weixin.qq.com"}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/agent/set?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsAgentSet{
		AgentID:            1000005,
		ReportLocationFlag: 0,
		LogoMediaID:        "j5Y8X5yocspvBHcgXMSS6z1Cn9RQKREEJr4ecgLHi4YHOYP-plvom-yD9zNI0vEl",
		Name:               "财经助手",
		Description:        "内部财经服务平台",
		RedirectDomain:     "open.work.weixin.qq.com",
		ISReportEnter:      0,
		HomeURL:            "https://open.work.weixin.qq.com",
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SetAgent(params))

	assert.Nil(t, err)
}
