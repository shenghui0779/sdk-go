package antchain

import (
	"context"
	"crypto"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/tidwall/gjson"

	"github.com/yiigo/sdk-go/internal"
	"github.com/yiigo/sdk-go/internal/xcrypto"
)

// Config 客户端配置
type Config struct {
	BizID      string              `json:"biz_id"`      // 链ID (a00e36c5)
	TenantID   string              `json:"tenant_id"`   // 租户ID
	AccessID   string              `json:"access_id"`   // AccessID
	AccessKey  *xcrypto.PrivateKey `json:"access_key"`  // AccessKey
	Account    string              `json:"account"`     // 链账户
	MyKmsKeyID string              `json:"mykmskey_id"` // 托管标识
}

// Client 发送请求使用的客户端
type Client interface {
	// CreateAccount 创建账户
	CreateAccount(ctx context.Context, account, kmsID string, gas int) (string, error)

	// Deposit 存证
	Deposit(ctx context.Context, content string, gas int) (string, error)

	// DeploySolidity 部署Solidity合约
	DeploySolidity(ctx context.Context, name, code string, gas int) (string, error)

	// AsyncCallSolidity 异步调用Solidity合约
	AsyncCallSolidity(ctx context.Context, contractName, methodSign, inputParams, outTypes string, gas int) (string, error)

	// QueryTransaction 查询交易
	QueryTransaction(ctx context.Context, xhash string) (string, error)

	// QueryReceipt 查询交易回执
	QueryReceipt(ctx context.Context, xhash string) (string, error)

	// QueryBlockHeader 查询块头
	QueryBlockHeader(ctx context.Context, blockNumber int64) (string, error)

	// QueryBlockBody 查询块体
	QueryBlockBody(ctx context.Context, blockNumber int64) (string, error)

	// QueryLastBlock 查询最新块高
	QueryLastBlock(ctx context.Context) (string, error)

	// QueryAccount 查询账户
	QueryAccount(ctx context.Context, account string) (string, error)
}

// ChainCallOption 链调用选项
type ChainCallOption func(params X)

func WithParam(key string, value any) ChainCallOption {
	return func(params X) {
		params[key] = value
	}
}

type client struct {
	endpoint string
	config   *Config
	httpCli  *resty.Client
	logger   func(ctx context.Context, err error, data map[string]string)
}

func (c *client) shakehand(ctx context.Context) (string, error) {
	timeStr := strconv.FormatInt(time.Now().UnixMilli(), 10)

	sign, err := c.config.AccessKey.Sign(crypto.SHA256, []byte(c.config.AccessID+timeStr))
	if err != nil {
		return "", err
	}

	params := X{
		"accessId": c.config.AccessID,
		"time":     timeStr,
		"secret":   hex.EncodeToString(sign),
	}
	return c.do(ctx, c.endpoint+SHAKE_HAND, params)
}

func (c *client) chainCall(ctx context.Context, method string, options ...ChainCallOption) (string, error) {
	token, err := c.shakehand(ctx)
	if err != nil {
		return "", err
	}

	params := X{}
	for _, f := range options {
		f(params)
	}
	params["bizid"] = c.config.BizID
	params["accessId"] = c.config.AccessID
	params["method"] = method
	params["token"] = token

	return c.do(ctx, c.endpoint+CHAIN_CALL, params)
}

func (c *client) chainCallForBiz(ctx context.Context, method string, options ...ChainCallOption) (string, error) {
	token, err := c.shakehand(ctx)
	if err != nil {
		return "", err
	}

	params := X{}
	for _, f := range options {
		f(params)
	}
	params["orderId"] = uuid.New().String()
	params["bizid"] = c.config.BizID
	params["account"] = c.config.Account
	params["mykmsKeyId"] = c.config.MyKmsKeyID
	params["method"] = method
	params["accessId"] = c.config.AccessID
	params["tenantid"] = c.config.TenantID
	params["token"] = token

	return c.do(ctx, c.endpoint+CHAIN_CALL_FOR_BIZ, params)
}

func (c *client) do(ctx context.Context, reqURL string, params X) (string, error) {
	log := internal.NewReqLog(http.MethodPost, reqURL)
	defer log.Do(ctx, c.logger)

	body, err := json.Marshal(params)
	if err != nil {
		log.SetError(err)
		return "", err
	}
	log.SetReqBody(string(body))

	resp, err := c.httpCli.R().
		SetContext(ctx).
		SetHeader(internal.HeaderContentType, internal.ContentJSON).
		SetBody(body).
		Post(reqURL)
	if err != nil {
		log.SetError(err)
		return "", err
	}
	log.SetRespHeader(resp.Header())
	log.SetStatusCode(resp.StatusCode())
	log.SetRespBody(string(resp.Body()))
	if !resp.IsSuccess() {
		return "", fmt.Errorf("HTTP Request Error, StatusCode = %d", resp.StatusCode())
	}

	ret := gjson.ParseBytes(resp.Body())
	if !ret.Get("success").Bool() {
		return "", fmt.Errorf("%s | %s", ret.Get("code").String(), ret.Get("data").String())
	}
	return ret.Get("data").String(), nil
}

// Option 自定义设置项
type Option func(c *client)

// WithHttpClient 设置自定义 HTTP Client
func WithHttpClient(cli *http.Client) Option {
	return func(c *client) {
		c.httpCli = resty.NewWithClient(cli)
	}
}

// WithLogger 设置日志记录
func WithLogger(fn func(ctx context.Context, err error, data map[string]string)) Option {
	return func(c *client) {
		c.logger = fn
	}
}

// NewClient 生成蚂蚁联盟链客户端
func NewClient(cfg *Config, options ...Option) Client {
	c := &client{
		endpoint: "https://rest.baas.alipay.com",
		config:   cfg,
		httpCli:  internal.NewClient(),
	}
	for _, f := range options {
		f(c)
	}
	return c
}
