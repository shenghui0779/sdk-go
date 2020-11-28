package mp

import (
	"encoding/json"
	"io/ioutil"
	"net/url"

	"github.com/shenghui0779/gochat/internal"
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
func OCRBankCard(mode OCRMode, filename string, dest *BankCard) internal.Action {
	query := url.Values{}

	query.Set("type", string(mode))

	return internal.NewOpenUploadAPI(OCRBankCardURL, query, internal.NewUploadBody("img", filename, func() ([]byte, error) {
		return ioutil.ReadFile(filename)
	}), func(resp []byte) error {
		dest.ID = gjson.GetBytes(resp, "id").String()

		return nil
	})
}

// OCRBankCardByURL 银行卡识别
func OCRBankCardByURL(mode OCRMode, imgURL string, dest *BankCard) internal.Action {
	query := url.Values{}

	query.Set("type", string(mode))
	query.Set("img_url", imgURL)

	return internal.NewOpenPostAPI(OCRBankCardURL, query, nil, func(resp []byte) error {
		dest.ID = gjson.GetBytes(resp, "id").String()

		return nil
	})
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
func OCRBusinessLicense(mode OCRMode, filename string, dest *BusinessLicense) internal.Action {
	query := url.Values{}

	query.Set("type", string(mode))

	return internal.NewOpenUploadAPI(OCRBusinessLicenseURL, query, internal.NewUploadBody("img", filename, func() ([]byte, error) {
		return ioutil.ReadFile(filename)
	}), func(resp []byte) error {
		return json.Unmarshal(resp, dest)
	})
}

// OCRBusinessLicenseByURL 营业执照识别
func OCRBusinessLicenseByURL(mode OCRMode, imgURL string, dest *BusinessLicense) internal.Action {
	query := url.Values{}

	query.Set("type", string(mode))
	query.Set("img_url", imgURL)

	return internal.NewOpenPostAPI(OCRBusinessLicenseURL, query, nil, func(resp []byte) error {
		return json.Unmarshal(resp, dest)
	})
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
func OCRDriverLicense(mode OCRMode, filename string, dest *DriverLicense) internal.Action {
	query := url.Values{}

	query.Set("type", string(mode))

	return internal.NewOpenUploadAPI(OCRDriverLicenseURL, query, internal.NewUploadBody("img", filename, func() ([]byte, error) {
		return ioutil.ReadFile(filename)
	}), func(resp []byte) error {
		return json.Unmarshal(resp, dest)
	})
}

// OCRDriverLicenseByURL 驾照识别
func OCRDriverLicenseByURL(mode OCRMode, imgURL string, dest *DriverLicense) internal.Action {
	query := url.Values{}

	query.Set("type", string(mode))
	query.Set("img_url", imgURL)

	return internal.NewOpenPostAPI(OCRDriverLicenseURL, query, nil, func(resp []byte) error {
		return json.Unmarshal(resp, dest)
	})
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
func OCRIDCardFront(mode OCRMode, filename string, dest *IDCardFront) internal.Action {
	query := url.Values{}

	query.Set("type", string(mode))

	return internal.NewOpenUploadAPI(OCRIDCardURL, query, internal.NewUploadBody("img", filename, func() ([]byte, error) {
		return ioutil.ReadFile(filename)
	}), func(resp []byte) error {
		return json.Unmarshal(resp, dest)
	})
}

// OCRIDCardFrontByURL 身份证前面识别
func OCRIDCardFrontByURL(mode OCRMode, imgURL string, dest *IDCardFront) internal.Action {
	query := url.Values{}

	query.Set("type", string(mode))
	query.Set("img_url", imgURL)

	return internal.NewOpenPostAPI(OCRIDCardURL, query, nil, func(resp []byte) error {
		return json.Unmarshal(resp, dest)
	})
}

// IDCardBack 身份证背面
type IDCardBack struct {
	ValidDate string `json:"valid_date"`
}

// OCRIDCardBack 身份证背面识别
func OCRIDCardBack(mode OCRMode, filename string, dest *IDCardBack) internal.Action {
	query := url.Values{}

	query.Set("type", string(mode))

	return internal.NewOpenUploadAPI(OCRBankCardURL, query, internal.NewUploadBody("img", filename, func() ([]byte, error) {
		return ioutil.ReadFile(filename)
	}), func(resp []byte) error {
		dest.ValidDate = gjson.GetBytes(resp, "valid_date").String()

		return nil
	})
}

// OCRIDCardBackByURL 身份证背面识别
func OCRIDCardBackByURL(mode OCRMode, imgURL string, dest *IDCardBack) internal.Action {
	query := url.Values{}

	query.Set("type", string(mode))
	query.Set("img_url", imgURL)

	return internal.NewOpenPostAPI(OCRBankCardURL, query, nil, func(resp []byte) error {
		dest.ValidDate = gjson.GetBytes(resp, "valid_date").String()

		return nil
	})
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
func OCRPrintedText(mode OCRMode, filename string, dest *PrintedText) internal.Action {
	query := url.Values{}

	query.Set("type", string(mode))

	return internal.NewOpenUploadAPI(OCRPrintedTextURL, query, internal.NewUploadBody("img", filename, func() ([]byte, error) {
		return ioutil.ReadFile(filename)
	}), func(resp []byte) error {
		return json.Unmarshal(resp, dest)
	})
}

// OCRPrintedTextByURL 通用印刷体识别
func OCRPrintedTextByURL(mode OCRMode, imgURL string, dest *PrintedText) internal.Action {
	query := url.Values{}

	query.Set("type", string(mode))
	query.Set("img_url", imgURL)

	return internal.NewOpenPostAPI(OCRPrintedTextURL, query, nil, func(resp []byte) error {
		return json.Unmarshal(resp, dest)
	})
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
func OCRVehicleLicense(mode OCRMode, filename string, dest *VehicleLicense) internal.Action {
	query := url.Values{}

	query.Set("type", string(mode))

	return internal.NewOpenUploadAPI(OCRVehicleLicenseURL, query, internal.NewUploadBody("img", filename, func() ([]byte, error) {
		return ioutil.ReadFile(filename)
	}), func(resp []byte) error {
		return json.Unmarshal(resp, dest)
	})
}

// OCRVehicleLicenseByURL 行驶证识别
func OCRVehicleLicenseByURL(mode OCRMode, imgURL string, dest *VehicleLicense) internal.Action {
	query := url.Values{}

	query.Set("type", string(mode))
	query.Set("img_url", imgURL)

	return internal.NewOpenPostAPI(OCRVehicleLicenseURL, query, nil, func(resp []byte) error {
		return json.Unmarshal(resp, dest)
	})
}
