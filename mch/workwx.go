package mch

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
	"github.com/shenghui0779/yiigo"
)

// WorkWXRedpackData 企业红包发放数据
type WorkWXRedpackData struct {
	// 必填参数
	MchBillNO   string // 商户订单号（每个订单号必须唯一。取值范围：0~9，a~z，A~Z）接口根据商户订单号支持重入，如出现超时可再调用
	ReOpenID    string // 接受红包的用户openid
	TotalAmount int    // 付款金额，单位：分，单笔最小金额默认为1元
	Wishing     string // 红包祝福语
	ActName     string // 活动名称
	Remark      string // 备注信息
	// 选填参数
	AgentID             string // 以企业应用的名义发红包，企业应用id，整型，可在企业微信管理端应用的设置页面查看。与sender_name互斥，二者只能填一个
	SenderName          string // 以个人名义发红包，红包发送者名称(需要utf-8格式)。与agentid互斥，二者只能填一个
	SenderHeaderMediaID string // 发送者头像素材id，通过企业微信开放上传素材接口获取
	SceneID             string // 发放红包使用场景，红包金额大于200或者小于1元时必传
}

// SendWorkWXRedpack 发放企业红包
func SendWorkWXRedpack(secret string, data *WorkWXRedpackData) wx.Action {
	return wx.NewPostAction(urls.MchRedpackWorkWX,
		wx.WithTLS(),
		wx.WithWXML(func(appid, mchid, nonce string) (wx.WXML, error) {
			body := wx.WXML{
				"wxappid":      appid,
				"mch_id":       mchid,
				"nonce_str":    nonce,
				"mch_billno":   data.MchBillNO,
				"re_openid":    data.ReOpenID,
				"total_amount": strconv.Itoa(data.TotalAmount),
				"wishing":      data.Wishing,
				"act_name":     data.ActName,
				"remark":       data.Remark,
				"sign_type":    string(SignMD5),
			}

			workwxSignStr := fmt.Sprintf("act_name=%s&mch_billno=%s&mch_id=%s&nonce_str=%s&re_openid=%s&total_amount=%d&wxappid=%s&secret=%s", data.ActName, data.MchBillNO, mchid, nonce, data.ReOpenID, data.TotalAmount, appid, secret)

			body["workwx_sign"] = strings.ToUpper(yiigo.MD5(workwxSignStr))

			if len(data.AgentID) != 0 {
				body["agentid"] = data.AgentID
			}

			if len(data.SenderName) != 0 {
				body["sender_name"] = data.SenderName
			}

			if len(data.SenderHeaderMediaID) != 0 {
				body["sender_header_media_id"] = data.SenderHeaderMediaID
			}

			if len(data.SceneID) != 0 {
				body["scene_id"] = data.SceneID
			}

			return body, nil
		}),
	)
}
