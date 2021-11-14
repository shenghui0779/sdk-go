package offia

import (
	"encoding/json"

	"github.com/shenghui0779/yiigo"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

// MaxUserListCount 关注列表的最大数目
const MaxUserListCount = 10000

// SubscribeScene 关注的渠道来源
type SubscribeScene string

// 微信支持的关注的渠道来源
const (
	SubscribeSearch           SubscribeScene = "ADD_SCENE_SEARCH"               // 公众号搜索
	SubscribeQRCode           SubscribeScene = "ADD_SCENE_QR_CODE"              // 扫描二维码
	SubscribeAccountMigration SubscribeScene = "ADD_SCENE_ACCOUNT_MIGRATION"    // 公众号迁移
	SubscribeProfileCard      SubscribeScene = "ADD_SCENE_PROFILE_CARD"         // 名片分享
	SubscribeProfileLink      SubscribeScene = "ADD_SCENE_PROFILE_LINK"         // 图文页内名称点击
	SubscribeProfileItem      SubscribeScene = "ADD_SCENE_PROFILE_ITEM"         // 图文页右上角菜单
	SubscribePaid             SubscribeScene = "ADD_SCENE_PAID"                 // 支付后关注
	SubscribeWechatAD         SubscribeScene = "ADD_SCENE_WECHAT_ADVERTISEMENT" // 微信广告
	SubscribeOthers           SubscribeScene = "ADD_SCENE_OTHERS"               // 其他
)

// UserInfo 关注用户信息
type UserInfo struct {
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

// UserList 关注列表
type UserList struct {
	Total      int          `json:"total"`
	Count      int          `json:"count"`
	Data       UserListData `json:"data"`
	NextOpenID string       `json:"next_openid"`
}

// UserListData 关注列表数据
type UserListData struct {
	OpenID []string `json:"openid"`
}

type ParamsUserGet struct {
	OpenID string `json:"open_id"`
	Lang   string `json:"lang,omitempty"`
}

// GetUser 获取关注用户信息
func GetUser(params *ParamsUserGet, result *UserInfo) wx.Action {
	options := []wx.ActionOption{
		wx.WithQuery("openid", params.OpenID),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	}

	if len(params.Lang) != 0 {
		options = append(options, wx.WithQuery("lang", params.Lang))
	}

	return wx.NewGetAction(urls.OffiaUserGet, options...)
}

type ParamsUserBatchGet struct {
	UserList []*ParamsUserGet `json:"user_list"`
}

type ResultUserBatchGet struct {
	UserInfoList []*UserInfo `json:"user_info_list"`
}

// BatchGetUser 批量关注用户信息
func BatchGetUser(params *ParamsUserBatchGet, result *ResultUserBatchGet) wx.Action {
	return wx.NewPostAction(urls.OffiaUserBatchGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// GetUserList 获取关注用户列表
func GetUserList(dest *UserList, nextOpenID ...string) wx.Action {
	return wx.NewGetAction(urls.OffiaUserList,
		wx.WithQuery("next_openid", nextOpenID[0]),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}

// GetBlackList 获取用户黑名单列表
func GetBlackList(dest *UserList, beginOpenID ...string) wx.Action {
	return wx.NewPostAction(urls.OffiaBlackListGet,
		wx.WithBody(func() ([]byte, error) {
			params := yiigo.X{
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

// BlackUsers 拉黑用户
func BlackUsers(openids ...string) wx.Action {
	return wx.NewPostAction(urls.OffiaBatchBlackList,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(yiigo.X{"openid_list": openids})
		}),
	)
}

// UnBlackUser 取消拉黑用户
func UnBlackUsers(openids ...string) wx.Action {
	return wx.NewPostAction(urls.OffiaBatchUnBlackList,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(yiigo.X{"openid_list": openids})
		}),
	)
}

// SetUserRemark 设置用户备注名（该接口暂时开放给微信认证的服务号）
func SetUserRemark(openid, remark string) wx.Action {
	return wx.NewPostAction(urls.OffiaUserRemarkSet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(yiigo.X{
				"openid": openid,
				"remark": remark,
			})
		}),
	)
}
