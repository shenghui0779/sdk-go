package offia

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

// MaxUserListCount 关注列表的最大数目
const MaxUserListCount = 10000

// SubscribeScene 关注的渠道来源
type SubscribeScene string

// 微信支持的关注的渠道来源
const (
	AddSceneSearch           SubscribeScene = "ADD_SCENE_SEARCH"               // 公众号搜索
	AddSceneQRCode           SubscribeScene = "ADD_SCENE_QR_CODE"              // 扫描二维码
	AddSceneAccountMigration SubscribeScene = "ADD_SCENE_ACCOUNT_MIGRATION"    // 公众号迁移
	AddSceneProfileCard      SubscribeScene = "ADD_SCENE_PROFILE_CARD"         // 名片分享
	AddSceneProfileLink      SubscribeScene = "ADD_SCENE_PROFILE_LINK"         // 图文页内名称点击
	AddSceneProfileItem      SubscribeScene = "ADD_SCENE_PROFILE_ITEM"         // 图文页右上角菜单
	AddScenePaid             SubscribeScene = "ADD_SCENE_PAID"                 // 支付后关注
	AddSceneWechatAD         SubscribeScene = "ADD_SCENE_WECHAT_ADVERTISEMENT" // 微信广告
	AddSceneOthers           SubscribeScene = "ADD_SCENE_OTHERS"               // 其他
)

type Tag struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type ParamsTag struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ParamsTagOpt struct {
	Tag *ParamsTag `json:"tag"`
}

type ResultTagCreate struct {
	Tag *Tag `json:"tag"`
}

// CreateTag 用户管理 - 用户标签管理 - 创建标签
func CreateTag(name string, result *ResultTagCreate) wx.Action {
	params := &ParamsTagOpt{
		Tag: &ParamsTag{
			Name: name,
		},
	}

	return wx.NewPostAction(urls.OffiaTagCreate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// UpdateTag 用户管理 - 用户标签管理 - 编辑标签
func UpdateTag(id int64, name string) wx.Action {
	params := &ParamsTagOpt{
		Tag: &ParamsTag{
			ID:   id,
			Name: name,
		},
	}

	return wx.NewPostAction(urls.OffiaTagUpdate,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ResultTagsGet struct {
	Tags []*Tag `json:"tags"`
}

// GetTags 用户管理 - 用户标签管理 - 获取公众号已创建的标签
func GetTags(result *ResultTagsGet) wx.Action {
	return wx.NewGetAction(urls.OffiaTagGet,
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// DeleteTag 用户管理 - 用户标签管理 - 删除标签
func DeleteTag(id int64) wx.Action {
	params := &ParamsTagOpt{
		Tag: &ParamsTag{
			ID: id,
		},
	}

	return wx.NewPostAction(urls.OffiaTagDelete,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsTagUsers struct {
	TagID      int64  `json:"tagid"`
	NextOpenID string `json:"next_openid"`
}

type ResultTagUsers struct {
	Count      int          `json:"count"`
	Data       *TagUserData `json:"data"`
	NextOpenID string       `json:"next_openid"`
}

type TagUserData struct {
	OpenID []string `json:"openid"`
}

// GetTagUsers 用户管理 - 用户标签管理 - 获取标签下粉丝列表
func GetTagUsers(tagID int64, nextOpenID string, result *ResultTagUsers) wx.Action {
	params := &ParamsTagUsers{
		TagID:      tagID,
		NextOpenID: nextOpenID,
	}

	return wx.NewPostAction(urls.OffiaTagUserGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsBatchTagging struct {
	TagID      int64    `json:"tagid"`
	OpenIDList []string `json:"openid_list"`
}

// BatchTagging 用户管理 - 用户标签管理 - 批量为用户打标签
func BatchTagging(tagID int64, openids ...string) wx.Action {
	params := &ParamsBatchTagging{
		TagID:      tagID,
		OpenIDList: openids,
	}

	return wx.NewPostAction(urls.OffiaBatchTagging,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsBatchUnTagging struct {
	TagID      int64    `json:"tagid"`
	OpenIDList []string `json:"openid_list"`
}

// BatchUnTagging 用户管理 - 用户标签管理 - 批量为用户取消标签
func BatchUnTagging(tagID int64, openids ...string) wx.Action {
	params := &ParamsBatchUnTagging{
		TagID:      tagID,
		OpenIDList: openids,
	}

	return wx.NewPostAction(urls.OffiaBatchUnTagging,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsUserTags struct {
	OpenID string `json:"openid"`
}

type ResultUserTags struct {
	TagIDList []int64 `json:"tagid_list"`
}

// GetUserTags 用户管理 - 用户标签管理 - 获取用户身上的标签列表
func GetUserTags(openid string, result *ResultUserTags) wx.Action {
	params := &ParamsUserTags{
		OpenID: openid,
	}

	return wx.NewPostAction(urls.OffiaTagGetIDList,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsUserRemark struct {
	OpenID string `json:"openid"`
	Remark string `json:"remark"`
}

// SetUserRemark 用户管理 - 设置用户备注名（该接口暂时开放给微信认证的服务号）
func SetUserRemark(openid, remark string) wx.Action {
	params := &ParamsUserRemark{
		OpenID: openid,
		Remark: remark,
	}

	return wx.NewPostAction(urls.OffiaUserRemarkSet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

// UserInfo 用户基本信息(UnionID机制)
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

type ParamsUserInfo struct {
	OpenID string `json:"openid"`
	Lang   string `json:"lang,omitempty"`
}

// GetUserInfo 用户管理 - 获取用户基本信息（包括UnionID机制）
func GetUserInfo(openid, lang string, result *UserInfo) wx.Action {
	params := &ParamsUserInfo{
		OpenID: openid,
		Lang:   lang,
	}

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

type ParamsBatchUserInfo struct {
	UserList []*ParamsUserInfo `json:"user_list"`
}

type ResultBatchUserInfo struct {
	UserInfoList []*UserInfo `json:"user_info_list"`
}

// BatchGetUserInfo 用户管理 - 批量获取用户基本信息
func BatchGetUserInfo(users []*ParamsUserInfo, result *ResultBatchUserInfo) wx.Action {
	params := &ParamsBatchUserInfo{
		UserList: users,
	}

	return wx.NewPostAction(urls.OffiaUserBatchGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type UserListData struct {
	OpenID []string `json:"openid"`
}

type ResultUserList struct {
	Total      int          `json:"total"`
	Count      int          `json:"count"`
	Data       UserListData `json:"data"`
	NextOpenID string       `json:"next_openid"`
}

// GetUserList 用户管理 - 获取用户列表
func GetUserList(nextOpenID string, result *ResultUserList) wx.Action {
	options := []wx.ActionOption{
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	}

	if len(nextOpenID) != 0 {
		options = append(options, wx.WithQuery("next_openid", nextOpenID))
	}

	return wx.NewGetAction(urls.OffiaUserList, options...)
}

type ParamsBlackList struct {
	BeginOpenID string `json:"begin_openid"`
}

type ResultBlackList struct {
	Total      int          `json:"total"`
	Count      int          `json:"count"`
	Data       UserListData `json:"data"`
	NextOpenID string       `json:"next_openid"`
}

// GetBlackList 用户管理 - 获取公众号的黑名单列表
func GetBlackList(beginOpenID string, result *ResultBlackList) wx.Action {
	params := &ParamsBlackList{
		BeginOpenID: beginOpenID,
	}

	return wx.NewPostAction(urls.OffiaBlackListGet,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsBlackUsers struct {
	OpenIDList []string `json:"openid_list"`
}

// BlackUsers 用户管理 - 拉黑用户
func BlackUsers(openids ...string) wx.Action {
	params := &ParamsBlackUsers{
		OpenIDList: openids,
	}

	return wx.NewPostAction(urls.OffiaBatchBlackList,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}

type ParamsUnBlackUsers struct {
	OpenIDList []string `json:"openid_list"`
}

// UnBlackUsers 用户管理 - 取消拉黑用户
func UnBlackUsers(openids ...string) wx.Action {
	params := &ParamsUnBlackUsers{
		OpenIDList: openids,
	}

	return wx.NewPostAction(urls.OffiaBatchUnBlackList,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
	)
}
