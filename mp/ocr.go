package mp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/shenghui0779/gochat/helpers"
	"github.com/tidwall/gjson"
)

// OCRMode 识别模式
type OCRMode string

var (
	// OCRPhoto 拍照模式
	OCRPhoto OCRMode = "photo"
	// OCRScan 扫描模式
	OCRScan OCRMode = "scan"
)

// BankCard 银行卡
type BankCard struct {
	ID string `json:"id"`
}

// OCRBankCard 银行卡识别
func OCRBankCard(mode OCRMode, filename string, receiver *BankCard) Action {
	return &WechatAPI{
		body: helpers.NewUploadBody("img", filename, func() ([]byte, error) {
			return ioutil.ReadFile(filename)
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("UPLOAD|%s?type=%s&access_token=%s", OCRBankCardURL, mode, accessToken)
		},
		decode: func(resp []byte) error {
			receiver.ID = gjson.GetBytes(resp, "id").String()

			return nil
		},
	}
}

// OCRBankCardByURL 银行卡识别
func OCRBankCardByURL(mode OCRMode, imgURL string, receiver *BankCard) Action {
	return &WechatAPI{
		url: func(accessToken string) string {
			return fmt.Sprintf("POST|%s?type=%s&img_url=%s&access_token=%s", OCRBankCardURL, mode, imgURL, accessToken)
		},
		decode: func(resp []byte) error {
			receiver.ID = gjson.GetBytes(resp, "id").String()

			return nil
		},
	}
}

// CertPosition 证书位置
type CertPosition struct {
	Pos ImagePosition `json:"pos"`
}

// BusinessLicense 营业执照
type BusinessLicense struct {
	RegNum              string       `json:"reg_num"`
	Serial              string       `json:"serial"`
	LegalRepresentative string       `json:"legal_representative"`
	EnterpriseName      string       `json:"enterprise_name"`
	TypeOfOrganization  string       `json:"type_of_organization"`
	Address             string       `json:"address"`
	TypeOfEnterprise    string       `json:"type_of_enterprise"`
	BusinessScope       string       `json:"business_scope"`
	RegisteredCapital   string       `json:"registered_capital"`
	PaidInCapital       string       `json:"paid_in_capital"`
	ValidPeriod         string       `json:"valid_period"`
	RegisteredDate      string       `json:"registered_date"`
	CertPosition        CertPosition `json:"cert_position"`
	ImgSize             ImageSize    `json:"img_size"`
}

// OCRBusinessLicense 营业执照识别
func OCRBusinessLicense(mode OCRMode, filename string, receiver *BusinessLicense) Action {
	return &WechatAPI{
		body: helpers.NewUploadBody("img", filename, func() ([]byte, error) {
			return ioutil.ReadFile(filename)
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("UPLOAD|%s?type=%s&access_token=%s", OCRBusinessLicenseURL, mode, accessToken)
		},
		decode: func(resp []byte) error {
			return json.Unmarshal(resp, receiver)
		},
	}
}

// OCRBusinessLicenseByURL 营业执照识别
func OCRBusinessLicenseByURL(mode OCRMode, imgURL string, receiver *BusinessLicense) Action {
	return &WechatAPI{
		url: func(accessToken string) string {
			return fmt.Sprintf("POST|%s?type=%s&img_url=%s&access_token=%s", OCRBusinessLicenseURL, mode, imgURL, accessToken)
		},
		decode: func(resp []byte) error {
			return json.Unmarshal(resp, receiver)
		},
	}
}

// DriverLicense 驾照
type DriverLicense struct {
	IDNum        string `json:"id_num"`
	Name         string `json:"name"`
	Sex          string `json:"sex"`
	Address      string `json:"address"`
	BirthDate    string `json:"birth_date"`
	IssueDate    string `json:"issue_date"`
	CarClass     string `json:"car_class"`
	ValidFrom    string `json:"valid_from"`
	ValidTo      string `json:"valid_to"`
	OfficialSeal string `json:"official_seal"`
}

// OCRDriverLicense 驾照识别
func OCRDriverLicense(mode OCRMode, filename string, receiver *DriverLicense) Action {
	return &WechatAPI{
		body: helpers.NewUploadBody("img", filename, func() ([]byte, error) {
			return ioutil.ReadFile(filename)
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("UPLOAD|%s?type=%s&access_token=%s", OCRDriverLicenseURL, mode, accessToken)
		},
		decode: func(resp []byte) error {
			return json.Unmarshal(resp, receiver)
		},
	}
}

// OCRDriverLicenseByURL 驾照识别
func OCRDriverLicenseByURL(mode OCRMode, imgURL string, receiver *DriverLicense) Action {
	return &WechatAPI{
		url: func(accessToken string) string {
			return fmt.Sprintf("POST|%s?type=%s&img_url=%s&access_token=%s", OCRDriverLicenseURL, mode, imgURL, accessToken)
		},
		decode: func(resp []byte) error {
			return json.Unmarshal(resp, receiver)
		},
	}
}

// IDCardFront 身份证前面
type IDCardFront struct {
	Name        string `json:"name"`
	ID          string `json:"id"`
	Addr        string `json:"addr"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}

// OCRIDCardFront 身份证前面识别
func OCRIDCardFront(mode OCRMode, filename string, receiver *IDCardFront) Action {
	return &WechatAPI{
		body: helpers.NewUploadBody("img", filename, func() ([]byte, error) {
			return ioutil.ReadFile(filename)
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("UPLOAD|%s?type=%s&access_token=%s", OCRIDCardURL, mode, accessToken)
		},
		decode: func(resp []byte) error {
			return json.Unmarshal(resp, receiver)
		},
	}
}

// OCRIDCardFrontByURL 身份证前面识别
func OCRIDCardFrontByURL(mode OCRMode, imgURL string, receiver *IDCardFront) Action {
	return &WechatAPI{
		url: func(accessToken string) string {
			return fmt.Sprintf("POST|%s?type=%s&img_url=%s&access_token=%s", OCRIDCardURL, mode, imgURL, accessToken)
		},
		decode: func(resp []byte) error {
			return json.Unmarshal(resp, receiver)
		},
	}
}

// IDCardBack 身份证背面
type IDCardBack struct {
	ValidDate string `json:"valid_date"`
}

// OCRIDCardBack 身份证背面识别
func OCRIDCardBack(mode OCRMode, filename string, receiver *IDCardBack) Action {
	return &WechatAPI{
		body: helpers.NewUploadBody("img", filename, func() ([]byte, error) {
			return ioutil.ReadFile(filename)
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("UPLOAD|%s?type=%s&access_token=%s", OCRBankCardURL, mode, accessToken)
		},
		decode: func(resp []byte) error {
			receiver.ValidDate = gjson.GetBytes(resp, "valid_date").String()

			return nil
		},
	}
}

// OCRIDCardBackByURL 身份证背面识别
func OCRIDCardBackByURL(mode OCRMode, imgURL string, receiver *IDCardBack) Action {
	return &WechatAPI{
		url: func(accessToken string) string {
			return fmt.Sprintf("POST|%s?type=%s&img_url=%s&access_token=%s", OCRBankCardURL, mode, imgURL, accessToken)
		},
		decode: func(resp []byte) error {
			receiver.ValidDate = gjson.GetBytes(resp, "valid_date").String()

			return nil
		},
	}
}

// PrintedText 通用印刷体
type PrintedText struct {
	Items   []PrintedTextItem `json:"items"`
	ImgSize ImageSize         `json:"img_size"`
}

// PrintedTextItem 通用印刷体内容项
type PrintedTextItem struct {
	Text string        `json:"text"`
	Pos  ImagePosition `json:"pos"`
}

// OCRPrintedText 通用印刷体识别
func OCRPrintedText(mode OCRMode, filename string, receiver *PrintedText) Action {
	return &WechatAPI{
		body: helpers.NewUploadBody("img", filename, func() ([]byte, error) {
			return ioutil.ReadFile(filename)
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("UPLOAD|%s?type=%s&access_token=%s", OCRPrintedTextURL, mode, accessToken)
		},
		decode: func(resp []byte) error {
			return json.Unmarshal(resp, receiver)
		},
	}
}

// OCRPrintedTextByURL 通用印刷体识别
func OCRPrintedTextByURL(mode OCRMode, imgURL string, receiver *PrintedText) Action {
	return &WechatAPI{
		url: func(accessToken string) string {
			return fmt.Sprintf("POST|%s?type=%s&img_url=%s&access_token=%s", OCRPrintedTextURL, mode, imgURL, accessToken)
		},
		decode: func(resp []byte) error {
			return json.Unmarshal(resp, receiver)
		},
	}
}

// VehicleLicense 行驶证
type VehicleLicense struct {
	VehicleType    string `json:"vehicle_type"`
	Owner          string `json:"owner"`
	Addr           string `json:"addr"`
	UseCharacter   string `json:"use_character"`
	Model          string `json:"model"`
	VIN            string `json:"vin"`
	EngineNum      string `json:"engine_num"`
	RegisterDate   string `json:"register_date"`
	IssueDate      string `json:"issue_date"`
	PlateNumB      string `json:"plate_num_b"`
	Record         string `json:"record"`
	PassengersNum  string `json:"passengers_num"`
	TotalQuality   string `json:"total_quality"`
	PrepareQuality string `json:"prepare_quality"`
}

// OCRVehicleLicense 行驶证识别
func OCRVehicleLicense(mode OCRMode, filename string, receiver *VehicleLicense) Action {
	return &WechatAPI{
		body: helpers.NewUploadBody("img", filename, func() ([]byte, error) {
			return ioutil.ReadFile(filename)
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("UPLOAD|%s?type=%s&access_token=%s", OCRVehicleLicenseURL, mode, accessToken)
		},
		decode: func(resp []byte) error {
			return json.Unmarshal(resp, receiver)
		},
	}
}

// OCRVehicleLicenseByURL 行驶证识别
func OCRVehicleLicenseByURL(mode OCRMode, imgURL string, receiver *VehicleLicense) Action {
	return &WechatAPI{
		url: func(accessToken string) string {
			return fmt.Sprintf("POST|%s?type=%s&img_url=%s&access_token=%s", OCRVehicleLicenseURL, mode, imgURL, accessToken)
		},
		decode: func(resp []byte) error {
			return json.Unmarshal(resp, receiver)
		},
	}
}
