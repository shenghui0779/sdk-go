package minip

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
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

// ResultIDCardFrontOCR 身份证前面识别结果
type ResultIDCardFrontOCR struct {
	Name        string `json:"name"`
	ID          string `json:"id"`
	Addr        string `json:"addr"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}

// OCRIDCardFront OCR - 身份证前面识别
func OCRIDCardFront(mode OCRMode, imgPath string, result *ResultIDCardFrontOCR) wx.Action {
	_, filename := filepath.Split(imgPath)

	return wx.NewPostAction(urls.MinipOCRIDCard,
		wx.WithQuery("type", string(mode)),
		wx.WithUpload(func() (wx.UploadForm, error) {
			path, err := filepath.Abs(filepath.Clean(imgPath))

			if err != nil {
				return nil, err
			}

			return wx.NewUploadForm(
				wx.WithFormFile("img", filename, func(w io.Writer) error {
					f, err := os.Open(path)

					if err != nil {
						return err
					}

					defer f.Close()

					if _, err = io.Copy(w, f); err != nil {
						return err
					}

					return nil
				}),
			), nil
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// OCRIDCardFrontByURL OCR - 身份证前面识别
func OCRIDCardFrontByURL(mode OCRMode, imgURL string, result *ResultIDCardFrontOCR) wx.Action {
	return wx.NewPostAction(urls.MinipOCRIDCard,
		wx.WithQuery("type", string(mode)),
		wx.WithQuery("img_url", imgURL),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// ResultIDCardBackOCR 身份证背面识别结果
type ResultIDCardBackOCR struct {
	ValidDate string `json:"valid_date"`
}

// OCRIDCardBack OCR - 身份证背面识别
func OCRIDCardBack(mode OCRMode, imgPath string, result *ResultIDCardBackOCR) wx.Action {
	_, filename := filepath.Split(imgPath)

	return wx.NewPostAction(urls.MinipOCRIDCard,
		wx.WithQuery("type", string(mode)),
		wx.WithUpload(func() (wx.UploadForm, error) {
			path, err := filepath.Abs(filepath.Clean(imgPath))

			if err != nil {
				return nil, err
			}

			return wx.NewUploadForm(
				wx.WithFormFile("img", filename, func(w io.Writer) error {
					f, err := os.Open(path)

					if err != nil {
						return err
					}

					defer f.Close()

					if _, err = io.Copy(w, f); err != nil {
						return err
					}

					return nil
				}),
			), nil
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// OCRIDCardBackByURL OCR - 身份证背面识别
func OCRIDCardBackByURL(mode OCRMode, imgURL string, result *ResultIDCardBackOCR) wx.Action {
	return wx.NewPostAction(urls.MinipOCRIDCard,
		wx.WithQuery("type", string(mode)),
		wx.WithQuery("img_url", imgURL),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// ResultBankCardOCR 银行卡识别结果
type ResultBankCardOCR struct {
	Number string `json:"number"`
}

// OCRBankCard OCR - 银行卡识别
func OCRBankCard(mode OCRMode, imgPath string, result *ResultBankCardOCR) wx.Action {
	_, filename := filepath.Split(imgPath)

	return wx.NewPostAction(urls.MinipOCRBankCard,
		wx.WithQuery("type", string(mode)),
		wx.WithUpload(func() (wx.UploadForm, error) {
			path, err := filepath.Abs(filepath.Clean(imgPath))

			if err != nil {
				return nil, err
			}

			return wx.NewUploadForm(
				wx.WithFormFile("img", filename, func(w io.Writer) error {
					f, err := os.Open(path)

					if err != nil {
						return err
					}

					defer f.Close()

					if _, err = io.Copy(w, f); err != nil {
						return err
					}

					return nil
				}),
			), nil
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// OCRBankCardByURL OCR - 银行卡识别
func OCRBankCardByURL(mode OCRMode, imgURL string, result *ResultBankCardOCR) wx.Action {
	return wx.NewPostAction(urls.MinipOCRBankCard,
		wx.WithQuery("type", string(mode)),
		wx.WithQuery("img_url", imgURL),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// ResultPlateNumberOCR 车牌号识别结果
type ResultPlateNumberOCR struct {
	Number string `json:"number"`
}

// OCRPlateNumber OCR - 车牌号识别
func OCRPlateNumber(mode OCRMode, imgPath string, result *ResultPlateNumberOCR) wx.Action {
	_, filename := filepath.Split(imgPath)

	return wx.NewPostAction(urls.MinipOCRPlateNumber,
		wx.WithQuery("type", string(mode)),
		wx.WithUpload(func() (wx.UploadForm, error) {
			path, err := filepath.Abs(filepath.Clean(imgPath))

			if err != nil {
				return nil, err
			}

			return wx.NewUploadForm(
				wx.WithFormFile("img", filename, func(w io.Writer) error {
					f, err := os.Open(path)

					if err != nil {
						return err
					}

					defer f.Close()

					if _, err = io.Copy(w, f); err != nil {
						return err
					}

					return nil
				}),
			), nil
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// OCRPlateNumberByURL OCR - 车牌号识别
func OCRPlateNumberByURL(mode OCRMode, imgURL string, result *ResultPlateNumberOCR) wx.Action {
	return wx.NewPostAction(urls.MinipOCRPlateNumber,
		wx.WithQuery("type", string(mode)),
		wx.WithQuery("img_url", imgURL),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// ResultDriverLicenseOCR 驾照识别结果
type ResultDriverLicenseOCR struct {
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

// OCRDriverLicense OCR - 驾照识别
func OCRDriverLicense(mode OCRMode, imgPath string, result *ResultDriverLicenseOCR) wx.Action {
	_, filename := filepath.Split(imgPath)

	return wx.NewPostAction(urls.MinipOCRDriverLicense,
		wx.WithQuery("type", string(mode)),
		wx.WithUpload(func() (wx.UploadForm, error) {
			path, err := filepath.Abs(filepath.Clean(imgPath))

			if err != nil {
				return nil, err
			}

			return wx.NewUploadForm(
				wx.WithFormFile("img", filename, func(w io.Writer) error {
					f, err := os.Open(path)

					if err != nil {
						return err
					}

					defer f.Close()

					if _, err = io.Copy(w, f); err != nil {
						return err
					}

					return nil
				}),
			), nil
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// OCRDriverLicenseByURL OCR - 驾照识别
func OCRDriverLicenseByURL(mode OCRMode, imgURL string, result *ResultDriverLicenseOCR) wx.Action {
	return wx.NewPostAction(urls.MinipOCRDriverLicense,
		wx.WithQuery("type", string(mode)),
		wx.WithQuery("img_url", imgURL),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// ResultVehicleLicenseOCR 行驶证识别结果
type ResultVehicleLicenseOCR struct {
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

// OCRVehicleLicense OCR - 行驶证识别
func OCRVehicleLicense(mode OCRMode, imgPath string, result *ResultVehicleLicenseOCR) wx.Action {
	_, filename := filepath.Split(imgPath)

	return wx.NewPostAction(urls.MinipOCRVehicleLicense,
		wx.WithQuery("type", string(mode)),
		wx.WithUpload(func() (wx.UploadForm, error) {
			path, err := filepath.Abs(filepath.Clean(imgPath))

			if err != nil {
				return nil, err
			}

			return wx.NewUploadForm(
				wx.WithFormFile("img", filename, func(w io.Writer) error {
					f, err := os.Open(path)

					if err != nil {
						return err
					}

					defer f.Close()

					if _, err = io.Copy(w, f); err != nil {
						return err
					}

					return nil
				}),
			), nil
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// OCRVehicleLicenseByURL OCR - 行驶证识别
func OCRVehicleLicenseByURL(mode OCRMode, imgURL string, result *ResultVehicleLicenseOCR) wx.Action {
	return wx.NewPostAction(urls.MinipOCRVehicleLicense,
		wx.WithQuery("type", string(mode)),
		wx.WithQuery("img_url", imgURL),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// ResultBusinessLicenseOCR 营业执照
type ResultBusinessLicenseOCR struct {
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

// OCRBusinessLicense OCR - 营业执照识别
func OCRBusinessLicense(mode OCRMode, imgPath string, result *ResultBusinessLicenseOCR) wx.Action {
	_, filename := filepath.Split(imgPath)

	return wx.NewPostAction(urls.MinipOCRBusinessLicense,
		wx.WithQuery("type", string(mode)),
		wx.WithUpload(func() (wx.UploadForm, error) {
			path, err := filepath.Abs(filepath.Clean(imgPath))

			if err != nil {
				return nil, err
			}

			return wx.NewUploadForm(
				wx.WithFormFile("img", filename, func(w io.Writer) error {
					f, err := os.Open(path)

					if err != nil {
						return err
					}

					defer f.Close()

					if _, err = io.Copy(w, f); err != nil {
						return err
					}

					return nil
				}),
			), nil
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// OCRBusinessLicenseByURL OCR - 营业执照识别
func OCRBusinessLicenseByURL(mode OCRMode, imgURL string, result *ResultBusinessLicenseOCR) wx.Action {
	return wx.NewPostAction(urls.MinipOCRBusinessLicense,
		wx.WithQuery("type", string(mode)),
		wx.WithQuery("img_url", imgURL),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// CommOCRItem 通用印刷体内容项
type CommOCRItem struct {
	Text string        `json:"text"`
	Pos  ImagePosition `json:"pos"`
}

// ResultCommOCR 通用印刷体识别结果
type ResultCommOCR struct {
	Items   []*CommOCRItem `json:"items"`
	ImgSize ImageSize      `json:"img_size"`
}

// OCRComm OCR - 通用印刷体识别
func OCRComm(mode OCRMode, imgPath string, result *ResultCommOCR) wx.Action {
	_, filename := filepath.Split(imgPath)

	return wx.NewPostAction(urls.MinipOCRComm,
		wx.WithQuery("type", string(mode)),
		wx.WithUpload(func() (wx.UploadForm, error) {
			path, err := filepath.Abs(filepath.Clean(imgPath))

			if err != nil {
				return nil, err
			}

			return wx.NewUploadForm(
				wx.WithFormFile("img", filename, func(w io.Writer) error {
					f, err := os.Open(path)

					if err != nil {
						return err
					}

					defer f.Close()

					if _, err = io.Copy(w, f); err != nil {
						return err
					}

					return nil
				}),
			), nil
		}),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// OCRCommByURL OCR - 通用印刷体识别
func OCRCommByURL(mode OCRMode, imgURL string, result *ResultCommOCR) wx.Action {
	return wx.NewPostAction(urls.MinipOCRComm,
		wx.WithQuery("type", string(mode)),
		wx.WithQuery("img_url", imgURL),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}
