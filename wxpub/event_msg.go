package wxpub

import (
	"encoding/base64"
	"encoding/xml"

	"github.com/iiinsomnia/yiigo"
	"go.uber.org/zap"
	"meipian.cn/printapi/wechat"
	"meipian.cn/printapi/wechat/utils"
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

// EncryptMsg 微信公众号加密消息
type EncryptMsg struct {
	ToUserName string `xml:"ToUserName"`
	Encrypt    string `xml:"Encrypt"`
}

// Decrypt 消息解密，参考微信文档[加密解密技术方案](https://open.weixin.qq.com/cgi-bin/showdocument?action=dir_list&t=resource/res_list&verify=1&id=open1419318482&token=&lang=zh_CN)
func (e *EncryptMsg) Decrypt() (*EventMsg, error) {
	settings := wechat.GetSettingsWithChannel(wechat.WXPub)

	key, err := base64.StdEncoding.DecodeString(settings.EncodingAESKey + "=")

	if err != nil {
		yiigo.Logger.Error("base64.decode wxpub EncodingAESKey error", zap.String("error", err.Error()), zap.String("encrypted_msg", e.Encrypt))

		return nil, err
	}

	cipherText, err := base64.StdEncoding.DecodeString(e.Encrypt)

	if err != nil {
		yiigo.Logger.Error("base64.decode wxpub encrypted_msg error", zap.String("error", err.Error()), zap.String("encrypted_msg", e.Encrypt))

		return nil, err
	}

	plainText, err := utils.AESCBCDecrypt(cipherText, key)

	if err != nil {
		yiigo.Logger.Error("decrypt wxpub encrypted_msg error", zap.String("error", err.Error()), zap.String("encrypted_msg", e.Encrypt))

		return nil, err
	}

	appidOffset := len(plainText) - len([]byte(settings.AppID))

	// 校验APPID
	if string(plainText[appidOffset:]) != settings.AppID {
		return nil, utils.ErrIllegaAppID
	}

	msg := new(EventMsg)

	if err := xml.Unmarshal(plainText[20:appidOffset], msg); err != nil {
		yiigo.Logger.Error("unmarshal wxpub decrypted_msg error", zap.String("error", err.Error()), zap.ByteString("decrypted_msg", plainText))

		return nil, err
	}

	return msg, nil
}
