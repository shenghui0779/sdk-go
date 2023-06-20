package mch

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/mock"
	"github.com/shenghui0779/gochat/wx"
	"github.com/stretchr/testify/assert"
)

func TestShortURL(t *testing.T) {
	body, err := wx.FormatMap2XMLForTest(wx.WXML{
		"appid":     "wx2421b1c4370ec43b",
		"mch_id":    "10000100",
		"long_url":  "weixin%3A%2F%2Fwxpay%2Fbizpayurl%3Fsign%3DXXXXX%26appid%3DXXXXX%26mch_id%3DXXXXX%26product_id%3DXXXXXX%26time_stamp%3DXXXXXX%26nonce_str%3DXXXXX",
		"nonce_str": "ec2316275641faa3aacf3cc599e8730f",
		"sign":      "423DE6471E12AB787CC23B87A6DBF449",
	})

	assert.Nil(t, err)

	resp := []byte(`<xml>
	<return_code>SUCCESS</return_code>
	<return_msg>OK</return_msg>
	<appid>wx2421b1c4370ec43b</appid>
	<mch_id>10000100</mch_id>
	<nonce_str>o5bAKF3o2ypC8hwa</nonce_str>
	<sign>48B30BC93E3190C8A969C173E4521427</sign>
	<result_code>SUCCESS</result_code>
	<short_url>weixin://wxpay/s/XXXXXX</short_url>
</xml>`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.mch.weixin.qq.com/tools/shorturl", body).Return(resp, nil)

	mch := New("10000100", "192006250b4c09247ec02edce69f6a2d", WithNonce(func() string {
		return "ec2316275641faa3aacf3cc599e8730f"
	}), WithMockClient(client))

	r, err := mch.Do(context.TODO(), ShortURL("wx2421b1c4370ec43b", "weixin://wxpay/bizpayurl?sign=XXXXX&appid=XXXXX&mch_id=XXXXX&product_id=XXXXXX&time_stamp=XXXXXX&nonce_str=XXXXX"))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
		"return_code": "SUCCESS",
		"return_msg":  "OK",
		"appid":       "wx2421b1c4370ec43b",
		"mch_id":      "10000100",
		"nonce_str":   "o5bAKF3o2ypC8hwa",
		"sign":        "48B30BC93E3190C8A969C173E4521427",
		"result_code": "SUCCESS",
		"short_url":   "weixin://wxpay/s/XXXXXX",
	}, r)
}

func TestAuthCodeToOpenID(t *testing.T) {
	body, err := wx.FormatMap2XMLForTest(wx.WXML{
		"appid":     "wx2421b1c4370ec43b",
		"mch_id":    "10000100",
		"auth_code": "120269300684844649",
		"nonce_str": "ec2316275641faa3aacf3cc599e8730f",
		"sign":      "A024BE089899DB35131FCB720FAE08A6",
	})

	assert.Nil(t, err)

	resp := []byte(`<xml>
	<return_code>SUCCESS</return_code>
	<return_msg>OK</return_msg>
	<appid>wx2421b1c4370ec43b</appid>
	<mch_id>10000100</mch_id>
	<nonce_str>o5bAKF3o2ypC8hwa</nonce_str>
	<sign>9F22C037E4219E3642443EF95D378FA7</sign>
	<result_code>SUCCESS</result_code>
	<openid>oUpF8uN95-Ptaags6E_roPHg7AG0</openid>
</xml>`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://api.mch.weixin.qq.com/tools/authcodetoopenid", body).Return(resp, nil)

	mch := New("10000100", "192006250b4c09247ec02edce69f6a2d", WithNonce(func() string {
		return "ec2316275641faa3aacf3cc599e8730f"
	}), WithMockClient(client))

	r, err := mch.Do(context.TODO(), AuthCodeToOpenID("wx2421b1c4370ec43b", "120269300684844649"))

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
		"return_code": "SUCCESS",
		"return_msg":  "OK",
		"appid":       "wx2421b1c4370ec43b",
		"mch_id":      "10000100",
		"nonce_str":   "o5bAKF3o2ypC8hwa",
		"sign":        "9F22C037E4219E3642443EF95D378FA7",
		"result_code": "SUCCESS",
		"openid":      "oUpF8uN95-Ptaags6E_roPHg7AG0",
	}, r)
}

func TestRSAPublicKey(t *testing.T) {
	body, err := wx.FormatMap2XMLForTest(wx.WXML{
		"mch_id":    "10000100",
		"nonce_str": "50780e0cca98c8c8e814883e5caa672e",
		"sign_type": "MD5",
		"sign":      "CA227C435D88EE017A9457B657FCA515",
	})

	assert.Nil(t, err)

	resp := []byte(`<xml>
	<return_code>SUCCESS</return_code>
	<mch_id>10000100</mch_id>
	<result_code>SUCCESS</result_code>
	<pub_key>-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAl1c+37GJSFSqbuHJ/wge
LzxLp7C2GYrjzVAnEF3xgjJVTltkQzdu3u+fcB3c/dgHX/Zdv5fqVoOqvoOMk4N4
zdGeaxN+Cm19c1gsxigNJDtm6Qno1s1T/qPph/zRArylM0N9Z3vWVEq4xI4B4NXk
6IoK/bXc1dwQe5UBzIZyzU5aWfqmTQilWEs7mqro43LTFkhN05QjC7IUFvWEhh6T
wvGYLBSAn+oNw/uSAu6B3c6dh+pslgORCzrIRs68GWsARGZkI/lmOJWEgzQ9KC7b
yHVqEnDDaWQFyQpq30JdP6YTXR/xlKyo8f1DingoSDXAhKMGRKaT4oIFkE6OA3jt
DQIDAQAB
-----END PUBLIC KEY-----</pub_key>
</xml>`)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodPost, "https://fraud.mch.weixin.qq.com/risk/getpublickey", body).Return(resp, nil)

	mch := New("10000100", "192006250b4c09247ec02edce69f6a2d", WithNonce(func() string {
		return "50780e0cca98c8c8e814883e5caa672e"
	}), WithMockClient(client))

	r, err := mch.Do(context.TODO(), RSAPublicKey())

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
		"return_code": "SUCCESS",
		"mch_id":      "10000100",
		"result_code": "SUCCESS",
		"pub_key": `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAl1c+37GJSFSqbuHJ/wge
LzxLp7C2GYrjzVAnEF3xgjJVTltkQzdu3u+fcB3c/dgHX/Zdv5fqVoOqvoOMk4N4
zdGeaxN+Cm19c1gsxigNJDtm6Qno1s1T/qPph/zRArylM0N9Z3vWVEq4xI4B4NXk
6IoK/bXc1dwQe5UBzIZyzU5aWfqmTQilWEs7mqro43LTFkhN05QjC7IUFvWEhh6T
wvGYLBSAn+oNw/uSAu6B3c6dh+pslgORCzrIRs68GWsARGZkI/lmOJWEgzQ9KC7b
yHVqEnDDaWQFyQpq30JdP6YTXR/xlKyo8f1DingoSDXAhKMGRKaT4oIFkE6OA3jt
DQIDAQAB
-----END PUBLIC KEY-----`,
	}, r)
}
