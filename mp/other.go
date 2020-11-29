package mp

import "github.com/shenghui0779/gochat/wx"

// ServiceData 服务数据
type ServiceData struct {
	ServiceID   string `json:"service"`       // 服务ID
	API         string `json:"api"`           // 接口名
	Data        wx.X   `json:"data"`          // 服务提供方接口定义的 JSON 格式的数据
	ClientMsgID string `json:"client_msg_id"` // 随机字符串 ID，调用方请求的唯一标识
}

// InvokeService 调用服务平台提供的服务
func InvokeService() {

}

type SoterSignature struct {
	OpenID        string `json:"open_id"`        // 用户 openid
	JSONString    string `json:"json_string"`    // 通过 wx.startSoterAuthentication 成功回调获得的 resultJSON 字段
	JSONSignature string `json:"json_signature"` // 通过 wx.startSoterAuthentication 成功回调获得的 resultJSONSignature 字段
}

// SoterVerify 生物认证秘钥签名验证
func SoterVerify() {

}
