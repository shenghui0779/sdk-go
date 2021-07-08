package mp

import (
	"encoding/json"

	"github.com/tidwall/gjson"

	"github.com/shenghui0779/gochat/wx"
)

// InvokeData 服务调用数据
type InvokeData struct {
	Service     string `json:"service"`       // 服务ID
	API         string `json:"api"`           // 接口名
	Data        wx.X   `json:"data"`          // 服务提供方接口定义的 JSON 格式的数据
	ClientMsgID string `json:"client_msg_id"` // 随机字符串ID，调用方请求的唯一标识
}

type InvokeResult struct {
	Data string `json:"data"`
}

// InvokeService 调用服务平台提供的服务
func InvokeService(dest *InvokeResult, data *InvokeData) wx.Action {
	return wx.NewAction(InvokeServiceURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(data)
		}),
		wx.WithDecode(func(resp []byte) error {
			dest.Data = gjson.GetBytes(resp, "data").String()

			return nil
		}),
	)
}

// SoterSignature 生物认证秘钥签名
type SoterSignature struct {
	OpenID        string `json:"openid"`         // 用户 openid
	JSONString    string `json:"json_string"`    // 通过 wx.startSoterAuthentication 成功回调获得的 resultJSON 字段
	JSONSignature string `json:"json_signature"` // 通过 wx.startSoterAuthentication 成功回调获得的 resultJSONSignature 字段
}

// SoterVerifyResult 生物认证秘钥签名验证结果
type SoterVerifyResult struct {
	OK bool
}

// SoterVerify 生物认证秘钥签名验证
func SoterVerify(dest *SoterVerifyResult, sign *SoterSignature) wx.Action {
	return wx.NewAction(SoterVerifyURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(sign)
		}),
		wx.WithDecode(func(resp []byte) error {
			dest.OK = gjson.GetBytes(resp, "is_ok").Bool()

			return nil
		}),
	)
}

// 风控场景
type RiskScene int

// 微信支持的风控场景值
const (
	RiskRegister RiskScene = 0 // 注册
	RiskCheat    RiskScene = 1 // 营销作弊
)

// UserRiskData 用户风控数据
type UserRiskData struct {
	AppID        string    `json:"appid"`                   // 小程序appid
	OpenID       string    `json:"openid"`                  // 用户的openid
	Scene        RiskScene `json:"scene"`                   // 场景值，0:注册，1:营销作弊
	MobileNO     string    `json:"mobile_no,omitempty"`     // 用户手机号
	ClientIP     string    `json:"client_ip"`               // 用户访问源ip
	EmailAddress string    `json:"email_address,omitempty"` // 用户邮箱地址
	ExtendedInfo string    `json:"extended_info,omitempty"` // 额外补充信息
	IsTest       bool      `json:"is_test,omitempty"`       // false：正式调用，true：测试调用
}

// UserRiskRank 用户风控结果
type UserRiskResult struct {
	RiskRank int
}

// GetUserRiskRank 获取用户的安全等级（无需用户授权）
func GetUserRiskRank(dest *UserRiskResult, data *UserRiskData) wx.Action {
	return wx.NewAction(UserRiskRankURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(data)
		}),
		wx.WithDecode(func(resp []byte) error {
			dest.RiskRank = int(gjson.GetBytes(resp, "risk_rank").Int())

			return nil
		}),
	)
}
