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
	"github.com/shenghui0779/gochat/wx"
	"github.com/tidwall/gjson"
)

// OA 微信公众号
type OA struct {
	appid          string
	appsecret      string
	originid       string
	token          string
	encodingAESKey string
	nonce          func(size int) string
	client         wx.Client
}

// New returns new OA
func New(appid, appsecret string) *OA {
	return &OA{
		appid:     appid,
		appsecret: appsecret,
		nonce: func(size int) string {
			nonce := make([]byte, size/2)
			io.ReadFull(rand.Reader, nonce)

			return hex.EncodeToString(nonce)
		},
		client: wx.NewHTTPClient(),
	}
}

// SetOriginID 设置原始ID（开发者微信号）
func (oa *OA) SetOriginID(originid string) {
	oa.originid = originid
}

// SetServerConfig 设置服务器配置
func (oa *OA) SetServerConfig(token, encodingAESKey string) {
	oa.token = token
	oa.encodingAESKey = encodingAESKey
}

// AuthURL 生成网页授权URL（请使用 URLEncode 对 redirectURL 进行处理）
func (oa *OA) AuthURL(scope AuthScope, redirectURL string) string {
	return fmt.Sprintf("%s?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect", AuthorizeURL, oa.appid, redirectURL, scope, oa.nonce(16))
}

// Code2AuthToken 获取公众号网页授权AccessToken
func (oa *OA) Code2AuthToken(ctx context.Context, code string, options ...wx.HTTPOption) (*AuthToken, error) {
	resp, err := oa.client.Get(ctx, fmt.Sprintf("%s?appid=%s&secret=%s&code=%s&grant_type=authorization_code", SnsCode2TokenURL, oa.appid, oa.appsecret, code), options...)

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
func (oa *OA) RefreshAuthToken(ctx context.Context, refreshToken string, options ...wx.HTTPOption) (*AuthToken, error) {
	resp, err := oa.client.Get(ctx, fmt.Sprintf("%s?appid=%s&grant_type=refresh_token&refresh_token=%s", SnsRefreshAccessTokenURL, oa.appid, refreshToken), options...)

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
func (oa *OA) AccessToken(ctx context.Context, options ...wx.HTTPOption) (*AccessToken, error) {
	resp, err := oa.client.Get(ctx, fmt.Sprintf("%s?grant_type=client_credential&appid=%s&secret=%s", CgiBinAccessTokenURL, oa.appid, oa.appsecret), options...)

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
func (oa *OA) Do(ctx context.Context, accessToken string, action wx.Action, options ...wx.HTTPOption) error {
	var (
		resp []byte
		err  error
	)

	switch action.Method() {
	case wx.MethodGet:
		resp, err = oa.client.Get(ctx, action.URL()(accessToken), options...)
	case wx.MethodPost:
		resp, err = oa.client.Post(ctx, action.URL()(accessToken), action.Body(), options...)
	case wx.MethodUpload:
		resp, err = oa.client.Upload(ctx, action.URL()(accessToken), action.Body(), options...)
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
func (oa *OA) DecryptEventMessage(msgEncrypt string) (*event.Message, error) {
	b, err := event.Decrypt(oa.appid, oa.encodingAESKey, msgEncrypt)

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
func (oa *OA) EncryptReplyMessage(openid string, reply Reply) (*ReplyMessage, error) {
	body, err := reply.Bytes(oa.originid, openid)

	if err != nil {
		return nil, err
	}

	// 消息加密
	cipherText, err := event.Encrypt(oa.appid, oa.encodingAESKey, oa.nonce(16), body)

	if err != nil {
		return nil, err
	}

	encryptMsg := base64.StdEncoding.EncodeToString(cipherText)

	// 签名
	now := time.Now().Unix()
	nonce := oa.nonce(16)

	signItems := []string{oa.token, strconv.FormatInt(now, 10), nonce, encryptMsg}

	sort.Strings(signItems)

	h := sha1.New()
	h.Write([]byte(strings.Join(signItems, "")))

	msg := &ReplyMessage{
		Encrypt:      wx.CDATA(encryptMsg),
		MsgSignature: wx.CDATA(hex.EncodeToString(h.Sum(nil))),
		TimeStamp:    now,
		Nonce:        wx.CDATA(nonce),
	}

	return msg, nil
}

// VerifyServer 验证消息来自微信服务器（若验证成功，请原样返回echostr参数内容）
func (oa *OA) VerifyServer(signature, timestamp, nonce string) bool {
	signItems := []string{oa.token, timestamp, nonce}

	sort.Strings(signItems)

	h := sha1.New()
	h.Write([]byte(strings.Join(signItems, "")))
	signStr := hex.EncodeToString(h.Sum(nil))

	return signStr == signature
}

// VerifyEvent 验证事件签名（注意：使用 msg_signature）
func (oa *OA) VerifyEvent(msgSignature, timestamp, nonce, msgEncrypt string) bool {
	signItems := []string{oa.token, timestamp, nonce, msgEncrypt}

	sort.Strings(signItems)

	h := sha1.New()
	h.Write([]byte(strings.Join(signItems, "")))
	signStr := hex.EncodeToString(h.Sum(nil))

	return signStr == msgSignature
}
