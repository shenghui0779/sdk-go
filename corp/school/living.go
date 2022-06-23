package school

import (
	"encoding/json"

	"github.com/chenghonour/gochat/urls"
	"github.com/chenghonour/gochat/wx"
)

type ParamsUserAllLivingID struct {
	UserID string `json:"userid"`
	Cursor string `json:"cursor"`
	Limit  int    `json:"limit"`
}

type ResultUserAllLivingID struct {
	NextCursor   string   `json:"next_cursor"`
	LivingIDList []string `json:"livingid_list"`
}

// GetUserAllLivingID 获取老师直播ID列表
func GetUserAllLivingID(userID, cursor string, limit int, result *ResultUserAllLivingID) wx.Action {
	params := &ParamsUserAllLivingID{
		UserID: userID,
		Cursor: cursor,
		Limit:  limit,
	}

	return wx.NewPostAction(urls.CorpSchoolGetUserAllLivingID,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
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

type ResultLivingInfo struct {
	LivingInfo *LivingInfo `json:"living_info"`
}

// GetLivingInfo 获取直播详情
func GetLivingInfo(livingID string, result *ResultLivingInfo) wx.Action {
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

type ParamsLivingWatchStat struct {
	LivingID string `json:"livingid"`
	NextKey  string `json:"next_key,omitempty"`
}

type ResultLivingWatchStat struct {
	Ending     int                  `json:"ending"`
	NextKey    string               `json:"next_key"`
	StatInfoes *LivingWatchStatInfo `json:"stat_infoes"`
}

// GetLivingWatchStat 获取观看直播统计
func GetLivingWatchStat(livingID, nextKey string, result *ResultLivingWatchStat) wx.Action {
	params := &ParamsLivingWatchStat{
		LivingID: livingID,
		NextKey:  nextKey,
	}

	return wx.NewPostAction(urls.CorpSchoolGetLivingWatchStat,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
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

type ParamsLivingUnwatchStat struct {
	LivingID string `json:"livingid"`
	NextKey  string `json:"next_key,omitempty"`
}

type ResultLivingUnwatchStat struct {
	Ending   int                    `json:"ending"`
	NextKey  string                 `json:"next_key"`
	StatInfo *LivingUnwatchStatInfo `json:"stat_info"`
}

// GetLivingUnwatchStat 获取未观看直播统计
func GetLivingUnwatchStat(livingID, nextKey string, result *ResultLivingUnwatchStat) wx.Action {
	params := &ParamsLivingUnwatchStat{
		LivingID: livingID,
		NextKey:  nextKey,
	}

	return wx.NewPostAction(urls.CorpSchoolGetLivingUnwatchStat,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsLivingReplayDataDelete struct {
	LivingID string `json:"livingid"`
}

// DeleteLivingReplayData 删除直播回放
func DeleteLivingReplayData(livingID string) wx.Action {
	params := &ParamsLivingReplayDataDelete{
		LivingID: livingID,
	}

	return wx.NewPostAction(urls.CorpSchoolDeleteLivingReplayData,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}
