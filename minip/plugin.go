package minip

import (
	"encoding/json"

	"github.com/shenghui0779/yiigo"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
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

type ParamsPluginApply struct {
	Action      PluginAction `json:"action"`
	PluginAppID string       `json:"plugin_appid"`
	Reason      string       `json:"reason"`
}

// ApplyPlugin 向插件开发者发起使用插件的申请
func ApplyPlugin(pluginAppID, reason string) wx.Action {
	params := &ParamsPluginApply{
		Action:      PluginApply,
		PluginAppID: pluginAppID,
		Reason:      reason,
	}

	return wx.NewPostAction(urls.MinipPluginManage,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

// PluginDevApplyInfo 插件使用方信息
type PluginDevApplyInfo struct {
	AppID      string    `json:"appid"`       // 使用者的appid
	Status     int       `json:"status"`      // 插件状态
	Nickname   string    `json:"nickname"`    // 使用者的昵称
	HeadImgURL string    `json:"headimgurl"`  // 使用者的头像
	Categories []yiigo.X `json:"categories"`  // 使用者的类目
	CreateTime string    `json:"create_time"` // 使用者的申请时间
	ApplyURL   string    `json:"apply_url"`   // 使用者的小程序码
	Reason     string    `json:"reason"`      // 使用者的申请说明
}

type ParamsPluginDevApplyList struct {
	Action PluginAction `json:"action"`
	Page   int          `json:"page"`
	Num    int          `json:"num"`
}

type ResultPluginDevApplyList struct {
	ApplyList []*PluginDevApplyInfo `json:"apply_list"`
}

// GetPluginDevApplyList 获取当前所有插件使用方（供插件开发者调用）
func GetPluginDevApplyList(page, num int, result *ResultPluginDevApplyList) wx.Action {
	params := &ParamsPluginDevApplyList{
		Action: PluginDevApplyList,
		Page:   page,
		Num:    num,
	}

	return wx.NewPostAction(urls.MinipPluginDevManage,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
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

type ParamsPluginList struct {
	Action PluginAction `json:"action"`
}

type ResultPluginList struct {
	PluginList []*PluginInfo `json:"plugin_list"`
}

// GetPluginList 查询已添加的插件
func GetPluginList(result *ResultPluginList) wx.Action {
	params := &ParamsPluginList{Action: PluginList}

	return wx.NewPostAction(urls.MinipPluginManage,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsDevPluginApplyStatus struct {
	Action PluginAction `json:"action"`
	AppID  string       `json:"appid,omitempty"`
	Reason string       `json:"reason,omitempty"`
}

// SetDevPluginApplyStatus 修改插件使用申请的状态（供插件开发者调用）
func SetDevPluginApplyStatus(params *ParamsDevPluginApplyStatus) wx.Action {
	return wx.NewPostAction(urls.MinipPluginDevManage,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsPluginUnbind struct {
	Action      PluginAction `json:"action"`
	PluginAppID string       `json:"plugin_appid"`
}

// UnbindPlugin 删除已添加的插件
func UnbindPlugin(pluginAppID string) wx.Action {
	params := &ParamsPluginUnbind{
		Action:      PluginUnbind,
		PluginAppID: pluginAppID,
	}

	return wx.NewPostAction(urls.MinipPluginManage,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}
