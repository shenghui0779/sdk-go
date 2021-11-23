package minip

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/shenghui0779/yiigo"
	"github.com/tidwall/gjson"

	"github.com/shenghui0779/gochat/event"
	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

// Minip 微信小程序
type Minip struct {
	appid          string
	appsecret      string
	token          string
	encodingAESKey string
	nonce          func() string
	client         wx.Client
}

// New returns new wechat mini program
func New(appid, appsecret string) *Minip {
	return &Minip{
		appid:     appid,
		appsecret: appsecret,
		nonce: func() string {
			return wx.Nonce(16)
		},
		client: wx.NewDefaultClient(),
	}
}

// SetServerConfig 设置服务器配置
// [参考](https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Access_Overview.html)
func (mp *Minip) SetServerConfig(token, encodingAESKey string) {
	mp.token = token
	mp.encodingAESKey = encodingAESKey
}

// SetClient set client
func (mp *Minip) SetClient(client wx.Client) {
	mp.client = client
}

// SetLogger set client logger
func (mp *Minip) SetLogger(logger wx.Logger) {
	mp.client.SetLogger(logger)
}

// AppID returns appid
func (mp *Minip) AppID() string {
	return mp.appid
}

// AppSecret returns app secret
func (mp *Minip) AppSecret() string {
	return mp.appsecret
}

// Code2Session 获取小程序授权的session_key
func (mp *Minip) Code2Session(ctx context.Context, code string, options ...yiigo.HTTPOption) (*AuthSession, error) {
	resp, err := mp.client.Get(ctx, fmt.Sprintf("%s?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", urls.MinipCode2Session, mp.appid, mp.appsecret, code), options...)

	if err != nil {
		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if code := r.Get("errcode").Int(); code != 0 {
		return nil, fmt.Errorf("%d|%s", code, r.Get("errmsg").String())
	}

	session := new(AuthSession)

	if err = json.Unmarshal(resp, session); err != nil {
		return nil, err
	}

	return session, nil
}

// AccessToken 获取小程序的access_token
func (mp *Minip) AccessToken(ctx context.Context, options ...yiigo.HTTPOption) (*AccessToken, error) {
	resp, err := mp.client.Get(ctx, fmt.Sprintf("%s?appid=%s&secret=%s&grant_type=client_credential", urls.MinipAccessToken, mp.appid, mp.appsecret), options...)

	if err != nil {
		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if code := r.Get("errcode").Int(); code != 0 {
		return nil, fmt.Errorf("%d|%s", code, r.Get("errmsg").String())
	}

	token := new(AccessToken)

	if err = json.Unmarshal(resp, token); err != nil {
		return nil, err
	}

	return token, nil
}

// DecryptAuthInfo 解密授权信息
func (mp *Minip) DecryptAuthInfo(sessionKey, iv, encryptedData string, result AuthInfo) error {
	key, err := base64.StdEncoding.DecodeString(sessionKey)

	if err != nil {
		return err
	}

	ivb, err := base64.StdEncoding.DecodeString(iv)

	if err != nil {
		return err
	}

	cipherText, err := base64.StdEncoding.DecodeString(encryptedData)

	if err != nil {
		return err
	}

	cbc := yiigo.NewCBCCrypto(key, ivb, yiigo.PKCS7)

	b, err := cbc.Decrypt(cipherText)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, result); err != nil {
		return err
	}

	if result.AppID() != mp.appid {
		return fmt.Errorf("appid mismatch, want: %s, got: %s", mp.appid, result.AppID())
	}

	return nil
}

// Do exec action
func (mp *Minip) Do(ctx context.Context, accessToken string, action wx.Action, options ...yiigo.HTTPOption) error {
	var (
		resp []byte
		err  error
	)

	switch action.Method() {
	case wx.MethodGet:
		resp, err = mp.client.Get(ctx, action.URL(accessToken), options...)
	case wx.MethodPost:
		body, berr := action.Body()

		if berr != nil {
			return err
		}

		resp, err = mp.client.Post(ctx, action.URL(accessToken), body, options...)
	case wx.MethodUpload:
		form, ferr := action.UploadForm()

		if ferr != nil {
			fmt.Println("[ERR]", ferr)
			return ferr
		}

		resp, err = mp.client.Upload(ctx, action.URL(accessToken), form, options...)
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

// VerifyEventSign 验证事件消息签名
// 验证消息来自微信服务器，使用：signature、timestamp、nonce；若验证成功，请原样返回echostr参数内容
// 验证事件消息签名，使用：msg_signature、timestamp、nonce、msg_encrypt
// [参考](https://developers.weixin.qq.com/miniprogram/dev/framework/server-ability/message-push.html)
func (mp *Minip) VerifyEventSign(signature string, items ...string) bool {
	signStr := event.SignWithSHA1(mp.token, items...)

	return signStr == signature
}

// DecryptEventMessage 事件消息解密
func (mp *Minip) DecryptEventMessage(encrypt string) (wx.WXML, error) {
	b, err := event.Decrypt(mp.appid, mp.encodingAESKey, encrypt)

	if err != nil {
		return nil, err
	}

	return wx.ParseXML2Map(b)
}
