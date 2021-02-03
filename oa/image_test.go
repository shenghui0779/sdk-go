package oa

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/wx"
	"github.com/stretchr/testify/assert"
)

func TestAICrop(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cv/img/aicrop?access_token=ACCESS_TOKEN", wx.NewUploadForm("img", "test.jpg", nil)).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"results": [
			{
				"crop_left": 112,
				"crop_top": 0,
				"crop_right": 839,
				"crop_bottom": 727
			},
			{
				"crop_left": 0,
				"crop_top": 205,
				"crop_right": 965,
				"crop_bottom": 615
			}
		],
		"img_size": {
			"w": 966,
			"h": 728
		}
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(AICropResult)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", AICrop(dest, "test.jpg"))

	assert.Nil(t, err)
	assert.Equal(t, &AICropResult{
		Results: []*CropPosition{
			{
				CropLeft:   112,
				CropTop:    0,
				CropRight:  839,
				CropBottom: 727,
			},
			{
				CropLeft:   0,
				CropTop:    205,
				CropRight:  965,
				CropBottom: 615,
			},
		},
		ImgSize: ImageSize{
			W: 966,
			H: 728,
		},
	}, dest)
}

func TestAICropByURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cv/img/aicrop?access_token=ACCESS_TOKEN&img_url=ENCODE_URL", nil).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"results": [
			{
				"crop_left": 112,
				"crop_top": 0,
				"crop_right": 839,
				"crop_bottom": 727
			},
			{
				"crop_left": 0,
				"crop_top": 205,
				"crop_right": 965,
				"crop_bottom": 615
			}
		],
		"img_size": {
			"w": 966,
			"h": 728
		}
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(AICropResult)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", AICropByURL(dest, "ENCODE_URL"))

	assert.Nil(t, err)
	assert.Equal(t, &AICropResult{
		Results: []*CropPosition{
			{
				CropLeft:   112,
				CropTop:    0,
				CropRight:  839,
				CropBottom: 727,
			},
			{
				CropLeft:   0,
				CropTop:    205,
				CropRight:  965,
				CropBottom: 615,
			},
		},
		ImgSize: ImageSize{
			W: 966,
			H: 728,
		},
	}, dest)
}

func TestScanQRCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cv/img/qrcode?access_token=ACCESS_TOKEN", wx.NewUploadForm("img", "test.jpg", nil)).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"code_results": [
			{
				"type_name": "QR_CODE",
				"data": "http://www.qq.com",
				"pos": {
					"left_top": {
						"x": 585,
						"y": 378
					},
					"right_top": {
						"x": 828,
						"y": 378
					},
					"right_bottom": {
						"x": 828,
						"y": 618
					},
					"left_bottom": {
						"x": 585,
						"y": 618
					}
				}
			},
			{
				"type_name": "QR_CODE",
				"data": "https://mp.weixin.qq.com",
				"pos": {
					"left_top": {
						"x": 185,
						"y": 142
					},
					"right_top": {
						"x": 396,
						"y": 142
					},
					"right_bottom": {
						"x": 396,
						"y": 353
					},
					"left_bottom": {
						"x": 185,
						"y": 353
					}
				}
			},
			{
				"type_name": "EAN_13",
				"data": "5906789678957"
			},
			{
				"type_name": "CODE_128",
				"data": "50090500019191"
			}
		],
		"img_size": {
			"w": 1000,
			"h": 900
		}
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(QRCodeScanResult)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", ScanQRCode(dest, "test.jpg"))

	assert.Nil(t, err)
	assert.Equal(t, &QRCodeScanResult{
		CodeResults: []*QRCodeScanData{
			{
				TypeName: "QR_CODE",
				Data:     "http://www.qq.com",
				Pos: ImagePosition{
					LeftTop: Position{
						X: 585,
						Y: 378,
					},
					RightTop: Position{
						X: 828,
						Y: 378,
					},
					RightBottom: Position{
						X: 828,
						Y: 618,
					},
					LeftBottom: Position{
						X: 585,
						Y: 618,
					},
				},
			},
			{
				TypeName: "QR_CODE",
				Data:     "https://mp.weixin.qq.com",
				Pos: ImagePosition{
					LeftTop: Position{
						X: 185,
						Y: 142,
					},
					RightTop: Position{
						X: 396,
						Y: 142,
					},
					RightBottom: Position{
						X: 396,
						Y: 353,
					},
					LeftBottom: Position{
						X: 185,
						Y: 353,
					},
				},
			},
			{
				TypeName: "EAN_13",
				Data:     "5906789678957",
				Pos: ImagePosition{
					LeftTop: Position{
						X: 0,
						Y: 0,
					},
					RightTop: Position{
						X: 0,
						Y: 0,
					},
					RightBottom: Position{
						X: 0,
						Y: 0,
					},
					LeftBottom: Position{
						X: 0,
						Y: 0,
					},
				},
			},
			{
				TypeName: "CODE_128",
				Data:     "50090500019191",
				Pos: ImagePosition{
					LeftTop: Position{
						X: 0,
						Y: 0,
					},
					RightTop: Position{
						X: 0,
						Y: 0,
					},
					RightBottom: Position{
						X: 0,
						Y: 0,
					},
					LeftBottom: Position{
						X: 0,
						Y: 0,
					},
				},
			},
		},
		ImgSize: ImageSize{
			W: 1000,
			H: 900,
		},
	}, dest)
}

func TestScanQRCodeByURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cv/img/qrcode?access_token=ACCESS_TOKEN&img_url=ENCODE_URL", nil).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"code_results": [
			{
				"type_name": "QR_CODE",
				"data": "http://www.qq.com",
				"pos": {
					"left_top": {
						"x": 585,
						"y": 378
					},
					"right_top": {
						"x": 828,
						"y": 378
					},
					"right_bottom": {
						"x": 828,
						"y": 618
					},
					"left_bottom": {
						"x": 585,
						"y": 618
					}
				}
			},
			{
				"type_name": "QR_CODE",
				"data": "https://mp.weixin.qq.com",
				"pos": {
					"left_top": {
						"x": 185,
						"y": 142
					},
					"right_top": {
						"x": 396,
						"y": 142
					},
					"right_bottom": {
						"x": 396,
						"y": 353
					},
					"left_bottom": {
						"x": 185,
						"y": 353
					}
				}
			},
			{
				"type_name": "EAN_13",
				"data": "5906789678957"
			},
			{
				"type_name": "CODE_128",
				"data": "50090500019191"
			}
		],
		"img_size": {
			"w": 1000,
			"h": 900
		}
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(QRCodeScanResult)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", ScanQRCodeByURL(dest, "ENCODE_URL"))

	assert.Nil(t, err)
	assert.Equal(t, &QRCodeScanResult{
		CodeResults: []*QRCodeScanData{
			{
				TypeName: "QR_CODE",
				Data:     "http://www.qq.com",
				Pos: ImagePosition{
					LeftTop: Position{
						X: 585,
						Y: 378,
					},
					RightTop: Position{
						X: 828,
						Y: 378,
					},
					RightBottom: Position{
						X: 828,
						Y: 618,
					},
					LeftBottom: Position{
						X: 585,
						Y: 618,
					},
				},
			},
			{
				TypeName: "QR_CODE",
				Data:     "https://mp.weixin.qq.com",
				Pos: ImagePosition{
					LeftTop: Position{
						X: 185,
						Y: 142,
					},
					RightTop: Position{
						X: 396,
						Y: 142,
					},
					RightBottom: Position{
						X: 396,
						Y: 353,
					},
					LeftBottom: Position{
						X: 185,
						Y: 353,
					},
				},
			},
			{
				TypeName: "EAN_13",
				Data:     "5906789678957",
				Pos: ImagePosition{
					LeftTop: Position{
						X: 0,
						Y: 0,
					},
					RightTop: Position{
						X: 0,
						Y: 0,
					},
					RightBottom: Position{
						X: 0,
						Y: 0,
					},
					LeftBottom: Position{
						X: 0,
						Y: 0,
					},
				},
			},
			{
				TypeName: "CODE_128",
				Data:     "50090500019191",
				Pos: ImagePosition{
					LeftTop: Position{
						X: 0,
						Y: 0,
					},
					RightTop: Position{
						X: 0,
						Y: 0,
					},
					RightBottom: Position{
						X: 0,
						Y: 0,
					},
					LeftBottom: Position{
						X: 0,
						Y: 0,
					},
				},
			},
		},
		ImgSize: ImageSize{
			W: 1000,
			H: 900,
		},
	}, dest)
}

func TestSuperreSolution(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cv/img/superresolution?access_token=ACCESS_TOKEN", wx.NewUploadForm("img", "test.jpg", nil)).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"media_id": "6WXsIXkG7lXuDLspD9xfm5dsvHzb0EFl0li6ySxi92ap8Vl3zZoD9DpOyNudeJGB"
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(SuperreSolutionResult)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SuperreSolution(dest, "test.jpg"))

	assert.Nil(t, err)
	assert.Equal(t, &SuperreSolutionResult{
		MediaID: "6WXsIXkG7lXuDLspD9xfm5dsvHzb0EFl0li6ySxi92ap8Vl3zZoD9DpOyNudeJGB",
	}, dest)
}

func TestSuperreSolutionByURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockHTTPClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cv/img/superresolution?access_token=ACCESS_TOKEN&img_url=ENCODE_URL", nil).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"media_id": "6WXsIXkG7lXuDLspD9xfm5dsvHzb0EFl0li6ySxi92ap8Vl3zZoD9DpOyNudeJGB"
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(SuperreSolutionResult)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SuperreSolutionByURL(dest, "ENCODE_URL"))

	assert.Nil(t, err)
	assert.Equal(t, &SuperreSolutionResult{
		MediaID: "6WXsIXkG7lXuDLspD9xfm5dsvHzb0EFl0li6ySxi92ap8Vl3zZoD9DpOyNudeJGB",
	}, dest)
}
