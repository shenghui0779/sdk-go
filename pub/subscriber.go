package pub

import (
	"encoding/json"
	"fmt"

	"github.com/shenghui0779/gochat/utils"
)

// MaxSubscriberListCount 公众号订阅者列表的最大数目
const MaxSubscriberListCount = 10000

// Subscriber 微信公众号订阅者
type Subscriber struct {
	pub     *WXPub
	options []utils.RequestOption
}

// SubscriberInfo 微信公众号订阅者信息
type SubscriberInfo struct {
	Subscribe      int     `json:"subscribe"`
	OpenID         string  `json:"openid"`
	NickName       string  `json:"nickName"`
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

// SubscriberList 微信公众号订阅者列表
type SubscriberList struct {
	Total      int                 `json:"total"`
	Count      int                 `json:"count"`
	Data       map[string][]string `json:"data"`
	NextOpenID string              `json:"next_openid"`
}

// GetSubscriberInfo 获取微信公众号订阅者信息
func (s *Subscriber) Get(accessToken, openid string) (*SubscriberInfo, error) {
	b, err := s.pub.get(fmt.Sprintf("%s?access_token=%s&openid=%s&lang=zh_CN", SubscriberGetURL, accessToken, openid), s.options...)

	if err != nil {
		return nil, err
	}

	resp := new(SubscriberInfo)

	if err := json.Unmarshal(b, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// BatchGet 批量获取微信公众号订阅者信息
func (s *Subscriber) BatchGet(accessToken string, openid ...string) ([]*SubscriberInfo, error) {
	l := len(openid)

	if l == 0 {
		return []*SubscriberInfo{}, nil
	}

	userList := make([]map[string]string, 0, l)

	for _, v := range openid {
		userList = append(userList, map[string]string{
			"openid": v,
			"lang":   "zh_CN",
		})
	}

	body, err := json.Marshal(map[string][]map[string]string{"user_list": userList})

	if err != nil {
		return nil, err
	}

	b, err := s.pub.post(fmt.Sprintf("%s?access_token=%s", SubscriberBatchGetURL, accessToken), body, s.options...)

	if err != nil {
		return nil, err
	}

	resp := make(map[string][]*SubscriberInfo)

	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}

	return resp["user_info_list"], nil
}

// GetList 获取微信公众号订阅者列表
func (s *Subscriber) GetList(accessToken string, nextOpenID ...string) (*SubscriberList, error) {
	url := fmt.Sprintf("%s?access_token=%s", SubscriberListURL, accessToken)

	if len(nextOpenID) > 0 {
		url = fmt.Sprintf("%s?access_token=%s&next_openid=%s", SubscriberListURL, accessToken, nextOpenID[0])
	}

	b, err := s.pub.get(url, s.options...)

	if err != nil {
		return nil, err
	}

	resp := new(SubscriberList)

	if err := json.Unmarshal(b, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
