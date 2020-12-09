package oa

import (
	"encoding/json"
	"net/url"

	"github.com/shenghui0779/gochat/wx"
)

// MenuButtonType 菜单按钮类型
type MenuButtonType string

// 微信支持的按钮
const (
	ButtonClick           MenuButtonType = "click"              // 点击推事件用户点击click类型按钮后，微信服务器会通过消息接口推送消息类型为event的结构给开发者（参考消息接口指南），并且带上按钮中开发者填写的key值，开发者可以通过自定义的key值与用户进行交互。
	ButtonView            MenuButtonType = "view"               // 跳转URL用户点击view类型按钮后，微信客户端将会打开开发者在按钮中填写的网页URL，可与网页授权获取用户基本信息接口结合，获得用户基本信息。
	ButtonScanCodePush    MenuButtonType = "scancode_push"      // 扫码推事件用户点击按钮后，微信客户端将调起扫一扫工具，完成扫码操作后显示扫描结果（如果是URL，将进入URL），且会将扫码的结果传给开发者，开发者可以下发消息。
	ButtonScanCodeWaitMsg MenuButtonType = "scancode_waitmsg"   // 扫码推事件且弹出“消息接收中”提示框用户点击按钮后，微信客户端将调起扫一扫工具，完成扫码操作后，将扫码的结果传给开发者，同时收起扫一扫工具，然后弹出“消息接收中”提示框，随后可能会收到开发者下发的消息。
	ButtonPicSysPhoto     MenuButtonType = "pic_sysphoto"       // 弹出系统拍照发图用户点击按钮后，微信客户端将调起系统相机，完成拍照操作后，会将拍摄的相片发送给开发者，并推送事件给开发者，同时收起系统相机，随后可能会收到开发者下发的消息。
	ButtonPicPhotoOrAlbum MenuButtonType = "pic_photo_or_album" // 弹出拍照或者相册发图用户点击按钮后，微信客户端将弹出选择器供用户选择“拍照”或者“从手机相册选择”。用户选择后即走其他两种流程。
	ButtonPicWeixin       MenuButtonType = "pic_weixin"         // 弹出微信相册发图器用户点击按钮后，微信客户端将调起微信相册，完成选择操作后，将选择的相片发送给开发者的服务器，并推送事件给开发者，同时收起相册，随后可能会收到开发者下发的消息。
	ButtonLocationSelect  MenuButtonType = "location_select"    // 弹出地理位置选择器用户点击按钮后，微信客户端将调起地理位置选择工具，完成选择操作后，将选择的地理位置发送给开发者的服务器，同时收起位置选择工具，随后可能会收到开发者下发的消息。
	ButtonMedia           MenuButtonType = "media_id"           // 下发消息（除文本消息）用户点击media_id类型按钮后，微信服务器会将开发者填写的永久素材id对应的素材下发给用户，永久素材类型可以是图片、音频、视频、图文消息。请注意：永久素材id必须是在“素材管理/新增永久素材”接口上传后获得的合法id。
	ButtonViewLimited     MenuButtonType = "view_limited"       // 跳转图文消息URL用户点击view_limited类型按钮后，微信客户端将打开开发者在按钮中填写的永久素材id对应的图文消息URL，永久素材类型只支持图文消息。请注意：永久素材id必须是在“素材管理/新增永久素材”接口上传后获得的合法id。​
	ButtonMinip           MenuButtonType = "miniprogram"        // 小程序页面跳转，不支持小程序的老版本客户端将打开指定的URL。
)

// MenuInfo 自定义菜单信息
type MenuInfo struct {
	Menu            Menu               `json:"menu"`
	ConditionalMenu []*ConditionalMenu `json:"conditionalmenu"`
}

// Menu 普通菜单
type Menu struct {
	Button []*MenuButton `json:"button"`
	MenuID int64         `json:"menuid"`
}

// ConditionalMenu 个性化菜单
type ConditionalMenu struct {
	Button    []*MenuButton `json:"button"`
	MatchRule MenuMatchRule `json:"matchrule"`
	MenuID    int64         `json:"menuid"`
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
	TagID              string `json:"tag_id,omitempty"`               // 用户标签的id，可通过用户标签管理接口获取，不填则不做匹配
	Sex                string `json:"sex,omitempty"`                  // 性别：男（1）女（2），不填则不做匹配
	Country            string `json:"country,omitempty"`              // 国家信息，是用户在微信中设置的地区，具体请参考地区信息表，不填则不做匹配
	Province           string `json:"province,omitempty"`             // 省份信息，是用户在微信中设置的地区，具体请参考地区信息表，不填则不做匹配
	City               string `json:"city,omitempty"`                 // 城市信息，是用户在微信中设置的地区，具体请参考地区信息表，不填则不做匹配
	ClientPlatformType string `json:"client_platform_type,omitempty"` // 客户端版本，当前只具体到系统型号：IOS(1), Android(2),Others(3)，不填则不做匹配
	Language           string `json:"language,omitempty"`             // 语言信息，是用户在微信中设置的语言，具体请参考语言表，不填则不做匹配
}

// CreateMenu 创建自定义菜单
func CreateMenu(buttons ...*MenuButton) wx.Action {
	return wx.NewPostAPI(MenuCreateURL, url.Values{}, func() ([]byte, error) {
		return json.Marshal(wx.X{"button": buttons})
	}, nil)
}

// CreateConditionalMenu 创建个性化菜单
func CreateConditionalMenu(matchRule *MenuMatchRule, buttons ...*MenuButton) wx.Action {
	return wx.NewPostAPI(MenuAddConditionalURL, url.Values{}, func() ([]byte, error) {
		return json.Marshal(wx.X{
			"button":    buttons,
			"matchrule": matchRule,
		})
	}, nil)
}

// GetMenu 查询自定义菜单
func GetMenu(dest *MenuInfo) wx.Action {
	return wx.NewGetAPI(MenuListURL, url.Values{}, func(resp []byte) error {
		return json.Unmarshal(resp, dest)
	})
}

// DeleteMenu 删除自定义菜单
func DeleteMenu() wx.Action {
	return wx.NewGetAPI(MenuDeleteURL, url.Values{}, nil)
}

// DeleteConditional 删除个性化菜单
func DeleteConditionalMenu(menuID string) wx.Action {
	return wx.NewPostAPI(MenuDeleteConditionalURL, url.Values{}, func() ([]byte, error) {
		return json.Marshal(wx.X{"menuid": menuID})
	}, nil)
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
		Type: ButtonClick,
		Name: name,
		Key:  key,
	}
}

// ViewButton 跳转URL按钮
func ViewButton(name, redirectURL string) *MenuButton {
	return &MenuButton{
		Type: ButtonView,
		Name: name,
		URL:  redirectURL,
	}
}

// ScanCodePushButton 扫码推事件按钮
func ScanCodePushButton(name, key string) *MenuButton {
	return &MenuButton{
		Type: ButtonScanCodePush,
		Name: name,
		Key:  key,
	}
}

// ScanCodeWaitMsgButton 扫码带提示按钮
func ScanCodeWaitMsgButton(name, key string) *MenuButton {
	return &MenuButton{
		Type: ButtonScanCodeWaitMsg,
		Name: name,
		Key:  key,
	}
}

// PicSysPhotoButton 系统拍照发图按钮
func PicSysPhotoButton(name, key string) *MenuButton {
	return &MenuButton{
		Type: ButtonPicSysPhoto,
		Name: name,
		Key:  key,
	}
}

// PicPhotoOrAlbum 拍照或者相册发图按钮
func PicPhotoOrAlbumButton(name, key string) *MenuButton {
	return &MenuButton{
		Type: ButtonPicPhotoOrAlbum,
		Name: name,
		Key:  key,
	}
}

// PicWeixinButton 微信相册发图按钮
func PicWeixinButton(name, key string) *MenuButton {
	return &MenuButton{
		Type: ButtonPicWeixin,
		Name: name,
		Key:  key,
	}
}

// LocationSelectButton 发送位置按钮
func LocationSelectButton(name, key string) *MenuButton {
	return &MenuButton{
		Type: ButtonLocationSelect,
		Name: name,
		Key:  key,
	}
}

// MediaButton 素材按钮
func MediaButton(name, mediaID string) *MenuButton {
	return &MenuButton{
		Type:    ButtonMedia,
		Name:    name,
		MediaID: mediaID,
	}
}

// ViewLimitedButton 图文消息按钮
func ViewLimitedButton(name, mediaID string) *MenuButton {
	return &MenuButton{
		Type:    ButtonViewLimited,
		Name:    name,
		MediaID: mediaID,
	}
}

// MinipButton 小程序跳转按钮
func MinipButton(name, appid, pagepath, redirectURL string) *MenuButton {
	return &MenuButton{
		Type:     ButtonMinip,
		Name:     name,
		URL:      redirectURL,
		AppID:    appid,
		PagePath: pagepath,
	}
}
