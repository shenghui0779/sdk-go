package minip

import (
	"encoding/json"

	"github.com/shenghui0779/yiigo"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

// ParamsServiceInvoke 服务调用参数
type ParamsServiceInvoke struct {
	Service     string  `json:"service"`       // 服务ID
	API         string  `json:"api"`           // 接口名
	Data        yiigo.X `json:"data"`          // 服务提供方接口定义的 JSON 格式的数据
	ClientMsgID string  `json:"client_msg_id"` // 随机字符串ID，调用方请求的唯一标识
}

// ResultServiceInvoke 服务调用结果
type ResultServiceInvoke struct {
	Data string `json:"data"`
}

// InvokeService 调用服务平台提供的服务
func InvokeService(params *ParamsServiceInvoke, result *ResultServiceInvoke) wx.Action {
	return wx.NewPostAction(urls.MinipInvokeService,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// ParamsSoterVerify 生物认证秘钥签名验证参数
type ParamsSoterVerify struct {
	OpenID        string `json:"openid"`         // 用户 openid
	JSONString    string `json:"json_string"`    // 通过 wx.startSoterAuthentication 成功回调获得的 resultJSON 字段
	JSONSignature string `json:"json_signature"` // 通过 wx.startSoterAuthentication 成功回调获得的 resultJSONSignature 字段
}

// ResultSoterVerify 生物认证秘钥签名验证结果
type ResultSoterVerify struct {
	IsOK bool `json:"is_ok"`
}

// SoterVerify 生物认证秘钥签名验证
func SoterVerify(params *ParamsSoterVerify, result *ResultSoterVerify) wx.Action {
	return wx.NewPostAction(urls.MinipSoterVerify,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
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

// ParamsUserRisk 用户风控参数
type ParamsUserRisk struct {
	AppID        string    `json:"appid"`                   // 小程序appid
	OpenID       string    `json:"openid"`                  // 用户的openid
	Scene        RiskScene `json:"scene"`                   // 场景值，0:注册，1:营销作弊
	MobileNO     string    `json:"mobile_no,omitempty"`     // 用户手机号
	ClientIP     string    `json:"client_ip"`               // 用户访问源ip
	EmailAddress string    `json:"email_address,omitempty"` // 用户邮箱地址
	ExtendedInfo string    `json:"extended_info,omitempty"` // 额外补充信息
	IsTest       bool      `json:"is_test,omitempty"`       // false：正式调用，true：测试调用
}

// ResultUserRisk 用户风控结果
type ResultUserRisk struct {
	RiskRank int `json:"risk_rank"`
}

// GetUserRiskRank 获取用户的安全等级（无需用户授权）
func GetUserRiskRank(params *ParamsUserRisk, result *ResultUserRisk) wx.Action {
	return wx.NewPostAction(urls.MinipUserRiskRank,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
