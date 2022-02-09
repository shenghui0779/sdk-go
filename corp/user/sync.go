package user

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
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

type SyncCallback struct {
	URL            string `json:"url,omitempty"`
	Token          string `json:"token,omitempty"`
	EncodingAESKey string `json:"encodingaeskey,omitempty"`
}

type ParamsUserBatchSync struct {
	MediaID  string        `json:"media_id"`
	ToInvite *bool         `json:"to_invite,omitempty"`
	Callback *SyncCallback `json:"callback,omitempty"`
}

func BatchSyncUser(params *ParamsUserBatchSync, result *ResultBatch) wx.Action {
	return wx.NewPostAction(urls.CorpUserBatchSyncUser,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsUserBatchReplace struct {
	MediaID  string        `json:"media_id"`
	ToInvite *bool         `json:"to_invite,omitempty"`
	Callback *SyncCallback `json:"callback,omitempty"`
}

func BatchReplaceUser(params *ParamsUserBatchReplace, result *ResultBatch) wx.Action {
	return wx.NewPostAction(urls.CorpUserBatchReplaceUser,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsPartyBatchReplace struct {
	MediaID  string        `json:"media_id"`
	Callback *SyncCallback `json:"callback,omitempty"`
}

func BatchReplaceParty(params *ParamsPartyBatchReplace, result *ResultBatch) wx.Action {
	return wx.NewPostAction(urls.CorpUserBatchReplaceParty,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
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

func GetBatchResult(jobID string, result *ResultBatchResult) wx.Action {
	return wx.NewGetAction(urls.CorpUserGetBatchResult,
		wx.WithQuery("jobid", jobID),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
