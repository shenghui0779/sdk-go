package addrbook

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

type ParamsExport struct {
	TagID          int64  `json:"tagid,omitempty"`
	EncodingAESKey string `json:"encoding_aeskey"`
	BlockSize      int64  `json:"block_size"`
}

type ResultExport struct {
	JobID string `json:"jobid"`
}

// ExportSimpleUser 导出成员
func ExportSimpleUser(encodingAESKey string, blockSize int64, result *ResultExport) wx.Action {
	params := &ParamsExport{
		EncodingAESKey: encodingAESKey,
		BlockSize:      blockSize,
	}

	return wx.NewPostAction(urls.CorpUserExportSimpleUser,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// ExportUser 导出成员详情
func ExportUser(encodingAESKey string, blockSize int64, result *ResultExport) wx.Action {
	params := &ParamsExport{
		EncodingAESKey: encodingAESKey,
		BlockSize:      blockSize,
	}

	return wx.NewPostAction(urls.CorpUserExportUser,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// ExportDepartment 导出部门
func ExportDepartment(encodingAESKey string, blockSize int64, result *ResultExport) wx.Action {
	params := &ParamsExport{
		EncodingAESKey: encodingAESKey,
		BlockSize:      blockSize,
	}

	return wx.NewPostAction(urls.CorpUserExportDepartment,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// ExportTagUser 导出标签成员
func ExportTagUser(tagID int64, encodingAESKey string, blockSize int64, result *ResultExport) wx.Action {
	params := &ParamsExport{
		TagID:          tagID,
		EncodingAESKey: encodingAESKey,
		BlockSize:      blockSize,
	}

	return wx.NewPostAction(urls.CorpUserExportTagUser,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

type ResultExportRet struct {
	Status   int           `json:"status"`
	DataList []*ExportData `json:"data_list"`
}

type ExportData struct {
	URL  string `json:"url"`
	Size int64  `json:"size"`
	MD5  string `json:"md5"`
}

// GetExportResult 获取导出结果
func GetExportResult(jobID string, result *ResultExportRet) wx.Action {
	return wx.NewGetAction(urls.CorpUserGetExportResult,
		wx.WithQuery("jobid", jobID),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}
