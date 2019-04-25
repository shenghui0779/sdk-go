package pub

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/iiinsomnia/gochat/utils"
)

type menuReply struct {
	DefaultMenu     *DefaultMenu       `json:"menu"`
	ConditionalMenu []*ConditionalMenu `json:"conditionalmenu"`
	ErrCode         int                `json:"errcode"`
	ErrMsg          string             `json:"errmsg"`
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
	accessToken string
	buttons     []Button
	matchRule   *MenuMatchRule
	reply       *menuReply
}

func (m *Menu) SetButtons(btns ...Button) *Menu {
	m.buttons = btns

	return m
}

func (m *Menu) SetMatchRule(rule *MenuMatchRule) *Menu {
	m.matchRule = rule

	return m
}

func (m *Menu) Create() error {
	body := map[string][]Button{"button": m.buttons}

	b, err := json.Marshal(body)

	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%s", m.accessToken)

	if m.matchRule != nil {
		url = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/menu/addconditional?access_token=%s", m.accessToken)
	}

	resp, err := utils.HTTPPost(url, b, utils.WithRequestHeader("Content-Type", "application/json; charset=utf-8"))

	if err != nil {
		return err
	}

	reply := new(menuReply)

	if err := json.Unmarshal(resp, reply); err != nil {
		return err
	}

	if reply.ErrCode != 0 {
		return errors.New(reply.ErrMsg)
	}

	return nil
}

func (m *Menu) Get() error {
	resp, err := utils.HTTPGet(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/menu/get?access_token=%s", m.accessToken))

	if err != nil {
		return err
	}

	reply := new(menuReply)

	if err := json.Unmarshal(resp, reply); err != nil {
		return err
	}

	if reply.ErrCode != 0 {
		return errors.New(reply.ErrMsg)
	}

	m.reply = reply

	return nil
}

func (m *Menu) Delete() error {
	resp, err := utils.HTTPGet(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=%s", m.accessToken))

	if err != nil {
		return err
	}

	reply := new(menuReply)

	if err := json.Unmarshal(resp, reply); err != nil {
		return err
	}

	if reply.ErrCode != 0 {
		return errors.New(reply.ErrMsg)
	}

	return nil
}

func (m *Menu) DefaultMenu() *DefaultMenu {
	return m.reply.DefaultMenu
}

func (m *Menu) ConditionalMenu() []*ConditionalMenu {
	return m.reply.ConditionalMenu
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
	Type    string `json:"type"`
	Name    string `json:"name"`
	MediaID string `json:"media_id"`
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
