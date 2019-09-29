package mch

import (
	"fmt"
	"strconv"
	"time"

	"github.com/iiinsomnia/gochat/utils"
)

// WXMch 微信商户
type WXMch struct {
	AppID     string
	MchID     string
	AppKey    string
	Client    *utils.HTTPClient
	SSLClient *utils.HTTPClient
}

// Order returns new order
func (wx *WXMch) Order() *Order {
	return &Order{wx}
}

// Refund returns new refund
func (wx *WXMch) Refund() *Refund {
	return &Refund{wx}
}

// Pappay returns new pappay
func (wx *WXMch) Pappay() *Pappay {
	return &Pappay{wx}
}

// APPAPI 用于APP拉起支付
func (wx *WXMch) APPAPI(prepayID string) utils.WXML {
	ch := utils.WXML{
		"appid":     wx.AppID,
		"partnerid": wx.MchID,
		"prepayid":  prepayID,
		"package":   "Sign=WXPay",
		"noncestr":  utils.NonceStr(),
		"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
	}

	ch["sign"] = SignWithMD5(ch, wx.AppKey)

	return ch
}

// JSAPI 用于JS拉起支付
func (wx *WXMch) JSAPI(prepayID string) utils.WXML {
	ch := utils.WXML{
		"appId":     wx.AppID,
		"nonceStr":  utils.NonceStr(),
		"package":   fmt.Sprintf("prepay_id=%s", prepayID),
		"signType":  SignMD5,
		"timeStamp": strconv.FormatInt(time.Now().Unix(), 10),
	}

	ch["paySign"] = SignWithMD5(ch, wx.AppKey)

	return ch
}

// VerifyWXReply 验证微信结果
func (wx *WXMch) VerifyWXReply(reply utils.WXML, signType string) error {
	signature := ""

	switch signType {
	case SignMD5:
		signature = SignWithMD5(reply, wx.AppKey)
	case SignHMacSHA256:
		signature = SignWithHMacSHA256(reply, wx.AppKey)
	}

	if signature != reply["sign"] {
		return fmt.Errorf("signature verified failed, want: %s, got: %s", signature, reply["sign"])
	}

	if appid, ok := reply["appid"]; ok {
		if appid != wx.AppID {
			return fmt.Errorf("appid mismatch, want: %s, got: %s", wx.AppID, reply["appid"])
		}
	}

	if mchid, ok := reply["mch_id"]; ok {
		if mchid != wx.MchID {
			return fmt.Errorf("mchid mismatch, want: %s, got: %s", wx.MchID, reply["mch_id"])
		}
	}

	return nil
}
