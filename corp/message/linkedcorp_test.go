package message

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

func TestSendLinkedcorpText(t *testing.T) {
	body := []byte(`{"touser":["userid1","userid2","CorpId1/userid1","CorpId2/userid2"],"toparty":["partyid1","partyid2","LinkedId1/partyid1","LinkedId2/partyid2"],"totag":["tagid1","tagid2"],"msgtype":"text","agentid":1,"text":{"content":"你的快递已到，请携带工卡前往邮件中心领取。\n出发前可查看<a href=\"http://work.weixin.qq.com\">邮件中心视频实况</a>，聪明避开排队。"}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invaliduser": [
		"userid1",
		"userid2",
		"CorpId1/userid1",
		"CorpId2/userid2"
	],
	"invalidparty": [
		"partyid1",
		"partyid2",
		"LinkedId1/partyid1",
		"LinkedId2/partyid2"
	],
	"invalidtag": [
		"tagid1",
		"tagid2"
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	extra := &LinkedcorpExtra{
		ToUser:  []string{"userid1", "userid2", "CorpId1/userid1", "CorpId2/userid2"},
		ToParty: []string{"partyid1", "partyid2", "LinkedId1/partyid1", "LinkedId2/partyid2"},
		ToTag:   []string{"tagid1", "tagid2"},
	}

	result := new(ResultLinkedcorpSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendLinkedcorpText(1, "你的快递已到，请携带工卡前往邮件中心领取。\n出发前可查看<a href=\"http://work.weixin.qq.com\">邮件中心视频实况</a>，聪明避开排队。", extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultLinkedcorpSend{
		InvalidUser:  []string{"userid1", "userid2", "CorpId1/userid1", "CorpId2/userid2"},
		InvalidParty: []string{"partyid1", "partyid2", "LinkedId1/partyid1", "LinkedId2/partyid2"},
		InvalidTag:   []string{"tagid1", "tagid2"},
	}, result)
}

func TestSendLinkedcorpImage(t *testing.T) {
	body := []byte(`{"touser":["userid1","userid2","CorpId1/userid1","CorpId2/userid2"],"toparty":["partyid1","partyid2","LinkedId1/partyid1","LinkedId2/partyid2"],"totag":["tagid1","tagid2"],"msgtype":"image","agentid":1,"image":{"media_id":"MEDIA_ID"}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invaliduser": [
		"userid1",
		"userid2",
		"CorpId1/userid1",
		"CorpId2/userid2"
	],
	"invalidparty": [
		"partyid1",
		"partyid2",
		"LinkedId1/partyid1",
		"LinkedId2/partyid2"
	],
	"invalidtag": [
		"tagid1",
		"tagid2"
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	extra := &LinkedcorpExtra{
		ToUser:  []string{"userid1", "userid2", "CorpId1/userid1", "CorpId2/userid2"},
		ToParty: []string{"partyid1", "partyid2", "LinkedId1/partyid1", "LinkedId2/partyid2"},
		ToTag:   []string{"tagid1", "tagid2"},
	}

	result := new(ResultLinkedcorpSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendLinkedcorpImage(1, "MEDIA_ID", extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultLinkedcorpSend{
		InvalidUser:  []string{"userid1", "userid2", "CorpId1/userid1", "CorpId2/userid2"},
		InvalidParty: []string{"partyid1", "partyid2", "LinkedId1/partyid1", "LinkedId2/partyid2"},
		InvalidTag:   []string{"tagid1", "tagid2"},
	}, result)
}

func TestSendLinkedcorpVoice(t *testing.T) {
	body := []byte(`{"touser":["userid1","userid2","CorpId1/userid1","CorpId2/userid2"],"toparty":["partyid1","partyid2","LinkedId1/partyid1","LinkedId2/partyid2"],"totag":["tagid1","tagid2"],"msgtype":"voice","agentid":1,"voice":{"media_id":"MEDIA_ID"}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invaliduser": [
		"userid1",
		"userid2",
		"CorpId1/userid1",
		"CorpId2/userid2"
	],
	"invalidparty": [
		"partyid1",
		"partyid2",
		"LinkedId1/partyid1",
		"LinkedId2/partyid2"
	],
	"invalidtag": [
		"tagid1",
		"tagid2"
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	extra := &LinkedcorpExtra{
		ToUser:  []string{"userid1", "userid2", "CorpId1/userid1", "CorpId2/userid2"},
		ToParty: []string{"partyid1", "partyid2", "LinkedId1/partyid1", "LinkedId2/partyid2"},
		ToTag:   []string{"tagid1", "tagid2"},
	}

	result := new(ResultLinkedcorpSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendLinkedcorpVoice(1, "MEDIA_ID", extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultLinkedcorpSend{
		InvalidUser:  []string{"userid1", "userid2", "CorpId1/userid1", "CorpId2/userid2"},
		InvalidParty: []string{"partyid1", "partyid2", "LinkedId1/partyid1", "LinkedId2/partyid2"},
		InvalidTag:   []string{"tagid1", "tagid2"},
	}, result)
}

func TestSendLinkedcorpVideo(t *testing.T) {
	body := []byte(`{"touser":["userid1","userid2","CorpId1/userid1","CorpId2/userid2"],"toparty":["partyid1","partyid2","LinkedId1/partyid1","LinkedId2/partyid2"],"totag":["tagid1","tagid2"],"msgtype":"video","agentid":1,"video":{"media_id":"MEDIA_ID","title":"Title","description":"Description"}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invaliduser": [
		"userid1",
		"userid2",
		"CorpId1/userid1",
		"CorpId2/userid2"
	],
	"invalidparty": [
		"partyid1",
		"partyid2",
		"LinkedId1/partyid1",
		"LinkedId2/partyid2"
	],
	"invalidtag": [
		"tagid1",
		"tagid2"
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	video := &Video{
		MediaID:     "MEDIA_ID",
		Title:       "Title",
		Description: "Description",
	}

	extra := &LinkedcorpExtra{
		ToUser:  []string{"userid1", "userid2", "CorpId1/userid1", "CorpId2/userid2"},
		ToParty: []string{"partyid1", "partyid2", "LinkedId1/partyid1", "LinkedId2/partyid2"},
		ToTag:   []string{"tagid1", "tagid2"},
	}

	result := new(ResultLinkedcorpSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendLinkedcorpVideo(1, video, extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultLinkedcorpSend{
		InvalidUser:  []string{"userid1", "userid2", "CorpId1/userid1", "CorpId2/userid2"},
		InvalidParty: []string{"partyid1", "partyid2", "LinkedId1/partyid1", "LinkedId2/partyid2"},
		InvalidTag:   []string{"tagid1", "tagid2"},
	}, result)
}

func TestSendLinkedcorpFile(t *testing.T) {
	body := []byte(`{"touser":["userid1","userid2","CorpId1/userid1","CorpId2/userid2"],"toparty":["partyid1","partyid2","LinkedId1/partyid1","LinkedId2/partyid2"],"totag":["tagid1","tagid2"],"msgtype":"file","agentid":1,"file":{"media_id":"1Yv-zXfHjSjU-7LH-GwtYqDGS-zz6w22KmWAT5COgP7o"}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invaliduser": [
		"userid1",
		"userid2",
		"CorpId1/userid1",
		"CorpId2/userid2"
	],
	"invalidparty": [
		"partyid1",
		"partyid2",
		"LinkedId1/partyid1",
		"LinkedId2/partyid2"
	],
	"invalidtag": [
		"tagid1",
		"tagid2"
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	extra := &LinkedcorpExtra{
		ToUser:  []string{"userid1", "userid2", "CorpId1/userid1", "CorpId2/userid2"},
		ToParty: []string{"partyid1", "partyid2", "LinkedId1/partyid1", "LinkedId2/partyid2"},
		ToTag:   []string{"tagid1", "tagid2"},
	}

	result := new(ResultLinkedcorpSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendLinkedcorpFile(1, "1Yv-zXfHjSjU-7LH-GwtYqDGS-zz6w22KmWAT5COgP7o", extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultLinkedcorpSend{
		InvalidUser:  []string{"userid1", "userid2", "CorpId1/userid1", "CorpId2/userid2"},
		InvalidParty: []string{"partyid1", "partyid2", "LinkedId1/partyid1", "LinkedId2/partyid2"},
		InvalidTag:   []string{"tagid1", "tagid2"},
	}, result)
}

func TestSendLinkedcorpTextCard(t *testing.T) {
	body := []byte(`{"touser":["userid1","userid2","CorpId1/userid1","CorpId2/userid2"],"toparty":["partyid1","partyid2","LinkedId1/partyid1","LinkedId2/partyid2"],"totag":["tagid1","tagid2"],"msgtype":"textcard","agentid":1,"textcard":{"title":"领奖通知","description":"<div class=\"gray\">2016年9月26日</div> <div class=\"normal\">恭喜你抽中iPhone 7一台，领奖码：xxxx</div><div class=\"highlight\">请于2016年10月10日前联系行政同事领取</div>","url":"URL","btntxt":"更多"}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invaliduser": [
		"userid1",
		"userid2",
		"CorpId1/userid1",
		"CorpId2/userid2"
	],
	"invalidparty": [
		"partyid1",
		"partyid2",
		"LinkedId1/partyid1",
		"LinkedId2/partyid2"
	],
	"invalidtag": [
		"tagid1",
		"tagid2"
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	card := &TextCard{
		Title:       "领奖通知",
		Description: "<div class=\"gray\">2016年9月26日</div> <div class=\"normal\">恭喜你抽中iPhone 7一台，领奖码：xxxx</div><div class=\"highlight\">请于2016年10月10日前联系行政同事领取</div>",
		URL:         "URL",
		BtnTxt:      "更多",
	}

	result := new(ResultLinkedcorpSend)

	extra := &LinkedcorpExtra{
		ToUser:  []string{"userid1", "userid2", "CorpId1/userid1", "CorpId2/userid2"},
		ToParty: []string{"partyid1", "partyid2", "LinkedId1/partyid1", "LinkedId2/partyid2"},
		ToTag:   []string{"tagid1", "tagid2"},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendLinkedcorpTextCard(1, card, extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultLinkedcorpSend{
		InvalidUser:  []string{"userid1", "userid2", "CorpId1/userid1", "CorpId2/userid2"},
		InvalidParty: []string{"partyid1", "partyid2", "LinkedId1/partyid1", "LinkedId2/partyid2"},
		InvalidTag:   []string{"tagid1", "tagid2"},
	}, result)
}

func TestSendLinkedcorpNews(t *testing.T) {
	body := []byte(`{"touser":["userid1","userid2","CorpId1/userid1","CorpId2/userid2"],"toparty":["partyid1","partyid2","LinkedId1/partyid1","LinkedId2/partyid2"],"totag":["tagid1","tagid2"],"msgtype":"news","agentid":1,"news":{"articles":[{"title":"中秋节礼品领取","description":"今年中秋节公司有豪礼相送","url":"URL","picurl":"http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png","btntxt":"更多"}]}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invaliduser": [
		"userid1",
		"userid2",
		"CorpId1/userid1",
		"CorpId2/userid2"
	],
	"invalidparty": [
		"partyid1",
		"partyid2",
		"LinkedId1/partyid1",
		"LinkedId2/partyid2"
	],
	"invalidtag": [
		"tagid1",
		"tagid2"
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	articles := []*NewsArticle{
		{
			Title:       "中秋节礼品领取",
			Description: "今年中秋节公司有豪礼相送",
			URL:         "URL",
			PicURL:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
			BtnTxt:      "更多",
		},
	}

	extra := &LinkedcorpExtra{
		ToUser:  []string{"userid1", "userid2", "CorpId1/userid1", "CorpId2/userid2"},
		ToParty: []string{"partyid1", "partyid2", "LinkedId1/partyid1", "LinkedId2/partyid2"},
		ToTag:   []string{"tagid1", "tagid2"},
	}

	result := new(ResultLinkedcorpSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendLinkedcorpNews(1, articles, extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultLinkedcorpSend{
		InvalidUser:  []string{"userid1", "userid2", "CorpId1/userid1", "CorpId2/userid2"},
		InvalidParty: []string{"partyid1", "partyid2", "LinkedId1/partyid1", "LinkedId2/partyid2"},
		InvalidTag:   []string{"tagid1", "tagid2"},
	}, result)
}

func TestSendLinkedcorpMPNews(t *testing.T) {
	body := []byte(`{"touser":["userid1","userid2","CorpId1/userid1","CorpId2/userid2"],"toparty":["partyid1","partyid2","LinkedId1/partyid1","LinkedId2/partyid2"],"totag":["tagid1","tagid2"],"msgtype":"mpnews","agentid":1,"mpnews":{"articles":[{"title":"Title","thumb_media_id":"MEDIA_ID","author":"Author","content_source_url":"URL","content":"Content","digest":"Digest description"}]}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invaliduser": [
		"userid1",
		"userid2",
		"CorpId1/userid1",
		"CorpId2/userid2"
	],
	"invalidparty": [
		"partyid1",
		"partyid2",
		"LinkedId1/partyid1",
		"LinkedId2/partyid2"
	],
	"invalidtag": [
		"tagid1",
		"tagid2"
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

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

	extra := &LinkedcorpExtra{
		ToUser:  []string{"userid1", "userid2", "CorpId1/userid1", "CorpId2/userid2"},
		ToParty: []string{"partyid1", "partyid2", "LinkedId1/partyid1", "LinkedId2/partyid2"},
		ToTag:   []string{"tagid1", "tagid2"},
	}

	result := new(ResultLinkedcorpSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendLinkedcorpMPNews(1, articles, extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultLinkedcorpSend{
		InvalidUser:  []string{"userid1", "userid2", "CorpId1/userid1", "CorpId2/userid2"},
		InvalidParty: []string{"partyid1", "partyid2", "LinkedId1/partyid1", "LinkedId2/partyid2"},
		InvalidTag:   []string{"tagid1", "tagid2"},
	}, result)
}

func TestSendLinkedcorpMarkdown(t *testing.T) {
	body := []byte(`{"touser":["userid1","userid2","CorpId1/userid1","CorpId2/userid2"],"toparty":["partyid1","partyid2","LinkedId1/partyid1","LinkedId2/partyid2"],"totag":["tagid1","tagid2"],"msgtype":"markdown","agentid":1,"markdown":{"content":"您的会议室已经预定，稍后会同步到邮箱\n>**事项详情**\n>事　项：<font color=\"info\">开会</font>\n>组织者：@miglioguan\n>参与者：@miglioguan、@kunliu、@jamdeezhou、@kanexiong、@kisonwang\n>\n>会议室：<font color=\"info\">广州TIT 1楼 301</font>\n>日　期：<font color=\"warning\">2018年5月18日</font>\n>时　间：<font color=\"comment\">上午9:00-11:00</font>\n>\n>请准时参加会议。\n>\n>如需修改会议信息，请点击：[修改会议信息](https://work.weixin.qq.com)"}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invaliduser": [
		"userid1",
		"userid2",
		"CorpId1/userid1",
		"CorpId2/userid2"
	],
	"invalidparty": [
		"partyid1",
		"partyid2",
		"LinkedId1/partyid1",
		"LinkedId2/partyid2"
	],
	"invalidtag": [
		"tagid1",
		"tagid2"
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	extra := &LinkedcorpExtra{
		ToUser:  []string{"userid1", "userid2", "CorpId1/userid1", "CorpId2/userid2"},
		ToParty: []string{"partyid1", "partyid2", "LinkedId1/partyid1", "LinkedId2/partyid2"},
		ToTag:   []string{"tagid1", "tagid2"},
	}

	result := new(ResultLinkedcorpSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendLinkedcorpMarkdown(1, "您的会议室已经预定，稍后会同步到邮箱\n>**事项详情**\n>事　项：<font color=\"info\">开会</font>\n>组织者：@miglioguan\n>参与者：@miglioguan、@kunliu、@jamdeezhou、@kanexiong、@kisonwang\n>\n>会议室：<font color=\"info\">广州TIT 1楼 301</font>\n>日　期：<font color=\"warning\">2018年5月18日</font>\n>时　间：<font color=\"comment\">上午9:00-11:00</font>\n>\n>请准时参加会议。\n>\n>如需修改会议信息，请点击：[修改会议信息](https://work.weixin.qq.com)", extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultLinkedcorpSend{
		InvalidUser:  []string{"userid1", "userid2", "CorpId1/userid1", "CorpId2/userid2"},
		InvalidParty: []string{"partyid1", "partyid2", "LinkedId1/partyid1", "LinkedId2/partyid2"},
		InvalidTag:   []string{"tagid1", "tagid2"},
	}, result)
}

func TestSendLinkedcorpMinipNotice(t *testing.T) {
	body := []byte(`{"touser":["userid1","userid2","CorpId1/userid1","CorpId2/userid2"],"toparty":["partyid1","partyid2","LinkedId1/partyid1","LinkedId2/partyid2"],"totag":["tagid1","tagid2"],"msgtype":"miniprogram_notice","miniprogram_notice":{"appid":"wx123123123123123","page":"pages/index?userid=zhangsan&orderid=123123123","title":"会议室预订成功通知","description":"4月27日 16:16","emphasis_first_item":true,"content_item":[{"key":"会议室","value":"402"},{"key":"会议地点","value":"广州TIT-402会议室"},{"key":"会议时间","value":"2018年8月1日 09:00-09:30"},{"key":"参与人员","value":"周剑轩"}]}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invaliduser": [
		"userid1",
		"userid2",
		"CorpId1/userid1",
		"CorpId2/userid2"
	],
	"invalidparty": [
		"partyid1",
		"partyid2",
		"LinkedId1/partyid1",
		"LinkedId2/partyid2"
	],
	"invalidtag": [
		"tagid1",
		"tagid2"
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	notice := &MinipNotice{
		AppID:             "wx123123123123123",
		Page:              "pages/index?userid=zhangsan&orderid=123123123",
		Title:             "会议室预订成功通知",
		Description:       "4月27日 16:16",
		EmphasisFirstItem: true,
		ContentItem: []*MsgKV{
			{
				Key:   "会议室",
				Value: "402",
			},
			{
				Key:   "会议地点",
				Value: "广州TIT-402会议室",
			},
			{
				Key:   "会议时间",
				Value: "2018年8月1日 09:00-09:30",
			},
			{
				Key:   "参与人员",
				Value: "周剑轩",
			},
		},
	}

	extra := &LinkedcorpExtra{
		ToUser:  []string{"userid1", "userid2", "CorpId1/userid1", "CorpId2/userid2"},
		ToParty: []string{"partyid1", "partyid2", "LinkedId1/partyid1", "LinkedId2/partyid2"},
		ToTag:   []string{"tagid1", "tagid2"},
	}

	result := new(ResultLinkedcorpSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendLinkedcorpMinipNotice(notice, extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultLinkedcorpSend{
		InvalidUser:  []string{"userid1", "userid2", "CorpId1/userid1", "CorpId2/userid2"},
		InvalidParty: []string{"partyid1", "partyid2", "LinkedId1/partyid1", "LinkedId2/partyid2"},
		InvalidTag:   []string{"tagid1", "tagid2"},
	}, result)
}
