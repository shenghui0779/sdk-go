# gochat

[![golang](https://img.shields.io/badge/Language-Go-green.svg?style=flat)](https://golang.org)
[![GitHub release](https://img.shields.io/github/release/iiinsomnia/gochat.svg)](https://github.com/iiinsomnia/gochat/releases/latest)
[![pkg.go.dev](https://img.shields.io/badge/dev-reference-007d9c?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/iiinsomnia/gochat)
[![MIT license](http://img.shields.io/badge/license-MIT-brightgreen.svg)](http://opensource.org/licenses/MIT)

这可能是目前最好的 微信 SDK for Go

| 目录 | 对应         | 功能                                               |
| ---- | ------------ | -------------------------------------------------- |
| /mch | 微信商户平台 | 下单、支付、退款、查询、委托代扣、企业付款 等 |
| /pub | 微信公众平台 | 网页授权、菜单、模板消息、消息回复、用户管理、消息转客服 等 |
| /mp  | 微信小程序   | 小程序授权、用户数据解析、消息发送、二维码生成 等 |

## 获取

```sh
go get github.com/iiinsomnia/gochat
```

## 文档

- [API Reference](https://pkg.go.dev/github.com/iiinsomnia/gochat)

## 说明

- 支持 Go1.11+
- 注意：因 `access_token` 每日获取次数有限且含有效期，故服务端应妥善保存 `access_token` 并定时刷新
- 配合 [yiigo](https://github.com/iiinsomnia/yiigo) 使用，可以更方便的操作 `MySQL`、`MongoDB` 与 `Redis` 等

**Enjoy 😊**

