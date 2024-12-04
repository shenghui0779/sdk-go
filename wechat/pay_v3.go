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
	"sync/atomic"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"

	"github.com/shenghui0779/sdk-go/lib"
	"github.com/shenghui0779/sdk-go/lib/value"
	"github.com/shenghui0779/sdk-go/lib/xcrypto"
)

// PayV3 微信支付V3
type PayV3 struct {
	host   string
	mchid  string
	apikey string
	prvSN  string
	prvKey *xcrypto.PrivateKey
	pubKey atomic.Value // map[string]*xcrypto.PublicKey
	client *resty.Client
	logger func(ctx context.Context, err error, data map[string]string)
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

func (p *PayV3) publicKey(serialNO string) (*xcrypto.PublicKey, error) {
	v := p.pubKey.Load()
	if v == nil {
		return nil, errors.New("public key is empty (forgotten auto load?)")
	}
	keyMap, ok := v.(map[string]*xcrypto.PublicKey)
	if !ok {
		return nil, errors.New("public key is not map[string]*PublicKey")
	}
	pk, ok := keyMap[serialNO]
	if !ok {
		return nil, fmt.Errorf("cert(serial_no=%s) not found", serialNO)
	}
	return pk, nil
}

func (p *PayV3) reloadCerts() error {
	ctx := context.Background()

	reqURL := p.url("/v3/certificates", nil)

	log := lib.NewReqLog(http.MethodGet, reqURL)
	defer log.Do(ctx, p.logger)

	authStr, err := p.Authorization(http.MethodGet, "/v3/certificates", nil, "")
	if err != nil {
		log.SetError(err)
		return err
	}
	log.Set(lib.HeaderAuthorization, authStr)

	resp, err := p.client.R().
		SetContext(ctx).
		SetHeader(lib.HeaderAccept, lib.ContentJSON).
		SetHeader(lib.HeaderAuthorization, authStr).
		Get(reqURL)
	if err != nil {
		log.SetError(err)
		return err
	}
	log.SetRespHeader(resp.Header())
	log.SetStatusCode(resp.StatusCode())
	log.SetRespBody(string(resp.Body()))

	if resp.StatusCode() >= 400 {
		text := string(resp.Body())
		log.SetError(errors.New(text))
		return errors.New(text)
	}

	keyMap := make(map[string]*xcrypto.PublicKey)
	headSerial := resp.Header().Get(HeaderPaySerial)

	ret := gjson.GetBytes(resp.Body(), "data")
	for _, v := range ret.Array() {
		serialNO := v.Get("serial_no").String()
		cert := v.Get("encrypt_certificate")

		nonce := cert.Get("nonce").String()
		data := cert.Get("ciphertext").String()
		aad := cert.Get("associated_data").String()

		block, err := xcrypto.AESDecryptGCM([]byte(p.apikey), []byte(nonce), []byte(data), []byte(aad), nil)
		if err != nil {
			log.SetError(err)
			return err
		}
		key, err := xcrypto.NewPublicKeyFromDerBlock(block)
		if err != nil {
			log.SetError(err)
			return err
		}
		keyMap[serialNO] = key

		// 签名验证
		if serialNO == headSerial {
			// 签名验证
			var builder strings.Builder

			builder.WriteString(resp.Header().Get(HeaderPayTimestamp))
			builder.WriteString("\n")
			builder.WriteString(resp.Header().Get(HeaderPayNonce))
			builder.WriteString("\n")
			builder.Write(resp.Body())
			builder.WriteString("\n")

			if err = key.Verify(crypto.SHA256, []byte(builder.String()), []byte(resp.Header().Get(HeaderPaySignature))); err != nil {
				log.SetError(err)
				return err
			}
		}
	}

	p.pubKey.Store(keyMap)
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
			log.SetError(err)
			return nil, err
		}
		log.SetReqBody(string(body))
	}

	authStr, err := p.Authorization(method, path, query, string(body))
	if err != nil {
		log.SetError(err)
		return nil, err
	}
	log.Set(lib.HeaderAuthorization, authStr)

	resp, err := p.client.R().
		SetContext(ctx).
		SetHeader(lib.HeaderAccept, lib.ContentJSON).
		SetHeader(lib.HeaderAuthorization, authStr).
		SetHeader(lib.HeaderContentType, lib.ContentJSON).
		SetBody(body).
		Execute(method, reqURL)
	if err != nil {
		log.SetError(err)
		return nil, err
	}
	log.SetRespHeader(resp.Header())
	log.SetStatusCode(resp.StatusCode())
	log.SetRespBody(string(resp.Body()))
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", resp.StatusCode())
	}

	// 签名校验
	if err = p.Verify(ctx, resp.Header(), resp.Body()); err != nil {
		log.SetError(err)
		return nil, err
	}

	ret := &APIResult{
		Code: resp.StatusCode(),
		Body: gjson.ParseBytes(resp.Body()),
	}
	return ret, nil
}

// AutoLoadCerts 自动加载平台证书
func (p *PayV3) AutoLoadCerts() error {
	if err := p.reloadCerts(); err != nil {
		return err
	}
	go func() {
		ticker := time.NewTicker(time.Hour * 24)
		defer ticker.Stop()
		for range ticker.C {
			p.reloadCerts()
		}
	}()
	return nil
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
func (p *PayV3) Upload(ctx context.Context, reqPath, fieldName, filePath, metadata string, query url.Values) (*APIResult, error) {
	reqURL := p.url(reqPath, nil)

	log := lib.NewReqLog(http.MethodPost, reqURL)
	defer log.Do(ctx, p.logger)

	authStr, err := p.Authorization(http.MethodPost, reqPath, query, metadata)
	if err != nil {
		log.SetError(err)
		return nil, err
	}
	log.Set(lib.HeaderAuthorization, authStr)

	resp, err := p.client.R().
		SetContext(ctx).
		SetHeader(lib.HeaderAuthorization, authStr).
		SetFile(fieldName, filePath).
		SetMultipartField("meta", "", lib.ContentJSON, strings.NewReader(metadata)).
		Post(reqURL)
	if err != nil {
		log.SetError(err)
		return nil, err
	}
	log.SetRespHeader(resp.Header())
	log.SetStatusCode(resp.StatusCode())
	log.SetRespBody(string(resp.Body()))

	// 签名校验
	if err = p.Verify(ctx, resp.Header(), resp.Body()); err != nil {
		log.SetError(err)
		return nil, err
	}

	ret := &APIResult{
		Code: resp.StatusCode(),
		Body: gjson.ParseBytes(resp.Body()),
	}
	return ret, nil
}

// UploadWithReader 上传资源
func (p *PayV3) UploadWithReader(ctx context.Context, reqPath, fieldName, fileName string, reader io.Reader, metadata string, query url.Values) (*APIResult, error) {
	reqURL := p.url(reqPath, nil)

	log := lib.NewReqLog(http.MethodPost, reqURL)
	defer log.Do(ctx, p.logger)

	authStr, err := p.Authorization(http.MethodPost, reqPath, query, metadata)
	if err != nil {
		log.SetError(err)
		return nil, err
	}
	log.Set(lib.HeaderAuthorization, authStr)

	resp, err := p.client.R().
		SetContext(ctx).
		SetHeader(lib.HeaderAuthorization, authStr).
		SetMultipartField(fieldName, fileName, "", reader).
		SetMultipartField("meta", "", lib.ContentJSON, strings.NewReader(metadata)).
		Post(reqURL)
	if err != nil {
		log.SetError(err)
		return nil, err
	}
	log.SetRespHeader(resp.Header())
	log.SetStatusCode(resp.StatusCode())
	log.SetRespBody(string(resp.Body()))

	// 签名校验
	if err = p.Verify(ctx, resp.Header(), resp.Body()); err != nil {
		log.SetError(err)
		return nil, err
	}

	ret := &APIResult{
		Code: resp.StatusCode(),
		Body: gjson.ParseBytes(resp.Body()),
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
		log.SetError(err)
		return err
	}
	log.Set(lib.HeaderAuthorization, authStr)

	resp, err := p.client.R().
		SetContext(ctx).
		SetHeader(lib.HeaderAuthorization, authStr).
		SetDoNotParseResponse(true).
		Get(downloadURL)
	if err != nil {
		log.SetError(err)
		return err
	}
	log.SetRespHeader(resp.Header())
	log.SetStatusCode(resp.StatusCode())

	_, err = io.Copy(w, resp.RawResponse.Body)
	return err
}

// Authorization 生成签名并返回 HTTP Authorization
func (p *PayV3) Authorization(method, path string, query url.Values, body string) (string, error) {
	if p.prvKey == nil {
		return "", errors.New("private key not found (forgotten configure?)")
	}

	nonce := lib.Nonce(32)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	var builder strings.Builder

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

	key, err := p.publicKey(serial)
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

// WithPayV3Client 设置支付(v3)请求的 HTTP Client
func WithPayV3Client(cli *http.Client) PayV3Option {
	return func(p *PayV3) {
		p.client = resty.NewWithClient(cli)
	}
}

// WithPayV3PrivateKey 设置支付(v3)商户RSA私钥
func WithPayV3PrivateKey(serialNO string, key *xcrypto.PrivateKey) PayV3Option {
	return func(p *PayV3) {
		p.prvSN = serialNO
		p.prvKey = key
	}
}

// WithPayV3Logger 设置支付(v3)日志记录
func WithPayV3Logger(fn func(ctx context.Context, err error, data map[string]string)) PayV3Option {
	return func(p *PayV3) {
		p.logger = fn
	}
}

// NewPayV3 生成一个微信支付(v3)实例
func NewPayV3(mchid, apikey string, options ...PayV3Option) *PayV3 {
	pay := &PayV3{
		host:   "https://api.mch.weixin.qq.com",
		mchid:  mchid,
		apikey: apikey,
		client: lib.NewClient(),
	}
	for _, f := range options {
		f(pay)
	}
	return pay
}
