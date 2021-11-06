package school

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsUserAllLivingIDGet struct {
	UserID string `json:"userid"`
	Cursor string `json:"cursor"`
	Limit  int    `json:"limit"`
}

type ResultUserAllLivingIDGet struct {
	NextCursor   string   `json:"next_cursor"`
	LivingIDList []string `json:"livingid_list"`
}

func GetUserAllLivingID(params *ParamsUserAllLivingIDGet, result *ResultUserAllLivingIDGet) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolGetUserAllLivingID,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type LivingInfo struct {
	Theme          string       `json:"theme"`
	LivingStart    int64        `json:"living_start"`
	LivingDuration int          `json:"living_duration"`
	AnchorUserID   string       `json:"anchor_userid"`
	LivingRange    *LivingRange `json:"living_range"`
	ViewerNum      int          `json:"viewer_num"`
	CommentNum     int          `json:"comment_num"`
	OpenReplay     int          `json:"open_replay"`
	PushStreamURL  string       `json:"push_stream_url"`
}

type LivingRange struct {
	PartyIDs   []int64  `json:"partyids"`
	GroupNames []string `json:"group_names"`
}

type ResultLivingInfoGet struct {
	LivingInfo *LivingInfo `json:"living_info"`
}

func GetLivingInfo(livingID string, result *ResultLivingInfoGet) wx.Action {
	return wx.NewGetAction(urls.CorpSchoolGetLivingInfo,
		wx.WithQuery("livingid", livingID),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type LivingWatchStudent struct {
	StudentUserID string  `json:"student_userid"`
	ParentUserID  string  `json:"parent_userid"`
	PartyIDs      []int64 `json:"partyids"`
	WatchTime     int     `json:"watch_time"`
	EnterTime     int64   `json:"enter_time"`
	LeaveTime     int64   `json:"leave_time"`
	IsComment     int     `json:"is_comment"`
}

type LivingVisitor struct {
	Nickname  string `json:"nickname"`
	WatchTime int    `json:"watch_time"`
	EnterTime int64  `json:"enter_time"`
	LeaveTime int64  `json:"leave_time"`
	IsComment int    `json:"is_comment"`
}

type LivingWatchStatInfo struct {
	Students []*LivingWatchStudent `json:"students"`
	Visitors []*LivingVisitor      `json:"visitors"`
}

type ParamsLivingWatchStatGet struct {
	LivingID string `json:"livingid"`
	NextKey  string `json:"next_key"`
}

type ResultLivingWatchStatGet struct {
	Ending     int                  `json:"ending"`
	NextKey    string               `json:"next_key"`
	StatInfoes *LivingWatchStatInfo `json:"stat_infoes"`
}

func GetLivingWatchStat(params *ParamsLivingWatchStatGet, result *ResultLivingWatchStatGet) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolGetLivingWatchStat,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type LivingUnwatchStudent struct {
	StudentUserID string  `json:"student_userid"`
	ParentUserID  string  `json:"parent_userid"`
	PartyIDs      []int64 `json:"partyids"`
}

type LivingUnwatchStatInfo struct {
	Students []*LivingUnwatchStudent `json:"students"`
}

type ParamsLivingUnwatchStatGet struct {
	LivingID string `json:"livingid"`
	NextKey  string `json:"next_key"`
}

type ResultLivingUnwatchStatGet struct {
	Ending   int                    `json:"ending"`
	NextKey  string                 `json:"next_key"`
	StatInfo *LivingUnwatchStatInfo `json:"stat_info"`
}

func GetLivingUnwatchStat(params *ParamsLivingUnwatchStatGet, result *ResultLivingUnwatchStatGet) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolGetLivingUnwatchStat,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsLivingReplayDataDelete struct {
	LivingID string `json:"livingid"`
}

func DeleteLivingReplayData(params *ParamsLivingReplayDataDelete) wx.Action {
	return wx.NewPostAction(urls.CorpSchoolDeleteLivingReplayData,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}
