package message

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type CardType string

const (
	CardTextNotice          CardType = "text_notice"
	CardNewsNotice          CardType = "news_notice"
	CardButtonInteraction   CardType = "button_interaction"
	CardVoteInteraction     CardType = "vote_interaction"
	CardMultipleInteraction CardType = "multiple_interaction"
)

type TemplateCard struct {
	CardType              CardType             `json:"card_type"`
	TaskID                string               `json:"task_id,omitempty"`
	Source                *CardSource          `json:"source"`
	ActionMenu            *ActionMenu          `json:"action_menu,omitempty"`
	MainTitle             *MainTitle           `json:"main_title"`
	QuoteArea             *QuoteArea           `json:"quote_area,omitempty"`
	EmphasisContent       *EmphasisContent     `json:"emphasis_content,omitempty"`
	SubTitleText          string               `json:"sub_title_text,omitempty"`
	ImageTextArea         *ImageTextArea       `json:"image_text_area,omitempty"`
	CardImage             *CardImage           `json:"card_image,omitempty"`
	VerticalContentList   []*VerticalContent   `json:"vertical_content_list,omitempty"`
	HorizontalContentList []*HorizontalContent `json:"horizontal_content_list,omitempty"`
	JumpList              []*CardJump          `json:"jump_list,omitempty"`
	CardAction            *CardAction          `json:"card_action,omitempty"`
	ButtonSelection       *ButtonSelection     `json:"button_selection,omitempty"`
	ButtonList            []*CardButton        `json:"button_list,omitempty"`
	CheckBox              *CheckBox            `json:"checkbox,omitempty"`
	SelectList            []*ButtonSelection   `json:"select_list,omitempty"`
	SubmitButton          *SubmitButton        `json:"submit_button,omitempty"`
	ReplaceText           string               `json:"replace_text,omitempty"`
}

type CardSource struct {
	IconURL   string `json:"icon_url,omitempty"`
	Desc      string `json:"desc,omitempty"`
	DescColor int    `json:"desc_color,omitempty"`
}

type ActionMenu struct {
	Desc       string        `json:"desc"`
	ActionList []*MenuAction `json:"action_list"`
}

type MenuAction struct {
	Text string `json:"text"`
	Key  string `json:"key"`
}

type MainTitle struct {
	Title string `json:"title"`
	Desc  string `json:"desc,omitempty"`
}

type QuoteArea struct {
	Type      int    `json:"type"`
	URL       string `json:"url"`
	Title     string `json:"title"`
	QuoteText string `json:"quote_text"`
}

type ImageTextArea struct {
	Type     int    `json:"type"`
	URL      string `json:"url"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	ImageURL string `json:"image_url"`
}

type CardImage struct {
	URL         string  `json:"url"`
	AspectRatio float64 `json:"aspect_ratio"`
}

type VerticalContent struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type HorizontalContent struct {
	Type    int    `json:"type,omitempty"`
	Keyname string `json:"keyname"`
	Value   string `json:"value"`
	UserID  string `json:"userid,omitempty"`
	MediaID string `json:"media_id,omitempty"`
	URL     string `json:"url,omitempty"`
}

type CardJump struct {
	Type     int    `json:"type"`
	Title    string `json:"title"`
	URL      string `json:"url,omitempty"`
	AppID    string `json:"appid,omitempty"`
	Pagepath string `json:"pagepath,omitempty"`
}

type CardAction struct {
	Type     int    `json:"type"`
	URL      string `json:"url,omitempty"`
	AppID    string `json:"appid,omitempty"`
	Pagepath string `json:"pagepath,omitempty"`
}

type EmphasisContent struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type CheckBox struct {
	QuestionKey string         `json:"question_key"`
	OptionList  []*CheckOption `json:"option_list"`
	Disable     bool           `json:"disable,omitempty"`
	Mode        int            `json:"mode,omitempty"`
}

type CheckOption struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	IsChecked bool   `json:"is_checked"`
}

type SubmitButton struct {
	Text string `json:"text"`
	Key  string `json:"key"`
}

type ButtonSelection struct {
	QuestionKey string          `json:"question_key"`
	Title       string          `json:"title,omitempty"`
	OptionList  []*SelectOption `json:"option_list"`
	SelectedID  string          `json:"selected_id,omitempty"`
	Disable     bool            `json:"disable,omitempty"`
}

type SelectOption struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type CardButton struct {
	Text  string `json:"text"`
	Style int    `json:"style"`
	Key   string `json:"key"`
}

type TextNoticeCard struct {
	Source                *CardSource
	ActionMenu            *ActionMenu
	MainTitle             *MainTitle
	QuoteArea             *QuoteArea
	EmphasisContent       *EmphasisContent
	SubTitleText          string
	HorizontalContentList []*HorizontalContent
	JumpList              []*CardJump
	CardAction            *CardAction
}

type NewsNoticeCard struct {
	Source                *CardSource
	ActionMenu            *ActionMenu
	MainTitle             *MainTitle
	QuoteArea             *QuoteArea
	ImageTextArea         *ImageTextArea
	CardImage             *CardImage
	VerticalContentList   []*VerticalContent
	HorizontalContentList []*HorizontalContent
	JumpList              []*CardJump
	CardAction            *CardAction
}

type ButtonInteractionCard struct {
	Source                *CardSource
	ActionMenu            *ActionMenu
	MainTitle             *MainTitle
	QuoteArea             *QuoteArea
	SubTitleText          string
	HorizontalContentList []*HorizontalContent
	CardAction            *CardAction
	ButtonSelection       *ButtonSelection
	ButtonList            []*CardButton
	ReplaceText           string
}

type VoteInteractionCard struct {
	Source       *CardSource
	ActionMenu   *ActionMenu
	MainTitle    *MainTitle
	CheckBox     *CheckBox
	SubmitButton *SubmitButton
	ReplaceText  string
}

type MultipleInteractionCard struct {
	Source       *CardSource
	ActionMenu   *ActionMenu
	MainTitle    *MainTitle
	SelectList   []*ButtonSelection
	SubmitButton *SubmitButton
	ReplaceText  string
}

type CardExtra struct {
	UserIDs  []string
	PartyIDs []int64
	TagIDs   []int64
	AtAll    int
}

type ParamsCardUpdate struct {
	UserIDs      []string          `json:"userids,omitempty"`
	PartyIDs     []int64           `json:"partyids,omitempty"`
	TagIDs       []int64           `json:"tagids,omitempty"`
	AtAll        int               `json:"atall,omitempty"`
	AgentID      int64             `json:"agentid"`
	ResponseCode string            `json:"response_code"`
	Button       *CardUpdateButton `json:"button,omitempty"`
	TemplateCard *TemplateCard     `json:"template_card,omitempty"`
}

type CardUpdateButton struct {
	ReplaceName string `json:"replace_name"`
}

type ResultCardUpdate struct {
	InvalidUser []string `json:"invaliduser"`
}

func UpdateCardButtonDisable(agentID int64, respCode, replaceName string, extra *CardExtra, result *ResultCardUpdate) wx.Action {
	params := &ParamsCardUpdate{
		AgentID:      agentID,
		ResponseCode: respCode,
		Button: &CardUpdateButton{
			ReplaceName: replaceName,
		},
	}

	if extra != nil {
		params.UserIDs = extra.UserIDs
		params.PartyIDs = extra.PartyIDs
		params.TagIDs = extra.TagIDs
		params.AtAll = extra.AtAll
	}

	return wx.NewPostAction(urls.CorpMessageUpdateTemplateCard,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func UpdateToTextNoticeCard(agentID int64, respCode string, card *TextNoticeCard, extra *CardExtra, result *ResultCardUpdate) wx.Action {
	params := &ParamsCardUpdate{
		AgentID:      agentID,
		ResponseCode: respCode,
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
		params.UserIDs = extra.UserIDs
		params.PartyIDs = extra.PartyIDs
		params.TagIDs = extra.TagIDs
		params.AtAll = extra.AtAll
	}

	return wx.NewPostAction(urls.CorpMessageUpdateTemplateCard,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func UpdateToNewsNoticeCard(agentID int64, respCode string, card *NewsNoticeCard, extra *CardExtra, result *ResultCardUpdate) wx.Action {
	params := &ParamsCardUpdate{
		AgentID:      agentID,
		ResponseCode: respCode,
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
		params.UserIDs = extra.UserIDs
		params.PartyIDs = extra.PartyIDs
		params.TagIDs = extra.TagIDs
		params.AtAll = extra.AtAll
	}

	return wx.NewPostAction(urls.CorpMessageUpdateTemplateCard,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func UpdateToButtonInteractionCard(agentID int64, respCode string, card *ButtonInteractionCard, extra *CardExtra, result *ResultCardUpdate) wx.Action {
	params := &ParamsCardUpdate{
		AgentID:      agentID,
		ResponseCode: respCode,
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
		params.UserIDs = extra.UserIDs
		params.PartyIDs = extra.PartyIDs
		params.TagIDs = extra.TagIDs
		params.AtAll = extra.AtAll
	}

	return wx.NewPostAction(urls.CorpMessageUpdateTemplateCard,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func UpdateToVoteInteractionCard(agentID int64, respCode string, card *VoteInteractionCard, extra *CardExtra, result *ResultCardUpdate) wx.Action {
	params := &ParamsCardUpdate{
		AgentID:      agentID,
		ResponseCode: respCode,
		TemplateCard: &TemplateCard{
			CardType:     CardVoteInteraction,
			Source:       card.Source,
			ActionMenu:   card.ActionMenu,
			MainTitle:    card.MainTitle,
			CheckBox:     card.CheckBox,
			SubmitButton: card.SubmitButton,
			ReplaceText:  card.ReplaceText,
		},
	}

	if extra != nil {
		params.UserIDs = extra.UserIDs
		params.PartyIDs = extra.PartyIDs
		params.TagIDs = extra.TagIDs
		params.AtAll = extra.AtAll
	}

	return wx.NewPostAction(urls.CorpMessageUpdateTemplateCard,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

func UpdateToMultipleInteractionCard(agentID int64, respCode string, card *MultipleInteractionCard, extra *CardExtra, result *ResultCardUpdate) wx.Action {
	params := &ParamsCardUpdate{
		AgentID:      agentID,
		ResponseCode: respCode,
		TemplateCard: &TemplateCard{
			CardType:     CardMultipleInteraction,
			Source:       card.Source,
			ActionMenu:   card.ActionMenu,
			MainTitle:    card.MainTitle,
			SelectList:   card.SelectList,
			SubmitButton: card.SubmitButton,
			ReplaceText:  card.ReplaceText,
		},
	}

	if extra != nil {
		params.UserIDs = extra.UserIDs
		params.PartyIDs = extra.PartyIDs
		params.TagIDs = extra.TagIDs
		params.AtAll = extra.AtAll
	}

	return wx.NewPostAction(urls.CorpMessageUpdateTemplateCard,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
