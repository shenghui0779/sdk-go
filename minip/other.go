package minip

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

// ParamsServiceInvoke 服务调用参数
type ParamsServiceInvoke struct {
	Service     string `json:"service"`       // 服务ID
	API         string `json:"api"`           // 接口名
	Data        wx.M   `json:"data"`          // 服务提供方接口定义的 JSON 格式的数据
	ClientMsgID string `json:"client_msg_id"` // 随机字符串ID，调用方请求的唯一标识
}

// ResultServiceInvoke 服务调用结果
type ResultServiceInvoke struct {
	Data string `json:"data"`
}

// InvokeService 服务市场 - 调用服务平台提供的服务
func InvokeService(params *ParamsServiceInvoke, result *ResultServiceInvoke) wx.Action {
	return wx.NewPostAction(urls.MinipInvokeService,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
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

// SoterVerify 生物认证 - 生物认证秘钥签名验证
func SoterVerify(openID, jsonStr, jsonSign string, result *ResultSoterVerify) wx.Action {
	params := &ParamsSoterVerify{
		OpenID:        openID,
		JSONString:    jsonStr,
		JSONSignature: jsonSign,
	}

	return wx.NewPostAction(urls.MinipSoterVerify,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsShortLink struct {
	PageURL     string `json:"page_url"`
	PageTitle   string `json:"page_title"`
	ISPermanent bool   `json:"is_permanent"`
}

type ResultShortLink struct {
	Link string `json:"link"`
}

// GenerateShortLink Short Link - 获取小程序 Short Link，适用于微信内拉起小程序的业务场景。目前只开放给电商类目(具体包含以下一级类目：电商平台、商家自营、跨境电商)。
func GenerateShortLink(pageURL, pageTitle string, isPermanent bool, result *ResultShortLink) wx.Action {
	params := &ParamsShortLink{
		PageURL:     pageURL,
		PageTitle:   pageTitle,
		ISPermanent: isPermanent,
	}

	return wx.NewPostAction(urls.MinipShortLink,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// RiskScene 风控场景
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

// GetUserRiskRank 安全风控 - 获取用户的安全等级（无需用户授权）
func GetUserRiskRank(params *ParamsUserRisk, result *ResultUserRisk) wx.Action {
	return wx.NewPostAction(urls.MinipUserRiskRank,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsSchemeGenerate struct {
	JumpWxa        *SchemeJumpWxa `json:"jump_wxa,omitempty"`
	IsExpire       bool           `json:"is_expire,omitempty"`
	ExpireType     int            `json:"expire_type,omitempty"`
	ExpireTime     int64          `json:"expire_time,omitempty"`
	ExpireInterval int            `json:"expire_interval,omitempty"`
}

type SchemeJumpWxa struct {
	Path       string     `json:"path,omitempty"`
	Query      string     `json:"query,omitempty"`
	EnvVersion EnvVersion `json:"env_version,omitempty"`
}

type ResultSchemeGenerate struct {
	OpenLink string `json:"openlink"`
}

// GenerateScheme 获取小程序 scheme 码，适用于短信、邮件、外部网页、微信内等拉起小程序的业务场景。
func GenerateScheme(params *ParamsSchemeGenerate, result *ResultSchemeGenerate) wx.Action {
	return wx.NewPostAction(urls.MinipGenerateScheme,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsSchemeQuery struct {
	Scheme string `json:"scheme"`
}

type ResultSchemeQuery struct {
	SchemeInfo  *SchemeInfo  `json:"scheme_info"`
	SchemeQuota *SchemeQuota `json:"scheme_quota"`
}

type SchemeInfo struct {
	AppID      string     `json:"appid"`
	Path       string     `json:"path"`
	Query      string     `json:"query"`
	CreateTime int64      `json:"create_time"`
	ExpireTime int64      `json:"expire_time"`
	EnvVersion EnvVersion `json:"env_version"`
}

type SchemeQuota struct {
	LongTimeUsed  int `json:"long_time_used"`
	LongTimeLimit int `json:"long_time_limit"`
}

// QueryScheme 查询小程序 scheme 码，及长期有效 quota。
func QueryScheme(scheme string, result *ResultSchemeQuery) wx.Action {
	params := &ParamsSchemeQuery{
		Scheme: scheme,
	}

	return wx.NewPostAction(urls.MinipQueryScheme,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type CloudBase struct {
	Env           string `json:"env"`
	Domain        string `json:"domain,omitempty"`
	Path          string `json:"path,omitempty"`
	Query         string `json:"query,omitempty"`
	ResourceAppID string `json:"resource_appid,omitempty"`
}

type ParamsURLLinkGenerate struct {
	Path           string     `json:"path,omitempty"`
	Query          string     `json:"query,omitempty"`
	IsExpire       bool       `json:"is_expire,omitempty"`
	ExpireType     int        `json:"expire_type,omitempty"`
	ExpireTime     int64      `json:"expire_time,omitempty"`
	ExpireInterval int        `json:"expire_interval,omitempty"`
	EnvVersion     EnvVersion `json:"env_version,omitempty"`
	CloudBase      *CloudBase `json:"cloud_base,omitempty"`
}

type ResultURLLinkGenerate struct {
	URLLink string `json:"url_link"`
}

// GenerateURLLink 获取小程序 URL Link，适用于短信、邮件、网页、微信内等拉起小程序的业务场景。
func GenerateURLLink(params *ParamsURLLinkGenerate, result *ResultURLLinkGenerate) wx.Action {
	return wx.NewPostAction(urls.MinipGenerateURLLink,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsURLLinkQuery struct {
	URLLink string `json:"url_link"`
}

type ResultURLLinkQuery struct {
	URLLinkInfo  *URLLinkInfo  `json:"url_link_info"`
	URLLinkQuota *URLLinkQuota `json:"url_link_quota"`
}

type URLLinkInfo struct {
	AppID      string     `json:"appid"`
	Path       string     `json:"path"`
	Query      string     `json:"query"`
	CreateTime int64      `json:"create_time"`
	ExpireTime int64      `json:"expire_time"`
	EnvVersion EnvVersion `json:"env_version"`
	CloudBase  *CloudBase `json:"cloud_base"`
}

type URLLinkQuota struct {
	LongTimeUsed  int `json:"long_time_used"`
	LongTimeLimit int `json:"long_time_limit"`
}

// QueryURLLink 查询小程序 url_link 配置，及长期有效 quota。
func QueryURLLink(urllink string, result *ResultURLLinkQuery) wx.Action {
	params := &ParamsURLLinkQuery{
		URLLink: urllink,
	}

	return wx.NewPostAction(urls.MinipQueryURLLink,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
