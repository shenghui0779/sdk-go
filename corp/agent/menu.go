package agent

import (
	"encoding/json"
	"strconv"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

// ButtonType 菜单按钮类型
type ButtonType string

// 微信支持的按钮
const (
	ButtonClick           ButtonType = "click"              // 成员点击click类型按钮后，企业微信服务器会通过消息接口推送消息类型为event 的结构给开发者（参考消息接口指南），并且带上按钮中开发者填写的key值，开发者可以通过自定义的key值与成员进行交互
	ButtonView            ButtonType = "view"               // 成员点击view类型按钮后，企业微信客户端将会打开开发者在按钮中填写的网页URL，可与网页授权获取成员基本信息接口结合，获得成员基本信息
	ButtonScanCodePush    ButtonType = "scancode_push"      // 成员点击按钮后，企业微信客户端将调起扫一扫工具，完成扫码操作后显示扫描结果（如果是URL，将进入URL），且会将扫码的结果传给开发者，开发者可用于下发消息
	ButtonScanCodeWaitMsg ButtonType = "scancode_waitmsg"   // 成员点击按钮后，企业微信客户端将调起扫一扫工具，完成扫码操作后，将扫码的结果传给开发者，同时收起扫一扫工具，然后弹出“消息接收中”提示框，随后可能会收到开发者下发的消息
	ButtonPicSysPhoto     ButtonType = "pic_sysphoto"       // 弹出系统拍照发图 成员点击按钮后，企业微信客户端将调起系统相机，完成拍照操作后，会将拍摄的相片发送给开发者，并推送事件给开发者，同时收起系统相机，随后可能会收到开发者下发的消息
	ButtonPicPhotoOrAlbum ButtonType = "pic_photo_or_album" // 成员点击按钮后，企业微信客户端将弹出选择器供成员选择“拍照”或者“从手机相册选择”。成员选择后即走其他两种流程
	ButtonPicWeixin       ButtonType = "pic_weixin"         // 成员点击按钮后，企业微信客户端将调起企业微信相册，完成选择操作后，将选择的相片发送给开发者的服务器，并推送事件给开发者，同时收起相册，随后可能会收到开发者下发的消息
	ButtonLocationSelect  ButtonType = "location_select"    // 成员点击按钮后，企业微信客户端将调起地理位置选择工具，完成选择操作后，将选择的地理位置发送给开发者的服务器，同时收起位置选择工具，随后可能会收到开发者下发的消息
	ButtonViewMinip       ButtonType = "view_miniprogram"   // 成员点击按钮后，企业微信客户端将会打开开发者在按钮中配置的小程序
)

// Button 菜单按钮
type Button struct {
	Type      ButtonType `json:"type,omitempty"`       // 菜单的响应动作类型，view表示网页类型，click表示点击类型，miniprogram表示小程序类型
	Name      string     `json:"name,omitempty"`       // 菜单标题，不超过16个字节，子菜单不超过60个字节
	Key       string     `json:"key,omitempty"`        // click等点击类型必须，菜单KEY值，用于消息接口推送，不超过128字节
	URL       string     `json:"url,omitempty"`        // view、miniprogram类型必须，网页 链接，用户点击菜单可打开链接，不超过1024字节。 type为miniprogram时，不支持小程序的老版本客户端将打开本url。
	PagePath  string     `json:"pagepath,omitempty"`   // miniprogram类型必须，小程序的页面路径
	AppID     string     `json:"appid,omitempty"`      // miniprogram类型必须，小程序的appid（仅认证公众号可配置）
	SubButton []*Button  `json:"sub_button,omitempty"` // 二级菜单数组，个数应为1~5个
}

type ParamsMenuCreate struct {
	Button []*Button `json:"button"`
}

// CreateMenu 创建菜单
func CreateMenu(agentID int64, params *ParamsMenuCreate) wx.Action {
	return wx.NewPostAction(urls.CorpMenuCreate,
		wx.WithQuery("agentid", strconv.FormatInt(agentID, 10)),
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type ResultMenuGet struct {
	Button []*Button `json:"button"`
}

// GetMenu 获取菜单
func GetMenu(agentID int64, result *ResultMenuGet) wx.Action {
	return wx.NewGetAction(urls.CorpMenuGet,
		wx.WithQuery("agentid", strconv.FormatInt(agentID, 10)),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

// DeleteMenu 删除菜单
func DeleteMenu(agentID int64) wx.Action {
	return wx.NewGetAction(urls.CorpMenuDelete,
		wx.WithQuery("agentid", strconv.FormatInt(agentID, 10)),
	)
}

// GroupButton 组合按钮
func GroupButton(name string, buttons ...*Button) *Button {
	btn := &Button{Name: name}

	if len(buttons) != 0 {
		btn.SubButton = buttons
	}

	return btn
}

// ClickButton 点击事件按钮
func ClickButton(name, key string) *Button {
	return &Button{
		Type: ButtonClick,
		Name: name,
		Key:  key,
	}
}

// ViewButton 跳转URL按钮
func ViewButton(name, redirectURL string) *Button {
	return &Button{
		Type: ButtonView,
		Name: name,
		URL:  redirectURL,
	}
}

// ScanCodePushButton 扫码推事件按钮
func ScanCodePushButton(name, key string) *Button {
	return &Button{
		Type: ButtonScanCodePush,
		Name: name,
		Key:  key,
	}
}

// ScanCodeWaitMsgButton 扫码带提示按钮
func ScanCodeWaitMsgButton(name, key string) *Button {
	return &Button{
		Type: ButtonScanCodeWaitMsg,
		Name: name,
		Key:  key,
	}
}

// PicSysPhotoButton 系统拍照发图按钮
func PicSysPhotoButton(name, key string) *Button {
	return &Button{
		Type: ButtonPicSysPhoto,
		Name: name,
		Key:  key,
	}
}

// PicPhotoOrAlbumButton 拍照或者相册发图按钮
func PicPhotoOrAlbumButton(name, key string) *Button {
	return &Button{
		Type: ButtonPicPhotoOrAlbum,
		Name: name,
		Key:  key,
	}
}

// PicWeixinButton 微信相册发图按钮
func PicWeixinButton(name, key string) *Button {
	return &Button{
		Type: ButtonPicWeixin,
		Name: name,
		Key:  key,
	}
}

// LocationSelectButton 发送位置按钮
func LocationSelectButton(name, key string) *Button {
	return &Button{
		Type: ButtonLocationSelect,
		Name: name,
		Key:  key,
	}
}

// ViewMinipButton 小程序跳转按钮
func ViewMinipButton(name, appid, pagepath string) *Button {
	return &Button{
		Type:     ButtonViewMinip,
		Name:     name,
		PagePath: pagepath,
		AppID:    appid,
	}
}
