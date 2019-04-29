package pub

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/iiinsomnia/gochat/utils"
	"github.com/tidwall/gjson"
)

type TplMsgBody map[string]map[string]string

// TplMsgData 公众号模板消息数据
type TplMsgData struct {
	TplID       string
	OpenID      string
	RedirectURL string
	MPAppID     string
	MPPagePath  string
	Body        TplMsgBody
}

// TplMsg 公众号模板消息
type TplMsg struct {
	client *utils.HTTPClient
}

// Send 发送模板消息
func (t *TplMsg) Send(accessToken string, data *TplMsgData) (int64, error) {
	body := utils.X{
		"touser":      data.OpenID,
		"template_id": data.TplID,
		"data":        data,
	}

	if data.RedirectURL != "" {
		body["url"] = data.RedirectURL
	}

	if data.MPAppID != "" {
		body["miniprogram"] = map[string]string{
			"appid":    data.MPAppID,
			"pagepath": data.MPPagePath,
		}
	}

	b, err := json.Marshal(body)

	if err != nil {
		return 0, err
	}

	resp, err := t.client.Post(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", accessToken), b, utils.WithRequestHeader("Content-Type", "application/json; charset=utf-8"))

	if err != nil {
		return 0, err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return 0, errors.New(r.Get("errmsg").String())
	}

	return r.Get("msgid").Int(), nil
}
