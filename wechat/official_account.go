package wechat

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/tidwall/gjson"

	"github.com/shenghui0779/sdk-go/lib"
	"github.com/shenghui0779/sdk-go/lib/curl"
	"github.com/shenghui0779/sdk-go/lib/value"
)

// ServerConfig 服务器配置
type ServerConfig struct {
	token  string
	aeskey string
}

// OfficialAccount 微信公众号
type OfficialAccount struct {
	host    string
	appid   string
	secret  string
	srvCfg  *ServerConfig
	httpCli curl.Client
	logger  func(ctx context.Context, data map[string]string)
}

// AppID returns appid
func (oa *OfficialAccount) AppID() string {
	return oa.appid
}

// Secret returns app secret
func (oa *OfficialAccount) Secret() string {
	return oa.secret
}

// URL 生成请求URL
func (oa *OfficialAccount) url(path string, query url.Values) string {
	var builder strings.Builder

	builder.WriteString(oa.host)
	if len(path) != 0 && path[0] != '/' {
		builder.WriteString("/")
	}
	builder.WriteString(path)
	if len(query) != 0 {
		builder.WriteString("?")
		builder.WriteString(query.Encode())
	}

	return builder.String()
}

func (oa *OfficialAccount) do(ctx context.Context, method, path string, query url.Values, params lib.X, options ...curl.Option) ([]byte, error) {
	reqURL := oa.url(path, query)

	log := lib.NewReqLog(method, reqURL)
	defer log.Do(ctx, oa.logger)

	var (
		body []byte
		err  error
	)

	if params != nil {
		body, err = json.Marshal(params)
		if err != nil {
			return nil, err
		}
		log.SetReqBody(string(body))
	}

	resp, err := oa.httpCli.Do(ctx, method, reqURL, body, options...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	log.SetRespHeader(resp.Header)
	log.SetStatusCode(resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", resp.StatusCode)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.SetRespBody(string(b))

	return b, nil
}

// OAuth2URL 生成网页授权URL
// [参考](https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html)
func (oa *OfficialAccount) OAuth2URL(scope AuthScope, redirectURI, state string) string {
	query := url.Values{}

	query.Set("appid", oa.appid)
	query.Set("redirect_uri", redirectURI)
	query.Set("response_type", "code")
	query.Set("scope", string(scope))
	query.Set("state", state)

	return fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?%s#wechat_redirect", query.Encode())
}

// SubscribeMsgAuthURL 公众号一次性订阅消息授权URL
// [参考](https://developers.weixin.qq.com/doc/offiaccount/Message_Management/One-time_subscription_info.html)
func (oa *OfficialAccount) SubscribeMsgAuthURL(scene, templateID, redirectURL, reserved string) string {
	query := url.Values{}

	query.Set("appid", oa.appid)
	query.Set("action", "get_confirm")
	query.Set("template_id", templateID)
	query.Set("redirect_url", redirectURL)
	query.Set("reserved", reserved)

	return fmt.Sprintf("https://oa.weixin.qq.com/oa/subscribemsg?%s#wechat_redirect", query.Encode())
}

// Code2OAuthToken 获取网页授权Token
func (oa *OfficialAccount) Code2OAuthToken(ctx context.Context, code string) (gjson.Result, error) {
	query := url.Values{}

	query.Set("appid", oa.appid)
	query.Set("secret", oa.secret)
	query.Set("code", code)
	query.Set("grant_type", "authorization_code")

	b, err := oa.do(ctx, http.MethodGet, "/sns/oauth2/access_token", query, nil)
	if err != nil {
		return lib.Fail(err)
	}

	ret := gjson.ParseBytes(b)
	if code := ret.Get("errcode").Int(); code != 0 {
		return lib.Fail(fmt.Errorf("%d | %s", code, ret.Get("errmsg").String()))
	}
	return ret, nil
}

// RefreshOAuthToken 刷新网页授权Token
func (oa *OfficialAccount) RefreshOAuthToken(ctx context.Context, refreshToken string) (gjson.Result, error) {
	query := url.Values{}

	query.Set("appid", oa.appid)
	query.Set("grant_type", "refresh_token")
	query.Set("refresh_token", refreshToken)

	b, err := oa.do(ctx, http.MethodGet, "/sns/oauth2/refresh_token", query, nil)
	if err != nil {
		return lib.Fail(err)
	}

	ret := gjson.ParseBytes(b)
	if code := ret.Get("errcode").Int(); code != 0 {
		return lib.Fail(fmt.Errorf("%d | %s", code, ret.Get("errmsg").String()))
	}
	return ret, nil
}

// AccessToken 获取接口调用凭据
func (oa *OfficialAccount) AccessToken(ctx context.Context) (gjson.Result, error) {
	query := url.Values{}

	query.Set("appid", oa.appid)
	query.Set("secret", oa.secret)
	query.Set("grant_type", "client_credential")

	b, err := oa.do(ctx, http.MethodGet, "/cgi-bin/token", query, nil)
	if err != nil {
		return lib.Fail(err)
	}

	ret := gjson.ParseBytes(b)
	if code := ret.Get("errcode").Int(); code != 0 {
		return lib.Fail(fmt.Errorf("%d | %s", code, ret.Get("errmsg").String()))
	}
	return ret, nil
}

// StableAccessToken 获取稳定版接口调用凭据，有两种调用模式:
// 1. 普通模式，access_token 有效期内重复调用该接口不会更新 access_token，绝大部分场景下使用该模式；
// 2. 强制刷新模式，会导致上次获取的 access_token 失效，并返回新的 access_token
func (oa *OfficialAccount) StableAccessToken(ctx context.Context, forceRefresh bool) (gjson.Result, error) {
	params := lib.X{
		"grant_type":    "client_credential",
		"appid":         oa.appid,
		"secret":        oa.secret,
		"force_refresh": forceRefresh,
	}

	b, err := oa.do(ctx, http.MethodPost, "/cgi-bin/stable_token", nil, params, curl.WithHeader(curl.HeaderContentType, curl.ContentJSON))
	if err != nil {
		return lib.Fail(err)
	}

	ret := gjson.ParseBytes(b)
	if code := ret.Get("errcode").Int(); code != 0 {
		return lib.Fail(fmt.Errorf("%d | %s", code, ret.Get("errmsg").String()))
	}
	return ret, nil
}

// GetJSON GET请求JSON数据
func (oa *OfficialAccount) GetJSON(ctx context.Context, accessToken, path string, query url.Values) (gjson.Result, error) {
	if query == nil {
		query = url.Values{}
	}
	query.Set(AccessToken, accessToken)

	b, err := oa.do(ctx, http.MethodGet, path, query, nil)
	if err != nil {
		return lib.Fail(err)
	}

	ret := gjson.ParseBytes(b)
	if code := ret.Get("errcode").Int(); code != 0 {
		return lib.Fail(fmt.Errorf("%d | %s", code, ret.Get("errmsg").String()))
	}
	return ret, nil
}

// PostJSON POST请求JSON数据
func (oa *OfficialAccount) PostJSON(ctx context.Context, accessToken, path string, params lib.X) (gjson.Result, error) {
	query := url.Values{}
	query.Set(AccessToken, accessToken)

	b, err := oa.do(ctx, http.MethodPost, path, query, params, curl.WithHeader(curl.HeaderContentType, curl.ContentJSON))
	if err != nil {
		return lib.Fail(err)
	}

	ret := gjson.ParseBytes(b)
	if code := ret.Get("errcode").Int(); code != 0 {
		return lib.Fail(fmt.Errorf("%d | %s", code, ret.Get("errmsg").String()))
	}
	return ret, nil
}

// GetBuffer GET请求获取buffer (如：获取媒体资源)
func (oa *OfficialAccount) GetBuffer(ctx context.Context, accessToken, path string, query url.Values) ([]byte, error) {
	if query == nil {
		query = url.Values{}
	}
	query.Set(AccessToken, accessToken)

	b, err := oa.do(ctx, http.MethodGet, path, query, nil)
	if err != nil {
		return nil, err
	}

	ret := gjson.ParseBytes(b)
	if code := ret.Get("errcode").Int(); code != 0 {
		return nil, fmt.Errorf("%d | %s", code, ret.Get("errmsg").String())
	}
	return b, nil
}

// PostBuffer POST请求获取buffer (如：获取二维码)
func (oa *OfficialAccount) PostBuffer(ctx context.Context, accessToken, path string, params lib.X) ([]byte, error) {
	query := url.Values{}
	query.Set(AccessToken, accessToken)

	b, err := oa.do(ctx, http.MethodPost, path, query, params, curl.WithHeader(curl.HeaderContentType, curl.ContentJSON))
	if err != nil {
		return nil, err
	}

	ret := gjson.ParseBytes(b)
	if code := ret.Get("errcode").Int(); code != 0 {
		return nil, fmt.Errorf("%d | %s", code, ret.Get("errmsg").String())
	}
	return b, nil
}

// Upload 上传媒体资源
func (oa *OfficialAccount) Upload(ctx context.Context, accessToken, path string, form curl.UploadForm) (gjson.Result, error) {
	query := url.Values{}
	query.Set(AccessToken, accessToken)

	reqURL := oa.url(path, query)

	log := lib.NewReqLog(http.MethodPost, reqURL)
	defer log.Do(ctx, oa.logger)

	resp, err := oa.httpCli.Upload(ctx, reqURL, form)
	if err != nil {
		return lib.Fail(err)
	}
	defer resp.Body.Close()

	log.SetRespHeader(resp.Header)
	log.SetStatusCode(resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		return lib.Fail(fmt.Errorf("HTTP Request Error, StatusCode = %d", resp.StatusCode))
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return lib.Fail(err)
	}
	log.SetRespBody(string(b))

	ret := gjson.ParseBytes(b)
	if code := ret.Get("errcode").Int(); code != 0 {
		return lib.Fail(fmt.Errorf("%d | %s", code, ret.Get("errmsg").String()))
	}
	return ret, nil
}

// VerifyURL 服务器URL验证，使用：signature、timestamp、nonce（若验证成功，请原样返回echostr参数内容）
// [参考](https://developers.weixin.qq.com/miniprogram/dev/framework/server-ability/message-push.html)
func (oa *OfficialAccount) VerifyURL(signature, timestamp, nonce string) error {
	if SignWithSHA1(oa.srvCfg.token, timestamp, nonce) != signature {
		return errors.New("signature verified fail")
	}
	return nil
}

// DecodeEventMsg 解析事件消息，使用：msg_signature、timestamp、nonce、msg_encrypt
// [参考](https://developers.weixin.qq.com/miniprogram/dev/framework/server-ability/message-push.html)
func (oa *OfficialAccount) DecodeEventMsg(signature, timestamp, nonce, encryptMsg string) (value.V, error) {
	if SignWithSHA1(oa.srvCfg.token, timestamp, nonce, encryptMsg) != signature {
		return nil, errors.New("signature verified fail")
	}

	b, err := EventDecrypt(oa.appid, oa.srvCfg.aeskey, encryptMsg)
	if err != nil {
		return nil, err
	}
	return ParseXMLToV(b)
}

// ReplyEventMsg 事件消息回复
func (oa *OfficialAccount) ReplyEventMsg(msg value.V) (value.V, error) {
	return EventReply(oa.appid, oa.srvCfg.token, oa.srvCfg.aeskey, msg)
}

// OAOption 公众号设置项
type OAOption func(oa *OfficialAccount)

// WithOASrvCfg 设置公众号服务器配置
// [参考](https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Access_Overview.html)
func WithOASrvCfg(token, aeskey string) OAOption {
	return func(oa *OfficialAccount) {
		oa.srvCfg.token = token
		oa.srvCfg.aeskey = aeskey
	}
}

// WithOAHttpCli 设置公众号请求的 HTTP Client
func WithOAHttpCli(c *http.Client) OAOption {
	return func(oa *OfficialAccount) {
		oa.httpCli = curl.NewHTTPClient(c)
	}
}

// WithOALogger 设置公众号日志记录
func WithOALogger(fn func(ctx context.Context, data map[string]string)) OAOption {
	return func(oa *OfficialAccount) {
		oa.logger = fn
	}
}

// NewOfficialAccount 生成一个公众号实例
func NewOfficialAccount(appid, secret string, options ...OAOption) *OfficialAccount {
	oa := &OfficialAccount{
		host:    "https://api.weixin.qq.com",
		appid:   appid,
		secret:  secret,
		srvCfg:  new(ServerConfig),
		httpCli: curl.NewDefaultClient(),
	}
	for _, fn := range options {
		fn(oa)
	}
	return oa
}
