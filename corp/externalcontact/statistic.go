package externalcontact

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsUserBehaviorDataGet struct {
	UserID    []string `json:"userid,omitempty"`
	PartyID   []string `json:"partyid,omitempty"`
	StartTime int64    `json:"start_time"`
	EndTime   int64    `json:"end_time"`
}

type UserBehaviorData struct {
	StatTime            int64   `json:"stat_time"`
	ChatCnt             int     `json:"chat_cnt"`
	MessageCnt          int     `json:"message_cnt"`
	ReplyPercentage     float64 `json:"reply_percentage"`
	AvgReplyTime        int     `json:"avg_reply_time"`
	NegativeFeedbackCnt int     `json:"negative_feedback_cnt"`
	NewApplyCnt         int     `json:"new_apply_cnt"`
	NewContactCnt       int     `json:"new_contact_cnt"`
}

type ResultUserBehaviorDataGet struct {
	BehaviorData []*UserBehaviorData `json:"behavior_data"`
}

func GetUserBehaviorData(params *ParamsUserBehaviorDataGet, result *ResultUserBehaviorDataGet) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactGetUserBehaviorData,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsGroupChatStatistic struct {
	DayBeginTime int64                 `json:"day_begin_time"`
	DayEndTime   int64                 `json:"day_end_time,omitempty"`
	OwnerFilter  *GroupChatOwnerFilter `json:"owner_filter"`
	OrderBy      int                   `json:"order_by,omitempty"`
	OrderAsc     int                   `json:"order_asc,omitempty"`
	Offset       int                   `json:"offset,omitempty"`
	Limit        int                   `json:"limit,omitempty"`
}

type GroupChatStatisticItem struct {
	Owner string                  `json:"owner"`
	Data  *GroupChatStatisticData `json:"data"`
}

type GroupChatStatisticData struct {
	NewChatCnt            int `json:"new_chat_cnt"`
	ChatTotal             int `json:"chat_total"`
	ChatHasMsg            int `json:"chat_has_msg"`
	NewMemberCnt          int `json:"new_member_cnt"`
	MemberTotal           int `json:"member_total"`
	MemberHasMsg          int `json:"member_has_msg"`
	MsgTotal              int `json:"msg_total"`
	MigrateTraineeChatCnt int `json:"migrate_trainee_chat_cnt"`
}

type ResultGroupChatStatistic struct {
	Total      int                       `json:"total"`
	NextOffset int                       `json:"next_offset"`
	Items      []*GroupChatStatisticItem `json:"items"`
}

func GetGroupChatStatistic(params *ParamsGroupChatStatistic, result *ResultGroupChatStatistic) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactGroupChatStatistic,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsGroupChatStatisticByDay struct {
	DayBeginTime int64                 `json:"day_begin_time"`
	DayEndTime   int64                 `json:"day_end_time,omitempty"`
	OwnerFilter  *GroupChatOwnerFilter `json:"owner_filter"`
}

type GroupChatStatisticByDayItem struct {
	StartTime int64                   `json:"start_time"`
	Data      *GroupChatStatisticData `json:"data"`
}

type ResultGroupChatStatisticByDay struct {
	Items []*GroupChatStatisticByDayItem `json:"items"`
}

func GetGroupChatStatisticByDay(params *ParamsGroupChatStatisticByDay, result *ResultGroupChatStatisticByDay) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactGroupChatStatisticByDay,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
