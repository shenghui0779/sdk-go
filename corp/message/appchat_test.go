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

func TestCreateAppchat(t *testing.T) {
	body := []byte(`{"name":"NAME","owner":"userid1","userlist":["userid1","userid2","userid3"],"chatid":"CHATID"}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"chatid": "CHATID"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/appchat/create?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsAppchatCreate{
		Name:     "NAME",
		Owner:    "userid1",
		UserList: []string{"userid1", "userid2", "userid3"},
		ChatID:   "CHATID",
	}
	result := new(ResultAppchartCreate)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", CreateAppchat(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultAppchartCreate{
		ChatID: "CHATID",
	}, result)
}

func TestUpdateAppchat(t *testing.T) {
	body := []byte(`{"chatid":"CHATID","name":"NAME","owner":"userid2","add_user_list":["userid1","userid2","userid3"],"del_user_list":["userid3","userid4"]}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/appchat/update?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsAppchatUpdate{
		ChatID:      "CHATID",
		Name:        "NAME",
		Owner:       "userid2",
		AddUserList: []string{"userid1", "userid2", "userid3"},
		DelUserList: []string{"userid3", "userid4"},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UpdateAppchat(params))

	assert.Nil(t, err)
}

func TestGetAppchat(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"chat_info": {
		"chatid": "CHATID",
		"name": "NAME",
		"owner": "userid2",
		"userlist": [
			"userid1",
			"userid2",
			"userid3"
		]
	}
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/appchat/get?access_token=ACCESS_TOKEN&chatid=CHATID", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultAppchatGet)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetAppchat("CHATID", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultAppchatGet{
		ChatInfo: &AppchatInfo{
			ChatID:   "CHATID",
			Name:     "NAME",
			Owner:    "userid2",
			UserList: []string{"userid1", "userid2", "userid3"},
		},
	}, result)
}

func TestSendAppchatText(t *testing.T) {
	body := []byte(`{"chatid":"CHATID","msgtype":"text","text":{"content":"你的快递已到\n请携带工卡前往邮件中心领取"}}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/appchat/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendAppchatText("CHATID", "你的快递已到\n请携带工卡前往邮件中心领取", 0))

	assert.Nil(t, err)
}

func TestSendAppchatImage(t *testing.T) {
	body := []byte(`{"chatid":"CHATID","msgtype":"image","image":{"media_id":"MEDIAID"}}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/appchat/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendAppchatImage("CHATID", "MEDIAID", 0))

	assert.Nil(t, err)
}

func TestSendAppchatVoice(t *testing.T) {
	body := []byte(`{"chatid":"CHATID","msgtype":"voice","voice":{"media_id":"MEDIA_ID"}}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/appchat/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendAppchatVoice("CHATID", "MEDIA_ID", 0))

	assert.Nil(t, err)
}

func TestSendAppchatVideo(t *testing.T) {
	body := []byte(`{"chatid":"CHATID","msgtype":"video","video":{"media_id":"MEDIA_ID","title":"Title","description":"Description"}}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/appchat/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	video := &Video{
		MediaID:     "MEDIA_ID",
		Title:       "Title",
		Description: "Description",
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendAppchatVideo("CHATID", video, 0))

	assert.Nil(t, err)
}

func TestSendAppchatFile(t *testing.T) {
	body := []byte(`{"chatid":"CHATID","msgtype":"file","file":{"media_id":"1Yv-zXfHjSjU-7LH-GwtYqDGS-zz6w22KmWAT5COgP7o"}}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/appchat/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendAppchatFile("CHATID", "1Yv-zXfHjSjU-7LH-GwtYqDGS-zz6w22KmWAT5COgP7o", 0))

	assert.Nil(t, err)
}

func TestSendAppchatTextCard(t *testing.T) {
	body := []byte(`{"chatid":"CHATID","msgtype":"textcard","textcard":{"title":"领奖通知","description":"<div class=\"gray\">2016年9月26日</div> <div class=\"normal\"> 恭喜你抽中iPhone 7一台，领奖码:520258</div><div class=\"highlight\">请于2016年10月10日前联系行 政同事领取</div>","url":"https://work.weixin.qq.com/","btntxt":"更多"}}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/appchat/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	card := &TextCard{
		Title:       "领奖通知",
		Description: "<div class=\"gray\">2016年9月26日</div> <div class=\"normal\"> 恭喜你抽中iPhone 7一台，领奖码:520258</div><div class=\"highlight\">请于2016年10月10日前联系行 政同事领取</div>",
		URL:         "https://work.weixin.qq.com/",
		BtnTxt:      "更多",
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendAppchatTextCard("CHATID", card, 0))

	assert.Nil(t, err)
}

func TestSendAppchatNews(t *testing.T) {
	body := []byte(`{"chatid":"CHATID","msgtype":"news","news":{"articles":[{"title":"中秋节礼品领取","description":"今年中秋节公司有豪礼相送","url":"https://work.weixin.qq.com/","picurl":"http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png"}]}}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/appchat/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	articles := []*NewsArticle{
		{
			Title:       "中秋节礼品领取",
			Description: "今年中秋节公司有豪礼相送",
			URL:         "https://work.weixin.qq.com/",
			PicURL:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
		},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendAppchatNews("CHATID", articles, 0))

	assert.Nil(t, err)
}

func TestSendAppchatMPNews(t *testing.T) {
	body := []byte(`{"chatid":"CHATID","msgtype":"mpnews","mpnews":{"articles":[{"title":"地球一小时","thumb_media_id":"biz_get(image)","author":"Author","content_source_url":"https://work.weixin.qq.com","content":"3月24日20:30-21:30 \n办公区将关闭照明一小时，请各部门同事相互转告","digest":"3月24日20:30-21:30 \n办公区将关闭照明一小时"}]}}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/appchat/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	articles := []*MPNewsArticle{
		{
			Title:            "地球一小时",
			ThumbMediaID:     "biz_get(image)",
			Author:           "Author",
			ContentSourceURL: "https://work.weixin.qq.com",
			Content:          "3月24日20:30-21:30 \n办公区将关闭照明一小时，请各部门同事相互转告",
			Digest:           "3月24日20:30-21:30 \n办公区将关闭照明一小时",
		},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendAppchatMPNews("CHATID", articles, 0))

	assert.Nil(t, err)
}

func TestSendAppchatMarkdown(t *testing.T) {
	body := []byte(`{"chatid":"CHATID","msgtype":"markdown","markdown":{"content":"您的会议室已经预定，稍后会同步到邮箱 >**事项详情** \n>事　项：<font color=\"info\">开会</font> \n>组织者：@miglioguan \n>参与者：@miglioguan、@kunliu、@jamdeezhou、@kanexiong、@kisonwang\n>\n>会议室：<font color=\"info\">广州TIT 1楼 301</font>\n>日　期：<font color=\"warning\">2018年5月18日</font>\n>时　间：<font color=\"comment\">上午9:00-11:00</font>\n>\n>请准时参加会议。\n>\n>如需修改会议信息，请点击：[修改会议信息](https://work.weixin.qq.com)"}}`)
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

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/appchat/send?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", SendAppchatMarkdown("CHATID", "您的会议室已经预定，稍后会同步到邮箱 >**事项详情** \n>事　项：<font color=\"info\">开会</font> \n>组织者：@miglioguan \n>参与者：@miglioguan、@kunliu、@jamdeezhou、@kanexiong、@kisonwang\n>\n>会议室：<font color=\"info\">广州TIT 1楼 301</font>\n>日　期：<font color=\"warning\">2018年5月18日</font>\n>时　间：<font color=\"comment\">上午9:00-11:00</font>\n>\n>请准时参加会议。\n>\n>如需修改会议信息，请点击：[修改会议信息](https://work.weixin.qq.com)", 0))

	assert.Nil(t, err)
}
