package mch

import (
	"strconv"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

// ParamsTransferBalance 付款到零钱参数
type ParamsTransferBalance struct {
	// 必填参数
	PartnerTradeNO string // 商户订单号，需保持唯一性 (只能是字母或者数字，不能包含有其它字符)
	OpenID         string // 商户appid下，某用户的openid
	CheckName      string // NO_CHECK：不校验真实姓名；FORCE_CHECK：强校验真实姓名
	Amount         int    // 企业付款金额，单位：分
	Desc           string // 企业付款备注，必填。注意：备注中的敏感词会被转成字符*
	// 选填参数
	ReUserName     string // 收款用户真实姓名。如果check_name设置为FORCE_CHECK，则必填用户真实姓名
	DeviceInfo     string // 微信支付分配的终端设备号
	SpbillCreateIP string // 该IP同在商户平台设置的IP白名单中的IP没有关联，该IP可传用户端或者服务端的IP
}

// ParamsTransferBankCard 付款到银行卡参数
type ParamsTransferBankCard struct {
	// 必填参数
	PartnerTradeNO string // 商户订单号，需保持唯一（只允许数字[0~9]或字母[action~Z]和[a~z]，最短8位，最长32位）
	EncBankNO      string // 收款方银行卡号（采用标准RSA_PKCS1_OAEP加密，公钥由微信侧提供，参考：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay_yhk.php?chapter=25_7&index=4）
	EncTrueName    string // 收款方用户名（采用标准RSA_PKCS1_OAEP加密，公钥由微信侧提供，参考：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay_yhk.php?chapter=25_7&index=4）
	BankCode       string // 银行卡所在开户行编号，参考：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_4
	Amount         int    // 付款金额：RMB分（支付总额，不含手续费）注：大于0的整数
	// 选填参数
	Desc string // 企业付款到银行卡付款说明，即订单备注（UTF8编码，允许100个字符以内）
}

// TransferToBalance 付款到零钱（需要证书）
// 注意：当返回错误码为“SYSTEMERROR”时，请务必使用原商户订单号重试，否则可能造成重复支付等资金风险。
func TransferToBalance(appid string, params *ParamsTransferBalance) wx.Action {
	return wx.NewPostAction(urls.MchTransferToBalance,
		wx.WithTLS(),
		wx.WithWXML(func(mchid, apikey, nonce string) (wx.WXML, error) {
			m := wx.WXML{
				"mch_appid":        appid,
				"mchid":            mchid,
				"nonce_str":        nonce,
				"partner_trade_no": params.PartnerTradeNO,
				"openid":           params.OpenID,
				"check_name":       params.CheckName,
				"amount":           strconv.Itoa(params.Amount),
				"desc":             params.Desc,
			}

			if params.ReUserName != "" {
				m["re_user_name"] = params.ReUserName
			}

			if params.DeviceInfo != "" {
				m["device_info"] = params.DeviceInfo
			}

			if params.SpbillCreateIP != "" {
				m["spbill_create_ip"] = params.SpbillCreateIP
			}

			// 签名
			m["sign"] = wx.SignMD5.Do(apikey, m, true)

			return m, nil
		}),
	)
}

// QueryTransferBalance 查询付款到零钱结果（需要证书）
func QueryTransferBalance(appid, partnerTradeNO string) wx.Action {
	return wx.NewPostAction(urls.MchTransferBalanceOrderQuery,
		wx.WithTLS(),
		wx.WithWXML(func(mchid, apikey, nonce string) (wx.WXML, error) {
			m := wx.WXML{
				"appid":            appid,
				"mch_id":           mchid,
				"partner_trade_no": partnerTradeNO,
				"nonce_str":        nonce,
			}

			// 签名
			m["sign"] = wx.SignMD5.Do(apikey, m, true)

			return m, nil
		}),
	)
}

// TransferToBankCard 付款到银行卡（需要证书）
// 注意：当返回错误码为“SYSTEMERROR”时，请务必使用原商户订单号重试，否则可能造成重复支付等资金风险。
func TransferToBankCard(appid string, params *ParamsTransferBankCard) wx.Action {
	return wx.NewPostAction(urls.MchTransferToBankCard,
		wx.WithTLS(),
		wx.WithWXML(func(mchid, apikey, nonce string) (wx.WXML, error) {
			m := wx.WXML{
				"mch_id":           mchid,
				"nonce_str":        nonce,
				"partner_trade_no": params.PartnerTradeNO,
				"enc_bank_no":      params.EncBankNO,
				"enc_true_name":    params.EncTrueName,
				"bank_code":        params.BankCode,
				"amount":           strconv.Itoa(params.Amount),
			}

			if params.Desc != "" {
				m["desc"] = params.Desc
			}

			// 签名
			m["sign"] = wx.SignMD5.Do(apikey, m, true)

			return m, nil
		}),
	)
}

// QueryTransferBankCard 查询付款到银行卡结果（需要证书）
func QueryTransferBankCard(appid, partnerTradeNO string) wx.Action {
	return wx.NewPostAction(urls.MchTransferBankCardOrderQuery,
		wx.WithTLS(),
		wx.WithWXML(func(mchid, apikey, nonce string) (wx.WXML, error) {
			m := wx.WXML{
				"mch_id":           mchid,
				"partner_trade_no": partnerTradeNO,
				"nonce_str":        nonce,
			}

			// 签名
			m["sign"] = wx.SignMD5.Do(apikey, m, true)

			return m, nil
		}),
	)
}
