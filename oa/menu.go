package oa

import (
	"encoding/json"
	"fmt"

	"github.com/shenghui0779/gochat/internal"
)

type MenuButtonType string

const (
	ButtonTypeClick           MenuButtonType = "click"
	ButtonTypeView            MenuButtonType = "view"
	ButtonTypeMP              MenuButtonType = "miniprogram"
	ButtonTypeScanCodePush    MenuButtonType = "scancode_push"
	ButtonTypeScanCodeWaitMsg MenuButtonType = "scancode_waitmsg"
	ButtonTypePicSysPhoto     MenuButtonType = "pic_sysphoto"
	ButtonTypePicPhotoOrAlbum MenuButtonType = "pic_photo_or_album"
	ButtonTypePicWeixin       MenuButtonType = "pic_weixin"
	ButtonTypeLocationSelect  MenuButtonType = "location_select"
	ButtonTypeMedia           MenuButtonType = "media_id"
	ButtonTypeViewLimited     MenuButtonType = "view_limited"
)

// MenuInfo 自定义菜单信息
type MenuInfo struct {
	DefaultMenu     *DefaultMenu       `json:"menu"`
	ConditionalMenu []*ConditionalMenu `json:"conditionalmenu"`
}

// DefaultMenu 默认菜单
type DefaultMenu struct {
	Button []*MenuButton `json:"button"`
	MenuID int64         `json:"menuid"`
}

// ConditionalMenu 个性化菜单
type ConditionalMenu struct {
	Button    []*MenuButton  `json:"button"`
	MatchRule *MenuMatchRule `json:"matchrule"`
	MenuID    int64          `json:"menuid"`
}

// MenuButton 菜单按钮
type MenuButton struct {
	Type      MenuButtonType `json:"type,omitempty"`
	Name      string         `json:"name,omitempty"`
	Key       string         `json:"key,omitempty"`
	URL       string         `json:"url,omitempty"`
	AppID     string         `json:"appid,omitempty"`
	PagePath  string         `json:"page_path,omitempty"`
	MediaID   string         `json:"media_id,omitempty"`
	SubButton []*MenuButton  `json:"sub_button,omitempty"`
}

// MenuMatchRule 菜单匹配规则
type MenuMatchRule struct {
	TagID              string `json:"tag_id,omitempty"`
	Sex                string `json:"sex,omitempty"`
	Country            string `json:"country,omitempty"`
	Province           string `json:"province,omitempty"`
	City               string `json:"city,omitempty"`
	ClientPlatformType string `json:"client_platform_type,omitempty"`
	Language           string `json:"language,omitempty"`
}

// CreateMenu 自定义菜单
func CreateMenu(buttons ...*MenuButton) internal.Action {
	return &WechatAPI{
		body: internal.NewPostBody(func() ([]byte, error) {
			return json.Marshal(internal.X{"button": buttons})
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("POST|%s?access_token=%s", MenuCreateURL, accessToken)
		},
	}
}

// CreateConditionalMenu 个性化菜单
func CreateConditionalMenu(matchRule *MenuMatchRule, buttons ...*MenuButton) internal.Action {
	return &WechatAPI{
		body: internal.NewPostBody(func() ([]byte, error) {
			return json.Marshal(internal.X{
				"button":    buttons,
				"matchrule": matchRule,
			})
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("POST|%s?access_token=%s", MenuAddConditionalURL, accessToken)
		},
	}
}

// GetMenu 查询自定义菜单
func GetMenu(dest *MenuInfo) internal.Action {
	return &WechatAPI{
		url: func(accessToken string) string {
			return fmt.Sprintf("GET|%s?access_token=%s", MenuListURL, accessToken)
		},
		decode: func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		},
	}
}

// DeleteMenu 删除自定义菜单
func DeleteMenu() internal.Action {
	return &WechatAPI{
		url: func(accessToken string) string {
			return fmt.Sprintf("GET|%s?access_token=%s", MenuDeleteURL, accessToken)
		},
	}
}

// DeleteConditional 删除个性化菜单
func DeleteConditionalMenu(menuID string) internal.Action {
	return &WechatAPI{
		body: internal.NewPostBody(func() ([]byte, error) {
			return json.Marshal(internal.X{"menuid": menuID})
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("POST|%s?access_token=%s", MenuDeleteConditionalURL, accessToken)
		},
	}
}

// GroupButton 组合按钮
func GroupButton(name string, buttons ...*MenuButton) *MenuButton {
	btn := &MenuButton{Name: name}

	if len(buttons) != 0 {
		btn.SubButton = buttons
	}

	return btn
}

// ClickButton 点击事件按钮
func ClickButton(name, key string) *MenuButton {
	return &MenuButton{
		Type: ButtonTypeClick,
		Name: name,
		Key:  key,
	}
}

// ViewButton 跳转URL按钮
func ViewButton(name, url string) *MenuButton {
	return &MenuButton{
		Type: ButtonTypeView,
		Name: name,
		URL:  url,
	}
}

// MPButton 小程序跳转按钮
func MPButton(name, appid, pagepath, url string) *MenuButton {
	return &MenuButton{
		Type:     ButtonTypeMP,
		Name:     name,
		URL:      url,
		AppID:    appid,
		PagePath: pagepath,
	}
}

// ScanCodePushButton 扫码推事件按钮
func ScanCodePushButton(name, key string) *MenuButton {
	return &MenuButton{
		Type: ButtonTypeScanCodePush,
		Name: name,
		Key:  key,
	}
}

// ScanCodeWaitMsgButton 扫码带提示按钮
func ScanCodeWaitMsgButton(name, key string) *MenuButton {
	return &MenuButton{
		Type: ButtonTypeScanCodeWaitMsg,
		Name: name,
		Key:  key,
	}
}

// PicSysPhotoButton 系统拍照发图按钮
func PicSysPhotoButton(name, key string) *MenuButton {
	return &MenuButton{
		Type: ButtonTypePicSysPhoto,
		Name: name,
		Key:  key,
	}
}

// PicPhotoOrAlbum 拍照或者相册发图按钮
func PicPhotoOrAlbumButton(name, key string) *MenuButton {
	return &MenuButton{
		Type: ButtonTypePicPhotoOrAlbum,
		Name: name,
		Key:  key,
	}
}

// PicWeixinButton 微信相册发图按钮
func PicWeixinButton(name, key string) *MenuButton {
	return &MenuButton{
		Type: ButtonTypePicWeixin,
		Name: name,
		Key:  key,
	}
}

// LocationSelectButton 发送位置按钮
func LocationSelectButton(name, key string) *MenuButton {
	return &MenuButton{
		Type: ButtonTypeLocationSelect,
		Name: name,
		Key:  key,
	}
}

// MediaButton 素材按钮
func MediaButton(name, mediaID string) *MenuButton {
	return &MenuButton{
		Type:    ButtonTypeMedia,
		Name:    name,
		MediaID: mediaID,
	}
}

// ViewLimitedButton 图文消息按钮
func ViewLimitedButton(name, mediaID string) *MenuButton {
	return &MenuButton{
		Type:    ButtonTypeViewLimited,
		Name:    name,
		MediaID: mediaID,
	}
}
