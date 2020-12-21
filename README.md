# gochat

[![golang](https://img.shields.io/badge/Language-Go-green.svg?style=flat)](https://golang.org)
[![GitHub release](https://img.shields.io/github/release/shenghui0779/gochat.svg)](https://github.com/shenghui0779/gochat/releases/latest)
[![pkg.go.dev](https://img.shields.io/badge/dev-reference-007d9c?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/shenghui0779/gochat)
[![MIT license](http://img.shields.io/badge/license-MIT-brightgreen.svg)](http://opensource.org/licenses/MIT)

微信 Go SDK（支付、公众号、小程序）

| 目录  | 对应                         | 功能                                           |
| ---- | ---------------------------- | --------------------------------------------- |
| /mch | 微信支付（普通商户直连模式）      | 下单、支付、退款、查询、委托代扣、企业付款、企业红包 等 |
| /oa  | 微信公众号（Official Accounts）| 网页授权、用户管理、模板消息、菜单管理、事件消息 等     |
| /mp  | 微信小程序（Mini Program）     | 小程序授权、数据解密、二维码、事件消息 等             |

## 获取

```sh
go get -u github.com/shenghui0779/gochat
```

## 文档

- [API Reference](https://pkg.go.dev/github.com/shenghui0779/gochat)
- [支付](https://github.com/shenghui0779/gochat/wiki/支付)
- [公众号](https://github.com/shenghui0779/gochat/wiki/公众号)
- [小程序](https://github.com/shenghui0779/gochat/wiki/小程序)

## 说明

- 支持 Go1.11+
- 注意：因 `access_token` 每日获取次数有限且含有效期，故服务端应妥善保存 `access_token` 并定时刷新
- 配合 [yiigo](https://github.com/shenghui0779/yiigo) 使用，可以更方便的操作 `MySQL`、`MongoDB` 与 `Redis` 等

**Enjoy 😊**

