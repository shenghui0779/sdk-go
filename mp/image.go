package mp

import (
	"encoding/json"
	"net/url"

	"github.com/shenghui0779/gochat/wx"
	"github.com/tidwall/gjson"
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
func AICrop(dest *AICropResult, filename string) wx.Action {
	return wx.NewAction(AICropURL,
		wx.WithMethod(wx.MethodUpload),
		wx.WithUploadForm("img", filename, nil),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// AICropByURL 图片智能裁切
func AICropByURL(dest *AICropResult, imgURL string) wx.Action {
	return wx.NewAction(AICropURL,
		wx.WithMethod(wx.MethodPost),
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
func ScanQRCode(dest *QRCodeScanResult, filename string) wx.Action {
	return wx.NewAction(ScanQRCodeURL,
		wx.WithMethod(wx.MethodUpload),
		wx.WithUploadForm("img", filename, nil),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// ScanQRCodeByURL 条码/二维码识别
func ScanQRCodeByURL(dest *QRCodeScanResult, imgURL string) wx.Action {
	return wx.NewAction(ScanQRCodeURL,
		wx.WithMethod(wx.MethodPost),
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
func SuperreSolution(dest *SuperreSolutionResult, filename string) wx.Action {
	return wx.NewAction(SuperreSolutionURL,
		wx.WithMethod(wx.MethodUpload),
		wx.WithUploadForm("img", filename, nil),
		wx.WithDecode(func(resp []byte) error {
			dest.MediaID = gjson.GetBytes(resp, "media_id").String()

			return nil
		}),
	)
}

// SuperreSolutionByURL 图片高清化
func SuperreSolutionByURL(dest *SuperreSolutionResult, imgURL string) wx.Action {
	query := url.Values{}

	query.Set("img_url", imgURL)

	return wx.NewAction(SuperreSolutionURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithQuery("img_url", imgURL),
		wx.WithDecode(func(resp []byte) error {
			dest.MediaID = gjson.GetBytes(resp, "media_id").String()

			return nil
		}),
	)
}
