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
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type TextCard struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	BtnTxt      string `json:"btntxt,omitempty"`
}

type News struct {
	Articles []*NewsArticle `json:"articles"`
}

type NewsArticle struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty"`
	PicURL      string `json:"picurl,omitempty"`
	AppID       string `json:"appid,omitempty"`
	PagePath    string `json:"pagepath,omitempty"`
	BtnTxt      string `json:"btntxt,omitempty"`
}

type MPNews struct {
	Articles []*MPNewsArticle `json:"articles"`
}

type MPNewsArticle struct {
	Title            string `json:"title"`
	ThumbMediaID     string `json:"thumb_media_id"`
	Author           string `json:"author,omitempty"`
	ContentSourceURL string `json:"content_source_url,omitempty"`
	Content          string `json:"content"`
	Digest           string `json:"digest,omitempty"`
}

type Miniprogram struct {
	AppID        string `json:"appid"`
	Title        string `json:"title,omitempty"`
	ThumbMediaID string `json:"thumb_media_id"`
	PagePath     string `json:"pagepath"`
}

type MinipNotice struct {
	AppID             string   `json:"appid"`
	Page              string   `json:"page,omitempty"`
	Title             string   `json:"title"`
	Description       string   `json:"description,omitempty"`
	EmphasisFirstItem bool     `json:"emphasis_first_item,omitempty"`
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
	AgentID                int64         `json:"agentid,omitempty"`
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

// SendText 发送应用消息（文本消息）
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
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// SendImage 发送应用消息（图片消息）
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
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// SendVoice 发送应用消息（语音消息）
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
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// SendVideo 发送应用消息（视频消息）
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
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// SendFile 发送应用消息（文件消息）
func SendFile(agentidID int64, mediaID string, extra *MsgExtra, result *ResultMsgSend) wx.Action {
	msg := &Messasge{
		MsgType: event.MsgFile,
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
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// SendTextCard 发送应用消息（文本卡片消息）
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
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// SendNews 发送应用消息（图文消息）
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
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// SendMPNews 发送应用消息（图文消息 - mpnews）
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
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// SendMarkdown 发送应用消息（markdown消息）
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
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// SendMinipNotice 发送应用消息（小程序通知消息）
func SendMinipNotice(notice *MinipNotice, extra *MsgExtra, result *ResultMsgSend) wx.Action {
	msg := &Messasge{
		MsgType:     event.MsgMinipNotice,
		MinipNotice: notice,
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
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// SendTextNoticeCard 发送应用消息（模板卡片消息 - 文本通知型）
func SendTextNoticeCard(agentidID int64, taskID string, card *TextNoticeCard, extra *MsgExtra, result *ResultMsgSend) wx.Action {
	msg := &Messasge{
		MsgType: event.MsgTemplateCard,
		AgentID: agentidID,
		TemplateCard: &TemplateCard{
			TaskID:                taskID,
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
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// SendNewsNoticeCard 发送应用消息（模板卡片消息 - 图文展示型）
func SendNewsNoticeCard(agentidID int64, taskID string, card *NewsNoticeCard, extra *MsgExtra, result *ResultMsgSend) wx.Action {
	msg := &Messasge{
		MsgType: event.MsgTemplateCard,
		AgentID: agentidID,
		TemplateCard: &TemplateCard{
			CardType:              CardNewsNotice,
			TaskID:                taskID,
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
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// SendButtonInteractionCard 发送应用消息（模板卡片消息 - 按钮交互型）
func SendButtonInteractionCard(agentidID int64, taskID string, card *ButtonInteractionCard, extra *MsgExtra, result *ResultMsgSend) wx.Action {
	msg := &Messasge{
		MsgType: event.MsgTemplateCard,
		AgentID: agentidID,
		TemplateCard: &TemplateCard{
			CardType:              CardButtonInteraction,
			TaskID:                taskID,
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
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// SendVoteInteractionCard 发送应用消息（模板卡片消息 - 投票选择型）
func SendVoteInteractionCard(agentidID int64, taskID string, card *VoteInteractionCard, extra *MsgExtra, result *ResultMsgSend) wx.Action {
	msg := &Messasge{
		MsgType: event.MsgTemplateCard,
		AgentID: agentidID,
		TemplateCard: &TemplateCard{
			CardType:     CardVoteInteraction,
			TaskID:       taskID,
			Source:       card.Source,
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
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// SendMultipleInteractionCard 发送应用消息（模板卡片消息 - 多项选择型）
func SendMultipleInteractionCard(agentidID int64, taskID string, card *MultipleInteractionCard, extra *MsgExtra, result *ResultMsgSend) wx.Action {
	msg := &Messasge{
		MsgType: event.MsgTemplateCard,
		AgentID: agentidID,
		TemplateCard: &TemplateCard{
			CardType:     CardMultipleInteraction,
			TaskID:       taskID,
			Source:       card.Source,
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
			return wx.MarshalNoEscapeHTML(msg)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsMsgRecall struct {
	MsgID string `json:"msgid"`
}

// Recall 撤回应用消息
func Recall(msgid string) wx.Action {
	params := &ParamsMsgRecall{
		MsgID: msgid,
	}

	return wx.NewPostAction(urls.CorpMessageRecall,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}
