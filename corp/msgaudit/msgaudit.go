package msgaudit

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsPermitUserListGet struct {
	Type int `json:"type"`
}

type ResultPermitUserListGet struct {
	IDs []string `json:"ids"`
}

func GetPermitUserList(params *ParamsPermitUserListGet, result *ResultPermitUserListGet) wx.Action {
	return wx.NewPostAction(urls.CorpMsgAuditGetPermitUserList,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type CheckInfo struct {
	UserID         string `json:"userid"`
	ExternalOpenID string `json:"externalopenid"`
}

type AgreeInfo struct {
	StatusChangeTime int64  `json:"status_change_time"`
	UserID           string `json:"userid"`
	ExternalOpenID   string `json:"externalopenid"`
	AgreeStatus      string `json:"agree_status"`
}

type ParamsSingleAgreeCheck struct {
	Info []*CheckInfo `json:"info"`
}

type ResultSingleAgreeCheck struct {
	AgreeInfo []*AgreeInfo `json:"agree_info"`
}

func CheckSingleAgree(params *ParamsSingleAgreeCheck, result *ResultSingleAgreeCheck) wx.Action {
	return wx.NewPostAction(urls.CorpMsgAuditCheckSingleAgree,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsRoomAgreeCheck struct {
	RoomID string `json:"roomid"`
}

type ResultRoomAgreeCheck struct {
	AgreeInfo []*AgreeInfo `json:"agree_info"`
}

func CheckRoomAgree(params *ParamsRoomAgreeCheck, result *ResultRoomAgreeCheck) wx.Action {
	return wx.NewPostAction(urls.CorpMsgAuditCheckRoomAgree,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type GroupMember struct {
	MemberID string `json:"memberid"`
	JoinTime int64  `json:"jointime"`
}

type ParamsGroupChatGet struct {
	RoomID string `json:"roomid"`
}

type ResultGroupChatGet struct {
	RoomName       string         `json:"roomname"`
	Creator        string         `json:"creator"`
	RoomCreateTime int64          `json:"room_create_time"`
	Notice         string         `json:"notice"`
	Members        []*GroupMember `json:"members"`
}

func GetGroupChat(params *ParamsGroupChatGet, result *ResultGroupChatGet) wx.Action {
	return wx.NewPostAction(urls.CorpMsgAuditGroupChatGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
