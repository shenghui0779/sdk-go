package oa

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/shenghui0779/yiigo"
	"github.com/tidwall/gjson"

	"github.com/shenghui0779/gochat/event"
	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

// AuthScope 应用授权作用域
type AuthScope string

// 公众号支持的应用授权作用域
const (
	ScopeSnsapiBase AuthScope = "snsapi_base"     // 静默授权使用，不弹出授权页面，直接跳转，只能获取用户openid
	ScopeSnsapiUser AuthScope = "snsapi_userinfo" // 弹出授权页面，可通过openid拿到昵称、性别、所在地。并且，即使在未关注的情况下，只要用户授权，也能获取其信息
)

// OA 微信公众号
type OA struct {
	appid          string
	appsecret      string
	originid       string
	token          string
	encodingAESKey string
	nonce          func(size uint) string
	client         wx.Client
}

// New returns new OA
func New(appid, appsecret string) *OA {
	return &OA{
		appid:     appid,
		appsecret: appsecret,
		nonce:     wx.Nonce,
		client:    wx.NewClient(wx.WithInsecureSkipVerify()),
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

// AppID returns appid
func (oa *OA) AppID() string {
	return oa.appid
}

// AppSecret returns app secret
func (oa *OA) AppSecret() string {
	return oa.appsecret
}

// AuthURL 生成网页授权URL（请使用 URLEncode 对 redirectURL 进行处理）
// [参考](https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html)
func (oa *OA) AuthURL(scope AuthScope, redirectURL string, state ...string) string {
	paramState := oa.nonce(16)

	if len(state) != 0 {
		paramState = state[0]
	}

	return fmt.Sprintf("%s?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect", urls.OAAuthorize, oa.appid, redirectURL, scope, paramState)
}

// Code2AuthToken 获取网页授权AccessToken
func (oa *OA) Code2AuthToken(ctx context.Context, code string, options ...yiigo.HTTPOption) (*AuthToken, error) {
	resp, err := oa.client.Get(ctx, fmt.Sprintf("%s?appid=%s&secret=%s&code=%s&grant_type=authorization_code", urls.OASnsCode2Token, oa.appid, oa.appsecret, code), options...)

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
func (oa *OA) RefreshAuthToken(ctx context.Context, refreshToken string, options ...yiigo.HTTPOption) (*AuthToken, error) {
	resp, err := oa.client.Get(ctx, fmt.Sprintf("%s?appid=%s&grant_type=refresh_token&refresh_token=%s", urls.OASnsRefreshAccessToken, oa.appid, refreshToken), options...)

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
func (oa *OA) AccessToken(ctx context.Context, options ...yiigo.HTTPOption) (*AccessToken, error) {
	resp, err := oa.client.Get(ctx, fmt.Sprintf("%s?grant_type=client_credential&appid=%s&secret=%s", urls.OACgiBinAccessToken, oa.appid, oa.appsecret), options...)

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
func (oa *OA) Do(ctx context.Context, accessToken string, action wx.Action, options ...yiigo.HTTPOption) error {
	var (
		resp []byte
		err  error
	)

	switch action.Method() {
	case wx.MethodGet:
		resp, err = oa.client.Get(ctx, action.URL(accessToken), options...)
	case wx.MethodPost:
		body, berr := action.Body()

		if berr != nil {
			return berr
		}

		resp, err = oa.client.Post(ctx, action.URL(accessToken), body, options...)
	case wx.MethodUpload:
		form, ferr := action.UploadForm()

		if ferr != nil {
			return ferr
		}

		resp, err = oa.client.Upload(ctx, action.URL(accessToken), form, options...)
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
