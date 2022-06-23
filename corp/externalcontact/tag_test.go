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

func TestListCorpTag(t *testing.T) {
	body := []byte(`{"tag_id":["etXXXXXXXXXX","etYYYYYYYYYY"],"group_id":["etZZZZZZZZZZZZZ","etYYYYYYYYYYYYY"]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "tag_group": [
        {
            "group_id": "TAG_GROUPID1",
            "group_name": "GOURP_NAME",
            "create_time": 1557838797,
            "order": 1,
            "deleted": false,
            "tag": [
                {
                    "id": "TAG_ID1",
                    "name": "NAME1",
                    "create_time": 1557838797,
                    "order": 1,
                    "deleted": false
                },
                {
                    "id": "TAG_ID2",
                    "name": "NAME2",
                    "create_time": 1557838797,
                    "order": 2,
                    "deleted": true
                }
            ]
        }
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_corp_tag_list?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	tagIDs := []string{"etXXXXXXXXXX", "etYYYYYYYYYY"}
	groupIDs := []string{"etZZZZZZZZZZZZZ", "etYYYYYYYYYYYYY"}

	result := new(ResultCorpTagList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListCorpTag(tagIDs, groupIDs, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultCorpTagList{
		TagGroup: []*CorpTagGroup{
			{
				GroupID:    "TAG_GROUPID1",
				GroupName:  "GOURP_NAME",
				CreateTime: 1557838797,
				Order:      1,
				Deleted:    false,
				Tag: []*CorpTag{
					{
						ID:         "TAG_ID1",
						Name:       "NAME1",
						CreateTime: 1557838797,
						Order:      1,
						Deleted:    false,
					},
					{
						ID:         "TAG_ID2",
						Name:       "NAME2",
						CreateTime: 1557838797,
						Order:      2,
						Deleted:    true,
					},
				},
			},
		},
	}, result)
}

func TestAddCorpTag(t *testing.T) {
	body := []byte(`{"group_id":"GROUP_ID","group_name":"GROUP_NAME","order":1,"tag":[{"name":"TAG_NAME_1","order":1},{"name":"TAG_NAME_2","order":2}],"agentid":1000014}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "tag_group": {
        "group_id": "TAG_GROUPID1",
        "group_name": "GOURP_NAME",
        "create_time": 1557838797,
        "order": 1,
        "tag": [
            {
                "id": "TAG_ID1",
                "name": "NAME1",
                "create_time": 1557838797,
                "order": 1
            },
            {
                "id": "TAG_ID2",
                "name": "NAME2",
                "create_time": 1557838797,
                "order": 2
            }
        ]
    }
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_corp_tag?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsCorpTagAdd{
		GroupID:   "GROUP_ID",
		GroupName: "GROUP_NAME",
		Order:     1,
		Tag: []*ParamsCorpTag{
			{
				Name:  "TAG_NAME_1",
				Order: 1,
			},
			{
				Name:  "TAG_NAME_2",
				Order: 2,
			},
		},
		AgentID: 1000014,
	}

	result := new(ResultCorpTagAdd)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", AddCorpTag(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultCorpTagAdd{
		TagGroup: &CorpTagGroup{
			GroupID:    "TAG_GROUPID1",
			GroupName:  "GOURP_NAME",
			CreateTime: 1557838797,
			Order:      1,
			Tag: []*CorpTag{
				{
					ID:         "TAG_ID1",
					Name:       "NAME1",
					CreateTime: 1557838797,
					Order:      1,
				},
				{
					ID:         "TAG_ID2",
					Name:       "NAME2",
					CreateTime: 1557838797,
					Order:      2,
				},
			},
		},
	}, result)
}

func TestEditCorpTag(t *testing.T) {
	body := []byte(`{"id":"TAG_ID","name":"NEW_TAG_NAME","order":1,"agentid":1000014}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/edit_corp_tag?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsCorpTagEdit{
		ID:      "TAG_ID",
		Name:    "NEW_TAG_NAME",
		Order:   1,
		AgentID: 1000014,
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", EditCorpTag(params))

	assert.Nil(t, err)
}

func TestDeleteCorpTag(t *testing.T) {
	body := []byte(`{"tag_id":["TAG_ID_1","TAG_ID_2"],"group_id":["GROUP_ID_1","GROUP_ID_2"],"agentid":1000014}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/del_corp_tag?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	tagIDs := []string{"TAG_ID_1", "TAG_ID_2"}
	groupIDs := []string{"GROUP_ID_1", "GROUP_ID_2"}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", DeleteCorpTag(tagIDs, groupIDs, 1000014))

	assert.Nil(t, err)
}

func TestListStrategyTag(t *testing.T) {
	body := []byte(`{"strategy_id":1,"tag_id":["etXXXXXXXXXX","etYYYYYYYYYY"],"group_id":["etZZZZZZZZZZZZZ","etYYYYYYYYYYYYY"]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "tag_group": [
        {
            "group_id": "TAG_GROUPID1",
            "group_name": "GOURP_NAME",
            "create_time": 1557838797,
            "order": 1,
            "strategy_id": 1,
            "tag": [
                {
                    "id": "TAG_ID1",
                    "name": "NAME1",
                    "create_time": 1557838797,
                    "order": 1
                },
                {
                    "id": "TAG_ID2",
                    "name": "NAME2",
                    "create_time": 1557838797,
                    "order": 2
                }
            ]
        }
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_strategy_tag_list?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	tagIDs := []string{"etXXXXXXXXXX", "etYYYYYYYYYY"}
	groupIDs := []string{"etZZZZZZZZZZZZZ", "etYYYYYYYYYYYYY"}

	result := new(ResultStrategyTagList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListStrategyTag(1, tagIDs, groupIDs, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultStrategyTagList{
		TagGroup: []*StrategyTagGroup{
			{
				GroupID:    "TAG_GROUPID1",
				GroupName:  "GOURP_NAME",
				CreateTime: 1557838797,
				Order:      1,
				StrategyID: 1,
				Tag: []*StrategyTag{
					{
						ID:         "TAG_ID1",
						Name:       "NAME1",
						CreateTime: 1557838797,
						Order:      1,
					},
					{
						ID:         "TAG_ID2",
						Name:       "NAME2",
						CreateTime: 1557838797,
						Order:      2,
					},
				},
			},
		},
	}, result)
}

func TestAddStrategyTag(t *testing.T) {
	body := []byte(`{"strategy_id":1,"group_id":"GROUP_ID","group_name":"GROUP_NAME","order":1,"tag":[{"name":"TAG_NAME_1","order":1},{"name":"TAG_NAME_2","order":2}]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "tag_group": {
        "group_id": "TAG_GROUPID1",
        "group_name": "GOURP_NAME",
        "create_time": 1557838797,
        "order": 1,
        "tag": [
            {
                "id": "TAG_ID1",
                "name": "NAME1",
                "create_time": 1557838797,
                "order": 1
            },
            {
                "id": "TAG_ID2",
                "name": "NAME2",
                "create_time": 1557838797,
                "order": 2
            }
        ]
    }
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_strategy_tag?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsStrategyTagAdd{
		StrategyID: 1,
		GroupID:    "GROUP_ID",
		GroupName:  "GROUP_NAME",
		Order:      1,
		Tag: []*ParamsStrategyTag{
			{
				Name:  "TAG_NAME_1",
				Order: 1,
			},
			{
				Name:  "TAG_NAME_2",
				Order: 2,
			},
		},
	}

	result := new(ResultStrategyTagAdd)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", AddStrategyTag(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultStrategyTagAdd{
		TagGroup: &StrategyTagGroup{
			GroupID:    "TAG_GROUPID1",
			GroupName:  "GOURP_NAME",
			CreateTime: 1557838797,
			Order:      1,
			Tag: []*StrategyTag{
				{
					ID:         "TAG_ID1",
					Name:       "NAME1",
					CreateTime: 1557838797,
					Order:      1,
				},
				{
					ID:         "TAG_ID2",
					Name:       "NAME2",
					CreateTime: 1557838797,
					Order:      2,
				},
			},
		},
	}, result)
}

func TestEditStrategyTag(t *testing.T) {
	body := []byte(`{"id":"TAG_ID","name":"NEW_TAG_NAME","order":1}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/edit_strategy_tag?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsStrategyTagEdit{
		ID:    "TAG_ID",
		Name:  "NEW_TAG_NAME",
		Order: 1,
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", EditStrategyTag(params))

	assert.Nil(t, err)
}

func TestDeleteStrategyTag(t *testing.T) {
	body := []byte(`{"tag_id":["TAG_ID_1","TAG_ID_2"],"group_id":["GROUP_ID_1","GROUP_ID_2"]}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/del_strategy_tag?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	tagIDs := []string{"TAG_ID_1", "TAG_ID_2"}
	groupIDs := []string{"GROUP_ID_1", "GROUP_ID_2"}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", DeleteStrategyTag(tagIDs, groupIDs))

	assert.Nil(t, err)
}

func TestMarkTag(t *testing.T) {
	body := []byte(`{"userid":"zhangsan","external_userid":"woAJ2GCAAAd1NPGHKSD4wKmE8Aabj9AAA","add_tag":["TAGID1","TAGID2"],"remove_tag":["TAGID3","TAGID4"]}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/mark_tag?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsTagMark{
		UserID:         "zhangsan",
		ExternalUserID: "woAJ2GCAAAd1NPGHKSD4wKmE8Aabj9AAA",
		AddTag:         []string{"TAGID1", "TAGID2"},
		RemoveTag:      []string{"TAGID3", "TAGID4"},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", MarkTag(params))

	assert.Nil(t, err)
}
