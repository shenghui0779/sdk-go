package minip

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

// AuthSession 小程序授权Session
type AuthSession struct {
	SessionKey string `json:"session_key"`
	OpenID     string `json:"openid"`
	UnionID    string `json:"unionid"`
}

// AccessToken 小程序access_token
type AccessToken struct {
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

// Watermark 水印
type Watermark struct {
	Timestamp int64  `json:"timestamp"`
	AppID     string `json:"appid"`
}

// AuthInfo 用户信息
type AuthInfo struct {
	OpenID    string    `json:"openId"`
	Language  string    `json:"language"`
	City      string    `json:"city"`
	Province  string    `json:"province"`
	AvatarURL string    `json:"avatarUrl"`
	Nickname  string    `json:"nickName"`
	Gender    int       `json:"gender"`
	Country   string    `json:"country"`
	UnionID   string    `json:"unionId"`
	Watermark Watermark `json:"watermark"`
}

type ResultPhoneNumber struct {
	PhoneInfo *PhoneInfo `json:"phone_info"`
}

type PhoneInfo struct {
	PhoneNumber     string    `json:"phoneNumber"`     // 用户绑定的手机号（国外手机号会有区号）
	PurePhoneNumber string    `json:"purePhoneNumber"` // 没有区号的手机号
	CountryCode     string    `json:"countryCode"`     // 区号
	Watermark       Watermark `json:"watermark"`       // 数据水印
}

type ParamsPhoneNumber struct {
	Code string `json:"code"`
}

func GetPhoneNumber(code string, result *ResultPhoneNumber) wx.Action {
	params := &ParamsPhoneNumber{
		Code: code,
	}

	return wx.NewPostAction(urls.MinipPhoneNumber,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsEncryptedDataCheck struct {
	EncryptedMsgHash string `json:"encrypted_msg_hash"` // 加密数据的sha256，通过Hex（Base16）编码后的字符串
}

type ResultEncryptedDataCheck struct {
	Valid      bool  `json:"vaild"`       // 是否是合法的数据
	CreateTime int64 `json:"create_time"` // 加密数据生成的时间戳
}

// CheckEncryptedData 用户信息 - 检查加密信息是否由微信生成（当前只支持手机号加密数据），只能检测最近3天生成的加密数据
func CheckEncryptedData(encryptedData string, result *ResultEncryptedDataCheck) wx.Action {
	params := &ParamsEncryptedDataCheck{
		EncryptedMsgHash: wx.SHA256(encryptedData),
	}

	return wx.NewPostAction(urls.MinipEncryptedDataCheck,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// ResultPaidUnionID 支付用户unionid
type ResultPaidUnionID struct {
	UnionID string `json:"unionid"`
}

// GetPaidUnionIDByTransactionID 用户信息 - 用户支付完成后，获取该用户的 UnionId，无需用户授权
func GetPaidUnionIDByTransactionID(openid, transactionID string, result *ResultPaidUnionID) wx.Action {
	return wx.NewGetAction(urls.MinipPaidUnion,
		wx.WithQuery("openid", openid),
		wx.WithQuery("transaction_id", transactionID),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// GetPaidUnionIDByOutTradeNO 用户信息 - 用户支付完成后，获取该用户的 UnionId，无需用户授权
func GetPaidUnionIDByOutTradeNO(openid, mchid, outTradeNO string, result *ResultPaidUnionID) wx.Action {
	return wx.NewGetAction(urls.MinipPaidUnion,
		wx.WithQuery("openid", openid),
		wx.WithQuery("mch_id", mchid),
		wx.WithQuery("out_trade_no", outTradeNO),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
