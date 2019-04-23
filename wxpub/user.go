package wxpub

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/iiinsomnia/yiigo"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
)

// MaxUserListCount 用户列表的最大数目
const MaxUserListCount = 10000

// WXPubUserInfo 微信公众号用户信息
type WXPubUserInfo struct {
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

// WXPubUserList 微信公众号用户列表
type WXPubUserList struct {
	Total      int                 `json:"total"`
	Count      int                 `json:"count"`
	Data       map[string][]string `json:"data"`
	NextOpenID string              `json:"next_openid"`
}

// GetWXPubUserInfo 获取微信公众号用户信息
func GetWXPubUserInfo(accessToken, openid string) (*WXPubUserInfo, error) {
	resp, err := yiigo.HTTPGet(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=zh_CN", accessToken, openid))

	if err != nil {
		yiigo.Logger.Error("get wxpub userinfo error", zap.String("error", err.Error()))

		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		yiigo.Logger.Error("get wxpub userinfo error", zap.ByteString("resp", resp))

		return nil, errors.New(r.Get("errmsg").String())
	}

	reply := new(WXPubUserInfo)

	if err := json.Unmarshal(resp, reply); err != nil {
		yiigo.Logger.Error("unmarshal wxpub userinfo error", zap.String("error", err.Error()))

		return nil, err
	}

	return reply, nil
}

// BatchGetWXPubUserInfo 批量获取微信公众号用户信息
func BatchGetWXPubUserInfo(accessToken string, openid ...string) ([]*WXPubUserInfo, error) {
	l := len(openid)

	if l == 0 {
		return []*WXPubUserInfo{}, nil
	}

	userList := make([]map[string]string, 0, l)

	for _, v := range openid {
		userList = append(userList, map[string]string{
			"openid": v,
			"lang":   "zh_CN",
		})
	}

	b, err := json.Marshal(map[string][]map[string]string{"user_list": userList})

	if err != nil {
		yiigo.Logger.Error("marshal batch get wxpub userinfo req body error", zap.String("error", err.Error()))

		return nil, err
	}

	resp, err := yiigo.HTTPPost(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/user/info/batchget?access_token=%s", accessToken), b, yiigo.WithHeader("Content-Type", "application/json; charset=utf-8"))

	if err != nil {
		yiigo.Logger.Error("batch get wxpub userinfo error", zap.String("error", err.Error()))

		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		yiigo.Logger.Error("batch get wxpub userinfo error", zap.ByteString("resp", resp))

		return nil, errors.New(r.Get("errmsg").String())
	}

	reply := make(map[string][]*WXPubUserInfo)

	if err := json.Unmarshal(resp, reply); err != nil {
		yiigo.Logger.Error("unmarshal batch wxpub userinfo error", zap.String("error", err.Error()))

		return nil, err
	}

	return reply["user_info_list"], nil
}

// GetUserList 获取微信公众号用户列表
func GetWXPubUserList(accessToken string, nextOpenID ...string) (*WXPubUserList, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/user/get?access_token=%s", accessToken)

	if len(nextOpenID) > 0 {
		url = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/user/get?access_token=%s&next_openid=%s", accessToken, nextOpenID[0])
	}

	resp, err := yiigo.HTTPGet(url)

	if err != nil {
		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		yiigo.Logger.Error("get wxpub user list error", zap.ByteString("resp", resp))

		return nil, errors.New(r.Get("errmsg").String())
	}

	reply := new(WXPubUserList)

	if err := json.Unmarshal(resp, reply); err != nil {
		yiigo.Logger.Error("unmarshal wxpub user list error", zap.String("error", err.Error()))

		return nil, err
	}

	return reply, nil
}
