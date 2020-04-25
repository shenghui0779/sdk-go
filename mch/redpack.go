package mch

import (
	"errors"
	"strconv"

	"github.com/iiinsomnia/gochat/utils"
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
	AmtType   string // 红包金额设置方式，适用于裂变红包，ALL_RAND — 全部随机,商户指定总金额和红包发放总人数，由微信支付随机计算出各红包金额
	NotifyWay string // 通过JSAPI方式领取红包，小程序红包固定传 MINI_PROGRAM_JSAPI
	SceneID   string // 发放红包使用场景，红包金额大于200或者小于1元时必传
	RiskInfo  string // 活动信息，urlencode(posttime=xx&mobile=xx&deviceid=xx。posttime：用户操作的时间戳；mobile：业务系统账号的手机号，国家代码-手机号，不需要+号；deviceid：MAC地址或者设备唯一标识；clientversion：用户操作的客户端版本
}

// Redpack 企业红包
type Redpack struct {
	mch     *WXMch
	options []utils.HTTPRequestOption
}

// SendNormal 发放普通红包
func (r *Redpack) SendNormal(data *RedpackData) (utils.WXML, error) {
	body := utils.WXML{
		"wxappid":      r.mch.AppID,
		"mch_id":       r.mch.MchID,
		"nonce_str":    utils.NonceStr(),
		"mch_billno":   data.MchBillNO,
		"send_name":    data.SendName,
		"re_openid":    data.ReOpenID,
		"total_amount": strconv.Itoa(data.TotalAmount),
		"total_num":    strconv.Itoa(data.TotalNum),
		"wishing":      data.Wishing,
		"client_ip":    data.ClientIP,
		"act_name":     data.ActName,
		"remark":       data.Remark,
	}

	if data.SceneID != "" {
		body["scene_id"] = data.SceneID
	}

	if data.RiskInfo != "" {
		body["risk_info"] = data.RiskInfo
	}

	body["sign"] = SignWithMD5(body, r.mch.ApiKey)

	resp, err := r.mch.SSLClient.PostXML(RedpackNormalURL, body, r.options...)

	if err != nil {
		return nil, err
	}

	if resp["return_code"] != ResultSuccess {
		return nil, errors.New(resp["return_msg"])
	}

	if err := r.mch.VerifyWXReply(resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// SendGroup 发放裂变红包
func (r *Redpack) SendGroup(data *RedpackData) (utils.WXML, error) {
	body := utils.WXML{
		"wxappid":      r.mch.AppID,
		"mch_id":       r.mch.MchID,
		"nonce_str":    utils.NonceStr(),
		"mch_billno":   data.MchBillNO,
		"send_name":    data.SendName,
		"re_openid":    data.ReOpenID,
		"total_amount": strconv.Itoa(data.TotalAmount),
		"total_num":    strconv.Itoa(data.TotalNum),
		"amt_type":     "ALL_RAND",
		"wishing":      data.Wishing,
		"act_name":     data.ActName,
		"remark":       data.Remark,
	}

	if data.AmtType != "" {
		body["amt_type"] = data.AmtType
	}

	if data.SceneID != "" {
		body["scene_id"] = data.SceneID
	}

	if data.RiskInfo != "" {
		body["risk_info"] = data.RiskInfo
	}

	body["sign"] = SignWithMD5(body, r.mch.ApiKey)

	resp, err := r.mch.SSLClient.PostXML(RedpackGroupURL, body, r.options...)

	if err != nil {
		return nil, err
	}

	if resp["return_code"] != ResultSuccess {
		return nil, errors.New(resp["return_msg"])
	}

	if err := r.mch.VerifyWXReply(resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// SendMinip 发放小程序红包
func (r *Redpack) SendMinip(data *RedpackData) (utils.WXML, error) {
	body := utils.WXML{
		"wxappid":      r.mch.AppID,
		"mch_id":       r.mch.MchID,
		"nonce_str":    utils.NonceStr(),
		"mch_billno":   data.MchBillNO,
		"send_name":    data.SendName,
		"re_openid":    data.ReOpenID,
		"total_amount": strconv.Itoa(data.TotalAmount),
		"total_num":    strconv.Itoa(data.TotalNum),
		"wishing":      data.Wishing,
		"act_name":     data.ActName,
		"remark":       data.Remark,
		"notify_way":   "MINI_PROGRAM_JSAPI",
	}

	if data.NotifyWay != "" {
		body["notify_way"] = data.NotifyWay
	}

	if data.SceneID != "" {
		body["scene_id"] = data.SceneID
	}

	body["sign"] = SignWithMD5(body, r.mch.ApiKey)

	resp, err := r.mch.SSLClient.PostXML(RedpackMinipURL, body, r.options...)

	if err != nil {
		return nil, err
	}

	if resp["return_code"] != ResultSuccess {
		return nil, errors.New(resp["return_msg"])
	}

	if err := r.mch.VerifyWXReply(resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// QueryByBillNO 查询红包记录
func (r *Redpack) QueryByBillNO(billNO string) (utils.WXML, error) {
	body := utils.WXML{
		"appid":      r.mch.AppID,
		"mch_id":     r.mch.MchID,
		"mch_billno": billNO,
		"bill_type":  "MCHT",
		"nonce_str":  utils.NonceStr(),
	}

	body["sign"] = SignWithMD5(body, r.mch.ApiKey)

	resp, err := r.mch.SSLClient.PostXML(RedpackQueryURL, body, r.options...)

	if err != nil {
		return nil, err
	}

	if resp["return_code"] != ResultSuccess {
		return nil, errors.New(resp["return_msg"])
	}

	if err := r.mch.VerifyWXReply(resp); err != nil {
		return nil, err
	}

	return resp, nil
}
