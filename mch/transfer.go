package mch

import (
	"errors"
	"strconv"

	"github.com/iiinsomnia/gochat/utils"
)

type TransferBalanceData struct {
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

type TransferBankCardData struct {
}

type Transfer struct {
	mch     *WXMch
	options []utils.HTTPRequestOption
}

func (t *Transfer) ToBalance(data *TransferBalanceData) (utils.WXML, error) {
	body := utils.WXML{
		"mch_appid":        t.mch.AppID,
		"mchid":            t.mch.MchID,
		"nonce_str":        utils.NonceStr(),
		"partner_trade_no": data.PartnerTradeNO,
		"openid":           data.OpenID,
		"check_name":       data.CheckName,
		"amount":           strconv.Itoa(data.Amount),
		"desc":             data.Desc,
	}

	if data.ReUserName != "" {
		body["re_user_name"] = data.ReUserName
	}

	if data.DeviceInfo != "" {
		body["device_info"] = data.DeviceInfo
	}

	if data.SpbillCreateIP != "" {
		body["spbill_create_ip"] = data.SpbillCreateIP
	}

	body["sign"] = SignWithMD5(body, t.mch.ApiKey)

	resp, err := t.mch.SSLClient.PostXML(TransferToBalanceURL, body, t.options...)

	if err != nil {
		return nil, err
	}

	if resp["return_code"] != ResultSuccess {
		return nil, errors.New(resp["return_msg"])
	}

	if err := t.mch.VerifyWXReply(resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *Transfer) QueryBalanceOrder(partnerTradeNO string) (utils.WXML, error) {
	body := utils.WXML{
		"appid":            t.mch.AppID,
		"mch_id":           t.mch.MchID,
		"partner_trade_no": partnerTradeNO,
		"nonce_str":        utils.NonceStr(),
	}

	body["sign"] = SignWithMD5(body, t.mch.ApiKey)

	resp, err := t.mch.Client.PostXML(TransferBalanceOrderQueryURL, body, t.options...)

	if err != nil {
		return nil, err
	}

	if resp["return_code"] != ResultSuccess {
		return nil, errors.New(resp["return_msg"])
	}

	if err := t.mch.VerifyWXReply(resp); err != nil {
		return nil, err
	}

	return resp, nil
}
