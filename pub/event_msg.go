package pub

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"

	"github.com/iiinsomnia/gochat/utils"
)

// EventMsg 微信公众号事件消息
type EventMsg struct {
	ToUserName   string  `xml:"ToUserName"`
	FromUserName string  `xml:"FromUserName"`
	CreateTime   int64   `xml:"CreateTime"`
	MsgType      string  `xml:"MsgType"`
	MsgID        int64   `xml:"MsgId"`
	Content      string  `xml:"Content"`
	PicURL       string  `xml:"PicUrl"`
	MediaID      string  `xml:"MediaId"`
	Format       string  `xml:"Format"`
	Recognition  string  `xml:"Recognition"`
	ThumbMediaID string  `xml:"ThumbMediaId"`
	LocationX    float64 `xml:"Location_X"`
	LocationY    float64 `xml:"Location_Y"`
	Scale        int     `xml:"Scale"`
	Label        string  `xml:"Label"`
	Title        string  `xml:"Title"`
	Description  string  `xml:"Description"`
	URL          string  `xml:"Url"`
	Event        string  `xml:"Event"`
	EventKey     string  `xml:"EventKey"`
	Ticket       string  `xml:"Ticket"`
	Latitude     float64 `xml:"Latitude"`
	Longitude    float64 `xml:"Longitude"`
	Precision    float64 `xml:"Precision"`
}

type MsgChiper struct {
	*WXPub
}

// Decrypt 消息解密，参考微信[加密解密技术方案](https://open.weixin.qq.com/cgi-bin/showdocument?action=dir_list&t=resource/res_list&verify=1&id=open1419318482&token=&lang=zh_CN)
func (c *MsgChiper) Decrypt(encrypt string) (*EventMsg, error) {
	key, err := base64.StdEncoding.DecodeString(c.EncodingAESKey + "=")

	if err != nil {
		return nil, err
	}

	cipherText, err := base64.StdEncoding.DecodeString(encrypt)

	if err != nil {
		return nil, err
	}

	plainText, err := utils.AESCBCDecrypt(cipherText, key)

	if err != nil {
		return nil, err
	}

	appidOffset := len(plainText) - len([]byte(c.AppID))

	// 校验APPID
	if appid := string(plainText[appidOffset:]); appid != c.AppID {
		return nil, fmt.Errorf("wxpub event_msg appid mismatch, want: %s, got: %s", c.AppID, appid)
	}

	m := new(EventMsg)

	if err := xml.Unmarshal(plainText[20:appidOffset], m); err != nil {
		return nil, err
	}

	return m, nil
}
