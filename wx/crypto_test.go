package wx

import (
	"crypto/aes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCBCCrypto(t *testing.T) {
	key := []byte("AES256Key-32Characters1234567890")
	iv := key[:aes.BlockSize]
	plainText := "IloveGochat"

	// ZERO_PADDING
	zero := NewCBCCrypto(key, iv, ZERO)

	e0b, err := zero.Encrypt([]byte(plainText))
	assert.Nil(t, err)

	d0b, err := zero.Decrypt(e0b)
	assert.Nil(t, err)
	assert.Equal(t, plainText, string(d0b))

	// PKCS5_PADDING
	pkcs5 := NewCBCCrypto(key, iv, PKCS5)

	e5b, err := pkcs5.Encrypt([]byte(plainText))
	assert.Nil(t, err)

	d5b, err := pkcs5.Decrypt(e5b)
	assert.Nil(t, err)
	assert.Equal(t, plainText, string(d5b))

	// PKCS7_PADDING
	pkcs7 := NewCBCCrypto(key, iv, PKCS7)

	e7b, err := pkcs7.Encrypt([]byte(plainText))
	assert.Nil(t, err)

	d7b, err := pkcs7.Decrypt(e7b)
	assert.Nil(t, err)
	assert.Equal(t, plainText, string(d7b))
}

func TestECBCrypto(t *testing.T) {
	key := []byte("AES256Key-32Characters1234567890")
	plainText := "IloveGochat"

	// ZERO_PADDING
	zero := NewECBCrypto(key, ZERO)

	e0b, err := zero.Encrypt([]byte(plainText))
	assert.Nil(t, err)

	d0b, err := zero.Decrypt(e0b)
	assert.Nil(t, err)
	assert.Equal(t, plainText, string(d0b))

	// PKCS5_PADDING
	pkcs5 := NewECBCrypto(key, PKCS5)

	e5b, err := pkcs5.Encrypt([]byte(plainText))
	assert.Nil(t, err)

	d5b, err := pkcs5.Decrypt(e5b)
	assert.Nil(t, err)
	assert.Equal(t, plainText, string(d5b))

	// PKCS7_PADDING
	pkcs7 := NewECBCrypto(key, PKCS7)

	e7b, err := pkcs7.Encrypt([]byte(plainText))
	assert.Nil(t, err)

	d7b, err := pkcs7.Decrypt(e7b)
	assert.Nil(t, err)
	assert.Equal(t, plainText, string(d7b))
}

func TestRSACrypto(t *testing.T) {
	eb, err := RSAEncrypt([]byte("ILoveWechatPay"), publicKey)

	assert.Nil(t, err)

	db, err := RSADecrypt(eb, privateKey)

	assert.Nil(t, err)
	assert.Equal(t, "ILoveWechatPay", string(db))
}

var (
	privateKey []byte
	publicKey  []byte
)

func TestMain(m *testing.M) {
	privateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAl1c+37GJSFSqbuHJ/wgeLzxLp7C2GYrjzVAnEF3xgjJVTltk
Qzdu3u+fcB3c/dgHX/Zdv5fqVoOqvoOMk4N4zdGeaxN+Cm19c1gsxigNJDtm6Qno
1s1T/qPph/zRArylM0N9Z3vWVEq4xI4B4NXk6IoK/bXc1dwQe5UBzIZyzU5aWfqm
TQilWEs7mqro43LTFkhN05QjC7IUFvWEhh6TwvGYLBSAn+oNw/uSAu6B3c6dh+ps
lgORCzrIRs68GWsARGZkI/lmOJWEgzQ9KC7byHVqEnDDaWQFyQpq30JdP6YTXR/x
lKyo8f1DingoSDXAhKMGRKaT4oIFkE6OA3jtDQIDAQABAoIBAEYkE2nNeJjjVJJL
Zzkh5At0YIP5rIwuCOJRMamuQI5dEZhdVxyoBmSgnj5yOMgVZWWeY1i26c7q+ymN
AowjtMt/SXLK9/GRSUE6LpYwXdbbCRkclKNpSnKMOWHjVGN2VwJpVyetB2rNrtC6
GDjCKXN09x8bOJyNf74nE0xdS7vGzDlmAhCwju34DuMhdj8GBtLZo8O0esaeqNuK
EhlQrur9KuyYJR63ZR306qJpVE7ZX6bFQZpwTrebnATHDnWcvVbVWWpfe8xmQwNa
b2Gsctv8Ght/Ka/OjbRP0d48ZnTGeOuC9eKjpUKi2nZiEiYsCUjTxO30Ib6Pw2Z3
lWMx7kECgYEAxM2UtYjTXFcIbRWSGx9b997xhPpnxLSPzO4JIM2WdQqlRBdgOi7u
BNIL19Z37d6CElEYJ+G/6lqs072xMWt4Nph2cgiKUzcOAAKfS0vna/IXir4oGhTb
auAsj7Ga7dQi23a3UTDb1bNavemo3SqYI1anud00TnyQdBvVJ1ZwADUCgYEAxNzv
zDLiABRETLtFU7zOEjYsB/+WV2cvofsqvq8NQDdyOP6UVZ8vE/DkG61uyMpWp0u/
3/A9krLTz9Gfgw4A7CFFDV3S+z1AY1T2N7I04+QQHMqfbcjotVEG7xouuEfjDN2P
Xi5M2zcmTAkuStO7Yx5UdGPdJNv6JgJyy2doBHkCgYAu6i8kI2z3W0wH7Rd6Xbxn
137Ny3/HNZ/+I1SLvFa8qgABvmzTEfLttUDbgCXwz5VEVo6imz9L17fRdivycwMi
SLAbuQt4kOxGdlmQ8pRFeF3CVlhq90PjM3OMAbPENEjm9mL2+OW/CNV95mC58Hh6
HCM5vJDGkQ1CkIv8p69lbQKBgAYRWULN/rFJ7qD+1LA0DZX6HXlRo2ymPY2clEC0
XJAyJU8kaaYJ9gWDU0SXH+cIdYtKhmt8mClBYc3yBByh/d1JWTuEPNCJnsZxA/XL
hF3R1b1NcYSMwL918+TCxdXgQVtQKO8aNjw7gu6tCcQ8qnXvpWLBATv1m8w4Hxmt
4kLhAoGAejdp4xTh6OYb4kfZA5EN/9wBO3l/7TwWrOe8qT1/FtWMfmcU62Y3LdXE
xuHKcd+Q3/PUQKM5lPFpXqyY/pCE9AQpjFmjo5eU99NNy/oS0P8IaCS2SyppGhF2
HsIxLjl3+jtjS8cptPO47qFnr7Pnvb7kA8MNVrI+ymny/WG/yfU=
-----END RSA PRIVATE KEY-----`)

	publicKey = []byte(`-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAl1c+37GJSFSqbuHJ/wge
LzxLp7C2GYrjzVAnEF3xgjJVTltkQzdu3u+fcB3c/dgHX/Zdv5fqVoOqvoOMk4N4
zdGeaxN+Cm19c1gsxigNJDtm6Qno1s1T/qPph/zRArylM0N9Z3vWVEq4xI4B4NXk
6IoK/bXc1dwQe5UBzIZyzU5aWfqmTQilWEs7mqro43LTFkhN05QjC7IUFvWEhh6T
wvGYLBSAn+oNw/uSAu6B3c6dh+pslgORCzrIRs68GWsARGZkI/lmOJWEgzQ9KC7b
yHVqEnDDaWQFyQpq30JdP6YTXR/xlKyo8f1DingoSDXAhKMGRKaT4oIFkE6OA3jt
DQIDAQAB
-----END PUBLIC KEY-----`)

	m.Run()
}
