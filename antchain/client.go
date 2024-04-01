package antchain

import (
	"context"
	"crypto"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/tidwall/gjson"

	"github.com/shenghui0779/sdk-go/lib"
	lib_crypto "github.com/shenghui0779/sdk-go/lib/crypto"
	"github.com/shenghui0779/sdk-go/lib/curl"
)

// Config 客户端配置
type Config struct {
	BizID      string                 `json:"biz_id"`      // 链ID (a00e36c5)
	TenantID   string                 `json:"tenant_id"`   // 租户ID
	AccessID   string                 `json:"access_id"`   // AccessID
	AccessKey  *lib_crypto.PrivateKey `json:"access_key"`  // AccessKey
	Account    string                 `json:"account"`     // 链账户
	MyKmsKeyID string                 `json:"mykmskey_id"` // 托管标识
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
	QueryTransaction(ctx context.Context, hash string) (string, error)

	// QueryReceipt 查询交易回执
	QueryReceipt(ctx context.Context, hash string) (string, error)

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
type ChainCallOption func(params lib.X)

func WithParam(key string, value any) ChainCallOption {
	return func(params lib.X) {
		params[key] = value
	}
}

type client struct {
	endpoint string
	config   *Config
	httpCli  curl.Client
	logger   func(ctx context.Context, data map[string]string)
}

func (c *client) shakehand(ctx context.Context) (string, error) {
	timeStr := strconv.FormatInt(time.Now().UnixMilli(), 10)

	sign, err := c.config.AccessKey.Sign(crypto.SHA256, []byte(c.config.AccessID+timeStr))
	if err != nil {
		return "", err
	}

	params := lib.X{
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

	params := lib.X{}

	for _, fn := range options {
		fn(params)
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

	params := lib.X{}

	for _, fn := range options {
		fn(params)
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

func (c *client) do(ctx context.Context, reqURL string, params lib.X) (string, error) {
	log := lib.NewReqLog(http.MethodPost, reqURL)
	defer log.Do(ctx, c.logger)

	body, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	log.SetReqBody(string(body))

	resp, err := c.httpCli.Do(ctx, http.MethodPost, reqURL, body, curl.WithHeader(curl.HeaderContentType, curl.ContentJSON))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	log.SetRespHeader(resp.Header)
	log.SetStatusCode(resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP Request Error, StatusCode = %d", resp.StatusCode)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	log.SetRespBody(string(b))

	ret := gjson.ParseBytes(b)
	if !ret.Get("success").Bool() {
		return "", fmt.Errorf("%s | %s", ret.Get("code").String(), ret.Get("data").String())
	}

	return ret.Get("data").String(), nil
}

// Option 自定义设置项
type Option func(c *client)

// WithHttpCli 设置自定义 HTTP Client
func WithHttpCli(httpCli *http.Client) Option {
	return func(c *client) {
		c.httpCli = curl.NewHTTPClient(httpCli)
	}
}

// WithLogger 设置日志记录
func WithLogger(fn func(ctx context.Context, data map[string]string)) Option {
	return func(c *client) {
		c.logger = fn
	}
}

// NewClient 生成蚂蚁联盟链客户端
func NewClient(cfg *Config, options ...Option) Client {
	c := &client{
		endpoint: "https://rest.baas.alipay.com",
		config:   cfg,
		httpCli:  curl.NewDefaultClient(),
	}

	for _, fn := range options {
		fn(c)
	}

	return c
}
