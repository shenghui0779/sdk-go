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

	"github.com/shenghui0779/gochat/event"
	"github.com/shenghui0779/gochat/internal"
	"github.com/tidwall/gjson"
)

// WechatMP 微信小程序
type WechatMP struct {
	appid          string
	appsecret      string
	signToken      string
	encodingAESKey string
	nonce          func(size int) string
	client         internal.Client
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
		client: internal.NewHTTPClient(),
	}
}

// SetServerConfig 设置服务器配置
func (w *WechatMP) SetServerConfig(token, encodingAESKey string) {
	w.signToken = token
	w.encodingAESKey = encodingAESKey
}

// Code2Session 获取小程序授权的session_key
func (w *WechatMP) Code2Session(ctx context.Context, code string, options ...internal.HTTPOption) (*AuthSession, error) {
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
func (w *WechatMP) AccessToken(ctx context.Context, options ...internal.HTTPOption) (*AccessToken, error) {
	resp, err := w.client.Get(ctx, fmt.Sprintf("%s?appid=%s&secret=%s&grant_type=client_credential", AccessTokenURL, w.appid, w.appsecret), options...)

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
func (w *WechatMP) DecryptAuthInfo(sessionKey, iv, encryptedData string, dest AuthInfo) error {
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

	cbc := internal.NewAESCBCCrypto(key, ivb)

	b, err := cbc.Decrypt(cipherText, internal.PKCS7)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, dest); err != nil {
		return err
	}

	if dest.AppID() != w.appid {
		return fmt.Errorf("appid mismatch, want: %s, got: %s", w.appid, dest.AppID())
	}

	return nil
}

// Do exec action
func (w *WechatMP) Do(ctx context.Context, accessToken string, action internal.Action, options ...internal.HTTPOption) error {
	var (
		resp []byte
		err  error
	)

	switch action.Method() {
	case internal.MethodGet:
		resp, err = w.client.Get(ctx, action.URL()(accessToken), options...)
	case internal.MethodPost:
		resp, err = w.client.Post(ctx, action.URL()(accessToken), action.Body(), options...)
	case internal.MethodUpload:
		resp, err = w.client.Upload(ctx, action.URL()(accessToken), action.Body(), options...)
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
