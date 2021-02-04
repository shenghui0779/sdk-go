package mp

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/wx"
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
func ApplyPlugin(pluginAppID, reason string) wx.Action {
	return wx.NewAction(PluginManageURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(wx.X{
				"action":       PluginApply,
				"plugin_appid": pluginAppID,
				"reason":       reason,
			})
		}),
	)
}

// PluginDevApplyInfo 插件使用方信息
type PluginDevApplyInfo struct {
	AppID      string `json:"appid"`       // 使用者的appid
	Status     int    `json:"status"`      // 插件状态
	Nickname   string `json:"nickname"`    // 使用者的昵称
	HeadImgURL string `json:"headimgurl"`  // 使用者的头像
	Categories []wx.X `json:"categories"`  // 使用者的类目
	CreateTime string `json:"create_time"` // 使用者的申请时间
	ApplyURL   string `json:"apply_url"`   // 使用者的小程序码
	Reason     string `json:"reason"`      // 使用者的申请说明
}

// GetPluginDevApplyList 获取当前所有插件使用方（供插件开发者调用）
func GetPluginDevApplyList(dest *[]*PluginDevApplyInfo, page, num int) wx.Action {
	return wx.NewAction(PluginDevManageURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(wx.X{
				"action": PluginDevApplyList,
				"page":   page,
				"num":    num,
			})
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal([]byte(gjson.GetBytes(resp, "apply_list").Raw), dest)
		}),
	)
}

// PluginInfo 插件信息
type PluginInfo struct {
	AppID      string `json:"appid"`      // 插件 appId
	Status     int    `json:"status"`     // 插件状态
	Nickname   string `json:"nickname"`   // 插件昵称
	HeadImgURL string `json:"headimgurl"` // 插件头像
}

// GetPluginList 查询已添加的插件
func GetPluginList(dest *[]*PluginInfo) wx.Action {
	return wx.NewAction(PluginManageURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(wx.X{"action": PluginList})
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal([]byte(gjson.GetBytes(resp, "plugin_list").Raw), dest)
		}),
	)
}

// SetDevPluginApplyStatus 修改插件使用申请的状态（供插件开发者调用）
func SetDevPluginApplyStatus(action PluginAction, appid, reason string) wx.Action {
	return wx.NewAction(PluginDevManageURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			params := wx.X{"action": action}

			if len(appid) != 0 {
				params["appid"] = appid
			}

			if len(reason) != 0 {
				params["reason"] = reason
			}

			return json.Marshal(params)
		}),
	)
}

// UnbindPlugin 删除已添加的插件
func UnbindPlugin(pluginAppID string) wx.Action {
	return wx.NewAction(PluginManageURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(wx.X{
				"action":       PluginUnbind,
				"plugin_appid": pluginAppID,
			})
		}),
	)
}
