package sandpay

import (
	"crypto"
	"encoding/base64"
	"errors"
	"net/url"
	"time"

	"github.com/shenghui0779/sdk-go/lib"
	lib_crypto "github.com/shenghui0779/sdk-go/lib/crypto"
	"github.com/shenghui0779/sdk-go/lib/value"
)

const OK = "000000"

// Form 数据表单
type Form struct {
	Head value.V `json:"Head"`
	Body value.V `json:"body"`
}

// URLEncode 数据表单格式化为POST表单
func (f *Form) URLEncode(mid string, key *lib_crypto.PrivateKey) (string, error) {
	if key == nil {
		return "", errors.New("private key is nil (forgotten configure?)")
	}

	f.Head["mid"] = mid

	b, err := lib.MarshalNoEscapeHTML(f)
	if err != nil {
		return "", err
	}

	sign, err := key.Sign(crypto.SHA1, b)
	if err != nil {
		return "", err
	}

	v := make(url.Values)

	v.Set("charset", "utf-8")
	v.Set("data", string(b))
	v.Set("signType", "01")
	v.Set("sign", base64.StdEncoding.EncodeToString(sign))

	return v.Encode(), nil
}

// HeadOption 报文头配置项
type HeadOption func(form *Form)

// WithVersion 设置版本号：默认：1.0；功能产品号为微信小程序或支付宝生活号，对账单需获取营销优惠金额字段传：3.0
func WithVersion(v string) HeadOption {
	return func(form *Form) {
		form.Head["version"] = v
	}
}

// WithPLMid 设置平台ID：接入类型为2时必填，在担保支付模式下填写核心商户号；在杉德宝平台终端模式下填写平台商户号
func WithPLMid(id string) HeadOption {
	return func(form *Form) {
		form.Head["plMid"] = id
	}
}

// WithAccessType 设置接入类型：1 - 普通商户接入（默认）；2 - 平台商户接入
func WithAccessType(v string) HeadOption {
	return func(form *Form) {
		form.Head["accessType"] = v
	}
}

// WithChannelType 设置渠道类型：07 - 互联网（默认）；08 - 移动端
func WithChannelType(v string) HeadOption {
	return func(form *Form) {
		form.Head["channelType"] = v
	}
}

// NewReqForm 生成请求数据表单
func NewReqForm(method, productID string, body value.V, options ...HeadOption) *Form {
	form := &Form{
		Head: value.V{
			"version":     "1.0",
			"method":      method,
			"productId":   productID,
			"accessType":  "1",
			"channelType": "07",
			"reqTime":     time.Now().In(lib.GMT8).Format("20060102150405"),
		},
		Body: body,
	}

	for _, fn := range options {
		fn(form)
	}

	return form
}
