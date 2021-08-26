/*
@Time : 2021/8/16 5:42 下午
@Author : 21
@File : oplatform_test
@Software: GoLand
*/
package oplatform

import (
	"fmt"
	"github.com/shenghui0779/gochat/wx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOplatform_DecryptEventMessage(t *testing.T) {
	var EventMsg = "lr/RV+4U21T5Ie+cP+NvRW7lJNLgXD+fVy7iZOSCnKagLdjKHoT5s0bv7EmSrOeTsdi13QA7yv1Sf22d/EWUCwF1D12ImCB9+qeaWsLtUUYEZoQLtALhjBZwp4VrSpBCxiw61MKYpuPV7i/qemZ5LjWcA0FwBp4Mce5Gd8zRikEQyCL+ThHf8Zi2wlG2fue+/ly8Xc/h+0MBCfyekp2JKFnpoAtgblbCOiEZDhKc6a8CP0OerEPann1S0Z5ZMQg5rgJeF+lX2Gyxx1QPIRX8Kxup3MqsC09VHa1vtS585QS9NxMSqOF4Ss3GrdcaWp+CLBa/J+Q35hrmChaDsivbeRIIBcx24ncymMOBIL/buyBG9IC8ezRhPFl7MH4q2Se5"
	op := New("wxc83d3daa98fe100c","dd8c33e9d4634923f70a77ada891f4f8")
	op.SetServerConfig("womeibanfale","zhinengxiugainimenle00000000000000000000001","123123")
	msg , err := op.DecryptEventMessage("wx77182eed6aa0cf1b",EventMsg)
	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
		//"ToUserName":   "gh_3ad31c0ba9b5",
		//"FromUserName": "oB4tA6ANthOfuQ5XSlkdPsWOVUsY",
		//"CreateTime":   "1606902602",
		//"MsgType":      "text",
		//"MsgId":        "10086",
		//"Content":      "ILoveGochat",
		//"URL":          "http://182.92.100.180/webhook",
	}, msg)

}

func TestOplatform_Reply(t *testing.T) {
	op := New("wxc83d3daa98fe100c","dd8c33e9d4634923f70a77ada891f4f8")
	op.SetServerConfig("womeibanfale","zhinengxiugainimenle00000000000000000000001","123123")
	res , err := op.Reply("1111111111","1111111",NewTextReply("121212"))
	assert.Nil(t,err)
	fmt.Println(res)


}
