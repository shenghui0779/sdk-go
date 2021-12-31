package offia

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
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

// Offia 微信公众号
type Offia struct {
	appid          string
	appsecret      string
	originid       string
	token          string
	encodingAESKey string
	nonce          func() string
	client         wx.Client
}

// New returns new Offia
func New(appid, appsecret string) *Offia {
	return &Offia{
		appid:     appid,
		appsecret: appsecret,
		nonce: func() string {
			return wx.Nonce(16)
		},
		client: wx.DefaultClient(),
	}
}

// SetOriginID 设置原始ID（开发者微信号）
func (oa *Offia) SetOriginID(originid string) {
	oa.originid = originid
}

// SetServerConfig 设置服务器配置
// [参考](https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Access_Overview.html)
func (oa *Offia) SetServerConfig(token, encodingAESKey string) {
	oa.token = token
	oa.encodingAESKey = encodingAESKey
}

// SetClient sets options for wechat client
func (oa *Offia) SetClient(options ...wx.ClientOption) {
	oa.client.Set(options...)
}

// AppID returns appid
func (oa *Offia) AppID() string {
	return oa.appid
}

// AppSecret returns app secret
func (oa *Offia) AppSecret() string {
	return oa.appsecret
}

// OAuth2URL 生成网页授权URL（请使用 URLEncode 对 redirectURL 进行处理）
// [参考](https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html)
func (oa *Offia) OAuth2URL(scope AuthScope, redirectURL, state string) string {
	return fmt.Sprintf("%s?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect", urls.Oauth2Authorize, oa.appid, redirectURL, scope, state)
}

// SubscribeMsgAuthURL 公众号一次性订阅消息授权URL（请使用 URLEncode 对 redirectURL 进行处理）
// [参考](https://developers.weixin.qq.com/doc/offiaccount/Message_Management/One-time_subscription_info.html)
func (oa *Offia) SubscribeMsgAuthURL(scene, templateID, redirectURL, reserved string) string {
	return fmt.Sprintf("%s?action=get_confirm&appid=%s&template_id=%s&redirect_url=%s&reserved=%s#wechat_redirect", urls.SubscribeMsgAuth, oa.appid, templateID, redirectURL, reserved)
}

// Code2AuthToken 获取网页授权AccessToken
func (oa *Offia) Code2OAuthToken(ctx context.Context, code string, options ...yiigo.HTTPOption) (*OAuthToken, error) {
	resp, err := oa.client.Do(ctx, http.MethodGet, fmt.Sprintf("%s?appid=%s&secret=%s&code=%s&grant_type=authorization_code", urls.OffiaSnsCode2Token, oa.appid, oa.appsecret, code), nil, options...)

	if err != nil {
		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if code := r.Get("errcode").Int(); code != 0 {
		return nil, fmt.Errorf("%d|%s", code, r.Get("errmsg").String())
	}

	token := new(OAuthToken)

	if err = json.Unmarshal(resp, token); err != nil {
		return nil, err
	}

	return token, nil
}

// RefreshAuthToken 刷新网页授权AccessToken
func (oa *Offia) RefreshOAuthToken(ctx context.Context, refreshToken string, options ...yiigo.HTTPOption) (*OAuthToken, error) {
	resp, err := oa.client.Do(ctx, http.MethodGet, fmt.Sprintf("%s?appid=%s&grant_type=refresh_token&refresh_token=%s", urls.OffiaSnsRefreshAccessToken, oa.appid, refreshToken), nil, options...)

	if err != nil {
		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if code := r.Get("errcode").Int(); code != 0 {
		return nil, fmt.Errorf("%d|%s", code, r.Get("errmsg").String())
	}

	token := new(OAuthToken)

	if err = json.Unmarshal(resp, token); err != nil {
		return nil, err
	}

	return token, nil
}

// AccessToken 获取普通AccessToken
func (oa *Offia) AccessToken(ctx context.Context, options ...yiigo.HTTPOption) (*AccessToken, error) {
	resp, err := oa.client.Do(ctx, http.MethodGet, fmt.Sprintf("%s?grant_type=client_credential&appid=%s&secret=%s", urls.OffiaCgiBinAccessToken, oa.appid, oa.appsecret), nil, options...)

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
func (oa *Offia) Do(ctx context.Context, accessToken string, action wx.Action, options ...yiigo.HTTPOption) error {
	var (
		resp []byte
		err  error
	)

	if action.IsUpload() {
		form, ferr := action.UploadForm()

		if ferr != nil {
			return ferr
		}

		resp, err = oa.client.Upload(ctx, action.URL(accessToken), form, options...)
	} else {
		body, berr := action.Body()

		if berr != nil {
			return berr
		}

		resp, err = oa.client.Do(ctx, action.Method(), action.URL(accessToken), body, options...)
	}

	if err != nil {
		return err
	}

	r := gjson.ParseBytes(resp)

	if code := r.Get("errcode").Int(); code != 0 {
		return fmt.Errorf("%d|%s", code, r.Get("errmsg").String())
	}

	return action.Decode(resp)
}

// VerifyEventSign 验证消息事件签名
// 验证消息来自微信服务器，使用：signature、timestamp、nonce；若验证成功，请原样返回echostr参数内容
// 验证事件消息签名，使用：msg_signature、timestamp、nonce、msg_encrypt
// [参考](https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Access_Overview.html)
func (oa *Offia) VerifyEventSign(signature string, items ...string) bool {
	signStr := event.SignWithSHA1(oa.token, items...)

	return signStr == signature
}

// DecryptEventMessage 事件消息解密
func (oa *Offia) DecryptEventMessage(encrypt string) (wx.WXML, error) {
	b, err := event.Decrypt(oa.appid, oa.encodingAESKey, encrypt)

	if err != nil {
		return nil, err
	}

	return wx.ParseXML2Map(b)
}

// Reply 消息回复
func (oa *Offia) Reply(openid string, reply event.Reply) (*event.ReplyMessage, error) {
	body, err := reply.Bytes(oa.originid, openid)

	if err != nil {
		return nil, err
	}

	// 消息加密
	cipherText, err := event.Encrypt(oa.appid, oa.encodingAESKey, oa.nonce(), body)

	if err != nil {
		return nil, err
	}

	return event.BuildReply(oa.token, oa.nonce(), base64.StdEncoding.EncodeToString(cipherText)), nil
}

// JSSDKSign 生成 JS-SDK 签名
func (oa *Offia) JSSDKSign(ticket, url string) *JSSDKSign {
	noncestr := oa.nonce()
	now := time.Now().Unix()

	h := sha1.New()
	h.Write([]byte(fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", ticket, noncestr, now, url)))

	return &JSSDKSign{
		Signature: hex.EncodeToString(h.Sum(nil)),
		Noncestr:  noncestr,
		Timestamp: now,
	}
}
