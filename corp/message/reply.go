package message

import (
	"encoding/xml"
	"time"

	"github.com/shenghui0779/gochat/event"
	"github.com/shenghui0779/gochat/wx"
)

type XMLText struct {
	Content wx.CDATA `xml:"Content,omitempty"`
}

type XMLMedia struct {
	MediaID wx.CDATA `xml:"MediaId,omitempty"`
}

type XMLVideo struct {
	MediaID     wx.CDATA `xml:"MediaId,omitempty"`
	Title       wx.CDATA `xml:"Title,omitempty"`
	Description wx.CDATA `xml:"Description,omitempty"`
}

type XMLNews struct {
	Articles []*XMLNewsArticle `xml:"item,omitempty"`
}

type XMLNewsArticle struct {
	Title       wx.CDATA `xml:"Title,omitempty"`
	Description wx.CDATA `xml:"Description,omitempty"`
	URL         wx.CDATA `xml:"Url,omitempty"`
	PicURL      wx.CDATA `xml:"PicUrl,omitempty"`
}

// Reply 消息回复
type Reply struct {
	XMLName      xml.Name         `xml:"xml"`
	FromUserName wx.CDATA         `xml:"FromUserName,omitempty"`
	ToUserName   wx.CDATA         `xml:"ToUserName,omitempty"`
	CreateTime   int64            `xml:"CreateTime,omitempty"`
	MsgType      wx.CDATA         `xml:"MsgType,omitempty"`
	Content      wx.CDATA         `xml:"Content,omitempty"`
	Image        *XMLMedia        `xml:"Image,omitempty"`
	Voice        *XMLMedia        `xml:"Voice,omitempty"`
	Video        *XMLVideo        `xml:"Video,omitempty"`
	ArticleCount int              `xml:"ArticleCount,omitempty"`
	Articles     *XMLNews         `xml:"Articles,omitempty"`
	Button       *XMLUpdateButton `xml:"Button,omitempty"`
	TemplateCard *XMLTemplateCard `xml:"TemplateCard,omitempty"`
}

func (r *Reply) Bytes(from, to string) ([]byte, error) {
	r.FromUserName = wx.CDATA(from)
	r.ToUserName = wx.CDATA(to)
	r.CreateTime = time.Now().Unix() // 执行 testing 前，请注释掉

	return xml.Marshal(r)
}

func ReplyText(content string) event.Reply {
	return &Reply{
		MsgType: wx.CDATA(event.MsgText),
		Content: wx.CDATA(content),
	}
}

func ReplyImage(mediaID string) event.Reply {
	return &Reply{
		MsgType: wx.CDATA(event.MsgImage),
		Image: &XMLMedia{
			MediaID: wx.CDATA(mediaID),
		},
	}
}

func ReplyVoice(mediaID string) event.Reply {
	return &Reply{
		MsgType: wx.CDATA(event.MsgVoice),
		Voice: &XMLMedia{
			MediaID: wx.CDATA(mediaID),
		},
	}
}

func ReplyVideo(mediaID, title, description string) event.Reply {
	return &Reply{
		MsgType: wx.CDATA(event.MsgVideo),
		Video: &XMLVideo{
			MediaID:     wx.CDATA(mediaID),
			Title:       wx.CDATA(title),
			Description: wx.CDATA(description),
		},
	}
}

func ReplyNews(articles ...*XMLNewsArticle) event.Reply {
	return &Reply{
		MsgType:      wx.CDATA(event.MsgNews),
		ArticleCount: len(articles),
		Articles: &XMLNews{
			Articles: articles,
		},
	}
}

type XMLUpdateButton struct {
	ReplaceName wx.CDATA `xml:"ReplaceName,omitempty"`
}

type XMLTemplateCard struct {
	CardType              wx.CDATA                `xml:"CardType,omitempty"`
	Source                *XMLCardSource          `xml:"Source,omitempty"`
	ActionMenu            *XMLActionMenu          `xml:"ActionMenu,omitempty"`
	MainTitle             *XMLMainTitle           `xml:"MainTitle,omitempty"`
	QuoteArea             *XMLQuoteArea           `xml:"QuoteArea,omitempty"`
	EmphasisContent       *XMLEmphasisContent     `xml:"EmphasisContent,omitempty"`
	SubTitleText          wx.CDATA                `xml:"SubTitleText,omitempty"`
	ImageTextArea         *XMLImageTextArea       `xml:"ImageTextArea,omitempty"`
	CardImage             *XMLCardImage           `xml:"CardImage,omitempty"`
	VerticalContentList   []*XMLVerticalContent   `xml:"VerticalContentList,omitempty"`
	HorizontalContentList []*XMLHorizontalContent `xml:"HorizontalContentList,omitempty"`
	JumpList              []*XMLCardJump          `xml:"JumpList,omitempty"`
	CardAction            *XMLCardAction          `xml:"CardAction,omitempty"`
	ButtonSelection       *XMLButtonSelection     `xml:"ButtonSelection,omitempty"`
	ButtonList            []*XMLCardButton        `xml:"ButtonList,omitempty"`
	CheckBox              *XMLCheckBox            `xml:"CheckBox,omitempty"`
	SelectList            []*XMLButtonSelection   `xml:"SelectList,omitempty"`
	SubmitButton          *XMLSubmitButton        `xml:"SubmitButton,omitempty"`
	ReplaceText           wx.CDATA                `xml:"ReplaceText,omitempty"`
}

type XMLCardSource struct {
	IconURL   wx.CDATA `xml:"IconUrl,omitempty"`
	Desc      wx.CDATA `xml:"Desc,omitempty"`
	DescColor int      `xml:"DescColor,omitempty"`
}

type XMLActionMenu struct {
	Desc       wx.CDATA         `xml:"Desc,omitempty"`
	ActionList []*XMLMenuAction `xml:"ActionList,omitempty"`
}

type XMLMenuAction struct {
	Text wx.CDATA `xml:"Text,omitempty"`
	Key  wx.CDATA `xml:"Key,omitempty"`
}

type XMLMainTitle struct {
	Title wx.CDATA `xml:"Title,omitempty"`
	Desc  wx.CDATA `xml:"Desc,omitempty"`
}

type XMLQuoteArea struct {
	Type      int      `xml:"Type,omitempty"`
	URL       wx.CDATA `xml:"Url,omitempty"`
	AppID     wx.CDATA `xml:"AppId,omitempty"`
	PagePath  wx.CDATA `xml:"PagePath,omitempty"`
	Title     wx.CDATA `xml:"Title,omitempty"`
	QuoteText wx.CDATA `xml:"QuoteText,omitempty"`
}

type XMLImageTextArea struct {
	Type     int      `xml:"Type,omitempty"`
	URL      wx.CDATA `xml:"Url,omitempty"`
	AppID    wx.CDATA `xml:"AppId,omitempty"`
	PagePath wx.CDATA `xml:"PagePath,omitempty"`
	Title    wx.CDATA `xml:"Title,omitempty"`
	Desc     wx.CDATA `xml:"Desc,omitempty"`
	ImageURL wx.CDATA `xml:"ImageUrl,omitempty"`
}

type XMLCardImage struct {
	URL         wx.CDATA `xml:"Url,omitempty"`
	AspectRatio float64  `xml:"AspectRatio,omitempty"`
}

type XMLVerticalContent struct {
	Title wx.CDATA `xml:"Title,omitempty"`
	Desc  wx.CDATA `xml:"Desc,omitempty"`
}

type XMLHorizontalContent struct {
	Type    int      `xml:"Type,omitempty"`
	KeyName wx.CDATA `xml:"KeyName,omitempty"`
	Value   wx.CDATA `xml:"Value,omitempty"`
	UserID  wx.CDATA `xml:"UserId,omitempty"`
	MediaID wx.CDATA `xml:"MediaId,omitempty"`
	URL     wx.CDATA `xml:"Url,omitempty"`
}

type XMLCardJump struct {
	Type     int      `xml:"Type,omitempty"`
	Title    wx.CDATA `xml:"Title,omitempty"`
	URL      wx.CDATA `xml:"Url,omitempty"`
	AppID    wx.CDATA `xml:"AppId,omitempty"`
	PagePath wx.CDATA `xml:"PagePath,omitempty"`
}

type XMLCardAction struct {
	Type     int      `xml:"Type,omitempty"`
	URL      wx.CDATA `xml:"Url,omitempty"`
	AppID    wx.CDATA `xml:"AppId,omitempty"`
	PagePath wx.CDATA `xml:"PagePath,omitempty"`
}

type XMLEmphasisContent struct {
	Title wx.CDATA `xml:"Title,omitempty"`
	Desc  wx.CDATA `xml:"Desc,omitempty"`
}

type XMLCheckBox struct {
	QuestionKey wx.CDATA          `xml:"QuestionKey,omitempty"`
	OptionList  []*XMLCheckOption `xml:"OptionList,omitempty"`
	Disable     bool              `xml:"Disable"`
	Mode        int               `xml:"Mode,omitempty"`
}

type XMLCheckOption struct {
	ID        wx.CDATA `xml:"Id,omitempty"`
	Text      wx.CDATA `xml:"Text,omitempty"`
	IsChecked bool     `xml:"IsChecked"`
}

type XMLSubmitButton struct {
	Text wx.CDATA `xml:"Text,omitempty"`
	Key  wx.CDATA `xml:"Key,omitempty"`
}

type XMLButtonSelection struct {
	QuestionKey wx.CDATA           `xml:"QuestionKey,omitempty"`
	Title       wx.CDATA           `xml:"Title,omitempty"`
	OptionList  []*XMLSelectOption `xml:"OptionList,omitempty"`
	SelectedID  wx.CDATA           `xml:"SelectedId,omitempty"`
	Disable     bool               `xml:"Disable"`
}

type XMLSelectOption struct {
	ID   wx.CDATA `xml:"Id,omitempty"`
	Text wx.CDATA `xml:"Text,omitempty"`
}

type XMLCardButton struct {
	Text  wx.CDATA `xml:"Text,omitempty"`
	Style int      `xml:"Style,omitempty"`
	Key   wx.CDATA `xml:"Key,omitempty"`
}

type XMLTextNoticeCard struct {
	Source                *XMLCardSource
	ActionMenu            *XMLActionMenu
	MainTitle             *XMLMainTitle
	QuoteArea             *XMLQuoteArea
	EmphasisContent       *XMLEmphasisContent
	SubTitleText          wx.CDATA
	HorizontalContentList []*XMLHorizontalContent
	JumpList              []*XMLCardJump
	CardAction            *XMLCardAction
}

type XMLNewsNoticeCard struct {
	Source                *XMLCardSource
	ActionMenu            *XMLActionMenu
	MainTitle             *XMLMainTitle
	QuoteArea             *XMLQuoteArea
	ImageTextArea         *XMLImageTextArea
	CardImage             *XMLCardImage
	VerticalContentList   []*XMLVerticalContent
	HorizontalContentList []*XMLHorizontalContent
	JumpList              []*XMLCardJump
	CardAction            *XMLCardAction
}

type XMLButtonInteractionCard struct {
	Source                *XMLCardSource
	ActionMenu            *XMLActionMenu
	MainTitle             *XMLMainTitle
	QuoteArea             *XMLQuoteArea
	SubTitleText          wx.CDATA
	HorizontalContentList []*XMLHorizontalContent
	CardAction            *XMLCardAction
	ButtonSelection       *XMLButtonSelection
	ButtonList            []*XMLCardButton
	ReplaceText           wx.CDATA
}

type XMLVoteInteractionCard struct {
	Source       *XMLCardSource
	MainTitle    *XMLMainTitle
	CheckBox     *XMLCheckBox
	SubmitButton *XMLSubmitButton
	ReplaceText  wx.CDATA
}

type XMLMultipleInteractionCard struct {
	Source       *XMLCardSource
	MainTitle    *XMLMainTitle
	SelectList   []*XMLButtonSelection
	SubmitButton *XMLSubmitButton
	ReplaceText  wx.CDATA
}

func ReplyUpdateCardButton(replaceName string) event.Reply {
	return &Reply{
		MsgType: wx.CDATA(event.MsgUpdateButton),
		Button: &XMLUpdateButton{
			ReplaceName: wx.CDATA(replaceName),
		},
	}
}

func ReplyUpdateTextNoticeCard(card *XMLTextNoticeCard) event.Reply {
	return &Reply{
		MsgType: wx.CDATA(event.MsgUpdateTemplateCard),
		TemplateCard: &XMLTemplateCard{
			CardType:              wx.CDATA(CardTextNotice),
			Source:                card.Source,
			ActionMenu:            card.ActionMenu,
			MainTitle:             card.MainTitle,
			QuoteArea:             card.QuoteArea,
			EmphasisContent:       card.EmphasisContent,
			SubTitleText:          card.SubTitleText,
			HorizontalContentList: card.HorizontalContentList,
			JumpList:              card.JumpList,
			CardAction:            card.CardAction,
		},
	}
}

func ReplyUpdateNewsNoticeCard(card *XMLNewsNoticeCard) event.Reply {
	return &Reply{
		MsgType: wx.CDATA(event.MsgUpdateTemplateCard),
		TemplateCard: &XMLTemplateCard{
			CardType:              wx.CDATA(CardNewsNotice),
			Source:                card.Source,
			ActionMenu:            card.ActionMenu,
			MainTitle:             card.MainTitle,
			QuoteArea:             card.QuoteArea,
			ImageTextArea:         card.ImageTextArea,
			CardImage:             card.CardImage,
			VerticalContentList:   card.VerticalContentList,
			HorizontalContentList: card.HorizontalContentList,
			JumpList:              card.JumpList,
			CardAction:            card.CardAction,
		},
	}
}

func ReplyUpdateButtonInteractionCard(card *XMLButtonInteractionCard) event.Reply {
	return &Reply{
		MsgType: wx.CDATA(event.MsgUpdateTemplateCard),
		TemplateCard: &XMLTemplateCard{
			CardType:              wx.CDATA(CardButtonInteraction),
			Source:                card.Source,
			ActionMenu:            card.ActionMenu,
			MainTitle:             card.MainTitle,
			QuoteArea:             card.QuoteArea,
			SubTitleText:          wx.CDATA(card.SubTitleText),
			HorizontalContentList: card.HorizontalContentList,
			CardAction:            card.CardAction,
			ButtonSelection:       card.ButtonSelection,
			ButtonList:            card.ButtonList,
			ReplaceText:           card.ReplaceText,
		},
	}
}

func ReplyUpdateVoteInteractionCard(card *XMLVoteInteractionCard) event.Reply {
	return &Reply{
		MsgType: wx.CDATA(event.MsgUpdateTemplateCard),
		TemplateCard: &XMLTemplateCard{
			CardType:     wx.CDATA(CardVoteInteraction),
			Source:       card.Source,
			MainTitle:    card.MainTitle,
			CheckBox:     card.CheckBox,
			SubmitButton: card.SubmitButton,
			ReplaceText:  wx.CDATA(card.ReplaceText),
		},
	}
}

func ReplyUpdateMultipleInteractionCard(card *XMLMultipleInteractionCard) event.Reply {
	return &Reply{
		MsgType: wx.CDATA(event.MsgUpdateTemplateCard),
		TemplateCard: &XMLTemplateCard{
			CardType:     wx.CDATA(CardMultipleInteraction),
			Source:       card.Source,
			MainTitle:    card.MainTitle,
			SelectList:   card.SelectList,
			SubmitButton: card.SubmitButton,
			ReplaceText:  wx.CDATA(card.ReplaceText),
		},
	}
}
