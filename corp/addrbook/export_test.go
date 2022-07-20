package addrbook

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/corp"
	"github.com/shenghui0779/gochat/mock"
	"github.com/shenghui0779/gochat/wx"
)

func TestExportSimpleUser(t *testing.T) {
	body := []byte(`{"encoding_aeskey":"IJUiXNpvGbODwKEBSEsAeOAPAhkqHqNCF6g19t9wfg2","block_size":1000000}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "jobid": "jobid_xxxxxxxxx"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/export/simple_user?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultExport)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ExportSimpleUser("IJUiXNpvGbODwKEBSEsAeOAPAhkqHqNCF6g19t9wfg2", 1000000, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultExport{
		JobID: "jobid_xxxxxxxxx",
	}, result)
}

func TestExportUser(t *testing.T) {
	body := []byte(`{"encoding_aeskey":"IJUiXNpvGbODwKEBSEsAeOAPAhkqHqNCF6g19t9wfg2","block_size":1000000}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "jobid": "jobid_xxxxxxxxx"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/export/user?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultExport)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ExportUser("IJUiXNpvGbODwKEBSEsAeOAPAhkqHqNCF6g19t9wfg2", 1000000, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultExport{
		JobID: "jobid_xxxxxxxxx",
	}, result)
}

func TestExportDepartment(t *testing.T) {
	body := []byte(`{"encoding_aeskey":"IJUiXNpvGbODwKEBSEsAeOAPAhkqHqNCF6g19t9wfg2","block_size":1000000}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "jobid": "jobid_xxxxxxxxx"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/export/department?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultExport)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ExportDepartment("IJUiXNpvGbODwKEBSEsAeOAPAhkqHqNCF6g19t9wfg2", 1000000, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultExport{
		JobID: "jobid_xxxxxxxxx",
	}, result)
}

func TestExportTagUser(t *testing.T) {
	body := []byte(`{"tagid":1,"encoding_aeskey":"IJUiXNpvGbODwKEBSEsAeOAPAhkqHqNCF6g19t9wfg2","block_size":1000000}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "jobid": "jobid_xxxxxxxxx"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/export/taguser?access_token=ACCESS_TOKEN", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultExport)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", ExportTagUser(1, "IJUiXNpvGbODwKEBSEsAeOAPAhkqHqNCF6g19t9wfg2", 1000000, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultExport{
		JobID: "jobid_xxxxxxxxx",
	}, result)
}

func TestGetExportResult(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
    "errcode": 0,
    "errmsg": "ok",
    "status": 2,
    "data_list": [
        {
            "url": "https://xxxxx",
            "size": 10240,
            "md5": "xxxxxxxxx"
        },
        {
            "url": "https://xxxxx",
            "size": 20480,
            "md5": "xxxxxxxx"
        }
    ]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/export/get_result?access_token=ACCESS_TOKEN&jobid=jobid_xxxxxxxxxxxxxxx", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultExportRet)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetExportResult("jobid_xxxxxxxxxxxxxxx", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultExportRet{
		Status: 2,
		DataList: []*ExportData{
			{
				URL:  "https://xxxxx",
				Size: 10240,
				MD5:  "xxxxxxxxx",
			},
			{
				URL:  "https://xxxxx",
				Size: 20480,
				MD5:  "xxxxxxxx",
			},
		},
	}, result)
}
