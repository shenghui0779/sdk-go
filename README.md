# gochat

[![golang](https://img.shields.io/badge/Language-Go-green.svg?style=flat)](https://golang.org)
[![GoDoc](https://godoc.org/github.com/iiinsomnia/gochat?status.svg)](https://godoc.org/github.com/iiinsomnia/gochat)
[![MIT license](http://img.shields.io/badge/license-MIT-brightgreen.svg)](http://opensource.org/licenses/MIT)

微信平台 Golang SDK

| 目录 | 对应                         | 功能                                               |
| ---- | ---------------------------- | -------------------------------------------------- |
| /mch | 微信商户平台                 | 统一下单、订单查询、订单关闭、申请退款、退款查询 等   |
| /pub | 微信公众平台(订阅号，服务号) | 网页授权、菜单、模板消息、消息回复、用户管理、消息转客服 等 |
| /mp  | 微信小程序                   | 小程序授权、用户数据解析、二维码生成 等               |

## 获取

```sh
go get github.com/iiinsomnia/gochat
```

## 文档

- [API Reference](https://godoc.org/github.com/iiinsomnia/gochat)

## 说明

- 支持 Go1.11+
- 注意：因 `access_token` 每日获取次数有限且含有效期，故服务端应妥善保存 `access_token` 并定时刷新
- 配合 [yiigo](https://github.com/iiinsomnia/yiigo) 使用，可以更方便的操作 `MySQL`、`MongoDB` 与 `Redis` 等

**Enjoy 😊**

