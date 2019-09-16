package mch

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/iiinsomnia/gochat/consts"
	"github.com/iiinsomnia/gochat/utils"
)

// Entrust 微信纯签约
type Entrust struct {
	// 必填字段
	PlanID                 string // 协议模板id，设置路径见开发步骤
	ContractCode           string // 商户侧的签约协议号，由商户生成
	RequestSerial          int64  // 商户请求签约时的序列号，要求唯一性，纯数字, 范围不能超过Int64的范围
	ContractDisplayAccount string // 签约用户的名称，用于页面展示，参数值不支持UTF8非3字节编码的字符，如表情符号，故请勿使用微信昵称
	ClientIP               string // 用户客户端的真实IP地址，H5签约必填
	NotifyURL              string // 用于接收签约成功消息的回调通知地址，对notify_url参数值需进行encode处理,注意是对参数值进行encode
	// 选填字段
	ReturnAPP   bool   // APP签约选填，签约后是否返回app，注：签约参数appid必须为发起签约的app所有，且在微信开放平台注册过
	ReturnWeb   bool   // 公众号签约选填，签约后是否返回签约页面的referrer url, 不填或获取不到referrer则不返回; 跳转referrer url时会自动带上参数from_wxpay=1
	OuterID     int64  // 小程序签约选填，用户在商户侧的标识
	ReturnAPPID string // H5签约选填，商户具有指定返回app的权限时，签约成功将返回appid指定的app应用，如不填且签约发起时的浏览器UA可被微信识别，则跳转到浏览器，否则留在微信
}

// PappayApply 扣款申请
type PappayApply struct {
	// 必填参数
	OutTradeNO     string // 商户系统内部的订单号，32个字符内、可包含字母，其他说明见商户订单号
	TotalFee       int    // 订单总金额，单位为分，详见支付金额
	SpbillCreateIP string // APP和网页支付提交用户端ip，Native支付填调用微信支付API的机器IP
	ContractID     string // 签约成功后，微信返回的委托代扣协议id
	Body           string // 商品或支付单简要描述
	NotifyURL      string // 接收微信支付异步通知回调地址，通知url必须为直接可访问的url，不能携带参数
	// 选填参数
	Detail   string // 商品名称明细列表
	Attach   string // 附加数据，在查询API和支付通知中原样返回，该字段主要用于商户携带订单的自定义数据
	FeeType  string // 符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	GoodsTag string // 商品标记，代金券或立减优惠功能的参数，说明详见代金券或立减优惠
	Receipt  bool   // 是否在支付成功消息和支付详情页中出现开票入口，注：需要在微信支付商户平台或微信公众平台开通电子发票功能
}

type Pappay struct {
	appid  string
	mchid  string
	apikey string
	client *utils.HTTPClient
}

// APPEntrust APP纯签约
func (p *Pappay) APPEntrust(e *Entrust) (utils.WXML, error) {
	body := utils.WXML{
		"appid":                    p.appid,
		"mch_id":                   p.mchid,
		"plan_id":                  e.PlanID,
		"contract_code":            e.ContractCode,
		"request_serial":           strconv.FormatInt(e.RequestSerial, 10),
		"contract_display_account": e.ContractDisplayAccount,
		"version":                  "1.0",
		"sign_type":                consts.MchSignMD5,
		"timestamp":                strconv.FormatInt(time.Now().Unix(), 10),
		"notify_url":               e.NotifyURL,
	}

	if e.ReturnAPP {
		body["return_app"] = "Y"
	}

	return p.doPost(consts.MchPappayAPPEntrustURL, body, consts.MchSignMD5)
}

// PubEntrust 公众号纯签约
func (p *Pappay) PubEntrust(e *Entrust) (utils.WXML, error) {
	body := utils.WXML{
		"appid":                    p.appid,
		"mch_id":                   p.mchid,
		"plan_id":                  e.PlanID,
		"contract_code":            e.ContractCode,
		"request_serial":           strconv.FormatInt(e.RequestSerial, 10),
		"contract_display_account": e.ContractDisplayAccount,
		"version":                  "1.0",
		"timestamp":                strconv.FormatInt(time.Now().Unix(), 10),
		"notify_url":               e.NotifyURL,
	}

	if e.ReturnWeb {
		body["return_web"] = "1"
	}

	return p.doGet(consts.MchPappayPubEntrustURL, body, consts.MchSignMD5)
}

// MPEntrust 小程序纯签约，返回小程序所需的 extraData 数据
func (p *Pappay) MPEntrust(e *Entrust) utils.WXML {
	extraData := utils.WXML{
		"appid":                    p.appid,
		"mch_id":                   p.mchid,
		"plan_id":                  e.PlanID,
		"contract_code":            e.ContractCode,
		"request_serial":           strconv.FormatInt(e.RequestSerial, 10),
		"contract_display_account": e.ContractDisplayAccount,
		"timestamp":                strconv.FormatInt(time.Now().Unix(), 10),
		"notify_url":               e.NotifyURL,
	}

	if e.OuterID != 0 {
		extraData["outerid"] = strconv.FormatInt(e.OuterID, 10)
	}

	extraData["sign"] = SignWithMD5(extraData, p.apikey)

	return extraData
}

// H5Entrust H5纯签约
func (p *Pappay) H5Entrust(e *Entrust) (utils.WXML, error) {
	body := utils.WXML{
		"appid":                    p.appid,
		"mch_id":                   p.mchid,
		"plan_id":                  e.PlanID,
		"contract_code":            e.ContractCode,
		"request_serial":           strconv.FormatInt(e.RequestSerial, 10),
		"contract_display_account": e.ContractDisplayAccount,
		"version":                  "1.0",
		"timestamp":                strconv.FormatInt(time.Now().Unix(), 10),
		"clientip":                 e.ClientIP,
		"notify_url":               e.NotifyURL,
	}

	if e.ReturnAPPID != "" {
		body["return_appid"] = e.ReturnAPPID
	}

	return p.doGet(consts.MchPappayH5EntrustURL, body, consts.MchSignHMacSHA256)
}

// QueryContractByID 根据微信返回的委托代扣协议id查询签约关系
func (p *Pappay) QueryContractByID(contractID string) (utils.WXML, error) {
	body := utils.WXML{
		"appid":       p.appid,
		"mch_id":      p.mchid,
		"contract_id": contractID,
		"version":     "1.0",
	}

	return p.doPost(consts.MchPappayContractQueryURL, body, consts.MchSignMD5)
}

// QueryContractByCode 根据签约协议号查询签约关系，需要商户平台配置的代扣模版id
func (p *Pappay) QueryContractByCode(planID, contractCode string) (utils.WXML, error) {
	body := utils.WXML{
		"appid":         p.appid,
		"mch_id":        p.mchid,
		"plan_id":       planID,
		"contract_code": contractCode,
		"version":       "1.0",
	}

	return p.doPost(consts.MchPappayContractQueryURL, body, consts.MchSignMD5)
}

// PayApply 申请扣款
func (p *Pappay) PayApply(apply *PappayApply) (utils.WXML, error) {
	body := utils.WXML{
		"appid":            p.appid,
		"mch_id":           p.mchid,
		"nonce_str":        utils.NonceStr(),
		"fee_type":         "CNY",
		"trade_type":       "PAP",
		"notify_url":       apply.NotifyURL,
		"body":             apply.Body,
		"out_trade_no":     apply.OutTradeNO,
		"total_fee":        strconv.Itoa(apply.TotalFee),
		"contract_id":      apply.ContractID,
		"spbill_create_ip": apply.SpbillCreateIP,
	}

	if apply.Detail != "" {
		body["detail"] = apply.Detail
	}

	if apply.Attach != "" {
		body["attach"] = apply.Attach
	}

	if apply.FeeType != "" {
		body["fee_type"] = apply.FeeType
	}

	if apply.GoodsTag != "" {
		body["goods_tag"] = apply.GoodsTag
	}

	if apply.Receipt {
		body["receipt"] = "Y"
	}

	return p.doPost(consts.MchPappayPayApplyURL, body, consts.MchSignMD5)
}

// DeleteContractByID 根据微信返回的委托代扣协议id解约
func (p *Pappay) DeleteContractByID(contractID, remark string) (utils.WXML, error) {
	body := utils.WXML{
		"appid":                       p.appid,
		"mch_id":                      p.mchid,
		"contract_id":                 contractID,
		"version":                     "1.0",
		"contract_termination_remark": remark,
	}

	return p.doPost(consts.MchPappayContractDeleteURL, body, consts.MchSignMD5)
}

// DeleteContractByCode 根据签约协议号解约，需要商户平台配置的代扣模版id
func (p *Pappay) DeleteContractByCode(planID, contractCode, remark string) (utils.WXML, error) {
	body := utils.WXML{
		"appid":                       p.appid,
		"mch_id":                      p.mchid,
		"plan_id":                     planID,
		"contract_code":               contractCode,
		"version":                     "1.0",
		"contract_termination_remark": remark,
	}

	return p.doPost(consts.MchPappayContractDeleteURL, body, consts.MchSignMD5)
}

// QueryOrderByTransactionID 根据微信订单号查询
func (p *Pappay) QueryOrderByTransactionID(transactionID string) (utils.WXML, error) {
	body := utils.WXML{
		"appid":          p.appid,
		"mch_id":         p.mchid,
		"transaction_id": transactionID,
		"nonce_str":      utils.NonceStr(),
	}

	return p.doPost(consts.MchPappayOrderQueryURL, body, consts.MchSignMD5)
}

// QueryOrderByOutTradeNO 根据商户订单号查询
func (p *Pappay) QueryOrderByOutTradeNO(outTradeNO string) (utils.WXML, error) {
	body := utils.WXML{
		"appid":        p.appid,
		"mch_id":       p.mchid,
		"out_trade_no": outTradeNO,
		"nonce_str":    utils.NonceStr(),
		"sign_type":    "MD5",
	}

	return p.doPost(consts.MchPappayOrderQueryURL, body, consts.MchSignMD5)
}

func (p *Pappay) doGet(uri string, body utils.WXML, signType string) (utils.WXML, error) {
	switch signType {
	case consts.MchSignMD5:
		body["sign"] = SignWithMD5(body, p.apikey)
	case consts.MchSignHMacSHA256:
		body["sign"] = SignWithHMacSHA256(body, p.apikey)
	}

	query := url.Values{}

	for k, v := range body {
		if k == "sign" || v == "" {
			continue
		}

		query.Add(k, v)
	}

	resp, err := p.client.GetXML(fmt.Sprintf("%s?%s", uri, query.Encode()))

	if err != nil {
		return nil, err
	}

	if resp["return_code"] != consts.MchReplySuccess {
		return nil, errors.New(resp["return_msg"])
	}

	if resp["result_code"] != consts.MchReplySuccess {
		return nil, errors.New(resp["err_code_des"])
	}

	signature := ""

	switch signType {
	case consts.MchSignMD5:
		signature = SignWithMD5(resp, p.apikey)
	case consts.MchSignHMacSHA256:
		signature = SignWithHMacSHA256(resp, p.apikey)
	}

	if signature != resp["sign"] {
		return nil, fmt.Errorf("papay resp signature verified failed, want: %s, got: %s", signature, resp["sign"])
	}

	if resp["appid"] != p.appid {
		return nil, fmt.Errorf("papay resp appid mismatch, want: %s, got: %s", p.appid, resp["appid"])
	}

	if resp["mch_id"] != p.mchid {
		return nil, fmt.Errorf("papay resp mchid mismatch, want: %s, got: %s", p.mchid, resp["mch_id"])
	}

	return resp, nil
}

func (p *Pappay) doPost(url string, body utils.WXML, signType string) (utils.WXML, error) {
	switch signType {
	case consts.MchSignMD5:
		body["sign"] = SignWithMD5(body, p.apikey)
	case consts.MchSignHMacSHA256:
		body["sign"] = SignWithHMacSHA256(body, p.apikey)
	}

	resp, err := p.client.PostXML(url, body)

	if err != nil {
		return nil, err
	}

	if resp["return_code"] != consts.MchReplySuccess {
		return nil, errors.New(resp["return_msg"])
	}

	if resp["result_code"] != consts.MchReplySuccess {
		return nil, errors.New(resp["err_code_des"])
	}

	signature := ""

	switch signType {
	case consts.MchSignMD5:
		signature = SignWithMD5(resp, p.apikey)
	case consts.MchSignHMacSHA256:
		signature = SignWithHMacSHA256(resp, p.apikey)
	}

	if signature != resp["sign"] {
		return nil, fmt.Errorf("papay resp signature verified failed, want: %s, got: %s", signature, resp["sign"])
	}

	if resp["appid"] != p.appid {
		return nil, fmt.Errorf("papay resp appid mismatch, want: %s, got: %s", p.appid, resp["appid"])
	}

	if resp["mch_id"] != p.mchid {
		return nil, fmt.Errorf("papay resp mchid mismatch, want: %s, got: %s", p.mchid, resp["mch_id"])
	}

	return resp, nil
}
