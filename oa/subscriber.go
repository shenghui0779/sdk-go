package oa

import (
	"encoding/json"
	"fmt"

	"github.com/shenghui0779/gochat/helpers"
	"github.com/tidwall/gjson"
)

// MaxSubscriberListCount 公众号订阅者列表的最大数目
const MaxSubscriberListCount = 10000

// SubscriberInfo 微信公众号订阅者信息
type SubscriberInfo struct {
	Subscribe      int     `json:"subscribe"`
	OpenID         string  `json:"openid"`
	NickName       string  `json:"nickname"`
	Sex            int     `json:"sex"`
	Language       string  `json:"language"`
	City           string  `json:"city"`
	Province       string  `json:"province"`
	Country        string  `json:"country"`
	AvatarURL      string  `json:"headimgurl"`
	SubscribeTime  int64   `json:"subscribe_time"`
	UnionID        string  `json:"unionid"`
	Remark         string  `json:"remark"`
	GroupID        int64   `json:"groupid"`
	TagidList      []int64 `json:"tagid_list"`
	SubscribeScene string  `json:"subscribe_scene"`
	QRScene        int64   `json:"qr_scene"`
	QRSceneStr     string  `json:"qr_scene_str"`
}

// BatchSubscriberInfo 微信公众号批量订阅者信息
type BatchSubscriberInfo struct {
	UserInfoList []*SubscriberInfo `json:"user_info_list"`
}

// SubscriberList 微信公众号订阅者列表
type SubscriberList struct {
	Total      int                 `json:"total"`
	Count      int                 `json:"count"`
	Data       *SubscriberListData `json:"data"`
	NextOpenID string              `json:"next_openid"`
}

// SubscriberListData 微信公众号订阅者列表数据
type SubscriberListData struct {
	OpenID []string `json:"openid"`
}

// GetSubscriberInfo 获取微信公众号订阅者信息
func GetSubscriberInfo(openid string, receiver *SubscriberInfo) Action {
	return &WechatAPI{
		url: func(accessToken string) string {
			return fmt.Sprintf("GET|%s?access_token=%s&openid=%s&lang=zh_CN", SubscriberGetURL, accessToken, openid)
		},
		decode: func(resp []byte) error {
			return json.Unmarshal(resp, receiver)
		},
	}
}

// BatchGetSubscriberInfo 批量获取微信公众号订阅者信息
func BatchGetSubscriberInfo(openids []string, receiver *[]SubscriberInfo) Action {
	return &WechatAPI{
		body: helpers.NewPostBody(func() ([]byte, error) {
			userList := make([]map[string]string, 0, len(openids))

			for _, v := range openids {
				userList = append(userList, map[string]string{
					"openid": v,
					"lang":   "zh_CN",
				})
			}

			return json.Marshal(map[string][]map[string]string{"user_list": userList})
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("POST|%s?access_token=%s", SubscriberBatchGetURL, accessToken)
		},
		decode: func(resp []byte) error {
			r := gjson.GetBytes(resp, "user_info_list")

			return json.Unmarshal([]byte(r.Raw), receiver)
		},
	}
}

// GetSubscriberList 获取微信公众号订阅者列表
func GetSubscriberList(nextOpenID string, receiver *SubscriberList) Action {
	return &WechatAPI{
		url: func(accessToken string) string {
			return fmt.Sprintf("GET|%s?access_token=%s&next_openid=%s", SubscriberListURL, accessToken, nextOpenID)
		},
		decode: func(resp []byte) error {
			return json.Unmarshal(resp, receiver)
		},
	}
}

// GetBlackList 获取用户黑名单列表
func GetBlackList(beginOpenID string, receiver *SubscriberList) Action {
	return &WechatAPI{
		body: helpers.NewPostBody(func() ([]byte, error) {
			return json.Marshal(helpers.X{
				"begin_openid": beginOpenID,
			})
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("POST|%s?access_token=%s", BlackListGetURL, accessToken)
		},
		decode: func(resp []byte) error {
			return json.Unmarshal(resp, receiver)
		},
	}
}

// BatchBlackList 拉黑用户
func BatchBlackList(openids []string) Action {
	return &WechatAPI{
		body: helpers.NewPostBody(func() ([]byte, error) {
			return json.Marshal(helpers.X{
				"openid_list": openids,
			})
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("POST|%s?access_token=%s", BatchBlackListURL, accessToken)
		},
	}
}

// BatchUnBlackList 取消拉黑用户
func BatchUnBlackList(openids []string) Action {
	return &WechatAPI{
		body: helpers.NewPostBody(func() ([]byte, error) {
			return json.Marshal(helpers.X{
				"openid_list": openids,
			})
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("POST|%s?access_token=%s", BatchUnBlackListURL, accessToken)
		},
	}
}

// SetUserRemark 设置用户备注名（该接口暂时开放给微信认证的服务号）
func SetUserRemark(openid, remark string) Action {
	return &WechatAPI{
		body: helpers.NewPostBody(func() ([]byte, error) {
			return json.Marshal(helpers.X{
				"openid": openid,
				"remark": remark,
			})
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("POST|%s?access_token=%s", UserRemarkSetURL, accessToken)
		},
	}
}
