package event

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"

	"github.com/shenghui0779/gochat/wx"
)

// Encrypt 参考微信[加密技术方案](https://open.work.weixin.qq.com/api/doc/90000/90139/90968)
func Encrypt(receiveid, encodingAESKey, nonce string, plainText []byte) ([]byte, error) {
	key, err := base64.StdEncoding.DecodeString(encodingAESKey + "=")

	if err != nil {
		return nil, err
	}

	contentLen := len(plainText)
	appidOffset := 20 + contentLen

	encryptData := make([]byte, appidOffset+len(receiveid))

	copy(encryptData[:16], nonce)
	copy(encryptData[16:20], wx.EncodeUint32ToBytes(uint32(contentLen)))
	copy(encryptData[20:], plainText)
	copy(encryptData[appidOffset:], receiveid)

	cbc := wx.NewCBCCrypto(key, key[:aes.BlockSize], wx.AES_PKCS7)
	cipherText, err := cbc.Encrypt(encryptData)

	if err != nil {
		return nil, err
	}

	return cipherText, nil
}

// Decrypt 参考微信[加密技术方案](https://open.work.weixin.qq.com/api/doc/90000/90139/90968)
func Decrypt(receiveid, encodingAESKey, cipherText string) ([]byte, error) {
	key, err := base64.StdEncoding.DecodeString(encodingAESKey + "=")

	if err != nil {
		return nil, err
	}

	decryptData, err := base64.StdEncoding.DecodeString(cipherText)

	if err != nil {
		return nil, err
	}

	cbc := wx.NewCBCCrypto(key, key[:aes.BlockSize], wx.AES_PKCS7)
	plainText, err := cbc.Decrypt(decryptData)

	if err != nil {
		return nil, err
	}

	appidOffset := len(plainText) - len([]byte(receiveid))

	// 校验 receiveid
	if v := string(plainText[appidOffset:]); v != receiveid {
		return nil, fmt.Errorf("receiveid mismatch, want: %s, got: %s", receiveid, v)
	}

	return plainText[20:appidOffset], nil
}
