package wxmp

import (
	"encoding/base64"
	"encoding/json"

	"github.com/iiinsomnia/yiigo"
	"go.uber.org/zap"
	"meipian.cn/printapi/wechat"
	"meipian.cn/printapi/wechat/utils"
)

// BizData ...
type BizData interface {
	GetAppID() string
}

// BizDataCrypt ...
type BizDataCrypt struct {
	sessionKey string
}

// Decrypt ...
func (wx *BizDataCrypt) Decrypt(cipherText string, iv string, bizData BizData) error {
	aesKey, err := base64.StdEncoding.DecodeString(wx.sessionKey)

	if err != nil {
		yiigo.Logger.Error("base64.decode wxmp sessionKey error", zap.String("error", err.Error()), zap.String("session_key", wx.sessionKey), zap.String("cipher_text", cipherText), zap.String("iv", iv))

		return err
	}

	aesData, err := base64.StdEncoding.DecodeString(cipherText)

	if err != nil {
		yiigo.Logger.Error("base64.decode wxmp cipherText error", zap.String("error", err.Error()), zap.String("session_key", wx.sessionKey), zap.String("cipher_text", cipherText), zap.String("iv", iv))

		return err
	}

	aesIV, err := base64.StdEncoding.DecodeString(iv)

	if err != nil {
		yiigo.Logger.Error("base64.decode wxmp iv error", zap.String("error", err.Error()), zap.String("session_key", wx.sessionKey), zap.String("cipher_text", cipherText), zap.String("iv", iv))

		return err
	}

	plainText, err := utils.AESCBCDecrypt(aesData, aesKey, aesIV...)

	if err != nil {
		yiigo.Logger.Error("decrypt wxmp cipherText error", zap.String("error", err.Error()), zap.String("session_key", wx.sessionKey), zap.String("cipher_text", cipherText), zap.String("iv", iv))

		return err
	}

	if err := json.Unmarshal(plainText, &bizData); err != nil {
		yiigo.Logger.Error("unmarshal wxmp plainText error", zap.String("error", err.Error()), zap.String("session_key", wx.sessionKey), zap.String("cipher_text", cipherText), zap.String("iv", iv), zap.ByteString("decrypted_data", aesData))

		return err
	}

	if bizData.GetAppID() != wechat.WXMPAppID {
		return utils.ErrIllegaAppID
	}

	return nil
}

// NewWXMPBizDataCrypt ...
func NewWXMPBizDataCrypt(sessionKey string) *BizDataCrypt {
	return &BizDataCrypt{sessionKey: sessionKey}
}

// WaterMark ...
type WaterMark struct {
	Timestamp int64  `json:"timestamp"`
	AppID     string `json:"appid"`
}

// WXUserData ...
type WXUserData struct {
	OpenID    string     `json:"openId"`
	Language  string     `json:"language"`
	City      string     `json:"city"`
	Province  string     `json:"province"`
	AvatarURL string     `json:"avatarUrl"`
	NickName  string     `json:"nickName"`
	Gender    int        `json:"gender"`
	Country   string     `json:"country"`
	UnionID   string     `json:"unionId"`
	Watermark *WaterMark `json:"watermark"`
}

// GetAppID get appid
func (wx *WXUserData) GetAppID() string {
	return wx.Watermark.AppID
}

// WXPhoneData ...
type WXPhoneData struct {
	PhoneNumber     string     `json:"phoneNumber"`
	PurePhoneNumber string     `json:"purePhoneNumber"`
	CountryCode     string     `json:"countryCode"`
	Watermark       *WaterMark `json:"watermark"`
}

// GetAppID get appid
func (wx *WXPhoneData) GetAppID() string {
	return wx.Watermark.AppID
}
