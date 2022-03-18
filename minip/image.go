package minip

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

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

// ResultAICrop 图片裁切结果
type ResultAICrop struct {
	Results []*CropPosition `json:"results"`
	ImgSize ImageSize       `json:"img_size"`
}

// AICrop 图像处理 - 图片智能裁切
func AICrop(imgPath string, result *ResultAICrop) wx.Action {
	_, filename := filepath.Split(imgPath)

	return wx.NewPostAction(urls.MinipAICrop,
		wx.WithUpload(func() (wx.UploadForm, error) {
			path, err := filepath.Abs(filepath.Clean(imgPath))

			if err != nil {
				return nil, err
			}

			body, err := ioutil.ReadFile(path)

			if err != nil {
				return nil, err
			}

			return wx.NewUploadForm(
				wx.WithFileField("img", filename, body),
			), nil
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// AICropByURL 图像处理 - 图片智能裁切
func AICropByURL(imgURL string, result *ResultAICrop) wx.Action {
	return wx.NewPostAction(urls.MinipAICrop,
		wx.WithQuery("img_url", imgURL),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// QRCodeScanData 二维码扫描数据
type QRCodeScanData struct {
	TypeName string        `json:"type_name"`
	Data     string        `json:"data"`
	Pos      ImagePosition `json:"pos"`
}

// ResultQRCodeScan 二维码扫描结果
type ResultQRCodeScan struct {
	CodeResults []*QRCodeScanData `json:"code_results"`
	ImgSize     ImageSize         `json:"img_size"`
}

// ScanQRCode 图像处理 - 条码/二维码识别
func ScanQRCode(imgPath string, result *ResultQRCodeScan) wx.Action {
	_, filename := filepath.Split(imgPath)

	return wx.NewPostAction(urls.MinipScanQRCode,
		wx.WithUpload(func() (wx.UploadForm, error) {
			path, err := filepath.Abs(filepath.Clean(imgPath))

			if err != nil {
				return nil, err
			}

			body, err := ioutil.ReadFile(path)

			if err != nil {
				return nil, err
			}

			return wx.NewUploadForm(
				wx.WithFileField("img", filename, body),
			), nil
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// ScanQRCodeByURL 图像处理 - 条码/二维码识别
func ScanQRCodeByURL(imgURL string, result *ResultQRCodeScan) wx.Action {
	return wx.NewPostAction(urls.MinipScanQRCode,
		wx.WithQuery("img_url", imgURL),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// ResultSuperreSolution 图片高清化结果
type ResultSuperreSolution struct {
	MediaID string `json:"media_id"`
}

// SuperreSolution 图像处理 - 图片高清化
func SuperreSolution(imgPath string, result *ResultSuperreSolution) wx.Action {
	_, filename := filepath.Split(imgPath)

	return wx.NewPostAction(urls.MinipSuperreSolution,
		wx.WithUpload(func() (wx.UploadForm, error) {
			path, err := filepath.Abs(filepath.Clean(imgPath))

			if err != nil {
				return nil, err
			}

			body, err := ioutil.ReadFile(path)

			if err != nil {
				return nil, err
			}

			return wx.NewUploadForm(
				wx.WithFileField("img", filename, body),
			), nil
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// SuperreSolutionByURL 图像处理 - 图片高清化
func SuperreSolutionByURL(imgURL string, result *ResultSuperreSolution) wx.Action {
	return wx.NewPostAction(urls.MinipSuperreSolution,
		wx.WithQuery("img_url", imgURL),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
