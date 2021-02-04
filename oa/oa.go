package oa

import (
	"context"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
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
	client         wx.HTTPClient
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
// [参考](https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Access_Overview.html)
func (oa *OA) SetServerConfig(token, encodingAESKey string) {
	oa.token = token
	oa.encodingAESKey = encodingAESKey
}

// AuthURL 生成网页授权URL（请使用 URLEncode 对 redirectURL 进行处理）
func (oa *OA) AuthURL(scope AuthScope, redirectURL string) string {
	return fmt.Sprintf("%s?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect", AuthorizeURL, oa.appid, redirectURL, scope, oa.nonce(16))
}

// Code2AuthToken 获取网页授权AccessToken
func (oa *OA) Code2AuthToken(ctx context.Context, code string, options ...wx.HTTPOption) (*AuthToken, error) {
	resp, err := oa.client.Get(ctx, fmt.Sprintf("%s?appid=%s&secret=%s&code=%s&grant_type=authorization_code", SnsCode2TokenURL, oa.appid, oa.appsecret, code), options...)

	if err != nil {
		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if code := r.Get("errcode").Int(); code != 0 {
		return nil, fmt.Errorf("%d|%s", code, r.Get("errmsg").String())
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

	if code := r.Get("errcode").Int(); code != 0 {
		return nil, fmt.Errorf("%d|%s", code, r.Get("errmsg").String())
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

	if code := r.Get("errcode").Int(); code != 0 {
		return nil, fmt.Errorf("%d|%s", code, r.Get("errmsg").String())
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
		resp, err = oa.client.Get(ctx, action.URL(accessToken), options...)
	case wx.MethodPost:
		body, err := action.Body()

		if err != nil {
			return err
		}

		resp, err = oa.client.Post(ctx, action.URL(accessToken), body, options...)
	case wx.MethodUpload:
		resp, err = oa.client.Upload(ctx, action.URL(accessToken), action.UploadForm(), options...)
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

// VerifyEventSign 验证消息事件签名
// 验证消息来自微信服务器，使用：signature、timestamp、nonce；若验证成功，请原样返回echostr参数内容
// 验证事件消息签名，使用：msg_signature、timestamp、nonce、msg_encrypt
// [参考](https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Access_Overview.html)
func (oa *OA) VerifyEventSign(signature string, items ...string) bool {
	signStr := event.SignWithSHA1(oa.token, items...)

	return signStr == signature
}

// DecryptEventMessage 事件消息解密
func (oa *OA) DecryptEventMessage(encrypt string) (wx.WXML, error) {
	b, err := event.Decrypt(oa.appid, oa.encodingAESKey, encrypt)

	if err != nil {
		return nil, err
	}

	return wx.ParseXML2Map(b)
}

// Reply 消息回复
func (oa *OA) Reply(openid string, reply event.Reply) (*event.ReplyMessage, error) {
	body, err := reply.Bytes(oa.originid, openid)

	if err != nil {
		return nil, err
	}

	// 消息加密
	cipherText, err := event.Encrypt(oa.appid, oa.encodingAESKey, oa.nonce(16), body)

	if err != nil {
		return nil, err
	}

	return event.BuildReply(oa.token, oa.nonce(16), base64.StdEncoding.EncodeToString(cipherText)), nil
}

// JSSDKSign 生成 JS-SDK 签名
func (oa *OA) JSSDKSign(jsapiTicket, url string) *JSSDKSign {
	noncestr := oa.nonce(16)
	now := time.Now().Unix()

	h := sha1.New()
	h.Write([]byte(fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", jsapiTicket, noncestr, now, url)))

	return &JSSDKSign{
		Signature: hex.EncodeToString(h.Sum(nil)),
		Noncestr:  noncestr,
		Timestamp: now,
	}
}
