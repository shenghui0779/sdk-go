package minip

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/yiigo"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/wx"
)

func TestApplyPlugin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/plugin?access_token=ACCESS_TOKEN", []byte(`{"action":"apply","plugin_appid":"aaaa","reason":"hello"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", ApplyPlugin("aaaa", "hello"))

	assert.Nil(t, err)
}

func TestGetPluginDevApplyList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/devplugin?access_token=ACCESS_TOKEN", []byte(`{"action":"dev_apply_list","page":1,"num":10}`)).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"apply_list": [{
			"appid": "xxxxxxxxxxxxx",
			"status": 1,
			"nickname": "名称",
			"headimgurl": "**********",
			"reason": "polo has gone",
			"apply_url": "*******",
			"create_time": "1536305096",
			"categories": [{
				"first": "IT科技",
				"second": "硬件与设备"
			}]
		}]
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	result := new(ResultPluginDevApplyList)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetPluginDevApplyList(1, 10, result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultPluginDevApplyList{
		ApplyList: []*PluginDevApplyInfo{
			{
				AppID:      "xxxxxxxxxxxxx",
				Status:     1,
				Nickname:   "名称",
				HeadImgURL: "**********",
				Categories: []yiigo.X{
					{
						"first":  "IT科技",
						"second": "硬件与设备",
					},
				},
				CreateTime: "1536305096",
				ApplyURL:   "*******",
				Reason:     "polo has gone",
			},
		},
	}, result)
}

func TestGetPluginList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/plugin?access_token=ACCESS_TOKEN", []byte(`{"action":"list"}`)).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"plugin_list": [{
			"appid": "aaaa",
			"status": 1,
			"nickname": "插件昵称",
			"headimgurl": "http://plugin.qq.com"
		}]
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	result := new(ResultPluginList)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetPluginList(result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultPluginList{
		PluginList: []*PluginInfo{
			{
				AppID:      "aaaa",
				Status:     1,
				Nickname:   "插件昵称",
				HeadImgURL: "http://plugin.qq.com",
			},
		},
	}, result)
}

func TestSetDevPluginApplyStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/devplugin?access_token=ACCESS_TOKEN", []byte(`{"action":"dev_agree","appid":"APPID"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	params := &ParamsDevPluginApplyStatus{
		Action: PluginDevAgree,
		AppID:  "APPID",
	}

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SetDevPluginApplyStatus(params))

	assert.Nil(t, err)
}

func TestUnbindPlugin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/plugin?access_token=ACCESS_TOKEN", []byte(`{"action":"unbind","plugin_appid":"APPID"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UnbindPlugin("APPID"))

	assert.Nil(t, err)
}
