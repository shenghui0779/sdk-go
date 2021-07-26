package corp

import (
	"encoding/json"

	"github.com/tidwall/gjson"

	"github.com/shenghui0779/gochat/wx"
)

type BatchSyncType string

const (
	SyncUser     BatchSyncType = "sync_user"
	ReplaceUser  BatchSyncType = "replace_user"
	ReplaceParty BatchSyncType = "replace_party"
)

type ParamsBatchSync struct {
	MediaID  string        `json:"media_id"`
	ToInvite bool          `json:"to_invite"`
	Callback *SyncCallback `json:"callback,omitempty"`
}

type SyncCallback struct {
	URL            string `json:"url,omitempty"`
	Token          string `json:"token,omitempty"`
	EncodingAESKey string `json:"encodingaeskey,omitempty"`
}

type BatchSyncJob struct {
	JobID string `json:"jobid"`
}

func BatchSyncUser(dest *BatchSyncJob, params *ParamsBatchSync) wx.Action {
	return wx.NewPostAction(BatchSyncUserURL,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			dest.JobID = gjson.GetBytes(resp, "jobid").String()

			return nil
		}),
	)
}

func BatchReplaceUser(dest *BatchSyncJob, params *ParamsBatchSync) wx.Action {
	return wx.NewPostAction(BatchReplaceUserURL,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			dest.JobID = gjson.GetBytes(resp, "jobid").String()

			return nil
		}),
	)
}

func BatchReplaceParty(dest *BatchSyncJob, params *ParamsBatchSync) wx.Action {
	return wx.NewPostAction(BatchReplacePartyURL,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			dest.JobID = gjson.GetBytes(resp, "jobid").String()

			return nil
		}),
	)
}

type BatchSyncResult struct {
	Status     int           `json:"status"`
	Type       BatchSyncType `json:"type"`
	Total      int           `json:"total"`
	Percentage int           `json:"percentage"`
}
