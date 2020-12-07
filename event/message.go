package event

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"sort"
	"strings"
)

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
	EventSubscribe                  EventType = "subscribe"                    // 订阅
	EventUnSubscribe                EventType = "unsubscribe"                  // 取消订阅
	EventScan                       EventType = "SCAN"                         // 扫码
	EventLocation                   EventType = "LOCATION"                     // 上报地理位置
	EventClick                      EventType = "CLICK"                        // 点击自定义菜单
	EventView                       EventType = "VIEW"                         // 点击菜单跳转链接
	EventTemplateSendJobFinish      EventType = "TEMPLATESENDJOBFINISH"        // 模板消息发送完成
	EventQualificationVerifySuccess EventType = "qualification_verify_success" // 资质认证成功
	EventQualificationVerifyFail    EventType = "qualification_verify_fail"    // 资质认证失败
	EventNamingVerifySuccess        EventType = "naming_verify_success"        // 名称认证成功
	EventNamingVerifyFail           EventType = "naming_verify_fail"           // 名称认证失败
	EventAnnualRenew                EventType = "annual_renew"                 // 年审通知
	EventVerifyExpired              EventType = "verify_expired"               // 认证过期失效通知审通知
	EventCardPassCheck              EventType = "card_pass_check"              // 卡券通过审核
	EventCardNotPassCheck           EventType = "card_not_pass_check"          // 卡券未通过审核
	EventUserGetCard                EventType = "user_get_card"                // 用户领取卡券
	EventUserGiftingCard            EventType = "user_gifting_card"            // 用户转赠卡券
	EventUserDelCard                EventType = "user_del_card"                // 用户删除卡券
	EventUserConsumeCard            EventType = "user_consume_card"            // 用户核销卡券
	EventUserPayFromPayCell         EventType = "user_pay_from_pay_cell"       // 用户微信买单
	EventUserViewCard               EventType = "user_view_card"               // 用户点击会员卡
	EventUserEnterSessionFromCard   EventType = "user_enter_session_from_card" // 用户从卡券进入公众号会话
	EventUpdateMemberCard           EventType = "update_member_card"           // 会员卡内容更新
	EventCardSkuRemind              EventType = "card_sku_remind"              // 库存报警
	EventCardPayOrder               EventType = "card_pay_order"               // 券点流水详情事件
	EventSubmitMemberCardUserInfo   EventType = "submit_membercard_user_info"  // 会员卡激活
	EventWxaMediaCheck              EventType = "wxa_media_check"              // 校验图片/音频是否含有违法违规内容
)

// EventMessage 微信公众平台事件推送加密消息（兼容/安全模式）
type EventMessage struct {
	XMLName    xml.Name `xml:"xml"`
	ToUserName string   `xml:"ToUserName"`
	Encrypt    string   `xml:"Encrypt"`
}

// SignWithSHA1 事件消息sha1签名
func SignWithSHA1(token string, items ...string) string {
	items = append(items, token)

	sort.Strings(items)

	h := sha1.New()

	h.Write([]byte(strings.Join(items, "")))

	return hex.EncodeToString(h.Sum(nil))
}
