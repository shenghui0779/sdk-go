package oa

import (
	"encoding/json"

	"github.com/tidwall/gjson"

	"github.com/shenghui0779/gochat/wx"
)

// MaxSubscriberListCount 关注列表的最大数目
const MaxSubscriberListCount = 10000

// SubscribeScene 关注的渠道来源
type SubscribeScene string

// 微信支持的关注的渠道来源
const (
	SceneSearch           SubscribeScene = "ADD_SCENE_SEARCH"               // 公众号搜索
	SceneQRCode           SubscribeScene = "ADD_SCENE_QR_CODE"              // 扫描二维码
	SceneAccountMigration SubscribeScene = "ADD_SCENE_ACCOUNT_MIGRATION"    // 公众号迁移
	SceneProfileCard      SubscribeScene = "ADD_SCENE_PROFILE_CARD"         // 名片分享
	SceneProfileLink      SubscribeScene = "ADD_SCENE_PROFILE_LINK"         // 图文页内名称点击
	SceneProfileItem      SubscribeScene = "ADD_SCENE_PROFILE_ITEM"         // 图文页右上角菜单
	ScenePaid             SubscribeScene = "ADD_SCENE_PAID"                 // 支付后关注
	SceneWechatAD         SubscribeScene = "ADD_SCENE_WECHAT_ADVERTISEMENT" // 微信广告
	SceneOthers           SubscribeScene = "ADD_SCENE_OTHERS"               // 其他
)

// SubscriberInfo 关注用户信息
type SubscriberInfo struct {
	Subscribe      int            `json:"subscribe"`       // 用户是否订阅该公众号标识，值为0时，代表此用户没有关注该公众号，拉取不到其余信息。
	OpenID         string         `json:"openid"`          // 用户的标识，对当前公众号唯一
	NickName       string         `json:"nickname"`        // 用户的昵称
	Sex            int            `json:"sex"`             // 用户的性别，男（1），女（2），未知（0）
	Country        string         `json:"country"`         // 用户所在国家
	City           string         `json:"city"`            // 用户所在城市
	Province       string         `json:"province"`        // 用户所在省份
	Language       string         `json:"language"`        // 用户的语言，简体中文为zh_CN
	HeadImgURL     string         `json:"headimgurl"`      // 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	SubscribeTime  int64          `json:"subscribe_time"`  // 用户关注时间，为时间戳。如果用户曾多次关注，则取最后关注时间
	UnionID        string         `json:"unionid"`         // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
	Remark         string         `json:"remark"`          // 公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
	GroupID        int64          `json:"groupid"`         // 用户所在的分组ID（兼容旧的用户分组接口）
	TagidList      []int64        `json:"tagid_list"`      // 用户被打上的标签ID列表
	SubscribeScene SubscribeScene `json:"subscribe_scene"` // 用户关注的渠道来源
	QRScene        int64          `json:"qr_scene"`        // 二维码扫码场景（开发者自定义）
	QRSceneStr     string         `json:"qr_scene_str"`    // 二维码扫码场景描述（开发者自定义）
}

// SubscriberList 关注列表
type SubscriberList struct {
	Total      int                `json:"total"`
	Count      int                `json:"count"`
	Data       SubscriberListData `json:"data"`
	NextOpenID string             `json:"next_openid"`
}

// SubscriberListData 关注列表数据
type SubscriberListData struct {
	OpenID []string `json:"openid"`
}

// GetSubscriberInfo 获取关注用户信息
func GetSubscriberInfo(dest *SubscriberInfo, openid string) wx.Action {
	return wx.NewAction(SubscriberGetURL,
		wx.WithMethod(wx.MethodGet),
		wx.WithQuery("openid", openid),
		wx.WithQuery("lang", "zh_CN"),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// BatchGetSubscribers 批量关注用户信息
func BatchGetSubscribers(dest *[]*SubscriberInfo, openids ...string) wx.Action {
	return wx.NewAction(SubscriberBatchGetURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			userList := make([]map[string]string, 0, len(openids))

			for _, v := range openids {
				userList = append(userList, map[string]string{
					"openid": v,
					"lang":   "zh_CN",
				})
			}

			return json.Marshal(wx.X{"user_list": userList})
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal([]byte(gjson.GetBytes(resp, "user_info_list").Raw), dest)
		}),
	)
}

// GetSubscriberList 获取关注用户列表
func GetSubscriberList(dest *SubscriberList, nextOpenID ...string) wx.Action {
	return wx.NewAction(SubscriberListURL,
		wx.WithMethod(wx.MethodGet),
		wx.WithQuery("next_openid", nextOpenID[0]),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// GetBlackList 获取用户黑名单列表
func GetBlackList(dest *SubscriberList, beginOpenID ...string) wx.Action {
	return wx.NewAction(BlackListGetURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			params := wx.X{
				"begin_openid": "",
			}

			if len(beginOpenID) != 0 {
				params["begin_openid"] = beginOpenID[0]
			}

			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// BlackSubscribers 拉黑用户
func BlackSubscribers(openids ...string) wx.Action {
	return wx.NewAction(BatchBlackListURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(wx.X{"openid_list": openids})
		}),
	)
}

// UnBlackSubscriber 取消拉黑用户
func UnBlackSubscribers(openids ...string) wx.Action {
	return wx.NewAction(BatchUnBlackListURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(wx.X{"openid_list": openids})
		}),
	)
}

// SetUserRemark 设置用户备注名（该接口暂时开放给微信认证的服务号）
func SetUserRemark(openid, remark string) wx.Action {
	return wx.NewAction(UserRemarkSetURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(wx.X{
				"openid": openid,
				"remark": remark,
			})
		}),
	)
}
