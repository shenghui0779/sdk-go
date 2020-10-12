package pub

import (
	"encoding/json"
	"fmt"

	"github.com/shenghui0779/gochat/utils"
	"github.com/tidwall/gjson"
)

// MsgBody 消息内容体
type MsgBody map[string]map[string]string

// TplMsg 公众号模板消息
type TplMsg struct {
	OpenID      string  // 接收者（用户）的 openid
	TplID       string  // 模板ID
	RedirectURL string  // 模板跳转链接（海外帐号没有跳转能力）
	MPAppID     string  // 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系，暂不支持小游戏）
	MPPagePath  string  // 所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar），要求该小程序已发布，暂不支持小游戏
	Data        MsgBody // 模板内容，格式形如：{"key1": {"value": any}, "key2": {"value": any}}
	Color       string  // 模板内容字体颜色，不填默认为黑色
}

// Message 公众号消息
type Message struct {
	pub     *WXPub
	options []utils.RequestOption
}

// Template 发送模板消息
func (m *Message) Template(data *TplMsg, accessToken string) (int64, error) {
	params := utils.X{
		"touser":      data.OpenID,
		"template_id": data.TplID,
		"data":        data.Data,
	}

	if data.RedirectURL != "" {
		params["url"] = data.RedirectURL
	}

	if data.MPAppID != "" {
		params["miniprogram"] = map[string]string{
			"appid":    data.MPAppID,
			"pagepath": data.MPPagePath,
		}
	}

	body, err := json.Marshal(params)

	if err != nil {
		return 0, err
	}

	b, err := m.pub.post(fmt.Sprintf("%s?access_token=%s", TplMsgSendURL, accessToken), body, m.options...)

	if err != nil {
		return 0, err
	}

	return gjson.GetBytes(b, "msgid").Int(), nil
}
