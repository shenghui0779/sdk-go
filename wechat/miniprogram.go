package wechat

import (
	"context"
	"crypto"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/tidwall/gjson"

	"github.com/shenghui0779/sdk-go/lib"
	"github.com/shenghui0779/sdk-go/lib/value"
	"github.com/shenghui0779/sdk-go/lib/xcrypto"
	"github.com/shenghui0779/sdk-go/lib/xhttp"
)

// SafeMode 安全鉴权模式配置
type SafeMode struct {
	aesSN  string
	aeskey string
	prvKey *xcrypto.PrivateKey
	pubSN  string
	pubKey *xcrypto.PublicKey
}

// MiniProgram 小程序
type MiniProgram struct {
	host    string
	appid   string
	secret  string
	srvCfg  *ServerConfig
	sfMode  *SafeMode
	token   atomic.Value
	httpCli xhttp.Client
	logger  func(ctx context.Context, data map[string]string)
}

// AppID 返回appid
func (mp *MiniProgram) AppID() string {
	return mp.appid
}

// Secret 返回secret
func (mp *MiniProgram) Secret() string {
	return mp.secret
}

func (mp *MiniProgram) url(path string, query url.Values) string {
	var builder strings.Builder

	builder.WriteString(mp.host)
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

func (mp *MiniProgram) do(ctx context.Context, method, path string, query url.Values, params lib.X, options ...xhttp.Option) ([]byte, error) {
	reqURL := mp.url(path, query)

	log := lib.NewReqLog(method, reqURL)
	defer log.Do(ctx, mp.logger)

	var (
		body []byte
		err  error
	)

	if params != nil {
		body, err = json.Marshal(params)
		if err != nil {
			log.Set("error", err.Error())
			return nil, err
		}
		log.SetReqBody(string(body))
	}

	resp, err := mp.httpCli.Do(ctx, method, reqURL, body, options...)
	if err != nil {
		log.Set("error", err.Error())
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
		log.Set("error", err.Error())
		return nil, err
	}
	log.SetRespBody(string(b))

	return b, nil
}

func (mp *MiniProgram) doSafe(ctx context.Context, method, path string, query url.Values, params lib.X) ([]byte, error) {
	reqURL := mp.url(path, query)

	log := lib.NewReqLog(method, reqURL)
	defer log.Do(ctx, mp.logger)

	now := time.Now().Unix()

	// 加密
	params, err := mp.encrypt(log, path, query, params, now)
	if err != nil {
		log.Set("error", err.Error())
		return nil, err
	}

	body, err := json.Marshal(params)
	if err != nil {
		log.Set("error", err.Error())
		return nil, err
	}
	log.SetReqBody(string(body))

	// 签名
	sign, err := mp.sign(path, now, body)
	if err != nil {
		log.Set("error", err.Error())
		return nil, err
	}

	reqHeader := http.Header{}

	reqHeader.Set(xhttp.HeaderContentType, xhttp.ContentJSON)
	reqHeader.Set(HeaderMPAppID, mp.appid)
	reqHeader.Set(HeaderMPTimestamp, strconv.FormatInt(now, 10))
	reqHeader.Set(HeaderMPSignature, sign)

	log.SetReqHeader(reqHeader)

	resp, err := mp.httpCli.Do(ctx, method, reqURL, body, lib.HeaderToHttpOption(reqHeader)...)
	if err != nil {
		log.Set("error", err.Error())
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
		log.Set("error", err.Error())
		return nil, err
	}
	log.SetRespBody(string(b))

	// 验签
	if err = mp.verify(path, resp.Header, b); err != nil {
		log.Set("error", err.Error())
		return nil, err
	}

	// 解密
	data, err := mp.decrypt(path, resp.Header, b)
	if err != nil {
		log.Set("error", err.Error())
		return nil, err
	}

	log.Set("origin_response_body", string(data))

	return data, nil
}

func (mp *MiniProgram) encrypt(log *lib.ReqLog, path string, query url.Values, params lib.X, timestamp int64) (lib.X, error) {
	if len(mp.sfMode.aeskey) == 0 {
		return nil, errors.New("aes-gcm key not found (forgotten configure?)")
	}

	if params == nil {
		params = lib.X{}
	}

	params["_n"] = base64.StdEncoding.EncodeToString(lib.NonceByte(16))
	params["_appid"] = mp.appid
	params["_timestamp"] = timestamp

	for k, v := range query {
		if k != AccessToken && len(v) != 0 {
			params[k] = v[0]
		}
	}

	data, err := json.Marshal(params)
	if err != nil {
		log.Set("error", err.Error())
		return nil, err
	}

	log.Set("origin_request_body", string(data))

	key, err := base64.StdEncoding.DecodeString(mp.sfMode.aeskey)
	if err != nil {
		log.Set("error", err.Error())
		return nil, err
	}

	iv := lib.NonceByte(12)
	aad := fmt.Sprintf("%s|%s|%d|%s", mp.url(path, nil), mp.appid, timestamp, mp.sfMode.aesSN)

	ct, err := xcrypto.AESEncryptGCM(key, iv, data, []byte(aad), nil)
	if err != nil {
		log.Set("error", err.Error())
		return nil, err
	}

	body := lib.X{
		"iv":      base64.StdEncoding.EncodeToString(iv),
		"data":    base64.StdEncoding.EncodeToString(ct.Data()),
		"authtag": base64.StdEncoding.EncodeToString(ct.Tag()),
	}
	return body, nil
}

func (mp *MiniProgram) sign(path string, timestamp int64, body []byte) (string, error) {
	if mp.sfMode.prvKey == nil {
		return "", errors.New("private key not found (forgotten configure?)")
	}

	var builder strings.Builder

	builder.WriteString(mp.url(path, nil))
	builder.WriteString("\n")
	builder.WriteString(mp.appid)
	builder.WriteString("\n")
	builder.WriteString(strconv.FormatInt(timestamp, 10))
	builder.WriteString("\n")
	builder.Write(body)

	b, err := mp.sfMode.prvKey.SignPSS(crypto.SHA256, []byte(builder.String()), &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthEqualsHash})
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func (mp *MiniProgram) verify(path string, header http.Header, body []byte) error {
	if mp.sfMode.pubKey == nil {
		return errors.New("public key not found (forgotten configure?)")
	}

	if appid := header.Get(HeaderMPAppID); appid != mp.appid {
		return fmt.Errorf("header appid mismatch, expect = %s", mp.appid)
	}

	sign := ""
	if serial := header.Get(HeaderMPSerial); serial == mp.sfMode.pubSN {
		sign = header.Get(HeaderMPSignature)
	} else {
		serialDeprecated := header.Get(HeaderMPSerialDeprecated)
		if serialDeprecated != mp.sfMode.pubSN {
			return fmt.Errorf("header serial mismatch, expect = %s", mp.sfMode.pubSN)
		}
		sign = header.Get(HeaderMPSignatureDeprecated)
	}
	b, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}

	var builder strings.Builder

	builder.WriteString(mp.url(path, nil))
	builder.WriteString("\n")
	builder.WriteString(mp.appid)
	builder.WriteString("\n")
	builder.WriteString(header.Get(HeaderMPTimestamp))
	builder.WriteString("\n")
	builder.Write(body)

	return mp.sfMode.pubKey.VerifyPSS(crypto.SHA256, []byte(builder.String()), b, &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthEqualsHash})
}

func (mp *MiniProgram) decrypt(path string, header http.Header, body []byte) ([]byte, error) {
	if len(mp.sfMode.aeskey) == 0 {
		return nil, errors.New("aes-gcm key not found (forgotten configure?)")
	}

	key, err := base64.StdEncoding.DecodeString(mp.sfMode.aeskey)
	if err != nil {
		return nil, err
	}

	ret := gjson.ParseBytes(body)

	iv, err := base64.StdEncoding.DecodeString(ret.Get("iv").String())
	if err != nil {
		return nil, err
	}

	data, err := base64.StdEncoding.DecodeString(ret.Get("data").String())
	if err != nil {
		return nil, err
	}

	tag, err := base64.StdEncoding.DecodeString(ret.Get("authtag").String())
	if err != nil {
		return nil, err
	}

	aad := fmt.Sprintf("%s|%s|%s|%s", mp.url(path, nil), mp.appid, header.Get(HeaderMPTimestamp), mp.sfMode.aesSN)

	return xcrypto.AESDecryptGCM(key, iv, append(data, tag...), []byte(aad), nil)
}

// Code2Session 通过临时登录凭证code完成登录流程
func (mp *MiniProgram) Code2Session(ctx context.Context, code string) (gjson.Result, error) {
	query := url.Values{}

	query.Set("appid", mp.appid)
	query.Set("secret", mp.secret)
	query.Set("js_code", code)
	query.Set("grant_type", "authorization_code")

	b, err := mp.do(ctx, http.MethodGet, "/sns/jscode2session", query, nil)
	if err != nil {
		return lib.Fail(err)
	}

	ret := gjson.ParseBytes(b)
	if code := ret.Get("errcode").Int(); code != 0 {
		return lib.Fail(fmt.Errorf("%d | %s", code, ret.Get("errmsg").String()))
	}
	return ret, nil
}

// loadAccessToken 获取稳定版接口调用凭据，有两种调用模式:
// 1. 普通模式，access_token有效期内重复调用该接口不会更新access_token，绝大部分场景下使用该模式；
// 2. 强制刷新模式，会导致上次获取的access_token失效，并返回新的access_token
func (mp *MiniProgram) reloadAccessToken() error {
	params := lib.X{
		"grant_type":    "client_credential",
		"appid":         mp.appid,
		"secret":        mp.secret,
		"force_refresh": false,
	}

	b, err := mp.do(context.Background(), http.MethodPost, "/cgi-bin/stable_token", nil, params, xhttp.WithHeader(xhttp.HeaderContentType, xhttp.ContentJSON))
	if err != nil {
		return err
	}

	ret := gjson.ParseBytes(b)
	if code := ret.Get("errcode").Int(); code != 0 {
		return fmt.Errorf("%d | %s", code, ret.Get("errmsg").String())
	}
	mp.token.Store(ret.Get("access_token").String())
	return nil
}

func (mp *MiniProgram) getAccessToken() (string, error) {
	v := mp.token.Load()
	if v == nil {
		return "", errors.New("access_token is empty (forgotten auto load?)")
	}
	token, ok := v.(string)
	if !ok {
		return "", errors.New("access_token is not a string")
	}
	return token, nil
}

// AutoLoadAccessToken 自动加载AccessToken
func (mp *MiniProgram) AutoLoadAccessToken() error {
	if err := mp.reloadAccessToken(); err != nil {
		return err
	}
	go func() {
		ticker := time.NewTicker(time.Minute * 5)
		defer ticker.Stop()
		for range ticker.C {
			mp.reloadAccessToken()
		}
	}()
	return nil
}

// GetJSON GET请求JSON数据
func (mp *MiniProgram) GetJSON(ctx context.Context, path string, query url.Values) (gjson.Result, error) {
	accessToken, err := mp.getAccessToken()
	if err != nil {
		return lib.Fail(err)
	}
	if query == nil {
		query = url.Values{}
	}
	query.Set(AccessToken, accessToken)

	b, err := mp.do(ctx, http.MethodGet, path, query, nil)
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
func (mp *MiniProgram) GetBuffer(ctx context.Context, path string, query url.Values) ([]byte, error) {
	accessToken, err := mp.getAccessToken()
	if err != nil {
		return nil, err
	}
	if query == nil {
		query = url.Values{}
	}
	query.Set(AccessToken, accessToken)

	b, err := mp.do(ctx, http.MethodGet, path, query, nil)
	if err != nil {
		return nil, err
	}

	ret := gjson.ParseBytes(b)
	if code := ret.Get("errcode").Int(); code != 0 {
		return nil, fmt.Errorf("%d | %s", code, ret.Get("errmsg").String())
	}
	return b, nil
}

// PostJSON POST请求JSON数据
func (mp *MiniProgram) PostJSON(ctx context.Context, path string, params lib.X) (gjson.Result, error) {
	accessToken, err := mp.getAccessToken()
	if err != nil {
		return lib.Fail(err)
	}
	query := url.Values{}
	query.Set(AccessToken, accessToken)

	b, err := mp.do(ctx, http.MethodPost, path, query, params, xhttp.WithHeader(xhttp.HeaderContentType, xhttp.ContentJSON))
	if err != nil {
		return lib.Fail(err)
	}

	ret := gjson.ParseBytes(b)
	if code := ret.Get("errcode").Int(); code != 0 {
		return lib.Fail(fmt.Errorf("%d | %s", code, ret.Get("errmsg").String()))
	}
	return ret, nil
}

// PostBuffer POST请求获取buffer (如：获取二维码)
func (mp *MiniProgram) PostBuffer(ctx context.Context, path string, params lib.X) ([]byte, error) {
	accessToken, err := mp.getAccessToken()
	if err != nil {
		return nil, err
	}
	query := url.Values{}
	query.Set(AccessToken, accessToken)

	b, err := mp.do(ctx, http.MethodPost, path, query, params, xhttp.WithHeader(xhttp.HeaderContentType, xhttp.ContentJSON))
	if err != nil {
		return nil, err
	}

	ret := gjson.ParseBytes(b)
	if code := ret.Get("errcode").Int(); code != 0 {
		return nil, fmt.Errorf("%d | %s", code, ret.Get("errmsg").String())
	}
	return b, nil
}

// SafePostJSON POST请求JSON数据
// 安全鉴权模式 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/getting_started/api_signature.html
// 支持的api可参考 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc
func (mp *MiniProgram) SafePostJSON(ctx context.Context, path string, params lib.X) (gjson.Result, error) {
	accessToken, err := mp.getAccessToken()
	if err != nil {
		return lib.Fail(err)
	}
	query := url.Values{}
	query.Set(AccessToken, accessToken)

	b, err := mp.doSafe(ctx, http.MethodPost, path, query, params)
	if err != nil {
		return lib.Fail(err)
	}

	ret := gjson.ParseBytes(b)
	if code := ret.Get("errcode").Int(); code != 0 {
		return lib.Fail(fmt.Errorf("%d | %s", code, ret.Get("errmsg").String()))
	}
	return ret, nil
}

// SafePostBuffer POST请求获取buffer (如：获取二维码)
// 安全鉴权模式 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/getting_started/api_signature.html
// 支持的api可参考 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc
func (mp *MiniProgram) SafePostBuffer(ctx context.Context, path string, params lib.X) ([]byte, error) {
	accessToken, err := mp.getAccessToken()
	if err != nil {
		return nil, err
	}
	query := url.Values{}
	query.Set(AccessToken, accessToken)

	b, err := mp.doSafe(ctx, http.MethodPost, path, query, params)
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
func (mp *MiniProgram) Upload(ctx context.Context, path string, form xhttp.UploadForm) (gjson.Result, error) {
	accessToken, err := mp.getAccessToken()
	if err != nil {
		return lib.Fail(err)
	}
	query := url.Values{}
	query.Set(AccessToken, accessToken)

	reqURL := mp.url(path, query)

	log := lib.NewReqLog(http.MethodPost, reqURL)
	defer log.Do(ctx, mp.logger)

	resp, err := mp.httpCli.Upload(ctx, reqURL, form)
	if err != nil {
		log.Set("error", err.Error())
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
		log.Set("error", err.Error())
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
func (mp *MiniProgram) VerifyURL(signature, timestamp, nonce string) error {
	if SignWithSHA1(mp.srvCfg.token, timestamp, nonce) != signature {
		return errors.New("signature verified fail")
	}
	return nil
}

// DecodeEncryptData 解析加密数据，如：授权的用户信息和手机号
// [参考](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html)
func (mp *MiniProgram) DecodeEncryptData(sessionKey, iv, encryptData string) ([]byte, error) {
	keyBlock, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return nil, fmt.Errorf("session_key base64.decode error: %w", err)
	}

	ivBlock, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, fmt.Errorf("iv base64.decode error: %w", err)
	}

	data, err := base64.StdEncoding.DecodeString(encryptData)
	if err != nil {
		return nil, fmt.Errorf("encrypt_data base64.decode error: %w", err)
	}

	ct, err := xcrypto.AESEncryptCBC(keyBlock, ivBlock, data)
	if err != nil {
		return nil, err
	}

	return ct.Bytes(), nil
}

// DecodeEventMsg 解析事件消息，使用：msg_signature、timestamp、nonce、msg_encrypt
// [参考](https://developers.weixin.qq.com/miniprogram/dev/framework/server-ability/message-push.html)
func (mp *MiniProgram) DecodeEventMsg(signature, timestamp, nonce, encryptMsg string) (value.V, error) {
	if SignWithSHA1(mp.srvCfg.token, timestamp, nonce, encryptMsg) != signature {
		return nil, errors.New("signature verified fail")
	}

	b, err := EventDecrypt(mp.appid, mp.srvCfg.aeskey, encryptMsg)
	if err != nil {
		return nil, err
	}
	return XMLToValue(b)
}

// ReplyEventMsg 事件消息回复
func (mp *MiniProgram) ReplyEventMsg(msg value.V) (value.V, error) {
	return EventReply(mp.appid, mp.srvCfg.token, mp.srvCfg.aeskey, msg)
}

// MPOption 小程序设置项
type MPOption func(mp *MiniProgram)

// WithMPSrvCfg 设置小程序服务器配置
// [参考](https://developers.weixin.qq.com/miniprogram/dev/framework/server-ability/message-push.html)
func WithMPSrvCfg(token, aeskey string) MPOption {
	return func(mp *MiniProgram) {
		mp.srvCfg.token = token
		mp.srvCfg.aeskey = aeskey
	}
}

// WithMPHttpCli 设置小程序请求的 HTTP Client
func WithMPHttpCli(c *http.Client) MPOption {
	return func(mp *MiniProgram) {
		mp.httpCli = xhttp.NewHTTPClient(c)
	}
}

// WithMPLogger 设置小程序日志记录
func WithMPLogger(fn func(ctx context.Context, data map[string]string)) MPOption {
	return func(mp *MiniProgram) {
		mp.logger = fn
	}
}

// WithMPAesKey 设置小程序 AES-GCM 加密Key
func WithMPAesKey(serialNO, key string) MPOption {
	return func(mp *MiniProgram) {
		mp.sfMode.aesSN = serialNO
		mp.sfMode.aeskey = key
	}
}

// WithMPPrivateKey 设置小程序RSA私钥
func WithMPPrivateKey(key *xcrypto.PrivateKey) MPOption {
	return func(mp *MiniProgram) {
		mp.sfMode.prvKey = key
	}
}

// WithMPPublicKey 设置小程序平台RSA公钥
func WithMPPublicKey(serialNO string, key *xcrypto.PublicKey) MPOption {
	return func(mp *MiniProgram) {
		mp.sfMode.pubSN = serialNO
		mp.sfMode.pubKey = key
	}
}

// NewMiniProgram 生成一个小程序实例
func NewMiniProgram(appid, secret string, options ...MPOption) *MiniProgram {
	mp := &MiniProgram{
		host:    "https://api.weixin.qq.com",
		appid:   appid,
		secret:  secret,
		srvCfg:  new(ServerConfig),
		httpCli: xhttp.NewDefaultClient(),
	}
	for _, fn := range options {
		fn(mp)
	}
	return mp
}
