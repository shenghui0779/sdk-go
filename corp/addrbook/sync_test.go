package addrbook

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/chenghonour/gochat/corp"
	"github.com/chenghonour/gochat/mock"
	"github.com/chenghonour/gochat/wx"
)

func TestBatchSyncUser(t *testing.T) {
	body := []byte(`{"media_id":"xxxxxx","to_invite":true,"callback":{"url":"xxx","token":"xxx","encodingaeskey":"xxx"}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "jobid": "xxxxx"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/batch/syncuser?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	callback := &BatchCallback{
		URL:            "xxx",
		Token:          "xxx",
		EncodingAESKey: "xxx",
	}

	result := new(ResultBatch)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", BatchSyncUser("xxxxxx", true, callback, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultBatch{
		JobID: "xxxxx",
	}, result)
}

func TestBatchReplaceUser(t *testing.T) {
	body := []byte(`{"media_id":"xxxxxx","to_invite":true,"callback":{"url":"xxx","token":"xxx","encodingaeskey":"xxx"}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "jobid": "xxxxx"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/batch/replaceuser?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	callback := &BatchCallback{
		URL:            "xxx",
		Token:          "xxx",
		EncodingAESKey: "xxx",
	}

	result := new(ResultBatch)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", BatchReplaceUser("xxxxxx", true, callback, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultBatch{
		JobID: "xxxxx",
	}, result)
}

func TestBatchReplaceParty(t *testing.T) {
	body := []byte(`{"media_id":"xxxxxx","callback":{"url":"xxx","token":"xxx","encodingaeskey":"xxx"}}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "jobid": "xxxxx"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/batch/replaceparty?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultBatch)

	callback := &BatchCallback{
		URL:            "xxx",
		Token:          "xxx",
		EncodingAESKey: "xxx",
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", BatchReplaceParty("xxxxxx", callback, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultBatch{
		JobID: "xxxxx",
	}, result)
}

func TestGetBatchResult(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "status": 1,
    "type": "replace_user",
    "total": 3,
    "percentage": 33,
    "result": [
        {
            "userid": "lisi",
            "errcode": 0,
            "errmsg": "ok"
        },
        {
            "action": 1,
            "partyid": 1,
            "errcode": 0,
            "errmsg": "ok"
        }
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/batch/getresult?access_token=ACCESS_TOKEN&jobid=JOBID", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultBatchResult)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetBatchResult("JOBID", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultBatchResult{
		Status:     1,
		Type:       "replace_user",
		Total:      3,
		Percentage: 33,
		Result: []*BatchResult{
			{
				UserID:  "lisi",
				ErrCode: 0,
				ErrMsg:  "ok",
			},
			{
				Action:  1,
				PartyID: 1,
				ErrCode: 0,
				ErrMsg:  "ok",
			},
		},
	}, result)
}
