package mch

import (
	"strconv"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

// ParamsRedpack 红包参数
type ParamsRedpack struct {
	// 必填参数
	MchBillNO   string // 商户订单号（每个订单号必须唯一。取值范围：0~9，a~z，A~Z）接口根据商户订单号支持重入，如出现超时可再调用
	SendName    string // 红包发送者名称；注意：敏感词会被转义成字符*
	ReOpenID    string // 接受红包的用户openid
	TotalAmount int    // 付款金额，单位：分
	TotalNum    int    // 红包发放总人数
	Wishing     string // 红包祝福语；注意：敏感词会被转义成字符*
	ClientIP    string // 调用接口的机器Ip地址
	ActName     string // 活动名称；注意：敏感词会被转义成字符*
	Remark      string // 备注信息
	// 选填参数
	SceneID  string // 发放红包使用场景，红包金额大于200或者小于1元时必传
	RiskInfo string // 活动信息，urlencode(posttime=xx&mobile=xx&deviceid=xx。posttime：用户操作的时间戳；mobile：业务系统账号的手机号，国家代码-手机号，不需要+号；deviceid：MAC地址或者设备唯一标识；clientversion：用户操作的客户端版本
}

// SendNormalRedpack 发放普通红包（需要证书）
// 注意：当返回错误码为“SYSTEMERROR”时，请务必使用原商户订单号重试，否则可能造成重复支付等资金风险。
func SendNormalRedpack(appid string, params *ParamsRedpack) wx.Action {
	return wx.NewPostAction(urls.MchRedpackNormal,
		wx.WithTLS(),
		wx.WithWXML(func(mchid, apikey, nonce string) (wx.WXML, error) {
			m := wx.WXML{
				"wxappid":      appid,
				"mch_id":       mchid,
				"nonce_str":    nonce,
				"mch_billno":   params.MchBillNO,
				"send_name":    params.SendName,
				"re_openid":    params.ReOpenID,
				"total_amount": strconv.Itoa(params.TotalAmount),
				"total_num":    strconv.Itoa(params.TotalNum),
				"wishing":      params.Wishing,
				"client_ip":    params.ClientIP,
				"act_name":     params.ActName,
				"remark":       params.Remark,
			}

			if params.SceneID != "" {
				m["scene_id"] = params.SceneID
			}

			if params.RiskInfo != "" {
				m["risk_info"] = params.RiskInfo
			}

			// 签名
			m["sign"] = wx.SignMD5.Do(apikey, m, true)

			return m, nil
		}),
	)
}

// SendGroupRedpack 发放裂变红包（需要证书）
// 注意：当返回错误码为“SYSTEMERROR”时，请务必使用原商户订单号重试，否则可能造成重复支付等资金风险。
func SendGroupRedpack(appid string, params *ParamsRedpack) wx.Action {
	return wx.NewPostAction(urls.MchRedpackGroup,
		wx.WithTLS(),
		wx.WithWXML(func(mchid, apikey, nonce string) (wx.WXML, error) {
			m := wx.WXML{
				"wxappid":      appid,
				"mch_id":       mchid,
				"nonce_str":    nonce,
				"mch_billno":   params.MchBillNO,
				"send_name":    params.SendName,
				"re_openid":    params.ReOpenID,
				"total_amount": strconv.Itoa(params.TotalAmount),
				"total_num":    strconv.Itoa(params.TotalNum),
				"amt_type":     "ALL_RAND",
				"wishing":      params.Wishing,
				"act_name":     params.ActName,
				"remark":       params.Remark,
			}

			if params.SceneID != "" {
				m["scene_id"] = params.SceneID
			}

			if params.RiskInfo != "" {
				m["risk_info"] = params.RiskInfo
			}

			// 签名
			m["sign"] = wx.SignMD5.Do(apikey, m, true)

			return m, nil
		}),
	)
}

// SendMinipRedpack 发放小程序红包（需要证书）
// 注意：当返回错误码为“SYSTEMERROR”时，请务必使用原商户订单号重试，否则可能造成重复支付等资金风险。
func SendMinipRedpack(appid string, params *ParamsRedpack) wx.Action {
	return wx.NewPostAction(urls.MchRedpackMinip,
		wx.WithTLS(),
		wx.WithWXML(func(mchid, apikey, nonce string) (wx.WXML, error) {
			m := wx.WXML{
				"wxappid":      appid,
				"mch_id":       mchid,
				"nonce_str":    nonce,
				"mch_billno":   params.MchBillNO,
				"send_name":    params.SendName,
				"re_openid":    params.ReOpenID,
				"total_amount": strconv.Itoa(params.TotalAmount),
				"total_num":    strconv.Itoa(params.TotalNum),
				"wishing":      params.Wishing,
				"act_name":     params.ActName,
				"remark":       params.Remark,
				"notify_way":   "MINI_PROGRAM_JSAPI",
			}

			if params.SceneID != "" {
				m["scene_id"] = params.SceneID
			}

			// 签名
			m["sign"] = wx.SignMD5.Do(apikey, m, true)

			return m, nil
		}),
	)
}

// QueryRedpack 查询红包记录（需要证书）
func QueryRedpack(appid, billNO string) wx.Action {
	return wx.NewPostAction(urls.MchRedpackQuery,
		wx.WithTLS(),
		wx.WithWXML(func(mchid, apikey, nonce string) (wx.WXML, error) {
			m := wx.WXML{
				"appid":      appid,
				"mch_id":     mchid,
				"mch_billno": billNO,
				"bill_type":  "MCHT",
				"nonce_str":  nonce,
			}

			// 签名
			m["sign"] = wx.SignMD5.Do(apikey, m, true)

			return m, nil
		}),
	)
}
