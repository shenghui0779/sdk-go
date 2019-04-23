package wxpay

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/iiinsomnia/yiigo"
	"go.uber.org/zap"
	"meipian.cn/printapi/wechat"
	"meipian.cn/printapi/wechat/utils"
)

// UnifiedOrder 统一下单
type UnifiedOrder struct {
	OutTradeNO string
	TotalFee   int
	ClientIP   string
	OpenID     string
	NotifyURL  string
	Channel    string
	Body       string
	wxResp     utils.WXML
}

// Do 统一下单
func (u *UnifiedOrder) Do() error {
	settings := wechat.GetSettingsWithChannel(u.Channel)

	body := utils.WXML{
		"appid":            settings.AppID,
		"mch_id":           settings.MchID,
		"nonce_str":        utils.NonceStr(),
		"sign_type":        "MD5",
		"fee_type":         "CNY",
		"trade_type":       settings.TradeType,
		"notify_url":       u.NotifyURL,
		"body":             u.Body,
		"out_trade_no":     u.OutTradeNO,
		"total_fee":        strconv.Itoa(u.TotalFee),
		"spbill_create_ip": u.ClientIP,
		"openid":           u.OpenID,
	}

	body["sign"] = utils.PaySign(body, settings.ApiKey)

	resp, err := utils.PostXML(defaultClient, "https://api.mch.weixin.qq.com/pay/unifiedorder", body)

	if err != nil {
		return err
	}

	if resp["return_code"] != "SUCCESS" {
		yiigo.Logger.Error("do wx unifiedorder error", zap.String("trade_no", u.OutTradeNO), zap.Any("resp", resp["return_msg"]))

		return errors.New(resp["return_msg"])
	}

	if resp["result_code"] != "SUCCESS" {
		yiigo.Logger.Error("do wx unifiedorder error", zap.String("trade_no", u.OutTradeNO), zap.String("code", resp["err_code"]), zap.String("error", resp["err_code_des"]))

		return errors.New(resp["err_code_des"])
	}

	u.wxResp = resp

	return nil
}

// Result 统一下单结果
func (u *UnifiedOrder) Result() utils.WXML {
	return u.wxResp
}

// PrepayID 预付款ID
func (u *UnifiedOrder) PrepayID() string {
	prepayID, ok := u.wxResp["prepay_id"]

	if !ok {
		return ""
	}

	return prepayID
}

// BuildWXCharge 生成微信支付对象
func BuildWXCharge(channel, prepayID string) utils.WXML {
	settings := wechat.GetSettingsWithChannel(channel)

	if channel == wechat.WXAPP {
		charge := utils.WXML{
			"appid":     settings.AppID,
			"partnerid": settings.MchID,
			"prepayid":  prepayID,
			"package":   "Sign=WXPay",
			"noncestr":  utils.NonceStr(),
			"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
		}

		charge["sign"] = utils.PaySign(charge, settings.ApiKey)

		return charge
	}

	charge := utils.WXML{
		"appId":     settings.AppID,
		"nonceStr":  utils.NonceStr(),
		"package":   fmt.Sprintf("prepay_id=%s", prepayID),
		"signType":  "MD5",
		"timeStamp": strconv.FormatInt(time.Now().Unix(), 10),
	}

	charge["paySign"] = utils.PaySign(charge, settings.ApiKey)

	return charge
}
