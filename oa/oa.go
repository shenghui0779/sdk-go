package oa

import (
	"context"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/shenghui0779/gochat/event"
	"github.com/shenghui0779/gochat/internal"
	"github.com/tidwall/gjson"
)

// WechatOA 微信公众号
type WechatOA struct {
	appid          string
	appsecret      string
	signToken      string
	encodingAESKey string
	nonce          func(size int) string
	client         internal.Client
}

// New returns new WechatOA
func New(appid, appsecret string) *WechatOA {
	return &WechatOA{
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
func (w *WechatOA) SetServerConfig(token, encodingAESKey string) {
	w.signToken = token
	w.encodingAESKey = encodingAESKey
}

// Code2AuthToken 获取公众号网页授权AccessToken
func (w *WechatOA) Code2AuthToken(ctx context.Context, code string, options ...internal.HTTPOption) (*AuthToken, error) {
	resp, err := w.client.Get(ctx, fmt.Sprintf("%s?appid=%s&secret=%s&code=%s&grant_type=authorization_code", SnsCode2TokenURL, w.appid, w.appsecret, code), options...)

	if err != nil {
		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return nil, errors.New(r.Get("errmsg").String())
	}

	token := new(AuthToken)

	if err = json.Unmarshal(resp, token); err != nil {
		return nil, err
	}

	return token, nil
}

// RefreshAuthToken 刷新网页授权AccessToken
func (w *WechatOA) RefreshAuthToken(ctx context.Context, refreshToken string, options ...internal.HTTPOption) (*AuthToken, error) {
	resp, err := w.client.Get(ctx, fmt.Sprintf("%s?appid=%s&grant_type=refresh_token&refresh_token=%s", SnsRefreshAccessTokenURL, w.appid, refreshToken), options...)

	if err != nil {
		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return nil, errors.New(r.Get("errmsg").String())
	}

	token := new(AuthToken)

	if err = json.Unmarshal(resp, token); err != nil {
		return nil, err
	}

	return token, nil
}

// AccessToken 获取普通AccessToken
func (w *WechatOA) AccessToken(ctx context.Context, options ...internal.HTTPOption) (*AccessToken, error) {
	resp, err := w.client.Get(ctx, fmt.Sprintf("%s?grant_type=client_credential&appid=%s&secret=%s", CgiBinAccessTokenURL, w.appid, w.appsecret), options...)

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

// Do exec action
func (w *WechatOA) Do(ctx context.Context, accessToken string, action internal.Action, options ...internal.HTTPOption) error {
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
func (w *WechatOA) DecryptEventMessage(cipherText string) (*event.Message, error) {
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

// EncryptReplyMessage 回复消息加密
func (w *WechatOA) EncryptReplyMessage(from, to string, reply Reply) (*ReplyMessage, error) {
	body, err := reply.Bytes(from, to)

	if err != nil {
		return nil, err
	}

	// 消息加密
	cipherText, err := event.Encrypt(w.appid, w.encodingAESKey, w.nonce(16), body)

	if err != nil {
		return nil, err
	}

	encryptData := base64.StdEncoding.EncodeToString(cipherText)

	now := time.Now().Unix()
	nonce := w.nonce(16)

	signItems := []string{w.signToken, strconv.FormatInt(now, 10), nonce, encryptData}

	sort.Strings(signItems)

	h := sha1.New()
	h.Write([]byte(strings.Join(signItems, "")))

	msg := &ReplyMessage{
		Encrypt:      internal.CDATA(encryptData),
		MsgSignature: internal.CDATA(hex.EncodeToString(h.Sum(nil))),
		TimeStamp:    now,
		Nonce:        internal.CDATA(nonce),
	}

	return msg, nil
}
