package addrbook

import (
	"encoding/json"

	"github.com/chenghonour/gochat/urls"
	"github.com/chenghonour/gochat/wx"
)

type BatchType string

const (
	SyncUser     BatchType = "sync_user"
	ReplaceUser  BatchType = "replace_user"
	ReplaceParty BatchType = "replace_party"
)

type ResultBatch struct {
	JobID string `json:"jobid"`
}

type BatchCallback struct {
	URL            string `json:"url,omitempty"`
	Token          string `json:"token,omitempty"`
	EncodingAESKey string `json:"encodingaeskey,omitempty"`
}

type ParamsUserBatchSync struct {
	MediaID  string         `json:"media_id"`
	ToInvite bool           `json:"to_invite"`
	Callback *BatchCallback `json:"callback,omitempty"`
}

// BatchSyncUser 增量更新成员
func BatchSyncUser(mediaID string, toInvite bool, callback *BatchCallback, result *ResultBatch) wx.Action {
	params := &ParamsUserBatchSync{
		MediaID:  mediaID,
		ToInvite: toInvite,
		Callback: callback,
	}

	return wx.NewPostAction(urls.CorpUserBatchSyncUser,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsUserBatchReplace struct {
	MediaID  string         `json:"media_id"`
	ToInvite bool           `json:"to_invite"`
	Callback *BatchCallback `json:"callback,omitempty"`
}

// BatchReplaceUser 全量覆盖成员
func BatchReplaceUser(mediaID string, toInvite bool, callback *BatchCallback, result *ResultBatch) wx.Action {
	params := &ParamsUserBatchReplace{
		MediaID:  mediaID,
		ToInvite: toInvite,
		Callback: callback,
	}

	return wx.NewPostAction(urls.CorpUserBatchReplaceUser,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsPartyBatchReplace struct {
	MediaID  string         `json:"media_id"`
	Callback *BatchCallback `json:"callback,omitempty"`
}

// BatchReplaceParty 全量覆盖部门
func BatchReplaceParty(mediaID string, callback *BatchCallback, result *ResultBatch) wx.Action {
	params := &ParamsPartyBatchReplace{
		MediaID:  mediaID,
		Callback: callback,
	}

	return wx.NewPostAction(urls.CorpUserBatchReplaceParty,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ResultBatchResult struct {
	Status     int            `json:"status"`
	Type       BatchType      `json:"type"`
	Total      int            `json:"total"`
	Percentage int            `json:"percentage"`
	Result     []*BatchResult `json:"result"`
}

type BatchResult struct {
	UserID  string `json:"userid"`
	Action  int    `json:"action"`
	PartyID int    `json:"partyid"`
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// GetBatchResult 获取异步任务结果
func GetBatchResult(jobID string, result *ResultBatchResult) wx.Action {
	return wx.NewGetAction(urls.CorpUserGetBatchResult,
		wx.WithQuery("jobid", jobID),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
