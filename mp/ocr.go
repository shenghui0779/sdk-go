package mp

import (
	"encoding/json"
	"net/url"

	"github.com/shenghui0779/gochat/wx"
	"github.com/tidwall/gjson"
)

// OCRMode 识别模式
type OCRMode string

// 微信支持的识别模式
var (
	OCRPhoto OCRMode = "photo" // 拍照模式
	OCRScan  OCRMode = "scan"  // 扫描模式
)

// OCRPosition 识别位置
type OCRPosition struct {
	Pos ImagePosition `json:"pos"`
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
func OCRIDCardFront(dest *IDCardFront, mode OCRMode, filename string) wx.Action {
	return wx.NewAction(OCRIDCardURL,
		wx.WithMethod(wx.MethodUpload),
		wx.WithQuery("type", string(mode)),
		wx.WithUploadForm("img", filename, nil),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// OCRIDCardFrontByURL 身份证前面识别
func OCRIDCardFrontByURL(dest *IDCardFront, mode OCRMode, imgURL string) wx.Action {
	return wx.NewAction(OCRIDCardURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithQuery("type", string(mode)),
		wx.WithQuery("img_url", imgURL),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// IDCardBack 身份证背面
type IDCardBack struct {
	ValidDate string `json:"valid_date"`
}

// OCRIDCardBack 身份证背面识别
func OCRIDCardBack(dest *IDCardBack, mode OCRMode, filename string) wx.Action {
	query := url.Values{}

	query.Set("type", string(mode))

	return wx.NewAction(OCRIDCardURL,
		wx.WithMethod(wx.MethodUpload),
		wx.WithQuery("type", string(mode)),
		wx.WithUploadForm("img", filename, nil),
		wx.WithDecode(func(resp []byte) error {
			dest.ValidDate = gjson.GetBytes(resp, "valid_date").String()

			return nil
		}),
	)
}

// OCRIDCardBackByURL 身份证背面识别
func OCRIDCardBackByURL(dest *IDCardBack, mode OCRMode, imgURL string) wx.Action {
	query := url.Values{}

	query.Set("type", string(mode))
	query.Set("img_url", imgURL)

	return wx.NewAction(OCRIDCardURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithQuery("type", string(mode)),
		wx.WithQuery("img_url", imgURL),
		wx.WithDecode(func(resp []byte) error {
			dest.ValidDate = gjson.GetBytes(resp, "valid_date").String()

			return nil
		}),
	)
}

// BankCard 银行卡
type BankCard struct {
	ID string `json:"id"`
}

// OCRBankCard 银行卡识别
func OCRBankCard(dest *BankCard, mode OCRMode, filename string) wx.Action {
	query := url.Values{}

	query.Set("type", string(mode))

	return wx.NewAction(OCRBankCardURL,
		wx.WithMethod(wx.MethodUpload),
		wx.WithQuery("type", string(mode)),
		wx.WithUploadForm("img", filename, nil),
		wx.WithDecode(func(resp []byte) error {
			dest.ID = gjson.GetBytes(resp, "number").String()

			return nil
		}),
	)
}

// OCRBankCardByURL 银行卡识别
func OCRBankCardByURL(dest *BankCard, mode OCRMode, imgURL string) wx.Action {
	query := url.Values{}

	query.Set("type", string(mode))
	query.Set("img_url", imgURL)

	return wx.NewAction(OCRBankCardURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithQuery("type", string(mode)),
		wx.WithQuery("img_url", imgURL),
		wx.WithDecode(func(resp []byte) error {
			dest.ID = gjson.GetBytes(resp, "number").String()

			return nil
		}),
	)
}

// PlateNumber 车牌号
type PlateNumber struct {
	ID string `json:"id"`
}

// OCRPlateNumber 车牌号识别
func OCRPlateNumber(dest *PlateNumber, mode OCRMode, filename string) wx.Action {
	query := url.Values{}

	query.Set("type", string(mode))

	return wx.NewAction(OCRPlateNumberURL,
		wx.WithMethod(wx.MethodUpload),
		wx.WithQuery("type", string(mode)),
		wx.WithUploadForm("img", filename, nil),
		wx.WithDecode(func(resp []byte) error {
			dest.ID = gjson.GetBytes(resp, "number").String()

			return nil
		}),
	)
}

// OCRPlateNumberByURL 车牌号识别
func OCRPlateNumberByURL(dest *PlateNumber, mode OCRMode, imgURL string) wx.Action {
	query := url.Values{}

	query.Set("type", string(mode))
	query.Set("img_url", imgURL)

	return wx.NewAction(OCRPlateNumberURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithQuery("type", string(mode)),
		wx.WithQuery("img_url", imgURL),
		wx.WithDecode(func(resp []byte) error {
			dest.ID = gjson.GetBytes(resp, "number").String()

			return nil
		}),
	)
}

// DriverLicense 驾照
type DriverLicense struct {
	IDNum        string `json:"id_num"`        // 证号
	Name         string `json:"name"`          // 姓名
	Sex          string `json:"sex"`           // 性别
	Nationality  string `json:"nationality"`   // 国籍
	Address      string `json:"address"`       // 住址
	BirthDate    string `json:"birth_date"`    // 出生日期
	IssueDate    string `json:"issue_date"`    // 初次领证日期
	CarClass     string `json:"car_class"`     // 准驾车型
	ValidFrom    string `json:"valid_from"`    // 有效期限起始日
	ValidTo      string `json:"valid_to"`      // 有效期限终止日
	OfficialSeal string `json:"official_seal"` // 印章文字
}

// OCRDriverLicense 驾照识别
func OCRDriverLicense(dest *DriverLicense, mode OCRMode, filename string) wx.Action {
	query := url.Values{}

	query.Set("type", string(mode))

	return wx.NewAction(OCRDriverLicenseURL,
		wx.WithMethod(wx.MethodUpload),
		wx.WithQuery("type", string(mode)),
		wx.WithUploadForm("img", filename, nil),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// OCRDriverLicenseByURL 驾照识别
func OCRDriverLicenseByURL(dest *DriverLicense, mode OCRMode, imgURL string) wx.Action {
	query := url.Values{}

	query.Set("type", string(mode))
	query.Set("img_url", imgURL)

	return wx.NewAction(OCRDriverLicenseURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithQuery("type", string(mode)),
		wx.WithQuery("img_url", imgURL),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// VehicleLicense 行驶证
type VehicleLicense struct {
	VehicleType       string      `json:"vhicle_type"`         // 车辆类型
	Owner             string      `json:"owner"`               // 所有人
	Addr              string      `json:"addr"`                // 住址
	UseCharacter      string      `json:"use_character"`       // 使用性质
	Model             string      `json:"model"`               // 品牌型号
	VIN               string      `json:"vin"`                 // 车辆识别代号
	EngineNum         string      `json:"engine_num"`          // 发动机号码
	RegisterDate      string      `json:"register_date"`       // 注册日期
	IssueDate         string      `json:"issue_date"`          // 发证日期
	PlateNum          string      `json:"plate_num"`           // 车牌号码
	PlateNumB         string      `json:"plate_num_b"`         // 车牌号码
	Record            string      `json:"record"`              // 号牌
	PassengersNum     string      `json:"passengers_num"`      // 核定载人数
	TotalQuality      string      `json:"total_quality"`       // 总质量
	PrepareQuality    string      `json:"prepare_quality"`     // 整备质量
	OverallSize       string      `json:"overall_size"`        // 外廓尺寸
	CardPositionFront OCRPosition `json:"card_position_front"` // 卡片正面位置（检测到卡片正面才会返回）
	CardPositionBack  OCRPosition `json:"card_position_back"`  // 卡片反面位置（检测到卡片反面才会返回）
}

// OCRVehicleLicense 行驶证识别
func OCRVehicleLicense(dest *VehicleLicense, mode OCRMode, filename string) wx.Action {
	query := url.Values{}

	query.Set("type", string(mode))

	return wx.NewAction(OCRVehicleLicenseURL,
		wx.WithMethod(wx.MethodUpload),
		wx.WithQuery("type", string(mode)),
		wx.WithUploadForm("img", filename, nil),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// OCRVehicleLicenseByURL 行驶证识别
func OCRVehicleLicenseByURL(dest *VehicleLicense, mode OCRMode, imgURL string) wx.Action {
	query := url.Values{}

	query.Set("type", string(mode))
	query.Set("img_url", imgURL)

	return wx.NewAction(OCRVehicleLicenseURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithQuery("type", string(mode)),
		wx.WithQuery("img_url", imgURL),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// BusinessLicense 营业执照
type BusinessLicense struct {
	RegNum              string      `json:"reg_num"`              // 注册号
	Serial              string      `json:"serial"`               // 编号
	LegalRepresentative string      `json:"legal_representative"` // 法定代表人姓名
	EnterpriseName      string      `json:"enterprise_name"`      // 企业名称
	TypeOfOrganization  string      `json:"type_of_organization"` // 组成形式
	Address             string      `json:"address"`              // 经营场所/企业住所
	TypeOfEnterprise    string      `json:"type_of_enterprise"`   // 公司类型
	BusinessScope       string      `json:"business_scope"`       // 经营范围
	RegisteredCapital   string      `json:"registered_capital"`   // 注册资本
	PaidInCapital       string      `json:"paid_in_capital"`      // 实收资本
	ValidPeriod         string      `json:"valid_period"`         // 营业期限
	RegisteredDate      string      `json:"registered_date"`      // 注册日期/成立日期
	CertPosition        OCRPosition `json:"cert_position"`        // 营业执照位置
	ImgSize             ImageSize   `json:"img_size"`             // 图片大小
}

// OCRBusinessLicense 营业执照识别
func OCRBusinessLicense(dest *BusinessLicense, mode OCRMode, filename string) wx.Action {
	query := url.Values{}

	query.Set("type", string(mode))

	return wx.NewAction(OCRBusinessLicenseURL,
		wx.WithMethod(wx.MethodUpload),
		wx.WithQuery("type", string(mode)),
		wx.WithUploadForm("img", filename, nil),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// OCRBusinessLicenseByURL 营业执照识别
func OCRBusinessLicenseByURL(dest *BusinessLicense, mode OCRMode, imgURL string) wx.Action {
	query := url.Values{}

	query.Set("type", string(mode))
	query.Set("img_url", imgURL)

	return wx.NewAction(OCRBusinessLicenseURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithQuery("type", string(mode)),
		wx.WithQuery("img_url", imgURL),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// PrintedText 通用印刷体
type PrintedText struct {
	Items   []*PrintedTextItem `json:"items"`
	ImgSize ImageSize          `json:"img_size"`
}

// PrintedTextItem 通用印刷体内容项
type PrintedTextItem struct {
	Text string        `json:"text"`
	Pos  ImagePosition `json:"pos"`
}

// OCRPrintedText 通用印刷体识别
func OCRPrintedText(dest *PrintedText, mode OCRMode, filename string) wx.Action {
	return wx.NewAction(OCRPrintedTextURL,
		wx.WithMethod(wx.MethodUpload),
		wx.WithQuery("type", string(mode)),
		wx.WithUploadForm("img", filename, nil),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// OCRPrintedTextByURL 通用印刷体识别
func OCRPrintedTextByURL(dest *PrintedText, mode OCRMode, imgURL string) wx.Action {
	query := url.Values{}

	query.Set("type", string(mode))
	query.Set("img_url", imgURL)

	return wx.NewAction(OCRPrintedTextURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithQuery("type", string(mode)),
		wx.WithQuery("img_url", imgURL),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}
