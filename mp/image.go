package mp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/shenghui0779/gochat/helpers"
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
	CropRignt  int `json:"crop_rignt"`
	CropBottom int `json:"crop_bottom"`
}

// AICropResult 图片裁切结果
type AICropResult struct {
	Result  []CropPosition `json:"result"`
	ImgSize ImageSize      `json:"img_size"`
}

// AICrop 图片智能裁切
func AICrop(filename string, dest *AICropResult) Action {
	return &WechatAPI{
		body: helpers.NewUploadBody("img", filename, func() ([]byte, error) {
			return ioutil.ReadFile(filename)
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("UPLOAD|%s?access_token=%s", AICropURL, accessToken)
		},
		decode: func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		},
	}
}

// AICropByURL 图片智能裁切
func AICropByURL(imgURL string, dest *AICropResult) Action {
	return &WechatAPI{
		url: func(accessToken string) string {
			return fmt.Sprintf("POST|%s?img_url=%s&access_token=%s", AICropURL, imgURL, accessToken)
		},
		decode: func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		},
	}
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
func ScanQRCode(filename string, dest *QRCodeScanResult) Action {
	return &WechatAPI{
		body: helpers.NewUploadBody("img", filename, func() ([]byte, error) {
			return ioutil.ReadFile(filename)
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("UPLOAD|%s?access_token=%s", ScanQRCodeURL, accessToken)
		},
		decode: func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		},
	}
}

// ScanQRCodeByURL 条码/二维码识别
func ScanQRCodeByURL(imgURL string, dest *QRCodeScanResult) Action {
	return &WechatAPI{
		url: func(accessToken string) string {
			return fmt.Sprintf("POST|%s?img_url=%s&access_token=%s", ScanQRCodeURL, imgURL, accessToken)
		},
		decode: func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		},
	}
}

// SuperreSolutionResult 图片高清化结果
type SuperreSolutionResult struct {
	MediaID string `json:"media_id"`
}

// SuperreSolution 图片高清化
func SuperreSolution(filename string, dest *SuperreSolutionResult) Action {
	return &WechatAPI{
		body: helpers.NewUploadBody("img", filename, func() ([]byte, error) {
			return ioutil.ReadFile(filename)
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("UPLOAD|%s?access_token=%s", SuperreSolutionURL, accessToken)
		},
		decode: func(resp []byte) error {
			dest.MediaID = gjson.GetBytes(resp, "media_id").String()

			return nil
		},
	}
}

// SuperreSolutionByURL 图片高清化
func SuperreSolutionByURL(imgURL string, dest *SuperreSolutionResult) Action {
	return &WechatAPI{
		url: func(accessToken string) string {
			return fmt.Sprintf("POST|%s?img_url=%s&access_token=%s", SuperreSolutionURL, imgURL, accessToken)
		},
		decode: func(resp []byte) error {
			dest.MediaID = gjson.GetBytes(resp, "media_id").String()

			return nil
		},
	}
}
