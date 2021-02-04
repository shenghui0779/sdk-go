package oa

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/wx"
	"github.com/tidwall/gjson"
)

// KFInviteStatus 客服邀请状态
type KFInviteStatus string

// 微信支持的客服邀请状态
const (
	InviteWaiting  KFInviteStatus = "waiting"  // 待确认
	InviteRejected KFInviteStatus = "rejected" // 被拒绝
	InviteExpired  KFInviteStatus = "expired"  // 已过期
)

// KFAccount 客服账号
type KFAccount struct {
	ID               string         `json:"kf_id"`              // 客服编号
	Account          string         `json:"kf_account"`         // 完整客服帐号，格式为：帐号前缀@公众号微信号
	Nickname         string         `json:"kf_nick"`            // 客服昵称
	HeadImgURL       string         `json:"kf_headimgurl"`      // 客服头像
	Weixin           string         `json:"kf_wx"`              // 如果客服帐号已绑定了客服人员微信号， 则此处显示微信号
	InviteWeixin     string         `json:"invite_wx"`          // 如果客服帐号尚未绑定微信号，但是已经发起了一个绑定邀请， 则此处显示绑定邀请的微信号
	InviteExpireTime int64          `json:"invite_expire_time"` // 如果客服帐号尚未绑定微信号，但是已经发起过一个绑定邀请， 邀请的过期时间，为unix 时间戳
	InviteStatus     KFInviteStatus `json:"invite_status"`      // 邀请的状态，有等待确认“waiting”，被拒绝“rejected”， 过期“expired”
}

// GetKFAccountList 获取客服列表
func GetKFAccountList(dest *[]*KFAccount) wx.Action {
	return wx.NewAction(KFAccountListURL,
		wx.WithMethod(wx.MethodGet),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal([]byte(gjson.GetBytes(resp, "kf_list").Raw), dest)
		}),
	)
}

// KFOnline 在线客服
type KFOnline struct {
	ID           string `json:"kf_id"`         // 客服编号
	Account      string `json:"kf_account"`    // 完整客服帐号，格式为：帐号前缀@公众号微信号
	Status       int    `json:"status"`        // 客服在线状态，目前为：1-web在线
	AcceptedCase int    `json:"accepted_case"` // 客服当前正在接待的会话数
}

// GetKFOnlineList 获取客服在线列表
func GetKFOnlineList(dest *[]*KFOnline) wx.Action {
	return wx.NewAction(KFOnlineListURL,
		wx.WithMethod(wx.MethodGet),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal([]byte(gjson.GetBytes(resp, "kf_online_list").Raw), dest)
		}),
	)
}

// AddKFAccount 添加客服账号
// 完整客服帐号，格式为：帐号前缀@公众号微信号，帐号前缀最多10个字符，必须是英文、数字字符或者下划线，后缀为公众号微信号，长度不超过30个字符
// 客服昵称，最长16个字
func AddKFAccount(account, nickname string) wx.Action {
	return wx.NewAction(KFAccountAddURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(wx.X{
				"kf_account": account,
				"nickname":   nickname,
			})
		}),
	)
}

// UpdateKFAccount 设置客服信息
// 完整客服帐号，格式为：帐号前缀@公众号微信号，帐号前缀最多10个字符，必须是英文、数字字符或者下划线，后缀为公众号微信号，长度不超过30个字符
// 客服昵称，最长16个字
func UpdateKFAccount(account, nickname string) wx.Action {
	return wx.NewAction(KFAccountUpdateURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(wx.X{
				"kf_account": account,
				"nickname":   nickname,
			})
		}),
	)
}

// InviteKFWorker 邀请绑定客服帐号
// 新添加的客服帐号是不能直接使用的，只有客服人员用微信号绑定了客服账号后，方可登录Web客服进行操作。
// 发起一个绑定邀请到客服人员微信号，客服人员需要在微信客户端上用该微信号确认后帐号才可用。
// 尚未绑定微信号的帐号可以进行绑定邀请操作，邀请未失效时不能对该帐号进行再次绑定微信号邀请。
// 完整客服帐号，格式为：帐号前缀@公众号微信号
func InviteKFWorker(account, inviteWeixin string) wx.Action {
	return wx.NewAction(KFInviteURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(wx.X{
				"kf_account": account,
				"invite_wx":  inviteWeixin,
			})
		}),
	)
}

// UploadKFAvatar 上传客服头像
// 完整客服帐号，格式为：帐号前缀@公众号微信号
func UploadKFAvatar(account, filename string) wx.Action {
	return wx.NewAction(KFAvatarUploadURL,
		wx.WithMethod(wx.MethodUpload),
		wx.WithQuery("kf_account", account),
		wx.WithUploadForm("media", filename, nil),
	)
}

// DeleteKFAccount 删除客服帐号
// 完整客服帐号，格式为：帐号前缀@公众号微信号
func DeleteKFAccount(account string) wx.Action {
	return wx.NewAction(KFDeleteURL,
		wx.WithMethod(wx.MethodGet),
		wx.WithQuery("kf_account", account),
	)
}

// KFSession 客服会话
type KFSession struct {
	KFAccount  string `json:"kf_account"`  // 正在接待的客服，为空表示没有人在接待
	OpenID     string `json:"openid"`      // 粉丝的openid
	CreateTime int64  `json:"createtime"`  // 会话接入的时间
	LatestTime int64  `json:"latest_time"` // 粉丝的最后一条消息的时间
}

// CreateKFSession 创建会话
// 完整客服帐号，格式为：帐号前缀@公众号微信号
func CreateKFSession(account, openid string) wx.Action {
	return wx.NewAction(KFSessionCreateURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(wx.X{
				"kf_account": account,
				"openid":     openid,
			})
		}),
	)
}

// CloseKFSession 关闭会话
// 完整客服帐号，格式为：帐号前缀@公众号微信号
func CloseKFSession(account, openid string) wx.Action {
	return wx.NewAction(KFSessionCloseURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(wx.X{
				"kf_account": account,
				"openid":     openid,
			})
		}),
	)
}

// GetKFSession 获取客户会话状态
// 获取一个客户的会话，如果不存在，则kf_account为空。
func GetKFSession(dest *KFSession, openid string) wx.Action {
	return wx.NewAction(KFSessionGetURL,
		wx.WithMethod(wx.MethodGet),
		wx.WithQuery("openid", openid),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// GetKFSessionList 获取客服会话列表
// 完整客服帐号，格式为：帐号前缀@公众号微信号
func GetKFSessionList(dest *[]*KFSession, account string) wx.Action {
	return wx.NewAction(KFSessionListURL,
		wx.WithMethod(wx.MethodGet),
		wx.WithQuery("kf_account", account),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal([]byte(gjson.GetBytes(resp, "sessionlist").Raw), dest)
		}),
	)
}

// KFWaitCase 客服未接入会话
type KFWaitCase struct {
	Count int          `json:"count"`
	List  []*KFSession `json:"waitcaselist"`
}

// GetKFWaitCase 获取未接入会话列表
// 最多返回100条数据，按照来访顺序
func GetKFWaitCase(dest *KFWaitCase) wx.Action {
	return wx.NewAction(KFWaitCaseURL,
		wx.WithMethod(wx.MethodGet),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// KFMsgRecord 客服聊天记录
type KFMsgRecord struct {
	Worker   string `json:"worker"`   // 完整客服帐号，格式为：帐号前缀@公众号微信号
	OpenID   string `json:"openid"`   // 用户标识
	OperCode int    `json:"opercode"` // 操作码（2002-客服发送信息，2003-客服接收消息）
	Text     string `json:"text"`     // 聊天记录
	Time     int64  `json:"time"`     // 操作时间，unix时间戳
}

// KFMsgRecordList 客服聊天记录列表
type KFMsgRecordList struct {
	MsgID      int64          `json:"msgid"`
	Number     int            `json:"number"`
	RecordList []*KFMsgRecord `json:"recordlist"`
}

// GetKFMsgRecordList 获取聊天记录（每次查询时段不能超过24小时）
// 聊天记录中，对于图片、语音、视频，分别展示成文本格式的[image]、[voice]、[video]。
func GetKFMsgRecordList(dest *KFMsgRecordList, msgid, starttime, endtime int64, number int) wx.Action {
	return wx.NewAction(KFMsgRecordListURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(wx.X{
				"msgid":     msgid,
				"starttime": starttime,
				"endtime":   endtime,
				"number":    number,
			})
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}
