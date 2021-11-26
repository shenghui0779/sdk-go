package message

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type TextMessage struct {
	Content string `json:"content"`
}

type ImageMessage struct {
	MediaID string `json:"media_id"`
}

type VoiceMessage struct {
	MediaID string `json:"media_id"`
}

type VideoMessage struct {
	MediaID     string `json:"media_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type FileMessage struct {
	MediaID string `json:"media_id"`
}

type TextCardMessage struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	BtnTxt      string `json:"btntxt"`
}

type NewsMessage struct {
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

type MPNewsMessage struct {
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

type MarkdownMessage struct {
	Content string `json:"content"`
}

type MinipNoticeMessage struct {
	AppID             string            `json:"appid"`
	Page              string            `json:"page"`
	Title             string            `json:"title"`
	Description       string            `json:"description"`
	EmphasisFirstItem bool              `json:"emphasis_first_item"`
	ContentItem       map[string]string `json:"content_item"`
}

type ParamsMessageStaticsGet struct {
	TimeType int `json:"time_type"`
}

type ResultMessageStaticsGet struct {
	Statics *MessageStatic `json:"statics"`
}

type MessageStatic struct {
	AgentID int64  `json:"agentid"`
	AppName string `json:"app_name"`
	Count   int64  `json:"count"`
}

func GetMessageStatics(params *ParamsMessageStaticsGet, result *ResultMessageStaticsGet) wx.Action {
	return wx.NewPostAction(urls.CorpMessageStaticsGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
