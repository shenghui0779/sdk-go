package mp

import (
	"encoding/json"
	"io/ioutil"
	"net/url"

	"github.com/shenghui0779/gochat/public"
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
	Results []CropPosition `json:"results"`
	ImgSize ImageSize      `json:"img_size"`
}

// AICrop 图片智能裁切
func AICrop(filename string, dest *AICropResult) public.Action {
	return public.NewOpenUploadAPI(AICropURL, url.Values{}, public.NewUploadBody("img", filename, func() ([]byte, error) {
		return ioutil.ReadFile(filename)
	}), func(resp []byte) error {
		return json.Unmarshal(resp, dest)
	})
}

// AICropByURL 图片智能裁切
func AICropByURL(imgURL string, dest *AICropResult) public.Action {
	query := url.Values{}

	query.Set("img_url", imgURL)

	return public.NewOpenPostAPI(AICropURL, query, nil, func(resp []byte) error {
		return json.Unmarshal(resp, dest)
	})
}

// QRCodeScanData 二维码扫描数据
type QRCodeScanData struct {
	TypeName string        `json:"type_name"`
	Data     string        `json:"data"`
	Pos      ImagePosition `json:"pos"`
}

// QRCodeScanResult 二维码扫描结果
type QRCodeScanResult struct {
	CodeResults []QRCodeScanData `json:"code_results"`
	ImgSize     ImageSize        `json:"img_size"`
}

// ScanQRCode 条码/二维码识别
func ScanQRCode(filename string, dest *QRCodeScanResult) public.Action {
	return public.NewOpenUploadAPI(ScanQRCodeURL, url.Values{}, public.NewUploadBody("img", filename, func() ([]byte, error) {
		return ioutil.ReadFile(filename)
	}), func(resp []byte) error {
		return json.Unmarshal(resp, dest)
	})
}

// ScanQRCodeByURL 条码/二维码识别
func ScanQRCodeByURL(imgURL string, dest *QRCodeScanResult) public.Action {
	query := url.Values{}

	query.Set("img_url", imgURL)

	return public.NewOpenPostAPI(ScanQRCodeURL, query, nil, func(resp []byte) error {
		return json.Unmarshal(resp, dest)
	})
}

// SuperreSolutionResult 图片高清化结果
type SuperreSolutionResult struct {
	MediaID string `json:"media_id"`
}

// SuperreSolution 图片高清化
func SuperreSolution(filename string, dest *SuperreSolutionResult) public.Action {
	return public.NewOpenUploadAPI(SuperreSolutionURL, url.Values{}, public.NewUploadBody("img", filename, func() ([]byte, error) {
		return ioutil.ReadFile(filename)
	}), func(resp []byte) error {
		dest.MediaID = gjson.GetBytes(resp, "media_id").String()

		return nil
	})
}

// SuperreSolutionByURL 图片高清化
func SuperreSolutionByURL(imgURL string, dest *SuperreSolutionResult) public.Action {
	query := url.Values{}

	query.Set("img_url", imgURL)

	return public.NewOpenPostAPI(SuperreSolutionURL, query, nil, func(resp []byte) error {
		dest.MediaID = gjson.GetBytes(resp, "media_id").String()

		return nil
	})
}
