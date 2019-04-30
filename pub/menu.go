package pub

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/iiinsomnia/gochat/utils"
	"github.com/tidwall/gjson"
)

type MenuList struct {
	DefaultMenu     *DefaultMenu       `json:"menu"`
	ConditionalMenu []*ConditionalMenu `json:"conditionalmenu"`
}

type DefaultMenu struct {
	Button []*menuButton `json:"button"`
	MenuID int64         `json:"menuid"`
}

type ConditionalMenu struct {
	Button    []*menuButton  `json:"button"`
	MenuID    int64          `json:"menuid"`
	MatchRule *MenuMatchRule `json:"matchrule"`
}

type menuButton struct {
	Type      string        `json:"type"`
	Name      string        `json:"name"`
	Key       string        `json:"key,omitempty"`
	URL       string        `json:"url,omitempty"`
	AppID     string        `json:"appid,omitempty"`
	PagePath  string        `json:"page_path,omitempty"`
	MediaID   string        `json:"media_id,omitempty"`
	SubButton []*menuButton `json:"sub_button,omitempty"`
}

type MenuMatchRule struct {
	TagID              int64  `json:"tag_id"`
	Sex                int    `json:"sex"`
	Country            string `json:"country"`
	Province           string `json:"province"`
	City               string `json:"city"`
	ClientPlatformType int    `json:"client_platform_type"`
}

type Menu struct {
	client *utils.HTTPClient
}

func (m *Menu) Create(accessToken string, btns ...Button) error {
	body := utils.X{"button": btns}

	b, err := json.Marshal(body)

	if err != nil {
		return err
	}

	resp, err := m.client.Post(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%s", accessToken), b, utils.WithRequestHeader("Content-Type", "application/json; charset=utf-8"))

	if err != nil {
		return err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return errors.New(r.Get("errmsg").String())
	}

	return nil
}

func (m *Menu) CreateConditional(accessToken string, matchRule *MenuMatchRule, btns ...Button) error {
	body := utils.X{
		"button":    btns,
		"matchrule": matchRule,
	}

	b, err := json.Marshal(body)

	if err != nil {
		return err
	}

	resp, err := m.client.Post(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/menu/addconditional?access_token=%s", accessToken), b, utils.WithRequestHeader("Content-Type", "application/json; charset=utf-8"))

	if err != nil {
		return err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return errors.New(r.Get("errmsg").String())
	}

	return nil
}

func (m *Menu) GetList(accessToken string) (*MenuList, error) {
	resp, err := m.client.Get(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/menu/get?access_token=%s", accessToken))

	if err != nil {
		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return nil, errors.New(r.Get("errmsg").String())
	}

	reply := new(MenuList)

	if err := json.Unmarshal(resp, reply); err != nil {
		return nil, err
	}

	return reply, nil
}

func (m *Menu) Delete(accessToken string) error {
	resp, err := m.client.Get(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=%s", accessToken))

	if err != nil {
		return err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return errors.New(r.Get("errmsg").String())
	}

	return nil
}

type Button interface {
	AddSubButton(btn ...Button)
}

type groupButton struct {
	Name      string   `json:"name"`
	SubButton []Button `json:"sub_button"`
}

func (b *groupButton) AddSubButton(btn ...Button) {
	b.SubButton = append(b.SubButton, btn...)
}

// eventButton 事件按钮
type eventButton struct {
	Type      string   `json:"type"`
	Name      string   `json:"name"`
	Key       string   `json:"key"`
	SubButton []Button `json:"sub_button"`
}

func (b *eventButton) AddSubButton(btn ...Button) {}

// viewButton 跳转链接按钮
type viewButton struct {
	Type      string   `json:"type"`
	Name      string   `json:"name"`
	URL       string   `json:"url"`
	SubButton []Button `json:"sub_button"`
}

func (b *viewButton) AddSubButton(btn ...Button) {}

// mpButton 小程序跳转按钮
type mpButton struct {
	Type      string   `json:"type"`
	Name      string   `json:"name"`
	URL       string   `json:"url,omitempty"`
	AppID     string   `json:"appid"`
	PagePath  string   `json:"pagepath"`
	SubButton []Button `json:"sub_button"`
}

func (b *mpButton) AddSubButton(btn ...Button) {}

// mediaButton 媒体素材按钮
type mediaButton struct {
	Type      string   `json:"type"`
	Name      string   `json:"name"`
	MediaID   string   `json:"media_id"`
	SubButton []Button `json:"sub_button"`
}

func (b *mediaButton) AddSubButton(btn ...Button) {}

// GroupButton 组合按钮
func GroupButton(name string) Button {
	return &groupButton{Name: name}
}

// ClickButton 点击事件按钮
func ClickButton(name, key string) Button {
	return &eventButton{
		Type: "click",
		Name: name,
		Key:  key,
	}
}

// ViewButton 跳转URL按钮
func ViewButton(name, url string) Button {
	return &viewButton{
		Type: "view",
		Name: name,
		URL:  url,
	}
}

// MPButton 小程序跳转按钮
func MPButton(name, appid, pagepath string, url ...string) Button {
	btn := &mpButton{
		Type:     "miniprogram",
		Name:     name,
		AppID:    appid,
		PagePath: pagepath,
	}

	if len(url) > 0 {
		btn.URL = url[0]
	}

	return btn
}

// ScancodePushButton 扫码推事件按钮
func ScancodePushButton(name, key string) Button {
	return &eventButton{
		Type: "click",
		Name: name,
		Key:  key,
	}
}

// ScancodeWaitMsgButton 扫码带提示按钮
func ScancodeWaitMsgButton(name, key string) Button {
	return &eventButton{
		Type: "scancode_waitmsg",
		Name: name,
		Key:  key,
	}
}

// PicSysphotoButton 系统拍照发图按钮
func PicSysphotoButton(name, key string) Button {
	return &eventButton{
		Type: "pic_sysphoto",
		Name: name,
		Key:  key,
	}
}

// PicPhotoOrAlbum 拍照或者相册发图按钮
func PicPhotoOrAlbum(name, key string) Button {
	return &eventButton{
		Type: "pic_photo_or_album",
		Name: name,
		Key:  key,
	}
}

// PicWeixin 微信相册发图按钮
func PicWeixin(name, key string) Button {
	return &eventButton{
		Type: "pic_weixin",
		Name: name,
		Key:  key,
	}
}

// LocationSelectButton 发送位置按钮
func LocationSelectButton(name, key string) Button {
	return &eventButton{
		Type: "location_select",
		Name: name,
		Key:  key,
	}
}

// MediaButton 图片按钮
func MediaButton(name, mediaID string) Button {
	return &mediaButton{
		Type:    "media_id",
		Name:    name,
		MediaID: mediaID,
	}
}

// ViewLimitedButton 图文消息按钮
func ViewLimitedButton(name, mediaID string) Button {
	return &mediaButton{
		Type:    "view_limited",
		Name:    name,
		MediaID: mediaID,
	}
}
