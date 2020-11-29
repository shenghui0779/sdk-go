package mp

import (
	"encoding/json"
	"net/url"

	"github.com/shenghui0779/gochat/public"
	"github.com/tidwall/gjson"
)

// PluginAction defines for plugin action params
type PluginAction string

// 微信支持的插件行为
var (
	PluginApply        PluginAction = "apply"          // 向插件开发者发起使用插件的申请
	PluginDevApplyList PluginAction = "dev_apply_list" // 获取当前所有插件使用方（供插件开发者调用）
	PluginList         PluginAction = "list"           // 查询已添加的插件
	PluginDevAgree     PluginAction = "dev_agree"      // 同意申请
	PluginDevRefuse    PluginAction = "dev_refuse"     // 拒绝申请
	PluginDevDelete    PluginAction = "dev_delete"     // 删除已拒绝的申请者
	PluginUnbind       PluginAction = "unbind"         // 删除已添加的插件
)

// ApplyPlugin 向插件开发者发起使用插件的申请
func ApplyPlugin(pluginAppID, reason string) public.Action {
	return public.NewOpenPostAPI(PluginManageURL, url.Values{}, public.NewPostBody(func() ([]byte, error) {
		return json.Marshal(public.X{
			"action":       PluginApply,
			"plugin_appid": pluginAppID,
			"reason":       reason,
		})
	}), nil)
}

// PluginDevApplyInfo 插件使用方信息
type PluginDevApplyInfo struct {
	AppID      string     `json:"appid"`
	Status     int        `json:"status"`
	Nickname   string     `json:"nickname"`
	HeadImgURL string     `json:"headimgurl"`
	Categories []public.X `json:"categories"`
	CreateTime string     `json:"create_time"`
	ApplyURL   string     `json:"apply_url"`
	Reason     string     `json:"reason"`
}

// GetPluginDevApplyList 获取当前所有插件使用方（供插件开发者调用）
func GetPluginDevApplyList(page, num int, dest *[]PluginDevApplyInfo) public.Action {
	return public.NewOpenPostAPI(PluginDevManageURL, url.Values{}, public.NewPostBody(func() ([]byte, error) {
		return json.Marshal(public.X{
			"action": PluginDevApplyList,
			"page":   page,
			"num":    num,
		})
	}), func(resp []byte) error {
		r := gjson.GetBytes(resp, "apply_list")

		return json.Unmarshal([]byte(r.Raw), dest)
	})
}

// PluginInfo 插件信息
type PluginInfo struct {
	AppID      string `json:"appid"`
	Status     int    `json:"status"`
	Nickname   string `json:"nickname"`
	HeadImgURL string `json:"headimgurl"`
}

// GetPluginList 查询已添加的插件
func GetPluginList(dest *[]PluginInfo) public.Action {
	return public.NewOpenPostAPI(PluginManageURL, url.Values{}, public.NewPostBody(func() ([]byte, error) {
		return json.Marshal(public.X{
			"action": PluginList,
		})
	}), func(resp []byte) error {
		r := gjson.GetBytes(resp, "plugin_list")

		return json.Unmarshal([]byte(r.Raw), dest)
	})
}

// SetDevPluginApplyStatus 修改插件使用申请的状态（供插件开发者调用）
func SetDevPluginApplyStatus(action PluginAction, appid, reason string) public.Action {
	return public.NewOpenPostAPI(PluginDevManageURL, url.Values{}, public.NewPostBody(func() ([]byte, error) {
		return json.Marshal(public.X{
			"action": action,
			"appid":  appid,
			"reason": reason,
		})
	}), nil)
}

// UnbindPlugin 删除已添加的插件
func UnbindPlugin(pluginAppID string) public.Action {
	return public.NewOpenPostAPI(PluginManageURL, url.Values{}, public.NewPostBody(func() ([]byte, error) {
		return json.Marshal(public.X{
			"action":       PluginUnbind,
			"plugin_appid": pluginAppID,
		})
	}), nil)
}
