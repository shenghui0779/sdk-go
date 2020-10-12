package pub

import (
	"crypto/aes"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"

	"github.com/shenghui0779/gochat/utils"
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

// MsgCrypto 公众号消息解析
type MsgCrypto struct {
	pub        *WXPub
	cipherText string
	body       []byte
}

// Decrypt 消息解密，参考微信[加密解密技术方案](https://open.weixin.qq.com/cgi-bin/showdocument?action=dir_list&t=resource/res_list&verify=1&id=open1419318482&token=&lang=zh_CN)
func (m *MsgCrypto) Decrypt() error {
	key, err := base64.StdEncoding.DecodeString(m.pub.encodingAESKey + "=")

	if err != nil {
		return err
	}

	cipherText, err := base64.StdEncoding.DecodeString(m.cipherText)

	if err != nil {
		return err
	}

	cbc := utils.NewAESCBCCrypto(key, key[:aes.BlockSize])
	plainText, err := cbc.Decrypt(cipherText, utils.PKCS7)

	if err != nil {
		return err
	}

	appidOffset := len(plainText) - len([]byte(m.pub.appid))

	// 校验 AppID
	if appid := string(plainText[appidOffset:]); appid != m.pub.appid {
		return fmt.Errorf("gochat: wxpub msg appid mismatch, want: %s, got: %s", m.pub.appid, appid)
	}

	m.body = plainText[20:appidOffset]

	return nil
}

// EventMsg 获取事件消息
func (m *MsgCrypto) EventMsg() (*EventMsg, error) {
	if len(m.body) == 0 {
		return nil, errors.New("gochat: empty msg (forgotten decrypted?)")
	}

	msg := new(EventMsg)

	if err := xml.Unmarshal(m.body, msg); err != nil {
		return nil, err
	}

	return msg, nil
}
