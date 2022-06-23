package msgaudit

import (
	"encoding/json"

	"github.com/chenghonour/gochat/urls"
	"github.com/chenghonour/gochat/wx"
)

type ParamsPermitUserList struct {
	Type int `json:"type,omitempty"`
}

type ResultPermitUserList struct {
	IDs []string `json:"ids"`
}

// ListPermitUser 获取会话内容存档开启成员列表
func ListPermitUser(listType int, result *ResultPermitUserList) wx.Action {
	params := &ParamsPermitUserList{
		Type: listType,
	}

	return wx.NewPostAction(urls.CorpMsgAuditGetPermitUserList,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsSingleAgreeCheck struct {
	Info []*ParamsSingleAgree `json:"info"`
}

type ParamsSingleAgree struct {
	UserID         string `json:"userid"`
	ExternalOpenID string `json:"exteranalopenid"`
}

type ResultSingleAgreeCheck struct {
	AgreeInfo []*SingleAgreeInfo `json:"agreeinfo"`
}

type SingleAgreeInfo struct {
	StatusChangeTime int64  `json:"status_change_time"`
	UserID           string `json:"userid"`
	ExternalOpenID   string `json:"exteranalopenid"`
	AgreeStatus      string `json:"agree_status"`
}

// CheckSingleAgree 获取会话同意情况（单聊）
func CheckSingleAgree(agrees []*ParamsSingleAgree, result *ResultSingleAgreeCheck) wx.Action {
	params := &ParamsSingleAgreeCheck{
		Info: agrees,
	}

	return wx.NewPostAction(urls.CorpMsgAuditCheckSingleAgree,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
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
	AgreeInfo []*RoomAgreeInfo `json:"agreeinfo"`
}

type RoomAgreeInfo struct {
	StatusChangeTime int64  `json:"status_change_time"`
	ExternalOpenID   string `json:"exteranalopenid"`
	AgreeStatus      string `json:"agree_status"`
}

// CheckRoomAgree 获取会话同意情况（群聊）
func CheckRoomAgree(roomID string, result *ResultRoomAgreeCheck) wx.Action {
	params := &ParamsRoomAgreeCheck{
		RoomID: roomID,
	}

	return wx.NewPostAction(urls.CorpMsgAuditCheckRoomAgree,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
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

type ParamsGroupChat struct {
	RoomID string `json:"roomid"`
}

type ResultGroupChat struct {
	RoomName       string         `json:"roomname"`
	Creator        string         `json:"creator"`
	RoomCreateTime int64          `json:"room_create_time"`
	Notice         string         `json:"notice"`
	Members        []*GroupMember `json:"members"`
}

// GetGroupChat 获取会话内容存档内部群信息
func GetGroupChat(roomID string, result *ResultGroupChat) wx.Action {
	params := &ParamsGroupChat{
		RoomID: roomID,
	}

	return wx.NewPostAction(urls.CorpMsgAuditGroupChatGet,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
