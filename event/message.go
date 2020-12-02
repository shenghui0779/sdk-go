package event

import "encoding/xml"

// MessageType 消息类型
type MessageType string

// 微信支持的消息类型
const (
	MessageText       MessageType = "text"       // 文本消息
	MessageImage      MessageType = "image"      // 图片消息
	MessageVoice      MessageType = "voice"      // 语音消息
	MessageVideo      MessageType = "video"      // 视频消息
	MessageShortVideo MessageType = "shortvideo" // 小视频消息
	MessageLocation   MessageType = "location"   // 地理位置消息
	MessageLink       MessageType = "link"       // 链接消息
	MessageMusic      MessageType = "music"      // 音乐消息
	MessageNews       MessageType = "news"       // 图文消息
	MessageWXCard     MessageType = "wxcard"     // 卡券，客服消息时使用
	MessageEvent      MessageType = "event"      // 事件推送
)

// EventType 事件类型
type EventType string

// 微信支持的事件类型
const (
	EventSubscribe             EventType = "subscribe"             // 订阅
	EventUnSubscribe           EventType = "unsubscribe"           // 取消订阅
	EventScan                  EventType = "SCAN"                  // 扫码
	EventLocation              EventType = "LOCATION"              // 上报地理位置
	EventClick                 EventType = "CLICK"                 // 点击自定义菜单
	EventView                  EventType = "VIEW"                  // 点击菜单跳转链接
	EventTemplateSendJobFinish EventType = "TEMPLATESENDJOBFINISH" // 模板消息发送完成
	EventWxaMediaCheck         EventType = "wxa_media_check"       // 校验图片/音频是否含有违法违规内容
)

// Message 微信公众平台事件推送消息
type Message struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   string      `xml:"ToUserName"`
	FromUserName string      `xml:"FromUserName"`
	CreateTime   int64       `xml:"CreateTime"`
	MsgType      MessageType `xml:"MsgType"`
	// 普通消息
	MsgID        int64   `xml:"MsgId"`
	Content      string  `xml:"Content"`
	PicURL       string  `xml:"PicUrl"`
	MediaID      string  `xml:"MediaId"`
	Format       string  `xml:"Format"`
	Recognition  string  `xml:"Recognition"`
	ThumbMediaID string  `xml:"ThumbMediaId"`
	LocationX    float64 `xml:"Location_X"`
	LocationY    float64 `xml:"Location_Y"`
	Scale        int     `xml:"Scale"`
	Label        string  `xml:"Label"`
	Title        string  `xml:"Title"`
	Description  string  `xml:"Description"`
	URL          string  `xml:"Url"`
	// 事件消息
	Event         EventType `xml:"Event"`
	EventKey      string    `xml:"EventKey"`
	Ticket        string    `xml:"Ticket"`
	Latitude      float64   `xml:"Latitude"`
	Longitude     float64   `xml:"Longitude"`
	Precision     float64   `xml:"Precision"`
	IsRisky       int       `xml:"isrisky"`
	ExtraInfoJSON string    `xml:"extra_info_json"`
	AppID         string    `xml:"app_id"`
	TraceID       string    `xml:"trace_id"`
	StatusCode    int       `xml:"status_code"`
}
