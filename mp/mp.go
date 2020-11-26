package mp

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/shenghui0779/gochat/event"
	"github.com/shenghui0779/gochat/helpers"
	"github.com/tidwall/gjson"
)

// Action is a interface that defines wechat mp action
type Action interface {
	Body() helpers.HTTPBody
	URL() func(accessToken string) string
	Decode() func(resp []byte) error
}

// WechatAPI is a Action implementation for wechat mp api
type WechatAPI struct {
	body   helpers.HTTPBody
	url    func(accessToken string) string
	decode func(resp []byte) error
}

// Body returns http body
func (a *WechatAPI) Body() helpers.HTTPBody {
	return a.body
}

// URL returns url closure
func (a *WechatAPI) URL() func(accessToken string) string {
	return a.url
}

// Decode returns decode closure
func (a *WechatAPI) Decode() func(resp []byte) error {
	return a.decode
}

// WechatMP 微信小程序
type WechatMP struct {
	appid          string
	appsecret      string
	signToken      string
	encodingAESKey string
	nonce          func(size int) string
	client         helpers.HTTPClient
}

// New returns new wechat mini program
func New(appid, appsecret string) *WechatMP {
	return &WechatMP{
		appid:     appid,
		appsecret: appsecret,
		nonce: func(size int) string {
			nonce := make([]byte, size/2)
			io.ReadFull(rand.Reader, nonce)

			return hex.EncodeToString(nonce)
		},
		client: helpers.NewHTTPClient(),
	}
}

// SetServerConfig 设置服务器配置
func (w *WechatMP) SetServerConfig(token, encodingAESKey string) {
	w.signToken = token
	w.encodingAESKey = encodingAESKey
}

// Code2Session 获取小程序授权的session_key
func (w *WechatMP) Code2Session(ctx context.Context, code string, options ...helpers.HTTPOption) (*AuthSession, error) {
	resp, err := w.client.Get(ctx, fmt.Sprintf("%s?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", Code2SessionURL, w.appid, w.appsecret, code), options...)

	if err != nil {
		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return nil, errors.New(r.Get("errmsg").String())
	}

	session := new(AuthSession)

	if err = json.Unmarshal(resp, session); err != nil {
		return nil, err
	}

	return session, nil
}

// AccessToken 获取小程序的access_token
func (w *WechatMP) AccessToken(ctx context.Context, options ...helpers.HTTPOption) (*AccessToken, error) {
	resp, err := w.client.Get(ctx, fmt.Sprintf("%s?grant_type=client_credential&appid=%s&secret=%s", AccessTokenURL, w.appid, w.appsecret), options...)

	if err != nil {
		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return nil, errors.New(r.Get("errmsg").String())
	}

	token := new(AccessToken)

	if err = json.Unmarshal(resp, token); err != nil {
		return nil, err
	}

	return token, nil
}

// DecryptAuthInfo 解密授权信息
func (w *WechatMP) DecryptAuthInfo(sessionKey, iv, encryptedData string, receiver AuthInfo) error {
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

	cbc := helpers.NewAESCBCCrypto(key, ivb)

	b, err := cbc.Decrypt(cipherText, helpers.PKCS7)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, receiver); err != nil {
		return err
	}

	if receiver.AppID() != w.appid {
		return fmt.Errorf("appid mismatch, want: %s, got: %s", w.appid, receiver.AppID())
	}

	return nil
}

// Do exec action
func (w *WechatMP) Do(ctx context.Context, accessToken string, action Action, options ...helpers.HTTPOption) error {
	var (
		resp []byte
		err  error
	)

	arr := strings.SplitN(action.URL()(accessToken), "|", 2)

	switch arr[0] {
	case "GET":
		resp, err = w.client.Get(ctx, arr[1], options...)
	case "POST":
		resp, err = w.client.Post(ctx, arr[1], action.Body(), options...)
	case "UPLOAD":
		resp, err = w.client.Upload(ctx, arr[1], action.Body(), options...)
	}

	if err != nil {
		return err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return errors.New(r.Get("errmsg").String())
	}

	if action.Decode() == nil {
		return nil
	}

	return action.Decode()(resp)
}

// DecryptEventMessage 事件消息解密
func (w *WechatMP) DecryptEventMessage(cipherText string) (*event.Message, error) {
	b, err := event.Decrypt(w.appid, w.encodingAESKey, cipherText)

	if err != nil {
		return nil, err
	}

	msg := new(event.Message)

	if err = xml.Unmarshal(b, msg); err != nil {
		return nil, err
	}

	return msg, nil
}
