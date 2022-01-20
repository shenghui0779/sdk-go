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
				UserList:      []string{"zhangshan", "lisi"},
				DeparmentList: []int64{2, 3},
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
				UserList:      []string{"zhangshan", "lisi"},
				DeparmentList: []int64{2, 3},
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
			{},
		},
	}, result)
}

func TestGetMomentTask(t *testing.T) {
	body := []byte(``)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/user/authsucc?access_token=ACCESS_TOKEN&userid=USERID", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN")

	assert.Nil(t, err)
}

func TestListMomentCustomer(t *testing.T) {
	body := []byte(``)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/user/authsucc?access_token=ACCESS_TOKEN&userid=USERID", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN")

	assert.Nil(t, err)
}

func TestGetMomentSendResult(t *testing.T) {
	body := []byte(``)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/user/authsucc?access_token=ACCESS_TOKEN&userid=USERID", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN")

	assert.Nil(t, err)
}

func TestGetMomentComments(t *testing.T) {
	body := []byte(``)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/user/authsucc?access_token=ACCESS_TOKEN&userid=USERID", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN")

	assert.Nil(t, err)
}
