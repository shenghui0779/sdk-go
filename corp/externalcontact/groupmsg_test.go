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
	"github.com/chenghonour/gochat/event"
	"github.com/chenghonour/gochat/mock"
	"github.com/chenghonour/gochat/wx"
)

func TestAddMsgTemplate(t *testing.T) {
	body := []byte(`{"chat_type":"single","external_userid":["woAJ2GCAAAXtWyujaWJHDDGi0mACAAAA","wmqfasd1e1927831123109rBAAAA"],"sender":"zhangsan","text":{"content":"文本消息内容"},"attachments":[{"msgtype":"image","image":{"media_id":"MEDIA_ID","pic_url":"http://p.qpic.cn/pic_wework/3474110808/7a6344sdadfwehe42060/0"}},{"msgtype":"link","link":{"title":"消息标题","picurl":"https://example.pic.com/path","desc":"消息描述","url":"https://example.link.com/path"}},{"msgtype":"miniprogram","miniprogram":{"title":"消息标题","pic_media_id":"MEDIA_ID","appid":"wx8bd80126147dfAAA","page":"/path/index.html"}},{"msgtype":"video","video":{"media_id":"MEDIA_ID"}},{"msgtype":"file","file":{"media_id":"MEDIA_ID"}}]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"fail_list": [
		"wmqfasd1e1927831123109rBAAAA"
	],
	"msgid": "msgGCAAAXtWyujaWJHDDGi0mAAAA"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_msg_template?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsMsgTemplateAdd{
		ChatType:       "single",
		ExternalUserID: []string{"woAJ2GCAAAXtWyujaWJHDDGi0mACAAAA", "wmqfasd1e1927831123109rBAAAA"},
		Sender:         "zhangsan",
		Text: &GroupText{
			Content: "文本消息内容",
		},
		Attachments: []*MsgAttachment{
			{
				MsgType: event.MsgImage,
				Image: &GroupImage{
					MediaID: "MEDIA_ID",
					PicURL:  "http://p.qpic.cn/pic_wework/3474110808/7a6344sdadfwehe42060/0",
				},
			},
			{
				MsgType: event.MsgLink,
				Link: &GroupLink{
					Title:  "消息标题",
					PicURL: "https://example.pic.com/path",
					Desc:   "消息描述",
					URL:    "https://example.link.com/path",
				},
			},
			{
				MsgType: event.MsgMinip,
				Minip: &GroupMinip{
					Title:      "消息标题",
					PicMediaID: "MEDIA_ID",
					AppID:      "wx8bd80126147dfAAA",
					Page:       "/path/index.html",
				},
			},
			{
				MsgType: event.MsgVideo,
				Video: &GroupVideo{
					MediaID: "MEDIA_ID",
				},
			},
			{
				MsgType: event.MsgFile,
				File: &GroupFile{
					MediaID: "MEDIA_ID",
				},
			},
		},
	}

	result := new(ResultMsgTemplateAdd)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", AddMsgTemplate(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMsgTemplateAdd{
		FailList: []string{"wmqfasd1e1927831123109rBAAAA"},
		MsgID:    "msgGCAAAXtWyujaWJHDDGi0mAAAA",
	}, result)
}

func TestListGroupMsg(t *testing.T) {
	body := []byte(`{"chat_type":"single","start_time":1605171726,"end_time":1605172726,"creator":"zhangshan","filter_type":1,"limit":50,"cursor":"CURSOR"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "next_cursor": "CURSOR",
    "group_msg_list": [
        {
            "msgid": "msgGCAAAXtWyujaWJHDDGi0mAAAA",
            "creator": "xxxx",
            "create_time": "xxxx",
            "create_type": 1,
            "text": {
                "content": "文本消息内容"
            },
            "attachments": [
                {
                    "msgtype": "image",
                    "image": {
                        "media_id": "MEDIA_ID",
                        "pic_url": "http://p.qpic.cn/pic_wework/3474110808/7a6344sdadfwehe42060/0"
                    }
                },
                {
                    "msgtype": "link",
                    "link": {
                        "title": "消息标题",
                        "picurl": "https://example.pic.com/path",
                        "desc": "消息描述",
                        "url": "https://example.link.com/path"
                    }
                },
                {
                    "msgtype": "miniprogram",
                    "miniprogram": {
                        "title": "消息标题",
                        "pic_media_id": "MEDIA_ID",
                        "appid": "wx8bd80126147dfAAA",
                        "page": "/path/index.html"
                    }
                },
                {
                    "msgtype": "video",
                    "video": {
                        "media_id": "MEDIA_ID"
                    }
                },
                {
                    "msgtype": "file",
                    "file": {
                        "media_id": "MEDIA_ID"
                    }
                }
            ]
        }
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_groupmsg_list_v2?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsGroupMsgList{
		ChatType:   ChatSingle,
		StartTime:  1605171726,
		EndTime:    1605172726,
		Creator:    "zhangshan",
		FilterType: 1,
		Limit:      50,
		Cursor:     "CURSOR",
	}

	result := new(ResultGroupMsgList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListGroupMsg(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultGroupMsgList{
		NextCursor: "CURSOR",
		GroupMsgList: []*GroupMsg{
			{
				MsgID:      "msgGCAAAXtWyujaWJHDDGi0mAAAA",
				Creator:    "xxxx",
				CreateTime: "xxxx",
				CreateType: 1,
				Text: &GroupText{
					Content: "文本消息内容",
				},
				Attachments: []*MsgAttachment{
					{
						MsgType: event.MsgImage,
						Image: &GroupImage{
							MediaID: "MEDIA_ID",
							PicURL:  "http://p.qpic.cn/pic_wework/3474110808/7a6344sdadfwehe42060/0",
						},
					},
					{
						MsgType: event.MsgLink,
						Link: &GroupLink{
							Title:  "消息标题",
							PicURL: "https://example.pic.com/path",
							Desc:   "消息描述",
							URL:    "https://example.link.com/path",
						},
					},
					{
						MsgType: event.MsgMinip,
						Minip: &GroupMinip{
							Title:      "消息标题",
							PicMediaID: "MEDIA_ID",
							AppID:      "wx8bd80126147dfAAA",
							Page:       "/path/index.html",
						},
					},
					{
						MsgType: event.MsgVideo,
						Video: &GroupVideo{
							MediaID: "MEDIA_ID",
						},
					},
					{
						MsgType: event.MsgFile,
						File: &GroupFile{
							MediaID: "MEDIA_ID",
						},
					},
				},
			},
		},
	}, result)
}

func TestGetGroupMsgTask(t *testing.T) {
	body := []byte(`{"msgid":"msgGCAAAXtWyujaWJHDDGi0mACAAAA","limit":50,"cursor":"CURSOR"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"next_cursor": "CURSOR",
	"task_list": [
		{
			"userid": "zhangsan",
			"status": 1,
			"send_time": 1552536375
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_groupmsg_task?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsGroupMsgTask{
		MsgID:  "msgGCAAAXtWyujaWJHDDGi0mACAAAA",
		Limit:  50,
		Cursor: "CURSOR",
	}

	result := new(ResultGroupMsgTask)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetGroupMsgTask(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultGroupMsgTask{
		NextCursor: "CURSOR",
		TaskList: []*GroupMsgTask{
			{
				UserID:   "zhangsan",
				Status:   1,
				SendTime: 1552536375,
			},
		},
	}, result)
}

func TestGetGroupMsgSendResult(t *testing.T) {
	body := []byte(`{"msgid":"msgGCAAAXtWyujaWJHDDGi0mACAAAA","userid":"zhangsan","limit":50,"cursor":"CURSOR"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"next_cursor": "CURSOR",
	"send_list": [
		{
			"external_userid": "wmqfasd1e19278asdasAAAA",
			"chat_id": "wrOgQhDgAAMYQiS5ol9G7gK9JVAAAA",
			"userid": "zhangsan",
			"status": 1,
			"send_time": 1552536375
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_groupmsg_send_result?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsGroupMsgSendResult{
		MsgID:  "msgGCAAAXtWyujaWJHDDGi0mACAAAA",
		UserID: "zhangsan",
		Limit:  50,
		Cursor: "CURSOR",
	}

	result := new(ResultGroupMsgSendResult)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetGroupMsgSendResult(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultGroupMsgSendResult{
		NextCursor: "CURSOR",
		SendList: []*GroupMsgSendResult{
			{
				ExternalUserID: "wmqfasd1e19278asdasAAAA",
				ChatID:         "wrOgQhDgAAMYQiS5ol9G7gK9JVAAAA",
				UserID:         "zhangsan",
				Status:         1,
				SendTime:       1552536375,
			},
		},
	}, result)
}

func TestSendWelcomeMsg(t *testing.T) {
	body := []byte(`{"welcome_code":"CALLBACK_CODE","text":{"content":"文本消息内容"},"attachments":[{"msgtype":"image","image":{"media_id":"MEDIA_ID","pic_url":"http://p.qpic.cn/pic_wework/3474110808/7a6344sdadfwehe42060/0"}},{"msgtype":"link","link":{"title":"消息标题","picurl":"https://example.pic.com/path","desc":"消息描述","url":"https://example.link.com/path"}},{"msgtype":"miniprogram","miniprogram":{"title":"消息标题","pic_media_id":"MEDIA_ID","appid":"wx8bd80126147dfAAA","page":"/path/index.html"}},{"msgtype":"video","video":{"media_id":"MEDIA_ID"}},{"msgtype":"file","file":{"media_id":"MEDIA_ID"}}]}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/send_welcome_msg?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsWelcomeMsgSend{
		WelcomeCode: "CALLBACK_CODE",
		Text: &GroupText{
			Content: "文本消息内容",
		},
		Attachments: []*MsgAttachment{
			{
				MsgType: event.MsgImage,
				Image: &GroupImage{
					MediaID: "MEDIA_ID",
					PicURL:  "http://p.qpic.cn/pic_wework/3474110808/7a6344sdadfwehe42060/0",
				},
			},
			{
				MsgType: event.MsgLink,
				Link: &GroupLink{
					Title:  "消息标题",
					PicURL: "https://example.pic.com/path",
					Desc:   "消息描述",
					URL:    "https://example.link.com/path",
				},
			},
			{
				MsgType: event.MsgMinip,
				Minip: &GroupMinip{
					Title:      "消息标题",
					PicMediaID: "MEDIA_ID",
					AppID:      "wx8bd80126147dfAAA",
					Page:       "/path/index.html",
				},
			},
			{
				MsgType: event.MsgVideo,
				Video: &GroupVideo{
					MediaID: "MEDIA_ID",
				},
			},
			{
				MsgType: event.MsgFile,
				File: &GroupFile{
					MediaID: "MEDIA_ID",
				},
			},
		},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendWelcomeMsg(params))

	assert.Nil(t, err)
}

func TestAddGroupWelcomeTemplate(t *testing.T) {
	body := []byte(`{"text":{"content":"亲爱的%NICKNAME%用户，你好"},"image":{"media_id":"MEDIA_ID","pic_url":"http://p.qpic.cn/pic_wework/3474110808/7a6344sdadfwehe42060/0"},"link":{"title":"消息标题","picurl":"https://example.pic.com/path","desc":"消息描述","url":"https://example.link.com/path"},"miniprogram":{"title":"消息标题","pic_media_id":"MEDIA_ID","appid":"wx8bd80126147dfAAA","page":"/path/index"},"file":{"media_id":"1Yv-zXfHjSjU-7LH-GwtYqDGS-zz6w22KmWAT5COgP7o"},"video":{"media_id":"MEDIA_ID"},"agentid":1000014,"notify":1}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"template_id": "msgXXXXXX"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/add?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsGroupWelcomeTemplateAdd{
		Text: &GroupText{
			Content: "亲爱的%NICKNAME%用户，你好",
		},
		Image: &GroupImage{
			MediaID: "MEDIA_ID",
			PicURL:  "http://p.qpic.cn/pic_wework/3474110808/7a6344sdadfwehe42060/0",
		},
		Link: &GroupLink{
			Title:  "消息标题",
			PicURL: "https://example.pic.com/path",
			Desc:   "消息描述",
			URL:    "https://example.link.com/path",
		},
		Minip: &GroupMinip{
			Title:      "消息标题",
			PicMediaID: "MEDIA_ID",
			AppID:      "wx8bd80126147dfAAA",
			Page:       "/path/index",
		},
		File: &GroupFile{
			MediaID: "1Yv-zXfHjSjU-7LH-GwtYqDGS-zz6w22KmWAT5COgP7o",
		},
		Video: &GroupVideo{
			MediaID: "MEDIA_ID",
		},
		AgentID: 1000014,
		Notify:  1,
	}

	result := new(ResultGroupWelcomeTemplateAdd)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", AddGroupWelcomeTemplate(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultGroupWelcomeTemplateAdd{
		TemplateID: "msgXXXXXX",
	}, result)
}

func TestEditGroupWelcomeTemplate(t *testing.T) {
	body := []byte(`{"template_id":"msgXXXXXXX","text":{"content":"文本消息内容"},"image":{"media_id":"MEDIA_ID","pic_url":"http://p.qpic.cn/pic_wework/3474110808/7a6344sdadfwehe42060/0"},"link":{"title":"消息标题","picurl":"https://example.pic.com/path","desc":"消息描述","url":"https://example.link.com/path"},"miniprogram":{"title":"消息标题","pic_media_id":"MEDIA_ID","appid":"wx8bd80126147df384","page":"/path/index"},"file":{"media_id":"1Yv-zXfHjSjU-7LH-GwtYqDGS-zz6w22KmWAT5COgP7o"},"video":{"media_id":"MEDIA_ID"},"agentid":1000014}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/edit?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsGroupWelcomeTemplateEdit{
		TemplateID: "msgXXXXXXX",
		Text: &GroupText{
			Content: "文本消息内容",
		},
		Image: &GroupImage{
			MediaID: "MEDIA_ID",
			PicURL:  "http://p.qpic.cn/pic_wework/3474110808/7a6344sdadfwehe42060/0",
		},
		Link: &GroupLink{
			Title:  "消息标题",
			PicURL: "https://example.pic.com/path",
			Desc:   "消息描述",
			URL:    "https://example.link.com/path",
		},
		Minip: &GroupMinip{
			Title:      "消息标题",
			PicMediaID: "MEDIA_ID",
			AppID:      "wx8bd80126147df384",
			Page:       "/path/index",
		},
		File: &GroupFile{
			MediaID: "1Yv-zXfHjSjU-7LH-GwtYqDGS-zz6w22KmWAT5COgP7o",
		},
		Video: &GroupVideo{
			MediaID: "MEDIA_ID",
		},
		AgentID: 1000014,
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", EditGroupWelcomeTemplate(params))

	assert.Nil(t, err)
}

func TestGetGroupWelcomeTemplate(t *testing.T) {
	body := []byte(`{"template_id":"msgXXXXXXX"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"text": {
		"content": "文本消息内容"
	},
	"image": {
		"pic_url": "http://p.qpic.cn/pic_wework/XXXXX"
	},
	"link": {
		"title": "消息标题",
		"picurl": "https://example.pic.com/path",
		"desc": "消息描述",
		"url": "https://example.link.com/path"
	},
	"miniprogram": {
		"title": "消息标题",
		"pic_media_id": "MEDIA_ID",
		"appid": "wx8bd80126147df384",
		"page": "/path/index"
	},
	"file": {
		"media_id": "1Yv-zXfHjSjU-7LH-GwtYqDGS-zz6w22KmWAT5COgP7o"
	},
	"video": {
		"media_id": "MEDIA_ID"
	}
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/get?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultGroupWelcomeTemplateGet)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetGroupWelcomeTemplate("msgXXXXXXX", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultGroupWelcomeTemplateGet{
		Text: &GroupText{
			Content: "文本消息内容",
		},
		Image: &GroupImage{
			PicURL: "http://p.qpic.cn/pic_wework/XXXXX",
		},
		Link: &GroupLink{
			Title:  "消息标题",
			PicURL: "https://example.pic.com/path",
			Desc:   "消息描述",
			URL:    "https://example.link.com/path",
		},
		Minip: &GroupMinip{
			Title:      "消息标题",
			PicMediaID: "MEDIA_ID",
			AppID:      "wx8bd80126147df384",
			Page:       "/path/index",
		},
		File: &GroupFile{
			MediaID: "1Yv-zXfHjSjU-7LH-GwtYqDGS-zz6w22KmWAT5COgP7o",
		},
		Video: &GroupVideo{
			MediaID: "MEDIA_ID",
		},
	}, result)
}

func TestDeleteGroupWelcomeTemplate(t *testing.T) {
	body := []byte(`{"template_id":"msgXXXXXXX","agentid":1000014}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/del?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", DeleteGroupWelcomeTemplate("msgXXXXXXX", 1000014))

	assert.Nil(t, err)
}
