package message

import (
	"encoding/json"

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
