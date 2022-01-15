package agent

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/corp"
	"github.com/shenghui0779/gochat/mock"
	"github.com/shenghui0779/gochat/wx"
	"github.com/stretchr/testify/assert"
)

func TestCreateMenu(t *testing.T) {
	body := []byte(`{"button":[{"name":"扫码","sub_button":[{"type":"scancode_waitmsg","name":"扫码带提示","key":"rselfmenu_0_0"},{"type":"scancode_push","name":"扫码推事件","key":"rselfmenu_0_1"},{"type":"view_miniprogram","name":"小程序","pagepath":"pages/lunar/index","appid":"wx4389ji4kAAA"}]},{"name":"发图","sub_button":[{"type":"pic_sysphoto","name":"系统拍照发图","key":"rselfmenu_1_0"},{"type":"pic_photo_or_album","name":"拍照或者相册发图","key":"rselfmenu_1_1"},{"type":"pic_weixin","name":"微信相册发图","key":"rselfmenu_1_2"}]},{"type":"location_select","name":"发送位置","key":"rselfmenu_2_0"}]}`)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://qyapi.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN&agentid=1000005", body).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	params := &ParamsMenuCreate{
		Button: []*Button{
			GroupButton("扫码",
				ScanCodeWaitMsgButton("扫码带提示", "rselfmenu_0_0"),
				ScanCodePushButton("扫码推事件", "rselfmenu_0_1"),
				ViewMinipButton("小程序", "wx4389ji4kAAA", "pages/lunar/index"),
			),
			GroupButton("发图",
				PicSysPhotoButton("系统拍照发图", "rselfmenu_1_0"),
				PicPhotoOrAlbumButton("拍照或者相册发图", "rselfmenu_1_1"),
				PicWeixinButton("微信相册发图", "rselfmenu_1_2"),
			),
			LocationSelectButton("发送位置", "rselfmenu_2_0"),
		},
	}

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", CreateMenu(1000005, params))

	assert.Nil(t, err)
}

func TestGetMenu(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok",
	"button": [
		{
			"name": "扫码",
			"sub_button": [
				{
					"type": "scancode_waitmsg",
					"name": "扫码带提示",
					"key": "rselfmenu_0_0"
				},
				{
					"type": "scancode_push",
					"name": "扫码推事件",
					"key": "rselfmenu_0_1"
				},
				{
					"type": "view_miniprogram",
					"name": "小程序",
					"pagepath": "pages/lunar/index",
					"appid": "wx4389ji4kAAA"
				}
			]
		},
		{
			"name": "发图",
			"sub_button": [
				{
					"type": "pic_sysphoto",
					"name": "系统拍照发图",
					"key": "rselfmenu_1_0"
				},
				{
					"type": "pic_photo_or_album",
					"name": "拍照或者相册发图",
					"key": "rselfmenu_1_1"
				},
				{
					"type": "pic_weixin",
					"name": "微信相册发图",
					"key": "rselfmenu_1_2"
				}
			]
		},
		{
			"type": "location_select",
			"name": "发送位置",
			"key": "rselfmenu_2_0"
		}
	]
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/menu/get?access_token=ACCESS_TOKEN&agentid=1000005", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	result := new(ResultMenuGet)

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", GetMenu(1000005, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultMenuGet{
		Button: []*Button{
			{
				Name: "扫码",
				SubButton: []*Button{
					{
						Type: "scancode_waitmsg",
						Name: "扫码带提示",
						Key:  "rselfmenu_0_0",
					},
					{
						Type: "scancode_push",
						Name: "扫码推事件",
						Key:  "rselfmenu_0_1",
					},
					{
						Type:     "view_miniprogram",
						Name:     "小程序",
						PagePath: "pages/lunar/index",
						AppID:    "wx4389ji4kAAA",
					},
				},
			},
			{
				Name: "发图",
				SubButton: []*Button{
					{
						Type: "pic_sysphoto",
						Name: "系统拍照发图",
						Key:  "rselfmenu_1_0",
					},
					{
						Type: "pic_photo_or_album",
						Name: "拍照或者相册发图",
						Key:  "rselfmenu_1_1",
					},
					{
						Type: "pic_weixin",
						Name: "微信相册发图",
						Key:  "rselfmenu_1_2",
					},
				},
			},
			{
				Type: "location_select",
				Name: "发送位置",
				Key:  "rselfmenu_2_0",
			},
		},
	}, result)
}

func TestDeleteMenu(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"errcode": 0,
	"errmsg": "ok"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://qyapi.weixin.qq.com/cgi-bin/menu/delete?access_token=ACCESS_TOKEN&agentid=1000005", nil).Return(resp, nil)

	cp := corp.New("CORPID")
	cp.SetClient(wx.WithHTTPClient(client))

	err := cp.Do(context.TODO(), "ACCESS_TOKEN", DeleteMenu(1000005))

	assert.Nil(t, err)
}
