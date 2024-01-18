# alipay

支付宝 Go SDK

### 使用说明

- 发送HTTP请求，使用 `Client.Do(...)`
- 以下场景使用 `Action.Encode(...)`
  * alipay.trade.app.pay(app支付接口2.0)
  * alipay.fund.auth.order.app.freeze(线上资金授权冻结接口)
- 以下场景使用 `Client.PageExecute(...)`
  * alipay.trade.wap.pay(手机网站支付接口2.0)
  * alipay.trade.page.pay(统一收单下单并支付页面接口)
  * alipay.user.certify.open.certify(身份认证开始认证)
- 验证回调通知，使用 `Client.VerifyNotify(...)`
- JSON结果均以 `gjson.Result` 返回，理论上支持所有 JSON API
- 解析加密数据，如：授权的用户信息和手机号，使用 `Client.DecodeEncryptData(...)`
