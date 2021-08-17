/*
@Time : 2021/8/16 4:56 下午
@Author : 21
@File : oplatform
@Software: GoLand
*/
package oplatform

import (
	"context"
	"errors"
	"fmt"
	"github.com/shenghui0779/gochat/event"
	"github.com/shenghui0779/gochat/wx"
	"github.com/shenghui0779/yiigo"
	"github.com/tidwall/gjson"
)

type Oplatform  struct {
	appid          string
	appsecret      string
	token          string
	encodingAESKey string
	componentVerifyTicket string //
	nonce          func(size uint) string
	client         wx.Client
}


// New returns new wechat mini program
func New(appid, appsecret string) *Oplatform {
	return &Oplatform{
		appid:     appid,
		appsecret: appsecret,
		nonce:     wx.Nonce,
		client:    wx.NewClient(wx.WithInsecureSkipVerify()),
	}
}


// SetServerConfig 设置服务器配置
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/component_verify_ticket.html
func (o *Oplatform) SetServerConfig(token, encodingAESKey ,componentVerifyTicket  string) {
	o.token = token
	o.encodingAESKey = encodingAESKey
	o.componentVerifyTicket = componentVerifyTicket
}

// AppID returns appid
func (o *Oplatform) AppID() string {
	return o.appid
}

// AppSecret returns app secret
func (o *Oplatform) AppSecret() string {
	return o.appsecret
}

// ComponentVerifyTicket returns app componentVerifyTicket
func (o *Oplatform)  ComponentVerifyTicket () string {
	return o.componentVerifyTicket
}

// DecryptEventMessage 事件消息解密
func (o *Oplatform) DecryptEventMessage(appId string,encrypt string) (wx.WXML, error) {
	b, err := event.Decrypt(appId, o.encodingAESKey, encrypt)

	if err != nil {
		return nil, err
	}

	return wx.ParseXML2Map(b)
}

// 获取 移动端授权链接的方法
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Before_Develop/Authorization_Process_Technical_Description.html
func (o *Oplatform) SafeBindComponent(preAuthCode string, redirectUri string, authType int, bizAppid string) (string, error) {
	if len(o.componentVerifyTicket) < 1 {
		return "", errors.New("component_verify_ticket is error")
	}

	safeBindComponentUrl := fmt.Sprintf("%s/safe/bindcomponent?action=bindcomponent&no_scan=1&component_appid=%s&pre_auth_code=%s&redirect_uri=%s&auth_type=%d&biz_appid=%s#wechat_redirect",
		BaseUrl, o.appid, preAuthCode, redirectUri, authType, bizAppid)
	return safeBindComponentUrl, nil
}

// Do exec action
func (o *Oplatform) Do(ctx context.Context,  action wx.Action, options ...yiigo.HTTPOption) error {
	var (
		resp []byte
		err  error
	)

	switch action.Method() {
	case wx.MethodGet:
		resp, err = o.client.Get(ctx, action.URL(), options...)
	case wx.MethodPost:
		body, berr := action.Body()

		if berr != nil {
			return berr
		}

		resp, err = o.client.Post(ctx, action.URL(), body, options...)
	case wx.MethodUpload:
		form, ferr := action.UploadForm()

		if ferr != nil {
			return ferr
		}

		resp, err = o.client.Upload(ctx, action.URL(), form, options...)
	}

	if err != nil {
		return err
	}

	r := gjson.ParseBytes(resp)

	if code := r.Get("errcode").Int(); code != 0 {
		return fmt.Errorf("%d|%s", code, r.Get("errmsg").String())
	}

	if action.Decode() == nil {
		return nil
	}

	return action.Decode()(resp)
}


