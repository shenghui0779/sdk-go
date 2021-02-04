package mp

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/wx"
	"github.com/stretchr/testify/assert"
)

func TestOCRIDCardFront(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cv/ocr/idcard?access_token=ACCESS_TOKEN&type=photo", wx.NewUploadForm("img", "test.jpg", nil)).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"type": "Front",
		"name": "张三",
		"id": "123456789012345678",
		"addr": "广东省广州市",
		"gender": "男",
		"nationality": "汉"
	  }`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(IDCardFront)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", OCRIDCardFront(dest, OCRPhoto, "test.jpg"))

	assert.Nil(t, err)
	assert.Equal(t, &IDCardFront{
		Name:        "张三",
		ID:          "123456789012345678",
		Addr:        "广东省广州市",
		Gender:      "男",
		Nationality: "汉",
	}, dest)
}

func TestOCRIDCardFrontByURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cv/ocr/idcard?access_token=ACCESS_TOKEN&img_url=ENCODE_URL&type=photo", nil).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"type": "Front",
		"name": "张三",
		"id": "123456789012345678",
		"addr": "广东省广州市",
		"gender": "男",
		"nationality": "汉"
	  }`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(IDCardFront)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", OCRIDCardFrontByURL(dest, OCRPhoto, "ENCODE_URL"))

	assert.Nil(t, err)
	assert.Equal(t, &IDCardFront{
		Name:        "张三",
		ID:          "123456789012345678",
		Addr:        "广东省广州市",
		Gender:      "男",
		Nationality: "汉",
	}, dest)
}

func TestOCRIDCardBack(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cv/ocr/idcard?access_token=ACCESS_TOKEN&type=photo", wx.NewUploadForm("img", "test.jpg", nil)).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"type": "Back",
		"valid_date": "20070105-20270105"
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(IDCardBack)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", OCRIDCardBack(dest, OCRPhoto, "test.jpg"))

	assert.Nil(t, err)
	assert.Equal(t, "20070105-20270105", dest.ValidDate)
}

func TestOCRIDCardBackByURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cv/ocr/idcard?access_token=ACCESS_TOKEN&img_url=ENCODE_URL&type=photo", nil).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"type": "Back",
		"valid_date": "20070105-20270105"
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(IDCardBack)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", OCRIDCardBackByURL(dest, OCRPhoto, "ENCODE_URL"))

	assert.Nil(t, err)
	assert.Equal(t, "20070105-20270105", dest.ValidDate)
}

func TestOCRBankCard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cv/ocr/bankcard?access_token=ACCESS_TOKEN&type=photo", wx.NewUploadForm("img", "test.jpg", nil)).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"number": "622213XXXXXXXXX"
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(BankCard)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", OCRBankCard(dest, OCRPhoto, "test.jpg"))

	assert.Nil(t, err)
	assert.Equal(t, "622213XXXXXXXXX", dest.ID)
}

func TestOCRBankCardByURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cv/ocr/bankcard?access_token=ACCESS_TOKEN&img_url=ENCODE_URL&type=photo", nil).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"number": "622213XXXXXXXXX"
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(BankCard)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", OCRBankCardByURL(dest, OCRPhoto, "ENCODE_URL"))

	assert.Nil(t, err)
	assert.Equal(t, "622213XXXXXXXXX", dest.ID)
}

func TestOCRPlateNumber(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cv/ocr/platenum?access_token=ACCESS_TOKEN&type=photo", wx.NewUploadForm("img", "test.jpg", nil)).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"number": "苏A123456"
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(PlateNumber)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", OCRPlateNumber(dest, OCRPhoto, "test.jpg"))

	assert.Nil(t, err)
	assert.Equal(t, "苏A123456", dest.ID)
}

func TestOCRPlateNumberByURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cv/ocr/platenum?access_token=ACCESS_TOKEN&img_url=ENCODE_URL&type=photo", nil).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"number": "苏A123456"
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(PlateNumber)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", OCRPlateNumberByURL(dest, OCRPhoto, "ENCODE_URL"))

	assert.Nil(t, err)
	assert.Equal(t, "苏A123456", dest.ID)
}

func TestOCRDriverLicense(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cv/ocr/drivinglicense?access_token=ACCESS_TOKEN&type=photo", wx.NewUploadForm("img", "test.jpg", nil)).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"id_num": "660601xxxxxxxx1234",
		"name": "张三",
		"sex": "男",
		"nationality": "中国",
		"address": "广东省东莞市xxxxx号",
		"birth_date": "1990-12-21",
		"issue_date": "2012-12-21",
		"car_class": "C1",
		"valid_from": "2018-07-06",
		"valid_to": "2020-07-01",
		"official_seal": "xx市公安局公安交通管理局"
	   }`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(DriverLicense)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", OCRDriverLicense(dest, OCRPhoto, "test.jpg"))

	assert.Nil(t, err)
	assert.Equal(t, &DriverLicense{
		IDNum:        "660601xxxxxxxx1234",
		Name:         "张三",
		Sex:          "男",
		Nationality:  "中国",
		Address:      "广东省东莞市xxxxx号",
		BirthDate:    "1990-12-21",
		IssueDate:    "2012-12-21",
		CarClass:     "C1",
		ValidFrom:    "2018-07-06",
		ValidTo:      "2020-07-01",
		OfficialSeal: "xx市公安局公安交通管理局",
	}, dest)
}

func TestOCRDriverLicenseByURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cv/ocr/drivinglicense?access_token=ACCESS_TOKEN&img_url=ENCODE_URL&type=photo", nil).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"id_num": "660601xxxxxxxx1234",
		"name": "张三",
		"sex": "男",
		"nationality": "中国",
		"address": "广东省东莞市xxxxx号",
		"birth_date": "1990-12-21",
		"issue_date": "2012-12-21",
		"car_class": "C1",
		"valid_from": "2018-07-06",
		"valid_to": "2020-07-01",
		"official_seal": "xx市公安局公安交通管理局"
	   }`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(DriverLicense)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", OCRDriverLicenseByURL(dest, OCRPhoto, "ENCODE_URL"))

	assert.Nil(t, err)
	assert.Equal(t, &DriverLicense{
		IDNum:        "660601xxxxxxxx1234",
		Name:         "张三",
		Sex:          "男",
		Nationality:  "中国",
		Address:      "广东省东莞市xxxxx号",
		BirthDate:    "1990-12-21",
		IssueDate:    "2012-12-21",
		CarClass:     "C1",
		ValidFrom:    "2018-07-06",
		ValidTo:      "2020-07-01",
		OfficialSeal: "xx市公安局公安交通管理局",
	}, dest)
}

func TestOCRVehicleLicense(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cv/ocr/driving?access_token=ACCESS_TOKEN&type=photo", wx.NewUploadForm("img", "test.jpg", nil)).Return([]byte(`{
		"vhicle_type": "小型普通客⻋",
		"owner": "东莞市xxxxx机械厂",
		"addr": "广东省东莞市xxxxx号",
		"use_character": "非营运",
		"model": "江淮牌HFCxxxxxxx",
		"vin": "LJ166xxxxxxxx51",
		"engine_num": "J3xxxxx3",
		"register_date": "2018-07-06",
		"issue_date": "2018-07-01",
		"plate_num": "粤xxxxx",
		"plate_num_b": "粤xxxxx",
		"record": "441xxxxxx3",
		"passengers_num": "7人",
		"total_quality": "2700kg",
		"prepare_quality": "1995kg",
		"overall_size": "4582x1795x1458mm"
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(VehicleLicense)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", OCRVehicleLicense(dest, OCRPhoto, "test.jpg"))

	assert.Nil(t, err)
	assert.Equal(t, &VehicleLicense{
		VehicleType:    "小型普通客⻋",
		Owner:          "东莞市xxxxx机械厂",
		Addr:           "广东省东莞市xxxxx号",
		UseCharacter:   "非营运",
		Model:          "江淮牌HFCxxxxxxx",
		VIN:            "LJ166xxxxxxxx51",
		EngineNum:      "J3xxxxx3",
		RegisterDate:   "2018-07-06",
		IssueDate:      "2018-07-01",
		PlateNum:       "粤xxxxx",
		PlateNumB:      "粤xxxxx",
		Record:         "441xxxxxx3",
		PassengersNum:  "7人",
		TotalQuality:   "2700kg",
		PrepareQuality: "1995kg",
		OverallSize:    "4582x1795x1458mm",
	}, dest)
}

func TestOCRVehicleLicenseByURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cv/ocr/driving?access_token=ACCESS_TOKEN&img_url=ENCODE_URL&type=photo", nil).Return([]byte(`{
		"vhicle_type": "小型普通客⻋",
		"owner": "东莞市xxxxx机械厂",
		"addr": "广东省东莞市xxxxx号",
		"use_character": "非营运",
		"model": "江淮牌HFCxxxxxxx",
		"vin": "LJ166xxxxxxxx51",
		"engine_num": "J3xxxxx3",
		"register_date": "2018-07-06",
		"issue_date": "2018-07-01",
		"plate_num": "粤xxxxx",
		"plate_num_b": "粤xxxxx",
		"record": "441xxxxxx3",
		"passengers_num": "7人",
		"total_quality": "2700kg",
		"prepare_quality": "1995kg",
		"overall_size": "4582x1795x1458mm"
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(VehicleLicense)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", OCRVehicleLicenseByURL(dest, OCRPhoto, "ENCODE_URL"))

	assert.Nil(t, err)
	assert.Equal(t, &VehicleLicense{
		VehicleType:    "小型普通客⻋",
		Owner:          "东莞市xxxxx机械厂",
		Addr:           "广东省东莞市xxxxx号",
		UseCharacter:   "非营运",
		Model:          "江淮牌HFCxxxxxxx",
		VIN:            "LJ166xxxxxxxx51",
		EngineNum:      "J3xxxxx3",
		RegisterDate:   "2018-07-06",
		IssueDate:      "2018-07-01",
		PlateNum:       "粤xxxxx",
		PlateNumB:      "粤xxxxx",
		Record:         "441xxxxxx3",
		PassengersNum:  "7人",
		TotalQuality:   "2700kg",
		PrepareQuality: "1995kg",
		OverallSize:    "4582x1795x1458mm",
	}, dest)
}

func TestOCRBusinessLicense(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cv/ocr/bizlicense?access_token=ACCESS_TOKEN&type=photo", wx.NewUploadForm("img", "test.jpg", nil)).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"reg_num": "123123",
		"serial": "123123",
		"legal_representative": "张三",
		"enterprise_name": "XX饮食店",
		"type_of_organization": "个人经营",
		"address": "XX市XX区XX路XX号",
		"type_of_enterprise": "xxx",
		"business_scope": "中型餐馆(不含凉菜、不含裱花蛋糕，不含生食海产品)。",
		"registered_capital": "200万",
		"paid_in_capital": "200万",
		"valid_period": "2019年1月1日",
		"registered_date": "2018年1月1日",
		"cert_position": {
			"pos": {
				"left_top": {
					"x": 155,
					"y": 191
				},
				"right_top": {
					"x": 725,
					"y": 157
				},
				"right_bottom": {
					"x": 743,
					"y": 512
				},
				"left_bottom": {
					"x": 164,
					"y": 525
				}
			}
		},
		"img_size": {
			"w": 966,
			"h": 728
		}
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(BusinessLicense)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", OCRBusinessLicense(dest, OCRPhoto, "test.jpg"))

	assert.Nil(t, err)
	assert.Equal(t, &BusinessLicense{
		RegNum:              "123123",
		Serial:              "123123",
		LegalRepresentative: "张三",
		EnterpriseName:      "XX饮食店",
		TypeOfOrganization:  "个人经营",
		Address:             "XX市XX区XX路XX号",
		TypeOfEnterprise:    "xxx",
		BusinessScope:       "中型餐馆(不含凉菜、不含裱花蛋糕，不含生食海产品)。",
		RegisteredCapital:   "200万",
		PaidInCapital:       "200万",
		ValidPeriod:         "2019年1月1日",
		RegisteredDate:      "2018年1月1日",
		CertPosition: OCRPosition{
			Pos: ImagePosition{
				LeftTop: Position{
					X: 155,
					Y: 191,
				},
				RightTop: Position{
					X: 725,
					Y: 157,
				},
				RightBottom: Position{
					X: 743,
					Y: 512,
				},
				LeftBottom: Position{
					X: 164,
					Y: 525,
				},
			},
		},
		ImgSize: ImageSize{
			W: 966,
			H: 728,
		},
	}, dest)
}

func TestOCRBusinessLicenseByURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cv/ocr/bizlicense?access_token=ACCESS_TOKEN&img_url=ENCODE_URL&type=photo", nil).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"reg_num": "123123",
		"serial": "123123",
		"legal_representative": "张三",
		"enterprise_name": "XX饮食店",
		"type_of_organization": "个人经营",
		"address": "XX市XX区XX路XX号",
		"type_of_enterprise": "xxx",
		"business_scope": "中型餐馆(不含凉菜、不含裱花蛋糕，不含生食海产品)。",
		"registered_capital": "200万",
		"paid_in_capital": "200万",
		"valid_period": "2019年1月1日",
		"registered_date": "2018年1月1日",
		"cert_position": {
			"pos": {
				"left_top": {
					"x": 155,
					"y": 191
				},
				"right_top": {
					"x": 725,
					"y": 157
				},
				"right_bottom": {
					"x": 743,
					"y": 512
				},
				"left_bottom": {
					"x": 164,
					"y": 525
				}
			}
		},
		"img_size": {
			"w": 966,
			"h": 728
		}
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(BusinessLicense)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", OCRBusinessLicenseByURL(dest, OCRPhoto, "ENCODE_URL"))

	assert.Nil(t, err)
	assert.Equal(t, &BusinessLicense{
		RegNum:              "123123",
		Serial:              "123123",
		LegalRepresentative: "张三",
		EnterpriseName:      "XX饮食店",
		TypeOfOrganization:  "个人经营",
		Address:             "XX市XX区XX路XX号",
		TypeOfEnterprise:    "xxx",
		BusinessScope:       "中型餐馆(不含凉菜、不含裱花蛋糕，不含生食海产品)。",
		RegisteredCapital:   "200万",
		PaidInCapital:       "200万",
		ValidPeriod:         "2019年1月1日",
		RegisteredDate:      "2018年1月1日",
		CertPosition: OCRPosition{
			Pos: ImagePosition{
				LeftTop: Position{
					X: 155,
					Y: 191,
				},
				RightTop: Position{
					X: 725,
					Y: 157,
				},
				RightBottom: Position{
					X: 743,
					Y: 512,
				},
				LeftBottom: Position{
					X: 164,
					Y: 525,
				},
			},
		},
		ImgSize: ImageSize{
			W: 966,
			H: 728,
		},
	}, dest)
}

func TestOCRPrintedText(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cv/ocr/comm?access_token=ACCESS_TOKEN&type=photo", wx.NewUploadForm("img", "test.jpg", nil)).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"items": [
			{
				"text": "腾讯",
				"pos": {
					"left_top": {
						"x": 575,
						"y": 519
					},
					"right_top": {
						"x": 744,
						"y": 519
					},
					"right_bottom": {
						"x": 744,
						"y": 532
					},
					"left_bottom": {
						"x": 573,
						"y": 532
					}
				}
			},
			{
				"text": "微信团队",
				"pos": {
					"left_top": {
						"x": 670,
						"y": 516
					},
					"right_top": {
						"x": 762,
						"y": 517
					},
					"right_bottom": {
						"x": 762,
						"y": 532
					},
					"left_bottom": {
						"x": 670,
						"y": 531
					}
				}
			}
		],
		"img_size": {
			"w": 1280,
			"h": 720
		}
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(PrintedText)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", OCRPrintedText(dest, OCRPhoto, "test.jpg"))

	assert.Nil(t, err)
	assert.Equal(t, &PrintedText{
		Items: []*PrintedTextItem{
			{
				Text: "腾讯",
				Pos: ImagePosition{
					LeftTop: Position{
						X: 575,
						Y: 519,
					},
					RightTop: Position{
						X: 744,
						Y: 519,
					},
					RightBottom: Position{
						X: 744,
						Y: 532,
					},
					LeftBottom: Position{
						X: 573,
						Y: 532,
					},
				},
			},
			{
				Text: "微信团队",
				Pos: ImagePosition{
					LeftTop: Position{
						X: 670,
						Y: 516,
					},
					RightTop: Position{
						X: 762,
						Y: 517,
					},
					RightBottom: Position{
						X: 762,
						Y: 532,
					},
					LeftBottom: Position{
						X: 670,
						Y: 531,
					},
				},
			},
		},
		ImgSize: ImageSize{
			W: 1280,
			H: 720,
		},
	}, dest)
}

func TestOCRPrintedTextByURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cv/ocr/comm?access_token=ACCESS_TOKEN&img_url=ENCODE_URL&type=photo", nil).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"items": [
			{
				"text": "腾讯",
				"pos": {
					"left_top": {
						"x": 575,
						"y": 519
					},
					"right_top": {
						"x": 744,
						"y": 519
					},
					"right_bottom": {
						"x": 744,
						"y": 532
					},
					"left_bottom": {
						"x": 573,
						"y": 532
					}
				}
			},
			{
				"text": "微信团队",
				"pos": {
					"left_top": {
						"x": 670,
						"y": 516
					},
					"right_top": {
						"x": 762,
						"y": 517
					},
					"right_bottom": {
						"x": 762,
						"y": 532
					},
					"left_bottom": {
						"x": 670,
						"y": 531
					}
				}
			}
		],
		"img_size": {
			"w": 1280,
			"h": 720
		}
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(PrintedText)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", OCRPrintedTextByURL(dest, OCRPhoto, "ENCODE_URL"))

	assert.Nil(t, err)
	assert.Equal(t, &PrintedText{
		Items: []*PrintedTextItem{
			{
				Text: "腾讯",
				Pos: ImagePosition{
					LeftTop: Position{
						X: 575,
						Y: 519,
					},
					RightTop: Position{
						X: 744,
						Y: 519,
					},
					RightBottom: Position{
						X: 744,
						Y: 532,
					},
					LeftBottom: Position{
						X: 573,
						Y: 532,
					},
				},
			},
			{
				Text: "微信团队",
				Pos: ImagePosition{
					LeftTop: Position{
						X: 670,
						Y: 516,
					},
					RightTop: Position{
						X: 762,
						Y: 517,
					},
					RightBottom: Position{
						X: 762,
						Y: 532,
					},
					LeftBottom: Position{
						X: 670,
						Y: 531,
					},
				},
			},
		},
		ImgSize: ImageSize{
			W: 1280,
			H: 720,
		},
	}, dest)
}
