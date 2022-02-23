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

func TestUpdateButtonDisable(t *testing.T) {
	body := []byte(`{"userids":["userid1","userid2"],"partyids":[2,3],"tagids":[44,55],"agentid":1,"response_code":"response_code","button":{"replace_name":"replace_name"}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invaliduser": [
		"userid1",
		"userid2"
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/message/update_template_card?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	extra := &CardExtra{
		UserIDs:  []string{"userid1", "userid2"},
		PartyIDs: []int64{2, 3},
		TagIDs:   []int64{44, 55},
	}

	result := new(ResultCardUpdate)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UpdateCardButtonDisable(1, "response_code", "replace_name", extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultCardUpdate{
		InvalidUser: []string{"userid1", "userid2"},
	}, result)
}

func TestUpdateToTextNoticeCard(t *testing.T) {
	body := []byte(`{"userids":["userid1","userid2"],"partyids":[2,3],"agentid":1,"response_code":"response_code","template_card":{"card_type":"text_notice","source":{"icon_url":"图片的url","desc":"企业微信","desc_color":1},"action_menu":{"desc":"卡片副交互辅助文本说明","action_list":[{"text":"接受推送","key":"A"},{"text":"不再推送","key":"B"}]},"main_title":{"title":"欢迎使用企业微信","desc":"您的好友正在邀请您加入企业微信"},"quote_area":{"type":1,"url":"https://work.weixin.qq.com","title":"企业微信的引用样式","quote_text":"企业微信真好用呀真好用"},"emphasis_content":{"title":"100","desc":"核心数据"},"sub_title_text":"下载企业微信还能抢红包！","horizontal_content_list":[{"keyname":"邀请人","value":"张三"},{"type":1,"keyname":"企业微信官网","value":"点击访问","url":"https://work.weixin.qq.com"},{"type":2,"keyname":"企业微信下载","value":"企业微信.apk","media_id":"文件的media_id"},{"type":3,"keyname":"员工信息","value":"点击查看","userid":"zhangsan"}],"jump_list":[{"type":1,"title":"企业微信官网","url":"https://work.weixin.qq.com"},{"type":2,"title":"跳转小程序","appid":"小程序的appid","pagepath":"/index.html"}],"card_action":{"type":2,"url":"https://work.weixin.qq.com","appid":"小程序的appid","pagepath":"/index.html"}}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invaliduser": [
		"userid1",
		"userid2"
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/message/update_template_card?access_token=ACCESS_TOKEN", body).Return(resp, nil)

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
				KeyName: "邀请人",
				Value:   "张三",
			},
			{
				Type:    1,
				KeyName: "企业微信官网",
				Value:   "点击访问",
				URL:     "https://work.weixin.qq.com",
			},
			{
				Type:    2,
				KeyName: "企业微信下载",
				Value:   "企业微信.apk",
				MediaID: "文件的media_id",
			},
			{
				Type:    3,
				KeyName: "员工信息",
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
				PagePath: "/index.html",
			},
		},
		CardAction: &CardAction{
			Type:     2,
			URL:      "https://work.weixin.qq.com",
			AppID:    "小程序的appid",
			PagePath: "/index.html",
		},
	}

	extra := &CardExtra{
		UserIDs:  []string{"userid1", "userid2"},
		PartyIDs: []int64{2, 3},
	}

	result := new(ResultCardUpdate)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UpdateToTextNoticeCard(1, "response_code", card, extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultCardUpdate{
		InvalidUser: []string{"userid1", "userid2"},
	}, result)
}

func TestUpdateToNewsNoticeCard(t *testing.T) {
	body := []byte(`{"userids":["userid1","userid2"],"partyids":[2,3],"agentid":1,"response_code":"response_code","template_card":{"card_type":"news_notice","source":{"icon_url":"图片的url","desc":"企业微信","desc_color":1},"action_menu":{"desc":"卡片副交互辅助文本说明","action_list":[{"text":"接受推送","key":"A"},{"text":"不再推送","key":"B"}]},"main_title":{"title":"欢迎使用企业微信","desc":"您的好友正在邀请您加入企业微信"},"quote_area":{"type":1,"url":"https://work.weixin.qq.com","title":"企业微信的引用样式","quote_text":"企业微信真好用呀真好用"},"image_text_area":{"type":1,"url":"https://work.weixin.qq.com","title":"企业微信的左图右文样式","desc":"企业微信真好用呀真好用","image_url":"https://img.iplaysoft.com/wp-content/uploads/2019/free-images/free_stock_photo_2x.jpg"},"card_image":{"url":"图片的url","aspect_ratio":1.3},"vertical_content_list":[{"title":"惊喜红包等你来拿","desc":"下载企业微信还能抢红包！"}],"horizontal_content_list":[{"keyname":"邀请人","value":"张三"},{"type":1,"keyname":"企业微信官网","value":"点击访问","url":"https://work.weixin.qq.com"},{"type":2,"keyname":"企业微信下载","value":"企业微信.apk","media_id":"文件的media_id"},{"type":3,"keyname":"员工信息","value":"点击查看","userid":"zhangsan"}],"jump_list":[{"type":1,"title":"企业微信官网","url":"https://work.weixin.qq.com"},{"type":2,"title":"跳转小程序","appid":"小程序的appid","pagepath":"/index.html"}],"card_action":{"type":2,"url":"https://work.weixin.qq.com","appid":"小程序的appid","pagepath":"/index.html"}}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invaliduser": [
		"userid1",
		"userid2"
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/message/update_template_card?access_token=ACCESS_TOKEN", body).Return(resp, nil)

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
				KeyName: "邀请人",
				Value:   "张三",
			},
			{
				Type:    1,
				KeyName: "企业微信官网",
				Value:   "点击访问",
				URL:     "https://work.weixin.qq.com",
			},
			{
				Type:    2,
				KeyName: "企业微信下载",
				Value:   "企业微信.apk",
				MediaID: "文件的media_id",
			},
			{
				Type:    3,
				KeyName: "员工信息",
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
				PagePath: "/index.html",
			},
		},
		CardAction: &CardAction{
			Type:     2,
			URL:      "https://work.weixin.qq.com",
			AppID:    "小程序的appid",
			PagePath: "/index.html",
		},
	}

	extra := &CardExtra{
		UserIDs:  []string{"userid1", "userid2"},
		PartyIDs: []int64{2, 3},
	}

	result := new(ResultCardUpdate)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UpdateToNewsNoticeCard(1, "response_code", card, extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultCardUpdate{
		InvalidUser: []string{"userid1", "userid2"},
	}, result)
}

func TestUpdateToButtonInteractionCard(t *testing.T) {
	body := []byte(`{"userids":["userid1","userid2"],"partyids":[2,3],"agentid":1,"response_code":"response_code","template_card":{"card_type":"button_interaction","source":{"icon_url":"图片的url","desc":"企业微信","desc_color":1},"action_menu":{"desc":"卡片副交互辅助文本说明","action_list":[{"text":"接受推送","key":"A"},{"text":"不再推送","key":"B"}]},"main_title":{"title":"欢迎使用企业微信","desc":"您的好友正在邀请您加入企业微信"},"quote_area":{"type":1,"url":"https://work.weixin.qq.com","title":"企业微信的引用样式","quote_text":"企业微信真好用呀真好用"},"sub_title_text":"下载企业微信还能抢红包！","horizontal_content_list":[{"keyname":"邀请人","value":"张三"},{"type":1,"keyname":"企业微信官网","value":"点击访问","url":"https://work.weixin.qq.com"},{"type":2,"keyname":"企业微信下载","value":"企业微信.apk","media_id":"文件的media_id"},{"type":3,"keyname":"员工信息","value":"点击查看","userid":"zhangsan"}],"card_action":{"type":2,"url":"https://work.weixin.qq.com","appid":"小程序的appid","pagepath":"/index.html"},"button_selection":{"question_key":"btn_question_key1","title":"企业微信评分","option_list":[{"id":"btn_selection_id1","text":"100分"},{"id":"btn_selection_id2","text":"101分"}],"selected_id":"btn_selection_id1"},"button_list":[{"text":"按钮1","style":1,"key":"button_key_1"},{"text":"按钮2","style":2,"key":"button_key_2"}],"replace_text":"已提交"}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invaliduser": [
		"userid1",
		"userid2"
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/message/update_template_card?access_token=ACCESS_TOKEN", body).Return(resp, nil)

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
				KeyName: "邀请人",
				Value:   "张三",
			},
			{
				Type:    1,
				KeyName: "企业微信官网",
				Value:   "点击访问",
				URL:     "https://work.weixin.qq.com",
			},
			{
				Type:    2,
				KeyName: "企业微信下载",
				Value:   "企业微信.apk",
				MediaID: "文件的media_id",
			},
			{
				Type:    3,
				KeyName: "员工信息",
				Value:   "点击查看",
				UserID:  "zhangsan",
			},
		},
		CardAction: &CardAction{
			Type:     2,
			URL:      "https://work.weixin.qq.com",
			AppID:    "小程序的appid",
			PagePath: "/index.html",
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
		ReplaceText: "已提交",
	}

	extra := &CardExtra{
		UserIDs:  []string{"userid1", "userid2"},
		PartyIDs: []int64{2, 3},
	}

	result := new(ResultCardUpdate)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UpdateToButtonInteractionCard(1, "response_code", card, extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultCardUpdate{
		InvalidUser: []string{"userid1", "userid2"},
	}, result)
}

func TestUpdateToVoteInteractionCard(t *testing.T) {
	body := []byte(`{"userids":["userid1","userid2"],"partyids":[2,3],"agentid":1,"response_code":"response_code","template_card":{"card_type":"vote_interaction","source":{"icon_url":"图片的url","desc":"企业微信"},"main_title":{"title":"欢迎使用企业微信","desc":"您的好友正在邀请您加入企业微信"},"checkbox":{"question_key":"question_key1","option_list":[{"id":"option_id1","text":"选择题选项1","is_checked":true},{"id":"option_id2","text":"选择题选项2","is_checked":false}],"mode":1},"submit_button":{"text":"提交","key":"key"},"replace_text":"已提交"}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invaliduser": [
		"userid1",
		"userid2"
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/message/update_template_card?access_token=ACCESS_TOKEN", body).Return(resp, nil)

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
		ReplaceText: "已提交",
	}

	extra := &CardExtra{
		UserIDs:  []string{"userid1", "userid2"},
		PartyIDs: []int64{2, 3},
	}

	result := new(ResultCardUpdate)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UpdateToVoteInteractionCard(1, "response_code", card, extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultCardUpdate{
		InvalidUser: []string{"userid1", "userid2"},
	}, result)
}

func TestUpdateToMultipleInteractionCard(t *testing.T) {
	body := []byte(`{"userids":["userid1","userid2"],"partyids":[2,3],"tagids":[44,55],"agentid":1,"response_code":"response_code","template_card":{"card_type":"multiple_interaction","source":{"icon_url":"图片的url","desc":"企业微信"},"main_title":{"title":"欢迎使用企业微信","desc":"您的好友正在邀请您加入企业微信"},"select_list":[{"question_key":"question_key1","title":"选择器标签1","option_list":[{"id":"selection_id1","text":"选择器选项1"},{"id":"selection_id2","text":"选择器选项2"}],"selected_id":"selection_id1"},{"question_key":"question_key2","title":"选择器标签2","option_list":[{"id":"selection_id3","text":"选择器选项3"},{"id":"selection_id4","text":"选择器选项4"}],"selected_id":"selection_id3"}],"submit_button":{"text":"提交","key":"key"},"replace_text":"已提交"}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"invaliduser": [
		"userid1",
		"userid2"
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/message/update_template_card?access_token=ACCESS_TOKEN", body).Return(resp, nil)

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
		ReplaceText: "已提交",
	}

	extra := &CardExtra{
		UserIDs:  []string{"userid1", "userid2"},
		PartyIDs: []int64{2, 3},
		TagIDs:   []int64{44, 55},
	}

	result := new(ResultCardUpdate)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", UpdateToMultipleInteractionCard(1, "response_code", card, extra, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultCardUpdate{
		InvalidUser: []string{"userid1", "userid2"},
	}, result)
}
