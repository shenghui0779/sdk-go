package mp

import (
	"encoding/base64"
	"encoding/json"

	"github.com/iiinsomnia/gochat/utils"
)

// UserData ...
type UserData struct {
	OpenID    string     `json:"openId"`
	Language  string     `json:"language"`
	City      string     `json:"city"`
	Province  string     `json:"province"`
	AvatarURL string     `json:"avatarUrl"`
	NickName  string     `json:"nickName"`
	Gender    int        `json:"gender"`
	Country   string     `json:"country"`
	UnionID   string     `json:"unionId"`
	WaterMark *WaterMark `json:"watermark"`
}

// PhoneData ...
type PhoneData struct {
	PhoneNumber     string     `json:"phoneNumber"`
	PurePhoneNumber string     `json:"purePhoneNumber"`
	CountryCode     string     `json:"countryCode"`
	WaterMark       *WaterMark `json:"watermark"`
}

// WaterMark ...
type WaterMark struct {
	Timestamp int64  `json:"timestamp"`
	AppID     string `json:"appid"`
}

// BizData ...
type BizDataCrypt struct {
	appID string
	data  []byte
}

// Decrypt ...
func (b *BizDataCrypt) Decrypt(sessionKey, iv, bizData string) error {
	aesKey, err := base64.StdEncoding.DecodeString(sessionKey)

	if err != nil {
		return err
	}

	aesIV, err := base64.StdEncoding.DecodeString(iv)

	if err != nil {
		return err
	}

	aesData, err := base64.StdEncoding.DecodeString(bizData)

	if err != nil {
		return err
	}

	plainText, err := utils.AESCBCDecrypt(aesData, aesKey, aesIV...)

	if err != nil {
		return err
	}

	b.data = plainText

	return nil
}

func (b *BizDataCrypt) GetUserData() (*UserData, error) {
	userData := new(UserData)

	if err := json.Unmarshal(b.data, &userData); err != nil {
		return nil, err
	}

	if userData.WaterMark.AppID != b.appID {
		return nil, utils.ErrIllegaAppID
	}

	return userData, nil
}

func (b *BizDataCrypt) GetPhoneData() (*PhoneData, error) {
	phoneData := new(PhoneData)

	if err := json.Unmarshal(b.data, &phoneData); err != nil {
		return nil, err
	}

	if phoneData.WaterMark.AppID != b.appID {
		return nil, utils.ErrIllegaAppID
	}

	return phoneData, nil
}
