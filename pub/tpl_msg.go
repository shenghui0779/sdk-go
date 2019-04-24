package pub

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tidwall/gjson"
)

// TplMsgData 公众号模板消息数据
type TplMsgData map[string]map[string]string

// TplMsg 公众号模板消息
type TplMsg struct {
	openID      string
	accessToken string
	redirectURL string
	mpAppID     string
	mpPagePath  string
}

// SetAccessToken 设置AccessToken
func (m *TplMsg) SetAccessToken(token string) {
	m.accessToken = token

	return
}

// SetRedirectURL 设置跳转URL
func (m *TplMsg) SetRedirectURL(url string) {
	m.redirectURL = url

	return
}

// SetMPPath 设置小程序跳转路径
func (m *TplMsg) SetMPPath(appid, path string) {
	m.mpAppID = appid
	m.mpPagePath = path

	return
}

// Send 发送模板消息
func (m *TplMsg) Send(tplID string, data TplMsgData) (int64, error) {
	body := utils.X{
		"touser":      m.openID,
		"template_id": tplID,
		"data":        data,
	}

	if m.redirectURL != "" {
		body["url"] = m.redirectURL
	}

	if m.mpPagePath != "" {
		body["miniprogram"] = map[string]string{
			"appid":    m.mpAppID,
			"pagepath": m.mpPagePath,
		}
	}

	b, err := json.Marshal(body)

	if err != nil {
		return 0, err
	}

	resp, err := utils.HTTPPost(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", m.accessToken), b, utils.WithRequestHeader("Content-Type", "application/json; charset=utf-8"))

	if err != nil {
		return 0, err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return 0, errors.New(r.Get("errmsg").String())
	}

	return r.Get("msgid").Int(), nil
}

// NewTplMsg ...
func NewTplMsg(openid string) *TplMsg {
	return &TplMsg{openID: openid}
}
