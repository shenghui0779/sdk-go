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

func TestSendText(t *testing.T) {
	body := []byte(`{"touser":"UserID1|UserID2|UserID3","toparty":"PartyID1|PartyID2","totag":"TagID1|TagID2","msgtype":"text","agentid":1,"text":{"content":"你的快递已到，请携带工卡前往邮件中心领取。\n出发前可查看<a href=\"http://work.weixin.qq.com\">邮件中心视频实况</a>，聪明避开排队。"},"duplicate_check_interval":1800}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invaliduser": "userid1|userid2",
	"invalidparty": "partyid1|partyid2",
	"invalidtag": "tagid1|tagid2",
	"msgid": "xxxx",
	"response_code": "xyzxyz"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	extra := &MsgExtra{
		ToUser:                 "UserID1|UserID2|UserID3",
		ToParty:                "PartyID1|PartyID2",
		ToTag:                  "TagID1|TagID2",
		DuplicateCheckInterval: 1800,
	}

	result := new(ResultMsgSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendText(1, "你的快递已到，请携带工卡前往邮件中心领取。\n出发前可查看<a href=\"http://work.weixin.qq.com\">邮件中心视频实况</a>，聪明避开排队。", extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMsgSend{
		InvalidUser:  "userid1|userid2",
		InvalidParty: "partyid1|partyid2",
		InvalidTag:   "tagid1|tagid2",
		MsgID:        "xxxx",
		ResponseCode: "xyzxyz",
	}, result)
}

func TestSendImage(t *testing.T) {
	body := []byte(`{"touser":"UserID1|UserID2|UserID3","toparty":"PartyID1|PartyID2","totag":"TagID1|TagID2","msgtype":"image","agentid":1,"image":{"media_id":"MEDIA_ID"},"duplicate_check_interval":1800}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invaliduser": "userid1|userid2",
	"invalidparty": "partyid1|partyid2",
	"invalidtag": "tagid1|tagid2",
	"msgid": "xxxx",
	"response_code": "xyzxyz"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	extra := &MsgExtra{
		ToUser:                 "UserID1|UserID2|UserID3",
		ToParty:                "PartyID1|PartyID2",
		ToTag:                  "TagID1|TagID2",
		DuplicateCheckInterval: 1800,
	}

	result := new(ResultMsgSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendImage(1, "MEDIA_ID", extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMsgSend{
		InvalidUser:  "userid1|userid2",
		InvalidParty: "partyid1|partyid2",
		InvalidTag:   "tagid1|tagid2",
		MsgID:        "xxxx",
		ResponseCode: "xyzxyz",
	}, result)
}

func TestSendVoice(t *testing.T) {
	body := []byte(`{"touser":"UserID1|UserID2|UserID3","toparty":"PartyID1|PartyID2","totag":"TagID1|TagID2","msgtype":"voice","agentid":1,"voice":{"media_id":"MEDIA_ID"},"duplicate_check_interval":1800}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invaliduser": "userid1|userid2",
	"invalidparty": "partyid1|partyid2",
	"invalidtag": "tagid1|tagid2",
	"msgid": "xxxx",
	"response_code": "xyzxyz"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	extra := &MsgExtra{
		ToUser:                 "UserID1|UserID2|UserID3",
		ToParty:                "PartyID1|PartyID2",
		ToTag:                  "TagID1|TagID2",
		DuplicateCheckInterval: 1800,
	}

	result := new(ResultMsgSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendVoice(1, "MEDIA_ID", extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMsgSend{
		InvalidUser:  "userid1|userid2",
		InvalidParty: "partyid1|partyid2",
		InvalidTag:   "tagid1|tagid2",
		MsgID:        "xxxx",
		ResponseCode: "xyzxyz",
	}, result)
}

func TestSendVideo(t *testing.T) {
	body := []byte(`{"touser":"UserID1|UserID2|UserID3","toparty":"PartyID1|PartyID2","totag":"TagID1|TagID2","msgtype":"video","agentid":1,"video":{"media_id":"MEDIA_ID","title":"Title","description":"Description"},"duplicate_check_interval":1800}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invaliduser": "userid1|userid2",
	"invalidparty": "partyid1|partyid2",
	"invalidtag": "tagid1|tagid2",
	"msgid": "xxxx",
	"response_code": "xyzxyz"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	video := &Video{
		MediaID:     "MEDIA_ID",
		Title:       "Title",
		Description: "Description",
	}

	extra := &MsgExtra{
		ToUser:                 "UserID1|UserID2|UserID3",
		ToParty:                "PartyID1|PartyID2",
		ToTag:                  "TagID1|TagID2",
		DuplicateCheckInterval: 1800,
	}

	result := new(ResultMsgSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendVideo(1, video, extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMsgSend{
		InvalidUser:  "userid1|userid2",
		InvalidParty: "partyid1|partyid2",
		InvalidTag:   "tagid1|tagid2",
		MsgID:        "xxxx",
		ResponseCode: "xyzxyz",
	}, result)
}

func TestSendFile(t *testing.T) {
	body := []byte(`{"touser":"UserID1|UserID2|UserID3","toparty":"PartyID1|PartyID2","totag":"TagID1|TagID2","msgtype":"file","agentid":1,"file":{"media_id":"1Yv-zXfHjSjU-7LH-GwtYqDGS-zz6w22KmWAT5COgP7o"},"duplicate_check_interval":1800}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invaliduser": "userid1|userid2",
	"invalidparty": "partyid1|partyid2",
	"invalidtag": "tagid1|tagid2",
	"msgid": "xxxx",
	"response_code": "xyzxyz"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	extra := &MsgExtra{
		ToUser:                 "UserID1|UserID2|UserID3",
		ToParty:                "PartyID1|PartyID2",
		ToTag:                  "TagID1|TagID2",
		DuplicateCheckInterval: 1800,
	}

	result := new(ResultMsgSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendFile(1, "1Yv-zXfHjSjU-7LH-GwtYqDGS-zz6w22KmWAT5COgP7o", extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMsgSend{
		InvalidUser:  "userid1|userid2",
		InvalidParty: "partyid1|partyid2",
		InvalidTag:   "tagid1|tagid2",
		MsgID:        "xxxx",
		ResponseCode: "xyzxyz",
	}, result)
}

func TestSendTextCard(t *testing.T) {
	body := []byte(`{"touser":"UserID1|UserID2|UserID3","toparty":"PartyID1|PartyID2","totag":"TagID1|TagID2","msgtype":"textcard","agentid":1,"textcard":{"title":"领奖通知","description":"<div class=\"gray\">2016年9月26日</div> <div class=\"normal\">恭喜你抽中iPhone 7一台，领奖码：xxxx</div><div class=\"highlight\">请于2016年10月10日前联系行政同事领取</div>","url":"URL","btntxt":"更多"},"duplicate_check_interval":1800}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invaliduser": "userid1|userid2",
	"invalidparty": "partyid1|partyid2",
	"invalidtag": "tagid1|tagid2",
	"msgid": "xxxx",
	"response_code": "xyzxyz"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	card := &TextCard{
		Title:       "领奖通知",
		Description: "<div class=\"gray\">2016年9月26日</div> <div class=\"normal\">恭喜你抽中iPhone 7一台，领奖码：xxxx</div><div class=\"highlight\">请于2016年10月10日前联系行政同事领取</div>",
		URL:         "URL",
		BtnTxt:      "更多",
	}

	extra := &MsgExtra{
		ToUser:                 "UserID1|UserID2|UserID3",
		ToParty:                "PartyID1|PartyID2",
		ToTag:                  "TagID1|TagID2",
		DuplicateCheckInterval: 1800,
	}

	result := new(ResultMsgSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendTextCard(1, card, extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMsgSend{
		InvalidUser:  "userid1|userid2",
		InvalidParty: "partyid1|partyid2",
		InvalidTag:   "tagid1|tagid2",
		MsgID:        "xxxx",
		ResponseCode: "xyzxyz",
	}, result)
}

func TestSendNews(t *testing.T) {
	body := []byte(`{"touser":"UserID1|UserID2|UserID3","toparty":"PartyID1|PartyID2","totag":"TagID1|TagID2","msgtype":"news","agentid":1,"news":{"articles":[{"title":"中秋节礼品领取","description":"今年中秋节公司有豪礼相送","url":"URL","picurl":"http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png","appid":"wx123123123123123","pagepath":"pages/index?userid=zhangsan&orderid=123123123"}]},"duplicate_check_interval":1800}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invaliduser": "userid1|userid2",
	"invalidparty": "partyid1|partyid2",
	"invalidtag": "tagid1|tagid2",
	"msgid": "xxxx",
	"response_code": "xyzxyz"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	articles := []*NewsArticle{
		{
			Title:       "中秋节礼品领取",
			Description: "今年中秋节公司有豪礼相送",
			URL:         "URL",
			PicURL:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
			AppID:       "wx123123123123123",
			Pagepath:    "pages/index?userid=zhangsan&orderid=123123123",
		},
	}

	extra := &MsgExtra{
		ToUser:                 "UserID1|UserID2|UserID3",
		ToParty:                "PartyID1|PartyID2",
		ToTag:                  "TagID1|TagID2",
		DuplicateCheckInterval: 1800,
	}

	result := new(ResultMsgSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendNews(1, articles, extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMsgSend{
		InvalidUser:  "userid1|userid2",
		InvalidParty: "partyid1|partyid2",
		InvalidTag:   "tagid1|tagid2",
		MsgID:        "xxxx",
		ResponseCode: "xyzxyz",
	}, result)
}

func TestSendMPNews(t *testing.T) {
	body := []byte(`{"touser":"UserID1|UserID2|UserID3","toparty":"PartyID1|PartyID2","totag":"TagID1|TagID2","msgtype":"mpnews","agentid":1,"mpnews":{"articles":[{"title":"Title","thumb_media_id":"MEDIA_ID","author":"Author","content_source_url":"URL","content":"Content","digest":"Digest description"}]},"duplicate_check_interval":1800}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invaliduser": "userid1|userid2",
	"invalidparty": "partyid1|partyid2",
	"invalidtag": "tagid1|tagid2",
	"msgid": "xxxx",
	"response_code": "xyzxyz"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

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

	extra := &MsgExtra{
		ToUser:                 "UserID1|UserID2|UserID3",
		ToParty:                "PartyID1|PartyID2",
		ToTag:                  "TagID1|TagID2",
		DuplicateCheckInterval: 1800,
	}

	result := new(ResultMsgSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendMPNews(1, articles, extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMsgSend{
		InvalidUser:  "userid1|userid2",
		InvalidParty: "partyid1|partyid2",
		InvalidTag:   "tagid1|tagid2",
		MsgID:        "xxxx",
		ResponseCode: "xyzxyz",
	}, result)
}

func TestSendMarkdown(t *testing.T) {
	body := []byte(`{"touser":"UserID1|UserID2|UserID3","toparty":"PartyID1|PartyID2","totag":"TagID1|TagID2","msgtype":"markdown","agentid":1,"markdown":{"content":"您的会议室已经预定，稍后会同步到邮箱\n>**事项详情**\n>事　项：<font color=\"info\">开会</font>\n>组织者：@miglioguan\n>参与者：@miglioguan、@kunliu、@jamdeezhou、@kanexiong、@kisonwang\n>\n>会议室：<font color=\"info\">广州TIT 1楼 301</font>\n>日　期：<font color=\"warning\">2018年5月18日</font>\n>时　间：<font color=\"comment\">上午9:00-11:00</font>\n>\n>请准时参加会议。\n>\n>如需修改会议信息，请点击：[修改会议信息](https://work.weixin.qq.com)"},"duplicate_check_interval":1800}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invaliduser": "userid1|userid2",
	"invalidparty": "partyid1|partyid2",
	"invalidtag": "tagid1|tagid2",
	"msgid": "xxxx",
	"response_code": "xyzxyz"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	extra := &MsgExtra{
		ToUser:                 "UserID1|UserID2|UserID3",
		ToParty:                "PartyID1|PartyID2",
		ToTag:                  "TagID1|TagID2",
		DuplicateCheckInterval: 1800,
	}

	result := new(ResultMsgSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendMarkdown(1, "您的会议室已经预定，稍后会同步到邮箱\n>**事项详情**\n>事　项：<font color=\"info\">开会</font>\n>组织者：@miglioguan\n>参与者：@miglioguan、@kunliu、@jamdeezhou、@kanexiong、@kisonwang\n>\n>会议室：<font color=\"info\">广州TIT 1楼 301</font>\n>日　期：<font color=\"warning\">2018年5月18日</font>\n>时　间：<font color=\"comment\">上午9:00-11:00</font>\n>\n>请准时参加会议。\n>\n>如需修改会议信息，请点击：[修改会议信息](https://work.weixin.qq.com)", extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMsgSend{
		InvalidUser:  "userid1|userid2",
		InvalidParty: "partyid1|partyid2",
		InvalidTag:   "tagid1|tagid2",
		MsgID:        "xxxx",
		ResponseCode: "xyzxyz",
	}, result)
}

func TestSendMinipNotice(t *testing.T) {
	body := []byte(`{"touser":"zhangsan|lisi","toparty":"1|2","totag":"1|2","msgtype":"miniprogram_notice","miniprogram_notice":{"appid":"wx123123123123123","page":"pages/index?userid=zhangsan&orderid=123123123","title":"会议室预订成功通知","description":"4月27日 16:16","emphasis_first_item":true,"content_item":[{"key":"会议室","value":"402"},{"key":"会议地点","value":"广州TIT-402会议室"},{"key":"会议时间","value":"2018年8月1日 09:00-09:30"},{"key":"参与人员","value":"周剑轩"}]},"duplicate_check_interval":1800}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invaliduser": "userid1|userid2",
	"invalidparty": "partyid1|partyid2",
	"invalidtag": "tagid1|tagid2",
	"msgid": "xxxx",
	"response_code": "xyzxyz"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

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

	extra := &MsgExtra{
		ToUser:                 "zhangsan|lisi",
		ToParty:                "1|2",
		ToTag:                  "1|2",
		DuplicateCheckInterval: 1800,
	}

	result := new(ResultMsgSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendMinipNotice(notice, extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMsgSend{
		InvalidUser:  "userid1|userid2",
		InvalidParty: "partyid1|partyid2",
		InvalidTag:   "tagid1|tagid2",
		MsgID:        "xxxx",
		ResponseCode: "xyzxyz",
	}, result)
}

func TestSendTextNoticeCard(t *testing.T) {
	body := []byte(`{"touser":"UserID1|UserID2|UserID3","toparty":"PartyID1|PartyID2","totag":"TagID1|TagID2","msgtype":"template_card","agentid":1,"template_card":{"card_type":"text_notice","task_id":"task_id","source":{"icon_url":"图片的url","desc":"企业微信","desc_color":1},"action_menu":{"desc":"卡片副交互辅助文本说明","action_list":[{"text":"接受推送","key":"A"},{"text":"不再推送","key":"B"}]},"main_title":{"title":"欢迎使用企业微信","desc":"您的好友正在邀请您加入企业微信"},"quote_area":{"type":1,"url":"https://work.weixin.qq.com","title":"企业微信的引用样式","quote_text":"企业微信真好用呀真好用"},"emphasis_content":{"title":"100","desc":"核心数据"},"sub_title_text":"下载企业微信还能抢红包！","horizontal_content_list":[{"keyname":"邀请人","value":"张三"},{"type":1,"keyname":"企业微信官网","value":"点击访问","url":"https://work.weixin.qq.com"},{"type":2,"keyname":"企业微信下载","value":"企业微信.apk","media_id":"文件的media_id"},{"type":3,"keyname":"员工信息","value":"点击查看","userid":"zhangsan"}],"jump_list":[{"type":1,"title":"企业微信官网","url":"https://work.weixin.qq.com"},{"type":2,"title":"跳转小程序","appid":"小程序的appid","pagepath":"/index.html"}],"card_action":{"type":2,"url":"https://work.weixin.qq.com","appid":"小程序的appid","pagepath":"/index.html"}},"duplicate_check_interval":1800}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invaliduser": "userid1|userid2",
	"invalidparty": "partyid1|partyid2",
	"invalidtag": "tagid1|tagid2",
	"msgid": "xxxx",
	"response_code": "xyzxyz"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	card := &TextNoticeCard{
		Source: &CardSource{
			IconURL:   "图片的url",
			Desc:      "企业微信",
			DescColor: 1,
		},
		ActionMenu: &ActionMenu{
			Desc: "卡片副交互辅助文本说明",
			ActionList: []*MenuAction{
				{
					Text: "接受推送",
					Key:  "A",
				},
				{
					Text: "不再推送",
					Key:  "B",
				},
			},
		},
		MainTitle: &MainTitle{
			Title: "欢迎使用企业微信",
			Desc:  "您的好友正在邀请您加入企业微信",
		},
		QuoteArea: &QuoteArea{
			Type:      1,
			URL:       "https://work.weixin.qq.com",
			Title:     "企业微信的引用样式",
			QuoteText: "企业微信真好用呀真好用",
		},
		EmphasisContent: &EmphasisContent{
			Title: "100",
			Desc:  "核心数据",
		},
		SubTitleText: "下载企业微信还能抢红包！",
		HorizontalContentList: []*HorizontalContent{
			{
				Keyname: "邀请人",
				Value:   "张三",
			},
			{
				Type:    1,
				Keyname: "企业微信官网",
				Value:   "点击访问",
				URL:     "https://work.weixin.qq.com",
			},
			{
				Type:    2,
				Keyname: "企业微信下载",
				Value:   "企业微信.apk",
				MediaID: "文件的media_id",
			},
			{
				Type:    3,
				Keyname: "员工信息",
				Value:   "点击查看",
				UserID:  "zhangsan",
			},
		},
		JumpList: []*CardJump{
			{
				Type:  1,
				Title: "企业微信官网",
				URL:   "https://work.weixin.qq.com",
			},
			{
				Type:     2,
				Title:    "跳转小程序",
				AppID:    "小程序的appid",
				Pagepath: "/index.html",
			},
		},
		CardAction: &CardAction{
			Type:     2,
			URL:      "https://work.weixin.qq.com",
			AppID:    "小程序的appid",
			Pagepath: "/index.html",
		},
	}

	extra := &MsgExtra{
		ToUser:                 "UserID1|UserID2|UserID3",
		ToParty:                "PartyID1|PartyID2",
		ToTag:                  "TagID1|TagID2",
		DuplicateCheckInterval: 1800,
	}

	result := new(ResultMsgSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendTextNoticeCard(1, "task_id", card, extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMsgSend{
		InvalidUser:  "userid1|userid2",
		InvalidParty: "partyid1|partyid2",
		InvalidTag:   "tagid1|tagid2",
		MsgID:        "xxxx",
		ResponseCode: "xyzxyz",
	}, result)
}

func TestSendNewsNoticeCard(t *testing.T) {
	body := []byte(`{"touser":"UserID1|UserID2|UserID3","toparty":"PartyID1|PartyID2","totag":"TagID1|TagID2","msgtype":"template_card","agentid":1,"template_card":{"card_type":"news_notice","task_id":"task_id","source":{"icon_url":"图片的url","desc":"企业微信","desc_color":1},"action_menu":{"desc":"卡片副交互辅助文本说明","action_list":[{"text":"接受推送","key":"A"},{"text":"不再推送","key":"B"}]},"main_title":{"title":"欢迎使用企业微信","desc":"您的好友正在邀请您加入企业微信"},"quote_area":{"type":1,"url":"https://work.weixin.qq.com","title":"企业微信的引用样式","quote_text":"企业微信真好用呀真好用"},"image_text_area":{"type":1,"url":"https://work.weixin.qq.com","title":"企业微信的左图右文样式","desc":"企业微信真好用呀真好用","image_url":"https://img.iplaysoft.com/wp-content/uploads/2019/free-images/free_stock_photo_2x.jpg"},"card_image":{"url":"图片的url","aspect_ratio":1.3},"vertical_content_list":[{"title":"惊喜红包等你来拿","desc":"下载企业微信还能抢红包！"}],"horizontal_content_list":[{"keyname":"邀请人","value":"张三"},{"type":1,"keyname":"企业微信官网","value":"点击访问","url":"https://work.weixin.qq.com"},{"type":2,"keyname":"企业微信下载","value":"企业微信.apk","media_id":"文件的media_id"},{"type":3,"keyname":"员工信息","value":"点击查看","userid":"zhangsan"}],"jump_list":[{"type":1,"title":"企业微信官网","url":"https://work.weixin.qq.com"},{"type":2,"title":"跳转小程序","appid":"小程序的appid","pagepath":"/index.html"}],"card_action":{"type":2,"url":"https://work.weixin.qq.com","appid":"小程序的appid","pagepath":"/index.html"}},"duplicate_check_interval":1800}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invaliduser": "userid1|userid2",
	"invalidparty": "partyid1|partyid2",
	"invalidtag": "tagid1|tagid2",
	"msgid": "xxxx",
	"response_code": "xyzxyz"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	card := &NewsNoticeCard{
		Source: &CardSource{
			IconURL:   "图片的url",
			Desc:      "企业微信",
			DescColor: 1,
		},
		ActionMenu: &ActionMenu{
			Desc: "卡片副交互辅助文本说明",
			ActionList: []*MenuAction{
				{
					Text: "接受推送",
					Key:  "A",
				},
				{
					Text: "不再推送",
					Key:  "B",
				},
			},
		},
		MainTitle: &MainTitle{
			Title: "欢迎使用企业微信",
			Desc:  "您的好友正在邀请您加入企业微信",
		},
		QuoteArea: &QuoteArea{
			Type:      1,
			URL:       "https://work.weixin.qq.com",
			Title:     "企业微信的引用样式",
			QuoteText: "企业微信真好用呀真好用",
		},
		ImageTextArea: &ImageTextArea{
			Type:     1,
			URL:      "https://work.weixin.qq.com",
			Title:    "企业微信的左图右文样式",
			Desc:     "企业微信真好用呀真好用",
			ImageURL: "https://img.iplaysoft.com/wp-content/uploads/2019/free-images/free_stock_photo_2x.jpg",
		},
		CardImage: &CardImage{
			URL:         "图片的url",
			AspectRatio: 1.3,
		},
		VerticalContentList: []*VerticalContent{
			{
				Title: "惊喜红包等你来拿",
				Desc:  "下载企业微信还能抢红包！",
			},
		},
		HorizontalContentList: []*HorizontalContent{
			{
				Keyname: "邀请人",
				Value:   "张三",
			},
			{
				Type:    1,
				Keyname: "企业微信官网",
				Value:   "点击访问",
				URL:     "https://work.weixin.qq.com",
			},
			{
				Type:    2,
				Keyname: "企业微信下载",
				Value:   "企业微信.apk",
				MediaID: "文件的media_id",
			},
			{
				Type:    3,
				Keyname: "员工信息",
				Value:   "点击查看",
				UserID:  "zhangsan",
			},
		},
		JumpList: []*CardJump{
			{
				Type:  1,
				Title: "企业微信官网",
				URL:   "https://work.weixin.qq.com",
			},
			{
				Type:     2,
				Title:    "跳转小程序",
				AppID:    "小程序的appid",
				Pagepath: "/index.html",
			},
		},
		CardAction: &CardAction{
			Type:     2,
			URL:      "https://work.weixin.qq.com",
			AppID:    "小程序的appid",
			Pagepath: "/index.html",
		},
	}

	extra := &MsgExtra{
		ToUser:                 "UserID1|UserID2|UserID3",
		ToParty:                "PartyID1|PartyID2",
		ToTag:                  "TagID1|TagID2",
		DuplicateCheckInterval: 1800,
	}

	result := new(ResultMsgSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendNewsNoticeCard(1, "task_id", card, extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMsgSend{
		InvalidUser:  "userid1|userid2",
		InvalidParty: "partyid1|partyid2",
		InvalidTag:   "tagid1|tagid2",
		MsgID:        "xxxx",
		ResponseCode: "xyzxyz",
	}, result)
}

func TestSendButtonInteractionCard(t *testing.T) {
	body := []byte(`{"touser":"UserID1|UserID2|UserID3","toparty":"PartyID1|PartyID2","totag":"TagID1|TagID2","msgtype":"template_card","agentid":1,"template_card":{"card_type":"button_interaction","task_id":"task_id","source":{"icon_url":"图片的url","desc":"企业微信","desc_color":1},"action_menu":{"desc":"卡片副交互辅助文本说明","action_list":[{"text":"接受推送","key":"A"},{"text":"不再推送","key":"B"}]},"main_title":{"title":"欢迎使用企业微信","desc":"您的好友正在邀请您加入企业微信"},"quote_area":{"type":1,"url":"https://work.weixin.qq.com","title":"企业微信的引用样式","quote_text":"企业微信真好用呀真好用"},"sub_title_text":"下载企业微信还能抢红包！","horizontal_content_list":[{"keyname":"邀请人","value":"张三"},{"type":1,"keyname":"企业微信官网","value":"点击访问","url":"https://work.weixin.qq.com"},{"type":2,"keyname":"企业微信下载","value":"企业微信.apk","media_id":"文件的media_id"},{"type":3,"keyname":"员工信息","value":"点击查看","userid":"zhangsan"}],"card_action":{"type":2,"url":"https://work.weixin.qq.com","appid":"小程序的appid","pagepath":"/index.html"},"button_selection":{"question_key":"btn_question_key1","title":"企业微信评分","option_list":[{"id":"btn_selection_id1","text":"100分"},{"id":"btn_selection_id2","text":"101分"}],"selected_id":"btn_selection_id1"},"button_list":[{"text":"按钮1","style":1,"key":"button_key_1"},{"text":"按钮2","style":2,"key":"button_key_2"}]},"duplicate_check_interval":1800}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invaliduser": "userid1|userid2",
	"invalidparty": "partyid1|partyid2",
	"invalidtag": "tagid1|tagid2",
	"msgid": "xxxx",
	"response_code": "xyzxyz"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	card := &ButtonInteractionCard{
		Source: &CardSource{
			IconURL:   "图片的url",
			Desc:      "企业微信",
			DescColor: 1,
		},
		ActionMenu: &ActionMenu{
			Desc: "卡片副交互辅助文本说明",
			ActionList: []*MenuAction{
				{
					Text: "接受推送",
					Key:  "A",
				},
				{
					Text: "不再推送",
					Key:  "B",
				},
			},
		},
		MainTitle: &MainTitle{
			Title: "欢迎使用企业微信",
			Desc:  "您的好友正在邀请您加入企业微信",
		},
		QuoteArea: &QuoteArea{
			Type:      1,
			URL:       "https://work.weixin.qq.com",
			Title:     "企业微信的引用样式",
			QuoteText: "企业微信真好用呀真好用",
		},
		SubTitleText: "下载企业微信还能抢红包！",
		HorizontalContentList: []*HorizontalContent{
			{
				Keyname: "邀请人",
				Value:   "张三",
			},
			{
				Type:    1,
				Keyname: "企业微信官网",
				Value:   "点击访问",
				URL:     "https://work.weixin.qq.com",
			},
			{
				Type:    2,
				Keyname: "企业微信下载",
				Value:   "企业微信.apk",
				MediaID: "文件的media_id",
			},
			{
				Type:    3,
				Keyname: "员工信息",
				Value:   "点击查看",
				UserID:  "zhangsan",
			},
		},
		CardAction: &CardAction{
			Type:     2,
			URL:      "https://work.weixin.qq.com",
			AppID:    "小程序的appid",
			Pagepath: "/index.html",
		},
		ButtonSelection: &ButtonSelection{
			QuestionKey: "btn_question_key1",
			Title:       "企业微信评分",
			OptionList: []*SelectOption{
				{
					ID:   "btn_selection_id1",
					Text: "100分",
				},
				{
					ID:   "btn_selection_id2",
					Text: "101分",
				},
			},
			SelectedID: "btn_selection_id1",
		},
		ButtonList: []*CardButton{
			{
				Text:  "按钮1",
				Style: 1,
				Key:   "button_key_1",
			},
			{
				Text:  "按钮2",
				Style: 2,
				Key:   "button_key_2",
			},
		},
	}

	extra := &MsgExtra{
		ToUser:                 "UserID1|UserID2|UserID3",
		ToParty:                "PartyID1|PartyID2",
		ToTag:                  "TagID1|TagID2",
		DuplicateCheckInterval: 1800,
	}

	result := new(ResultMsgSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendButtonInteractionCard(1, "task_id", card, extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMsgSend{
		InvalidUser:  "userid1|userid2",
		InvalidParty: "partyid1|partyid2",
		InvalidTag:   "tagid1|tagid2",
		MsgID:        "xxxx",
		ResponseCode: "xyzxyz",
	}, result)
}

func TestSendVoteInteractionCard(t *testing.T) {
	body := []byte(`{"touser":"UserID1|UserID2|UserID3","toparty":"PartyID1|PartyID2","totag":"TagID1|TagID2","msgtype":"template_card","agentid":1,"template_card":{"card_type":"vote_interaction","task_id":"task_id","source":{"icon_url":"图片的url","desc":"企业微信"},"main_title":{"title":"欢迎使用企业微信","desc":"您的好友正在邀请您加入企业微信"},"checkbox":{"question_key":"question_key1","option_list":[{"id":"option_id1","text":"选择题选项1","is_checked":true},{"id":"option_id2","text":"选择题选项2","is_checked":false}],"mode":1},"submit_button":{"text":"提交","key":"key"}},"duplicate_check_interval":1800}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invaliduser": "userid1|userid2",
	"invalidparty": "partyid1|partyid2",
	"invalidtag": "tagid1|tagid2",
	"msgid": "xxxx",
	"response_code": "xyzxyz"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	card := &VoteInteractionCard{
		Source: &CardSource{
			IconURL: "图片的url",
			Desc:    "企业微信",
		},
		MainTitle: &MainTitle{
			Title: "欢迎使用企业微信",
			Desc:  "您的好友正在邀请您加入企业微信",
		},
		CheckBox: &CheckBox{
			QuestionKey: "question_key1",
			OptionList: []*CheckOption{
				{
					ID:        "option_id1",
					Text:      "选择题选项1",
					IsChecked: true,
				},
				{
					ID:        "option_id2",
					Text:      "选择题选项2",
					IsChecked: false,
				},
			},
			Mode: 1,
		},
		SubmitButton: &SubmitButton{
			Text: "提交",
			Key:  "key",
		},
	}

	extra := &MsgExtra{
		ToUser:                 "UserID1|UserID2|UserID3",
		ToParty:                "PartyID1|PartyID2",
		ToTag:                  "TagID1|TagID2",
		DuplicateCheckInterval: 1800,
	}

	result := new(ResultMsgSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendVoteInteractionCard(1, "task_id", card, extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMsgSend{
		InvalidUser:  "userid1|userid2",
		InvalidParty: "partyid1|partyid2",
		InvalidTag:   "tagid1|tagid2",
		MsgID:        "xxxx",
		ResponseCode: "xyzxyz",
	}, result)
}

func TestSendMultipleInteractionCard(t *testing.T) {
	body := []byte(`{"touser":"UserID1|UserID2|UserID3","toparty":"PartyID1|PartyID2","totag":"TagID1|TagID2","msgtype":"template_card","agentid":1,"template_card":{"card_type":"multiple_interaction","task_id":"task_id","source":{"icon_url":"图片的url","desc":"企业微信"},"main_title":{"title":"欢迎使用企业微信","desc":"您的好友正在邀请您加入企业微信"},"select_list":[{"question_key":"question_key1","title":"选择器标签1","option_list":[{"id":"selection_id1","text":"选择器选项1"},{"id":"selection_id2","text":"选择器选项2"}],"selected_id":"selection_id1"},{"question_key":"question_key2","title":"选择器标签2","option_list":[{"id":"selection_id3","text":"选择器选项3"},{"id":"selection_id4","text":"选择器选项4"}],"selected_id":"selection_id3"}],"submit_button":{"text":"提交","key":"key"}},"duplicate_check_interval":1800}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invaliduser": "userid1|userid2",
	"invalidparty": "partyid1|partyid2",
	"invalidtag": "tagid1|tagid2",
	"msgid": "xxxx",
	"response_code": "xyzxyz"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	card := &MultipleInteractionCard{
		Source: &CardSource{
			IconURL: "图片的url",
			Desc:    "企业微信",
		},
		MainTitle: &MainTitle{
			Title: "欢迎使用企业微信",
			Desc:  "您的好友正在邀请您加入企业微信",
		},
		SelectList: []*ButtonSelection{
			{
				QuestionKey: "question_key1",
				Title:       "选择器标签1",
				OptionList: []*SelectOption{
					{
						ID:   "selection_id1",
						Text: "选择器选项1",
					},
					{
						ID:   "selection_id2",
						Text: "选择器选项2",
					},
				},
				SelectedID: "selection_id1",
			},
			{
				QuestionKey: "question_key2",
				Title:       "选择器标签2",
				OptionList: []*SelectOption{
					{
						ID:   "selection_id3",
						Text: "选择器选项3",
					},
					{
						ID:   "selection_id4",
						Text: "选择器选项4",
					},
				},
				SelectedID: "selection_id3",
			},
		},
		SubmitButton: &SubmitButton{
			Text: "提交",
			Key:  "key",
		},
	}

	extra := &MsgExtra{
		ToUser:                 "UserID1|UserID2|UserID3",
		ToParty:                "PartyID1|PartyID2",
		ToTag:                  "TagID1|TagID2",
		DuplicateCheckInterval: 1800,
	}

	result := new(ResultMsgSend)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendMultipleInteractionCard(1, "task_id", card, extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMsgSend{
		InvalidUser:  "userid1|userid2",
		InvalidParty: "partyid1|partyid2",
		InvalidTag:   "tagid1|tagid2",
		MsgID:        "xxxx",
		ResponseCode: "xyzxyz",
	}, result)
}

func TestRecall(t *testing.T) {
	body := []byte(`{"msgid":"vcT8gGc-7dFb4bxT35ONjBDz901sLlXPZw1DAMC_Gc26qRpK-AK5sTJkkb0128t"}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/message/recall?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", Recall("vcT8gGc-7dFb4bxT35ONjBDz901sLlXPZw1DAMC_Gc26qRpK-AK5sTJkkb0128t"))

	assert.Nil(t, err)
}
