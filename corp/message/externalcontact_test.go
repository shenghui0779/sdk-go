package message

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

func TestSendExternalContactText(t *testing.T) {
	body := []byte(`{"to_external_user":["external_userid1","external_userid2"],"to_parent_userid":["parent_userid1","parent_userid2"],"to_student_userid":["student_userid1","student_userid2"],"to_party":["partyid1","partyid2"],"msgtype":"text","agentid":1,"text":{"content":"你的快递已到，请携带工卡前往邮件中心领取。\n出发前可查看<a href=\"http://work.weixin.qq.com\">邮件中心视频实况</a>，聪明避开排队。"},"duplicate_check_interval":1800}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invalid_external_user": [
		"external_userid1"
	],
	"invalid_parent_userid": [
		"parent_userid1"
	],
	"invalid_student_userid": [
		"student_userid1"
	],
	"invalid_party": [
		"party1"
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	extra := &ExternalContactExtra{
		ToExternalUser:         []string{"external_userid1", "external_userid2"},
		ToParentUserID:         []string{"parent_userid1", "parent_userid2"},
		ToStudentUserID:        []string{"student_userid1", "student_userid2"},
		ToParty:                []string{"partyid1", "partyid2"},
		DuplicateCheckInterval: 1800,
	}

	result := new(ResultExternalContactSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendExternalContactText(1, "你的快递已到，请携带工卡前往邮件中心领取。\n出发前可查看<a href=\"http://work.weixin.qq.com\">邮件中心视频实况</a>，聪明避开排队。", extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultExternalContactSend{
		InvalidExternalUser:  []string{"external_userid1"},
		InvalidParentUserID:  []string{"parent_userid1"},
		InvalidStudentUserID: []string{"student_userid1"},
		InvalidParty:         []string{"party1"},
	}, result)
}

func TestSendExternalContactImage(t *testing.T) {
	body := []byte(`{"to_external_user":["external_userid1","external_userid2"],"to_parent_userid":["parent_userid1","parent_userid2"],"to_student_userid":["student_userid1","student_userid2"],"to_party":["partyid1","partyid2"],"msgtype":"image","agentid":1,"image":{"media_id":"MEDIA_ID"},"duplicate_check_interval":1800}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invalid_external_user": [
		"external_userid1"
	],
	"invalid_parent_userid": [
		"parent_userid1"
	],
	"invalid_student_userid": [
		"student_userid1"
	],
	"invalid_party": [
		"party1"
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	extra := &ExternalContactExtra{
		ToExternalUser:         []string{"external_userid1", "external_userid2"},
		ToParentUserID:         []string{"parent_userid1", "parent_userid2"},
		ToStudentUserID:        []string{"student_userid1", "student_userid2"},
		ToParty:                []string{"partyid1", "partyid2"},
		DuplicateCheckInterval: 1800,
	}

	result := new(ResultExternalContactSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendExternalContactImage(1, "MEDIA_ID", extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultExternalContactSend{
		InvalidExternalUser:  []string{"external_userid1"},
		InvalidParentUserID:  []string{"parent_userid1"},
		InvalidStudentUserID: []string{"student_userid1"},
		InvalidParty:         []string{"party1"},
	}, result)
}

func TestSendExternalContactVoice(t *testing.T) {
	body := []byte(`{"to_external_user":["external_userid1","external_userid2"],"to_parent_userid":["parent_userid1","parent_userid2"],"to_student_userid":["student_userid1","student_userid2"],"to_party":["partyid1","partyid2"],"msgtype":"voice","agentid":1,"voice":{"media_id":"MEDIA_ID"},"duplicate_check_interval":1800}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invalid_external_user": [
		"external_userid1"
	],
	"invalid_parent_userid": [
		"parent_userid1"
	],
	"invalid_student_userid": [
		"student_userid1"
	],
	"invalid_party": [
		"party1"
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	extra := &ExternalContactExtra{
		ToExternalUser:         []string{"external_userid1", "external_userid2"},
		ToParentUserID:         []string{"parent_userid1", "parent_userid2"},
		ToStudentUserID:        []string{"student_userid1", "student_userid2"},
		ToParty:                []string{"partyid1", "partyid2"},
		DuplicateCheckInterval: 1800,
	}

	result := new(ResultExternalContactSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendExternalContactVoice(1, "MEDIA_ID", extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultExternalContactSend{
		InvalidExternalUser:  []string{"external_userid1"},
		InvalidParentUserID:  []string{"parent_userid1"},
		InvalidStudentUserID: []string{"student_userid1"},
		InvalidParty:         []string{"party1"},
	}, result)
}

func TestSendExternalContactVideo(t *testing.T) {
	body := []byte(`{"to_external_user":["external_userid1","external_userid2"],"to_parent_userid":["parent_userid1","parent_userid2"],"to_student_userid":["student_userid1","student_userid2"],"to_party":["partyid1","partyid2"],"msgtype":"video","agentid":1,"video":{"media_id":"MEDIA_ID","title":"Title","description":"Description"},"duplicate_check_interval":1800}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invalid_external_user": [
		"external_userid1"
	],
	"invalid_parent_userid": [
		"parent_userid1"
	],
	"invalid_student_userid": [
		"student_userid1"
	],
	"invalid_party": [
		"party1"
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	video := &Video{
		MediaID:     "MEDIA_ID",
		Title:       "Title",
		Description: "Description",
	}

	extra := &ExternalContactExtra{
		ToExternalUser:         []string{"external_userid1", "external_userid2"},
		ToParentUserID:         []string{"parent_userid1", "parent_userid2"},
		ToStudentUserID:        []string{"student_userid1", "student_userid2"},
		ToParty:                []string{"partyid1", "partyid2"},
		DuplicateCheckInterval: 1800,
	}

	result := new(ResultExternalContactSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendExternalContactVideo(1, video, extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultExternalContactSend{
		InvalidExternalUser:  []string{"external_userid1"},
		InvalidParentUserID:  []string{"parent_userid1"},
		InvalidStudentUserID: []string{"student_userid1"},
		InvalidParty:         []string{"party1"},
	}, result)
}

func TestSendExternalContactFile(t *testing.T) {
	body := []byte(`{"to_external_user":["external_userid1","external_userid2"],"to_parent_userid":["parent_userid1","parent_userid2"],"to_student_userid":["student_userid1","student_userid2"],"to_party":["partyid1","partyid2"],"msgtype":"file","agentid":1,"file":{"media_id":"1Yv-zXfHjSjU-7LH-GwtYqDGS-zz6w22KmWAT5COgP7o"},"duplicate_check_interval":1800}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invalid_external_user": [
		"external_userid1"
	],
	"invalid_parent_userid": [
		"parent_userid1"
	],
	"invalid_student_userid": [
		"student_userid1"
	],
	"invalid_party": [
		"party1"
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	extra := &ExternalContactExtra{
		ToExternalUser:         []string{"external_userid1", "external_userid2"},
		ToParentUserID:         []string{"parent_userid1", "parent_userid2"},
		ToStudentUserID:        []string{"student_userid1", "student_userid2"},
		ToParty:                []string{"partyid1", "partyid2"},
		DuplicateCheckInterval: 1800,
	}

	result := new(ResultExternalContactSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendExternalContactFile(1, "1Yv-zXfHjSjU-7LH-GwtYqDGS-zz6w22KmWAT5COgP7o", extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultExternalContactSend{
		InvalidExternalUser:  []string{"external_userid1"},
		InvalidParentUserID:  []string{"parent_userid1"},
		InvalidStudentUserID: []string{"student_userid1"},
		InvalidParty:         []string{"party1"},
	}, result)
}

func TestSendExternalContactNews(t *testing.T) {
	body := []byte(`{"to_external_user":["external_userid1","external_userid2"],"to_parent_userid":["parent_userid1","parent_userid2"],"to_student_userid":["student_userid1","student_userid2"],"to_party":["partyid1","partyid2"],"msgtype":"news","agentid":1,"news":{"articles":[{"title":"中秋节礼品领取","description":"今年中秋节公司有豪礼相送","url":"URL","picurl":"http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png"}]},"duplicate_check_interval":1800}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invalid_external_user": [
		"external_userid1"
	],
	"invalid_parent_userid": [
		"parent_userid1"
	],
	"invalid_student_userid": [
		"student_userid1"
	],
	"invalid_party": [
		"party1"
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	articles := []*NewsArticle{
		{
			Title:       "中秋节礼品领取",
			Description: "今年中秋节公司有豪礼相送",
			URL:         "URL",
			PicURL:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
		},
	}

	extra := &ExternalContactExtra{
		ToExternalUser:         []string{"external_userid1", "external_userid2"},
		ToParentUserID:         []string{"parent_userid1", "parent_userid2"},
		ToStudentUserID:        []string{"student_userid1", "student_userid2"},
		ToParty:                []string{"partyid1", "partyid2"},
		DuplicateCheckInterval: 1800,
	}

	result := new(ResultExternalContactSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendExternalContactNews(1, articles, extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultExternalContactSend{
		InvalidExternalUser:  []string{"external_userid1"},
		InvalidParentUserID:  []string{"parent_userid1"},
		InvalidStudentUserID: []string{"student_userid1"},
		InvalidParty:         []string{"party1"},
	}, result)
}

func TestSendExternalContactMPNews(t *testing.T) {
	body := []byte(`{"to_external_user":["external_userid1","external_userid2"],"to_parent_userid":["parent_userid1","parent_userid2"],"to_student_userid":["student_userid1","student_userid2"],"to_party":["partyid1","partyid2"],"msgtype":"mpnews","agentid":1,"mpnews":{"articles":[{"title":"Title","thumb_media_id":"MEDIA_ID","author":"Author","content_source_url":"URL","content":"Content","digest":"Digest description"}]},"duplicate_check_interval":1800}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invalid_external_user": [
		"external_userid1"
	],
	"invalid_parent_userid": [
		"parent_userid1"
	],
	"invalid_student_userid": [
		"student_userid1"
	],
	"invalid_party": [
		"party1"
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	articles := []*MPNewsArticle{
		{
			Title:            "Title",
			ThumbMediaID:     "MEDIA_ID",
			Author:           "Author",
			ContentSourceURL: "URL",
			Content:          "Content",
			Digest:           "Digest description",
		},
	}

	extra := &ExternalContactExtra{
		ToExternalUser:         []string{"external_userid1", "external_userid2"},
		ToParentUserID:         []string{"parent_userid1", "parent_userid2"},
		ToStudentUserID:        []string{"student_userid1", "student_userid2"},
		ToParty:                []string{"partyid1", "partyid2"},
		DuplicateCheckInterval: 1800,
	}

	result := new(ResultExternalContactSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendExternalContactMPNews(1, articles, extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultExternalContactSend{
		InvalidExternalUser:  []string{"external_userid1"},
		InvalidParentUserID:  []string{"parent_userid1"},
		InvalidStudentUserID: []string{"student_userid1"},
		InvalidParty:         []string{"party1"},
	}, result)
}

func TestSendExternalContactMiniprogram(t *testing.T) {
	body := []byte(`{"to_external_user":["external_userid1","external_userid2"],"to_parent_userid":["parent_userid1","parent_userid2"],"to_student_userid":["student_userid1","student_userid2"],"to_party":["partyid1","partyid2"],"msgtype":"miniprogram","agentid":1,"miniprogram":{"appid":"APPID","title":"欢迎报名夏令营","thumb_media_id":"MEDIA_ID","pagepath":"PAGE_PATH"},"duplicate_check_interval":1800}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invalid_external_user": [
		"external_userid1"
	],
	"invalid_parent_userid": [
		"parent_userid1"
	],
	"invalid_student_userid": [
		"student_userid1"
	],
	"invalid_party": [
		"party1"
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	minip := &Miniprogram{
		AppID:        "APPID",
		Title:        "欢迎报名夏令营",
		ThumbMediaID: "MEDIA_ID",
		Pagepath:     "PAGE_PATH",
	}

	extra := &ExternalContactExtra{
		ToExternalUser:         []string{"external_userid1", "external_userid2"},
		ToParentUserID:         []string{"parent_userid1", "parent_userid2"},
		ToStudentUserID:        []string{"student_userid1", "student_userid2"},
		ToParty:                []string{"partyid1", "partyid2"},
		DuplicateCheckInterval: 1800,
	}

	result := new(ResultExternalContactSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendExternalContactMiniprogram(1, minip, extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultExternalContactSend{
		InvalidExternalUser:  []string{"external_userid1"},
		InvalidParentUserID:  []string{"parent_userid1"},
		InvalidStudentUserID: []string{"student_userid1"},
		InvalidParty:         []string{"party1"},
	}, result)
}
