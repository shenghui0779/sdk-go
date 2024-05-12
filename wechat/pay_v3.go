package wechat

import (
	"context"
	"crypto"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/tidwall/gjson"

	"github.com/shenghui0779/sdk-go/lib"
	lib_crypto "github.com/shenghui0779/sdk-go/lib/crypto"
	"github.com/shenghui0779/sdk-go/lib/curl"
	"github.com/shenghui0779/sdk-go/lib/value"
)

// PayV3 微信支付V3
type PayV3 struct {
	host    string
	mchid   string
	apikey  string
	prvSN   string
	prvKey  *lib_crypto.PrivateKey
	pubKey  map[string]*lib_crypto.PublicKey
	mutex   sync.Mutex
	httpCli curl.Client
	logger  func(ctx context.Context, data map[string]string)
}

// MchID 返回mchid
func (p *PayV3) MchID() string {
	return p.mchid
}

// ApiKey 返回apikey
func (p *PayV3) ApiKey() string {
	return p.apikey
}

func (p *PayV3) url(path string, query url.Values) string {
	var builder strings.Builder

	builder.WriteString(p.host)
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

func (p *PayV3) publicKey(ctx context.Context, serialNO string) (*lib_crypto.PublicKey, error) {
	key, ok := p.pubKey[serialNO]
	if ok {
		return key, nil
	}

	p.mutex.Lock()
	defer p.mutex.Unlock()

	// 二次确认
	key, ok = p.pubKey[serialNO]
	if ok {
		return key, nil
	}

	if err := p.httpCerts(ctx); err != nil {
		return nil, err
	}

	pk, ok := p.pubKey[serialNO]
	if !ok {
		return nil, fmt.Errorf("cert(serial_no=%s) not found", serialNO)
	}
	return pk, nil
}

func (p *PayV3) httpCerts(ctx context.Context) error {
	reqURL := p.url("/v3/certificates", nil)

	log := lib.NewReqLog(http.MethodGet, reqURL)
	defer log.Do(ctx, p.logger)

	authStr, err := p.Authorization(http.MethodGet, "/v3/certificates", nil, "")
	if err != nil {
		return err
	}

	log.Set(curl.HeaderAuthorization, authStr)

	resp, err := p.httpCli.Do(ctx, http.MethodGet, reqURL, nil, curl.WithHeader(curl.HeaderAccept, "application/json"), curl.WithHeader(curl.HeaderAuthorization, authStr))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.SetRespHeader(resp.Header)
	log.SetStatusCode(resp.StatusCode)

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.SetRespBody(string(b))

	if resp.StatusCode >= 400 {
		return errors.New(string(b))
	}

	serial := resp.Header.Get(HeaderPaySerial)

	ret := gjson.GetBytes(b, "data")
	for _, v := range ret.Array() {
		cert := v.Get("encrypt_certificate")

		nonce := cert.Get("nonce").String()
		data := cert.Get("ciphertext").String()
		aad := cert.Get("associated_data").String()

		block, err := lib_crypto.AESDecryptGCM([]byte(p.apikey), []byte(nonce), []byte(data), []byte(aad), nil)
		if err != nil {
			return err
		}
		key, err := lib_crypto.NewPublicKeyFromDerBlock(block)
		if err != nil {
			return err
		}
		p.pubKey[cert.Get("serial_no").String()] = key

		// 签名验证
		if v.Get("serial_no").String() == serial {
			// 签名验证
			var builder strings.Builder

			builder.WriteString(resp.Header.Get(HeaderPayTimestamp))
			builder.WriteString("\n")
			builder.WriteString(resp.Header.Get(HeaderPayNonce))
			builder.WriteString("\n")
			builder.Write(b)
			builder.WriteString("\n")

			if err = key.Verify(crypto.SHA256, []byte(builder.String()), []byte(resp.Header.Get(HeaderPaySignature))); err != nil {
				return err
			}
		}
	}

	return nil
}

func (p *PayV3) do(ctx context.Context, method, path string, query url.Values, params lib.X) (*APIResult, error) {
	reqURL := p.url(path, query)

	log := lib.NewReqLog(method, reqURL)
	defer log.Do(ctx, p.logger)

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

	authStr, err := p.Authorization(method, path, query, string(body))
	if err != nil {
		return nil, err
	}

	log.Set(curl.HeaderAuthorization, authStr)

	resp, err := p.httpCli.Do(ctx, method, reqURL, body,
		curl.WithHeader(curl.HeaderAccept, "application/json"),
		curl.WithHeader(curl.HeaderAuthorization, authStr),
		curl.WithHeader(curl.HeaderContentType, curl.ContentJSON),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	log.SetRespHeader(resp.Header)
	log.SetStatusCode(resp.StatusCode)

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.SetRespBody(string(b))

	// 签名校验
	if err = p.Verify(ctx, resp.Header, b); err != nil {
		return nil, err
	}

	ret := &APIResult{
		Code: resp.StatusCode,
		Body: gjson.ParseBytes(b),
	}

	return ret, nil
}

// GetJSON GET请求JSON数据
func (p *PayV3) GetJSON(ctx context.Context, path string, query url.Values) (*APIResult, error) {
	return p.do(ctx, http.MethodGet, path, query, nil)
}

// PostJSON POST请求JSON数据
func (p *PayV3) PostJSON(ctx context.Context, path string, params lib.X) (*APIResult, error) {
	return p.do(ctx, http.MethodPost, path, nil, params)
}

// Upload 上传资源
func (p *PayV3) Upload(ctx context.Context, path string, form curl.UploadForm) (*APIResult, error) {
	reqURL := p.url(path, nil)

	log := lib.NewReqLog(http.MethodPost, reqURL)
	defer log.Do(ctx, p.logger)

	authStr, err := p.Authorization(http.MethodPost, path, nil, form.Field("meta"))
	if err != nil {
		return nil, err
	}

	log.Set(curl.HeaderAuthorization, authStr)

	resp, err := p.httpCli.Upload(ctx, reqURL, form, curl.WithHeader(curl.HeaderAuthorization, authStr))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	log.SetRespHeader(resp.Header)
	log.SetStatusCode(resp.StatusCode)

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.SetRespBody(string(b))

	// 签名校验
	if err = p.Verify(ctx, resp.Header, b); err != nil {
		return nil, err
	}

	ret := &APIResult{
		Code: resp.StatusCode,
		Body: gjson.ParseBytes(b),
	}

	return ret, nil
}

// Download 下载资源 (需先获取download_url)
func (p *PayV3) Download(ctx context.Context, downloadURL string, w io.Writer) error {
	log := lib.NewReqLog(http.MethodGet, downloadURL)
	defer log.Do(ctx, p.logger)

	// 获取 download_url
	authStr, err := p.Authorization(http.MethodGet, downloadURL, nil, "")
	if err != nil {
		return err
	}

	log.Set(curl.HeaderAuthorization, authStr)

	resp, err := p.httpCli.Do(ctx, http.MethodGet, downloadURL, nil, curl.WithHeader(curl.HeaderAuthorization, authStr))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	log.SetRespHeader(resp.Header)
	log.SetStatusCode(resp.StatusCode)

	_, err = io.Copy(w, resp.Body)

	return err
}

// Authorization 生成签名并返回 HTTP Authorization
func (p *PayV3) Authorization(method, path string, query url.Values, body string) (string, error) {
	if p.prvKey == nil {
		return "", errors.New("private key not found (forgotten configure?)")
	}

	var builder strings.Builder

	nonce := lib.Nonce(32)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	builder.WriteString(method)
	builder.WriteString("\n")
	builder.WriteString(path)
	if len(query) != 0 {
		builder.WriteString("?")
		builder.WriteString(query.Encode())
	}
	builder.WriteString("\n")
	builder.WriteString(timestamp)
	builder.WriteString("\n")
	builder.WriteString(nonce)
	builder.WriteString("\n")
	if len(body) != 0 {
		builder.WriteString(body)
	}
	builder.WriteString("\n")

	sign, err := p.prvKey.Sign(crypto.SHA256, []byte(builder.String()))
	if err != nil {
		return "", err
	}

	auth := fmt.Sprintf(`WECHATPAY2-SHA256-RSA2048 mchid="%s",nonce_str="%s",signature="%s",timestamp="%s",serial_no="%s"`, p.mchid, nonce, base64.StdEncoding.EncodeToString(sign), timestamp, p.prvSN)

	return auth, nil
}

// Verify 验证微信签名
func (p *PayV3) Verify(ctx context.Context, header http.Header, body []byte) error {
	nonce := header.Get(HeaderPayNonce)
	timestamp := header.Get(HeaderPayTimestamp)
	serial := header.Get(HeaderPaySerial)
	sign := header.Get(HeaderPaySignature)

	key, err := p.publicKey(ctx, serial)
	if err != nil {
		return err
	}

	var builder strings.Builder

	builder.WriteString(timestamp)
	builder.WriteString("\n")
	builder.WriteString(nonce)
	builder.WriteString("\n")
	if len(body) != 0 {
		builder.Write(body)
	}
	builder.WriteString("\n")

	return key.Verify(crypto.SHA256, []byte(builder.String()), []byte(sign))
}

// APPAPI 用于APP拉起支付
func (p *PayV3) APPAPI(appid, prepayID string) (value.V, error) {
	nonce := lib.Nonce(32)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	v := value.V{}

	v.Set("appid", appid)
	v.Set("partnerid", p.mchid)
	v.Set("prepayid", prepayID)
	v.Set("package", "Sign=WXPay")
	v.Set("noncestr", nonce)
	v.Set("timestamp", timestamp)

	var builder strings.Builder

	builder.WriteString(appid)
	builder.WriteString("\n")
	builder.WriteString(timestamp)
	builder.WriteString("\n")
	builder.WriteString(nonce)
	builder.WriteString("\n")
	builder.WriteString(prepayID)
	builder.WriteString("\n")

	sign, err := p.prvKey.Sign(crypto.SHA256, []byte(builder.String()))
	if err != nil {
		return nil, err
	}

	v.Set("sign", string(sign))

	return v, nil
}

// JSAPI 用于JS拉起支付
func (p *PayV3) JSAPI(appid, prepayID string) (value.V, error) {
	nonce := lib.Nonce(32)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	v := value.V{}

	v.Set("appId", appid)
	v.Set("nonceStr", nonce)
	v.Set("package", "prepay_id="+prepayID)
	v.Set("signType", "RSA")
	v.Set("timeStamp", timestamp)

	var builder strings.Builder

	builder.WriteString(appid)
	builder.WriteString("\n")
	builder.WriteString(timestamp)
	builder.WriteString("\n")
	builder.WriteString(nonce)
	builder.WriteString("\n")
	builder.WriteString("prepay_id=" + prepayID)
	builder.WriteString("\n")

	sign, err := p.prvKey.Sign(crypto.SHA256, []byte(builder.String()))
	if err != nil {
		return nil, err
	}

	v.Set("sign", string(sign))

	return v, nil
}

// PayV3Option 微信支付(v3)设置项
type PayV3Option func(p *PayV3)

// WithPayV3HttpCli 设置支付(v3)请求的 HTTP Client
func WithPayV3HttpCli(c *http.Client) PayV3Option {
	return func(p *PayV3) {
		p.httpCli = curl.NewHTTPClient(c)
	}
}

// WithPayV3PrivateKey 设置支付(v3)商户RSA私钥
func WithPayV3PrivateKey(serialNO string, key *lib_crypto.PrivateKey) PayV3Option {
	return func(p *PayV3) {
		p.prvSN = serialNO
		p.prvKey = key
	}
}

// WithPayV3Logger 设置支付(v3)日志记录
func WithPayV3Logger(fn func(ctx context.Context, data map[string]string)) PayV3Option {
	return func(p *PayV3) {
		p.logger = fn
	}
}

// NewPayV3 生成一个微信支付(v3)实例
func NewPayV3(mchid, apikey string, options ...PayV3Option) *PayV3 {
	pay := &PayV3{
		host:    "https://api.mch.weixin.qq.com",
		mchid:   mchid,
		apikey:  apikey,
		httpCli: curl.NewDefaultClient(),
	}

	for _, fn := range options {
		fn(pay)
	}

	return pay
}
