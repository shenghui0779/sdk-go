package alipay

import (
	"strings"

	"github.com/shenghui0779/sdk-go/lib/xcrypto"
	"github.com/tidwall/gjson"
)

const CodeOK = "10000" // API请求成功

const (
	HeaderMethodOverride = "x-http-method-override"
	HeaderRequestID      = "alipay-request-id"
	HeaderTraceID        = "alipay-trace-id"
	HeaderRootCertSN     = "alipay-root-cert-sn"
	HeaderNonce          = "alipay-nonce"
	HeaderTimestamp      = "alipay-timestamp"
	HeaderEncryptType    = "alipay-encrypt-type"
	HeaderAppAuthToken   = "alipay-app-auth-token"
	HeaderSignature      = "alipay-signature"
)

type GrantType string

const (
	OAuthCode    GrantType = "authorization_code"
	RefreshToken GrantType = "refresh_token"
)

// APIResult API结果 (支付v3)
type APIResult struct {
	Code int // HTTP状态码
	Body gjson.Result
}

// FormatPKCS1PrivateKey 格式化支付宝应用私钥(PKCS#1)
func FormatPKCS1PrivateKey(pemStr string) (xcrypto.RSAPadding, []byte) {
	rawLen := 64
	keyLen := len(pemStr)

	raws := keyLen / rawLen
	temp := keyLen % rawLen
	if temp > 0 {
		raws++
	}

	start := 0
	end := start + rawLen

	var builder strings.Builder

	builder.WriteString("-----BEGIN RSA PRIVATE KEY-----\n")
	for i := 0; i < raws; i++ {
		if i == raws-1 {
			builder.WriteString(pemStr[start:])
		} else {
			builder.WriteString(pemStr[start:end])
		}

		builder.WriteByte('\n')

		start += rawLen
		end = start + rawLen
	}
	builder.WriteString("-----END RSA PRIVATE KEY-----\n")

	return xcrypto.RSA_PKCS1, []byte(builder.String())
}

// FormatPKCS8PrivateKey 格式化支付宝应用私钥(PKCS#8)
func FormatPKCS8PrivateKey(pemStr string) (xcrypto.RSAPadding, []byte) {
	rawLen := 64
	keyLen := len(pemStr)

	raws := keyLen / rawLen
	temp := keyLen % rawLen
	if temp > 0 {
		raws++
	}

	start := 0
	end := start + rawLen

	var builder strings.Builder

	builder.WriteString("-----BEGIN PRIVATE KEY-----\n")
	for i := 0; i < raws; i++ {
		if i == raws-1 {
			builder.WriteString(pemStr[start:])
		} else {
			builder.WriteString(pemStr[start:end])
		}

		builder.WriteByte('\n')

		start += rawLen
		end = start + rawLen
	}
	builder.WriteString("-----END PRIVATE KEY-----\n")

	return xcrypto.RSA_PKCS8, []byte(builder.String())
}

// FormatPKCS1PublicKey 格式化支付宝应用公钥(PKCS#1)
func FormatPKCS1PublicKey(pemStr string) (xcrypto.RSAPadding, []byte) {
	rawLen := 64
	keyLen := len(pemStr)

	raws := keyLen / rawLen
	temp := keyLen % rawLen
	if temp > 0 {
		raws++
	}

	start := 0
	end := start + rawLen

	var builder strings.Builder

	builder.WriteString("-----BEGIN RSA PUBLIC KEY-----\n")
	for i := 0; i < raws; i++ {
		if i == raws-1 {
			builder.WriteString(pemStr[start:])
		} else {
			builder.WriteString(pemStr[start:end])
		}

		builder.WriteByte('\n')

		start += rawLen
		end = start + rawLen
	}
	builder.WriteString("-----END RSA PUBLIC KEY-----\n")

	return xcrypto.RSA_PKCS1, []byte(builder.String())
}

// FormatPKCS8PublicKey 格式化支付宝应用公钥(PKCS#8)
func FormatPKCS8PublicKey(pemStr string) (xcrypto.RSAPadding, []byte) {
	rawLen := 64
	keyLen := len(pemStr)

	raws := keyLen / rawLen
	temp := keyLen % rawLen
	if temp > 0 {
		raws++
	}

	start := 0
	end := start + rawLen

	var builder strings.Builder

	builder.WriteString("-----BEGIN PUBLIC KEY-----\n")
	for i := 0; i < raws; i++ {
		if i == raws-1 {
			builder.WriteString(pemStr[start:])
		} else {
			builder.WriteString(pemStr[start:end])
		}
		builder.WriteByte('\n')

		start += rawLen
		end = start + rawLen
	}
	builder.WriteString("-----END PUBLIC KEY-----\n")

	return xcrypto.RSA_PKCS8, []byte(builder.String())
}
