# gochat

[![golang](https://img.shields.io/badge/Language-Go-green.svg?style=flat)](https://golang.org) [![GitHub release](https://img.shields.io/github/release/shenghui0779/gochat.svg)](https://github.com/shenghui0779/gochat/releases/latest) [![pkg.go.dev](https://img.shields.io/badge/dev-reference-007d9c?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/shenghui0779/gochat) [![Apache 2.0 license](http://img.shields.io/badge/license-Apache%202.0-brightgreen.svg)](http://opensource.org/licenses/apache2.0)

微信 Go SDK（支持多账号配置）

- 支付.v2
- 公众号
- 小程序
- 企业微信

| 目录   | 对应                            | 功能                                                      |
| ------ | ------------------------------- | --------------------------------------------------------- |
| /mch   | 微信支付.v2（普通商户直连模式） | 下单、支付、退款、查询、委托代扣、企业付款、企业红包 等   |
| /offia | 微信公众号（Official Accounts） | 网页授权、用户管理、模板消息、菜单管理、客服、事件消息 等 |
| /minip | 微信小程序（Mini Program）      | 小程序授权、数据解密、二维码、消息发送、事件消息 等       |
| /corp  | 企业微信小程序（Work Wechat）   | 支持几乎全部服务端API                                     |

## 获取

```sh
go get -u github.com/shenghui0779/gochat
```

## 使用须知

- 微信API被封装成 `Action` 接口（授权 和 AccessToken 的部分API除外）
- 每个API对应一个 `Action`，统一由 `Do` 方法执行，返回结果以 `Result` 为前缀的结构体指针接收
- 对于微信通知的事件消息，提供了三个方法：
  - 验签 - `VerifyEventSign`
  - 解密 - `DecryptEventMessage`
  - 回复 - `Reply`
- 企业微信按照不同功能模块划分了相应的目录，根据URL可以找到对应的目录和文件
- 由于自身账号限制，所有API均采用Mock单元测试（Mock数据来源于微信官方文档，如遇问题，欢迎提[Issue](https://github.com/shenghui0779/gochat/issues)）

## 支付.v2

```go
import (
    "github.com/shenghui0779/gochat"
    "github.com/shenghui0779/gochat/mch"
)

// 创建实例
pay := gochat.NewMch("mchid", "apikey", tls.Certificate...)

// 统一下单
action := mch.UnifyOrder("appid", &mch.ParamsUnifyOrder{...})
result, err := mch.Do(ctx, action)

if err != nil {
    fmt.Println(err)

    return
}

fmt.Println(result)
```

## 公众号

```go
import (
    "github.com/shenghui0779/gochat"
    "github.com/shenghui0779/gochat/offia"
)

// 创建实例
oa := gochat.NewOffia("appid", "appsecret")

// 设置服务器配置
oa.SetServerConfig("token", "encodingAESKey")

// 设置 debug 模式（目前支持记录，支持自定义日志）
oa.SetClient(wx.WithDedug(), wx.WithLogger(wx.Logger))

// --------- 生成网页授权URL ---------------------

url := oa.OAuth2URL("scope", "redirectURL", "state")

fmt.Println(url)

// --------- 获取网页授权Token ---------------------

result, err := oa.Code2OAuthToken(ctx, "code")

if err != nil {
    fmt.Println(err)

    return
}

fmt.Println(result)

// --------- 获取AccessToken ---------------------

result, err := oa.AccessToken(ctx)

if err != nil {
    fmt.Println(err)

    return
}

fmt.Println(result)

// --------- 获取关注的用户列表 ---------------------

result := new(ResultUserList)
action := offia.GetUserList("nextOpenID", result)

if err := oa.Do(ctx, action); err != nil {
    fmt.Println(err)

    return
}

fmt.Println(result)
```

## 小程序

```go
import (
    "github.com/shenghui0779/gochat"
    "github.com/shenghui0779/gochat/minip"
)

// 创建实例
oa := gochat.NewMinip("appid", "appsecret")

// 设置服务器配置
oa.SetServerConfig("token", "encodingAESKey")

// 设置 debug 模式（目前支持记录，支持自定义日志）
oa.SetClient(wx.WithDedug(), wx.WithLogger(wx.Logger))

// --------- 获取小程序授权的SessionKey ---------------------

result, err := oa.Code2Session(ctx, "code")

if err != nil {
    fmt.Println(err)

    return
}

fmt.Println(result)

// --------- 获取AccessToken ---------------------

result, err := oa.AccessToken(ctx)

if err != nil {
    fmt.Println(err)

    return
}

fmt.Println(result)

// --------- 解密授权的用户信息 ---------------------

result := new(minip.UserInfo)

if err := DecryptAuthInfo("sessionKey", "iv", "encryptedData", result); err != nil {
    log.Println(err)

    return
}

fmt.Println(result)

// --------- 创建小程序二维码（数量有限） ---------------------

qrcode := new(minip.QRCode)
action := minip.CreateQRCode("pagepath", 120, qrcode)

if err := minip.Do(ctx, action); err != nil {
    log.Println(err)

    return
}

fmt.Println(base64.StdEncoding.EncodeToString(qrcode.Buffer))
```

## 企业微信

```go
import (
    "github.com/shenghui0779/gochat"
    "github.com/shenghui0779/gochat/corp/user"
)

// 创建实例
cp := gochat.NewCorp("corpid")

// 设置服务器配置
cp.SetServerConfig("token", "encodingAESKey")

// 设置 debug 模式（支持自定义日志）
cp.SetClient(wx.WithDedug(), wx.WithLogger(wx.Logger))

// --------- 生成网页授权URL ---------------------

url := cp.OAuth2URL("scope", "redirectURL", "state")

fmt.Println(url)

// --------- 生成扫码授权URL ---------------------

url := cp.QRCodeAuthURL("agentID", "redirectURL", "state")

fmt.Println(url)

// --------- 获取AccessToken ---------------------

result, err := cp.AccessToken(ctx, "secret")

if err != nil {
    fmt.Println(err)

    return
}

fmt.Println(result)

// --------- 获取获取部门成员详情 ---------------------

result := new(ResultUserList)
action := user.ListUser(1, 1, result)

if err := cp.Do(ctx, action); err != nil {
    fmt.Println(err)

    return
}

fmt.Println(result)
```

## 说明

- [API Reference](https://pkg.go.dev/github.com/shenghui0779/gochat)
- 注意：因 `access_token` 每日获取次数有限且含有效期，故服务端应妥善保存 `access_token` 并定时刷新
- 配合 [yiigo](https://github.com/shenghui0779/yiigo) 使用，可以更方便的操作 `MySQL`、`MongoDB` 与 `Redis` 等

**Enjoy 😊**
