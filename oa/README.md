# 微信公众号（Official Accounts

```go
"github.com/shenghui0779/gochat"
"github.com/shenghui0779/gochat/oa"
```

### 初始化小程序实例

```go
wxoa := gochat.NewOA(appid, appsecret)

// 如果需要消息回复，需要设置原始ID（开发者微信号）
wxoa.SetOriginID(originID)

// 如果启用了服务器配置，需要设置配置项
wxoa.SetServerConfig(token, encodingAESKey)
```