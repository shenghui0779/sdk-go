package externalcontact

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/corp"
	"github.com/shenghui0779/gochat/event"
	"github.com/shenghui0779/gochat/mock"
	"github.com/shenghui0779/gochat/wx"
)

func TestAddMomentTask(t *testing.T) {
	body := []byte(`{"text":{"content":"文本消息内容"},"attachments":[{"msgtype":"image","image":{"media_id":"MEDIA_ID"}},{"msgtype":"video","video":{"media_id":"MEDIA_ID"}},{"msgtype":"link","link":{"title":"消息标题","url":"https://example.link.com/path","media_id":"MEDIA_ID"}}],"visible_range":{"sender_list":{"user_list":["zhangshan","lisi"],"department_list":[2,3]},"external_contact_list":{"tag_list":["etXXXXXXXXXX","etYYYYYYYYYY"]}}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "jobid": "xxxx"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_moment_task?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsMomentTaskAdd{
		Text: &MomentText{
			Content: "文本消息内容",
		},
		Attachments: []*MomentAttachment{
			{
				MsgType: event.MsgImage,
				Image: &MomentImage{
					MediaID: "MEDIA_ID",
				},
			},
			{
				MsgType: event.MsgVideo,
				Video: &MomentVideo{
					MediaID: "MEDIA_ID",
				},
			},
			{
				MsgType: event.MsgLink,
				Link: &MomentLink{
					Title:   "消息标题",
					URL:     "https://example.link.com/path",
					MediaID: "MEDIA_ID",
				},
			},
		},
		VisibleRange: &MomentVisibleRange{
			SenderList: &MomentSenderList{
				UserList:       []string{"zhangshan", "lisi"},
				DepartmentList: []int64{2, 3},
			},
			ExternalContactList: &MomentExternalContactList{
				TagList: []string{"etXXXXXXXXXX", "etYYYYYYYYYY"},
			},
		},
	}

	result := new(ResultMomentTaskAdd)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", AddMomentTask(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMomentTaskAdd{
		JobID: "xxxx",
	}, result)
}

func TestGetMomentTaskResult(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "status": 1,
    "type": "add_moment_task",
    "result": {
        "errcode": 0,
        "errmsg": "ok",
        "moment_id": "xxxx",
        "invalid_sender_list": {
            "user_list": [
                "zhangshan",
                "lisi"
            ],
            "department_list": [
                2,
                3
            ]
        },
        "invalid_external_contact_list": {
            "tag_list": [
                "xxx"
            ]
        }
    }
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_moment_task_result?access_token=ACCESS_TOKEN&jobid=JOBID", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultMomentTaskResult)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetMomentTaskResult("JOBID", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMomentTaskResult{
		Status: 1,
		Type:   "add_moment_task",
		Result: &MomentTaskResult{
			ErrCode:  0,
			ErrMsg:   "ok",
			MomentID: "xxxx",
			InvalidSenderList: &MomentInvalidSenderList{
				UserList:       []string{"zhangshan", "lisi"},
				DepartmentList: []int64{2, 3},
			},
			InvalidExternalContactList: &MomentInvalidExternalContactList{
				TagList: []string{"xxx"},
			},
		},
	}, result)
}

func TestListMoment(t *testing.T) {
	body := []byte(`{"start_time":1605000000,"end_time":1605172726,"creator":"zhangsan","filter_type":1,"cursor":"CURSOR","limit":10}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "next_cursor": "CURSOR",
    "moment_list": [
        {
            "moment_id": "momxxx",
            "creator": "xxxx",
            "create_time": "xxxx",
            "create_type": 1,
            "visible_type": 1,
            "text": {
                "content": "test"
            },
            "image": [
                {
                    "media_id": "WWCISP_xxxxx"
                }
            ],
            "video": {
                "media_id": "WWCISP_xxxxx",
                "thumb_media_id": "WWCISP_xxxxx"
            },
            "link": {
                "title": "腾讯网-QQ.COM",
                "url": "https://www.qq.com"
            },
            "location": {
                "latitude": "23.10647",
                "longitude": "113.32446",
                "name": "广州市 · 广州塔"
            }
        }
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_moment_list?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsMomentList{
		StartTime:  1605000000,
		EndTime:    1605172726,
		Creator:    "zhangsan",
		FilterType: 1,
		Cursor:     "CURSOR",
		Limit:      10,
	}

	result := new(ResultMomentList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListMoment(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMomentList{
		NextCursor: "CURSOR",
		MomentList: []*MomentListData{
			{
				MomentID:    "momxxx",
				Creator:     "xxxx",
				CreateTime:  "xxxx",
				CreateType:  1,
				VisibleType: 1,
				Text: &MomentText{
					Content: "test",
				},
				Image: []*MomentImage{
					{
						MediaID: "WWCISP_xxxxx",
					},
				},
				Video: &MomentVideo{
					MediaID:      "WWCISP_xxxxx",
					ThumbMediaID: "WWCISP_xxxxx",
				},
				Link: &MomentLink{
					Title: "腾讯网-QQ.COM",
					URL:   "https://www.qq.com",
				},
				Location: &MomentLocation{
					Latitude:  "23.10647",
					Longitude: "113.32446",
					Name:      "广州市 · 广州塔",
				},
			},
		},
	}, result)
}

func TestGetMomentTask(t *testing.T) {
	body := []byte(`{"moment_id":"momxxx","cursor":"CURSOR","limit":10}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "next_cursor": "CURSOR",
    "task_list": [
        {
            "userid": "zhangsan",
            "publish_status": 1
        }
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_moment_task?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultMomentTaskGet)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetMomentTask("momxxx", "CURSOR", 10, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMomentTaskGet{
		NextCursor: "CURSOR",
		TaskList: []*MomentTask{
			{
				UserID:        "zhangsan",
				PublishStatus: 1,
			},
		},
	}, result)
}

func TestListMomentCustomer(t *testing.T) {
	body := []byte(`{"moment_id":"momxxx","userid":"xxx","cursor":"CURSOR","limit":10}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode":0,
	"errmsg":"ok",
	"next_cursor":"CURSOR",
	"customer_list":[
		{
			"userid":"xxx",
			"external_userid":"woAJ2GCAAAXtWyujaWJHDDGi0mACCCC"
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_moment_customer_list?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultMomentCustomerList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListMomentCustomer("momxxx", "xxx", "CURSOR", 10, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMomentCustomerList{
		NextCursor: "CURSOR",
		CustomerList: []*MomentCustomer{
			{
				UserID:         "xxx",
				ExternalUserID: "woAJ2GCAAAXtWyujaWJHDDGi0mACCCC",
			},
		},
	}, result)
}

func TestGetMomentSendResult(t *testing.T) {
	body := []byte(`{"moment_id":"momxxx","userid":"xxx","cursor":"CURSOR","limit":100}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "next_cursor": "CURSOR",
    "customer_list": [
        {
            "external_userid": "woAJ2GCAAAXtWyujaWJHDDGi0mACCCC"
        }
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_moment_send_result?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultMomentSendResult)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetMomentSendResult("momxxx", "xxx", "CURSOR", 100, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMomentSendResult{
		NextCursor: "CURSOR",
		CustomerList: []*MomentCustomer{
			{
				ExternalUserID: "woAJ2GCAAAXtWyujaWJHDDGi0mACCCC",
			},
		},
	}, result)
}

func TestGetMomentComments(t *testing.T) {
	body := []byte(`{"moment_id":"momxxx","userid":"xxx"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "comment_list": [
        {
            "external_userid": "woAJ2GCAAAXtWyujaWJHDDGi0mACAAAA",
            "create_time": 1605172726
        },
        {
            "userid": "zhangshan",
            "create_time": 1605172729
        }
    ],
    "like_list": [
        {
            "external_userid": "woAJ2GCAAAXtWyujaWJHDDGi0mACBBBB",
            "create_time": 1605172726
        },
        {
            "userid": "zhangshan",
            "create_time": 1605172720
        }
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_moment_comments?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultMomentComments)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetMomentComments("momxxx", "xxx", "", 0, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMomentComments{
		CommentList: []*MomentCommentData{
			{
				ExternalUserID: "woAJ2GCAAAXtWyujaWJHDDGi0mACAAAA",
				CreateTime:     1605172726,
			},
			{
				UserID:     "zhangshan",
				CreateTime: 1605172729,
			},
		},
		LikeList: []*MomentLikeData{
			{
				ExternalUserID: "woAJ2GCAAAXtWyujaWJHDDGi0mACBBBB",
				CreateTime:     1605172726,
			},
			{
				UserID:     "zhangshan",
				CreateTime: 1605172720,
			},
		},
	}, result)
}
