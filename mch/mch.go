package mch

import (
	"context"
	"crypto/rand"
	"crypto/tls"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"strconv"

	"github.com/shenghui0779/gochat/internal"
	"golang.org/x/crypto/pkcs12"
)

// Mch 微信支付
type Mch struct {
	appid     string
	mchid     string
	apikey    string
	nonce     func(size int) string
	client    internal.Client
	tlsClient internal.Client
}

// New returns new wechat pay
func New(appid, mchid, apikey string) *Mch {
	c := internal.NewHTTPClient(&tls.Config{InsecureSkipVerify: true})

	return &Mch{
		appid:  appid,
		mchid:  mchid,
		apikey: apikey,
		nonce: func(size int) string {
			nonce := make([]byte, size/2)
			io.ReadFull(rand.Reader, nonce)

			return hex.EncodeToString(nonce)
		},
		client:    c,
		tlsClient: c,
	}
}

// LoadCertFromP12File load cert from p12(pfx) file
func (mch *Mch) LoadCertFromP12File(path string) error {
	p12, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	cert, err := mch.pkcs12ToPem(p12)

	if err != nil {
		return err
	}

	mch.tlsClient = internal.NewHTTPClient(&tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	})

	return nil
}

// LoadCertFromPemFile load cert from PEM file
func (mch *Mch) LoadCertFromPemFile(certFile, keyFile string) error {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)

	if err != nil {
		return err
	}

	mch.tlsClient = internal.NewHTTPClient(&tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	})

	return nil
}

// LoadCertFromPemBlock load cert from a pair of PEM encoded data
func (mch *Mch) LoadCertFromPemBlock(certPEMBlock, keyPEMBlock []byte) error {
	cert, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)

	if err != nil {
		return err
	}

	mch.tlsClient = internal.NewHTTPClient(&tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	})

	return nil
}

// Do exec action
func (mch *Mch) Do(ctx context.Context, action internal.Action, options ...internal.HTTPOption) (internal.WXML, error) {
	body, err := action.WXML()(mch.appid, mch.mchid, mch.apikey, mch.nonce(16))

	if err != nil {
		return nil, err
	}

	reqURL := action.URL()()

	if len(reqURL) == 0 {
		return body, nil
	}

	var resp internal.WXML

	if action.TLS() {
		resp, err = mch.tlsClient.PostXML(ctx, reqURL, body, options...)
	} else {
		resp, err = mch.client.PostXML(ctx, reqURL, body, options...)
	}

	if err != nil {
		return nil, err
	}

	if resp["return_code"] != ResultSuccess {
		return nil, errors.New(resp["return_msg"])
	}

	if err := mch.VerifyWXReply(resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// APPAPI 用于APP拉起支付
func (mch *Mch) APPAPI(prepayID string, timestamp int64) internal.WXML {
	m := internal.WXML{
		"appid":     mch.appid,
		"partnerid": mch.mchid,
		"prepayid":  prepayID,
		"package":   "Sign=WXPay",
		"noncestr":  mch.nonce(16),
		"timestamp": strconv.FormatInt(timestamp, 10),
	}

	m["sign"] = internal.SignWithMD5(m, mch.apikey, true)

	return m
}

// JSAPI 用于JS拉起支付
func (mch *Mch) JSAPI(prepayID string, timestamp int64) internal.WXML {
	m := internal.WXML{
		"appId":     mch.appid,
		"nonceStr":  mch.nonce(16),
		"package":   fmt.Sprintf("prepay_id=%s", prepayID),
		"signType":  SignMD5,
		"timeStamp": strconv.FormatInt(timestamp, 10),
	}

	m["paySign"] = internal.SignWithMD5(m, mch.apikey, true)

	return m
}

// MPRedpackJSAPI 小程序领取红包
func (mch *Mch) MPRedpackJSAPI(pkg string, timestamp int64) internal.WXML {
	m := internal.WXML{
		"appId":     mch.appid,
		"nonceStr":  mch.nonce(16),
		"package":   url.QueryEscape(pkg),
		"timeStamp": strconv.FormatInt(timestamp, 10),
	}

	m["paySign"] = internal.SignWithMD5(m, mch.apikey, false)

	delete(m, "appId")
	m["signType"] = SignMD5

	return m
}

// VerifyWXReply 验证微信结果
func (mch *Mch) VerifyWXReply(m internal.WXML) error {
	if wxsign, ok := m["sign"]; ok {
		signature := ""

		if v, ok := m["sign_type"]; ok && v == SignHMacSHA256 {
			signature = internal.SignWithHMacSHA256(m, mch.apikey, true)
		} else {
			signature = internal.SignWithMD5(m, mch.apikey, true)
		}

		if wxsign != signature {
			return fmt.Errorf("signature verified failed, want: %s, got: %s", signature, wxsign)
		}
	}

	if appid, ok := m["appid"]; ok {
		if appid != mch.appid {
			return fmt.Errorf("appid mismatch, want: %s, got: %s", mch.appid, m["appid"])
		}
	}

	if mchid, ok := m["mch_id"]; ok {
		if mchid != mch.mchid {
			return fmt.Errorf("mchid mismatch, want: %s, got: %s", mch.mchid, m["mch_id"])
		}
	}

	return nil
}

func (mch *Mch) pkcs12ToPem(p12 []byte) (tls.Certificate, error) {
	blocks, err := pkcs12.ToPEM(p12, mch.mchid)

	if err != nil {
		return tls.Certificate{}, err
	}

	pemData := make([]byte, 0)

	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}

	// then use PEM data for tls to construct tls certificate:
	return tls.X509KeyPair(pemData, pemData)
}
