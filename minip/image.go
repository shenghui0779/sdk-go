package minip

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/tidwall/gjson"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

// ImageSize 图片尺寸
type ImageSize struct {
	W int `json:"w"`
	H int `json:"h"`
}

// Position 位置信息
type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// ImagePosition 图片位置
type ImagePosition struct {
	LeftTop     Position `json:"left_top"`
	RightTop    Position `json:"right_top"`
	RightBottom Position `json:"right_bottom"`
	LeftBottom  Position `json:"left_bottom"`
}

// CropPosition 裁切位置
type CropPosition struct {
	CropLeft   int `json:"crop_left"`
	CropTop    int `json:"crop_top"`
	CropRight  int `json:"crop_right"`
	CropBottom int `json:"crop_bottom"`
}

// AICropResult 图片裁切结果
type AICropResult struct {
	Results []*CropPosition `json:"results"`
	ImgSize ImageSize       `json:"img_size"`
}

// AICrop 图片智能裁切
func AICrop(dest *AICropResult, path string) wx.Action {
	_, filename := filepath.Split(path)

	return wx.NewUploadAction(urls.MinipAICrop,
		wx.WithUploadField(&wx.UploadField{
			FileField: "img",
			Filename:  filename,
		}),
		wx.WithBody(func() ([]byte, error) {
			path, err := filepath.Abs(filepath.Clean(path))

			if err != nil {
				return nil, err
			}

			return ioutil.ReadFile(path)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// AICropByURL 图片智能裁切
func AICropByURL(dest *AICropResult, imgURL string) wx.Action {
	return wx.NewPostAction(urls.MinipAICrop,
		wx.WithQuery("img_url", imgURL),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// QRCodeScanData 二维码扫描数据
type QRCodeScanData struct {
	TypeName string        `json:"type_name"`
	Data     string        `json:"data"`
	Pos      ImagePosition `json:"pos"`
}

// QRCodeScanResult 二维码扫描结果
type QRCodeScanResult struct {
	CodeResults []*QRCodeScanData `json:"code_results"`
	ImgSize     ImageSize         `json:"img_size"`
}

// ScanQRCode 条码/二维码识别
func ScanQRCode(dest *QRCodeScanResult, path string) wx.Action {
	_, filename := filepath.Split(path)

	return wx.NewUploadAction(urls.MinipScanQRCode,
		wx.WithUploadField(&wx.UploadField{
			FileField: "img",
			Filename:  filename,
		}),
		wx.WithBody(func() ([]byte, error) {
			path, err := filepath.Abs(filepath.Clean(path))

			if err != nil {
				return nil, err
			}

			return ioutil.ReadFile(path)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// ScanQRCodeByURL 条码/二维码识别
func ScanQRCodeByURL(dest *QRCodeScanResult, imgURL string) wx.Action {
	return wx.NewPostAction(urls.MinipScanQRCode,
		wx.WithQuery("img_url", imgURL),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// SuperreSolutionResult 图片高清化结果
type SuperreSolutionResult struct {
	MediaID string `json:"media_id"`
}

// SuperreSolution 图片高清化
func SuperreSolution(dest *SuperreSolutionResult, path string) wx.Action {
	_, filename := filepath.Split(path)

	return wx.NewUploadAction(urls.MinipSuperreSolution,
		wx.WithUploadField(&wx.UploadField{
			FileField: "img",
			Filename:  filename,
		}),
		wx.WithBody(func() ([]byte, error) {
			path, err := filepath.Abs(filepath.Clean(path))

			if err != nil {
				return nil, err
			}

			return ioutil.ReadFile(path)
		}),
		wx.WithDecode(func(resp []byte) error {
			dest.MediaID = gjson.GetBytes(resp, "media_id").String()

			return nil
		}),
	)
}

// SuperreSolutionByURL 图片高清化
func SuperreSolutionByURL(dest *SuperreSolutionResult, imgURL string) wx.Action {
	return wx.NewPostAction(urls.MinipSuperreSolution,
		wx.WithQuery("img_url", imgURL),
		wx.WithDecode(func(resp []byte) error {
			dest.MediaID = gjson.GetBytes(resp, "media_id").String()

			return nil
		}),
	)
}