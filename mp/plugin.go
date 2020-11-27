package mp

import (
	"encoding/json"
	"fmt"

	"github.com/shenghui0779/gochat/helpers"
	"github.com/tidwall/gjson"
)

// PluginAction defines for plugin action params
type PluginAction string

var (
	// PluginApply 向插件开发者发起使用插件的申请
	PluginApply PluginAction = "apply"
	// PluginDevApplyList 获取当前所有插件使用方（供插件开发者调用）
	PluginDevApplyList PluginAction = "dev_apply_list"
	// PluginList 查询已添加的插件
	PluginList = "list"
	// PluginDevAgree 同意申请
	PluginDevAgree = "dev_agree"
	// PluginDevRefuse 拒绝申请
	PluginDevRefuse = "dev_refuse"
	// PluginDevDelete 删除已拒绝的申请者
	PluginDevDelete = "dev_delete"
	// PluginUnbind 删除已添加的插件
	PluginUnbind = "unbind"
)

// ApplyPlugin 向插件开发者发起使用插件的申请
func ApplyPlugin(pluginAppID, reason string) Action {
	return &WechatAPI{
		body: helpers.NewPostBody(func() ([]byte, error) {
			return json.Marshal(helpers.X{
				"action":       PluginApply,
				"plugin_appid": pluginAppID,
				"reason":       reason,
			})
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("POST|%s?access_token=%s", PluginManageURL, accessToken)
		},
	}
}

// PluginDevApplyInfo 插件使用方信息
type PluginDevApplyInfo struct {
	AppID      string      `json:"appid"`
	Status     int         `json:"status"`
	Nickname   string      `json:"nickname"`
	HeadImgURL string      `json:"headimgurl"`
	Categories []helpers.X `json:"categories"`
	CreateTime string      `json:"create_time"`
	ApplyURL   string      `json:"apply_url"`
	Reason     string      `json:"reason"`
}

// GetPluginDevApplyList 获取当前所有插件使用方（供插件开发者调用）
func GetPluginDevApplyList(page, num int, dest *[]PluginDevApplyInfo) Action {
	return &WechatAPI{
		body: helpers.NewPostBody(func() ([]byte, error) {
			return json.Marshal(helpers.X{
				"action": PluginDevApplyList,
				"page":   page,
				"num":    num,
			})
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("POST|%s?access_token=%s", PluginDevManageURL, accessToken)
		},
		decode: func(resp []byte) error {
			r := gjson.GetBytes(resp, "apply_list")

			return json.Unmarshal([]byte(r.Raw), dest)
		},
	}
}

// PluginInfo 插件信息
type PluginInfo struct {
	AppID      string `json:"appid"`
	Status     int    `json:"status"`
	Nickname   string `json:"nickname"`
	HeadImgURL string `json:"headimgurl"`
}

// GetPluginList 查询已添加的插件
func GetPluginList(dest *[]PluginInfo) Action {
	return &WechatAPI{
		body: helpers.NewPostBody(func() ([]byte, error) {
			return json.Marshal(helpers.X{
				"action": PluginList,
			})
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("POST|%s?access_token=%s", PluginManageURL, accessToken)
		},
		decode: func(resp []byte) error {
			r := gjson.GetBytes(resp, "plugin_list")

			return json.Unmarshal([]byte(r.Raw), dest)
		},
	}
}

// SetDevPluginApplyStatus 修改插件使用申请的状态（供插件开发者调用）
func SetDevPluginApplyStatus(action PluginAction, appid, reason string) Action {
	return &WechatAPI{
		body: helpers.NewPostBody(func() ([]byte, error) {
			return json.Marshal(helpers.X{
				"action": action,
				"appid":  appid,
				"reason": reason,
			})
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("POST|%s?access_token=%s", PluginDevManageURL, accessToken)
		},
	}
}

// UnbindPlugin 删除已添加的插件
func UnbindPlugin(pluginAppID string) Action {
	return &WechatAPI{
		body: helpers.NewPostBody(func() ([]byte, error) {
			return json.Marshal(helpers.X{
				"action":       PluginUnbind,
				"plugin_appid": pluginAppID,
			})
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("POST|%s?access_token=%s", PluginManageURL, accessToken)
		},
	}
}
