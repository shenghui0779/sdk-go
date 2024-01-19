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
	lib_http "github.com/shenghui0779/sdk-go/lib/http"
	"github.com/shenghui0779/sdk-go/lib/value"
)

// Corp 企业微信(企业内部开发)
type Corp struct {
	host    string
	corpid  string
	secret  string
	srvCfg  *ServerConfig
	httpCli lib_http.Client
	logger  func(ctx context.Context, data map[string]string)
}

// AppID 返回AppID
func (c *Corp) CorpID() string {
	return c.corpid
}

// Secret 返回Secret
func (c *Corp) Secret() string {
	return c.secret
}

func (c *Corp) url(path string, query url.Values) string {
	var builder strings.Builder

	builder.WriteString(c.host)
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

func (c *Corp) do(ctx context.Context, method, path string, query url.Values, params lib.X, options ...lib_http.Option) ([]byte, error) {
	reqURL := c.url(path, query)

	log := lib.NewReqLog(method, reqURL)
	defer log.Do(ctx, c.logger)

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

	resp, err := c.httpCli.Do(ctx, method, reqURL, body, options...)
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

// OAuthURL 生成网页授权URL
// [参考](https://developer.work.weixin.qq.com/document/path/91022)
func (c *Corp) OAuthURL(scope AuthScope, redirectURI, state, agentID string) string {
	query := url.Values{}

	query.Set("appid", c.corpid)
	query.Set("redirect_uri", redirectURI)
	query.Set("response_type", "code")
	query.Set("scope", string(scope))
	query.Set("state", state)
	query.Set("agentid", agentID)

	return fmt.Sprintf("https://open.weixin.qq.com/connect/cuth2/authorize?%s#wechat_redirect", query.Encode())
}

// AccessToken 获取接口调用凭据 (开发者应在 WithAccessToken 回调函数中使用该方法，并自行实现存/取)
func (c *Corp) AccessToken(ctx context.Context) (gjson.Result, error) {
	query := url.Values{}

	query.Set("corpid", c.corpid)
	query.Set("corpsecret", c.secret)

	b, err := c.do(ctx, http.MethodGet, "/cgi-bin/gettoken", query, nil)
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
func (c *Corp) GetJSON(ctx context.Context, accessToken, path string, query url.Values) (gjson.Result, error) {
	if query == nil {
		query = url.Values{}
	}
	query.Set(AccessToken, accessToken)

	b, err := c.do(ctx, http.MethodGet, path, query, nil)
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
func (c *Corp) PostJSON(ctx context.Context, accessToken, path string, params lib.X) (gjson.Result, error) {
	query := url.Values{}
	query.Set(AccessToken, accessToken)

	b, err := c.do(ctx, http.MethodPost, path, query, params, lib_http.WithHeader(lib_http.HeaderContentType, lib_http.ContentJSON))
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
func (c *Corp) GetBuffer(ctx context.Context, accessToken, path string, query url.Values) ([]byte, error) {
	if query == nil {
		query = url.Values{}
	}
	query.Set(AccessToken, accessToken)

	b, err := c.do(ctx, http.MethodGet, path, query, nil)
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
func (c *Corp) PostBuffer(ctx context.Context, accessToken, path string, params lib.X) ([]byte, error) {
	query := url.Values{}
	query.Set(AccessToken, accessToken)

	b, err := c.do(ctx, http.MethodPost, path, query, params, lib_http.WithHeader(lib_http.HeaderContentType, lib_http.ContentJSON))
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
func (c *Corp) Upload(ctx context.Context, accessToken, path string, form lib_http.UploadForm) (gjson.Result, error) {
	query := url.Values{}
	query.Set(AccessToken, accessToken)

	reqURL := c.url(path, query)

	log := lib.NewReqLog(http.MethodPost, reqURL)
	defer log.Do(ctx, c.logger)

	resp, err := c.httpCli.Upload(ctx, reqURL, form)
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

// VerifyURL 服务器URL验证，使用：msg_signature、timestamp、nonce、echostr（若验证成功，解密echostr后返回msg字段内容）
// [参考](https://developer.work.weixin.qq.com/document/path/90930)
func (c *Corp) VerifyURL(signature, timestamp, nonce, echoStr string) (string, error) {
	if SignWithSHA1(c.srvCfg.token, timestamp, nonce, echoStr) != signature {
		return "", errors.New("signature verified fail")
	}

	b, err := EventDecrypt(c.corpid, c.srvCfg.aeskey, echoStr)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// DecodeEventMsg 解析事件消息，使用：msg_signature、timestamp、nonce、msg_encrypt
// [参考](https://developer.work.weixin.qq.com/document/path/90930)
func (c *Corp) DecodeEventMsg(signature, timestamp, nonce, encryptMsg string) (value.V, error) {
	if SignWithSHA1(c.srvCfg.token, timestamp, nonce, encryptMsg) != signature {
		return nil, errors.New("signature verified fail")
	}

	b, err := EventDecrypt(c.corpid, c.srvCfg.aeskey, encryptMsg)
	if err != nil {
		return nil, err
	}

	return ParseXMLToV(b)
}

// ReplyEventMsg 事件消息回复
func (c *Corp) ReplyEventMsg(msg value.V) (value.V, error) {
	return EventReply(c.corpid, c.srvCfg.token, c.srvCfg.aeskey, msg)
}

// CorpOption 企业微信设置项
type CorpOption func(c *Corp)

// WithCorpSrvCfg 设置企业微信服务器配置
// [参考](https://developer.work.weixin.qq.com/document/path/90968)
func WithCorpSrvCfg(token, aeskey string) CorpOption {
	return func(c *Corp) {
		c.srvCfg.token = token
		c.srvCfg.aeskey = aeskey
	}
}

// WithCorpHttpCli 设置企业微信请求的 HTTP Client
func WithCorpHttpCli(cli *http.Client) CorpOption {
	return func(c *Corp) {
		c.httpCli = lib_http.NewHTTPClient(cli)
	}
}

// WithCorpLogger 设置企业微信日志记录
func WithCorpLogger(f func(ctx context.Context, data map[string]string)) CorpOption {
	return func(c *Corp) {
		c.logger = f
	}
}

// NewCorp 生成一个企业微信(企业内部开发)实例
func NewCorp(corpid, secret string, options ...CorpOption) *Corp {
	c := &Corp{
		host:    "https://qyapi.weixin.qq.com",
		corpid:  corpid,
		secret:  secret,
		srvCfg:  new(ServerConfig),
		httpCli: lib_http.NewDefaultClient(),
	}

	for _, f := range options {
		f(c)
	}

	return c
}
