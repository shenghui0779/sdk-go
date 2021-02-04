package mch

import (
	"strconv"

	"github.com/shenghui0779/gochat/wx"
)

// RedpackData 红包发放数据
type RedpackData struct {
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

// SendNormalRedpack 发放普通红包
func SendNormalRedpack(data *RedpackData) wx.Action {
	return wx.NewAction(RedpackNormalURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithTLS(),
		wx.WithWXML(func(appid, mchid, nonce string) (wx.WXML, error) {
			body := wx.WXML{
				"wxappid":      appid,
				"mch_id":       mchid,
				"nonce_str":    nonce,
				"mch_billno":   data.MchBillNO,
				"send_name":    data.SendName,
				"re_openid":    data.ReOpenID,
				"total_amount": strconv.Itoa(data.TotalAmount),
				"total_num":    strconv.Itoa(data.TotalNum),
				"wishing":      data.Wishing,
				"client_ip":    data.ClientIP,
				"act_name":     data.ActName,
				"remark":       data.Remark,
				"sign_type":    SignMD5,
			}

			if data.SceneID != "" {
				body["scene_id"] = data.SceneID
			}

			if data.RiskInfo != "" {
				body["risk_info"] = data.RiskInfo
			}

			return body, nil
		}),
	)
}

// SendGroupRedpack 发放裂变红包
func SendGroupRedpack(data *RedpackData) wx.Action {
	return wx.NewAction(RedpackGroupURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithTLS(),
		wx.WithWXML(func(appid, mchid, nonce string) (wx.WXML, error) {
			body := wx.WXML{
				"wxappid":      appid,
				"mch_id":       mchid,
				"nonce_str":    nonce,
				"mch_billno":   data.MchBillNO,
				"send_name":    data.SendName,
				"re_openid":    data.ReOpenID,
				"total_amount": strconv.Itoa(data.TotalAmount),
				"total_num":    strconv.Itoa(data.TotalNum),
				"amt_type":     "ALL_RAND",
				"wishing":      data.Wishing,
				"act_name":     data.ActName,
				"remark":       data.Remark,
				"sign_type":    SignMD5,
			}

			if data.SceneID != "" {
				body["scene_id"] = data.SceneID
			}

			if data.RiskInfo != "" {
				body["risk_info"] = data.RiskInfo
			}

			return body, nil
		}),
	)
}

// SendMinipRedpack 发放小程序红包
func SendMinipRedpack(data *RedpackData) wx.Action {
	return wx.NewAction(RedpackMinipURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithTLS(),
		wx.WithWXML(func(appid, mchid, nonce string) (wx.WXML, error) {
			body := wx.WXML{
				"wxappid":      appid,
				"mch_id":       mchid,
				"nonce_str":    nonce,
				"mch_billno":   data.MchBillNO,
				"send_name":    data.SendName,
				"re_openid":    data.ReOpenID,
				"total_amount": strconv.Itoa(data.TotalAmount),
				"total_num":    strconv.Itoa(data.TotalNum),
				"wishing":      data.Wishing,
				"act_name":     data.ActName,
				"remark":       data.Remark,
				"notify_way":   "MINI_PROGRAM_JSAPI",
				"sign_type":    SignMD5,
			}

			if data.SceneID != "" {
				body["scene_id"] = data.SceneID
			}

			return body, nil
		}),
	)
}

// QueryRedpackByBillNO 查询红包记录
func QueryRedpackByBillNO(billNO string) wx.Action {
	return wx.NewAction(RedpackQueryURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithTLS(),
		wx.WithWXML(func(appid, mchid, nonce string) (wx.WXML, error) {
			return wx.WXML{
				"appid":      appid,
				"mch_id":     mchid,
				"mch_billno": billNO,
				"bill_type":  "MCHT",
				"nonce_str":  nonce,
				"sign_type":  SignMD5,
			}, nil
		}),
	)
}
