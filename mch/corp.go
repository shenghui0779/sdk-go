package mch

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

// ParamsCorpRedpack 企业红包参数
type ParamsCorpRedpack struct {
	// 必填参数
	MchBillNO   string // 商户订单号（每个订单号必须唯一。取值范围：0~9，a~z，A~Z）接口根据商户订单号支持重入，如出现超时可再调用
	ReOpenID    string // 接受红包的用户openid
	TotalAmount int    // 付款金额，单位：分，单笔最小金额默认为1元
	Wishing     string // 红包祝福语
	ActName     string // 活动名称
	Remark      string // 备注信息
	// 选填参数
	AgentID             int64  // 以企业应用的名义发红包，企业应用id，整型，可在企业微信管理端应用的设置页面查看。与sender_name互斥，二者只能填一个
	SenderName          string // 以个人名义发红包，红包发送者名称(需要utf-8格式)。与agentid互斥，二者只能填一个
	SenderHeaderMediaID string // 发送者头像素材id，通过企业微信开放上传素材接口获取
	SceneID             string // 发放红包使用场景，红包金额大于200或者小于1元时必传
}

type ParamsCorpTransfer struct {
	// 必填参数
	PartnerTradeNO string
	OpenID         string
	CheckName      string
	Amount         int
	Desc           string
	SpbillCreateIP string
	WWMsgType      string
	ActName        string
	// 选填参数
	AgentID        int64
	ReUserName     string
	ApprovalType   int
	ApprovalNumber string
	DeviceInfo     string
}

// SendCorpRedpack 发放企业红包（需要证书）
// 注意：当返回错误码为“SYSTEMERROR”时，请务必使用原商户订单号重试，否则可能造成重复支付等资金风险。
func SendCorpRedpack(appid, secret string, params *ParamsCorpRedpack) wx.Action {
	return wx.NewPostAction(urls.MchRedpackCorp,
		wx.WithTLS(),
		wx.WithWXML(func(mchid, apikey, nonce string) (wx.WXML, error) {
			m := wx.WXML{
				"wxappid":      appid,
				"mch_id":       mchid,
				"nonce_str":    nonce,
				"mch_billno":   params.MchBillNO,
				"re_openid":    params.ReOpenID,
				"total_amount": strconv.Itoa(params.TotalAmount),
				"wishing":      params.Wishing,
				"act_name":     params.ActName,
				"remark":       params.Remark,
			}

			signStr := fmt.Sprintf("act_name=%s&mch_billno=%s&mch_id=%s&nonce_str=%s&re_openid=%s&total_amount=%d&wxappid=%s&secret=%s",
				params.ActName,
				params.MchBillNO,
				mchid, nonce,
				params.ReOpenID,
				params.TotalAmount,
				appid,
				secret,
			)

			m["workwx_sign"] = strings.ToUpper(wx.MD5(signStr))

			if params.AgentID != 0 {
				m["agentid"] = strconv.FormatInt(params.AgentID, 10)
			}

			if len(params.SenderName) != 0 {
				m["sender_name"] = params.SenderName
			}

			if len(params.SenderHeaderMediaID) != 0 {
				m["sender_header_media_id"] = params.SenderHeaderMediaID
			}

			if len(params.SceneID) != 0 {
				m["scene_id"] = params.SceneID
			}

			// 签名
			m["sign"] = wx.SignMD5.Sign(apikey, m, true)

			return m, nil
		}),
	)
}

// QueryCorpRedpack 查询企业红包记录（需要证书）
func QueryCorpRedpack(appid, billNO string) wx.Action {
	return wx.NewPostAction(urls.MchRedpackCorpQuery,
		wx.WithTLS(),
		wx.WithWXML(func(mchid, apikey, nonce string) (wx.WXML, error) {
			m := wx.WXML{
				"appid":      appid,
				"mch_id":     mchid,
				"mch_billno": billNO,
				"nonce_str":  nonce,
			}

			// 签名
			m["sign"] = wx.SignMD5.Sign(apikey, m, true)

			return m, nil
		}),
	)
}

// TransferToPocket 向员工付款（需要证书）
// 注意：当返回错误码为“SYSTEMERROR”时，请务必使用原商户订单号重试，否则可能造成重复支付等资金风险。
func TransferToPocket(appid, secret string, params *ParamsCorpTransfer) wx.Action {
	return wx.NewPostAction(urls.MchTransferToPocket,
		wx.WithTLS(),
		wx.WithWXML(func(mchid, apikey, nonce string) (wx.WXML, error) {
			m := wx.WXML{
				"appid":            appid,
				"mch_id":           mchid,
				"nonce_str":        nonce,
				"partner_trade_no": params.PartnerTradeNO,
				"openid":           params.OpenID,
				"check_name":       params.CheckName,
				"amount":           strconv.Itoa(params.Amount),
				"desc":             params.Desc,
				"spbill_create_ip": params.SpbillCreateIP,
				"ww_msg_type":      params.WWMsgType,
				"act_name":         params.ActName,
			}

			signStr := fmt.Sprintf("amount=%d&appid=%s&desc=%s&mch_id=%s&nonce_str=%s&openid=%s&partner_trade_no=%s&ww_msg_type=%s&secret=%s",
				params.Amount,
				appid,
				params.Desc,
				mchid,
				nonce,
				params.OpenID,
				params.PartnerTradeNO,
				params.WWMsgType,
				secret,
			)

			m["workwx_sign"] = strings.ToUpper(wx.MD5(signStr))

			if params.AgentID != 0 {
				m["agentid"] = strconv.FormatInt(params.AgentID, 10)
			}

			if len(params.ReUserName) != 0 {
				m["re_user_name"] = params.ReUserName
			}

			if params.ApprovalType != 0 {
				m["approval_type"] = strconv.Itoa(params.ApprovalType)
			}

			if len(params.ApprovalNumber) != 0 {
				m["approval_number"] = params.ApprovalNumber
			}

			if len(params.DeviceInfo) != 0 {
				m["device_info"] = params.DeviceInfo
			}

			// 签名
			m["sign"] = wx.SignMD5.Sign(apikey, m, true)

			return m, nil
		}),
	)
}

// QueryTransferPocket 查询向员工付款结果（需要证书）
func QueryTransferPocket(appid, partnerTradeNO string) wx.Action {
	return wx.NewPostAction(urls.MchTransferPocketOrderQuery,
		wx.WithTLS(),
		wx.WithWXML(func(mchid, apikey, nonce string) (wx.WXML, error) {
			m := wx.WXML{
				"appid":            appid,
				"mch_id":           mchid,
				"partner_trade_no": partnerTradeNO,
				"nonce_str":        nonce,
			}

			// 签名
			m["sign"] = wx.SignMD5.Sign(apikey, m, true)

			return m, nil
		}),
	)
}
