package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplyText(t *testing.T) {
	b, err := ReplyText("this is a test").Bytes("fromUser", "toUser")

	assert.Nil(t, err)

	expected := `<xml><FromUserName><![CDATA[fromUser]]></FromUserName><ToUserName><![CDATA[toUser]]></ToUserName><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[this is a test]]></Content></xml>`

	assert.Equal(t, expected, string(b))
}

func TestReplyImage(t *testing.T) {
	b, err := ReplyImage("media_id").Bytes("fromUser", "toUser")

	assert.Nil(t, err)

	expected := `<xml><FromUserName><![CDATA[fromUser]]></FromUserName><ToUserName><![CDATA[toUser]]></ToUserName><MsgType><![CDATA[image]]></MsgType><Image><MediaId><![CDATA[media_id]]></MediaId></Image></xml>`

	assert.Equal(t, expected, string(b))
}

func TestReplyVoice(t *testing.T) {
	b, err := ReplyVoice("media_id").Bytes("fromUser", "toUser")

	assert.Nil(t, err)

	expected := `<xml><FromUserName><![CDATA[fromUser]]></FromUserName><ToUserName><![CDATA[toUser]]></ToUserName><MsgType><![CDATA[voice]]></MsgType><Voice><MediaId><![CDATA[media_id]]></MediaId></Voice></xml>`

	assert.Equal(t, expected, string(b))
}

func TestReplyVideo(t *testing.T) {
	b, err := ReplyVideo("media_id", "title", "description").Bytes("fromUser", "toUser")

	assert.Nil(t, err)

	expected := `<xml><FromUserName><![CDATA[fromUser]]></FromUserName><ToUserName><![CDATA[toUser]]></ToUserName><MsgType><![CDATA[video]]></MsgType><Video><MediaId><![CDATA[media_id]]></MediaId><Title><![CDATA[title]]></Title><Description><![CDATA[description]]></Description></Video></xml>`

	assert.Equal(t, expected, string(b))
}

func TestReplyNews(t *testing.T) {
	articles := []*XMLNewsArticle{
		{
			Title:       "title1",
			Description: "description1",
			URL:         "url",
			PicURL:      "picurl",
		},
		{
			Title:       "title",
			Description: "description",
			URL:         "url",
			PicURL:      "picurl",
		},
	}

	b, err := ReplyNews(articles...).Bytes("fromUser", "toUser")

	assert.Nil(t, err)

	expected := `<xml><FromUserName><![CDATA[fromUser]]></FromUserName><ToUserName><![CDATA[toUser]]></ToUserName><MsgType><![CDATA[news]]></MsgType><ArticleCount>2</ArticleCount><Articles><item><Title><![CDATA[title1]]></Title><Description><![CDATA[description1]]></Description><Url><![CDATA[url]]></Url><PicUrl><![CDATA[picurl]]></PicUrl></item><item><Title><![CDATA[title]]></Title><Description><![CDATA[description]]></Description><Url><![CDATA[url]]></Url><PicUrl><![CDATA[picurl]]></PicUrl></item></Articles></xml>`

	assert.Equal(t, expected, string(b))
}

func TestReplyUpdateCardButton(t *testing.T) {
	b, err := ReplyUpdateCardButton("ReplaceName").Bytes("fromUser", "toUser")

	assert.Nil(t, err)

	expected := `<xml><FromUserName><![CDATA[fromUser]]></FromUserName><ToUserName><![CDATA[toUser]]></ToUserName><MsgType><![CDATA[update_button]]></MsgType><Button><ReplaceName><![CDATA[ReplaceName]]></ReplaceName></Button></xml>`

	assert.Equal(t, expected, string(b))
}

func TestReplyUpdateTextNoticeCard(t *testing.T) {
	card := &XMLTextNoticeCard{
		Source: &XMLCardSource{
			IconURL:   "source_url",
			Desc:      "更新后的卡片",
			DescColor: 2,
		},
		ActionMenu: &XMLActionMenu{
			Desc: "您可以使用以下功能",
			ActionList: []*XMLMenuAction{
				{
					Text: "您将收到A回调",
					Key:  "A",
				},
				{
					Text: "您将收到B回调",
					Key:  "B",
				},
			},
		},
		MainTitle: &XMLMainTitle{
			Title: "更新后的卡片标题",
			Desc:  "更新后的卡片副标题",
		},
		QuoteArea: &XMLQuoteArea{
			Type:      1,
			URL:       "quote_area_url",
			Title:     "企业微信",
			QuoteText: "企业微信真好用呀",
		},
		EmphasisContent: &XMLEmphasisContent{
			Title: "100万",
			Desc:  "核心数据实例",
		},
		SubTitleText: "更新后的卡片二级标题",
		HorizontalContentList: []*XMLHorizontalContent{
			{
				KeyName: "应用名称",
				Value:   "企业微信",
			},
			{
				Type:    1,
				KeyName: "跳转企业微信",
				Value:   "跳转企业微信",
				URL:     "url",
			},
		},
		JumpList: []*XMLCardJump{
			{
				Type:  1,
				Title: "跳转企业微信",
				URL:   "jump_url",
			},
		},
		CardAction: &XMLCardAction{
			Type: 1,
			URL:  "jump_url",
		},
	}

	b, err := ReplyUpdateTextNoticeCard(card).Bytes("FromUserName", "ToUserName")

	assert.Nil(t, err)

	expected := `<xml><FromUserName><![CDATA[FromUserName]]></FromUserName><ToUserName><![CDATA[ToUserName]]></ToUserName><MsgType><![CDATA[update_template_card]]></MsgType><TemplateCard><CardType><![CDATA[text_notice]]></CardType><Source><IconUrl><![CDATA[source_url]]></IconUrl><Desc><![CDATA[更新后的卡片]]></Desc><DescColor>2</DescColor></Source><ActionMenu><Desc><![CDATA[您可以使用以下功能]]></Desc><ActionList><Text><![CDATA[您将收到A回调]]></Text><Key><![CDATA[A]]></Key></ActionList><ActionList><Text><![CDATA[您将收到B回调]]></Text><Key><![CDATA[B]]></Key></ActionList></ActionMenu><MainTitle><Title><![CDATA[更新后的卡片标题]]></Title><Desc><![CDATA[更新后的卡片副标题]]></Desc></MainTitle><QuoteArea><Type>1</Type><Url><![CDATA[quote_area_url]]></Url><Title><![CDATA[企业微信]]></Title><QuoteText><![CDATA[企业微信真好用呀]]></QuoteText></QuoteArea><EmphasisContent><Title><![CDATA[100万]]></Title><Desc><![CDATA[核心数据实例]]></Desc></EmphasisContent><SubTitleText><![CDATA[更新后的卡片二级标题]]></SubTitleText><HorizontalContentList><KeyName><![CDATA[应用名称]]></KeyName><Value><![CDATA[企业微信]]></Value></HorizontalContentList><HorizontalContentList><Type>1</Type><KeyName><![CDATA[跳转企业微信]]></KeyName><Value><![CDATA[跳转企业微信]]></Value><Url><![CDATA[url]]></Url></HorizontalContentList><JumpList><Type>1</Type><Title><![CDATA[跳转企业微信]]></Title><Url><![CDATA[jump_url]]></Url></JumpList><CardAction><Type>1</Type><Url><![CDATA[jump_url]]></Url></CardAction></TemplateCard></xml>`

	assert.Equal(t, expected, string(b))
}

func TestReplyUpdateNewsNoticeCard(t *testing.T) {
	card := &XMLNewsNoticeCard{
		Source: &XMLCardSource{
			IconURL:   "source_url",
			Desc:      "更新后的卡片",
			DescColor: 2,
		},
		ActionMenu: &XMLActionMenu{
			Desc: "您可以使用以下功能",
			ActionList: []*XMLMenuAction{
				{
					Text: "您将收到A回调",
					Key:  "A",
				},
				{
					Text: "您将收到B回调",
					Key:  "B",
				},
			},
		},
		MainTitle: &XMLMainTitle{
			Title: "更新后的卡片标题",
			Desc:  "更新后的卡片副标题",
		},
		QuoteArea: &XMLQuoteArea{
			Type:      1,
			URL:       "quote_area_url",
			Title:     "企业微信",
			QuoteText: "企业微信真好用呀",
		},
		ImageTextArea: &XMLImageTextArea{
			Type:     1,
			URL:      "image_text_area_url",
			Title:    "企业微信",
			Desc:     "企业微信真好用呀",
			ImageURL: "image_url",
		},
		CardImage: &XMLCardImage{
			URL:         "image_url",
			AspectRatio: 1.3,
		},
		VerticalContentList: []*XMLVerticalContent{
			{
				Title: "卡片二级标题1",
				Desc:  "卡片二级内容1",
			},
			{
				Title: "卡片二级标题2",
				Desc:  "卡片二级内容2",
			},
		},
		HorizontalContentList: []*XMLHorizontalContent{
			{
				KeyName: "应用名称",
				Value:   "企业微信",
			},
			{
				Type:    1,
				KeyName: "跳转企业微信",
				Value:   "跳转企业微信",
				URL:     "url",
			},
		},
		JumpList: []*XMLCardJump{
			{
				Type:  1,
				Title: "跳转企业微信",
				URL:   "jump_url",
			},
		},
		CardAction: &XMLCardAction{
			Type: 1,
			URL:  "jump_url",
		},
	}

	b, err := ReplyUpdateNewsNoticeCard(card).Bytes("FromUserName", "ToUserName")

	assert.Nil(t, err)

	expected := `<xml><FromUserName><![CDATA[FromUserName]]></FromUserName><ToUserName><![CDATA[ToUserName]]></ToUserName><MsgType><![CDATA[update_template_card]]></MsgType><TemplateCard><CardType><![CDATA[news_notice]]></CardType><Source><IconUrl><![CDATA[source_url]]></IconUrl><Desc><![CDATA[更新后的卡片]]></Desc><DescColor>2</DescColor></Source><ActionMenu><Desc><![CDATA[您可以使用以下功能]]></Desc><ActionList><Text><![CDATA[您将收到A回调]]></Text><Key><![CDATA[A]]></Key></ActionList><ActionList><Text><![CDATA[您将收到B回调]]></Text><Key><![CDATA[B]]></Key></ActionList></ActionMenu><MainTitle><Title><![CDATA[更新后的卡片标题]]></Title><Desc><![CDATA[更新后的卡片副标题]]></Desc></MainTitle><QuoteArea><Type>1</Type><Url><![CDATA[quote_area_url]]></Url><Title><![CDATA[企业微信]]></Title><QuoteText><![CDATA[企业微信真好用呀]]></QuoteText></QuoteArea><ImageTextArea><Type>1</Type><Url><![CDATA[image_text_area_url]]></Url><Title><![CDATA[企业微信]]></Title><Desc><![CDATA[企业微信真好用呀]]></Desc><ImageUrl><![CDATA[image_url]]></ImageUrl></ImageTextArea><CardImage><Url><![CDATA[image_url]]></Url><AspectRatio>1.3</AspectRatio></CardImage><VerticalContentList><Title><![CDATA[卡片二级标题1]]></Title><Desc><![CDATA[卡片二级内容1]]></Desc></VerticalContentList><VerticalContentList><Title><![CDATA[卡片二级标题2]]></Title><Desc><![CDATA[卡片二级内容2]]></Desc></VerticalContentList><HorizontalContentList><KeyName><![CDATA[应用名称]]></KeyName><Value><![CDATA[企业微信]]></Value></HorizontalContentList><HorizontalContentList><Type>1</Type><KeyName><![CDATA[跳转企业微信]]></KeyName><Value><![CDATA[跳转企业微信]]></Value><Url><![CDATA[url]]></Url></HorizontalContentList><JumpList><Type>1</Type><Title><![CDATA[跳转企业微信]]></Title><Url><![CDATA[jump_url]]></Url></JumpList><CardAction><Type>1</Type><Url><![CDATA[jump_url]]></Url></CardAction></TemplateCard></xml>`

	assert.Equal(t, expected, string(b))
}

func TestReplyUpdateButtonInteractionCard(t *testing.T) {
	card := &XMLButtonInteractionCard{
		Source: &XMLCardSource{
			IconURL:   "source_url",
			Desc:      "更新后的卡片",
			DescColor: 2,
		},
		ActionMenu: &XMLActionMenu{
			Desc: "您可以使用以下功能",
			ActionList: []*XMLMenuAction{
				{
					Text: "您将收到A回调",
					Key:  "A",
				},
				{
					Text: "您将收到B回调",
					Key:  "B",
				},
			},
		},
		MainTitle: &XMLMainTitle{
			Title: "更新后的卡片标题",
			Desc:  "更新后的卡片副标题",
		},
		QuoteArea: &XMLQuoteArea{
			Type:      1,
			URL:       "quote_area_url",
			Title:     "企业微信",
			QuoteText: "企业微信真好用呀",
		},
		SubTitleText: "更新后的卡片二级标题",
		HorizontalContentList: []*XMLHorizontalContent{
			{
				KeyName: "应用名称",
				Value:   "企业微信",
			},
			{
				Type:    1,
				KeyName: "跳转企业微信",
				Value:   "跳转企业微信",
				URL:     "url",
			},
		},
		CardAction: &XMLCardAction{
			Type: 1,
			URL:  "jump_url",
		},
		ButtonSelection: &XMLButtonSelection{
			QuestionKey: "QuestionKey1",
			Title:       "下拉式选择器",
			OptionList: []*XMLSelectOption{
				{
					ID:   "option_id2",
					Text: "选择题选项2",
				},
				{
					ID:   "option_id2",
					Text: "选择题选项2",
				},
			},
			SelectedID: "option_id2",
		},
		ButtonList: []*XMLCardButton{
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

	b, err := ReplyUpdateButtonInteractionCard(card).Bytes("FromUserName", "ToUserName")

	assert.Nil(t, err)

	expected := `<xml><FromUserName><![CDATA[FromUserName]]></FromUserName><ToUserName><![CDATA[ToUserName]]></ToUserName><MsgType><![CDATA[update_template_card]]></MsgType><TemplateCard><CardType><![CDATA[button_interaction]]></CardType><Source><IconUrl><![CDATA[source_url]]></IconUrl><Desc><![CDATA[更新后的卡片]]></Desc><DescColor>2</DescColor></Source><ActionMenu><Desc><![CDATA[您可以使用以下功能]]></Desc><ActionList><Text><![CDATA[您将收到A回调]]></Text><Key><![CDATA[A]]></Key></ActionList><ActionList><Text><![CDATA[您将收到B回调]]></Text><Key><![CDATA[B]]></Key></ActionList></ActionMenu><MainTitle><Title><![CDATA[更新后的卡片标题]]></Title><Desc><![CDATA[更新后的卡片副标题]]></Desc></MainTitle><QuoteArea><Type>1</Type><Url><![CDATA[quote_area_url]]></Url><Title><![CDATA[企业微信]]></Title><QuoteText><![CDATA[企业微信真好用呀]]></QuoteText></QuoteArea><SubTitleText><![CDATA[更新后的卡片二级标题]]></SubTitleText><HorizontalContentList><KeyName><![CDATA[应用名称]]></KeyName><Value><![CDATA[企业微信]]></Value></HorizontalContentList><HorizontalContentList><Type>1</Type><KeyName><![CDATA[跳转企业微信]]></KeyName><Value><![CDATA[跳转企业微信]]></Value><Url><![CDATA[url]]></Url></HorizontalContentList><CardAction><Type>1</Type><Url><![CDATA[jump_url]]></Url></CardAction><ButtonSelection><QuestionKey><![CDATA[QuestionKey1]]></QuestionKey><Title><![CDATA[下拉式选择器]]></Title><OptionList><Id><![CDATA[option_id2]]></Id><Text><![CDATA[选择题选项2]]></Text></OptionList><OptionList><Id><![CDATA[option_id2]]></Id><Text><![CDATA[选择题选项2]]></Text></OptionList><SelectedId><![CDATA[option_id2]]></SelectedId><Disable>false</Disable></ButtonSelection><ButtonList><Text><![CDATA[按钮1]]></Text><Style>1</Style><Key><![CDATA[button_key_1]]></Key></ButtonList><ButtonList><Text><![CDATA[按钮2]]></Text><Style>2</Style><Key><![CDATA[button_key_2]]></Key></ButtonList><ReplaceText><![CDATA[已提交]]></ReplaceText></TemplateCard></xml>`

	assert.Equal(t, expected, string(b))
}

func TestReplyUpdateVoteInteractionCard(t *testing.T) {
	card := &XMLVoteInteractionCard{
		Source: &XMLCardSource{
			IconURL: "source_url",
			Desc:    "更新后的卡片",
		},
		MainTitle: &XMLMainTitle{
			Title: "更新后的卡片标题",
			Desc:  "更新后的卡片副标题",
		},
		CheckBox: &XMLCheckBox{
			QuestionKey: "QuestionKey1",
			OptionList: []*XMLCheckOption{
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
		SubmitButton: &XMLSubmitButton{
			Text: "提交",
			Key:  "Key",
		},
		ReplaceText: "已提交",
	}

	b, err := ReplyUpdateVoteInteractionCard(card).Bytes("FromUserName", "ToUserName")

	assert.Nil(t, err)

	expected := `<xml><FromUserName><![CDATA[FromUserName]]></FromUserName><ToUserName><![CDATA[ToUserName]]></ToUserName><MsgType><![CDATA[update_template_card]]></MsgType><TemplateCard><CardType><![CDATA[vote_interaction]]></CardType><Source><IconUrl><![CDATA[source_url]]></IconUrl><Desc><![CDATA[更新后的卡片]]></Desc></Source><MainTitle><Title><![CDATA[更新后的卡片标题]]></Title><Desc><![CDATA[更新后的卡片副标题]]></Desc></MainTitle><CheckBox><QuestionKey><![CDATA[QuestionKey1]]></QuestionKey><OptionList><Id><![CDATA[option_id1]]></Id><Text><![CDATA[选择题选项1]]></Text><IsChecked>true</IsChecked></OptionList><OptionList><Id><![CDATA[option_id2]]></Id><Text><![CDATA[选择题选项2]]></Text><IsChecked>false</IsChecked></OptionList><Disable>false</Disable><Mode>1</Mode></CheckBox><SubmitButton><Text><![CDATA[提交]]></Text><Key><![CDATA[Key]]></Key></SubmitButton><ReplaceText><![CDATA[已提交]]></ReplaceText></TemplateCard></xml>`

	assert.Equal(t, expected, string(b))
}

func TestReplyUpdateMultipleInteractionCard(t *testing.T) {
	card := &XMLMultipleInteractionCard{
		Source: &XMLCardSource{
			IconURL: "source_url",
			Desc:    "更新后的卡片",
		},
		MainTitle: &XMLMainTitle{
			Title: "更新后的卡片标题",
			Desc:  "更新后的卡片副标题",
		},
		SelectList: []*XMLButtonSelection{
			{
				QuestionKey: "QuestionKey1",
				Title:       "下拉式选择器",
				OptionList: []*XMLSelectOption{
					{
						ID:   "option_id2",
						Text: "选择题选项2",
					},
					{
						ID:   "option_id2",
						Text: "选择题选项2",
					},
				},
				SelectedID: "option_id2",
			},
		},
		SubmitButton: &XMLSubmitButton{
			Text: "提交",
			Key:  "Key",
		},
		ReplaceText: "已提交",
	}

	b, err := ReplyUpdateMultipleInteractionCard(card).Bytes("FromUserName", "ToUserName")

	assert.Nil(t, err)

	expected := `<xml><FromUserName><![CDATA[FromUserName]]></FromUserName><ToUserName><![CDATA[ToUserName]]></ToUserName><MsgType><![CDATA[update_template_card]]></MsgType><TemplateCard><CardType><![CDATA[multiple_interaction]]></CardType><Source><IconUrl><![CDATA[source_url]]></IconUrl><Desc><![CDATA[更新后的卡片]]></Desc></Source><MainTitle><Title><![CDATA[更新后的卡片标题]]></Title><Desc><![CDATA[更新后的卡片副标题]]></Desc></MainTitle><SelectList><QuestionKey><![CDATA[QuestionKey1]]></QuestionKey><Title><![CDATA[下拉式选择器]]></Title><OptionList><Id><![CDATA[option_id2]]></Id><Text><![CDATA[选择题选项2]]></Text></OptionList><OptionList><Id><![CDATA[option_id2]]></Id><Text><![CDATA[选择题选项2]]></Text></OptionList><SelectedId><![CDATA[option_id2]]></SelectedId><Disable>false</Disable></SelectList><SubmitButton><Text><![CDATA[提交]]></Text><Key><![CDATA[Key]]></Key></SubmitButton><ReplaceText><![CDATA[已提交]]></ReplaceText></TemplateCard></xml>`

	assert.Equal(t, expected, string(b))
}
