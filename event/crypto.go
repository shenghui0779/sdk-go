package event

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"

	"github.com/shenghui0779/gochat/wx"
)

// Encrypt 参考微信[加密技术方案](https://open.weixin.qq.com/cgi-bin/showdocument?action=dir_list&t=resource/res_list&verify=1&id=open1419318482&token=&lang=zh_CN)
func Encrypt(appid, encodingAESKey, nonce string, plainText []byte) ([]byte, error) {
	key, err := base64.StdEncoding.DecodeString(encodingAESKey + "=")

	if err != nil {
		return nil, err
	}

	contentLen := len(plainText)
	appidOffset := 20 + contentLen

	encryptData := make([]byte, appidOffset+len(appid))

	copy(encryptData[:16], nonce)
	copy(encryptData[16:20], wx.EncodeUint32ToBytes(uint32(contentLen)))
	copy(encryptData[20:], plainText)
	copy(encryptData[appidOffset:], appid)

	cbc := wx.NewCBCCrypto(key, key[:aes.BlockSize], wx.PKCS7)
	cipherText, err := cbc.Encrypt(encryptData)

	if err != nil {
		return nil, err
	}

	return cipherText, nil
}

// Decrypt 参考微信[加密技术方案](https://open.weixin.qq.com/cgi-bin/showdocument?action=dir_list&t=resource/res_list&verify=1&id=open1419318482&token=&lang=zh_CN)
func Decrypt(appid, encodingAESKey, cipherText string) ([]byte, error) {
	key, err := base64.StdEncoding.DecodeString(encodingAESKey + "=")

	if err != nil {
		return nil, err
	}

	decryptData, err := base64.StdEncoding.DecodeString(cipherText)

	if err != nil {
		return nil, err
	}

	cbc := wx.NewCBCCrypto(key, key[:aes.BlockSize], wx.PKCS7)
	plainText, err := cbc.Decrypt(decryptData)

	if err != nil {
		return nil, err
	}

	appidOffset := len(plainText) - len([]byte(appid))

	// 校验 AppID
	if v := string(plainText[appidOffset:]); v != appid {
		return nil, fmt.Errorf("appid mismatch, want: %s, got: %s", appid, v)
	}

	return plainText[20:appidOffset], nil
}
