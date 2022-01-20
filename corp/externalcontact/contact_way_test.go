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
	"github.com/shenghui0779/gochat/mock"
	"github.com/shenghui0779/gochat/wx"
)

func TestListFollowUser(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"follow_user": [
		"zhangsan",
		"lissi"
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_follow_user_list?access_token=ACCESS_TOKEN", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultFollowUserList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListFollowUser(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultFollowUserList{
		FollowUser: []string{"zhangsan", "lissi"},
	}, result)
}

func TestAddContactWay(t *testing.T) {
	body := []byte(`{"type":1,"scene":1,"style":1,"remark":"渠道客户","skip_verify":true,"state":"teststate","user":["zhangsan","lisi","wangwu"],"party":[2,3],"is_temp":true,"expires_in":86400,"chat_expires_in":86400,"unionid":"oxTWIuGaIt6gTKsQRLau2M0AAAA","conclusions":{"text":{"content":"文本消息内容"},"image":{"media_id":"MEDIA_ID"},"link":{"title":"消息标题","picurl":"https://example.pic.com/path","desc":"消息描述","url":"https://example.link.com/path"},"miniprogram":{"title":"消息标题","pic_media_id":"MEDIA_ID","appid":"wx8bd80126147dfAAA","page":"/path/index.html"}}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"config_id": "42b34949e138eb6e027c123cba77fAAA",
	"qr_code": "http://p.qpic.cn/wwhead/duc2TvpEgSdicZ9RrdUtBkv2UiaA/0"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_contact_way?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsContactWayAdd{
		Type:          1,
		Scene:         1,
		Style:         1,
		Remark:        "渠道客户",
		SkipVerify:    true,
		State:         "teststate",
		User:          []string{"zhangsan", "lisi", "wangwu"},
		Party:         []int64{2, 3},
		IsTemp:        true,
		ExpiresIn:     86400,
		ChatExpiresIn: 86400,
		UnionID:       "oxTWIuGaIt6gTKsQRLau2M0AAAA",
		Conclusions: &Conclusions{
			Text: &TextConclusion{
				Content: "文本消息内容",
			},
			Image: &ImageConclusion{
				MediaID: "MEDIA_ID",
			},
			Link: &LinkConclusion{
				Title:  "消息标题",
				PicURL: "https://example.pic.com/path",
				Desc:   "消息描述",
				URL:    "https://example.link.com/path",
			},
			Minip: &MinipConclusion{
				Title:      "消息标题",
				PicMediaID: "MEDIA_ID",
				AppID:      "wx8bd80126147dfAAA",
				Page:       "/path/index.html",
			},
		},
	}

	result := new(ResultContactWayAdd)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", AddContactWay(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultContactWayAdd{
		ConfigID: "42b34949e138eb6e027c123cba77fAAA",
		QRCode:   "http://p.qpic.cn/wwhead/duc2TvpEgSdicZ9RrdUtBkv2UiaA/0",
	}, result)
}

func TestUpdateContactWay(t *testing.T) {
	body := []byte(`{"config_id":"42b34949e138eb6e027c123cba77fAAA","remark":"渠道客户","skip_verify":true,"style":1,"state":"teststate","user":["zhangsan","lisi","wangwu"],"party":[2,3],"expires_in":86400,"chat_expires_in":86400,"unionid":"oxTWIuGaIt6gTKsQRLau2M0AAAA","conclusions":{"text":{"content":"文本消息内容"},"image":{"media_id":"MEDIA_ID"},"link":{"title":"消息标题","picurl":"https://example.pic.com/path","desc":"消息描述","url":"https://example.link.com/path"},"miniprogram":{"title":"消息标题","pic_media_id":"MEDIA_ID","appid":"wx8bd80126147dfAAA","page":"/path/index"}}}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/update_contact_way?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsContactWayUpdate{
		ConfigID:      "42b34949e138eb6e027c123cba77fAAA",
		Remark:        "渠道客户",
		SkipVerify:    true,
		Style:         1,
		State:         "teststate",
		User:          []string{"zhangsan", "lisi", "wangwu"},
		Party:         []int64{2, 3},
		ExpiresIn:     86400,
		ChatExpiresIn: 86400,
		UnionID:       "oxTWIuGaIt6gTKsQRLau2M0AAAA",
		Conclusions: &Conclusions{
			Text: &TextConclusion{
				Content: "文本消息内容",
			},
			Image: &ImageConclusion{
				MediaID: "MEDIA_ID",
			},
			Link: &LinkConclusion{
				Title:  "消息标题",
				PicURL: "https://example.pic.com/path",
				Desc:   "消息描述",
				URL:    "https://example.link.com/path",
			},
			Minip: &MinipConclusion{
				Title:      "消息标题",
				PicMediaID: "MEDIA_ID",
				AppID:      "wx8bd80126147dfAAA",
				Page:       "/path/index",
			},
		},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UpdateContactWay(params))

	assert.Nil(t, err)
}

func TestGetContactWay(t *testing.T) {
	body := []byte(`{"config_id":"42b34949e138eb6e027c123cba77fad7"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"contact_way": {
		"config_id": "42b34949e138eb6e027c123cba77fAAA",
		"type": 1,
		"scene": 1,
		"style": 2,
		"remark": "test remark",
		"skip_verify": true,
		"state": "teststate",
		"qr_code": "http://p.qpic.cn/wwhead/duc2TvpEgSdicZ9RrdUtBkv2UiaA/0",
		"user": [
			"zhangsan",
			"lisi",
			"wangwu"
		],
		"party": [
			2,
			3
		],
		"is_temp": true,
		"expires_in": 86400,
		"chat_expires_in": 86400,
		"unionid": "oxTWIuGaIt6gTKsQRLau2M0AAAA",
		"conclusions": {
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
				"appid": "wx8bd80126147dfAAA",
				"page": "/path/index"
			}
		}
	}
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_contact_way?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultContactWayGet)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetContactWay("42b34949e138eb6e027c123cba77fad7", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultContactWayGet{
		ContactWay: &ContactWay{
			ConfigID:      "42b34949e138eb6e027c123cba77fAAA",
			Type:          1,
			Scene:         1,
			Style:         2,
			Remark:        "test remark",
			SkipVerify:    true,
			State:         "teststate",
			QRCode:        "http://p.qpic.cn/wwhead/duc2TvpEgSdicZ9RrdUtBkv2UiaA/0",
			User:          []string{"zhangsan", "lisi", "wangwu"},
			Party:         []int64{2, 3},
			IsTemp:        true,
			ExpiresIn:     86400,
			ChatExpiresIn: 86400,
			UnionID:       "oxTWIuGaIt6gTKsQRLau2M0AAAA",
			Conclusions: &Conclusions{
				Text: &TextConclusion{
					Content: "文本消息内容",
				},
				Image: &ImageConclusion{
					MediaID: "http://p.qpic.cn/pic_wework/XXXXX",
				},
				Link: &LinkConclusion{
					Title:  "消息标题",
					PicURL: "https://example.pic.com/path",
					Desc:   "消息描述",
					URL:    "https://example.link.com/path",
				},
				Minip: &MinipConclusion{
					Title:      "消息标题",
					PicMediaID: "MEDIA_ID",
					AppID:      "wx8bd80126147dfAAA",
					Page:       "/path/index",
				},
			},
		},
	}, result)
}

func TestListContactWay(t *testing.T) {
	body := []byte(`{"start_time":1622476800,"end_time":1625068800,"cursor":"CURSOR","limit":1000}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"contact_way": [
		{
			"config_id": "534b63270045c9ABiKEE814ef56d91c62f"
		},
		{
			"config_id": "87bBiKEE811c62f63270041c62f5c9A4ef"
		}
	],
	"next_cursor": "NEXT_CURSOR"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/list_contact_way?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsContactWayList{
		StartTime: 1622476800,
		EntTime:   1625068800,
		Cursor:    "CURSOR",
		Limit:     1000,
	}

	result := new(ResultContactWayList)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ListContactWay(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultContactWayList{
		ContactWay: []*ContactWayListData{
			{
				ConfigID: "534b63270045c9ABiKEE814ef56d91c62f",
			},
			{
				ConfigID: "87bBiKEE811c62f63270041c62f5c9A4ef",
			},
		},
		NextCursor: "NEXT_CURSOR",
	}, result)
}

func TestDeleteContactWay(t *testing.T) {
	body := []byte(`{"config_id":"42b34949e138eb6e027c123cba77fAAA"}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/del_contact_way?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", DeleteContactWay("42b34949e138eb6e027c123cba77fAAA"))

	assert.Nil(t, err)
}

func TestCloseTempChat(t *testing.T) {
	body := []byte(`{"userid":"zhangyisheng","external_userid":"woAJ2GCAAAXtWyujaWJHDDGi0mACHAAA"}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/close_temp_chat?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsTempChatClose{
		UserID:         "zhangyisheng",
		ExternalUserID: "woAJ2GCAAAXtWyujaWJHDDGi0mACHAAA",
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", CloseTempChat(params))

	assert.Nil(t, err)
}
