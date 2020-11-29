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
	"github.com/shenghui0779/gochat/public"
	"github.com/tidwall/gjson"
)

// OA 微信公众号
type OA struct {
	appid          string
	appsecret      string
	signToken      string
	encodingAESKey string
	nonce          func(size int) string
	client         public.Client
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
		client: public.NewHTTPClient(),
	}
}

// SetServerConfig 设置服务器配置
func (oa *OA) SetServerConfig(token, encodingAESKey string) {
	oa.signToken = token
	oa.encodingAESKey = encodingAESKey
}

// Code2AuthToken 获取公众号网页授权AccessToken
func (oa *OA) Code2AuthToken(ctx context.Context, code string, options ...public.HTTPOption) (*AuthToken, error) {
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
func (oa *OA) RefreshAuthToken(ctx context.Context, refreshToken string, options ...public.HTTPOption) (*AuthToken, error) {
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
func (oa *OA) AccessToken(ctx context.Context, options ...public.HTTPOption) (*AccessToken, error) {
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
func (oa *OA) Do(ctx context.Context, accessToken string, action public.Action, options ...public.HTTPOption) error {
	var (
		resp []byte
		err  error
	)

	switch action.Method() {
	case public.MethodGet:
		resp, err = oa.client.Get(ctx, action.URL()(accessToken), options...)
	case public.MethodPost:
		resp, err = oa.client.Post(ctx, action.URL()(accessToken), action.Body(), options...)
	case public.MethodUpload:
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
func (oa *OA) DecryptEventMessage(cipherText string) (*event.Message, error) {
	b, err := event.Decrypt(oa.appid, oa.encodingAESKey, cipherText)

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
func (oa *OA) EncryptReplyMessage(from, to string, reply Reply) (*ReplyMessage, error) {
	body, err := reply.Bytes(from, to)

	if err != nil {
		return nil, err
	}

	// 消息加密
	cipherText, err := event.Encrypt(oa.appid, oa.encodingAESKey, oa.nonce(16), body)

	if err != nil {
		return nil, err
	}

	encryptData := base64.StdEncoding.EncodeToString(cipherText)

	now := time.Now().Unix()
	nonce := oa.nonce(16)

	signItems := []string{oa.signToken, strconv.FormatInt(now, 10), nonce, encryptData}

	sort.Strings(signItems)

	h := sha1.New()
	h.Write([]byte(strings.Join(signItems, "")))

	msg := &ReplyMessage{
		Encrypt:      public.CDATA(encryptData),
		MsgSignature: public.CDATA(hex.EncodeToString(h.Sum(nil))),
		TimeStamp:    now,
		Nonce:        public.CDATA(nonce),
	}

	return msg, nil
}
