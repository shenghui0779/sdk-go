package message

type CardType string

const (
	CardTextNotice          CardType = "text_notice"
	CardNewsNotice          CardType = "news_notice"
	CardButtonInteraction   CardType = "button_interaction"
	CardVoteInteraction     CardType = "vote_interaction"
	CardMultipleInteraction CardType = "multiple_interaction"
)

type TemplateCard struct {
	CardType CardType
	Source   *CardSource
}

type CardSource struct {
	IconURL   string `json:"icon_url"`
	Desc      string `json:"desc"`
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
	Desc  string `json:"desc"`
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
	UserID  string `json:"user_id,omitempty"`
	MediaID string `json:"media_id,omitempty"`
	URL     string `json:"url,omitempty"`
}

type CardJump struct {
	Type     int    `json:"type"`
	Title    string `json:"title"`
	URL      string `json:"url,omitempty"`
	AppID    string `json:"app_id,omitempty"`
	Pagepath string `json:"pagepath,omitempty"`
}

type CardAction struct {
	Type     int    `json:"type"`
	URL      string `json:"url,omitempty"`
	AppID    string `json:"app_id,omitempty"`
	Pagepath string `json:"pagepath,omitempty"`
}

type EmphasisContent struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type CheckBox struct {
	QuestionKey string         `json:"question_key"`
	OptionList  []*CheckOption `json:"option_list"`
	Mode        int            `json:"mode"`
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
	Title       string          `json:"title"`
	OptionList  []*SelectOption `json:"option_list"`
	SelectID    string          `json:"select_id"`
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
