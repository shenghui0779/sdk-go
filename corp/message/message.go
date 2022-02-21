package message

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/event"
	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type Text struct {
	Content string `json:"content"`
}

type Media struct {
	MediaID string `json:"media_id"`
}

type Video struct {
	MediaID     string `json:"media_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TextCard struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	BtnTxt      string `json:"btntxt"`
}

type News struct {
	Articles []*NewsArticle `json:"articles"`
}

type NewsArticle struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	PicURL      string `json:"picurl"`
	AppID       string `json:"appid"`
	PagePath    string `json:"pagepath"`
}

type MPNews struct {
	Articles []*MPNewsArticle `json:"articles"`
}

type MPNewsArticle struct {
	Title            string `json:"title"`
	ThumbMediaID     string `json:"thumb_media_id"`
	Author           string `json:"author"`
	ContentSourceURL string `json:"content_source_url"`
	Content          string `json:"content"`
	Digest           string `json:"digest"`
}

type MinipNotice struct {
	AppID             string   `json:"appid"`
	Page              string   `json:"page"`
	Title             string   `json:"title"`
	Description       string   `json:"description"`
	EmphasisFirstItem bool     `json:"emphasis_first_item"`
	ContentItem       []*MsgKV `json:"content_item"`
}

type MsgKV struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type MsgExtra struct {
	ToUser                 string
	ToParty                string
	ToTag                  string
	Safe                   int
	EnableIDTrans          int
	EnableDuplicateCheck   int
	DuplicateCheckInterval int
}

type Messasge struct {
	ToUser                 string        `json:"touser,omitempty"`
	ToParty                string        `json:"toparty,omitempty"`
	ToTag                  string        `json:"totag,omitempty"`
	MsgType                event.MsgType `json:"msgtype"`
	AgentID                int64         `json:"agentid"`
	Text                   *Text         `json:"text,omitempty"`
	Image                  *Media        `json:"image,omitempty"`
	Voice                  *Media        `json:"voice,omitempty"`
	Video                  *Video        `json:"video,omitempty"`
	File                   *Media        `json:"file,omitempty"`
	TextCard               *TextCard     `json:"textcard,omitempty"`
	News                   *News         `json:"news,omitempty"`
	MPNews                 *MPNews       `json:"mpnews,omitempty"`
	Markdown               *Text         `json:"markdown,omitempty"`
	MinipNotice            *MinipNotice  `json:"miniprogram_notice,omitempty"`
	TemplateCard           *TemplateCard `json:"template_card,omitempty"`
	Safe                   int           `json:"safe,omitempty"`
	EnableIDTrans          int           `json:"enable_id_trans,omitempty"`
	EnableDuplicateCheck   int           `json:"enable_duplicate_check,omitempty"`
	DuplicateCheckInterval int           `json:"duplicate_check_interval,omitempty"`
}

type ResultMsgSend struct {
	InvalidUser  string `json:"invaliduser"`
	InvalidParty string `json:"invalidparty"`
	InvalidTag   string `json:"invalidtag"`
	MsgID        string `json:"msgid"`
	ResponseCode string `json:"response_code"`
}

func SendText(agentidID int64, content string, extra *MsgExtra, result *ResultMsgSend) wx.Action {
	msg := &Messasge{
		MsgType: event.MsgText,
		AgentID: agentidID,
		Text: &Text{
			Content: content,
		},
	}

	if extra != nil {
		msg.ToUser = extra.ToUser
		msg.ToParty = extra.ToParty
		msg.ToTag = extra.ToTag
		msg.Safe = extra.Safe
		msg.EnableIDTrans = extra.EnableIDTrans
		msg.EnableDuplicateCheck = extra.EnableDuplicateCheck
		msg.DuplicateCheckInterval = extra.DuplicateCheckInterval
	}

	return wx.NewPostAction(urls.CorpMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendImage(agentidID int64, mediaID string, extra *MsgExtra, result *ResultMsgSend) wx.Action {
	msg := &Messasge{
		MsgType: event.MsgImage,
		AgentID: agentidID,
		Image: &Media{
			MediaID: mediaID,
		},
	}

	if extra != nil {
		msg.ToUser = extra.ToUser
		msg.ToParty = extra.ToParty
		msg.ToTag = extra.ToTag
		msg.Safe = extra.Safe
		msg.EnableIDTrans = extra.EnableIDTrans
		msg.EnableDuplicateCheck = extra.EnableDuplicateCheck
		msg.DuplicateCheckInterval = extra.DuplicateCheckInterval
	}

	return wx.NewPostAction(urls.CorpMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendVoice(agentidID int64, mediaID string, extra *MsgExtra, result *ResultMsgSend) wx.Action {
	msg := &Messasge{
		MsgType: event.MsgVoice,
		AgentID: agentidID,
		Voice: &Media{
			MediaID: mediaID,
		},
	}

	if extra != nil {
		msg.ToUser = extra.ToUser
		msg.ToParty = extra.ToParty
		msg.ToTag = extra.ToTag
		msg.Safe = extra.Safe
		msg.EnableIDTrans = extra.EnableIDTrans
		msg.EnableDuplicateCheck = extra.EnableDuplicateCheck
		msg.DuplicateCheckInterval = extra.DuplicateCheckInterval
	}

	return wx.NewPostAction(urls.CorpMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendVideo(agentidID int64, video *Video, extra *MsgExtra, result *ResultMsgSend) wx.Action {
	msg := &Messasge{
		MsgType: event.MsgVideo,
		AgentID: agentidID,
		Video:   video,
	}

	if extra != nil {
		msg.ToUser = extra.ToUser
		msg.ToParty = extra.ToParty
		msg.ToTag = extra.ToTag
		msg.Safe = extra.Safe
		msg.EnableIDTrans = extra.EnableIDTrans
		msg.EnableDuplicateCheck = extra.EnableDuplicateCheck
		msg.DuplicateCheckInterval = extra.DuplicateCheckInterval
	}

	return wx.NewPostAction(urls.CorpMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendFile(agentidID int64, mediaID string, extra *MsgExtra, result *ResultMsgSend) wx.Action {
	msg := &Messasge{
		MsgType: event.MsgText,
		AgentID: agentidID,
		File: &Media{
			MediaID: mediaID,
		},
	}

	if extra != nil {
		msg.ToUser = extra.ToUser
		msg.ToParty = extra.ToParty
		msg.ToTag = extra.ToTag
		msg.Safe = extra.Safe
		msg.EnableIDTrans = extra.EnableIDTrans
		msg.EnableDuplicateCheck = extra.EnableDuplicateCheck
		msg.DuplicateCheckInterval = extra.DuplicateCheckInterval
	}

	return wx.NewPostAction(urls.CorpMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendTextCard(agentidID int64, card *TextCard, extra *MsgExtra, result *ResultMsgSend) wx.Action {
	msg := &Messasge{
		MsgType:  event.MsgTextCard,
		AgentID:  agentidID,
		TextCard: card,
	}

	if extra != nil {
		msg.ToUser = extra.ToUser
		msg.ToParty = extra.ToParty
		msg.ToTag = extra.ToTag
		msg.Safe = extra.Safe
		msg.EnableIDTrans = extra.EnableIDTrans
		msg.EnableDuplicateCheck = extra.EnableDuplicateCheck
		msg.DuplicateCheckInterval = extra.DuplicateCheckInterval
	}

	return wx.NewPostAction(urls.CorpMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendNews(agentidID int64, articles []*NewsArticle, extra *MsgExtra, result *ResultMsgSend) wx.Action {
	msg := &Messasge{
		MsgType: event.MsgNews,
		AgentID: agentidID,
		News: &News{
			Articles: articles,
		},
	}

	if extra != nil {
		msg.ToUser = extra.ToUser
		msg.ToParty = extra.ToParty
		msg.ToTag = extra.ToTag
		msg.Safe = extra.Safe
		msg.EnableIDTrans = extra.EnableIDTrans
		msg.EnableDuplicateCheck = extra.EnableDuplicateCheck
		msg.DuplicateCheckInterval = extra.DuplicateCheckInterval
	}

	return wx.NewPostAction(urls.CorpMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendMPNews(agentidID int64, articles []*MPNewsArticle, extra *MsgExtra, result *ResultMsgSend) wx.Action {
	msg := &Messasge{
		MsgType: event.MsgMPNews,
		AgentID: agentidID,
		MPNews: &MPNews{
			Articles: articles,
		},
	}

	if extra != nil {
		msg.ToUser = extra.ToUser
		msg.ToParty = extra.ToParty
		msg.ToTag = extra.ToTag
		msg.Safe = extra.Safe
		msg.EnableIDTrans = extra.EnableIDTrans
		msg.EnableDuplicateCheck = extra.EnableDuplicateCheck
		msg.DuplicateCheckInterval = extra.DuplicateCheckInterval
	}

	return wx.NewPostAction(urls.CorpMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendMarkdown(agentidID int64, content string, extra *MsgExtra, result *ResultMsgSend) wx.Action {
	msg := &Messasge{
		MsgType: event.MsgMarkdown,
		AgentID: agentidID,
		Markdown: &Text{
			Content: content,
		},
	}

	if extra != nil {
		msg.ToUser = extra.ToUser
		msg.ToParty = extra.ToParty
		msg.ToTag = extra.ToTag
		msg.Safe = extra.Safe
		msg.EnableIDTrans = extra.EnableIDTrans
		msg.EnableDuplicateCheck = extra.EnableDuplicateCheck
		msg.DuplicateCheckInterval = extra.DuplicateCheckInterval
	}

	return wx.NewPostAction(urls.CorpMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendMinipNotice(agentidID int64, minipNotice *MinipNotice, extra *MsgExtra, result *ResultMsgSend) wx.Action {
	msg := &Messasge{
		MsgType:     event.MsgMinipNotice,
		AgentID:     agentidID,
		MinipNotice: minipNotice,
	}

	if extra != nil {
		msg.ToUser = extra.ToUser
		msg.ToParty = extra.ToParty
		msg.ToTag = extra.ToTag
		msg.Safe = extra.Safe
		msg.EnableIDTrans = extra.EnableIDTrans
		msg.EnableDuplicateCheck = extra.EnableDuplicateCheck
		msg.DuplicateCheckInterval = extra.DuplicateCheckInterval
	}

	return wx.NewPostAction(urls.CorpMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendTextNoticeCard(agentidID int64, card *TextNoticeCard, extra *MsgExtra, result *ResultMsgSend) wx.Action {
	msg := &Messasge{
		MsgType: event.MsgMinipNotice,
		AgentID: agentidID,
		TemplateCard: &TemplateCard{
			CardType:              CardTextNotice,
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

	if extra != nil {
		msg.ToUser = extra.ToUser
		msg.ToParty = extra.ToParty
		msg.ToTag = extra.ToTag
		msg.Safe = extra.Safe
		msg.EnableIDTrans = extra.EnableIDTrans
		msg.EnableDuplicateCheck = extra.EnableDuplicateCheck
		msg.DuplicateCheckInterval = extra.DuplicateCheckInterval
	}

	return wx.NewPostAction(urls.CorpMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendNewsNoticeCard(agentidID int64, card *NewsNoticeCard, extra *MsgExtra, result *ResultMsgSend) wx.Action {
	msg := &Messasge{
		MsgType: event.MsgMinipNotice,
		AgentID: agentidID,
		TemplateCard: &TemplateCard{
			CardType:              CardNewsNotice,
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

	if extra != nil {
		msg.ToUser = extra.ToUser
		msg.ToParty = extra.ToParty
		msg.ToTag = extra.ToTag
		msg.Safe = extra.Safe
		msg.EnableIDTrans = extra.EnableIDTrans
		msg.EnableDuplicateCheck = extra.EnableDuplicateCheck
		msg.DuplicateCheckInterval = extra.DuplicateCheckInterval
	}

	return wx.NewPostAction(urls.CorpMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendButtonInteractionCard(agentidID int64, card *ButtonInteractionCard, extra *MsgExtra, result *ResultMsgSend) wx.Action {
	msg := &Messasge{
		MsgType: event.MsgMinipNotice,
		AgentID: agentidID,
		TemplateCard: &TemplateCard{
			CardType:              CardButtonInteraction,
			Source:                card.Source,
			ActionMenu:            card.ActionMenu,
			MainTitle:             card.MainTitle,
			QuoteArea:             card.QuoteArea,
			SubTitleText:          card.SubTitleText,
			HorizontalContentList: card.HorizontalContentList,
			CardAction:            card.CardAction,
			ButtonSelection:       card.ButtonSelection,
			ButtonList:            card.ButtonList,
			ReplaceText:           card.ReplaceText,
		},
	}

	if extra != nil {
		msg.ToUser = extra.ToUser
		msg.ToParty = extra.ToParty
		msg.ToTag = extra.ToTag
		msg.Safe = extra.Safe
		msg.EnableIDTrans = extra.EnableIDTrans
		msg.EnableDuplicateCheck = extra.EnableDuplicateCheck
		msg.DuplicateCheckInterval = extra.DuplicateCheckInterval
	}

	return wx.NewPostAction(urls.CorpMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendVoteInteractionCard(agentidID int64, card *VoteInteractionCard, extra *MsgExtra, result *ResultMsgSend) wx.Action {
	msg := &Messasge{
		MsgType: event.MsgMinipNotice,
		AgentID: agentidID,
		TemplateCard: &TemplateCard{
			CardType:     CardButtonInteraction,
			Source:       card.Source,
			ActionMenu:   card.ActionMenu,
			MainTitle:    card.MainTitle,
			CheckBox:     card.CheckBox,
			SubmitButton: card.SubmitButton,
			ReplaceText:  card.ReplaceText,
		},
	}

	if extra != nil {
		msg.ToUser = extra.ToUser
		msg.ToParty = extra.ToParty
		msg.ToTag = extra.ToTag
		msg.Safe = extra.Safe
		msg.EnableIDTrans = extra.EnableIDTrans
		msg.EnableDuplicateCheck = extra.EnableDuplicateCheck
		msg.DuplicateCheckInterval = extra.DuplicateCheckInterval
	}

	return wx.NewPostAction(urls.CorpMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func SendMultipleInteractionCard(agentidID int64, card *MultipleInteractionCard, extra *MsgExtra, result *ResultMsgSend) wx.Action {
	msg := &Messasge{
		MsgType: event.MsgMinipNotice,
		AgentID: agentidID,
		TemplateCard: &TemplateCard{
			CardType:     CardButtonInteraction,
			Source:       card.Source,
			ActionMenu:   card.ActionMenu,
			MainTitle:    card.MainTitle,
			SelectList:   card.SelectList,
			SubmitButton: card.SubmitButton,
			ReplaceText:  card.ReplaceText,
		},
	}

	if extra != nil {
		msg.ToUser = extra.ToUser
		msg.ToParty = extra.ToParty
		msg.ToTag = extra.ToTag
		msg.Safe = extra.Safe
		msg.EnableIDTrans = extra.EnableIDTrans
		msg.EnableDuplicateCheck = extra.EnableDuplicateCheck
		msg.DuplicateCheckInterval = extra.DuplicateCheckInterval
	}

	return wx.NewPostAction(urls.CorpMessageSend,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsMsgRecall struct {
	MsgID string `json:"msgid"`
}

func Recall(msgid string) wx.Action {
	params := &ParamsMsgRecall{
		MsgID: msgid,
	}

	return wx.NewPostAction(urls.CorpMessageRecall,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsMsgStatics struct {
	TimeType int `json:"time_type"`
}

type ResultMsgStatics struct {
	Statics *MsgStatic `json:"statics"`
}

type MsgStatic struct {
	AgentID int64  `json:"agentid"`
	AppName string `json:"app_name"`
	Count   int64  `json:"count"`
}

func GetMessageStatics(timeType int, result *ResultMsgStatics) wx.Action {
	params := &ParamsMsgStatics{
		TimeType: timeType,
	}

	return wx.NewPostAction(urls.CorpMessageStaticsGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
