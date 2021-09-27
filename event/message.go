package event

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"sort"
	"strings"
)

// MsgType 消息类型
type MsgType string

// 微信支持的消息类型
const (
	MsgText                 MsgType = "text"                   // 文本消息
	MsgImage                MsgType = "image"                  // 图片消息
	MsgVoice                MsgType = "voice"                  // 语音消息
	MsgVideo                MsgType = "video"                  // 视频消息
	MsgShortVideo           MsgType = "shortvideo"             // 小视频消息
	MsgLocation             MsgType = "location"               // 地理位置消息
	MsgLink                 MsgType = "link"                   // 链接消息
	MsgMusic                MsgType = "music"                  // 音乐消息
	MsgNews                 MsgType = "news"                   // 图文消息
	MsgWXCard               MsgType = "wxcard"                 // 卡券，客服消息时使用
	MsgFile                 MsgType = "file"                   // 文件消息
	MsgMinip                MsgType = "miniprogram"            // 小程序消息
	MsgMinipPage            MsgType = "miniprogrampage"        // 小程序卡片消息
	MsgUserEnterTempSession MsgType = "user_enter_tempsession" // 进入会话事件
	MsgMenu                 MsgType = "menu"                   // 菜单消息
	MsgBussinessCard        MsgType = "business_card"          // 企业名片消息
	MsgTextCard             MsgType = "textcard"               // 文本卡片消息
	MsgMPNews               MsgType = "mpnews"                 // 图文消息，跟普通的图文消息一致，唯一的差异是图文内容存储在企业微信
	MsgMarkdown             MsgType = "markdown"               // markdown消息
	MsgMinipNotice          MsgType = "miniprogram_notice"     // 小程序通知消息
	MsgTemplateCard         MsgType = "template_card"          // 模板卡片消息
	MsgEvent                MsgType = "event"                  // 事件推送
)

// EventType 事件类型
type EventType string

// 微信支持的事件类型（统一小写匹配）
const (
	EventSubscribe                  EventType = "subscribe"                    // 关注
	EventUnSubscribe                EventType = "unsubscribe"                  // 取消关注
	EventScan                       EventType = "scan"                         // 扫码
	EventLocation                   EventType = "location"                     // 上报地理位置
	EventClick                      EventType = "click"                        // 点击自定义菜单
	EventView                       EventType = "view"                         // 点击菜单跳转链接
	EventTemplateSendJobFinish      EventType = "templatesendjobfinish"        // 模板消息发送完成
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
	EventKFMsgOREvent               EventType = "kf_msg_or_event"              // 企业微信客服
	EventEnterSession               EventType = "enter_session"                // 用户进入会话
	EventMsgSendFail                EventType = "msg_send_fail"                // 消息发送失败
	EventServicerStatusChange       EventType = "servicer_status_change"       // 客服人员接待状态变更
	EventSessionStatusChange        EventType = "session_status_change"        // 会话状态变更
	EventSwitchWorkbenchMode        EventType = "switch_workbench_mode"        // 切换工作台自定义模式
	EventEnterAgent                 EventType = "enter_agent"                  // 进入应用
	EventBatchJobResult             EventType = "batch_job_result"             // 异步任务完成事件推送
	EventChangeContact              EventType = "change_contact"               // 通讯录变更
	EventScanCodePush               EventType = "scancode_push"                // 扫码推事件
	EventScanCodeWaitMsg            EventType = "scancode_waitmsg"             // 扫码推事件且弹出“消息接收中”提示框
	EventPicSysPhoto                EventType = "pic_sysphoto "                // 弹出系统拍照发图
	EventPicPhotoOrAlbum            EventType = "pic_photo_or_album"           // 弹出拍照或者相册发图
	EventPicWeixin                  EventType = "pic_weixin"                   // 弹出微信相册发图器
	EventLocationSelect             EventType = "location_select"              // 弹出地理位置选择器
	EventOpenApprovalChange         EventType = "open_approval_change"         // 审批状态通知
	EventSysApprovalChange          EventType = "sys_approval_change"          // 审批申请状态变化回调
	EventShareAgentChange           EventType = "share_agent_change"           // 共享应用
	EventTemplateCard               EventType = "template_card_event"          // 模板卡片事件推送
	EventModifyCalendar             EventType = "modify_calendar"              // 修改日历
	EventDeleteCalendar             EventType = "delete_calendar"              // 删除日历
	EventAddSchedule                EventType = "add_schedule"                 // 添加日程
	EventModifySchedule             EventType = "modify_schedule"              // 修改日程
	EventDeleteSchedule             EventType = "delete_schedule"              // 删除日程
)

// JobType 任务类型
type JobType string

// 微信支持的任务类型
const (
	JobSyncUser     JobType = "sync_user"     // 增量更新成员
	JobReplaceUser  JobType = "replace_user"  // 全量覆盖成员
	JobInviteUser   JobType = "invite_user"   // 邀请成员关注
	JobReplaceParty JobType = "replace_party" // 全量覆盖部门
)

// CorpCardType 企业微信模板卡片消息类型
type CorpCardType string

const (
	CardTextNotice          CorpCardType = "text_notice"          // 文本通知型
	CardNewsNotice          CorpCardType = "news_notice"          // 图文展示型
	CardButtonInteraction   CorpCardType = "button_interaction"   // 按钮交互型
	CardVoteInteraction     CorpCardType = "vote_interaction"     // 投票选择型
	CardMultipleInteraction CorpCardType = "multiple_interaction" // 多项选择型
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
