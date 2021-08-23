package common

type AttrType int

const (
	TextAttr  AttrType = 0
	WebAttr   AttrType = 1
	MinipAttr AttrType = 2
)

type ExternalProfile struct {
	ExternalCorpName string  `json:"external_corp_name"`
	WechatChannels   string  `json:"wechat_channels"`
	ExternalAttr     []*Attr `json:"external_attr"`
}

type ExtAttr struct {
	Attrs []*Attr `json:"attrs"`
}

type Attr struct {
	Type        AttrType   `json:"type"`
	Name        string     `json:"name"`
	Text        *AttrText  `json:"text,omitempty"`
	Web         *AttrWeb   `json:"web,omitempty"`
	Miniprogram *AttrMinip `json:"miniprogram,omitempty"`
}

type AttrText struct {
	Value string `json:"value"`
}

type AttrWeb struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type AttrMinip struct {
	Title    string `json:"title"`
	AppID    string `json:"appid"`
	Pagepath string `json:"pagepath"`
}
