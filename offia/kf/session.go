package kf

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsSessionCreate struct {
	KFAccount string `json:"kf_account"` // 完整客服帐号，格式为：帐号前缀@公众号微信号
	OpenID    string `json:"openid"`     // 粉丝的openid
}

// CreateSession 创建会话
func CreateSession(account, openid string) wx.Action {
	params := &ParamsSessionCreate{
		KFAccount: account,
		OpenID:    openid,
	}

	return wx.NewPostAction(urls.OffiaKFSessionCreate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type ParamsSessionClose struct {
	KFAccount string `json:"kf_account"` // 完整客服帐号，格式为：帐号前缀@公众号微信号
	OpenID    string `json:"openid"`     // 粉丝的openid
}

// CloseSession 关闭会话
func CloseSession(account, openid string) wx.Action {
	params := &ParamsSessionClose{
		KFAccount: account,
		OpenID:    openid,
	}

	return wx.NewPostAction(urls.OffiaKFSessionClose,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

// Session 客服会话
type Session struct {
	KFAccount  string `json:"kf_account"`  // 正在接待的客服，为空表示没有人在接待
	OpenID     string `json:"openid"`      // 粉丝的openid
	CreateTime int64  `json:"createtime"`  // 会话接入的时间
	LatestTime int64  `json:"latest_time"` // 粉丝的最后一条消息的时间
}

// GetSession 获取客户会话状态
// 获取一个客户的会话，如果不存在，则kf_account为空。
func GetSession(openid string, result *Session) wx.Action {
	return wx.NewGetAction(urls.OffiaKFSessionGet,
		wx.WithQuery("openid", openid),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ResultSessionList struct {
	SessionList []*Session `json:"sessionlist"`
}

// GetSessionList 获取客服会话列表
// 完整客服帐号，格式为：帐号前缀@公众号微信号
func GetSessionList(account string, result *ResultSessionList) wx.Action {
	return wx.NewGetAction(urls.OffiaKFSessionList,
		wx.WithQuery("kf_account", account),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// WaitCase 客服未接入会话
type WaitCase struct {
	Count int        `json:"count"`
	List  []*Session `json:"waitcaselist"`
}

// GetWaitCase 获取未接入会话列表
// 最多返回100条数据，按照来访顺序
func GetWaitCase(result *WaitCase) wx.Action {
	return wx.NewGetAction(urls.OffiaKFWaitCase,
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// MsgRecord 客服聊天记录
type MsgRecord struct {
	Worker   string `json:"worker"`   // 完整客服帐号，格式为：帐号前缀@公众号微信号
	OpenID   string `json:"openid"`   // 用户标识
	OperCode int    `json:"opercode"` // 操作码（2002-客服发送信息，2003-客服接收消息）
	Text     string `json:"text"`     // 聊天记录
	Time     int64  `json:"time"`     // 操作时间，unix时间戳
}

type ParamsMsgRecordList struct {
	MsgID     int64 `json:"msgid"`
	StartTime int64 `json:"starttime"`
	EndTime   int64 `json:"endtime"`
	Number    int   `json:"number"`
}

type ResultMsgRecordList struct {
	MsgID      int64        `json:"msgid"`
	Number     int          `json:"number"`
	RecordList []*MsgRecord `json:"recordlist"`
}

// GetMsgRecordList 获取聊天记录（每次查询时段不能超过24小时）
// 聊天记录中，对于图片、语音、视频，分别展示成文本格式的[image]、[voice]、[video]。
func GetMsgRecordList(msgID, starttime, endtime int64, number int, result *ResultMsgRecordList) wx.Action {
	params := &ParamsMsgRecordList{
		MsgID:     msgID,
		StartTime: starttime,
		EndTime:   endtime,
		Number:    number,
	}

	return wx.NewPostAction(urls.OffiaKFMsgRecordList,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}
