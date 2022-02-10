package tools

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type LivingActivity struct {
	Description string   `json:"description,omitempty"`
	ImageList   []string `json:"image_list,omitempty"`
}

type ParamsLivingCreate struct {
	AnchorUserID         string          `json:"anchor_userid"`
	Theme                string          `json:"theme"`
	LivingStart          int64           `json:"living_start"`
	LivingDuration       int             `json:"living_duration"`
	Description          string          `json:"description,omitempty"`
	Type                 int             `json:"type,omitempty"`
	AgentID              int64           `json:"agentid,omitempty"`
	RemindTime           int             `json:"remind_time,omitempty"`
	ActivityCoverMediaID string          `json:"activity_cover_mediaid,omitempty"`
	ActivityShareMediaID string          `json:"activity_share_mediaid,omitempty"`
	ActivityDetail       *LivingActivity `json:"activity_detail,omitempty"`
}

type ResultLivingCreate struct {
	LivingID string `json:"livingid"`
}

func CreateLiving(params *ParamsLivingCreate, result *ResultLivingCreate) wx.Action {
	return wx.NewPostAction(urls.CorpToolsLivingCreate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsLivingModify struct {
	LivingID       string `json:"livingid"`
	Theme          string `json:"theme,omitempty"`
	LivingStart    int64  `json:"living_start,omitempty"`
	LivingDuration int    `json:"living_duration,omitempty"`
	Description    string `json:"description,omitempty"`
	Type           int    `json:"type,omitempty"`
	RemindTime     int    `json:"remind_time,omitempty"`
}

func ModifyLiving(params *ParamsLivingModify) wx.Action {
	return wx.NewPostAction(urls.CorpToolsLivingModify,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsLivingCancel struct {
	LivingID string `json:"livingid"`
}

func CancelLiving(livingID string) wx.Action {
	params := &ParamsLivingCancel{
		LivingID: livingID,
	}

	return wx.NewPostAction(urls.CorpToolsLivingCancel,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsLivingReplayDataDelete struct {
	LivingID string `json:"livingid"`
}

func DeleteLivingReplayData(livingID string) wx.Action {
	params := &ParamsLivingReplayDataDelete{
		LivingID: livingID,
	}

	return wx.NewPostAction(urls.CorpToolsLivingDeleteReplayData,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsLivingCode struct {
	LivingID string `json:"livingid"`
	OpenID   string `json:"openid"`
}

type ResultLivingCode struct {
	LivingCode string `json:"living_code"`
}

func GetLivingCode(livingID, openID string, result *ResultLivingCode) wx.Action {
	params := &ParamsLivingCode{
		LivingID: livingID,
		OpenID:   openID,
	}

	return wx.NewPostAction(urls.CorpToolsLivingGetCode,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsUserAllLivingID struct {
	UserID string `json:"userid"`
	Cursor string `json:"cursor,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}

type ResultUserAllLivingID struct {
	NextCursor   string   `json:"next_cursor"`
	LivingIDList []string `json:"livingid_list"`
}

func GetUserAllLivingID(userID, cursor string, limit int, result *ResultUserAllLivingID) wx.Action {
	params := &ParamsUserAllLivingID{
		UserID: userID,
		Cursor: cursor,
		Limit:  limit,
	}

	return wx.NewPostAction(urls.CorpToolsLivingGetUserAllLivingID,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type LivingInfo struct {
	Theme                 string `json:"theme"`
	LivingStart           int64  `json:"living_start"`
	LivingDuration        int    `json:"living_duration"`
	Status                int    `json:"status"`
	ReserveStart          int64  `json:"reserve_start"`
	ReserveLivingDuration int    `json:"reserve_living_duration"`
	Description           string `json:"description"`
	AnchorUserID          string `json:"anchor_userid"`
	MainDepartment        int64  `json:"main_department"`
	ViewerNum             int    `json:"viewer_num"`
	CommentNum            int    `json:"comment_num"`
	MicNum                int    `json:"mic_num"`
	OpenReplay            int    `json:"open_replay"`
	ReplayStatus          int    `json:"replay_status"`
	Type                  int    `json:"type"`
	PushStreamURL         string `json:"push_stream_url"`
	OnlineCount           int    `json:"online_count"`
	SubscribeCount        int    `json:"subscribe_count"`
}

type ResultLivingInfo struct {
	LivingInfo *LivingInfo `json:"living_info"`
}

func GetLivingInfo(livingID string, result *ResultLivingInfo) wx.Action {
	return wx.NewGetAction(urls.CorpToolsLivingGetInfo,
		wx.WithQuery("livingid", livingID),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type LivingStatInfo struct {
	Users         []*LivingUser         `json:"users"`
	ExternalUsers []*LivingExternalUser `json:"external_users"`
}

type LivingUser struct {
	UserID    string `json:"userid"`
	WatchTime int64  `json:"watch_time"`
	IsComment int    `json:"is_comment"`
	IsMic     int    `json:"is_mic"`
}

type LivingExternalUser struct {
	ExternalUserID string `json:"external_userid"`
	Type           int    `json:"type"`
	Name           string `json:"name"`
	WatchTime      int64  `json:"watch_time"`
	IsComment      int    `json:"is_comment"`
	IsMic          int    `json:"is_mic"`
}

type ParamsLivingWatchStat struct {
	LivingID string `json:"livingid"`
	NextKey  string `json:"next_key"`
}

type ResultLivingWatchStat struct {
	Ending   int             `json:"ending"`
	NextKey  string          `json:"next_key"`
	StatInfo *LivingStatInfo `json:"stat_info"`
}

func GetLivingWatchStat(livingID, nextKey string, result *ResultLivingWatchStat) wx.Action {
	params := &ParamsLivingWatchStat{
		LivingID: livingID,
		NextKey:  nextKey,
	}

	return wx.NewPostAction(urls.CorpToolsLivingGetWatchStat,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsLivingShareInfo struct {
	WWShareCode string `json:"ww_share_code"`
}

type ResultLivingShareInfo struct {
	LivingID              string `json:"livingid"`
	ViewerUserID          string `json:"viewer_userid"`
	ViewerExternalUserID  string `json:"viewer_external_userid"`
	InvitorUserID         string `json:"Invitor_userid"`
	InvitorExternalUserID string `json:"Invitor_external_userid"`
}

func GetLivingShareInfo(wwshareCode string, result *ResultLivingShareInfo) wx.Action {
	params := &ParamsLivingShareInfo{
		WWShareCode: wwshareCode,
	}

	return wx.NewPostAction(urls.CorpToolsLivingGetShareInfo,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
